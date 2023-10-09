package entity

// SearchRequest ...
type SearchRequest struct {
	SearchWord string `json:"search_word"`
	Pagination int    `json:"pagination"`
}

// SearchResponse ...
type SearchResponse struct {
	Search        []SearchDetail `json:"Search"`
	TotalResponse string         `json:"totalResults"`
	Response      string         `json:"Response"`
	Error         string         `json:"Error"`
}

// SearchDetail ...
type SearchDetail struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

// GetRequest ...
type GetRequest struct {
	ID string `json:"id"`
}

// GetResponse ...
type GetResponse struct {
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Rated      string   `json:"Rated"`
	Released   string   `json:"Released"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	Writer     string   `json:"Writer"`
	Actors     string   `json:"Actors"`
	Plot       string   `json:"Plot"`
	Language   string   `json:"Language"`
	Country    string   `json:"Country"`
	Awards     string   `json:"Awards"`
	Poster     string   `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	Metascore  string   `json:"Metascore"`
	ImdbRating string   `json:"imdbRating"`
	ImdbVotes  string   `json:"imdbVotes"`
	ImdbID     string   `json:"imdbID"`
	Type       string   `json:"Type"`
	DVD        string   `json:"DVD"`
	BoxOffice  string   `json:"BoxOffice"`
	Production string   `json:"Production"`
	Website    string   `json:"Website"`
	Response   string   `json:"Response"`
	Error      string   `json:"Error"`
}

// Rating ...
type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
