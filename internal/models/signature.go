package models

type Signatures []Signature

type Signature struct {
	V string `json:"v,omitempty" bson:"v"`
	R string `json:"r,omitempty" bson:"r"`
	S string `json:"s,omitempty" bson:"s"`
}
