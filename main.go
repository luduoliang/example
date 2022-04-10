package main

import (
	"example/common"
	"example/config"
	"example/controller/rpc/grpctest"
	"example/controller/rpc/rpctest"
	"example/router"
	"fmt"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	//初始化logrus日志
	config.InitLog()

	//初始化配置
	err := config.InitConfig()
	if err != nil {
		logrus.Errorf("加载配置文件失败：%v", err.Error())
		return
	}

	//初始化mongoDb
	go common.InitMongo(config.Cfg.MongoUrl, config.Cfg.MongoDb)

	//初始化redis
	common.InitRedis(config.Cfg.RedisAddress, config.Cfg.RedisPasswrod, config.Cfg.RedisDb)

	//初始化缓存
	common.InitCache()

	//初始化bolt数据库
	/*boltDb, err := orm.InitBolt(config.Cfg.BoltDb)
	defer func() {
		if boltDb != nil {
			boltDb.Sync()
			boltDb.Close()
		}
	}()
	if err != nil {
		logrus.Errorf("初始化boltDB失败：%v", err.Error())
		return
	}

	//初始化sqllite数据库，sqllite只能在linux下打包运行，windows下跑不起来
	sqlliteDb, err := orm.InitSqllite(config.Cfg.SqlliteDb)
	defer func() {
		if sqlliteDb != nil {
			sqlliteDb.Close()
		}
	}()
	if err != nil {
		logrus.Errorf("初始化sqlliteDB失败：%v", err.Error())
		return
	}*/

	//初始化mysql数据库
	/*mysqlDb, err := orm.InitMysql(config.Cfg.MysqlDb)
	defer func() {
		if mysqlDb != nil {
			mysqlDb.Close()
		}
	}()
	if err != nil {
		logrus.Errorf("初始化mysqlDb失败：%v", err.Error())
		return
	}*/

	//初始化定时任务
	common.InitCron()
	common.AddCronTask("*/1 * * * * ?", func() {
		//fmt.Println("222")
	})

	//创建线程池
	threadPool, err := common.ThreadPoolCreate(300, func(i interface{}) {
		fmt.Println(i)
	})
	defer func() {
		if threadPool != nil {
			threadPool.Release()
		}
	}()
	if err != nil {
		logrus.Errorf("初始化线程池失败：%v", err.Error())
	}
	threadPool.Invoke("aaa")

	//初始化http服务
	go router.InitHttpServer(config.Cfg.HttpPort)

	//开启TCP服务端
	go common.TcpServer(5555, 2048, func(data []byte, conn net.Conn) {
		fmt.Println(string(data))
		fmt.Println(conn.RemoteAddr().String())
	})

	//开启UDP服务端
	go common.UdpServer(6666, 2048, func(data []byte, client *net.UDPAddr) {
		fmt.Println(string(data))
		fmt.Println(fmt.Sprintf("%v:%v", client.IP.String(), client.Port))
	})

	//开启grpc服务端
	go grpctest.GrpcServer(7777)

	//开启基于http2服务的rpc服务端
	go rpctest.RpcHttpServer(8888)

	//开启基于tcp服务的rpc服务端
	go rpctest.RpcTcpServer(9999)

	fmt.Println(111)
	//阻止主线程退出
	tiker := time.NewTicker(time.Hour)
	for {
		<-tiker.C
	}
}
