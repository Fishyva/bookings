package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-available", "/availability", "POST", []postData{
		{key: "start",value:"2022-08-09",},
		{key: "end", value:"2022-08-10"},
	}, http.StatusOK},
	{"post-availableJ", "/availability-json", "POST", []postData{
		{key: "start",value:"2022-08-09",},
		{key: "end", value:"2022-08-10"},
	}, http.StatusOK},
	{"make-reservationPost", "/make-reservation", "POST", []postData{
		{key: "first_name",value:"Tommy",},
		{key: "last_name", value:"Smith"},
		{key: "email", value:"tsmither@gmail.com"},
		{key: "phone", value:"4159093454"},
		
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := GetRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal()
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d",e.name,e.expectedStatusCode,resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _,x := range e.params{
				values.Add(x.key,x.value)
			}
			resp,err := ts.Client().PostForm(ts.URL + e.url,values)

			if err != nil {
				t.Log(err)
				t.Fatal()
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d",e.name,e.expectedStatusCode,resp.StatusCode)
			}

		}
	}
}
