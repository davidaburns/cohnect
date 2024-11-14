package server

import (
	"encoding/binary"
	"fmt"

	"github.com/google/uuid"
)

const (
	PACKET_ACCEPTED_VERSION uint8 = 1
	PACKET_BUFFER_TOTAL uint16 = 1472;
	PACKET_VERSION_SIZE uint16 = 1;
	PACKET_UUID_SIZE uint16 = 16;
	PACKET_OPCODE_SIZE uint16 = 2;
	PACKET_BODY_LENGTH_SIZE uint16 = 2;
	PACKET_HEADER_SIZE uint16 = PACKET_VERSION_SIZE + PACKET_UUID_SIZE + PACKET_OPCODE_SIZE + PACKET_BODY_LENGTH_SIZE;
	PACKET_BODY_SIZE uint16 = PACKET_BUFFER_TOTAL - (PACKET_VERSION_SIZE + PACKET_UUID_SIZE + PACKET_OPCODE_SIZE + PACKET_BODY_LENGTH_SIZE);
)

type Packet struct {
	Version uint8 			
	UUID uuid.UUID			
	Opcode Opcode
	BodyLength uint16
	Body []byte
}

func FromBytesBuffer(data []byte) (*Packet, error) {
	if len(data) < int(PACKET_HEADER_SIZE) || len(data) > int(PACKET_BUFFER_TOTAL) {
		return nil, fmt.Errorf("invalid packet size: %d", len(data))
	}

	var packet *Packet = &Packet{}
	var offset uint16 = 0

	packet.Version = data[offset]
	offset += 1

	if packet.Version != PACKET_ACCEPTED_VERSION {
		return nil, fmt.Errorf("invalid packet version: %d", packet.Version)
	}

	uuid, err := uuid.FromBytes(data[offset:offset+PACKET_UUID_SIZE])
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %v", err)
	}

	packet.UUID=uuid
	offset += PACKET_UUID_SIZE

	packet.Opcode = Opcode(binary.BigEndian.Uint16(data[offset:offset+PACKET_OPCODE_SIZE]))
	offset += PACKET_OPCODE_SIZE

	packet.BodyLength = binary.BigEndian.Uint16(data[offset:offset+PACKET_BODY_LENGTH_SIZE])
	offset += PACKET_BODY_LENGTH_SIZE

	packet.Body = data[offset:offset+packet.BodyLength+1]
	return packet, nil
}