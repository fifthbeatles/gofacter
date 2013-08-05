package facter

type hostname_collector struct {
}

func NewHostnameCollector() Collector {
	return &hostname_collector{}
}

func (hc *hostname_collector) Collect() (facts []Fact) {
	hostname, domain := getHostnameAndDomain()
	if len(hostname) > 0 {
		facts = append(facts, Fact{"hostname", hostname})
		if len(domain) > 0 {
			facts = append(facts, Fact{"domain", domain})
			facts = append(facts, Fact{"fqdn", hostname + "." + domain})
		}
	}
	return
}
