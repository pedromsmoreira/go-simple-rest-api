package handlers_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/pedromsmoreira/go-simple-rest-api/handlers"
	"github.com/pedromsmoreira/go-simple-rest-api/model"
)

type dummyRepository struct {
	pong string
	err  error
}

func (r dummyRepository) Ping() (string, error) {
	return r.pong, r.err
}

var enc = json.NewEncoder(os.Stdout)

func TestShallowHandler(t *testing.T) {
	dummyRepo := dummyRepository{
		err:  nil,
		pong: "PONGTest",
	}

	hcHandler := handlers.NewHealthCheckHandler(dummyRepo)

	handler := http.HandlerFunc(hcHandler.Shallow)
	ts := httptest.NewServer(handler)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expResp := []model.Shallow{model.NewShallow("Redis", "PONGTest", true)}
	exp, _ := json.Marshal(expResp)

	slurp, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if !bytes.Equal(slurp, exp) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(slurp), string(exp))
	}
}

func TestShallowHandlerParallel(t *testing.T) {
	dummyRepo := dummyRepository{
		err:  nil,
		pong: "PONGTest",
	}

	hcHandler := handlers.NewHealthCheckHandler(dummyRepo)

	handler := http.HandlerFunc(hcHandler.Shallow)
	ts := httptest.NewServer(handler)

	defer ts.Close()
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := http.Get(ts.URL)
			if err != nil {
				t.Error(err)
				return
			}

			slurp, err := ioutil.ReadAll(res.Body)
			defer res.Body.Close()
			if err != nil {
				t.Error(err)
				return
			}
			t.Logf("Got: %s", slurp)

		}()
	}
	wg.Wait()
}

func BenchmarkShallowHandler(b *testing.B) {
	b.ReportAllocs()
	rec := httptest.NewRecorder()
	r := req(b, "GET /healthchecks/shallow HTTP/1.0\r\n\r\n")

	dummyRepo := dummyRepository{
		err:  nil,
		pong: "PONGTest",
	}

	hcHandler := handlers.NewHealthCheckHandler(dummyRepo)

	for i := 0; i < b.N; i++ {
		hcHandler.Shallow(rec, r)
		reset(rec)
	}
}

func reset(rw *httptest.ResponseRecorder) {
	m := rw.HeaderMap
	for k := range m {
		delete(m, k)
	}
	body := rw.Body
	body.Reset()
	*rw = httptest.ResponseRecorder{
		Body:      body,
		HeaderMap: m,
	}
}

func req(t testing.TB, v string) *http.Request {
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(v)))
	if err != nil {
		t.Fatal(err)
	}
	return req
}
