package config_test

import (
	"fmt"

	"github.com/leonkaihao/cosmos/pkg/config"
)

type service struct {
	Trading  string `json:"trading"`
	Member   string `json:"member"`
	Currency string `json:"currency"`
}

func ExampleGetLocalConfig() {
	cfg, err := config.GetLocalConfig("../../testdata/config/service.abc.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := map[string]interface{}{}
	err = cfg.Unmarshal(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	var svc service
	err = cfg.UnmarshalKey("service", &svc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(
		data["kafka"].(map[string]interface{})["brokers"].([]interface{})[0].(string),
		svc.Trading)
	// Output: localhost:9092 localhost:8028
}
