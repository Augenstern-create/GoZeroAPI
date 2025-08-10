package handler

import (
	"net/http"
	"strconv"

	"GoZeroAPI/internal/logic"
	"GoZeroAPI/internal/svc"
	"GoZeroAPI/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// CreateUserHandler 处理 POST /users 请求
func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateUser(&req)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		httpx.OkJson(w, resp)
	}
}

// GetUserHandler 处理 GET /users/:id 请求
func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserIdReq
		if err := httpx.ParsePath(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		httpx.OkJson(w, resp)
	}
}

// UpdateUserHandler 处理 PUT /users/:id 请求
func UpdateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		var pathReq types.UserIdReq
		if err := httpx.ParsePath(r, &pathReq); err != nil {
			httpx.Error(w, err)
			return
		}

		id, err := strconv.ParseInt(pathReq.Id, 10, 64)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUser(id, &req)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		httpx.OkJson(w, resp)
	}
}

// DeleteUserHandler 处理 DELETE /users/:id 请求
func DeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserIdReq
		if err := httpx.ParsePath(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		id, err := strconv.ParseInt(req.Id, 10, 64)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
		err = l.DeleteUser(id)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		httpx.Ok(w) // 返回空 200 OK
	}
}

// ListUsersHandler 处理 GET /users 请求，返回所有用户
func ListUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.ListUsers()
		if err != nil {
			httpx.Error(w, err)
			return
		}

		httpx.OkJson(w, resp)
	}
}
