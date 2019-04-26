package handler

import (
	"context"

	example "sss/PostRet/proto/example"
	"fmt"
	"sss/IhomeWeb/utils"
	"github.com/garyburd/redigo/redis"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostRet(ctx context.Context, req *example.Request, rsp *example.Response) error {

	fmt.Println("PostRet 	/api/v1.0/users	发送注册信息服务 ")
	//1.初始化返回值
	rsp.Errno=utils.RECODE_OK
	rsp.Errmsg=utils.RecodeText(rsp.Errno)
	//2.连接redis
	bm,err:=utils.Openredis(utils.G_server_name,utils.G_redis_addr,utils.G_redis_port,utils.G_redis_dbnum)
	if err!=nil{

		rsp.Errno=utils.RECODE_DBERR
		rsp.Errmsg=utils.RecodeText(rsp.Errno)
		fmt.Println("redis连接失败")
		return nil
	}


	//3.查询短信验证码
	valus:=bm.Get(req.Mobile)
	value_string,_:=redis.String(valus,nil)
	//4.判断短信验证码是否正确
	if value_string!=req.SmsCode{
		rsp.Errno=utils.RECODE_DBERR
		rsp.Errmsg=utils.RecodeText(rsp.Errno)
		fmt.Println("短信验证码不正确")
		return nil
	}
	//5.密码加密

	//6.数据存入数据库

	//7.生成sessionid


	//8.通过sessionid拼接key，存入数据库


	return nil
}

