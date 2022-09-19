package statistics

import (
	"testing"

	"github.com/zhangrt/voyager1_platform/model/statistics"
)

var (
	service StatisticsFacilityService
)

func Test_StatisticsFacilityByType(t *testing.T) {

	tests := [2]struct {
		req     statistics.StatisticsFacilityByType
		want    []statistics.StatisticsFacilityType
		wantErr bool
	}{
		{
			req: statistics.StatisticsFacilityByType{
				Types: "0,1",
			},
			want:    []statistics.StatisticsFacilityType{},
			wantErr: false,
		},
		{
			req: statistics.StatisticsFacilityByType{
				Types: "",
			},
			want:    []statistics.StatisticsFacilityType{},
			wantErr: false,
		},
	}

	for _, data := range tests {
		t.Run("StatisticsFacilityByType", func(t *testing.T) {
			if got, err := service.StatisticsFacilityByType(data.req); got == nil {
				if (err != nil) != data.wantErr {
					t.Errorf("StatisticsFacilityByType() error = %v, wantErr %v", err, data.wantErr)
				} else {
					t.Errorf("StatisticsFacilityByType() = %v, want %v", got, data.want)
				}
			}
		})
	}
}
