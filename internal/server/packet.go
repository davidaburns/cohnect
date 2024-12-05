package server

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/davidaburns/cohnect/internal/server/buffers"
	"github.com/google/uuid"
)

type RequestPacket struct {
	CorrelationId uuid.UUID
	Opcode buffers.RequestOp
	BodyLength uint16
	Body map[string]any
	ClientAddr *net.UDPAddr
}

func requestPacketFromBuffer(data []byte, addr *net.UDPAddr) (*RequestPacket, error) {
	if !buffers.RequestPacketBufferHasIdentifier(data) {
		return nil, fmt.Errorf("RequestPacket buffer has invalid identifier");
	}

	buf := buffers.GetRootAsRequestPacket(data, 0)
	uuid, err := uuid.FromBytes(buf.CorrelationIdBytes())
	if err != nil {
		return nil, fmt.Errorf("unable to parse uuid: %v", err)
	}

	opcode := buf.Opcode()
	if !validRequestOpcode(opcode) {
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

func validRequestOpcode(op buffers.RequestOp) bool {
	switch op {
	case buffers.RequestOpPING, buffers.RequestOpSESSION_START, buffers.RequestOpSESSION_END, buffers.RequestOpEVENT, buffers.RequestOpREGISTER_CLIENT_TAGS:
		return true;
	default:
		return false;
	}
}