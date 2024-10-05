package authHelper

import (
	"context"
	"middlend-template/pkg/domain/innerDomain/response"
	"middlend-template/pkg/useCases/Helpers/responseHelper"
	"net/http"
)

type UserKey string

// a middleware function to check that a user is authenticated
// and has the correct role
func WithClaims(expectedRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Extract token from Authorization header
			token := r.Header.Get("Authorization")
			if token == "" {
				status := response.Status{
					Text: "Unauthorized: Missing token",
					Code: http.StatusUnauthorized,
				}
				responseHelper.WriteResponse(w, status, nil)
				return
			}

			// 2. Verify the token
			t, err := VerifyIdToken(token)
			if err != nil {
				status := response.Status{
					Text: "Unauthorized: Invalid token",
					Code: http.StatusUnauthorized,
				}
				responseHelper.WriteResponse(w, status, nil)
				return
			}

			// 3. Check user role against expected role
			user, err := GetUserWithClaims(t.UID, expectedRole)
			if err != nil {
				// Handle error: user not found, or invalid role
				status := response.Status{
					Text: "Forbidden: Invalid role or user not found",
					Code: http.StatusUnauthorized,
				}
				responseHelper.WriteResponse(w, status, nil)
				return
			}

			// 4. User is authorized, proceed to next handler
			ctx := context.WithValue(r.Context(), UserKey("user"), user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func WithUid() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Extract token from Authorization header
			token := r.Header.Get("Authorization")
			if token == "" {
				status := response.Status{
					Text: "Unauthorized: Missing token",
					Code: http.StatusUnauthorized,
				}
				responseHelper.WriteResponse(w, status, nil)
				return
			}

			// 2. Verify the token
			t, err := VerifyIdToken(token)
			if err != nil {
				status := response.Status{
					Text: "Unauthorized: Invalid token",
					Code: http.StatusUnauthorized,
				}
				responseHelper.WriteResponse(w, status, nil)
				return
			}

			// 3. User is authorized, get user by uid
			user, err := GetUserByUid(t.UID)
			if err != nil {
				// Handle error: user not found
				status := response.Status{
					Text: "Forbidden: Invalid role or user not found",
					Code: http.StatusUnauthorized,
				}
				responseHelper.WriteResponse(w, status, nil)
				return
			}

			ctx := context.WithValue(r.Context(), UserKey("user"), user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
