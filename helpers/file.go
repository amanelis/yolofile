package helpers

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

// File interface
type File interface {
	readBinary(string) ([]byte, error)

	GetBody() []byte
	GetPath() string

	GetMD5() string
	GetSHA() string
}

type file struct {
	Body []byte
	Path string

	MD5 [16]byte
	SHA [32]byte
}

// NewFile returns a new object of File interface
func NewFile(path string) (File, error)  {
	if !FileExists(path) {
		return nil, fmt.Errorf("%s", RFgB("invalid or missing file"))
	}

	//  Create our new instance of file
	f := &file{
		Path: path,
	}

	data, err := f.readBinary(path)
	if err != nil {
		return nil, fmt.Errorf("%s", RFgB("error reading file"))
	}

	f.Body = data
	f.MD5 = md5.Sum(data)
	f.SHA = sha256.Sum256(data)

	return f, nil
}

func (f *file) GetBody() []byte {
	return f.Body
}
func (f *file) GetPath() string {
	return f.Path
}

func (f *file) GetMD5() string {
	return hex.EncodeToString(f.MD5[:])
}
func (f *file) GetSHA() string {
	return fmt.Sprintf("%x", f.SHA[:])
}

// ReadBinary ...
func (f *file) readBinary(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	size := stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// FileExists ...
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// GetYolofile returns the valid and present yolofile
func GetYolofile() string {
	path, _ := os.Getwd()
	yolo := fmt.Sprintf("%s/Yolofile", path)

	return yolo
}

// YolofileExists checks current working directory for a Yolofile
func YolofileExists() bool {
	path, _ := os.Getwd()

	yolo := fmt.Sprintf("%s/Yolofile", path)

	if _, err := os.Stat(yolo); err != nil {
    if os.IsNotExist(err) {
    	return false
    }
  }

  return true
}
