package search

import (
	"context"
	"math/rand"
	"time"
)

type Provider interface {
	Search(ctx context.Context) []Offer
}

func NewProvider(name string) Provider {
	return &provider{
		name: name,
	}
}

type provider struct {
	name string
}

func (p provider) Search(ctx context.Context) []Offer {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))

	return []Offer{
		{
			Provider: p.name,
			Price:    int64(rand.Intn(10000)),
		},
	}
}
