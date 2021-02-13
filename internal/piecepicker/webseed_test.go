package piecepicker

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uhthomas/rain/internal/piece"
)

func TestPickLastPieceOfSmallestGap(t *testing.T) {
	pieces := make([]piece.Piece, numPieces)
	for i := range pieces {
		pieces[i] = newPiece(i)
	}
	pieces[1].Done = true
	peer := newPeer(0)
	pp := New(pieces, 2, nil)
	assert.Nil(t, pp.pickLastPieceOfSmallestGap(peer))
}
