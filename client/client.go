package client

import (
	"context"
    "errors"

	"github.com/google/uuid"
	"google.golang.org/api/iterator"

	firebaseWrapper "go-oauth-lite/util/firebaseWrapper"
)

type Client struct {
	ClientID    uuid.UUID `firestore:"client_id"`
	ValidScopes []string  `firestore:"valid_scope"`
}

func IsValidClientAuthRequest(clientIDParam string, scopeParam string) (bool, error) {
	_, err := uuid.Parse(clientIDParam)

	if err != nil {
		return false, err
	}

	iter := firebaseWrapper.GetFirestoreClient().
		Collection("clients").
		Where("client_id", "==", clientIDParam).
		Limit(1).
		Documents(context.Background())

	doc, err := iter.Next()
	if err == iterator.Done {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	var client Client
	err = doc.DataTo(&client)
    validScopes, ok := doc.Data()["valid_scope"].([]interface{})
    if !ok {
        return false, errors.New("ValidScopes type mismatch")
    }
    for _, validScope := range validScopes {
            if validScope.(string) == scopeParam {
                return true, nil
            }
        }
    return false, nil

}
