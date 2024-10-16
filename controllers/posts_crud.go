package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/goravel-community/goravel-admin/fields/field_text"
	"github.com/goravel/framework/contracts/http"
)

type CrudApi struct{}

func (a *CrudApi) AddField(name string, def any) {
	b, _ := json.Marshal(def)
	fmt.Println(string(b))
}

type PostsCrud struct {
	crud CrudApi
}

func NewPostsCrud() *PostsCrud {
	return &PostsCrud{
		crud: CrudApi{},
	}
}

func (r *PostsCrud) Setup(ctx http.Context) {
	// method 1, functional options
	r.crud.AddField("title", field_text.Make(
		field_text.WithPlaceholder("Title"),
		field_text.WithReadOnly(false),
		func(def *field_text.Type) {
			if true {
				def.ReadOnly = true
			}
		},
	))

	// method 2, struct options
	r.crud.AddField("body", field_text.Type{
		ReadOnly:    false,
		Placeholder: "Body",
	})

	// method 3, expressive struct options
	description := field_text.Type{
		ReadOnly:    false,
		Placeholder: "Description",
	}
	if true {
		description.ReadOnly = true
	}
	r.crud.AddField("description", &description)

	// method 4, update existing
	r.crud.AddField("title", field_text.Make(func(def *field_text.Type) {
		if true {
			def.ReadOnly = true
		}
	}))
}

type Type struct {
}
