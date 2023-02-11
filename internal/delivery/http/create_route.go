package http

import (
	. "SocialNetHTTPService/internal/service"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type UserCreateRequest struct {
	Name    string `json:"name"`
	Age     string `json:"age"`
	Friends []int  `json:"friends,omitempty"`
}

type UserCreateResponse struct {
	Id string `json:"id"`
}

func (u *UserCreateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewUserCreateResponse(id string) render.Renderer {
	return &UserCreateResponse{Id: id}
}

func (u *UserCreateRequest) Bind(r *http.Request) error {
	if u.Name == "" || u.Age == "" {
		return errors.New("missing required Name or Age fields")
	}
	return nil
}

func (h *Handler) initCreateRoutes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.CreateUser)
	return r
}

// @Summary CreateUser
// @Description Сreating a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param input body UserCreateRequest true "json information user"
// @Success 200 {object} UserCreateResponse
// @Failure 400 {object} service.ErrResponse
// @Router /create [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		userId  int
		userAge int
		err     error
	)
	data := UserCreateRequest{}

	if err = render.Bind(r, &data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if userAge, err = strconv.Atoi(data.Age); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if userId, err = h.usersService.CreateUser(r.Context(), data.Name, userAge, data.Friends); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, NewUserCreateResponse(strconv.Itoa(userId)))
}
