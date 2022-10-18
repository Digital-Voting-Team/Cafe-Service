package helpers

import (
	"Cafe-Service/internal/data"
	"context"

	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	addressesQCtxKey
	cafesQCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxAddressesQ(entry data.AddressesQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, addressesQCtxKey, entry)
	}
}

func AddressesQ(r *http.Request) data.AddressesQ {
	return r.Context().Value(addressesQCtxKey).(data.AddressesQ).New()
}

func CtxCafesQ(entry data.CafesQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, cafesQCtxKey, entry)
	}
}

func CafesQ(r *http.Request) data.CafesQ {
	return r.Context().Value(cafesQCtxKey).(data.CafesQ).New()
}
