# Forever Sixty Addon v2 Implementation Plan (lanes)

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking. Each lane is its own worktree and branch; lanes merge to `main` at their week gate.

**Goal:** Replace the six-page addon window with marks on the client's own surfaces, a HUD chip, a character-panel tab, guild and group-leader tools, and the Forever Score, fed end to end by the pipeline, the sim, the API, the companion and the site, in 30 days.

**Architecture:** The addon keeps its model layer and gains four new modules (`Rules`, `Progress`, `Score`, `Comms`) that every surface reads; one file per surface; no surface computes a number. The pipeline ships default builds, drop chances, DPS-per-point weights, encounter needs and ability levels in `Data.lua`. The API computes the Forever Score, writes the nightly data addon, and serves inbox v2 messages; the companion carries the outbox up and the inbox down; the site shows the live score, profile and "Copy for the game" buttons.

**Tech Stack:** Lua 5.1 (addon, luacheck, busted for pure functions only), Python 3.12 (`data/pipeline`, pytest), Go (`sim`, `api`, `companion`; `go test`), Astro + Svelte + TypeScript (`web`; vitest, playwright), Cloud Run jobs, WoW 1.60 beta client for every addon verification.

**Spec:** `docs/superpowers/specs/2026-10-11-addon-v2-design.md`

## Global Constraints

- Addon verification is in the real client, never the Lua mock: every surface task ends with a capture checklist run on the 1.60 beta client, reviewed by the `ux-designer` and `wow-player` agents on screenshots. `addon/tests/wow_mock.lua` stays only for the legacy specs already green; no new spec may depend on it.
- busted is used only for pure functions (`Rules`, `Progress`, `Score`, `Comms` encode/decode, codec); `luacheck ForeverSixty tests` is zero warnings; `luac -p` under Lua 5.1 passes for every shipped file.
- No network call from the addon; no load-time frames; idle memory under 2 MB; every client API behind a capability check in `Theme`/`Compat` that degrades, never errors.
- Every number comes from one function (spec §2.1); a surface with nothing to say is absent (§2.2); a surface that fails `selfCheck()` disables itself and writes one line to `/fs diag` (§2.4).
- Leader tools inform only: no kick, decline, whisper or announcement is ever sent by the addon (§2.6).
- Copy rules: plain words, no spell or item ids on screen, no build numbers, no "data addon" or "companion" without a four-word explanation beside it, dates as "as of 4 Nov".
- Tokens: bg `#0d111a`, border `#262e40`, gold `#e5b955`, muted `#9a9484`, body `#e9e4d8`, warning `#ff6b5c` (errors only), success `#1eff00`; quality colours from the client; class colours from `RAID_CLASS_COLORS`.
- Data keyed by client build id; `python -m pipeline addon-data --check` is the drift gate; `Data.lua` is committed.
- Message protocol prefix `FS1`, version field first, unknown versions ignored silently.
- Score formula lives in the API (`api/internal/rating`); the addon displays and dates only.

## Review Focus

1. A character the site has never logged, hovered in a raid: the player tooltip must show "no logs yet" and the group panel row must show inspect data with "no score", never a 0. Pinned in Lane 2, `Score.lua` tests, and Lane 5's capture checklist.
2. A pasted build authored for an older client build: the dialog must refuse with the codec's reason and keep the current build; the chip must not change state. Pinned in Lane 3's `BuildDialog` capture checklist and Lane 2's `Progress` tests (nil build → zero progress, no next point).
3. Addon messages from a newer protocol version or a malformed payload: `Comms.decode` must return nil and count it, never raise inside the event handler. Pinned in Lane 2's `Comms` tests with a fuzz list of 50 payloads.
4. A level-20 shaman hovering a mail chest, a rogue hovering a mace, an Undead paladin reading a plan: `Rules` must answer "until 40", "yes" and "yes", and the tooltip must never say "upgrade" for an item the character cannot equip. Pinned in Lane 2's `Rules` fixtures and Lane 3's tooltip captures.
5. Combat lockdown: the chip's equip action, the roll mark and the group panel's inspect queue must all stand down in combat and resume on `PLAYER_REGEN_ENABLED`. Pinned in Lane 3 and Lane 5 capture checklists with a combat capture.

---

## Lanes, owners and gates

| Lane | Branch | Depends on | Week gate |
|---|---|---|---|
| L0 Contracts | `lane/contracts` | none | week 1, day 2: merged before any other lane starts coding against it |
| L1 Pipeline and sim data | `lane/data-v2` | L0 | week 1 |
| L2 Addon core modules | `lane/addon-core` | L0 | week 1 |
| L3 Addon surfaces A: chip, talent marks, item tooltip, build dialog, settings | `lane/addon-surfaces-a` | L1, L2 | week 1 |
| L4 Addon surfaces B: slot marks, reward and roll marks, panel tab, player tooltip, chat badges | `lane/addon-surfaces-b` | L2, L3 | week 2 |
| L5 Addon surfaces C: guild tab, group panel, ready check, loot council, last raid | `lane/addon-surfaces-c` | L2, L6 (inbox v2), L7 | week 3 |
| L6 API: score v2, nightly data v2, inbox v2, outbox ingestion, character linking | `lane/api-v2` | L0 | week 2 (score, inbox), week 3 (loot values, last raid) |
| L7 Companion: outbox up, inbox v2 down, ack | `lane/companion-v2` | L0, L6 | week 2 |
| L8 Website: score profile, "Copy for the game", in-game panel, addon page, guild roster | `lane/web-v2` | L6 | week 2 (buttons, profile), week 3 (guild) |
| L9 Release and QA: weekly capture loop, performance pass, CurseForge and Wago, docs | `lane/release-v2` | all | every week; final gate week 4 |

Each lane: `git worktree add ../fs-<lane> -b lane/<name>`, commits per task, PR to `main` at its gate with both persona verdicts attached for addon lanes.

---

## Lane L0: Contracts

### Task L0.1: Write the shared contract document

**Files:**
- Create: `docs/contracts/addon-v2.md`
- Modify: `docs/superpowers/specs/2026-09-14-phase-3-interfaces.md` (append a pointer line to the new file)

**Interfaces:**
- Produces: the five contracts every other lane codes against, verbatim:

1. `Data.lua` additions (Lane L1 writes, Lane L2 reads):

```lua
ns.Data = {
  build = "1.60.1.70291",
  classes = { ... },            -- unchanged
  weights = { ... },            -- unchanged
  -- NEW: DPS (or HPS for healers) per 1.0 of the spec's reference stat, from the site's sim
  dpsPerPoint = { ["warrior-arms"] = { reference = "attack_power", dps = 0.42, healing = false }, ... },
  -- NEW: default builds per spec, as the codec's own triples (tab, tier, column), 1-based
  defaultBuilds = {
    ["warrior-arms"] = {
      leveling = { name = "Arms leveling", order = { {1,1,1}, {1,1,1}, ... } },  -- 51 triples
      raid     = { name = "Arms raid",     order = { ... } },
    }, ...
  },
  -- NEW: per BiS pick, the drop chance and two alternatives (addonbis renders these beside each pick)
  --   bis[spec][band][faction][slot] = { item_id, kind, source, chance = 0.14 | nil,
  --                                      alternatives = { { item_id, kind, source }, { ... } } }
  -- NEW: what each raid encounter needs, for the ready-check coverage list
  encounterNeeds = { [<npcId>] = { buffs = { 1126, ... }, debuffs = { 7386, ... }, resist = "fire" | nil } },
  -- NEW: the level each rotation ability is learned, so the rotation region can say "at level N"
  abilityLevels = { [<spellId>] = 24, ... },
  rotations = { ... }, bis = { ... },        -- unchanged shapes plus the fields above
}
```

2. FS1 addon-message protocol (Lane L2 implements, L4/L5 use):

```
prefix "FS1"; payload = "<v>|<type>|<fields...>", fields "|"-separated, values URL-safe (no "|", "\n")
v = 1
hello     | level | specSlug | gearGap(int) | buildName | scoreDate(YYYY-MM-DD or "-")
share     | buildCode
push      | buildCode | buildName           (sender must be guild rank index <= 1 or raid leader; receivers verify with GetGuildInfo / UnitIsGroupLeader)
loot?     | itemId | lootId
loot!     | lootId | value(float, DPS or score) | unit("dps"|"score") | currentItemId | planned(0|1)
snapshot? | (no fields)
snapshot! | same fields as hello
```
Channels: GUILD for hello/share/push/snapshot; RAID or PARTY for loot and snapshot. Rate: at most one `hello` per 60 s per sender; `loot!` only in reply to a `loot?` seen in the last 30 s.

3. Companion inbox v2 (`ForeverSixtyInbox`, Lane L6 emits, L7 writes, L2/L4/L5 read):

```lua
ForeverSixtyInbox = {
  format = 2, generated_at = "2026-11-04T21:03:00Z",
  messages = {
    { type = "build",      character = "Obnoxious-Yell", name = "Arms raid", code = "FSB1:..." },
    { type = "upgrade",    character = ..., slot = "chest", item_id = 6627, item_name = ..., source = ..., delta = 4.2, unit = "dps" },
    { type = "weights",    character = ..., spec = "warrior-arms", weights = { {stat="attack_power", weight=1}, ... }, caps = {"hit"}, dps_per_point = 0.42 },
    { type = "sim_result", character = ..., spec = ..., dps = 412.6, error = 3.1, generated_at = ..., slots = { {slot="chest", item_id=6627, delta=4.2}, ... } },
    { type = "last_raid",  character = ..., report_id = "r_...", date = "2026-12-09", bosses = { {name="Lucifron", percentile=71, biggest_loss="downtime", loss_seconds=18, deaths=0, consumables_missed={"flask"}, adherence=0.82}, ... } },
    { type = "loot_values", character = ..., raid = "Molten Core", items = { [17076] = { value = 6.3, unit = "dps", planned = true }, ... } },
    { type = "guild",      character = ..., guild_name = ..., claim_state = ..., pending_approvals = 2, rank = "officer" },
    { type = "ack",        character = ..., export_seen_at = "2026-11-04T14:03:00Z" },
  },
}
```

4. Companion outbox (`ForeverSixtyDB.outbox`, Lane L2 writes on `PLAYER_LOGOUT`, L7 reads):

```lua
ForeverSixtyDB.outbox = {
  format = 1,
  export = { character = "Obnoxious-Yell", code = "FS1:...", saved_at = "2026-11-04 14:02" },
  sim_request = { character = ..., spec = "warrior-arms", requested_at = ..., gear_code = "FS1:..." } | nil,
  roster_snapshot = { guild = "OLYMPUS XXVII", members = { { name=..., level=20, spec="warrior-arms", gear_gap=14, build="Arms leveling", seen_at=... }, ... } } | nil,
}
```

5. Nightly data addon v2 (`ForeverSixtyData`, Lane L6 writes, L2 `Score` reads):

```lua
ForeverSixtyData = {
  format = 2, generated = "2026-11-04",
  characters = {
    ["Obnoxious-Yell"] = { score = 1830, tier = "epic", as_of = "2026-11-04",
      components = { gear = 412, experience = 620, parses = 540, attendance = 180, discipline = 78 },
      gear_score = 412, item_level = 31, spec = "warrior-arms",
      experience = { dungeons = 14, raid_kills = 0, alts = 2 },
      best_parses = { ["Lucifron"] = 71 }, attendance = 0.9, fights = 12, profile = "/character/us/forever/obnoxious-yell" },
  },
  guilds = { ["OLYMPUS XXVII"] = { standing = 3, nights = 4, roster = 31, attendance = 0.8, members = { "Obnoxious-Yell", ... } } },
}
```
Tiers by score: `common` < 500, `uncommon` < 1000, `rare` < 1500, `epic` < 2000, `legendary` ≥ 2000 (final thresholds set by L6 from the distribution; the addon reads `tier`, never computes it).

- [ ] **Step 1: Write `docs/contracts/addon-v2.md`** with the five blocks above verbatim, each with a one-paragraph "who writes, who reads, how versioned" note.
- [ ] **Step 2: Add the pointer line** to the phase-3 interfaces spec: `See docs/contracts/addon-v2.md for the addon v2 wire contracts (2026-10-11).`
- [ ] **Step 3: Review** with one `ux-designer` pass on copy fields (names shown to players: `buildName`, `biggest_loss`, `source`) and fix wording inline.
- [ ] **Step 4: Commit** `git add docs/contracts/addon-v2.md docs/superpowers/specs/2026-09-14-phase-3-interfaces.md && git commit -m "docs: addon v2 wire contracts"`

### Task L0.2: Pin the contracts with one golden test per side

**Files:**
- Create: `addon/tests/contracts_spec.lua` (pure Lua, no mock: loads the Lua literals from the doc's fixtures)
- Create: `api/internal/addon/contracts_test.go`
- Create: `companion/internal/addon/contracts_test.go`
- Create: `docs/contracts/fixtures/inbox-v2.json`, `outbox-v1.json`, `data-addon-v2.lua`, `fs1-messages.txt`

**Interfaces:**
- Produces: the fixture files every lane's tests load; a change to a fixture fails all three sides until each is updated.

- [ ] **Step 1: Write the fixtures** from L0.1's blocks (one message of each type, one character, one guild, every FS1 message type).
- [ ] **Step 2: Addon golden test** (`contracts_spec.lua`): `loadfile("docs/contracts/fixtures/data-addon-v2.lua")` parses, `format == 2`, every character has `score`, `tier`, `as_of`; every inbox message type in the JSON fixture appears in `L.inboxTypes` (added in L2).
- [ ] **Step 3: API golden test**: `json.Unmarshal` the inbox fixture into `addon.MessageEntry` (extended in L6.3) with `DisallowUnknownFields`; marshal back; byte-equal after normalisation.
- [ ] **Step 4: Companion golden test**: decode the same fixture with `companion/internal/addon.Message`; render to Lua through `lua.go`; the rendered chunk `loadstring`s under a Lua 5.1 interpreter in CI (`lua5.1 -e "assert(loadfile(...))"`).
- [ ] **Step 5: Run** `cd addon && busted tests/contracts_spec.lua`, `cd api && go test ./internal/addon/`, `cd companion && go test ./internal/addon/`. Expected: all pass.
- [ ] **Step 6: Commit** `git commit -am "test: golden contract fixtures for addon v2"`

---

## Lane L1: Pipeline and sim data

### Task L1.1: Default builds per spec from the sim's guide builds

**Files:**
- Create: `sim/cmd/addon-builds/main.go` (reads every spec's guide leveling and raid builds, emits `data/builds/<build>/addon-builds.json`)
- Modify: `sim/leveling/talents.go` (export `LadderPointOrder`)
- Create: `sim/leveling/point_order.go`, `sim/leveling/point_order_test.go`
- Modify: `data/pipeline/addondata.py` (read `addon-builds.json` into `AddonData.default_builds`), `data/pipeline/models.py` (`AddonDefaultBuild`), `data/pipeline/addonlua.py` (`_default_builds_lua`), `data/tests/test_addondata.py`

**Interfaces:**
- Produces (Go): `func LadderPointOrder(activeTrees []TalentTree, targets map[int]int, ownTreeIndex int) ([]Point, error)` where `type Point struct{ Tab, Tier, Column int }` (1-based, the codec's own triple), 51 points for a 60 build, derived as the per-level diff of `LadderTalentString(…, level)` for level 10..60; an error if any level's diff is not exactly one point.
- Produces (JSON): `addon-builds.json` = `{ "<spec>": { "leveling": { "name": "Arms leveling", "order": [[1,1,1], ...] }, "raid": { ... } } }`.
- Produces (Lua): `ns.Data.defaultBuilds` per L0 contract 1.

- [ ] **Step 1: Write the failing Go test** in `point_order_test.go`: for the warrior trees in `data/builds/<active>/talents/warrior.json` and the Arms guide targets, `LadderPointOrder` returns 51 points, point 1 is in tab 1 tier 1, and replaying the order through `TalentRanksFromString` at level 60 equals the guide's own ranks.
- [ ] **Step 2: Run** `cd sim && go test ./leveling/ -run PointOrder`. Expected: FAIL, undefined `LadderPointOrder`.
- [ ] **Step 3: Implement** `point_order.go`: loop level 10..60, call `LadderTalentString` for level-1 and level, diff the digit strings per tree, map each changed digit index to (tab = tree position+1, tier, column) from the tree's talents, append; return an error naming the level if the diff is not one point.
- [ ] **Step 4: Run** the test. Expected: PASS.
- [ ] **Step 5: Write `sim/cmd/addon-builds/main.go`**: for every spec in `sim/specs`, read `GuideBuildTalents` and `GuideRaidBuildTalents`, `RequireGuideBuildMatchesActive`, compute both orders, write `addon-builds.json` with names `"<Spec> leveling"` and `"<Spec> raid"`; exit non-zero listing specs with no guide.
- [ ] **Step 6: Pipeline**: add `AddonDefaultBuild(BaseModel): name: str; order: list[tuple[int,int,int]]` and `AddonData.default_builds: dict[str, dict[str, AddonDefaultBuild]]`; read the JSON in `build_addon_data`; render in `addonlua.py` as nested tables of `{t,r,c}` triples; test in `test_addondata.py` that every spec present in `weights` has a leveling build with 51 points and that the Lua renders and parses.
- [ ] **Step 7: Run** `cd data && uv run pytest tests/test_addondata.py` and `uv run python -m pipeline addon-data --build <active> && uv run python -m pipeline addon-data --check`. Expected: pass, `Data.lua` regenerated.
- [ ] **Step 8: Commit** both repos' changes: `git commit -am "feat(data): default leveling and raid builds per spec from the sim's guide builds"`

### Task L1.2: Drop chance and alternatives per BiS pick

**Files:**
- Modify: `data/pipeline/addonbis.py` (`_one_band`: carry `chance` and the two next-ranked candidates as `alternatives`), `data/pipeline/models.py` (`AddonBisItem.chance: float | None`, `alternatives: list[AddonBisAlternative]`), `data/pipeline/addonlua.py` (`_bis_lua`: pick becomes `{ id, "D", "source", 0.14, { {id,"Q","source"}, {...} } }`), `data/tests/test_addonbis.py`
- Modify: `sim/cmd/leveling-bis/report.go` (emit `chance` and `runners_up` per slot in the JSON the pipeline reads; the loot tables already feed `pipeline/loot`)

**Interfaces:**
- Produces: `ns.Data.bis[spec][band][faction][slot] = { item_id, kind, source, chance, alternatives }` per contract 1; `Gear.pick(spec, band, faction, slot)` in L2 reads it.

- [ ] **Step 1: Failing pytest**: a band fixture with `runners_up` renders a Lua pick table whose fifth field is a table of two entries and whose fourth is `0.14`; a pick with no chance renders `nil` in the fourth field.
- [ ] **Step 2: Run** `uv run pytest tests/test_addonbis.py -k chance`. Expected: FAIL.
- [ ] **Step 3: Implement** the model fields, the reader and the renderer; `report.go` writes `chance` from the loot table's drop rate (`pipeline/loot` already has it per item and source) and `runners_up` as the next two by score.
- [ ] **Step 4: Run** the tests and `addon-data --check`. Expected: pass.
- [ ] **Step 5: Commit** `git commit -am "feat(data): drop chance and two alternatives per BiS pick"`

### Task L1.3: DPS per weight point, encounter needs, ability levels

**Files:**
- Modify: `api/internal/sims/score.go` (expose `DPSPerPoint(result) float64`: the sim's own DPS derivative along the reference stat, already computed for weights), `sim/cmd/forever-sim` (flag `-dps-per-point` writes `data/curated/dps_per_point.json` per spec from the guide build's sim)
- Create: `data/curated/encounter_needs.json` (per raid boss npc id: buffs, debuffs, resist), `data/pipeline/encounters.py` (loader + validation against `dungeons.json` ids)
- Modify: `data/pipeline/addonrotation.py` (collect `ability_levels` from `spellranks`: the first rank's learn level per spell family), `data/pipeline/addondata.py`, `models.py`, `addonlua.py`, `data/tests/test_addondata.py`

**Interfaces:**
- Produces: `ns.Data.dpsPerPoint`, `ns.Data.encounterNeeds`, `ns.Data.abilityLevels` per contract 1.

- [ ] **Step 1: Failing pytest**: `Data.lua` renders `dpsPerPoint["warrior-arms"].dps > 0` with `reference == "attack_power"`; `abilityLevels[<Execute spellId>] == 24`; `encounterNeeds` has an entry for Lucifron's npc id with a non-empty `debuffs` list.
- [ ] **Step 2: Run** the pytest. Expected: FAIL.
- [ ] **Step 3: Implement** the three readers and renderers; `encounters.py` refuses an npc id not in the build's `dungeons.json` and a spell id not in `spellranks`.
- [ ] **Step 4: Run** `uv run pytest tests/` and `addon-data --check`. Expected: pass.
- [ ] **Step 5: Commit** `git commit -am "feat(data): dps per point, encounter needs, ability levels for the addon"`

---

## Lane L2: Addon core modules

### Task L2.1: `Rules.lua`, the Forever rules engine

**Files:**
- Create: `addon/ForeverSixty/Rules.lua`, `addon/tests/rules_spec.lua` (pure Lua)
- Modify: `addon/ForeverSixty/ForeverSixty.toc` (load after `Compat.lua`), `Locale.lua` (`rulesUntilLevel = "Not for you until %d: %s"`, `rulesNever = "Not for you: %s"`)

**Interfaces:**
- Produces:
```lua
Rules.ARMOR_AT = { WARRIOR = { mail = 1, plate = 40 }, PALADIN = { mail = 1, plate = 40 }, HUNTER = { leather = 1, mail = 40 }, SHAMAN = { leather = 1, mail = 40 }, ROGUE = { leather = 1 }, DRUID = { leather = 1 }, PRIEST = { cloth = 1 }, MAGE = { cloth = 1 }, WARLOCK = { cloth = 1 } }
Rules.WEAPONS = { SHAMAN = { axe1=true, axe2=true, mace1=true, mace2=true, fist=true, staff=true, dagger=true, shield=true, offhand=true, dualWield=false }, ROGUE = { dagger=true, sword1=true, mace1=true, fist=true, dualWield=true }, ... }
--- @return "ok" | "until" | "never", level|nil, reason string|nil
function Rules.canEquip(classToken, level, itemClass, itemSubClass, equipLoc)
--- hit and crit are one stat in Forever; the key the weights table uses
function Rules.statKey(itemStatKey) -- "ITEM_MOD_CRIT_RATING_SHORT" -> "crit", "ITEM_MOD_HIT_RATING_SHORT" -> "hit"
function Rules.racePairAllowed(raceToken, classToken) -- the six new pairs plus every vanilla pair
function Rules.questRewardFloor(questLevel) -- questLevel - 3
```
- Consumes: `Compat.itemInfoInstant(link)` for itemClass/subClass/equipLoc.

- [ ] **Step 1: Failing spec** `rules_spec.lua` (no mock; `require("Rules")` directly): shaman level 20 + mail chest → `"until", 40`; shaman + second one-hand weapon → `"never"`; rogue + mace1 → `"ok"`; warrior level 39 + plate → `"until", 40`; Undead + PALADIN → `true`; Tauren + MAGE → `false`; `questRewardFloor(23) == 20`.
- [ ] **Step 2: Run** `cd addon && busted tests/rules_spec.lua`. Expected: FAIL, module not found.
- [ ] **Step 3: Implement** the tables and functions; `canEquip` maps `equipLoc` and the client's `itemSubClass` ids through a small table, returns `"never"` for a dual-wield second weapon on a shaman.
- [ ] **Step 4: Run** the spec and `luacheck ForeverSixty`. Expected: pass, zero warnings.
- [ ] **Step 5: In-client check**: `/fs diag` prints nothing new after `/reload`; hovering a mail item on the level-20 warrior shows no refusal (plate is 40, mail is 1). Capture one tooltip.
- [ ] **Step 6: Commit** `git commit -am "feat(addon): Forever rules engine"`

### Task L2.2: `Progress.lua`, one function per number

**Files:**
- Create: `addon/ForeverSixty/Progress.lua`, `addon/tests/progress_spec.lua` (pure: builds and ranks as plain tables)
- Modify: `Follow.lua` (move `Follow.progress`, `Follow.pointsSpent`, `Follow.nextPoint` bodies here; `Follow` keeps thin wrappers), `Gear.lua` (`Gear.slotVerdicts` moves here), `GearView.lua` functions `slotRow`, `noteFor` are deleted with the view in L3.5

**Interfaces:**
- Produces:
```lua
--- @return { matched, total, offPlan, point = {tab,tier,column,index}|nil, nextLevel = 9+index|nil, done }
function Progress.build(build, ranks)
--- @return { slot, plannedItemId, equippedItemId, delta (worn - planned, nil without stats), state = "match"|"better"|"behind"|"empty"|"noplan", unit = "dps"|"score" }
function Progress.slotVerdict(slot, plan, worn, weights, dpsPerPoint)   -- same id → "match" always
--- @return array of slotVerdict for every slot the plan or the paper doll names, paper-doll order, weapons first for melee specs
function Progress.slots(data, build, ranks, equipped, dpsPerPoint)
--- @return { slot, delta, item } for the slot furthest behind the plan, or nil
function Progress.biggestGap(slots)
--- @return array of { link, itemId, slot, delta } bag items beating the worn item, sorted by delta desc; worn items never listed
function Progress.bagUpgrades(data, build, ranks, bagLinks, equipped, dpsPerPoint)
--- @return { level = N, text = "Take at level N" }
function Progress.levelForPoint(index)
```
- Consumes: `Talents.cellKey`, `Gear.score`, `Gear.statsOf`, `Rules.statKey`.

- [ ] **Step 1: Failing spec** `progress_spec.lua`: the 51-point warrior fixture with ranks {IHS 2, Rend 3, Overpower 2, Cruelty 4} → `matched == 11`, `offPlan == 0`, `point.index == 3`, `nextLevel == 12`; same item id on both sides with worse plan stats → `"match"`; worn nil with a planned stat line → `"empty"` with `delta == -planScore`; `biggestGap` returns the empty ring over a `"behind"` chest when its delta is lower; `bagUpgrades` never returns a link that is in `equipped`.
- [ ] **Step 2: Run** `busted tests/progress_spec.lua`. Expected: FAIL.
- [ ] **Step 3: Implement**, moving the existing bodies; convert deltas to DPS when `dpsPerPoint` is given (`delta * dps`), keep `unit` on every verdict.
- [ ] **Step 4: Run** the new spec plus the whole legacy suite `busted`. Expected: all pass (legacy specs call the `Follow`/`Gear` wrappers).
- [ ] **Step 5: Commit** `git commit -am "feat(addon): Progress, the one function per number"`

### Task L2.3: `Score.lua`, reading the nightly record

**Files:**
- Create: `addon/ForeverSixty/Score.lua`, `addon/tests/score_spec.lua`
- Modify: `Ratings.lua` (becomes a thin adapter over `Score` for the legacy tooltip line; delete after L4.4)

**Interfaces:**
- Produces:
```lua
--- @return nil when the data addon is absent, unreadable or newer than Score.FORMAT (and why, via Theme.note once)
function Score.status()                         -- { available, generated, stale, reason }
--- @return nil for a character the record does not have (callers print L.scoreNoLogs), else the record plus `asOfText` ("as of 4 Nov") and `tierColor`
function Score.forCharacter(name, realm)
function Score.lines(record)                    -- tooltip lines: "Forever Score 1830 · Epic · as of 4 Nov", then components
function Score.guild(name)                      -- the guild record or nil
Score.FORMAT = 2
```

- [ ] **Step 1: Failing spec**: with `_G.ForeverSixtyData` set to the L0 fixture, `forCharacter("Obnoxious", "Yell")` returns `score == 1830`, `tier == "epic"`, `asOfText == "as of 4 Nov"`; an unknown name returns nil; `format = 3` makes `status().available == false` with `reason == L.ratingsUpdateAddon`; a record with `score = 0` and `fights = 0` returns nil (never a zero).
- [ ] **Step 2: Run**. Expected: FAIL.
- [ ] **Step 3: Implement**; dates through `Compat.shortDate("2026-11-04") == "4 Nov"`.
- [ ] **Step 4: Run** the spec and the suite. Expected: pass.
- [ ] **Step 5: Commit** `git commit -am "feat(addon): Score reader for the nightly record"`

### Task L2.4: `Comms.lua`, the FS1 protocol

**Files:**
- Create: `addon/ForeverSixty/Comms.lua`, `addon/tests/comms_spec.lua` (pure: `encode`/`decode` only; sending is verified in-client)
- Modify: `Options.lua` (register `CHAT_MSG_ADDON`, `C_ChatInfo.RegisterAddonMessagePrefix("FS1")` at login behind a capability check), `Theme.lua` (`Theme.sendAddonMessage(prefix, text, channel)` guarded)

**Interfaces:**
- Produces:
```lua
Comms.PREFIX, Comms.VERSION = "FS1", 1
function Comms.encode(kind, fields)             -- "1|hello|20|warrior-arms|14|Arms leveling|2026-11-04"; escapes "|" as "%7C"
function Comms.decode(text)                     -- { version, kind, fields = {...} } or nil; nil for other versions, malformed text, unknown kind
function Comms.send(kind, fields, channel)      -- channel "GUILD"|"RAID"|"PARTY"; rate-limited per kind (hello 60 s); returns true when sent
function Comms.on(kind, handler)                -- handler(fields, sender, channel)
function Comms.hello()                          -- builds the hello fields from Progress/Score; called at login and on TRAIT_CONFIG_UPDATED
Comms.seen = {}                                 -- sender -> last hello fields + time, read by GuildTab/GroupPanel
Comms.dropped = 0                               -- malformed or foreign-version count for /fs diag
```

- [ ] **Step 1: Failing spec**: round trip of every kind in `docs/contracts/fixtures/fs1-messages.txt`; a field containing `|` survives; `decode("2|hello|...")` is nil and increments `dropped`; a fuzz list of 50 malformed strings never raises (`pcall` wrapper in the spec asserts `ok`).
- [ ] **Step 2: Run**. Expected: FAIL.
- [ ] **Step 3: Implement**; `push` handlers verify sender rank (`GetGuildInfo` rank index ≤ 1) or `UnitIsGroupLeader`, else ignore and count.
- [ ] **Step 4: Run** spec and suite. Expected: pass.
- [ ] **Step 5: In-client check** with two characters in the same guild on the beta: `/fs hello` (a diag subcommand added here) sends; the other client's `/fs diag` lists the sender in `Comms.seen`. Capture both chats.
- [ ] **Step 6: Commit** `git commit -am "feat(addon): FS1 addon-message protocol"`

### Task L2.5: `Inspect.lua`, the inspect queue

**Files:**
- Create: `addon/ForeverSixty/Inspect.lua`
- Modify: `Options.lua` (register `INSPECT_READY`, `GROUP_ROSTER_UPDATE`)

**Interfaces:**
- Produces:
```lua
function Inspect.request(unit)                  -- queues; one NotifyInspect at a time; skipped in combat and for units out of range
function Inspect.get(name)                      -- { spec, level, itemLevel, items = { [slot] = link }, enchants = { [slot] = bool }, at = GetTime() } or nil
function Inspect.onReady(guid)                  -- reads GetInventoryItemLink(unit, slot) for every slot, talents via C_Traits when the inspected config is exposed, else spec from the class's highest tree
Inspect.TTL = 300
```

- [ ] **Step 1: Implement** with a 2 s spacing between requests and a `PLAYER_REGEN_DISABLED` pause; no spec (client-only behaviour).
- [ ] **Step 2: In-client check**: in a party of two on the beta, `/fs inspect` (diag subcommand) prints the partner's item level and missing enchants within 5 s; in combat it prints "paused".
- [ ] **Step 3: Commit** `git commit -am "feat(addon): inspect queue"`

### Task L2.6: Remove the window; outbox; wiring

**Files:**
- Delete: `Window.lua`, `Cards.lua` navigation functions, `views/OverviewView.lua`, `views/GuildView.lua`, `views/ExportView.lua`, `views/SettingsView.lua` (the Settings page is rebuilt in L3.5), `views/FollowView.lua`, `views/GearView.lua`, their specs
- Modify: `Options.lua` (`/fs` opens the build dialog; `/fs diag`, `/fs hello`, `/fs inspect`, `/fs record` stay; `PLAYER_LOGOUT` writes `ForeverSixtyDB.outbox` per contract 4), `Export.lua` (`Export.saveOutbox(data)`), `Prefs.lua` (new keys: `chip.shown`, `chip.combatHide`, `badges`, `broadcast`), `ForeverSixty.toc`
- Keep: `Tracker.lua` deleted after L3.1 ships the chip (the chip replaces it)

- [ ] **Step 1: Delete** the files and their specs; fix `luacheck`.
- [ ] **Step 2: Outbox**: on logout write `export`, pending `sim_request` (set by L4.3's "Sim me"), and `roster_snapshot` from `Comms.seen` when in a guild.
- [ ] **Step 3: Run** `busted` (remaining specs) and `luacheck`. Expected: pass.
- [ ] **Step 4: In-client check**: `/reload` shows no Lua error with `scriptErrors 1`; `/fs` opens nothing yet (L3.4 adds the dialog) and prints `L.slashHint`; logout and `WTF/.../ForeverSixty.lua` contains `outbox` with `format = 1` and an `export`.
- [ ] **Step 5: Commit** `git commit -am "refactor(addon): remove the six-page window; outbox on logout"`

---

## Lane L3: Addon surfaces A (week 1)

Every task here ends with a capture checklist; the captures go to the week's review folder and both personas review them before the lane's PR.

### Task L3.1: The chip

**Files:**
- Create: `addon/ForeverSixty/surfaces/Chip.lua`
- Modify: `Theme.lua` (`Theme.SIZES.chip = 28`, `chipWide = 220`, `chipText = 12`), `Locale.lua` (`chipTake = "Take %s"`, `chipEquip = "Equip %s, %s"`, `chipHere = "Here for you: %s%s"`, `chipHereMore = ", %d more"`, `chipRaidCheck = "Missing: %s"`, `chipLastNight = "%s: %dth percentile, biggest loss %s"`, `chipComplete = "Build complete"`, `chipInCombat = "In combat"`, `chipHideUntilLevel = "Hide until next level"`), `Options.lua` (events: `PLAYER_LEVEL_UP`, `TRAIT_CONFIG_UPDATED`, `BAG_UPDATE_DELAYED`, `ZONE_CHANGED_NEW_AREA`, `READY_CHECK`, `PLAYER_REGEN_DISABLED/ENABLED` all call `Chip.refresh`), `Prefs.lua`
- Delete: `Tracker.lua`, `Toast.lua` (the toast's `unspentPoints` moves into `Progress`)

**Interfaces:**
- Produces: `Chip.state()` → `{ kind = "point"|"upgrade"|"here"|"raidcheck"|"lastnight"|"complete"|"idle", icon, text, onClick, waiting = N }`; `Chip.refresh()`; `Chip.selfCheck()`.
- Consumes: `Progress.build`, `Progress.bagUpgrades`, `Progress.levelForPoint`, `Gear.picksHere(instanceId)` (L4.2), `Score`, inbox `last_raid`.

- [ ] **Step 1: Build the frame**: a 28×28 `Button` on `UIParent`, strata `MEDIUM`, default anchor `TOPRIGHT` of `Minimap` offset (-(Minimap width + 12), 0), saved position in `Prefs.chip`; a 20 px icon, a 1-line `FontString` (12 px, measured with `GetStringWidth`, truncated with "…" until it fits 220 - 36), the gold ring from `Theme.outline` with the pulse from `Theme.showGlow`.
- [ ] **Step 2: State machine** exactly as spec §4.1, priority order; `waiting` badge when more than one state is true; combat: `SetAlpha(0.25)`, `EnableMouse(false)`; `OnClick` per state (`point` → `ToggleTalentFrame()` or `PlayerSpellsUtil.TogglePlayerSpellsFrame()` whichever exists; `upgrade` → `Theme.equip(link)` unless `Theme.inCombat()`; `here` → opens the panel tab's targets region (L4.3, until then prints the list); `lastnight` → the panel tab's last-raid region).
- [ ] **Step 3: selfCheck**: after `refresh`, `text:GetStringWidth() <= wide - 36` and the frame is on screen (`GetLeft() >= 0`); on failure `Theme.note(L.diagChipOff)` and `Hide()`.
- [ ] **Step 4: Capture checklist** (level 20 warrior, beta): idle state; point state (respec one point with `/run` disabled, so use the level-up path: a level-9 alt dinging 10 on the beta, or load a build whose first cell is unmet); upgrade state (put a better bag item on: lower the plan's chest stats in the pasted code); here state (enter The Deadmines); combat (attack a dummy): alpha 25 %; right-click hide; `/fs diag` empty. Ten screenshots named `chip-<state>.jpg` in `docs/superpowers/reviews/week1/`.
- [ ] **Step 5: Personas review** the ten captures; fix until both SHIP.
- [ ] **Step 6: Commit** `git commit -am "feat(addon): the chip"`

### Task L3.2: Talent marks

**Files:**
- Create: `addon/ForeverSixty/surfaces/TalentMarks.lua`
- Modify: `TalentGlow.lua` (keep `watch`, `openFrame`, `traitButton`; marks call it), `Locale.lua` (`talentStrip = "%s · %d of %d · next at level %d · %d off plan"`), `Theme.lua` (`Theme.SIZES.orderBadge = 14`)

**Interfaces:**
- Produces: `TalentMarks.refresh(data)` (ring on the next node via `TalentGlow`, a 14 px gold-on-dark number badge on every planned node = the order index of the cell's last point, and a 24 px strip anchored `BOTTOMLEFT` of the talent frame with the `talentStrip` text), `TalentMarks.selfCheck()` (every badge found a button; else the badges are off and one diag line names how many were missing).

- [ ] **Step 1: Implement** badges through `TalentGlow.traitButton(frame, node)` for every cell in `Progress`'s cell list; strip through `Widgets.label` on a frame parented to the talent frame; refresh on the frame's `OnShow` and on `TRAIT_CONFIG_UPDATED`.
- [ ] **Step 2: Capture checklist**: talent window with the default Arms leveling build: ring on Improved Heroic Strike, badges on every Arms and Fury node the plan names, strip reads "Arms leveling · 11 of 51 · next at level 21 · 0 off plan"; spend one point off plan on the beta (any free talent) and capture the strip in warning colour; close and reopen the window: marks persist; `/fs diag` empty.
- [ ] **Step 3: Personas review**; fix until SHIP.
- [ ] **Step 4: Commit** `git commit -am "feat(addon): talent marks"`

### Task L3.3: Item tooltip verdict and source

**Files:**
- Modify: `addon/ForeverSixty/Tooltip.lua` (rename to `surfaces/ItemTooltip.lua`; keep the processor hook, the compare tooltip and the recursion guard; replace `verdictFor` and the BiS lines with `Progress.slotVerdict` output), `Locale.lua` (`tooltipVerdictDps = "Upgrade +%.1f DPS over %s"`, `tooltipVerdictScore = "Upgrade +%d score over %s"`, `tooltipSide = "Sidegrade"`, `tooltipDown = "Downgrade −%.1f %s"`, `tooltipPlanned = "Planned for %s at %d"`, `tooltipSource = "%s · %s · %d%%"`, `tooltipAlternatives = "Also: %s, %s"`, `tooltipTagSim = "Your sim"`, `tooltipTagBuild = "Your build"`, `tooltipTagDefault = "Default weights"`, `tooltipUntil = L.rulesUntilLevel`)

**Interfaces:**
- Consumes: `Progress.slotVerdict`, `Rules.canEquip`, `Gear.pick`, inbox `weights`/`sim_result` via `Gear.currentWeights(spec)` → `{ weights, dpsPerPoint, tag }`.

- [ ] **Step 1: Implement** the section: tag line, verdict line (DPS when `dpsPerPoint` present), planned line, source line with chance, alternatives line; "Not for you until 40: plate" replaces the verdict when `Rules.canEquip` is not `"ok"`; unchanged compare tooltip beside it.
- [ ] **Step 2: Capture checklist**: bag item that is an upgrade; worn item; planned item in a vendor window; a mail item on the warrior at 20 (no refusal); a plate item on the warrior at 20 ("until 40"); a quest reward; with the companion's weights message present (copy the L0 inbox fixture into `WTF` for the character) the tag reads "Your sim" and the unit is DPS.
- [ ] **Step 3: Personas review**; fix until SHIP.
- [ ] **Step 4: Commit** `git commit -am "feat(addon): item tooltip verdict, source and alternatives"`

### Task L3.4: Build dialog

**Files:**
- Create: `addon/ForeverSixty/surfaces/BuildDialog.lua`
- Modify: `Options.lua` (`/fs` → `BuildDialog.open()`), `Locale.lua` (`dialogPaste = "Paste a build code"`, `dialogUse = "Use this build"`, `dialogCopy = "Copy my character"`, `dialogCopied = "Copied %s, paste it on the site"`, `dialogSeen = "Seen by the site %s"`, `dialogPreview = "%s · %s · %d points"`)

**Interfaces:**
- Produces: `BuildDialog.open(mode)` with `mode = "paste"|"copy"`; `BuildDialog.selfCheck()`.
- Consumes: `Codec.loadBuild`, `Follow.load`, `Export.string`, inbox `ack`.

- [ ] **Step 1: Implement**: 420×300 `Widgets.panel`, Escape-closable, a focused `EditBox`; `OnTextChanged` with user input validates through `Codec.loadBuild` and swaps the box for the preview (class icon, name, spec split from the order's tab counts, points, three gear picks with `Compat.itemIcon`) and one `Cards.primaryButton` "Use this build" (`OnEnterPressed` too); a refusal shows the codec's reason muted under the box and keeps the current build; copy mode selects the export and shows the copied line for 3 s, or the `ack` line when the inbox has one newer than `savedAt`.
- [ ] **Step 2: Capture checklist**: `/fs` empty dialog; paste the Arms code → preview; Enter → chip and strip change; paste a code with `FSB1:1.60.1.69893` (older build) → reason shown, current build kept; copy mode with and without an ack fixture; Escape closes; keyboard works after close (walk).
- [ ] **Step 3: Personas review**; fix until SHIP.
- [ ] **Step 4: Commit** `git commit -am "feat(addon): build dialog"`

### Task L3.5: Settings page and launcher

**Files:**
- Create: `addon/ForeverSixty/surfaces/Settings.lua` (client Settings API only), keep `Minimap.lua` (compartment + optional button)
- Modify: `Locale.lua` (five switch labels and hints; `settingsCommands` = the `/fs` list)

- [ ] **Step 1: Implement** with `Settings.RegisterCanvasLayoutCategory`; five `Widgets.toggle` (on = gold square with the client's `UI-CheckBox-Check` tinted dark; off = 1 px border) with wrapped hints (`SetWidth(panelWidth - 40)`), the command list as muted lines; `Reset positions` button.
- [ ] **Step 2: Capture checklist**: Settings page top and bottom (Reset visible); each toggle on and off; the compartment entry; the minimap button on and off.
- [ ] **Step 3: Personas review**; fix until SHIP.
- [ ] **Step 4: Commit** `git commit -am "feat(addon): settings page"`

---

## Lane L4: Addon surfaces B (week 2)

### Task L4.1: Paper-doll and bag marks

**Files:**
- Create: `addon/ForeverSixty/surfaces/SlotMarks.lua`
- Modify: `Options.lua` (`PLAYER_EQUIPMENT_CHANGED`, `BAG_UPDATE_DELAYED` → `SlotMarks.refresh`), `Theme.lua` (`Theme.MEDIA.arrow = "Interface\\Buttons\\UI-MicroStream-Green"`, `Theme.SIZES.mark = 12`)

- [ ] **Step 1: Implement**: for each `CharacterFrame` slot button (`CharacterHeadSlot` … `CharacterRangedSlot`) a 12 px corner texture at `TOPRIGHT` when `Progress.bagUpgrades` names the slot, and a target corner at `BOTTOMRIGHT` when `Gear.pick` names a next target; for bags, hook `ContainerFrame` item buttons (`ContainerFrameItemButton_OnUpdate` or the `BagUpdate` mixin callback on this client) and mark the upgrading item. Marks are textures on the client's buttons, hidden with them.
- [ ] **Step 2: Capture checklist**: character panel with one marked slot; open bags with the matching item marked; equip it: both marks gone; Baganator installed (the beta character has it): the mark appears on its buttons or, if the hook does not reach them, one diag line names the bag addon and the bag mark is off.
- [ ] **Step 3: Personas review**; commit `git commit -am "feat(addon): paper-doll and bag marks"`

### Task L4.2: Quest reward, roll window and dungeon line

**Files:**
- Create: `addon/ForeverSixty/surfaces/RewardMarks.lua`
- Modify: `Gear.lua` (`Gear.picksHere(instanceId)` → the BiS picks for the current band whose source is this instance, from `Data.bis` and `Data.dungeons` ids), `Locale.lua` (`rewardBest = "Best for %s"`, `rewardVendor = "Vendor"`)
- Modify: `Chip.lua` (the `here` state reads `Gear.picksHere`)

- [ ] **Step 1: Implement**: `QUEST_COMPLETE` → score each `GetQuestItemLink("choice", i)` through `Progress.slotVerdict`; a gold "Best for Arms" label on the best, "Vendor" when none beats what is worn; `START_LOOT_ROLL` → the verdict line beside the roll frame's item; `ZONE_CHANGED_NEW_AREA` → `Chip` `here` state when `picksHere` is non-empty.
- [ ] **Step 2: Capture checklist**: a quest turn-in with three rewards; a dungeon roll frame; entering The Deadmines at band 20.
- [ ] **Step 3: Personas review**; commit `git commit -am "feat(addon): quest reward, roll and dungeon marks"`

### Task L4.3: The "Forever Sixty" character-panel tab

**Files:**
- Create: `addon/ForeverSixty/surfaces/PanelTab.lua`
- Modify: `Theme.lua` (`Theme.addCharacterTab(name, frame)` using `PanelTemplates_SetNumTabs` on `CharacterFrame` where present), `Locale.lua` (region titles, `panelSimMe = "Sim me"`, `panelSimQueued = "Sim queued; results after your next login"`, `panelRotationAt = "at level %d"`)
- Modify: `Rotation.lua` (`Rotation.plainCondition(text)`: strips parentheticals and cuts at the first sentence end; `Rotation.learnedAt(spellId)` from `Data.abilityLevels`)

- [ ] **Step 1: Implement** the three regions (targets, rotation, build) and the two 60-only regions (sim, last raid) per spec §4.5; regions absent when empty; "Sim me" sets `ForeverSixtyDB.outbox.sim_request` and shows `panelSimQueued`.
- [ ] **Step 2: Capture checklist**: level 20 with the default build; level 60 character on the beta (the account has one) with the L0 inbox fixture (sim_result, last_raid) copied into `WTF`; rotation conditions two lines max with no parentheses; Execute marked "at level 24" and last.
- [ ] **Step 3: Personas review**; commit `git commit -am "feat(addon): character panel tab"`

### Task L4.4: Player tooltip score and chat badges

**Files:**
- Create: `addon/ForeverSixty/surfaces/PlayerTooltip.lua`, `addon/ForeverSixty/surfaces/ChatBadges.lua`
- Delete: `Ratings.lua` (replaced by `Score`)
- Modify: `Locale.lua` (`scoreLine = "Forever Score %d · %s · %s"`, `scoreNoLogs = "Forever Score: no logs yet"`, `scoreLive = "Level %d %s · gap %d (live)"`, `badge = "[%d]"`)

- [ ] **Step 1: Implement** the unit tooltip section (score line in the tier colour, components on `IsShiftKeyDown`, live `Comms.seen` line when present, "no logs yet" otherwise); `ChatBadges` through `ChatFrame_AddMessageEventFilter` for `CHAT_MSG_CHANNEL` (LFG, Trade), `CHAT_MSG_GUILD`, `CHAT_MSG_WHISPER`: prefix the name with `[1830]` in the tier colour when `Score.forCharacter` has it; off by the Settings switch.
- [ ] **Step 2: Capture checklist**: tooltip of a logged character (seed the nightly fixture with a guildmate), of an unlogged one, of a guildmate with the addon (live line), shift held; Trade chat with two badged names; badges switched off.
- [ ] **Step 3: Personas review**; commit `git commit -am "feat(addon): player tooltip score and chat badges"`

---

## Lane L5: Addon surfaces C (week 3)

### Task L5.1: Guild tab

**Files:**
- Create: `addon/ForeverSixty/surfaces/GuildTab.lua`
- Modify: `Theme.lua` (`Theme.addGuildTab`: a tab on `GuildFrame`/`CommunitiesFrame`, whichever this client has; recorded in `/fs diag`)

- [ ] **Step 1: Implement** the roster (`GetNumGuildMembers`/`GetGuildRosterInfo` rows joined with `Score.forCharacter` and `Comms.seen`), sortable columns, the officer region (inbox `guild` message: approvals, "Push the raid build" → `Comms.send("push", …)` after a confirm), and the loot council region placeholder that L5.3 fills.
- [ ] **Step 2: Capture checklist**: guild tab as a member; as an officer (the account's guild master character); sorted by gap; push confirm dialog.
- [ ] **Step 3: Personas review**; commit `git commit -am "feat(addon): guild tab"`

### Task L5.2: Group panel and ready check

**Files:**
- Create: `addon/ForeverSixty/surfaces/GroupPanel.lua`
- Modify: `Options.lua` (`GROUP_ROSTER_UPDATE`, `READY_CHECK`, `READY_CHECK_FINISHED`, `UNIT_AURA` for group units), `Locale.lua` (`groupWeakest = "%s: %s"`, `groupMissing = "Missing: %s"`, `groupCoverage = "%s covered by %s"`)

- [ ] **Step 1: Implement**: one row per `party`/`raid` unit from `Inspect.get`, `Score`, `Comms.seen`; the weakest-slot line from `Progress.slots` over inspected gear; ready check: on `READY_CHECK` scan every unit's auras (`AuraUtil.ForEachAura` or `UnitAura`) against the flask, food, imbue and class-buff spell id lists in `Data.encounterNeeds` for the current instance's bosses, list who lacks what; coverage checklist per boss; applicant hover reads `Score`; the leader's floor pref flags rows under it in muted text, nothing else.
- [ ] **Step 2: Capture checklist** (a test raid on the beta with at least five players, two running the addon): panel before the pull; ready check with a missing imbue; coverage list; a non-addon member's row from inspect; combat: inspect paused.
- [ ] **Step 3: Personas review**; commit `git commit -am "feat(addon): group panel and ready check"`

### Task L5.3: Loot council helper and last raid

**Files:**
- Modify: `GuildTab.lua`, `GroupPanel.lua` (loot region), `Comms.lua` handlers for `loot?`/`loot!`, `PanelTab.lua` (last-raid region already built in L4.3 reads inbox `last_raid`)
- Modify: `Options.lua` (`LOOT_OPENED` for the master looter → `Comms.send("loot?", …)` per epic item)

- [ ] **Step 1: Implement**: the master looter's addon sends `loot?` per item; every member replies `loot!` from `Progress.slotVerdict` (DPS when sim weights exist); the master looter's panel ranks replies within 5 s, fills non-addon members from the inbox `loot_values` message or `Inspect` where available, "no data" otherwise.
- [ ] **Step 2: Capture checklist**: a drop in the test raid: the ranked list; a member without the addon shown "no data"; the chip's `lastnight` state and the last-raid region after copying a `last_raid` fixture.
- [ ] **Step 3: Personas review**; commit `git commit -am "feat(addon): loot council helper"`

---

## Lane L6: API

### Task L6.1: Forever Score v2 with components

**Files:**
- Create: `api/internal/rating/score_v2.go`, `score_v2_test.go`, migration `api/internal/db/migrations/<next>_score_v2.sql` (`rating_scores` gains `gear_score`, `experience`, `attendance`, `discipline`, `tier`)
- Modify: `api/internal/rating/rater.go` (compute components per rated fight and roll up per character), `dto.go` (`characterRatingDTO.components`), `handler.go`, `api/openapi.yaml` (`/v1/characters/{region}/{ruleset}/{name}` gains `score`, `tier`, `components`, `as_of`, `experience`)
- Modify: `api/internal/character/character.go` (`LinkedCharacters(userID)`: mains and alts from `/v1/me`'s verified characters)

**Interfaces:**
- Produces: `rating.ScoreV2(ctx, characterKey) (ScoreV2, error)` with `ScoreV2{ Score int; Tier string; AsOf time.Time; Components{ Gear, Experience, Parses, Attendance, Discipline int }; GearScore int; ItemLevel int; Experience{ Dungeons, RaidKills, Alts int } }`; `Tier(score int) string` with the thresholds in L0 contract 5.
- Consumes: `sims.GearScore(gearCode, spec)` (new in L6.2), reports' consumables and enchant audit (`api/internal/reports` already parses buffs), `rankings` for parse percentiles.

- [ ] **Step 1: Failing Go tests**: a character with twelve rated fights, 71st percentile best, full attendance, one missing flask fight, two linked alts with dungeon clears → each component within the documented ranges and `Tier(1830) == "epic"`; a character with no fights → `ErrNoLogs` (never a zero score).
- [ ] **Step 2: Run** `cd api && go test ./internal/rating/ -run ScoreV2`. Expected: FAIL.
- [ ] **Step 3: Implement**: gear = `GearScore` scaled 0–500; experience = 25 per distinct dungeon cleared + 40 per distinct raid boss killed across linked characters, capped 700; parses = best-percentile mean scaled 0–600; attendance = fraction of guild nights attended in 90 days × 200; discipline = fights with full consumables and enchants / fights × 100; score = sum; `as_of` = the newest fight's date.
- [ ] **Step 4: Run** tests; migrate the test db (`docker-compose.test.yml`). Expected: PASS.
- [ ] **Step 5: Commit** `git commit -am "feat(api): Forever Score v2 with components and tiers"`

### Task L6.2: Gear score and sim request job

**Files:**
- Create: `api/internal/sims/gearscore.go`, `gearscore_test.go`, `api/internal/sims/request_job.go`
- Modify: `api/internal/addon/addon.go` (`POST /v1/addon/exports` accepts `outbox` v1: stores the export, enqueues a sim run for `sim_request`, stores `roster_snapshot` for the guild page), `api/openapi.yaml`

- [ ] **Step 1: Failing tests**: `GearScore` of the band-20 BiS set under Arms weights equals the sum of per-item scores from `Gear.score`'s Lua definition (shared fixture `docs/contracts/fixtures/gearscore.json`); an outbox with `sim_request` creates one `sims.Run` row keyed by the character and the request time; a second identical request within an hour is deduplicated.
- [ ] **Step 2: Run** `go test ./internal/sims/ ./internal/addon/`. Expected: FAIL.
- [ ] **Step 3: Implement**; the sim job reuses `sims.Run` with the character's `sim-input`.
- [ ] **Step 4: Run** tests. Expected: PASS.
- [ ] **Step 5: Commit** `git commit -am "feat(api): gear score and sim requests from the addon outbox"`

### Task L6.3: Inbox v2 messages

**Files:**
- Modify: `api/internal/addon/messages.go` (`format = 2`; new `MessageEntry` fields per L0 contract 3; builders `simResultMessage`, `lastRaidMessage`, `lootValuesMessage`, `ackMessage`), `messages_test.go`, `api/openapi.yaml` (`/v1/addon/inbox` response schema v2)
- Modify: `api/internal/reports/ingest.go` (after rating, compute the per-boss summary: percentile, biggest loss category from the engine's downtime/uptime analysis, deaths, consumables missed, adherence from `engine` cast order vs the sim's APL), `api/internal/sims/parses.go` (`LootValues(raidID, characterKey)`: the value of each boss's loot table item to the character from the latest sim)

- [ ] **Step 1: Failing tests**: the golden fixture from L0.2 round-trips; `lastRaidMessage` for a seeded report names Lucifron with `percentile 71` and `biggest_loss "downtime"`; `ackMessage` carries the newest export's `saved_at`; a user without the companion paired gets no `ack`.
- [ ] **Step 2: Run**. Expected: FAIL.
- [ ] **Step 3: Implement**.
- [ ] **Step 4: Run** `go test ./...` in `api`. Expected: PASS.
- [ ] **Step 5: Commit** `git commit -am "feat(api): inbox v2 messages"`

### Task L6.4: Nightly data addon v2

**Files:**
- Modify: `api/internal/dataaddon/aggregate.go`, `write.go`, `guild_aggregate.go`, `golden_test.go` (format 2 per L0 contract 5: score, tier, components, gear_score, item_level, spec, experience, best_parses, attendance, fights, profile; guild members list), `.github/workflows/addon-data-release.yml` (unchanged trigger)

- [ ] **Step 1: Failing golden test**: the rendered `Data.lua` for the seeded db byte-equals `docs/contracts/fixtures/data-addon-v2.lua` after date normalisation; a character under the 4 MiB split keeps its full record.
- [ ] **Step 2: Run**. Expected: FAIL.
- [ ] **Step 3: Implement**; keep the 90-day window; characters with no rated fight but a verified guild membership get a record with `score = nil` so the addon shows "no logs yet" and the guild roster still lists them.
- [ ] **Step 4: Run** tests. Expected: PASS.
- [ ] **Step 5: Commit** `git commit -am "feat(api): nightly data addon format 2"`

---

## Lane L7: Companion

### Task L7.1: Outbox up, inbox v2 down, ack

**Files:**
- Modify: `companion/internal/addon/sync.go` (read `ForeverSixtyDB.outbox` on `SavedVariables` change; `POST /v1/addon/exports` with the whole outbox; `GET /v1/addon/inbox` v2; write `ForeverSixtyInbox` with `format = 2`), `lua.go` (render every v2 message type), `addon_test.go`, `companion/internal/fakeapi` (v2 fixtures), `companion/internal/app` (tray status line "Seen by the site 14:03")

**Interfaces:**
- Consumes: L0 contracts 3 and 4; the API endpoints from L6.2 and L6.3.

- [ ] **Step 1: Failing tests**: a saved-variables file with `outbox.sim_request` produces one POST whose body includes it; an inbox v2 response renders a Lua chunk that `lua5.1` loads and whose `messages[1].type == "sim_result"`; an `ack` message appears after a successful export POST.
- [ ] **Step 2: Run** `cd companion && go test ./internal/addon/`. Expected: FAIL.
- [ ] **Step 3: Implement**; the writer never touches `ForeverSixtyDB` (the game owns it), only `ForeverSixtyInbox`.
- [ ] **Step 4: Run** tests; manual check on the beta: log out with a sim request pending, see the tray line change, log in and see the chip's `lastnight` or the panel's sim region.
- [ ] **Step 5: Commit** `git commit -am "feat(companion): outbox upload and inbox v2"`

---

## Lane L8: Website

### Task L8.1: "Copy for the game" everywhere a build or result exists

**Files:**
- Create: `web/src/components/CopyForGame.svelte`, `CopyForGame.test.ts`
- Modify: `web/src/pages/planner.astro`, `web/src/pages/sim/[id].astro`, `web/src/pages/sim/gear.astro` (Top Gear), `web/src/lib/addon-export.ts` (one `buildCodeFor(build | simResult)` that emits `FSB1` with the sim's weights as the stat lines)

- [ ] **Step 1: Failing vitest**: `buildCodeFor(simResult)` emits a code whose gear stat lines equal the sim's per-item deltas rounded per the codec's integer rule; the button copies to the clipboard and reads "Copied, paste it in game with /fs" for 3 s.
- [ ] **Step 2: Run** `cd web && npm test -- CopyForGame`. Expected: FAIL.
- [ ] **Step 3: Implement**; one component, one copy string, the addon's `/fs` named in it.
- [ ] **Step 4: Run** vitest and `npm run lint`. Expected: PASS.
- [ ] **Step 5: Commit** `git commit -am "feat(web): copy for the game on planner, sim and top gear"`

### Task L8.2: The Forever Score profile and search

**Files:**
- Modify: `web/src/components/CharacterRatingPanel.svelte` (score, tier, components with plain-word labels, "as of", history sparkline from `/v1/characters/...` v2), `web/src/pages/character/[...path].astro` (the "In game" section: addon status from the ack, last export seen, "Copy my character" instructions), `web/src/pages/ratings.astro` (what the score is, how each component moves, the tier thresholds, in plain words), `web/src/pages/index.astro` (character search box routes to the profile)
- Create: `web/src/components/ScoreBadge.svelte` (the same number and tier colour the addon shows), `ScoreBadge.test.ts`

- [ ] **Step 1: Failing tests**: `ScoreBadge` renders `1830` in the epic colour; the ratings page lists five components with the exact labels the addon's `Score.lines` uses (shared fixture `docs/contracts/fixtures/score-labels.json`); the character page shows "no logs yet" for a character the API returns 404 for, never a zero.
- [ ] **Step 2: Run**. Expected: FAIL.
- [ ] **Step 3: Implement**; playwright test for the search box → profile.
- [ ] **Step 4: Run** vitest, playwright, lighthouse budget. Expected: PASS.
- [ ] **Step 5: Commit** `git commit -am "feat(web): Forever Score profile, badge and explainer"`

### Task L8.3: Addon page and guild roster

**Files:**
- Modify: `web/src/pages/addon` (install, what it does on each surface with real client screenshots from the review folders, the `/fs` commands, companion handshake), `web/src/components/Guild.svelte` (roster from the roster snapshot: level, spec, gear gap, score; standing), `web/src/components/GuildRosterHandoff.svelte`

- [ ] **Step 1: Failing tests**: the addon page lists every surface named in the spec's §4 table; the guild roster renders the snapshot fixture sorted by gap.
- [ ] **Step 2: Run**. Expected: FAIL.
- [ ] **Step 3: Implement** with the week-3 captures as images (`web/public/addon/*.jpg`, under 200 KB each).
- [ ] **Step 4: Run** tests and lighthouse. Expected: PASS.
- [ ] **Step 5: Commit** `git commit -am "feat(web): addon page and guild roster from the addon"`

---

## Lane L9: Release and QA

### Task L9.1: The weekly capture loop

**Files:**
- Create: `docs/superpowers/reviews/README.md` (the checklist below), `docs/superpowers/reviews/week{1,2,3,4}/` folders
- Create: `addon/tools/capture_checklist.md` (every surface × every state × level 1, 20, 60, in combat)

- [ ] **Step 1: Write the checklist** from the spec's §4 table and each lane's capture steps, one line per capture with its file name.
- [ ] **Step 2: Each week**: run the checklist on the beta (the `ux-designer` and `wow-player` agents receive the folder and the manifest exactly as in the Oct 11 session); both must return SHIP for the week's lanes before their PRs merge.
- [ ] **Step 3: Commit** the folder and the two reports each week.

### Task L9.2: Performance pass and release

**Files:**
- Modify: `addon/ForeverSixty/ForeverSixty.toc` (version, `X-Curse-Project-ID`, `X-Wago-ID` present), `addon/ForeverSixty/CHANGELOG.md` (`## 2.0.0`), `addon/README.md` (spike rows for every new client surface: talent frame name, guild frame name, character tab API, chat filter, inspect)

- [ ] **Step 1: Measure** with the client's addon memory list after 30 minutes idle and after a raid: under 2 MB idle, no frame created at load (`/fs diag` lists the frames created and when).
- [ ] **Step 2: Fix** anything over budget (inspect queue spacing, `Comms.seen` pruning after 1 h, tooltip cache cap 200).
- [ ] **Step 3: Release**: `git tag addon-v2.0.0 && git push origin addon-v2.0.0`; verify the CurseForge and Wago listings show the LX icon and the changelog.
- [ ] **Step 4: Commit** `git commit -am "release: addon 2.0.0"`

---

## Self-review notes

- Spec coverage: §4.1 L3.1; §4.2 L3.2; §4.3 L3.3; §4.4 L4.1–L4.2; §4.5 L4.3; §4.6 L4.4 + L6.1 + L8.2; §4.7 L5.1; §4.8 L5.2; §4.9 L2.4 + L5.3; §4.10 L3.4; §4.11 L3.5; §5 loops are the union of the above plus L1; §6.1 L1; §6.2 L6.4; §6.3 L6.3 + L7; §6.4 L2.6 + L6.2 + L7; §6.5 L2.1; §7 architecture is the file list of L2–L5; §8 testing is the per-task steps and L9.1; §9 weeks are the lane gates.
- Type consistency: `Progress.slotVerdict` fields (`state`, `delta`, `unit`) are what L3.3, L4.1, L4.2, L5.2 and L5.3 read; `Comms.seen` shape equals the `hello` fields; `Score.forCharacter` returns the contract-5 record plus `asOfText` and `tierColor`.
- Review Focus pins: 1 → L2.3 step 1 and L5.2 captures; 2 → L3.4 captures and L2.2 nil-build case; 3 → L2.4 fuzz; 4 → L2.1 fixtures and L3.3 captures; 5 → L3.1 and L5.2 combat captures.
