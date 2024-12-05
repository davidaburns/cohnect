package server

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/davidaburns/cohnect/internal/cache"
	"github.com/davidaburns/cohnect/internal/server/buffers"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type RequestPacket struct {
	CorrelationId uuid.UUID
	Opcode buffers.RequestOp
	BodyLength uint16
	Body map[string]any
	ClientAddr *net.UDPAddr
}

type PacketProcessor struct {
	Log *zap.SugaredLogger
	Cache *cache.InMemoryCache
	Connection *net.UDPConn
}

func NewPacketFromBuffer(data []byte, addr *net.UDPAddr) (*RequestPacket, error) {
	if !buffers.RequestPacketBufferHasIdentifier(data) {
		return nil, fmt.Errorf("RequestPacket buffer has invalid identifier");
	}

	buf := buffers.GetRootAsRequestPacket(data, 0)
	uuid, err := uuid.FromBytes(buf.CorrelationIdBytes())
	if err != nil {
		return nil, fmt.Errorf("unable to parse uuid: %v", err)
	}

	opcode := buf.Opcode()
	switch opcode {
	case buffers.RequestOpPING, buffers.RequestOpSESSION_START, buffers.RequestOpSESSION_END, buffers.RequestOpEVENT, buffers.RequestOpREGISTER_CLIENT_TAGS:
	default:
		return nil, fmt.Errorf("invalid request opcode: %v", opcode)
	}

	var body map[string]any;
	if err := json.Unmarshal(buf.BodyBytes(), &body); err != nil {
		return nil, fmt.Errorf("error while unmarshaling json from packet body: %v", err)
	}

	return &RequestPacket{
		CorrelationId: uuid,
		Opcode: opcode,
		BodyLength: buf.Length(),
		Body: body,
		ClientAddr: addr,
	}, nil
}

func NewPacketProcessor(log *zap.SugaredLogger, cache *cache.InMemoryCache, connection *net.UDPConn) *PacketProcessor {
	return &PacketProcessor{
		Log: log,
		Cache: cache,
		Connection: connection,
	}
}

func (pp *PacketProcessor) Process(packet *RequestPacket) error {
	switch packet.Opcode {
	case buffers.RequestOpPING:
		pp.ProcessPing(packet)
	case buffers.RequestOpSESSION_START:
		pp.ProcessSessionStart(packet)
	case buffers.RequestOpSESSION_END:
		pp.ProcessSessionEnd(packet)
	case buffers.RequestOpEVENT:
		pp.ProcessEvent(packet)
	case buffers.RequestOpREGISTER_CLIENT_TAGS:
		pp.ProcessRegisterTags(packet)
	default:
		return fmt.Errorf("unknown request opcode: %v", packet.Opcode)
	}

	return nil
}

func (pp *PacketProcessor) ProcessPing(packet *RequestPacket) error {
	pp.Log.Info("Processing PING request")
	return nil
}

func (pp *PacketProcessor) ProcessSessionStart(packet *RequestPacket) error {
	pp.Log.Info("Processing SESSION_START request")
	return nil;
}

func (pp *PacketProcessor) ProcessSessionEnd(packet *RequestPacket) error {
	pp.Log.Info("Processing SESSION_END request")
	return nil;
}

func (pp *PacketProcessor) ProcessEvent(packet *RequestPacket) error {
	pp.Log.Info("Processing EVENT request")
	return nil;
}

func (pp *PacketProcessor) ProcessRegisterTags(packet *RequestPacket) error {
	pp.Log.Info("Processing REGISTER_CLIENT_TAGS request")
	return nil;
}