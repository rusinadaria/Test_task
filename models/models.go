package models

// Couplet provides data of a song.
//
// swagger:model Couplet
type Verse struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

// Song provides data of a song.
//
// swagger:model Song
type Song struct {
	Id string
	// required: true
	Song string `json:"song"`
	// required: true
	GroupName string `json:"group"`
	// required: true
	ReleaseDate string `json:"releaseDate"`
	// required: true
	Text string `json:"text"`
	// required: true
	Link string `json:"link"`
}

// ErrorResponse represents a structure for API errors.

// swagger:model ErrorResponse
type ErrorResponse struct {
	Message string `json:"message"`
}

// type SongDetail struct {
// 	ReleaseDate string `json:"releaseDate"`
// 	Text string `json:"text"`
// 	Link string `json:"link"`
// }