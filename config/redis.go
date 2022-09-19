package config

type Redis struct {
	ClusterMod  bool     `mapstructure:"cluster-mod" json:"cluster-mod" yaml:"cluster-mod"` // cluster
	Addr        string   `mapstructure:"addr" json:"addr" yaml:"addr"`                      // 服务器地址:单机
	ClusterAddr []string `mapstructure:"addrs" json:"addrs" yaml:"addrs"`                   // 服务器地址:集群
	Username    string   `mapstructure:"username" json:"username" yaml:"username"`          // 用户
	Password    string   `mapstructure:"password" json:"password" yaml:"password"`          // 密码
}
