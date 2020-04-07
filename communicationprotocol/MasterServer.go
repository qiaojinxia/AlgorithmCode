package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

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
			fmt.Println("连接关闭",&conn)
			return
		}
		length := BytesToInt(buf1) //字节转化为整数
		buf2 := make([]byte,length)
		n,err = conn.Read(buf2)
		if err != nil{
			panic(err)
		}
		fmt.Printf("Master收到字节 %s \n",buf2)
	}
}



func main(){
	server_listener,err:= net.Listen("tcp","0.0.0.0:8848")
	if err != nil{
		panic(err)
	}
	defer server_listener.Close()
	for {
		new_conn,err := server_listener.Accept()
		if err != nil{
			panic(err)
		}
		go Message_Handler(new_conn)
		for i:=0;i<10;i++{
			mystr := "hello," + strconv.Itoa(i)
			mystrlength := len(mystr)
			myb := IntToBytes(mystrlength)
			fmt.Println(mystr,mystrlength)
			new_conn.Write(myb)
			new_conn.Write([]byte(mystr))
		}

	}

	fmt.Println(IntToBytes(1))
	fmt.Println(BytesToInt(IntToBytes(1)))


}