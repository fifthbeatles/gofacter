package facter

type hostname_collector struct {
}

func NewHostnameCollector() Collector {
	return &hostname_collector{}
}

func (hc *hostname_collector) Collect() (facts [][2]string) {
	hostname, domain := getHostnameAndDomain()
	if len(hostname) > 0 {
		facts = append(facts, [2]string{"hostname", hostname})
		if len(domain) > 0 {
			facts = append(facts, [2]string{"domain", domain})
			facts = append(facts, [2]string{"fqdn", hostname + "." + domain})
		}
	}
	return
}
