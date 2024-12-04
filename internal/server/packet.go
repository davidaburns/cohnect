package server

import (
	"encoding/json"
	"fmt"

	"github.com/davidaburns/cohnect/internal/server/buffers"
	"github.com/google/uuid"
)

type RequestPacket struct {
	UUID uuid.UUID
	Opcode buffers.RequestOp
	BodyLength uint16
	Body map[string]any
}

func requestPacketFromBuffer(data []byte) (*RequestPacket, error) {
	if !buffers.RequestPacketBufferHasIdentifier(data) {
		return nil, fmt.Errorf("RequestPacket buffer has invalid identifier");
	}

	buf := buffers.GetRootAsRequestPacket(data, 0)
	uuid, err := uuid.FromBytes(buf.UuidBytes())
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
		UUID: uuid,
		Opcode: opcode,
		BodyLength: buf.Length(),
		Body: body,
	}, nil
}

func validRequestOpcode(op buffers.RequestOp) bool {
	switch op {
	case buffers.RequestOpGET, buffers.RequestOpSET, buffers.RequestOpSET_CLIENT_TAGS, buffers.RequestOpEXECUTE, buffers.RequestOpSUBSCRIBE, buffers.RequestOpBROADCAST:
		return true;
	default:
		return false;
	}
}