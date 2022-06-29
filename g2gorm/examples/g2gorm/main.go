package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"

	lc "github.com/dijkvy/geocentric/config"
	"github.com/dijkvy/geocentric/g2gorm"
	"github.com/dijkvy/geocentric/log/zaplog"
)

var (
	configPath = flag.String("config", "g2gorm/examples/g2gorm/configs/config.toml", "配置文件路径")
)

type Config struct {
	LoggerConfig lc.ZapConfig    `toml:"logger_config"`
	DBConfig     lc.EngineConfig `toml:"db_config"`
}

type UserInfo struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type g2Hook struct {
}

type anyKey int

var ak anyKey

// Fire 回调钩子函数
func (c *g2Hook) Fire(e g2gorm.Entry) (err error) {
	ctx := e.Context()
	if ctx.Value(ak) == nil {
		return nil
	}
	return e.AppendField(zap.String("trace", fmt.Sprintf("%v", ctx.Value(ak))))
}

func (c *g2Hook) Levels() (levels []logger.LogLevel) {

	return append(levels, logger.Info, logger.Warn, logger.Warn, logger.Error)
}

func main() {

	flag.Parse()
	var conf Config
	var err error
	if _, err = toml.DecodeFile(*configPath, &conf); err != nil {
		panic(err)
	}
	fmt.Println(conf.DBConfig.String())
	fmt.Println(conf.LoggerConfig.String())
	var helper g2gorm.DBHelper
	var zl *zap.Logger
	if zl, err = zaplog.NewZapLogger(&conf.LoggerConfig); err != nil {
		panic(err)
	}
	defer zl.Sync()

	if helper, err = g2gorm.NewDBHelper(&conf.DBConfig,
		g2gorm.WithLoggerV2(zl, conf.DBConfig.GetOrmLogConfig(), &g2Hook{})); err != nil {
		panic(err)
	}

	// 清空表格
	ctx := context.TODO()
	helper.WithContext(context.TODO()).Exec("truncate  user_info")

	// 向表格中插入数据
	helper.WithContext(context.WithValue(ctx, ak, rand.ExpFloat64())).Table("user_info").Create(&UserInfo{CreatedAt: time.Now(), DeletedAt: time.Now(), UpdatedAt: time.Now()})
	// 触发慢查询
	_ = helper.WithContext(context.WithValue(ctx, ak, rand.ExpFloat64())).Exec("select sleep(4);").Error

}
