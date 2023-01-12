package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/zhangrt/voyager1_platform/model/system"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	builder      SQLBuilder
	dbtest       *gorm.DB
	dbsafetytest *gorm.DB
)

func CraeteDB() *gorm.DB {

	pgsqlConfig := postgres.Config{
		DSN:                  "host=" + "172.30.0.112" + " user=" + "postgres" + " password=" + "postgres" + " dbname=" + "voyager1" + " port=" + "25431" + " " + "TimeZone=Asia/Shanghai", // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}
func TestSQLBuilder(t *testing.T) {

	dbtest = CraeteDB()
	t.Log(dbtest)
	db := SQLAdapterObj.Adapter("test", dbtest).
		Model("test", &system.Vo1Person{}).
		Where("test", "id", "=", nil).
		Where("test", "name", "-like", "_1 or ^3=3%").
		Where("test", "age", "like-", "").
		Where("test", "gender", "in", []string{"1", "2", "5"}).
		Where("test", "last_login_lime", "<=", time.Now()).
		Where("test", "last_login_lime", ">=", time.Time{}).
		Where("test", "phone", "in", "0,1").
		Page("test", 1, 10).
		Order("test", "code-desc,type-asc").
		Go("test")
	var personList []system.Vo1Person
	var total int64

	err := db.Count(&total).Find(&personList).Error
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Total:  %v \n", total)
	t.Log(db)
	db2 := builder.Get("test")
	t.Log(db2)
}

func TestSQLBuilderSafety(t *testing.T) {

	dbsafetytest = CraeteDB()
	db := builder.Adapter("test", dbsafetytest).
		Model("test", &system.Vo1Person{}).
		Where("test", "id", "=", nil).
		Where("test", "name", "-like", "_1 or ^3=3%").
		Where("test", "age", "like-", "").
		Where("test", "gender", "in", []string{"1", "2", "5"}).
		Where("test", "last_login_lime", "<=", time.Now()).
		Where("test", "last_login_lime", ">=", time.Time{}).
		Where("test", "phone", "in", "0,1").
		Page("test", 1, 10).
		Order("test", "code-desc,type-asc").
		Go("test")
	var personList []system.Vo1Person
	var total int64

	err := db.Count(&total).Find(&personList).Error
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Total:  %v \n", total)
	t.Log(db)
	db2 := builder.Get("test")
	t.Log(db2)
}
