package server

import (
	"fmt"
	"net"

	"github.com/davidaburns/cohnect/config"
	"go.uber.org/zap"
)

type Server struct {
	Config *config.Config
	Log *zap.SugaredLogger
}

func CreateNew(config *config.Config, log *zap.SugaredLogger) *Server {
	return &Server {
		Config: config,
		Log: log,
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

		defer conn.Close()
		ready <- true

		server.Log.Infof("UDP Server listening on %s:%d", server.Config.Server.Host, server.Config.Server.Port)

		buffer := make([]byte, 1024)
		for {
			select {
			case <-done:
				server.Log.Info("Shutting down UDP server...")
				return
			default:
				bytes, client, err := conn.ReadFromUDP(buffer)
				if err != nil {
					server.Log.Errorf("Error reading from UDP connection: %s", err)
					continue
				}

				server.Log.Infof("Recieved %d bytes from %s: %s", bytes, client, string(buffer[:bytes]))
				packet, err := FromBytesBuffer(buffer)
				if err != nil {
					server.Log.Errorf("Error parsing packet: %s", err)
				}
				
				server.Log.Infof("Parsed Packet: %v", packet)
			}
		}
	}()

	if success := <-ready; !success {
		return fmt.Errorf("failed to start UDP listener")
	}

	return nil
}