package blockTests

import (
	helper "github.com/dangula/goDockerTest/helpers"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

const FILE_PATH = "/tmp/rook2"
const FILE_DATA = "writing in folder provisioned by rook"

func TestIfBlockStorageMountIsPresent(t *testing.T) {
	t.Log("Make sure File storage mount is present")
	if _, err := helper.IsDirectoryPresent(FILE_PATH); err != nil {
		t.Log(err)
		t.Errorf("Expected file storage mount to be presnet")
		t.Fail()

	}
}

func TestWritingToBlockStorageMount(t *testing.T) {
	BLOCK_FILE_NAME := uuid.NewV4().String()
	t.Log("Make sure you can write to file storage mount")
	if _, err := helper.WriteToFile(FILE_PATH, BLOCK_FILE_NAME, FILE_DATA); err != nil {
		t.Log(err)
		t.Errorf("Cannot Write to file Storage Mount")
		t.Fail()
	}

}

func TestReadingFromBlockStorageMount(t *testing.T) {
	BLOCK_FILE_NAME := uuid.NewV4().String()
	_, err := helper.WriteToFile(FILE_PATH, BLOCK_FILE_NAME, FILE_DATA)
	if err != nil {
		panic(err)
	}
	t.Log("Make sure you can read from file storage mount")
	if readData, err := helper.ReadFile(FILE_PATH, BLOCK_FILE_NAME); err != nil {
		t.Log(err)
		t.Errorf("Cannot Read from file Strorage Mount")
		t.Fail()
	} else {
		assert.Equal(t, readData, FILE_DATA, "content of the file should be unchanged")

	}

}
