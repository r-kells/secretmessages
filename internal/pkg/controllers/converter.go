package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/r-kells/secretmessages/internal/pkg/views"
	"github.com/r-kells/secretmessages/pkg/crypto"
	"github.com/r-kells/secretmessages/pkg/crypto/aes"
)

type Index struct {
	IndexPage *views.View
}
type ConverterForm struct {
	crypto.Message
	crypto.Secret
	Direction string
}

//NewIndex is called in main() and parses the templates on startup.
func NewIndex() *Index {
	return &Index{
		IndexPage: views.NewView("static/index"),
	}
}

// Landing handles the index page of the webserver.
// It pre-populates some of the apps form fields to demonstrate for the user.
func (i *Index) Landing(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	form := &ConverterForm{
		Message:   crypto.Message("The recipe for Heinz Ketchup is..."),
		Secret:    crypto.Secret(crypto.RandStringBytes(16)),
		Direction: "decrypt"}

	i.IndexPage.Render(w, form)
}

// Convert will POST to "/" and render the same page with the encrypted or decrypted message
func (i *Index) Convert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var form ConverterForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	switch form.Direction {
	case "encrypt":
		encryptedMsg, err := aes.Encrypt(form.Secret, form.Message)
		if err != nil {
			panic(err)
		}
		form.Message = crypto.Message(encryptedMsg)

	case "decrypt":
		decryptedMsg, err := aes.Decrypt(form.Secret, form.Message)
		if err != nil {
			panic(err)
		}
		form.Message = crypto.Message(decryptedMsg)
	}

	i.IndexPage.Render(w, form)
}
