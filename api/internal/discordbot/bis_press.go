package discordbot

import "context"

// bisOptions are the resolved options a /bis answer is computed from (and Share stores).
func bisOptions(specKey string, p bisPicks) map[string]string {
	return map[string]string{optSpec: specKey, optPhase: p.Range, optSlot: p.Slot}
}

// picksOfOptions are the card picks of a resolved /bis answer; empty class and spec when the
// spec is unknown.
func picksOfOptions(opts map[string]string) bisPicks {
	p := bisPicks{Range: opts[optPhase], Slot: opts[optSlot]}
	if sp, ok := specByKey(opts[optSpec]); ok {
		p.Class, p.Spec = sp.ClassSlug, sp.SpecSlug
	}
	return p
}

// withChangePicks puts Change picks in the answer's first row, just before its last button
// (the link).
func withChangePicks(msg Message, p bisPicks) Message {
	if len(msg.Components) == 0 || len(msg.Components[0]) == 0 {
		return msg
	}
	change := Component{Kind: KindButton, Style: StyleSecondary, Label: guideCopy.BisChange,
		CustomID: cardCustomID(prefixBis, "chg", p.state()...)}
	rows := append([]ActionRow(nil), msg.Components...)
	row := rows[0]
	last := len(row) - 1
	rows[0] = append(append(append(ActionRow(nil), row[:last]...), change), row[last:]...)
	msg.Components = rows
	return msg
}

// decorateOwn adds a command's private controls to the presser's own copy of an answer: on
// /bis the page button and Change picks. The public follow-up is never decorated.
func decorateOwn(command string, opts map[string]string, msg Message) Message {
	if command == "bis" {
		p := picksOfOptions(opts)
		return withPageButton(withChangePicks(msg, p), opts, p)
	}
	return msg
}

// bisAnswer is the /bis answer for resolved options with its Share and Change picks buttons,
// ephemeral. The typed command and the card presses both use it. update edits the card in place.
func (s *Service) bisAnswer(ctx context.Context, in Interaction, guildID int64, viewer *int64, opts map[string]string, update bool) (Reply, error) {
	msg, found, err := s.answerBis(ctx, guildID, viewer, opts)
	if err != nil {
		return Reply{}, err
	}
	if !found {
		rep := failureReply()
		rep.Update = update
		return rep, nil
	}
	button, err := s.offerShare(ctx, in, guildID, "bis", opts)
	if err != nil {
		return Reply{}, err
	}
	msg = withShare(decorateOwn("bis", opts, msg), button)
	msg.Ephemeral = true
	return Reply{Message: msg, Update: update}, nil
}

// bisPress answers the presses on a /bis card.
func (s *Service) bisPress() Handler {
	return s.cardHandler("/bis", map[string]cardPart{
		"c": s.bisPickClass, "s": s.bisPickSpec, "r": s.bisPickRange,
		"go": s.bisShow, "me": s.bisUseMain, "chg": s.bisChange, partPage: s.bisPage,
	})
}

// bisSeat is who pressed: the server's guild, the linked account and its main.
type bisSeat struct {
	guildID int64
	viewer  *int64
	main    *bisMain
}

// bisSeatOf reads the presser; ok is false with the answer to give when the server is unbound.
func (s *Service) bisSeatOf(ctx context.Context, in Interaction, withMain bool) (bisSeat, Reply, bool, error) {
	guildID, viewer, bound, err := s.lookupCaller(ctx, in)
	if err != nil {
		return bisSeat{}, Reply{}, false, err
	}
	if !bound {
		return bisSeat{}, pressReply(commonCopy.Unbound, true), false, nil
	}
	seat := bisSeat{guildID: guildID, viewer: viewer}
	if withMain {
		seat, err = s.seatWithMain(ctx, seat)
	}
	return seat, Reply{}, err == nil, err
}

// seatWithMain reads the presser's main. Presses that answer for picks already made skip it.
func (s *Service) seatWithMain(ctx context.Context, seat bisSeat) (bisSeat, error) {
	main, err := s.mainPicks(ctx, seat.guildID, seat.viewer)
	seat.main = main
	return seat, err
}

// bisCardUpdate re-renders the card in place.
func (s *Service) bisCardUpdate(seat bisSeat, p bisPicks, lead string) Reply {
	card := bisCard{Picks: p, Lead: lead, Main: seat.main, Linked: seat.viewer != nil}
	return Reply{Message: s.bisCardMessage(card), Update: true}
}

// picked is the first value a select sent.
func picked(in Interaction) string {
	if v := in.Values(); len(v) > 0 {
		return v[0]
	}
	return ""
}

// bisPresser runs a card part for the presser's seat.
func (s *Service) bisPresser(withMain bool, run func(ctx context.Context, in Interaction, seat bisSeat, p bisPicks) (Reply, error)) cardPart {
	return func(ctx context.Context, in Interaction, press cardPress) (Reply, error) {
		seat, refused, ok, err := s.bisSeatOf(ctx, in, withMain)
		if err != nil || !ok {
			return refused, err
		}
		return run(ctx, in, seat, bisPicksOf(press.State))
	}
}

func (s *Service) bisPickClass(ctx context.Context, in Interaction, press cardPress) (Reply, error) {
	return s.bisPresser(true, func(_ context.Context, in Interaction, seat bisSeat, p bisPicks) (Reply, error) {
		p = bisPicksOf([]string{picked(in), "", p.Range, p.Slot})
		if seat.main != nil && p.Class != "" && p.Class == seat.main.Picks.Class {
			p.Spec = seat.main.Picks.Spec
			return s.bisCardUpdate(seat, p, fill(guideCopy.BisLeadMain, "{Name}", memberText(seat.main.Name))), nil
		}
		return s.bisCardUpdate(seat, p, bisLeadFor(p, seat.main, seat.viewer != nil)), nil
	})(ctx, in, press)
}

func (s *Service) bisPickSpec(ctx context.Context, in Interaction, press cardPress) (Reply, error) {
	return s.bisPresser(true, func(_ context.Context, in Interaction, seat bisSeat, p bisPicks) (Reply, error) {
		p = bisPicksOf([]string{p.Class, picked(in), p.Range, p.Slot})
		lead := bisLeadFor(p, seat.main, seat.viewer != nil)
		if p.Spec != "" {
			lead = guideCopy.BisLeadReady
		}
		return s.bisCardUpdate(seat, p, lead), nil
	})(ctx, in, press)
}

func (s *Service) bisPickRange(ctx context.Context, in Interaction, press cardPress) (Reply, error) {
	return s.bisPresser(true, func(_ context.Context, in Interaction, seat bisSeat, p bisPicks) (Reply, error) {
		p = bisPicksOf([]string{p.Class, p.Spec, picked(in), p.Slot})
		return s.bisCardUpdate(seat, p, bisLeadFor(p, seat.main, seat.viewer != nil)), nil
	})(ctx, in, press)
}

// bisShow is Show BiS: the answer when a spec is picked, else the nudge with every pick kept.
func (s *Service) bisShow(ctx context.Context, in Interaction, press cardPress) (Reply, error) {
	return s.bisPresser(false, func(ctx context.Context, in Interaction, seat bisSeat, p bisPicks) (Reply, error) {
		key, hasSpec := p.specKey()
		if hasSpec {
			return s.bisAnswer(ctx, in, seat.guildID, seat.viewer, bisOptions(key, p), true)
		}
		lead := guideCopy.BisNudgeSpec
		if p.Class == "" {
			lead = guideCopy.BisNudgeBoth
		}
		seat, err := s.seatWithMain(ctx, seat)
		if err != nil {
			return Reply{}, err
		}
		return s.bisCardUpdate(seat, p, lead), nil
	})(ctx, in, press)
}

// bisUseMain answers for the main's spec and range. The main is read from the link, never
// from the id.
func (s *Service) bisUseMain(ctx context.Context, in Interaction, press cardPress) (Reply, error) {
	return s.bisPresser(true, func(ctx context.Context, in Interaction, seat bisSeat, p bisPicks) (Reply, error) {
		if seat.main == nil {
			rep := failureReply()
			rep.Update = true
			return rep, nil
		}
		picks := seat.main.Picks
		picks.Slot = p.Slot
		key, _ := picks.specKey()
		return s.bisAnswer(ctx, in, seat.guildID, seat.viewer, bisOptions(key, picks), true)
	})(ctx, in, press)
}

// bisChange reopens the card with the answer's picks.
func (s *Service) bisChange(ctx context.Context, in Interaction, press cardPress) (Reply, error) {
	return s.bisPresser(true, func(_ context.Context, _ Interaction, seat bisSeat, p bisPicks) (Reply, error) {
		return s.bisCardUpdate(seat, p, bisLeadFor(p, seat.main, seat.viewer != nil)), nil
	})(ctx, in, press)
}
