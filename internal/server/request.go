package server

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/davidaburns/cohnect/internal/cache"
	"github.com/davidaburns/cohnect/internal/server/buffers"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type RequestPacket struct {
	CorrelationId uuid.UUID
	Opcode buffers.RequestOp
	BodyLength uint16
	Body map[string]any
	ClientAddr *net.UDPAddr
	ClientResponsePort int
}

type RequestProcessor struct {
	Log *zap.SugaredLogger
	Cache *cache.InMemoryCache
	Connection *net.UDPConn
}


// NOTE: This seems messy, lets try and find a better way of constructing a packet
func NewRequestFromBuffer(data []byte, addr *net.UDPAddr, responsePort int) (*RequestPacket, error) {
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
		ClientResponsePort: responsePort,
	}, nil
}

func NewRequestProcessor(log *zap.SugaredLogger, cache *cache.InMemoryCache, connection *net.UDPConn) *RequestProcessor {
	return &RequestProcessor{
		Log: log,
		Cache: cache,
		Connection: connection,
	}
}

func (pp *RequestProcessor) Process(packet *RequestPacket) error {
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

func (pp *RequestProcessor) SendClientResponse(packet *RequestPacket, responseType buffers.ClientMessageType, body []byte) error {
	if len(body) > 1452 {
		return fmt.Errorf("response body exceeds limit of 1452 bytes")
	}

	builder := flatbuffers.NewBuilder(0)
	correlationIdOffset := builder.CreateByteVector(packet.CorrelationId[:])
	bodyOffset := builder.CreateByteVector(body[:])

	buffers.ClientMessagePacketStart(builder)
	buffers.ClientMessagePacketAddCorrelationId(builder, correlationIdOffset)
	buffers.ClientMessagePacketAddType(builder, responseType)
	buffers.ClientMessagePacketAddLength(builder, uint16(len(body)))
	buffers.ClientMessagePacketAddBody(builder, bodyOffset)
	
	response := buffers.ClientMessagePacketEnd(builder)
	builder.Finish(response)

	responseBytes := builder.FinishedBytes()

	addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(packet.ClientAddr.IP.String(), strconv.Itoa(packet.ClientResponsePort)))
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}

	defer conn.Close()

	pp.Log.Infof("Sending %v to client %v %v", responseType.String(), addr.IP, addr.Port)
	_, err = conn.Write(responseBytes)
	if err != nil {
		return err
	}

	return nil
}

func (pp *RequestProcessor) ProcessPing(packet *RequestPacket) error {
	pp.Log.Info("Processing PING request")
	response, err := NewResponse(packet.CorrelationId, buffers.ClientMessageTypePING_SUCCESS, ResponseSuccess { Status: "OK"})
	if err != nil {
		return err
	}

	err = SendClientResponse(packet.ClientAddr.IP.String(), packet.ClientResponsePort, response)
	if err != nil {
		pp.Log.Errorf("Error sending response: %v", err)
		return err
	}

	return nil
}

func (pp *RequestProcessor) ProcessSessionStart(packet *RequestPacket) error {
	pp.Log.Info("Processing SESSION_START request")
	return nil;
}

func (pp *RequestProcessor) ProcessSessionEnd(packet *RequestPacket) error {
	pp.Log.Info("Processing SESSION_END request")
	return nil;
}

func (pp *RequestProcessor) ProcessEvent(packet *RequestPacket) error {
	pp.Log.Info("Processing EVENT request")
	return nil;
}

func (pp *RequestProcessor) ProcessRegisterTags(packet *RequestPacket) error {
	pp.Log.Info("Processing REGISTER_CLIENT_TAGS request")
	return nil;
}