package system

import (
	"testing"

	"github.com/zhangrt/voyager1_platform/model/system/response"
)

var (
	weatherService WeatherService
)

func Test_GetWeatherInfo(t *testing.T) {

	tests := [2]struct {
		code    string
		want    response.Vo1WeatherInfo
		wantErr bool
	}{
		{
			code: "101010100",
			want: response.Vo1WeatherInfo{
				Code:        "101010100",
				Temperature: "27.9",
				Location:    "北京",
				Humidness:   "28%",
				Info:        "多云转阴",
				Time:        "18:00",
			},
			wantErr: false,
		},
		{
			code: "101220101",
			want: response.Vo1WeatherInfo{
				Code:        "101220101",
				Temperature: "20.5",
				Location:    "合肥",
				Humidness:   "98%",
				Info:        "中雨转小雨",
				Time:        "18:00",
			},
			wantErr: false,
		},
	}

	for _, data := range tests {
		t.Run(data.code, func(t *testing.T) {
			if got, err := weatherService.GetWeatherInfo(data.code); got != data.want {
				if (err != nil) != data.wantErr {
					t.Errorf("GetWeatherInfo() error = %v, wantErr %v", err, data.wantErr)
				} else {
					t.Errorf("GetWeatherInfo() = %v, want %v", got, data.want)
				}
			}
		})
	}
}
