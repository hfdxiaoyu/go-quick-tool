package file

import "testing"

func TestReadFile(t *testing.T) {
	fileData, err := ReadFile("./test.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(fileData))
}

func TestReadFileByBuffer(t *testing.T) {
	fileData, err := ReadFileByBuffer("./test.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fileData)
}

func TestWriteFile(t *testing.T) {
	data := []byte("good good")
	if err := WriteFile(data, "./test.txt"); err != nil {
		t.Fatal(err)
	}
}

func TestWriteFileByBuffer(t *testing.T) {
	data := []byte(" lucking lucking")
	if err := WriteFileByBuffer(data, "./test.txt"); err != nil {
		t.Fatal(err)
	}
}
