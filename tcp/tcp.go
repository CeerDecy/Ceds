package tcp

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"Ceds/conf"
	"Ceds/log"
)

// ListenAndServe 监听并启动服务
func ListenAndServe() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.Config.Host(), conf.Config.Port()))
	if err != nil {
		log.Log.Error("listen err : ", err)
	}
	defer listen.Close()
	log.Log.Info("server", "listen success")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Log.Error("conn err : ", err)
		}
		go handle(conn)
	}
}

// handle 处理客户端请求
func handle(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Log.Info("server", "client closed")
				break
			} else {
				log.Log.Error("read err : ", err)
				break
			}
			return
		}

		if msg == "exit\r\n" {
			log.Log.Info("server", "client exit")
			conn.Close()
			break
		}
		log.Log.Info("receive msg ", msg[:len(msg)-2])
		// 将收到的信息发送给客户端
		conn.Write([]byte(msg))
	}
}
