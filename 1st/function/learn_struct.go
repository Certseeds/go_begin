/**
 * @Github: https://github.com/Certseeds/go_begin
 * @Organization: SUSTech
 * @Author: nanoseeds
 * @Date: 2020-06-26 17:19:56
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
	. "fmt"
	"unsafe"
)

type Student struct {
	GPA        float64 // 8 byts
	student_id int32   // 4bytes
	order      uint32  // 4 bytes
	in_school  bool    // 1 bit
}

func Initial() {
	var s1 Student
	Printf("%+v\n", s1)
	var s1_pointer *Student = &s1
	Println(&s1)
	Println(s1_pointer)
	Println(*s1_pointer)
	Println(unsafe.Sizeof(s1))
	Println(unsafe.Alignof(s1))
	// 内存对齐按最大的对齐,建议从大到小排,或者说尽量使得空间利用率最高
	s2 := Student{
		4.00,
		11451419,
		19,
		false,
	}
	Println(&s2, s2.GPA, s2.in_school, s2.order, s2.student_id)
    Printf("% \n",&s2)
	s3 := struct {
		Money      float64
		Student_id int32
	}{
		Money:
		114,
		Student_id: 11451419,
	} // 匿名sturct
	Println(&s3, unsafe.Alignof(s3), unsafe.Sizeof(s3), s3)
	s4 := s3
	Println(&s4, unsafe.Alignof(s4), unsafe.Sizeof(s4), s4)
}
