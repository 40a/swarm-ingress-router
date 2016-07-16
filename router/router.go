package router

import (
	"crypto/tls"
	"log"
	"net/url"
)

type Router struct {
	routes map[string]Routable
}

func (r *Router) RouteToService(address string) (*url.URL, bool) {
	route, ok := r.routes[address]
	if !ok {
		log.Printf("Failed to lookup service for %s", address)
		return &url.URL{}, false
	}

	serviceURL, err := url.Parse(route.URL())

	if err != nil {
		log.Printf("Failed to parse URL for service %s", address)
		return &url.URL{}, false
	}

	return serviceURL, true
}

func (r *Router) CertificateForService(address string) (*tls.Certificate, bool) {
	route, ok := r.routes[address]
	if !ok {
		log.Printf("Failed to lookup service for %s", address)
		return &tls.Certificate{}, false
	}

	return route.Certificate()
}

func (r *Router) UpdateTable(services []Routable) {
	newTable := make(map[string]Routable)
	for _, s := range services {
		log.Printf("Registering service for %s", s.DNSName())
		newTable[s.DNSName()] = s
	}

	r.routes = newTable
}

func NewRouter() *Router {
	return &Router{}
}
