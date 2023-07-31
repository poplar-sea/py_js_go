package main

import (
	"fmt"
)

type Author struct {
	Name         int      `json:"Name"`
	Publications []string `json:"Publication",omitempty`
}

func main() {
	// t := reflect.TypeOf(Author{})
	// for i := 0; i < t.NumField(); i++ {
	// 	name := t.Field(i).Name
	// 	// Type := t.Field(i).Type
	// 	s, _ := t.FieldByName(name)
	// 	fmt.Println(name, s.Tag.Get("json"))
	// 	// fmt.Println(Type)
	// }
	var author Author
	fmt.Printf("%v\n", author)
	fmt.Printf("%+v\n", author)
	fmt.Printf("%#v\n", author)
}
