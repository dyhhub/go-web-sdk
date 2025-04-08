package parser

import (
	"encoding/json"
	"fmt"
)

type JSONParser struct{}

func (p *JSONParser) Parse(input string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, fmt.Errorf("invalid JSON: %v", err)
	}
	return result, nil
}

func (p *JSONParser) Match(input string) bool {
	return json.Valid([]byte(input))

}
