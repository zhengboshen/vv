package logic

import (
	"github.com/zhengboshen/vv/logic/service"
)

func main() {
	// 初始化一个grpc对象
	// service := grpc.NewServer()
	// // 注册服务
	// vvproto.Tester()
	// // 设置监听
	// listener, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer listener.Close()
	// //启动服务
	// err = service.Serve(listener)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	service.NewEventService2()
	// vvproto.Tester()
}
