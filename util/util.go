/**
** @创建时间: 2020/9/7 9:46 上午
** @作者　　: return
** @描述　　:
 */
package util

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/gincmf/feieSdk"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//封装请求库
func Request(method string, url string, body io.Reader, header map[string]string) (int, []byte) {
	client := &http.Client{}
	switch method {
	case "get", "GET":
		method = "GET"
	case "post", "POST":
		method = "POST"
	case "put", "PUT":
		method = "PUT"
	case "delete", "DELETE":
		method = "POST"
	}
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println("http错误", err)
	}

	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Accept-Encoding", "gzip, deflate, br")
	r.Header.Add("Content-Length", "0")
	r.Header.Add("Cache-Control", "no-cache")
	r.Header.Add("Content-Type","application/x-www-form-urlencoded; param=value")

	for k, v := range header {
		r.Header.Add(k, v)
	}
	response, err := client.Do(r)

	defer response.Body.Close()

	var data []byte = nil

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(response.Body)
		for {
			buf := make([]byte, 1024)
			n, err := reader.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}
			data = append(data, buf...)
		}
	default:
		data, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("err", err.Error())
		}
	}

	index := bytes.IndexByte(data, 0)

	if index > 0 {
		data = data[:index]
	}

	return response.StatusCode, data
}

func GetResult(paramsMap map[string]string) []byte {
	return request(nil, paramsMap, nil)
}

func request(params map[string]string, body map[string]string, header map[string]string) []byte {

	itime := time.Now().Unix()
	stime := strconv.FormatInt(itime, 10)
	options := feieSdk.Options()
	user := options.User
	ukey := options.Ukey
	sig := SHA1(user + ukey + stime) //生成签名

	baseUrl := options.Url // 网关

	body["user"] = options.User
	body["stime"] = stime
	body["sig"] = sig

	body["debug"] = "0"


	data := url.Values{}
	for k,v := range body{
		data.Add(k,v)
	}

	_, result := Request("POST", baseUrl, bytes.NewBufferString(data.Encode()), header)
	return result
}
