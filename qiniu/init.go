package qiniu

import (
	"fmt"
	"os"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var UPTOkEN string

func QiniuAuth() {
	accessKey := os.Getenv("QINIU_AK")
	secretKey := os.Getenv("QINIU_SK")
	mac := qbox.NewMac(accessKey, secretKey)
	bucket := "cuiheran"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 // 2小时有消息

	upToken := putPolicy.UploadToken(mac)
	UPTOkEN = upToken
	fmt.Println("qiniu conn success")
}
