/**
 * @Github: https://github.com/Certseeds/go_begin
 * @Organization: SUSTech
 * @Author: nanoseeds
 * @Date: 2020-06-26 10:51:00
 * @LastEditors  : nanoseeds
 */
/*  go_begin
    Copyright (C) 2020  nanoseeds

    go_begin is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as
    published by the Free Software Foundation, either version 3 of the
    License, or (at your option) any later version.

    go_begin is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"./function"
	. "fmt"
	"math"
)

const MATH_PI = 3.1415926
const Math_E = 2.718281828
const (
	i = iota
	wonna
	be
	a
	_
	belt
	_
)

/* 1st go program for nanoseeds */
func main() {
	Println("Hello, World!")
	Println("google " + "Inc")
	var age = 12
	var fruit = 114514
	Println(fruit + age)
	// var king = "nameless"
	type newType int
	type stueent struct{}
	type can_dir interface{}
	Println(math.E)
	Println(math.Sqrt(114514))
	Println(math.Pi)
	// Print(get_ppid(), Get_ppid())
	vars()
	Print(MATH_PI, Math_E, MATH_PI*Math_E)
	bitt_operators_int8()
	bitt_operators_int16()
	bitt_operators_int32()
	bitt_operators_int64()
	switch_command("死亡搁浅")
	swtich_type(114, 51.4, "Debugging", complex(1, 2), nil, true)
	死亡搁浅 := int32(114)
	var 美国末日2 string = " US !"
	Println(死亡搁浅, 美国末日2)
	// Print(get_ppid()) is protect
	Println(function.Get_ppid()) // it is public
	function.Initial()
	function.Test_pointer()
	function.Test_constant()
    function.Test_array()
}
func vars() {
	var judge bool = true
	var number1 int64 = 0x3f3f3f3f3f3f
	var number2 int32 = 0x3f3f3f3f
	var number3 uint16 = 0x3f3f
	var number4 uint8 = 0x3f
	var number5 float32 = 0.114511
	var number6 float64 = 0.51419191919
	var number7 complex64 = 1 + 2i
	var number8 complex128 = 114514 + 0x3f3f3f3f3fi
	var str1 = "0x3f3f3f3f"
	Println(str1)
	Println(number1, number2, number3, number4, number5, number6, number7, number8, judge)
	Println(&number1, &number2, &number3, &number4, &number5, &number6, &number7, &number8, &judge)
	_, number9, _ := block_return()
	Println(number9)
}

func block_return() (int32, int64, float64) {
	return 114, 514, 19.18
}

// bit operators
func bitt_operators_int8() {
	var transfer int8 = 1
	for i := 0; i < 9; i++ {
		Println(transfer << i)
	}
}

func bitt_operators_int16() {
	var transfer int16 = 1
	for i := 0; i < 17; i++ {
		Println(transfer << i)
	}
}

func bitt_operators_int32() {
	var transfer int32 = 1
	for i := 0; i < 33; i++ {
		Println(transfer << i)
	}
}

func bitt_operators_int64() {
	var transfer int64 = 1
	for i := 0; i < 65; i++ {
		Println(transfer << i)
	}
}

// if ifelse is ignored

func switch_command(v string) {
	switch v {
	case "The_last_of_us_2":
		{
			Println("奇迹!")
		}
	case "The_Witcher_3":
		{
			Println("还是奇迹!")
		}
	default:
		{
			Println("? 不是奇迹")
		}
	}
}

func swtich_type(items ...interface{}) {
	// 可变参数?
	for index, value := range items {
		switch value.(type) {
		case bool:
			{
				Printf("Param #%d is a bool\n", index)
			}
		case float64:
			{
				Printf("Param #%d is a float64\n", index)

			}
		case int, int64:
			{
				Printf("Param #%d is a int\n", index)
			}
		case nil:
			{
				Printf("Param #%d is a nil\n", index)

			}
		case string:
			{
				Printf("Param #%d is a string\n", index)

			}
		default:
			{
				Printf("Param #%d is unknown\n", index)
			}
		}
	}
}
