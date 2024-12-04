// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package buffers

import "strconv"

type RequestOp uint16

const (
	RequestOpGET             RequestOp = 0
	RequestOpSET             RequestOp = 1
	RequestOpSET_CLIENT_TAGS RequestOp = 2
	RequestOpEXECUTE         RequestOp = 3
	RequestOpSUBSCRIBE       RequestOp = 4
	RequestOpBROADCAST       RequestOp = 5
)

var EnumNamesRequestOp = map[RequestOp]string{
	RequestOpGET:             "GET",
	RequestOpSET:             "SET",
	RequestOpSET_CLIENT_TAGS: "SET_CLIENT_TAGS",
	RequestOpEXECUTE:         "EXECUTE",
	RequestOpSUBSCRIBE:       "SUBSCRIBE",
	RequestOpBROADCAST:       "BROADCAST",
}

var EnumValuesRequestOp = map[string]RequestOp{
	"GET":             RequestOpGET,
	"SET":             RequestOpSET,
	"SET_CLIENT_TAGS": RequestOpSET_CLIENT_TAGS,
	"EXECUTE":         RequestOpEXECUTE,
	"SUBSCRIBE":       RequestOpSUBSCRIBE,
	"BROADCAST":       RequestOpBROADCAST,
}

func (v RequestOp) String() string {
	if s, ok := EnumNamesRequestOp[v]; ok {
		return s
	}
	return "RequestOp(" + strconv.FormatInt(int64(v), 10) + ")"
}