package piecewriter

import (
	"crypto/sha1"

	"github.com/rcrowley/go-metrics"
	"github.com/uhthomas/rain/internal/bufferpool"
	"github.com/uhthomas/rain/internal/piece"
	"github.com/uhthomas/rain/internal/semaphore"
)

// PieceWriter writes the data in the buffer to disk.
type PieceWriter struct {
	Piece  *piece.Piece
	Source interface{}
	Buffer bufferpool.Buffer

	HashOK bool
	Error  error
}

// New returns new PieceWriter for a given piece.
func New(p *piece.Piece, source interface{}, buf bufferpool.Buffer) *PieceWriter {
	return &PieceWriter{
		Piece:  p,
		Source: source,
		Buffer: buf,
	}
}

// Run checks the hash, then writes the data in the buffer to the disk.
func (w *PieceWriter) Run(resultC chan *PieceWriter, closeC chan struct{}, writesPerSecond, writeBytesPerSecond metrics.Meter, sem *semaphore.Semaphore) {
	w.HashOK = w.Piece.VerifyHash(w.Buffer.Data, sha1.New())
	if w.HashOK {
		writesPerSecond.Mark(1)
		writeBytesPerSecond.Mark(int64(len(w.Buffer.Data)))
		sem.Wait()
		_, w.Error = w.Piece.Data.Write(w.Buffer.Data)
		sem.Signal()
	}
	select {
	case resultC <- w:
	case <-closeC:
	}
}
