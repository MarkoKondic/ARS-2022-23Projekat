package main

import s "github.com/MihajloJankovic/Alati/Dao"

// swagger:parameters deleteConfig
type DeleteConfig struct {
	// Post ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters deleteConfigGroup
type DeleteConfigGroup struct {
	// Post ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters getConfigById
type GetConfigById struct {
	// Post ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters createConfigGroup
type CreateConfigGroup struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/RequestPost"
	//  required: true
	Body s.Config `json:"body"`
}

// swagger:parameters createConfig
type CreateConfig struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/RequestPost"
	//  required: true
	Body s.Config `json:"body"`
}

// swagger:parameters getConfigs
type GetConfigs struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/RequestPost"
	//  required: true

}

// swagger:parameters getConfigGroups
type GetConfigGroups struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/RequestPost"
	//  required: true

}

// swagger:parameters getConfigGroupById
type GetConfigGroupById struct {
	// Post ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters addConfigInConfigGroup
type AddConfigInConfigGroup struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/RequestPost"
	//  required: true
	Body s.Config `json:"body"`
}
