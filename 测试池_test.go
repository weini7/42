package main

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

type 结构体 struct {
	Filed          string
	CamelUnderline string
}

func 能不能拿出来啊行不行(i interface{}) {
	fmt.Println(" \033[46;31m这个反人类颜色看得清？？为什么不闪！！ \033[5m")
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Struct {
		return
	}
	v := reflect.ValueOf(i)
	structLen := t.NumField()
	for i := 0; i < structLen; i++ {
		field := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", field.Name, field.Type, val)
	}
	log.Println(reflect.TypeOf(i).Field(0).Name)
	log.Println(reflect.TypeOf(&i).Elem())
}

func Test宇宙的答案是什么(t *testing.T) {
	t.Log("42")
	fmt.Println("//")
	fmt.Println("//")

	//
	{
		集美 := 结构体{
			Filed:          "111",
			CamelUnderline: "222",
		}

		能不能拿出来啊行不行(集美)
		fmt.Println("//")
		fmt.Println("//")

		log.Println(reflect.TypeOf(集美).Field(0).Tag)
		log.Println(reflect.TypeOf(集美).Field(0).Index)
		log.Println(reflect.TypeOf(集美).Field(1).Index)

		//log.Println(reflect.ValueOf(集美))
		//log.Println(reflect.ValueOf(集美).Type())
		//log.Println(reflect.ValueOf(集美).Type().String())
		//log.Println(reflect.ValueOf(集美).Type().Name())
		//
		//log.Println("--------------------------------------")
		//
		//log.Println(reflect.ValueOf(集美.Filed))
		//log.Println(reflect.ValueOf(集美.Filed).Type())
		//log.Println(reflect.ValueOf(集美.Filed).Type().String())
		//log.Println(reflect.ValueOf(集美.Filed).Type().Name())
		//
		//log.Println("--------------------------------------")
		//
		//log.Println(utils.Camel2Underline("AbCdEfGh"))

	}

	//

	//

	//

	//

	//

	//

	//

	//

	//

	//

	//

	//

	//

	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")

}
