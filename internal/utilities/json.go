package utilities

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sparkeexd/hoyoapi/internal/errors"
)

// Return JSON unmarshalled response with the given model.
func UnmarshalJSON[V any](res []byte, model V) error {
	var err error

	if err = json.Unmarshal(res, &model); err != nil {
		return errors.NewError(
			errors.JSON_DESERIALIZATION_ERROR,
			fmt.Sprintf("Error: %s", err.Error()),
		)
	}

	return err
}

// Print JSON response returned from API.
// Used for development/debugging purposes.
func PrintJSON(response any, err error) {
	jsonString, _ := json.Marshal(response)

	data := make(map[string]interface{})
	json.Unmarshal(jsonString, &data)

	if err != nil {
		log.Println(err)
	} else {
		data, _ := json.MarshalIndent(data, "", "    ")
		log.Println(string(data))
	}
}
