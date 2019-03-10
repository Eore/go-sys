package broker

import (
	"errors"
	"fmt"
)

type newMap struct {
	domain string `json:"domain,omitempty"`
	ip     string `json:"ip,omitempty"`
}

func (p *Pool) ListClient() []newMap {
	var ret []newMap
	for _, el := range p.Clients {
		ret = append(ret, newMap{
			domain: el.Domain,
			ip:     el.IP,
		})
	}
	return ret
}

func (p *Pool) AddClient(c Client) (*Pool, error) {
	for _, el := range p.Clients {
		if el.Domain == c.Domain {
			msg := fmt.Sprintf("%s exist\n", c.Domain)
			return nil, errors.New(msg)
		}
	}
	p.Clients = append(p.Clients, c)
	return p, nil
}

func (p *Pool) RemoveClient(IP string) (*Pool, error) {
	for i, el := range p.Clients {
		if el.IP == IP {
			p.Clients = append(p.Clients[:i], p.Clients[i+1:]...)
			return p, nil
		}
	}
	msg := fmt.Sprintf("%s not exist\n", IP)
	return nil, errors.New(msg)
}

func (p *Pool) SendMessage(m Message) {
	for _, el := range p.Clients {
		if el.Domain == m.Domain {
			el.Connection.WriteJSON(m)
		}
	}
}
