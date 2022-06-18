package object

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var (
	imageBucket = "https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/"
	videoBucket = "https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/"
	filePath    = "E:/golang/tiktok-lite/public/"
)

func getCos(bucketUrl string) *cos.Client {
	var u, _ = url.Parse(bucketUrl)
	var b = &cos.BaseURL{BucketURL: u}
	var c = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 COS_SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: SECRETID,
			// 环境变量 COS_SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: SECRETKEY,
			// Debug 模式，把对应 请求头部、请求内容、响应头部、响应内容 输出到标准输出
			//Transport: &debug.DebugRequestTransport{
			//	RequestHeader: true,
			//	// Notice when put a large file and set need the request body, might happen out of memory error.
			//	RequestBody:    false,
			//	ResponseHeader: true,
			//	ResponseBody:   false,
			//},
		},
	})

	return c
}

func logStatus(err error) {
	if err == nil {
		return
	}
	if cos.IsNotFoundError(err) {
		// WARN
		fmt.Println("WARN: Resource is not existed")
	} else if e, ok := cos.IsCOSError(err); ok {
		fmt.Printf("ERROR: Code: %v\n", e.Code)
		fmt.Printf("ERROR: Message: %v\n", e.Message)
		fmt.Printf("ERROR: Resource: %v\n", e.Resource)
		fmt.Printf("ERROR: RequestId: %v\n", e.RequestID)
		// ERROR
	} else {
		fmt.Printf("ERROR: %v\n", err)
		// ERROR
	}
}

//func example() {
//	c := getCos(imageBucket)
//
//	// Case1 上传对象
//	name := "test/example"
//	f := strings.NewReader("test")
//
//	_, err := c.Object.Put(context.Background(), name, f, nil)
//	logStatus(err)
//
//	// Case2 使用options上传对象
//	f = strings.NewReader("test xxx")
//	opt := &cos.ObjectPutOptions{
//		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
//			ContentType: "text/html",
//		},
//		ACLHeaderOptions: &cos.ACLHeaderOptions{
//			//XCosACL: "public-read",
//			XCosACL: "private",
//		},
//	}
//	_, err = c.Object.Put(context.Background(), name, f, opt)
//	logStatus(err)
//
//	// Case3 通过本地文件上传对象
//	_, err = c.Object.PutFromFile(context.Background(), name, "./test", nil)
//	logStatus(err)
//
//	// Case4 查看上传进度
//	opt.ObjectPutHeaderOptions.Listener = &cos.DefaultProgressListener{}
//	_, err = c.Object.PutFromFile(context.Background(), name, "./test", opt)
//}

// UploadImgFile 上传图片到cos服务
func UploadImgFile(fileName string) (url string, err error) {
	c := getCos(imageBucket)

	// 通过本地文件上传对象
	_, err = c.Object.PutFromFile(context.Background(), fileName, filePath+fileName, nil)
	if err != nil {
		logStatus(err)
		return "", err
	}
	return imageBucket + fileName, nil
}

// UploadVideoFile 上传视频到cos服务
func UploadVideoFile(fileName string) (url string, err error) {
	c := getCos(videoBucket)
	// 通过本地文件上传对象
	_, err = c.Object.PutFromFile(context.Background(), fileName, filePath+fileName, nil)
	if err != nil {
		logStatus(err)
		return "", err
	}
	return videoBucket + fileName, nil
}
