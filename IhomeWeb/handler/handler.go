package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	example "github.com/micro/examples/template/srv/proto/example"
	GETAREA "sss/GetArea/proto/example"
	GETIMAGECD "sss/GetImageCd/proto/example"
	GETSMSCD "sss/GetSmscd/proto/example"
	POSTRET	"sss/PostRet/proto/example"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	"sss/IhomeWeb/models"
	"sss/IhomeWeb/utils"
	"fmt"
	"image"
	"github.com/afocus/captcha"
	"image/png"
	"regexp"
)

//func ExampleCall(w http.ResponseWriter, r *http.Request) {
//	// decode the incoming request as json
//	var request map[string]interface{}
//	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//
//	// call the backend service
//	exampleClient := example.NewExampleService("go.micro.srv.template", client.DefaultClient)
//	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
//		Name: request["name"].(string),
//	})
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//
//	// we want to augment the response
//	response := map[string]interface{}{
//		"msg": rsp.Msg,
//		"ref": time.Now().UnixNano(),
//	}
//
//	// encode and write the response as json
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//}


func ExampleCall(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {
	fmt.Println("GetIndex 	api/v1.0/house/index	请求首页轮播图 ")

	// decode the incoming request as json
	//接受殿端发送过来的内容
	var request map[string]interface{}
	//将body内容解析到map
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//创建GRPC句柄
	cli:=grpc.NewService()
	cli.Init()

	// call the backend service
	//连接服务
	exampleClient := example.NewExampleService("go.micro.srv.template", cli.Client())
	//调用方法
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	//准备返回数据的map
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}
	//设置表现为json格式
	w.Header().Set("Content-Type","application/json")
	// encode and write the response as json
	//将返回数据写入w
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//请求地域信息
func GetArea(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {
	fmt.Println("GetArea 	api/v1.0/areas	请求地域信息 ")

	//创建GRPC句柄
	cli:=grpc.NewService()
	cli.Init()

	// call the backend service
	//连接服务
	exampleClient := GETAREA.NewExampleService("go.micro.srv.GetArea", cli.Client())
	//调用方法
	rsp, err := exampleClient.GetArea(context.TODO(), &GETAREA.Request{})
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	//接受数据
	area_list:=[]models.Area{}
	for _,value:=range rsp.Data{
		temp:=models.Area{Id:int(value.Aid),Name:value.Aname}
		area_list= append(area_list, temp)
	}


	// we want to augment the response
	//准备返回数据的map
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":area_list,
	}
	//设置表现为json格式
	w.Header().Set("Content-Type","application/json")
	// encode and write the response as json
	//将返回数据写入w
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 502)
		return
	}
}


//请求首页轮播图
func GetIndex(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {
	fmt.Println("GetIndex 	api/v1.0/house/index	请求首页轮播图 ")

	// we want to augment the response
	//准备返回数据的map
	response := map[string]interface{}{
		"errno": utils.RECODE_OK,
		"errmsg": utils.RecodeText(utils.RECODE_OK),
	}
	//设置表现为json格式
	w.Header().Set("Content-Type","application/json")
	// encode and write the response as json
	//将返回数据写入w
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//请求session
func GetSession(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {
	fmt.Println("GetSession 	api/v1.0/session	请求session ")


	// we want to augment the response
	//准备返回数据的map
	response := map[string]interface{}{
		"errno": utils.RECODE_SESSIONERR,
		"errnum": utils.RecodeText(utils.RECODE_SESSIONERR),
	}
	//设置表现为json格式
	w.Header().Set("Content-Type","application/json")
	// encode and write the response as json
	//将返回数据写入w
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}


//请求验证码图片
func GetImageCd(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	fmt.Println("GetImageCd 	api/v1.0/imagecode/:uuid	请求验证码图片 ")


	//创建GRPC句柄
	cli:=grpc.NewService()
	cli.Init()

	// call the backend service
	//连接服务
	exampleClient := GETIMAGECD.NewExampleService("go.micro.srv.GetImageCd", cli.Client())
	//调用方法
	rsp, err := exampleClient.GetImageCd(context.TODO(), &GETIMAGECD.Request{
		Uuid:ps.ByName("uuid"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if rsp.Errno!="0"{
		// we want to augment the response
		//准备返回数据的map
		response := map[string]interface{}{
			"errno": rsp.Errno,
			"errmsg": rsp.Errmsg,
		}
		//设置表现为json格式
		w.Header().Set("Content-Type","application/json")
		// encode and write the response as json
		//将返回数据写入w
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	//拼接数据
	var img image.RGBA

	for _,value:=range rsp.Pix{
		img.Pix= append(img.Pix, uint8(value))
	}
	img.Stride=int(rsp.Stride)

	img.Rect.Max.X=int(rsp.Max.X)
	img.Rect.Min.X=int(rsp.Min.X)
	img.Rect.Max.Y=int(rsp.Max.Y)
	img.Rect.Min.Y=int(rsp.Min.Y)

	var image captcha.Image

	image.RGBA=&img
	//将图片数据发送给前端
	png.Encode(w,image)

}

//请求短信验证码服务
func GetSmscd(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	fmt.Println("GetSmscd 	GET	api/v1.0/smscode/:mobile	获取短信验证码服务 ")

	//获取手机号
	mobile:=ps.ByName("mobile")
	//验证手机号
	reg:=regexp.MustCompile(`0?(13|14|15|17|18|19)[0-9]{9}`)
	bo:=reg.MatchString(mobile)
	if bo==false{
		// we want to augment the response
		//准备返回数据的map
		response := map[string]interface{}{
			"errno": utils.RECODE_MOBILEERR,
			"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		}
		//设置表现为json格式
		w.Header().Set("Content-Type","application/json")
		// encode and write the response as json
		//将返回数据写入w
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//获取图片验证码，uuid\
	fmt.Println(r.URL.Query())
	text:=r.URL.Query()["text"][0]
	id:=r.URL.Query()["id"][0]

	//创建GRPC句柄
	cli:=grpc.NewService()
	cli.Init()

	// call the backend service
	//连接服务
	exampleClient := GETSMSCD.NewExampleService("go.micro.srv.GetSmscd", cli.Client())
	//调用方法
	rsp, err := exampleClient.GetSmscd(context.TODO(), &GETSMSCD.Request{
		Mobile:mobile,
		Text:text,
		Uuid:id	,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	//准备返回数据的map
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	//设置表现为json格式
	w.Header().Set("Content-Type","application/json")
	// encode and write the response as json
	//将返回数据写入w
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//发送注册信息服务
func PostRet(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {
	fmt.Println("PostRet 	/api/v1.0/users	发送注册信息服务 ")

	// decode the incoming request as json
	//接受殿端发送过来的内容
	var request map[string]interface{}
	//将body内容解析到map
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}


	//mobile: "123",
	//	password: "123",
	//		sms_code: "123
	//校验数据
	if request["mobile"].(string)==""||request["password"].(string)==""||request["sms_code"].(string)==""{
		response := map[string]interface{}{
			"errno":utils.RECODE_NODATA,
			"errmsg":utils.RecodeText(utils.RECODE_NODATA),
		}
		//设置表现为json格式
		w.Header().Set("Content-Type","application/json")
		// encode and write the response as json
		//将返回数据写入w
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}



	//创建GRPC句柄
	cli:=grpc.NewService()
	cli.Init()

	// call the backend service
	//连接服务
	exampleClient := POSTRET.NewExampleService("go.micro.srv.PostRet", cli.Client())
	//调用方法
	rsp, err := exampleClient.PostRet(context.TODO(), &POSTRET.Request{
		Mobile:request["mobile"].(string),
		Password:request["password"].(string),
		SmsCode:request["sms_code"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}


	// we want to augment the response
	//准备返回数据的map
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	//将sessionid设置到cookie中
	//ihomeloginname
	//读取cookie
	cookie,err:=r.Cookie("ihomeloginname")
	if cookie==nil||err!=nil{
		cookie:=http.Cookie{Name:"ihomeloginname",Value:rsp.Sessionid,Path:"/",MaxAge:600}

		http.SetCookie(w,&cookie)
	}


	//设置表现为json格式
	w.Header().Set("Content-Type","application/json")
	// encode and write the response as json
	//将返回数据写入w
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}