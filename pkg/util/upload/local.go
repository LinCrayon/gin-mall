package upload

import (
	"fmt"
	"gin-mall/config"
	"gin-mall/pkg/util/log"
	"io/ioutil"

	"mime/multipart"
	"os"
	"strconv"
)

// ProductUploadToLocalStatic 上传到本地文件中
func ProductUploadToLocalStatic(file multipart.File, bossId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(bossId))
	basePath := "." + config.AvatarPath + "boss" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := fmt.Sprintf("%s%s.jpg", basePath, productName)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.LogrusObj.Error(err)
		return "", err
	}
	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil {
		log.LogrusObj.Error(err)
		return "", err
	}
	return fmt.Sprintf("boss%s/%s.jpg", bId, productName), err
}

// AvatarUploadToLocalStatic 上传头像
func AvatarUploadToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId)) //int--->string
	basePath := "." + config.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := fmt.Sprintf("%s%s.jpg", basePath, userName)
	/*ReadAll(file) 把客户端上传的文件的所有内容读取到内存中，以便后续的处理。*/
	content, err := ioutil.ReadAll(file) //读取上传文件的内容，返回切片
	if err != nil {
		log.LogrusObj.Error(err)
		return "", err
	}
	//从客户端上传的文件内容写入到本地的头像文件中。
	err = ioutil.WriteFile(avatarPath, content, 0666) //将切片写入文件
	if err != nil {
		log.LogrusObj.Error(err)
		return "", err
	}
	return fmt.Sprintf("user%s/%s.jpg", bId, userName), err
}

// DirExistOrNot 检查指定路径下的文件是否存在，并判断该文件是否是一个目录
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr) //获取指定文件的信息
	if err != nil {
		fmt.Println(err)
		return false
	}
	return s.IsDir() //判断文件是否为目录
}

// CreateDir 创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

/*
所有者（Owner）权限：读、写、执行（rwx，二进制 111）。
同组用户（Group）权限：读、执行（r-x，二进制 101）。
其他用户（Other）权限：读、执行（r-x，二进制 101）。
*/
