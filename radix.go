package radix

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrNotFound 不存在
	ErrNotFound = errors.New("Node Not Found")
	// ErrExisted 不存在
	ErrExisted = errors.New("Node Existed")
)

// NodeValue 节点数据
type NodeValue interface{}

// Radix  接口
type Radix interface {
	Add(string, NodeValue) error
	Get(string) (NodeValue, error)
	Delete(string) error
}

type node struct {
	key      string
	indices  []byte
	children []*node
	value    NodeValue
}

// New Tree
func New() Radix {
	return &node{}
}

// Add 添加元素
func (n *node) Add(key string, value NodeValue) error {
	if n.key == "" && (len(n.children) == 0) {
		n.key = key
		n.value = value
	} else {
	START:
		end := min(len(key), len(n.key))

		pos := 0
		for pos < end && key[pos] == n.key[pos] {
			pos++
		}

		if pos < len(n.key) {
			child := &node{
				key:      n.key[pos:],
				indices:  n.indices,
				children: n.children,
				value:    n.value,
			}
			n.indices = []byte{n.key[pos]}
			n.key = n.key[:pos]
			n.children = []*node{child}
			n.value = nil
		}

		if pos < len(key) {
			key = key[pos:]
			for i := 0; i < len(n.indices); i++ {
				if key[0] == n.indices[i] {
					n = n.children[i]
					goto START
				}
			}
			child := &node{
				key:   key,
				value: value,
			}
			n.children = append(n.children, child)
			n.indices = append(n.indices, key[0])
		} else if pos == len(key) {
			if n.value != nil {
				return ErrExisted
			}
			n.value = value
		}
	}
	return nil
}

func (n *node) Get(key string) (NodeValue, error) {
START:
	if n.key == key {
		return n.value, nil
	} else if len(n.key) < len(key) && n.key == key[:len(n.key)] {
		key = key[len(n.key):]
		for i := 0; i < len(n.indices); i++ {
			if key[0] == n.indices[i] {
				n = n.children[i]
				goto START
			}
		}

	}
	return n, ErrNotFound
}

func (n *node) Delete(key string) error {
	var parent *node
START:
	if n.key == key {
		if len(n.children) != 0 {
			n.value = nil
		} else if parent != nil {
			children := []*node{}
			for _, child := range parent.children {
				if child.key == key {
					continue
				}
				children = append(children, child)
			}
			parent.children = children
		}
		return nil
	} else if len(n.key) < len(key) && key[:len(n.key)] == n.key {
		parent = n
		key = key[len(n.key):]
		for i := 0; i < len(n.indices); i++ {
			if key[0] == n.indices[i] {
				n = n.children[i]
				goto START
			}
		}
	}
	return ErrNotFound
}

func (n *node) walk(pos, depth int) string {

	if pos < 0 {
		pos = 0
	}

	var builder strings.Builder

	if depth > 0 {
		builder.WriteString(strings.Repeat(" ", pos+(depth-1)*2))
		builder.WriteString("|-")
	}

	builder.WriteString(fmt.Sprintf("%s(%v)", n.key, n.value))
	builder.WriteString("\n")

	for _, child := range n.children {
		builder.WriteString(child.walk(pos+len(n.key)-1, depth+1))
	}

	return builder.String()
}

func (n *node) String() string {
	return n.walk(0, 0)
}
