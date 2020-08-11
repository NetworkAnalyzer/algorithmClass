package problem016

import (
	"errors"
	"github.com/thoas/go-funk"
)

// This problem was asked by Twitter.
//
// You run an e-commerce website and want to record the last N order ids in a log.
//
// Implement a data structure to accomplish this, with the following API:
//
// record(order_id): adds the order_id to the log
// get_last(i): gets the ith last element from the log. i is guaranteed to be smaller than or equal to N.
// You should be as efficient with time and space as possible.
//
// log takes O(N) Space due to size
// record and getLast is O(1) Time complexity

type logger struct {
	log []log
}

type log struct {
	orderID int
}

func newLogger() *logger {
	return &logger{
		log: nil,
	}
}

func (l *logger) record(orderID int) error {
	if funk.Contains(l.log, log{orderID: orderID}) {
		return errors.New("orderID already exists")
	}

	l.log = append(l.log, log{orderID: orderID})
	return nil
}

func (l *logger) getLast(i int) []log {
	return l.log[len(l.log) - i:]
}
