package models

import ()

type Car struct {
	id      int64 `db:"article_id"`
	created int64
	make    string
	model   string
	year    string
}
