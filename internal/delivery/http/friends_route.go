package http

import (
	. "SocialNetHTTPService/internal/service"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type FriendInformation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

type FriendsResponse struct {
	FriendsArray []FriendInformation `json:"friends"`
}

func (h *Handler) FriendsCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			friends []int
			err     error
			userID  int
		)

		if userIDStr := chi.URLParam(r, "user_id"); userIDStr != "" {
			userID, _ = strconv.Atoi(userIDStr)
			friends, err = h.friendsService.GetFriendsUser(r.Context(), userID)
		} else {
			render.Render(w, r, ErrNotFound)
			return
		}
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "friends", friends)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) initFriendsRoutes() chi.Router {
	r := chi.NewRouter()
	r.Route("/{user_id}", func(r chi.Router) {
		r.Use(h.FriendsCtx)
		r.Get("/", h.GetAllFriendsByID)
	})

	return r
}

// @Summary GetAllFriendsByID
// @Description Getting information about user's friends
// @Tags Friends
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} FriendsResponse
// @Failure 400 {object} service.ErrResponse
// @Router /friends/{user_id} [get]
func (h *Handler) GetAllFriendsByID(w http.ResponseWriter, r *http.Request) {
	var (
		friendsUser     []int
		friendsInfo     []FriendInformation
		friendsResponse FriendsResponse
	)

	friendsUser = r.Context().Value("friends").([]int)
	for _, friendID := range friendsUser {
		var (
			friend  FriendInformation
			ageUser int
		)

		friend.ID = strconv.Itoa(friendID)
		friend.Name, _ = h.usersService.GetNameUser(r.Context(), friendID)
		ageUser, _ = h.usersService.GetAgeUser(r.Context(), friendID)
		friend.Age = strconv.Itoa(ageUser)
		friendsInfo = append(friendsInfo, friend)
	}
	friendsResponse.FriendsArray = friendsInfo
	render.JSON(w, r, friendsInfo)
}
