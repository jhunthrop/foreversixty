package main

import (
	"context"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
)

// simsAccountView adapts *sims.Store to auth.SimsView. auth cannot import sims (sims
// imports auth), so this is the one place the two row types meet.
type simsAccountView struct {
	store *sims.Store
}

func (v simsAccountView) FinishedUnseen(ctx context.Context, userID int64, limit int) ([]auth.FinishedSim, error) {
	runs, err := v.store.FinishedUnseen(ctx, userID, limit)
	if err != nil {
		return nil, err
	}
	out := make([]auth.FinishedSim, 0, len(runs))
	for _, r := range runs {
		out = append(out, auth.FinishedSim{
			SimID: r.SimID, Spec: r.Spec, Title: r.Title, DPS: r.DPS,
			FinishedAt: r.FinishedAt.UTC().Format(time.RFC3339),
		})
	}
	return out, nil
}

func (v simsAccountView) RunsThisMonth(ctx context.Context, userID int64) (int, error) {
	return v.store.RunsThisMonth(ctx, userID)
}
