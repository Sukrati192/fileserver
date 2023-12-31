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
        "/file/{fileID}": {
            "get": {
                "description": "Checks if a file exists in IPFS storage with given id, then stream the file",
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "public"
                ],
                "summary": "API to get the uploaded file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the file",
                        "name": "fileID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/contracts.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "description": "Uploads the file in IPFS storage and returns id, size and timestamp in response",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "public"
                ],
                "summary": "API to upload a file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to be uploaded",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/contracts.File"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/contracts.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "contracts.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "contracts.File": {
            "type": "object",
            "properties": {
                "fileID": {
                    "type": "string"
                },
                "size": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "FileServer",
	Description:      "A Service that facilitates file upload and download using IPFS",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
