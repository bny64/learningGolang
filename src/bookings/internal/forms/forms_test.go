package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length valid for non-existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedData := url.Values{}
	postedData.Add("some_field", "some value")

	form = New(postedData)
	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows min length of 100 met when field is shorter")
	}

	isError = form.Errors.Get("some_field")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("another_field", "abc12345")
	form = New(postedData)

	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows min length of 1 is not met when field is 8 characters long")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("email", "invalid_email")
	form := New(postedData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("shows valid email for invalid email address")
	}

	postedData = url.Values{}
	postedData.Add("email", "me@here.com")
	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("shows invalid email for valid email address")
	}
}
