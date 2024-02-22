package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/elvinlari/docker-golang/internal/user/domain"
	"github.com/elvinlari/docker-golang/internal/user/mock"
)

type App struct {
	*http.Server
	r  *gin.Engine
}

func TestGetByID(t *testing.T) {
	var ts mock.Service
	tsHTTP := Service{Service: &ts}

	// Mock GetByID() call.
	ts.GetByIDFn = func(id int) (*domain.User, error) {
		if id != 100 {
			t.Fatalf("unexpected id: %d", id)
		}
		return &domain.User{ID: 100, Username: "my-user-1"}, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/users/100", nil)

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	RegisterRoutes(router, &tsHTTP)
	router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.GetByIDInvoked {
		t.Fatal("expected GetByID() to be invoked")
	}
}

func TestList(t *testing.T) {
	var ts mock.Service
	tsHTTP := Service{Service: &ts}

	// Mock List() call.
	ts.ListFn = func() ([]*domain.User, error) {
		return []*domain.User{{ID: 100, Username: "my-user-1"}}, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/users/", nil)

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	RegisterRoutes(router, &tsHTTP)
	router.ServeHTTP(w, r)


	// Validate mock.
	if !ts.ListInvoked {
		t.Fatal("expected List() to be invoked")
	}
}

func TestCreate(t *testing.T) {
	var ts mock.Service
	tsHTTP := Service{Service: &ts}

	// Mock our Create() call.
	ts.CreateFn = func(user *domain.User) (*domain.User, error) {
		if user.Username != "my-user-1" {
			t.Fatalf("unexpected name: %s", user.Username)
		}
		return &domain.User{Username: "my-user-1"}, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	request, err := json.Marshal(&Request{&User{Username: "my-user-1"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
		return
	}
	reader := strings.NewReader(string(request))
	r, _ := http.NewRequest("POST", "/users/", reader)

	// Set JWT token in the request header.
	token := os.Getenv("TEST_JWT_TOKEN")
	if token == "" {
        t.Fatal("TEST_JWT_TOKEN environment variable is not set or has expired")
    }

	r.Header.Set("Authorization", token)

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	RegisterRoutes(router, &tsHTTP)
	router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.CreateInvoked {
		t.Fatal("expected Create() to be invoked")
	}
}

func TestUpdate(t *testing.T) {
	var ts mock.Service
	tsHTTP := Service{Service: &ts}

	// Mock our Update() call.
	ts.UpdateFn = func(user *domain.User) (*domain.User, error) {
		if user.Username != "my-user-1" {
			t.Fatalf("unexpected name: %s", user.Username)
		}
		return &domain.User{Username: "my-user-1"}, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	request, err := json.Marshal(&Request{&User{Username: "my-user-1"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
		return
	}
	reader := strings.NewReader(string(request))
	r, _ := http.NewRequest("PUT", "/users/", reader)

	// Set JWT token in the request header.
	token := os.Getenv("TEST_JWT_TOKEN")
	if token == "" {
        t.Fatal("TEST_JWT_TOKEN environment variable is not set or has expired")
    }

	r.Header.Set("Authorization", token)

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	RegisterRoutes(router, &tsHTTP)
	router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.UpdateInvoked {
		t.Fatal("expected Update() to be invoked")
	}
}

func TestDelete(t *testing.T) {
	var ts mock.Service
	tsHTTP := Service{Service: &ts}

	// Mock Delete() call.
	ts.DeleteFn = func(id int) error {
		if id != 100 {
			t.Fatalf("unexpected id: %d", id)
		}
		return nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/users/100", nil)

	// Set JWT token in the request header.
	token := os.Getenv("TEST_JWT_TOKEN")
    if token == "" {
        t.Fatal("TEST_JWT_TOKEN environment variable is not set or has expired")
    }

	r.Header.Set("Authorization", token)

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	RegisterRoutes(router, &tsHTTP)
	router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.DeleteInvoked {
		t.Fatal("expected Delete() to be invoked")
	}
}
