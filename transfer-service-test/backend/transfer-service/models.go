package main

type Transfer struct {
	ID           int    `json:"id"`
	PalletID     string `json:"pallet_id"`
	FromLocation string `json:"from_location"`
	ToLocation   string `json:"to_location"`
	Note         string `json:"note"`
	Status       string `json:"status"`
}

type Outbox struct {
	ID        int
	EventType string
	Payload   []byte
	Processed bool
}
