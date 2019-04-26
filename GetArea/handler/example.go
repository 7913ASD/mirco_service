package handler

import (
	"context"

	example "sss/GetArea/proto/example"
	"sss/IhomeWeb/utils"
	"sss/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"fmt"
	"encoding/json"

	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetArea(ctx context.Context, req *example.Request, rsp *example.Response) error {
	//1.初始化返回值
	rsp.Errno=utils.RECODE_OK
	rsp.Errmsg=utils.RecodeText(rsp.Errno)

	area_list:=[]models.Area{}

	//1.4连接redis
	//{"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}
	//key: Redis collection 的名称
	//conn: Redis 连接信息
	//dbNum: 连接 Redis 时的 DB 编号. 默认是0.
	//password: 用于连接有密码的 Redis 服务器.
	//redis_config:=map[string]string{
	//	"key":utils.G_server_name,
	//	"conn":utils.G_redis_addr+":"+utils.G_redis_port,
	//	"dbNum":utils.G_redis_dbnum,
	//}
	//redis_conf,_:=json.Marshal(redis_config)
	//bm,err:=cache.NewCache("redis",string(redis_conf))
	//if err!=nil{
	//	rsp.Errno=utils.RECODE_DBERR
	//	rsp.Errmsg=utils.RecodeText(rsp.Errno)
	//	fmt.Println("redis连接失败")
	//	return nil
	//}
	bm,err:=utils.Openredis(utils.G_server_name,utils.G_redis_addr,utils.G_redis_port,utils.G_redis_dbnum)
	if err!=nil{
		rsp.Errno=utils.RECODE_DBERR
		rsp.Errmsg=utils.RecodeText(rsp.Errno)
		return nil
	}

	//1.5从redis中获取数据，发送给web

	area_json:=bm.Get("area_info")

	if area_json!=nil{
		fmt.Println("获取redis 数据")

		json.Unmarshal(area_json.([]byte),&area_list)
		for _, value := range area_list {
			temp:=example.Address{Aid:int32(value.Id),Aname:value.Name}
			rsp.Data= append(rsp.Data, &temp)
		}
		return nil
	}

	//2.查询数据

	o:=orm.NewOrm()
	qs:=o.QueryTable("area")
	num,err:=qs.All(&area_list)
	if err!=nil{
		rsp.Errno=utils.RECODE_DBERR
		rsp.Errmsg=utils.RecodeText(rsp.Errno)
		fmt.Println("数据库查询失败")
		return nil
	}
	if 0==num{
		rsp.Errno=utils.RECODE_NODATA
		rsp.Errmsg=utils.RecodeText(rsp.Errno)
		fmt.Println("数据为空 ", err)
		return nil
	}

	//2.5将查询到的数据存入缓存
	area_json,_=json.Marshal(area_list)
	err=bm.Put("area_info",area_json,time.Second*3600)
	if err!=nil{
		rsp.Errno=utils.RECODE_DBERR
		rsp.Errmsg=utils.RecodeText(rsp.Errno)
		fmt.Println("redis存入失败")
		return nil
	}

	//3.返回数据
	for _, value := range area_list {
		temp:=example.Address{Aid:int32(value.Id),Aname:value.Name}
		rsp.Data= append(rsp.Data, &temp)
	}

	return nil
}


