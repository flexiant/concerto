package utils

import (
	"fmt"
	"github.com/codegangsta/cli"
	"reflect"
)

// FlagConvertParams converts cli parameters in API callable params
func FlagConvertParams(c *cli.Context) *map[string]string {
	v := make(map[string]string)
	for _, flag := range c.FlagNames() {
		if c.IsSet(flag) {
			v[flag] = c.String(flag)
		}
	}
	return &v
}

// ItemConvertParams converts API items into map of strings
func ItemConvertParams(item interface{}) (*map[string]string, error) {

	it := reflect.ValueOf(item)
	nf := it.NumField()
	v := make(map[string]string)

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
