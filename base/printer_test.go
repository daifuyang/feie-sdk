/**
** @创建时间: 2020/12/26 1:31 上午
** @作者　　: return
** @描述　　:
 */
package base

import (
	"fmt"
	"github.com/gincmf/feieSdk"
	"math"
	"strings"
	"testing"
)

func TestPrinter_Add(t *testing.T) {

	op := map[string]string{
		"user": "1140444693@qq.com",
		"ukey": "FGmRHjDFZeFDZaHg",
		"url":  "http://api.feieyun.cn/Api/Open/",
	}

	feieSdk.NewOptions(op)

	new(Printer).Add("960511811#3aps6y8r")
}

func TestPrinter_Delete(t *testing.T) {

	op := map[string]string{
		"user": "1140444693@qq.com",
		"ukey": "FGmRHjDFZeFDZaHg",
		"url":  "http://api.feieyun.cn/Api/Open/",
	}

	feieSdk.NewOptions(op)
	new(Printer).Delete("960511811")
}

func TestPrinter_Printer(t *testing.T) {
	op := map[string]string{
		"user": "1140444693@qq.com",
		"ukey": "FGmRHjDFZeFDZaHg",
		"url":  "http://api.feieyun.cn/Api/Open/",
	}

	feieSdk.NewOptions(op)

	order := []map[string]string{
		{
			"title": "超级大鸡排",
			"count": "1",
			"fee":   "20.00",
		},
		{
			"title": "黄焖鸡米饭(小份,微辣)",
			"count": "1",
			"fee":   "26.05",
		},
		{
			"title": "麻辣鲜香秘制火腿鸡排咖喱饭(大份,可乐,微辣)",
			"count": "100",
			"fee":   "2000.00",
		},
	}

	content := "<CB>堂食</CB><BR>"
	content += "<CB>取餐号：2819</CB><BR>"
	content += "--------------------------------<BR>"

	// 定义标题的长度
	tLen := 16 // 16个字符 （8个字符）
	cLen := 6  // 8个字符 （4个字符）
	FLen := 10  // 8个字符 （4个字符）

	for _, v := range order {

		// 统计标题总长度
		// 标题所在位置索引
		tTotal, titleArr := new(Printer).typesetting(v["title"], tLen)
		_, countArr := new(Printer).typesetting(v["count"], cLen)
		_, feeArr := new(Printer).typesetting(v["fee"], FLen)

		// 判断需要打印的行数
		pLine := math.Floor(float64(tTotal / tLen))

		for i := 0; i < int(pLine)+1; i++ {

			title := titleArr[i].Content
			titleCount := titleArr[i].Count

			// 后补空白
			var indent []string
			if titleCount < tLen {
				indentLen := tLen - titleCount

				for i := 0; i < indentLen; i++ {
					indent = append(indent, " ")
				}
				title += strings.Join(indent, "")
			}

			if i == 0 {

				countContent := new(Printer).Indent(countArr, cLen, "X")

				title += countContent

				feeContent := new(Printer).Indent(feeArr, FLen, "￥")

				title += feeContent

			}

			content += "<BOLD>" + title + "</BOLD><BR>"
			fmt.Println(title)
		}

	}

	//content += "<BOLD>这是一个十个字长标题    10    10</BOLD><BR>"
	//content += "<BOLD>招牌地瓜套餐            10   100</BOLD><BR>"
	//content += "<BOLD>123456789012345678901234567890</BOLD><BR>"

	// result :=  new(Printer).Printer("960511811", content, "1")
}
