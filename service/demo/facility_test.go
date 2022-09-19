package demo

import (
	"testing"
	"time"

	"github.com/zhangrt/voyager1_platform/model/demo"
)

type TestFacility struct {
	param   demo.Facility
	want    demo.Facility
	wantErr bool
}

var (
	service FacilityService
	adds    []TestFacility
	remove  TestFacility
)

func Test_AddFacility(t *testing.T) {

	now := time.Now()
	adds := [1]TestFacility{
		{
			param: demo.Facility{
				Name:   now.Format("2006-01-02 15:04:05"),
				Code:   now.Format("2006-01-02 15:04:05"),
				Type:   "0",
				Status: "0",
			},
			want: demo.Facility{
				Name:   now.Format("2006-01-02 15:04:05"),
				Code:   now.Format("2006-01-02 15:04:05"),
				Type:   "0",
				Status: "0",
			},
			wantErr: false,
		},
	}

	for _, data := range adds {
		t.Run("AddFacility", func(t *testing.T) {
			if got, err := service.AddFacility(data.param); got != data.want {
				if (err != nil) != data.wantErr {
					t.Errorf("AddFacility() error = %v, wantErr %v", err, data.wantErr)
				} else {
					t.Errorf("AddFacility() = %v, want %v", got, data.want)
				}
			}
		})

	}
}

func Test_RemoveFacility(t *testing.T) {
	remove := TestFacility{
		param: demo.Facility{
			Type:   "3",
			Status: "2",
		},
		want: demo.Facility{
			Type:   "3",
			Status: "2",
		},
		wantErr: false,
	}
	t.Run("RemoveFacility", func(t *testing.T) {
		if got, err := service.RemoveFacility(remove.param); got != remove.want {
			if (err != nil) != remove.wantErr {
				t.Errorf("RemoveFacility() error = %v, wantErr %v", err, remove.wantErr)
			} else {
				t.Errorf("RemoveFacility() = %v, want %v", got, remove.want)
			}
		}
	})
}
