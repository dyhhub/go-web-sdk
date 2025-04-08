package parser

import "fmt"

// Parser 定义解析器接口
type Parser interface {
	Parse(input string) (map[string]interface{}, error)
	Match(input string) bool
}

// NewParser 创建解析器工厂方法
func NewParser(parserType string) (Parser, error) {
	switch parserType {
	case "regex":
		return &RegexParser{}, nil
	case "json":
		return &JSONParser{}, nil
	default:
		return nil, fmt.Errorf("unsupported parser type: %s", parserType)
	}
}
