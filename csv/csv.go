package csv

import (
	"encoding/csv"
	"os"
	"strings"
)

// LoadCsv 导入csv文件
func LoadCsv(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

// ExportCsv csv导出
func ExportCsv(filePath string, head []string, data [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write(head)
	for _, datum := range data {
		writer.Write(datum)
	}
	writer.Flush()

	return nil
}

// ExportCsvToUri 浏览器下载
func ExportCsvToUri(head []string, data [][]string) (strings.Builder, error) {
	var b strings.Builder
	_, err := b.WriteString("\xEF\xBB\xBF") //设置utf8字符集防止乱码
	if err != nil {
		return b, err
	}
	writer := csv.NewWriter(&b)
	writer.Write(head)
	for _, datum := range data {
		writer.Write(datum)
	}
	writer.Flush()

	return b, nil
}
