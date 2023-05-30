package main

// swagger:response ResponsePost
type ResponsePost struct {
	// Id of the config
	// in: string
	Id      string            `json:"id"`

	// Version of the config
	// in: string
	Version string            `json:"version"`

}

// swagger:response ErrorResponse
type ErrorResponse struct {
	// Error status code
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
 }

 // swagger:response NoContentResponse
type NoContentResponse struct {}
 
