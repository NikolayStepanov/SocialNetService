package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"sync"
)

var (
	instance *Router
	once     sync.Once
)

type Router struct {
	mux *chi.Mux
}

func (r *Router) GetMux() *chi.Mux {
	return r.mux
}

func (r *Router) Init() {
	r.mux = chi.NewRouter()
	r.mux.Use(middleware.RequestID)
	r.mux.Use(middleware.Logger)
	r.mux.Use(middleware.Recoverer)
	r.mux.Use(middleware.URLFormat)
	r.mux.Use(render.SetContentType(render.ContentTypeJSON))
}

func (h *Handler) MountRoutes(r *Router) {
	r.GetMux().Mount("/user", h.initUserRoutes())
	r.GetMux().Mount("/create", h.initCreateRoutes())
	r.GetMux().Mount("/make_friends", h.initMakeFriendsRoutes())
	r.GetMux().Mount("/friends", h.initFriendsRoutes())
	r.GetMux().Mount("/", h.initRootRoutes())
}

func GetRouter() *Router {
	once.Do(func() {
		instance = new(Router)
		instance.Init()
	})
	return instance
}
