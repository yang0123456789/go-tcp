package znet

import (
	"fmt"
	"go-tcp/ziface"
	"net"
)

type Server struct {
	//服务器的名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器监听的IP
	IP string
	//服务器监听的端口
	Port int
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP :%s,Port %d,is starting\n", s.IP, s.Port)
	go func() {
		// 1获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Printf("err:", err)
			return
		}
		//2监听服务器的地址
		lister, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf(s.IPVersion, "err:", err)
			return
		}
		fmt.Printf("start success ")
		//3 阻塞的等待客户端链接，处理客户端链接业务(读写)
		for {
			conn, err := lister.AcceptTCP()
			if err != nil {
				fmt.Printf("Accept err:", err)
				continue
			}
			//建立连接，做一些业务，
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Printf("buf receive err:", err)
						continue
					}
					//	回显
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Printf("write back err:", err)
						continue
					}

				}
			}()

		}
	}()
}
func (s *Server) Stop() {
	//TODO 停止服务，释放资源
}
func (s *Server) Server() {
	//启动server服务
	s.Start()
	//TODO 启动服务器之外的事情

	//阻塞状态
	select {}

}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
