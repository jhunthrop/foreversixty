# Log coach v1: a coach for your spec, reading your own log (product design)

Status: product design for the owner, 2026-10-27, revised the same day after the owner's
direction: "I'd rather have the LLM be a coach with the world's best player persona for each
class and spec." Base: the architect's map in
`docs/superpowers/specs/2026-10-27-log-coach-feasibility.md` (its §6 v1 slice and §7 risks).
Player evidence: `design/reviews/2026-10-27-premium-differentiators/README.md` and
`design/reviews/2026-10-27-launch-pricing/appeal-*.md`. The Premium line this feature unlocks
is held in `design/specs/2026-10-27-launch-pricing.md` §B5 (`premium.featuresHeld.coach`).
Decisions marked **(decided)** are the controller's calls of 2026-10-27 and are not open.

This is the product design. Each player-facing surface still needs its experience spec and
mock from the `ux-designer`, reviewed by `wow-player`, before it is built (tenet 6). Those
specs build on the shapes, copy patterns and rules here. They do not reopen them.

---

## 1. Intent, and why now

After a raid, a player opens their fight and gets a review from **their spec's coach**: a
mentor written as the best player of that spec on Forever. The coach says what to fix first
and why, how the mistakes connect, what they did well, and the one drill to practise. Every
fact in that review comes from an **evidence pack**: rule-based findings measured from the
player's own log against the best logs on the site, the engine's written rotation or the
spell data. A language model writes the words. It never supplies the facts.

Why now:

- **It is what players said they would pay for.** Four of five personas in the
  differentiator round asked for "coach me on my own log." The Fury warrior wanted to hear
  "you clipped 14 Bloodthirsts". The Holy paladin wanted "what to fix after Onyxia". Today's
  Premium is "more of the free tools", and none of them would buy that on its own. The coach
  is Premium's main differentiator.
- **A list of findings tells a player what went wrong. A coach also tells them what to do
  about it.** It sets the priority, gives the reason, links the mistakes together, gives the
  encouragement a mentor gives and names one thing to practise. That is the owner's
  direction, and it is what turns a parse into a lesson.
- **We already own the facts.** The parser stores casts with timestamps, aura and resource
  tracks, deaths and gear. The engine holds a written rotation per spec. The rating already
  turns deaths and dropped debuffs into anchored moments. The missing pieces are the
  comparison table, the evidence pack, the personas, the model client and the surfaces.
- **The calendar.** Forever launches on 4 November and raids open on 9 December. The coach
  should be live for the first raid night, and the plans page's held line goes live with it.

## 2. Goals, non-goals, anti-goals

Goals

- After any logged boss pull, a player can open their own review: a short opening, one thing
  to fix first (what, why, drill), up to three more, one or two things done well, and what
  couldn't be checked. Every item links to its evidence in the report.
- Every fact in the review (number, spell, time, comparison) is in the evidence pack, and
  the API checks that before anything is shown.
- One coach per spec, 28 in all, each honest about what it is: "Your Fury coach".
- Damage dealers get the full set of findings. Healers and tanks get the findings the data
  supports, framed for their role (§6.3).
- Free accounts see the count and the first thing to fix on their own character. Premium
  (`log_coach`) sees the whole review.
- One line in Discord and one on the account page carry "what to fix first" to the player.
- 10-, 20- and 40-player raids each have their own comparison pool.

Non-goals (v1, named again in §22)

- Talents, race and a sim of the player's own character. Forever logs carry neither talents
  nor race.
- Movement, positioning, facing and range.
- Mechanics beyond the curated tables that exist.
- Trends over weeks, and a coach that remembers past nights.
- Trash pulls. Gear and consumables are only in the log on boss pulls.
- A chat. The coach writes one review per fight and does not answer questions in v1.
- Personas of real people.

Anti-goals (a build that does any of these fails review)

- **A finding that is wrong.** If the evidence pack doesn't contain it, the review can't say
  it. Accuracy is absolute (tenet 8).
- **A review that shames.** No ranking of raiders against each other, nothing personal in a
  channel others read, no coaching of another player's character except by their guild's
  officers, never "heal more", never a grade of the person.
- **A wall of numbers, or a wall of prose.** The review has a fixed shape with length caps
  (§9.3). The full numbers stay on the report's own tabs, one click away (tenet 11).
- **A coach pretending to be someone.** No names, no "pro" claims, no invented biography.

## 3. Who sees what (decided)

| Viewer | Report page row | Discord | Account page |
|---|---|---|---|
| Signed out, or not the character's owner | Nothing about the coach | Nothing | n/a |
| Owner of the character, free | The finding count and the title of the first thing to fix, plus an upgrade link | Nothing (the Discord line is for Premium only) | Last raid line: count and the first thing to fix |
| Owner, Premium (`log_coach`) | The full coach review, and coach history | `/report` answer: one "what to fix first" line, ephemeral. Raid-night post: one sentence that names nobody | Last raid line: what to fix first, with a link |
| Officer of the report's guild with `officer_views` for that guild, viewing a verified member's row | The full review, marked "Officer view" | Nothing extra | n/a |
| Moderator | As an officer, on any report they can open | Nothing | n/a |

"Owner of the character" means `auth.Store.OwnsCharacter(user, player_key)`, the test the
Premium history route already uses (`api/internal/rating/handler.go:258`). `log_coach` is a
standard personal Premium feature (§17.4). A guild seat holder has it.

**Generation cost only follows Premium (decided).** Free teasers, the Discord line and the
account line are rendered from the evidence pack alone. The model runs only when a Premium
owner or an officer opens a review (§15).

## 4. How it works

```
fight summary (bucket) ─┐
coach_profiles pool ────┤                         ┌─ persona file data/curated/coach/<spec>.md
sim reference ──────────┼─► Evidence pack ────────┼─► model (OpenAI-compatible) ─► review JSON
spell data, catalogue ──┤   F1–F8 findings,        │                                   │
curated mechanics ──────┘   strengths, links,      └─ coach rules (system prompt)      ▼
                            couldn't-check, anchors                       validate against pack
                                    │                                       │ pass   │ fail
                                    │                                       ▼        ▼ retry once
                                    └──────────── rule-only rendering ◄──── fail again, cap hit,
                                                                            kill switch, error
```

1. **The evidence pack** (§7) is built by pure Go code. It holds the eight rule findings
   (§6), the basis and comparison for each, the strengths (what the player did at or above
   the comparison), the computed links between findings, the "couldn't check" list and the
   evidence anchors. It is the single source of truth for the review.
2. **The model** gets the spec's persona and the coach rules as the system prompt, and the
   evidence pack as the only user content. It returns the fixed review shape (§9.3) as JSON.
3. **The validator** (§10) checks every number, spell, boss, time, anchor id, link id and
   drill id in the output against the pack. If anything fails, the model gets one retry with
   the failures listed. If that fails too, the player gets the **rule-only rendering**: the
   same review shape, filled from the pack's own copy table with no model at all.
4. The result is stored per (report, fight, player, model version, persona version) and
   generated again only when a version moves (§14).

The rule-only rendering always exists. If the model is down, over a cap or switched off, the
player still gets a correct review in plain words.

## 5. The player's experience

### 5.1 The review, on screen

The web renders a fixed structure. It never renders free prose.

- **Header:** `Your {Spec} coach`, e.g. `Your Fury coach`. For a starter persona (§8.3),
  `Your Enhancement coach · early version`, with the hint `This coach reads your log against
  our guide's Enhancement rotation.` On the rule-only rendering the header is
  `Your Fury review`, with no persona voice claimed.
- **Opening:** at most two sentences.
- **Fix first:** a title, why it matters, and the one drill to practise.
- **Then:** up to three more items, each with a title and a why.
- **Doing well:** one or two items.
- **Couldn't check:** the pack's lines, word for word.
- **Footer:** `Reviewed {date} · Based on: {labels}.`, e.g. `Reviewed 10 Dec · Based on: Top 25
  Fury warriors on Onyxia (40), your log.` The basis labels the pack used, deduplicated.

Every item carries the evidence anchor ids it used. The web turns each one into a link into
the report: the Casts row, the aura row, the death, or the resource graph at a time. They are
built the way `web/src/lib/rating/moments.ts` builds the rating's moment links today.
Abilities show icon and rank, and items show icon, quality colour and tooltip (tenets 2 and 3).
Fight times are `m:ss` from the pull.

States: loading (`Reading the pull…`), the rule-only rendering shown while the coach writes
(`Your coach is writing…`, then swapped in place without layout shift), a reference still
preparing, coach busy (§15.4), no findings, not reviewable, the free teaser, officer view, and
error (`The coach couldn't read this pull just now.`). The experience spec
(`design/specs/2026-10-28-log-coach.md` §8) holds every surface's final words.

### 5.2 The evidence pack's finding pattern (also the rule-only rendering)

Each pack finding renders as four parts from one copy table (`api/internal/coach/copy.go`):

1. **Title:** what happened, under eight words, at most one number: `Bloodthirst: 9 casts lost`.
2. **Detail:** one sentence of fact with the proving numbers and the worst moments.
3. **Basis:** what it is measured against, in plain words, with the comparison.
4. **Evidence anchors.**

The rule-only rendering maps the pack onto the review shape mechanically. The opening is
`Here is what your log shows for {Boss}.` Fix first is pack rank 1, with its detail as the
why, and the drill is the persona file's drill for that finding code. Then is ranks 2 to 4.
Doing well is the top two strengths. Couldn't check is copied as is.

When there are no findings, the opening becomes `Nothing to fix on this pull that the log can
show.` Doing well and couldn't check follow.

### 5.3 Worked examples

**All values below are invented examples for this document. None of them is a measurement.**
Each example shows the evidence pack first, then a review the model could write from it.
Every number, spell and time in the review is in the pack above it.

**Fury warrior, Onyxia (40), kill, 3:52**

Pack findings, ranked:

| Rank | Id | Title | Detail | Basis | Anchors |
|---|---|---|---|---|---|
| 1 (major) | `f1` | `Bloodthirst: 9 casts lost` | `9 of a possible 40 went unused while you had the rage. Longest gaps 1:12–1:25 and 2:40–2:51.` | `Measured against Bloodthirst's 6-second cooldown and your own rage bar. Top 25 Fury warriors on Onyxia (40): 9.1 a minute. You: 8.0.` | `a1` Casts row, `a2` 1:12, `a3` 2:40 |
| 2 (major) | `f6` | `No flask at the pull` | `You pulled with Grilled Squid and an Elemental Sharpening Stone, and no flask.` | `From the character sheet in your log, against the site's consumable list for damage dealers.` | `a4` Buffs tab |
| 3 (minor) | `f4` | `Rage full for 14 seconds` | `Your rage sat at its cap for 14 s outside the air phase, longest at 2:41–2:50.` | `Measured against your own rage bar, air phase left out. Top 25 Fury warriors here: 3 s.` | `a5` rage graph at 2:41 |

Strengths: `s1` `Alive the whole fight.` `s2` `Battle Shout up 100% of the fight. Top 25: 97%.`
Links: `l1` (`f1`, `f4`): `The Bloodthirst gap at 2:40–2:51 overlaps the rage cap at 2:41–2:50.`
Couldn't check: `Talents and race: Forever logs don't record them.` · `Range and movement: the
log doesn't say where you stood.`

Review:
- **Opening:** `Solid kill: you were alive the whole fight and Battle Shout was up 100% of it.
  Almost everything you left behind sits in one button.`
- **Fix first, `Bloodthirst: 9 casts lost`:** `Bloodthirst is the button the Fury rotation is
  built around, and 9 of a possible 40 went unused while you had the rage, longest at
  1:12–1:25. At 2:40–2:51 you sat on a full bar instead, so that one gap cost you twice.`
  Drill: `On a target dummy, press Bloodthirst the moment it comes off cooldown, before
  anything else; Whirlwind and Heroic Strike wait their turn.` Anchors `a1 a2 a3 a5`, link `l1`.
- **Then:** `No flask at the pull`: `A flask is the cheapest damage you will ever buy; bring
  one to every 40-player boss.` (`a4`) · `Rage full for 14 seconds`: `When the bar is full,
  Heroic Strike turns rage into damage instead of losing it.` (`a5`)
- **Doing well:** `s1`, `s2`, each in one sentence.

**Holy priest, Onyxia (40), kill, 3:52** (efficiency and survival, never "heal more")

| Rank | Id | Title | Detail | Basis | Anchors |
|---|---|---|---|---|---|
| 1 (minor) | `f7` | `Died at 3:20 to Flame Breath` | `Flame Breath hit for 2,140 at 3:20. No healing potion or Healthstone used in the 10 s before.` | `From your log. Flame Breath is on this boss's list of damage you can avoid.` | `a1` the death |
| 2 (minor) | `f4` | `Out of mana from 3:05` | `Your mana stayed under 5% from 3:05 until you died at 3:20, 15 s in all.` | `Measured against your own mana bar.` | `a2` mana graph at 3:05 |
| 3 (minor) | `f5` | `Greater Heal (Rank 4): 44% overheal` | `44% of what Greater Heal (Rank 4) healed was overheal, over 23 casts. Lower ranks cost less mana.` | `Top 21 Holy priests on Onyxia (40): 22% on the same rank.` | `a3` Healing tab row |

Strengths: `s1` `Flask of Distilled Wisdom at the pull.` `s2` `Inner Fire up 100% of the fight.`
Links: `l1` (`f4`, `f7`): `You died while out of mana.` `l2` (`f5`, `f4`): `Greater Heal
(Rank 4) is your most-cast heal with the most overheal.`

Review:
- **Opening:** `Your preparation was right: flask at the pull and Inner Fire up 100% of the
  fight. This one is about lasting to the end.`
- **Fix first, `Died at 3:20 to Flame Breath`:** `A dead healer heals nothing for the rest of
  the pull, and Flame Breath is damage you can step out of. You were also out of mana when it
  landed, so there was no answer left.` Drill: `Before the pull, pick your spot at her side,
  never her front, and check it each time she lands.` Anchors `a1 a2`, link `l1`.
- **Then:** `Out of mana from 3:05`: `The mana ran out 15 s before the death; the fix is upstream,
  in how you spend it.` (`a2`) · `Greater Heal (Rank 4): 44% overheal`: `Top Holy priests here
  overheal 22% with this rank. A lower rank costs less mana, and that is where your last 15 s
  went.` (`a3`, link `l2`)
- **Doing well:** `s1`, `s2`.

**Protection warrior, Onyxia (40), kill, 3:52**

| Rank | Id | Title | Detail | Basis | Anchors |
|---|---|---|---|---|---|
| 1 (major) | `f7` | `Died at 2:50 with Shield Wall ready` | `Onyxia's melee killed you at 2:50. Shield Wall had been ready since the pull.` | `From your log and Shield Wall's cooldown.` | `a1` the death, `a2` Shield Wall row |
| 2 (minor) | `f2` | `Shield Block up 18% of the fight` | `Shield Block was up for 42 s of 3:52.` | `Top 20 Protection warriors on Onyxia (40): 34%.` | `a3` Buffs row |
| 3 (minor) | `f5` | `Sunder Armor: 22 casts of Rank 4` | `Rank 5 is trainable at your level. Rank 4 was cast 22 times.` | `The spell ranks a level 60 warrior can train.` | `a4` Casts row |

Strengths: `s1` `Sunder Armor at 5 stacks from 0:21 to the end.` Links: none.

Review:
- **Opening:** `You had Sunder Armor at 5 stacks from 0:21 to the end, so the raid had what it
  needed from you early. The pull turned on one button you never pressed.`
- **Fix first, `Died at 2:50 with Shield Wall ready`:** `Shield Wall had been ready since the
  pull, and a tank who dies takes the raid with them. Pressing it is never wasted.` Drill:
  `Put Shield Wall on a key you can hit without looking, and say "wall" out loud when your
  health drops past half on a dummy pull with a friend.` (`a1 a2`)
- **Then:** `Shield Block up 18% of the fight`: `Top Protection warriors here keep it up 34%;
  it is your cheapest mitigation.` (`a3`) · `Sunder Armor: 22 casts of Rank 4`: `Rank 5 is
  trainable at your level; visit your trainer.` (`a4`)
- **Doing well:** `s1`.

### 5.4 Surface 1 (primary): the report page player row

- The player's row in the report's combatant list carries one coach line, only for viewers
  who may coach that row (§3). Premium: `Fix first · Bloodthirst: 9 casts lost`. Free:
  `3 things to fix · First: Bloodthirst: 9 casts lost · Read the full review with Premium`.
- The line opens the **Coach** panel for that player and fight (§5.1). Premium also gets a
  **History** link to the character's past reviews. The panel's placement (a tab beside
  Rating, or a drawer from the row) is the ux-designer's call. It belongs on the report page,
  which owns the log (one job per page).

### 5.5 Surface 2: Discord (decided: Premium only, never shaming in a channel)

- **`/report` answer (ephemeral).** When the viewer holds `log_coach` and owns a character in
  the matched log, the answer gains one line per character (three at most): `Fix first:
  Bloodthirst casts lost — on 5 of 8 bosses. Open your coach's review`, the last words linking
  to the coach panel. One boss in the night reads `Fix first: {stem}.`, every boss `on all {m}
  bosses`. `Share in channel` re-renders without it.
- **Raid-night post (`report_ready`, in the guild channel).** That post is counts only and
  never names a raider (rule R-A3, `api/internal/notify/report_embed.go`). The coach keeps the
  rule. When at least one raider in the log holds `log_coach` on a character they own, the
  description gains one sentence that names and counts nobody: `Coach reviews for tonight's
  pulls are ready on Forever Sixty and in /report.` The coach writes only when a review is
  opened, so the sentence says the reviews are ready, never that they were written.
- **Night line rule.** Read from the evidence packs, with no model call. Take the rank-1
  findings across one character's fights that night. Choose by severity, then by the finding
  type that showed on the most bosses, then the earliest boss. It is worded from the title stem
  without the per-fight number, plus `on N of M bosses`.
- **Timing.** Packs come from cache. If they aren't all cached within the interaction's budget,
  the line becomes `[Your coach's review is on the site.]({url})`, and the packs warm in the
  background. Discord never waits.

### 5.6 Surface 3: the account page's "last raid" line

- One line, for the account's most recent raid. "Last raid" has the definition the addon
  already uses (`addon.lastRaidFights`: newest evidence report with a rated boss fight, within
  `addon.LastRaidWindow`, 14 days). It reuses that query.
- Premium: `Last raid · Molten Core, Thu 23 Oct · Fix first: Bloodthirst casts lost, on 5 of 8
  bosses · Open review`. Free: `Last raid · Molten Core, Thu 23 Oct · 11 things to fix · First:
  Bloodthirst casts lost · Read the full review with Premium`.
- With no raid in 14 days it reads `Last raid · no log in the last 14 days`, muted: shown,
  never hidden, so its reserved slot never collapses. It loads separately (`GET /v1/me/coach/recent`)
  so a cold cache never slows the page. It likely sits in
  `web/src/components/account/CharacterList.svelte`. The ux-designer places it.

## 6. The eight findings (the evidence pack's measurements)

These are unchanged from the rule design. They are now the facts the coach is allowed to use.

### 6.1 Shared rules

- **Reviewable fights:** boss pulls only (`encounter_id` set), kills and wipes, at least 30 s
  long. On a wipe, deaths in the last 15 s before the end are part of the wipe and are not
  reviewed.
- **Spec and role** come from `RosterRow`. A row with no known spec is not reviewed: `We
  couldn't tell this player's spec from the log.`
- **Level** comes from the advanced-logging fields (`AdvLevel`). Without it, F5's rank check
  and the sim band go to "Couldn't check".
- **Never counted against a player:** time dead, the encounter's curated downtime (the windows
  `rating.ActivityIdleMS` subtracts), and the time before their first cast of the fight.
- **Talent abilities** are assessed only if cast at least once in the same fight. Forever logs
  carry no talents, and one cast proves the talent.
- **Pets and guardians** fold into their owner (`CastRow.OwnerGUID`, `Ability.Via`).
- **Per-spec rule data:** `data/curated/coach/<spec>.json` lists core cooldown abilities, owned
  auras, deliberate downranks, defensive cooldowns and efficiency cooldowns. Every entry
  carries a `verified` stamp against the client's tables, and a test refuses an unstamped
  entry. Missing lists send their findings to "Couldn't check".
- **Severity** is `major` or `minor`. **Rank** is major before minor, then the role's
  precedence (§6.3). Raid buffs are never rank 1, because the player doesn't control them.
- **One finding per type** per fight, eight at most. F1 and F2 name the worst ability in the
  title and up to two more in the detail.
- **Gain estimates** stay null and are never shown.

### 6.2 The findings

Thresholds are named constants in `api/internal/coach/rules.go`, with the v1 values below.

**F1. Core ability under-used** (`core_casts`). For each core ability with a cooldown of 4 s
to 10 min, walk `CastRow.Sequence`. A gap counts as ready time when the ability was off
cooldown, the player was alive and outside downtime, and their resource was at or above the
spell's base `cost` (`spellconst`) for at least one whole second of `ResourceTrack.Series`.
Talents only lower costs, so base cost never invents a lost cast. Each gap gives up a 1.5 s
grace. Lost = floor(remaining ready time / cooldown). Possible = actual + lost. *Fires* at
lost ≥ 2 and ≥ 10% of possible. *Major* at ≥ 20%. The basis is always the cooldown and the
player's own bar. The comparison follows the basis order (§11), skipping the sim for abilities
on the spec's `WorstActions` list (`api/internal/sims/specs.go`). If cost is unknown, the gaps
count without the resource check, and the detail says `Couldn't check whether you had the
rage.` *Roles:* DPS damage cooldowns; tank threat cooldowns; healer **efficiency cooldowns
only** (Inner Focus, Mana Tide Totem, Nature's Swiftness), never framed as missing healing.

**F2. Key buff or debuff uptime** (`uptime`). Owned auras, uptime from first application to
the end, excluding time dead and downtime. A debuff is judged only when the player is its
**sole applier** (`AuraTrack.Appliers`). Otherwise it goes to "Couldn't check": `Your share of
Sunder Armor: 3 warriors applied it, and the log merges them.` *Fires* ≥ 10 points under the
reference and under 90%. *Major* at ≥ 25 points under. *Basis:* top logs, else the sim, else
`not enough logs yet`. *Roles:* DPS own DoTs, debuffs, self-buffs; tank defensive uptime
(Shield Block, Demoralizing Shout/Roar, Thunder Clap), top logs only; healer own self-buffs.

**F3. Dead time** (`idle`). `rating.ActivityIdleMS`, plus its three longest windows (a windows
variant of the same code in `logs/engine/rating/activity.go`). *Fires* at ≥ 10 s and ≥ 5% of
counted time. *Major* at ≥ 10%. *Basis:* top logs' median active %, else the threshold.
*Roles:* DPS and tank. **Off for healers.**

**F4. Resource waste** (`resource`). Rage and energy: seconds at cap outside downtime and time
dead. Mana: seconds under 5% of `Max`, and the first such second. *Fires:* capped ≥ 10 s and
≥ 5%; mana under 5% for ≥ 10 s, starting before the last 10% of the fight. *Major* at ≥ 15%
capped, or out of mana before 75% of the fight. *Roles:* all. Healer copy is about running out
early, never about leftover mana.

**F5. Ranks** (`ranks`). *DPS and tank:* casts below the highest rank trainable at the
player's level, for rotation spells (`Rotation.Family`, `spellranks.json`, `trainables/`),
excluding listed deliberate downranks. Fires at ≥ 3 casts or ≥ 10%. Always minor. *Healer
(rank choice):* overheal % per heal rank (`Ability.Overheal / Ability.Total`; the spell id is
the rank). Fires when a rank with ≥ 10 casts is ≥ 15 points over the top logs' median for the
same rank. Major at ≥ 30. Top logs only. The detail may add `Lower ranks cost less mana.` and
nothing more.

**F6. Pull preparation** (`preparation`). From the seeded aura tracks (segments on the player
starting at 0 ms) against the role's consumables catalogue, the shared reader of §13. Covers
flask, food, weapon oil or stone, combat potion uses against `MaxUses`, missing permanent
enchants (`score_v2.enchanted`), and raid buffs. **Completeness guard:** a "missing" finding
only fires for a catalogue category marked `complete`. Otherwise: `Food: the coach's food list
isn't complete yet.` Weapon oils and stones stay unchecked until their COMBATANT_INFO form is
verified. **Raid buffs (decided):** flagged only when a class that provides the buff was in
that fight's roster. They are minor and never rank 1. The catalogue gains `raid_buffs` with
the providing class and every rank's id. **World buffs** are listed as facts, never a finding.
*Severity:* flask, food or weapon enchant missing is major in 40-player raids and minor in 10
and 20. A missing weapon enchant is major in every size. Potions and raid buffs are minor.
*Disabled* without a character sheet: `Consumables and enchants: your log has no character
sheet for this pull.`

**F7. Death review** (`death`). The first reviewable death: time, killing blow, whether it is on
the encounter's avoidable list, and the coach file's defensive cooldowns that were off cooldown
then. Baseline abilities count at their trainable level. Talent abilities count only if cast
that fight. Also whether a healing potion or Healthstone was used in the 10 s before (a fact
about casts, never about bags). It reuses the rating's death `Moments`. *Fires* before the last
10% of a kill. *Major* before 75%.

**F8. Rotation adherence** (`adherence`). `rating.fightAdherence` plus the three most-cast
off-rotation spells. *Fires* ≥ 10 points under the top logs' median, or under 85% against the
written rotation alone. *Major* at ≥ 25 points under. *Roles:* DPS and tank. Off for healers.

### 6.3 Which specs get which findings (v1)

DPS (20): Balance and Feral (cat) druid; Beast Mastery, Marksmanship and Survival hunter
(Survival is melee in Forever); Arcane, Fire and Frost mage; Retribution paladin; Shadow
priest; Assassination, Combat and Subtlety rogue; Elemental and Enhancement shaman;
Affliction, Demonology and Destruction warlock; Arms and Fury warrior.
Healers (5): Restoration druid, Holy paladin, Discipline priest, Holy priest, Restoration shaman.
Tanks (3): Feral (bear) druid, Protection paladin, Protection warrior.

28 specs in all (20 + 5 + 3), matching `data/curated/specs.json`. Corrected 2026-10-27 from specs.json: this section and §2, §8.1 and §22 read 19, 27 and 24 before.

| Finding | DPS | Healer | Tank |
|---|---|---|---|
| F1 core casts | damage cooldowns | efficiency cooldowns only | threat cooldowns |
| F2 uptime | own DoTs, debuffs, self-buffs | own self-buffs | defensive uptime (top logs only) |
| F3 dead time | yes | **off** | yes |
| F4 resource | rage, energy, mana | mana (running out early) | rage, mana |
| F5 ranks | below trainable rank | overheal by rank (top logs only) | below trainable rank |
| F6 preparation | yes | yes | yes |
| F7 death | yes | yes, survival first | yes, survival first |
| F8 adherence | yes | **off** | yes |

Precedence within a tier: DPS F7, F1, F3, F6, F2, F4, F8, F5. Healer F7, F4, F5, F1, F6, F2.
Tank F7, F2, F1, F3, F6, F4, F8, F5.

## 7. The evidence pack

A JSON document built by `coach.Pack(fight, row, refs, tables)`, which is pure Go. It is the
model's only input about the fight, and the validator's only allowed source.

- `context`: spec label, class, role, boss name, raid size, kill or wipe, duration (`m:ss`),
  level. **No character name, account, guild or realm.** The model addresses the player as "you".
- `findings[]`, ranked: `id` (`f1`…), `code`, `severity`, `rank`, `title`, `detail`, `basis`
  (label, kind, n), `comparison` (label, kind, n) when present, `anchors[]` (ids), and `facts`.
  `facts` lists every number in the finding with its unit and meaning (`{"v": 31, "unit":
  "casts", "of": "Bloodthirst"}`), every spell (id, name, rank), and every time (`m:ss`).
- `strengths[]`: measurements at or above their comparison, computed by the same rules. They
  include uptime ≥ the reference median, no death, every checkable consumable present, adherence
  ≥ median, casts per minute ≥ median, and overheal ≤ median. Each has `id` (`s1`…), a sentence,
  facts and anchors. `doingWell` may only use these.
- `links[]`: deterministic connections between findings, the only connections the review may
  state. A link has `id` (`l1`…), its two finding ids, a sentence and a kind. v1 kinds: a cast
  gap that overlaps a resource cap; out of mana before or at a death; a death inside a cast gap
  (the gap is the death); a rank's overheal on the spec's most-cast heal; a missing consumable
  on a finding about the resource or survival it covers (flask with mana, healing potion with
  death). A link exists only when its time windows or ids actually meet in this fight.
- `couldNotCheck[]`: the lines, verbatim.
- `anchors[]`: `id` (`a1`…) mapped to the report-page target (kind, spell id, time, death anchor).
- `allowed`: a precomputed set for the validator. Every number as it may appear (`31`, `9.1`,
  `44%`, `2,140`, `3:05`, `1:12–1:25`), every spell, item and boss name, and every id.

Packs are cached by (report, fight, player, pack model version) (§14). The free teaser, the
night line and the rule-only rendering read only the pack.

## 8. Personas (decided)

### 8.1 What a persona is

One file per spec, `data/curated/coach/<spec>.md`, 28 files. Each is versioned prose with front
matter: `spec`, `status` (`full` or `starter`), `version`, `sources` (the guide, APL and data
files it was written from) and `reviewed` (who checked it, when). The persona is **unnamed**:
"the best Fury warrior on Forever". It is never a real person, and it never claims a biography,
a guild or a ranking.

Sections, in this order:

1. **Who they are.** One paragraph: what playing this spec at the top looks like on Forever.
2. **What they know, for Forever's rules.** The rotation's priorities in words, from the
   site's APL notes and guide. Cooldown windows and how they line up. The mistakes they see
   most, keyed to finding codes. The consumables and world buffs that matter for the spec. How
   a 10-player night differs from a 40 (fewer buffs, different expectations). Forever-specific
   rules (Survival hunters fight in melee; shamans never dual wield; armour at 40).
3. **Drills.** A short catalogue keyed by finding code (`drill.core_casts`, `drill.death`, …).
   Each is one or two sentences of something to practise. Drill ids are what the review cites.
4. **Voice.** Direct, specific, a mentor who has done it. Short sentences. Never a lecture,
   never "heal more", never shaming, never a grade of the person, no exclamation marks. Praise
   only what the pack proves.

### 8.2 Rules for persona content

- Persona knowledge comes from the site's own guides, APLs (`data/curated/apl/<spec>.json`
  and their per-line notes) and the engine's data. No outside claims.
- No numbers that change nightly (no scale factors, DPS or counts). Those are the owner's rule
  for prose, and the validator would reject them anyway (§10).
- Drills may contain small fixed numbers. Any number in the chosen drill's text is allowed in
  the output.
- Each file is reviewed like a guide: `wow-player` reads it as a raider of that spec, and the
  facts are checked against the APL and guide. The review is recorded in `reviewed`.

### 8.3 Full and starter (decided)

- **Full at launch (3):** Fury warrior, Holy priest, Protection warrior. These are written by
  the implementer from the guides, APLs and engine data, then reviewed.
- **Starter at launch (25):** every other spec. A generator (`tools/coach/persona_starter`)
  writes a starter from that spec's guide and APL only. It covers the rotation priorities,
  the consumables list, the default drills per finding code and the shared voice section. A
  person reads each starter before it ships. A later lane deepens them to full.
- The player sees which one they have (§5.1): `early version` on starters. No starter claims
  more than its sources.
- `persona_version` is the file's content hash. Editing a persona regenerates its reviews on
  next view (§14).

## 9. The review

### 9.1 The model (decided: DeepSeek V4.1 Flash)

The model is reached through the **Vercel AI Gateway** (decided): one key, any model behind it,
billed at the gateway's own per-model prices. `LLM_BASE_URL` is the gateway's OpenAI-compatible
endpoint, `https://ai-gateway.vercel.sh/v1`. `LLM_MODEL` is a gateway model slug in
`provider/model` form. The key is already in Secret Manager as `LLM_API_KEY`, readable by
api-runtime.

**The chosen model (owner, 2026-10-27): DeepSeek V4.1 Flash.** `LLM_MODEL` is its slug on the
Vercel AI Gateway. The builder confirms the exact slug from the gateway's model list, for
example `deepseek/deepseek-v4.1-flash`. On the owner's board that day it scored Reasoning 89.1
and Data Analysis 79.5, the best Data Analysis score on the board, at $0.181 per successful
task. Kimi K3, the Reasoning leader, scored 90.7 and 78.7 at $2.189. The $0.181 is the board's
per-task cost. Our cost comes from the gateway's per-token prices for the slug (§15.3).

**The rule for picking again later** (a new board, a price change, or an evaluation that goes
against the model) uses the columns of the owner's benchmark board:

1. Start from the board's current S-tier. Candidates seen on 2026-10-27: Kimi K3, GLM-5.3,
   DeepSeek-V4.1-Flash, MiMo-V2.6-Pro, GLM-5.2, DeepSeek-V4-Pro, Kimi K2.6. Keep only models the
   Vercel AI Gateway lists.
2. Rank them by the **Reasoning** column.
3. Drop any that are weak on **Data Analysis**. The coach reads a structured evidence pack and
   must not mix up numbers. "Weak" means below the median Data Analysis score of the S-tier
   candidates on the board that day.
4. Ignore the Coding, Agentic, Browser Use and Design columns and the Overall score.
5. Of the models left, take the **cheapest by the gateway's output-token price**. A tie goes to
   the higher Reasoning score.

For DeepSeek V4.1 Flash, the builder records the confirmed slug, the scores above with the
board's date, and the gateway's `{in}`/`{out}` prices in the plan's build notes, then runs the
evaluation in §21 before launch. A later re-pick records the same facts. The code names no model, so a switch is a config change.

### 9.2 The prompt

- **System:** the coach rules (shared by every spec: the output schema, the evidence-only rule,
  the tone rules, the length caps, and "fixFirst is the pack's rank-1 finding"), then the spec's
  persona file.
- **User:** the evidence pack, as JSON.
- **Output:** JSON matching the schema below. The client asks for JSON mode
  (`response_format`), and uses the JSON-schema form when the provider supports it. The server
  validates the shape either way.
- Temperature low (configurable, default 0.3). One completion per call.

### 9.3 The output shape (decided)

```json
{
  "opening": "≤ 2 sentences, ≤ 280 characters",
  "fixFirst": { "finding": "f1", "title": "≤ 60 chars", "why": "≤ 320 chars", "drill": "drill.core_casts",
                "drillText": "≤ 240 chars", "anchors": ["a1","a2"], "links": ["l1"] },
  "thenFix": [ { "finding": "f6", "title": "≤ 60 chars", "why": "≤ 200 chars", "anchors": ["a4"], "links": [] } ],
  "doingWell": [ { "strength": "s1", "text": "≤ 160 chars", "anchors": [] } ],
  "couldntCheck": ["verbatim pack lines"]
}
```

Rules the validator enforces:
- `fixFirst.finding` is the pack's rank-1 finding. The model explains the priority. It doesn't
  choose it. That keeps the free teaser, the Discord line and the account line consistent with
  the full review at no model cost.
- `thenFix` holds at most three items, distinct pack findings other than rank 1, in any order
  the model argues for.
- `doingWell` holds one or two items when the pack has strengths, none when it has none. Each
  cites a strength id.
- `couldntCheck` equals the pack's list, word for word, in order.
- `title` for a finding item equals the pack title, so a player sees the same words on every
  surface. The `why` is the model's own.
- A sentence that relates two findings must cite a link id that joins them.
- `drill` is an id from the persona's drill list. `drillText` may reword the drill but adds no
  facts.
- Tone: no exclamation marks; none of a short banned-word list (`fail`, `bad`, `terrible`,
  `noob`, `heal more`, `pump`, `lazy`, …, in `api/internal/coach/tone.go`); no second player
  mentioned; no claim about a real person.

What the model adds, and the pack alone can't: why the first thing comes first, how the
findings connect (through the pack's links), the plain-language reason a fix matters, the
encouragement a mentor gives, and the one drill to practise.

## 10. Validation and fallback (decided)

1. **Shape:** JSON parses; required fields; caps; ids exist in the pack; the rank-1 rule;
   `couldntCheck` verbatim; titles equal to the pack's.
2. **Numbers:** every digit sequence in the prose (integers, decimals, `%`, thousands
   separators, `m:ss`, `m:ss–m:ss`) is extracted, normalised and looked up in the pack's
   `allowed` set, plus the chosen drill's text. Number words above "three" are rejected, so a
   number can't slip past as a word.
3. **Names:** every spell, item and boss name from the build's name dictionaries found in the
   prose must be in the pack's `allowed` names. The class and spec names of the reviewed player
   are allowed.
4. **Times:** every time must be in the pack. A range must match a pack range.
5. **Links and tone** as in §9.3.

On failure, the call is retried **once**, with the failures listed back to the model (`These
were not in the evidence pack: "12 casts", "Whirlwind". Use only the pack.`). On a second
failure, the player gets the rule-only rendering, which is stored with `source: "rules"` and
the reason. A failure is logged with the pack and both outputs for review, and never shown.

## 11. Where the comparison comes from (decided basis order)

For each comparison, the first that exists wins:

1. **Top logs on the site**: same spec, encounter, difficulty and size, with at least 20 kills.
   `Top {n} {Spec} {class} on {Boss} ({size}): …`.
2. **The engine's written rotation** (the sim reference, §12). DPS specs whose sim is
   `validated` only (`sims.Store.Validated`). `The written {Spec} rotation, simulated at level
   {band}: …`.
3. **Spell data**: cooldowns, costs and ranks. `Measured against {Spell}'s {n}-second cooldown…`.

Uncompared facts carry `From your log`. If a comparison finding has no basis at its level, it
isn't a finding. "Couldn't check" says `not enough logs yet — {n} {Spec} kills of {Boss}
({size}) on the site, the coach needs 20.`

### 11.1 `coach_profiles`, the pool

- Partitioned by range on `fought_at` like `fight_metrics` and `rating_scores`. Added to the
  partition job's tables (`api/internal/db/partitions_job.go`).
- Columns: `report_id, fight_index, player_key, class, spec, role, encounter_id, difficulty,
  size, kill, duration_ms, phase, metric, fought_at, model_version, profile jsonb`. Primary
  key `(report_id, fight_index, player_key, fought_at)`.
- `metric` is what "top" means: DPS for damage dealers, and the rating's overall score for
  healers and tanks.
- `phase` is the content phase of `fought_at` (`data/curated/phases.json`).
- `profile` holds the same measurements the pack uses: alive and counted ms, active %, casts per
  minute per family, uptime per owned aura (sole applier only), capped %, mana-under-5% ms and
  first time, overheal per heal rank, defensive uptime, adherence, and consumables at the pull.
  About 1 KB.
- Written by the rater in the rating transaction, for kills on reports that are ranked
  (`reports.Ranked`) and evidence (`presence.EvidenceReport`). Removed on `Ranker.RemoveReport`.
- Indexes: `(encounter_id, difficulty, size, spec, phase, kill, metric desc)` and
  `(player_key, fought_at desc)`.
- **The comparison:** one best kill per player in the same (encounter, difficulty, size, spec,
  phase), with the reviewed player left out. At least 20 distinct players. Compared with the
  **median** of the top 25 by `metric`. Nobody in the pool is ever named or linked.
  Medians are cached in memory for 1 hour.

## 12. The sim reference cache

The spec's APL run by the engine at the player's level band, with the band's preset gear and
the same preset and race the BiS ladder uses, at the fight length rounded to 30 s. The result's
`Summary` gives casts per minute and uptimes. It is used for validated DPS specs only.
`coach_sim_refs` is keyed `(engine_version, build, spec, band, duration_bucket)`. A miss queues
a job (about 3 s of compute), and the pack answers `preparing` until the reference is in. A
nightly warm is optional. The basis line says "the written rotation", never "the best possible".

## 13. Fix first: consumables at the pull (decided)

`CombatantRow.Consumables` is never filled in production (feasibility §1). So
`rating.consumablesMissed`, `fightDiscipline` / `carries` and the engine's `scorePreparation`
are wrong today. This lane fixes it before any coach code:

1. One engine function, e.g. `consumables.AtPull(fight, guid, catalogue)`, reads the player's
   seeded aura tracks against the catalogue. It needs no companion release.
2. The three readers and the coach's F6 all use it.
3. Raid buffs use the same path. `MissingBuffs` is not used anywhere new.
4. Bump `ratingengine.DefaultModelVersion`. The existing backfill re-rates stored fights, and
   the preparation and discipline digests are rebuilt.
5. Known-good fixtures pin a real Forever pull with and without a flask.

Ruling 2026-10-27: "at the pull" means the player's own COMBATANT_INFO listed the aura (the
summary marks the track `seeded`, whatever millisecond the snapshot lands on), or the aura was
applied within the first 10 s of the pull (`consumables.AtPullGraceMS`) for flasks, food and raid
buffs. 10 s grace; potions are F6's use rule.

## 14. Computing, caching and versions

- **Pack:** the cache row `coach_packs (report_id, fight_index, player_key, pack_version)`.
  `pack_version` = `pack-v1.` + a hash of the rules version constant, the `<spec>.json` rule
  files and the catalogue.
- **Review:** the row `coach_reviews (report_id, fight_index, player_key, pack_version,
  model_version, persona_version)`, holding the output JSON, `source` (`model` or `rules`),
  `tokens_in`, `tokens_out`, `cost_usd` (as the gateway reported it), `cost_estimated` (bool),
  the model slug, the gateway request id and `generated_at`. `model_version` = `LLM_MODEL` + a hash of the coach
  rules prompt. **One generation per key (decided)**. It is regenerated only when a version
  moves, on the next view by a viewer entitled to the full review.
- **A review stored as `rules` because a cap was hit or the switch was off** is generated
  properly on the next view once the cap or switch allows. One stored as `rules` after two
  failed validations is not retried until a version moves.
- Rows are deleted with their report (a foreign key with cascade).
- **Flow:** check access (§19), then load or build the pack (bucket read of the summary through
  `sims.fightSummary`, cached medians, sim reference). Free owners get the teaser from the pack
  here. Premium or officer: load the review. On a miss, return the rule-only rendering with
  `status: "writing"`, queue the generation on the jobs runner, and the panel polls and swaps
  the review in place.

## 15. The model client, caps and cost

### 15.1 `api/internal/llm`

A small client for OpenAI-compatible chat completions:
- Config: `LLM_BASE_URL` (the Vercel AI Gateway, `https://ai-gateway.vercel.sh/v1`),
  `LLM_MODEL` (a gateway slug, `provider/model`), and `LLM_API_KEY` from Secret Manager (already
  there; never logged; required at startup when the coach is enabled, failing fast with a clear
  error). `LLM_PRICE_IN_PER_MTOK` and `LLM_PRICE_OUT_PER_MTOK` are used only as a fallback when
  the gateway reports no cost for a request.
- `Complete(ctx, Request) (Response, error)`: messages, JSON response format, temperature and
  max tokens (default 1,200). It returns the content, the usage (prompt and completion tokens),
  and the **cost the gateway reports** for the request (from its usage block or cost headers;
  the builder confirms which against the gateway docs), plus the gateway's request id.
- Timeouts: 60 s per call, through the request context.
- Retries: up to 2 on 429, 5xx and network errors, with exponential backoff and jitter,
  honouring `Retry-After`. Validation retries (§10) are separate and counted once.
- Spend accounting from real numbers: each call's gateway-reported cost (or, only when the
  gateway reports none, its tokens × the fallback prices, flagged `estimated`) is added to
  `llm_spend (day, usd, calls, estimated_calls)` in one upsert. Every generation's tokens in,
  tokens out, cost, model slug and gateway request id are also stored on its `coach_reviews` row
  (§14), retries included.
- No provider SDK; plain `net/http` and JSON. The interface is small enough to fake in tests.

### 15.2 Caps (decided, values configurable)

- `COACH_DAILY_SPEND_CAP_USD` (default 25). Enforced from the recorded gateway costs in
  `llm_spend`, not from estimates. Once the UTC day's total reaches it, no new generations start
  that day.
- `COACH_DAILY_REVIEWS_PER_ACCOUNT` (default 30). This counts generations started for an
  account (as owner or officer) per UTC day. Cached reviews never count.
- `COACH_LLM_ENABLED` (default true): the kill switch (§20).

### 15.3 Cost (parametric)

Prices are left as `{in}` and `{out}` in US dollars per million tokens: the Vercel AI Gateway's
listed per-token prices for the DeepSeek V4.1 Flash slug at build time (§9.1). The builder fills
them in the build notes. The board's $0.181 per successful task is not our cost and is not used
here.
These figures are for planning only. The cap runs on the gateway's reported cost per request.

Token estimate per generated review (assumption, to be measured in the evaluation):
- Input: about **7,000 tokens**: coach rules about 1,500, persona about 3,000, evidence pack
  about 2,500.
- Output: about **800 tokens**.
- Validation retries: assume 10% of reviews need one, so ×1.1.

Cost per review = 1.1 × (7,000 × `{in}` + 800 × `{out}`) / 1,000,000
= **0.0077 × `{in}` + 0.00088 × `{out}`** dollars.

| Generated reviews per month | Input tokens | Output tokens | Monthly cost |
|---|---|---|---|
| 1 per review | 7,700 | 880 | 0.0077·`{in}` + 0.00088·`{out}` |
| 1,000 | 7.7 M | 0.88 M | 7.7·`{in}` + 0.88·`{out}` |
| 10,000 | 77 M | 8.8 M | 77·`{in}` + 8.8·`{out}` |
| 100,000 | 770 M | 88 M | 770·`{in}` + 88·`{out}` |
| One heavy Premium raider (≈ 80 fights opened a month) | 0.62 M | 0.07 M | 0.616·`{in}` + 0.0704·`{out}` |

If the provider discounts cached input, the persona and rules (about 4,500 of the 7,000 input
tokens) are a stable prefix. The real input cost falls accordingly, and the table above is the
upper bound.

### 15.4 When a cap is hit

The player gets the rule-only rendering with one quiet line above it: `Your coach is busy right
now. This is the review from your log alone; the coach's version will be here when you next
open it.` Nothing is queued while the cap holds. The next view after the cap resets generates.

## 16. Personas and voice per surface

The persona speaks only in the Coach panel and the history. The report row teaser, the Discord
line and the account line use the pack's own titles, in the site's voice, so no surface outside
the panel speaks as the coach.

## 17. API

### 17.1 Routes

- `GET /v1/reports/{id}/fights/{n}/coach?player={player_key}`: one player's review of one fight.
- `GET /v1/reports/{id}/coach/night?player={player_key}`: the night line (pack only).
- `GET /v1/me/coach/recent`: the account's last-raid line, plus the newest 20 stored reviews
  for Premium.
- `GET /v1/characters/{region}/{ruleset}/{name}/coach/history`: Premium, owner only. Stored
  reviews, newest first, 20 per page with a cursor.

### 17.2 Shapes

Full: `{ report_id, fight_index, player_key, spec, role, size, kill, status
("ready"|"writing"|"preparing"|"not_reviewable"), source ("model"|"rules"), busy (bool),
persona { label, status ("full"|"starter") }, reviewed_at, review { opening, fixFirst,
thenFix[], doingWell[], couldntCheck[] }, anchors { "a1": { kind, spell_id?, at_ms?, death? } },
basis_labels[] }`. Teaser: `{ status, access: "teaser", finding_count, fix_first: { title } }`.
Nothing else is sent to a free viewer. The gate is on the server.

### 17.3 Gating and errors

Report visibility, then the row (§19), then the feature. A caller who may not see the report
or coach the row gets the rating route's 404 (`not_found`). A flagged report answers `flagged`.
History without `log_coach` answers 403 `no_plan`. Responses are `Cache-Control: private,
no-store`.

### 17.4 Entitlement

Add `FeatureLogCoach Feature = "log_coach"` in `api/internal/entitlements/entitlements.go`.
`Standard()` already treats it as personal Premium. Add it to `auth.EntitlementsView`
(`log_coach`), the web's types and fixtures (`web/src/lib/account/api.ts`,
`web/src/test-support/billing-fixtures.ts`) and the launch-resolution test's feature list.

## 18. Discord and the account line, as built

- `discordbot/cmd_report.go` calls the coach service in-process for the night line. `share.go`
  renders without it.
- `notify.ReportReady` gains `CoachReady bool`. `guilds.sendReport` sets it when any owned
  character in the log belongs to an account holding `log_coach`. The embed appends the §5.5
  sentence.
- Copy lives in the existing copy files (`discordbot/copy_*.go`, `notify/report_embed.go`).

Performance: a cached pack or review is one indexed read. A pack miss is one bucket read plus
pure code, a few hundred milliseconds. Generation runs on the jobs runner, off the request.
Its time depends on the provider, and the panel shows the rule-only rendering meanwhile.
Profiles add one ~1 KB insert per rated kill row.

## 19. Privacy

- **Who may ask:** only viewers who may see the report. Export one `reports.MayView(actor,
  report, accounts)` and use it from rating and the coach, instead of rating's hand copy.
- **Whose row:** your own characters. Another player's only as an officer of the report's guild
  with `officer_views` for that guild (`Store.CanForGuild`), for a verified member of that
  guild. Moderators as officers. Nobody else sees anything about the coach on someone else's row.
- **Anonymised accounts:** shown only to the player and their guild's officers. Profiles store
  no names. Comparisons are medians of 20 or more players.
- **What leaves our systems:** the evidence pack goes through the Vercel AI Gateway to the
  model provider behind the chosen slug. It holds spec,
  boss, size and the numbers of one fight, and **no character name, account, guild, realm or
  Battle.net data**. The provider's training and retention stance is recorded with the model
  choice (see Owner decisions) and disclosed on the privacy page; it is not a gate. Owner,
  2026-10-27: "I don't care if they train on this — it's open data on our site anyway."
  The gate that holds is the one above: nothing personal leaves.
- **Channels:** nothing personal is ever posted to one.
- **Private reports** are reviewable by their owner, and never enter the pool.

## 20. Accuracy rules

1. Every number, spell, time and comparison in a review is in the evidence pack, and the
   validator proves it before the review is shown (§10).
2. Every pack finding carries at least one evidence anchor, and a test refuses one without.
3. Every review carries the "Couldn't check" list verbatim. Talents and race are on it in v1.
4. A finding fires only when the data proves it. Missing proof drops or narrows it, and says so.
5. Thin basis means no finding: `not enough logs yet`, with the real count.
6. Rule data, catalogue entries and personas are stamped and reviewed. Known-good fixtures pin
   the obvious answers (a Fury warrior's Bloodthirst count on a real Forever log, a flask that
   was there, a death to an avoidable ability), and a change that moves one fails the build
   until a person explains it (tenet 8).
7. **Fallback:** a review that fails validation twice is replaced by the rule-only rendering.
   The player never sees an unvalidated sentence.
8. **Nightly sample:** a job picks 20 reviews generated that day at random (spread across specs
   where possible). It writes each one's evidence pack, output, source and validation result to
   `coach_review_samples` and one structured log entry, for the owner to read. It carries no
   personal data beyond the character name. A moderator route lists them
   (`GET /v1/admin/coach/samples?day=`).
9. **Kill switch:** `COACH_LLM_ENABLED=false` forces the rule-only rendering everywhere, at once,
   with no deploy beyond the config change. Stored model reviews stay in the table but are not
   served while it is off.
10. Before the plans page renders the held line, each clause is checked against the shipped
    coach (§21).

## 21. Left to the builder

- **The model is DeepSeek V4.1 Flash** (§9.1). Confirm its slug from the gateway's model list
  (e.g. `deepseek/deepseek-v4.1-flash`) and set `LLM_MODEL` to it. Record the slug, its
  Reasoning and Data Analysis scores with the board's date, the provider behind it, its training
  and retention terms, and the gateway's `{in}`/`{out}` prices in the plan's build notes. Set `LLM_BASE_URL` to
  `https://ai-gateway.vercel.sh/v1` and `LLM_MODEL` to the slug. `LLM_API_KEY` is already in
  Secret Manager and readable by api-runtime. Confirm from the gateway docs where it reports
  per-request cost (usage block or headers), and record it.
- **Evaluation before launch:** run the same 20 sample fights (mixed specs, including the three
  full personas, a 10-player and a 20-player night, a wipe) through DeepSeek V4.1 Flash and
  MiniMax M3.1 Flash (the owner ruled Kimi K3 too expensive on 2026-10-27: stay with cheap
  models that score high on Reasoning and Data Analysis). Record both outputs, validation pass rates and measured tokens, and update §15.3's
  estimate. The owner reads both sets. Switch to MiniMax M3.1 Flash only if DeepSeek V4.1 Flash's
  quality is clearly worse.
- **Order of work:** §13's consumables fix and the rating bump. Then the `<spec>.json` rule files
  (DPS first), the catalogue's `complete` flags and `raid_buffs`. Then `Profile` and
  `coach_profiles` in the rater. Then `Pack` with fixtures and the rule-only rendering. Then the
  `llm` client, validator, caps and kill switch. Then the three full personas and the starter
  generator plus a person's read of each starter. Then the API. Then the ux-designer specs and
  the surfaces. Then the evaluation, then the plans-page line and changelog.
- **Migration number:** the next free number after `0054` when it lands. Only the deployed API
  migrates.
- **Client build tables per fight:** the build the log was recorded on if the report records
  it, else the build active at `fought_at`.
- **Anchors:** extend `momentHref` with `cast`, `time` and `resource` kinds, using the report
  page's existing tab and anchor conventions.
- **The banned-word list and the number normaliser** (`tone.go`, `validate.go`), with tests.
- **The held Premium line:** check each clause, "what to fix first", "the casts you missed"
  (F1), "the buffs you let drop" (F2), "the consumables you skipped" (F6), "learned from the
  engine" (§12), "the top logs on the site" (§11). Cut any clause the shipped coach doesn't
  deliver. The scope note changes per B5 (Owner decisions).

## 22. v2

- Talents from the synced addon export, and unspent points.
- Race.
- Personal sim: the pull simulated with the player's own gear, race and talents, with the gap
  split into gear, buffs and execution.
- Movement and positioning from advanced logging.
- More curated mechanics: avoidable damage and interrupt duty for every raid boss.
- Trend over weeks: the same finding across raid nights.
- **Learn from a player's own history across weeks:** the coach reads its own past packs for
  the character and says what improved.
- **Voice and tone per guild:** a guild chooses how blunt its raiders' coaches are.
- The 25 starter personas deepened to full.
- Watch a top log: a link to a public pool log, once naming and anonymity are settled.
- Asking the coach a follow-up question.
- Wipe review for raid leaders, on the guild plan.

---

## 23. Builder check

Re-read as a builder new to the codebase, with the answers the codebase or the brief gives.

1. *Where is the fight data?* The posted summary in the bucket, read with `sims.fightSummary`.
   The rater has it in memory at `reports/ingest.go` `rate`.
2. *Is `log_coach` automatically Premium?* `Standard()` says yes, but `EntitlementsView`, the web
   types and fixtures and the test list are explicit (§17.4).
3. *How do I partition?* Monthly ranges on `fought_at` via `db.EnsurePartitions`. Add the table
   to the partition job (§11.1).
4. *Player key and ownership?* `region/ruleset/name`, and `auth.Store.OwnsCharacter`.
5. *Who may see a report?* `reports.mayView`. Export it and use it in both places (§19).
6. *What is "evidence"?* `presence.EvidenceReport`: not private, not flagged.
7. *Officers?* `entitlements.Store.CanForGuild(guild, user, FeatureOfficerViews)`.
8. *Can the channel post carry a per-player line?* No (R-A3). See §5.5.
9. *Is there an LLM client already?* No. Build `api/internal/llm` (§15.1) on plain `net/http`
   against the Vercel AI Gateway's OpenAI-compatible endpoint. The key `LLM_API_KEY` is already
   in Secret Manager.
10. *Which model, and what does it cost?* DeepSeek V4.1 Flash through the gateway. The §9.1 rule
    is for picking again later. Prices go
    in the build notes, and §15.3 is parametric. The cap runs on the gateway's reported cost.
11. *Does every provider support JSON-schema output?* No. Ask for JSON mode, use the schema form
    where supported, and validate on the server always (§9.2).
12. *What if the free teaser's "first thing" disagrees with the model's priority?* It can't.
    `fixFirst` must be the pack's rank 1 (§9.3).
13. *Where do drills come from, and may they hold numbers?* From the persona's drill list, and
    numbers in the chosen drill's text are allowed (§8.1, §10).
14. *How are causal claims checked?* Only through pack links, which are deterministic overlaps.
    A sentence relating two findings must cite one (§7, §9.3).
15. *Where does generation run?* On the existing jobs runner. The panel shows the rule-only
    rendering and polls (§14).
16. *What happens to old wrong preparation scores?* The model bump re-rates, and the digests are
    rebuilt (§13).
17. *Account "last raid" line?* New on the web. Reuse `addon.lastRaidFights` (§5.6).
18. *Which specs get a sim reference?* Validated DPS specs only (§11).
19. *Raid buff list?* Add `raid_buffs` with provider class to the catalogue (§6.2 F6).
20. *Design step for surfaces?* The ux-designer spec and mock, and two SHIP verdicts on
    screenshots (tenet 6).

## Owner decisions

1. **The held Premium line says "after every raid".** The coach needs a log of the raid. Keep
   it, or change it to "after every logged raid"?
2. **The Premium scope note** must change when the line renders (B5). Proposed: `Sims cover
   damage specs; healing and tanking specs are not simulated yet. The coach reviews every spec;
   damage specs get the most findings.`
3. **"Top" for healers and tanks.** I chose the rating's overall score. Confirm, or name another
   measure, such as effective HPS.
4. **The model provider's data terms.** Decided 2026-10-27. DeepSeek makes no training or
   retention commitment for API traffic and is not on the gateway's no-training list (nor are
   MiniMax, Zai, Xiaomi, NVIDIA or StepFun; only Alibaba Cloud and Moonshot AI qualify among the
   cheap models). The owner ruled that training on the pack is acceptable because the pack holds
   only data already public on the site and nothing personal (§19). The stance is recorded in the
   plan's build notes and disclosed on the privacy page; `disallowPromptTraining` is NOT set.
5. **The privacy page** (legal pages finalised 2026-10-24) should name Vercel (the AI Gateway)
   and the model provider as processors of fight data sent for coaching. The pack carries no
   names or account data. Approve the addition.
6. **Cap defaults:** $25 a day total spend and 30 generated reviews per account per day. Confirm,
   or set the numbers once §15.3 is filled with real prices.
