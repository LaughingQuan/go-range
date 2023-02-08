package ldaptargets

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap/v3"
	"xmirror.cn/iast/goat/config"
	"xmirror.cn/iast/goat/util"
)

var confLdap config.LdapConf
var once sync.Once

func SearchSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		once.Do(initData)
		input := util.ExtractInput(c, "input")
		url := fmt.Sprintf("ldap://%s:%s", confLdap.Url, confLdap.Port)
		l, err := ldap.DialURL(url)
		if err != nil {
			log.Println(err)
		}
		defer l.Close()
		l.SimpleBind(&ldap.SimpleBindRequest{
			Username: confLdap.UserName,
			Password: confLdap.Password,
		})
		replacer := strings.NewReplacer(
			"\\", "\\5c",
			"*", "\\2a",
			"(", "\\28",
			")", "\\29",
			"\\x00", "\\00",
		)
		input = replacer.Replace(input)
		filter := fmt.Sprintf("(&(cn=%s))", input)
		searchRequest := ldap.NewSearchRequest(
			confLdap.BaseDN, //					 The base dn to search
			ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
			filter,               // The filter to apply
			[]string{"dn", "cn"}, // A list attributes to retrieve
			nil,
		)
		sr, err := l.Search(searchRequest) ///
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"Entries": sr.Entries,
			},
		})
	}
}

func SearchUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		once.Do(initData)
		input := util.ExtractInput(c, "input")
		url := fmt.Sprintf("ldap://%s:%s", confLdap.Url, confLdap.Port)
		l, err := ldap.DialURL(url)
		if err != nil {
			log.Println(err)
		}
		defer l.Close()
		l.SimpleBind(&ldap.SimpleBindRequest{
			Username: confLdap.UserName,
			Password: confLdap.Password,
		})
		searchRequest := ldap.NewSearchRequest(
			confLdap.BaseDN, //					 The base dn to search
			ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
			fmt.Sprintf("(&(cn=%s))", input), // The filter to apply
			[]string{"dn", "cn"},             // A list attributes to retrieve
			nil,
		)
		sr, err := l.Search(searchRequest) ///
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"Entries": sr.Entries,
			},
		})
	}

}
