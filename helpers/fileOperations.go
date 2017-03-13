package helpers

import (
	"io/ioutil"
	"os"
)

func IsDirectoryPresent(path string) (bool,error){
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false,err
	}else{
		return true,nil
	}

}

func WriteToFile(path string, filename string, data string) (bool, error) {
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

func ReadFile(path string, filename string) (string, error) {

	str, err := ioutil.ReadFile(path + "/" + filename)
	if err != nil {
		return "fail", err
	}

	return string(str), nil

}
