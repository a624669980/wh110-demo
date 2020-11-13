// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/wh/getfriendlist": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wh"
                ],
                "summary": "查询友链列表",
                "parameters": [
                    {
                        "enum": [
                            1,
                            2
                        ],
                        "type": "integer",
                        "description": "项目类型",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/wh/getlist": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wh"
                ],
                "summary": "查询新闻列表",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "页码,默认为1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "每页数量,默认为10",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "enum": [
                            1,
                            2
                        ],
                        "type": "integer",
                        "description": "项目类型",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "是否为升序",
                        "name": "isasc",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/wh/getnewsdetailbyid": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wh"
                ],
                "summary": "查询新闻详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "新闻id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/wh/getnewsdetailbyorder": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wh"
                ],
                "summary": "查询新闻详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "新闻order",
                        "name": "order",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/wh/getrrnews": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wh"
                ],
                "summary": "根据新闻id获取随机推荐",
                "parameters": [
                    {
                        "type": "string",
                        "description": "新闻标识",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "推荐的数量",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "enum": [
                            1,
                            2
                        ],
                        "type": "integer",
                        "description": "项目类型",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/wh/gettoplist": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wh"
                ],
                "summary": "查询新闻推荐列表",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "页码,默认为1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "每页数量,默认为10",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "enum": [
                            1,
                            2
                        ],
                        "type": "integer",
                        "description": "项目类型",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/wh/getudnews": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wh"
                ],
                "summary": "根据新闻id获取上下篇",
                "parameters": [
                    {
                        "type": "string",
                        "description": "新闻标识",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            1,
                            2
                        ],
                        "type": "integer",
                        "description": "项目类型",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/wh/getudnewsbo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wh"
                ],
                "summary": "根据新闻order获取上下篇",
                "parameters": [
                    {
                        "type": "string",
                        "description": "新闻标识",
                        "name": "order",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            1,
                            2
                        ],
                        "type": "integer",
                        "description": "项目类型",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample Server pets",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
