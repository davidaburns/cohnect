package server

import (
	"fmt"
	"net"

	"github.com/davidaburns/cohnect/internal/cache"
	"github.com/davidaburns/cohnect/internal/server/buffers"
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
	switch packet.Opcode {
	case buffers.RequestOpGET:
		key, ok := packet.Body["Key"].(string)
		if !ok {
			return fmt.Errorf("GET operation body does not include 'Key' property")
		}

		proc.processGET(key)
	case buffers.RequestOpSET:
		key, ok := packet.Body["Key"].(string)
		if !ok {
			return fmt.Errorf("SET operation body does not include 'Key' property")
		}

		data, ok := packet.Body["Data"]
		if !ok {
			return fmt.Errorf("SET operation body does not include 'Data' property")
		}

		proc.processSET(key, data)
	case buffers.RequestOpSET_CLIENT_TAGS:
		proc.Log.Warn("SET_CLIENT_TAGS operation not implemented")
	case buffers.RequestOpEXECUTE:
		proc.Log.Warn("EXECUTE operation not implemented")
	case buffers.RequestOpSUBSCRIBE:
		proc.Log.Warn("SUBSCRIBE operation not implemented")
	case buffers.RequestOpBROADCAST:
		proc.Log.Warn("BROADCAST operation not implemented")
	default:
		return fmt.Errorf("invalid opcode: %v", packet.Opcode)
	}

	return nil
}

func (proc *PacketProcessor) processGET(key string) {
	value, ok := proc.Cache.Get(key)
	if !ok {
		proc.Log.Warnf("Cache key '%v' has not been set or has expired", key)
		return
	}

	proc.Log.Infof("Cache value [%v]: %v", key, value)
}

func (proc *PacketProcessor) processSET(key string, value any) {
	proc.Cache.Set(key, value)
	proc.Log.Infof("Cache [%v]=%v", key, value)
}