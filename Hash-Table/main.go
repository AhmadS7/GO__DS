package main

import "fmt"

// size of the hashTable array
const ArraySize = 8

// HashTable structure / will hold an array
type HashTable struct {
	array [ArraySize]*Bucket

}

// Bucket structure
type Bucket struct {
	 head *BucketNode
}

// BucketNode structure
type BucketNode struct {
	key  string
	next *BucketNode
}

// Insert
// func (h *HashTable) Insert(key string) {
// 	index := hash(key)
// 	h.array[index].Insert(key )

// }

// Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search() 

}

// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)

}

// Insert
func (b *Bucket) Insert(key string) {

	if b.search(k) {
		newNode := &BucketNode{key: k}
		newNode.next = b.head
	
	
		b.head = newNode
	}  else {
		fmt.Println("Already Exists")
	}



}

// Search
func (b *Bucket) Search(key string) {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k 
	}

}
// Delete
func (b* Bucket) delete(k string) {

	if .b.head.key == k{
		b.head = b.head.next
	}

	previousNode := b.head
		for previousNode.next != nil {
			if previousNode.new.key == k{
				previousNode.next = previousNode.next.next
			}
			previousNode = previousNode.next
		}
}

// Hash
func Hash(key string) int {
	sum := 0
	for _, v := range key {
		s += int(v)
	}
	return sum % ArraySize

}

// Init
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}

	}
	return result
}

func main() {
	testHashTable := Init()
	fmt.Println(testHashTable)

}
