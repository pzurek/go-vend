package vend

type Migration struct {
	Version *string `json:"version"`
	Table   *string `json:"table"`
	Query   *string `json:"query"`
}
