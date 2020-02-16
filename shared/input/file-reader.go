package input

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type FileReader struct {
	f *os.File
	s *bufio.Scanner
}

func NewFileReader(fn string) (fr *FileReader, err error) {
	f, err := os.Open(fn)
	if nil != err {
		return
	}

	s := bufio.NewScanner(f)
	return &FileReader{f, s}, nil
}

func (fr *FileReader) ReadLine() (s string, err error) {
	if nil == fr || nil == fr.s {
		return s, errors.New("Invalid receiver to FileReader.Read")
	}

	if !fr.s.Scan() {
		if nil == fr.s.Err() {
			err = io.EOF
			fr.f.Close()
		}
		fr.f = nil
		fr.s = nil
		return
	}

	return fr.s.Text(), nil
}

func (fr *FileReader) Scan() bool {
	if nil == fr || nil == fr.s {
		return false
	}
	b := fr.s.Scan()
	if !b {
		if nil == fr.s.Err() {
			fr.f.Close()
		}
		fr.f = nil
		fr.s = nil
	}
	return b
}

func (fr *FileReader) Text() string {
	if nil == fr || nil == fr.s {
		return ""
	}
	return fr.s.Text()
}
