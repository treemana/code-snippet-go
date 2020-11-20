package upper

import (
	"log"

	"github.com/upper/db/v4/adapter/sqlite"
)

type DagInfo struct {
	SystemID     uint64 `db:"system_id"`
	DateDelay    int8   `db:"date_delay"`
	DagID        string `db:"dag_id"`
	Type         string `db:"type"`
	SiteID       string `db:"site_id"`
	Brand        string `db:"brand"`
	City         string `db:"city"`
	Store        string `db:"store"`
	Notice       bool   `db:"notice"`
	Status       bool   `db:"status"`
	DoneTemplate string `db:"done_template"`
	CreateTime   int64  `db:"create_time"`
	UpdateTime   int64  `db:"update_time"`
}

func Main() {
	sd, err := sqlite.Open(sqlite.ConnectionURL{
		Database: "/Users/treeman-zhou/temp/airflow.db",
		Options:  nil,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = sd.Close()
	}()

	var dagInfos []DagInfo
	if err = sd.Collection("dag_info").Find().All(&dagInfos); err != nil {
		log.Fatal(err)
	}
	for _, info := range dagInfos {
		log.Println(info)
	}

}
