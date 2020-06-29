package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var endpoint string
var accessKey string
var accessSecret string

func init() {
	endpoint = "https://oss-cn-shenzhen.aliyuncs.com"
	accessKey = "LTAI4G76fco"
	accessSecret = "pOEaamen3J"
}
func main() {
	fmt.Println("OSS Go SDK Version: ", oss.Version)
	client, err := oss.New(endpoint, accessKey, accessSecret)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	// 列举存储空间。
	UploadFile(client, "enkeshu-prd")

}

func ListBucket(client *oss.Client) {
	marker := ""
	for {
		lsRes, err := client.ListBuckets(oss.Marker(marker))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		// 默认情况下一次返回100条记录。
		for _, bucket := range lsRes.Buckets {
			fmt.Println("Bucket: ", bucket.Name)
		}

		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
}

func UploadFile(client *oss.Client, bucketName string) {
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("获取bucke失败")
		os.Exit(-1)
	}

	//storageType := oss.ObjectStorageClass(oss.StorageStandard)
	err = bucket.PutObjectFromFile("2312321231", "D:\\Image\\123.jpg")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func StsToken(client *oss.Client) {

}
