package config

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v3"
	"xmirror.cn/iast/goat/util"
)

const CONFIGNAME string = "config.yml"

type Conf struct {
	Mong  MogConf
	Ldap  LdapConf
	Mysql MysqlConf
}
type MogConf struct {
	Host       string
	Port       string
	Database   string
	Collection string
	Username   string
	Password   string
}
type LdapConf struct {
	Url      string
	Port     string
	UserName string
	Password string
	BaseDN   string
}

type MysqlConf struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func GetConfMog() (MogConf, error) {
	conf, err := parseConfig()
	if err != nil {
		return conf.Mong, err
	}
	if reflect.DeepEqual(conf.Mong, MogConf{}) {
		return conf.Mong, errors.New("fhe configuration item was not resolved  Procedure")
	}
	return conf.Mong, err
}

func GetConfLdap() (LdapConf, error) {
	conf, err := parseConfig()
	if err != nil {
		return conf.Ldap, err
	}
	if reflect.DeepEqual(conf.Ldap, LdapConf{}) {
		return conf.Ldap, errors.New("fhe configuration item was not resolved  Procedure")
	}
	return conf.Ldap, err
}

func GetConfMysql() (MysqlConf, error) {
	conf, err := parseConfig()
	if err != nil {
		return conf.Mysql, err
	}
	if reflect.DeepEqual(conf.Ldap, LdapConf{}) {
		return conf.Mysql, errors.New("fhe configuration item was not resolved  Procedure")
	}
	return conf.Mysql, err
}

func parseConfig() (Conf, error) {

	exPath := util.StartDir()
	configPath, err := filepath.Abs(filepath.Join(exPath, CONFIGNAME))
	if err != nil {
		panic(err)
	}
	var conf Conf
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return conf, err
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
