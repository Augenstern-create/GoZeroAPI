package handler

import (
	"net/http"

	"GoZeroAPI/internal/logic"
	"GoZeroAPI/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// UploadFileHandler 处理 POST /upload 请求
func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 解析 multipart/form-data 文件
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		defer file.Close()

		// 2. 调用业务逻辑
		l := logic.NewUploadFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadFile(file, fileHeader)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		// 3. 返回上传成功结果
		httpx.OkJson(w, resp)
	}
}
