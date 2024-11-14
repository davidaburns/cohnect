package server

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Opcode uint16

const (
	GET Opcode = iota 		// 0x01
	SET						// 0x02
	SET_CLIENT_TAGS			// 0x03
	EXECUTE					// 0x04
	SUBSCRIBE				// 0x05
	BROADCAST				// 0x06
)

type Op struct {
	Op Opcode
	UUID uuid.UUID
	Body map[string]any
}

func OpFromPacket(packet *Packet) (*Op, error) {
	var body map[string]any;
	if err := json.Unmarshal(packet.Body[1:], &body); err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %v", err)
	}

	return &Op{
		Op: packet.Opcode,
		UUID: packet.UUID,
		Body: body,
	}, nil
}

func ValidOpcode(op Opcode) bool {
	switch op {
	case GET, SET, SET_CLIENT_TAGS, EXECUTE, SUBSCRIBE, BROADCAST: return true;
	default: return false;
	}
}

func OpToString(op Opcode) string {
	switch op {
	case GET: return "GET";
	case SET: return "SET";
	case SET_CLIENT_TAGS: return "SET_CLIENT_TAGS";
	case EXECUTE: return "EXECUTE";
	case SUBSCRIBE: return "SUBSCRIBE";
	case BROADCAST: return "BROADCAST";
	default: return ""
	}
}