package utils

import (
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	_"github.com/garyburd/redigo/redis"
	_"github.com/gomodule/redigo/redis"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

/* 将url加上 http://IP:PROT/  前缀 */
//http:// + 127.0.0.1 + ：+ 8080 + 请求
//https://img.alicdn.com/tps/i4/TB1L7lExXzqK1RjSZFoSuvfcXXa.jpg_q90_.webp
func AddDomain2Url(url string) (domain_url string) {
	domain_url = "http://" + G_fastdfs_addr + ":" + G_fastdfs_port + "/" + url

	return domain_url
}


//后续还需要添加

func Openredis(key,addr,port,dbnum string)(bm cache.Cache,err error){
	redis_config:=map[string]string{
		"key":key,
		"conn":addr+":"+port,
		"dbNum":dbnum,
	}
	redis_conf,_:=json.Marshal(redis_config)
	bm,err=cache.NewCache("redis",string(redis_conf))
	if err!=nil{
		fmt.Println("redis连接失败")
		return nil,err
	}
	return bm,nil
}

func GetMd5string(s string)string{
	h:=md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}