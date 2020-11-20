package mail

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

const (
	addr     = "x.y.com:25"       // smtp 服务器地址
	from     = "no-reply@abc.com" // 发件人
	boundary = "xxx xxx xxx"      // 分割线
)

// Attachment 附件数据结构, 在下列模式中二选一, 同时存在时
// 1 文件路径模式 : FilePath 存在, 默认附件显示名字为路径中的文件名, 可指定 FileName 为附件显示名字
// 2 文件名+文件内容模式 : FileName 和 Data 同时存在
type Attachment struct {
	FilePath string  // 文件绝对路径
	FileName string  // 附件名
	Data     *[]byte // 文件字节数组

	fileName string // 附件名
	data     []byte // 文件字节数组
}

type Email struct {
	Subject     string        // 邮件主题, 必须
	Content     string        // 邮件内容, 必须
	Attachments []*Attachment // 附件
	To          []string      // 接收地址, 必须
	Cc          []string      // 抄送地址
}

var contentDefault string

func Init() {
	hostName := os.Getenv("MATRIX_INSTANCE_ID")
	if hostName == "" {
		hostName, _ = os.Hostname()
	}

	buffer := bytes.NewBuffer(nil)
	buffer.WriteString(fmt.Sprintf("Host : %s\n", hostName))
	contentDefault = string(buffer.Bytes())
}

// Send 发送邮件
func (e *Email) Send() error {

	if err := e.check(); err != nil {
		return err
	}

	// Head
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString(fmt.Sprintf("From: %s\r\n", from))
	buffer.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(e.To, ",")))
	if len(e.Cc) > 0 {
		buffer.WriteString(fmt.Sprintf("Cc: %s\r\n", strings.Join(e.Cc, ",")))
		e.To = append(e.To, e.Cc...)
	}
	buffer.WriteString(fmt.Sprintf("Subject: %s\r\n", e.Subject))
	buffer.WriteString("MIME-Version: 1.0\r\n")
	buffer.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n", boundary))

	// Content
	buffer.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	buffer.WriteString("Content-Type: text/plain; charset=utf-8\r\n\r\n")
	buffer.WriteString(contentDefault)
	buffer.WriteString(e.Content)
	buffer.WriteString("\r\n")

	// Attachment
	for _, attachment := range e.Attachments {
		buffer.WriteString(fmt.Sprintf("--%s\r\n", boundary))
		buffer.WriteString("Content-Type: application/octet-stream\r\n")
		buffer.WriteString("Content-Transfer-Encoding: base64\r\n")
		buffer.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n\r\n", attachment.fileName))
		b := make([]byte, base64.StdEncoding.EncodedLen(len(attachment.data)))
		base64.StdEncoding.Encode(b, attachment.data)
		buffer.Write(b)
		buffer.WriteString("\r\n")
	}

	// Tail
	buffer.WriteString(fmt.Sprintf("--%s--\r\n\r\n", boundary))

	return smtp.SendMail(addr, nil, from, e.To, buffer.Bytes())
}

// check 参数检查
func (e *Email) check() error {

	if len(e.Subject) == 0 {
		return errors.New("invalid parameter subject")
	}

	if len(e.Content) == 0 {
		return errors.New("invalid parameter content")
	}

	if len(e.To) == 0 {
		return errors.New("invalid parameter to")
	}

	if e.Attachments == nil {
		e.Attachments = make([]*Attachment, 0)
	}

	for _, attachment := range e.Attachments {
		if len(attachment.FilePath) == 0 && (len(attachment.FileName) == 0 || attachment.Data == nil || len(*attachment.Data) == 0) {
			return errors.New("invalid parameter attachments")
		}

		// 生成附件名
		if len(attachment.FileName) > 0 {
			attachment.fileName = attachment.FileName
		} else {
			attachment.fileName = filepath.Base(attachment.FilePath)
		}

		// 生成文件内容
		if len(attachment.FilePath) > 0 {
			data, err := ioutil.ReadFile(attachment.FilePath)
			if err != nil {
				return err
			}
			attachment.data = data
		} else {
			attachment.data = *attachment.Data
		}
	}

	return nil
}
