package dumb_broker

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

var (
	Enc = binary.BigEndian
)

const (
	numberOfBytesInRecord = 8
)

type store struct {
	File   *os.File
	mu     sync.Mutex
	buffer *bufio.Writer
	size   uint64
}

func NewStore(f *os.File) (*store, error) {
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	size := uint64(fi.Size())
	return &store{
		File:   f,
		size:   size,
		buffer: bufio.NewWriter(f),
	}, nil
}

func (s *store) Name() string {
	return s.File.Name()
}

func (s *store) Append(p []byte) (n uint64, pos uint64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	pos = s.size
	if err := binary.Write(s.buffer, Enc, uint64(len(p))); err != nil {
		return 0, 0, err
	}
	w, err := s.buffer.Write(p)
	if err != nil {
		return 0, 0, err
	}
	w += numberOfBytesInRecord
	s.size += uint64(w)
	return uint64(w), pos, nil
}

func (s *store) Read(position uint64) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.buffer.Flush(); err != nil {
		return nil, err
	}

	size := make([]byte, numberOfBytesInRecord)
	if _, err := s.File.ReadAt(size, int64(position)); err != nil {
		return nil, err
	}

	b := make([]byte, Enc.Uint64(size))
	if _, err := s.File.ReadAt(b, int64(position+numberOfBytesInRecord)); err != nil {
		return nil, err
	}

	return b, nil
}

func (s *store) ReadAt(p []byte, off int64) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.buffer.Flush(); err != nil {
		return 0, err
	}
	return s.File.ReadAt(p, off)
}

func (s *store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := s.buffer.Flush()
	if err != nil {
		return err
	}
	return s.File.Close()
}
