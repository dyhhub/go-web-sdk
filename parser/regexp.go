package parser

import (
	"fmt"
	"regexp"
)

type RegexParser struct {
	patterns map[string]string
}

func (p *RegexParser) Parse(input string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for name, pattern := range p.patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(input)

		if len(matches) > 1 {
			result[name] = matches[1]
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no matches found")
	}

	return result, nil
}

func (p *RegexParser) Match(input string) bool {
	for _, pattern := range p.patterns {
		matched, _ := regexp.MatchString(pattern, input)
		if matched {
			return true
		}
	}
	return false
}

// WithPatterns 设置正则模式
func (p *RegexParser) WithPatterns(patterns map[string]string) *RegexParser {
	p.patterns = patterns
	return p
}
