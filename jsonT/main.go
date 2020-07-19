package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	simplejson "github.com/bitly/go-simplejson"
)

var jsonStr string = `{"rc" : 0,
		"error" : "Success",
		"type" : "stats",
		"progress" : 100,
		"job_status" : "COMPLETED",
		"result" : {
			"total_hits" : 803254,
			"starttime" : 1528434707000,
			"endtime" : 1528434767000,
			"fields" : [ ],
			"timeline" : {
			"interval" : 1000,
			"startTs" : 1528434707000,
			"end_ts" : 1528434767000,
			"rows" : [ {
				"startTs" : 1528434707000,
				"end_ts" : 1528434708000,
				"number" : "x12887"
			}, {
				"startTs" : 1528434720000,
				"end_ts" : 1528434721000,
				"number" : "x13028"
			}, {
				"startTs" : 1528434721000,
				"end_ts" : 1528434722000,
				"number" : "x12975"
			}, {
				"startTs" : 1528434722000,
				"end_ts" : 1528434723000,
				"number" : "x12879"
			}, {
				"startTs" : 1528434723000,
				"end_ts" : 1528434724000,
				"number" : "x13989"
			} ],
			"total" : 803254
			},
			"total" : 8
		}
}`

func testJSON() {
	res, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	//获取json字符串中的 result 下的 timeline 下的 rows 数组
	rows, err := res.Get("result").Get("timeline").Get("rows").Array()

	//遍历rows数组
	for _, row := range rows {
		//对每个row获取其类型，每个row相当于 C++/Golang 中的map、Python中的dict
		//每个row对应一个map，该map类型为map[string]interface{}，也即key为string类型，value是interface{}类型
		if eachMap, ok := row.(map[string]interface{}); ok {
			//可以看到eachMap["startTs"]类型是json.Number
			//而json.Number是golang自带json库中decode.go文件中定义的: type Number string
			//因此json.Number实际上是个string类型
			fmt.Println(reflect.TypeOf(eachMap["startTs"]))
			if startTs, ok := eachMap["startTs"].(json.Number); ok {
				startTsInt, err := strconv.ParseInt(string(startTs), 10, 0)
				if err == nil {
					fmt.Println(startTsInt)
				}
			}

			if number, ok := eachMap["number"].(string); ok {
				fmt.Println(number)
			}
		}
	}
}

func main() {
	testJSON()
}
