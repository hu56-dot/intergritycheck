package dealwith

import (
	"integritycheck/param"
	"integritycheck/webshiyj"
)

func Contrastbool(a param.Param) map[uint64]bool {

	//获取目的主机B的MD5值
	outB := webshiyj.Md5OutB(a)
	// //测试打印B的结果
	// fmt.Println("打印目的机")
	// fmt.Println(outB)

	//获取本机A的md5值
	outA := param.ObtainMd5.Commandout(param.ObtainMd5{}, a)
	// //测试打印A结果
	// fmt.Println("打印本地机")
	// fmt.Println(outA)

	result := Contrast(outA, outB)

	return result
}

func Contrast(A map[uint64]string, B map[uint64]string) map[uint64]bool {
	result := make(map[uint64]bool)
	for i := 0; i < len(A); i++ {
		if A[uint64(i)] == B[uint64(i)] {
			result[uint64(i)] = true
		} else {
			result[uint64(i)] = false
		}
	}

	return result
}
