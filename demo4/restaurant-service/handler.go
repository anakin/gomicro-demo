package main

import (
	"context"
	"demo4/lib/tracer"
	"demo4/restaurant-service/proto/restaurant"

	"github.com/sirupsen/logrus"

	"github.com/go-log/log"
)

type service struct {
	repo Repository
}

func (s *service) Book(ctx context.Context, req *restaurant.Request, rsp *restaurant.Response) (err error) {
	defer tracer.Trace(ctx, "Book", req, rsp, err)
	id := req.Id
	log.Logf("receiveid id:", id)
	res, err := s.repo.Book(id)
	if err != nil {
		logrus.Error("err:", err)
		return err
	}
	rsp.Msg = res
	return nil
}
