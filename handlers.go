package main

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"github.com/pinzolo/linefido2"
	"net/http"
)

const (
	sessionIdHeader = "Fido2-Session-Id"
)

var client linefido2.Client

func handleRegistration(c echo.Context) error {
	ctx := c.Request().Context()

	var req RegistrationRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	fidoReq := &linefido2.RegistrationOptionsRequest{
		Rp: &linefido2.PublicKeyCredentialRpEntity{
			Id:   RpId,
			Name: "linefido2-example",
		},
		User: &linefido2.PublicKeyCredentialUserEntity{
			Id:          convertUserId(req.Login),
			Name:        req.Login,
			DisplayName: req.Login,
		},
		AuthenticatorSelection: &linefido2.AuthenticatorSelectionCriteria{
			AuthenticatorAttachment: linefido2.AuthenticatorAttachmentPlatform,
			RequireResidentKey:      false,
			UserVerification:        linefido2.UserVerificationRequirementRequired,
		},
		Attestation: linefido2.AttestationConveyancePreferenceNone,
	}
	res, err := client.GetRegistrationOptions(ctx, fidoReq)
	if err != nil {
		return err
	}

	session := RegistrationSession{
		Id:    res.SessionId,
		Login: req.Login,
	}
	SaveSession(session)
	setSessionId(c, session.Id)
	return c.JSON(http.StatusOK, res)
}

func handleAttestation(c echo.Context) error {
	ctx := c.Request().Context()

	var req AttestationRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	sessionId := getSessionId(c)

	fidoReq := &linefido2.RegisterCredentialRequest{
		PublicKeyCredential: &linefido2.RegistrationPublicKeyCredential{
			Id:       req.Id,
			Type:     req.Type,
			Response: req.Response,
		},
		RpId:      RpId,
		SessionId: sessionId,
		Origin:    RpOrigin,
	}

	res, err := client.RegisterCredential(ctx, fidoReq)
	if err != nil {
		return err
	}

	session := GetSession(sessionId)
	saveUser(session.Login)
	RemoveSession(sessionId)

	return c.JSON(http.StatusOK, res)
}

func handleAuthentication(c echo.Context) error {
	ctx := c.Request().Context()
	var req AuthenticationRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	fidoReq := &linefido2.AuthenticationOptionsRequest{
		RpId:             RpId,
		UserId:           convertUserId(req.Login),
		UserVerification: linefido2.UserVerificationRequirementRequired,
	}

	res, err := client.GetAuthenticationOptions(ctx, fidoReq)
	if err != nil {
		return err
	}

	setSessionId(c, res.SessionId)
	return c.JSON(http.StatusOK, res)
}

func handleAssertion(c echo.Context) error {
	ctx := c.Request().Context()
	var req AssertionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	sessionId := getSessionId(c)

	regReq := &linefido2.VerifyCredentialRequest{
		PublicKeyCredential: &linefido2.AuthenticationPublicKeyCredential{
			Id:       req.Id,
			Type:     req.Type,
			Response: req.Response,
		},
		RpId:      RpId,
		SessionId: sessionId,
		Origin:    RpOrigin,
	}

	res, err := client.VerifyCredential(ctx, regReq)
	if err != nil {
		return err
	}

	// application code related finding user.

	return c.JSON(http.StatusOK, res)
}

func convertUserId(login string) string {
	b := sha256.Sum256([]byte(login))
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b[:])
}

func setSessionId(c echo.Context, sessionId string) {
	c.Response().Header().Set(sessionIdHeader, sessionId)
}

func getSessionId(c echo.Context) string {
	return c.Request().Header.Get(sessionIdHeader)
}

func saveUser(login string) {
	// Persist user to datastore.
}
