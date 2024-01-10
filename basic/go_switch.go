package basic

import (
	"fmt"
)

func GoSwitch() {

	//支持多条件匹配
	//不同的case之间不使用break分割 默认只执行一个case
	//如果想要执行多个case 需要使用 fallthrough关键字 也可以使用break终止

	a := 100
	switch a {
	case 100:
		fmt.Println("100")
		// break 阻断
		fallthrough

	case 200:
		// break 阻断
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	default:
		fmt.Println("other")
	}

}
