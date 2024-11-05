package server

type PacketOpcode uint16

const (
	GET PacketOpcode = iota
	SET
	SET_CLIENT_TAGS
	EXECUTE
	STREAM
	BROADCAST
)

func ValidOpcode(op PacketOpcode) bool {
	switch op {
	case GET, SET, SET_CLIENT_TAGS, EXECUTE, STREAM, BROADCAST: return true;
	default: return false;
	}
}