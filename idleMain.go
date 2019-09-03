package gdzzIdle

import (
	"fmt"
	"golang.org/x/net/websocket"
	"time"
)

var origin = "http://127.0.0.1:8080/"
var serverip = "ws://tapandroid3.maobugames.com:35003"

var Users = make(map[string]interface{})

func main() {

	go guajiTask("mao0403603bu")
	time.Sleep(time.Duration(5) * time.Second)
	time.Sleep(time.Duration(6000) * time.Second)
}


func guajiTask(loginName string) {
	stop := make(chan int)
	ws, err := websocket.Dial(serverip, "", origin)
	if err != nil {
		fmt.Println(err)
	}
	Users[loginName] = ws
	go sendPacketThread(ws, loginName, stop)

	for {
		if (Users[loginName] == nil) {
			stop <- 1
			return
		}
		var msg [512]byte
		_, err := ws.Read(msg[:]) //此处阻塞，等待有数据可读
		if err != nil {
			fmt.Println("read:", err)
			time.Sleep(time.Duration(6) * time.Second)
			guajiTask(loginName)
		}

		fmt.Printf(" %s\n", msg)
	}

}

func sendPacketThread(conn *websocket.Conn, loginName string, stop chan int) {
	index := 1
	sleepTime := 0
	SendLoginPacket(conn, loginName, &index)
	for {
		select {
		case msg := <-stop:
			switch msg {
			case 1:
				fmt.Println("received message stop")
				return
			}

		default:
		}
		time.Sleep(time.Second * 5)
		sleepTime += 5
		SendHeartPacket(conn, &index)
		if sleepTime%15 == 0 { // 15s
			SendGuaJiPacket(conn, &index)

		}

	}
}
