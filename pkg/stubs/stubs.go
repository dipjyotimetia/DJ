package stubs

import (
	"bytes"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

func ListenAndServe(cfg *Config) error {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	fmt.Printf("Running the stub server in, %d", cfg.Port)

	for _, service := range cfg.Services {
		for _, endpoint := range service.Endpoints {
			service := service
			endpoint := endpoint
			router.MethodFunc(endpoint.Method, "/"+service.Prefix+endpoint.Name, func(w http.ResponseWriter, r *http.Request) {
				log.Println("request:", endpoint.Method, "/"+service.Prefix+endpoint.Name)

				for k, v := range cfg.Header {
					w.Header().Set(k, v)
				}

				if len(endpoint.Matches) != 0 {
					body, err := ioutil.ReadAll(r.Body)
					if err != nil {
						log.Fatal(err)
					}

					for _, match := range endpoint.Matches {
						reqBody, err := ioutil.ReadFile(cfg.ResponseDir + "/" + match.RequestBody)
						if err != nil {
							log.Fatal(err)
						}

						if !bytes.Contains(body, reqBody) {
							continue
						}

						for k, v := range match.Response.Header {
							w.Header().Set(k, v)
						}
						w.WriteHeader(match.Response.Status)

						responseBody, err := ioutil.ReadFile(cfg.ResponseDir + "/" + match.Response.Body)
						if err != nil {
							log.Fatal(err)
						}

						_, err = w.Write(responseBody)
						if err != nil {
							fmt.Println(err)
						}

						return
					}
					http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				} else {
					for k, v := range endpoint.Response.Header {
						w.Header().Set(k, v)
					}
					w.WriteHeader(endpoint.Response.Status)

					body, err := ioutil.ReadFile(cfg.ResponseDir + "/" + endpoint.Response.Body)
					if err != nil {
						log.Fatal(err)
					}

					_, err = w.Write(body)
					if err != nil {
						log.Fatal(err)
					}
				}
			})
		}
	}
	return http.ListenAndServe(net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)), router)
}
