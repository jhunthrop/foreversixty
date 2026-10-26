// Package refusal holds the refusal sentences several packages answer with, so the
// web reads one wording wherever the same refusal is raised.
package refusal

const (
	// ClaimContested is the claim_contested body of every frozen-claim officer action.
	ClaimContested = "this guild's claim is contested; officer actions are frozen until a moderator resolves it"
	// RaidPlanRequired is the guild_plan_required body of the officer raid writes and the channel routes.
	RaidPlanRequired = "Planning and running raids are part of the guild plan."
	// AwardPlanRequired is the guild_plan_required body of loot Award and Undo.
	AwardPlanRequired = "Award and Undo are part of the guild plan."
)
