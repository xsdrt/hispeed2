package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_page(t *testing.T) {
	r, err := http.NewRequest("Get", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error renderering page", err)
	}

	testRenderer.Renderer = "jet"
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error renderering page", err)
	}
}
