package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ehrktia/demo/service"
)

type WebServer struct {
	userService *service.UserService
	mux         *http.ServeMux
	server      *http.Server
}

func NewWebServer(us *service.UserService, port string) *WebServer {
	mux := http.NewServeMux()
	srv := &http.Server{}
	srv.Addr = fmt.Sprintf("[::]:%s", port)
	srv.Handler = mux
	return &WebServer{
		mux:         mux,
		userService: us,
		server:      srv,
	}
}

func (ws *WebServer) HTTPServer() *http.Server {
	return ws.server
}


func (ws *WebServer) RegisterRoutes() {
	ws.mux.Handle("GET /user", ws.getUserByIdHandler())
}

func (ws *WebServer) getUserByIdHandler() http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		values := u.Query()
		userId := values.Get("id")
		id, err := strconv.Atoi(userId)
		if err != nil {
			http.Error(w, errors.New("invalid id received in url").Error(), http.StatusBadRequest)
			return
		}
		if id < 1 || id > math.MaxInt {
			http.Error(w, errors.New("invalid user id received").Error(), http.StatusBadRequest)
			return
		}
		user, err := ws.userService.GetUserById(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resultData, err := json.Marshal(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(resultData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	return http.HandlerFunc(f)
}
