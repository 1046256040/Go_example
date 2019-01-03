package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

var pn = fmt.Println

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	pn(string(bolB))

	intB, _ := json.Marshal(1)
	pn(string(intB))

	fltB, _ := json.Marshal(2.34)
	pn(string(fltB))

	strB, _ := json.Marshal("gopher")
	pn(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	pn(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	pn(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"peach", "peanut", "apple"}}
	res1B, _ := json.Marshal(res1D)
	pn(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"peach", "peanut", "apple"}}
	res2B, _ := json.Marshal(res2D)
	pn(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var data map[string]interface{}

	if err := json.Unmarshal(byt, &data); err != nil {
		panic("!!! json.Unmarsha1 failed !!!")
	}
	pn(data["num"])

	strs := data["strs"].([]interface{})
	str1 := strs[0].(string)
	pn(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	pn(res)
	pn(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
