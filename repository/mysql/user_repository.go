package mysql

import (
	"errors"
	"sync"

	"github.com/micro-community/auth/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//userRepository data
type UserRepository struct {
	mu        *sync.Mutex
	db        *gorm.DB
	tableName string
}

func (u UserRepository) TableName() string {
	if u.tableName == "" {
		u.tableName = "user"
	}
	return u.tableName
}

func (u UserRepository) Table() *gorm.DB {
	return u.Table()
}

// Get 校验获取用户数据
func (u UserRepository) Get(user models.User) error {
	table := u.Table().Select([]string{"user.*", "role.role_name"})

	table = table.Joins("left join role on user.role_id=role.role_id")
	if user.ID != 0 {
		table = table.Where("user_id = ?", user.ID)
	}
	if user.Name != "" {
		table = table.Where("username = ?", user.Name)
	}
	if user.Password != "" {
		table = table.Where("password = ?", user.Password)
	}
	if user.RoleId != 0 {
		table = table.Where("role_id = ?", user.RoleId)
	}
	if user.DeptId != 0 {
		table = table.Where("dept_id = ?", user.DeptId)
	}
	if user.PostionId != 0 {
		table = table.Where("post_id = ?", user.PostionId)
	}
	if err := table.First(&user).Error; err != nil {
		return err
	}
	return nil
}

//加密
func (u *UserRepository) Encrypt(password string) (string, error) {
	if password == "" {
		return "", nil
	}
	var hash []byte
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hash), nil

}

//Insert 添加 user
func (u UserRepository) Insert(user models.User) (id int64, err error) {
	encryptedPassword, err := u.Encrypt(user.Password)

	if encryptedPassword == "" || err != nil {
		return -1, errors.New("password encrypted error")
	}
	// check 用户名
	var count int64
	u.Table().Where("username = ?", user.Name).Count(&count)
	if count > 0 {
		err = errors.New("user account exist")
		return
	}

	//添加数据
	if err = u.Table().Create(&u).Error; err != nil {
		return -1, err
	}
	return
}

//Update 修改
func (u *UserRepository) Update(user models.User) (updatedUser models.User, err error) {

	var encryptedPassword string
	if user.Password != "" {
		encryptedPassword, err = u.Encrypt(user.Password)
		if encryptedPassword == "" || err != nil {
			return models.User{}, errors.New("password encrypted error")
		}
	}
	if err = u.Table().First(&updatedUser, user.ID).Error; err != nil {
		return
	}

	if err = u.Table().Model(&updatedUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

func (u *UserRepository) BatchDelete(id []int) (result bool, err error) {
	if err = u.Table().Where("user_id in (?)", id).Delete(&models.User{}).Error; err != nil {
		return
	}
	result = true
	return
}
