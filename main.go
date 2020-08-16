package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Options main options
type Options struct {
	Host        string
	MQTTPrefix  string
	VirtualUser int
	UC          UserConfiguration
	DC          DeviceConfiguration
	SC          ServiceConfiguration
}

// UserConfiguration represent to configuration of users
type UserConfiguration struct {
	RunServiceDelay    float64
	GetDataInterval    float64
	RealtimeInterval   float64
	RealtimePeriod     float64
	RealtimeHBInterval int
	StartServiceDelay  float64
}

// DeviceConfiguration represent to configuration of devices
type DeviceConfiguration struct {
	RealtimeInterval float64
	HistoryInterval  float64
}

// ServiceConfiguration represent to configuration of services
type ServiceConfiguration struct {
	RealtimeLength int
}

// ConfigureOptions main config initial
func ConfigureOptions(vu int) (*Options, error) {
	testcase := fmt.Sprintf("testing%d", vu)
	opts := &Options{
		Host:        viper.GetString("api.endpoint"),
		MQTTPrefix:  viper.GetString("mqtt.prefix"),
		VirtualUser: viper.GetInt(fmt.Sprintf("%s.vu", testcase)),
		UC: UserConfiguration{
			RunServiceDelay:    viper.GetFloat64(fmt.Sprintf("%s.user.runServiceDelay", testcase)),
			GetDataInterval:    viper.GetFloat64(fmt.Sprintf("%s.user.getDataInterval", testcase)),
			RealtimeInterval:   viper.GetFloat64(fmt.Sprintf("%s.user.realtime.interval", testcase)),
			RealtimePeriod:     viper.GetFloat64(fmt.Sprintf("%s.user.realtime.period", testcase)),
			RealtimeHBInterval: viper.GetInt(fmt.Sprintf("%s.user.realtime.heartbeatInterval", testcase)),
			StartServiceDelay:  viper.GetFloat64(fmt.Sprintf("%s.user.startServiceDelay", testcase)),
		},
		DC: DeviceConfiguration{
			RealtimeInterval: viper.GetFloat64(fmt.Sprintf("%s.device.realtimeInterval", testcase)),
			HistoryInterval:  viper.GetFloat64(fmt.Sprintf("%s.device.historyInterval", testcase)),
		},
		SC: ServiceConfiguration{
			RealtimeLength: viper.GetInt(fmt.Sprintf("%s.service.realtimeLength", testcase)),
		},
	}

	return opts, nil
}

func init() {
	// read config from yaml file
	viper.SetConfigName("default") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")                                                      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/go/pkg/mod/github.com/go-benchmark/config@v0.0.2") // path to look for the config file in
	err := viper.ReadInConfig()                                                   // Find and read the config file
	if err != nil {                                                               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
}
