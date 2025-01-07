package server

import (
	"awesomeProject/library-app/global"
	"awesomeProject/library-app/managers"
	"awesomeProject/library-app/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func Static() http.Handler {
	return http.FileServer(http.Dir("./public"))
}

func (server *Server) Public(next http.HandlerFunc) http.HandlerFunc {
	return server.Protected(next, global.PUBLIC, global.PUBLIC)
}

func (server *Server) Protected(next http.HandlerFunc, resource string, role string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.EqualFold(resource, global.PUBLIC) {
			// Parse token
			authHeader := r.Header.Get("Authorization")
			if len(authHeader) < 7 {
				global.HttpError(
					w,
					"[Protected] Missing or incomplete Authorization header",
					"Unauthorized request: Authorization header is missing or incomplete",
					http.StatusUnauthorized,
					fmt.Errorf("unauthorized, missing bearer authorization header"),
				)
				return
			}

			authType := strings.ToLower(authHeader[:6])
			if authType != "bearer" {
				global.HttpError(
					w,
					"[Protected] Invalid Authorization type",
					"Unauthorized request: Authorization header must use Bearer token",
					http.StatusUnauthorized,
					fmt.Errorf("unauthorized, invalid bearer authorization header"),
				)
				return
			}

			// Verify token is valid
			tokenString := strings.TrimSpace(authHeader[7:])
			err := server.AuthClient.RetrospectToken(r.Context(), tokenString)
			if err != nil {
				global.HttpError(
					w,
					"[Protected] Invalid or expired token",
					"Unauthorized request: Token is invalid or expired",
					http.StatusUnauthorized,
					err,
				)
				return
			}

			// Create user if not exists
			userFromInfo, err := server.AuthClient.GetUserFromToken(r.Context(), tokenString)
			if err != nil {
				global.HttpError(
					w,
					"[Protected] Failed to extract user from token",
					"Unauthorized request: Could not extract user information from token",
					http.StatusUnauthorized,
					err,
				)
				return
			}

			loadedUser, _ := server.DBLoadUser(userFromInfo.ID.String()) // Ignore error as user might not exist
			if loadedUser == nil {
				log.Info(userFromInfo)
				err := server.DBSaveUser(*userFromInfo)
				if err != nil {
					global.HttpError(
						w,
						"[Protected] Failed to save new user",
						"Unauthorized request: Could not save user to the database",
						http.StatusUnauthorized,
						err,
					)
					return
				}
			}

			loadedUser, err = server.DBLoadUser(userFromInfo.ID.String())
			if err != nil {
				global.HttpError(
					w,
					"[Protected] Failed to load user from database",
					"Unauthorized request: Could not retrieve user details from the database",
					http.StatusUnauthorized,
					err,
				)
				return
			}

			// Create new context with current user
			newCtx := context.WithValue(r.Context(), global.CURRENT_USER_ID, loadedUser.ID.String())

			// Get roles from token
			roles, err := server.AuthClient.GetRolesFromToken(r.Context(), tokenString)
			if err != nil {
				global.HttpError(
					w,
					"[Protected] Failed to retrieve roles from token",
					"Unauthorized request: Could not retrieve roles from the token",
					http.StatusUnauthorized,
					err,
				)
				return
			}

			// Check for global scope for this resource and if exists add it to the enriched context
			if havePermission(resource, global.GLOBAL, roles) {
				newCtx = context.WithValue(newCtx, global.GLOBAL_SCOPE, global.GLOBAL)
			}

			// Replace request context
			rWithUpdatedContext := r.WithContext(newCtx)

			// Check permissions
			if havePermission(resource, role, roles) {
				next(w, rWithUpdatedContext)
			} else {
				global.HttpError(
					w,
					fmt.Sprintf("[Protected] Insufficient permissions for resource '%s' and role '%s'", resource, role),
					fmt.Sprintf("Unauthorized request: You lack the required permissions for resource '%s' with role '%s'", resource, role),
					http.StatusUnauthorized,
					fmt.Errorf("unauthorized, required scope (%s.%s)", resource, role),
				)
				return
			}
		} else {
			// NO authorization is required
			next(w, r)
		}
	}
}

// TODO: move manager out
func (server *Server) DBLoadUser(userID string) (*models.User, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	manager := managers.NewUserManager(server.DB)
	user, err := manager.Get(uid)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// TODO: move manager out
func (server *Server) DBSaveUser(user models.User) error {
	manager := managers.NewUserManager(server.DB)
	_, err := manager.Create(user)

	if err != nil {
		return err
	}
	return nil
}

func havePermission(resource, role string, roles []string) bool {
	for _, currentRole := range roles {
		resourceRole := fmt.Sprintf("%s.%s", resource, role)
		if strings.EqualFold(currentRole, resourceRole) {
			return true
		}
	}
	return false
}
