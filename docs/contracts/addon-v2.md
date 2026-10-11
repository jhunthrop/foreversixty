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
- `character` is `Name-Realm` and selects which logged character a message is for; the addon
  reads only the messages whose `character` is the one logged in.
- `delta` is always the change for the character, in the message's `unit`; `unit` names what
  the number is measured in (`"dps"`, which for a healer means healing per second, the same
  meaning as `dpsPerPoint.healing` in contract 1). The addon prints the number from this field
  and never recomputes it.
- `last_raid.date` is `YYYY-MM-DD`, shown as "as of 9 Dec". `biggest_loss` is a short plain
  phrase ("downtime", "standing in fire", "dead early") written by the API; the addon shows it
  as given. `consumables_missed` entries are plain item names ("flask"), not ids.
  `adherence` is a fraction of the rotation followed (0.82 is 82 percent).
- `loot_values.items` is keyed by item id (a lookup key, never printed). `planned = true` means
  the piece is on the character's plan.
- `ack.export_seen_at` is when the site last saw the character's export, an RFC 3339 time in
  UTC; the build dialog shows it as "Seen by the site 14:03" in local time.

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
  `gear_gap` is the same integer. `seen_at` is `YYYY-MM-DD HH:MM` local. A snapshot carries only
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
    ["Obnoxious-Yell"] = { score = 1830, tier = "epic", as_of = "2026-11-04",
      components = { gear = 412, experience = 620, parses = 540, attendance = 180, discipline = 78 },
      gear_score = 412, item_level = 31, spec = "warrior-arms",
      experience = { dungeons = 14, raid_kills = 0, alts = 2 },
      best_parses = { ["Lucifron"] = 71 }, attendance = 0.9, fights = 12, profile = "/character/us/forever/obnoxious-yell" },
  },
  guilds = { ["OLYMPUS XXVII"] = { standing = 3, nights = 4, roster = 31, attendance = 0.8, members = { "Obnoxious-Yell", ... } } },
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
  its name and meaning. `gear_score`, `item_level`, `spec`, `experience`, `best_parses`,
  `attendance` and `profile` are new. Lane L2 updates `Ratings` (or replaces it with `Score`) in
  step; the old six-component tooltip line is retired with format 1.
- `as_of` is the date the score was last calculated, `YYYY-MM-DD`; the addon shows it as "as of
  4 Nov" beside the score and never shows a score without its date. `generated` is when the file
  was written; a file whose `generated` is more than three days old is shown but marked stale,
  as today.
- `tier` is one of the five names above, which are the client's own item-quality names and use
  the client's quality colours. A missing or unknown `tier` shows the score uncoloured.
- `components` values sum to `score`. `best_parses` keys are boss names, shown as given.
  `attendance` is a fraction (0.9 is 90 percent). `profile` is a site path, shown to the player
  only as a copyable address.
- `guilds[*].members` holds the `Name-Realm` keys of `characters`; the other guild fields are
  the site's own numbers, shown as given.
