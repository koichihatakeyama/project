package sql

import (
	"regexp"
	"strings"
)

type SQLParser struct {
	query  string
	params map[string]interface{}
}

type ParsedSQL struct {
	Query      string
	Parameters []interface{}
}

func NewSQLParser(query string) *SQLParser {
	return &SQLParser{
		query:  query,
		params: make(map[string]interface{}),
	}
}

func (p *SQLParser) SetParameter(name string, value interface{}) {
	p.params[name] = value
}

func (p *SQLParser) Parse() (*ParsedSQL, error) {
	result := &ParsedSQL{
		Parameters: make([]interface{}, 0),
	}

	// 2-way SQLのパラメータを解析
	re := regexp.MustCompile(`/\*\s*(.+?)\s*\*/`)
	paramCount := 1

	query := re.ReplaceAllStringFunc(p.query, func(match string) string {
		param := re.FindStringSubmatch(match)[1]
		if value, exists := p.params[param]; exists {
			result.Parameters = append(result.Parameters, value)
			return "$" + string(paramCount)
		}
		return "NULL"
	})

	result.Query = strings.TrimSpace(query)
	return result, nil
}
