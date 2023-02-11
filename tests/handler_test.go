package tests

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

const cType = "application/json; charset=utf-8"
const LocalhostURL = "http://localhost:8080"

type Response struct {
	Status int
	Body   string
}

func TestCreateUser(t *testing.T) {
	client := &http.Client{}
	resWant := Response{201, "{\"id\":\"11\"}\n"}
	r := strings.NewReader(`{"name":"Nikolay Stepanov", "age":"27"}`)
	req, err := http.NewRequest(http.MethodPost, LocalhostURL+"/create", r)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", cType)
	res, err := client.Do(req)
	sc := res.StatusCode
	b, _ := io.ReadAll(res.Body)
	res.Body.Close()
	if sc != resWant.Status {
		t.Fatalf("unexpected status, want %v, got %v\n%s\n", resWant.Status, sc, b)
	}
	if string(b) != resWant.Body {
		t.Fatalf("unexpected response body, want:\n%s\ngot:\n%s\n", resWant.Body, b)
	}
}

func TestDeleteUser(t *testing.T) {
	client := &http.Client{}
	resWant := Response{200, "{\"name\":\"Vassily Petrov\"}\n"}
	r := strings.NewReader(`{"target_id":"1"}`)
	req, err := http.NewRequest(http.MethodDelete, LocalhostURL+"/user", r)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", cType)
	res, err := client.Do(req)
	sc := res.StatusCode
	b, _ := io.ReadAll(res.Body)
	res.Body.Close()
	if sc != resWant.Status {
		t.Fatalf("unexpected status, want %v, got %v\n%s\n", resWant.Status, sc, b)
	}
	if string(b) != resWant.Body {
		t.Fatalf("unexpected response body, want:\n%s\ngot:\n%s\n", resWant.Body, b)
	}
}

func TestUpdateAgeUser(t *testing.T) {
	client := &http.Client{}
	resWant := Response{200, "возраст пользователя успешно обновлён"}
	r := strings.NewReader(`{"new age":"31"}`)
	req, err := http.NewRequest(http.MethodPut, LocalhostURL+"/8", r)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", cType)
	res, err := client.Do(req)
	sc := res.StatusCode
	b, _ := io.ReadAll(res.Body)
	res.Body.Close()
	if sc != resWant.Status {
		t.Fatalf("unexpected status, want %v, got %v\n%s\n", resWant.Status, sc, b)
	}
	if string(b) != resWant.Body {
		t.Fatalf("unexpected response body, want:\n%s\ngot:\n%s\n", resWant.Body, b)
	}
}

func TestGetInformationUser(t *testing.T) {
	client := &http.Client{}
	resWant := Response{200, "{\"name\":\"David Freeman\",\"age\":43,\"friends\":" +
		"{\"friends\":[{\"id\":\"2\",\"name\":\"Mark Anderson\",\"age\":\"30\"},{\"id\":\"4\",\"" +
		"name\":\"Frank Fields\",\"age\":\"24\"},{\"id\":\"5\",\"name\":\"Richard Hardy\",\"age\":\"29\"}]}}\n"}
	r := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodGet, LocalhostURL+"/3", r)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", cType)
	res, err := client.Do(req)
	sc := res.StatusCode
	b, _ := io.ReadAll(res.Body)
	res.Body.Close()
	if sc != resWant.Status {
		t.Fatalf("unexpected status, want %v, got %v\n%s\n", resWant.Status, sc, b)
	}
	if string(b) != resWant.Body {
		t.Fatalf("unexpected response body, want:\n%s\ngot:\n%s\n", resWant.Body, b)
	}
}

func TestGetAllFriends(t *testing.T) {
	client := &http.Client{}
	resWant := Response{200, "[{\"id\":\"2\",\"name\":\"Mark Anderson\",\"age\":\"30\"},{\"id\":\"3\"," +
		"\"name\":\"David Freeman\",\"age\":\"43\"},{\"id\":\"5\",\"name\":\"Richard Hardy\",\"age\":\"29\"}]\n"}
	r := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodGet, LocalhostURL+"/friends/4", r)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", cType)
	res, err := client.Do(req)
	sc := res.StatusCode
	b, _ := io.ReadAll(res.Body)
	res.Body.Close()
	if sc != resWant.Status {
		t.Fatalf("unexpected status, want %v, got %v\n%s\n", resWant.Status, sc, b)
	}
	if string(b) != resWant.Body {
		t.Fatalf("unexpected response body, want:\n%s\ngot:\n%s\n", resWant.Body, b)
	}
}

func TestMakeFriends(t *testing.T) {
	client := &http.Client{}
	resWant := Response{200, "Jennifer Cruz и Anna Adams теперь друзья"}
	r := strings.NewReader(`{"source_id": "9", "target_id": "10"}`)
	req, err := http.NewRequest(http.MethodPost, LocalhostURL+"/make_friends", r)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", cType)
	res, err := client.Do(req)
	sc := res.StatusCode
	b, _ := io.ReadAll(res.Body)
	res.Body.Close()
	if sc != resWant.Status {
		t.Fatalf("unexpected status, want %v, got %v\n%s\n", resWant.Status, sc, b)
	}
	if string(b) != resWant.Body {
		t.Fatalf("unexpected response body, want:\n%s\ngot:\n%s\n", resWant.Body, b)
	}
}
