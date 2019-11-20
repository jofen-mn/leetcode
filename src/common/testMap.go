package main

import (
	"fmt"
	"encoding/json"
	"reflect"
	"log"
)

func main() {
	var m = make(map[int]string)
	a := 1
	b := a + 2
	c := b - 2
	m[a] = "aa"
	m[b] = "bb"
	m[c] = "cc"
	for v, _ := range m {
		fmt.Println(v)
	}

	type NameValue struct {
		Name string `json:"name"`
		Value int `json:"value"`
	}
	jsonStr := `{"name":"xxx", "value":233}`
	mydata := NameValue{}
	//jsonStrB, _ := json.Marshal(jsonStr)
	err := json.Unmarshal([]byte(jsonStr), &mydata)
	if err != nil {
		log.Println(err)
	}
	log.Println(reflect.TypeOf(mydata.Value))

	//jsonStr := map[string]interface{}{"name": "  ", "value": 233}
	////jsonStr := `{"name":"xxx", "value":233.00000}`
	//mydata := &serializer.NameValue{}
	//jsonStrB, _ := json.Marshal(jsonStr)
	//err := json.Unmarshal(jsonStrB, mydata)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(reflect.TypeOf(mydata.Value))

	var tts []TT
	log.Println(tts)


	aaaaa := 0.26
	log.Println(aaaaa)
	log.Println(aaaaa * 100)

}

type TT struct {
	Te string `json:"te"`
}


