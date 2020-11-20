package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Id      int64  `bson:"id"`
	ClassId int64  `bson:"score,omitempty"`
	Name    string `bson:"name,omitempty"`
}

var (
	collection *mongo.Collection
)

func Init() {
	ctx := context.Background()
	uri := "mongo://username:password@localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
		return
	}

	collection = client.Database("database_name").Collection("collection_name")
}

func Insert(ctx context.Context, student Student) (err error) {
	_, err = collection.InsertOne(ctx, &student)
	return
}

func Update(ctx context.Context, student Student) (modified int64, err error) {

	filter := bson.M{
		"id": student.Id,
	}

	var result *mongo.UpdateResult
	result, err = collection.UpdateOne(ctx, filter, bson.M{"$set": &student})
	if err != nil {
		panic(err)
		return
	}

	if result != nil {
		modified = result.ModifiedCount
	}

	return
}

func Delete(ctx context.Context, student Student) (deleted int64, err error) {
	filter := bson.M{
		"id": student.Id,
	}

	var result *mongo.DeleteResult
	result, err = collection.DeleteMany(ctx, filter)
	if err != nil {
		panic(err)
		return
	}

	if result != nil {
		deleted = result.DeletedCount
	}

	return
}

// ReplaceOne
// 未命中不报错, modified = 0
func ReplaceOne(ctx context.Context, student Student) (modified int64, err error) {

	// 一般情况可以使用非 bson.M{}
	filter := map[string]interface{}{
		"id": student.Id,
	}

	var result *mongo.UpdateResult
	result, err = collection.ReplaceOne(ctx, filter, student)
	if err != nil {
		panic(err)
		return
	}

	if result != nil {
		modified = result.ModifiedCount
	}

	return
}

// FindByFilter 根据条件查找记录
// 注 : filter 需要大于 0
// 未命中不报错, len(userComments) = 0
func FindByFilter(ctx context.Context, filter map[string]interface{}) (students []Student, err error) {

	// 参数集无效, 不允许全量查
	if len(filter) == 0 {
		return
	}

	var cursor *mongo.Cursor
	cursor, err = collection.Find(ctx, filter)
	if err != nil {
		panic(err)
		return
	}

	// 遍历结果
	for cursor.Next(ctx) {
		var student = Student{}
		if err = cursor.Decode(&student); err != nil {
			panic(err)
			return
		}
		students = append(students, student)
	}

	return
}
