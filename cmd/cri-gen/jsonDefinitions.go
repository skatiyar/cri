package main

type Version struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
}

type Item struct {
	Ref  string `json:"$ref"`
	Type string `json:"type"`
}

type Parameter struct {
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Ref          string   `json:"$ref"`
	Enum         []string `json:"enum"`
	Items        Item     `json:"items"`
	Optional     bool     `json:"optional"`
	Description  string   `json:"description"`
	Experimental bool     `json:"experimental"`
}

type Type struct {
	ID           string      `json:"id"`
	Type         string      `json:"type"`
	Enum         []string    `json:"enum"`
	Properties   []Parameter `json:"properties"`
	Items        Item        `json:"items"`
	Experimental bool        `json:"experimental"`
	Description  string      `json:"description"`
}

type Command struct {
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Parameters   []Parameter `json:"parameters"`
	Returns      []Parameter `json:"returns"`
	Experimental bool        `json:"experimental"`
}

type Event struct {
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Parameters   []Parameter `json:"parameters"`
	Experimental bool        `json:"experimental"`
}

type Domain struct {
	Domain       string    `json:"domain"`
	Description  string    `json:"description"`
	Dependencies []string  `json:"dependencies"`
	Experimental bool      `json:"experimental"`
	Events       []Event   `json:"events"`
	Types        []Type    `json:"types"`
	Commands     []Command `json:"commands"`
}

type BrowserProtocol struct {
	Version Version  `json:"version"`
	Domains []Domain `json:"domains"`
}

type JSProtocol struct {
	Version Version  `json:"version"`
	Domains []Domain `json:"domains"`
}
