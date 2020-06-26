/**
 * @Github: https://github.com/Certseeds/go_begin
 * @Organization: SUSTech
 * @Author: nanoseeds
 * @Date: 2020-06-26 19:22:39
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

func Test_array() {
	var strs [5]string
	strs[0] = "Dark angels"
	strs[1] = "Blood angles"
	strs[2] = "Ultramarines"
	strs[3] = "Imperial Fists"
	strs[4] = "White Scars"
	for x := 0; x < 5; x++ {
		fmt.Println(strs[x])
	}
	for index, value := range strs {
		fmt.Println(index, strs[index], value)
	}
	numbers := [4]int{114, 514, 1919, 810}
	for x := 0; x < len(numbers); x++ {
		fmt.Println(numbers[x])
	}
	fmt.Println(numbers)
	for index, value := range strs {
		fmt.Println(value, &strs[index], &value)
	}
    fmt.Println(reflect.TypeOf(numbers),reflect.TypeOf(strs))
}
