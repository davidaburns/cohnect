package server

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/davidaburns/cohnect/internal/server/buffers"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/google/uuid"
)

const RESPONSE_BODY_MAX_SIZE int = 1452;

type ResponseSuccess struct {
	Status string `json:"status"`
}

type ResponseFailure struct {
	Status string		 	`json:"status"`
	ErrorCode string		`json:"errorCode"`
	ErrorMessage string		`json:"errorMessage"`
}

type ResponseNotification struct {
	Type string 			`json:"type"`
	Body map[string]any 	`json:"body"`
}

func NewResponse(correlationId uuid.UUID, responseType buffers.ClientMessageType, body any) ([]byte, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	if (len(bodyBytes) > RESPONSE_BODY_MAX_SIZE) {
		return nil, fmt.Errorf("response body exceeds limit of %v bytes", RESPONSE_BODY_MAX_SIZE)
	}

	builder := flatbuffers.NewBuilder(0)
	correlationIdOffset := builder.CreateByteVector(correlationId[:])
	bodyOffset := builder.CreateByteVector(bodyBytes[:])

	buffers.ClientMessagePacketStart(builder)
	buffers.ClientMessagePacketAddCorrelationId(builder, correlationIdOffset)
	buffers.ClientMessagePacketAddType(builder, responseType)
	buffers.ClientMessagePacketAddLength(builder, uint16(len(bodyBytes)))
	buffers.ClientMessagePacketAddBody(builder, bodyOffset)

	response := buffers.ClientMessagePacketEnd(builder)
	builder.Finish(response)

	return builder.FinishedBytes(), nil
}

func SendClientResponse(host string, port int, response []byte) error {
	addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, strconv.Itoa(port)))
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(response)
	if err != nil {
		return err
	}
	
	return nil
}