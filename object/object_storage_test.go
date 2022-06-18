package object

import (
	"fmt"
	"testing"
)

func TestObject(t *testing.T) {
	url1, err := UploadImgFile("img.png")
	if err != nil {
		error.Error(err)
	}
	fmt.Println("imgUrl:", url1)

	url2, err := UploadVideoFile("bear.mp4")
	if err != nil {
		error.Error(err)
	}
	fmt.Println("videoUrl:", url2)
}
