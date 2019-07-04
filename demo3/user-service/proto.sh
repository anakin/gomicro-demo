#!/bin/sh
protoc --micro_out=. --gogofaster_out=. proto/user/user.proto
