// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package buffers

import "strconv"

type RequestOp uint16

const (
	RequestOpPING                 RequestOp = 0
	RequestOpSESSION_START        RequestOp = 1
	RequestOpSESSION_END          RequestOp = 2
	RequestOpEVENT                RequestOp = 3
	RequestOpREGISTER_CLIENT_TAGS RequestOp = 4
)

var EnumNamesRequestOp = map[RequestOp]string{
	RequestOpPING:                 "PING",
	RequestOpSESSION_START:        "SESSION_START",
	RequestOpSESSION_END:          "SESSION_END",
	RequestOpEVENT:                "EVENT",
	RequestOpREGISTER_CLIENT_TAGS: "REGISTER_CLIENT_TAGS",
}

var EnumValuesRequestOp = map[string]RequestOp{
	"PING":                 RequestOpPING,
	"SESSION_START":        RequestOpSESSION_START,
	"SESSION_END":          RequestOpSESSION_END,
	"EVENT":                RequestOpEVENT,
	"REGISTER_CLIENT_TAGS": RequestOpREGISTER_CLIENT_TAGS,
}

func (v RequestOp) String() string {
	if s, ok := EnumNamesRequestOp[v]; ok {
		return s
	}
	return "RequestOp(" + strconv.FormatInt(int64(v), 10) + ")"
}
