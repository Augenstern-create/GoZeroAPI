package logic

import (
	"context"
	"errors"

	"GoZeroAPI/internal/svc"
	"GoZeroAPI/internal/types"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// UserModel gorm 用户模型
type UserModel struct {
	Id   int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	db     *gorm.DB
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	db, err := gorm.Open(sqlite.Open("./users.db"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}
	if err = db.AutoMigrate(&UserModel{}); err != nil {
		panic("自动迁移失败: " + err.Error())
	}

	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		db:     db,
	}
}

// CreateUser 创建用户
func (l *UserLogic) CreateUser(req *types.CreateUserReq) (*UserModel, error) {
	user := UserModel{
		Name: req.Name,
		Age:  req.Age,
	}
	if err := l.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser 根据 ID 获取用户
func (l *UserLogic) GetUser(req *types.UserIdReq) (*UserModel, error) {
	var user UserModel
	if err := l.db.First(&user, req.Id).Error; err != nil {
		return nil, errors.New("用户未找到")
	}
	return &user, nil
}

// UpdateUser 根据 ID 更新用户信息
func (l *UserLogic) UpdateUser(id int64, req *types.CreateUserReq) (*UserModel, error) {
	var user UserModel
	if err := l.db.First(&user, id).Error; err != nil {
		return nil, errors.New("用户未找到")
	}

	user.Name = req.Name
	user.Age = req.Age

	if err := l.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser 根据 ID 删除用户
func (l *UserLogic) DeleteUser(id int64) error {
	if err := l.db.Delete(&UserModel{}, id).Error; err != nil {
		return err
	}
	return nil
}

// ListUsers 获取所有用户
func (l *UserLogic) ListUsers() ([]UserModel, error) {
	var users []UserModel
	if err := l.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
