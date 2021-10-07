package main
import (
	"fmt"
	"net"
)

func process(conn net.Conn) {

	defer conn.Close()
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Server Read err=", err)
			break
		}
		fmt.Print(string(buf[:n]))
	}

}

func main() {

	fmt.Println("Start Listening...")
	listen, err := net.Listen("tcp","0.0.0.0:7000")
	if err != nil {
		fmt.Println("listen err=", err)
		return 
	}
	defer listen.Close()
	for {
		fmt.Println("Waiting...")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=",err)
		} else {
			fmt.Printf("Accept() suc con=%v client ip=%v\n", conn, conn.RemoteAddr())
		}
		go process(conn)
	}
}