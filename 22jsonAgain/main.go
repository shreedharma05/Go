package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON in Golang")
	encodeJson()
	decodeJson()
}

func encodeJson() {
	// create a slice of courses available using the type Course (struct)
	lcoCourses := []Course{
		{"NestJS", 299, "learnCodeOnline", "iamNest", []string{"backend-dev"}},
		{"Gin", 299, "leanrCodeOnline", "iamGin", []string{"backend-dev"}},
		{"ReactJS", 299, "learnCodeOnline", "iamReact", []string{"frontend-dev"}},
		{"Kafka", 299, "learnCodeOnline", "iamKafka", nil},
	}

	jsonData, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	// fmt.Println(jsonData)
	fmt.Println(string(jsonData))
}

func decodeJson() {
	JsonFromWeb := []byte(`
	{
		"coursename": "Gin",
		"Price": 299,
		"website": "leanrCodeOnline",
		"tags": ["backend-dev"]
	}
	`)

	var lcoCourse Course

	checkValid := json.Valid(JsonFromWeb)

	if checkValid {
		fmt.Println("JSON is Valid")
		json.Unmarshal(JsonFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON IS INVALID")
	}

	// when you want it as key value pair

	var myWebData map[string]interface{}
	json.Unmarshal(JsonFromWeb, &myWebData)
	fmt.Printf("%#v\n", myWebData)

	for k, v := range myWebData {
		fmt.Printf("The key is: %v, The value is: %v and the type is: %T\n", k, v, v)
	}
}
