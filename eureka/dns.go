package eureka

// golang's net.LookupTXT has "bug" (comments below), so here use miekg/dns to implement lookupTxt()
// Refer to net.LookupTXT():
// Multiple strings in one TXT record need to be
// concatenated without separator to be consistent
// with previous Go resolver.
//
// parameters:
// domain, e.g: txt.zone-cn-hz-1.dev.ms-registry.xf.io or txt.zone-cn-hz-1.dev.ms-registry.xf.io.
// dnsAddr, e.g:  "192.168.20.238:53","192.168.20.239:53"
//func lookupTXT(domain string, dnsAddr ...string) ([]string, time.Duration, error) {
//	// format params
//	domain = strings.TrimRight(domain, ".") + "."
//
//	if len(dnsAddr) > 0 {
//		for i, _ := range dnsAddr {
//			if !strings.Contains(dnsAddr[i], ":") { // validate whether contains port
//				dnsAddr[i] += ":53"
//			}
//		}
//
//	} else {
//		var err error
//		dnsAddr = []string{"192.168.1.6:8761"}
//		if err != nil {
//			log.Errorf("Failed to get DNS Server address from system conf err=%s", err.Error())
//		}
//	}
//
//	// loop lookup txt from all dns server address till success / finish.
//	for _, dnsSvr := range dnsAddr {
//		// query txt record
//		query := new(dns.Msg)
//		query.SetQuestion(domain, dns.TypeTXT)
//		response, err := dns.Exchange(query, dnsSvr)
//		if err != nil {
//			log.Errorf("Failure resolving name %s err=%s, dns=%s", domain, err.Error(), dnsSvr)
//			continue
//		}
//
//		if len(response.Answer) < 1 {
//			err := fmt.Errorf("no Eureka discovery TXT record returned for name=%s, dns=%s", domain, dnsSvr)
//			log.Errorf("no answer for name=%s err=%s", domain, err.Error())
//			continue
//		}
//
//		if response.Answer[0].Header().Rrtype != dns.TypeTXT {
//			err := fmt.Errorf("did not receive TXT record back from query specifying TXT record. This should never happen.")
//			log.Errorf("Failure resolving name %s err=%s, dns=%s", domain, err.Error(), dnsSvr)
//			continue
//		}
//		txt := response.Answer[0].(*dns.TXT)
//		ttl := response.Answer[0].Header().Ttl
//		if ttl < 60 {
//			ttl = 60
//		}
//
//		return txt.Txt, time.Duration(ttl) * time.Second, nil
//	}
//
//	err := errors.New(fmt.Sprintf("Failed to lookup TXT records, dns=%v", dnsAddr))
//	return nil, 0, err
//}
