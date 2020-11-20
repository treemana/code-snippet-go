package mongo

import (
	"context"
)

func Main() {
	Init()

	var (
		err      error
		modified int64
	)
	ctx := context.Background()
	student := Student{
		Id:      1,
		ClassId: 2,
		Name:    "JoJo",
	}

	err = Insert(ctx, student)
	if err != nil {
		panic(err)
	}

	student.Name = "MoMo"
	modified, err = Update(ctx, student)
	if err != nil {
		panic(err)
	}
	if modified < 1 {
		panic("not match")
	}

	student.ClassId = 12
	modified, err = ReplaceOne(ctx, student)
	if err != nil {
		panic(err)
	}
	if modified < 1 {
		panic("not match")
	}

	modified, err = Delete(ctx, student)
	if err != nil {
		panic(err)
	}
	if modified < 1 {
		panic("not match")
	}
}
