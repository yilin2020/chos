package redblack

type Color uint8

const (
	RED Color = iota
	BLACK
)

func (c Color) isRed() bool {
	return c == RED
}
func (c Color) isBlack() bool {
	return c == BLACK
}

type RBT struct {
	root *Node
	size int
}

func New() *RBT {
	return &RBT{}
}

type Node struct {
	key   string
	value string
	right *Node
	left  *Node
	color Color
}

func genratorNode(key, value string) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

func (b *RBT) Add(key, val string) {
	b.root = b.add(b.root, key, val)
	b.root.color = BLACK

}

func (b *RBT) add(node *Node, key, val string) *Node {
	if node == nil {
		b.size++
		return genratorNode(key, val)
	}
	if key < node.key {
		node.left = b.add(node.left, key, val)
	}

	if key > node.key {
		node.right = b.add(node.right, key, val)
	}
	return node
}

//左旋转
//   node                     x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2
func (b *RBT) lRotate(node *Node) *Node {
	x := node.right
	node.right = x.left
	x.left = node
	x.color = node.color
	node.color = RED
	return x
}

//     node                   x
//    /   \     右旋转       /  \
//   x    T2   ------->   T3   node
//  / \                       /  \
// T3  T1                    T1  T2
func rRotate(node *Node) *Node {
	x := node.left
	node.left = x.right
	x.right = node
	x.color = node.color
	node.color = RED
	return x
}
