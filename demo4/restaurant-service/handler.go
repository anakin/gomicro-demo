package main

import (
	"context"
	"demo4/restaurant-service/proto/restaurant"
	"demo4/tracer"

	"github.com/go-log/log"
)

type service struct {
	repo Repository
}

func (s *service) Book(ctx context.Context, req *restaurant.Request, rsp *restaurant.Response) error {
	id := req.Id
	log.Logf("receiveid id:", id)
	res, err := s.repo.Book(id)
	if err != nil {
		log.Logf("err:", err)
		return err
	}
	tracer.Trace(ctx, req, res)
	rsp.Msg = res
	return nil
}
