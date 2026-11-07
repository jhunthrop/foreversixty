// api/internal/account/confirm.go
//
// Package account is self-service account deletion, DELETE /v1/me: the typed confirmation,
// the guild and staff refusals, the Stripe cancel, the Discord revoke, one transaction that
// removes or unlinks everything the account holds, and the storage clean-up after it. What
// is deleted and what is kept, and why, is written beside each step in store.go.
package account

import "strings"

// ConfirmWord is the word anyone may type instead of their BattleTag.
const ConfirmWord = "delete"

// Confirmed is the typed-confirmation rule every client mirrors: the account's BattleTag or
// the word "delete", trimmed and in any case. An empty answer never confirms.
func Confirmed(typed string, battletag *string) bool {
	typed = strings.TrimSpace(typed)
	if typed == "" {
		return false
	}
	if strings.EqualFold(typed, ConfirmWord) {
		return true
	}
	return battletag != nil && *battletag != "" && strings.EqualFold(typed, strings.TrimSpace(*battletag))
}
