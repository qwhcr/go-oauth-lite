package oauth2

import (
	"errors"
	"net/http"

	configUtil "go-oauth-lite/util/config"
	responseUtil "go-oauth-lite/util/response"

	//firebaseWrapper "go-oauth-lite/util/firebaseWrapper"
	oauth2client "go-oauth-lite/client"

	routing "github.com/qiangxue/fasthttp-routing"
)

var parameterClientID = "client_id"
var parameterResponseType = "response_type"
var parameterScope = "scope"
var parameterState = "state"
var parameterRedirectURI = "redirect_uri"

type authRequest struct {
	ClientID     string
	ResponseType string
	Scope        string
	State        string
	RedirectURI  string
}

func AuthorizationHandler(ctx *routing.Context) error {
	authReq, err := extractParam(ctx)
	if err != nil {
		return responseUtil.RespondError(ctx,
			err,
			"malformed request",
			http.StatusBadRequest)
	}

	isValidClient, err := oauth2client.IsValidClientAuthRequest(authReq.ClientID, authReq.Scope)
	if err != nil {
		return responseUtil.RespondError(ctx, err, "internal error", http.StatusInternalServerError)
	}
	if !isValidClient {
		return responseUtil.RespondError(ctx, nil, "invalid client", http.StatusForbidden)
	}

	//checked clientID is valid and scope is valid for the client

	ctx.SendFile(configUtil.GetConfig().LoginAssetsURL)
	return nil
}

// extractParam extracts necessary paramters used for oauth2 flow from the request
func extractParam(ctx *routing.Context) (*authRequest, error) {
	var authReq authRequest
	if clientID := string(ctx.QueryArgs().Peek(parameterClientID)); clientID != "" {
		authReq.ClientID = clientID
	} else {
		return nil, errors.New("expected client_id get request paramter to be not empty")
	}

	if responseType := string(ctx.QueryArgs().Peek(parameterResponseType)); responseType != "" {
		authReq.ResponseType = responseType
	} else {
		return nil, errors.New("expected client_id get request paramter to be not empty")
	}

	if scope := string(ctx.QueryArgs().Peek(parameterScope)); scope != "" {
		authReq.Scope = scope
	} else {
		return nil, errors.New("expected client_id get request paramter to be not empty")
	}

	// state is not required
	if state := string(ctx.QueryArgs().Peek(parameterState)); state != "" {
		authReq.State = state
	}

	if redirectURI := string(ctx.QueryArgs().Peek(parameterRedirectURI)); redirectURI != "" {
		authReq.RedirectURI = redirectURI
	} else {
		return nil, errors.New("expected client_id get request paramter to be not empty")
	}
	return &authReq, nil
}
