package cw

import (
	"github.com/samuel/go-zookeeper/zk"
	"strings"
)

type EventType int

const (
	EventNodeUnknown     EventType = 0
	EventNodeDeleted     EventType = 2
	EventNodeDataChanged EventType = 3
)

func (et EventType) String() string {
	switch et {
	case EventNodeDeleted:
		return "EventNodeDeleted"
	case EventNodeDataChanged:
		return "EventNodeDataChanged"
	default:
		return "EventNodeUnknown"
	}
}

func Existed(c *zk.Conn, path string) (bool, error) {

	existed, _, err := c.Exists(path)

	return existed, err
}

// 确保路径存在
func MakeSurePath(c *zk.Conn, path string) error {

	pathArr := strings.Split(path, "/")
	curPath := ""
	for _, node := range pathArr {
		if len(node) <= 0 {
			continue
		}

		curPath += "/"
		curPath += node
		existed, err := Existed(c, curPath)
		if err != nil {
			return err
		}

		if existed {
			continue
		}

		err = Create(c, curPath, []byte(""), 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func Create(c *zk.Conn, path string, data []byte, flags int32) error {

	acls := zk.WorldACL(zk.PermAll)
	_, err := c.Create(path, data, flags, acls)
	if err != nil {
		return err
	}

	return nil
}
