package common

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"reflect"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/tatsushid/go-fastping"
)

//获取设备ID
func GetDeviceID() (devid string, err error) {
	var macString string
	var macs []string
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error : " + err.Error())
		return
	}
	for _, inter := range interfaces {
		mac := inter.HardwareAddr //获取本机MAC地址
		//fmt.Println("MAC = ", mac)
		if mac.String() != "" {
			macs = append(macs, mac.String())
		}
	}
	sort.Strings(macs)
	for _, v := range macs {
		macString += v
		macString += ","
	}
	//fmt.Println("macString=", macString)
	h := md5.New()
	h.Write([]byte(macString))
	devid = hex.EncodeToString(h.Sum(nil))
	devid = strings.ToUpper(devid)
	// FB5290764EEE488301ABFB10B0A28FD8
	if len(devid) == 32 {
		devid = devid[0:16]
	}
	return
}

//md5
func Md5(password string) string {
	w := md5.New()
	io.WriteString(w, password)                   //将str写入到w中
	password_md5 := fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	return password_md5
}

//获取当前时间
func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location).UTC()
}

//随机一段整数
func RandNumber(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

//生成干扰字符串
func RandString(len int) string {
	str := "1123456789QWERTYUIOPASDFGHJKLZXVBNMqwertyuioplkjhgfdsamnbvcxzz"
	var return_str string = ""
	for i := 0; i < len; i++ {
		index := RandNumber(i+1, 57)
		return_str = return_str + str[index:index+1]
	}
	return return_str
}

//字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// stripslashes() 函数删除由 addslashes() 函数添加的反斜杠。
func Stripslashes(str string) string {
	dstRune := []rune{}
	strRune := []rune(str)
	strLenth := len(strRune)
	for i := 0; i < strLenth; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}

//截取字符串，含中文
func Substr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}
		if sl+rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}

//判断是否在数组中,arr为数组,item为元素
func InArray(item interface{}, arr interface{}) bool {
	itemType := strings.ToLower(reflect.TypeOf(item).Kind().String())
	switch itemType {
	case "string":
		if arrTemp, ok := arr.([]string); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(string); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "int":
		if arrTemp, ok := arr.([]int); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(int); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "uint":
		if arrTemp, ok := arr.([]uint); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(uint); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "complex128":
		if arrTemp, ok := arr.([]complex128); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(complex128); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "complex64":
		if arrTemp, ok := arr.([]complex64); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(complex64); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "float64":
		if arrTemp, ok := arr.([]float64); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(float64); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "float32":
		if arrTemp, ok := arr.([]float32); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(float32); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "int64":
		if arrTemp, ok := arr.([]int64); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(int64); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "int32":
		if arrTemp, ok := arr.([]int32); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(int32); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "int16":
		if arrTemp, ok := arr.([]int16); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(int16); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "int8":
		if arrTemp, ok := arr.([]int8); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(int8); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "uint64":
		if arrTemp, ok := arr.([]uint64); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(uint64); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "uint32":
		if arrTemp, ok := arr.([]uint32); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(uint32); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "uint16":
		if arrTemp, ok := arr.([]uint16); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(uint16); ok2 && v == itemStr {
					return true
				}
			}
		}
	case "uint8":
		if arrTemp, ok := arr.([]uint8); ok {
			for _, v := range arrTemp {
				if itemStr, ok2 := item.(uint8); ok2 && v == itemStr {
					return true
				}
			}
		}
	default:
		return false
	}
	return false
}

//获取本地IP
func GetLocalIp() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			//ip := getIpFromAddr(addr)
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip.String(), nil
		}
	}
	return "", fmt.Errorf("connected to the network?")
}

//检查网络是否能拼通，及延时时间
//isOn是否能拼通，timeout为延时时间(毫秒)
func NetworkCheck(ip string) (canPing bool, timeout int64) {
	var wg sync.WaitGroup
	wg.Add(1)
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ip)
	if err != nil {
		return
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		canPing = true
		timeout = rtt.Milliseconds()
	}
	p.OnIdle = func() {
		wg.Done()
	}
	err = p.Run()
	if err != nil {
		return
	}
	wg.Wait()
	return
}

//数字转换字节数组
func IntToBytes(n uint32) []byte {
	x := uint32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节数组转换数字
func BytesToInt(b []byte) int {
	buf := bytes.NewBuffer(b)
	var tmp uint32
	binary.Read(buf, binary.BigEndian, &tmp)
	return int(tmp)
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
