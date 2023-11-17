package Service

import (
	"webBlog/app/Models"
	"errors"
)

func Login(UserName, PassWord string) (uint, error) {

	user := Models.Users_xiazhaoxuan{
		Username: UserName, Password: PassWord,
	}
	if err := user.Signin(); err != nil {
		return 0, err
	}
	if user.Lockstate == 1 {
		return 0, errors.New("该账户已被限制登录")
	}
	return user.ID, nil
}

func Register(UserName, PassWord, Nickname string) error {

	user := Models.Users_xiazhaoxuan{
		Username: UserName, Nickname: Nickname, Password: PassWord,
	}
	if err := user.Register(); err != nil {
		return err
	}
	return nil
}

func GetUserBaseInfo(user_id int) (user *Models.UserInfo_xiazhaoxuan, err error) {
	user = &Models.UserInfo_xiazhaoxuan{ID: user_id}
	err = user.GetUserBaseInfo()
	if err != nil {
		return
	}
	return
}

func GetUserInfo(user_id int) (user Models.Users_xiazhaoxuan, err error) {
	user.ID = uint(user_id)
	err = user.GetUserInfo()
	return user, err
}