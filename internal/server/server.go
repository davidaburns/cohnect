package server

import (
	"fmt"
	"net"

	"github.com/davidaburns/cohnect/config"
	"github.com/davidaburns/cohnect/internal/cache"
	"go.uber.org/zap"
)

type Server struct {
	Config *config.Config
	Log *zap.SugaredLogger
	Cache *cache.InMemoryCache
	Connection *net.UDPConn
}

func CreateNew(config *config.Config, log *zap.SugaredLogger, cache *cache.InMemoryCache) *Server {
	return &Server {
		Config: config,
		Log: log,
		Cache: cache,
	}
}

func (server *Server) Start(done chan bool) error {
	ready := make(chan bool)
	go func() {
		addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", server.Config.Server.Host, server.Config.Server.Port))
		if err != nil {
			server.Log.Errorf("Error resolving udp address: %s", err)
			ready <- false
			return
		}

		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			server.Log.Errorf("Error setting up udp server: %s", err)
			ready <- false
			return
		}

		server.Connection = conn;

		defer server.Connection.Close()
		ready <- true

		server.Log.Infof("UDP Server listening on %s:%d", server.Config.Server.Host, server.Config.Server.Port)
		processor := NewPacketProcessor(server.Log, server.Cache, server.Connection)

		buffer := make([]byte, 1024)
		for {
			select {
			case <-done:
				server.Log.Info("Shutting down UDP server...")
				return
			default:
				_, addr, err := server.Connection.ReadFromUDP(buffer)
				if err != nil {
					server.Log.Errorf("Error reading from UDP connection: %s", err)
					continue
				}

				packet, err := requestPacketFromBuffer(buffer, addr)
				if err != nil {
					server.Log.Errorf("Error parsing packet: %s", err)
					continue
				}

				if err := processor.ProcessPacket(packet); err != nil {
					server.Log.Errorf("Error processing packet: %s", err)
				}
			}
		}
	}()

	if success := <-ready; !success {
		return fmt.Errorf("failed to start UDP listener")
	}

	return nil
}