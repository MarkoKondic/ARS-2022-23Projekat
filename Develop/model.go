package main

type Config struct {
	Id      string            `json:"id"`
	Entries map[string]string `json:"entries"`
}

type ConfigGroup struct {
	Id      string             `json:"id"`
	Configs map[string]*Config `json:"configs"`
}