package system

import (
	"encoding/json"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/system/response"

	// Import resty into your code and refer it as `resty`.
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type WeatherService struct{}

/* Output
Response Info:
  Error      : <nil>
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/2.0
  Time       : 457.034718ms
  Received At: 2020-09-14 15:35:29.784681 -0700 PDT m=+0.458137045
  Body       :
  {
    "args": {},
    "headers": {
      "Accept-Encoding": "gzip",
      "Host": "httpbin.org",
      "User-Agent": "go-resty/2.4.0 (https://github.com/go-resty/resty)",
      "X-Amzn-Trace-Id": "Root=1-5f5ff031-000ff6292204aa6898e4de49"
    },
    "origin": "0.0.0.0",
    "url": "https://httpbin.org/get"
  }

Request Trace Info:
  DNSLookup     : 4.074657ms
  ConnTime      : 381.709936ms
  TCPConnTime   : 77.428048ms
  TLSHandshake  : 299.623597ms
  ServerTime    : 75.414703ms
  ResponseTime  : 79.337µs
  TotalTime     : 457.034718ms
  IsConnReused  : false
  IsConnWasIdle : false
  ConnIdleTime  : 0s
  RequestAttempt: 1
  RemoteAddr    : 3.221.81.55:443
*/
func (weatherService *WeatherService) GetWeatherInfo(cityCode string) (sysWeatherInfo response.Vo1WeatherInfo, err error) {
	sxUrl := "http://www.weather.com.cn/data/sk/" + cityCode + ".html"
	cityInfoUrl := "http://www.weather.com.cn/data/cityinfo/" + cityCode + ".html"

	client := resty.New()
	/*
		{
		  "weatherinfo": {
		    "city": "北京", //城市名称
		    "cityid": "101010100", // 城市代码
		    "temp": "27.9", // 温度
		    "WD": "南风", //风向
		    "WS": "小于3级", //风力
		    "SD": "28%", //湿度
		    "AP": "1002hPa", //气压
		    "njd": "暂无实况",
		    "WSE": "<3",//风速
		    "time": "17:55", //发布时间
		    "sm": "2.1",
		    "isRadar": "1",
		    "Radar": "JC_RADAR_AZ9010_JB"
		  }
		}
	*/
	sxResp, err := client.R().
		EnableTrace().
		Get(sxUrl)

	jsonSx := sxResp.Body()
	var sxMap map[string]map[string]string
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal(jsonSx, &sxMap); err != nil {
		global.GS_LOG.Error("Unmarshal sx faield", zap.Error(err))
	}
	/*
			{
		  "weatherinfo": {
		    "city": "北京",//城市名称
		    "cityid": "101010100",// 城市代码
		    "temp1": "18℃", //最低气温
		    "temp2": "31℃", //最高气温
		    "weather": "多云转阴", //天气情况
		    "img1": "n1.gif",
		    "img2": "d2.gif",
		    "ptime": "18:00" //发布时间
		 }
		}
	*/
	cityResp, err := client.R().
		EnableTrace().
		Get(cityInfoUrl)
	jsonCity := cityResp.Body()
	var cityMap map[string]map[string]string
	if err := json.Unmarshal(jsonCity, &cityMap); err != nil {
		global.GS_LOG.Error("Unmarshal cityinfo faield", zap.Error(err))
	}
	client.GetClient().CloseIdleConnections()
	// global.GS_LOG.Info("status", zap.String(sxUrl, sxResp.Status()))
	// global.GS_LOG.Info("status", zap.String(cityInfoUrl, cityResp.Status()))

	sysWeatherInfo = response.Vo1WeatherInfo{
		Code:        cityCode,
		Temperature: sxMap["weatherinfo"]["temp"],
		Location:    sxMap["weatherinfo"]["city"],
		Humidness:   sxMap["weatherinfo"]["SD"],
		Info:        cityMap["weatherinfo"]["weather"],
		Time:        cityMap["weatherinfo"]["ptime"],
	}

	return sysWeatherInfo, err
}
