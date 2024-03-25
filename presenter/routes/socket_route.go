package routes

import (
	"fmt"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
)

func SocketRoutes(r *mux.Router, serv *socketio.Server) {

	serv.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		serv.JoinRoom("/", "GeneralRoom", s)
		fmt.Println("connected:", s.ID())
		return nil
	})

	serv.OnEvent("/", "NewFeed", func(s socketio.Conn, msg string) {
		fmt.Println("NewFeed event is hit")
		serv.BroadcastToNamespace("/", "NewFeed")
	})

	serv.OnEvent("/", "StartChat", func(s socketio.Conn, msg string) {

		serv.JoinRoom("/", "<roodId>", s)
	})

	serv.OnEvent("/", "ChatMsg", func(s socketio.Conn, msg string) {

		serv.BroadcastToRoom("/", "<roomId>", "chat message", msg)
	})

	serv.OnEvent("/", "LeaveChat", func(s socketio.Conn, msg string) {

		serv.LeaveRoom("/", "<roomId>", s)
	})

	serv.OnError("/", func(s socketio.Conn, e error) {
		// server.Remove(s.ID())
		fmt.Println("meet error:", e)
	})

	serv.OnDisconnect("/", func(c socketio.Conn, s string) {
		serv.LeaveAllRooms("/", c)

	})

	go serv.Serve()
	defer serv.Close()

	http.Handle("/socket.io/", serv)
}
