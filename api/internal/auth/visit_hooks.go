package auth

// VisitHooks is several VisitHooks heard in order (the Battle.net refresh and the creator
// program's second-day stamp). Each must return at once, as VisitHook says.
type VisitHooks []VisitHook

// Visited tells every hook about the account.
func (h VisitHooks) Visited(u User) {
	for _, hook := range h {
		hook.Visited(u)
	}
}

// AddVisitHook answers current with hook added: the hook alone when there was none.
func AddVisitHook(current, hook VisitHook) VisitHook {
	switch c := current.(type) {
	case nil:
		return hook
	case VisitHooks:
		return append(append(VisitHooks{}, c...), hook)
	default:
		return VisitHooks{current, hook}
	}
}
