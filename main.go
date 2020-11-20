package main

import (
	"github.com/treeman-zhou/code-snippet-go/aws"
	"github.com/treeman-zhou/code-snippet-go/cron"
	"github.com/treeman-zhou/code-snippet-go/db/kafka"
	"github.com/treeman-zhou/code-snippet-go/db/mongo"
	"github.com/treeman-zhou/code-snippet-go/db/mysql/gorm"
	"github.com/treeman-zhou/code-snippet-go/db/redis"
	"github.com/treeman-zhou/code-snippet-go/db/sqlite/upper"
	"github.com/treeman-zhou/code-snippet-go/internal/mail"
	"github.com/treeman-zhou/code-snippet-go/localcache"
	"github.com/treeman-zhou/code-snippet-go/log/zap"
	"github.com/treeman-zhou/code-snippet-go/web"
)

func main() {
	aws.Init()

	cron.Main()

	kafka.Main()
	mongo.Main()
	gorm.Init()
	redis.Init()
	upper.Main()
	mail.Init()

	localcache.Main()

	zap.Init(zap.Config{})

	web.CHI()
	web.ECHO()
	web.GIN()
	web.IRIS()
}
