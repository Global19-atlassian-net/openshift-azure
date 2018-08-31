// Code generated by go-bindata. DO NOT EDIT. @generated
// sources:
// data/azuredeploy.json
// data/master-startup.sh
// data/node-startup.sh
package arm

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _azuredeployJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x5b\xeb\x6f\xdb\x38\x12\xff\xde\xbf\x82\xd0\x15\x48\x5b\xc4\xaf\x3c\xba\x87\x00\xf7\x21\x8f\xb6\x31\x9a\x34\x46\x94\xee\x7e\x58\x04\x05\x4d\x8e\x6d\x6e\x24\x52\x47\x52\x4e\x5d\x9f\xff\xf7\x03\xf5\xb2\x64\x49\xa6\xe4\x78\x77\x7b\x38\xf6\x43\x12\x89\x33\x24\x67\x7e\x33\x9c\x87\xba\x7c\x85\x10\x42\xce\x6b\x45\x66\xe0\x63\xe7\x0c\x39\x33\xad\x03\x75\xd6\xeb\xc5\x4f\xba\x3e\xe6\x78\x0a\x3e\x70\xdd\xc5\x3f\x42\x09\x5d\x22\xfc\xe4\x9d\xea\x1d\xf5\x07\xa7\x9d\xfe\xa0\xd3\x1f\xf4\x28\x04\x9e\x58\x98\x79\x0f\xe0\x07\x1e\xd6\xd0\xfd\x43\x09\xfe\x0f\xe7\x30\x5e\x81\x08\xae\x81\xeb\x5f\x41\x2a\x26\xb8\x59\x68\xd0\xed\x9b\x7f\xe9\x84\x00\x4b\xec\x83\x06\xa9\x9c\x33\xb4\x5c\x25\x4f\xe7\x58\x32\x3c\xf6\xa0\xf0\x50\x82\x12\xa1\x24\xd1\xc3\xdf\xa3\x47\x66\x2c\xb3\xdf\xa2\x49\x7a\x11\x80\x59\xe6\x96\x11\x29\x94\x98\xe8\xee\x5d\x00\x12\x6b\x26\x38\xf6\x86\x5c\xb1\xe9\x4c\xab\xde\xb3\x90\x4f\x2a\xc0\x86\xd7\x61\x91\x9e\x63\x3f\xa2\x5f\x2e\x51\xf7\x52\xf0\x09\x9b\x76\x6f\xc4\x74\xca\xf8\xf4\xb7\x94\x06\xfd\x07\xfd\xa1\xd0\x6a\xb5\x49\x8a\x03\x96\x3b\xe7\x51\x7f\xf0\x4b\xa7\x7f\xdc\x19\x9c\x76\x02\x09\x73\x06\xcf\x9b\xf3\x3d\x41\xa2\x7d\xe5\x96\xd3\x98\x71\x90\x2e\xc8\x39\x23\xd0\xbd\x49\x26\xd4\x2d\x18\x48\x11\x80\xd4\x2c\x16\x53\xe1\x5d\xf4\x5e\x3d\x85\x95\x2f\xa2\x97\x5f\x92\x83\x8e\x40\x7e\xba\x38\xea\x0f\xfe\xe9\x94\xe6\xad\x0e\xcb\x3c\x27\x80\x75\x28\x6b\x56\x8c\x57\x05\x2c\xc9\x6c\x2d\x89\x41\x99\xef\xab\xea\xbf\x56\x87\xaf\x96\xcb\x0e\x62\x13\xc4\x85\x46\x6f\x18\xa7\xf0\xbd\x42\x2c\xa3\xec\xd8\xdd\xf3\x29\x70\x3d\x12\xc2\x1b\x49\x31\x61\x1e\x28\xd4\x7f\xdb\xfd\x95\x83\x76\xc3\x31\x07\x3d\xbc\x42\xab\x55\x73\xa4\x7c\x01\x6d\x80\xd1\x9b\x33\xa9\x43\xec\x25\x7f\x96\x20\x52\xd2\xf3\x7b\xa3\xe7\xe3\xfe\xde\xf5\x9b\x62\x71\xce\x41\xb7\xd5\x3d\xa6\x54\x82\x52\xae\x01\x6c\xbd\xae\x92\x59\x23\x09\x13\xf6\x7d\xc3\xb0\x4a\x93\x07\xfd\xd8\x74\x7b\x15\x50\x31\xe3\xb1\x11\x80\x54\xa4\x99\xfa\xa5\xaa\x77\x8a\xf2\xf2\xa0\x30\xc1\xa1\xb7\x29\x92\xc2\x54\x8b\x78\x0a\x73\x0b\x42\x88\x5c\x54\x7a\xd0\xa3\x93\xea\x93\xa2\x12\x88\xeb\x9f\x3e\x6e\x07\x3b\x70\x6a\x30\xba\x5c\xf6\xde\xa1\x87\xbb\xab\xbb\x33\xf4\x0c\xc8\xc7\x0b\x34\x06\x64\x3c\x20\xd2\x02\xe1\xb9\x60\x14\xe9\x19\xa0\xe1\x08\x61\x4e\xd1\xcd\x05\x9a\x81\x04\x63\x29\xcf\x80\x94\x16\x01\x0a\x15\xe3\x53\xf4\xae\x97\xf2\xba\xc2\xe0\x0b\xee\x82\x56\x68\x22\x24\x72\xaf\xbe\x74\x11\x7a\x98\x01\x47\xb7\x58\x69\x90\x88\x0b\x0a\x0a\xa9\x99\x08\x3d\x8a\x88\xf0\x01\x85\x01\xc2\x0a\xdd\x03\xa6\x8b\x1c\x23\x1c\x6a\xe1\x63\xcd\x08\xf6\xbc\xc5\x61\xb4\xfc\x0c\x38\x01\xb3\x41\xe0\x1a\x24\x50\xc4\xb8\x16\x88\x08\xae\x18\x4d\x7c\x6d\xb4\x28\xce\xb1\xb9\x11\x98\x5e\x60\x0f\x73\x02\x12\x25\x06\x80\x26\x52\x70\x6d\xf6\x6d\xce\x76\x3e\x1a\x22\x05\x72\x0e\x32\x26\x4b\x05\xd5\xd8\x6c\x83\x70\xec\x31\x32\x1c\x9d\xc7\xfa\x2c\xfb\xf6\xbf\xdc\x70\x59\xd0\xc1\x01\x8b\x0f\xd5\xd6\x80\xb3\xd3\x78\xe9\x76\x6e\x41\xcf\x04\x35\x7c\xaf\x16\x1c\xfb\x8c\x54\x18\x80\xc3\xa8\x07\x0f\xcc\x07\x11\xea\x21\xbf\x65\x3c\xd4\xd1\x02\x83\xd3\x8a\xb9\x94\x2b\x17\xb4\x51\xc0\x16\x77\x4e\x85\x8f\x19\x37\xb7\xc5\x0d\x1e\x83\xb7\x71\x37\xc6\x60\xba\xb9\xb8\x34\x13\x62\x0b\xca\x84\x62\xf3\xfb\x1b\x02\xa9\xbb\xad\x32\x69\x5e\x60\xc5\x88\x53\x6f\x4e\xe9\xaf\x8d\x01\xe3\xe5\x30\xb9\x33\x58\x28\x04\xc0\xa9\xba\xe3\x95\x1e\xcd\xf9\x3d\x8d\x5a\x86\xf4\xcd\x41\x03\xcc\x1e\x1c\xa2\x83\x3c\x6c\x0e\xde\x3e\x16\x8f\xfc\xf8\x67\x81\xd5\x1b\xef\x0e\xd6\x31\x26\x4f\xc0\x69\x72\x0a\x73\x21\xbf\xc8\xc1\x27\xec\xaa\x9d\x6f\x85\x8b\xad\x8a\x53\x64\x14\x79\xd2\xe1\x28\x46\x6a\x18\xfb\xa5\x17\x6d\x2b\xe5\xb9\xaf\x8b\x27\x90\x6c\x8e\x35\xb4\xb5\xf1\x22\x8f\x22\x80\xac\x8b\xa2\xd8\x49\x98\x15\xf6\x0e\xcd\xcd\x51\x7d\x45\xd6\xbf\x69\xa6\x59\xc6\xc7\x22\xe4\xf4\x0b\xd6\x19\xce\xb6\x4f\xbb\x0f\xe3\x44\xa2\x72\xda\xda\x07\x30\x3e\xcd\x66\xee\x8a\x90\x40\x48\xdd\x39\x39\x39\xde\x17\x42\xca\x76\xd5\x4a\xc1\x44\x70\x82\xf5\x9b\xed\x7a\x2e\x78\x41\xa3\xe3\xbc\x23\x38\x78\x7b\x88\x0e\x7a\x15\xe6\x9d\x3e\xb3\x83\xc0\x02\xe0\x84\xcf\x48\x48\xed\x9c\xa1\x93\x93\x63\xcb\x7c\xe0\x26\x2c\xfa\xe8\x09\x6c\x2e\xae\xe1\xc8\x39\x43\x13\xec\x29\xb0\x90\xd5\xf8\x83\xbf\x45\x9c\x75\xbe\x29\x7b\xf1\x62\xa1\xa6\x8c\x1a\x4b\xb5\x45\xd8\x50\xa0\x33\xa7\xbd\x62\x4a\x4b\x36\x0e\xd3\x5b\xe8\xca\x1a\x9e\xa3\xc4\x0e\xc6\xf5\xa9\xc9\xc6\xee\xf6\x2a\xff\x68\x65\xd5\x4b\x8d\xf5\xc5\xd2\x0e\xa4\xd0\x82\x44\xb6\xe9\x3c\x90\x60\x0f\x39\x43\x85\xa3\x12\xa1\x6e\xe4\xd0\xe2\xc3\xfd\x4c\x4e\x8c\x99\x2c\x61\x8e\xbd\x21\x77\x81\x08\x4e\x0d\x89\x0d\x57\x3c\xf4\xc7\x20\xef\x26\xa3\xf4\x34\x47\x36\x1d\x34\x45\xfa\xfe\x95\xf5\x93\x44\xb5\xae\x16\x12\x4f\xa1\xa7\xe2\x9f\xe7\x84\x88\x90\x6b\x7b\x5c\x7b\xda\xe9\xbf\xef\x0c\x4e\xff\xb4\x24\x28\x97\x2d\xdc\xc3\xd4\xf8\x8a\x85\x5b\xd8\xe2\xae\xe5\x2d\x1c\x93\x3f\x24\x92\x70\x35\xe6\x14\x4b\xfa\xed\xe6\xde\xfd\x3f\x91\x67\xfc\xe3\xaf\x96\xe6\x06\x63\x8d\x6b\xd2\xc7\x4c\xa6\x24\xda\x66\xad\x4e\xa2\x3a\x88\xc4\x7c\x0a\xe8\x35\x0e\x02\x74\xf6\xaf\xb6\x45\xbf\xd5\x2e\xd9\x1f\x8f\x7f\xba\x40\x42\xc9\xf4\xe2\x93\x14\x61\xb0\xaf\x92\xc1\xeb\xdd\xd5\xcb\xd5\xb4\x63\x38\xe0\x20\xe8\x9a\xa4\x7a\x97\xba\x6f\x72\xa4\x17\x87\xb3\xd8\xf3\xc4\xf3\x37\xa5\x66\x7b\x2b\xb5\x11\x12\x27\x29\x8e\x49\x79\x36\x2b\xe2\xa5\xe9\x14\x14\x91\x2c\x48\x05\x1b\xd1\x20\xd7\xbd\x46\x5a\xe2\xc9\xc4\x9e\x1c\x51\x50\x9a\xf1\x48\xec\xe7\x9b\x45\xbe\x77\x2d\x88\x4d\x18\x75\x6f\x00\x1a\x81\xe0\xa8\x73\x74\x64\x25\x66\x12\x48\xba\xef\x61\x9c\x8a\xd8\x03\x22\x26\x8c\xda\x4c\xdc\xd5\x1f\xb4\xbc\xc6\x2c\xd3\xe3\x88\xa9\xbd\x10\x62\xba\xc2\xf9\xdf\xb5\xbf\x32\x93\xaa\x3e\xfc\x3b\x86\xf5\xbd\xf0\x00\x39\x7e\x54\x3a\x72\x0a\xc6\x9b\x1f\x8d\xf1\x19\xb5\xae\x7e\x26\x84\x5e\x3f\x3c\x8c\xdc\xbf\x15\xa3\x27\x27\xc7\x96\x08\x0e\xed\x05\xa5\xd6\xa8\xec\x7f\x0d\xa5\x49\x39\x7e\xf3\xe5\x96\x22\x7e\xfa\xab\xf5\xde\xb9\x14\x7e\x10\x6a\x48\xbb\x4b\xb7\x98\xcc\x18\x07\x97\x60\x0f\x5c\x68\x10\x57\xfc\xd2\x19\x1c\x75\xfa\x83\xfd\xdf\x3c\xc5\x0a\x66\xae\x03\x17\x19\x6b\x5d\x37\x2d\x23\xb7\x94\x91\x36\x9a\x69\x26\x23\x9b\x73\xd0\x26\xe9\xda\xe8\x80\x6c\xf5\x11\xad\xd7\xb5\xe4\x81\x1b\xab\xb7\xe5\x5e\x19\x3c\x98\x55\x4a\xd7\x77\x22\x74\x6b\x1d\x37\xf5\x69\x4a\x55\xd2\x6f\x2a\xad\x36\xaf\xd0\x0c\x64\x3e\x80\xab\x6a\x12\x10\x1c\x60\x12\xdb\x70\xba\xd6\x65\x3e\x76\xac\x20\xc9\x45\x9f\x31\x2c\x6e\x5d\xf6\x03\xaa\x6b\xfd\xeb\x3e\xae\xd1\x66\x1a\xac\x0e\x7d\x3c\x85\xfb\x44\xaa\xd1\xd1\x9c\x92\x62\x9d\xc0\xc3\xd5\x35\x99\xc2\x06\x0a\x2c\xdd\xcf\x5f\x6b\x84\x84\xb2\x02\xa9\x9a\xc5\x42\x29\x11\x8f\xd2\xb7\x5b\x59\x48\x41\x43\xa2\x2b\x19\xdc\x4d\x26\x39\xe2\x2a\x31\x54\x20\xcc\x1a\xc1\x89\x39\xc8\x40\x8a\x39\x4b\xec\xbf\xa6\xc2\xe5\x84\xc1\x54\x62\x0a\x23\xe1\x31\xb2\xa8\xef\xe1\xf8\x82\xc6\xce\x08\xf3\x10\x7b\xcd\x1a\xfd\x45\x37\x95\x84\xda\xf5\x4b\x08\x65\x9b\x82\xe2\x5e\xab\xcf\xf8\x57\x05\x32\x55\x27\xf1\x44\x48\x3b\xa1\x2a\xf5\x1c\x0a\x64\x24\xf6\x9d\x72\xdd\x65\xca\x63\x31\x6f\x27\x9d\x6d\x6c\x3c\xc6\xc3\xef\xed\xea\x7f\x0e\x65\x0a\x8f\x3d\x18\x61\xa5\x9e\x85\xa4\xe7\xa1\x9e\x01\xd7\x2c\x73\xba\x5a\x86\xb6\xda\xa3\x09\xa0\x1b\xd5\xb9\xe2\x92\xfb\x67\x58\x6c\xef\xc4\xe7\x87\x9d\x6b\xc6\xfd\x09\x16\x57\x58\xe3\x44\x72\xae\x7b\x3d\x4a\x97\x3b\x57\xae\x96\x8c\x4f\xd7\xd8\x76\xdd\xeb\xcf\xb0\xe8\x66\x33\xb6\x98\x47\xfd\x69\xb0\x36\xe7\x76\x7a\x33\xe1\x43\x6f\xad\xe8\x5e\x57\xa9\x59\x0f\x87\x7a\x26\x24\xfb\x01\xf4\xdb\x93\x39\x70\x23\xbe\xf5\x1d\x85\x74\x94\x3f\x42\x68\x46\x5f\x13\x14\x54\x1f\xd7\x49\x6a\x02\x8d\x20\xcf\x62\xbf\x37\x01\x09\x3c\xf9\x16\x63\x47\xef\x58\x62\x2d\x8c\xf3\x69\xe0\x96\x9a\xf4\x90\x5e\xe4\x22\x8b\xc2\x89\xae\xa6\xb6\x8e\xba\xc0\x62\xbe\x8e\x7c\x4a\x6c\x92\xa8\x68\xed\x73\x23\x27\xeb\x29\xb0\x8a\xab\xa2\xed\x55\xe2\x9e\xaa\x20\xba\xd1\xb3\xcb\xfb\x10\x1d\x94\x03\xb9\x48\xb1\xd1\xa5\x5f\xcb\xa5\x14\x01\x6c\x09\x39\xd2\xb1\xa5\xda\xec\x08\x75\xc5\xd4\x93\xdd\x69\x91\xc8\x6b\x4f\xcd\x71\xef\x01\xd3\xdf\x24\xd3\x60\x93\x39\x91\x80\x35\xdc\x65\xb9\xcc\x47\x29\xfc\xe8\x30\x36\xc2\xf8\x5b\x41\xda\x68\x67\x28\x67\x3d\xe7\xc5\x8a\xd3\x48\x82\xcf\x42\xbf\x5c\x70\xda\x1c\xdb\x8c\x78\xa7\x4c\x33\xda\x14\xc5\x1a\x9b\x23\xd8\x5d\x6f\x83\x13\x52\xa6\x9e\x4c\x7c\xf4\xe9\xc2\x39\x43\xc7\x96\x2c\x09\x55\x49\xff\x83\x1f\xe8\x45\x03\x6f\xeb\x78\xa1\x99\xdf\xdf\x51\x62\x8f\x36\x44\xd6\x79\xc0\x24\x00\x6e\xe4\x01\x93\xb9\x43\xae\x41\x4e\x30\x81\x86\x1d\xf9\x74\x34\x90\x77\x56\x3c\xb3\xe6\xdb\xa8\x65\x41\x20\x47\xc3\x7c\x2c\x17\x8d\x6e\xfc\x8c\x28\x6e\x56\x0e\x47\x1f\x85\x7c\xc6\x92\xc6\x26\xd9\x82\xbe\x2a\xcd\x68\xbc\x65\xd4\xb8\xd5\xdf\x3c\x9b\xa9\x48\x64\xea\x86\xa5\x6b\xb6\xde\x61\xd0\x12\x0f\xf9\xd1\x5c\x12\x08\x15\x3e\xcc\x4a\x6a\xd2\xcd\xc3\x19\xb4\x23\x70\x36\xe8\xdb\x83\xa8\xc0\x20\xfe\x7a\x72\xa7\xc5\xd1\x1a\x0e\xcb\x25\x9b\x94\xb3\xfb\xd5\x6a\xb9\x2c\xa7\xfc\xc9\xed\xb5\x5c\x9a\x2b\x76\xb5\x6a\xd6\x84\xad\x4d\xfa\x0f\xd1\x41\x2f\xf9\x02\xb4\x97\x7c\xc6\x79\xf0\xf6\x71\xb9\x04\x4e\xab\xbe\x18\xb3\x8d\x86\x18\x2b\x88\x20\xf9\xae\x25\x48\xca\x4c\xed\x3f\x07\xa8\xe4\xba\x46\x56\x4b\x4c\xad\xf7\xf5\x32\x6c\x65\x7c\x6a\x5b\xf9\x3b\xb1\xb4\x47\xdc\x65\x8a\x9d\x2f\xe1\xba\xe1\xe4\x8b\x39\x17\x2d\xbe\x35\xb3\x8d\x17\x89\xf9\xa7\xf9\xc4\xa6\x6e\xb4\xd7\x9d\x35\x1c\x78\xf9\x52\xcd\x66\x6e\x4f\xe3\xec\x7c\xb6\x45\x3c\x95\xf3\x6b\xc2\x1c\xf8\xae\x81\x9b\x74\xa3\x51\xa0\x93\xcd\xde\x6b\x50\x43\x94\x2d\x04\x47\xbb\x06\x35\x38\xd4\xe2\x6b\x5c\x43\xba\x65\x5c\xc8\x75\xc5\xb9\x45\x90\x12\x48\xa1\x81\x68\xa0\xd6\x8f\x89\x2b\xc9\xe3\x06\x4a\x92\xea\x5d\x60\x05\xef\x4f\x3e\x70\x22\x28\xa0\x37\xae\xc6\x52\x87\x41\xce\x8d\xbc\xad\xff\xba\xb8\x6a\x34\x8d\x41\x0a\x19\xf0\xda\x7e\xcf\xa3\xff\xfa\xf4\x61\xad\xd5\x86\xec\x54\x4e\x10\x4d\xb7\x90\x36\x0c\x2e\x43\xa5\x85\xef\xc6\x42\x69\x41\x7b\x8d\x39\xf5\x40\xe6\x7b\x06\xdd\xbe\x5d\x4a\x7b\x36\xa3\x72\x69\xb1\xae\x73\xb2\xe9\x69\x92\x7a\xb8\x23\x42\x1d\x84\x3a\x16\xdd\xab\xd5\xab\xff\x06\x00\x00\xff\xff\x20\xd8\x4d\xf4\xaa\x36\x00\x00")

func azuredeployJsonBytes() ([]byte, error) {
	return bindataRead(
		_azuredeployJson,
		"azuredeploy.json",
	)
}

func azuredeployJson() (*asset, error) {
	bytes, err := azuredeployJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "azuredeploy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\xfd\x73\x1a\xb9\x92\x3f\xbf\xf9\x2b\xfa\xe1\x7d\x97\xf7\x52\x19\xc0\x8e\x93\x4d\xd8\xf8\x55\x61\x8c\x13\x2a\xc4\xe6\x01\xc9\xde\xde\xde\x96\x4b\xcc\x34\xa0\xf5\x20\xcd\x4a\x1a\x12\xf2\xf1\xbf\x5f\x49\x9a\x0f\x0d\x0c\x18\x13\xef\x5d\xdd\xd5\x79\xab\xb2\x20\xf5\x97\xba\x5b\xad\xee\xd6\x0c\x47\x7f\x6d\x4c\x28\x6b\x4c\x88\x9c\x83\x8f\x9f\x3c\xef\x08\xc6\xd7\x17\xd7\x2d\x68\xa0\x0a\x1a\x21\x93\x0b\x22\xff\xa8\x87\x0d\x2e\xe8\x8c\x32\x3f\x89\xa5\x12\x48\x16\x7e\xc8\x64\x3d\xe0\x6c\x0a\x54\x42\x90\x08\x81\x4c\x45\x2b\x98\x13\x11\x06\x3c\xc4\xf0\x27\xa0\xca\x3b\x82\x58\xf0\x09\x99\x44\x2b\x90\x73\x9e\x44\x21\x7b\xa4\x60\x82\x9e\x37\xea\x0e\x3f\xf4\x3a\xdd\x9b\xf1\x2f\x83\xee\x99\xa5\xec\xd1\x29\xfc\x0a\xfe\x14\x6a\x86\xb1\x5c\x49\x4d\x9d\xce\x1a\x44\xf1\x05\x0d\x7c\x1e\x23\x93\x73\x3a\x55\x3e\xe3\x21\xd6\xe0\xb7\x9f\x40\xcd\x91\x79\x00\x00\x25\x72\xeb\xf0\xde\x94\x7a\x9a\xf8\x5f\x61\x26\x30\x86\xc6\x92\x88\x46\x44\x27\x8d\x90\x07\xb7\x28\xec\x32\xa7\x52\x91\x49\x4e\x70\x71\x3b\x95\xf5\x4f\x53\xa9\xa5\x69\x84\xb8\x6c\x84\x54\xde\x36\xc8\xe7\x44\x60\x43\xa0\xe4\x89\x08\xd0\x8f\x89\x50\xc7\x1e\x00\x06\x73\x0e\x8f\x76\x83\xc1\x06\x57\xd0\xe4\x61\x26\xe2\x3f\x12\xae\x08\x40\x13\x9a\x8f\xe0\x9f\xff\x2c\x84\xf1\x00\xe4\x4a\x2a\x5c\x04\x2a\x02\xa9\x78\x0c\x16\xb3\x2e\x51\x2c\x69\x80\x5a\x4c\x9e\x30\xb5\x4e\xd9\x03\x10\x28\x15\x17\x18\x70\x06\xfe\xb0\x62\x3e\x20\x0a\x2c\x27\x3b\xd4\x08\x09\x2e\x38\xab\xff\x2e\x39\x83\x57\xaf\x1e\x75\xaf\x2f\x1f\x79\x5f\x3c\x80\x5a\xc4\x67\x7e\x28\xe8\x12\x45\xad\x05\xb5\xdf\x79\x22\x18\x89\xc2\x9a\xf7\xcd\xeb\x5e\x5f\xae\x49\x48\x84\x5a\x17\x51\x2b\x3e\xf3\xa6\x80\x33\x49\x43\x14\x30\x25\x81\x02\x35\x27\x6a\x43\xb5\x32\x90\xf4\xb8\x11\x25\xac\xb9\xcd\xa7\xb4\x43\x11\xa1\xa8\xa2\x9c\xed\x40\xff\xc9\xe1\x96\xa8\x44\x20\x48\x25\x88\xc2\xd9\x0a\xa6\x5c\x68\xfd\xd0\xcf\x28\x81\x4e\xbd\x23\x60\x88\x21\x86\x55\xfe\x81\x2a\x08\x77\x7a\xc7\x0e\xf9\xbf\x7e\x05\x25\x12\xdc\xea\x1e\x0e\xe8\x1a\x43\xeb\x18\x21\x4e\x49\x12\x29\x59\xed\x18\x6b\x86\xd7\x78\xdb\xcd\x6e\x66\xb5\x25\x8c\x24\xb5\xf3\xeb\xeb\xf1\x68\x3c\x6c\x0f\x6e\x3a\xd7\x57\x97\xbd\xd7\x37\x57\xed\x77\xdd\x33\xbd\xa5\x7c\xbb\xdf\xfc\x05\x91\x0a\x45\x2d\x63\x5a\x6c\xc4\x1f\xbe\xb8\xfb\xec\x9b\xd9\x87\x9e\x27\x31\x04\x9f\x82\x8f\x50\x93\x47\x17\xdd\xf3\xf7\xaf\x6f\xfa\xd7\xaf\xfb\xdd\x0f\xdd\xfe\xd9\xc9\xfa\xc0\xe9\x51\x0d\xf6\xa2\xaa\xed\x14\x4a\x05\x94\x81\x0a\xe2\x27\x27\x4f\x5f\x34\x7f\x82\x90\x7b\x47\x1b\x13\x3f\xbe\xcc\x21\xcc\x87\x17\xa7\xa7\x4f\xb3\x0f\xa7\xf6\x43\xf3\xd9\x53\x48\xc2\xf4\x83\x1e\x79\xd9\x7c\x69\xc9\xfd\x25\x16\x5c\xf1\xb3\x1f\xbe\x84\x52\xfd\xed\x6f\x4f\x1e\x7f\xf3\xfe\x12\x73\xa1\xec\xc0\xd1\xd1\xe3\x27\xdf\xbc\xbf\xd0\x58\x91\x49\x84\x12\xfc\x36\x5c\x8f\x6e\x2e\x7b\xc3\xee\xcf\xed\x7e\xff\xa6\xdd\xef\x5f\xff\x0c\x7e\x0c\x3f\x18\x22\xe0\x2f\xf4\x2e\x50\x08\xbe\x6f\xff\x7f\xd5\xfd\x59\x0f\x66\xd3\x7e\xa8\x49\xc3\x0f\xe6\x5f\xff\x77\x68\x77\x3a\xdd\xc1\xd8\x0b\x39\x43\xcf\xcb\x98\xf8\x92\x2c\x11\xd6\x35\x9f\xcd\x7a\x9e\x58\x80\x2f\xa6\x56\x87\xda\xb2\x8d\xc7\xf6\xb3\x8d\xa1\x0d\x6b\xbb\xc6\x63\xcf\x9b\x10\x89\xcf\x4f\xc1\x0f\xe1\xd5\xab\x57\xf0\xe5\x0b\x74\x50\xa8\xb6\x3c\x5f\x29\x94\x50\xef\x18\xba\x75\x3d\x46\xa7\x34\x20\x0a\x65\xbd\xab\x82\xb0\x43\xcc\x18\x7c\x85\x73\x83\xdf\x65\x7a\xe3\xc1\xb7\x6f\xa9\x48\x86\x65\x40\xea\x81\x50\x07\x72\x18\xa1\x58\xa2\xd8\x83\x8b\xb4\x80\x95\x9c\x06\x82\x2e\x89\xc2\xb7\xb8\xda\x97\xdf\x5b\x5c\xed\xc5\xee\x16\x57\x07\x2e\x6c\x80\x7b\x2d\x2b\xc6\x07\x58\x94\xe1\x75\xe7\x92\x0c\x2b\xbd\xa0\x4d\x5e\xbf\x90\x45\xf4\x8e\x08\x39\x27\x51\xce\xa5\x1d\x2e\x28\x7b\x9b\x4c\xd0\x3a\xdd\x56\xda\xa9\xab\xe9\x7d\x6a\xfe\xa9\xdf\xe6\x38\x07\xa8\xee\x2e\x8f\x73\xb9\xa5\x8e\xe7\x2d\x6e\x43\x2a\xf4\xce\xab\x70\x7d\x46\x16\x18\xfe\x49\xde\x5f\xe6\x64\xff\x57\xd7\xaa\xf6\x0f\xde\x12\xf7\x63\xb9\x8d\xcd\x9e\xae\xd3\x21\x3b\x9d\x66\x83\xd7\x61\x9b\xe1\x3a\x4b\xbc\x3a\x9c\x49\x1e\xe1\x7d\x16\x68\xac\xd7\x08\x52\xc4\xef\x59\xeb\x86\x14\xfb\xaf\xbc\x2c\xc4\x61\x4a\xb8\x14\x9c\xa9\x81\xe0\x9f\x56\xf7\xb3\xf0\x54\xe3\xf9\xb1\x46\x3c\xdc\xa9\x46\x36\xff\x1a\xd1\x19\xa3\x6c\x76\x3f\x01\xd2\xdc\xcd\x97\x74\xc6\xbe\x33\x52\x6d\x88\xb1\xbf\x09\xd6\xa4\x38\x3c\x2a\x77\x22\x8a\x4c\x1d\xbc\xad\x2d\xf6\xf7\x86\xeb\x54\x88\xfd\x97\x5f\x21\xc3\x61\x2a\x68\xcf\x66\x02\x67\x44\x71\x51\x38\xe4\x7d\x94\x41\x72\x7c\xdf\xf1\xcc\xef\x52\x48\xa5\x48\xfb\xab\x66\x8b\x44\x87\xa9\xe7\x9d\xa1\xa9\xcf\xbc\x08\xd5\xc1\xae\x72\x6b\xf1\x1f\xc2\x5b\xaa\x04\xba\xb7\xdb\xac\xc9\xf3\x3d\xaa\xb1\x11\xec\x50\xc5\xa4\x61\xec\xa1\xd4\xe2\x0a\x73\x6f\xa5\x94\x64\xf9\x1e\x95\xec\x93\xc0\x56\x4a\xf0\x00\x09\x6d\x49\x82\x7b\xab\x60\x57\x8a\x3b\x48\x26\x11\x0d\x2a\xf8\xa7\x41\xbc\x1d\x04\xba\xe8\x7c\x8b\xab\x7a\x0e\x7a\xbf\x58\x4e\x2c\x05\x59\x8f\x0d\xfe\x16\x31\xb6\xea\x61\x43\x8e\x43\xb9\x5b\x0e\xd5\xec\x33\x66\x6f\xd4\x80\x48\xf9\x31\xdc\x93\xc7\x5c\xc5\x06\xfc\x4f\xc9\xb3\xb3\xb8\xa7\xb1\x76\xa6\xda\x55\xbc\x8a\x78\x72\x2f\x66\x45\xaf\xcd\x09\x29\x29\x5f\xaf\x68\x1f\x99\x2a\x43\xff\x63\x1b\x81\xaf\x5e\x75\xaf\x2f\xbd\xee\xb8\x73\x71\xd3\xbe\xf8\xd0\x1d\x8e\x7b\xa3\xee\x4d\xa7\xdf\xeb\x5e\x8d\x6f\xde\x0f\xfb\xa3\xb3\xb9\x52\xb1\x6c\x35\x1a\x3f\xfc\x7d\xce\xa5\xd2\x69\xd6\x3f\x5a\xba\x80\xb7\x38\x9d\xee\x70\x7c\x73\xd9\xeb\x77\xcf\x2a\x8b\x40\x0b\x63\xa9\x19\xd0\xf6\xfb\xf1\x9b\x33\xd3\x65\x31\x53\x17\xed\x71\xfb\xe6\xa2\x37\x3c\x2b\x77\x3e\xcc\x5c\xb7\xdf\xed\x8c\x7b\xd7\x57\x37\xe3\xde\xbb\xee\xf5\xfb\xf1\xd9\xc9\xb3\x66\xd3\x4e\xbd\xe9\xb6\x87\xe3\xf3\x6e\x7b\x7c\xd3\xbb\x1a\x77\x87\x1f\xda\xfd\xb3\x7c\xae\x77\xd5\x1b\xf7\xda\x7d\x67\x35\x83\x6e\x77\xb8\x6b\x2d\x2f\xd6\x30\x3b\xfd\xf7\xa3\x71\x77\x78\x66\xd5\xe8\x37\xcd\x5f\x8e\x5b\x1a\x35\xd8\x4f\xdc\xa1\xe3\x4a\xc0\xe3\x4d\xc0\x93\x4a\xc0\x13\x47\x9e\xb7\xdd\x5f\xb6\xa8\x56\xef\x03\x03\xd2\xef\x8d\xc6\xdd\xab\x4a\x7b\x35\xeb\xe6\x3f\xc7\x56\x29\xf0\xa6\x3a\x0a\xd0\x8c\xb5\xe9\x2f\x39\x5a\xb2\xa3\x06\xb3\xca\xe2\x79\x7d\xec\x80\x6d\x37\xba\x99\xaf\x58\x5c\x5e\xfa\x16\x50\xe3\xa1\x36\xc5\xc5\x4d\xa7\xbd\x0e\x9c\xe6\xd9\x06\xf4\x5f\xef\xaf\xc7\xed\x9b\xf3\x76\xe7\x6d\xf7\xea\xe2\xe6\xfc\x97\x71\x77\x74\x76\x7a\xf2\xf2\xf4\xe5\xf3\x1f\x4f\x5e\x3e\xb7\x30\x77\x53\xba\xbe\xf4\xbc\xa0\x5c\x9e\x3a\x05\x6c\xc5\xb8\x39\x96\xb2\x84\xff\x2e\xcc\xf8\x96\x36\x02\xe2\x2b\x91\x48\xd5\xb0\x6d\xe6\x06\x61\xc1\x9c\x0b\xe9\xec\xdc\x94\x58\x12\x87\x44\xa1\x9f\xc1\xbb\xdb\xb7\xea\x90\x48\x9b\x80\xf5\x15\x59\x44\xe9\x86\xd6\x91\x47\x4a\xca\x99\x8d\x29\x2d\x0f\x20\x8e\x92\x19\x75\xbe\x03\xb4\xa3\x8f\x64\x25\x07\x49\x14\xf5\x16\x64\x86\xd2\x8e\x02\x58\x72\x89\x20\x8a\x72\x96\x0d\x02\xdc\x52\x16\xb6\xe0\xc2\x76\x39\xdb\x65\x06\x39\x10\x89\xe9\x07\x14\x7a\xa2\x05\xcb\xe3\x7c\x38\xa4\x92\x4c\x22\x6c\xc1\x94\x44\x12\xcd\xf0\x79\x42\xa3\x30\xa5\x76\x17\xeb\x2d\x54\xad\x44\x25\x42\x8e\x38\x66\xfc\x7a\x89\x42\xd0\xf0\xce\xc5\xdd\xcd\x21\xa7\xe4\xb0\x18\xf0\x70\x20\x50\xa2\xfa\x1e\xea\x3b\x34\x9a\x7b\x46\x9d\xf2\x86\x31\xd2\x80\x47\x34\x58\x1d\xc6\x0e\x3f\x61\x90\x68\xc8\x61\x12\x15\x0a\x01\xf0\x61\x41\x54\x30\x37\xf4\xdb\x8c\x71\x65\xc8\x39\x00\x1a\xe4\x16\x57\x2d\xa0\xc6\x4f\xea\x25\xb1\x42\x64\x2b\x3f\x27\xed\xe0\x00\x2c\x49\x94\x60\x0b\x6a\x7a\xef\xd7\x9c\x19\x1d\x53\x5a\x85\x38\x7e\x88\x8c\x62\xe8\x00\x70\x36\x4c\xef\x62\xd6\xa4\xc8\xae\x68\x5a\x10\xf3\x50\x6e\x99\x9a\x68\x73\xb9\x93\x02\x7f\xc7\x40\xb5\xb2\xf6\x7e\xf6\x27\x6f\x69\x7c\x6d\x38\x45\x46\x8e\x4b\x42\xa3\x44\xe0\x1a\x9c\x35\x92\xa3\xfc\xd4\x3e\x24\x09\xa9\x2a\xb6\x13\x32\xed\xe0\x61\x8a\x5c\x54\x3c\xce\x06\x2c\x52\xe1\x1e\x9b\x72\xbb\xb0\x00\x85\xba\xa4\x7a\x67\xec\x28\xdb\x8c\x18\xb8\xda\x09\xa7\x03\x27\x89\x69\x1f\x97\x18\xc9\x96\xe7\x6b\xc3\xaf\xf9\x01\x49\xd4\xbc\x10\x47\xe0\x1f\x09\x4a\xf5\x06\x49\x88\x22\x15\xc6\x08\xd7\x69\xb7\xa0\xa2\xa5\xe1\x00\xf0\xc5\x82\xb3\x2b\xb2\xc8\xac\xe3\x6f\x11\xca\xb3\x5e\xa7\x04\xb1\x5c\x06\x02\xa7\xf4\x53\x81\xf5\xef\xfe\x10\x17\x5c\xa1\xdf\xd5\x30\xbe\x19\x9d\x09\x9e\xc4\x16\x7c\x13\xee\xb5\x9e\x34\x83\x89\x44\xa1\xdd\x68\x1b\xe4\x7b\x89\xc2\x0b\x38\x53\x82\x47\x11\x3a\x56\xc0\x08\x83\x62\xb7\x44\x3c\xb8\xbd\x32\xde\xb8\x9e\x3e\xf9\x05\xb2\x76\xa5\x34\x25\x35\x19\x2d\x9b\xe9\x04\xdf\x12\xb0\xcd\x8e\x7c\x3f\xe6\xd6\xac\x68\xc9\xa4\xee\x94\xd9\xb1\xa2\x5d\xe2\xb0\x6c\x41\xed\x71\xcd\x0b\xb8\x90\xed\x28\xe2\x1f\x31\xbc\x36\x81\x5f\xb6\xbc\x90\xc9\x62\x35\x13\xca\xc2\x76\x18\x0a\x94\xb2\x05\xd9\x39\xfe\xa2\xf9\xec\x69\x3a\x77\x85\xea\x23\x17\xb7\x2d\x50\x41\x7c\xea\x61\xde\xd7\xc8\x1c\x30\x20\x2d\xa8\xe8\x89\xba\x2b\xd9\xd2\x5b\x71\x56\xb2\xa5\xf3\x01\x90\x88\xc8\x58\xc6\x87\xad\x59\xa4\x46\x1a\x29\x2e\xc8\x0c\x8b\x55\xe9\xcc\x55\x30\x54\x28\xd3\x29\xeb\x38\x2d\x67\xa2\x4e\x79\x15\x60\x39\xec\x69\x9b\x8e\xb4\x4d\xd7\xc8\xb8\xf1\xab\x02\xcc\x25\x62\x22\x5e\x21\xd9\x94\x8b\x05\x51\x2d\xb7\xee\xe8\x15\x10\x97\x66\x16\xbe\x02\xca\x80\xc4\x3a\x53\xb7\xf8\x6e\xdc\xd0\x54\x28\x53\xda\x7b\xa3\x21\xce\xa8\x54\x62\xf5\x26\xd5\x49\x2b\xbd\x8c\xf5\x45\x3a\x51\x4f\xef\x14\xeb\x72\x19\xb4\x9e\x35\x9b\x4d\xcf\x46\x23\x5b\x22\xa4\x81\xe8\xd6\xed\x3c\xb8\x86\xdd\x6e\xcc\x8a\xee\xc7\xa6\x3d\x2b\x5a\x12\x00\x31\x17\xaa\x05\xc7\xcd\x93\x67\x4d\xaf\xd0\xbe\x2b\x8f\xe6\x4e\x62\x6a\x0b\xde\xb6\x98\x25\x0b\x64\xd9\x01\x1f\x44\x3c\x09\xd3\x74\x25\xdb\xb2\x6e\x5a\x63\xe6\x63\xc1\x97\x34\xd4\xe5\xd3\xe7\x44\xa0\x29\x53\x1c\xe4\x6c\x36\x8f\x3c\x1a\xc8\x7c\x16\x09\x53\x74\x81\x6b\xe4\x25\x2a\x45\xd9\x4c\xd6\x6f\x5f\x68\xa7\x69\x2c\x8f\x49\x14\xcf\xc9\xf1\x59\x1e\xe4\xa5\xb5\xba\x3f\x21\xc1\x2d\xb2\x30\x43\xd4\x9e\xf9\xb4\x04\xb0\xc0\x90\x12\x5f\xad\x62\xcc\x99\xc7\x71\xa4\xcb\x7c\xca\x59\x63\xc9\xc2\xba\xe3\x9f\xe6\x02\x71\x92\x68\xd1\x8b\x6d\xfd\xdf\xa9\x8e\x20\x4a\x4c\x1c\x93\xb6\x95\xeb\x6b\x27\xf0\xa7\xda\xc0\x15\x9c\xca\x77\x15\x55\xe8\xb7\xb8\xda\x03\xdb\x3a\x89\xfd\xde\x1b\xb4\xe0\xf8\xe4\x47\x13\x94\x8e\xef\x3e\xff\xb6\xf5\x9f\x4a\x41\x73\x5b\x63\x08\x40\x06\x73\x0c\x93\x3c\xd4\x5b\xf0\xaa\xde\x42\x06\x67\x1e\x9d\x28\x22\xbb\x1c\x25\x13\x86\xda\xb7\x7f\x3c\xa9\x3f\x35\x81\xb4\x71\xfc\xdc\xb3\x58\x56\x6a\x63\xb5\x3c\x76\xf4\x39\x8f\xb5\xcb\x74\xd2\x33\x91\x31\x7b\xb0\xac\x65\x9c\x24\x08\x30\xd6\xd3\x0a\x99\x1a\xaf\x62\x94\xad\x7d\xdc\xe6\x89\x0b\x93\x4a\x0a\x30\x49\x84\x54\x2d\x78\xde\x6c\x7a\x69\xfa\x97\x51\xdd\x8b\xa8\x41\xfa\x23\x96\x2d\x78\x6a\x28\x6c\xac\xe5\x6d\x32\xc9\x82\xdd\xc6\x81\xe8\xf6\x13\xec\x88\x6d\x2a\xbd\x1f\xf6\x4d\x3c\x8c\x05\x65\x0a\x6a\x59\xa4\xaf\x99\x00\xa9\x08\x65\xb6\xf9\x45\x03\xac\x0f\x04\x8f\x51\x28\x8a\x32\x6d\x48\x65\x81\x0f\xbe\xc2\x1f\x09\x57\x26\x62\x32\x7b\x64\x15\xc1\x24\xf5\xc4\xf4\x28\x4b\x0f\x94\x80\x86\x42\xc7\xa1\xfa\xf1\xc9\x0b\x6b\xab\x53\xb3\x3a\x7d\xbc\x58\x4b\xf6\x91\xcd\xd4\xbc\x05\x2f\x3d\x93\x83\x98\x80\xdb\x1b\xa4\x54\x3a\xbd\x8b\x61\x4a\x29\x3d\x35\x1b\x5a\x21\x29\xef\x81\x29\x95\x6c\x5e\x20\x30\x9c\x13\xe5\xd4\x68\x7c\x29\x7d\x69\x38\x14\xce\xe3\x50\x5d\x77\x20\x5e\xce\xba\x88\x94\xa8\x1e\x50\x71\xb5\xec\xce\xab\x51\x73\x95\xa8\xf3\x28\xc2\x94\x5b\xed\x2d\x50\xcd\x79\xd8\x02\x92\x28\x7d\xe8\xd1\x10\x99\xa2\x6a\x35\x48\x03\x48\xaa\x8d\x88\xcf\x28\x73\x32\xe0\x05\x89\x63\xca\x66\xef\x52\xe4\x20\x22\x74\xe1\x15\x39\x7c\x5b\x87\x1b\x68\x5f\x98\xa1\x72\x2c\xda\x52\x86\x18\x0a\x4e\x56\x8f\x0b\x42\x23\xb7\x16\x31\x03\xf9\x77\x1a\xba\x73\x32\x99\x78\xa5\x22\xc2\x99\xd3\xdf\xf3\xaf\xb1\xc0\x29\x0a\x81\xe1\xfb\x34\x4f\x74\x21\x13\x46\xff\x48\xf0\xc6\x41\xb0\x91\xa4\x77\x61\xec\xf1\x77\xca\x42\xfc\xb4\xdb\x0a\xed\x44\xcd\x07\x82\xeb\x80\x58\xef\xad\x6b\x12\x9a\xff\xa8\x67\x5f\xea\x69\xc0\xbb\x28\x5b\xa7\x60\x3a\xc2\x40\xa0\xfa\x13\x18\x5b\xc2\x9b\x6c\x6d\x06\x71\x1d\x23\xeb\x5d\xac\x53\x48\x41\xb2\xbc\x2d\x35\x63\xa2\xe6\x5c\xd0\xcf\x58\xe5\xad\xc6\x5f\xea\x0b\x1a\x08\x2e\xf9\x54\x71\x16\x51\xa6\x0f\xa9\x45\xea\xc7\x3a\x43\x1a\x23\x23\x46\x05\xb5\x86\xd9\x0d\x27\x8d\x9c\x64\x6d\x53\x3e\x00\xc5\x6f\x91\x3d\x1c\x33\x43\x6e\x8d\x91\x0f\xc1\x9c\x44\x11\xb2\x99\x5b\xee\xdd\xd3\xf9\xfb\x3c\x20\x11\x98\x36\x32\x17\xe1\xfe\x5b\x60\xba\xed\x6c\xca\x7b\xd2\xae\xa5\xde\x8c\x6d\x5f\x7b\x90\xf2\xa9\xb0\x59\x7a\x48\xb5\x9d\x94\xef\xcf\x8c\xd0\x19\xf5\xc3\xe8\x5e\xfe\xeb\xe2\xaa\x4c\x4d\xe2\x5a\xab\x2a\x1f\x7a\x47\x3e\xb5\x67\x38\xd2\xc7\x4e\xa8\x4f\xad\xec\xe0\x4b\xa7\x6d\x84\x96\x92\xb9\x83\xd6\xef\xe5\xf6\x04\xc0\x82\xf9\xd2\xc2\x99\xae\x99\x97\x3a\x9d\x2b\x82\x3e\xb8\xa5\x1c\xeb\xe1\x35\x31\x5e\x3c\x3f\x4d\xe5\xc8\xfd\xb8\x0a\xec\x59\xb3\xe9\xc5\x82\xff\x8e\x81\x13\x87\xd3\x84\xfe\x8a\x87\x38\x32\xa5\x28\x17\x2d\x30\x4f\xf2\x09\xf3\xe0\x82\x5b\xe0\x34\x02\xbe\x88\x13\x85\x59\xaa\x2a\x31\x48\x04\x55\x2b\x5d\x11\x06\xba\xd6\x4e\x03\x7b\x20\xf3\x91\x21\x31\x1e\x2d\x9b\xad\xc6\x49\x36\xd9\x27\x13\x8c\xe4\xc0\xdc\xc5\xd9\x3e\xc8\x33\x5b\x42\xd3\x70\x1d\xef\xb8\x99\xfd\xf9\xc7\x2f\xb3\xbf\x86\x19\xf5\x04\x4f\x74\x22\x5d\x2c\x45\x26\x93\x90\x2f\x08\x65\x7b\x07\xaf\x21\x4f\xec\x95\xa0\xde\x01\x36\x5a\x19\x27\x1b\x65\x94\x5c\xbf\x90\xa5\xcb\xa3\x82\xed\x82\x30\x32\xc3\x30\x6f\x3f\xf8\x99\x4e\xcd\x67\xd3\xfb\x31\x9b\x42\x8f\xc7\x11\x5f\x6d\xdb\x21\x71\x7e\x6d\x55\xaa\xc7\x2b\x2f\x9d\x00\xe2\xec\x02\x4d\x03\xa7\x7c\x77\x5c\x92\x49\xdb\x27\xc8\xf2\xdc\xca\x32\xfd\xf4\xb4\xba\x4a\xaf\x48\x8a\x9d\x2b\x15\xb7\x4b\x93\xaf\x65\x3d\x3f\x76\xee\x09\xf4\xda\x3f\x0d\x6d\xaf\x47\xf6\xd8\x65\x44\x67\x73\x65\x9d\x33\xef\x01\x8d\xe9\x02\x79\xa2\xd6\xf7\x99\x79\xa6\xc6\xbd\xcb\x34\xfe\xe6\x3b\xe2\xed\xf5\x34\x50\xb9\xed\xb1\xd7\xb3\x3b\xf9\x01\x9f\x1f\x43\x7e\x56\x72\x1f\x10\xaf\x96\x3c\x4a\x16\x4e\x11\x1f\xae\x18\x59\xd0\xc0\x04\x50\x1d\x0a\x28\x9b\x75\x4b\xdd\x3b\x7b\x17\xb0\xa5\xe7\x5e\x15\x3f\xf2\x27\xb0\xd7\xe2\xbd\x8d\xe0\xa3\x52\x60\xf2\x52\xc4\x96\xe7\x9b\xf8\xa1\xa3\xb9\x4d\xdb\x4b\x6d\x85\x12\x92\x3e\xf5\xd7\xaf\x01\x4b\x21\x14\x59\x20\x56\xf1\x6e\x22\x5d\x16\xec\xa0\xb1\x7b\xcd\xa5\xa2\xa9\xfc\xbc\x79\xb1\xe4\x5a\x0b\x6a\xcb\xe3\xda\x13\x3d\xaa\x57\xae\xbf\xdb\xd6\x87\x1d\x8b\x05\x86\xd6\x91\x6a\x2d\xf8\xd5\x18\xf6\x4b\x6a\xde\x9a\xb6\x99\x86\xbf\xe2\x1f\x8c\xb5\xfe\x83\x33\x63\xb1\x88\x06\xca\xb6\x8e\xbf\x3d\xa9\xc6\x78\x47\x3e\x75\xcf\x47\x1f\x52\x1b\x27\xec\x6e\xf0\xd7\x9d\xee\xe0\xe2\x3e\x08\x26\xcb\xbd\xa0\xf2\xf6\x1e\x48\x2a\x98\xf7\x98\x8e\x76\x3c\x6c\x4f\xa7\x94\x51\xb5\xda\x8d\x72\xc5\x35\x87\xfd\xd6\xfc\x1a\x19\x0a\x12\x0d\x0a\x85\xee\x04\x1f\xf0\x70\xcc\x23\x14\x1a\x52\x9f\x3c\x63\x42\x99\xba\x03\xa7\x33\xc7\xe0\x56\x03\xbf\xc3\x05\x17\xab\x81\x8e\x5d\x89\xc0\x3d\x91\xf4\x52\xee\x81\x62\xf5\x7a\x4e\x59\x48\xd9\xac\x1a\x9e\xa4\xad\x93\x5a\x2b\x1f\x03\xa8\x65\x87\x44\xa6\x61\x77\x12\xa0\x16\x99\xa3\x2f\xf7\xb6\x7c\x5c\xe0\x4c\xfb\xab\x33\xf8\x5b\xfe\x39\xcb\x43\x53\x01\x1c\x51\x87\x0e\x96\x86\xfa\x2d\xf5\x6a\xca\x05\xd5\x21\x68\xbb\x57\x67\x47\xfd\x28\x16\x48\xc2\x81\xc5\xb0\xbb\xc2\xc0\x7d\x44\x1d\x95\x6b\x2d\x38\xde\xa9\xab\x75\x87\x3a\x98\x50\x1f\x89\x54\xe9\xa1\x80\x87\xcb\x73\x4e\x22\xc2\x02\x0c\xb3\xcb\x9a\x34\x9d\xd0\x4a\xba\x2f\x29\xed\x34\x03\x53\xb4\xb5\x97\x9c\x86\x03\x1e\xca\x5d\x62\x99\x8c\xe4\x2e\x7a\xdf\xad\x26\xb3\x4f\xd2\x9d\x43\x39\xbb\x37\x9d\xdd\x3e\xcb\x14\xdd\xe9\xb7\x5a\x82\xcf\x9c\x39\x37\x67\xdb\x5d\x53\x87\xcb\x0a\xb1\x4e\x0a\x5f\x4d\xdf\xff\xb1\x6f\x94\x3c\x32\x67\xab\x49\x12\xe0\xf8\xf9\x8b\xfa\xf3\xa7\xf5\xe3\x93\x97\xf5\xe3\xe7\x8f\x2a\x9e\xe0\x16\x28\x79\xb4\xb4\x8d\xc7\xca\xa7\xb8\x4b\x6d\xca\x8a\x43\x64\x5b\x1b\xb3\x38\x47\x8a\x13\xab\xa3\x61\xb3\xd2\x46\x8f\xc1\x57\x18\x29\x41\xd9\x2c\x3f\xa4\xb2\xb7\x93\x28\x5b\xa2\x54\x74\x46\x14\x82\x9a\x23\xf8\xfe\x82\x30\x3a\x45\xa9\xfc\x44\x44\x90\x3e\x90\x07\x31\x11\x64\x81\x0a\x05\x10\x16\x82\x44\x04\x3a\x05\xaa\x60\xa1\x75\xe4\x1d\xc1\x1c\xa3\x18\x12\x09\x44\x01\x89\xa2\x4d\xe9\x8d\x0e\x62\x1e\x4a\xfb\xd4\x8c\x7b\xc9\x5e\x75\xd8\x0f\x78\xe8\x2d\x50\x91\x90\x28\xd2\xf2\xf2\xbb\x4e\xfb\xba\x8f\x51\x7b\x4c\x02\xb4\x57\x19\xbe\x7d\x17\xcb\x93\x31\x06\xad\xb4\x51\x6c\xd2\x9b\x34\xc1\x24\x62\x96\xdf\x6c\x7d\x4d\x8d\x2b\x51\x81\x4f\xd2\x2f\x75\xa8\x78\xa6\x27\x9d\xc3\x4f\x18\x64\x7c\x35\xe9\xc5\x82\x14\xcd\x6d\xf3\xfa\xa0\x9c\xa7\xdf\xfc\xc0\x7c\x30\x77\x15\xa5\xfc\xc1\xf6\xf5\xbb\x2a\x08\xcd\x2d\xc7\x7a\xb1\x6e\x2f\x37\x92\x28\x4a\x6f\xa5\xd3\x27\x0a\x6c\x2d\x4d\x97\xc8\x50\xca\x81\xe0\x93\xbc\xfb\xa2\x65\x2a\x1a\x0b\x25\x99\x20\x6f\xba\x07\x2a\x72\x46\x7c\x3f\x20\xa6\xf1\xec\x8c\x6d\x3c\xbc\x51\x02\xcf\xfa\xdc\x95\x08\xf9\xa3\x2a\x2e\x4a\xd6\xdb\xde\x8e\x91\x27\xa5\x29\x06\xb2\x30\xe6\xfa\x10\x75\x46\xb7\xde\x71\x15\x20\x59\x53\x7d\x8e\x24\x52\xf3\x74\x42\x47\x00\x4a\xa2\x0b\x8c\xc8\x2a\x4f\xc1\x4f\x9f\x39\x8d\x86\xdc\x8a\x59\x05\x68\xda\xc9\x9f\xf2\xa7\x0e\x74\xa9\x42\x23\x9c\xe5\x09\xac\x1e\xb4\x59\xef\x3b\x53\x9a\x64\x66\x37\x2f\xa1\x0d\x88\x9a\xb7\x8a\x05\xa6\x34\x0a\x4e\xe9\x85\x44\x3a\xae\x0f\xac\x6b\x16\xad\x1c\xca\x65\x3a\x6b\x6f\xb4\xad\xd1\xd2\xdb\xc0\x8c\xea\x02\x87\xb2\xd9\x05\x15\x9b\x38\x5a\x5f\x45\x0d\x64\xd9\x58\xf1\xb3\xeb\x43\x2e\x2d\xbb\x6c\xc1\x15\x4b\xa8\x5a\xc0\x56\xcc\x0d\xa1\xd7\x45\xae\x4e\x87\x8b\x58\x40\x62\xba\x47\xe6\xbf\x1e\x0c\x6c\x3e\x62\x65\x29\x3d\x44\xa1\x2b\x7d\xce\x90\xa9\x16\x90\x98\xe6\x61\xa3\xf8\x7c\x60\xd4\x30\xaf\x78\x66\x46\x33\x9b\x39\xfd\x62\x29\xa7\x3b\xc6\xa8\xeb\x6c\xaf\xe7\x8b\x72\xa4\x88\xcf\x22\x5c\x62\x74\x76\x5a\x15\x5f\x8a\x17\x79\xab\x03\x4b\xc7\x5e\x89\x0d\x22\xc2\xf0\x61\x42\x8b\xde\x7d\xaf\x8b\x07\x71\x32\x4b\xdb\x9d\xf6\xb9\x18\x35\xf7\x97\xb6\xf4\x4e\x03\x6a\x30\x47\xad\xeb\x37\xe3\xf1\x60\xb4\xc7\x96\x04\x50\x6b\xe5\xf2\x71\xd3\x71\xa1\x4c\xb1\x7a\xdf\xd0\x7b\x4a\xd9\xd0\x48\xab\x87\x90\x35\x15\x69\xab\xac\x0f\x1d\x48\x4a\x6e\x53\x8a\x02\x25\x17\xda\x37\xa6\x6c\x4d\x2e\xaa\x28\x97\xee\x41\xb7\x71\xf8\x9e\x10\xb3\xb9\xb8\xea\xa5\xed\x43\x64\x73\x39\xbb\x16\x73\x57\x14\x72\x1e\x18\x39\x20\x1a\x59\xce\xe5\xe7\x5c\xfe\x8c\x58\x53\xe6\xf0\x3d\x31\x87\x4a\x85\x6c\xe3\xb9\xd5\xd3\xd3\xd3\xff\x2b\x61\xe9\x74\xc7\x56\xaf\x32\xd7\xff\xef\xe4\xff\x4d\x3b\x99\x4e\xe1\xd7\x5f\xa1\xe6\x24\x88\x35\x38\x3b\x83\x5a\xe9\x99\xf2\x1a\xfc\x56\xfc\x4c\xc6\x8e\xbd\x2f\x57\x2c\x38\x78\xd3\x6b\xe4\xfb\xef\xf6\xcd\x0d\x33\x5a\xb1\xe0\xde\x1b\xa5\x24\xc2\xc3\x7a\xf0\x8d\x5e\x66\xe3\x86\x27\xea\x7f\xd4\xc1\xee\xef\x1b\x53\xea\xbe\x22\x9d\xa5\xa7\x11\x9f\x4d\x04\x0d\x67\xb8\xeb\x10\xc8\x81\x0e\xf6\x86\x82\xcd\x03\xb8\x44\x9f\xcf\xce\x0d\xb1\x03\xfd\xc2\x15\xe6\x61\x9d\xc3\xfc\xc2\x43\xc9\x2f\x8a\x91\x32\xe4\x03\x85\xaa\x8a\x60\xb8\x5e\x21\xdd\xab\xca\xe2\xb3\x12\xfa\x92\x08\xbf\x18\x7b\x30\x3f\xde\x74\xbe\x4d\x85\x3d\xfc\x26\xd8\x4d\xd3\x5b\x57\xda\x6e\xe9\x53\xa5\x94\xd5\x64\x42\x70\xf6\xcb\x27\xc2\xfc\xf4\xc9\xdf\xe1\xb1\xed\x82\xb5\xe0\x1f\xf5\xc7\x47\xff\x79\x5c\x51\xc3\x67\xbf\x7e\xb2\xff\xdb\x56\xdf\xcd\xa3\xf2\x67\x1a\x8e\xe0\x4d\xbb\xf3\x56\x27\x00\xf1\x0a\xd6\x26\x41\x71\x98\x70\xae\xa4\x12\x24\x76\xc7\x25\xb7\x3f\xdb\x93\x8b\x9b\xfe\xe8\x8f\xc6\x87\x44\x0b\x4a\x99\x69\x9e\x99\x68\x7c\x04\xf6\x57\x85\x24\x2a\xf8\x48\xa3\x08\x18\x57\x30\x25\x34\x4a\x5b\x67\xca\x80\xda\x05\x5b\x12\xb6\xa0\x85\x80\x0b\x81\x81\x8a\x56\xf5\xca\x17\x61\xd6\xa5\xdd\x00\xa8\x92\xdd\xf3\x8e\x34\x7f\x6c\x41\xd5\xaf\xcd\x40\x20\x88\x9c\x43\xc4\x79\x2c\x21\x61\x8a\x46\x99\x5c\x54\x42\x12\x7b\xc5\xaf\x1c\xd9\x07\xf3\x2b\x89\xe4\x3f\x7a\xb4\xfe\x9b\x48\xbb\x80\xe1\xdf\xdc\x48\x2d\x38\x57\x0d\x23\xf5\xfa\xca\xb7\xbc\x03\xe8\xa2\x34\xd2\x95\xfe\x57\x00\x00\x00\xff\xff\x0c\x21\x8d\xcc\xcf\x4b\x00\x00")

func masterStartupShBytes() ([]byte, error) {
	return bindataRead(
		_masterStartupSh,
		"master-startup.sh",
	)
}

func masterStartupSh() (*asset, error) {
	bytes, err := masterStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x7f\x6f\xdb\x36\x10\xfd\x9f\x9f\xe2\xea\x0c\xcb\x56\x80\x52\xd2\x15\xc1\xa6\xa6\x05\x92\xd4\x2d\x8a\xa6\xb1\x91\x66\x05\x8a\x61\x08\x68\xf2\x6c\xb3\xa6\x48\xf5\x48\xb9\xd1\x5c\x7f\xf7\x81\xa2\xe3\x9f\x69\x92\x61\xff\xd9\xc7\x77\xef\x1e\x1f\x8f\x47\xed\x3d\xc9\x07\xda\xe6\x03\xe1\xc7\xc0\xf1\x86\xb1\x3d\xb8\xea\xbd\xee\x15\x90\x63\x90\xb9\xb2\xbe\x14\xfe\x6b\xa6\x72\x47\x7a\xa4\x2d\xaf\x2b\x1f\x08\x45\xc9\x95\xf5\x99\x74\x76\x08\xda\x83\xac\x89\xd0\x06\xd3\xc0\x58\x90\x92\x4e\xa1\x7a\x01\x3a\xb0\x3d\xa8\xc8\x0d\xc4\xc0\x34\xe0\xc7\xae\x36\xca\xee\x07\x18\x20\x63\x1f\xbb\x97\x9f\xde\x9d\x75\xaf\xaf\x3e\xf7\xbb\x2f\x13\x33\xd3\x43\xf8\x0b\xf8\x10\x3a\x6d\x61\xdf\xf8\xc8\xae\x47\xb9\x08\xae\xd4\x92\xbb\x0a\xad\x1f\xeb\x61\xe0\xd6\x29\xec\xc0\xdf\x2f\x20\x8c\xd1\x32\x00\x80\x0d\xba\x6d\x3c\x1b\x6a\x16\xc9\x9f\xc0\x88\xb0\x82\x7c\x2a\x28\x37\x7a\x90\x2b\x27\x27\x48\x69\x9b\x43\x1f\xc4\x60\x49\x58\x4e\x86\x3e\xbb\x19\xfa\xa8\x26\x57\x38\xcd\x95\xf6\x93\x5c\xfc\x53\x13\xe6\x84\xde\xd5\x24\x91\x57\x82\xc2\x21\x03\x40\x39\x76\xb0\x7f\x3f\x0c\x76\xaa\x42\xa4\x87\x11\x55\x5f\x6b\x17\x04\xc0\x01\x1c\xec\xc3\xab\x57\x2b\x31\x0c\xc0\x37\x3e\x60\x29\x83\x01\x1f\x5c\x05\x29\x33\xf3\x48\x53\x2d\x31\xca\x74\xb5\x0d\xdb\xcc\x0c\x80\xd0\x07\x47\x28\x9d\x05\x7e\xb9\xb3\x3e\x9b\x71\xd0\x43\xc0\xaf\x90\x75\x6f\x02\x89\xec\xd2\x19\x84\x8e\xb6\x43\x12\x1d\x98\xcf\x19\x80\x14\x01\x92\x94\x94\x93\x2b\x81\xa5\xb3\xd9\x17\xef\x2c\x1c\x1f\xef\x77\x7b\x6f\xf6\xd9\x8c\x01\x74\x8c\x1b\x71\x45\x7a\x8a\xd4\x29\xa0\xf3\xc5\xd5\x64\x85\x51\x1d\x36\x67\xdd\xde\x9b\xb6\x14\x5a\x95\x48\xd7\x77\x23\x28\x6c\x6f\x27\x1e\x52\xb2\xf2\xb4\xd7\xbb\xfa\x78\x75\x79\xd2\xbf\x3e\xeb\x5d\xbc\x79\xf7\xf6\xfa\xe2\xe4\x43\xf7\x65\x3c\x74\x9e\x3a\x82\xcf\x66\x1b\xda\xe7\xf3\xa5\x75\xab\xae\xf9\x69\xb6\xde\x14\xf3\xb6\x69\x18\xf3\xa8\x80\x6b\xe0\x08\x1d\xbf\xf7\xba\x7b\xfa\xe7\xdb\xeb\xf3\xde\xdb\xf3\xee\xa7\xee\xf9\xcb\x67\xdb\x81\xe7\x7b\x1d\x78\x14\x2b\x95\xc0\x69\x98\xb0\x18\xa4\xca\x9f\xa6\xdf\xa9\xb1\xf3\x52\xf8\x80\x94\x3f\x65\x6c\x20\x3c\x1e\x3d\x07\xae\xe0\xf8\xf8\x18\x66\x33\x38\x6d\x03\x5d\x1b\xaf\x0c\xfc\xf2\x59\x94\xe6\x83\x20\x3f\x16\x06\xb2\xb3\xb6\x62\x76\xe1\x14\x9e\x3a\x17\x7c\x20\x51\xbd\xaf\x07\x98\x94\xfc\x0a\xf3\xf9\xe2\x8c\x16\x55\xa2\x94\x7c\x70\x8b\xcc\x26\x4b\xe8\x43\x55\xcf\x90\xc2\x89\x3f\x6d\x02\xfa\x65\xd5\x18\xd3\x43\x2d\x45\x40\xbf\x29\xa1\x5d\xfa\x41\xf5\xf6\x8c\x96\x12\x2a\xa4\x4c\x52\x78\xa8\x7c\x9f\xf4\x54\x04\x7c\x8f\xcd\x7f\x10\xf1\x1e\x9b\x47\x6b\x98\x60\xc3\xe4\xb8\x74\x0a\x0e\x8e\x0e\x0e\xe0\x71\x19\xbb\xb0\x3b\xad\xfd\xdf\xde\x9e\x89\xfb\x0c\x95\xa2\x75\x50\x56\xbb\x72\xd2\x52\x8a\x57\x13\x9d\x4b\xc1\x03\xd5\x3e\xe4\x69\xee\xe4\xc2\xca\xb1\x23\x9f\xaf\xc6\xe6\x82\xac\xae\x94\x08\xc8\x6f\xf1\xb7\xb7\xce\x8a\x12\xe3\x5d\x44\x82\xc3\xa3\xdf\xb3\xa3\xdf\xb2\xc3\x67\x7f\x64\x87\x47\xfb\x77\xc8\x8a\xc3\xcd\x4c\xdb\xe9\xcf\xca\x89\xd2\x04\x7c\x53\xa1\x34\xae\x56\x15\xb9\xa9\x56\x48\x8c\xad\xe6\xc9\x5d\xeb\x69\x62\xa6\xb7\x64\x39\x5c\x66\x2b\xbb\x22\xb6\xbf\xc0\xc6\x18\x7c\x87\x8f\x81\xb4\x1d\xc5\xb9\x12\xc7\xcc\x9a\x86\xdb\x59\x67\xdc\x68\x40\x5a\x8d\x70\xb7\x76\xbb\x83\xca\x29\xbf\x02\x65\x8d\x28\xcd\xb2\xb6\xa8\xf4\x27\x24\xaf\x9d\x2d\x60\x7a\xc8\x26\xda\xaa\x02\xfa\x4e\xb1\x12\x83\x50\x22\x88\x82\x01\x44\xbb\x0a\x58\x95\x49\x11\x5f\x09\x89\x05\xc4\x06\xe1\x69\xde\x31\x5f\xa1\x8c\x09\xd2\xd9\x20\xb4\x45\xf2\xf1\x1f\x07\x5d\x8a\x11\x16\xb0\xb6\xd1\x73\x37\x3a\x6d\xc9\xde\xc5\x25\xf8\x0e\xf1\x69\xc0\x34\x3c\x21\xe1\xfb\xb5\x31\x7d\x67\xb4\x6c\x0a\x38\x31\xdf\x44\xe3\xdb\xb5\x5d\x31\x00\x1e\x65\x4d\x3a\x34\x67\xce\x06\xbc\x09\x45\x1b\x04\xa8\x48\x4f\xb5\xc1\x11\xaa\x02\x02\xd5\x09\x3b\x75\xa6\x2e\xf1\x43\x7c\x4e\x7c\x02\xf2\xf4\xb8\xf4\x45\x18\x17\x90\xfb\x20\x02\x2e\x08\x52\xad\x55\x64\x13\xb9\x79\xf0\xeb\x19\x69\x08\xf2\x16\xc0\xb7\x10\x84\x42\xf5\xac\x69\xd6\x24\x6d\xd2\x62\x90\x1b\x64\xab\xff\x0f\xa6\xb6\x1d\xe1\x46\x1b\xe9\x53\x41\x7c\x15\xdb\xa1\x18\x3b\x1f\x2e\x30\x7c\x73\x34\x59\xc6\x92\x45\x8b\xb3\x8b\x80\x96\xfe\xd6\xd4\xb5\x52\x1b\xcd\xb7\x6b\xd8\x0f\x93\x7f\x78\x7b\x1e\xb6\xf0\x3e\x4e\xb6\x6d\xda\xfd\xea\x17\xa6\x6c\xda\xd4\xde\xb1\x3d\xb0\x2e\x60\x01\x77\xbd\x80\x20\x29\x7e\x3d\x1a\xe7\x2a\x0f\xb5\x0d\xda\x2c\xb4\xc6\x6f\xc3\xba\x62\xab\xa7\x1f\xad\x18\x18\xbc\x93\x64\xf9\x25\xb0\xfd\xa1\x70\x1f\x18\x7e\x66\xff\x06\x00\x00\xff\xff\xae\x0f\xee\xfa\xc0\x0a\x00\x00")

func nodeStartupShBytes() ([]byte, error) {
	return bindataRead(
		_nodeStartupSh,
		"node-startup.sh",
	)
}

func nodeStartupSh() (*asset, error) {
	bytes, err := nodeStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"azuredeploy.json":  azuredeployJson,
	"master-startup.sh": masterStartupSh,
	"node-startup.sh":   nodeStartupSh,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"azuredeploy.json":  {azuredeployJson, map[string]*bintree{}},
	"master-startup.sh": {masterStartupSh, map[string]*bintree{}},
	"node-startup.sh":   {nodeStartupSh, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
