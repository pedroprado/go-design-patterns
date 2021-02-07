package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

// Once é: thread safe e lazy
// Singletons também podem ser criados com init(): mas este é apenas thread safe (não lazy)
var once sync.Once
var instance *singletonDatabase

func main() {
	db := GetSingletonDatabase()

	pop := db.GetPopulation("Tokyo")

	fmt.Println(pop)
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, err := readData("capitals.txt")
		db := singletonDatabase{capitals: caps}
		if err != nil {
			fmt.Println(err)
		}
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})

	return instance
}

func readData(path string) (map[string]int, error) {
	// ex, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// exPath := filepath.Dir(ex)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}
