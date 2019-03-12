// Code generated; DO NOT EDIT.
// file ./cmd/areacn/router_gen.go

package main

import (
	json "encoding/json"
	mux "github.com/gorilla/mux"
	githubComWzshimingAreacn "github.com/wzshiming/areacn"
	githubComWzshimingAreacnServiceAreacn "github.com/wzshiming/areacn/service/areacn"
	ui "github.com/wzshiming/openapi/ui"
	redoc "github.com/wzshiming/openapi/ui/redoc"
	swaggereditor "github.com/wzshiming/openapi/ui/swaggereditor"
	swaggerui "github.com/wzshiming/openapi/ui/swaggerui"
	http "net/http"
)

// Router is all routing for package
// generated do not edit.
func Router() http.Handler {
	router := mux.NewRouter()

	// AreacnService Define the method scope
	var _areacnService githubComWzshimingAreacnServiceAreacn.AreacnService
	RouteAreacnService(router, &_areacnService)

	router = RouteOpenAPI(router)

	return router
}

// RouteAreacnService is routing for AreacnService
func RouteAreacnService(router *mux.Router, _areacnService *githubComWzshimingAreacnServiceAreacn.AreacnService, fs ...mux.MiddlewareFunc) *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}
	subrouter := router.PathPrefix("/areacn").Subrouter()
	if len(fs) != 0 {
		subrouter.Use(fs...)
	}

	// Registered routing GET /areacn/{area_id}
	var __operationGetAreacnAreaID http.Handler
	__operationGetAreacnAreaID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetAreacnAreaID(_areacnService, w, r)
	})
	subrouter.Methods("GET").Path("/{area_id}").Handler(__operationGetAreacnAreaID)

	return router
}

// _requestPathAreaID Parsing the path for of area_id
func _requestPathAreaID(w http.ResponseWriter, r *http.Request) (_areaID string, err error) {

	var _raw_areaID = mux.Vars(r)["area_id"]
	_areaID = string(_raw_areaID)

	return
}

// _operationGetAreacnAreaID Is the route of Get
func _operationGetAreacnAreaID(s *githubComWzshimingAreacnServiceAreacn.AreacnService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _areaID string
	var _areas []*githubComWzshimingAreacn.Area

	// Parsing area_id.
	_areaID, err = _requestPathAreaID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/areacn/service/areacn AreacnService.Get.
	_areas, err = s.Get(_areaID)

	// Response code 200 OK for areas.
	if _areas != nil {
		var __areas []byte
		__areas, err = json.Marshal(_areas)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__areas)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __areas []byte
	__areas, err = json.Marshal(_areas)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__areas)

	return
}

var OpenAPI4YAML = []byte(`openapi: 3.0.1
info:
  title: OpenAPI Demo
  description: Automatically generated
  contact:
    name: wzshiming
    url: https://github.com/wzshiming/gen
  version: 0.0.1
servers:
- url: /
- url: '{scheme}{host}{port}{path}'
  variables:
    host:
      enum:
      - localhost
      default: localhost
    path:
      enum:
      - /
      default: /
    port:
      enum:
      - ""
      default: ""
    scheme:
      enum:
      - http://
      - https://
      default: http://
paths:
  /areacn/{area_id}:
    get:
      tags:
      - AreacnService
      summary: 'Get #route:"GET /{area_id}"# 获取行政区划分信息 总共5级 获取第一级省份传0'
      description: 'Get #route:"GET /{area_id}"# 获取行政区划分信息 总共5级 获取第一级省份传0'
      parameters:
      - $ref: '#/components/parameters/area_id'
      responses:
        "200":
          description: Response code is 200
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Area'
        "400":
          description: Response code is 400
          content:
            text/plain:
              schema:
                type: string
                format: error
components:
  schemas:
    Area:
      type: object
      properties:
        area_id:
          type: string
        name:
          type: string
  responses:
    areas:
      description: Response code is 200
      content:
        application/json:
          schema:
            type: array
            items:
              $ref: '#/components/schemas/Area'
    err:
      description: Response code is 400
      content:
        text/plain:
          schema:
            type: string
            format: error
  parameters:
    area_id:
      name: area_id
      in: path
      description: '#name:"area_id"#'
      required: true
      schema:
        type: string
tags:
- name: AreacnService
  description: 'AreacnService #path:"/areacn/"#'
`)
var OpenAPI4JSON = []byte(`{"openapi":"3.0.1","info":{"title":"OpenAPI Demo","description":"Automatically generated","contact":{"name":"wzshiming","url":"https://github.com/wzshiming/gen"},"version":"0.0.1"},"servers":[{"url":"/"},{"url":"{scheme}{host}{port}{path}","variables":{"host":{"enum":["localhost"],"default":"localhost"},"path":{"enum":["/"],"default":"/"},"port":{"enum":[""],"default":""},"scheme":{"enum":["http://","https://"],"default":"http://"}}}],"paths":{"/areacn/{area_id}":{"get":{"tags":["AreacnService"],"summary":"Get #route:\"GET /{area_id}\"# 获取行政区划分信息 总共5级 获取第一级省份传0","description":"Get #route:\"GET /{area_id}\"# 获取行政区划分信息 总共5级 获取第一级省份传0","parameters":[{"$ref":"#/components/parameters/area_id"}],"responses":{"200":{"description":"Response code is 200","content":{"application/json":{"schema":{"type":"array","items":{"$ref":"#/components/schemas/Area"}}}}},"400":{"description":"Response code is 400","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}}}}}},"components":{"schemas":{"Area":{"type":"object","properties":{"area_id":{"type":"string"},"name":{"type":"string"}}}},"responses":{"areas":{"description":"Response code is 200","content":{"application/json":{"schema":{"type":"array","items":{"$ref":"#/components/schemas/Area"}}}}},"err":{"description":"Response code is 400","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}}},"parameters":{"area_id":{"name":"area_id","in":"path","description":"#name:\"area_id\"#","required":true,"schema":{"type":"string"}}}},"tags":[{"name":"AreacnService","description":"AreacnService #path:\"/areacn/\"#"}]}`)

// RouteOpenAPI
func RouteOpenAPI(router *mux.Router) *mux.Router {
	openapi := map[string][]byte{
		"openapi.json": OpenAPI4JSON,
		"openapi.yml":  OpenAPI4YAML,
		"openapi.yaml": OpenAPI4YAML,
	}
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger", ui.HandleWithFiles(openapi, swaggerui.Asset)))
	router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui", ui.HandleWithFiles(openapi, swaggerui.Asset)))
	router.PathPrefix("/swaggereditor/").Handler(http.StripPrefix("/swaggereditor", ui.HandleWithFiles(openapi, swaggereditor.Asset)))
	router.PathPrefix("/redoc/").Handler(http.StripPrefix("/redoc", ui.HandleWithFiles(openapi, redoc.Asset)))
	return router
}
