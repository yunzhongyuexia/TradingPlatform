package model

import (
	"gorm.io/gorm"
	"server/db"
	"time"
)

type User struct {
	Id          int64     `json:"id" gorm:"id"`                     // 主键id
	Uid         int64     `json:"uid" gorm:"uid"`                   // 用户uid
	Name        string    `json:"name" gorm:"name"`                 // 用户账号
	Password    string    `json:"password" gorm:"password"`         // 用户密码
	Phone       string    `json:"phone" gorm:"phone"`               // 用户电话
	CreatedTime time.Time `json:"created_time" gorm:"created_time"` // 用户创建时间
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"` // 用户信息更新时间
}

// TableName 表名称
func (*User) TableName() string {
	return "user"
}

func GetUserByNameAndPassword(name, password string) *User {
	var ret User
	err := db.Mysql.Raw("SELECT * FROM user WHERE name = ? and password =? LIMIT 1", name, password).First(&ret).Error
	if err != nil {
		return nil
	}
	return &ret
}

func GetUserByPhone(phone string) *User {
	var ret User
	err := db.Mysql.Raw("SELECT * FROM user WHERE phone = ?  LIMIT 1", phone).First(&ret).Error
	if err != nil {
		return nil
	}
	return &ret
}

func GetUserByName(name string) (*User, error) {
	var ret User
	// 使用 Raw 执行 SQL 查询
	err := db.Mysql.Raw("SELECT * FROM user WHERE name = ?", name).First(&ret).Error
	if err != nil {
		// 检查是否是记录未找到的错误
		if err == gorm.ErrRecordNotFound {
			// 没有找到记录，返回零值 User 的指针
			return &ret, nil
		} else {
			// 其他错误，返回错误
			return nil, err
		}
	}
	// 查询成功且记录存在，返回非零 User 的指针
	return &ret, nil
}

func GetUserByUid(uid int64) (*User, error) {
	var ret User
	err := db.Mysql.Raw("SELECT * FROM user WHERE uid = ?  LIMIT 1", uid).First(&ret).Error
	if err != nil {
		// 检查是否是记录未找到的错误
		if err == gorm.ErrRecordNotFound {
			// 没有找到记录，返回零值 User 的指针
			return &ret, nil
		} else {
			// 其他错误，返回错误
			return nil, err
		}
	}
	// 查询成功且记录存在，返回非零 User 的指针
	return &ret, nil
}

func RegistrationUser(user *User) error {
	// 编写 SQL 插入语句
	query := "INSERT INTO user (uid,name,password,created_time,updated_time) VALUES (?, ?, ?,?,?)"

	// 执行插入操作
	err := db.Mysql.Exec(query, user.Uid, user.Name, user.Password, user.CreatedTime, user.UpdatedTime)
	if err != nil {
		return err.Error
	}
	return nil
}

func UpdateUserPhone(phone string, uid int64) error {
	// 准备 SQL 更新语句
	query := "UPDATE user SET phone = ? WHERE uid = ?"

	// 执行 SQL 更新语句
	err := db.Mysql.Exec(query, phone, uid)
	if err != nil {
		return err.Error
	}
	return nil
}
