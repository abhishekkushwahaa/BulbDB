package files

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randomInt() int {
	// Seed the random number generator to ensure different results each time
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(100) // Generates a random integer from 0 to 99
}

func SaveData2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer func() {
		fp.Close()
		if err != nil {
			os.Remove(tmp)
		}
	}()

	_, err = fp.Write(data)
	if err != nil {
		return err
	}

	err = fp.Sync()
	if err != nil {
		return err
	}

	return os.Rename(tmp, path)
}
