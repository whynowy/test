package main

import (
	"bytes"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

const (
	conf = `
producer:
  maxMessageBytes: 899
consumer:
  group:
    rebalance:
      retry:
        max: 321
`
)

func GetSaramaConfigFromYAMLString(yaml string) (*sarama.Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	if err := v.ReadConfig(bytes.NewBufferString(yaml)); err != nil {
		return nil, err
	}

	c := sarama.NewConfig()
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("failed unmarshal configuration. %w", err)
	}
	return c, nil
}

func main() {
	if s, err := GetSaramaConfigFromYAMLString(conf); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("producer.maxMessageBytes: %v\n", s.Producer.MaxMessageBytes)                     // should be 899
		fmt.Printf("consumer.group.rebalance.retry.max: %v\n", s.Consumer.Group.Rebalance.Retry.Max) // should be 321
		fmt.Printf("admin.retry.max: %v\n", s.Admin.Retry.Max)                                       // should be 5, which is an unchanged property from sarama.NewConfig()
	}

}
