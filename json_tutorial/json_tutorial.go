package json_tutorial

//import "encoding/json"

type CONFIG struct {
	URL   string `json:"url"`
	hits   int `json:"hits"`
}

func GetJSON(path string) CONFIG {
	return CONFIG{"https://www.google.com", 100}
}
