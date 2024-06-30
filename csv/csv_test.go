package csv

import (
	"testing"
)

func TestLoadCsv(t *testing.T) {
	csvData, err := LoadCsv("./test.csv")
	if err != nil {
		t.Fatal(err)
	}
	for _, datum := range csvData {
		t.Log(datum)
	}
}

func TestCsvExport(t *testing.T) {
	head := []string{"id", "name"}
	data := [][]string{{"1", "张三"}, {"2", "李四"}}
	if err := ExportCsv("./test.csv", head, data); err != nil {
		t.Fatal(err)
	}
}
