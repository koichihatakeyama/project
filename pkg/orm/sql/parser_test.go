package sql

import (
	"reflect"
	"testing"
)

func TestSQLParser_Parse(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		params   map[string]interface{}
		expected *ParsedSQL
	}{
		{
			name:  "Simple query with one parameter",
			query: "SELECT * FROM users WHERE id = /* userId */1",
			params: map[string]interface{}{
				"userId": 5,
			},
			expected: &ParsedSQL{
				Query:      "SELECT * FROM users WHERE id = $1",
				Parameters: []interface{}{5},
			},
		},
		{
			name:  "Query with multiple parameters",
			query: "SELECT * FROM users WHERE name = /* userName */'test' AND age > /* userAge */20",
			params: map[string]interface{}{
				"userName": "John",
				"userAge":  30,
			},
			expected: &ParsedSQL{
				Query:      "SELECT * FROM users WHERE name = $1 AND age > $2",
				Parameters: []interface{}{"John", 30},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := NewSQLParser(tt.query)
			for k, v := range tt.params {
				parser.SetParameter(k, v)
			}

			result, err := parser.Parse()
			if err != nil {
				t.Errorf("Parse() error = %v", err)
				return
			}

			if result.Query != tt.expected.Query {
				t.Errorf("Query = %v, want %v", result.Query, tt.expected.Query)
			}

			if !reflect.DeepEqual(result.Parameters, tt.expected.Parameters) {
				t.Errorf("Parameters = %v, want %v", result.Parameters, tt.expected.Parameters)
			}
		})
	}
}
