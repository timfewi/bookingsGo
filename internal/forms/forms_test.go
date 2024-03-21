package forms

import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/some-url", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/some-url", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "b")
	postData.Add("c", "c")

	r = httptest.NewRequest("POST", "/some-url", strings.NewReader(postData.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.PostForm = postData
	form = New(r.PostForm)

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/some-url", nil)
	form := New(r.PostForm)

	has := form.Has("whatever", r)

	if has {
		t.Error("form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a", r)
	if !has {
		t.Error("form shows form does not have field when it does")
	}

}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/some-url", nil)
	form := New(r.PostForm)

	postData := url.Values{}
	postData.Add("a", "a")

	r = httptest.NewRequest("POST", "/some-url", strings.NewReader(postData.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.PostForm = postData
	form = New(r.PostForm)

	isMinLength := form.MinLength("a", 3, r)
	if isMinLength {
		t.Error("form shows min length for a field when it does not")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	postData = url.Values{}
	postData.Add("a", "aaaa")

	r = httptest.NewRequest("POST", "/some-url", strings.NewReader(postData.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.PostForm = postData
	form = New(r.PostForm)

	isMinLength = form.MinLength("a", 3, r)
	if !isMinLength {
		t.Error("form shows does not have min length for a field when it does")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/some-url", nil)
	form := New(r.PostForm)

	postData := url.Values{}
	postData.Add("email", "user@example.com")

	r = httptest.NewRequest("POST", "/some-url", strings.NewReader(postData.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.PostForm = postData
	form = New(r.PostForm)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("email is not valid when it should be")
	}

	postData = url.Values{}
	postData.Add("email", "user")

	r = httptest.NewRequest("POST", "/some-url", strings.NewReader(postData.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.PostForm = postData
	form = New(r.PostForm)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("email is valid when it should not be")
	}

}
