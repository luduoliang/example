package config

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

//配置信息
type Config struct {
	XMLName             xml.Name      `xml:"config"json:"-"`
	MongoUrl            string        `xml:"mongoUrl"json:"mongoUrl"`                       //mongo连接地址
	MongoDb             string        `xml:"mongoDb"json:"mongoDb"`                         //mongo数据库
	RedisAddress        string        `xml:"redisAddress"json:"redisAddress"`               //redis地址
	RedisPasswrod       string        `xml:"redisPasswrod"json:"redisPasswrod"`             //redis密码
	RedisDb             int           `xml:"redisDb"json:"redisDb"`                         //redis库
	BoltDb              string        `xml:"boltDb"json:"boltDb"`                           //bolt数据文件
	SqlliteDb           string        `xml:"sqlliteDb"json:"sqlliteDb"`                     //sqllite数据文件
	MysqlDb             string        `xml:"mysqlDb"json:"mysqlDb"`                         //mysql链接串
	HttpPort            string        `xml:"httpPort"json:"httpPort"`                       //本地http开启端口号
	JwtSecretKey        string        `xml:"jwtSecretKey"json:"jwtSecretKey"`               //jwt密钥
	JwtTokenExpriseHour int           `xml:"jwtTokenExpriseHour"json:"jwtTokenExpriseHour"` //jwt token过期小时数
	TestArray           TestArrayItem `xml:"testArray"json:"testArray"`                     //测试数组
}

type TestArrayItem struct {
	TestArrayItem []string `xml:"testArrayItem"json:"testArrayItem"` //测试数组
}

//配置信息
var Cfg *Config

//初始化日志
func InitLog() {
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	//日志文件存放路径
	path := "./log/date_"
	/* 日志轮转相关函数%Y%m%d%H%M"
	   `WithLinkName` 为最新的日志建立软连接
	   `WithRotationTime` 设置日志分割的时间，隔多久分割一次
	   WithMaxAge 和 WithRotationCount二者只能设置一个
	    `WithMaxAge` 设置文件清理前的最长保存时间
	    `WithRotationCount` 设置文件清理前最多保存的个数
	*/
	// 下面配置日志每隔 1 分钟轮转一个新文件，保留最近 3 分钟的日志文件，多余的自动清理掉。
	fileWriter, _ := rotatelogs.New(
		path+"%Y%m%d.log",
		//rotatelogs.WithLinkName(path),
		//rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationCount(15),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)

	//设置多输出，控制台和文件都输出
	writers := []io.Writer{
		fileWriter,
		os.Stdout}
	writersAll := io.MultiWriter(writers...)
	//logrus.SetOutput(os.Stdout)
	logrus.SetOutput(writersAll)
	//日志添加文件和行号
	logrus.AddHook(&GlobalHook{})
	//设置日志输入样式
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true, DisableColors: true, FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05.000"})
}

//初始化配置
func InitConfig() error {
	Cfg = new(Config)
	var err error
	var bsData []byte
	var configPath string = "./config/config.xml"
	bsData, err = ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	if len(bsData) == 0 {
		return err
	}
	err = xml.Unmarshal(bsData, &Cfg)
	if err != nil {
		return err
	}
	return nil
}

// logrus日志全局HOOK，输出日志中增加代码的所属文件与行号
type GlobalHook struct {
}

func (h *GlobalHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *GlobalHook) Fire(e *logrus.Entry) error {
	var fileLine string
	var skip = 5
	var i = skip
	var find = 0
	for i = 0; i < 13; i++ {
		funcName, file, line, ok := runtime.Caller(i)
		if ok {
			var nIndex = strings.Index(file, "logrus")
			if nIndex > 0 {
				continue
			} else if find == 0 {
				find = 1
				continue
			}
			// 取长路径的文件名
			file = filepath.Base(file)
			fileLine = fmt.Sprintf("%s:%s:%d", file, runtime.FuncForPC(funcName).Name(), line)
			break
		}
	}
	e.Data["file"] = fileLine
	return nil
}
