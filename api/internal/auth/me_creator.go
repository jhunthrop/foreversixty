package auth

import (
	"context"
	"fmt"
)

// CreatorFacts answers GET /v1/me's creator program fields (creators.Store satisfies it).
type CreatorFacts interface {
	IsCreator(ctx context.Context, userID int64) (bool, error)
	RequestCount(ctx context.Context) (int64, error)
}

// CreatorView is /v1/me's creator program fields, each omitted while it says nothing: creator
// is true for an account in the program (the menu's Creator stats item); creator_requests is
// the waiting applications, for the site's owner only (the menu badge).
type CreatorView struct {
	Creator         bool   `json:"creator,omitempty"`
	CreatorRequests *int64 `json:"creator_requests,omitempty"`
}

func (s *Service) buildCreatorView(ctx context.Context, u User) (CreatorView, error) {
	var v CreatorView
	if s.Creators == nil {
		return v, nil
	}
	creator, err := s.Creators.IsCreator(ctx, u.ID)
	if err != nil {
		return v, fmt.Errorf("auth: me: creator: %w", err)
	}
	v.Creator = creator
	if u.Role != "admin" {
		return v, nil
	}
	n, err := s.Creators.RequestCount(ctx)
	if err != nil {
		return v, fmt.Errorf("auth: me: creator requests: %w", err)
	}
	if n > 0 {
		v.CreatorRequests = &n
	}
	return v, nil
}
