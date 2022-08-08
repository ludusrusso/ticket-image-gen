package main

import (
	"context"

	"github.com/ludusrusso/ticket-image-gen/pkg/server"
)

const (
	outputFilename = "./output.png"
)

func main() {
	// t := ticket.Ticket{
	// 	BaseColor:     "#ec4899",
	// 	UserName:      "Ludovico Russo",
	// 	UserAvatar:    "https://res.cloudinary.com/dbdvy5b2z/image/upload/c_scale,w_60,f_auto/fy/authors/silvia_weqxvf.jpg",
	// 	EventName:     "Picnic in GxP",
	// 	TicketNum:     10,
	// 	EventLocation: "Parso Sempione, Milano",
	// 	EventDate:     "Domenica 04 Settembre 2022",
	// 	EventHours:    "17:00 - 20:00",
	// }
	// img, err := t.Draw()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s\n", err)
	// 	os.Exit(1)
	// }

	// if err := img.SavePNG(outputFilename); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s\n", err)
	// 	os.Exit(1)
	// }

	server.Run(context.Background())
}
