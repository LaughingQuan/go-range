SET CHARACTER_SET_CLIENT = utf8;
SET CHARACTER_SET_CONNECTION = utf8;
CREATE DATABASE goat_db;
CREATE TABLE goat_db.vulnerability (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `severity` bigint DEFAULT NULL,
  `desc` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;


INSERT INTO goat_db.vulnerability(name,severity,`desc`) VALUES ("不安全的加密算法", 1, "应用使用了不安全的或加密强度弱的加密算法，使被加密数据有可能被攻击者破解。");
INSERT INTO goat_db.vulnerability(name,severity,`desc`) VALUES ("违反信任边界规则", 2, "应用将未处理的外部数据被存到可信区域中，如Session中。");
INSERT INTO goat_db.vulnerability(name,severity,`desc`) VALUES ("XStream反序列化", 3, "XStream可以将对象序列化成XML或将XML反序列化为对象。在使用XStream进行反序列化时，如果程序在对外部数据反序列化时，没有校验，会导致反序列化漏洞。");
INSERT INTO goat_db.vulnerability(name,severity,`desc`) VALUES ("Headers敏感信息泄露", 0, "网站请求响应中Headers的字段会显示网站web中间件或开发框架版本信息。");
INSERT INTO goat_db.vulnerability(name,severity,`desc`) VALUES ("服务端模板注入漏洞", 3, "服务端模板注入和常见Web注入的成因一样，也是服务端接收了用户的输入，将其作为 Web 应用模板内容的一部分，在进行目标编译渲染的过程中，执行了用户插入的恶意内容，因而可能导致了敏感信息泄露、代码执行、GetShell等问题，其影响范围主要取决于模版引擎的复杂性。");
INSERT INTO goat_db.vulnerability(name,severity,`desc`) VALUES ("不安全的加密算法", 1, "应用使用了不安全的或加密强度弱的加密算法，使被加密数据有可能被攻击者破解。");
INSERT INTO goat_db.vulnerability(name,severity,`desc`) VALUES ("违反信任边界规则", 2, "应用将未处理的外部数据被存到可信区域中，如Session中。");
INSERT INTO goat_db.vulnerability(name,severity,`desc`) VALUES ("XStream反序列化", 3, "XStream可以将对象序列化成XML或将XML反序列化为对象。在使用XStream进行反序列化时，如果程序在对外部数据反序列化时，没有校验，会导致反序列化漏洞。");
INSERT INTO goat_db.vulnerability(name,severity,`desc`) VALUES ("Headers敏感信息泄露", 0, "网站请求响应中Headers的字段会显示网站web中间件或开发框架版本信息。");
INSERT INTO goat_db.vulnerability(name,severity,`desc`) VALUES ("服务端模板注入漏洞", 3, "服务端模板注入和常见Web注入的成因一样，也是服务端接收了用户的输入，将其作为 Web 应用模板内容的一部分，在进行目标编译渲染的过程中，执行了用户插入的恶意内容，因而可能导致了敏感信息泄露、代码执行、GetShell等问题，其影响范围主要取决于模版引擎的复杂性。");


