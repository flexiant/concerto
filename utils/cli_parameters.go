package utils

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"reflect"
)

// FlagConvertParams converts cli parameters in API callable params
func FlagConvertParams(c *cli.Context) *map[string]interface{} {
	v := make(map[string]interface{})
	for _, flag := range c.FlagNames() {
		if c.IsSet(flag) {
			v[flag] = c.String(flag)
		}
	}
	return &v
}

// ItemConvertParams converts API items into map of interface
func ItemConvertParams(item interface{}) (*map[string]interface{}, error) {

	it := reflect.ValueOf(item)
	nf := it.NumField()
	v := make(map[string]interface{})

	for i := 0; i < nf; i++ {
		v[it.Type().Field(i).Name] = fmt.Sprintf("%s", it.Field(i).Interface())
		// if value, ok :=  it.Field(i).Interface().(string); ok {
		// 	v[it.Type().Field(i).Name] = value
		// } else {
		// 	return nil, fmt.Errorf("Interface couldn't be converted to map of strings. Field: %s", it.Type().Field(i).Name)
		// }
	}
	return &v, nil
}

// JSONParam parses parameter to be
func JSONParam(param string) (interface{}, error) {
	var p interface{}
	err := json.Unmarshal([]byte(param), &p)
	return p, err
}
