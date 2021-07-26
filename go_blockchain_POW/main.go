package main

// imports
import (
	"crypto/sha256"
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
}

// BC = Linked list of block structs
var blockchain []Block

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

// Add Transaction
func (block Block) addTransaction(from [32]byte, to [32]byte, data string, amount float32) Block {
	block.transactions = append(block.transactions, Transaction{from, to, data, amount})
	return block
}

// HERE BEGINS THE MAGIC
func main() {

	addr_1 := sha256.Sum256([]byte("ADRESS 1"))
	addr_2 := sha256.Sum256([]byte("ADRESS 2"))
	now := time.Now()

	// Create genesis block
	genesis_block := initBlock(0, now.UnixNano()).addTransaction(addr_1, addr_2, "This is the first transaction", 1)

	// Push genesis block
	blockchain = append(blockchain, genesis_block)
	println(genesis_block.nonce, genesis_block.timestamp, genesis_block.transactions, blockchain)

}
