package main

import (
	"context"
	"demo4/middleware"
	"demo4/restaurant-service/proto/restaurant"

	"github.com/go-log/log"
)

type service struct {
	repo Repository
}

func (s *service) Book(ctx context.Context, req *restaurant.Request, rsp *restaurant.Response) (err error) {
	defer middleware.Trace(ctx, req, rsp, err)
	id := req.Id
	log.Logf("receiveid id:", id)
	res, err := s.repo.Book(id)
	if err != nil {
		log.Logf("err:", err)
		return err
	}
	rsp.Msg = res
	return nil
}
