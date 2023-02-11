package http

import (
	. "SocialNetHTTPService/internal/service"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type UserInformationResponse struct {
	Name    string          `json:"name"`
	Age     int             `json:"age"`
	Friends FriendsResponse `json:"friends"`
}

type UserUpdateRequest struct {
	Name string `json:"new name,omitempty"`
	Age  string `json:"new age,omitempty"`
}

func (h *Handler) UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			err    error
			userID int
		)

		if userIDStr := chi.URLParam(r, "user_id"); userIDStr != "" {
			userID, err = strconv.Atoi(userIDStr)
		} else {
			render.Render(w, r, ErrNotFound)
			return
		}
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) initRootRoutes() chi.Router {
	r := chi.NewRouter()
	r.Route("/{user_id}", func(r chi.Router) {
		r.Use(h.UserCtx)
		r.Get("/", h.GetUserInformation)
		r.Put("/", h.UserUpdate)
	})
	return r
}

// @Summary GetUserInformation
// @Description Getting information about the user
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} UserInformationResponse
// @Failure 400 {object} service.ErrResponse
// @Router /{user_id} [get]
func (h *Handler) GetUserInformation(w http.ResponseWriter, r *http.Request) {
	var (
		err         error
		userID      int
		userName    string
		userAge     int
		userFriends []int
		userInf     UserInformationResponse
	)
	userID = r.Context().Value("user_id").(int)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if userName, err = h.usersService.GetNameUser(r.Context(), userID); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	userInf.Name = userName
	if userAge, err = h.usersService.GetAgeUser(r.Context(), userID); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	userInf.Age = userAge
	if userFriends, err = h.friendsService.GetFriendsUser(r.Context(), userID); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	for _, friendID := range userFriends {
		var (
			friend  FriendInformation
			ageUser int
		)

		friend.ID = strconv.Itoa(friendID)
		friend.Name, _ = h.usersService.GetNameUser(r.Context(), friendID)
		ageUser, _ = h.usersService.GetAgeUser(r.Context(), friendID)
		friend.Age = strconv.Itoa(ageUser)
		userInf.Friends.FriendsArray = append(userInf.Friends.FriendsArray, friend)
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, userInf)
}

// @Summary UpdateUser
// @Description User information update
// @Tags Users
// @Accept json
// @Produce html
// @Param user_id path int true "User ID"
// @Param updateUserRequest body UserUpdateRequest false "json update information user"
// @Success 200 {string} string
// @Failure 400 {object} service.ErrResponse
// @Router /{user_id} [put]
func (h *Handler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		userID          int
		userAge         int
		messageResponse string = "возраст пользователя успешно обновлён"
		err             error
		dataRequest     UserUpdateRequest
	)
	userID = r.Context().Value("user_id").(int)
	if err = render.DefaultDecoder(r, &dataRequest); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if userAge, err = strconv.Atoi(dataRequest.Age); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if err = h.usersService.UpdateUserAge(r.Context(), userID, userAge); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.HTML(w, r, messageResponse)
}
