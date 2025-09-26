package main

//unit testing for network.go

import (
	"testing"
)

func TestIsServiceUp(t *testing.T) {
	url := "http://jsonplaceholder.typicode.com/posts/1"
	if !IsServiceUp(url) {
		t.Errorf("Expected service to be up")
	}
}

func TestIsLatencyAcceptable(t *testing.T) {
	url := "http://jsonplaceholder.typicode.com/users/1"
	threshold := int64(100)
	if !IsLatencyAcceptable(url, threshold) {
		t.Errorf("Expected latency to be acceptable")
	}

}
func TestPacketLossRate(t *testing.T) {
	url := "http://jsonplaceholder.typicode.com/comments/1"
	threshold := 0.1
	if !PacketLossRate(url, threshold) {
		t.Errorf("Expected packet loss rate to be acceptable")
	}

}


func TestGETOTP(t *testing.T) {
	min := 100000
	max := 999999
	otp := GETOTP(min, max)
	if otp < min || otp > max {
		t.Errorf("Expected OTP to be between %d and %d, got %d", min, max, otp)
	}
	t.Logf("Generated OTP: %d", otp)
}