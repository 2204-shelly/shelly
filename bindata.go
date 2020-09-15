// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// public/index.html
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xc1\x6e\xdb\x46\x10\x3d\x73\xbf\x62\xb2\x45\x0b\xa9\x96\x48\xc5\xc8\xa1\xa0\x48\x16\x4d\xec\x02\x01\x6a\xd8\xa8\x73\x29\x92\x1c\x56\xdc\x91\xb8\x09\xb9\xcb\xcc\x0e\x23\x2b\x85\xff\xbd\xd8\x25\xe5\xc8\x41\x62\xa0\x07\x89\xe4\xec\x9b\xf7\xde\xbc\x21\x8b\x67\x17\xd7\xaf\xde\xfc\x73\x73\x09\x0d\x77\x6d\x25\x44\xc1\x86\x5b\xac\x6e\x1b\x6c\xdb\x43\x91\x8d\x4f\xa2\x78\xb6\x5c\xc2\x8d\xdb\x23\xa1\x86\xcd\x21\x87\x86\xb9\xf7\x79\x96\xed\x0c\x37\xc3\x26\xad\x5d\x97\x9d\x9f\xaf\x5e\x2c\x7d\xec\xcb\xc6\x0b\x2c\x97\x95\x28\x5a\x63\x3f\x02\x61\x5b\x4a\xcf\x87\x16\x7d\x83\xc8\x12\x1a\xc2\x6d\x29\x8f\x34\x5b\x67\xd9\xa7\x3b\xe7\x76\x2d\xaa\xde\xf8\x48\x58\x7b\x7f\xfe\xfb\x56\x75\xa6\x3d\x94\xaf\x5f\x5e\x9d\xdd\xb4\x78\x77\x76\xe5\xac\xfb\x45\x1b\xdf\xb7\xea\x50\xfa\xbd\xea\x65\x30\x1d\x99\x2b\xf1\x2b\xfc\x2b\x92\x8d\xbb\x5b\x7a\xf3\xc5\xd8\x5d\x0e\x1b\x47\x1a\x69\xb9\x71\x77\x6b\x71\x2f\xc2\x88\x01\x11\xd4\x96\x23\x71\x0e\xf2\xf5\xcb\x2b\x08\xd4\x10\xa8\xe5\x02\x3a\x67\x9d\xef\x55\x8d\xa1\x65\xe3\xf4\x21\xb4\x74\x8a\x76\xc6\xe6\xb0\x0a\x45\x61\x6c\x3f\xf0\x5b\x3e\xf4\x58\x32\xde\xf1\xfb\x05\x84\x8b\x22\x54\x01\xdb\x2b\xad\xa3\xfa\x2a\x3d\x27\xec\x60\x95\xbe\x20\xec\xd6\x8f\x48\x92\xd1\xd9\x74\xaf\xea\x8f\x3b\x72\x83\xd5\x39\x58\x67\x71\x2d\xbe\xf1\x68\x6c\x83\x64\x78\x3d\x95\xbd\xf9\x82\x27\x45\x91\xec\x8d\xe6\x26\x87\xe7\xab\xd5\xcf\xd1\xdf\x4f\x6e\xe0\x7e\xe0\xa7\xcc\xfc\xaf\x14\x8a\x6c\x0a\x58\x14\x9d\x32\xb6\x12\x49\xd1\x13\x82\xd1\xa5\x1c\x95\x64\x55\x64\x3d\x61\x38\xd8\x3a\xea\xe2\x49\xed\xba\x4e\x59\x2d\xa1\x43\x6e\x9c\x2e\xe5\xcd\xf5\xed\x1b\x09\xaa\x66\xe3\x6c\x29\x33\xd5\x9b\xec\xf3\xf3\xcc\x0f\x3d\x52\xe6\xb1\x1e\x08\x33\xaf\xac\xde\xb8\x3b\x59\x89\x24\x29\x62\xca\x10\x53\x96\x21\x5f\x09\x56\x75\x58\xca\xba\xd3\x12\xfa\x56\xd5\xd8\xb8\x56\x23\x7d\x95\x8a\x6d\x0f\xab\x18\xd1\x9e\xb5\xb1\xdf\xe0\xa7\x1a\xb9\xbd\x2f\xe5\x6f\xc1\xfc\xb1\x29\x4c\x90\x85\x11\x2a\x51\x64\xe3\xac\xa2\xf0\x35\x99\x9e\x2b\x51\x3b\xeb\x19\xa6\x6c\x4b\xd0\xae\x1e\x3a\xb4\x9c\x7e\x1a\x90\x0e\xb7\xd8\x62\xcd\x8e\x66\x72\x4a\x5f\xce\xd7\x53\xc7\x64\xef\xa9\x96\xe3\x04\xf3\xb5\x10\xd3\x7d\xaa\xb4\xbe\xfc\x8c\x96\xff\x32\x9e\xd1\x22\xcd\xa4\x1f\x36\x9d\x61\xb9\x00\x84\xb2\x0a\xdb\xc5\xb4\x27\x0c\x98\x0b\xdc\xaa\xa1\xe5\x59\xe8\x4f\xb2\x0c\xfe\x46\xa5\x21\xae\xa2\x27\xd7\x23\xb1\x41\x2f\x92\xd1\x4e\x28\xff\x11\xb7\x00\x25\x3c\x88\xc5\xc2\xfa\x14\x73\x15\xf7\x76\x82\x19\x17\xf9\x08\x73\xa1\x58\x41\x09\x16\xf7\xf0\xe7\xf4\x38\x9b\xe0\x47\x2b\x37\x48\xd1\x08\xe1\xa7\x01\x3d\x8b\x64\x8b\x5c\x37\xb3\xaf\x2e\x16\x61\x92\x64\xa0\x36\x3f\xb1\xb6\x10\x49\x32\xea\xe5\x27\x66\x42\x35\x7c\x92\xf9\x83\xf8\x42\x24\xf7\xf3\x94\x1b\xb4\x33\x42\xdf\x4f\xb9\x44\x5d\x45\x1e\x41\x2b\x56\x22\x49\xcc\x16\xe2\x79\xea\x59\xf1\xe0\xa1\x2c\x4b\x38\x5f\xad\xe6\x11\x9c\x10\xf2\x40\x16\x22\xe0\x83\x77\x36\xc4\x98\x24\xf7\xe2\xf1\x49\x78\x47\x66\xdf\xd1\x4a\xb8\x21\xb7\x8f\x19\x5c\x12\x39\x8a\x87\x23\x43\xf8\xff\x81\xbd\xeb\xf1\x35\x0a\x55\x67\x3d\x8a\x24\x19\x5f\x9b\x28\xf3\xca\x59\x46\xcb\x70\x56\x82\x7c\x67\xa7\x40\x73\x78\x2b\xe1\x6c\xf4\xf2\x6a\x5a\xc9\x07\x67\xec\x4c\x82\x9c\xc3\x19\xc8\xf7\x72\xfd\x14\x8d\x67\xed\x06\xce\xdf\xd9\xc0\xa2\xd8\x6d\xc6\x44\x6e\x63\x39\xfa\x7d\x48\xe9\x96\x35\x12\xa5\x2d\xda\x1d\x37\x50\xc1\x31\xa8\xa7\xb8\x91\xe8\x7b\xdc\x48\x74\x4c\xf3\x7e\x9e\xd6\x2a\xec\x1e\x89\x8e\x49\xfc\x90\x31\x10\x05\x0f\x1d\x7a\xaf\x76\x18\x83\x5c\x8b\xf0\x2b\xb2\xe3\x07\xf9\x5f\x00\x00\x00\xff\xff\x28\x3b\x2e\x9f\xbe\x06\x00\x00")

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

	info := bindataFileInfo{name: "index.html", size: 1726, mode: os.FileMode(420), modTime: time.Unix(1600053776, 0)}
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
