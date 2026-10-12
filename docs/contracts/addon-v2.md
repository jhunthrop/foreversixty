# Forever Sixty addon v2: shared wire contracts

Date: 2026-10-11. Branch: `addon/v2-foundation`.
Spec: `docs/superpowers/specs/2026-10-11-addon-v2-design.md` (sections 4, 6 and 9).

Five contracts, each stated once, here. Every lane of the addon v2 programme codes against these
blocks and not against another lane's code. If a lane needs a field that is not here, it changes
this file first, in its own commit, and says so in the lane report.

Reading order: contract 1 is what ships inside the addon. Contract 2 is what players' addons say
to each other. Contracts 3 and 4 are the private two-way file exchange with the companion app.
Contract 5 is the public nightly data file.

## Rules that apply to every contract

- **Versioning.** Each contract carries its own version number in its first field, and a reader
  that meets a version it does not know ignores the whole thing silently and shows nothing in
  its place (spec 2.2: a surface with nothing to say is absent). A writer only ever adds fields
  inside a version; it bumps the version to remove or change one.
- **Names already in use stay.** The existing inbox messages (`api/internal/addon/messages.go`,
  `companion/internal/addon/addon.go`) and the existing nightly reader
  (`addon/ForeverSixty/Ratings.lua`) fix some field names. Where a block below keeps one, it
  keeps it unchanged. Where a field is new, or replaces an older one, the contract's
  versioning note says so.
- **Player-visible text is already plain.** Any string a player will read (`buildName`, `name`,
  `source`, `biggest_loss`, `consumables_missed`, `item_name`, `guild_name`, and the like) is
  written by its producer as plain words ready to show: no spell or item ids, no build numbers,
  no internal codes. The addon shows it as given and never composes a sentence from a code.
  Dates travel as ISO `YYYY-MM-DD` and the addon shows them as "as of 4 Nov". Numbers on screen
  always come from the one function that owns them (spec 2.1). The words "data addon" and
  "companion" never appear in player copy without a four-word explanation beside them.
- **Ids are for machines.** `item_id`, `spellId`, `npcId`, `lootId` and the keys of
  `encounterNeeds`, `abilityLevels` and `loot_values` items are lookup keys only. They are never
  printed.
- **No network call from the addon.** Everything below travels in shipped data, in the client's
  addon-message channels, or in saved-variable files the companion reads and writes.

---

## 1. `Data.lua` additions

**Who writes, who reads, how versioned.** Lane L1 writes it: the nightly pipeline (`python -m
pipeline addon-data`) regenerates `Data.lua`, which is committed, and `--check` is the drift
gate. Lane L2 reads it through `ns.Data`. The file is keyed by the client build id in `build`,
so a new client patch means a new generated file; the addon compares `build` to the running
client and says "data is for an older patch" when they differ (the build id itself is never
shown). The existing keys (`classes`, `weights`, `rotations`, `bis`) keep their current shapes
and names; the keys marked NEW are added in this version, and an addon that does not know a NEW
key simply never reads it.

```lua
ns.Data = {
  build = "1.60.1.70291",
  classes = { ... },            -- unchanged
  weights = { ... },            -- unchanged
  -- NEW: DPS (or HPS for healers) per 1.0 of the spec's reference stat, from the site's sim
  dpsPerPoint = { ["warrior-arms"] = { reference = "attack_power", dps = 0.42, healing = false }, ... },
  -- NEW: default builds per spec, as the codec's own triples (tab, tier, column), 1-based
  -- The addon loads defaultBuilds[<spec>].leveling on first login when no build is saved for the
  -- character (raid at 60), and marks it default = true so the build dialog can say
  -- "Using the default Arms leveling build".
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

Notes for the reader:

- `dpsPerPoint` is per spec slug; `healing = true` means the number is healing per second and
  the addon says "healing" instead of "damage". A spec with no entry shows no DPS wording at all.
- `defaultBuilds` order lists are the 51 point spends in learning order, one `{tab, tier,
  column}` triple per point, 1-based, exactly what `Codec` encodes. The saved copy of a default
  carries `default = true`; the first time the player saves, imports or is pushed a build for
  that character the flag is dropped.
- `chance` is a fraction (0.14 is 14 percent) or nil when the pick is not a drop. `source` is
  the plain name a player knows ("Lucifron", "Blackrock Depths", "Tailoring"), never an id.
- Physical encoding of a bis pick: on disk, `Data.lua` stores each pick as one packed string,
  not as a table, so the loaded data stays inside the addon's 2 MB idle memory budget. The
  fields are joined by `|` in this order:
  `item_id|kind|source|chance|alt1_id|alt1_kind|alt1_source|alt2_id|alt2_kind|alt2_source`.
  `kind` is the one-letter source code. `chance` is a plain number, or empty when it is nil. A
  missing field is empty, so the positions never shift. Trailing empty fields are trimmed, and
  `source` never contains `|`. Readers never split the string themselves:
  `Gear.pick(spec, band, faction, slot)` returns the logical table shown above,
  `{ item_id, kind, source, chance, alternatives = { { item_id, kind, source }, ... } }`, and
  `Gear.decodePick(packed)` decodes a single string. Positions 1-3 of that table also hold
  the id, the kind and the source.
- `encounterNeeds` ids are spell ids for buffs and debuffs and are used only to check the raid's
  auras; the coverage list shows spell names from the client. `resist` is a school name or nil.

---

## 2. FS1 addon-message protocol

**Who writes, who reads, how versioned.** Lane L2 implements the encode, decode and rate
limiting in `Comms` (pure functions, tested with busted); Lanes L4 and L5 send and receive
through it. The prefix is `FS1` and the first payload field is the protocol version `v`, so a
message from a future version (`v` other than 1) is dropped silently and never shown as an
error. No message ever causes an action on the recipient's behalf: `push` queues a build the
player may use, `loot!` is shown to the master looter, nothing is applied automatically.

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

Notes for the reader:

- "Values URL-safe" means a sender percent-encodes `|`, `\n` and `%` in `buildName` and
  `buildCode` and a receiver decodes them. A message is one client addon message, so a
  `buildName` that would push it past the client's 255-byte limit is shortened by the sender,
  never split across messages.
- `buildCode` in `share` and `push` is an FSB1 build code: the site's own `FSB1:` form that
  `Codec.decodeFSB1` reads. It is never an FS1 export string (the character export of contract
  4). The `FS1` message prefix and the `FS1:` export-string prefix are unrelated names that
  happen to match.
- `snapshot?` and `snapshot!` go out on GUILD when the guild tab asks and on RAID or PARTY when
  the group panel asks. A receiver answers on the channel the question arrived on.
- `scoreDate` is the date of the sender's last Forever Score as `YYYY-MM-DD`, or `-` for none.
  The receiver shows it as "as of 4 Nov".
- `push` is honoured only after the receiver has checked the sender's standing itself, with
  `GetGuildInfo` for the guild channel and `UnitIsGroupLeader` for raid. A failed check drops the
  message without a reply.
- `loot?` and `loot!` carry numbers computed by the one scoring function (spec 2.1). `unit`
  says which, and a `loot!` for a different `lootId` than the open request is ignored.

---

## 3. Companion inbox v2 (`ForeverSixtyInbox`)

**Who writes, who reads, how versioned.** Lane L6 (the API) emits the messages, Lane L7 (the
companion) renders them into `ForeverSixtyInbox.lua` every ten minutes, and Lanes L2, L4 and L5
read them. The file is a private, per-account exchange and is never shown to anyone else. The
top-level `format` field is new and is the version this contract keys on; an addon that reads a
`format` greater than it knows ignores the file and says so in one line in `/fs diag`. Each
message is selected by `type`; an unknown `type` is skipped, so a new message type never needs a
new format number.

```lua
ForeverSixtyInbox = {
  format = 2, generated_at = "2026-11-04T21:03:00Z",
  messages = {
    { type = "build",      character = "us/normal/obnoxious", name = "Arms raid", code = "FSB1:..." },
    { type = "upgrade",    character = ..., slot = "chest", item_id = 6627, item_name = ..., source = ..., delta = 4.2, unit = "dps" },
    { type = "weights",    character = ..., spec = "warrior-arms", weights = { {stat="attack_power", weight=1}, ... }, caps = {"hit"}, dps_per_point = 0.42 },
    { type = "sim_result", character = ..., spec = ..., dps = 412.6, error = 3.1, generated_at = ..., slots = { {slot="chest", item_id=6627, delta=4.2}, ... } },
    { type = "last_raid",  character = ..., report_id = "r_...", date = "2026-12-09", bosses = { {name="Lucifron", percentile=71, biggest_loss="downtime", loss_seconds=18, deaths=0, consumables_missed={"flask"}, adherence=0.82}, ... } },
    { type = "loot_values", character = ..., raid = "Molten Core", items = { [17076] = { value = 6.3, unit = "dps", planned = true }, ... } },
    { type = "guild",      character = ..., guild_name = ..., claim_state = ..., pending_approvals = 2, rank = "officer" },
    { type = "ack",        character = ..., export_seen_at = "2026-11-04T14:03:00Z", sim_request_state = "queued" },
  },
}
```

Notes for the reader:

- **Versioning against the file the companion writes today.** Today's file
  (`companion/internal/addon/addon.go`, `RenderInbox`) writes `version = 2` (informational
  only, nothing reads it), `generated_at`, a top-level `builds` list of `{id, name, character,
  code}`, and `messages` of type `upgrade`, `weights` and `guild`. v2 keeps `generated_at`,
  `messages`, and the field names of `upgrade`, `weights` and `guild` unchanged. New in v2:
  the top-level `format`; the `build` message type (the same `character`, `name`, `code` as
  today's `builds` entries, which the companion keeps emitting alongside it until the addon
  reads only `build`); `unit` on `upgrade` (absent means `"dps"`); `dps_per_point` on
  `weights`; and the `sim_result`, `last_raid`, `loot_values` and `ack` types.
- `character` is the character key `region/ruleset/name-slug` (`"us/normal/obnoxious"`), the
  form the `upgrade`, `weights` and `guild` messages and the `builds` entries have always
  carried, and selects which logged character a message is for; the addon reads only the
  messages whose `character` names the one logged in, matched segment by segment, case-folded
  (`sameCharacter` in `Codec.lua` and `Follow.lua`). It is not `Name-Realm`: the site stores the
  ruleset a realm maps to, not the realm.
- **The API's body.** `GET /v1/addon/inbox` answers exactly the shape the companion writes:
  `{ format = 2, generated_at, builds, messages }`, `generated_at` an RFC 3339 time in UTC.
  `docs/contracts/fixtures/inbox-v2.json` is that body.
- `delta` is always the change for the character, in the message's `unit`; `unit` names what
  the number is measured in (`"dps"`, which for a healer means healing per second, the same
  meaning as `dpsPerPoint.healing` in contract 1). The addon prints the number from this field
  and never recomputes it.
- Which `unit` values each message may carry: `upgrade` and `loot_values` items carry `"dps"` or
  `"score"`; `sim_result` slots carry no `unit` field and are always `"dps"`.
- `last_raid.date` is `YYYY-MM-DD`, shown as "as of 9 Dec". `biggest_loss` is a short plain
  phrase ("downtime", "standing in fire", "dead early") written by the API; the addon shows it
  as given. It is absent on a perfect card, where no component cost a point. `consumables_missed` entries are plain item names ("flask"), not ids.
  `adherence` is a fraction of the rotation followed (0.82 is 82 percent), to two decimals: of
  the player's own ability casts in the fight (not white hits, consumables, mounts, movement or
  utility, and not a pet's), the share that are abilities the spec's shipped rotation lists
  for the character's level band, at any rank. The spec is the rated card's; the level is the
  one the site stores for the character when the fight is rated (Battle.net import or newest
  export, else 60), so a fight re-rated later uses the level of that day. It is optional:
  absent when the spec has no shipped rotation, when the player cast fewer than 10 abilities,
  and for a fight rated before migration 0037. It is shown, never scored.
  `percentile` is the fight card's own overall (0 to 100). `loss_seconds` is the downtime
  seconds when the loss is downtime, else the seconds the player lay dead, else 0.
  `consumables_missed` is a list even when empty. One `last_raid` per character, for its
  newest report that is not private with a rated fight in the last 14 days; none without one. A
  fight rated before migration 0036 has none of the stored facts a boss line shows and is left
  out of `last_raid` rather than shown as zeros.
- `weights.dps_per_point` is the DPS (healing per second for a healer) one point of the stat
  whose weight is exactly 1 is worth, from the data build's `dps-per-point.json` for the spec;
  absent for a spec that file has no entry for (the tanks), and absent when the run's
  weight-1 stat is not that file's `reference` stat (its points are not the file's points).
- `sim_result.slots` carries per-slot deltas only when the run had per-item figures (a Top
  Gear run); for the plain run an in-game request queues it is an empty list, and the addon
  prices a slot itself from `weights` (`dps_per_point` times the weight delta). `dps` and
  `error` are always present, `error = 0` included.
- `loot_values.items` is keyed by item id (a lookup key, never printed). `planned = true` means
  the piece is on the character's plan. One `loot_values` per character with a `last_raid`, for
  the raid whose name matches that report's zone (case, spaces and punctuation ignored); none
  when no raid matches. `value` is the item's stat-line score less what the character wears in
  that slot, in `"dps"` (score times `dps_per_point`) when the spec has one and the weights are in its
  reference stat's points, else `"score"`; an off-hand item is left out while the character
  wears a two-hander, and a two-hander is weighed against both hands together;
  `planned` is true for the spec's level-60 best-in-slot pick.
- `ack.export_seen_at` is when the site last saw the character's export, an RFC 3339 time in
  UTC; the build dialog shows it as "Seen by the site 14:03" in local time. `ack` is sent only
  to an account with a companion paired. Its optional `sim_request_state` is what became of
  the character's newest in-game sim request: `"queued"` (a run started), `"needs_premium"`
  (server sims are a premium feature), `"duplicate"` (another request started a run within
  the hour) or `"invalid"` (the request could not be simmed); absent when the character never
  asked.

---

## 4. Companion outbox (`ForeverSixtyDB.outbox`)

**Who writes, who reads, how versioned.** Lane L2 writes it from the addon on `PLAYER_LOGOUT`,
so it holds what the player's last session left; Lane L7 (the companion) reads it from the
`ForeverSixty.lua` saved-variables file and sends it to the site. The addon never reads it back
and cannot see whether an upload happened. `format` starts at 1 and is the version this contract
keys on; the companion ignores an outbox whose `format` it does not know, leaving the file
untouched. Each of `sim_request` and `roster_snapshot` is `nil` when there is nothing to send,
and the companion treats a missing key and `nil` the same. `export.code` is the codec's own
`FS1:` export string, unchanged from what the addon already writes per character; the outbox
holds it under one key so the companion has one place to look.

```lua
ForeverSixtyDB.outbox = {
  format = 1,
  export = { character = "Obnoxious-Yell", code = "FS1:...", saved_at = "2026-11-04 14:02" },
  sim_request = { character = ..., spec = "warrior-arms", requested_at = ..., gear_code = "FS1:..." } | nil,
  roster_snapshot = { guild = "OLYMPUS XXVII", members = { { name=..., level=20, spec="warrior-arms", gear_gap=14, build="Arms leveling", seen_at=... }, ... } } | nil,
}
```

Notes for the reader:

- `export.saved_at` is the player's local wall time as `YYYY-MM-DD HH:MM`. It is what the build
  dialog shows as "Copied 14:02, paste it on the site"; the site's own time is in the inbox
  `ack`, not here.
- `sim_request` exists only after the player asked for a sim; `requested_at` and `gear_code` are
  as of that request, so a request made at level 20 is not answered with gear from level 24.
- `roster_snapshot.members[*]` mirrors the `hello` fields of contract 2 plus `name` and
  `seen_at`; `build` is the member's build name as that member's addon announced it, and
  `gear_gap` is the same integer. The member fields are exactly `name`, `level`, `spec`,
  `gear_gap`, `build` and `seen_at`; there is no `scoreDate` here, because the site holds the
  score itself. `seen_at` is `YYYY-MM-DD HH:MM` local. A snapshot carries only
  the members the guild tab has live data for.

---

## 5. Nightly data addon v2 (`ForeverSixtyData`)

**Who writes, who reads, how versioned.** Lane L6 writes it nightly from the API (the file
players install as the Forever Sixty data addon, which holds the site's scores); the addon's
`Score` module (Lane L2) reads it. Today's reader (`addon/ForeverSixty/Ratings.lua`) is
`FORMAT = 1`. This block is `format = 2`: an older addon that meets `format = 2` refuses it with
one line, and this addon refuses anything newer than 2 the same way. The score formula lives in
the API (`api/internal/rating`); the addon displays and dates the score and never computes it.

```lua
ForeverSixtyData = {
  format = 2, generated = "2026-11-04",
  characters = {
    ["Obnoxious Yell-ClassicBetaPvE2"] = { score = 1560, tier = "epic", as_of = "2026-11-04",
      components = { gear = 412, experience = 350, parses = 540, attendance = 180, discipline = 78 },
      gear_score = 412, item_level = 31, spec = "warrior-arms",
      experience = { dungeons = 14, raid_kills = 0, alts = 2 },
      best_parses = { ["Lucifron"] = 71 }, attendance = 0.9, fights = 12, profile = "/character/us/normal/obnoxious-yell" },
    -- a verified guild member with no logs yet
    ["Quiet Fern-ClassicBetaPvE2"] = { score = nil, item_level = 22, spec = "warrior-fury", profile = "/character/us/normal/quiet-fern" },
  },
  guilds = { ["OLYMPUS XXVII"] = { standing = 3, nights = 4, roster = 31, attendance = 0.8, members = { "Obnoxious Yell-ClassicBetaPvE2", "Quiet Fern-ClassicBetaPvE2", ... } } },
}
```
Tiers by score: `common` < 500, `uncommon` < 1000, `rare` < 1500, `epic` < 2000, `legendary` ≥ 2000. These thresholds are provisional: the API lane sets the final numbers from the live distribution before week 2 closes, and the addon reads `tier` from the record, never computing it.

Notes for the reader:

- **What changes from format 1.** Format 1 keys characters and guilds by the slug
  `region:realm:name` (`Ratings.characterKey`), carries `generated` as an RFC 3339 timestamp
  that `Ratings.ageOf` reads, a single `rating`, and six components (`output`, `survival`,
  `mechanics`, `utility`, `preparation`, `activity`) taken from the old log-based rating. Format
  2 keys characters by `Name-Realm` and guilds by guild name, carries `generated` as `YYYY-MM-DD`, and replaces `rating` and
  the six components with `score`, `tier`, `as_of` and the five components above. `fights` keeps
  its name (it now counts the rated fights the score rests on). `gear_score`, `item_level`, `spec`, `experience`, `best_parses`,
  `attendance` and `profile` are new. Lane L2 updates `Ratings` (or replaces it with `Score`) in
  step; the old six-component tooltip line is retired with format 1.
- `as_of` is the date of the newest rated fight the score rests on (`rating.ScoreV2`'s own),
  `YYYY-MM-DD`; the addon shows it as "as of 4 Nov" beside the score and never shows a score without its date. `generated` is when the file
  was written, and in format 2 it is a date-only string (`YYYY-MM-DD`). Today's `Ratings.ageOf`
  requires an RFC 3339 timestamp, so the format 2 reader must parse a date-only value (treat it
  as the start of that day, UTC). A file whose `generated` is more than three days old is shown but marked stale,
  as today.
- `tier` is one of the five names above, which are the client's own item-quality names and use
  the client's quality colours. A missing or unknown `tier` shows the score uncoloured.
- `components` values sum to `score`. `best_parses` keys are boss names, shown as given.
  `attendance` is a fraction (0.9 is 90 percent). `profile` is a site path, shown to the player
  only as a copyable address.
- `guilds[*].members` holds the `Name-Realm` keys of `characters`; the other guild fields are
  the site's own numbers, shown as given.
- **Keys.** A character's key is `Name-<realm>`: its name as the game shows it, and the
  client's own realm normalised by removing every space and hyphen (`Classic Beta PvP` is
  written `ClassicBetaPvP`, `Classic Beta PvE 2` is `ClassicBetaPvE2`), the second form
  `Score.forCharacter` looks a character up by and the one every spelling of the realm reduces
  to. The realm comes from the first source that has one: the Battle.net import's realm name
  (`characters.realm_name`), else the stored export's `who=` realm, else the realm half of the
  newest rated row's logged unit (`Name-Realm`). A character with no realm from any of them,
  or no stored name, has no record (logged and counted) rather than a guessed key. A guild's
  key is its name as `GetGuildInfo` returns it. Neither key carries the region (nor the guild
  key the ruleset), so two characters or guilds that would share one are both left out rather
  than one read as the other.
- **Who is in the file.** Every character with a rated fight in a public report in the last 90
  days, and every verified guild member, has a record; `rating.ScoreV2` alone decides which
  kind: a score is a scored record, and a character with no rated fight in a public report
  has a no-logs record, `{ score = nil, profile = ... }` plus `item_level` and `spec` when
  they are known, and no `tier`, `as_of` or `components`: the addon shows "no logs yet",
  never a zero. A character whose account asked to be anonymous has no record and is in no
  guild's `members`. `fights` is how many rated fights in public reports the score rests on
  (not only the last 90 days'),
  and `best_parses` the best overall percentile per boss from the rated fights of the last 90
  days (an empty table when none is that recent). `item_level` and `spec` are absent when
  the site does not know them. Fractions are written to the hundredth.
- **Guild fields.** `standing` is the guild's rank among every guild with a verified member, by
  distinct raid bosses killed in the last 90 days, ties by `attendance` (guilds still tied share
  a rank). `nights` is the distinct UTC dates of its public reports in the last 90 days.
  `roster` is its verified member count. `attendance` is the mean of the attendance fractions
  of its listed members that have a score (a no-logs member has no fraction and is not
  counted), absent when none has one.
- `docs/contracts/fixtures/data-addon-v2.lua` is the file the API's nightly job writes for its
  golden test's seeded database, byte for byte once the run's dates are moved onto the
  fixture's (`api/internal/dataaddon/golden_test.go`).
