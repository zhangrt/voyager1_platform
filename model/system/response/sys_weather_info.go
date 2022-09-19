package response

type SysWeatherInfo struct {
	Code        string `json:"code"`
	Temperature string `json:"temperature"`
	Location    string `json:"location"`
	Humidness   string `json:"humidness"`
	Info        string `json:"info"`
	Time        string `json:"time"`
}
