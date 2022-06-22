package main

import "fmt"

const ArraySize = 7

// HashTable structure
type HashTable struct {
	array [ArraySize]*bucket
}

//Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// //Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// //Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// bucket (Linked List) structure
type bucket struct {
	head *bucketNode
}
type bucketNode struct {
	key  string
	next *bucketNode
}

//insert
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, " already exists")
	}
}

//search
func (b *bucket) search(k string) bool {
	currNode := b.head
	for currNode != nil {
		if currNode.key == k {
			return true
		}
		currNode = currNode.next
	}
	return false
}

//delete
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == k {
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

//hash function
func hash(k string) int {
	var sum int
	for _, v := range k {
		sum += int(v)
	}
	return sum % ArraySize
}

//Init function
func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

//custom print formating
func (h *HashTable) show() {
	for i := range h.array {
		currNode := h.array[i].head
		for currNode != nil {
			fmt.Print(currNode.key, " ")
			currNode = currNode.next
		}
		fmt.Print("\n")
	}
}
func main() {
	testHashTable := Init()

	list := []string{
		"ERIC",
		"RANDY",
		"ADAM",
		"NURI",
	}

	for _, v := range list {
		testHashTable.Insert(v)
	}

	testHashTable.show()
	testHashTable.Delete("ADAM")
	testHashTable.show()
	// fmt.Printf("Type %T. Value %v\n", testHashTable, testHashTable)
	// // fmt.Println(hash("RANDY"))

	// testBucket := &bucket{}
	// testBucket.insert("RANDY")
	// testBucket.insert("ADAM")
	// testBucket.insert("NURI")
	// fmt.Println(testBucket)
	// testBucket.delete("RANDY")
	// // fmt.Println(testBucket.head.key)
	// fmt.Println(testBucket.search("RANDY"))
	// fmt.Println(testBucket.search("RAN"))
}
