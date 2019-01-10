// This package contains files for the Relay API
package relay

/*
"models.go" defines API models in this file.
*/

type Dispatch struct {
	DevID string `json:"devid"`
}

/*
"Relay" is a JSON structure that specifies information to complete reads and writes to other blockchains
*/
type Relay struct {
	Blockchain string `json:"blockchain"`
	NetworkID  string `json:"netid"`
	Version	   string `json:"version"`
	Data       string `json:"data"`
}

// NOTE: This is for the centralized dispatcher of Pocket core mvp, may be removed for production
type Report struct {
	GID string 		`json:"gid"`
	Message string 	`json:"message"`
}
