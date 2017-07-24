//作者:李志强
//email:love@lizhiqiang.name
//基于Http的远程控制服务端程序
package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	_ "net/http/pprof"
	"runtime"

	"github.com/kbinani/screenshot"
	"github.com/valyala/fasthttp"
)

var cpu *int
var ip, port *string

/**
获取第一个屏幕的图像数据
*/
func CaptureScreen() *image.RGBA {
	// Capture each displays.
	//n := screenshot.NumActiveDisplays()
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	return img
}
func init() {
	ip = flag.String("ip", "0.0.0.0", "ip address to listen on")
	port = flag.String("port", "1218", "port to listen on")
	cpu = flag.Int("cpu", runtime.NumCPU(), "cpu number for httpmq")
	flag.Parse()
}

func main() {
	runtime.GOMAXPROCS(*cpu)

	m := func(ctx *fasthttp.RequestCtx) {

		model := string(ctx.FormValue("model"))
		// name := string(ctx.FormValue("name"))
		// opt := string(ctx.FormValue("opt"))
		// pos := string(ctx.FormValue("pos"))
		// num := string(ctx.FormValue("num"))
		charset := string(ctx.FormValue("charset"))
		//data = string(ctx.FormValue("data"))
		method := string(ctx.Method())
		if method == "GET" {
			//	data = string(ctx.FormValue("data"))
		} else if method == "POST" {
			if string(ctx.Request.Header.ContentType()) == "application/x-www-form-urlencoded" {
				//		data = string(ctx.FormValue("data"))
			} else {
				//		buf = ctx.PostBody()
			}
		}
		if model == "png" {
			ctx.Response.Header.Set("Content-Type", "image/png")
		}
		ctx.Response.Header.Set("Connection", "keep-alive")
		ctx.Response.Header.Set("Cache-Control", "no-cache")
		ctx.Response.Header.Set("Content-type", "image/png")

		if len(charset) > 0 {
			ctx.Response.Header.Set("Content-type", "text/plain; charset="+charset)
		}
		w := ctx.Response.BodyWriter()
		png.Encode(w, CaptureScreen())

	}
	fmt.Println("端口：" + *port)
	log.Fatal(fasthttp.ListenAndServe(*ip+":"+*port, m))

}
