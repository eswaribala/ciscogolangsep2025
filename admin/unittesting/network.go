package main

import (
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
)

func IsServiceUp(url string) bool {

	//is service up
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func IsLatencyAcceptable(url string, threshold int64) bool {
	//is latency acceptable
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	latency := resp.Header.Get("X-Response-Time")
	return latency != ""
}

func PacketLossRate(url string, threshold float64) bool {
	//is packet loss rate acceptable
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	packetLoss := resp.Header.Get("X-Packet-Loss")
	return packetLoss != ""
}


func GETOTP(min int, max int) int {
	return gofakeit.Number(min, max)
}