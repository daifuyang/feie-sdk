/**
** @创建时间: 2020/12/26 12:47 上午
** @作者　　: return
** @描述　　:
 */
package feieSdk

import "reflect"

type baseOptions struct {
	User string `json:"user"`
	Ukey string `json:"ukey"`
	Sn   string `json:"sn"`
	Url  string `json:"url"`
}

var options *baseOptions

func NewOptions(params map[string]string) baseOptions {
	options = &baseOptions{
		User: params["user"],
		Ukey: params["ukey"],
		Url:  params["url"],
	}
	return *options
}

func SetOption(key string, val string) {
	oPoint := reflect.ValueOf(options)
	field := oPoint.Elem().FieldByName(key)
	field.SetString(val)
}

func Options() baseOptions {
	if options == nil {
		panic("配置为初始化！")
	}
	return *options

}
