package dto

type Response struct {
	Uno string `json:"un"`
}

type Request struct {
	Three string `param:"three"`
	Fang  string `query:"fang"`
}
