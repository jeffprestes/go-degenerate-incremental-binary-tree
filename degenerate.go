package godegenerateincrementalbinarytree

/*
function insertLeaf(
        bytes32 leaf,
        bytes32 root,
        bool isFirstLeaf
    ) internal returns (bytes32 newRoot) {
        newRoot = isFirstLeaf ? leaf : hash(root, leaf);
    }
*/

import (
	"log"
	"math/big"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

type DegenerateAndIncrementalBinaryTree struct {
	ActualRoot *big.Int
	Leaves     map[int]*big.Int
	LastIndex  int
	Zero       *big.Int
}

func NewDegenerateAndIncrementalBinaryTree(zero *big.Int) (tree DegenerateAndIncrementalBinaryTree) {
	tree.Leaves = make(map[int]*big.Int)
	tree.Zero = zero
	return
}

func (daib *DegenerateAndIncrementalBinaryTree) InsertLeaf(leaf *big.Int) (err error) {
	daib.LastIndex++
	daib.Leaves[daib.LastIndex] = leaf
	if daib.ActualRoot == nil {
		daib.ActualRoot = leaf
		return
	}
	daib.ActualRoot, err = InsertLeafWithPreviousRoot(leaf, daib.ActualRoot, false)
	if err != nil {
		log.Printf("InsertLeaf - error InsertLeafWithPreviousRoot - Tree: %+v\nError: %sd\n", daib, err.Error())
		return
	}
	return
}

func (daib *DegenerateAndIncrementalBinaryTree) InsertBatchLeaves(leaves []*big.Int) (err error) {
	for _, leaf := range leaves {
		err = daib.InsertLeaf(leaf)
		if err != nil {
			log.Printf("InsertBatchLeaf - error InsertLeaf - Leaf: %s\nError: %s\n", leaf.Text(10), err.Error())
			return
		}
	}
	return
}

func (daib *DegenerateAndIncrementalBinaryTree) FillWithZeros(maxLeaves uint) (err error) {
	missingLeaves := maxLeaves - uint(daib.LastIndex)
	if maxLeaves < 1 {
		return
	}
	for i := 0; i < int(missingLeaves); i++ {
		err = daib.InsertLeaf(daib.Zero)
		if err != nil {
			return err
		}
	}
	return
}

func InsertLeafWithPreviousRoot(leaf, oldRoot *big.Int, isFirstLeaf bool) (root *big.Int, err error) {
	if isFirstLeaf {
		root = leaf
	} else {
		root, err = poseidon.Hash([]*big.Int{oldRoot, leaf})
		if err != nil {
			log.Println("InsertLeafWithPreviousRoot - error hashing", oldRoot.Text(10), " and ", leaf.Text(10), " - error is: ", err.Error())
			return
		}
	}
	return
}
