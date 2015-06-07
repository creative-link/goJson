/*
Package goJson implements a simple library for json parsing.
*/
package goJson

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
)

func ParseJson(b []byte) (*Root, error) {
	r := bytes.NewReader(b)
	j := new(Root)
	var Val map[string]interface{}
	err := json.NewDecoder(r).Decode(&Val)
	j.Val = Val
	return j, err
}

type Root struct {
	Val interface{}
}

func (r *Root) Get(Name string) *Root {
	currentItem := r.Val.(map[string]interface{})[Name]
	root := new(Root)
	root.Val = currentItem.(map[string]interface{})
	return root
}

func (r *Root) GetAll(Name string) ([]*Root, error) {
	currentItem := r.Val.(map[string]interface{})[Name]
	result := []*Root{}
	if reflect.ValueOf(currentItem).Kind() == reflect.Map {
            return nil, errors.New(Name+" is not array!")
        }
	items := currentItem.([]interface{})
	for _, elem := range items {
		root := new(Root)
		root.Val = elem
		result = append(result, root)
	}
	return result, nil
}

func (r *Root) String(Name string) string {
	return r.Val.(map[string]interface{})[Name].(string)
}

func (r *Root) Float(Name string) float64 {
	return r.Val.(map[string]interface{})[Name].(float64)
}

func (r *Root) Int(Name string) int {
	return int(r.Val.(map[string]interface{})[Name].(float64))
}

func (r *Root) Strings(Name string) []string {
	s := reflect.ValueOf(r.Val.(map[string]interface{})[Name])
	arr := make([]string, s.Len())
	for i := 0; i < s.Len(); i++ {
		arr[i] = s.Index(i).Interface().(string)
	}
	return arr
}
