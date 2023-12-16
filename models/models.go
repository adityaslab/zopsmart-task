package models

type Train struct {
	TrainNumber int    `json:"train_number"`
	Name        string `json:"name"`
	Status      string `json:"status"`
}

type Platform struct {
	PlatformNumber int `json:"platform_number"`
	TrainNumber    int `json:"train_number"`
}
