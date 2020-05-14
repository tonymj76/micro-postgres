package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __000001_create_users_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2d\x4e\x2d\x2a\x56\x70\x76\x0c\x76\x76\x74\x71\xb5\xe6\x82\xa8\x89\x0c\x70\x55\x28\xca\xcf\x49\xb5\x06\x04\x00\x00\xff\xff\x22\xd6\xd1\x93\x33\x00\x00\x00")

func _000001_create_users_down_sql() ([]byte, error) {
	return bindata_read(
		__000001_create_users_down_sql,
		"000001_create_users.down.sql",
	)
}

var __000001_create_users_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xce\xc1\x6a\x84\x40\x0c\x06\xe0\xbb\x4f\xf1\xdf\x54\xd8\x37\xe8\x69\xba\x9b\xd2\xa1\x3a\xca\x98\xa1\xb5\x17\xb1\x3b\x41\x84\xaa\xcb\xa8\x87\xbe\x7d\xe9\xd0\x4a\x61\x2f\x21\x21\xf9\x92\x9c\x2d\x29\x26\xd0\x1b\x93\x69\x74\x65\x70\x1b\xae\xe1\xeb\xb6\x2d\x0f\x49\xf2\xdb\xe3\xb6\x26\x84\xe5\x53\xa0\x1a\x90\x71\x25\xb2\x74\xd8\x65\xdd\xd2\x13\xd2\x49\xa6\x0f\x09\x3f\x59\xef\xa7\x71\x4e\xf3\x7f\x4e\x3d\x16\x84\x7d\x95\xb0\x22\x4b\x00\x60\xf4\x70\x4e\x5f\x50\x5b\x5d\x2a\xdb\xe2\x85\x5a\x5c\xe8\x49\xb9\x82\x31\xc8\xdc\x85\x7e\xf6\xcb\xd4\xed\xfb\xe8\xb3\xfc\x14\x49\xbc\x1b\x83\xa9\x18\xc6\x15\xc5\x21\xfe\x9e\x88\x73\xd7\x20\xfd\x26\xbe\xeb\x37\xb0\x2e\xa9\x61\x55\xd6\x78\xd5\xfc\x1c\x4b\xbc\x57\x86\xee\x37\x9c\x9d\xb5\x64\xb8\x3b\x44\x92\x7f\x07\x00\x00\xff\xff\x07\x0d\x84\x9c\x10\x01\x00\x00")

func _000001_create_users_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_create_users_up_sql,
		"000001_create_users.up.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"000001_create_users.down.sql": _000001_create_users_down_sql,
	"000001_create_users.up.sql": _000001_create_users_up_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"000001_create_users.down.sql": &_bintree_t{_000001_create_users_down_sql, map[string]*_bintree_t{
	}},
	"000001_create_users.up.sql": &_bintree_t{_000001_create_users_up_sql, map[string]*_bintree_t{
	}},
}}
