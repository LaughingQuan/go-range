var config ={
	"info": {
		"_postman_id": "0c2df0d0-f20d-4b11-bbfc-5281de32e7c2",
		"name": "靶机前端",
		"schema": "https://schema.getpostman.com/json/collection/v2.0.0/collection.json"
	},
	"item": [
		{
			"name": "sql查询",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "1 or 1 = 1",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/sql/query/body/unsafe/"
			},
			"response": []
		},
		{
			"name": "sql 执行",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "Robert'; select * from vulnerability;--",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/sql/exec/body/unsafe/"
			},
			"response": []
		},
		{
			"name": "上下文命令执行",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "ifconfig",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/exec/commandctx/body/unsafe/"
			},
			"response": []
		},
		{
			"name": "命令执行",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "ifconfig",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/exec/command/body/safe/"
			},
			"response": []
		},
		{
			"name": "gin框架文件读取",
			"request": {
				"method": "POST",
				"header": [
					{
						"key": "Content-Type",
						"type": "text",
						"value": "form-data",
						"disabled": true
					}
				],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "./../../../../../../../../etc/passwd",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/ginframework/ginfile/body/unsafe/"
			},
			"response": []
		},
		{
			"name": "xss漏洞",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "{{baseUrl}}gin/xss/script/query/unsafe/?input=<script>alert(1)</script>",
					"host": [
						"{{baseUrl}}gin"
					],
					"path": [
						"xss",
						"script",
						"query",
						"unsafe",
						""
					],
					"query": [
						{
							"key": "input",
							"value": "<script>alert(1)</script>"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "模板注入",
			"protocolProfileBehavior": {
				"disableBodyPruning": true
			},
			"request": {
				"method": "GET",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": []
				},
				"url": {
					"raw": "{{baseUrl}}iris/ssti/execute/query/safe?input={{.ID}}用户的用户名是 {{.Username }} 密码是{{.Password}} 电话是{{.Phone}}",
					"host": [
						"{{baseUrl}}iris"
					],
					"path": [
						"ssti",
						"execute",
						"query",
						"safe"
					],
					"query": [
						{
							"key": "input",
							"value": "{{.ID}}用户的用户名是 {{.Username }} 密码是{{.Password}} 电话是{{.Phone}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "ssrf url渗透",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "http://example.com",
							"type": "default"
						}
					]
				},
				"url": "{{baseUrl}}echo/ssrf/request/body/unsafe/"
			},
			"response": []
		},
		{
			"name": "ssrf参数渗透",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "@www.baidu.com",
							"type": "default"
						}
					]
				},
				"url": "{{baseUrl}}gin/ssrf/request/body/unsafe/"
			},
			"response": []
		},
		{
			"name": "重定向",
			"protocolProfileBehavior": {
				"followRedirects": false
			},
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "{{baseUrl}}gin/unvalidated/redirect/query/unsafe/?input=http://example.com",
					"host": [
						"{{baseUrl}}gin"
					],
					"path": [
						"unvalidated",
						"redirect",
						"query",
						"unsafe",
						""
					],
					"query": [
						{
							"key": "input",
							"value": "http://example.com"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "ladp查询漏洞",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "*",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/ldap/search/body/unsafe/"
			},
			"response": []
		},
		{
			"name": "文件夹创建",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "/tmp/Sample",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/file/mkdir/body/safe/"
			},
			"response": []
		},
		{
			"name": "文件创建",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "/tmp/Sample.txt",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/file/openfile/body/safe/"
			},
			"response": []
		},
		{
			"name": "文件上传",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "formdata",
					"formdata": [
						{
							"key": "input",
							"type": "file",
							"src": ""
						}
					]
				},
				"url": "{{baseUrl}}gin/file/download/body/unsafe/"
			},
			"response": []
		},
		{
			"name": "文件删除",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "/tmp/Sample.txt",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/file/remove/body/safe/"
			},
			"response": []
		},
		{
			"name": "文件重命名",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "/tmp/Sample.txt",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/file/rename/body/safe/"
			},
			"response": []
		},
		{
			"name": "iris框架-命令执行",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "ls",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}iris/exec/command/body/safe"
			},
			"response": []
		},
		{
			"name": "iris框架-命令执行-qurey 参数",
			"protocolProfileBehavior": {
				"disableBodyPruning": true
			},
			"request": {
				"method": "GET",
				"header": [],
				"body": {
					"mode": "formdata",
					"formdata": []
				},
				"url": {
					"raw": "{{baseUrl}}iris/exec/command/body/safe?input=ls",
					"host": [
						"{{baseUrl}}iris"
					],
					"path": [
						"exec",
						"command",
						"body",
						"safe"
					],
					"query": [
						{
							"key": "input",
							"value": "ls"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "身份证信息泄露",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "23018219801225165X",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/response-info/id-number/body/unsafe/"
			},
			"response": []
		},
		{
			"name": "手机号泄露",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "15514143533",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/response-info/cell/body/unsafe/"
			},
			"response": []
		},
		{
			"name": "银行卡号泄露",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "6214881179982233",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/response-info/bank-id/body/unsafe/"
			},
			"response": []
		},
		{
			"name": "自定义敏感信息（请输入要检测的信息）",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "敏感信息体",
							"type": "text"
						}
					]
				},
				"url": "{{baseUrl}}gin/response-info/custom/body/safe/"
			},
			"response": []
		},
		{
			"name": "xpath",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "WEB']/../../books[@type='chinese",
							"type": "default"
						}
					]
				},
				"url": "{{baseUrl}}chi/xml/xpath/body/unsafe/"
			},
			"response": []
		},
		{
			"name": "Cookie注入",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "cookie value",
							"type": "default"
						}
					]
				},
				"url": "{{baseUrl}}chi/cookie/cookie/body/unsafe/"
			},
			"response": []
		},
		{
			"name": "Cookie未设置HttpOnly",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "cookieHttpOnly",
							"type": "default"
						}
					]
				},
				"url": "{{baseUrl}}chi/cookie/httponly/body/safe/"
			},
			"response": []
		},
		{
			"name": "不安全的随机数",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "100",
							"type": "default"
						}
					]
				},
				"url": "{{baseUrl}}beego/unsec/body/rand/safe"
			},
			"response": []
		},
		{
			"name": "不安全的哈希",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "input",
							"value": "Hash，一般翻译做散列、杂凑，或音译为哈希.",
							"type": "default"
						}
					]
				},
				"url": "{{baseUrl}}beego/unsec/body/hash/unsafe"
			},
			"response": []
		}
	],
	"event": [
		{
			"listen": "prerequest",
			"script": {
				"type": "text/javascript",
				"exec": [
					""
				]
			}
		},
		{
			"listen": "test",
			"script": {
				"type": "text/javascript",
				"exec": [
					""
				]
			}
		}
	],
	"variable": [
		{
			"key": "baseUrl",
			"value": "http://192.168.172.137:8888/"
		}
	]
}