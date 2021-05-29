package ginshiyj

import (
	"integritycheck/param"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Gintarget() {

	//利用gin做API来获取severB的md5值

	r := gin.Default()

	r.GET("/:Bs/:Counts/:Skip", func(ctx *gin.Context) {
		//解析相应参数
		bs, _ := strconv.ParseUint(ctx.Param("Bs"), 10, 64)
		count, _ := strconv.ParseUint(ctx.Param("Counts"), 10, 64)
		skip, _ := strconv.ParseUint(ctx.Param("Skip"), 10, 64)
		//buffer := ctx.Param("Buffer")

		//---------------------------------------------------响应请求返回severB的md5值------------------------------
		a := param.Param{SdName: "sda", Bs: bs, Counts: count, Skip: skip, Conv: true, Buffer: 40}
		//---------------------------------------------------------------------------------------------------------

		ctx.JSON(http.StatusOK, param.ObtainMd5.Commandout(param.ObtainMd5{}, a))
	})

	log.Fatal(r.Run("192.168.19.180:8865"))

}
