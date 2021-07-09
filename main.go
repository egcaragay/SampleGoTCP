package main

import (
	"log"
	"net"
	"sample/datastream"
	"time"
)

var ConnectedPlayers = map[int]*Player{}

func main() {

	l,err := net.Listen("tcp",":9999")
	if err != nil {
		log.Println(err,"Listen Error")
		return
	}

	defer func() {
		l.Close()
	}()

	for {
		c, err := l.Accept()
		if err != nil {
			return
		}
		go Connect(c)
	}

}

func Connect(conn net.Conn) {
	var player *Player

	defer func() {
		if player != nil {
			delete(ConnectedPlayers, player.id)
		}

		conn.Close()
	}()
	b := make([]byte,256)
	for {
		//Timeout 20 secs
		conn.SetDeadline(time.Now().Add(20 * time.Second))
		n, err := conn.Read(b)
		if err != nil {
			break
		} else if n > 0 {
			bytes := b[0:n]

			dsr := datastream.NewDataStreamReader(bytes)
			messageType := dsr.ReadByte()
			if messageType == 0 {
				id := dsr.ReadInt()
				player = &Player{id: id,conn: conn}
				ConnectedPlayers[id] = player
				log.Println("Connected")
				dw := datastream.NewDataStreamWriter()
				dw.WriteByte(0)
				conn.Write(dw.GetBuffer())

				dsw := datastream.DataStream{}
				dsw.WriteByte(1)
				dsw.WriteInt(id)
				for i := range ConnectedPlayers {
					p := ConnectedPlayers[i]
					if p.id != id {
						p.conn.Write(dsw.GetBuffer())
					}
				}
			} else if messageType == 2 {

				player.x = dsr.ReadFloat()
				player.y = dsr.ReadFloat()
				player.z = dsr.ReadFloat()

				log.Println("RUN",player.id)
				dsw := datastream.DataStream{}

				dsw.WriteByte(2)
				dsw.WriteInt(player.id)
				dsw.WriteFloat(player.x)
				dsw.WriteFloat(player.y)
				dsw.WriteFloat(player.z)
				for i := range ConnectedPlayers {
					p := ConnectedPlayers[i]
					p.conn.Write(dsw.GetBuffer())
				}
				
			} else {

			}
		}
	}
}