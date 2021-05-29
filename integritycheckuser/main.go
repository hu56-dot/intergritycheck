package main

import (
	"fmt"
	"integritycheck/dealwith"
	"integritycheck/param"
)

func main() {

	//初始化参数
	a := param.Param{
		SdName: "sda",
		Bs:     5,
		Counts: 400,
		Skip:   22,
		Conv:   true, //conv代表块不足时用nul代替
		Buffer: 40,   //buffer代表单次响应的发送量
	}
	fmt.Println(dealwith.Contrastbool(a))

}
