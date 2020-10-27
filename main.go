package main

import (
	"fmt"
	"go_server/types"
	"net/http"

	"github.com/gorilla/mux"
)

/*
#cgo LDFLAGS: -L /lib64 -lcal-api
#include <stdlib.h>
#include "include/cal_e.h"
*/
import "C"

var (
	httpsrv httpServer
)


// Api general interface
type Api interface {
}

type httpServer struct {
	router    *mux.Router
	eventChan chan types.HttpResponse
}

// New construct new http server
func New(eventChan chan types.HttpResponse) Api {
	srv := &httpServer{
		router:    mux.NewRouter().StrictSlash(false),
		eventChan: eventChan,
	}

	srv.router.HandleFunc("/cal", srv.show).Methods("GET")

	srv.runLocalhost()
	return srv
}

func (s *httpServer) runLocalhost() {
	http.ListenAndServe("127.0.0.1:5003", s.router)
}

func (s *httpServer) show(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	a_s := params.Get("item1")
	b_s := params.Get("item2")
	//if b, err := strconv.Atoi(b_s); err == nil {
	c := C.startcal(C.int(len(a_s)), C.int(len(b_s)))
	fmt.Println(c)

	s.eventChan <- types.HttpResponse{Json: "{\"OK\"}"}
	s.handleResponse(w, r, s.eventChan)
}

func (s *httpServer) handleResponse(w http.ResponseWriter, r *http.Request, rch chan types.HttpResponse) {
	resp := <-rch

	r.Close = true // indicate that the connection should be closed after responding

	if resp.Error == nil {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, resp.Json)
	} else {
		var errorCode = http.StatusInternalServerError

		switch resp.Error {
		case types.ErrProcessNotFound, types.ErrQueryNotFound, types.ErrAvailSvcNotFound:
			errorCode = http.StatusNotFound
		case types.ErrNameAlreadyExists, types.ErrIdAlreadyExists, types.ErrNameAndUrlDoesNotMatch:
			errorCode = http.StatusConflict
		default:
			errorCode = http.StatusInternalServerError
		}

		http.Error(w, resp.Error.Error(), errorCode)
	}
}

func main() {
	eventChan := make(chan types.HttpResponse, 10000)
	httpsrv = New(eventChan).(httpServer)
}
