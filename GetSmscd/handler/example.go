package handler

import (
	"context"

	example "sss/GetSmscd/proto/example"
	"fmt"
	"sss/IhomeWeb/utils"
	"github.com/astaxie/beego/orm"
	"sss/IhomeWeb/models"
	"reflect"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"time"
	"github.com/SubmailDemo/submail"
	"strconv"
	"strings"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetSmscd(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("GetSmscd 	GET	api/v1.0/smscode/:mobile	获取短信验证码服务 ")

	//1.初始化返回值
	rsp.Errno=utils.RECODE_OK
	rsp.Errmsg=utils.RecodeText(rsp.Errno)
	//2.手机号验证，带入数据库查询
	o:=orm.NewOrm()
	user:=models.User{Mobile:req.Mobile}
	err:=o.Read(&user)
	if err==nil{
		fmt.Println("用户已经注册过了")
		rsp.Errno=utils.RECODE_USERONERR
		rsp.Errmsg=utils.RecodeText(rsp.Errno)
		return nil
	}

	//3.连接redis
	bm,err:=utils.Openredis(utils.G_server_name,utils.G_redis_addr,utils.G_redis_port,utils.G_redis_dbnum)
	if err!=nil	{
		rsp.Errno=utils.RECODE_DBERR
		rsp.Errmsg=utils.RecodeText(rsp.Errno)
		return nil
	}

	//4.获取之前缓存的图片验证码的值
	value:=bm.Get(req.Uuid)
	fmt.Println(reflect.TypeOf(value),value)

	redis_string,_:=redis.String(value,nil)
	fmt.Println(reflect.TypeOf(redis_string),redis_string)

	//5.数据对比，验证值是否正确
	if req.Text!=redis_string{
		fmt.Println("图片验证码不正确")
		rsp.Errno=utils.RECODE_DBERR
		rsp.Errmsg=utils.RecodeText(rsp.Errno)
		return nil
	}

	//6.生成随机数
	t:=rand.New(rand.NewSource(time.Now().UnixNano()))
	//四位数的取之范围：1000-9999
	//rand.Intn(9999) 0-9999
	size:=t.Intn(8999)+1000
	fmt.Println("随机数验证码：",size)
	//7.发送短信

	//发送短信的配置信息
	messageconfig := make(map[string]string)
	//预先创建好的appid
	messageconfig["appid"] = "29672"
	//预先获得的app的key
	messageconfig["appkey"] = "89d90165cbea8cae80137d7584179bdb"
	//加密方式默认
	messageconfig["signtype"] = "md5"



	//messagexsend
	//创建短信发送的句柄
	messagexsend := submail.CreateMessageXSend()
	//短信发送的手机号
	submail.MessageXSendAddTo(messagexsend, req.Mobile)
	//短信发送的模板
	submail.MessageXSendSetProject(messagexsend, "NQ1J94")
	//验证码
	submail.MessageXSendAddVar(messagexsend, "code", strconv.Itoa(size))
	//发送短信的请求
	send:=submail.MessageXSendRun(submail.MessageXSendBuildRequest(messagexsend), messageconfig)
	fmt.Println("MessageXSend ", send)


	//8.验证是否成功
	bo:=strings.Contains(send,"success")
	if bo!=true{
		fmt.Println("短信发送失败")
		rsp.Errno=utils.RECODE_SMSERR
		rsp.Errmsg=utils.RecodeText(rsp.Errno)
		return nil
	}


	//9.将随机数存入缓存，key手机号，value随机数
	err=bm.Put(req.Mobile,strconv.Itoa(size),time.Second*300)
	if err!=nil{
		fmt.Println("缓存失败")
		rsp.Errno=utils.RECODE_DBERR
		rsp.Errmsg=utils.RecodeText(rsp.Errno)
		return nil
	}

	return nil
}

