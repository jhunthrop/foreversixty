// api/internal/guilds/application_post.go
//
// A new application posts to the guild's recruiting channel through the bot (guild onboarding
// spec 2026-11-09 §2.3): free for any claimed guild. Off the request, never failing the apply.
package guilds

import (
	"context"
	"strings"
	"time"
)

// applicationPostBudget bounds the recruiting post: it runs after the apply has answered.
const applicationPostBudget = 8 * time.Second

// ApplicationPost is a new application as the recruiting channel shows it. ReviewURL is the
// applications page; SpecClass the character's "Holy Priest" line, empty when unknown.
type ApplicationPost struct {
	ID                                                                int64
	GuildName, Character, SpecClass, Message, Availability, ReviewURL string
	Level, ItemLevel                                                  *int
}

// ApplicationPoster is the bot's recruiting post; Service.ApplicationBot satisfies it when the
// deployment runs the bot. The outcome is OutcomeSent, OutcomeNoBot (no install or no recruiting
// channel) or OutcomeDiscordRefuse.
type ApplicationPoster interface {
	PostApplication(ctx context.Context, guildID int64, post ApplicationPost) (string, error)
}

// postApplication sends application a to the guild's recruiting channel, off the request. A
// failure is logged: the application is filed either way.
func (s *Service) postApplication(ctx context.Context, guildID int64, a Application) {
	poster, ok := s.ApplicationBot.(ApplicationPoster)
	if !ok {
		return
	}
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), applicationPostBudget)
	go func() {
		defer cancel()
		g, err := s.Store.getGuild(ctx, guildID)
		if err != nil {
			s.logger().Error("guilds", "op", "application_post", "guild_id", guildID, "err", err)
			return
		}
		word, err := poster.PostApplication(ctx, guildID, s.applicationPostOf(g, a))
		if err != nil {
			s.logger().Warn("guilds", "op", "application_post", "guild_id", guildID, "err", err)
			return
		}
		s.logger().Info("guilds", "op", "application_post", "guild_id", guildID, "application_id", a.ID, "outcome", word)
	}()
}

func (s *Service) applicationPostOf(g Guild, a Application) ApplicationPost {
	return ApplicationPost{
		ID: a.ID, GuildName: g.Name, Character: a.Character.Name, SpecClass: specClassOf(a.Character),
		Message: a.Message, Availability: a.Availability, Level: a.Character.Level, ItemLevel: a.Character.ItemLevel,
		ReviewURL: strings.TrimRight(s.PublicBaseURL, "/") + GuildPagePath(g) + "/applications",
	}
}

// specClassOf is the applicant's "Holy Priest" (specClassName); empty when nothing is known.
func specClassOf(c ApplicantCharacter) string {
	class := ""
	if c.Class != nil {
		class = *c.Class
	}
	return specClassName(c.SpecName, class)
}

// specClassName is a spec and class in the game's words: the spec name with the class after it
// (unless the name already says the class), the class alone without a spec.
func specClassName(specName *string, classSlug string) string {
	class := className(classSlug)
	if specName == nil || *specName == "" {
		return class
	}
	if class == "" || strings.Contains(strings.ToLower(*specName), strings.ToLower(class)) {
		return *specName
	}
	return *specName + " " + class
}

// className is a class slug in the game's words: "priest" reads "Priest".
func className(slug string) string {
	if slug == "" {
		return ""
	}
	return strings.ToUpper(slug[:1]) + slug[1:]
}
