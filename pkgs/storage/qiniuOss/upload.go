package qiniuOss

import (
	"bytes"
	"context"
	"errors"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"path/filepath"
)

// IQiNiuOssUpload 七牛云上传文件
type IQiNiuOssUpload interface {
	// UploadResourceByte 上传文件([]byte)
	UploadResourceByte(param *UploadResourceByteParam) (string, error)
	// UploadLocalFile 上传本地文件
	UploadLocalFile(param *UploadLocalFileParam) (string, error)
}

// QiNiuOssUpload 又拍云上传文件
type QiNiuOssUpload struct {
	AccessKey string `json:"accessKey"` // AccessKey
	SecretKey string `json:"secretKey"` // SecretKey
	Pipeline  string `json:"pipeline"`  // Pipeline
}

// UploadResourceByte 上传文件([]byte)
func (h *QiNiuOssUpload) UploadResourceByte(param *UploadResourceByteParam) (string, error) {
	key := param.SavePath + param.FileName
	mac := qbox.NewMac(h.AccessKey, h.SecretKey)

	cfg := storage.Config{
		UseHTTPS: false,
		Region:   &storage.Zone_z0,
	}

	// 强制重新执行数据处理任务
	putPolicy := storage.PutPolicy{
		Scope:               param.Bucket,
		PersistentNotifyURL: "http://49b6-69-172-67-65.ngrok.io",
		PersistentPipeline:  h.Pipeline,
	}

	upToken := putPolicy.UploadToken(mac)
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(param.ResourceByte), int64(len(param.ResourceByte)), &putExtra)
	if err != nil {
		return "", err
	}

	if ret.Hash == "" {
		return "", errors.New("上传文件失败")
	}

	return key, nil
}

// UploadLocalFile 上传本地文件
func (h *QiNiuOssUpload) UploadLocalFile(param *UploadLocalFileParam) (string, error) {
	key := ""
	if param.FileName == "" {
		_, fileName := filepath.Split(param.LocalFilePath)
		key = param.SavePath + "/" + fileName
	} else {
		key = param.SavePath + "/" + param.FileName
	}

	mac := qbox.NewMac(h.AccessKey, h.SecretKey)

	cfg := storage.Config{
		UseHTTPS: false,
		Region:   &storage.Zone_z0,
	}

	putPolicy := storage.PutPolicy{
		Scope:               param.Bucket,
		PersistentNotifyURL: "http://49b6-69-172-67-65.ngrok.io",
		PersistentPipeline:  h.Pipeline,
	}
	upToken := putPolicy.UploadToken(mac)
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	err := formUploader.PutFile(context.Background(), &ret, upToken, key, param.LocalFilePath, &putExtra)
	if err != nil {
		return "", err
	}

	if ret.Hash == "" {
		return "", errors.New("上传文件失败")
	}

	return key, nil
}

// UploadResourceByteParam 上传文件([]byte)请求参数
type UploadResourceByteParam struct {
	ResourceByte []byte `json:"resource_byte"` // ResourceByte 文件[]byte流
	FileName     string `json:"fileName"`      // FileName 文件名
	SavePath     string `json:"save_path"`     // SavePath 保存目录
	Bucket       string `json:"bucket"`        // Bucket 存放空间
}

// UploadLocalFileParam 上传本地文件请求参数
type UploadLocalFileParam struct {
	LocalFilePath string `json:"local_file_path" example:"ltest"` // LocalFilePath 本地文件路径(最前面不要带 / )
	SavePath      string `json:"save_path"`                       // SavePath 保存目录
	Bucket        string `json:"bucket"`                          // Bucket 存放空间
	FileName      string `json:"file_name"`                       // FileName 文件名
}
