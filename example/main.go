package main

import (
	"fmt"

	"github.com/SolarisNeko/sugar/generic"
)

// 示例代码
func main() {
	stream := Stream{1, 2, 3, 4, 5}

	// 使用 Filter 过滤偶数
	predicate := func(num T) bool {
		if number, ok := num.(int); ok {
			return number%2 == 0
		}
		return false
	}
	// 使用 Map 将每个元素加倍
	mapper := func(n T) T {
		toNumber := generic.ToNumber(n)
		number := generic.ToNumber(2)
		return toNumber * number
	}

	reducer := func(n1 T, n2 T) T {
		return generic.ToNumber(n1) + generic.ToNumber(n2)
	}

	sum := stream.
		Filter(predicate).
		Map(mapper).
		Reduce(0, reducer)

	// 输出: 12
	fmt.Println(sum)
}
