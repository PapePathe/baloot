package dumb_broker

import (
	"encoding/binary"
	"io"
	"os"

	"github.com/tysonmote/gommap"
)

var (
	offsetWidth    uint64 = 4
	postitionWidth uint64 = 8
	entryWidth            = offsetWidth + postitionWidth
)

type Index struct {
	file *os.File
	mmap gommap.MMap
	size uint64
}

func NewIndex(f *os.File, c Config) (*Index, error) {
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}

	idx := &Index{
		file: f,
	}

	idx.size = uint64(fi.Size())
	if err = os.Truncate(
		f.Name(), int64(c.Segment.MaxIndexBytes),
	); err != nil {
		return nil, err
	}

	if idx.mmap, err = gommap.Map(
		idx.file.Fd(),
		gommap.PROT_READ|gommap.PROT_WRITE,
		gommap.MAP_SHARED,
	); err != nil {
		return nil, err
	}
	return idx, nil
}

func (i *Index) Read(offset int64) (out uint32, pos uint64, err error) {
	if i.size == 0 {
		return 0, 0, io.EOF
	}

	if offset == -1 {
		out = uint32((i.size / entryWidth) - 1)
	} else {
		out = uint32(offset)
	}

	pos = uint64(out) * entryWidth
	if i.size < pos+entryWidth {
		return 0, 0, io.EOF
	}
	out = binary.LittleEndian.Uint32(i.mmap[pos : pos+offsetWidth])
	pos = binary.LittleEndian.Uint64(i.mmap[pos+offsetWidth : pos+entryWidth])

	return out, pos, nil
}

func (i *Index) Write(offset uint32, position uint64) error {
	if uint64(len(i.mmap)) < i.size+entryWidth {
		return io.EOF
	}

	binary.LittleEndian.PutUint32(i.mmap[i.size:i.size+offsetWidth], offset)
	binary.LittleEndian.PutUint64(i.mmap[i.size+offsetWidth:i.size+entryWidth], position)

	i.size += uint64(entryWidth)

	return nil
}

func (i *Index) Close() error {
	if err := i.mmap.Sync(gommap.MS_SYNC); err != nil {
		return err
	}

	if err := i.file.Sync(); err != nil {
		return err
	}

	if err := i.file.Truncate(int64(i.size)); err != nil {
		return err
	}

	return i.file.Close()
}

func (i *Index) Name() string {
	return i.file.Name()
}
