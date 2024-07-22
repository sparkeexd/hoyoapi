package utilities

import (
	"encoding/json"
	"fmt"
)

// Print JSON response returned from API.
// Used for development/debugging purposes.
func PrintJSON(response map[string]interface{}, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	}
}
