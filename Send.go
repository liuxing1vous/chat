package main

import (
	"fmt"
	"os"
	"net"
	"io"
)

func SendFile(path string,conn net.Conn)  {
	//以只读方式打开文件
	f,err:=os.Open(path)
	if err!=nil{
		fmt.Println("os.Open err:",err)
		return
	}
	defer f.Close()
	buf:=make([]byte,1024*4)
	//读文件，读多少发多少
	for {
		n,err:=f.Read(buf)//从文件读取内容
		if err!=nil{
			if err == io.EOF{
				fmt.Println("文件发送完毕")
			}else {
				fmt.Println("err:",err)
			}
			return
		}
		//发送内容
		conn.Write(buf[:n])
	}

}
func main() {
	//提示输入文件
	fmt.Println("请输入需要传输的文件：")
	var path string
	fmt.Scan(&path)
	//获取文件名info.Name()
	info,err1:=os.Stat(path)
	if err1!=nil{
		fmt.Println("os.Stat err1:",err1)
		return
	}
	//主动连接服务器
	conn,err2 :=net.Dial("tcp","127.0.0.1:8000")
	if err2!=nil{
		fmt.Println("net.Dial err2:",err2)
		return
	}
	//程序结束之后要关闭connect
	defer conn.Close()

	//给接收方。先发送文件名
	_,err3:=conn.Write([]byte(info.Name()))
	if err3!=nil{
		fmt.Println("conn.Write err3:",err3)
		return
	}
	//接收对方的回复，如果回复OK，则开始发送内容
	var n int
	buf:=make([]byte,1024)
	n,err4:=conn.Read(buf)
	if err4!=nil{
		fmt.Println("conn.Read err4:",err4)
		return
	}
	if "ok"==string(buf[:n]){
		SendFile(path,conn)
	}//发送文件内容
}
