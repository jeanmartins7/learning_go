package hash_table

import "fmt"

func TestHashTable() {
	pessoas := []Pessoa{
		{"Maria", "Da silva", 12, "F"},
		{"Joao", "Fulano", 31, "M"},
	}

	table := HashTable{}
	keys := make([]int, len(pessoas))

	for i, pessoa := range pessoas {
		keys[i] = table.Put(pessoa)
	}

	for _, key := range keys {
		ps := table.Get(key)
		for _, pessoa := range ps {
			fmt.Println(pessoa)
		}
	}
}
