package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"strings"
	"sync"
	"time"
)

import (
	"log"
)

type Argss struct {
	Op string
}

var RunningPc []string
var wg sync.WaitGroup

func main() {
	fmt.Println("服务启动中...")
	wg = sync.WaitGroup{}

	// 查询运行中服务端ip清单
	// 服务运行后 进行注册
	fmt.Println("正在查询运行中的电脑,请稍后...")

	CheckServer()
	if len(RunningPc) == 0 {
		fmt.Println("没有电脑在运行")
	} else {
		var IsShutdown string
		fmt.Println("找到以下电脑:", RunningPc)
		fmt.Print("是否要全部关机？(y|n)")
		fmt.Scanln(&IsShutdown)
		if strings.ToLower(IsShutdown) == "y" || strings.ToLower(IsShutdown) == "yes" {
			fmt.Println("正在关闭计算机,请稍后...")
			ShutDown(RunningPc)
		} else {
			fmt.Println("选择不关闭计算机")
		}

	}
	fmt.Println("5秒后自动退出...")
	time.Sleep(5 * time.Second)
}

func ShutDown(pcs []string) {
	for _, pc := range pcs {
		client, err := rpc.Dial("tcp", pc+":1234")
		if err != nil {
			fmt.Println("无法连接", pc, "关机失败")
			return
		}
		defer client.Close()

		var reply bool
		err = client.Call("Pmanager.Do", Argss{Op: "shutdown"}, &reply)
		if err != nil {
			fmt.Println("fail", err)
		}
		log.Println("关机" + pc + "成功！")
	}

}

func CheckServer() {
	for i := 41; i <= 41; i++ {
		wg.Add(1)
		go CheckPort(i, &wg)
	}
	wg.Wait()

}

func CheckPort(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	ips := "192.168.1." + strconv.Itoa(i)
	ipp := net.ParseIP(ips)
	tcpAddr := net.TCPAddr{
		IP:   ipp,
		Port: 1234,
	}
	client, err := net.DialTCP("tcp", nil, &tcpAddr)
	if err != nil {
		return
	}
	defer client.Close()
	RunningPc = append(RunningPc, ips)

}

func CleanRepeat(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}
	for v := range elements {
		if encountered[elements[v]] == true {
		} else {
			encountered[elements[v]] = true
			if strings.TrimSpace(elements[v]) != "" {
				result = append(result, strings.TrimSpace(elements[v]))
			}
		}
	}
	return result
}
