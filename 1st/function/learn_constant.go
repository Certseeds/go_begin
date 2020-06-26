/**
 * @Github: https://github.com/Certseeds/go_begin
 * @Organization: SUSTech
 * @Author: nanoseeds
 * @Date: 2020-06-26 18:40:35
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

package function

import (
	"fmt"
	"reflect"
)

const dollar = 6.907
const Tokyo = "Japan_Tykio"

const lubu float64 = 0.11
const Segu string = "Japan_T"

func Test_constant() {
	fmt.Print(reflect.TypeOf(dollar))
	fmt.Print(reflect.TypeOf(Tokyo))
	fmt.Print(reflect.TypeOf(lubu))
	fmt.Print(reflect.TypeOf(Segu))
	var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)

	fmt.Println(answer)

	// Constant third will be of kind floating point.
	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

	fmt.Println(third)

	// Constant zero will be of kind integer.
	const zero = 1 / 3 // KindInt(1) / KindInt(3)

	fmt.Println(zero)

	// This is an example of constant arithmetic between typed and
	// untyped constants. Must have like types to perform math.
	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)

	fmt.Println(one)
	fmt.Println(two)

	// Max integer value on 64 bit architecture.
	const maxInt = 9223372036854775807

	fmt.Println(maxInt)
	// fmt.Println(maxInt + 1 ) // overflow
	// fmt.Println(-maxInt-2) // overflow in another way
	const bigger = 9223372036854775808543522345 // only this line can run
	// fmt.Println(reflect.TypeOf(bigger)) // after add this line it can not run

}
