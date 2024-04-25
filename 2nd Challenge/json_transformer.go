package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Read input JSON file
	inputJSON, err := os.ReadFile("input.json")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	var inputMap map[string]interface{}
	err = json.Unmarshal(inputJSON, &inputMap)
	if err != nil {
		fmt.Println("Error parsing input JSON:", err)
		return
	}

	transformedOutput := JSONtransformer(inputMap)
	// Converting the output to JSON format
	outputJSON, err := json.MarshalIndent([]interface{}{transformedOutput}, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling output JSON:", err)
		return
	}
	fmt.Println(string(outputJSON))
}
func JSONtransformer(inputMap map[string]interface{}) map[string]interface{} {
	var transformedOutput = make(map[string]interface{}, 0)
	for key, value := range inputMap {
		// Remove leading and trailing whitespace from keys
		key = strings.TrimSpace(key)
		if key != "" {
			switch valueType := value.(type) {
			case map[string]interface{}:
				if val, ok := valueType["N"].(string); ok {
					// For Number data type
					if numValue, err := strconv.ParseFloat(strings.TrimSpace(val), 64); err == nil {
						transformedOutput[key] = numValue
					}
				} else if val, ok := valueType["S"].(string); ok {
					// For String data type
					value := strings.TrimSpace(val)
					if value != "" {
						if timestamp, err := time.Parse(time.RFC3339, value); err == nil {
							transformedOutput[key] = timestamp.Unix()
						} else {
							transformedOutput[key] = value
						}
					}
				} else if val, ok := valueType["BOOL"].(string); ok {
					// For Boolean data type
					boolValue := strings.TrimSpace(strings.ToLower(val))
					if boolValue == "1" || boolValue == "t" || boolValue == "true" {
						transformedOutput[key] = true
					} else if boolValue == "0" || boolValue == "f" || boolValue == "false" {
						transformedOutput[key] = false
					}
				} else if val, ok := valueType["NULL"].(string); ok {
					// For Null data type
					nullValue := strings.TrimSpace(strings.ToLower(val))
					if nullValue == "1" || nullValue == "t" || nullValue == "true" {
						transformedOutput[key] = nil
					}
				} else if val, ok := valueType["L"].([]interface{}); ok {
					// For List data type
					if len(val) > 0 {
						transformedList := []interface{}{}
						for _, item := range val {
							itemMap, ok := item.(map[string]interface{})
							if !ok {
								continue
							}
							for itemType, itemValue := range itemMap {
								switch itemType {
								case "N":
									if numVal, err := strconv.ParseFloat(strings.TrimSpace(fmt.Sprintf("%v", itemValue)), 64); err == nil {
										transformedList = append(transformedList, numVal)
									}
								case "S":
									strVal := strings.TrimSpace(fmt.Sprintf("%v", itemValue))
									if strVal != "" {
										transformedList = append(transformedList, strVal)
									}
								case "BOOL":
									boolVal := strings.TrimSpace(strings.ToLower(fmt.Sprintf("%v", itemValue)))
									if boolVal == "1" || boolVal == "t" || boolVal == "true" {
										transformedList = append(transformedList, true)
									} else if boolVal == "0" || boolVal == "f" || boolVal == "false" {
										transformedList = append(transformedList, false)
									}
								case "NULL":
									nullVal := strings.TrimSpace(strings.ToLower(fmt.Sprintf("%v", itemValue)))
									if nullVal == "1" || nullVal == "t" || nullVal == "true" {
										transformedList = append(transformedList, nil)
									}
								}
							}
						}
						if len(transformedList) > 0 {
							transformedOutput[key] = transformedList
						}
					}
				} else if val, ok := valueType["M"].(map[string]interface{}); ok {
					// For Map data type
					transformedMap := JSONtransformer(val)
					if len(transformedMap) > 0 {
						transformedOutput[key] = transformedMap
					}
				}
			}
		}
	}
	return transformedOutput
}
