package common

// gorm 定义示例
// Start WrapTime `gorm:"type:datetime(3);default:null;comment:时间"`
/*
pgsql
timestamp [(p)] [without time zone]
timestamp [(p)] with time zone
date
time [(p)] [without time zone]
time [(p)] with time zone
interval [fields] [(p)]
*/

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

const (
	YYYY_MM_DD          = "2006-01-02"
	YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
)

type WrapTime sql.NullTime

func (t *WrapTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		t.Valid = false
		return nil
	}

	var now time.Time
	if len(string(data)) == len(YYYY_MM_DD)+2 {
		now, err = time.ParseInLocation(`"`+YYYY_MM_DD+`"`, string(data), time.Local)
		t.Valid = true
		t.Time = now
	} else {
		now, err = time.ParseInLocation(`"`+YYYY_MM_DD_HH_MM_SS+`"`, string(data), time.Local)
		t.Valid = true
		t.Time = now
	}
	return
}

func (t WrapTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	b := make([]byte, 0, len(YYYY_MM_DD_HH_MM_SS)+2)
	b = append(b, '"')
	b = t.Time.AppendFormat(b, YYYY_MM_DD_HH_MM_SS)
	b = append(b, '"')
	return b, nil
}

func (t WrapTime) String() string {
	if !t.Valid {
		return "null"
	}
	return t.Time.Format(YYYY_MM_DD_HH_MM_SS)
}

// Value insert timestamp into mysql need this function.
func (t WrapTime) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value time.Time
func (t *WrapTime) Scan(v interface{}) error {
	return (*sql.NullTime)(t).Scan(v)
}

func NewWrapTime(t time.Time) WrapTime {
	if t.IsZero() {
		return WrapTime{Valid: false}
	}
	return WrapTime{Valid: true, Time: t}
}
