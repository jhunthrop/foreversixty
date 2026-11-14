# Guild later: comp and bench call, recruiting search

2026-11-15. Short briefs, not full experience specs. Inputs: 00-facts.md, 01-plan.md, 02-officer-plan.md, 03-brief.md. Fixture: Olympus XXVII, Onyxia's Lair, Wed 9 Dec.
Both items build on a surface that already exists (the raid page's cover Grid; the character directory). Neither starts before the owner decisions below.

## 1. Comp and bench call

The question: "I have 30 signed up for 40 spots, or 44 for 40. Which buffs and debuffs do we lose or gain with each player, and who sits?"
Answer with no scroll (360 and 1280): one line, `Missing: Windfury Totem, Sunder Armor`, then the bench list with the effect of each swap, `Bench Brenn: lose Windfury Totem (counted per raid, not per group)`. Coverage count at the head of the Grid as a figure, `16 of 24 covered`.
Reference: the Raid-Helper class tally and the raid planner's buff Grid (web-guild-raid mock), plus the in-game raid frame by group. We add the cost of each swap, which neither shows.

What exists (verified in code):
- Cover catalogue: api/internal/guilds/raids_cover_catalogue.go, generated from data/curated/raid-cover.json. 24 entries: buffs (Windfury, Grace of Air, Strength of Earth, Mana Spring, Battle Shout, Trueshot, Leader of the Pack, Moonkin, Fortitude, Shadow Protection, Divine Spirit, Arcane Intellect, Mark of the Wild, Kings, Might, Wisdom) and debuffs (Sunder, Faerie Fire, Curse of Recklessness, Curse of the Elements, Demoralizing Shout, Thunder Clap, Judgement of Wisdom, Judgement of Light).
  Each entry has providers (class, spec, talent), a need (`1 per melee group`, `1 per caster group`, `1 per raid`) and a pool (one shaman air totem, one paladin blessing per paladin).
- raids_cover.go computes cover from sign-ups, with `sure` and `likely` providers. The raid page shows it as the buff Grid.
- Raid-wide reach in Forever: the code says Strength of Earth is `melee_group`; the owner says it reaches the whole raid. They disagree and neither is verified. K1 checks in game (and against the client's spell data) and fixes the catalogue before K3, for every totem and aura, not one: Strength of Earth, Windfury, Grace of Air, Mana Spring, Trueshot, Leader of the Pack, Moonkin. No bench example in this brief depends on the answer.
- Missing from the catalogue (raiders count them): Improved Scorch, Expose Armor, Shadow Weaving, Curse of Shadow, Judgement of the Crusader, Mana Tide, Prayer of Fortitude. K1 adds each or the brief names it out of scope in the owner's G2 answer.
- The sim: dps-per-point.json and the engine give a spec's DPS per stat point; simbuffs.json holds consumables only. No raid-buff overlay on a sim run was found. A marginal sim gain per extra player (their own DPS plus the buffs they bring to everyone else) needs engine support for external buffs. That is unverified; the first lane is a spike.

What does not exist: group assignments (who stands in which group), so group-scoped buffs (Windfury, Grace of Air) can only be counted per raid, not per group. Comp targets per raid are a gap (2026-11-05-web-layout.md 7.4).

Shape:
- Surface: the raid page's Raid tab (officer), under the Grid. Rows: each signed-up extra player (beyond 40, or on Bench) with `Brings` (cover only they provide), `Adds` (cover the raid gains by fielding them) and `Cover after swap` as before and after counts (`Buffs 14 of 16, debuffs 7 of 8`), all from raids_cover.go providers (`sure`, `likely`). Windfury and Grace of Air rows say `counted per raid, not per group`. A marginal DPS band `30 to 40 DPS` joins only after K4, one unit (raid DPS), healers coverage only.
- Same rules as the loot plan: bands, not numbers; names only to officers.
- Unknown shown as `No gear on file` for a raider and `Unknown` for a provider whose talent is unknown (the `likely` provider class), never `missing`.
- Discord: one line in the officer post, `Bench call: Brenn sits, no cover lost.`, sent only when an officer presses send; no names in the raid channel.
- App: read only under Officer. Addon: none.

Lanes: K1 verify cover catalogue against client spell data and fix the reach (S, 1 d). K2 spike: external buff overlay in the engine, one buff, one spec, compare to a hand sim (M, 2 d, decides everything after). K3 API plus web: bench and swap rows from coverage alone, no DPS band (M, 3 d; ships without K2). K4 DPS band from K2 (L, 4 to 5 d, only if K2 holds). K5 app read only (S, 1 d).

## 2. Recruiting search over the character directory

The question: "Who is a 60 Fury Warrior who is not in a guild and would fit us?" Answer without scrolling: a filtered list, first row visible: crest, name in class colour, `60 · Fury Warrior`, item level, realm, `Profile 2 d`, with `Copy whisper` as the only action.
Reference: the in-game Who list and the site guild finder (the site's guild finder at /guild; the guild-tab spec's `Guilds recruiting` rows). Borrow: one line per character, class colour, level, no card.

What exists: `character_directory` (migration 0106; public profiles only; equipment summary without enchants; profile re-read at 72 h; rosters 24 h; rows unconfirmed for 30 days are deleted), `directory_optouts` (0108, honoured by removing the character forever), the opt-out flow (api/internal/directory/optout*.go).
Not there: guild membership of a directory character is only known where the roster pass saw it; a character in a guild we have not scanned reads `Guild unknown`, never `No guild`. There is no contact path: a player cannot be messaged from the site. The action is `Copy whisper` (a `/w Name` line on CopyButton), not a site message.

Shape:
- Surface: the guild page's Recruiting tab (officers). Components: ControlBar (class, spec, level, item level, realm; sort by item level or profile age; result count; a saved filter per guild), RankList rows, ShowAll, EmptyState `No characters match.`, LoadError. The spec filter is valid: character_directory.spec is the active talent tree from the specializations endpoint (directory/parse.go `activeSpec`). It is empty when no point is spent, and then the row has no spec.
- Row: item level, the count of rare-or-better pieces from the directory `equipment` summary, and `Fills: Windfury Totem`, computed from the guild's open cover gaps (raids_cover.go) for the character's class and spec.
- Guild membership: show only characters the roster pass has scanned, else every row says `Guild unknown`; the list never claims `No guild`.
- Rules: opted-out characters never appear and are not counted; the list shows only public data already on the directory; every row carries `Profile 2 d` (the profile's age, up to 72 h); no enchants, no gear ranking, no `bad` signal; no export, no bulk copy.
- Rate limit per guild; no scraping of the directory by anyone without a guild plan seat.
- App: none in v1.

Privacy flags, both need the owner before any build:
1. design/legal/2026-11-14-privacy-additions.md section 2 lists three uses of Battle.net data (who is in a guild, guild rosters, Discord sign-ups, plus realm type). Recruiting search over the directory is not one of them. The paragraph must name it before launch, and so must the opt-out page (`Remove your character` says what removal prevents).
2. The loot plan (design/specs/2026-11-15-loot-plan.md section 4) uses directory gear for a guild's loot plan, also not named. The same edit covers it: one new sentence naming `recruiting searches by guilds and loot planning by a guild a raider has signed up to`.
Until both are added, the directory path in the loot plan ships off and this search is not built.

Lanes: R1 notice and opt-out copy, owner sign-off (S, 0.5 d, blocks R2). R2 API: search endpoint over directory with opt-out join, guild-plan gate, rate limit (M, 2 d). R3 web Recruiting tab (M, 3 d). R4 tests: opt-out honoured end to end, no opted-out row ever returned (S, 1 d).

## Open owner decisions
- G1. Add recruiting search and loot planning from directory gear to the privacy notice (blocks R2 and the loot plan's L3 directory path).
- G2. Confirm Strength of Earth's raid-wide reach in Forever and whether the catalogue's per-group needs change.
- G3. Marginal DPS for bench calls: wait for the engine spike (K2), or ship coverage-only first (recommended).
- G4. Recruiting plan gate: all guilds or paid guild plans only.

## Review fixes
Strength of Earth disagreement stated as code against owner, K1 checked in game before K3, bench example replaced. Missing buffs and debuffs listed for K1. Bench rows gain `Adds`, before and after cover counts, a coverage figure, per-raid labels and an officer-send rule. Recruiting: action is `Copy whisper`, spec filter kept (the directory carries spec), item level, rare piece count and `Fills:` added, `Profile 2 d` replaces `Seen`, `Guild unknown` rule, saved filter and sort. One disclaimer rule: the leader line is gone; unknown wording matches the loot plan (`No gear on file`).
