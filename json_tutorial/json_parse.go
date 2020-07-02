package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


type Config struct {
	URL []string `json:"URL"`
	Hits int    `json:"Hits"`
}

func main(){
	// Open the file and defer a close
	jsonFile, err := os.Open("stuff.json")

	if err != nil{
		fmt.Println(err)
	}

	defer jsonFile.Close()

	FileInBytes, _ := ioutil.ReadAll(jsonFile)

	a := ParseJSON(FileInBytes)

	fmt.Println(a.URL)
	fmt.Println(a.Hits)


}



func ParseJSON(bytesData []byte) *Config {

	parsed := &Config{}

	err := json.Unmarshal(bytesData, parsed)

	if err != nil {
		fmt.Println("error", err)
	}

	return parsed
}
