package statistics

import (
	"time"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/statistics"
)

type StatisticsFacilityRepository struct{}

func (r *StatisticsFacilityRepository) StatisticsFacilityByType(param statistics.StatisticsFacilityByType) (array []statistics.StatisticsFacilityType, err error) {
	var db = global.GS_DB
	var results []statistics.StatisticsFacilityType = []statistics.StatisticsFacilityType{}

	if db == nil {
		return results, err
	}

	originSql := "SELECT type, COUNT(type) AS num FROM facility WHERE deleted_at is NULL "
	nilTime := time.Time{}
	var sqlType int
	if param.StartTime == nilTime && param.EndTime == nilTime {
		sqlType = 0
	} else if param.StartTime != nilTime && param.EndTime == nilTime {
		originSql = originSql + " AND created_at > ? "
		sqlType = 1
	} else if param.EndTime != nilTime && param.StartTime == nilTime {
		originSql = originSql + " AND created_at < ? "
		sqlType = 2
	} else {
		originSql = originSql + " AND created_at > ? " + " AND created_at < ? "
		sqlType = 3
	}
	originSql = originSql + " GROUP BY type ORDER BY type ASC"

	switch sqlType {
	case 0:
		db = db.Raw(originSql)
	case 1:
		db = db.Raw(originSql, param.StartTime)
	case 2:
		db = db.Raw(originSql, param.EndTime)
	case 3:
		db = db.Raw(originSql, param.StartTime, param.EndTime)
	}
	rows, err := db.Rows()
	if err != nil {
		global.GS_LOG.Error(err.Error())
	}
	defer rows.Close()

	if rows != nil {
		for rows.Next() {

			db.ScanRows(rows, &results)

		}
	}

	return results, err
}
