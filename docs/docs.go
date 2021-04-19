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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/guest-registration": {
            "post": {
                "description": "Sign in like a guest (without progress saving)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Guest sign in",
                "parameters": [
                    {
                        "description": "Body for guest sign in",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GuestSignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "Sign in with login and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign in",
                "parameters": [
                    {
                        "description": "Body for sign in",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/api/novel/create": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Create a new novel with title and content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Novels"
                ],
                "summary": "Create a new novel",
                "parameters": [
                    {
                        "description": "body for a new novel creation",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateNovelRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NovelResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/api/novel/delete": {
            "delete": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Delete user novels by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Novels"
                ],
                "summary": "Delete novel",
                "parameters": [
                    {
                        "type": "string",
                        "description": "novel id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/api/novel/list": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Get novels list by search parameter, sorting and pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Novels"
                ],
                "summary": "Novels list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search by any fields in datagrid",
                        "name": "search",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name of sorting field",
                        "name": "sort_field",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "asc or desc",
                        "name": "sort_order",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "page number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "limit of items on page",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.NovelResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/api/novel/update": {
            "put": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Update novel title or data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Novels"
                ],
                "summary": "Update a novel",
                "parameters": [
                    {
                        "description": "body for a novel updating",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateNovelRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NovelResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/api/registration": {
            "post": {
                "description": "User registraton by email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign up",
                "parameters": [
                    {
                        "description": "Body for sign up",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/api/reset_password_request": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Reset your account password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Reset password request",
                "parameters": [
                    {
                        "description": "email for reset password",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/api/socials-login": {
            "post": {
                "description": "User login by socials (Facebook, Google, Apple, etc.). If user doesn't exist in DB, new account will be created.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Socials sign in",
                "parameters": [
                    {
                        "description": "body for sign up",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SocialsSignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/api/user-info": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Get user info by user ID from bearer token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Get user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ErrResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "message": {
                    "type": "string",
                    "example": "INTERNAL_SERVER_ERROR"
                }
            }
        },
        "dto.AuthResponse": {
            "type": "object",
            "properties": {
                "avatarData": {
                    "type": "string",
                    "example": "avatar_data"
                },
                "dateOfBith": {
                    "type": "integer",
                    "example": 12345672
                },
                "email": {
                    "type": "string",
                    "example": "my@testmail.com"
                },
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "id": {
                    "type": "string",
                    "example": "some_id"
                },
                "membership": {
                    "type": "string",
                    "example": "some_info"
                },
                "rate": {
                    "type": "integer",
                    "example": 0
                },
                "token": {
                    "type": "string",
                    "example": "someSuperseCretToken.ForuseRAuthoriZATIon"
                },
                "username": {
                    "type": "string",
                    "example": "awesome_user"
                }
            }
        },
        "dto.CreateNovelRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string",
                    "example": "My awesome true story!"
                },
                "title": {
                    "type": "string",
                    "example": "My new novel"
                }
            }
        },
        "dto.GuestSignInRequest": {
            "type": "object",
            "properties": {
                "deviceId": {
                    "type": "string",
                    "example": "thisIsMyDeviceId"
                }
            }
        },
        "dto.NovelResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "integer",
                    "example": 121342424
                },
                "data": {
                    "type": "string",
                    "example": "My awesome true story!"
                },
                "id": {
                    "type": "string",
                    "example": "some_id"
                },
                "participatedInCompetition": {
                    "type": "boolean",
                    "example": false
                },
                "title": {
                    "type": "string",
                    "example": "My new novel"
                },
                "updatedAt": {
                    "type": "integer",
                    "example": 1654726235
                }
            }
        },
        "dto.ResetPasswordRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "myemail@mail.com"
                }
            }
        },
        "dto.SignInRequest": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string",
                    "example": "test_login"
                },
                "password": {
                    "type": "string",
                    "example": "supersecretpassword"
                }
            }
        },
        "dto.SignUpRequest": {
            "type": "object",
            "properties": {
                "deviceId": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.SocialsSignInRequest": {
            "type": "object",
            "properties": {
                "deviceId": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "social": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateNovelRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string",
                    "example": "My awesome true story!"
                },
                "id": {
                    "type": "string",
                    "example": "some_id"
                },
                "title": {
                    "type": "string",
                    "example": "My new novel"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Version:     "0.0.2",
	Host:        "localhost:8000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Novels REST API",
	Description: "REST API for Novels app.\nNew in version:<br> - socials sign up was deleted. Now we have 1 endpoint for signin/signup.<br> - some minor fixes",
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
