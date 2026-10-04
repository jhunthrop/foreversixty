# Guild tab — would a rank-and-file raider use it? — wow-player — 2026-10-04

Reviewed as a guild member, not an officer. Read: `GuildView.lua` (whole file), §4.5.5/§10 of
`design/specs/2026-10-04-addon.md`, the guild membership spec, `ForeverSixtyData/Data.lua`,
`companion/README.md`, `web/src/components/Guild*.svelte`, and
`design/reviews/2026-10-04-addon-value-wow-player.md`. No `addon-guild.png` board exists yet in
`design/mocks/renders/` as of this review.

## 1. What the Guild tab shows today

Header is my guild's name and my in-game rank (`GetGuildInfo`, `L.guildRank`), then a public
standing block from the Forever Sixty Data addon if it's installed and knows my guild: a
progress string, nights logged, roster size (`Ratings.forGuild`, `Data.lua`'s
`guilds[region:ruleset:slug]`). Below that, a private line from the companion's inbox — claim
state (`unclaimed`/`pending`/`claimed`/`contested`) and, only if I'm an officer, a pending-
approvals count (`guildMessage`, `stateLine`, `L.guildClaimPending` etc). Then a roster list,
sorted by rating, unrated members reading "No rating yet" (`rosterOf`, `Ratings.forCharacter`).
No data addon → `ratingsNotInstalled` warning. Guild not tracked on the site → `guildNotOnSite`.

## 2. What I'd actually look at more than once a week

The roster-by-rating list, maybe — once, out of curiosity, the week I install the Data addon.
Everything else here is a glance-once-per-raid-cycle thing at best: my own rank line never
changes, and the claim/approval line is for officers, not me.

## 3. What I want as a member, and where the data already is

- **Where I stand on my spec vs. my raiders:** the roster list gives rating only, no parse, no
  gear check, no "you're 4th of 6 Fury Warriors." `Ratings.forCharacter` has rating/output/
  survival/mechanics/utility/preparation/activity/fights per character — the addon shows none of
  those breakdown numbers, just the one rating. The site's guild page (`Guild.svelte`) has the
  roster and reports but I saw no per-spec ranking there either.
- **What the guild needs from me before Thursday:** nothing here. No gear-gap, consumable, or
  "bring this build" line anywhere in `GuildView.lua`. The data clearly exists elsewhere (BiS gap
  is the planner's whole job per tenet 14's "one job per page," the companion knows my addon
  export) but it isn't surfaced on this tab at all.
- **Did my logs upload:** not shown here. The companion's own Reports page and `/fs inbox`
  cover upload status; the Guild tab doesn't reference it, and `GuildRosterHandoff.svelte` on the
  site (open in sim/planner per roster row) is an officer/web-side view, not something I see
  in-game.
- **Everything else I'd want** (my trend over the tier, who's behind on enchants) doesn't exist
  on any of the three surfaces yet.

## 4. What I never want to see

The claim-state line (`guildClaimUnclaimed`/`pending`/`claimed`/`contested`) and the pending-
approvals count. I'm not an officer; this is loot-council/admin plumbing bolted onto what should
be my personal gear helper. The earlier addon-value review already called this out exactly
("Cut: the Guild tab's officer-facing claim/pending-approval state on the default surface... put
it behind an officer-only toggle") — I'd make the same cut.

## 5. Verdict

NOT YET. I'd open the Guild tab once out of curiosity, then stop — there's nothing here that
changes what I do before Thursday.

The one thing that would make me open it every raid night: tell me, in one line, where I stand
against my guild's other same-spec raiders this week and what gear gap I should chase before the
next pull — everything else here is plumbing I don't care about.
