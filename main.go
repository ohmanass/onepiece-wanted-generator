package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("test")
	pirates, err := ParseCSV("Pirates.csv")
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	for _, p := range pirates {
		fmt.Println(p.Name, p.Prime, p.Img)
	}
}


func ParseCSV(path string) ([]*Pirate, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var Pirates []*Pirate

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Datas extraction
		line = strings.TrimSpace(line)
		line = strings.Trim(line, `"`)
		line = strings.TrimSuffix(line, ";;")

		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			return nil, fmt.Errorf("ligne invalide : %s", line)
		}

		// Split line into unic data
		name := parts[0]
		prime := parts[1]
		img := parts[2]

		// Creation/integration of our data
		p, err := New(name, prime, img)
		if err != nil {
			return nil, err
		}

		Pirates = append(Pirates, p)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return Pirates, nil
}