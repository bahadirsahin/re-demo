package structs

type Env struct {
	ApiToken string
}

type PackRequest struct {
	Items     int   `json:"items"`
	PackSizes []int `json:"pack_sizes"`
}

type PackResponse struct {
	NumOfPacks int      `json:"num_of_packs"`
	Packs      []string `json:"packs"`
}
