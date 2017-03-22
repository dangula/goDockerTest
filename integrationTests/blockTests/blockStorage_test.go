package blockTests

import (
	helper "github.com/dangula/goDockerTest/helpers"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

const BLOCK_PATH = "/tmp/rook1"
const BLOCK_DATA = "writing in folder provisioned by rook"

func TestIfBlockStorageMountIsPresent(t *testing.T) {
	t.Log("Make sure block mount is present")
	if _, err := helper.IsDirectoryPresent(BLOCK_PATH); err != nil {
		t.Log(err)
		t.Errorf("Expected Block mount to be presnet")
		t.Fail()

	}
}

func TestWritingToBlockStorageMount(t *testing.T) {
	BLOCK_FILE_NAME := uuid.NewV4().String()
	t.Log("Make sure you can write to block storage mount")
	if _, err := helper.WriteToFile(BLOCK_PATH, BLOCK_FILE_NAME, BLOCK_DATA); err != nil {
		t.Log(err)
		t.Errorf("Cannot Write to Block Storage Mount")
		t.Fail()
	}

}

func TestReadingFromBlockStorageMount(t *testing.T) {
	BLOCK_FILE_NAME := uuid.NewV4().String()
	_, err := helper.WriteToFile(BLOCK_PATH, BLOCK_FILE_NAME, BLOCK_DATA)
	if err != nil {
		panic(err)
	}
	t.Log("Make sure you can read from block storage mount")
	if readData, err := helper.ReadFile(BLOCK_PATH, BLOCK_FILE_NAME); err != nil {
		t.Log(err)
		t.Errorf("Cannot Read From Block Strorage Mount")
		t.Fail()
	} else {
		assert.Equal(t, readData, BLOCK_DATA, "content of the file should be unchanged")

	}

}
