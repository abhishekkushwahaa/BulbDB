package main

import "github.com/abhishekkushwahaa/BulbDB/files"

func main() {
	files.SaveData1("data/test.txt", []byte("Hello from BulbDB"))
	files.SaveData2("data/test.txt", []byte("Hello from Atomically BulbDB Files"))
}
