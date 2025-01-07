package security

import (
	"awesomeProject/library-app/models"
	"context"
	"errors"
	"github.com/Nerzal/gocloak/v13"
	"github.com/google/uuid"

	"github.com/Nerzal/gocloak/v13/pkg/jwx"
)

type AuthClient struct {
	Client       *gocloak.GoCloak
	URL          string
	Realm        string
	ClientID     string
	ClientSecret string
}

func (authClient *AuthClient) Initialize(authURL, authRealm, authClientID, authClientSecret string) {
	authClient.Client = gocloak.NewClient(authURL)
	authClient.URL = authURL
	authClient.Realm = authRealm
	authClient.ClientID = authClientID
	authClient.ClientSecret = authClientSecret
}

func (authClient *AuthClient) RetrospectToken(ctx context.Context, accessToken string) error {
	rptResult, err := authClient.Client.RetrospectToken(ctx, accessToken, authClient.ClientID, authClient.ClientSecret, authClient.Realm)
	if err != nil {
		return err
	}
	if !*rptResult.Active {
		return errors.New("token is not active")
	}

	return nil
}

func (authClient *AuthClient) GetRolesFromToken(ctx context.Context, accessToken string) ([]string, error) {
	jwxClaims := &jwx.Claims{}
	_, err := authClient.Client.DecodeAccessTokenCustomClaims(ctx, accessToken, authClient.Realm, jwxClaims)
	if err != nil {
		result := make([]string, 0)
		return result, err
	}
	return jwxClaims.RealmAccess.Roles, nil
}

func (authClient *AuthClient) GetUserFromToken(ctx context.Context, accessToken string) (*models.User, error) {
	jwxClaims := &jwx.Claims{}
	_, err := authClient.Client.DecodeAccessTokenCustomClaims(ctx, accessToken, authClient.Realm, jwxClaims)
	if err != nil {
		return nil, err
	}

	uid, err := uuid.Parse(jwxClaims.Subject)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		BaseModel: models.BaseModel{
			ID: uid,
		},
		PreferredUsername: jwxClaims.PreferredUsername,
		GivenName:         jwxClaims.GivenName,
		FamilyName:        jwxClaims.FamilyName,
		Email:             jwxClaims.Email,
	}

	return user, nil
}
