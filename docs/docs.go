// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/zouyonghe",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/zouyonghe",
            "email": "1259085392z@gmail.com"
        },
        "license": {
            "name": "GPLv3",
            "url": "https://www.gnu.org/licenses/gpl-3.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/admin": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List all users account include id, username, encrypted password, etc",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/admin"
                ],
                "summary": "List all users account",
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"totalCount\":1,\"userList\":[{\"id\":1,\"username\":\"admin\",\"ShortId\":\"5P9Ia4QnR\",\"password\":\"$2a$10$Fv9BWzqsiQ.JuuGdcXdvN.Fx3ml.dVR47W22GoJMWQAlm9wHQIMVe\",\"role\":\"admin\",\"createdAt\":\"2021-04-18 15:40:33\",\"updatedAt\":\"2021-04-18 15:40:33\"}]}}",
                        "schema": {
                            "$ref": "#/definitions/user.SwaggerListResponse"
                        }
                    }
                }
            }
        },
        "/user/admin/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a user account specified by user ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/admin"
                ],
                "summary": "Get a user information",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "the ID of the specified user to update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"username\":\"傅秀英\",\"password\":\"$2a$10$5pLrLpEQ1HAD2Hcm3Bnud.Shhmf5bTaf1yTWYloot0i5nvn1Td4hq\",\"role\":\"general\"}}",
                        "schema": {
                            "$ref": "#/definitions/user.SwaggerGetResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a user account specified by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/admin"
                ],
                "summary": "Update a user account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "the ID of the specified user to update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user information include username and password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"userId\":5}}",
                        "schema": {
                            "$ref": "#/definitions/user.SwaggerUpdateResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete a user by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/admin"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "the ID of the specified user to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"userId\":5}}",
                        "schema": {
                            "$ref": "#/definitions/user.SwaggerDeleteResponse"
                        }
                    }
                }
            }
        },
        "/user/common": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update the current user information by username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/common"
                ],
                "summary": "Update the current user information",
                "parameters": [
                    {
                        "description": "Create a new user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.SelfUpdRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"userId\":6,\"username\":\"夏秀兰\"}}",
                        "schema": {
                            "$ref": "#/definitions/user.SwaggerSelfUpdResponse"
                        }
                    }
                }
            }
        },
        "/user/common/": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "SelfDel deletes the user of token specified",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/common"
                ],
                "summary": "SelfDel deletes the user of token specified",
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"userId\":8}}",
                        "schema": {
                            "$ref": "#/definitions/user.SwaggerSelfDelResponse"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login a user account with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login  a user account",
                "parameters": [
                    {
                        "description": "Login account",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"userId\":7,\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NTA0NTkzODEsImlkIjo3LCJuYmYiOjE2NTA0NTkzODEsInJvbGUiOiJnZW5lcmFsIiwidXNlcm5hbWUiOiLkuIHno4oifQ.0kA4whaE9bZjXu4bN3Sw0DgrKwYzJ7kZenaGDOcdFRQ\"}}",
                        "schema": {
                            "$ref": "#/definitions/user.SwaggerLoginResponse"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Create a new user by username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "user information include username and password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"userId\":\"7\",\"username\":\"顾磊\"}}",
                        "schema": {
                            "$ref": "#/definitions/user.SwaggerCreateResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "ShortId": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.CreateRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.CreateResponse": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.DeleteResponse": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "integer"
                }
            }
        },
        "user.GetResponse": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "user.SelfDelResponse": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "integer"
                }
            }
        },
        "user.SelfUpdRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.SelfUpdResponse": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.SwaggerCreateResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/user.CreateResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "user.SwaggerDeleteResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/user.DeleteResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "user.SwaggerGetResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/user.GetResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "user.SwaggerListResponse": {
            "type": "object",
            "properties": {
                "totalCount": {
                    "type": "integer"
                },
                "userList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UserInfo"
                    }
                }
            }
        },
        "user.SwaggerLoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/user.LoginResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "user.SwaggerSelfDelResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/user.SelfDelResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "user.SwaggerSelfUpdResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/user.SelfUpdResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "user.SwaggerUpdateResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/user.UpdateResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "user.UpdateRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.UpdateResponse": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.3",
	Host:             "127.0.0.1:8081",
	BasePath:         "/v1",
	Schemes:          []string{"https"},
	Title:            "Jinshuzhai-Bookstore",
	Description:      "The jinshuzhai bookstore api server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
