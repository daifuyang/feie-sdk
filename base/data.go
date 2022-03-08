/**
** @创建时间: 2020/12/31 4:11 下午
** @作者　　: return
** @描述　　:
 */
package base

type Result struct {
	Msg                string      `json:"msg"`
	Ret                int         `json:"ret"`
	ServerExecutedTime int         `json:"server_executed_time"`
}
