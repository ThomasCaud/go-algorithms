package main

import "fmt"

type block struct {
	key string
	value string
	next *block
}

func getNewBlock(key string, value string) block {
	res := block{key: key, value:value, next: nil}
	return res;
}

func getHash(key string) int {
	return len(key);
}

func append(hashTable map[int]*block, key string, value string) {
	newBlock := getNewBlock(key, value);
	hash := getHash(key);

	c, ok := hashTable[hash];
	if ok {
                for c.next != nil {
			c = c.next;
                }
                c.next = &newBlock;
	} else {
		hashTable[hash] = &newBlock;
	}
}

func displayHashTable(hashTable map[int]*block) {
	for key, b := range hashTable {
		fmt.Println("Key ", key);
		fmt.Printf("	Value: %v", b.value);

		for b.next != nil {
			b = b.next;
			fmt.Printf(" -> %s", b.value);
		}
		fmt.Println();
	}
}

func main() {
	h := make(map[int]*block);
	append(h,"testKey", "test value");
	append(h,"1", "Hello");
	append(h,"2", "World");
	append(h,"3", "=)");
	displayHashTable(h);
	fmt.Printf("Ended");
}
