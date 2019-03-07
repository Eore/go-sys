package broker

import (
	"errors"
	"fmt"
	"log"
)

func (p *Pool) AddClient(c Client) (*Pool, error) {
	log.Printf("Adding %+v", c)
	for _, el := range p.Clients {
		if el.Domain == c.Domain {
			msg := fmt.Sprintf("%s exist\n", c.Domain)
			return nil, errors.New(msg)
		}
	}
	p.Clients = append(p.Clients, c)
	return p, nil
}

func (p *Pool) RemoveClient(domain string) (*Pool, error) {
	for i, el := range p.Clients {
		if el.Domain == domain {
			p.Clients = append(p.Clients[:i], p.Clients[i+1:]...)
			return p, nil
		}
	}
	msg := fmt.Sprintf("%s not exist\n", domain)
	return nil, errors.New(msg)
}

func (p *Pool) SendMessage(m Message) {
	for _, el := range p.Clients {
		if el.Domain == m.Domain {
			el.Connection.WriteJSON(m)
		}
	}
}
