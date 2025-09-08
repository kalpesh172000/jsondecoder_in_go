package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func classify(parents string, result interface{}) {
	obj, ok := result.(map[string]interface{})
	if !ok {
		fmt.Println("Top level of json is not a object")
		return
	}
	for k, v := range obj {
		switch v := v.(type) {
		case string:
			fmt.Println(parents+k, " : ", v)
		case float64:
			fmt.Println(parents+k, " : ", v)
		case bool:
			fmt.Println(parents+k, " : ", v)
		case []any:
			fmt.Println(parents+k, " : ", v)
		case map[string]any:

			dummy := parents
			parents = parents + k + "."
			classify(parents, v)
			parents = dummy

		default:
			fmt.Println(k, " is an unknown type")
		}
	}
}

func main() {
	filename := os.Args[1]
	data,err1 := os.ReadFile(filename)
	if err1 != nil {
		fmt.Println("error occured", err1)
	}

	var result interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	//fmt.Println("Raw : ", result)
	parent := ""
	classify(parent, result)
}
