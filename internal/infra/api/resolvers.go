package api

import (
	"errors"
	"github.com/graphql-go/graphql"
	"scheduler/internal/domain"
)

type Resolver struct {
	service domain.Service
}

func (r *Resolver) ClassesResolver(p graphql.ResolveParams)(interface{}, error) {
	group, ok := p.Args["group"].(string)
	if !ok {
		return nil, errors.New("incorrect group")
	}
	day, ok := p.Args["day"].(int)
	if !ok {
		return nil, errors.New("incorrect day")
	}

	return r.service.GetDayScheduleForGroup(group, day)
}

func (r *Resolver) ScheduleResolver(p graphql.ResolveParams)(interface{}, error) {
	group, ok := p.Args["group"].(string)
	if !ok {
		return nil, errors.New("incorrect group")
	}

	data, err := r.service.GetScheduleForGroup(group)
	return data, err
}
