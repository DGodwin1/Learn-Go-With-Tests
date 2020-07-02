package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestJsonEncodesToStruct(t *testing.T){
	t.Run("Open and decode a JSON file", func(t *testing.T) {
		jsonFile, err := os.Open("stuff.json")

		if err != nil{
			fmt.Println(err)
		}

		defer jsonFile.Close()

		FileInBytes, _ := ioutil.ReadAll(jsonFile)

		got := ParseJSON(FileInBytes)

		want := &Config{"https://www.google.com", 100}

		if !reflect.DeepEqual(got, want){
			t.Errorf("Got %v, want %v", got, want)
		}
	})

}
