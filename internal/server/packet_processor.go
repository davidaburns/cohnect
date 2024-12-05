package server

import (
	"fmt"
	"net"

	"github.com/davidaburns/cohnect/internal/cache"
	"go.uber.org/zap"
)

type PacketProcessor struct {
	Log *zap.SugaredLogger
	Cache *cache.InMemoryCache
	Connection *net.UDPConn
}

func NewPacketProcessor(log *zap.SugaredLogger, cache *cache.InMemoryCache, connection *net.UDPConn) *PacketProcessor {
	return &PacketProcessor{
		Log: log,
		Cache: cache,
		Connection: connection,
	}
}

func (proc *PacketProcessor) ProcessPacket(packet *RequestPacket) error {
	var err error
	switch packet.Opcode {
	default:
		err = fmt.Errorf("invalid opcode: %v", packet.Opcode)
	}

	return err
}