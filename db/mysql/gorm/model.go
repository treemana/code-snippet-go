package gorm

type ThisIsModel struct {
	Id   uint64 `gorm:"column:id;primary_key"`
	Data string `gorm:"column:data"`
}

func (tim *ThisIsModel) TableName() string {
	return "this_is_table"
}
