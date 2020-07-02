package json_tutorial

import (
	"encoding/json"
	"fmt"
)


type Config struct {
	URL string `json:"URL"`
	Hits int    `json:"Hits"`
}

func ParseJSON() *Config {

	blob := []byte(`{
		"URL": "https://www.google.com",
		"Hits": 100
	}`)

	f := &Config{}

	err := json.Unmarshal(blob, f)

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Printf("Here it is decoded: %+v", f)

	return f
}
