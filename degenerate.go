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

func (daib *DegenerateAndIncrementalBinaryTree) InsertLeaf(leaf *big.Int) (err error) {
	daib.LastIndex++
	daib.Leaves[daib.LastIndex] = leaf
	if daib.ActualRoot == nil {
		//In case the leaf value is 0, big.Int deals with it as nil, so it's need to hash as []byte
		if leaf.Int64() == 0 {
			daib.ActualRoot, err = poseidon.HashBytes([]byte("0"))
		} else {
			daib.ActualRoot, err = poseidon.Hash([]*big.Int{leaf})
		}
		if err != nil {
			log.Println("InsertLeaf - error InsertLeafWithPreviousRoot - error hashing", leaf.Text(10), " - error is: ", err.Error())
			return
		}
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

func NewDegenerateAndIncrementalBinaryTree(zero *big.Int) (tree DegenerateAndIncrementalBinaryTree) {
	tree.Leaves = make(map[int]*big.Int)
	tree.Zero = zero
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
