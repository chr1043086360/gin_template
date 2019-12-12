package qiniu

import (
	"context"
	"fmt"
	// "project2019/qiniu"

	"github.com/qiniu/api.v7/storage"
)

var RET string
var HASH string

func UpFile() {
	localFile := "/home/chr/godemo/project2019/static/1.jpg"
	// bucket := "cuiheran"
	key := "1.jpg"

	// putPolicy := storage.PutPolicy{
	// 	Scope: bucket,
	// }
	// mac := qbox.NewMac(os.Getenv("QINIU_AK"), os.Getenv("QINIU_SK"))
	// upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, UPTOkEN, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	RET = ret.Key
	HASH = ret.Hash
	fmt.Println(RET, HASH)
	// fmt.Println(UPTOkEN)
}
