package coach

import "slices"

// RenderRules maps the pack onto the review shape mechanically (§5.2), with no model:
// the opening names the boss (or says there is nothing to fix); fix first is rank 1 with
// its detail as the why and drills["drill."+code] as the drill; then fix is ranks 2–4;
// doing well is the top two strengths; couldn't check is copied as is. With no rank 1 (a
// raid buff alone ranks 2) there is no fix first.
func RenderRules(p Pack, drills map[string]string) ReviewDoc {
	doc := ReviewDoc{
		Opening:      Fill(RulesOpening, map[string]string{"Boss": p.Context.Boss}),
		ThenFix:      []ThenItem{},
		DoingWell:    []WellItem{},
		CouldntCheck: slices.Clone(p.CouldNotCheck),
	}
	if len(p.Findings) == 0 {
		doc.Opening = RulesNothing
	}
	for _, f := range p.Findings {
		switch {
		case f.Rank == 1:
			doc.FixFirst = fixItem(p, f, drills)
		case f.Rank <= 1+CapThenItems:
			doc.ThenFix = append(doc.ThenFix, ThenItem{Finding: f.ID, Title: f.Title, Why: f.Detail,
				Anchors: slices.Clone(f.Anchors), Links: linksOf(p, f.ID)})
		}
	}
	for _, s := range p.Strengths[:min(len(p.Strengths), CapWellItems)] {
		doc.DoingWell = append(doc.DoingWell, WellItem{Strength: s.ID, Text: s.Sentence, Anchors: slices.Clone(s.Anchors)})
	}
	return doc
}

func fixItem(p Pack, f Finding, drills map[string]string) *FixItem {
	drill := drillIDPrefix + string(f.Code)
	return &FixItem{Finding: f.ID, Title: f.Title, Why: f.Detail, Drill: drill, DrillText: drills[drill],
		Anchors: slices.Clone(f.Anchors), Links: linksOf(p, f.ID)}
}

// linksOf is the ids of the pack's links that join finding id to another.
func linksOf(p Pack, id string) []string {
	out := []string{}
	for _, l := range p.Links {
		if l.Findings[0] == id || l.Findings[1] == id {
			out = append(out, l.ID)
		}
	}
	return out
}
