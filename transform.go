// Copyright 2021 taichi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kra

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

type DefaultTransformer struct {
	core *Core
}

func NewDefaultTransformer(core *Core) *DefaultTransformer {
	return &DefaultTransformer{core}
}

var ErrNoPointer = errors.New("kra: destination is not a pointer")
var ErrNilPointer = errors.New("kra: destination is typed nil pointer")
var ErrNoColumns = errors.New("kra: no scannable columns")
var ErrInvalidMapKeyType = errors.New("kra: invalid map key type")
var ErrNoRecord = errors.New("kra: no record")
var ErrNoSlice = errors.New("kra: destination is not a slice")

func validate(src Rows, value reflect.Value) error {
	if src == nil {
		return sql.ErrNoRows
	}
	if src.Err() != nil {
		return src.Err()
	}

	if value.Kind() != reflect.Ptr {
		return ErrNoPointer
	}

	if value.IsNil() {
		return ErrNilPointer
	}

	return nil
}

func seekColumns(src Rows) ([]string, error) {
	columns, err := src.Columns()

	if err != nil {
		return nil, err
	}

	if len(columns) < 1 {
		return nil, ErrNoColumns
	}

	return columns, nil
}

func (transformer *DefaultTransformer) Transform(src Rows, dest interface{}) error {
	value := reflect.ValueOf(dest)

	if err := validate(src, value); err != nil {
		return err
	}

	defer src.Close()

	if src.Next() == false {
		return ErrNoRecord
	}

	directValue := reflect.Indirect(value)

	if isScanner(directValue.Type()) {
		return src.Scan(dest)
	}

	columns, err := seekColumns(src)

	if err != nil {
		return err
	}

	if directValue.Kind() == reflect.Map {
		mapType := directValue.Type()
		if mapType.Key().Kind() == reflect.String {
			return ScanMap(src, columns, directValue)
		} else {
			return fmt.Errorf("KeyType:%v %w", mapType.Key(), ErrInvalidMapKeyType)
		}
	}

	if directValue.Kind() == reflect.Struct {
		return transformer.ScanStruct(src, columns, value)
	}

	return src.Scan(dest)
}

var typeOfScanner = reflect.TypeOf((*sql.Scanner)(nil)).Elem()

func isScanner(elementType reflect.Type) bool {
	return reflect.PtrTo(elementType).Implements(typeOfScanner)
}

func ScanMap(src Rows, columns []string, value reflect.Value) error {

	if value.IsNil() {
		value.Set(reflect.MakeMap(value.Type()))
	}

	length := len(columns)
	receivers := make([]interface{}, length)
	for i := 0; i < length; i++ {
		receivers[i] = new(interface{})
	}

	if err := src.Scan(receivers...); err != nil {
		return err
	}

	for i, col := range columns {
		val := *(receivers[i].(*interface{}))
		value.SetMapIndex(reflect.ValueOf(col), reflect.ValueOf(val))
	}

	return src.Err()
}

func (transformer *DefaultTransformer) ScanStruct(src Rows, columns []string, root reflect.Value) error {
	structDef, err := transformer.core.Repository.Lookup(root.Type())

	if err != nil {
		return err
	}

	receivers := make([]interface{}, len(columns))
	for i, col := range columns {
		def, value, err := structDef.ByName(root, col)
		if err != nil {
			return err
		}
		var recv interface{}
		if def.Options["name"] == "-" {
			recv = new(interface{})
		} else {
			recv = value.Addr().Interface()
		}
		receivers[i] = recv
	}

	if err := src.Scan(receivers...); err != nil {
		return err
	}

	return src.Err()
}

func (transformer *DefaultTransformer) TransformAll(src Rows, dest interface{}) error {
	value := reflect.ValueOf(dest)

	if err := validate(src, value); err != nil {
		return err
	}

	directValue := reflect.Indirect(value)
	if directValue.Kind() != reflect.Slice {
		return fmt.Errorf("type:%v %w", value.Type(), ErrNoSlice)
	}

	defer src.Close()

	elementType := directValue.Type().Elem()
	appender := selectAppender(directValue, elementType)
	directElementType := Indirect(elementType)

	if isScanner(directElementType) {
		return ScanAll(src, directElementType, appender)
	}

	columns, err := seekColumns(src)

	if err != nil {
		return err
	}

	if directElementType.Kind() == reflect.Map {
		mapType := directElementType
		if mapType.Key().Kind() == reflect.String {
			return ScanAllMap(src, columns, mapType, appender)
		} else {
			return fmt.Errorf("KeyType:%v %w", mapType.Key(), ErrInvalidMapKeyType)
		}
	}

	if directElementType.Kind() == reflect.Struct {
		return transformer.ScanAllStruct(src, columns, directElementType, appender)
	}

	return ScanAll(src, directElementType, appender)
}

func selectAppender(directValue reflect.Value, elementType reflect.Type) func(reflect.Value) {
	if elementType.Kind() == reflect.Ptr {
		return func(newValue reflect.Value) {
			appendPointer(directValue, newValue)
		}
	} else {
		return func(newValue reflect.Value) {
			appendDirect(directValue, newValue)
		}
	}
}

func appendPointer(container reflect.Value, newValue reflect.Value) {
	container.Set(reflect.Append(container, newValue))
}

func appendDirect(container reflect.Value, newValue reflect.Value) {
	container.Set(reflect.Append(container, reflect.Indirect(newValue)))
}

func ScanAll(src Rows, elementType reflect.Type, appender func(reflect.Value)) error {
	for src.Next() {
		newValue := reflect.New(elementType)
		if err := src.Scan(newValue.Interface()); err != nil {
			return err
		}
		appender(newValue)
	}

	return src.Err()
}

func ScanAllMap(src Rows, columns []string, mapType reflect.Type, appender func(reflect.Value)) error {
	for src.Next() {
		newValue := reflect.MakeMap(mapType)
		if err := ScanMap(src, columns, newValue); err != nil {
			return err
		}
		appender(newValue)
	}

	return src.Err()
}

func (transformer *DefaultTransformer) ScanAllStruct(src Rows, columns []string, elementType reflect.Type, appender func(reflect.Value)) error {
	for src.Next() {
		newValue := reflect.New(elementType)
		if err := transformer.ScanStruct(src, columns, newValue); err != nil {
			return err
		}
		appender(newValue)
	}

	return src.Err()
}
