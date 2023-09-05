package tcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"

	"Ceds/conf"
)

func ListenAndServe() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.Config.Host(), conf.Config.Port()))
	if err != nil {
		log.Println("listen err : ", err)
	}
	defer listen.Close()
	log.Println("listen success")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("conn err : ", err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("client closed")
				break
			} else {
				log.Println("read err : ", err)
				break
			}
			return
		}
		if msg == "exit\r\n" {
			log.Println("client exit")
			conn.Close()
			break
		}
		// 将收到的信息发送给客户端
		conn.Write([]byte(msg))
	}
}
