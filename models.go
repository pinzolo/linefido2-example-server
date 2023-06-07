package main

import "github.com/pinzolo/linefido2"

type ErrorResponse struct {
	Message string `json:"message"`
}

type RegistrationRequest struct {
	Login string `json:"login"`
}

type AttestationRequest struct {
	Id         string                                           `json:"id"`
	Type       linefido2.PublicKeyCredentialType                `json:"type"`
	RawId      string                                           `json:"rawId"`
	Response   *linefido2.AuthenticatorAttestationResponse      `json:"response"`
	Extensions *linefido2.AuthenticationExtensionsClientOutputs `json:"extensions"`
}

type AuthenticationRequest struct {
	Login string `json:"login"`
}

type AssertionRequest struct {
	Id         string                                           `json:"id"`
	Type       linefido2.PublicKeyCredentialType                `json:"type"`
	RawId      string                                           `json:"rawId"`
	Response   *linefido2.AuthenticatorAssertionResponse        `json:"response"`
	Extensions *linefido2.AuthenticationExtensionsClientOutputs `json:"extensions"`
}

type User struct {
	Id    int64
	Login string
}
