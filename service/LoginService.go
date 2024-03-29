package service

import (
	"github.com/songzhonghuasongzhonghua/gogodancing/dao"
	"github.com/songzhonghuasongzhonghua/gogodancing/model"
	"github.com/songzhonghuasongzhonghua/gogodancing/param"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
	"log"
)

type LoginEngine struct {
	*tool.Orm
}

func LoginService(param *param.LoginParam) (*model.Member, error) {

	db := LoginEngine{tool.DbOrm}
	member := new(model.Member)
	bool, err := db.Alias("l").Where("l.mobile = ? and l.password", param.Mobile, param.Password).Get(member)
	if err != nil {
		return nil, err
	}
	if bool {
		return member, nil
	} else {
		insertParam := new(model.Member)
		insertParam.Mobile = param.Mobile
		insertParam.Password = param.Password

		result, err := dao.LoginDao(insertParam)
		if result > 0 {
			dao.LoginDao(insertParam)
			return insertParam, nil
		} else {
			return nil, err
		}

	}
}

func LoginPwdService(name string, password string) *model.Member {

	member := dao.QueryMember(name, password)
	//如果存在
	if member.Id != 0 {
		return member
	}

	newMember := new(model.Member)
	newMember.UserName = name
	newMember.Password = password
	//如果不存在则插入
	err := dao.InsertMember(newMember)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return newMember
}
