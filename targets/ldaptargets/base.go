package ldaptargets

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
	"xmirror.cn/iast/goat/config"
)

func initData() {
	var err error
	confLdap, err = config.GetConfLdap()
	if err != nil {
		log.Println(err)
	}
	url := fmt.Sprintf("ldap://%s:%s", confLdap.Url, confLdap.Port)
	l, err := ldap.DialURL(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()
	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: confLdap.UserName,
		Password: confLdap.Password,
	})
	if err != nil {
		log.Println(err)
		return
	}

	username := "zhangsan"
	UserSearchBase := confLdap.BaseDN
	deleteRequest := ldap.NewDelRequest(fmt.Sprintf("uid=%s,%s", username, UserSearchBase), nil)

	l.Del(deleteRequest)

	userCreateRequest := ldap.NewAddRequest(fmt.Sprintf("uid=%s,%s", username, UserSearchBase), nil)
	userCreateRequest.Attribute("objectClass", []string{"inetOrgPerson", "posixAccount", "top"})
	userCreateRequest.Attribute("sn", []string{"default"})
	userCreateRequest.Attribute("cn", []string{username})
	userCreateRequest.Attribute("uidNumber", []string{"123213"})
	userCreateRequest.Attribute("gidNumber", []string{"500"})
	userCreateRequest.Attribute("homeDirectory", []string{"/home/test"})
	userCreateRequest.Attribute("mail", []string{"test@123.com"})
	err = l.Add(userCreateRequest)
	if err != nil {
		fmt.Println(err)
	}
}
