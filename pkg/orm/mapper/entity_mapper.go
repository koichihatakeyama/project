package mapper

import (
	"database/sql"
	"reflect"
	"time"
)

type EntityMapper struct {
	typeMap map[string]reflect.Type
}

func NewEntityMapper() *EntityMapper {
	return &EntityMapper{
		typeMap: make(map[string]reflect.Type),
	}
}

func (m *EntityMapper) RegisterEntity(entity interface{}) {
	t := reflect.TypeOf(entity)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	m.typeMap[t.Name()] = t
}

func (m *EntityMapper) MapToEntity(rows *sql.Rows, entityName string) (interface{}, error) {
	t, exists := m.typeMap[entityName]
	if !exists {
		return nil, ErrEntityNotRegistered
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]interface{}, len(cols))
	entity := reflect.New(t).Interface()
	entityValue := reflect.ValueOf(entity).Elem()

	for i := range values {
		field := entityValue.FieldByName(cols[i])
		if !field.IsValid() {
			continue
		}

		switch field.Type().Kind() {
		case reflect.Int64:
			var v sql.NullInt64
			values[i] = &v
		case reflect.String:
			var v sql.NullString
			values[i] = &v
		case reflect.Struct:
			if field.Type() == reflect.TypeOf(time.Time{}) {
				var v sql.NullTime
				values[i] = &v
			}
		}
	}

	if err := rows.Scan(values...); err != nil {
		return nil, err
	}

	return entity, nil
}
