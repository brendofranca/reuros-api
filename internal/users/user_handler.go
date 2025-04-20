package users

import (
	"encoding/json"
	"net/http"
	"reuros-api/pkg"
)

// CreateUserHandler godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body users.UserRequest true "User details"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func CreateUserHandler(repo *UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			pkg.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{
				"error": "Invalid request payload",
			})
			return
		}

		if err := repo.CreateUser(&user); err != nil {
			pkg.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{
				"error": "Failed to create user",
			})
			return
		}

		pkg.WriteJSONResponse(w, http.StatusCreated, map[string]any{
			"message": "User created successfully",
			"user_id": user.ID,
		})
	}
}
