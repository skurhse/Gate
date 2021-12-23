package httpx

import (
	"github.com/mrNobody95/Gate/services/dolphin/internal/pkg/defaults"
	"net/http"
	"strings"
)

func ContentType(req *http.Request) string {
	ct := defaults.String(req.Header.Get("Content-Type"), req.Header.Get("Accept"))
	ct = strings.TrimSpace(ct)
	var cts []string
	if strings.Contains(ct, ",") {
		cts = strings.Split(ct, ",")
	} else {
		cts = strings.Split(ct, ";")
	}
	for _, c := range cts {
		c = strings.TrimSpace(c)
		if strings.HasPrefix(c, "*/*") {
			continue
		}
		return strings.ToLower(c)
	}
	if ct == "*/*" {
		return ""
	}
	return ct
}