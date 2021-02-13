package peerreader

import (
	"github.com/uhthomas/rain/internal/bufferpool"
	"github.com/uhthomas/rain/internal/peerprotocol"
)

// Piece message that is read from peers.
// Data of the piece is wrapped with a bufferpool.Buffer object.
type Piece struct {
	peerprotocol.PieceMessage
	Buffer bufferpool.Buffer
}
