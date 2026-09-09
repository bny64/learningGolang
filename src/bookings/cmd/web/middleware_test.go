package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoServer(t *testing.T) {
	var myHandler myHandler

	h := NoSurf(&myHandler)

	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, got %v", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myHandler myHandler

	h := SessionLoad(&myHandler)

	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, got %v", v))
	}
}
