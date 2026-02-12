package main

import "github.com/Raitfolt/url-shortener/internal/config"

func main() {
	cfg := config.MustLoad()
	_ = cfg
}
