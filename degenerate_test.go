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

func TestTreeBatchInsertion(t *testing.T) {
	var inputsBigInt []*big.Int
	inputs := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42}
	for i := 0; i < len(inputs); i++ {
		inputsBigInt = append(inputsBigInt, big.NewInt(inputs[i]))
	}
	tree := NewDegenerateAndIncrementalBinaryTree(big.NewInt(0))
	assert.NotNil(t, tree)
	assert.Nil(t, tree.ActualRoot)
	tree.InsertBatchLeaves(inputsBigInt)
	assert.NotNil(t, tree.ActualRoot)
	t.Log(tree.ActualRoot.Text(10))
	t.Logf("0x%064s", tree.ActualRoot.Text(16))
}
