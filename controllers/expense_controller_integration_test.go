//go:build integration

package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/walkmanrd/assessment/models"
)

func TestGetAllExpense(t *testing.T) {
	seedExpense(t)
	var es []models.Expense

	res := request(http.MethodGet, uri("expenses"), nil)
	err := res.Decode(&es)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.Greater(t, len(es), 0)
}

func TestCreateUser(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`)
	var e models.Expense

	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&e)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, e.ID)
	assert.Equal(t, "strawberry smoothie", e.Title)
	assert.Equal(t, float64(79), float64(e.Amount))
	assert.Equal(t, "night market promotion discount 10 bath", e.Note)
	assert.Equal(t, []string{"food", "beverage"}, e.Tags)
}

func TestGetExpenseByID(t *testing.T) {
	e := seedExpense(t)

	var latest models.Expense
	res := request(http.MethodGet, uri("expenses", e.ID), nil)
	err := res.Decode(&latest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, e.ID, latest.ID)
	assert.NotEmpty(t, latest.Title)
	assert.NotEmpty(t, latest.Note)
}

func seedExpense(t *testing.T) models.Expense {
	var e models.Expense
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`)
	err := request(http.MethodPost, uri("expenses"), body).Decode(&e)
	if err != nil {
		t.Fatal("can't create expense: ", err)
	}
	return e
}

func TestUpdateUserByID(t *testing.T) {
	e := seedExpense(t)

	var latest models.Expense
	res := request(http.MethodGet, uri("expenses", e.ID), nil)
	err := res.Decode(&latest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var e2 models.Expense
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie for update",
		"amount": 80,
		"note": "night market promotion discount 10 bath for update", 
		"tags": ["food", "beverage", "for update"]
	}`)
	res2 := request(http.MethodPut, uri("expenses/"+latest.ID), body)
	err2 := res2.Decode(&e2)

	assert.Nil(t, err2)
	assert.Equal(t, http.StatusOK, res2.StatusCode)
	assert.NotEqual(t, 0, e2.ID)
	assert.Equal(t, "strawberry smoothie for update", e2.Title)
	assert.Equal(t, float64(80), float64(e2.Amount))
	assert.Equal(t, "night market promotion discount 10 bath for update", e2.Note)
	assert.Equal(t, []string{"food", "beverage", "for update"}, e2.Tags)
}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}
