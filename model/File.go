package model

import (
	"bytes"
	"os"
	"time"
)

type MyFileInfo struct {
	NameInfo string
	Data     []byte
}

func (mif MyFileInfo) Name() string       { return mif.NameInfo }
func (mif MyFileInfo) Size() int64        { return int64(len(mif.Data)) }
func (mif MyFileInfo) Mode() os.FileMode  { return 0444 }        // Read for all
func (mif MyFileInfo) ModTime() time.Time { return time.Time{} } // Return whatever you want
func (mif MyFileInfo) IsDir() bool        { return false }
func (mif MyFileInfo) Sys() interface{}   { return nil }

type MyFile struct {
	*bytes.Reader
	Mif MyFileInfo
}

func (mf *MyFile) Close() error {
	return nil // Noop, nothing to do
}

func (mf *MyFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil // We are not a directory but a single file
}

func (mf *MyFile) Stat() (os.FileInfo, error) {
	return mf.Mif, nil
}
