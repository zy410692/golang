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
	dirPath = "/data/minio-data/yunwei/backup/mysql/10_5_0_126"
)

func cleandir() {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		log.Println(path)
		log.Println(info)
		log.Println(info.ModTime())
		if !info.IsDir() && time.Since(info.ModTime()) > 100*24*time.Hour {
			err := os.Remove(path)
			if err != nil {
				return err
			}
		}
		if info.IsDir() && time.Since(info.ModTime()) > 100*24*time.Hour {
			files, err := ioutil.ReadDir(dirPath + "/" + info.Name())
			if err != nil {
				log.Fatalf("Failed to read directory: %v", err)
			}

			if len(files) == 0 {
				err := os.Remove(dirPath + "/" + info.Name())
				if err != nil {
					log.Fatalf("Failed to remove directory: %v", err)
				} else {
					log.Printf("Directory %v removed successfully", info.Name())
				}
			} else {
				log.Printf("Directory %v is not empty", info.Name())
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("error walking the path %v: %v", dirPath, err)
	}
}

func main() {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for t := range ticker.C {
			cleandir()
			fmt.Println("Task runs at:", t)
		}
	}()

	// 运行无限循环，否则主线程可能在第一个ticker之前就退出
	select {}
}
