package server

import (
	"bytes"
	"context"
	"image/jpeg"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ludusrusso/ticket-image-gen/pkg/palette"
	"github.com/ludusrusso/ticket-image-gen/pkg/ticket"
	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context) {
	http.HandleFunc("/ticket", handleTicket)

	logrus.Infof("running server on %s", "http://localhost:8080")

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		logrus.Fatal(err)
	}
}

func handleTicket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	t, err := buildTicketFromUrl(r.URL)
	if err != nil {
		logrus.Errorf("error parsing ticket: %s", err)
		http.Error(w, "error parsing ticket number", http.StatusBadRequest)
		return
	}

	img, err := t.Draw()
	if err != nil {
		logrus.Errorf("err drawing: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img.Image(), nil)
	w.Write(buf.Bytes())
}

func buildTicketFromUrl(url *url.URL) (ticket.Ticket, error) {
	q := url.Query()
	name := q.Get("name")
	eventName := q.Get("event")
	eventLoc := q.Get("loc")
	eventDate := q.Get("date")
	eventTime := q.Get("time")
	ticketNum := q.Get("ticket")
	avatar := q.Get("avatar")
	color := q.Get("color")
	transparent := q.Get("transparent")

	num, err := strconv.ParseInt(ticketNum, 10, 64)
	if err != nil {
		return ticket.Ticket{}, err
	}

	pl, ok := palette.ColorPalettes[color]
	if !ok {
		pl = palette.ColorPalettes["pink"]
	}

	return ticket.Ticket{
		UserName:      name,
		UserAvatar:    avatar,
		EventName:     eventName,
		TicketNum:     uint(num),
		EventLocation: eventLoc,
		EventDate:     eventDate,
		EventHours:    eventTime,
		Palette:       pl,
		BgTransparent: transparent != "",
	}, nil
}
