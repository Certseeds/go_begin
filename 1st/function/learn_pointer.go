/**
 * @Github: https://github.com/Certseeds/go_begin
 * @Organization: SUSTech
 * @Author: nanoseeds
 * @Date: 2020-06-26 18:29:13
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
import "fmt"
type Ns struct {
	Price float64
	loca  string
}

func Set_Ns(NS Ns, price float64, loca string) {
	NS.Price = price
	NS.loca = loca
}

func Set_Ns_p(NS *Ns, price float64, loca string) {
	(*NS).Price = price
	(*NS).loca = loca
}

func Test_pointer() {
	game_play := Ns{300,"Guangzhou"}
	fmt.Println(game_play)
	Set_Ns(game_play,400,"Tokyo")
	fmt.Println(game_play)
	Set_Ns_p(&game_play,200,"Beijing")
	fmt.Println(game_play)
	// 指针只是用来区别传值还是传引用的方式?
	// 而且也没有delete,由GC.
}
