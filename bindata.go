// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// public/index.html
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\x51\x6f\xdb\x36\x10\x7e\x26\x7f\xc5\x8d\xc3\x06\x7b\x71\x24\x37\xd8\xc3\x20\x4b\x1e\xd6\x24\x03\x0a\x2c\x48\xb0\xf4\x65\x68\xfb\x40\x8b\x67\x8b\xad\x44\xaa\xc7\x53\x13\x77\xc8\x7f\x1f\x48\xc9\xa9\x53\xb4\x01\xf6\x90\x58\x3a\x7e\xf7\x7d\xdf\x7d\x47\x95\x3f\x5c\x5c\x9f\xbf\xfe\xe7\xe6\x12\x1a\xee\xda\xb5\x94\x25\x5b\x6e\x71\x7d\xdb\x60\xdb\xee\xcb\x7c\x7c\x93\x65\x6b\xdd\x07\x20\x6c\x2b\x15\x78\xdf\x62\x68\x10\x59\x41\x43\xb8\xad\x54\xc3\xdc\x87\x22\xcf\xb7\xde\x71\xc8\x76\xde\xef\x5a\xd4\xbd\x0d\x59\xed\xbb\xbc\x0e\xe1\xec\xf7\xad\xee\x6c\xbb\xaf\x5e\xbd\xbc\x3a\xb9\x69\xf1\xfe\xe4\xca\x3b\xff\xb3\xb1\xa1\x6f\xf5\xbe\x0a\x77\xba\x57\x51\x39\x31\xaf\xe5\x2f\xf0\xaf\x14\x1b\x7f\x7f\x1a\xec\x67\xeb\x76\x05\x6c\x3c\x19\xa4\xd3\x8d\xbf\x5f\xc9\x07\x19\x7d\x46\x44\x54\x3b\x1d\x89\x0b\x50\xaf\x5e\x5e\x41\xa4\x86\x48\xad\x16\xd0\x79\xe7\x43\xaf\x6b\x8c\x2d\x1b\x6f\xf6\xb1\xa5\xd3\xb4\xb3\xae\x80\x65\x2c\x4a\xeb\xfa\x81\xdf\xf0\xbe\xc7\x8a\xf1\x9e\xdf\x2d\x20\xfe\x68\x42\x1d\xb1\xbd\x36\x26\xa9\x2f\xb3\x33\xc2\x0e\x96\xd9\xaf\x84\xdd\xea\x09\x89\x18\x9d\x4d\xcf\xba\xfe\xb0\x23\x3f\x38\x53\x80\xf3\x0e\x57\xf2\x2b\x8f\xd6\x35\x48\x96\x57\x53\x39\xd8\xcf\x78\x54\x94\xe2\xce\x1a\x6e\x0a\x78\xb1\x5c\xfe\x94\xfc\xfd\xe8\x07\xee\x07\x7e\xce\xcc\xff\x4a\xa1\xcc\xa7\x80\x65\xd9\x69\xeb\xd6\x52\x94\x3d\x21\x58\x53\xa9\x51\x49\xad\xcb\xbc\x27\x8c\x07\x5b\x4f\x5d\x3a\xa9\x7d\xd7\x69\x67\x14\x74\xc8\x8d\x37\x95\xba\xb9\xbe\x7d\xad\x40\xd7\x6c\xbd\xab\x54\xae\x7b\x9b\x7f\x7a\x91\x87\xa1\x47\xca\x03\xd6\x03\x61\x1e\xb4\x33\x1b\x7f\xaf\xd6\x52\x88\x32\xa5\x0c\x29\x65\x15\xf3\x55\xe0\x74\x87\x95\xaa\x3b\xa3\xa0\x6f\x75\x8d\x8d\x6f\x0d\xd2\x17\xa9\xd4\xf6\xb8\x8a\x11\x1d\xd8\x58\xf7\x15\x7e\xaa\x91\xbf\x0b\x95\xfa\x2d\x9a\x3f\x34\xc5\x09\xf2\x38\xc2\x5a\x96\xf9\x38\xab\x2c\x43\x4d\xb6\xe7\xb5\xac\xbd\x0b\x0c\x53\xb6\x15\x18\x5f\x0f\x1d\x3a\xce\x3e\x0e\x48\xfb\x5b\x6c\xb1\x66\x4f\x33\x35\xa5\xaf\xe6\xab\xa9\x63\xb2\xf7\x5c\xcb\x61\x82\xf9\x4a\xca\xe9\x39\xd3\xc6\x5c\x7e\x42\xc7\x7f\xd9\xc0\xe8\x90\x66\x2a\x0c\x9b\xce\xb2\x5a\x00\x42\xb5\x8e\xdb\xc5\xac\x27\x8c\x98\x0b\xdc\xea\xa1\xe5\x59\xec\x17\x79\x0e\x7f\xa3\x36\x90\x56\xd1\x93\xef\x91\xd8\x62\x90\x62\xb4\x13\xcb\x7f\xa4\x2d\x40\x05\x8f\x62\xa9\xb0\x3a\xc6\x5c\xa5\xbd\x1d\x61\xc6\x45\x3e\xc1\x5c\x68\xd6\x50\x81\xc3\x3b\xf8\x73\x7a\x9d\x4d\xf0\x83\x95\x1b\xa4\x64\x84\xf0\xe3\x80\x81\xa5\xd8\x22\xd7\xcd\xec\x8b\x8b\x45\x9c\x44\x0c\xd4\x16\x47\xd6\x16\x52\x88\x51\xaf\x38\x32\x13\xab\xf1\x93\x2c\x1e\xc5\x17\x52\x3c\xcc\x33\x6e\xd0\xcd\x08\x43\x3f\xe5\x92\x74\x35\x05\x04\xa3\x59\x4b\x21\xec\x16\xd2\x79\x16\x58\xf3\x10\xa0\xaa\x2a\x38\x5b\x2e\xe7\x09\x2c\x08\x79\x20\x07\x09\xf0\x3e\x78\x17\x63\x14\xe2\x41\x3e\x3d\x89\x77\x64\xf6\x0d\x2d\xc1\x0d\xf9\xbb\x94\xc1\x25\x91\xa7\x74\x38\x32\xc4\xff\xdf\xb1\x77\x3d\x5e\xa3\x58\xf5\x2e\xa0\x14\x62\xbc\x36\x49\xe6\xdc\x3b\x46\xc7\x70\x52\x81\x7a\xeb\xa6\x40\x0b\x78\xa3\xe0\x64\xf4\x72\x3e\xad\xe4\xbd\xb7\x6e\xa6\x40\xcd\xe1\x04\xd4\x3b\xb5\x7a\x8e\x26\xb0\xf1\x03\x17\x6f\x5d\x64\xd1\xec\x37\x63\x22\xb7\xa9\x9c\xfc\x3e\xa6\x74\xcb\x06\x89\xb2\x16\xdd\x8e\x1b\x58\xc3\x21\xa8\xe7\xb8\x91\xe8\x5b\xdc\x48\x74\x48\xf3\x61\x9e\xd5\x3a\xee\x1e\x89\x0e\x49\x7c\x97\x31\x12\x45\x0f\x1d\x86\xa0\x77\x98\x82\x5c\xc9\xf8\x57\xe6\x87\x0f\xf2\xbf\x00\x00\x00\xff\xff\xb8\x07\x1d\xeb\x83\x06\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 1667, mode: os.FileMode(420), modTime: time.Unix(1600052923, 0)}
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
	"index.html": indexHtml,
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
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
