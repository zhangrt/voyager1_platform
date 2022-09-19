package utils

import (
	"fmt"
	"github.com/zhangrt/voyager1_platform/model/demo"
	"testing"
	"time"

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
	db := builder.Adapter(dbtest).
		Model(&demo.Facility{}).
		Where("id", "=", nil).
		Where("name", "-like", "_1 or ^3=3%").
		Where("code", "like-", "").
		Where("type", "in", []string{"1", "2", "5"}).
		Where("alarm_time", "<=", time.Now()).
		Where("alarm_time", ">=", time.Time{}).
		Where("status", "in", "0,1").
		Page(1, 10).
		Order("code-desc,type-asc").
		Go()
	var facilityList []demo.Facility
	var total int64

	err := db.Count(&total).Find(&facilityList).Error
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Total:  %v \n", total)
	t.Log(db)
	db2 := builder.Get()
	t.Log(db2)
}

func TestSQLBuilderSafety(t *testing.T) {

	dbsafetytest = CraeteDB()
	db := builder.AdapterSafety("test", dbsafetytest).
		ModelSafety("test", &demo.Facility{}).
		WhereSafety("test", "id", "=", nil).
		WhereSafety("test", "name", "like", "2").
		WhereSafety("test", "code", "-like-", "1").
		WhereSafety("test", "type", "in", []string{"1", "2", "5"}).
		WhereSafety("test", "alarm_time", "<=", time.Now()).
		WhereSafety("test", "alarm_time", ">=", time.Time{}).
		WhereSafety("test", "status", "in", "0,1").
		PageSafety("test", 1, 10).
		OrderSafety("test", "code-desc,type-asc").
		GoSafety("test")
	var facilityList []demo.Facility
	var total int64

	err := db.Count(&total).Find(&facilityList).Error
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Total:  %v \n", total)
	t.Log(db)
	db2 := builder.Get()
	t.Log(db2)
}
