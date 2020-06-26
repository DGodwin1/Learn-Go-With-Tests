package websiteracerSELECT

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url1, url2 string) (fastURL string){
	start1 := time.Now()
	http.Get(url1)
	url1Duration := time.Since(start1)
	fmt.Println(url1Duration)

	startB := time.Now()
	http.Get(url2)
	url2Duration := time.Since(startB)
	fmt.Println(url2Duration)

	if url1Duration < url2Duration{
		return url1
	}
	return url2
}