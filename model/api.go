package model

type ApiInfo struct {
	Info     Info       `json:"info"`
	Item     []Item     `json:"item"`
	Event    []Event    `json:"event"`
	Variable []Variable `json:"variable"`
}

type Info struct {
	PostmanID string `json:"_postman_id"`
	Name      string `json:"name"`
	Schema    string `json:"schema"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Urlencoded struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Body struct {
	Mode       string       `json:"mode"`
	Urlencoded []Urlencoded `json:"urlencoded"`
}

type Query struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Request struct {
	Method string      `json:"method"`
	Header []Header    `json:"header"`
	Body   Body        `json:"body"`
	URL    interface{} `json:"url"`
}

type Item struct {
	Name     string        `json:"name"`
	Request  Request       `json:"request"`
	Response []interface{} `json:"response"`
}

type Script struct {
	Type string   `json:"type"`
	Exec []string `json:"exec"`
}

type Event struct {
	Listen string `json:"listen"`
	Script Script `json:"script"`
}

type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
