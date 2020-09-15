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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xdd\x6e\xdb\x46\x13\xbd\xe6\x3e\xc5\x64\x3f\x7c\x85\x54\x4b\xa4\x62\x04\x45\x41\x91\x2a\x9a\xd8\x05\x02\xd4\xb0\x51\xe7\xa6\x88\x73\xb1\xe4\x8e\xc4\x4d\xc8\x5d\x66\x76\x18\x59\x29\xfc\xee\xc5\x2e\x29\xff\x04\x89\x81\x5e\x48\x24\xe7\xe7\xcc\x99\x73\x76\x8b\x17\x67\x97\x6f\xde\xfd\x7d\x75\x0e\x0d\x77\xed\x46\x88\x82\x0d\xb7\xb8\xb9\x6e\xb0\x6d\x0f\x45\x36\x7e\x89\xe2\xc5\x72\x09\x57\x6e\x8f\x84\x1a\xaa\x43\x0e\x0d\x73\xef\xf3\x2c\xdb\x19\x6e\x86\x2a\xad\x5d\x97\x9d\x9e\xae\x5e\x2d\x7d\xec\xcb\xc6\x07\x2c\x97\x1b\x51\xb4\xc6\x7e\x02\xc2\xb6\x94\x9e\x0f\x2d\xfa\x06\x91\x25\x34\x84\xdb\x52\x1e\x61\xb6\xce\xb2\x4f\x77\xce\xed\x5a\x54\xbd\xf1\x11\xb0\xf6\xfe\xf4\xb7\xad\xea\x4c\x7b\x28\xdf\xbe\xbe\x38\xb9\x6a\xf1\xf6\xe4\xc2\x59\xf7\x93\x36\xbe\x6f\xd5\xa1\xf4\x7b\xd5\xcb\x40\x3a\x22\x6f\xc4\xcf\xf0\x8f\x48\x2a\x77\xbb\xf4\xe6\xab\xb1\xbb\x1c\x2a\x47\x1a\x69\x59\xb9\xdb\xb5\xb8\x13\x61\xc5\x50\x11\xa6\x2d\x47\xe0\x1c\xe4\xdb\xd7\x17\x10\xa0\x21\x40\xcb\x05\x74\xce\x3a\xdf\xab\x1a\x43\x4b\xe5\xf4\x21\xb4\x74\x8a\x76\xc6\xe6\xb0\x0a\x41\x61\x6c\x3f\xf0\x7b\x3e\xf4\x58\x32\xde\xf2\x87\x05\x84\x87\x22\x54\xa1\xb6\x57\x5a\xc7\xe9\xab\xf4\x94\xb0\x83\x55\xfa\x8a\xb0\x5b\x3f\x01\x49\x46\x66\xd3\xbb\xaa\x3f\xed\xc8\x0d\x56\xe7\x60\x9d\xc5\xb5\xf8\x86\xa3\xb1\x0d\x92\xe1\xf5\x14\xf6\xe6\x2b\x3e\x0a\x8a\x64\x6f\x34\x37\x39\xbc\x5c\xad\xfe\x1f\xf9\xfd\xcf\x0d\xdc\x0f\xfc\x1c\x99\xff\xa4\x42\x91\x4d\x02\x8b\xa2\x53\xc6\x6e\x44\x52\xf4\x84\x60\x74\x29\xc7\x49\x72\x53\x64\x3d\x61\x48\x6c\x1d\x75\x31\x53\xbb\xae\x53\x56\x4b\xe8\x90\x1b\xa7\x4b\x79\x75\x79\xfd\x4e\x82\xaa\xd9\x38\x5b\xca\x4c\xf5\x26\xfb\xf2\x32\xf3\x43\x8f\x94\x79\xac\x07\xc2\xcc\x2b\xab\x2b\x77\x2b\x37\x22\x49\x8a\xa8\x32\x44\x95\x65\xd0\x57\x82\x55\x1d\x96\xb2\xee\xb4\x84\xbe\x55\x35\x36\xae\xd5\x48\x0f\xa3\x62\xdb\xbd\x15\x63\xb5\x67\x6d\xec\x37\xf5\x53\x8c\xdc\xde\x97\xf2\xd7\x40\xfe\xd8\x14\x36\xc8\xc2\x0a\x1b\x51\x64\xe3\xae\xa2\xf0\x35\x99\x9e\x37\xa2\x76\xd6\x33\x4c\xda\x96\xa0\x5d\x3d\x74\x68\x39\xfd\x3c\x20\x1d\xae\xb1\xc5\x9a\x1d\xcd\xe4\xa4\xbe\x9c\xaf\xa7\x8e\x89\xde\x73\x2d\xc7\x0d\xe6\x6b\x21\xa6\xf7\x54\x69\x7d\xfe\x05\x2d\xff\x69\x3c\xa3\x45\x9a\x49\x3f\x54\x9d\x61\xb9\x00\x84\x72\x13\xdc\xc5\xb4\x27\x0c\x35\x67\xb8\x55\x43\xcb\xb3\xd0\x9f\x64\x19\xfc\x85\x4a\x43\xb4\xa2\x27\xd7\x23\xb1\x41\x2f\x92\x91\x4e\x08\xff\x1e\x5d\x80\x12\xee\x87\xc5\xc0\xfa\x71\xcd\x45\xf4\xed\x51\xcd\x68\xe4\x93\x9a\x33\xc5\x0a\x4a\xb0\xb8\x87\x3f\xa6\xcf\xd9\x54\x7e\xa4\x72\x6e\x6b\xa7\x11\xa2\xe6\xa0\x3c\x54\xca\xe3\x2f\xaf\x20\x1c\xc0\xb1\x21\xf5\xc8\xb3\xc9\x93\x05\x54\xec\xd4\xec\x3e\xb5\x7b\x48\xcd\x53\xc2\xe8\xe2\x4c\xde\xd0\x4d\x28\x95\x37\x56\xce\xe7\xc7\x39\x57\x48\x71\x61\xc2\xcf\x03\x7a\x16\xc9\x16\xb9\x6e\x66\x0f\xdb\x2e\x82\x62\xc9\x40\x6d\xfe\x48\x82\x85\x48\x92\x71\xaf\xfc\xd1\xd2\x21\x1a\xae\x7e\x7e\xbf\xe4\x42\x24\x77\xf3\x94\x1b\xb4\x33\x42\xdf\x4f\xfa\xc7\xb9\x8a\x3c\x82\x56\xac\x44\x92\x98\x2d\xc4\x7c\xea\x59\xf1\xe0\xa1\x2c\x4b\x38\x5d\xad\xe6\xb1\x38\x21\xe4\x81\x2c\xc4\x82\x8f\xde\xd9\x60\x57\x92\xdc\x89\xa7\x99\x70\x16\x67\xdf\x99\x95\x70\x43\x6e\x1f\xb5\x3e\x27\x72\x14\x93\x23\x42\xf8\xff\x01\xbd\xcb\xf1\xb8\x86\xa8\xb3\x1e\x45\x92\x8c\xc7\x33\x8e\x79\xe3\x2c\xa3\x65\x38\x29\x83\x98\x93\x71\x39\xbc\x97\x70\x32\x72\x79\x33\x59\xff\xd1\x19\x3b\x93\x20\xe7\x70\x02\xf2\x83\x5c\x3f\x07\xe3\x59\xbb\x81\xf3\x1b\x1b\x50\x14\xbb\x6a\x54\xe4\x3a\x86\x23\xdf\x7b\x95\xae\x59\x23\x51\xda\xa2\xdd\x71\x03\x1b\x38\x0a\xf5\x1c\x36\x12\x7d\x0f\x1b\x89\x8e\x6a\xde\xcd\xd3\x5a\x05\xef\x91\xe8\xa8\xc4\x0f\x11\x03\x50\xe0\xd0\xa1\xf7\x6a\x87\x51\xc8\xb5\x08\xbf\x22\x3b\x5e\xfc\x7f\x03\x00\x00\xff\xff\x06\x90\x3d\xbc\x26\x07\x00\x00")

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

	info := bindataFileInfo{name: "index.html", size: 1830, mode: os.FileMode(420), modTime: time.Unix(1600160533, 0)}
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
