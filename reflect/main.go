package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

const (
	sqlColumnTag string = "sql_column"
)

type ZOption func(*Z) error

func AddCustomTag(tag string) ZOption {
	return func(g *Z) error {
		g.tags = append(g.tags, tag)
		return nil
	}
}

type Z struct {
	tags []string
}

func newZ() *Z {
	return &Z{
		tags: []string{sqlColumnTag},
	}
}

func NewZ(opts ...ZOption) *Z {
	g := newZ()
	for _, opt := range opts {
		opt(g)
	}
	return g
}

func (z *Z) Decode(dst any) error {
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("error")
	}
	fmt.Println("success")
	return nil
}

func main() {
	g := NewZ(AddCustomTag("default"))
	var s S
	if err := g.Decode(&s); err != nil {
		log.Fatal(err)
	}
}

type S struct {
	F string `sql_column:"F,default:test"`
	A string `json:"a"`
}

func getTag(s S) {
	st := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		a := st.Field(i)
		fmt.Println(strings.Split(a.Tag.Get("sql_column"), ","))
	}
}

func mapIter(s S) {
	iter := reflect.ValueOf(s).MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		fmt.Println(k, v)
	}
}
