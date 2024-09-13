package http_test

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestGetHandler() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080")
	s.Equal(http.StatusOK, r.StatusCode)
}

func (s *EndToEndSuite) TestPostNoResult() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:8080/health")
	s.Equal(http.StatusOK, r.StatusCode)
}
