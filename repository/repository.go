package repository

import (
	"fmt"
	"github.com/megoo/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
	iLog "log"

	"github.com/micro/go-micro/debug/log"
	"os"
	"time"
)

/**
 * mysql数据库管理
 * 初始化mysql
 */

const (
	sqlDSNFormat = "%s:%s@tcp(%s)/%s?" +
		"clientFoundRows=false&parseTime=true&timeout=5s&charset=utf8mb4&collation=utf8mb4_general_ci&loc=Local"
)

// NewGormDB creates gorm Database data
func NewGormDB(user, password, host, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(sqlDSNFormat, user, password, host, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gLogger.New(
			iLog.New(os.Stdout, "\r\n", iLog.LstdFlags), // io writer
			gLogger.Config{
				SlowThreshold: time.Second,  // 慢 SQL 阈值
				LogLevel:      gLogger.Info, // Log level TODO 设置可配
				Colorful:      true,         // 禁用彩色打印
			},
		),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(config.NewConfig().GetConf().DB.MaxIdleConn)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(config.NewConfig().GetConf().DB.MaxOpenConn)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(config.NewConfig().GetConf().DB.MaxLifetime))

	return db, nil
}

var (
	repo Repo
)

type Repo struct {
	db *gorm.DB
}

// NewRepo new repository
func NewRepo() *Repo {
	return &repo
}

// Init init repository
func Init(dbConf *config.Database) error {
	db, err := NewGormDB(dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Database)
	if err != nil {
		log.Errorf("init gorm fail, %s", err.Error())
		return err
	}

	db.AutoMigrate(
		&User{},
	)
	repo.db = db
	return nil
}

// Release  repository
func Release() {
	if repo.db != nil {
		sqlDB, err := repo.db.DB()
		if err != nil {
			log.Errorf("release database fail err: %s", err.Error())
		}
		err = sqlDB.Close()
		if err != nil {
			log.Errorf("release database fail err: %s", err.Error())
		}
	}
}

func (r *Repo) DB() *gorm.DB {
	return r.db
}

// User 用户
type User struct {
	gorm.Model
	UserName string `gorm:"column:user_name; type:varchar(50)" json:"UserName"` // 用户名称
	UserType string `gorm:"column:user_type; type:varchar(10)" json:"UserType"` // 用户类型
	Role     string `gorm:"column:role; type:varchar(50)" json:"Role"`          // 用户角色
}

func (User) TableName() string {
	return "tbl_user"
}
