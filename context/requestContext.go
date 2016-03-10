package context

import (
    "errors"
    "net/http"
    "strconv"
)

const (
    API_VERSION = "api-version"
)

type RequestContext struct {
    Version float64
}

func NewRequesetContext(r *http.Request) (*RequestContext, error) {
      h := r.Header.Get("api-version")
      if h == "" {
          return nil, errors.New("version header not provided")
      } else {
          v, err := strconv.ParseFloat(h, 64)
          if err != nil {
              return nil, errors.New("invalid version")
          }
          return &RequestContext{ Version: v }, nil
      }
}
