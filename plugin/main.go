package main

import (
	coaredis "github.com/go-accounting/coa-redis"
)

func NewKeyValueStore(settings map[string]interface{}, sp *string) (interface{}, error) {
	v := settings["Addresses"].([]interface{})
	addrs := make([]string, len(v))
	for i, a := range v {
		addrs[i] = a.(string)
	}
	master, _ := settings["Master"].(string)
	return coaredis.NewStore(master, addrs, sp)
}
