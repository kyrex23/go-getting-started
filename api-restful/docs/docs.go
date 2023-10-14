// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Kyrex",
            "url": "github.com/kyrex23"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/heroes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Heroes"
                ],
                "summary": "Get the list of heroes",
                "responses": {
                    "200": {
                        "description": "List of heroes retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handler.heroResponse"
                            }
                        }
                    }
                }
            }
        },
        "/heroes/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Heroes"
                ],
                "summary": "Get a hero by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Hero ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Hero found successfully",
                        "schema": {
                            "$ref": "#/definitions/handler.heroResponse"
                        }
                    },
                    "404": {
                        "description": "Hero not found"
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.heroResponse": {
            "type": "object",
            "properties": {
                "actual_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Marvel API",
	Description:      "API about the Marvel Cinematic Universe (MCU)",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
