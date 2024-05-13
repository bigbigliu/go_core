package upyunOss

import (
	"github.com/upyun/go-sdk/v3/upyun"
	"path/filepath"
)

// IUpYunOssUpload 又拍云上传文件
type IUpYunOssUpload interface {
	// UploadLocalFile 上传本地文件
	UploadLocalFile(param *UploadLocalFileParam) (string, error)
}

// UpYunOssUpload 又拍云上传文件
type UpYunOssUpload struct {
	Operator string `json:"operator"` // Operator
	Password string `json:"password"` // Password
	Secret   string `json:"secret"`   // Secret
}

// UploadLocalFile 上传本地文件
func (h *UpYunOssUpload) UploadLocalFile(param *UploadLocalFileParam) (string, error) {
	upNew := upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   param.Bucket,
		Operator: h.Operator,
		Password: h.Password,
	})

	path := ""
	if param.FileName == "" {
		_, fileName := filepath.Split(param.LocalFilePath)
		path = param.SavePath + "/" + fileName
	} else {
		path = param.SavePath + "/" + param.FileName
	}
	uploadParam := &upyun.PutObjectConfig{
		Path:      path,
		LocalPath: param.LocalFilePath,
	}
	err := upNew.Put(uploadParam)
	if err != nil {
		return "", err
	}

	return uploadParam.Path, nil
}

type UploadLocalFileParam struct {
	Bucket        string // Bucket
	SavePath      string // SavePath 云存储中的保存目录
	LocalFilePath string // LocalFilePath 本地文件路径
	FileName      string // FileName 文件名
}

//type PutObjectConfig struct {
//	Path              string            // 云存储中的路径
//	LocalPath         string            // 待上传文件在本地文件系统中的路径
//	Reader            io.Reader         // 待上传的内容
//	Headers           map[string]string // 额外的 HTTP 请求头
//	UseMD5            bool              // 是否需要 MD5 校验
//	UseResumeUpload   bool              // 是否使用断点续传
//	AppendContent     bool              // 是否需要追加文件内容
//	ResumePartSize    int64             // 断点续传块大小
//	MaxResumePutTries int               // 断点续传最大重试次数
//}
