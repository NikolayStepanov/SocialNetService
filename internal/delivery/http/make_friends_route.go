package http

import (
	. "SocialNetHTTPService/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type MakeFriendsRequest struct {
	SourceID string `json:"source_id"`
	TargetID string `json:"target_id"`
}

func (h *Handler) initMakeFriendsRoutes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.MakeFriends)
	return r
}

// @Summary MakeFriends
// @Description Сreating a new city entry
// @Tags Friends
// @Accept json
// @Produce html
// @Param input body MakeFriendsRequest true "json SourceID,TargetID make friends"
// @Success 200 {string} string
// @Failure 400 {object} service.ErrResponse
// @Router /make_friends [post]
func (h *Handler) MakeFriends(w http.ResponseWriter, r *http.Request) {
	var (
		SourceID        int
		TargetID        int
		nameSourceUser  string
		nameTargetUser  string
		messageResponse string
		err             error
		dataRequest     MakeFriendsRequest
	)

	render.DefaultDecoder(r, &dataRequest)
	SourceID, err = strconv.Atoi(dataRequest.SourceID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	TargetID, err = strconv.Atoi(dataRequest.TargetID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if nameSourceUser, err = h.usersService.GetNameUser(r.Context(), SourceID); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if nameTargetUser, err = h.usersService.GetNameUser(r.Context(), TargetID); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	err = h.friendsService.MakeFriend(r.Context(), SourceID, TargetID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	messageResponse = nameSourceUser + " и " + nameTargetUser + " теперь друзья"
	render.Status(r, http.StatusOK)
	render.HTML(w, r, messageResponse)
}
