package search

import (
	"context"
	"math/rand"
	"time"
)

type Provider interface {
	Search(ctx context.Context) ([]Offer, error)
}

func NewProvider(name string) Provider {
	return &provider{
		name: name,
	}
}

type provider struct {
	name string
}

func (p provider) Search(ctx context.Context) ([]Offer, error) {
	timer := time.NewTimer(time.Duration(rand.Intn(10)) * time.Second)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-timer.C:
		return []Offer{
			{
				Provider: p.name,
				Price:    int64(rand.Intn(10000)),
			},
		}, nil
	}
}
