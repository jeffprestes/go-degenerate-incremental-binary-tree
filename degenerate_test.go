package godegenerateincrementalbinarytree

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTree(t *testing.T) {
	tree := NewDegenerateAndIncrementalBinaryTree(big.NewInt(1))
	assert.NotNil(t, tree)
	assert.Nil(t, tree.ActualRoot)
	tree.InsertLeaf(big.NewInt(9))
	assert.Equal(t, tree.LastIndex, 1)
	assert.NotNil(t, tree.ActualRoot)
	assert.Equal(t, tree.ActualRoot.Text(10), "5199363853932272446084541931873785938987820779897294035064941545455873932186")
}

func TestTreeWithZero(t *testing.T) {
	tree := NewDegenerateAndIncrementalBinaryTree(big.NewInt(0))
	assert.NotNil(t, tree)
	assert.Nil(t, tree.ActualRoot)
	tree.InsertLeaf(big.NewInt(0))
	assert.NotNil(t, tree.ActualRoot)
	t.Log(tree.ActualRoot.Text(10))
}
