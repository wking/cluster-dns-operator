// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/dns/cluster-role-binding.yaml (223B)
// assets/dns/cluster-role.yaml (210B)
// assets/dns/configmap.yaml (434B)
// assets/dns/daemonset.yaml (4.667kB)
// assets/dns/metrics/cluster-role-binding.yaml (279B)
// assets/dns/metrics/cluster-role.yaml (246B)
// assets/dns/metrics/role-binding.yaml (293B)
// assets/dns/metrics/role.yaml (284B)
// assets/dns/namespace.yaml (369B)
// assets/dns/service-account.yaml (85B)
// assets/dns/service.yaml (381B)

package manifests

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _assetsDnsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x31\x8e\x83\x40\x0c\x05\xd0\x7e\x4e\xe1\x0b\xc0\x6a\xbb\xd5\x74\x9b\xdc\x80\x48\xe9\xcd\x8c\x09\x0e\x60\xa3\xb1\x87\x22\xa7\x8f\x10\x4a\x45\x3a\x17\xfe\xff\xfd\x89\x25\x47\xb8\xce\xd5\x9c\x4a\xa7\x33\x5d\x58\x32\xcb\x23\xe0\xca\x77\x2a\xc6\x2a\x11\x4a\x8f\xa9\xc5\xea\xa3\x16\x7e\xa1\xb3\x4a\x3b\xfd\x59\xcb\xfa\xb3\xfd\x86\x85\x1c\x33\x3a\xc6\x00\x00\x20\xb8\x50\x04\x5d\x49\x6c\xe4\xc1\x9b\x2c\x16\xac\xf6\x4f\x4a\x6e\x31\x34\x70\x78\x37\x2a\x1b\x27\xfa\x4f\x49\xab\x78\xf8\xc4\xf6\xe7\xe3\xb6\x15\xd3\xa9\xa7\xe8\x4c\x1d\x0d\x3b\x74\x9a\x1d\xbe\xd3\xef\x00\x00\x00\xff\xff\xfa\x62\xe7\x50\xdf\x00\x00\x00")

func assetsDnsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleBindingYaml,
		"assets/dns/cluster-role-binding.yaml",
	)
}

func assetsDnsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role-binding.yaml", size: 223, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd9, 0xf6, 0x2a, 0x3b, 0x84, 0xd7, 0x3e, 0xc4, 0xe1, 0x70, 0x66, 0x31, 0xda, 0xc4, 0x2f, 0x53, 0x27, 0x29, 0x13, 0xfe, 0x80, 0x36, 0xc5, 0xa1, 0x70, 0xdc, 0x2d, 0xef, 0xcf, 0xe0, 0xc4, 0xeb}}
	return a, nil
}

var _assetsDnsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8d\xbd\x6e\xc5\x30\x08\x46\x77\x3f\x05\xba\x7b\x52\x75\xab\xbc\x76\xe8\xde\xa1\x3b\xd7\xa6\x0a\x8a\x03\x16\xe0\x54\xea\xd3\x5f\xe5\x67\x3b\xe7\x08\xf4\xad\x2c\x35\xc3\x67\x1b\x1e\x64\xdf\xda\x28\x61\xe7\x1f\x32\x67\x95\x0c\xf6\xc4\x32\xe3\x88\x45\x8d\xff\x31\x58\x65\x5e\x3f\x7c\x66\x7d\xdb\xdf\xd3\x46\x81\x15\x03\x73\x02\x10\xdc\x28\x83\x76\x12\x5f\xf8\x37\xa6\x2a\x9e\x6c\x34\xf2\x9c\x26\xc0\xce\x5f\xa6\xa3\xfb\x71\x39\xc1\xe3\x91\x00\x8c\x5c\x87\x15\xba\x1b\x49\xed\xca\x12\x7e\x9a\x93\xed\x5c\xe8\x92\xae\xf5\x82\x63\xc3\x3b\x5e\x7d\x27\x7b\xde\xbf\x8d\x3d\x4e\xf8\xc3\x28\x4b\x7a\x05\x00\x00\xff\xff\xcb\xdd\xd7\x2a\xd2\x00\x00\x00")

func assetsDnsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleYaml,
		"assets/dns/cluster-role.yaml",
	)
}

func assetsDnsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role.yaml", size: 210, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x37, 0xb2, 0x0, 0x7d, 0x4a, 0xd9, 0xf, 0x8, 0x44, 0xe7, 0xab, 0x82, 0xe4, 0x50, 0x94, 0xaa, 0x4e, 0xfd, 0xa0, 0x63, 0xba, 0x18, 0xcf, 0xeb, 0xa6, 0xe4, 0x2d, 0x4, 0x35, 0xd5, 0xc7, 0xd}}
	return a, nil
}

var _assetsDnsConfigmapYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x4e\xcd\x4e\xf3\x30\x10\xbc\xe7\x29\x46\xfa\xce\x5f\x4a\x15\x15\x89\x5c\x7b\xe6\xca\x7d\xb1\x27\x8d\x55\xc7\x36\xeb\x75\x11\x02\xde\x1d\xb5\x95\x82\x2a\x3a\xa7\xf9\xd3\xec\x1e\x43\xf2\x23\xf6\x39\x4d\xe1\xf0\x2c\xa5\x93\x12\x5e\xa8\x35\xe4\x34\xe2\xb4\xed\xfe\x21\xc9\x42\x48\xf2\x17\x52\x8b\x38\x42\x94\xa8\x34\x88\x41\x5b\xb2\xb0\xb0\xf3\x62\x32\x76\xc0\x3e\x2b\xa7\x10\x39\xe2\xab\x03\x80\x7e\xdc\x0d\xbb\x01\x9f\x17\x71\x06\x55\xb3\xd6\x55\xce\x94\x68\xf3\x2a\x8f\xed\x95\x9a\x68\xac\x70\xb1\x55\xa3\xf6\x31\x3b\x89\x08\xe9\xbf\x78\xaf\xbd\x68\x11\x84\xf2\x78\x25\xbf\xb3\x67\x94\xec\x2b\x42\xaa\x74\x4d\x79\x93\xb4\x52\x4d\x29\xcb\x8d\x39\x49\x8c\x36\x6b\x6e\x87\xf9\xfe\xfc\xda\xfe\x5e\x59\xd1\xbc\xd0\x66\xb6\x8a\xf1\x69\xbb\x1b\xd6\x60\xca\xfa\x2e\xea\xd1\x63\x43\x73\x1b\x65\xcd\xf1\xd4\xbb\x9c\xa6\x3f\x4f\xc6\xe0\x3e\x50\xf9\xd6\x98\x2c\x48\xbc\x73\xc5\x89\x9b\x89\xe1\x61\x35\x94\x31\x8b\xef\xae\xad\x9f\x00\x00\x00\xff\xff\x08\x52\xc6\xab\xb2\x01\x00\x00")

func assetsDnsConfigmapYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsConfigmapYaml,
		"assets/dns/configmap.yaml",
	)
}

func assetsDnsConfigmapYaml() (*asset, error) {
	bytes, err := assetsDnsConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/configmap.yaml", size: 434, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0x76, 0xcf, 0x25, 0x4e, 0xbd, 0x12, 0xda, 0x28, 0x18, 0x55, 0x70, 0xc2, 0x65, 0xe, 0x94, 0x40, 0x52, 0x25, 0x71, 0x45, 0xb8, 0xdf, 0xfe, 0x66, 0x8a, 0xa6, 0x6c, 0xa6, 0x90, 0xc7, 0x29}}
	return a, nil
}

var _assetsDnsDaemonsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xdf\x53\xe3\x38\xf2\x7f\xe7\xaf\xe8\x75\xa8\x61\xa6\xbe\x78\x08\xdf\x2d\xf6\xf6\x3c\xc3\xde\x72\x10\x6e\xa8\x9b\x40\x8a\x64\xf6\x1e\x28\x6a\x4a\x91\xdb\x71\x1f\xb2\xa4\x95\x64\x67\x52\x0c\xff\xfb\x95\x6c\x9c\xc8\x4e\x2e\x57\x57\xf7\xb2\x3c\x40\xa2\xee\xfe\xa8\x7f\x77\x8b\x27\x92\x69\x02\x57\x0c\x0b\x25\xa7\xe8\x0e\x98\xa6\xdf\xd0\x58\x52\x32\x01\xa6\xb5\x3d\xa9\x4e\x0f\x06\x20\x59\x81\xc7\xf5\x6f\xab\x19\x47\x60\x32\x05\xc1\xe6\x28\x2c\x30\x83\x60\xd1\x01\x73\x60\x4a\xe9\xa8\xc0\x03\xab\x91\x27\x07\x00\x0e\x0b\x2d\x98\x43\xff\x19\xa0\x3d\xad\x3f\xa3\xa9\x88\xe3\x05\xe7\xaa\x94\xee\x96\x15\x98\x40\x2a\xed\x2b\x55\x1b\x52\x86\xdc\xea\x52\x30\x6b\x1b\xa2\x5d\x59\x87\x45\x2c\x55\x8a\x31\x37\xe4\x88\x33\xf1\xca\xcd\x95\x74\x8c\x24\x1a\xdb\xa2\xc7\xb5\xa6\x21\x22\xc0\x00\xa8\x60\x0b\x04\xb2\x7d\x6d\x5b\x8e\x9a\x3e\x29\x85\x98\x28\x41\x7c\x95\xc0\x4d\x76\xab\xdc\xc4\xa0\x45\xe9\xd6\x5c\x0e\x4d\x41\x92\x39\x52\x72\x8c\xd6\x7a\x91\x57\xf6\x6b\x26\xc4\x9c\xf1\xa7\x99\xfa\xac\x16\xf6\x4e\x8e\x8c\x51\x66\x2d\xc7\x55\x51\x30\xef\xea\x07\x88\xb8\x32\x98\x4a\x1b\xc1\xe3\x9a\xcc\xcc\xc2\xd6\xb4\x98\x2b\x99\x45\xc7\x10\x9d\xa0\xe3\x27\xaf\x9c\x27\x97\xca\x60\x46\x02\x43\x91\x4a\x89\xb2\xc0\xb1\x77\xe0\xda\xf2\x8d\xed\x1e\x86\x16\x71\xc3\xb4\xa6\x02\x14\x9e\x7f\xc2\x5c\x9e\x40\x78\x43\xc0\x61\x90\xa5\x77\x52\xac\x12\x70\xa6\xdc\x88\x6a\x65\xba\xf7\xac\xfd\x3e\x51\xc6\x25\x70\xf6\xe3\xd9\x8f\x01\xca\x76\x04\x7c\x5c\x95\x53\x5c\x89\x04\xbe\x5c\x4d\xfe\x7b\xa4\xd8\x71\xbd\x13\x6d\x76\xb9\x07\xed\xcf\xa7\x3b\xd0\x0a\x74\x86\xf8\x6e\xdd\x42\x34\xef\x0b\x92\x68\xed\xc4\xa8\x39\x26\x01\x7f\xee\x9c\xfe\x1b\xba\xf0\x08\x40\x37\x7e\xcd\x91\x09\x97\x77\x29\xb5\x2e\x3f\x0f\x7f\x1e\x76\x8e\x2d\xcf\xd1\xeb\xf3\x69\x36\x9b\x04\x04\x92\xe4\x88\x89\x2b\x14\x6c\x35\x45\xae\x64\x6a\x13\x38\x0d\x45\x35\x1a\x52\xe9\x6e\x9a\x2d\x39\x47\x6b\x67\xb9\x41\x9b\x2b\x91\x26\x70\x1a\x50\x33\x46\xa2\x34\x18\x50\x43\xf7\xf8\x8a\x50\xa5\xdb\x05\x2c\xa8\xc2\x3f\x88\x2b\x7e\x1a\xee\x51\xf9\xec\x7f\x70\xc5\x59\x10\x79\xab\x4a\xc3\xd1\x86\x66\x09\x2a\xc8\xd9\xae\xa1\x05\x16\xca\xac\x12\x38\x3b\xfd\xff\x31\x75\xca\xe8\xf7\x12\x6d\x9f\x9b\xeb\xd2\x3b\x75\x58\xec\xc4\xf8\xd3\x70\x0d\x11\xf4\xb0\xa6\xeb\x79\x85\x44\x85\xe6\x0f\xd3\xd1\x2c\xf2\xb2\xee\xd2\x4a\x3a\xfc\xd6\x09\xbf\x36\x54\x91\xc0\x05\xa6\xbd\x26\xb2\xbf\x67\xe5\xca\x3a\x1b\xfb\x3e\xb7\xa7\x61\xd5\x4c\x81\x13\x50\x56\x70\x7b\x31\x1e\x4d\x47\xf7\xbf\x8d\xee\xeb\xc9\x74\xf9\xf9\xcb\x74\x36\xba\xff\x7a\x75\x37\xbe\xb8\xb9\xdd\x35\xa1\x5a\x71\x94\xd5\xb6\x1a\x1e\xe9\xe6\x72\x34\x0d\x94\x18\xc0\xa5\xef\xdf\xa0\x0c\x34\x03\xd0\xa2\x66\x86\x39\x4c\x41\x90\x75\xa0\xb2\x76\xa4\xd9\x8e\xd4\xed\xdd\x6c\x94\xc0\xb5\x32\x20\xd5\xf2\x18\x50\xda\xd2\x20\xb8\x1c\x2d\xd6\x6a\x19\x14\xcc\x51\x85\xcd\x68\xfd\x00\x99\x32\x80\x8c\xe7\x5d\xc2\x71\x07\x93\x49\x60\x82\x98\x85\x25\xb9\xdc\x63\xf5\xed\xb5\x65\x96\xd1\x37\x58\x92\x10\xc0\x84\x55\x30\x47\x60\x69\x8a\xe9\xfb\x00\xa7\x62\xa2\xc4\x04\xa2\x3a\x47\x62\x83\x0b\xb2\xce\xac\xde\x2b\x8d\xd2\xe6\x94\xb9\xb8\x47\xb0\x15\x8f\xb6\x86\x59\xe0\xba\x93\x39\xc9\x93\x39\xb3\x79\x70\x16\xf3\xe0\xcb\xf7\xd0\x88\x1f\xb6\xd9\xa1\x8e\x51\x5c\x2a\xd0\xa4\xd1\x97\xe6\x41\x58\xe4\x86\x69\x38\x7a\xf2\x26\x1d\xbe\xfd\xa7\x9a\x5b\x88\xf5\xbb\x0f\x80\xdf\xc8\xc1\xf0\x08\x66\xa3\xfb\x71\xc8\x7e\x37\x19\xdd\x4e\x3f\xdd\x5c\xcf\xbe\x8e\x2f\xee\xff\x3e\xba\x3f\x8f\x36\x86\x2d\x50\x62\x1d\xba\x6e\x5d\x45\x81\xf8\xa7\xbb\xe9\x6c\xfa\xf5\xfa\xe6\xf3\xe8\x3c\xda\x24\x5d\xc8\x31\x1b\x8d\x27\x5b\x0c\xef\x5d\xa1\xa3\x50\x8d\x9b\xeb\xe9\xf9\xd1\x31\x1c\xd5\x73\x04\x62\x03\x31\x5b\xe7\x09\x7c\xfc\xf8\x11\xa2\xc3\xe7\x36\xdb\x5e\x3a\x92\x03\x18\xb3\x27\x04\x56\xef\x50\xca\x30\xb3\x02\x5f\x17\x9b\x98\x2b\x91\x36\xf5\x52\x9f\x1f\x59\x60\xce\x19\x9a\x97\x0e\x6d\x18\x66\xae\x21\xce\x20\x8e\x37\xd4\x58\x49\xb1\xf2\x17\x6f\x8c\x7c\x89\xfc\xf7\xb5\x49\x5d\x4d\x96\xb9\xbf\xd7\x57\xf1\x07\x48\x55\xa7\x6d\xa5\xc8\x85\xcf\xe2\xf8\x02\x6c\xc5\xbf\x92\xb6\x1d\xb2\x4f\x66\x5b\x71\x20\xe9\xe1\x5b\xbb\x1f\x7e\x7d\x7c\x89\xb6\xa0\xbc\xc5\xd7\xe8\x78\xde\xfa\x07\x6e\x26\x90\x19\x55\x00\x17\xa5\x75\x68\x7c\x23\x04\xca\x40\xf7\xba\x57\xf3\x43\xda\x9e\xbf\x3d\x7c\x9b\xd2\x02\x7e\x8d\x0e\x9f\x37\xfd\xe0\x25\x82\xff\xb3\xb9\x32\xae\x56\xa1\xe2\x2f\xef\x0f\x9f\xbb\xe5\xf2\x12\xbd\x7b\x77\xd0\x87\xcb\xe0\xe1\x01\xa2\xc3\xbf\x44\x10\xe3\xef\x30\x84\x37\x6f\xbc\xfc\x80\x74\xa3\x3e\xc4\x12\x61\x08\x8f\x8f\x1f\x7c\x2c\x64\x4f\x1a\x5a\x6f\x3c\xbc\xde\x19\x3d\x9e\x47\x87\xcf\xad\x70\x8f\x3b\xa3\xae\x4f\x95\xc4\xae\x3a\x03\xf8\xa2\x53\xe6\x30\x68\x7f\x50\xc7\x90\x32\x58\x22\x2c\xd0\xf9\x62\xa6\x34\xf0\x9c\xed\x01\xfc\x03\x9b\x6e\x20\x95\x83\x72\x0b\x6c\x99\xa3\xf4\x76\x98\x7a\x96\xbc\x2e\x84\x6b\x34\x55\x3a\x3f\x65\x94\x01\xa6\x09\x4a\xc9\x2a\x46\x82\xcd\x49\x90\x5b\xf5\xae\x99\x3a\x26\x10\x50\x3a\x43\xe8\x81\x4a\x91\xfa\x0a\xb5\xce\xa7\x40\x70\x21\x65\x75\x0a\xb7\x37\x90\x85\x14\x05\x3a\x4c\x0f\x76\x05\xe1\x79\xd0\xba\xf3\x3f\xbb\x7e\x00\x7f\x2d\x49\xa4\xc0\x40\xe2\x32\xa8\x8f\x26\x95\x42\x9b\x7d\x1d\xa9\xd2\x00\x2f\xad\x53\xc5\x5a\xe9\x8c\x84\x43\x83\xa9\x37\xbb\x87\xbd\x30\xa8\x21\xae\x20\x1a\xc0\xe1\x73\xbf\xc1\x34\x25\xd4\x29\xa9\x5f\xf6\x14\x55\xa3\xeb\x85\xd6\x28\x53\x68\x3b\xd0\x46\x09\x5f\x38\xdb\xe3\x04\xb6\x6a\xea\x87\xd0\x33\x3b\x6a\xaa\xe1\x27\xed\xd9\xeb\x54\xac\x99\x9b\xa4\x7c\x7c\xd9\x29\x00\x80\x3c\x57\x50\xe7\xeb\x4b\x23\xd4\xfe\xd9\x2e\x1d\xf8\x37\xae\xf8\x65\xcb\xf6\xfe\x25\x75\x9a\xef\x38\xda\xf2\xd1\xec\xee\xea\x2e\xd9\x51\x01\xcc\xa9\xc2\x3f\x02\xc5\x0a\x9c\x02\x56\x29\x4a\x81\xc9\x15\x90\xe4\x4a\x5a\xb2\x0e\xa5\x83\x39\xe6\xac\xa2\x60\x67\x69\x51\xef\x51\x0b\x3f\xc6\x77\x65\x44\xa1\x52\xca\x08\x53\xa8\x9a\x77\xb0\x4f\x44\x89\x98\xf6\xd2\x13\x80\x17\xba\x67\xe6\x56\x0e\x7c\xff\xfe\xda\x81\xf7\xf3\x6d\x5b\xdd\xf2\xfa\xe2\xf0\x55\x6b\xb0\x50\x15\xa6\x1b\x5b\xeb\xac\xe6\x06\x99\xc3\x93\xa6\x7a\xea\xd5\x67\xd3\xe7\x81\x2b\xbd\x02\x9e\x97\xa6\x5b\x24\xbd\x7e\x63\x05\xa2\x86\x9f\x86\xf0\x06\x96\x8c\xba\x39\x5f\x4a\x3f\x92\xb7\x5b\x7b\x27\x78\x3b\x17\xe5\xbd\xcb\x6f\xbb\xfb\xa6\xd2\xb6\x9b\xe6\x15\x66\xac\x14\xed\xed\x7e\x2e\x4f\x51\x20\x77\xca\x6c\x10\x9e\xca\x39\x1a\x89\x7e\xc0\x91\x3a\x51\x36\x01\x41\xb2\xfc\xd6\x10\x5f\xb9\x9a\xfd\x72\xeb\x1f\x01\xbb\x1f\xc3\xcd\xe9\x98\xe9\x24\x58\x27\x6f\x59\xb1\x6f\xa5\x06\x20\x87\x45\xc7\xae\x18\x9e\x70\x95\x40\xfb\x44\xdf\xf1\x0a\xea\x91\xf6\xac\xbb\xfe\xa8\xde\x75\x0f\xfa\x18\x3b\x76\x5f\x00\xb7\xd2\x98\xc0\xf5\x06\xc1\x29\xe1\x77\x1b\x52\x72\xad\xe2\xa0\x3d\x44\x60\x42\x80\x7f\x1a\x3b\x0b\x56\x81\xcb\x99\x83\xab\xdb\xa9\xb7\x96\x89\x25\x5b\xd9\x76\xb8\x82\x92\x35\xaf\x8f\x83\x5d\x2b\xad\xb4\x47\x51\x26\x81\x91\xef\xe9\xf6\xe0\x5f\x01\x00\x00\xff\xff\x1f\x1c\x56\x01\x3b\x12\x00\x00")

func assetsDnsDaemonsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsDaemonsetYaml,
		"assets/dns/daemonset.yaml",
	)
}

func assetsDnsDaemonsetYaml() (*asset, error) {
	bytes, err := assetsDnsDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/daemonset.yaml", size: 4667, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x63, 0x20, 0x66, 0xa6, 0x1c, 0x81, 0x32, 0x84, 0xce, 0xec, 0x68, 0xe6, 0xa3, 0x60, 0x59, 0xe7, 0x6e, 0x8a, 0x70, 0x85, 0x25, 0x5b, 0xd5, 0x68, 0xb5, 0xbb, 0xb4, 0x59, 0x7, 0x2a, 0x20, 0x94}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xb1\x4a\x04\x41\x0c\x86\xfb\x79\x8a\xbc\xc0\xae\xd8\x1d\xd3\xa9\x85\xfd\x09\xf6\xb9\x99\x9c\x1b\x77\x27\x19\x92\xcc\x16\x3e\xbd\x2c\x8a\x08\xe2\xb5\x81\x7c\xdf\xff\xad\x2c\x35\xc3\xd3\x36\x3c\xc8\xce\xba\xd1\x23\x4b\x65\x79\x4b\xd8\xf9\x95\xcc\x59\x25\x83\x5d\xb0\xcc\x38\x62\x51\xe3\x0f\x0c\x56\x99\xd7\x93\xcf\xac\x77\xfb\x7d\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xaa\xf8\xd4\x54\x38\xd4\x0e\x92\x8f\xcb\x3b\x95\xf0\x9c\x26\xf8\xd2\xbd\x90\xed\x5c\xe8\xa1\x14\x1d\x12\x3f\x7f\xdd\xb4\x51\x2c\x34\x7c\x5a\x4f\xfe\x7d\xf6\x8e\x85\x32\x68\x27\xf1\x85\xaf\xf1\x9b\x6c\xba\xd1\x99\xae\x87\xf9\x4f\xc7\x7f\x6b\x00\xb0\xf3\xb3\xe9\xe8\x37\xba\xd2\x67\x00\x00\x00\xff\xff\x5b\x52\x00\xaa\x17\x01\x00\x00")

func assetsDnsMetricsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleBindingYaml,
		"assets/dns/metrics/cluster-role-binding.yaml",
	)
}

func assetsDnsMetricsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role-binding.yaml", size: 279, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x79, 0x95, 0x6f, 0xa4, 0xd5, 0xed, 0x48, 0x27, 0x41, 0x56, 0x5c, 0xea, 0x5c, 0x89, 0xdc, 0xc1, 0x44, 0x91, 0xd4, 0xb, 0x18, 0x85, 0x79, 0x75, 0xaa, 0x6e, 0xb5, 0x98, 0xbe, 0xc6, 0x33, 0x43}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcd\x31\x4b\x34\x41\x0c\x87\xf1\x7e\x3e\x45\xe0\xad\x77\x5f\xec\x64\x5a\x05\x3b\x0b\x05\xfb\xec\xce\xdf\xdb\x70\x3b\xc9\x90\x64\x0e\xf4\xd3\x8b\x70\xb6\x0f\x3f\x78\xfe\xd1\xd3\x39\x23\xe1\xe4\x76\x22\x48\x81\x86\x46\xdb\x17\x0d\xb7\x8e\x3c\x30\x83\xd2\x28\x76\xe7\x01\x7a\x7e\x7d\xa7\x8e\x74\xd9\x83\xa0\x6d\x98\x68\x16\x1e\xf2\x01\x0f\x31\xad\xe4\x1b\xef\x2b\xcf\x3c\xcc\xe5\x9b\x53\x4c\xd7\xeb\x63\xac\x62\xff\x6f\x0f\xe5\x2a\xda\xea\xdf\xf0\xcd\x4e\x94\x8e\xe4\xc6\xc9\xb5\x10\x29\x77\x54\x6a\x1a\x4b\x37\x95\x34\x17\xbd\x14\x9f\x27\xa2\x96\x85\x78\xc8\x8b\xdb\x1c\xf1\x4b\x17\xb2\x01\xe7\x34\x5f\x6d\x40\xe3\x90\xcf\x5c\xc5\x0a\x91\x23\x6c\xfa\x8e\x3b\x6b\x1a\x88\x42\x74\x83\x6f\xf7\x74\x41\x96\x9f\x00\x00\x00\xff\xff\x9f\xa8\x4d\x6c\xf6\x00\x00\x00")

func assetsDnsMetricsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleYaml,
		"assets/dns/metrics/cluster-role.yaml",
	)
}

func assetsDnsMetricsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role.yaml", size: 246, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x64, 0xdb, 0xe0, 0x95, 0x65, 0xae, 0x53, 0x96, 0x3a, 0x5f, 0x5e, 0x8b, 0x69, 0xe2, 0x7d, 0x5, 0xbf, 0x1f, 0x3a, 0xf, 0xff, 0xd0, 0x6b, 0x23, 0x4f, 0xfd, 0x11, 0x7f, 0x57, 0xd4, 0x4a, 0x8b}}
	return a, nil
}

var _assetsDnsMetricsRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\xb1\x4e\xc4\x40\x0c\x04\xd0\x7e\xbf\xc2\x3f\x90\x20\xba\xd3\x76\xd0\xd0\x1f\x12\xbd\x6f\xd7\x97\x98\x64\xed\x95\xed\x4d\xc1\xd7\x23\xa4\x48\x54\x20\x5d\x3b\x9a\xd1\x1b\xec\xfc\x41\xe6\xac\x92\xc1\x6e\x58\x66\x1c\xb1\xaa\xf1\x17\x06\xab\xcc\xdb\xc5\x67\xd6\xa7\xe3\x39\x6d\x2c\x35\xc3\x55\x77\x7a\x65\xa9\x2c\x4b\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xba\x69\xa3\x58\x69\xf8\xb4\x5d\xfc\x8c\xbd\x63\xa1\x0c\xda\x49\x7c\xe5\x7b\x4c\x55\x3c\x99\xee\x74\xa5\xfb\xcf\x14\x3b\xbf\x99\x8e\xfe\x8f\x9f\x00\x7e\xf9\xbf\x34\x1f\xb7\x4f\x2a\xe1\x39\x4d\x67\xfb\x9d\xec\xe0\x42\x2f\xa5\xe8\x90\x78\xf0\x65\x53\xe1\x50\x63\x59\x20\x7d\x07\x00\x00\xff\xff\xb9\xd9\xab\x8d\x25\x01\x00\x00")

func assetsDnsMetricsRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleBindingYaml,
		"assets/dns/metrics/role-binding.yaml",
	)
}

func assetsDnsMetricsRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role-binding.yaml", size: 293, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc, 0x7d, 0xc7, 0x45, 0x33, 0xc4, 0xd8, 0xf, 0x8d, 0x89, 0x8d, 0x6, 0x47, 0xa7, 0xa, 0x6b, 0x17, 0xf5, 0x5f, 0x5a, 0x2f, 0xd8, 0xf9, 0x6, 0x71, 0xaa, 0x78, 0x8d, 0xb5, 0x7a, 0xf6, 0x99}}
	return a, nil
}

var _assetsDnsMetricsRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xb1\x4e\xec\x40\x0c\x45\xfb\xf9\x0a\x6b\x5f\x9d\x7d\xa2\x5b\x4d\x8d\x44\x47\x01\x12\xbd\x77\xe6\x42\xac\x24\xe3\x91\xed\x04\xc1\xd7\xa3\xcd\x06\x89\xca\xf7\x1e\x59\x3e\xfe\x47\x2f\x3a\xc3\xa9\x01\x15\x95\xae\x5f\xd4\x4d\x17\xc4\x88\xd5\x29\x94\xbc\x18\x77\xd0\xe3\xf3\x2b\x2d\x08\x93\xe2\x84\x56\xbb\x4a\x8b\xc4\x5d\xde\x60\x2e\xda\x32\xd9\x95\xcb\x99\xd7\x18\xd5\xe4\x9b\x43\xb4\x9d\xa7\x8b\x9f\x45\xff\x6f\x0f\x69\x92\x56\xf3\x2e\x4a\x0b\x82\x2b\x07\xe7\x44\xd4\x78\x41\xfe\xe3\x1b\xa6\x8b\x1f\xd8\x3b\x17\x64\xd2\x8e\xe6\xa3\xbc\xc7\x50\x9b\x27\x5b\x67\x78\x4e\x03\x71\x97\x27\xd3\xb5\xfb\xed\xca\x40\xa7\x53\x22\x32\xb8\xae\x56\x70\x30\x87\x6d\x52\xe0\x7b\xf9\xfd\xf8\xde\xba\xd6\x5b\xd8\x60\xd7\x63\xf9\x03\xb1\xcf\x59\xfc\x1e\x3e\x39\xca\x98\x7e\x02\x00\x00\xff\xff\x29\x39\xda\x05\x1c\x01\x00\x00")

func assetsDnsMetricsRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleYaml,
		"assets/dns/metrics/role.yaml",
	)
}

func assetsDnsMetricsRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role.yaml", size: 284, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0xf2, 0x4e, 0x40, 0x91, 0xd8, 0x5e, 0x1c, 0x98, 0xb6, 0x2f, 0x11, 0x2a, 0x15, 0x8f, 0xe4, 0x7c, 0xfe, 0xc6, 0x31, 0xf3, 0xb2, 0xa0, 0x38, 0xb2, 0x3f, 0x15, 0x5a, 0x33, 0x12, 0xd2, 0x88}}
	return a, nil
}

var _assetsDnsNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xcd\x4e\xc4\x30\x0c\x84\xef\x79\x8a\x51\x38\x2f\x3f\xd7\xbc\x03\x5c\x90\xb8\xbb\x8d\x97\x35\x4d\xed\x2a\x76\xcb\xeb\xa3\xb2\x15\xac\xb4\xc7\x68\x46\xf3\x7d\xf1\x24\x5a\x0b\xde\x68\x66\x5f\x68\xe4\x44\x8b\x7c\x70\x77\x31\x2d\xd8\x5e\xd2\xcc\x41\x95\x82\x4a\x02\x48\xd5\x82\x42\x4c\x7d\x7f\x02\xb6\xb0\xfa\x45\xce\xf1\x28\xf6\xa4\x56\xf9\xe4\xdc\x78\x0c\xeb\x05\x39\x27\x40\x69\xe6\xf2\x5f\x3b\x55\xf5\x04\x34\x1a\xb8\x1d\x13\x0f\x70\x0e\x6c\xd4\x56\x46\x18\x68\x33\xa9\xa8\xbc\xb0\x56\xd1\x4f\x98\x62\x5a\x07\x06\xd5\x59\x7c\x97\x42\x5c\x28\x8e\x82\xef\xf1\xdf\x38\x68\x11\xbf\xd7\xea\xab\x9e\x1a\x6f\xdc\x0a\xf2\x73\x3e\x98\xd4\x9a\x7d\xdf\x78\xcd\xa6\x12\xd6\x77\x62\x18\x9a\xd9\x84\xb3\x75\xbc\x73\xdf\x64\xe4\xd7\x6b\x0a\x1b\xbe\x78\x0c\x87\xec\x16\xe2\xbf\xbf\xbb\x1e\xed\x8e\x3a\xb6\xd5\x83\xfb\xcd\x70\x41\x8e\xbe\x72\x4e\x3f\x01\x00\x00\xff\xff\x82\x6d\x29\x03\x71\x01\x00\x00")

func assetsDnsNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsNamespaceYaml,
		"assets/dns/namespace.yaml",
	)
}

func assetsDnsNamespaceYaml() (*asset, error) {
	bytes, err := assetsDnsNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/namespace.yaml", size: 369, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0xab, 0x50, 0x84, 0x61, 0x5f, 0x41, 0xf4, 0x17, 0x3b, 0x6, 0x84, 0xc0, 0x5f, 0x4f, 0xbb, 0xd8, 0x1d, 0xae, 0x26, 0x3e, 0x1f, 0x29, 0x2c, 0x84, 0x6d, 0x5e, 0xc1, 0x87, 0x97, 0x5f, 0xc9}}
	return a, nil
}

var _assetsDnsServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\xb1\x09\xc4\x30\x0c\x05\xd0\xde\x53\x68\x81\x2b\xae\x55\x77\x33\x1c\xa4\x17\xf2\x0f\x11\xc1\xb2\xb1\x14\xcf\x1f\x02\xe9\x1e\xbc\xd3\xbc\x32\xfd\x31\x97\x29\x7e\xaa\xfd\xf2\x2c\x32\x6c\xc3\x0c\xeb\xce\xb4\xbe\xa5\x21\xa5\x4a\x0a\x17\x22\x97\x06\xa6\xea\xf1\x3a\x86\x28\x98\xfa\x80\xc7\x61\x7b\x7e\x9e\xba\x03\x00\x00\xff\xff\x8e\x2c\xf1\x2e\x55\x00\x00\x00")

func assetsDnsServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceAccountYaml,
		"assets/dns/service-account.yaml",
	)
}

func assetsDnsServiceAccountYaml() (*asset, error) {
	bytes, err := assetsDnsServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service-account.yaml", size: 85, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x57, 0x12, 0x50, 0x4d, 0x67, 0x2f, 0x1b, 0x74, 0xa0, 0xa4, 0xbb, 0xa7, 0x59, 0xe9, 0x5a, 0xc6, 0xc1, 0x1a, 0xf8, 0x5f, 0xff, 0x5, 0xdb, 0xc, 0x10, 0x8b, 0xc1, 0x0, 0xcc, 0xf, 0x9f, 0x3a}}
	return a, nil
}

var _assetsDnsServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xce\x3d\x4b\x04\x31\x10\xc6\xf1\x3e\x9f\xe2\x81\x6b\x3d\xe1\x10\x0b\xd3\x6a\x63\xb7\xe0\x4b\x3f\x97\x1d\x8e\xe0\xe4\x85\x99\xd9\x13\xbf\xbd\x18\xe1\xdc\x15\xc1\x26\x90\xe4\xcf\x8f\xe7\x2d\xd7\x39\xe2\x89\xf5\x9c\x13\x07\xea\xf9\x95\xd5\x72\xab\x11\xe7\x43\xd8\xa1\x52\xe1\xab\x71\x5a\xa7\xc4\xa0\x3a\x43\xe8\xc8\x62\x20\x65\x18\x3b\xc8\xa1\x4b\xf5\x5c\x38\x58\xe7\x14\x03\xb0\x43\x92\xc5\x9c\xf5\x71\xc2\x7b\x16\xc1\x91\x41\x8b\xb7\x42\x9e\x13\x89\x7c\xa0\x50\xa5\x13\xcf\xd7\x23\x36\x16\x4e\xde\x14\xd9\x7e\x8b\x40\x6f\xea\xf6\x85\xee\xc7\x8c\x88\xb9\x5a\x00\xbe\x3f\x22\x6e\x6f\xc6\xc5\x49\x4f\xec\xd3\x78\xba\x04\xda\xbc\xa5\x26\x11\x2f\x0f\xd3\x16\xd8\x7b\xea\xff\x22\x3f\xd1\x05\x7a\xbe\x5f\x43\x85\x5d\x73\x5a\xaf\xb9\x3b\xfc\x41\x6d\xb2\x0d\xf5\x19\x00\x00\xff\xff\xc5\xcb\x88\xa8\x7d\x01\x00\x00")

func assetsDnsServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceYaml,
		"assets/dns/service.yaml",
	)
}

func assetsDnsServiceYaml() (*asset, error) {
	bytes, err := assetsDnsServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service.yaml", size: 381, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa9, 0x6d, 0x77, 0x15, 0x61, 0x5f, 0x22, 0x3d, 0xd7, 0x33, 0xf2, 0xcf, 0xe0, 0xbd, 0x5d, 0xc6, 0xf1, 0x17, 0x84, 0xff, 0x81, 0xdb, 0x33, 0xd0, 0x30, 0x3b, 0xe3, 0x8e, 0x8e, 0x39, 0x8d, 0x48}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"assets/dns/cluster-role-binding.yaml": assetsDnsClusterRoleBindingYaml,

	"assets/dns/cluster-role.yaml": assetsDnsClusterRoleYaml,

	"assets/dns/configmap.yaml": assetsDnsConfigmapYaml,

	"assets/dns/daemonset.yaml": assetsDnsDaemonsetYaml,

	"assets/dns/metrics/cluster-role-binding.yaml": assetsDnsMetricsClusterRoleBindingYaml,

	"assets/dns/metrics/cluster-role.yaml": assetsDnsMetricsClusterRoleYaml,

	"assets/dns/metrics/role-binding.yaml": assetsDnsMetricsRoleBindingYaml,

	"assets/dns/metrics/role.yaml": assetsDnsMetricsRoleYaml,

	"assets/dns/namespace.yaml": assetsDnsNamespaceYaml,

	"assets/dns/service-account.yaml": assetsDnsServiceAccountYaml,

	"assets/dns/service.yaml": assetsDnsServiceYaml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"assets": {nil, map[string]*bintree{
		"dns": {nil, map[string]*bintree{
			"cluster-role-binding.yaml": {assetsDnsClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml":         {assetsDnsClusterRoleYaml, map[string]*bintree{}},
			"configmap.yaml":            {assetsDnsConfigmapYaml, map[string]*bintree{}},
			"daemonset.yaml":            {assetsDnsDaemonsetYaml, map[string]*bintree{}},
			"metrics": {nil, map[string]*bintree{
				"cluster-role-binding.yaml": {assetsDnsMetricsClusterRoleBindingYaml, map[string]*bintree{}},
				"cluster-role.yaml":         {assetsDnsMetricsClusterRoleYaml, map[string]*bintree{}},
				"role-binding.yaml":         {assetsDnsMetricsRoleBindingYaml, map[string]*bintree{}},
				"role.yaml":                 {assetsDnsMetricsRoleYaml, map[string]*bintree{}},
			}},
			"namespace.yaml":       {assetsDnsNamespaceYaml, map[string]*bintree{}},
			"service-account.yaml": {assetsDnsServiceAccountYaml, map[string]*bintree{}},
			"service.yaml":         {assetsDnsServiceYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
