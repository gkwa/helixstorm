package core

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func Run(dirPath string) {
	if dirPath == "" {
		fmt.Println("Please specify a directory path using the -dir flag.")
		os.Exit(1)
	}

	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
		os.Exit(1)
	}

	totalFiles := len(files)
	deleteCount := totalFiles / 10

	rand.New(rand.NewSource(time.Now().UnixNano())).Shuffle(len(files), func(i, j int) {
		files[i], files[j] = files[j], files[i]
	})

	var filesToDelete []string
	for i := 0; i < deleteCount; i++ {
		filesToDelete = append(filesToDelete, files[i])
	}

	for _, file := range filesToDelete {
		err := os.Remove(file)
		if err != nil {
			fmt.Printf("Error deleting file %s: %v\n", file, err)
		} else {
			fmt.Printf("Deleted file: %s\n", file)
		}
	}

	fmt.Printf("Total files: %d\n", totalFiles)
	fmt.Printf("Files to delete: %d\n", deleteCount)

	GenerateRandomFiles(dirPath)
}

func GenerateRandomFiles(dirPath string) {
	words := []string{"unzip", "list"}
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	for i := 0; i < 1; i++ {
		nestedPath := dirPath
		depth := rand.Intn(10) + 1

		for j := 0; j < depth; j++ {
			subfolder := fmt.Sprintf("folder_%d", rand.Intn(1000))
			nestedPath = filepath.Join(nestedPath, subfolder)
			err := os.MkdirAll(nestedPath, os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating directory: %v\n", err)
				continue
			}
		}

		filePath := filepath.Join(nestedPath, "test.md")
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			continue
		}
		defer file.Close()

		fmt.Printf("Writing to file %s\n", filePath)

		_, err = file.WriteString(words[0] + "\n" + words[1] + "\n")
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		}
	}

	fmt.Println("Random files generated.")
}
