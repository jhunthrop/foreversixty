# Officer gate on the four guild specs (2026-11-15)

Spec gate, no screenshots. Read as the raid leader and as a raider, against 00-facts.md, my plan review (02) and the code. Specs are not edited.

## Summary

| Spec | Verdict | Plan objections answered |
|---|---|---|
| loot-plan | SHIP WITH FIXES | 1, 2, 5 yes; 3, 4 correctly out of scope |
| live-pulls | SHIP WITH FIXES (site and Discord); addon block deferred | 3 yes; consumables yes |
| prep-reminder | SHIP WITH FIXES for lanes 0 to 2 and 4; relay (L-prep-3) NOT SHIPPABLE yet | 4 only half answered |
| guild-later | SHIP WITH FIXES (K1, K3, R1 only) | consent flag honest |

## 1. Loot plan: SHIP WITH FIXES

Objections: the three big ones are answered. Directory gear is off by default behind D1, a per-guild setting and directory_optouts (section 4). Ranking is per role with bands and an `estimate` caption (rules 1 and 2). Every row carries awards this tier, last-4 attendance, MS/OS, source and age, and Drop, Swap and Lock are one click with undo (5.1). Raiders get only their own rows, stripped server-side (rule 7). Recommendation wording everywhere. This is the page that replaces my sheet.

Must fix:
1. Section 3 rule 1 says role is `roleOf(row)` with the sign-up role winning. loot_candidates.go roleOf returns only the row's logged role. Either name the sign-up override as new code in L1 or drop the claim. A tank signed up as DPS must land in the right group.
2. Raider status `Likely upgrade` needs "band above under 10", but tank rows have no band (rule 2) and D3 (the tank floor) is open. Tanks and healers without a number get no status. State the tank rule or ship tank status as `No estimate` until D3 is answered. Same hole in the Discord and addon strips.
3. Rule 3 sorts inside a band by fairness, but Lock "pins at its current rank" and Swap exchanges with the row above. Define what Swap does across a band or tie boundary and whether a locked row can sit above a higher band. Without it the council will see a ranking that contradicts the stated sort.
4. Rule 4 and "Hide dropped" imply a directory row can be above a fresher row in a lower band. Say so in the caption or add the age flag to the band edge; state it once: the band wins, source breaks ties only within a band.
5. Section 9 `Won 0 this tier` is only true if the guild has any award on file. Good. Also say what "this tier" means for a guild that imported history; the council counts ItemsWon per tier already, so cite that field.
6. Section 4 gates signed-in Blizzard gear on "gear consent, and not for a member who set guild sharing off". Name the field. If no such flag exists for Blizzard-source exports, that is new work in L3, not "yes by default".
7. Plan marks: `loot_plan_marks` must survive a recompute when the raider signs off the raid. Say what happens to a lock on a raider who withdraws.

Would not trust: the gain band on a hit-capped or set-bonus item (the spec does say estimate; keep it on every surface including the app). A raid-channel post that lists roles per drop still tells raiders who might win; acceptable because no names.

Needs data, not today: hit cap awareness, a tank floor number.

## 2. Live pulls: SHIP WITH FIXES

Objections: naming is answered well (section 3: officers named, raiders own lines through the existing ownOnly, Discord counts only, lock screen nobody, `Name them` per post). Consumables are suppressed behind `consumablesProven` flipped by a golden test, which is exactly what I asked. Latency is honest: only the 1 s tick and 5 s live cadence are marked verified; the rest is labelled estimate and the page claims no seconds until L0 measures. The addon block is split out as its own product (D1). Several companions: earliest wins with a heartbeat takeover; Heartbeat exists in companion/internal/client/ingest.go.

Must fix:
1. Section 2 ranks deaths, downtime and rotation in one "lost seconds" scale and shows the top three. Downtime and rotation are the lines that make a raider feel accused. Rotation is estimates from weights (D7). Ship deaths-by-mechanic only in v1; hold downtime and rotation until D7 is decided and a real log has been read by an officer. Section 2 also says downtime is "never a ranking" but ranks it among costs. Pick one.
2. Section 1 puts `Thandrel died first at 2:41 to Deep Breath` in the first-screen answer. Fine for officers; confirm the raider response replaces it with `a raider` server-side and add that to the acceptance test (section 12 has the fixture, good).
3. Section 7 rule 2: a second companion's report is dropped from the live list. If the first companion stands out of range of the raid (a healer in the back), pulls show missing deaths with no warning. Add a visible `Coverage` line only if the report can say how many raiders it logged; if it cannot, say `Deaths seen by Kaelra only` in the meta. Do not claim a full raid.
4. 5.5 states no live read route exists and the poll is not built. Make L2 own ETag and the 304, and make the first lane gate the page, since "10 s, ceiling 20 s" is only a target.
5. Section 5.3 names a Discord checkbox `Name them` per post. Confirm it needs officer role server-side, not a hidden control.

Would not trust: rotation lines at all until a spec audit. The Night summary line `Most costly raider-side loss: downtime, 5 raiders over 15 s` reads as a count of slackers; remove it or give it the same hold as downtime.

## 3. Prep reminder: SHIP WITH FIXES (L-prep-0, 1, 2, 4); relay NOT SHIPPABLE as written

Objection 4 asked: ship the member in-game check and the officer Cannot check list first, drop the relay and reload button until the ReloadUI spike answers, show bag age plainly. The spec puts the order right (lane order 0, 1 and 2, then 3) and shows ages everywhere, `Not a live check`, and `Enchants unknown`. But it still builds the relay as lane 3 on an approved design, and the spec's own gate sentence says "bag age under a day is not a gate".

Must fix:
1. Section 6.1 trust rule is garbled: "accepted only from a sender who passes Comms.mayPush (officer rank or group leader), or whose character the roster shows as a raider in the guild and who is also an officer's relay". Comms.mayPush (Comms.lua line 676) is officer rank or group leader only. Pick one rule. Letting any raider be a relay lets any raider inject prep rows for others. Rule: the relay must pass mayPush, and the server accepts `prep_checks` rows only for characters in the raid.
2. Section 6.3 uploads rows `{name, checkedAt, supplies...}` that other raiders sent over FS1. A raider's bags go to the site through someone else's reload. That bypasses their sharing setting unless the sender honoured Comms.sharing. Say the relay drops rows from anyone not sharing, and the server rechecks consent before storing.
3. Section 4.3 officer block shows directory-source rows under a per-guild setting, but 4.3 text says the Checked group shows `Gear 3 h · Blizzard` chips with no mention of the setting for the default. State the default: directory rows are `Could not check: Profile is private or not signed in` unless the setting is on. Also the loot plan and later spec both depend on the same privacy notice edit; link them.
4. 4.1 sends a DM at 7:00 PM saying `all slots enchanted` from a Blizzard read up to 20 h old. A swapped cloak at 4 PM reads wrong. Show the gear age in the field (it does in the footer; move to the field) and say `on file`, not `all slots enchanted`.
5. Section 4.5 Liveness uses Comms.seen with SEEN_FOR 3600: a raider silent for an hour reads `No addon`. At T-1h that is every raider who logs in at T-5 min. Use the check stamp, not liveness, for the reason chip, or the Cannot check list will name half the guild.
6. 4.2 and the DM: the DM sender does not exist (facts: notify has none for other kinds). It is in L-prep-2 but sized with the web block. Break it out; the loot plan DM (D4) wants the same sender.
7. Setting default: DM on for In or Late raiders (decision 1) is opt-out spam in a first-run guild. Make it off until the raider links Discord and has seen it once.

Would not trust: any claim of "checked" for a raider whose bags are four days old. The count line `31 of 38 checked` mixes game checks, uploads and Blizzard-only reads. Split it: `12 checked in game, 19 from uploads, 7 could not be checked`. As written an officer will read 31 as ready.

Relay: leave L-prep-3 gated on the ReloadUI spike, on bag freshness evidence, and on one real raid where the officer list matched /ready check. The button reloads the leader's UI at T-1h and the whole path is large for one list.

## 4. Guild later: SHIP WITH FIXES (K1, K3, R1 now; K2, K4, R2 to R4 wait)

Both are briefs, honest about what is missing.

Comp and bench: K3 (bench and swap rows from coverage alone) is useful and shippable. The catalogue shows Strength of Earth as `Per: melee_group` (api/internal/guilds/raids_cover_catalogue.go line 12) while the brief says it is raid-wide in Forever. K1 must come first and fix the catalogue, or the officer bench call is wrong on the first buff. Windfury and Grace of Air are group-scoped and there are no group assignments, so the spec correctly counts them per raid; label that on the row (`counted per raid, not per group`). Marginal DPS (K2, K4) is unverified; keep it out of v1.

Recruiting search: not built until R1 (privacy notice). One more fix: `Guild unknown` is honest, but a list that cannot say `no guild` will show characters already in guilds. Filter to rows the roster pass has scanned, or say `Guild unknown` on every row. Also the only action is a copied `/w` line; officers will accept that.

## What officers still would not trust

- Any number older than the last raid night, unless the row says so (the loot plan does, the prep DM does in the footer only).
- A recommendation that looks like a decision. Every surface says it is one; keep it in the app and the addon too.
- A first raid with no real Forever combat log: consumables and rotation lines stay suppressed until 9 Dec produces one.
- Four specs depend on the same unnamed privacy edit (privacy-additions section 2). Land that edit first; it gates loot directory gear, recruiting and prep directory rows.

## Value test

Loot plan with fixes 1 to 4 flips us from the sheet. Live pulls (deaths and mechanics only) replaces my after-raid Logs scrub for causes, not for parses. Prep reminder is a nice reminder, not a reason to pay, until the relay path proves out. Comp bench call from coverage is the second reason to open the raid page.
