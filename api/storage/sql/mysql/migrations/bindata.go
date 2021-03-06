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

var __1_initial_schema_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _1_initial_schema_down_sql() ([]byte, error) {
	return bindata_read(
		__1_initial_schema_down_sql,
		"1_initial_schema.down.sql",
	)
}

var __1_initial_schema_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x4d\x53\xdb\x3a\x14\x5d\xbf\xfc\x0a\xed\x70\x66\x58\x10\xde\x83\x61\x86\x61\x61\x12\x85\xe7\x79\xc1\xe6\x39\x4e\xa7\x74\xa3\x11\xb6\x70\xd5\xda\x72\xc6\x56\x42\xe1\xd7\x77\x24\x27\x8e\x24\x7f\xe0\x14\x4a\x4b\xb3\x8a\xaf\x8e\xae\xaf\xa5\x73\xce\x95\x3d\x98\xf8\xde\x0d\x08\xec\xcb\x19\x04\xce\x14\xc0\x8f\xce\x3c\x98\x83\x55\x41\xf2\xe2\x7c\x30\xf6\xa1\x1d\xc0\xcd\xa8\x8c\x01\x6b\x00\x00\x8d\x80\xf1\xbb\xa3\x31\x65\x7c\x7b\xe5\x7a\x01\x70\x17\xb3\x19\xb0\x17\x81\x87\x1c\x77\xec\xc3\x6b\xe8\x06\x87\x03\x00\xa2\xda\xdc\x35\xce\xc3\xcf\x38\xb7\x8e\x4f\x4e\x87\xbb\xa9\x02\x2b\x6e\xc8\x70\x4a\xea\xd8\x91\x80\x6a\x58\x92\x62\x9a\x34\xe6\x3d\xfd\x47\x60\x27\x70\x6a\x2f\x66\x3b\xbc\x9e\x57\xc5\xff\x7d\x6c\xe6\xbe\xa3\x59\x4b\xcd\x27\x23\x01\x36\x73\xd3\x34\xee\x7c\x46\x13\xbf\xcc\x69\x48\x10\xcf\x50\x4a\x8a\x02\xc7\xc4\x5c\x4e\x13\x1f\xe6\x04\x73\x12\x3d\xbb\xfc\x72\x0d\x97\x51\x27\xd6\xcc\x7d\xe3\x3b\xd7\xb6\x7f\x0b\xfe\x83\xb7\xc0\xa2\xd1\x50\xc4\x16\xae\xf3\xff\x02\xca\x90\xd8\x3e\x2b\xaa\xc7\x45\xb8\x1e\xad\x36\xd0\xda\xfe\x1b\x0e\x86\x00\xba\x57\x8e\x0b\x2f\x1c\xc6\xb2\xc9\xa5\x41\x91\x8b\xa3\xaa\xa2\xf1\xbf\xb6\x3f\x87\xc1\xc5\x8a\xdf\x9f\x81\xb1\x37\x9b\xd9\x01\x94\x17\x28\x26\x8c\xe4\x38\x41\x21\x3d\x1f\x34\xd3\x37\xcc\xd2\x74\xc5\x28\xa7\xc4\x24\xb1\x32\xf2\x52\x2a\x3f\xed\x41\xe5\x76\xba\x8d\x8e\xcf\x0c\x6c\xf6\xc0\x48\x8e\x54\xa1\xb4\xe7\x2d\xb1\xaa\x50\xda\x25\x12\x91\x22\xcc\xe9\x92\xd3\x8c\x35\xd2\x58\x93\x53\x11\xe6\xd9\x03\xc2\x69\xb6\xda\xac\x45\x3b\xc5\xda\xe9\xde\x2c\x8f\x04\x17\x1c\xe1\x90\xd3\x35\xe9\x45\xc9\x52\x1e\x4b\x92\x57\xfa\x68\xaf\xa5\x92\xd2\x97\x8c\xb2\x67\xea\x5e\xae\xee\x12\x1a\x6a\x5b\x9f\x65\x09\xc1\xcc\xa8\x63\xf4\xb3\x35\x17\x91\x84\x18\xf8\xc6\x42\x8e\x7a\x08\xb4\x59\x88\xa5\x08\xa5\x00\x8d\x11\x41\x62\xeb\x69\x33\x63\xec\xb9\xf3\xc0\xb7\x1d\x37\x50\x75\x22\xe9\x25\xf8\x28\xff\x14\x68\x42\x23\x74\x9f\xe5\x84\xc6\x0c\x4c\x3d\x1f\x3a\x57\x6e\x59\x4d\x45\xdc\x21\xf0\xe1\x14\xfa\xd0\x1d\xc3\xf9\xb6\x6d\xc8\xb0\xe7\x82\x09\x9c\xc1\x00\x02\x1f\xce\x03\xdf\x19\x07\x22\xb4\xb8\x99\xd8\x4a\xe8\xad\x2c\x82\xad\x49\x5e\x60\x21\x87\xba\x49\x28\x63\x0d\x36\xf1\x83\x16\xd1\xdd\x05\xb6\x2b\xfe\x88\xc4\xac\xee\xae\xb8\x73\x87\x76\x1c\x27\xdf\x78\xdf\x7b\x27\x94\x7d\xed\x8b\xd5\xe5\xde\x8d\x5d\xd3\x88\x64\x3d\xb1\xba\x1c\x3b\xa4\x58\x02\x91\x94\x7a\x9b\xb4\x0e\x8e\x0e\xea\xaa\xed\xab\xd8\x7d\xd4\xfa\xda\x4a\xd5\xf4\x28\x02\x3a\x2d\x2c\xed\xb2\x2e\x5a\x85\xb7\x48\x83\x22\x55\xd0\x9f\xda\x14\xac\x67\x57\x55\xac\xf5\xcd\xa7\xde\x5a\xee\x28\xb0\xb7\xad\x6c\x71\xef\xc1\x55\x38\x8e\x4d\x33\x11\x21\xe9\x21\x1c\xc7\xe0\x83\xed\x8b\xf4\xf2\x50\x5a\xf9\x86\x42\x94\x5a\x95\xaf\x75\x1a\x22\x8c\x37\x1d\x85\x44\xf8\x2d\xcf\x41\x1a\x01\xe4\xc4\xbe\x46\xb7\x8f\xd9\xed\x6b\x78\xcf\xe3\xdf\x61\xfb\x2f\xcd\xc3\x58\x6e\xcb\x8c\x54\xd0\x6a\xb5\x77\x7a\x6b\x38\x12\x08\xbe\x20\x33\x07\xd2\x7d\xa7\xc3\x5d\x8c\x7b\xeb\x06\xa3\xf5\xdc\x17\x58\xcc\xa6\xca\x3f\xd1\x5e\x76\xfe\xdc\x60\x34\xfa\xa0\x54\xb5\x79\xaa\x00\x95\x01\x95\x34\xd7\x78\x5b\x3a\x94\xfa\x53\xdd\x4a\x03\xff\x05\xfa\xf6\x0d\xb5\x71\x54\x7d\xcd\x98\xcd\x71\x3c\x54\xee\xaa\xce\x16\x4f\x22\xc7\xcb\x69\x1b\xa6\xd7\x1b\xa3\x78\x64\xa3\x8c\x43\x31\xf9\x8d\xdf\x37\x1f\x51\xd3\x87\x13\x63\xb4\xd1\x6f\x35\x8b\xe8\x30\x5b\x73\x3f\x35\xdb\xea\x72\xcf\x56\xa0\x78\x53\x22\x11\x12\x9e\xd5\x5e\x89\xf4\x4c\x72\xcf\x77\x30\x03\x58\x33\x58\x01\xce\x09\x2e\x36\x2f\x9b\xda\xd7\x98\xae\x8f\x0d\x1d\x5e\xb4\xff\x71\x48\x5b\xf7\xdf\xf3\x40\xa4\x97\xf8\x8b\x4d\x6b\x1f\x2d\x0c\xbe\x07\x00\x00\xff\xff\xdd\x02\x28\x6e\x3a\x14\x00\x00")

func _1_initial_schema_up_sql() ([]byte, error) {
	return bindata_read(
		__1_initial_schema_up_sql,
		"1_initial_schema.up.sql",
	)
}

var __2_stored_procedure_get_communities_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _2_stored_procedure_get_communities_down_sql() ([]byte, error) {
	return bindata_read(
		__2_stored_procedure_get_communities_down_sql,
		"2_stored_procedure_get_communities.down.sql",
	)
}

var __2_stored_procedure_get_communities_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x93\x4d\x6e\xdb\x30\x10\x85\xd7\xe2\x29\x66\x27\x0b\x30\x7c\x81\xa2\x0b\xd9\x66\x8c\x18\x89\x64\xc8\xf2\x22\x2b\x82\xa5\x58\x81\x85\xfe\x20\x92\x0d\xea\xd3\x17\x1c\xb2\x36\xfd\xd3\x64\x35\xf3\xde\x47\x86\x7a\x33\xde\x54\x34\xaf\x29\x1c\xaa\x72\x43\xb7\xa7\x8a\x42\x2b\x0d\x13\x63\xdf\xdb\x41\x19\x25\xf5\x22\x23\x6b\xba\x7b\x2d\x08\x39\xd2\x37\xba\xa9\x81\x24\xfb\x63\x59\xb0\xbc\xaa\xf2\x8f\x7c\xb7\x5b\x60\x57\xae\xf7\x74\x53\x2f\x48\x92\xa4\xaa\x49\x97\x20\x56\xaa\x59\x02\x81\xf0\x97\x0e\xbc\x97\x28\xbb\x22\x36\xce\x01\x3f\xdf\xf2\xe3\xe7\x20\xe7\x6d\xf0\xb0\x61\xcd\x13\xe2\xa4\xe5\x7c\xb9\xda\x63\x36\x48\x31\xdb\x48\x2d\x66\x35\x19\x35\x0e\x48\x46\x7d\x8c\x39\x75\xfc\xcc\xfb\xd1\x0e\x06\x39\x2f\x30\x8e\x4a\x4c\xaa\xbe\xf5\x1f\xd9\xb7\xb1\xdc\x71\x6d\xb8\x30\xea\xb7\x7f\x90\x6b\x99\xef\x63\x6a\x9a\x95\x90\x07\x39\xbf\x4b\xad\x79\xeb\x51\xd4\xd8\x24\x67\xd6\x7b\xf5\xe1\x40\x3d\xee\x47\x35\x44\xb0\x19\xd9\xaf\x51\xdd\x7c\xc0\x64\x7f\x74\x4a\x78\x06\xcb\xd8\x14\xb3\xe4\x46\xfa\x44\x43\xbd\x04\x37\x30\x3b\x35\x17\x23\xd4\xb7\xe1\x75\xf2\x9f\x1d\xea\xd8\x36\xbc\xd5\xe9\x12\xcc\xca\x15\xb1\xe1\x06\xe1\x1c\xbb\xc2\x2a\x23\x49\x06\x5c\x43\x5a\x49\x6d\x3b\x93\x92\x97\xaa\x7c\x07\x92\x44\xab\x06\x82\xbc\xd1\x97\x1a\xf6\xe5\x6b\x01\x0b\x92\x5c\x16\xee\x02\xfd\x61\x7e\x4f\x92\xbb\x1d\x14\xf8\xef\x33\xc8\x8f\xe0\x9e\x41\x12\xbc\xfc\x7a\xca\x89\x20\x0c\x49\x76\x55\x79\x3a\xc0\xfa\x03\x84\x59\xdd\x5c\x4a\x32\x30\x50\x16\x70\x27\xc3\x77\xbf\x9b\xff\x79\xd8\xc3\xbb\xbe\xfe\x69\xb8\x20\xc2\x52\xfb\x54\x70\xa9\xaf\x91\xb9\x79\xca\x66\xcb\x8d\xf4\x88\xef\x99\x1b\x49\x44\x75\xf2\xa7\xb9\x32\xae\x7b\x46\x54\x92\x6b\xbf\xee\x81\x99\x51\xc0\x31\xe4\x47\xc0\x99\xe0\x89\xbb\xa8\xd0\x00\x61\xd1\xbb\xc6\x65\x1f\xe2\xb2\x2e\x2e\xfb\x3c\xae\x6f\x84\xd0\x62\x4b\xfe\x06\x00\x00\xff\xff\xf4\xe0\xcf\x06\x5c\x04\x00\x00")

func _2_stored_procedure_get_communities_up_sql() ([]byte, error) {
	return bindata_read(
		__2_stored_procedure_get_communities_up_sql,
		"2_stored_procedure_get_communities.up.sql",
	)
}

var __3_stored_procedure_get_users_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _3_stored_procedure_get_users_down_sql() ([]byte, error) {
	return bindata_read(
		__3_stored_procedure_get_users_down_sql,
		"3_stored_procedure_get_users.down.sql",
	)
}

var __3_stored_procedure_get_users_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xcf\x6e\x83\x30\x0c\x87\xcf\xf1\x53\xf8\x16\x90\xd0\x5e\x60\x27\x4a\x3d\xb4\x6a\x85\x29\x65\x87\x9d\x50\xda\x58\x28\x52\x33\x2a\x42\xde\x7f\x5a\xc8\xba\xac\x39\xfd\xfc\x7d\x71\xfe\xb8\x51\x54\x0f\x84\xef\xaa\x6f\x68\xff\xa1\x08\x27\x5e\xc7\xe0\x79\xf1\x45\x09\x3b\x6a\x5f\x3b\x80\x13\xbd\x51\x33\x20\x88\xc3\xa9\xef\xc6\x5a\xa9\xfa\xb3\x6e\xdb\x22\x56\xfd\xee\x40\xcd\x50\x80\x10\xd2\x1a\x59\x61\x78\xb2\xa6\x42\xc0\xb4\xa4\x49\xd4\xfc\xc7\x3f\x37\x7c\x69\xc7\xd1\xfd\x16\xd5\x9f\x67\xa7\xed\x35\xca\x98\x32\x73\xef\x7a\xe8\x38\xdb\x39\xe2\xb3\x9d\x33\x6a\xdd\xb4\x3d\xca\x4d\x19\xbd\x2d\xf6\xc2\xc3\x7c\x64\xef\xf5\xb4\x9d\x16\xd1\xb8\xce\xa3\xdb\x60\xb6\xfb\xb2\xb0\x5e\x79\xfb\x46\xca\x99\x0d\x37\x73\xb7\x29\x97\x20\x4a\xd4\x1e\xa5\x62\x1f\xae\xab\x84\x17\xd5\x1f\x11\x44\x1c\x2b\x86\x67\x00\xea\xf6\xdf\x01\x00\x00\xff\xff\xde\x59\xe8\x0a\x79\x01\x00\x00")

func _3_stored_procedure_get_users_up_sql() ([]byte, error) {
	return bindata_read(
		__3_stored_procedure_get_users_up_sql,
		"3_stored_procedure_get_users.up.sql",
	)
}

var __4_stored_procedure_get_conversations_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _4_stored_procedure_get_conversations_down_sql() ([]byte, error) {
	return bindata_read(
		__4_stored_procedure_get_conversations_down_sql,
		"4_stored_procedure_get_conversations.down.sql",
	)
}

var __4_stored_procedure_get_conversations_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x93\x41\xaf\xa2\x30\x14\x85\xd7\xbd\xbf\xe2\xee\x80\xc4\xf8\x07\x26\xb3\x40\xac\x44\xe3\x80\xa9\xb8\x70\x36\xc4\x29\x8d\x69\x06\xa8\x81\x62\xde\xf3\xd7\xbf\x50\xc0\xd7\xa2\x79\xac\xce\x3d\xe7\x2e\xb8\xe7\x4b\x23\x46\xc3\x8c\xe2\x81\xa5\x11\x5d\x9f\x18\xc5\xab\xd0\x39\x57\xf5\x5d\x34\xed\x45\x4b\x55\xb7\x7e\x00\x2b\x1a\x6f\x13\x80\x23\xdd\xd3\x28\x43\x20\xbb\x63\x9a\xe4\x21\x63\xe1\x39\x8c\x63\xdf\x4c\xe9\x6a\x47\xa3\xcc\x07\x42\x3c\x59\x78\x0b\xe4\x4b\x59\x2c\x10\x70\xfc\xbc\xc7\xe8\x3e\x64\xb1\xf8\x76\xb9\xaa\xaa\xae\x96\xfa\xf3\xef\x18\x3f\x8d\xdc\x5d\xec\x5a\xd1\xac\xc7\x9d\x5e\xe7\x85\x13\x6b\xf1\xa1\x4d\xd6\x0b\xcb\x2f\x65\xfd\xdf\xf8\xbd\xb0\x7c\x59\x5d\x87\x7f\xac\xae\x96\x7b\x97\x85\x50\xc6\x37\xca\x4a\x6e\xdd\xbf\x52\x72\x13\x0d\xf2\x25\x3b\x34\x92\x0b\x6b\x21\xbf\xf5\x86\x7d\x6a\x23\x2e\x5a\x8c\x57\x0e\xda\xbe\xef\x56\x3c\xd3\x51\x5b\x69\x21\x4a\x31\xa5\xa3\x9e\x95\x28\x6a\xdd\xf6\xb1\x5a\x4e\x53\x00\x24\xc0\x4b\x8b\x1e\x13\x6d\x57\x6a\x0f\x36\x2c\xfd\x83\x40\x1c\xb6\xc8\x61\x4f\x37\x19\xee\xd2\x6d\x82\x3e\x90\x27\x61\x67\xcd\xb0\xe8\xcd\x1f\xb9\x4f\xe0\x95\x4b\xde\xa6\xaf\x5c\xfc\x73\xb2\xea\x0d\x5a\x07\xaf\x9a\xf1\x75\x18\xab\x19\xe4\x79\xeb\xea\xb5\xf6\x79\xf5\xea\xb5\xfb\x79\xff\x6a\x02\x10\x00\x21\x01\x86\x47\x9c\x0a\x07\x62\x1a\x9e\x46\xe4\x0a\x48\xcc\xd2\xd3\x01\x57\xe7\x01\x8c\xdb\x28\x04\xc8\x15\xa6\xc9\xbb\x0c\x7f\x0f\x4f\xe5\x17\x00\x4d\xd6\x5f\x01\x00\x00\xff\xff\x96\x4d\x72\x64\xa4\x03\x00\x00")

func _4_stored_procedure_get_conversations_up_sql() ([]byte, error) {
	return bindata_read(
		__4_stored_procedure_get_conversations_up_sql,
		"4_stored_procedure_get_conversations.up.sql",
	)
}

var _bindata_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x99\x4b\x6f\x5c\xc7\x11\x85\xd7\x9c\x5f\x31\x21\xe0\x80\x04\x04\xe5\xbe\x1f\x06\xb4\xb1\x91\x6d\x16\xd9\xa6\x03\xa2\xef\xed\xbe\xcc\x20\xe2\xc3\x33\xa4\xdd\x76\xe0\xff\x1e\x7c\x55\x67\x40\xc5\xa0\x64\x29\x36\xe0\x85\x34\x9c\x99\xdb\xdd\xd5\x55\xa7\x4e\x9d\xaa\x79\x8c\xeb\xbf\xe3\x6d\xde\xdf\x1d\x6e\x8f\xf1\xe9\xf0\x70\x7f\xda\xed\x0e\x77\x8f\x0f\xc7\xa7\xfd\xd5\xee\xe2\x72\xf9\xf1\x29\x9f\x2e\x77\x17\x97\xeb\xc3\xdd\xe3\x31\x9f\x4e\x7f\xb9\xfd\xe9\xf0\xc8\x07\xdb\xdd\x13\x2f\x87\x07\xfe\x3f\x3d\x1d\x0f\xf7\xb7\xa7\xcb\xdd\xf5\x6e\xb7\x3d\xdf\xaf\xfb\xe5\x70\x9f\xe2\x53\xbc\x39\xe6\x98\xae\xf8\x6b\xff\x8f\x7f\xb2\xd7\x9b\xfd\x7d\xbc\xcb\x7b\x7f\xfe\x7a\x7f\x75\xfe\x34\x1f\x8f\x0f\xc7\xeb\xfd\x7f\x76\x17\xb7\x3f\xd9\xbb\xfd\xd7\xef\xf6\x1c\xf5\xf6\x6f\xf9\x87\xbf\xe7\x98\xf2\xf1\xca\x6c\xe1\xfd\x37\xcf\xdb\x96\x8f\xb6\xed\xf5\xf5\xee\xe2\xb0\xd9\x82\x3f\xbd\xdb\xdf\x1f\xde\xb3\xc5\xc5\x31\x3f\x3d\x1f\xef\x79\xfb\x66\xbf\xdd\x3d\xbd\xfd\x2b\xbb\x6f\x57\x97\x6c\xb4\xff\xea\xbb\xaf\xf7\x5f\x7d\x7f\xe9\x96\xd8\x59\xd7\xbb\x8b\x9f\x77\xbb\x8b\xef\xe3\x71\xbf\x3c\x6f\x7b\x3f\xc7\x0f\xd9\x5d\xdc\xb8\x39\xef\xf6\x87\x87\xb7\xdf\x3e\x3c\xfe\x78\xf5\xe7\xe5\x79\x7b\xb3\xbf\xfd\xe9\x1a\x5b\xdf\x7e\xfb\xfe\xe1\x94\xaf\xae\x77\xbf\x9f\x19\x5a\xb5\x3c\x6f\x6f\xbf\xc1\x92\xab\xeb\x37\x6c\xb1\xfb\x79\xb7\xc3\xc2\x9b\x9b\xfa\xe6\x70\x7f\x78\x3a\xc4\xf7\x37\xa7\xf5\x5f\xf9\x2e\xde\xa4\x87\x1f\xee\x6f\x4e\xdf\xbd\xdf\xbf\x93\x97\xaf\x2e\x43\xa9\xb7\x50\xa6\x25\x94\x6a\x0a\xa5\xaa\x5e\xff\xb7\x6d\xa1\x54\xf5\xff\xbe\xb7\xcf\x3e\xf2\xbc\xff\xbb\x3c\x07\xf9\xe3\x96\x5c\xbd\x1a\xda\xf3\xcd\x3e\xc4\xc6\xee\xe2\xe2\x13\x37\x7a\xb3\xbb\xb8\xb8\xfc\xe5\xb7\x6f\xf9\xf6\xed\xe9\xbb\xf7\x97\x6f\x76\x17\xd7\x9f\xf0\xcb\xf3\xe3\xff\xed\x95\xb4\x86\xd2\x4f\xa1\x74\x29\x94\xbe\x0d\x25\x2d\xa1\xb4\x31\x94\xba\x0b\xa5\x4f\xa1\x2c\x78\x6a\x0d\xa5\x8a\xa1\xe4\x14\xca\x58\x85\x32\x0c\xbe\xa6\xae\x42\x49\x39\x94\xa9\x0d\x65\xa8\x43\x99\x06\x7f\xe5\x5f\xdd\x84\x32\xf5\xa1\xe4\x31\x94\x71\x0e\x65\xad\x43\xc9\x43\x28\xed\x1c\x4a\x97\x43\x89\x7c\xde\x85\x12\xdb\x50\xea\x3a\x94\x65\xf0\xbd\x53\x1f\x4a\x8a\xa1\x8c\x4d\x28\x2b\xe7\x0c\xa1\x74\x4d\x28\xb9\x0e\x25\xb1\x66\x0c\xa5\xe9\x42\x69\xc6\x50\xa6\xec\x7f\x8f\x5b\x28\xb9\x72\x9b\xbb\x18\x4a\xb7\x84\xb2\xb4\xa1\x4c\x31\x94\xb8\xf9\x73\x31\xfb\xdf\xb1\x0f\x65\x6c\x43\x59\x73\x28\x73\x1f\x4a\x9b\x42\x99\xa7\x50\xb6\xc9\xef\x52\x25\xf7\x59\x5e\x43\x59\x97\x50\xea\x39\x94\xaa\xf3\xe7\xd9\x7f\xad\x42\x99\x36\x7f\xdf\xae\xbe\x96\xfb\xf7\x7d\x28\x5d\x1d\xca\x86\xad\x4d\x28\xe3\x1a\x4a\x8b\x8f\x87\x50\x22\xfe\x48\xbe\x76\x4d\xa1\x44\xd6\xac\x8e\xc8\x61\x79\x89\xc7\x94\x42\x99\x58\x83\x3f\x16\xf7\x4d\xcb\x33\xbd\xef\x37\x2e\xa1\x64\xfe\x8e\xbe\x16\x7f\xd5\xa3\xdf\x15\x3b\x17\xbd\x9f\x88\xc5\x18\x4a\xbd\xba\xbf\xec\x2e\xad\x9f\x95\xc1\xc3\xe0\xdf\x57\xad\x9f\x1b\x1b\xf7\x39\x58\x68\x7b\xbf\x1b\xcf\xaf\x5b\x28\xed\x14\xca\x22\x5f\x77\x9b\xc7\x8e\xb5\x66\x1f\xbe\x69\x42\x69\x96\x50\x06\xd6\x4c\x6e\x13\x31\xc8\xec\x37\x85\x32\xd7\x7e\xa7\x21\x3a\x66\xe6\x26\x94\xa1\x09\x65\x8e\xa1\xb4\x9d\x30\x91\x42\xd9\x52\x28\xdd\x18\xca\x50\x79\x6c\x47\xad\x69\x36\xc7\x5b\x4b\x4c\xf1\x77\x0e\xa5\x1f\x43\x59\x7b\xc7\xef\x98\x42\x19\x56\xdf\x67\xc9\xee\xb3\x7e\x76\x0c\xe0\x6b\xf6\x6a\x5a\xf7\x57\x3b\x78\xec\x53\xab\xb3\xb3\xfb\xb5\x1b\x42\xa9\x5b\xc7\xfa\xba\x86\x32\xb0\x5e\xf8\xc6\x07\xbd\xee\x04\x9e\xb8\xef\xca\xf9\xe4\xc7\xea\xfe\x60\x2d\xf9\x86\x0d\x60\x85\x33\xc0\x3f\xf7\xc2\x4f\xe4\x0f\x78\x06\x5b\xf3\x18\x4a\x5f\xbb\x5d\x69\xf0\xf3\xb8\x43\x6e\xfd\x8e\x60\xdc\x72\x01\x76\x62\xad\xf2\x8b\x18\xac\xe7\x58\xd5\x6e\x13\x36\xd4\x83\x63\x7b\x6b\xc5\x68\xd8\x17\x43\x99\x2b\xf7\x7f\x9b\x3d\x3e\x31\xba\xdd\x5d\x1f\xca\xc2\xdd\xf0\x55\x72\x9f\xa4\xca\xe3\xcf\x59\x35\x78\xcc\x9e\xcb\x60\x63\x89\xee\x6f\xe2\x8f\x3d\xc4\x03\x5c\x91\x9b\x0b\x67\xcd\x9e\x53\x60\xdf\x62\x2d\xdc\x72\x1e\xb6\x1a\x96\x37\xbf\x53\xbb\x39\x76\xe0\x02\xe2\x87\x3f\xd3\x16\xca\xbc\x39\x66\x17\x62\x34\x86\x32\xe3\x87\xd9\x63\x36\x77\x8e\xe9\x66\x0a\xa5\x19\x3c\x6f\x72\xe7\xb8\xe4\xcc\xa6\x76\x3f\xe3\x7f\x62\x6a\xb1\x1d\xdd\xaf\xa9\x71\x8e\xe1\x2c\xd6\x47\xc5\x0d\xce\xb0\xb3\x1a\xcf\xd7\x1e\x0c\x47\xb7\x0f\x8e\x63\x5f\xcb\xb3\xc1\xef\x6f\x79\x99\x1d\x0b\xd5\xea\x78\x23\xaf\x89\x07\x38\xe1\xbb\xbe\xf3\x3c\x5f\x3a\xe7\xbc\x26\xf9\x19\xf8\xa1\x97\x7f\x9b\xde\x71\x81\x9f\x9a\xc6\x6d\xcc\xb3\xe7\x04\x76\x4f\xca\x3f\xec\xb6\x5c\x69\x9d\x93\x6b\x71\xa7\xf9\x84\xbc\x6c\x1c\xa3\xf8\x94\x7d\xe3\xea\x71\x05\x3f\xab\xf8\x93\x7d\xe1\xb3\x66\x75\xbe\xaa\xe4\x47\x62\x07\x97\x1a\x3e\x5a\x3f\x8f\x67\x59\x4f\x2c\xf1\x05\xd8\x20\x7f\xb0\x8b\xd8\x6f\x31\x94\x61\x7a\xe1\x4f\xb3\xb7\x71\x2c\x63\x33\x38\x18\x46\xcf\xf9\x5e\x5c\x4b\x6e\x71\xd7\x3a\x2a\xf7\x17\x8f\x33\x5c\x56\xa9\x1e\x70\x7f\xb8\x70\x00\xbb\x9d\xf3\x18\xf6\x60\xa7\x61\x09\x9b\xc5\xcd\xab\xea\x00\xbe\x36\xfe\x9b\xdc\xd7\xe4\xfa\x34\xb9\x5d\xd5\xf9\x55\x5c\x8a\xdf\x88\x25\x7c\x93\x17\xf7\x59\xdb\xfa\x79\x70\xe2\x32\x7b\xde\x10\x47\xf2\x0b\x8e\x06\x3b\xc4\x6b\xec\x3d\x46\xf8\x86\xb8\x61\x03\x18\x22\xc7\xa8\x03\xf8\x03\x7b\xe0\x16\x30\xc6\xdd\xc0\x0c\x36\x63\x2b\x38\xe9\xd6\x0f\xb0\xa9\x3c\x82\x27\x8d\x9f\x92\xfb\x1c\xce\x05\xc3\x5b\xe5\x71\x61\x6f\xfc\x4d\x4e\x5a\xcd\xc0\x37\x8a\xf5\x90\x9c\x63\xe1\x79\x30\x36\x35\xee\xcb\x79\x75\x7f\x5a\x1d\x69\xfc\x3c\x38\x1c\xac\x90\x5f\x56\xc3\x46\xbf\x0f\x3e\x23\xff\xe7\xd9\xf9\xc1\xb8\x7a\x72\x2e\x30\xac\x34\xfe\x77\x37\x7b\x2c\x8d\x8f\x55\x4f\xc8\x05\x3e\x6f\x14\xb3\x4a\xfc\x41\x8e\xc3\x51\x70\x11\xdc\x94\xa8\xa1\xbd\xdb\xdc\x9c\x39\x6c\xf2\xdc\x03\xab\x59\x18\xc1\x0e\xb8\x92\xbd\xc9\x35\xb0\x69\xbc\x08\xaf\x4f\x7a\x5d\xfc\x33\x78\xc0\x78\x32\x39\xc7\xa0\x13\xf0\xb7\xe9\x8f\xe4\x5a\xc5\x72\x08\xdb\x3a\x71\x6f\xed\x36\xb3\x2f\xb9\x11\x3b\xcf\xab\x4a\x39\x68\x7b\x46\xe7\x1b\xfc\x62\xbc\x97\xfd\x15\x2c\x47\x69\x9e\xb8\xb8\xaf\x38\x3f\x9d\xe3\x13\xdd\x77\xd8\x01\x76\x1b\xed\x4b\x9c\x1b\xf9\x92\xb5\xe4\x46\x43\xcc\xa3\xeb\x21\x70\xd6\x49\x33\x10\x0f\x74\x03\xb1\x06\x3f\x49\xba\x0b\x1f\x91\x4f\xe0\x85\xef\x79\xbf\xf5\xd2\x0e\xa3\xdf\x99\xfd\xf0\x63\x23\x3c\xb0\x6f\x54\x2d\x1c\xe1\x8b\xcd\x31\xc1\xfa\x56\x1c\x05\xe6\x5a\xf1\x09\x31\x35\x6c\xd7\xe2\xf3\xd1\x35\x14\x76\x9d\x7d\xcc\xbd\x8c\x1f\x2b\x3f\x0f\x1c\x4e\xca\x43\x7c\xc7\x67\x68\x28\x78\x81\x7c\x21\xd7\xb9\x93\xf1\x98\xf2\x89\x3b\x77\x9d\xf2\x27\x39\x87\x92\x93\xb5\x34\x0c\x5c\x0e\x2e\xa8\xdd\xd4\x61\xd3\x1a\x9b\x73\x1d\x76\xa0\x59\xc0\x02\xeb\xb2\xf0\x8c\xbe\xdc\x54\x8b\x56\xe5\x08\x71\x35\x0d\xb2\xf8\xf9\xf0\x71\x92\x86\x58\x7a\xd7\x41\xc4\x9a\xfd\xf0\x07\xdc\x64\x1c\x82\x0f\x2b\xf7\x0b\x39\x62\x35\x43\x7e\xb6\x7d\x93\xef\x67\xda\x6b\x71\xed\x44\x7d\xdb\x94\xbb\x70\x15\x67\x1b\x97\x0d\x5e\x7b\xa9\x23\x56\xa3\xb3\xe3\x1a\xfc\xf3\x2c\xb9\xd4\x0a\x3f\xb1\x72\x1b\xcd\xb6\xd9\xb1\x46\xfe\xe2\x4f\xab\x7b\xbd\xea\xd9\xe8\xf7\xef\xa5\xcb\x67\xad\x37\xad\x9e\x1c\x83\xd5\xb9\xf6\xf7\xfe\x1d\xb5\x89\xdc\xb4\xbd\x6b\x3f\x87\x3b\xad\xd2\xba\xbc\xf2\x79\x2b\x3f\x59\x3e\x46\x69\xdc\xc1\x75\x5e\x73\xd6\x39\xe2\x4d\x74\x01\xf7\xc3\x9e\x5a\x35\x7e\x54\xae\x45\xd5\xae\x4e\x7a\xd5\x70\x92\x3c\x56\x86\x83\xcd\x35\x4f\x52\xae\xd7\xe2\x60\x72\x0e\x3f\x0d\x9b\xf3\xfe\x24\x7d\x8a\x6e\xe6\xbe\x70\x43\xab\xef\xe8\x4b\xf2\xb9\x2f\xd8\xbc\x2f\xc0\x87\xe6\x1b\xe5\x52\xad\xbe\x03\x8d\x48\x5e\xe2\xff\x69\x76\x1c\x77\xeb\xcb\x5a\xb0\xc0\x19\x95\x70\xd3\xea\x15\xbf\x1b\x6f\xce\x9e\xb3\x70\x1b\x18\x33\x3d\x12\xfd\xee\xf8\x0f\xee\xb2\x9a\x9a\x1c\xe3\x9c\xc3\x5d\x4c\x8f\xd6\x9e\xdf\xd8\xbf\x89\xcb\xa8\x31\x5d\xe5\xfc\x43\xcd\x24\x17\xb8\x6b\x27\x7c\xe3\x1f\x72\x83\x9a\x01\x1e\xc1\xe1\x2f\x3b\x55\xe3\xd4\x46\xdc\x9b\x5f\x7a\xb4\x5f\xed\x54\xbd\x37\xfc\xad\x7d\xaa\xef\xf2\x7a\x97\xfa\xfc\xf8\x4a\x8f\xda\xdc\x9c\x9e\x1e\x8e\x39\xdd\x3c\x1e\x1f\xd6\x9c\x9e\x8f\xf9\xe6\x36\x3f\xdd\xac\x0f\x77\x77\xcf\x2c\xcf\xa7\x3f\xae\x9b\xff\x02\xdb\xbe\xd0\x6f\x5f\xb0\xb3\xf9\xf2\xd7\x9f\xff\xd8\x0c\xe0\x33\x4e\xfa\x0d\x53\x01\xba\x22\xd4\xa7\x21\x34\x6b\x2a\x50\x79\x47\x04\xa3\x93\x09\x74\xb9\xcd\xec\xdd\x19\x15\xba\xd2\x33\xac\x45\xcd\xa3\xfe\x2b\x29\xf9\x41\xca\x1d\x35\x47\x46\xc2\x1a\x28\x19\xaa\x4b\xa3\x0e\xd2\x94\x47\xef\x99\x68\x1d\x6c\xef\xcc\x09\x3b\xcf\xea\x9e\xac\xab\x6c\x5f\x94\x2c\xd5\xd5\x14\x88\xaa\xbf\x55\xfe\xde\x33\x8e\x2a\x49\x47\x49\xc5\x82\xf9\x5b\x7d\x06\xe3\xc1\xac\x54\x2d\xec\x67\x9f\x28\x66\xe9\xd4\x69\xc1\xfe\x74\x27\x9d\x6c\xb3\x4c\x17\x6b\x91\xd1\xf8\xc3\xaa\xd8\x59\x1d\x4a\x2d\x36\xea\xb2\x51\x1d\xa6\x5c\x34\x4d\x68\x67\xaf\x2c\xad\x54\x14\x9d\x33\x77\xa4\x13\xa1\x92\x0d\xea\x5a\xad\x92\xcb\x1e\x7c\x33\x29\x16\xd8\x04\xa3\x50\x41\x50\x05\xd6\xf1\x69\x82\x62\xec\xb7\x39\xeb\xe0\x27\x58\x26\x4a\xa9\xc1\x36\xf8\xcf\xd8\x32\xfa\x19\xb0\xc8\x24\xc5\xc9\xf7\xa8\x1d\xeb\xac\xd5\x65\xd1\xb1\x36\xf2\xab\xb1\x61\xad\x8a\xb6\x69\xba\xd1\xaa\xc3\x91\x0a\x45\x95\xb2\x17\xb6\x53\x99\x57\xf9\x2a\xcb\x96\xac\xe9\x05\x6a\x1b\x16\xe5\xae\xec\x8d\xbd\x54\x0b\xf3\xa5\xd4\x81\x75\xfa\xbd\xfb\x95\x57\x6c\xe3\x3b\x94\xc1\xa8\x4a\x8c\x2f\x51\x1b\x74\xd5\x54\x5d\xfc\x47\x15\xe1\xd9\x4a\x13\x26\xfc\x01\x16\x4d\xf5\x44\x67\x7f\xee\x0f\xbb\x12\xab\x5a\x0a\x03\xbf\xa2\x12\x93\x54\x13\x77\x59\x54\x81\xc0\x1a\xdf\xd3\x89\xd3\x2d\x11\x57\x54\x33\xca\x8a\x3b\xa0\x2c\x66\x29\x20\xbe\x87\xb5\x39\x9b\x6e\x97\xee\x09\x35\x46\x65\x47\x35\x58\xf5\xeb\x5c\x45\x12\xf3\x41\xdd\x3f\x18\xc1\x1e\x94\x3f\x15\x02\xd5\x9a\x85\xdb\x56\xca\xc9\xa6\x3a\xb3\xdb\x1b\xcf\x95\x34\xb9\x4f\xe9\x5c\x06\x4d\xc0\x46\x4d\xbf\x46\x4d\x78\x56\xa9\xd6\x5a\x9d\x15\x4a\x0b\x3f\xa2\x10\xd8\xd7\xb0\x9e\xbc\x33\x6c\x55\x11\x51\xc9\xe7\x0e\x0c\x2c\x63\x33\x3e\xc0\x3e\xf3\x79\xe3\x4a\xb5\x4e\x2f\x13\x1e\x62\x09\x7e\x92\x3a\xc9\x5a\x9d\x2d\x4a\xa8\xee\xdd\xbf\xac\xc7\xa6\x2c\xb5\x09\xc6\x56\x75\xdd\x9c\x03\x56\xf9\xbe\x97\x8a\xe8\x85\x5f\xe2\x81\x4d\x70\xd1\xac\x8e\xd3\x72\x75\x75\x7f\xf3\x19\x77\x89\x52\x66\x70\x88\xd9\x52\xfb\xb3\xb5\xd4\x10\x18\xb4\x49\x5b\xed\x9c\x64\x9d\xed\xea\x6a\x07\x7f\x93\x13\x70\x08\xf8\x1d\x35\xe9\xe2\x19\xec\xe4\xae\xb5\x14\x2c\xb1\x87\x47\xe0\xae\x49\x53\x47\xd4\x46\x27\xe5\xc1\xd9\x36\xed\x89\x7e\x7f\xb8\x8e\x3c\xb1\x6e\x70\x76\x35\xd8\x4b\xf1\x51\xc1\x89\xdf\x26\x05\x37\xa9\x9b\x25\xb6\xe6\x87\xe8\x13\x0a\x53\x43\x95\xdb\x3a\x6a\x3a\x86\x32\x45\xc5\x2c\x52\x27\xc6\xd9\xea\xe8\xf8\x0e\xbb\x0d\xeb\x59\x93\x82\xc9\xed\x58\x35\xed\x20\x67\x4c\x15\x56\xce\x45\xe0\x8a\x8e\x88\x1c\xb5\x69\x6d\xf4\x5c\x34\x35\x9d\x5c\x41\x4d\xba\xff\x2c\x35\xc8\x7d\xf0\xfd\x19\x17\x70\x26\xeb\xb1\x19\x7e\x02\xdb\xe0\x8d\x73\x6c\xea\xb7\xb8\xbd\xf8\x1b\xfe\x61\x1f\xb0\x8c\x0f\x5a\x75\x02\x9d\xa6\x5c\x96\x67\x95\xfb\xcf\x26\x8d\x93\xab\xf9\x4a\x1c\x93\x34\x21\xa5\xd3\x1e\x94\xcf\xf0\xcd\xb9\xdb\xe2\x1f\x77\x44\xe1\x62\x0b\x0a\x73\x3a\x77\x04\x8d\xe7\xe7\x26\xbe\xf8\xa5\x8a\xe0\xae\x59\x93\xbf\x33\x26\xab\x57\x94\xd6\x67\x57\xe0\xdf\x5f\x43\x7c\xa0\xc6\x3e\x43\x41\xbc\xaa\xcf\xda\xd7\xd7\x3d\x9f\xf2\xf1\x0f\x54\x66\x9f\x65\xd5\x17\xfa\xf3\xb3\xf6\x34\x5f\x7e\xea\xc9\x8f\xe9\xb0\x4f\xee\xfe\x1b\x14\x18\xa8\x9b\x85\xc2\x41\x33\xe0\x56\x4c\x62\x33\xd4\xcd\xe7\x2e\x54\x7f\x98\x09\x76\xb0\x99\x5e\xf5\xc1\xdc\x66\x14\x8b\x24\xef\x4f\x06\xcd\x20\x4c\xb5\xf5\xbe\xcf\x9c\x5e\xe6\x93\xbd\x66\x5c\x30\x04\x2c\xd7\x48\x01\xa1\x9a\xa8\x30\xb0\xb2\xcd\x9f\x34\x73\xa4\xda\xc1\x9a\x9b\x66\xd8\x54\xd0\x4d\x73\x9d\x5e\x7d\x67\xa5\xcc\x83\x39\x60\x6e\x32\x71\xd0\x1c\x2d\xd6\xee\x07\xec\xc4\xe6\xf3\x0c\x02\x06\xe2\x6e\x9d\x66\x85\x54\x00\x9b\x21\x6e\xfe\xf7\x54\x69\x5e\x9f\xfc\x1c\xb3\xb5\xd2\x8c\xa9\x75\x46\xb3\xf9\xc7\xe0\x36\xf3\x1e\x96\xa5\x0a\x9e\xd5\xac\xcd\x5e\xa4\x1c\xad\xe2\xea\xb7\x19\x7c\x31\x49\xe9\x26\xcd\x2f\x61\x65\xab\x0c\x93\x33\x4c\xd4\x0c\xc4\xe6\x86\x9d\xab\x1c\xd8\xd0\x66\x43\x9a\xa5\xb4\xea\x7d\x47\xfd\x9e\xc0\x67\xb0\x16\xe7\xd9\x9c\xac\xd7\x3c\x54\x6a\x33\x27\xfd\xf6\x12\xbd\xc7\x84\x45\xe9\x7f\xb9\xc7\xa8\x19\x5d\x3b\x49\x89\xcf\x62\xbb\xd6\xe3\xd7\x4a\x8d\x24\xc5\x99\x78\xb3\x0f\xea\xdd\x7e\x17\x4a\xae\xae\x6c\x66\xdd\xa8\x7a\xac\xae\x1c\xd8\x1f\x05\xc2\xba\x39\xbd\xcc\xf0\x88\xe1\xaa\xdf\x53\x92\x94\xe7\x32\x38\xab\xe2\x83\x45\x38\xb3\x58\x89\xe9\xcd\xf6\x59\x73\xc1\xd1\xed\x47\x3d\x34\x1f\x28\x51\x9b\x99\x75\xfa\xfd\xa1\x77\x36\x26\x27\xec\xf7\xb9\x45\x33\xdf\xd1\xfd\xc3\xe7\x36\xc3\xed\x34\x1b\x5c\xf4\x1b\xa3\xfa\xe3\xac\xea\x9e\x3e\xc2\x3a\x49\xb3\x56\xfb\xcd\xeb\xbf\x01\x00\x00\xff\xff\x52\x0b\x37\x68\x00\x20\x00\x00")

func bindata_go() ([]byte, error) {
	return bindata_read(
		_bindata_go,
		"bindata.go",
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
	"1_initial_schema.down.sql": _1_initial_schema_down_sql,
	"1_initial_schema.up.sql": _1_initial_schema_up_sql,
	"2_stored_procedure_get_communities.down.sql": _2_stored_procedure_get_communities_down_sql,
	"2_stored_procedure_get_communities.up.sql": _2_stored_procedure_get_communities_up_sql,
	"3_stored_procedure_get_users.down.sql": _3_stored_procedure_get_users_down_sql,
	"3_stored_procedure_get_users.up.sql": _3_stored_procedure_get_users_up_sql,
	"4_stored_procedure_get_conversations.down.sql": _4_stored_procedure_get_conversations_down_sql,
	"4_stored_procedure_get_conversations.up.sql": _4_stored_procedure_get_conversations_up_sql,
	"bindata.go": bindata_go,
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
	"1_initial_schema.down.sql": &_bintree_t{_1_initial_schema_down_sql, map[string]*_bintree_t{
	}},
	"1_initial_schema.up.sql": &_bintree_t{_1_initial_schema_up_sql, map[string]*_bintree_t{
	}},
	"2_stored_procedure_get_communities.down.sql": &_bintree_t{_2_stored_procedure_get_communities_down_sql, map[string]*_bintree_t{
	}},
	"2_stored_procedure_get_communities.up.sql": &_bintree_t{_2_stored_procedure_get_communities_up_sql, map[string]*_bintree_t{
	}},
	"3_stored_procedure_get_users.down.sql": &_bintree_t{_3_stored_procedure_get_users_down_sql, map[string]*_bintree_t{
	}},
	"3_stored_procedure_get_users.up.sql": &_bintree_t{_3_stored_procedure_get_users_up_sql, map[string]*_bintree_t{
	}},
	"4_stored_procedure_get_conversations.down.sql": &_bintree_t{_4_stored_procedure_get_conversations_down_sql, map[string]*_bintree_t{
	}},
	"4_stored_procedure_get_conversations.up.sql": &_bintree_t{_4_stored_procedure_get_conversations_up_sql, map[string]*_bintree_t{
	}},
	"bindata.go": &_bintree_t{bindata_go, map[string]*_bintree_t{
	}},
}}
