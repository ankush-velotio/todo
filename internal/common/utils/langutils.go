package utils

import (
	"encoding/json"
)

func Contains(slice []string, element string) bool {
	for _, val := range slice {
		if element == val {
			return true
		}
	}
	return false
}

// GetCustomJSON takes an object and the names of fields to ignore from the JSON representation
// of the struct
func GetCustomJSON(obj interface{}, ignoreFields ...string) (interface{}, error) {
	toJson, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	if len(ignoreFields) == 0 {
		return toJson, nil
	}

	toMap := map[string]interface{}{}
	err = json.Unmarshal(toJson, &toMap)
	if err != nil {
		return nil, err
	}

	for _, field := range ignoreFields {
		delete(toMap, field)
	}

	toJson, err = json.Marshal(toMap)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(toJson, &obj)
	return obj, nil
}
