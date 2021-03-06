// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// tpl/class_option.gohtml
package main

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

var _tplClass_optionGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x51\x41\x6a\xc3\x30\x10\x3c\x4b\xaf\x18\x42\x29\x71\x29\x7a\x40\xc1\xd7\x9e\x4a\xda\x83\x69\x8f\x45\x38\x4a\x6a\xea\x4a\x42\x56\x08\x61\xd9\xbf\x17\x49\xb6\x12\x93\x9b\x98\x9d\xd9\x9d\x19\x79\xdd\xff\xea\xa3\x01\x91\xfa\x28\x4f\xb5\xd3\x7f\x86\x59\x4a\x22\x3c\xf4\xa3\x9e\x26\xbc\xb4\x50\xcc\x57\xe0\x53\x87\x84\x75\xee\xcd\x9d\x4d\x98\xc1\x95\xce\xf9\xd8\x25\x8a\x0f\x83\x8d\xc8\x23\x6c\xde\x7d\x1c\x9c\xdd\x20\x91\xe2\xc5\xa7\xa3\x99\xc8\x8c\xc3\xc9\xf6\xdb\x27\xa2\x79\x49\x23\x65\x42\xb0\x33\xe7\x8a\x6d\x9d\x8f\x13\x94\x52\x55\xd5\x60\x4b\x54\x1d\x31\xe3\x66\x01\x48\x8a\xf5\xb0\xc5\x63\x1d\x13\x4b\x71\x70\x01\xdf\xcf\x70\x3e\x26\xa3\x41\xdb\xa3\x41\x3e\x41\x52\x08\xe7\xe3\x7a\x77\x23\x05\x4b\x11\x4c\x3c\x05\x2b\x4b\xc8\x22\x51\xaf\x83\x19\xf7\xd3\x5c\xcf\x98\x0b\xb9\xe9\xa6\x44\x67\x2e\x79\xbe\x86\xf8\xb3\xac\x5d\xac\xd4\x7c\x44\x45\xce\x9c\x7e\xa3\xbb\x78\xb3\x64\x59\xda\x59\x09\x73\xc4\x62\xa8\x10\xee\xba\xb8\xa7\xaf\x2b\x51\xf5\x36\x5a\x5c\xaf\xa7\xa4\x39\x8d\xb1\x7b\x30\xff\x07\x00\x00\xff\xff\x8e\x18\x19\xb9\x22\x02\x00\x00")

func tplClass_optionGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_tplClass_optionGohtml,
		"tpl/class_option.gohtml",
	)
}

func tplClass_optionGohtml() (*asset, error) {
	bytes, err := tplClass_optionGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/class_option.gohtml", size: 546, mode: os.FileMode(420), modTime: time.Unix(1557671235, 0)}
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
	"tpl/class_option.gohtml": tplClass_optionGohtml,
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
	"tpl": &bintree{nil, map[string]*bintree{
		"class_option.gohtml": &bintree{tplClass_optionGohtml, map[string]*bintree{}},
	}},
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
