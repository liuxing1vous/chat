package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//接收文件内容
func RecvFile(filename string,conn net.Conn)  {
	//新建文件
	f,err1:=os.Create(filename)
	if err1!=nil{
		fmt.Println("os.Create err1:",err1)
		return
	}

	buf:=make([]byte,1024*4)
	//接受多少，写多少
	for   {
		n,err2:=conn.Read(buf)
		if err2!=nil{
			if err2==io.EOF{
				fmt.Println("文件接收完毕")
			}else {
				fmt.Println("conn.Read() err2=", err2)
			}
			return
		}
		if n==0 {
			fmt.Println("n==0 文件接收完毕")
			break

		}
		f.Write(buf[:n])//往文件写入内容

	}
}
func main() {
	//监听
	listenner,err1:=net.Listen("tcp","127.0.0.1:8000")
	if err1!=nil{
		fmt.Println("net.Listen err1:",err1)
		return
	}
	defer listenner.Close()
	//阻塞等待用户连接
	conn,err2:=listenner.Accept()
	if err2!=nil{
		fmt.Println("listenner.Accept err2:",err2)
		return
	}

	buf:=make([]byte,1024)
	var n int
	n,err3:=conn.Read(buf)
	if err3!=nil{
		fmt.Println("conn.Read() err3=",err3)
		return
	}
	defer conn.Close()

	filename:=string(buf[:n])
	//回复"OK"
	conn.Write([]byte("OK"))
	//接收内容
	RecvFile(filename,conn)
}
