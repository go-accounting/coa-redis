package main

import (
	coaredis "github.com/go-accounting/coa-redis"
)

func NewKeyValueStore(config map[string]interface{}, ss ...*string) (interface{}, error) {
	v := config["NewKeyValueStore/Addresses"].([]interface{})
	addrs := make([]string, len(v))
	for i, a := range v {
		addrs[i] = a.(string)
	}
	master, _ := config["NewKeyValueStore/Master"].(string)
	return coaredis.NewStore(master, addrs, ss[0])
}
