package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	dirPath = "****"
)

func cleanfile() {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		log.Println(path)
		log.Println(info)
		log.Println(info.ModTime())

		if !info.IsDir() && time.Since(info.ModTime()) > 30*24*time.Hour {
			err := os.Remove(path)
			go cleandir(path)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("error walking the path %v: %v", dirPath, err)
	}
}

func cleandir(file string) {
	dir := filepath.Dir(file)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("Failed to read directory: %v", err)
	}

	if len(files) == 0 {
		err := os.Remove(dir)
		if err != nil {
			log.Printf("Failed to remove directory: %v", err)
		} else {
			log.Printf("Directory %v removed successfully", dir)
		}
	} else {
		log.Printf("Directory %v is not empty", dir)
	}

}

func main() {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for t := range ticker.C {
			cleanfile()
			fmt.Println("Task runs at:", t)
		}
	}()

	// 运行无限循环，否则主线程可能在第一个ticker之前就退出
	select {}
}
