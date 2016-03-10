package handler

import (
    "github.com/jjosephy/goauth_service/context"
    "github.com/jjosephy/goauth_service/logging"
    "net/http"
    "strconv"
)

func TokenHandler() http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        switch r.Method {
            case "POST":
                c, e := context.NewRequesetContext(r)

                if e != nil {
                    w.Write([]byte("errr"))
                    return
                }

                v := strconv.FormatFloat(c.Version, 'E', -1, 64)
                logging.LogEvent(v)

                w.Write([]byte("token"))

            default:
                w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }
    }
}
