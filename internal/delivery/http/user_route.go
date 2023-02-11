package http

import (
	. "SocialNetHTTPService/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type DeleteUserRequest struct {
	TargetID string `json:"target_id"`
}

type DeleteUserResponse struct {
	Name string `json:"name"`
}

func (h *Handler) initUserRoutes() chi.Router {
	r := chi.NewRouter()
	r.Delete("/", h.UserDelete)
	return r
}

// @Summary UserDelete
// @Description Delete user information
// @Tags Users
// @Accept json
// @Produce json
// @Param requestDelete body DeleteUserRequest true "json delete targetID User"
// @Success 200 {object} DeleteUserResponse
// @Failure 400 {object} service.ErrResponse
// @Router /user/ [delete]
func (h *Handler) UserDelete(w http.ResponseWriter, r *http.Request) {
	var (
		nameUserDelete string
		deleteUserId   int
		err            error
		dataRequest    DeleteUserRequest
		dataResponse   DeleteUserResponse
	)

	render.Decode(r, &dataRequest)
	if deleteUserId, err = strconv.Atoi(dataRequest.TargetID); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if nameUserDelete, err = h.usersService.DeleteUser(r.Context(), deleteUserId); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	dataResponse.Name = nameUserDelete
	render.Status(r, http.StatusOK)
	render.JSON(w, r, dataResponse)
}
