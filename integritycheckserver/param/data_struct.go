package param

import (
	"fmt"
	"integritycheck/systeminputcommand"
)

// Param This is the Params about the command "dd if=/dev/zero of=/dev/block bs=512 seek=1 count=1 conv=sync | md5sum"
type Param struct {
	SdName string //需要解析的磁盘名字  eg:"sda"
	Bs     uint64 //解析块的大小block size eg: 1M  每次解析1M
	Counts uint64 //解析块的数量 eg: 50
	Skip   uint64 //跳跃块的数量，注意索引是从0开始的
	Cbs    uint64 //一次转换的字节数，指定缓冲区的大小
	Conv   bool   //如果某个块的大小不够，则补充为空
	Buffer int
}

type ObtainMd5 struct {
	Param
	//AllMd5sum map[uint32]string
}

type ShellCommand interface {
	commandout(param Param) string
}

//this function will return the md5sum of the single block
func (O ObtainMd5) Commandout(param Param) map[uint64]string {
	//dd if=/dev/sda count=1  skip=k bs=1M | md5sum
	buffers := make(map[uint64]string)
	for k := param.Skip; k <= param.Counts+param.Skip; k++ {
		s := "dd if=/dev/sda count=1 " + "skip=" + fmt.Sprint(k) + " bs=" + fmt.Sprint(param.Bs) + "K | md5sum "
		//fmt.Println(s)
		buffers[k-param.Skip] = systeminputcommand.InputCommand(s)
	}
	//fmt.Println(buffers)
	return buffers
}
