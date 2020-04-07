package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/LuaProject/pkg/mod/golang.org/x/net@v0.0.0-20180906233101-161cd47e91fd/html"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)



func visit(links []string,n *html.Node) []string{
	if n.Type == html.ElementNode && n.Data == "a"{
		for _,a := range n.Attr{
			if a.Key == "href"{
				links = append(links,a.Val)
			}
		}
	}
	for c:= n.FirstChild;c != nil;c =c.NextSibling{
		links = visit(links,c)
	}
	return links
}



/**
 * Created by @CaomaoBoy on 2020/2/20.
 *  email:<115882934@qq.com>
 */

func IntToBytes(n int)[]byte {
	data := int64(n)
	bytebuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytebuffer,binary.BigEndian,data)
	return bytebuffer.Bytes()
}

func BytesToInt(btys []byte)int{
	bytebuffer := bytes.NewBuffer(btys)
	var data int64
	binary.Read(bytebuffer,binary.BigEndian,&data)
	return int(data)
}


func Message_Handler(conn net.Conn){
	defer conn.Close()
	for {
		buf1 := make([]byte,8)
		n,err := conn.Read(buf1)
		if err != nil || n!= 8{
			fmt.Println("连接关闭",conn)
			return
		}
		length := BytesToInt(buf1) //字节转化为整数
		buf2 := make([]byte,length)
		n,err = conn.Read(buf2)
		if err != nil{
			panic(err)
		}
		fmt.Printf("Salver收到字节 %s \n",buf2)
		reponse := "[确认收到Master字节]:" + string(buf2)
		strlength := len(reponse)
		mybytes := IntToBytes(strlength)
		conn.Write(mybytes)
		conn.Write([]byte(reponse))

	}
}

func main(){
	tcpaddr,err := net.ResolveTCPAddr("tcp4","0.0.0.0:8848")
	if err != nil{
		panic(err)
	}
	conn,err := net.DialTCP("tcp",nil,tcpaddr)
	if err != nil{
		panic(err)
	}
	go Message_Handler(conn)
	time.Sleep(time.Second * 10)
}
