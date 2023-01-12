package DBReader

import "io"

type DBReader interface {
	Parse(reader io.Reader) error
	Convert() ([]byte, error)
	CreateAnotherExt(data []byte) error
}
