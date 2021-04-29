package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["zhangsan"] = 100
	scoreMap["laowang"] = 90
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["laowang"])
	fmt.Println(len(scoreMap))
	fmt.Printf("type of a:%T\n", scoreMap)

	userInfo := map[string]string{
		"username": "admin",
		"password": "123456",
	}
	fmt.Println(userInfo)
	fmt.Println("=========判断是否存在===========")
	v, ok := scoreMap["zhangsan"]
	fmt.Println(v, ok)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
	fmt.Println("=======遍历=========")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	// 只遍历key
	for k := range scoreMap {
		fmt.Println(k)
	}
	fmt.Println("==========删除==============")
	score := make(map[string]int)
	score["张三"] = 90
	score["小明"] = 100
	score["娜扎"] = 60
	delete(score, "小明") //将小明:100从map中删除
	for k, v := range score {
		fmt.Println(k, v)
	}
	fmt.Println("======按照指定顺序遍历map=========")
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreDict = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreDict[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreDict {
		keys = append(keys, key)
	}
	// 对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreDict[key])
	}
	fmt.Println("=======元素为map类型的切片===========")
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index=%d, value=:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	// mapSlice[0] = make(map[string]string, 10)
	// mapSlice[0]["name"] = "小王子"
	// mapSlice[0]["password"] = "123456"
	// mapSlice[0]["address"] = "沙河"

	mapSlice[0] = map[string]string{
		"name":     "laowang",
		"password": "123456",
		"address":  "深圳",
	}

	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println(mapSlice)
	fmt.Println("======值为切片类型的map========")
	var silceMap = make(map[string][]string, 3)
	fmt.Println(silceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := silceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	silceMap[key] = value
	fmt.Println(silceMap)
	fmt.Println("========值为map类型的map=========")
	var mapMap = make(map[string]map[string]string, 0)
	mapMap["001"] = map[string]string{
		"name": "laowang",
		"age":  "18",
	}
	fmt.Println(mapMap)
}
