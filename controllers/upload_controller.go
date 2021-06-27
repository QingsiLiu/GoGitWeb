package controllers

import (
	"GoGitWeb/models"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type UploadController struct {
	BaseController
}

func (u *UploadController) Post() {
	fmt.Println("file-uploading...")
	fileData, fileHeader, err := u.GetFile("upload")
	if err != nil {
		u.responseErr(err)
		return
	}
	fmt.Println("name:", fileHeader.Filename, fileHeader.Size)
	fmt.Println(fileData)
	now := time.Now()
	//返回文件的扩展名
	fmt.Println("ext:", filepath.Ext(fileHeader.Filename))
	fileType := "other"
	//判断后缀为图片的文件，存入数据库中
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}

	//文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		u.responseErr(err)
		return
	}

	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathstr := filepath.Join(fileDir, fileName)
	desFile, err := os.Create(filePathstr)
	if err != nil {
		u.responseErr(err)
		return
	}

	//将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		u.responseErr(err)
		return
	}

	if fileType == "img" {
		album := models.Album{0, filePathstr, fileName, 0, timeStamp}
		models.
	}
}

func (u *UploadController) responseErr(err error) {
	u.Data["json"] = map[string]interface{}{"code" : 0, "message" : err}
	u.ServeJSON()
}
