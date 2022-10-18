package config

import (
	config "github.com/zhangrt/voyager1_core/config"
)

type Server struct {
	JWT     config.JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	AUTHKey config.AUTHKey `mapstructure:"auth-key" json:"auth-key" yaml:"auth-key"`
	Zap     Zap            `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis          `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email   Email          `mapstructure:"email" json:"email" yaml:"email"`
	System  System         `mapstructure:"system" json:"system" yaml:"system"`
	Casbin  config.Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`

	// gorm
	Mysql     Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql     Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Cockroach Cockroach       `mapstructure:"cockroach" json:"cockroach" yaml:"cockroach"`
	DBList    []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`

	// cache
	Cache config.Cache `mapstructure:"cache" json:"cache" yaml:"cache"`

	// oss
	Minio config.Minio `mapstructure:"minio" json:"minio" yaml:"minio"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`

	// Zinx
	Zinx config.Zinx `mapstructure:"zinx" json:"zinx" yaml:"zinx"`

	Grpc config.Grpc `mapstructure:"grpc" json:"grpc" yaml:"grpc"`

	// 服务网格
	ServiceMesh ServiceMesh `mapstructure:"service-mesh" json:"service-mesh" yaml:"service-mesh"`
}
