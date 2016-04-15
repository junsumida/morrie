package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
)

type ApiVersion struct {
	IOSVersions     []interface{} `json:"ios"`
	AndroidVersions []interface{} `json:"android"`
}

func response(rw http.ResponseWriter, request *http.Request) {
	buf, err := ioutil.ReadFile("../../yaml/test.yml")
	if err != nil {
		fmt.Println("file read error")
		return
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}

	ivrs := m["defaults"].(map[interface{}]interface{})["api_version_normal"].([]interface{})
	avrs := m["defaults"].(map[interface{}]interface{})["android_api_version_normal"].([]interface{})
	fmt.Println(ivrs)
	fmt.Println(avrs)

	apivr := new(ApiVersion)
	apivr.IOSVersions = ivrs
	apivr.AndroidVersions = avrs

	js, err := json.Marshal(apivr)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func main() {
	http.HandleFunc("/api_version", response)
	http.ListenAndServe(":30303", nil)
}
