# Prep reminder, one hour before the raid: experience spec (lane prep-reminder)

2026-11-15. Inputs: design/reviews/2026-11-15-guild-diff/00-facts.md (binding), 01-plan.md, 02-officer-plan.md, 03-brief.md (relay design owner approved, specced in full in section 6). Fixture world from design/specs/2026-11-10-guild-tab.md: Olympus XXVII, Thandrel (Fury Warrior 60, raider), Kaelra (officer, runs the companion), Thalrin (raider), Felsnap (36, not in the raid). The raid: Onyxia, Wed 8:00 PM server, so T-1h is 7:00 PM.

Vocabulary. A check is one reading of one character's prep with a stamp (GetServerTime for in game, the source's own time off game). A source is Game (the raider's own addon), Relay (a relay addon's collection), Upload (last companion upload of SavedVariables), Blizzard (signed-in profile), Public (directory profile). Unknown is never shown as missing. Only a source that can read a thing may say it is missing.

## 1. The one question and the answer above the fold

The raider asks: am I ready for tonight, and what do I fix in the next hour. The officer asks: who is ready, who could I not check, and why.

- Raider, 360 px phone and 1280 px desktop, no scroll: the raid line (`Onyxia · Wed 8:00 PM · in 1 h`), the verdict (`2 to fix before 8:00 PM`, or `All set on what we check`), the first two fix rows, then the stamp (`Checked 7:00 PM` or `Blizzard 3 h · bags 4 days`).
- Officer, same viewports: the split count line (`12 in game · 6 relay · 13 profile · 7 could not be checked`), the relay line (`Relay online: Kaelra` or `No relay online.`), then the first group header and rows. On a phone Could not check follows Checked and is never collapsed.

Tenet 14: the page leads with the raider's own character and the next fix. No pitch, no explainer.

## 2. Reference

- In-game: the raid frame's ready check and the existing /fs ready check (addon ReadyCheck.lua), the Details! "who is missing what" list. Borrowed: one row per person, class colour on the name, a plain verdict per row, a time. Not borrowed: chat output.
- Web: the Wowhead BiS table's row rhythm and the site's own ReadinessBoard.svelte and GuildReadiness.svelte (guild/dense), the ready-check lines (api/internal/guilds/ready_check.go ReadyItem: key, state, title, line, where, icon). The new surfaces are those lines plus a stamp and a source, not a new layout.
- Phone: the app's Guild tab and Activity inbox (design/specs/2026-11-10-guild-tab.md, 2026-11-08-app-flows.md).
- Discord: the raid reminder embed (api/internal/notify/reminder.go WebhookReminderEmbed); the DM is one embed in the same family, class colour on the left bar.

## 3. Data paths and latency

Freshness limits are the ones in 00-facts.md. Blizzard is not hourly.

| Fact shown | Source | Best freshness | Worst freshness | Shown when unknown |
|---|---|---|---|---|
| Worn gear, item level | Blizzard signed-in (bnetimport, addon_exports source blizzard) | 3 h (account active in 14 days) | 20 h | `Gear not on file` |
| Worn gear | Directory (character_directory, public profile) | 72 h re-read | 72 h plus the next hourly pass | `Gear not on file` or `Profile is private` |
| Enchants, suffixes | Signed-in Blizzard, or addon export | 3 h / 20 h; game: seconds | 20 h | Directory: `Enchants unknown`. Never `missing` |
| Resistance from gear | Derived from equipped item ids against our item data | as the gear | as the gear | `Resistance unknown` |
| Resistance enchant (cloak, libram style) | Signed-in Blizzard only | 3 h / 20 h | 20 h | `Enchants unknown` for directory characters |
| Resistance total | Game: UnitResistance (RaidPrep.resistances) | seconds | 1 h (the T-1h check) | `Not read` |
| Flask, elixir, food, protection potions in bags | Game: C_Container (RaidPrep.supplies) | seconds | the T-1h check | `Bags not read` |
| Same, off game | Last companion upload (addon export, `exported_at`) | minutes after a logout, /reload or relay reload | days | `No bag upload yet`; with an upload: `Had 2 flasks 4 days ago` |
| Attunement | Game (quests, items); off game from the upload | as bags | as bags | `Attunement not read` |
| Raid time | Companion inbox `raid_signup` (every 10 min); else a GUILD announcement from an officer or relay addon | inbox: fetched every 10 min (sync.go InboxEvery), read by the addon only at login or /reload | announcement: one per 5 min per relay | In game: `No raid time known`, no check at T-1h |
| Relay collection to the site | Relay addon, SavedVariables, relay companion | after the button: the reload writes SavedVariables, the companion reads it on its 1 s tick (companion/internal/addon/sync.go Poll, memoised on mtime) and uploads within seconds; the spike times it end to end | until the button is pressed | `Not sent yet` |

Off game paths (a) never contain a live bag reading. Every bag line off game carries its age. A row older than 24 h is muted and ends `Check in game`.

Paths:
- (a) Off game. Server job at T-1h minus 2 min per raid (notify, new `prep.reminder` kind) reads each In or Late raider's character: Blizzard gear and enchants (signed-in 3 h / 20 h), directory gear (up to 72 h, no enchants), latest addon export bags with its age. One DM per raider (Discord), one push (phone), same body. The raider reads their own data, so no guild consent gate applies to their own message.
- (b) In game. The raider's addon at T-1h reads its own bags; nothing leaves the game.
- (c) Relay. Section 6.

## 4. Surfaces

Row set, shared by every surface, in this order (a row appears only when its source can read it):
1. Attunement (first: the one row that can end a night). Source Game, or the export off game.
2. Enchants, listing only the slots that take one (chest, bracers, gloves, boots, cloak, weapon, and shoulders or head where a hand-in kind exists). Weapon enchant is its own line. Directory characters: `Enchants unknown`.
3. Resistance, only for a raid whose encounter data names a resistance need for that role (encounters.json mechanics). Onyxia names none, so the row is absent tonight. Where shown: `Fire 21, needs 315 on Ragnaros tanks`, no target when the data has none, just the total.
4. Flask or elixir, Food, Protection potions by school (the snapshot supplies kinds). Game and export only.
5. Encounter items per boss (`encounterItems`, snapshot! v2) where the raid has any.
6. Gear gap (`gear_gap`, `item_level_delta` to the guild median) and unspent talent points (`talent_points_unspent` where `talent_points_read`).
Out of v1 and said once in the panel foot, not per row: class reagents, ammo, weapon stones and oils. The verdict therefore reads `All set on what we check`, never `All set`.
Gold only for a known failure. Unknown is muted. A raider who signed Late is reminded at start minus 1 h like everyone, with the line `Late` replaced by their own expected time when the sign-up has one, else skipped.

### 4.1 Discord DM (path a)

Sent at 7:00 PM to the linked Discord account of a raider who answered In or Late, who turned the reminder on, and whose addon has not run the in-game check in the last hour (the raider is online, the toast is enough; same rule for push and the Activity row). Built on the Discord bot OpenDM (api/internal/discordbot/rest.go) plus a new DM sender in notify (own lane, L-prep-2a, shared with the loot plan DM). One embed, class colour bar, circular crest thumbnail.

Verbatim, Thandrel, signed-in, gear 3 h old, upload 4 days old:
- Title: `Thandrel, Onyxia in 1 hour`
- Line 1: `Wed 8:00 PM · Olympus XXVII`
- Field `Attuned`: `Yes` (or `Not yet` in gold with the `Where` line of the ready check, for example `Quest: Drakefire Amulet, Orgrimmar`).
- Field `Enchants` (signed-in): `Boots not enchanted` plus `Where: Enchanting, 8 Boots`; none to fix: `Chest, bracers, gloves, boots, cloak, weapon enchanted`.
- Field `Gear`: `Item level 61, 2 below the guild median` (from the export or profile).
- Field `Bags`: present only when the upload is under 24 h old: `2 flasks, 1 food`. Older or none: the field is absent and the footer carries the single sentence.
- Footer: `Blizzard 3 h · bags 4 days · off: Settings`. When the Bags field is absent: `Blizzard 3 h · bags not checked until you log in`.
- Button (link): `Open prep`.
Each age appears once. Directory character (40 h old profile): Gear `Item level 59`, Enchants `Enchants unknown`, footer `Public profile 40 h`. No upload: footer `bags not checked until you log in`. No gear at all: one line, `No gear on file for Thandrel. Log in at 7:00 PM and the addon checks.`
States: sent; DM closed (Discord 50007): no retry, Settings shows `Discord DMs are closed` once; not linked: skipped; the guild's channel gets nothing. The raider sees only their own data.

### 4.2 Phone push and Activity (path a)

- Push title `Onyxia in 1 hour`, body `Thandrel: 1 to fix. Boots not enchanted.` Nothing to fix: `Thandrel: nothing to fix on file.` Never names another raider. Tap opens the Guild tab raid page at Your prep.
- Activity tab: the same text as one row, stamp `7:00 PM`, gold when there is a fix, muted otherwise. Written even when no push is delivered (decision 5). Skipped, with the DM, when the in-game check ran in the last hour.
- Android: FCM when credentials are present. iOS push is off in production (no APNs keys). The iOS Settings row reads `iPhone alerts are off.`
- Setting, both apps and web Settings: `Prep reminder, 1 h before a raid`, default off. The first raid after a raider links Discord shows the toggle on the raid page's Your prep block once (`Remind me on Discord` link, no DM until chosen). States: on, off, disabled `Link Discord or install the app`.

### 4.3 Web guild raid page (paths a, b, c)

Components from web/src/components/guild: GuildRaid.svelte hosts the block; rows reuse ReadinessBoard.svelte and GuildReadiness.svelte, RaidOfficerRow.svelte, FactLine.svelte, the Tag tones, the circular crest, item icons from `withIcons`. No new control.

Member block `Your prep` (T-3h to raid start, per own character in the raid):
- Head: crest, `Thandrel`, `Fury Warrior 60`, `Item level 61 (-2)`, verdict `1 to fix` (gold) or `All set on what we check` (muted).
- Rows: the row set above, fix rows first, right-aligned age on each row only where it differs from the foot (`now`, `4 d`).
- Foot, once: `Blizzard 3 h · bags 4 days` or `Checked 7:00 PM`, and `Not checked: reagents, ammo, oils`.
- Raiders see counts only: `31 of 38 checked`.

Officer block `Prep` above Who is coming:
- Count line, split by source: `12 in game · 6 relay · 13 profile · 7 could not be checked`. Never one merged number. Tiles under it, counts only: `Attuned 36 of 38`, `Flasks 22 of 31 checked`, `Food 28`, `Enchants 25`. Phone: 2 by 2 wrap.
- Relay line: `Relay online: Kaelra`, or `No relay online.` with the action `Ask a raider with the companion to log in`. Last send: `Sent 7:04 PM`.
- Controls: sort (fixes first, name, class), class filter, `Copy as text` for the officer channel. Phone: sort and copy in the overflow.
- Group `Checked`: crest, class-coloured name, what is off, one stamp (`7:02 PM`). Source in the row tooltip and screen reader label (`In game`, `Relay Kaelra`, `Upload 4 d`, `Blizzard 3 h`). A directory-only row shows `Enchants unknown`, never a failure.
- Group `Could not check`: name and reason from facts about the raider's data and the stamps, not liveness: `Profile is private`, `Not signed in, no addon seen`, `Not sharing with guild`, `Opted out`, `Not seen since 6:10 PM` (the last hello or snapshot! inside the raid window, else `Not seen`). No red.
- Directory rows: default `Could not check: Profile is private or not signed in` unless the guild setting `Use public profiles for prep` is on (default off; same privacy edit as the loot plan and recruiting, land it first).
- Mobile 360: two-line rows, both groups always open, stamps never truncate.
States: loading (skeleton rows 40 px, count line 20 px reserved), before T-3h (absent), no signups (`Nobody has signed up`), partial (a source older than 24 h is muted and its row reads `Check in game`), stale (`Blizzard 31 h` muted), error (`Could not load prep. Try again`), no relay, viewer without companion (foot reads `bags not checked until you log in`).

### 4.4 Addon raider view (path b)

At T-1h, the addon with a raid time from the inbox or a GUILD announcement, out of combat, after a jitter of 0 to 120 s, reads its bags (RaidPrep.read) and shows:
- A static toast under the minimap button: gold border, fades in over 0.2 s and out after 8 s, no pulse, no chat line, no sound: `Onyxia in 1 hour: 2 to fix` or `Onyxia in 1 hour: all set on what we check`. In combat it waits for combat to end. Click opens the Prep panel.
- Prep panel (Guild tab, above the raid line): head `Prep for Onyxia · Wed 8:00 PM`, header figure `Item level 61 (-2)`, verdict, the row set with icons, quality colour on item names, counts, `Where` on fix rows. Foot `Checked 7:00 PM`, `Not checked: reagents, ammo, oils`. Button `Check again` (10 s cooldown).
- Raid time from an announcement: foot adds `Raid time from Kaelra`. No raid time: `No raid time known. Run the companion or wait for an officer's addon.` and no toast. Not sharing: foot adds `Not sent to the guild`.
Keyboard: opens from `/fs`, rows focusable, Escape closes.

### 4.5 Addon officer list (paths b and c)

Officer.lua readiness gains section `Prep` from T-1h: the split count line (`12 in game · 6 relay · 13 profile · 7 could not be checked`), tiles as text rows (`Attuned 36 of 38`, `Flasks 22 of 31`), the relay line, group `Checked` (name, what is off, stamp, source on hover), group `Could not check` (name, reason as in 4.3; `Not seen since 6:10 PM` comes from the last hello or snapshot! inside the raid window, never from the one hour SEEN_FOR expiry, which is not used), a class filter and `Copy as text` (opens a selected edit box). A relay also sees the button (section 6.3). Roster is the sign-ups; Felsnap never appears. A row keeps its last check and stamp, never blank.

## 5. States of every interactive element

| Element | default | hover | focus-visible | active | disabled | loading | empty | error |
|---|---|---|---|---|---|---|---|---|
| Open prep (DM link, web) | gold text link | underline | 2 px gold ring | pressed ink | n/a | n/a | n/a | n/a |
| Prep reminder toggle | track by state | track brighter | 2 px ring | thumb moves | greyed, caption `Link Discord or install the app` | spinner replaces thumb while saving | n/a | thumb returns, `Could not save. Try again` |
| Check again (addon) | `Check again` | highlight | ring (keyboard binding) | pressed | `Checked 7:00 PM` for 10 s; `In combat` while in combat | n/a | n/a | `Could not read bags` |
| Send tonight's prep (addon) | `Send tonight's prep` | tooltip below | ring | click starts the reload | `In combat`, `Ready check running`, `Loot session open`, `No companion seen`, `Sent 7:04 PM` | the client's loading screen | n/a | `Reload refused by the game. Type /reload.` |
| Retry (web) | the shared retry | as shared | as shared | as shared | n/a | skeleton | n/a | n/a |
Touch parity: every web control is 44 px tall on the phone; the addon has no touch surface.

## 6. Relay design (owner approved)

Why: the addon makes no network calls. Data leaves the game only through SavedVariables written at logout or reload, read by a companion. Addon messages (FS1) work only between players online in game.

### 6.1 Wire (Comms.lua, 255 bytes, never split)

Reuse `snapshot!` v2 for the raider's check (it carries addonVersion, attuned, encounterItems, supplies; already at most one message, Comms.encode refuses over 255). Add three kinds to Comms.FIELDS and Comms.CHANNELS:

| Kind | Fields in wire order | Channels | Limit | Meaning |
|---|---|---|---|---|
| `relay!` | relayVersion, raidId, startsAt (epoch) | GUILD | 300 s per sender | I run the companion. Also announces the raid time. |
| `prep?` | raidId | GUILD, RAID | 60 s per sender | Relay asks for checks (late arrivals). Answered by `snapshot!`. |
| `prep!` | raidId, checkedAt (epoch) | GUILD, RAID | 10 s | Marks the sender's next `snapshot!` as the T-1h check, so the relay stamps with the sender's own time. Sent right before `snapshot!`. |

A raider addon (sharing on, per Comms.sharing) sends `prep!` then `snapshot!` once at T-1h plus jitter 0 to 120 s, and again on a received `prep?` (jittered 0 to 20 s). 38 raiders in 120 s is 76 messages, under the client's addon-message throttle when each addon spreads its two by 1 s. Receivers drop a kind older than a minute of clock skew (checkedAt compared with GetServerTime).
Trust: one rule. `relay!` and `prep?` are accepted only from a sender who passes Comms.mayPush (officer rank or group leader). A raider who is not an officer is never a relay. A `relay!` raid time never overwrites the companion inbox time. Messages from a sender outside the guild are dropped (drop(), counted in /fs diag).
Nothing is written to a visible chat frame: addon messages are not chat, the addon prints no line for them, the toast is local.

### 6.2 Relay behaviour

- A player is a relay when the addon sees companion data in its own SavedVariables (an inbox message younger than 24 h) and the character passes Comms.mayPush. It announces `relay!` on GUILD every 300 s while the raid is within 3 h.
- The relay's addon collects every `snapshot!` after a `prep!` into its collection, key = character name-realm, keeping the newest `checkedAt`.
- Several relays online: each collects independently. Dedupe at three levels: (1) per character the newest stamp wins, in the collection and on the server; (2) one primary relay is chosen deterministically: the earliest alphabetical name among relays heard in the last 10 min; only the primary's button is the prominent one (gold), others read `Another relay is online: Kaelra` and the button is secondary; (3) the server accepts uploads idempotently by (raid, character, checkedAt), so a second upload changes nothing.
- No relay online: officers read `No relay online.` with the action `Ask a raider with the companion to log in` (web and addon). Off-game rows still show.

### 6.3 The button and the reload (v1)

`Send tonight's prep` shows in the addon officer panel and the raider Prep panel for a relay only (a character passing Comms.mayPush), from T-1h to raid start. Tooltip: `Reloads your UI so the companion can send everyone's checks. Out of combat only.` The collection holds only rows from raiders whose own sharing setting is on (Comms.sharing; a sender with sharing off sends no `snapshot!`, and a row without a `prep!` stamp is dropped). The outbox kind `prep_checks` carries {raidId, relay character, rows {name, checkedAt, supplies, attuned, encounterItems, version}}. The companion posts it to `POST /v1/guilds/{guild}/raids/{raid}/prep-checks`; the server accepts rows only for characters in the raid and rechecks each character's consent before storing, dropping the rest. After the reload the addon shows `Sent 7:04 PM`. Disabled also while a summon is running or the raid is mid-pull (boss engaged).

Rules, all disabled states named in section 5. ReloadUI spike (build prerequisite, 0.5 day, in the Forever client on the Windows box): does ReloadUI taint-block from addon code outside a click, is it protected in combat, does it need a hardware event. The button is a click and so is a hardware event either way; the spike decides whether an automatic reload is ever possible. Decision recorded: v1 ships click only, no automatic reload at T-1h, ever, until the spike and an owner call both say yes.

### 6.4 Verification task: does WoWChatLog.txt carry addon traffic

Run on the Windows test box over ssh jh@192.168.1.182 (not the Mac), a client with two characters. Steps: `/chatlog` on, send an FS1 message on GUILD from character A to B, then on a private channel, then close the client and read Logs/WoWChatLog.txt. Record: whether CHAT_MSG_ADDON lines appear, which channels, the time the line is flushed to disk (buffered until logout or per line). Results go to design/reviews/2026-11-15-guild-diff/05-chatlog-verification.md. If yes the companion could tail the file for a live path without a reload; build lane L-prep-4 then replaces the button with a tail. Notes:
- ToS: the companion already reads WoWCombatLog and SavedVariables; reading WoWChatLog.txt is the same class of local file read, but /chatlog is a player action the addon cannot toggle (no automated `/chatlog` by addon: that would be a macro of a protected command and could draw a ToS question). The player must turn it on; the setup text says so.
- Spam: do not create an extra channel and never send prep traffic to a visible channel to make it loggable. A joined custom channel shows in chat frames; excluded.
- Privacy: the file holds every chat line, whispers included; the companion must filter to FS1 lines and discard the rest at read, never upload or store the rest. This is an owner decision if the answer is yes.

## 7. Phone and desktop behaviour

- What collapses: on phone the chips wrap to a second line; the Raid count and Relay line stack. What never collapses: the verdict line, the stamp and source of every row, the Cannot check group.
- Desktop 1280: member block 2 columns (rows left, stamp column right, 96 px); officer groups side by side at 1280 and above only when each has 8 or more rows, else stacked. At 2000 px every container stays centred, no 1-up row stretches past its card width.
- Tokens: type scale of the Readiness board (name 14/20 semibold, chip 12/16), spacing 8, 12, 16, row height 40 (desktop) 56 (phone two-line), quality colours on item names only, class colour on character names only, gold for to fix, muted for old or unknown, red never used here (nothing off game is a known failure except a signed-in enchant missing or an in-game missing flask).

## 8. Performance and polish

No layout shift on hydration: reserved heights (count line 20 px, rows 40 px by the skeleton count from `viewer.raid.signed_up`), icons come from `withIcons`, so no missing-icon flash; the block is not in the first paint on non-raid days. Lighthouse budgets in web/lighthouserc.json hold: the block is server-rendered, one island, no new font. The reminder job sends 38 DMs in under 60 s (Discord rate limit: 1 DM open per second per bot in bursts, spread by the job at 2 per second). The addon reads bags once per 10 s at most.

## 9. Officer objections (02-officer-plan.md, section 3)

1. Public gear without guild consent. Directory gear is used for a raider's own DM and push only. In the officer list it is dark by default (`Profile is private or not signed in`) until the guild setting `Use public profiles for prep` is on (default off, same privacy edit as the loot plan and recruiting, landed first). Opt-outs remove the row. The relay path honours each raider's own sharing setting and the server rechecks consent.
2. Mixed units. Not in this spec: the prep carries no gain numbers, so no ranking in score, HPS or DPS.
3. Naming after wipes. Not applicable to prep, but the same rule holds: names in the Cannot check and fix lists are officer-only; raiders and the Discord channel see counts.
4. The reminder promises more than the data. Answered by: the DM and push carry Blizzard data and a bag list only when the upload is under 24 h old, each age shown once; enchants from the directory read `Enchants unknown`; the count is split by source so 31 is never read as ready; the verdict is `All set on what we check`; iOS push is off and said so; the raider's own in-game check ships first (L-prep-1), the officer Cannot check list second (L-prep-2), the relay third (L-prep-3), gated on the ReloadUI spike, bag-freshness evidence and one real raid where the officer list matched /ready check. A bag list over 24 h is dropped from the DM and muted on the web (`Check in game`).
5. Loot fairness: not this spec.
Shame list (section 4 of the review): the count to the raid, names to officers. The Cannot check chips are reasons (`Not seen since 6:10 PM`, `Not sharing with guild`, `Profile is private`), never `slipping` or `not ready`; the raid-facing line is `31 of 38 checked`; the Discord channel gets nothing; a push or DM body names only the recipient. A raider with no data is `Could not check`, not `Not ready`.

## 10. Copy, all strings

Web and app: `Your prep`, `Prep`, `All set on what we check`, `{n} to fix`, `{n} in game · {n} relay · {n} profile · {n} could not be checked`, `{n} of {m} checked` (raiders), `Checked`, `Could not check`, `Relay online: {names}`, `No relay online.`, `Ask a raider with the companion to log in`, `Boots not enchanted`, `Enchants unknown`, `Gear not on file`, `Item level {n} ({delta})`, `{n} unspent talent points`, `Check in game`, `Blizzard {age} · bags {age}`, `Public profile {age}`, `Checked {time}`, `Bags not checked until you log in`, `Not checked: reagents, ammo, oils`, `Not seen since {time}`, `Profile is private`, `Not signed in, no addon seen`, `Not sharing with guild`, `Opted out`, `Copy as text`, `Nobody has signed up`, `Could not load prep. Try again`, `Discord DMs are closed`, `Prep reminder, 1 h before a raid`, `Remind me on Discord`, `Link Discord or install the app`, `iPhone alerts are off.`
Addon: `Onyxia in 1 hour: {n} to fix`, `Onyxia in 1 hour: all set on what we check`, `Prep for {raid} · {day} {time}`, `Check again`, `Send tonight's prep`, `Another relay is online: {name}`, `Raid time from {name}`, `No raid time known. Run the companion or wait for an officer's addon.`, `Not sent to the guild`, `In combat`, `Ready check running`, `Loot session open`, `Sent {time}`, `Reload refused by the game. Type /reload.`, tooltip in 6.3. Ages: `now`, `{n} min`, `{n} h`, `{n} days`. One clock (`8:00 PM`, the shared locale clock).

## 11. Acceptance screenshots (design/reviews/2026-11-15-prep-reminder/)

1. Web raid page member block, 360 and 1280, states: to fix, all set, directory-only (enchants unknown), stale bags 4 d, loading, error.
2. Web officer block 360, 1280 and 2000: relay online, two relays, no relay, 31 of 38.
3. Discord DM render, three variants (signed-in, directory, no gear), desktop and phone client.
4. App Activity row and push on Android, iOS Settings row with the not-on sentence.
5. Addon raider panel and toast, in-game capture at 1080p: to fix, all set, no raid time.
6. Addon officer list: both groups, no relay, button disabled in combat.
7. Verification notes for the spike and the chat log task (text, with transcripts).

## 12. Build lanes

| Lane | Work | Size |
|---|---|---|
| L-prep-0 | ReloadUI spike and WoWChatLog.txt verification on the Windows box | S, 1 day |
| L-prep-1 | Addon raider check at T-1h, static toast, Prep panel, row set, raid time announcement (`relay!` receive only) | M, 3 days |
| L-prep-2a | notify DM sender (OpenDM, channel cache, closed-DM handling), shared with the loot plan DM | M, 2 days |
| L-prep-2 | API PrepReminderOf, `prep.reminder` job (skips when the in-game check ran), web member and officer blocks, tiles, sort, filter, copy, Settings toggle | L, 5 days |
| L-prep-4 | Phone push, Activity row, app Settings row | M, 2 days |
| L-prep-3 | Relay: wire kinds, collection, election, outbox, endpoint with consent recheck, officer list, button. Gated on L-prep-0, bag-freshness evidence, and one real raid where the list matched /ready check | L, 6 days |
| L-prep-5 | Chat log tail, only if L-prep-0 says yes | M, 3 days |
Order: 0; then 1, 2a, 2 in parallel; then 4; then 3.

## 13. Open owner decisions

1. DM and push default: off (this spec) or on for In or Late raiders after they link Discord.
2. Whether directory gear may appear in the officer list at all (the setting in 4.3, default off).
3. The privacy edit that gates directory rows lands first, shared with the loot plan and recruiting.
4. Whether an automatic reload at T-1h is ever allowed after the spike (recommended: never).
5. Whether the Activity row is written when iOS push is off (recommended: yes).
6. If WoWChatLog.txt carries addon traffic: a player-enabled `/chatlog` tail with filtering, or stay on the button.
7. Whether the relay gate (one real raid matched against /ready check) is the owner's bar to build L-prep-3.

Review fixes (2026-11-15): relay trust is Comms.mayPush only; relay rows honour each raider's sharing and the server rechecks consent; no `all slots enchanted`, only enchantable slots, directory reads unknown; reasons use the hello or snapshot! inside the raid window, not the 1 h liveness; count split by source; DM default off, own sender lane; attunement first, resistance only where a raid needs it; stale bags dropped from the DM; DM and push skipped when the in-game check ran; reagents and ammo stated out of v1, verdict softened; gear gap, talents, weapon enchant, protection potions, encounter items added; officer tiles, sort, filter, copy; static toast; ages shown once; no-relay line shortened.
