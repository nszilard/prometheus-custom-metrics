// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/alive": {
            "get": {
                "description": "Responds to the Kubernetes alive requests",
                "produces": [
                    "text/text"
                ],
                "tags": [
                    "Common"
                ],
                "summary": "Kubernetes Alive probe",
                "operationId": "alive",
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
        "/metrics": {
            "get": {
                "description": "Metrics is an http.Handler instance to expose Prometheus metrics via HTTP.",
                "tags": [
                    "Common"
                ],
                "summary": "Prometheus Metrics",
                "operationId": "metrics",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/ready": {
            "get": {
                "description": "Responds to the Kubernetes ready requests",
                "produces": [
                    "text/text"
                ],
                "tags": [
                    "Common"
                ],
                "summary": "Kubernetes Ready probe",
                "operationId": "ready",
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
        "/v1/delay": {
            "get": {
                "produces": [
                    "text/text"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Responds with a 200 HTTP status code but with a random delay",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/error": {
            "get": {
                "tags": [
                    "v1"
                ],
                "summary": "Responds with a 500 HTTP status code",
                "responses": {
                    "500": {
                        "description": "Oh no, something went wrong!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/ok": {
            "get": {
                "produces": [
                    "text/text"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Responds with a 200 HTTP status code",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Prometheus Custom Metrics",
	Description:      "API documentation for the 'Prometheus Custom Metrics' application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}