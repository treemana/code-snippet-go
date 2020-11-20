package aws

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type s3Config struct {
	host           string
	accessKey      string
	secretKey      string
	forcePathStyle bool
}

type DownloadInfo struct {
	DeleteInfo
	LocalPrefix string
}

type DeleteInfo struct {
	Bucket string
	Key    string
}

type Empty struct {
}

var (
	s3Client        *s3.S3
	s3Session       *session.Session
	ErrDataNotFound = errors.New("data not found")
	// "face_prod_daily_data_20200401"
	host      = "http://s3-prod.xxx:7480"
	accessKey = "accessKey"
	secretKey = "secretKey"
)

// initialize ceph client for specific city
func Init() {
	log.Println("ceph.Init start")

	s3Session = session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:           aws.String("default"),
			Endpoint:         &host,
			S3ForcePathStyle: aws.Bool(true),
			Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		},
	}))

	s3Client = s3.New(s3Session)
}

func AllBuckets() {
	outPut, err := s3Client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(len(outPut.Buckets))
}

func BucketList(bucketName string) {

	var keyMap = make(map[string]Empty)
	var keyList = make([]string, 0)
	var listInput = s3.ListObjectsInput{
		Bucket: &bucketName,
	}
	var err error
	var outPut *s3.ListObjectsOutput
	for {
		outPut, err = s3Client.ListObjects(&listInput)
		if err != nil {
			log.Println(err)
		}
		for _, content := range outPut.Contents {
			keyMap[*content.Key] = Empty{}
			keyList = append(keyList, *content.Key)
		}
		if outPut.NextMarker == nil || !*outPut.IsTruncated || outPut.MaxKeys == nil || *outPut.MaxKeys == 0 {
			break
		}

		listInput.Marker = outPut.NextMarker
	}

	log.Println("keyMap", len(keyMap))
	log.Println("keyList", len(keyList))
}

func DownloadBucket(processor int, bucket, imageDir string) error {
	log.Println("ceph.DownloadBucket start")
	// 当前目录以 Bucket 命名
	if _, err := os.Stat(imageDir); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err = os.MkdirAll(imageDir, 0755); err != nil {
			return err
		}
	}

	downloadChan := make(chan DownloadInfo)
	var wg sync.WaitGroup
	wg.Add(processor)
	for i := 0; i < processor; i++ {
		go func(idx int) {
			log.Println("start multipleDownload", idx)
			multipleDownload(downloadChan)
			wg.Done()
			log.Println("finish multipleDownload", idx)
		}(i)
	}

	var (
		num    = 0
		err    error
		output *s3.ListObjectsOutput
		input  = s3.ListObjectsInput{
			Bucket: &bucket,
		}
	)

	// 遍历 Bucket 时下载图片
	for {
		output, err = s3Client.ListObjects(&input)
		if err != nil {
			log.Println(err)
			return err
		}

		for _, content := range output.Contents {
			num++
			if num%100000 == 0 {
				log.Println("download object", num)
			}

			downloadChan <- DownloadInfo{
				DeleteInfo: DeleteInfo{
					Bucket: bucket,
					Key:    *content.Key,
				},
				LocalPrefix: filepath.Join(imageDir, *content.Key),
			}
		}

		// 没有偏移量了; 此次返回没有截取; MaxKeys 无效
		if output.NextMarker == nil || !*output.IsTruncated || output.MaxKeys == nil || *output.MaxKeys == 0 {
			break
		}

		// ceph 每次最多返回 1000 个内容
		// 用 marker 标记偏移量
		input.Marker = output.NextMarker
	}

	close(downloadChan)
	wg.Wait()

	log.Println("ceph.DownloadBucket finish", num)
	return nil
}

func multipleDownload(dc chan DownloadInfo) {
	for dInfo := range dc {
		if err := downloadImage(dInfo.Bucket, dInfo.Key, dInfo.LocalPrefix, false); err != nil {
			// 出错了就停止所有的协程, 让主函数停止, 抛出错误
			close(dc)
			panic(err)
		}
	}
}

func multipleDelete(dc chan DeleteInfo) {
	for dInfo := range dc {
		if err := DeleteObject(dInfo.Bucket, dInfo.Key); err != nil {
			// 出错了就停止所有的协程, 让主函数停止, 抛出错误
			close(dc)
			panic(err)
		}
	}
}

// 下载 ceph 图片数据到本地
// city       : ceph 数据所在的集群
// date       : 数据日期, 按照日期区分 Bucket
// remoteName : ceph 数据的 full path name, 不包括 Bucket
// localName  : 下载到本地的路径及名称
// force      : 是否强制覆盖本地数据(若本地同名文件存在)
func downloadImage(bucketName, key, prefixName string, force bool) error {

	// remoteName 的文件数据以二进制留的形式存在 resp.Body
	resp, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &key,
	})
	if err != nil {
		// 类型转换会丢数据, 等到升级 v2 版本再试试
		if cephErr := err.(awserr.Error); cephErr.Code() == "NotFound" {
			return ErrDataNotFound
		} else {
			return err
		}
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	localName := prefixName + getSuffix(*resp.ContentType)

	if _, err := os.Stat(localName); err == nil {
		if !force {
			// 不强制覆盖
			return nil
		}
		// 强制覆盖就删掉已有文件
		if err = os.Remove(localName); err != nil {
			return err
		}
	} else if !os.IsNotExist(err) {
		// 嗅探本地文件出错
		return err
	}

	// 运行到这里, 本地文件应该不存在
	// 准备接收文件
	idx := strings.LastIndex(localName, "/")
	if idx > 0 {
		dirPath := localName[0:idx]
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}

	var localFile *os.File
	if localFile, err = os.OpenFile(localName, os.O_CREATE|os.O_RDWR, 0644); err != nil {
		return err
	}
	defer func() {
		_ = localFile.Close()
	}()

	// 将 remote 的数据流 copy 到 local
	if _, err = io.Copy(localFile, resp.Body); err != nil {
		return err
	}

	return nil
}

func newS3Client(config s3Config) *s3.S3 {
	return s3.New(session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:           aws.String("default"),
			Endpoint:         &config.host,
			S3ForcePathStyle: &config.forcePathStyle,
			Credentials:      credentials.NewStaticCredentials(config.accessKey, config.secretKey, ""),
		},
	})))
}

func getSuffix(cType string) string {
	switch cType {
	case "image/jpeg":
		return ".jpg"
	}
	panic(cType)
}

func DownloadObject(sess *session.Session, filename, key, bucket *string) error {
	// snippet-start:[s3.go.download_object.create]
	file, err := os.Create(*filename)
	// snippet-end:[s3.go.download_object.create]
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	// snippet-start:[s3.go.download_object.call]
	downloader := s3manager.NewDownloader(sess)

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: bucket,
			Key:    key,
		})
	// snippet-end:[s3.go.download_object.call]
	if err != nil {
		return err
	}

	return nil
}

func DeleteBucket(processor int, bucket string) error {
	log.Println("ceph.DeleteBucket start", bucket)
	var (
		num    = 0
		err    error
		output *s3.ListObjectsOutput
		input  = s3.ListObjectsInput{
			Bucket: &bucket,
		}
	)

	deleteChan := make(chan DeleteInfo)
	var wg sync.WaitGroup
	wg.Add(processor)
	for i := 0; i < processor; i++ {
		go func(idx int) {
			log.Println("start multipleDelete", idx)
			multipleDelete(deleteChan)
			wg.Done()
			log.Println("finish multipleDelete", idx)
		}(i)
	}

	// 遍历 Bucket 删除
	for {
		output, err = s3Client.ListObjects(&input)
		if err != nil {
			log.Println(err)
			return err
		}

		if len(output.Contents) == 0 {
			break
		}

		for _, content := range output.Contents {
			num++
			if num%100000 == 0 {
				log.Println("delete object", num)
			}
			deleteChan <- DeleteInfo{
				Bucket: bucket,
				Key:    *content.Key,
			}
		}
	}

	wg.Wait()

	// Delete the S3 Bucket
	// It must be empty or else the call fails
	if _, err = s3Client.DeleteBucket(&s3.DeleteBucketInput{Bucket: aws.String(bucket)}); err != nil {
		return err
	}

	// Wait until Bucket is deleted before finishing
	if err = s3Client.WaitUntilBucketNotExists(&s3.HeadBucketInput{Bucket: aws.String(bucket)}); err != nil {
		return err
	}

	log.Println("ceph.DeleteBucket finish", bucket)
	return nil
}

func DeleteObject(bucket, key string) error {

	// Delete the item
	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(key)})
	if err != nil {
		return err
	}

	err = s3Client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	return nil
}

func Count(bucket string) (int, error) {
	log.Println("ceph.Count start")

	var (
		num    = 0
		err    error
		output *s3.ListObjectsOutput
		input  = s3.ListObjectsInput{
			Bucket: &bucket,
		}
	)

	// 遍历 Bucket 计数
	for {
		output, err = s3Client.ListObjects(&input)
		if err != nil {
			log.Println(err)
			return 0, err
		}

		for range output.Contents {
			num++
			if num%100000 == 0 {
				log.Println("count", num)
			}
		}

		// 没有偏移量了; 此次返回没有截取; MaxKeys 无效
		if output.NextMarker == nil || !*output.IsTruncated || output.MaxKeys == nil || *output.MaxKeys == 0 {
			break
		}

		// ceph 每次最多返回 1000 个内容
		// 用 marker 标记偏移量
		input.Marker = output.NextMarker
	}

	log.Println("ceph.Count finish", num)
	return num, nil
}

func ExistBucket(bucket string) (bool, error) {
	_, err := s3Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		if cephErr := err.(awserr.Error); cephErr.Code() == "NotFound" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
