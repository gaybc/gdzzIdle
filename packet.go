package gdzzIdle

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/json-iterator/go"
	"golang.org/x/net/websocket"
	"strings"
	"time"
)
var salt = "askj8789kldksiewkszkm2323lkkl"

//index 发包序列号
func SendPacket(conn *websocket.Conn, data map[string]interface{}, index *int) {
	ts := fmt.Sprintf("%d", time.Now().Unix())
	sign := calSha1(salt + ts)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	data["keyTime"] = ts
	data["key"] = sign
	data["pIn"] = *index
	ret, _ := json.Marshal(data)
	retstr := string(ret)
	_ = websocket.Message.Send(conn, retstr)
	*index++
}

func calSha1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	ret := strings.ToUpper(hex.EncodeToString(h.Sum(nil))[0:6])
	return ret
}

func SendLoginPacket(conn *websocket.Conn, username string, index *int) {

	m := make(map[string]interface{})
	m["pktId"] = 0
	m["userName"] = username
	m["password"] = "ljs"
	m["plat"] = 0
	m["youke"] = false
	m["idfa"] = ""
	SendPacket(conn, m, index)
}
//5s
func SendHeartPacket(conn *websocket.Conn, index *int) {
		m := make(map[string]interface{})
		m["pktId"] = -1
		SendPacket(conn, m, index)
}
//15s
func SendGuaJiPacket(conn *websocket.Conn, index *int) {
		m := make(map[string]interface{})
		m["pktId"] = 5
		m["levelId"] = 0
		m["operate"] = 5
		m["danci"] = 0
		SendPacket(conn, m, index)
		time.Sleep(time.Second * 15)
}
