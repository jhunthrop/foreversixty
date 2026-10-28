// api/internal/rating/viewer.go
package rating

import (
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// Accounts is the one method this package needs from auth.Store: whether userID is a
// verified member of guildID, and at what rank. It is reports.GuildRanker, so a value passes
// straight to reports.MayView.
type Accounts = reports.GuildRanker
