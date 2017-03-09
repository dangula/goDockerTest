package main

import (
	"io/ioutil"
	"os"
	"testing"
)

const path = "/tmp/rook1"
const filename = "file1"
const data = "writing in folder provisioned by rook"

func TestFileWriteAndRead(t *testing.T) {
	t.Log("Write and Read Test")

	if _, err := writeToFile(path, filename, data); err != nil {
		t.Log(err)
		t.Errorf("Expected Block to be created")
		t.Fail()
	}

	if readData, err := readFile(path, filename); err != nil {
		t.Log(err)
		t.Errorf("Expected Block to be created")
		t.Fail()
	} else {
		t.Log(readData)
	}

}

func writeToFile(path string, filename string, data string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, err
	}
	fileCreate, _ := os.Create(path + "/" + filename)
	if err != nil {
		return false, err
	}
	defer fileCreate.Close()

	file, err := os.OpenFile(path+"/"+filename, os.O_RDWR, 0644)
	if err != nil {
		return false, err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return false, err
	}

	return true, nil

}

func readFile(path string, filename string) (string, error) {

	str, err := ioutil.ReadFile(path + "/" + filename)
	if err != nil {
		return "fail", err
	}

	return string(str), nil

}
