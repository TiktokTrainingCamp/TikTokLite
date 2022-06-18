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

	filePath = "E:/golang/tiktok-lite/publics/"
	url1, err = UploadImgFile("img.png")
	if err != nil {
		error.Error(err)
	}
	url2, err = UploadVideoFile("bear.mp4")
	if err != nil {
		error.Error(err)
	}

	filePath = "E:/golang/tiktok-lite/public/"
	videoBucket = "https://tiktok-video-130517493.cos.ap-guangzhou.myqcloud.com/"
	url2, err = UploadVideoFile("bear.mp4")
	if err != nil {
		error.Error(err)
	}
}
