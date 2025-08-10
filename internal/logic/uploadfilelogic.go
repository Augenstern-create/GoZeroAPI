package logic

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"GoZeroAPI/internal/svc"
	"GoZeroAPI/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UploadFileLogic 用于处理文件上传业务逻辑
type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUploadFileLogic 构造函数
func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UploadFile 接收并保存上传文件
func (l *UploadFileLogic) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (*types.UploadResp, error) {
	// 1. 确保上传目录存在
	saveDir := "uploads"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("创建上传目录失败: %v", err)
	}

	// 2. 拼接保存路径
	savePath := filepath.Join(saveDir, fileHeader.Filename)

	// 3. 创建文件
	dst, err := os.Create(savePath)
	if err != nil {
		return nil, fmt.Errorf("保存文件失败: %v", err)
	}
	defer dst.Close()

	// 4. 将上传文件内容写入目标文件
	if _, err := dst.ReadFrom(file); err != nil {
		return nil, fmt.Errorf("写入文件失败: %v", err)
	}

	// 5. 返回文件保存成功信息
	return &types.UploadResp{
		Message: "文件上传成功",
		Path:    savePath,
	}, nil
}
