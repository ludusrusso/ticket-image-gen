package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTicket(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/upper?name=user&ticket=10", nil)
	ticket, err := buildTicketFromUrl(req.URL)
	assert.Nil(t, err)
	assert.Equal(t, ticket.UserName, "user")
}

func TestGenerateImage(t *testing.T) {
	name := "user"
	color := "blue"
	event := "Test%20Event"
	loc := "Somewhere%20in%20the%20world"
	avatar := "https://dummyimage.com/500"
	date := strings.ReplaceAll("Domenica 04 Settembre 2022", " ", "%20")
	time := strings.ReplaceAll("17:00 - 20:00", " ", "%20")
	ticketNum := 123
	url := fmt.Sprintf("/ticket?name=%s&color=%s&event=%s&loc=%s&avatar=%s&ticket=%d&date=%v&time=%v", name, color, event, loc, avatar, ticketNum, date, time)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	handleTicket(w, req)
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, http.StatusOK)
	assert.Equal(t, "image/png", res.Header.Get("Content-Type"))

	_, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
}
