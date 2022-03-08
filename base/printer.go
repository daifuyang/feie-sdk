/**
** @创建时间: 2020/12/26 12:46 上午
** @作者　　: return
** @描述　　:
 */
package base

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/feieSdk/util"
)

type Printer struct {
}

type Data struct {
	Ok []string `json:"ok"`
	No []string `json:"no"`
}

type reqResult struct {
	Result
	Data `json:"data"`
}

type printResult struct {
	Result
	Data string `json:"data"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 新增打印机
 * @Date 2021/4/11 1:24:4
 * @Param
 * @return
 **/
func (rest Printer) Add(snlist string) reqResult {
	param := map[string]string{
		"apiname":        "Open_printerAddlist",
		"printerContent": snlist,
	}
	data := util.GetResult(param)

	result := reqResult{}
	err := json.Unmarshal(data, &result)

	if err != nil {
		fmt.Println(err)
	}

	return result
}


/**
 * @Author return <1140444693@qq.com>
 * @Description 删除打印机
 * @Date 2021/4/11 1:23:57
 * @Param
 * @return
 **/
func (rest Printer) Delete(snlist string) reqResult {
	param := map[string]string{
		"apiname": "Open_printerDelList",
		"snlist":  snlist,
	}
	data := util.GetResult(param)

	result := reqResult{}
	err := json.Unmarshal(data, &result)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	return result
}



func (rest Printer) Printer(sn string, content string, times int) printResult {

	param := map[string]string{
		"apiname": "Open_printMsg",
		"sn":      sn,
		"content": content,
		"times":   string(times),
	}

	data := util.GetResult(param)

	result := printResult{}
	err := json.Unmarshal(data, &result)

	if err != nil {
		fmt.Println(err)
	}

	return result

}