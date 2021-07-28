package main

// imports
import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Create transaction struct
type Transaction struct {
	address_from [32]byte
	address_to   [32]byte
	data         string
	amount       float32
}

// Create block struct
type Block struct {
	nonce        uint64
	transactions []Transaction
	timestamp    int64
	prev_hash    [32]byte
}

type Blockchain struct {
	blocks []Block
}

// Mine function
// Create new block
// Verify chain

/// --- Block Function ---

// Block initialization
func initBlock(nonce uint64, timestamp int64) Block {
	b := Block{}
	b.nonce = nonce
	b.timestamp = timestamp
	return b
}

func transactionsHash(trans_list []Transaction) uint64 {
	return 0
}

// Block Hash
func (b Block) getBlockHash() [32]byte {
	block_str := string(b.nonce) + string(b.prev_hash[:]) + string(b.timestamp) + string(transactionsHash(b.transactions))
	return sha256.Sum256([]byte(block_str))
}

func (bc Blockchain) verifyBlock(index uint) bool {
	if index == 0 {
		panic("Cannot verify genesis block!")
	}
	block_1 := bc.blocks[index-1]
	block_2 := bc.blocks[index]

	// verify that the prev_hash field is equal to actual last block hash
	if block_2.prev_hash != block_1.getBlockHash() {
		return false
	}

	// every block hash should begin with this pattern "0" here
	if block_2.getBlockHash()[0] != 8 {
		return false
	}

	// blockchain should be ordered in time
	if block_2.timestamp < block_1.timestamp {
		return false
	}

	return true
}

// set block hash
func (block *Block) setHash(hash [32]byte) {
	block.prev_hash = hash
}

// set block nonce
func (block *Block) setNonce(nonce uint64) {
	block.nonce = nonce
}

// Add Transaction
func (block Block) addTransaction(from [32]byte, to [32]byte, data string, amount float32) Block {
	block.transactions = append(block.transactions, Transaction{from, to, data, amount})
	return block
}

func (bc *Blockchain) addBlock(block Block) {
	bc.blocks = append(bc.blocks, block)
}

func (bs *Blockchain) mineLast() {
	l := len(bs.blocks)
	if l < 2 {
		return
	}

	var i uint64 = 0

	for !bs.verifyBlock(uint(l - 1)) {
		i++
		bs.blocks[l-1].setNonce(i)
	}

	println(i)
}

// HERE BEGINS THE MAGIC
func main() {

	blockchain := Blockchain{}

	addr_1 := sha256.Sum256([]byte("ADRESS 1"))
	addr_2 := sha256.Sum256([]byte("ADRESS 2"))

	// Create genesis block
	genesis_block := initBlock(0, time.Now().UnixNano()).addTransaction(addr_1, addr_2, "This is the first transaction", 1)
	block_2 := initBlock(0, time.Now().UnixNano()).addTransaction(addr_2, addr_1, "This is the second transaction", 1)
	block_2.setHash(genesis_block.getBlockHash())
	// Push genesis block
	blockchain.addBlock(genesis_block)
	// Push second block
	blockchain.addBlock(block_2)

	/*  Verification and proof of work

	- each block should contain the hash of the previous block in the "prev_hash" field
	- the nonce of each block has to be chosen to make the block's hash begin by the null character (proof of work by spoofing)
	- TODO : adjustable difficulty = number of hash bytes that should be 0 or another constant pattern
	- TODO : refactor code
	- TODO :  add MerkleTrees for transactions checking and better data visualisation/manipulation

	*/

	// is the second block correct ?
	println(blockchain.verifyBlock(1))

	// then mine the block
	blockchain.mineLast()

	// it should be correct now
	println(blockchain.verifyBlock(1))
	hash := blockchain.blocks[1].getBlockHash()
	fmt.Println(hash)

}
