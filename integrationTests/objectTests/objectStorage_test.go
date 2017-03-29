package objectTests

import (
	helper "github.com/dangula/goDockerTest/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/satori/go.uuid"
	"testing"
)

func Test_Creating_New_Bucket_On_RookObjectStore(t *testing.T){
	created,err := helper.CreateBucket("mybucket1")
	if err !=nil{
		t.Log(err)
		t.Errorf("Error Creating new bucket on rook object store")
		t.Fail()
	} else {
		assert.True(t,created,"new bucket should be created successfully")
	}

	isBucketPresent,err := helper.IsBucketPresent("mybucket1")
	if err !=nil{
		t.Log(err)
		t.Errorf("Error getting bucket from rook object store")
		t.Fail()
	} else {
		assert.True(t,isBucketPresent,"new bucket should be present on rook object store")
	}
}

func Test_Putting_Object_On_A_Bucket_In_RookObjectStore(t *testing.T){
	isBucketPresent,err := helper.IsBucketPresent("mybucket1")
	if err !=nil{
		t.Log(err)
		t.Errorf("Error getting bucket information from rook object store")
		t.Fail()
	}
	if !isBucketPresent {
		created,err := helper.CreateBucket("mybucket1")
		if err !=nil{
			t.Log(err)
			t.Errorf("Error Creating new bucket on rook object store")
			t.Fail()
		}else {
			assert.True(t,created,"new bucket should be created successfully")
		}
	}
	OBJ_NAME := uuid.NewV4().String()
	putSuccess,err := helper.PutObjectInBucket("mybucket1",OBJ_NAME,"object 1 datat for rook")
	if err !=nil{
		t.Log(err)
		t.Errorf("Error putting data in rook object store")
		t.Fail()
	} else {
		assert.True(t,putSuccess,"Object successfully uploaded to rook object store")
	}

}

func Test_Getting_Object_From_A_Bucket_In_RookObjectStore(t *testing.T){
	isBucketPresent,err := helper.IsBucketPresent("mybucket1")
	if err !=nil{
		t.Log(err)
		t.Errorf("Error getting bucket information from rook object store")
		t.Fail()
	}
	if !isBucketPresent {
		created,err := helper.CreateBucket("mybucket1")
		if err !=nil{
			t.Log(err)
			t.Errorf("Error Creating new bucket on rook object store")
			t.Fail()
		}else {
			assert.True(t,created,"new bucket should be created successfully")
		}
	}
	OBJ_NAME := uuid.NewV4().String()
	OBJ_DATA := "Object data for rook object store"
	putSuccess,err := helper.PutObjectInBucket("mybucket1",OBJ_NAME,OBJ_DATA)
	if err !=nil{
		t.Log(err)
		t.Errorf("Error putting data in rook object store")
		t.Fail()
	} else {
		assert.True(t,putSuccess,"Object successfully uploaded to rook object store")
	}

	str,err :=helper.GetObjectFromBucket("mybucket1",OBJ_NAME)
	if err != nil{
		t.Log(err)
		t.Errorf("Error Getting Object from rook object store")
		t.Fail()
	}else {
		assert.Equal(t,OBJ_DATA,str,"Make sure data on object is unchanged")
	}


}


