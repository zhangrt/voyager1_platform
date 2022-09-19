package config

type ServiceMesh struct {
	Namespace []Namespace `mapstructure:"namespaces" json:"namespaces" yaml:"namespaces"`
}

type Namespace struct {
	Name        string    `mapstructure:"name" json:"name" yaml:"name"`
	LoadBalance string    `mapstructure:"loadbalance" json:"loadbalance" yaml:"loadbalance"`
	Services    []Service `mapstructure:"services" json:"services" yaml:"services"`
}

type Service struct {
	Name   string `mapstructure:"name" json:"name" yaml:"name"`
	Host   string `mapstructure:"host" json:"host" yaml:"host"`
	Subset string `mapstructure:"subset" json:"subset" yaml:"subset"`
	Port   string `mapstructure:"port" json:"port" yaml:"port"`
	Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
}
