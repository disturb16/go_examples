package api

import (
	"encoding/json"
	"gotests/api/dto"
	"gotests/internal/books"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

var a *API

func TestMain(m *testing.M) {

	service := &books.MockBookService{}
	a = New(service)

	code := m.Run()
	os.Exit(code)
}

func TestBooks(t *testing.T) {
	testCases := []struct {
		Name               string
		Params             map[string]string
		ExpectedStatusCode int
	}{
		{
			Name: "ListBooks",
			Params: map[string]string{
				"limit": "10",
			},
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name:               "no params",
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name: "invalid limit",
			Params: map[string]string{
				"limit": "invalid",
			},
			ExpectedStatusCode: http.StatusBadRequest,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			q := make(url.Values)
			for k, v := range tc.Params {
				q.Add(k, v)
			}

			e := echo.New()
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			r.URL.RawQuery = q.Encode()
			w := httptest.NewRecorder()

			c := e.NewContext(r, w)

			err := a.Books(c)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if w.Code != tc.ExpectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.ExpectedStatusCode, w.Code)
			}
		})
	}
}

func TestBookByID(t *testing.T) {
	testCases := []struct {
		Name               string
		BookID             string
		ExpectedStatusCode int
	}{
		{
			Name:               "BookByID",
			BookID:             "1",
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name:               "Invalid Book ID",
			BookID:             "invalid",
			ExpectedStatusCode: http.StatusBadRequest,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()

			c := e.NewContext(r, w)

			c.SetParamNames("id")
			c.SetParamValues(tc.BookID)

			err := a.BookByID(c)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if w.Code != tc.ExpectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.ExpectedStatusCode, w.Code)
			}
		})
	}
}

func TestSaveBook(t *testing.T) {
	testCases := []struct {
		Name               string
		Data               dto.SaveBook
		ExpectedStatusCode int
	}{
		{
			Name: "SaveBook",
			Data: dto.SaveBook{
				Title:  "Test",
				Author: "Test",
			},
			ExpectedStatusCode: http.StatusCreated,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			// setup request
			body, err := json.Marshal(tc.Data)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			e := echo.New()
			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(body)))
			w := httptest.NewRecorder()

			c := e.NewContext(r, w)

			err = a.SaveBook(c)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if w.Code != tc.ExpectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.ExpectedStatusCode, w.Code)
			}
		})
	}

}
