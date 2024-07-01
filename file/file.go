package file

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"strings"
)

// ReadFile 读取文件
func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var fileData []byte
	var temp = make([]byte, 128)
	for {
		n, err := file.Read(temp)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		fileData = append(fileData, temp[:n]...)
	}
	return fileData, nil
}

// ReadFileByBuffer 使用缓冲区读文件的方法
func ReadFileByBuffer(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var fileStr strings.Builder
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
		if _, err = fileStr.WriteString(str); err != nil {
			return "", err
		}
	}
	return fileStr.String(), nil
}

// WriteFile 写文件到磁盘
func WriteFile(data []byte, path string) error {
	fileExisted := FileExist(path)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, fs.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	if fileExisted {
		file.WriteString("\xEF\xBB\xBF")
	}
	if _, err = file.Write(data); err != nil {
		return err
	}

	return nil
}

// WriteFileByBuffer 使用缓冲区写数据到磁盘
func WriteFileByBuffer(data []byte, path string) error {
	fileExisted := FileExist(path)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, fs.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	if fileExisted {
		write.WriteString("\xEF\xBB\xBF")
	}
	if _, err = write.Write(data); err != nil {
		return err
	}
	write.Flush()
	return nil
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}
