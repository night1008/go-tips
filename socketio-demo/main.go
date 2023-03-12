package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
)

type UUIDGenerator struct{}

func (g *UUIDGenerator) NewID() string {
	return uuid.Must(uuid.NewV4()).String()
}

type SocketioServer struct {
	sync.Mutex
	server               *socketio.Server
	connectionRequestIDs map[string]map[string]context.CancelFunc // 测试 manager session 数量是否一致, {connection_id:request_id:cancelFunc}
}

func main() {
	router := gin.New()

	server := SocketioServer{
		server: socketio.NewServer(&engineio.Options{
			SessionIDGenerator: &UUIDGenerator{},
		}),
		connectionRequestIDs: make(map[string]map[string]context.CancelFunc),
	}

	server.server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		server.Lock()
		server.connectionRequestIDs[s.ID()] = make(map[string]context.CancelFunc)
		log.Println("connected:", s.ID(), server.server.Count(), len(server.connectionRequestIDs))
		server.Unlock()
		return nil
	})

	server.server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		// 1. 不需要再另起 goroutine
		// 2. 即使连接断开，也会执行完该监听事件，因此需要有取消机制
		server.Lock()
		request_id := uuid.Must(uuid.NewV4()).String()
		if cancel, ok := server.connectionRequestIDs[s.ID()][request_id]; ok {
			cancel()
			// TODO 如果是前端生成 request_id，需要确认重连情况下不会重新生成
			// server.Unlock()
			// s.Emit("reply", "request_id "+request_id+" already existed")
			// return
		}
		ctx, cancel := context.WithCancel(context.Background())
		server.connectionRequestIDs[s.ID()][request_id] = cancel
		server.Unlock()

		select {
		case <-ctx.Done():
			return
		default:
			go func(ctx context.Context) {
				// 查询完成删除记录
				defer func() {
					server.Lock()
					log.Println("===> clean session request_id start", request_id, len(server.connectionRequestIDs[s.ID()]))
					delete(server.connectionRequestIDs[s.ID()], request_id)
					log.Println("===> clean session request_id end", request_id, len(server.connectionRequestIDs[s.ID()]))
					server.Unlock()
				}()

				select {
				case err := <-ctx.Done():
					log.Println("===> receive ctx done singal", err)
					return
				case <-time.After(10 * time.Second):
					log.Println("===> finish deal msg", msg)
					s.Emit("reply", "have "+msg)
				}
			}(ctx)
		}
	})

	server.server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason, s.ID())

		server.Lock()
		for _, cancel := range server.connectionRequestIDs[s.ID()] {
			cancel()
		}
		delete(server.connectionRequestIDs, s.ID())
		log.Println("disconnected:", s.ID(), server.server.Count(), len(server.connectionRequestIDs))
		server.Unlock()

		go func() {
			log.Println("===> clean disconnected start")
			time.Sleep(5 * time.Second)
			log.Println("===> clean disconnected end")
		}()
	})

	go func() {
		if err := server.server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.server.Close()

	router.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/socket.io/*any", gin.WrapH(server.server))
	router.POST("/socket.io/*any", gin.WrapH(server.server))
	router.StaticFS("/public", http.Dir("./asset"))

	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
