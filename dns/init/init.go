package init

import (
	"net"

	"github.com/gdziwoki/go/dns"
)

func init() {
	net.DefaultResolver = dns.New()
}
