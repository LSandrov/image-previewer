package previewer

import (
	"bufio"
	"fmt"
	"os"
)

func loadImage(imgName string) []byte {
	fileToBeUploaded := "./test_img/" + imgName
	file, err := os.Open(fileToBeUploaded)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	bytes := make([]byte, fileInfo.Size())

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, _ = buffer.Read(bytes)

	return bytes
}
