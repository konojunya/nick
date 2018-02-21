package model

type Json struct {
	Dependencies    []string `json:"dependencies"`
	DevDependencies []string `json:"devDependencies"`
}
