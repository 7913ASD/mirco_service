package handler

import (
	"context"

	example "sss/GetImageCd/proto/example"
	"sss/IhomeWeb/utils"
	"github.com/afocus/captcha"
	"image/color"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetImageCd(ctx context.Context, req *example.Request, rsp *example.Response) error {
	//1.初始化返回值
	rsp.Errno=utils.RECODE_OK
	rsp.Errmsg=utils.RecodeText(rsp.Errno)
	//2.生成图片和随机数
	cap := captcha.New()

	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}

	cap.SetSize(90, 41)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})



	img, str := cap.Create(4, captcha.NUM)

	println("图片验证码的值",str)

	//3.返回图片
	//解引用
	a:=*img
	b:=*(a.RGBA)
	for _,value:=range b.Pix{
		rsp.Pix= append(rsp.Pix, uint32(value))
	}
	rsp.Stride=int64(b.Stride)
	rsp.Max=&example.Point{X:int64(b.Rect.Max.X),Y:int64(b.Rect.Max.Y)}
	rsp.Min=&example.Point{X:int64(b.Rect.Min.X),Y:int64(b.Rect.Min.Y)}

	//4.连接redis
	bm,err:=utils.Openredis(utils.G_server_name,utils.G_redis_addr,utils.G_redis_port,utils.G_redis_dbnum)
	if err!=nil{
		rsp.Errno=utils.RECODE_DBERR
		rsp.Errmsg=utils.RecodeText(rsp.Errno)
		return err
	}
	//5.把uuid和随机数存入redis
	bm.Put(req.Uuid,str,300*time.Second)

	return nil
}

