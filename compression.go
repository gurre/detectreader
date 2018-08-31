package detectreader

import (
	"bufio"
	"bytes"
	"compress/bzip2"
	"compress/gzip"
	"io"

	"xi2.org/x/xz"
)

// Compression represents the supported compress readers
type Compression int

const (
	// Uncompressed is the same reader you passed in
	Uncompressed Compression = iota
	// Bzip2 is a reader
	Bzip2
	// Gzip is a reader
	Gzip
	// Xz is a reader
	Xz
)

// Extension returns the file extension for the detected type
func (compression *Compression) Extension() string {
	switch *compression {
	case Uncompressed:
		return ""
	case Bzip2:
		return ".bz2"
	case Gzip:
		return ".gz"
	case Xz:
		return ".xz"
	}
	return "[unknown]"
}

// DetectCompression compares for any magic bytes
func DetectCompression(source []byte) Compression {
	for compression, m := range map[Compression][]byte{
		Bzip2: {0x42, 0x5A, 0x68},
		Gzip:  {0x1F, 0x8B, 0x08},
		Xz:    {0xFD, 0x37, 0x7A, 0x58, 0x5A, 0x00},
	} {
		if bytes.Compare(m, source[:len(m)]) == 0 {
			return compression
		}
	}
	return Uncompressed
}

// Decompress takes a reader, detects magic bytes and returns another reader for that compression
func Decompress(stream io.Reader) (io.Reader, error) {
	buf := bufio.NewReader(stream)
	bs, err := buf.Peek(10)
	if err != nil {
		return nil, err
	}

	compression := DetectCompression(bs)
	switch compression {
	case Uncompressed:
		//log.Info("Detected uncompressed file")
		return buf, nil
	case Gzip:
		//log.Info("Detected Gzip compressed file")
		return gzip.NewReader(buf)
	case Bzip2:
		//log.Info("Detected Bzip2 compressed file")
		return bzip2.NewReader(buf), nil
	case Xz:
		//log.Info("Detected Xz compressed file")
		return xz.NewReader(buf, 0)
	default:
		return stream, nil
	}
}
