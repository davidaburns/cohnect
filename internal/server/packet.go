package server

import (
	"encoding/binary"
	"fmt"

	"github.com/google/uuid"
)

const ACCEPTED_VERSION uint8 = 1

const PACKET_BUFFER_TOTAL uint16 = 1472;
const PACKET_VERSION_SIZE uint16 = 1;
const PACKET_UUID_SIZE uint16 = 16;
const PACKET_OPCODE_SIZE uint16 = 2;
const PACKET_BODY_SIZE uint16 = PACKET_BUFFER_TOTAL - (PACKET_VERSION_SIZE + PACKET_UUID_SIZE + PACKET_OPCODE_SIZE);

type Packet struct {
	Version uint8
	UUID uuid.UUID
	Opcode PacketOpcode
	Body []byte
}

func FromBytesBuffer(data []byte) (*Packet, error) {
	if len(data) < int(PACKET_VERSION_SIZE + PACKET_UUID_SIZE + PACKET_OPCODE_SIZE) || len(data) > int(PACKET_BUFFER_TOTAL) {
		return nil, fmt.Errorf("invalid packet size: %d", len(data))
	}

	packet := &Packet{
		Version: uint8(data[0]),
		UUID: uuid.UUID(data[1:17]),
		Opcode: PacketOpcode(binary.BigEndian.Uint16(data[17:19])),
		Body: data[19:],
	}

	if packet.Version != ACCEPTED_VERSION {
		return nil, fmt.Errorf("invalid packet version: %d", packet.Version)
	}

	return packet, nil
}