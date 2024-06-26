// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/users/{id}/wallets": {
            "get": {
                "description": "Get user wallets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get user wallets",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "Savings",
                            "Credit Card",
                            "Crypto Wallet"
                        ],
                        "type": "string",
                        "description": "wallet type",
                        "name": "wallet_type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/wallet.Wallet"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.Err"
                        }
                    }
                }
            }
        },
        "/api/v1/wallets": {
            "get": {
                "description": "Get all wallets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get all wallets",
                "parameters": [
                    {
                        "enum": [
                            "Savings",
                            "Credit Card",
                            "Crypto Wallet"
                        ],
                        "type": "string",
                        "description": "wallet type",
                        "name": "wallet_type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/wallet.Wallet"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.Err"
                        }
                    }
                }
            },
            "post": {
                "description": "Create wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Create wallet",
                "parameters": [
                    {
                        "description": "wallet request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/wallet.Request"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/wallet.Wallet"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.Err"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.Err"
                        }
                    }
                }
            }
        },
        "/api/v1/wallets/{id}": {
            "put": {
                "description": "Update wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Update wallet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wallet id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "wallet request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/wallet.Request"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.Err"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.Err"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.Err"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Delete wallet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wallet id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.Err"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.Err"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.Err"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errs.Err": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "wallet.Request": {
            "type": "object",
            "required": [
                "balance",
                "user_id",
                "user_name",
                "wallet_name",
                "wallet_type"
            ],
            "properties": {
                "balance": {
                    "type": "number",
                    "example": 100
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                },
                "user_name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "wallet_name": {
                    "type": "string",
                    "example": "John's Wallet"
                },
                "wallet_type": {
                    "type": "string",
                    "enum": [
                        "Savings",
                        "Credit Card",
                        "Crypto Wallet"
                    ],
                    "example": "Credit Card"
                }
            }
        },
        "wallet.Wallet": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number",
                    "example": 100
                },
                "created_at": {
                    "type": "string",
                    "example": "2024-03-25T14:19:00.729237Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                },
                "user_name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "wallet_name": {
                    "type": "string",
                    "example": "John's Wallet"
                },
                "wallet_type": {
                    "type": "string",
                    "example": "Create Card"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1323",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Wallet API",
	Description:      "Sophisticated Wallet API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
