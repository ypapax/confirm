package main

import (
	"fmt"
	"strings"

	"github.com/golang/glog"
)

func Yes(msg string) (bool, error) {
	var s string

	fmt.Printf(msg + " (y/N): ")
	_, err := fmt.Scan(&s)
	if err != nil {
		glog.Error(err)
		return false, err
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true, nil
	}
	return false, nil
}
