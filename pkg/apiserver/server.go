package apiserver

import (
	"github.com/expectedsh/expected/pkg/util/cors"
	"github.com/expectedsh/expected/pkg/util/github"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	gh "golang.org/x/oauth2/github"
)

type ApiServer struct {
	Addr         string
	Secret       string
	DashboardURL string
	OAuth        *oauth2.Config
}

func New(addr, secret, dashboardUrl string, githubConfig github.Config) *ApiServer {
	return &ApiServer{
		Addr:         addr,
		Secret:       secret,
		DashboardURL: dashboardUrl,
		OAuth: &oauth2.Config{
			ClientID:     githubConfig.ClientID,
			ClientSecret: githubConfig.ClientSecret,
			Endpoint:     gh.Endpoint,
			Scopes:       []string{"user", "user:email"},
		},
	}
}

func (s *ApiServer) Start() error {
	router := mux.NewRouter()

	router.Use(s.authMiddleware)

	router.HandleFunc("v1/account", s.GetAccount).Methods("GET")
	router.HandleFunc("v1/account/sync", s.SyncAccount).Methods("POST")
	router.HandleFunc("v1/account/regenerate_apikey", s.RegenerateAPIKeyAccount).Methods("POST")

	router.HandleFunc("v1/containers", s.GetContainers).Methods("GET")
	router.HandleFunc("v1/containers", s.CreateContainer).Methods("POST")

	if err := cors.ApplyMiddleware(router); err != nil {
		return err
	}
	return http.ListenAndServe(s.Addr, router)
}
