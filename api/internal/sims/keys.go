package sims

import "github.com/jhunthrop/foreversixty/logs/engine/store"

// Keys are a sim's object keys in the bucket, laid out the way a
// report's are. There is only one: the request lives in the row,
// because the envelope carries no protobuf and a SimRequest is a few
// hundred bytes of JSON.
type Keys struct{ SimID string }

// Result is sims/<id>/result.json: the whole SimResult, including the
// summary the report components render.
func (k Keys) Result() string { return "sims/" + k.SimID + "/result.json" }

// ResultPut is how a stored result is written: it never changes once
// written, and the sim it belongs to is public.
var ResultPut = store.PutOptions{
	ContentType:  "application/json; charset=utf-8",
	CacheControl: "public, max-age=31536000, immutable",
}
