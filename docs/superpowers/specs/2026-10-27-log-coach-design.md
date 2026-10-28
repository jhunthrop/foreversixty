# Log coach v1: what to fix first, from your own log (product design)

Status: product design for the owner, 2026-10-27. Base: the architect's map in
`docs/superpowers/specs/2026-10-27-log-coach-feasibility.md` (its §6 v1 slice and §7 risks).
Player evidence: `design/reviews/2026-10-27-premium-differentiators/README.md` and
`design/reviews/2026-10-27-launch-pricing/appeal-*.md`. The Premium line this feature unlocks
is held in `design/specs/2026-10-27-launch-pricing.md` §B5 (`premium.featuresHeld.coach`).
Decisions marked **(decided)** are the controller's calls of 2026-10-27 and are not open.

This is the product design. Each player-facing surface below still needs its experience spec
and mock from the `ux-designer`, reviewed by `wow-player`, before it is built (tenet 6). Those
specs build on the copy patterns and rules here. They do not reopen them.

---

## 1. Intent, and why now

The owner's direction: "learn from the best players in the world and create a coach to review
logs." After a raid, the coach reads one player's fight and says, in WoW words, what to fix
first: the casts they lost, the buffs they let drop, the consumables they skipped, the death
that cost the most. It measures against the best logs on the site for the same spec, boss and
raid size. Where there aren't enough of those yet, it measures against the engine's written
rotation, and failing that, against what the spell data allows.

Why now:

- **It is what players said they would pay for.** Four of five personas in the
  differentiator round asked for the same thing: "coach me on my own log." The Fury warrior
  in the appeal round said he would pay if Premium told him "you clipped 14 Bloodthirsts".
  The Holy paladin asked for "what to fix after Onyxia". Today's Premium is "more of the free
  tools", and none of them would buy that alone. The coach is Premium's main differentiator.
- **We already own nearly all of it.** The parser stores casts with timestamps, aura
  tracks, resource tracks, deaths and gear. The engine holds a written rotation for every
  spec. The rating already turns deaths and dropped debuffs into anchored moments. The
  missing joins are a comparison table, a pure review function and the surfaces.
- **The calendar.** Forever launches on 4 November and raids open on 9 December. The coach
  should be live for the first raid night, and the plans page's held line goes live with it.

## 2. Goals, non-goals, anti-goals

Goals

- After any logged boss pull, a player can open their own review: up to eight findings,
  ordered so the first one is what to fix first, each with evidence they can click.
- Every finding states what it is measured against, in plain words.
- Every review lists what it couldn't check, and why.
- Damage dealers get the full set. Healers and tanks get the findings the data supports,
  framed for their role (§5.3).
- Free accounts see that the coach found something on their own character (the count and
  the top finding's title). Premium (`log_coach`) sees all of it.
- One line in Discord and one line on the account page carry "what to fix first" to the
  player without them opening a report.
- 10-, 20- and 40-player raids are equal citizens: each size has its own comparison pool.

Non-goals (v1, and named again in §16)

- Talents, race and a sim of the player's own character. Forever logs record neither talents
  nor race (feasibility §1). Personal sim needs the addon-export builder.
- Movement, positioning, facing and range.
- Mechanics beyond the curated tables that already exist.
- Trends over weeks.
- Trash pulls. Gear and consumables are only in the log on boss pulls.
- A coach that writes free text. Every sentence comes from a copy table filled with numbers
  from the data. There is no language model in the loop.

Anti-goals (a build that does any of these fails review)

- **A finding that is wrong.** If the data can't prove it, it is not a finding. It becomes a
  line under "Couldn't check", or the finding is left out. Accuracy is absolute (tenet 8).
- **A finding that shames.** No ranking of raiders against each other. No named finding in a
  channel anyone else can read. No coaching of another player's character except by their
  guild's officers. No "heal more" for healers. The copy states facts and never grades the
  person ("you failed", "bad", "terrible" never appear).
- **A wall of numbers.** At most eight findings per fight, one per finding type. Each is a
  title, one line of detail, one basis line and its evidence links. The full numbers sit on
  the report's existing tabs, one click away through the evidence links (tenet 11).

## 3. Who sees what (decided)

| Viewer | Report page row | Discord | Account page |
|---|---|---|---|
| Signed out, or signed in but not the character's owner | Nothing about the coach | Nothing | n/a |
| Owner of the character, free | The finding count and the top finding's title, plus an upgrade link | Nothing (the Discord line is for Premium members only) | Last raid line: count and top finding's title |
| Owner of the character, Premium (`log_coach`) | Every finding with detail, basis, comparison, timestamps and evidence. Coach history | `/report` answer: one "what to fix first" line, ephemeral. Raid-night post: one line that names nobody | Last raid line: what to fix first, with a link to the review |
| Officer of the report's guild, with `officer_views` for that guild, viewing a verified member's row | The full review, marked "Officer view" | Nothing extra (officer coaching stays on the site) | n/a |
| Moderator | As an officer, on any report they can open | Nothing | n/a |

"Owner of the character" means `auth.Store.OwnsCharacter(user, player_key)` is true: the
same test the Premium history route uses (`api/internal/rating/handler.go:258`).

`log_coach` is a new standard personal Premium feature (§10.4). A guild seat holder has it,
because seats resolve standard features.

## 4. The player's experience, per surface

### 4.1 The finding, in one pattern

Every finding on every surface is built from the same four parts, all from one copy table
(`api/internal/coach/copy.go`):

1. **Title.** What happened, under eight words, with at most one number: `Bloodthirst: 7 casts lost`.
2. **Detail.** One sentence of fact with the numbers that prove the title, and the one or two
   worst moments as fight times (`m:ss` from the pull): `31 of a possible 38 with rage in
   your bar. Longest gaps 1:12–1:25 and 2:40–2:51.`
3. **Basis.** What it is measured against, in plain words, with the comparison when one
   exists: `Measured against Bloodthirst's 6-second cooldown and your own rage bar. Top 25
   Fury warriors on Onyxia (40): 9.1 a minute. You: 8.0.`
4. **Evidence anchors.** Links into the report page that show the proof: the Casts row, the
   aura row, the death, the resource graph at a time. They are built the way
   `web/src/lib/rating/moments.ts` builds the rating's moment links today.

Under the findings, every review ends with:

- **Couldn't check:** one line per thing the coach could not read, and why. For example:
  `Talents and race: Forever logs don't record them.` and `Range and movement: the log
  doesn't say where you stood.` This list is never empty. Talents and race are always on it in v1.
- **Reviewed:** the date and the coach version, e.g. `Reviewed 28 Oct.`

When there are no findings: `Nothing to fix on this pull that the log can show.` The
"Couldn't check" list follows.

When the basis is too thin for a comparison finding (§6.4), the finding is not shown. A line
takes its place under "Couldn't check": `Overheal by rank: not enough logs yet — 12 Holy
priest kills of Onyxia (40) on the site, the coach needs 20.`

All figures are rendered from the data at request time. The design and the copy table never
hard-code a nightly number (the owner's rule).

### 4.2 Worked examples

**All values below are invented examples for this document. None of them is a measurement.**

**Fury warrior, Onyxia (40), kill, 3:52**

| # | Title | Detail | Basis | Evidence |
|---|---|---|---|---|
| 1 | `Bloodthirst: 7 casts lost` | `31 of a possible 38 with rage in your bar. Longest gaps 1:12–1:25 and 2:40–2:51.` | `Measured against Bloodthirst's 6-second cooldown and your own rage bar. Top 25 Fury warriors on Onyxia (40): 9.1 a minute. You: 8.0.` | Casts row for Bloodthirst; two links to the timeline at 1:12 and 2:40 |
| 2 | `Rage full for 14 seconds` | `Your rage sat at its cap for 14 s outside Onyxia's air phase, longest at 3:05–3:12. Rage past the cap is lost.` | `Measured against your own rage bar, with the air phase left out. Top 25 Fury warriors here: 3 s.` | Rage graph at 3:05 |
| 3 | `No flask at the pull` | `You pulled with Grilled Squid and an Elemental Sharpening Stone, and no flask.` | `From the character sheet in your log, against the site's consumable list for damage dealers.` | Buffs tab, your row |

What to fix first: `Bloodthirst: 7 casts lost`.

**Holy priest, Onyxia (40), kill, 3:52** (efficiency and survival, never "heal more")

| # | Title | Detail | Basis | Evidence |
|---|---|---|---|---|
| 1 | `Out of mana from 3:17` | `Your mana stayed under 5% from 3:17 to the kill, 35 s in all.` | `Measured against your own mana bar.` | Mana graph at 3:17 |
| 2 | `Greater Heal (Rank 4): 44% overheal` | `44% of what Greater Heal (Rank 4) healed was overheal, over 23 casts. Lower ranks cost less mana.` | `Top 21 Holy priests on Onyxia (40): 22% on the same rank.` | Healing tab, the ability row |
| 3 | `Died at 2:58 to Flame Breath` | `Flame Breath hit for 2,140 at 2:58. No healing potion or Healthstone used in the 10 s before.` | `From your log. Flame Breath is on this boss's list of damage you can avoid.` | Deaths tab, this death |

What to fix first: `Died at 2:58 to Flame Breath`.

**Protection warrior, Onyxia (40), kill, 3:52**

| # | Title | Detail | Basis | Evidence |
|---|---|---|---|---|
| 1 | `Died at 3:30 with Shield Wall ready` | `Onyxia's melee killed you at 3:30. Shield Wall had been ready since the pull.` | `From your log and Shield Wall's cooldown.` | Deaths tab, this death; Casts row for Shield Wall |
| 2 | `Shield Block up 18% of the fight` | `Shield Block was up for 42 s of 3:52.` | `Top 20 Protection warriors on Onyxia (40): 34%.` | Buffs tab, Shield Block row |
| 3 | `Sunder Armor: 22 casts of Rank 4` | `Rank 5 is trainable at your level. Rank 4 was cast 22 times.` | `Measured against the spell ranks a level 60 warrior can train.` | Casts row for Sunder Armor (Rank 4) |

What to fix first: `Died at 3:30 with Shield Wall ready`.

### 4.3 Surface 1 (primary): the report page player row

- The player's row in the report's combatant list carries one coach line, but only for
  viewers who may coach that row (§3). Premium: `What to fix first: Bloodthirst: 7 casts
  lost · 4 findings`. Free: `4 findings · First: Bloodthirst: 7 casts lost · See them all
  with Premium`.
- The line opens the **Coach** panel for that player and fight. It holds the findings in
  order, the "Couldn't check" list and the "Reviewed" line. Premium also gets a **History**
  link listing this character's past reviews (§10.2). Where the panel sits (a tab beside
  Rating, or a drawer from the row) is the ux-designer's call. It belongs on the report page:
  the report page owns the log, and the coach is a reading of it (one job per page).
- Abilities carry their icon and rank, and items carry icon, quality colour and tooltip
  (tenets 2 and 3). Fight times are links (§4.1).
- States: loading (`Reading the fight…`), a reference still preparing (§9.2), no findings,
  not reviewable (§7.1), the free teaser, the officer view, and an error (`The coach couldn't
  read this fight just now.`).

### 4.4 Surface 2: Discord (decided: Premium members only, never shaming in a channel)

- **`/report` answer (ephemeral).** When the viewer holds `log_coach` and owns a character
  in the matched log, the answer gains one line for that character's night: `What to fix
  first: Bloodthirst casts lost — on 5 of 8 bosses. Open your review` (a link to the
  coach panel on the report page). The line is never in the version `Share in channel`
  reposts. Share re-renders the answer without it. A viewer with several characters in the
  log gets one line each, at most three.
- **Raid-night post (`report_ready`, in the guild channel).** The post is counts only and
  never names a raider (rule R-A3, `api/internal/notify/report_embed.go`). The coach keeps
  that rule. When at least one raider in the matched log holds `log_coach` on a character
  they own, the description gains one sentence that names and counts nobody: `Premium
  raiders: your coach review of tonight is ready on Forever Sixty and in /report.` The
  personal "what to fix first" stays in the ephemeral answer and on the site. That is what
  "Premium members only, never shaming in a channel" comes to under R-A3.
- **Night line rule.** The top finding across one character's fights that night: the
  highest severity first, then the finding type that showed on the most bosses, then the
  earliest boss. It is worded from the finding's title stem without the per-fight number
  (`Bloodthirst casts lost`), plus `on N of M bosses`.
- **Timing.** The night line is read from cached reviews (§9). If they are not all cached
  within the interaction's budget, the line becomes `Your coach review is on the site.` with
  the link, and the reviews warm in the background. Discord never waits on the coach.

### 4.5 Surface 3: the account page's "last raid" line

- One line, for the account's most recent raid. "Last raid" has the definition the addon
  already uses: the newest evidence report holding a rated boss fight of the character, within
  `addon.LastRaidWindow` (14 days), read with the same query (`api/internal/addon/inbox_reads.go`).
  The coach reuses that query rather than inventing a second meaning of "last raid".
- Premium: `Last raid · Molten Core, Thu 23 Oct · What to fix first: Bloodthirst casts lost,
  on 5 of 8 bosses · Open`. Free: `Last raid · Molten Core, Thu 23 Oct · 11 findings ·
  First: Bloodthirst casts lost · See them all with Premium`.
- No last raid in 14 days: the line is not shown. The account page's character list
  (`web/src/components/account/CharacterList.svelte`) is the likely home. The ux-designer
  places it.
- The account page loads the line separately from the rest of the page (`GET
  /v1/me/coach/recent`), so a cold cache never slows the page.

## 5. The eight findings

### 5.1 Shared rules

- **Reviewable fights.** Boss pulls only (`encounter_id` set), kills and wipes, at least 30 s
  long. On a wipe, deaths in the last 15 s before the fight ends are part of the wipe and are
  not reviewed.
- **Spec and role** come from the roster row (`RosterRow`). A row with no known spec is not
  reviewed: `We couldn't tell this player's spec from the log.`
- **Level** comes from the advanced-logging fields (`AdvLevel`). Without it, findings that
  need the level (F5's rank check, the sim band) go to "Couldn't check".
- **Time that never counts against a player:** time dead, the curated downtime windows of
  the encounter (the same windows `rating.ActivityIdleMS` subtracts), and the time before the
  player's first cast of the fight. The opening wait for threat is not a mistake.
- **Talent-granted abilities** are assessed only if the player cast them at least once in
  the same fight. One cast proves they have the talent. Forever logs carry no talents.
- **Pets and guardians** fold into their owner: `CastRow.OwnerGUID`, and abilities with `Via`.
- **Per-spec coach data.** Every spec gets a curated file, `data/curated/coach/<spec>.json`.
  It lists the spec's core cooldown abilities, the auras it owns (its own buffs and debuffs),
  deliberate downranks, defensive cooldowns and efficiency cooldowns. Every entry carries a
  `verified` stamp against the client's own tables, like `consumables/catalogue.json` does.
  A test refuses an unstamped entry. If a spec has no file, or a list is missing, the findings
  that need it go to "Couldn't check". Nothing is guessed.
- **Severity** is `major` or `minor`, set by each finding's thresholds below.
- **Order.** Major before minor. Within a tier, the role's precedence list (§5.3) decides. The
  first finding is "what to fix first". Findings about things the player doesn't control (raid
  buffs) are never first.
- **One finding per type** per fight, so eight at most. Where a type could fire for several
  abilities (F1, F2), the title names the worst one and the detail names up to two more.
- **Gain estimates** (`GainEstimate`) stay null in v1 and are never shown. A DPS figure the
  coach can't prove is a finding that's wrong.

### 5.2 The findings

Thresholds are named constants in `api/internal/coach/rules.go`. The values here are the v1
values.

**F1. Core ability under-used** (`core_casts`)
- *Rule.* For each core ability in the spec's coach file with a cooldown of 4 s to 10 min
  (`spellconst` `cooldown_ms`; longer cooldowns are saved across the night and are not judged
  per pull), walk `CastRow.Sequence`. A gap counts as **ready time** when the ability was off
  cooldown, the player was alive and outside downtime, and the player's resource was at or
  above the spell's base `cost` in at least one whole second of the per-second
  `ResourceTrack.Series`. Base cost is an upper bound, because talents can only lower it, so
  this never invents a lost cast. Each gap first gives up a 1.5 s grace for GCD collisions.
  Lost casts = floor(sum of remaining ready time / cooldown). "Possible" = actual + lost.
- *Fires* when lost ≥ 2 and lost ≥ 10% of possible. *Major* at ≥ 20%.
- *Data.* `CastRow.Sequence`, `ResourceTrack.Series`, deaths, curated downtime, `spellconst`.
- *Basis.* The cooldown and the player's own resource bar always decide whether it fires. The
  comparison line comes from the basis order (§6): top logs' casts per minute, else the
  written rotation's sim casts per minute, else none ("of a possible N" stands alone). The
  sim comparison is skipped for any ability on that spec's `WorstActions` list on the sim
  support page (`api/internal/sims/specs.go`), where the sim and real logs disagree.
- *Disabled* (goes to "Couldn't check") with no coach file, an ability with no `spellconst`
  row, or no resource track for the player. If an ability's cost is unknown, gaps count
  without the resource check, and the detail adds `Couldn't check whether you had the rage.`
- *Roles.* DPS: damage cooldowns (Bloodthirst, Mortal Strike, Aimed Shot, Death Wish…). Tank:
  threat cooldowns (Shield Slam, and Mangle where the spec has it). Healer: **efficiency
  cooldowns only** (for example Inner Focus, Mana Tide Totem, Nature's Swiftness), worded as
  a free or cheaper spell left unused, never as missing healing.

**F2. Key buff or debuff uptime** (`uptime`)
- *Rule.* For each aura the spec's coach file says the player owns: on the boss (debuffs) or
  on the player (self-buffs), uptime from the first application to the fight's end, leaving
  out time dead and downtime. A debuff is judged only when the player is its **sole applier**
  on that target (`AuraTrack.Appliers` holds only the player). Otherwise it goes to "Couldn't
  check": `Your share of Sunder Armor: 3 warriors applied it, and the log merges them.`
  (feasibility §7).
- *Fires* when uptime is ≥ 10 points under the reference and under 90%. *Major* at ≥ 25 points under.
- *Basis.* Top logs' median, else the written rotation's sim uptime. Spell data has no uptime
  ceiling, so with neither of those the finding becomes `not enough logs yet`.
- *Roles.* DPS: own DoTs, debuffs and self-buffs (Corruption, Serpent Sting, Slice and Dice,
  Battle Shout). Tank: **defensive uptime** (Shield Block, Demoralizing Shout/Roar, Thunder
  Clap). Tanks have no validated sim, so this needs top logs. Healer: own self-buffs (Inner Fire).

**F3. Dead time** (`idle`)
- *Rule.* `rating.ActivityIdleMS` (alive, outside downtime and assignments, nothing done; the
  1.5 s gap rule), plus its three longest windows with timestamps. That needs a windows variant
  of the same code in `logs/engine/rating/activity.go`, sharing its logic rather than copying it.
- *Fires* at ≥ 10 s and ≥ 5% of counted time. *Major* at ≥ 10%.
- *Basis.* Top logs' median active %, else the threshold alone ("Measured against your own
  casts, with downtime left out").
- *Roles.* DPS and tank. **Off for healers**: waiting so as not to overheal is part of healing (rule g).

**F4. Resource waste** (`resource`)
- *Rule.* Rage and energy: seconds at the cap (`Series` against `Max`) outside downtime and
  time dead. Mana: seconds under 5% of `Max`, and the first such second as a fight time.
- *Fires.* Capped ≥ 10 s and ≥ 5% of counted time. Mana under 5% for ≥ 10 s, starting
  before the last 10% of the fight. *Major* at ≥ 15% capped, or out of mana before 75% of the fight.
- *Basis.* The player's own bar always, with top logs' median capped seconds when available.
- *Roles.* DPS and tank rage/energy; casters and healers mana. Healer copy is about running
  out early (efficiency), never about leftover mana.

**F5. Ranks** (`ranks`)
- *DPS and tank rule.* Casts of a rank lower than the highest rank trainable at the player's
  level, for spells on the spec's rotation (`Rotation.Family`, `spellranks.json`, `trainables/`).
  Deliberate downranks listed in the coach file are left out. *Fires* at ≥ 3 casts or ≥ 10% of
  that spell's casts. Always *minor*. *Basis:* the spell ranks a player of that level can train.
- *Healer rule (rank choice).* Overheal % per heal **rank**: `Ability.Overheal` / `Ability.Total`,
  and the spell id is the rank. *Fires* when a rank with ≥ 10 casts is ≥ 15 points above the
  top logs' median for the same rank. *Major* at ≥ 30 points. *Basis:* top logs only. Below
  the minimum it becomes `not enough logs yet`. The detail may add `Lower ranks cost less
  mana.` (spell data), and never more than that.
- *Disabled* without the player's level (DPS/tank rule) or without a top-logs pool (healer rule).

**F6. Pull preparation** (`preparation`)
- *Rule.* What the player had at the pull, from the **seeded aura tracks**: segments on the
  player that start at 0 ms, seeded from COMBATANT_INFO. They are read against the consumables
  catalogue for the player's role, which is the same fix as §8. Categories: flask, food, weapon
  oil or stone, combat potion uses (from `Casts`, against `MaxUses`), missing permanent enchants
  (from the log's gear, through `score_v2.enchanted`), and raid buffs.
- *Completeness guard.* A category only produces a "missing" finding when the catalogue marks
  that category `complete`: every buff of that kind in the client is listed and stamped. A
  food category that lists only some foods would call a fed player unfed. While a category is
  incomplete it goes to "Couldn't check": `Food: the coach's food list isn't complete yet.`
  Weapon oils and stones stay in "Couldn't check" until someone verifies how they show in
  COMBATANT_INFO (feasibility §7).
- *Raid buffs* (decided, the 10-man rule) are flagged only when a class that provides the buff
  was in that fight's roster. They are always *minor* and never first, because the player
  doesn't control them: `Pulled without Arcane Intellect (a mage was in the raid).` The
  catalogue gains a `raid_buffs` list with the providing class and every rank's spell id.
- *World buffs* are never a finding. Premium's detail lists them as facts (`World buffs at the
  pull: Rallying Cry of the Dragonslayer, Songflower Serenade`).
- *Severity.* Flask, food or a weapon enchant missing: *major* in 40-player raids, *minor* in
  10 and 20. Potions and raid buffs: *minor*. A missing weapon enchant: *major* in every size.
- *Disabled* when the fight has no character sheet for the player (no COMBATANT_INFO, a trash
  pull, or advanced logging off): `Consumables and enchants: your log has no character sheet
  for this pull.`

**F7. Death review** (`death`)
- *Rule.* The first reviewable death (§5.1): its time, the killing blow, whether that damage is
  on the encounter's avoidable list (curated mechanics), and the defensive cooldowns from the
  coach file that were off cooldown at the moment of death. Baseline abilities count once the
  player is at their trainable level. Talent abilities count only if cast in the same fight. It
  also states whether a healing potion or Healthstone was used in the 10 s before. That is a
  fact about casts, never a claim about what was in their bags. It reuses the rating's death
  `Moments` and anchors.
- *Fires* on any death before the last 10% of a kill. *Major* before 75% of the fight.
- *Basis.* `From your log`, plus the cooldown for "ready" claims.
- *Roles.* All. Tank and healer copy leads with survival: `Died at 3:30 with Shield Wall ready`.

**F8. Rotation adherence** (`adherence`)
- *Rule.* `rating.fightAdherence` (the share of casts that are on the shipped rotation; nil
  under `AdherenceMinCasts`), plus the three off-rotation spells cast most.
- *Fires* when adherence is ≥ 10 points under the top logs' median, or under 85% when only the
  written rotation is available. *Major* at ≥ 25 points under.
- *Basis.* Top logs' median adherence, else `the written {Spec} rotation`.
- *Roles.* DPS and tank. Off for healers, whose written lists aren't rotations in this sense.

### 5.3 Which specs get which findings (v1)

DPS specs (19): Balance and Feral (cat) druid; Beast Mastery, Marksmanship and Survival hunter
(Survival is melee in Forever); Arcane, Fire and Frost mage; Retribution paladin; Shadow priest;
Assassination, Combat and Subtlety rogue; Elemental and Enhancement shaman; Affliction,
Demonology and Destruction warlock; Arms and Fury warrior.

Healer specs (5): Restoration druid, Holy paladin, Discipline priest, Holy priest, Restoration shaman.

Tank specs (3): Feral (bear) druid, Protection paladin, Protection warrior.

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

Precedence within a severity tier:
- DPS: F7, F1, F3, F6, F2, F4, F8, F5.
- Healer: F7, F4, F5, F1, F6, F2.
- Tank: F7, F2, F1, F3, F6, F4, F8, F5.

A spec with no coach file still gets F3, F4, F6 and F7, which need no per-spec list. F1, F2
and F5's DPS rule go to "Couldn't check" for it.

## 6. Where the comparison comes from (decided basis order)

### 6.1 The order

For each comparison a finding makes, the first one that exists wins:

1. **Top logs on the site**: the same spec, encounter, difficulty and size, when the pool
   holds at least 20 kills (§6.3). Basis line: `Top {n} {Spec} {class} on {Boss} ({size}): …`.
2. **The engine's written rotation** (the sim reference profile, §7). DPS specs only, and only
   specs whose sim is `validated` on the support page (`sims.Store.Validated`). Basis line:
   `The written {Spec} rotation, simulated at level {band}: …`.
3. **Spell data**: cooldowns, costs and ranks. Basis line: `Measured against {Spell}'s
   {n}-second cooldown…`.

Facts that compare with nothing (a death, a missing flask) carry `From your log`. Every finding
names its basis in plain words. The API carries it as data (§10).

### 6.2 `coach_profiles`, the comparison pool

Today the per-ability numbers of other players live only in per-fight summaries in the bucket.
Reading those per request is too slow (feasibility §4). The rater writes a compact profile
beside each `rating_scores` row instead.

- **Table** `coach_profiles`, partitioned by range on `fought_at` like `fight_metrics` and
  `rating_scores`. It is added to the partition job's table list (`api/internal/db/partitions_job.go`).
- **Columns:** `report_id text, fight_index int, player_key text, class text, spec text, role
  text, encounter_id int, difficulty int, size int, kill boolean, duration_ms int, phase text,
  metric numeric, fought_at timestamptz, model_version text, profile jsonb`, primary key
  `(report_id, fight_index, player_key, fought_at)`.
- **`metric`** is what "top" means: DPS for damage dealers. For healers and tanks it is the
  rating's overall score at the current rating model. HPS rewards overheal, and damage taken
  rewards nothing.
- **`phase`** is the content phase holding `fought_at`, from `data/curated/phases.json`.
- **`profile`**: alive ms and counted ms; active %; casts per minute per ability family; uptime %
  per owned aura (sole applier only); capped % per resource; mana-under-5% ms and first time;
  overheal % per heal rank; defensive uptime %; adherence; consumables present at the pull.
  Roughly 1 KB per row.
- **Written** by the rater, in the same transaction as the rating rows, from the summary already
  in memory. It is written for kills only, on reports that are ranked (`reports.Ranked`: not
  private) and evidence (`presence.EvidenceReport`: not flagged). Profiles are removed on the
  same path that withdraws ranking rows (`Ranker.RemoveReport`: moderation, a report turned
  private, deletion).
- **Indexes:** `(encounter_id, difficulty, size, spec, phase, kill, metric desc)` and
  `(player_key, fought_at desc)`.
- **Backfill.** The rating model bump in §8 makes the existing backfill re-rate every fight,
  and that writes profiles for history too.

### 6.3 The comparison itself

- **The pool** for a fight: kills of the same `(encounter_id, difficulty, size, spec, phase)`
  at the current profile `model_version`, with **one kill per player**: each `player_key`'s
  best by `metric`. The reviewed player's own key is left out, so nobody is compared with
  themselves.
- **The 20-kill minimum** counts those distinct players. Ten clears by one guild's warrior are
  one kill. This is the existing `MinSample = 20` rule, applied to people.
- **The comparison set** is the top 25 of the pool by `metric` (all of them between 20 and 25).
  The coach compares with the **median** of that set, never the best single log.
- **Names never appear.** The basis says `Top 25 Fury warriors on Onyxia (40)`. No pool member
  is named or linked in v1.
- Pool medians per bracket are cached in memory for an hour (§12).

### 6.4 When the basis is thin

If a comparison finding has no basis at its level (no pool and no validated sim for an uptime;
no pool for healer overheal), it is not shown. "Couldn't check" says so, with the real count:
`not enough logs yet — {n} {Spec} kills of {Boss} ({size}) on the site, the coach needs 20.`
The coach never drops to a weaker standard than the finding's rule allows.

## 7. The sim reference cache

- **What it is.** The spec's written rotation (`data/curated/apl/<spec>.json`) run by the engine
  at the player's level band. It uses the band's preset gear and the same preset and race the
  BiS ladder runs for that band. Fight length is rounded to a 30 s bucket. The result's
  `Summary` gives casts per minute and uptime per aura, averaged per iteration (feasibility §2).
- **When it is used.** DPS specs whose sim is `validated`, as step 2 of the basis order.
- **Cache table** `coach_sim_refs`: key `(engine_version, build, spec, band, duration_bucket)`,
  value the reference profile (the same shape as `coach_profiles.profile`) and `computed_at`.
  Each key is computed once per engine version. About 3 s of native compute at `ScoreIterations`.
- **A miss** queues one reference job on the existing jobs runner. The review answers `status:
  "preparing"`, and the panel says `Reading the fight against the written rotation…` and polls.
  The review is not cached until the reference is in. A nightly warm for level 60 × the DPS specs
  × the duration buckets of curated encounters is allowed but optional (§15).
- **Honesty.** The basis line says "the written rotation", never "the best possible". APLs are
  curated and some talents are not modelled (feasibility §7).

## 8. Fix first: consumables at the pull (decided)

`CombatantRow.Consumables` is never filled in production, because neither the API nor the
companion passes the catalogue into the parser (feasibility §1). So three readers are wrong
today: `rating.consumablesMissed` (`api/internal/rating/last_raid.go`), `fightDiscipline` /
`carries` (`api/internal/rating/score_v2.go`) and the engine's `scorePreparation`
(`logs/engine/rating/preparation.go`). This lane fixes that first, before any coach code:

1. One function in the engine, e.g. `consumables.AtPull(fight, guid, catalogue)`, reads the
   player's seeded aura tracks (segments on that player starting at 0 ms) against the catalogue
   and returns what was up at the pull. It needs no companion release.
2. The three readers use it instead of `CombatantRow.Consumables`. So does the coach's F6.
   There is one reader of "what did they bring", and it is shared.
3. Raid buffs follow the same path. `MissingBuffs` is not used anywhere new, because it ignores
   who was in the raid (feasibility §5).
4. Bump `ratingengine.DefaultModelVersion`. The existing backfill (`staleFights`) re-rates every
   stored fight. The preparation and discipline percentile digests are rebuilt from the
   re-rated rows, so old zeros don't stay in the brackets.
5. Known-good fixtures pin it: a real Forever pull with a flask and food reads as prepared, and
   one without reads as missing (tenet 8).

The addon's last-raid "missing consumables" phrase and the rating card's Preparation and
Discipline become right with this fix, before the coach ships.

## 9. Computing a review

### 9.1 The pure core

`api/internal/coach`:
- `Review(fight summary.Summary, row summary.RosterRow, ref References, tables Tables) Review`
  is a pure function. `Tables` holds the coach file, APL and rotation, `spellconst`,
  `spellranks`, `trainables`, the consumables catalogue and the encounter's curated mechanics.
  `References` holds the pool medians and the sim reference, each nil when absent.
- `Profile(fight, row, tables) Profile`: the same measurements, written to `coach_profiles`
  by the rater. Both sides call the same measuring code, so the player and the pool are measured
  identically.
- `Finding { Code, Severity, Params, Basis {Kind, N, Label}, Evidence []Anchor }`. The stored
  review keeps codes and params. Copy is rendered at read time from the copy table, so a wording
  change doesn't invalidate the cache.

### 9.2 On request

1. Check access (§13). Then look up the cache row `(report_id, fight_index, player_key,
   model_version)` in `coach_reviews`.
2. On a miss: read the fight summary from the bucket (`sims.fightSummary`), read the pool
   medians (cached) and the sim reference (cached, or queue it and answer `preparing`), run
   `Review`, and store the row.
3. Render for the viewer: the full review for Premium and officers, the count and top title
   for free owners.

### 9.3 `model_version`

`coach_reviews` and `coach_profiles` carry `model_version` = `coach-v1.` + a short hash of the
rules version constant, the curated coach files and the consumables catalogue. It is computed
at startup. Changing a rule, a threshold or any curated list moves the version. Cached reviews
are then recomputed on the next view, and profiles are rewritten by a backfill that re-runs
`Profile` over rated fights. Copy-only changes don't move it. One row per key is kept: a new
version replaces the old.

Reviews are immutable within a version. A review computed on launch day against the sim
reference keeps that basis until the version moves. The "Reviewed {date}" line makes that
honest.

## 10. API

### 10.1 Routes

- `GET /v1/reports/{id}/fights/{n}/coach?player={player_key}`: one player's review of one fight.
- `GET /v1/reports/{id}/coach/night?player={player_key}`: the night line for one character
  across the report's reviewable fights (for Discord and the account line).
- `GET /v1/me/coach/recent`: the signed-in account's last-raid line (§4.5), plus, for Premium,
  the newest 20 stored reviews of its characters (coach history).
- `GET /v1/characters/{region}/{ruleset}/{name}/coach/history`: Premium, owner only. One
  character's stored reviews, newest first, 20 per page with a cursor.

### 10.2 Shapes

Review (full):
```json
{
  "report_id": "…", "fight_index": 3, "player_key": "eu/forever/Kaedra",
  "spec": "warrior-fury", "role": "dps", "size": 40, "kill": true,
  "status": "ready",
  "model_version": "coach-v1.3f9a2c", "reviewed_at": "2026-10-28T09:12:00Z",
  "access": "full",
  "findings": [
    {
      "code": "core_casts", "severity": "major",
      "title": "Bloodthirst: 7 casts lost",
      "detail": "31 of a possible 38 with rage in your bar. Longest gaps 1:12–1:25 and 2:40–2:51.",
      "basis": { "kind": "spell_data", "label": "Measured against Bloodthirst's 6-second cooldown and your own rage bar.",
                 "comparison": { "kind": "top_logs", "n": 25, "label": "Top 25 Fury warriors on Onyxia (40): 9.1 a minute. You: 8.0." } },
      "evidence": [ { "kind": "cast", "spell_id": 23881 }, { "kind": "time", "at_ms": 72000 }, { "kind": "time", "at_ms": 160000 } ],
      "params": { "spell_id": 23881, "actual": 31, "possible": 38 }
    }
  ],
  "could_not_check": [ { "code": "talents_race", "line": "Talents and race: Forever logs don't record them." } ]
}
```
`basis.kind` is one of `top_logs`, `sim`, `spell_data`, `log`. `status` is `ready`,
`preparing` or `not_reviewable` (with a `reason` line). The values in the example are invented.

Free owner (`access: "teaser"`): `{ status, access, finding_count, top: { title } }`. Nothing
else is sent. The gate is on the server, never by hiding fields in the client.

### 10.3 Gating and errors

- Access is checked in this order: report visibility, then the row (§13), then the feature.
  A caller who may not see the report, or may not coach that row, gets the same 404 the rating
  route gives (`not_found`), so no route leaks a report's existence. A flagged report answers
  `flagged`, as the rating route does.
- Free owner: `access: "teaser"`. Premium owner or officer: `full`.
- `night` and `history` without `log_coach` answer the teaser form, or 403 `no_plan` for history.

### 10.4 Entitlement

Add `FeatureLogCoach Feature = "log_coach"` in `api/internal/entitlements/entitlements.go`.
`Standard()` already counts it as a personal Premium feature. It must also be added to
`auth.EntitlementsView` (`log_coach` JSON), to the web's entitlement types and fixtures
(`web/src/lib/account/api.ts`, `web/src/test-support/billing-fixtures.ts`), and to the
launch-resolution test's feature list.

### 10.5 Caching

- The server cache is `coach_reviews` (§9). Responses are `Cache-Control: private, no-store`,
  because every answer depends on who asks.
- Pool medians: in memory per `(encounter, difficulty, size, spec, phase, model_version)`, for 1 hour.
- Sim references: `coach_sim_refs`, per engine version (§7).

## 11. The Discord line and the account line, as built

- Both use the night-line logic behind `GET /v1/reports/{id}/coach/night` in-process: `discordbot` calls the coach
  service, not HTTP. Both render the night rule of §4.4.
- `discordbot/cmd_report.go`: after `reportMessage`, if the viewer holds `log_coach` and owns
  characters in `ans.Raid.Match.ReportID`, append the line (or the fallback link) to the head
  embed. `share.go` renders without it.
- `notify.ReportReady` gains a `CoachReady bool`. `guilds.sendReport` sets it when any owned
  character in the log belongs to an account holding `log_coach`. The embed appends the one
  sentence of §4.4 to the description. No names, no counts.
- Copy for both lives in the existing copy files (`discordbot/copy_*.go`,
  `notify/report_embed.go`'s `reportCopy`), not in the coach package. The coach supplies the
  finding stem.

## 12. Performance and cost

- A cached review is one indexed Postgres read.
- A miss is one bucket read of the fight summary (the same read `sims.Score` does), plus one
  cached pool-median lookup and a pure review in memory: about a few hundred milliseconds,
  dominated by the bucket read.
- A night line on a cold cache is up to one miss per boss fight. Discord never waits for it
  (§4.4), and the account page loads it on its own.
- Sim references cost about 3 s of compute once per (engine version, build, spec, band,
  duration bucket). That is a bounded set, and it is nothing at Cloud Run prices.
- Writing `coach_profiles` adds one ~1 KB insert per rated kill row, inside the rater's
  existing transaction.
- Nothing is computed at ingest except the profile. A review costs only when someone asks for it.

## 13. Privacy

- **Who may ask.** Only a viewer who may see the report under the existing visibility rule
  (public and unlisted: anyone; private: owner and moderators; guild: also verified members).
  The coach needs that rule a third time. Instead of copying it again (rating's `visible`
  already copies `reports.mayView` by hand), export one `reports.MayView(actor, report,
  accounts)` and make rating and the coach both call it.
- **Whose row.** Your own characters (`OwnsCharacter`). Another player's row only as an officer
  of the report's guild holding `officer_views` for that guild (`Store.CanForGuild`), and only
  for a verified member of that guild. Moderators as officers. Everyone else sees nothing about
  the coach on someone else's row: no teaser, no count.
- **Anonymised accounts.** An anonymised player's row is already left out of public rating
  views. The coach is shown to that player on their own row and to their guild's officers only.
  `coach_profiles` stores no names, and comparisons are medians of 20 or more players that never
  name or link anyone.
- **Channels.** Nothing personal is ever posted to a channel (§4.4).
- **Private reports** are reviewable by their owner, and never enter the pool.
- **Retention.** `coach_reviews` rows are deleted with their report (a foreign key with cascade).
  Profiles are removed on the withdrawal path (§6.2).

## 14. Accuracy rules

1. Every number on a finding comes from the log, the client's tables or the pool, through
   `Review`. No number is typed into the copy table.
2. Every finding carries at least one evidence anchor that shows its proof on the report page.
   A finding without one is a bug, and a test refuses it.
3. Every review carries the "Couldn't check" list. Talents and race are on it in v1.
4. A finding fires only when its rule can be proven from the data. Where part of the proof is
   missing (cost unknown, shared debuff, incomplete catalogue category), the finding is dropped
   or narrowed, and the gap is said out loud.
5. Thin basis means no finding: `not enough logs yet`, with the real count.
6. Curated coach files and catalogue entries carry `verified` stamps against the client's own
   tables. Known-good fixtures pin the obvious answers: a Fury warrior's Bloodthirst count on a
   real Forever log, a flask that was there, a death to an avoidable ability. A change that moves
   a fixture fails the build until a person explains why the new answer is right (tenet 8).
7. Before the plans page renders the held line, each clause is checked against the shipped coach
   (§15).

## 15. Left to the builder

- **Order of work:** §8's consumables fix and the rating bump; then coach files for the 27 specs
  (DPS first) and the catalogue's `complete` flags and `raid_buffs`; then `Profile` and
  `coach_profiles` in the rater; then `Review` with fixtures; then the API; then the
  ux-designer specs and the three surfaces; then the plans-page line and the changelog.
- **Migration number:** the next free number after `0054` when the lane lands
  (`migration_order_test.go` guards it). Only the deployed API migrates.
- **Which client build's tables to use for a fight:** the build the log was recorded on when the
  report records it, else the build active at `fought_at`.
- **Evidence anchors:** extend `momentHref` with `cast`, `time` and `resource` kinds, using the
  report page's existing tab and anchor conventions (`cast-key.ts`, the aura and death anchors).
- **The exact copy strings** within the patterns of §4.1, in one copy table.
- **The nightly sim-reference warm:** optional.
- **The held Premium line** (`premium.featuresHeld.coach`). Check each clause: "what to fix
  first" (§4), "the casts you missed" (F1), "the buffs you let drop" (F2), "the consumables you
  skipped" (F6, after §8), "learned from the engine" (§7) and "the top logs on the site" (§6).
  Cut any clause the shipped coach doesn't deliver, and land the line with the changelog entry.
  The coach covers healers and tanks, so the Premium scope note must say so, per B5 (see Owner
  decisions).

## 16. v2

- **Talents** from the synced addon export: talent findings and unspent points.
- **Race**, with it.
- **Personal sim:** the pull simulated with the player's own gear from the log and race and
  talents from the export. The gap is split into gear, buffs and execution. DPS first
  (feasibility §2, `sims.Score`).
- **Movement and positioning** from the advanced-logging fields.
- **More curated mechanics:** avoidable damage and interrupt duty for every raid boss.
- **Trend over weeks:** the same finding across raid nights ("Bloodthirst casts lost: better
  three weeks running").
- **Watch a top log:** a link to a public top log in the pool, once naming and anonymity are settled.
- **Wipe review** for raid leaders (what started the chain), on the guild plan.

---

## 17. Builder check

I re-read this as a builder who has never seen the codebase. These are the questions I would
have asked, with the answers the codebase or the brief gives.

1. *Where do I get the fight data?* From the companion's posted summary JSON in the bucket, read
   with `sims.fightSummary` (`api/internal/sims/summaries.go`). The rater gets the same summary
   at `api/internal/reports/ingest.go` (`rate`).
2. *Is `log_coach` automatically Premium?* `Standard()` returns true for it. But
   `EntitlementsView`, the web types and fixtures, and the test's feature list are explicit
   lists, so add it to each (§10.4).
3. *How do I partition a new table?* Monthly range partitions on `fought_at`, created by
   `db.EnsurePartitions`. Add the table to the partition job's list (§6.2).
4. *What is a "player key", and how do I know the viewer owns it?* `region/ruleset/name`, as in
   `/v1/characters/{region}/{ruleset}/{name}`. Ownership is `auth.Store.OwnsCharacter`.
5. *Who may see a report?* `reports.mayView` (`api/internal/reports/handler.go`). Rating keeps a
   hand copy. Export one and use it in both places (§13).
6. *What is "evidence"?* `presence.EvidenceReport`: not private, not flagged.
7. *How do officers qualify?* `entitlements.Store.CanForGuild(guild, user, FeatureOfficerViews)`.
8. *Can the raid-night post carry a per-player line?* No. R-A3 makes it counts only and names
   nobody. §4.4 resolves this with a sentence that names nobody, and the personal line goes in
   the ephemeral `/report`.
9. *What does `/report` look like today?* `api/internal/discordbot/cmd_report.go`. It is
   ephemeral, with `Share in channel` only for public logs. The coach line is never shared (§11).
10. *Does the account page have a "last raid" line now?* No. The definition already exists for
    the addon (`addon.lastRaidFights`, 14 days, evidence reports). Reuse it (§4.5).
11. *Which specs can the sim reference cover?* DPS specs only (`sims.Score` skips tanks and
    healers), and only those marked `validated` in `sim_specs` (§6.1).
12. *What does "ready with resource" read?* The per-second `ResourceTrack.Series`, and the base
    `cost` / `cost_type` from `data/builds/<build>/spellconst/<class>.json` (§5.2 F1).
13. *How are raid buffs listed?* They aren't in the catalogue today. Add `raid_buffs` with
    provider class and verified ids (§5.2 F6).
14. *How do phases bound the pool?* `data/curated/phases.json` start dates (§6.2).
15. *How do I keep the rating and the coach from measuring differently?* `Profile` and `Review`
    share their measuring code. Idle time reuses `ActivityIdleMS`. Adherence reuses `fightAdherence` (§9.1).
16. *What happens to existing wrong preparation scores?* The model bump re-rates every stored
    fight, and the digests are rebuilt (§8).
17. *Is there a design step before building the surfaces?* Yes. The ux-designer writes the
    experience spec and mock, and two SHIP verdicts on screenshots come before merge (tenet 6).

## Owner decisions

1. **The held Premium line says "after every raid".** The coach needs a log of the raid. Keep
   the line as drafted, or change it to "after every logged raid"? (B5 lets the builder cut a
   clause but not reword it.)
2. **The Premium scope note.** The coach covers every spec, so B5 requires the note to change
   when the line renders. Proposed: `Sims cover damage specs; healing and tanking specs are not
   simulated yet. The coach reviews every spec; damage specs get the most findings.` Approve, or
   give the wording.
3. **Who counts as a "top" healer or tank.** I decided on the rating's overall score (HPS rewards
   overheal, and damage taken rewards nothing). You may want healers ranked by something else,
   such as effective HPS. Confirm or override.
