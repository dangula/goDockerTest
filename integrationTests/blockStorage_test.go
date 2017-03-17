package integrationTests

import (
	"testing"
	helper "github.com/dangula/goDockerTest/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/satori/go.uuid"
)

const BLOCK_PATH = "/tmp/rook1"
const BLOCK_DATA = "writing in folder provisioned by rook"

func TestBlockStoage_integrationTest(t *testing.T) {
	BLOCK_FILE_NAME := uuid.NewV4().String()
	t.Log("Test Block Storage mount - provioned by rook")

	t.Log("Make sure mount directory is present")
	if _,err := helper.IsDirectoryPresent(BLOCK_PATH); err != nil{
		t.Log(err)
		t.Errorf("Expected Block mount to be presnet")
		t.Fail()

	}

	t.Log("Make sure you can write to block storage mount")
	if _, err := helper.WriteToFile(BLOCK_PATH, BLOCK_FILE_NAME, BLOCK_DATA); err != nil {
		t.Log(err)
		t.Errorf("Cannot Write to Block Storage Mount")
		t.Fail()
	}
	t.Log("Make sure you can read from block storage mount")
	if readData, err := helper.ReadFile(BLOCK_PATH, BLOCK_FILE_NAME); err != nil {
		t.Log(err)
		t.Errorf("Cannot Read From Block Strorage Mount")
		t.Fail()
	} else {
		assert.Equal(t,readData,BLOCK_DATA,"content of the file should be unchanged")

	}

}