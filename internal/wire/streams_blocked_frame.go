package wire

import (
	"bytes"

	"github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/lucas-clemente/quic-go/internal/utils"
)

// A StreamsBlockedFrame is a STREAMS_BLOCKED frame
type StreamsBlockedFrame struct {
	StreamID protocol.StreamID
}

func parseStreamIDBlockedFrame(r *bytes.Reader, _ protocol.VersionNumber) (*StreamsBlockedFrame, error) {
	if _, err := r.ReadByte(); err != nil {
		return nil, err
	}
	streamID, err := utils.ReadVarInt(r)
	if err != nil {
		return nil, err
	}
	return &StreamsBlockedFrame{StreamID: protocol.StreamID(streamID)}, nil
}

func (f *StreamsBlockedFrame) Write(b *bytes.Buffer, _ protocol.VersionNumber) error {
	typeByte := uint8(0x0a)
	b.WriteByte(typeByte)
	utils.WriteVarInt(b, uint64(f.StreamID))
	return nil
}

// Length of a written frame
func (f *StreamsBlockedFrame) Length(_ protocol.VersionNumber) protocol.ByteCount {
	return 1 + utils.VarIntLen(uint64(f.StreamID))
}
