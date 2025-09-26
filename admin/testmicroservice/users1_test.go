package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Users Service", func() {
	var srv *httptest.Server
	var svc *Service

	after := func() {
		if srv != nil {
			srv.Close()
		}
	}

	When("FetchUsers returns 200", func() {
		BeforeEach(func() {
			payload := []User{
				{ID: 1, Name: "Leanne Graham", Username: "Bret", Email: "Sincere@april.biz"},
			}
			srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Expect(r.URL.Path).To(Equal("/users"))
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(payload)
			}))
			svc = New(srv.URL, &http.Client{Timeout: time.Second})
		})
		AfterEach(after)

		It("decodes users", func(ctx SpecContext) {
			child, cancel := context.WithTimeout(ctx, 800*time.Millisecond)
			defer cancel()

			out, err := svc.FetchUsers(child)
			Expect(err).NotTo(HaveOccurred())
			Expect(out).To(HaveLen(1))
			Expect(out[0].Username).To(Equal("Bret"))
		}, SpecTimeout(2*time.Second))
	})

	When("Using set and hex", func() {
		It("extracts unique domains and hex-encodes IDs", func() {
			us := []User{
				{ID: 15, Email: "a@foo.io"},
				{ID: 16, Email: "b@bar.io"},
				{ID: 17, Email: "c@foo.io"},
			}
			s := UniqueEmailDomains(us)
			Expect(s.Contains("foo.io")).To(BeTrue())
			Expect(s.Contains("bar.io")).To(BeTrue())
			Expect(s.Cardinality()).To(Equal(2))
			Expect(HexID(15)).To(Equal("0xf"))
		})
	})
})
