package webshiyj

import (
	"encoding/json"
	"fmt"
	"integritycheck/param"
	"io/ioutil"
	"log"
	"net/http"
)

func Webhttp(a param.Param) string {
	//发起web请求

	// bs := strconv.FormatUint(a.Bs, 64)
	// count := strconv.FormatUint(a.Counts, 64)
	// skip := strconv.FormatUint(a.Skip, 64)
	s := "http://192.168.19.180:8866/" + fmt.Sprint(a.Bs) + "/" + fmt.Sprint(a.Counts) + "/" + fmt.Sprint(a.Skip)
	// + "/" + count + "/" + skip

	response, err := http.Get(s)
	//如果发起失败则打印日志
	if err != nil {
		log.Fatal(err)
		fmt.Println("==================连接主机B的请求失败==============")
	}

	//最后要关闭响应
	defer response.Body.Close()

	//获取响应结果
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		fmt.Println("==================获取响应body失败==================")
	}
	//fmt.Println("成功获取响应，下面打印响应结果+++++++++++++++++++++")
	//fmt.Printf("%s\n", body)

	return string(body)
}

func Md5OutB(a param.Param) map[uint64]string {

	var md5B map[uint64]string

	md5BB := Webhttp(a)

	//定义一个结构体变量
	err := json.Unmarshal([]byte(md5BB), &md5B) //第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		fmt.Println("解析响应结果失败========================上面是解析问题")
		return nil
	}
	//fmt.Printf("---成功拿到解析结果--------------------- md5sumB=\n %+v\n", md5B)

	return md5B

}
