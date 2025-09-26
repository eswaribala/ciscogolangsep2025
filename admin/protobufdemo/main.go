package main

import (
	"fmt"
	"time"

	v1 "github.com/cisco/admin/protobufdemo/protobufsrc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	site := &v1.Site{
		SiteId:      12345,
		HostName:    "example.com",
		Location:    "Datacenter 1",
		CreatedAt:   timestamppb.Now(),
		Status:      "active",
		Description: "Primary site",
	}

	// Encode (marshal) to []byte
	b, err := proto.Marshal(site)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Serialized bytes (%d): %v\n", len(b), b[:min(16, len(b))])

	// Decode (unmarshal) to struct
	var out v1.Site
	if err := proto.Unmarshal(b, &out); err != nil {
		panic(err)
	}
	fmt.Printf("Decoded: id=%d host=%s location=%s created_at=%T\n", out.SiteId, out.HostName, out.Location, out.CreatedAt)

	// Presence checks (proto3 with well-known types or oneofs)
	if out.CreatedAt != nil && out.CreatedAt.AsTime().After(time.Time{}) {
		fmt.Println("created_at present:", out.CreatedAt.AsTime())
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
