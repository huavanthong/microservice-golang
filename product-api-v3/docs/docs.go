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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/products": {
            "get": {
                "description": "List all existing products of store",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "List all existing products",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Page",
                        "name": "page",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Limit",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Currency",
                        "name": "currency",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    }
                }
            }
        },
        "/products/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "User create a product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Create a product",
                "parameters": [
                    {
                        "description": "New Product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.RequestCreateProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GetProductSuccess"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    }
                }
            }
        },
        "/products/category/{category}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get product by Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get product by Category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category",
                        "name": "category",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Currency",
                        "name": "currency",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GetProductSuccess"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    }
                }
            }
        },
        "/products/name/{name}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get product by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get product by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of Product",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Currency",
                        "name": "currency",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GetProductSuccess"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Get a product by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get a product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Currency",
                        "name": "currency",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GetProductSuccess"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
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
                "description": "User delete the product by product ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Delete a post by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "User update product info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Update a product",
                "parameters": [
                    {
                        "description": "Update post",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Product": {
            "type": "object",
            "required": [
                "category",
                "description",
                "imageFile",
                "name",
                "price",
                "sku",
                "summary"
            ],
            "properties": {
                "category": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "imageFile": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string",
                    "minLength": 0
                },
                "sku": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "payload.GetProductSuccess": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 201
                },
                "data": {
                    "$ref": "#/definitions/models.Product"
                },
                "message": {
                    "type": "string",
                    "example": "Get product success"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "payload.RequestCreateProduct": {
            "type": "object",
            "required": [
                "category",
                "description",
                "imageFile",
                "name",
                "price",
                "summary"
            ],
            "properties": {
                "category": {
                    "type": "string",
                    "example": "Phone"
                },
                "description": {
                    "type": "string",
                    "example": "Iphone 14 Pro Gold 256GB"
                },
                "imageFile": {
                    "type": "string",
                    "example": "default.png"
                },
                "name": {
                    "type": "string",
                    "example": "Iphone 14 Pro"
                },
                "price": {
                    "type": "string",
                    "minLength": 0,
                    "example": "1400$"
                },
                "summary": {
                    "type": "string",
                    "example": "Iphone 14 Pro Gold"
                }
            }
        },
        "payload.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "Error message"
                },
                "status": {
                    "type": "string",
                    "example": "failed"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9090",
	BasePath:         "/api/v3",
	Schemes:          []string{},
	Title:            "UserManagement Service API Document",
	Description:      "List APIs of UserManagement Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
