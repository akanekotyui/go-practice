//関数の定義方法　func
//fanc calc (変数リスト) 返り値の型
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("calc(2, 3) = %d\n", calc(2, 3))
  //quoとremainのふたつの変数を用意し、それぞれに代入
	quo, remain := calcQuoAndRemain(10, 3)
	fmt.Printf("quo = %d, remain = %d\n", quo, remain)

	// 関数を変数に代入して利用することができる（オブジェクトとして使う）
  　　//ローカルの関数なので関数名をいれずに使える
	f1 := func(x int) int {
		return 5 * x
	}
	fmt.Printf("f1(1) = %d\n", f1(1))
	fmt.Printf("f1(3) = %d\n", f1(3))
}

func calc(a int, b int) int {
	return a + b
}
//値の組みを返り値としてとれる
func calcQuoAndRemain(a, b int) (int, int) {
	return a / b, a % b
}
