package authHelper

import (
	"context"
	"os"

	"errors"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var client *auth.Client

// Initialize a firebase client
func InitFirebaseClient() error {
	credentials := os.Getenv("PROJECT_FIREBASE_AUTH")
	opts := option.WithCredentialsJSON([]byte(credentials))

	app, err := firebase.NewApp(context.Background(), nil, opts)
	if err != nil {
		return err
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		return err
	}

	client = auth
	return nil
}

func VerifyIdToken(token string) (*auth.Token, error) {
	if client == nil {
		return nil, errors.New("firebase client not initialized")
	}

	t, err := client.VerifyIDToken(context.Background(), token)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func GetUserByUid(uid string) (*auth.UserRecord, error) {
	if client == nil {
		return nil, errors.New("firebase client not initialized")
	}

	user, err := client.GetUser(context.Background(), uid)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(email string) (*auth.UserRecord, error) {
	if client == nil {
		return nil, errors.New("firebase client not initialized")
	}

	user, err := client.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserWithClaims(uid string, claims ...string) (*auth.UserRecord, error) {
	if client == nil {
		return nil, errors.New("firebase client not initialized")
	}

	user, err := client.GetUser(context.Background(), uid)
	if err != nil {
		return nil, err
	}

	for _, claim := range claims {
		if user.CustomClaims[claim] == nil {
			return nil, errors.New("user does not have the required claims")
		}
	}

	return user, nil
}

func GetUserFromContext(ctx context.Context) *auth.UserRecord {
	user := ctx.Value(UserKey("user"))
	if user == nil {
		return nil
	}

	return user.(*auth.UserRecord)
}

func CreateUserWithPassword(email, password string) (*auth.UserRecord, error) {
	if client == nil {
		return nil, errors.New("firebase client not initialized")
	}

	params := (&auth.UserToCreate{}).
		Email(email)

	if password != "" {
		params = params.Password(password)
	}

	user, err := client.CreateUser(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func AddCustomClaims(uid string, claims ...string) error {
	if client == nil {
		return errors.New("firebase client not initialized")
	}

	user, err := client.GetUser(context.Background(), uid)
	if err != nil {
		return err
	}

	if user.CustomClaims == nil {
		user.CustomClaims = make(map[string]interface{})
	}

	for _, claim := range claims {
		user.CustomClaims[claim] = true
	}

	err = client.SetCustomUserClaims(context.Background(), uid, user.CustomClaims)

	if err != nil {
		return err
	}

	return nil
}

func GetPasswordResetLink(email string) (string, error) {
	if client == nil {
		return "", errors.New("firebase client not initialized")
	}

	link, err := client.PasswordResetLink(context.Background(), email)
	if err != nil {
		return "", err
	}

	return link, nil
}

func VerifyEmail(uid string) error {
	if client == nil {
		return errors.New("firebase client not initialized")
	}

	_, err := client.UpdateUser(context.Background(), uid, (&auth.UserToUpdate{}).EmailVerified(true))
	if err != nil {
		return err
	}

	return nil
}
