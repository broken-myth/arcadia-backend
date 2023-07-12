package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

// read json file
func ReadJSON(url string) map[string][]map[string]interface{} {
	jsonfile, err := os.Open(url)
	if err != nil {
		fmt.Println(color.RedString("Error loading " + url + " file"))
	}
	defer jsonfile.Close()

	byteValue, _ := io.ReadAll(jsonfile)
	jsonInfoBytes := byteValue

	var result map[string][]map[string]interface{}

	if err2 := json.Unmarshal(jsonInfoBytes, &result); err2 != nil {
		fmt.Println(color.RedString("Error in unmarshal of JSON object"))
	}

	return result
}
