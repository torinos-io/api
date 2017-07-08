package service

import (
	"github.com/matcornic/hermes"
)

// Request is a request object for sending a mail
type Request struct {
}

func (s *service) Perform(req *Request) error {
	h := s.NewHermes()
	email := s.BuildEmail()

	body, err := h.GenerateHTML(email)
	if err != nil {
		return err
	}

	println(body)
	return nil
}

func (s *service) NewHermes() hermes.Hermes {
	return hermes.Hermes{
		Product: hermes.Product{
			Name:      "Torinos",
			Link:      "https://torinos.io",
			Logo:      "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
			Copyright: "Copyright © 2017 Torinos. All rights reserved.",
		},
	}
}

func (s *service) BuildEmail() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Confirm your account",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
}
