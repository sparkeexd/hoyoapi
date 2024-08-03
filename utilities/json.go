package utilities

import (
	"encoding/json"
	"log"
)

// Print JSON response returned from API.
// Used for development/debugging purposes.
func PrintJSON(response map[string]interface{}, err error) {
	if err != nil {
		log.Println(err)
	} else {
		data, _ := json.MarshalIndent(response, "", "    ")
		log.Println(string(data))
	}
}
