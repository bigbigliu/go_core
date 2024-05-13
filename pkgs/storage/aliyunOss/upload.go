package aliyunOss

import (
	"bytes"
	"crypto/tls"
	"net/http"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// IAliyunOssUpload 阿里云oss上传方法
type IAliyunOssUpload interface {
	// UploadResourceByte 上传Byte数组
	UploadResourceByte(param *UploadResourceByteReq) (path string, err error)
	// UploadLocalFile 上传本地文件
	UploadLocalFile(param *UploadLocalFileReq) (path string, err error)
}

// AliyunOss ...
type AliyunOssUpload struct{}

// UploadResourceByte 上传Byte数组
func (h *AliyunOssUpload) UploadResourceByte(param *UploadResourceByteReq) (path string, err error) {
	resourceByte := param.UploadConf.ResourceByte
	if err != nil {
		return "", err
	}

	client, err := oss.New("https://oss-cn-hangzhou.aliyuncs.com",
		param.OssConf.AccessKeyId,
		param.OssConf.AccessKeySecret,
		oss.HTTPClient(&http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}))
	if err != nil {
		return "", err
	}

	bucket, err := client.Bucket(param.OssConf.Bucket)
	if err != nil {
		return "", err
	}

	storagePath := param.UploadConf.SavePath + param.UploadConf.FileName
	err = bucket.PutObject(storagePath, bytes.NewReader(resourceByte))
	if err != nil {
		return "", err
	}

	return storagePath, nil
}

// UploadLocalFile 上传本地文件
func (h *AliyunOssUpload) UploadLocalFile(param *UploadLocalFileReq) (path string, err error) {
	client, err := oss.New("https://oss-cn-hangzhou.aliyuncs.com",
		param.OssConf.AccessKeyId,
		param.OssConf.AccessKeySecret,
		oss.HTTPClient(&http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}))
	if err != nil {
		return "", err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client.HTTPClient.Transport = tr

	bucket, err := client.Bucket(param.OssConf.Bucket)
	if err != nil {
		return "", err
	}

	err = bucket.PutObjectFromFile(param.UploadConf.SavePath, param.UploadConf.LocalPath)
	if err != nil {
		return "", err
	}

	return param.UploadConf.SavePath, nil
}

// UploadResourceByteReq ...
type UploadResourceByteReq struct {
	OssConf    *AliyunOssConf `json:"oss_conf"`    // OssConf ...
	UploadConf *UploadFile    `json:"upload_conf"` // UploadConf ...
}

// UploadFile ...
type UploadFile struct {
	ResourceByte []byte `json:"resource_byte"`
	FileName     string `json:"fileName"`
	SavePath     string `json:"save_path"`
}

// UploadLocalFileReq 上传本地文件
type UploadLocalFileReq struct {
	OssConf    *AliyunOssConf `json:"oss_conf"`    // OssConf ...
	UploadConf *LocalFile     `json:"upload_conf"` // UploadConf ...
}

// LocalFile ...
type LocalFile struct {
	LocalPath string `json:"local_path"`
	FileName  string `json:"fileName"`
	SavePath  string `json:"save_path"`
}

// AliyunOssConf 阿里云oss配置
type AliyunOssConf struct {
	AccessKeyId     string `json:"AccessKeyId"`     // AccessKeyId
	AccessKeySecret string `json:"AccessKeySecret"` // AccessKeySecret
	Bucket          string `json:"bucket"`          // Bucket
}
