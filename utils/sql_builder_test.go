package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/zhangrt/voyager1_platform/model/demo"

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
		DSN:                  "host=" + "192.168.244.142" + " user=" + "root" + " password=" + "123123" + " dbname=" + "test" + " port=" + "26257" + " " + "TimeZone=Asia/Shanghai", // DSN data source name
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
	db := SQLAdapter.Adapter("test", dbtest).
		Model("test", &demo.Facility{}).
		Where("test", "id", "=", nil).
		Where("test", "name", "-like", "_1 or ^3=3%").
		Where("test", "code", "like-", "").
		Where("test", "type", "in", []string{"1", "2", "5"}).
		Where("test", "alarm_time", "<=", time.Now()).
		Where("test", "alarm_time", ">=", time.Time{}).
		Where("test", "status", "in", "0,1").
		Page("test", 1, 10).
		Order("test", "code-desc,type-asc").
		Go("test")
	var facilityList []demo.Facility
	var total int64

	err := db.Count(&total).Find(&facilityList).Error
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
		Model("test", &demo.Facility{}).
		Where("test", "id", "=", nil).
		Where("test", "name", "like", "2").
		Where("test", "code", "-like-", "1").
		Where("test", "type", "in", []string{"1", "2", "5"}).
		Where("test", "alarm_time", "<=", time.Now()).
		Where("test", "alarm_time", ">=", time.Time{}).
		Where("test", "status", "in", "0,1").
		Page("test", 1, 10).
		Order("test", "code-desc,type-asc").
		Go("test")
	var facilityList []demo.Facility
	var total int64

	err := db.Count(&total).Find(&facilityList).Error
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Total:  %v \n", total)
	t.Log(db)
	db2 := builder.Get("test")
	t.Log(db2)
}
