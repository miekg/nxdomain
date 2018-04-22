package dump

import (
	"context"
	"fmt"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"

	"github.com/mholt/caddy"
	"github.com/miekg/dns"
)

// Nxdomain implement the plugin interface.
type Nxdomain struct {
	Next  plugin.Handler
	names []string
}

func init() {
	caddy.RegisterPlugin("nxdomain", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	names := []string{}
	for c.Next() {
		args := c.RemainingArgs()
		if len(args) == 0 {
			return plugin.Error("nxdomain", c.ArgErr())
		}
		// I'll bet these are not fully qualified
		for _, a := range args {
			names = append(names, dns.Fqdn(a))
		}
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Nxdomain{Next: next, names: names}
	})

	return nil
}

// ServeDNS implements the plugin.Handler interface.
func (n Nxdomain) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {

	state := request.Request{W: w, Req: r}

	for _, n := range n.names {
		if dns.IsSubDomain(n, state.Name()) {
			m := new(dns.Msg)
			m.SetRcode(r, dns.RcodeNameError)
			m.Ns = []dns.RR{soa(n)}
			w.WriteMsg(m)
			return 0, nil
		}
	}

	return plugin.NextOrFailure(n.Name(), n.Next, ctx, w, r)
}

// Name implements the Handler interface.
func (n Nxdomain) Name() string { return "nxdomain" }

func soa(name string) dns.RR {
	s := fmt.Sprintf("%s 60 IN SOA ns1.%s postmaster.%s 1524370381 14400 3600 604800 60", name, name, name)
	soa, _ := dns.NewRR(s)
	return soa
}