# Officer re-gate, round 2 (2026-11-15)

Spec read only against my round-1 must-fixes (10-officer-specs.md). Specs not edited.

| Spec | Round 1 | Round 2 |
|---|---|---|
| loot-plan | SHIP WITH FIXES | SHIP |
| live-pulls | SHIP WITH FIXES | SHIP |
| prep-reminder | SHIP WITH FIXES, relay NOT SHIPPABLE | SHIP (relay gated, acceptable) |
| guild-later | SHIP WITH FIXES | SHIP |

## loot-plan (all 7 resolved)
1 Sign-up override named as new L1 code (section 3 rule 1, lane table). 2 Tank and healer bands are relative (`Top`, `Close`, `Smaller`) until D3; status words use them. 3 Swap across a band makes the moved row locked with a glyph; a lock prints its true band so the contradiction is visible. 4 Band wins, source only inside a band. 5 `this tier` is now `this phase`. 6 Consent gate named (`hasGearConsent`) and L3 checks it for Blizzard exports; Blizzard rows stay off until it lands. 7 Withdrawn raider keeps a muted lock, ignored by the ranking.
Watch (minor): rule text says Lock "pins a row above any band" while 5.1 says "at its current rank". Make them one sentence at build.

## live-pulls (all 5 resolved)
1 Downtime and rotation held behind D7, out of v1 and out of the Night summary. 2 Raider sees `a raider` for others, and a test asserts no other name in a member response. 3 Meta reads `Deaths seen by Kaelra only`, no full-raid claim. 4 Read route with ETag and 304 is L2 and gates the page. 5 `Name them` posts to the officer channel only. Nothing broken: latency still labelled estimate until measured.

## prep-reminder (all 7 resolved)
1 Relay trust is mayPush only; non-officers are never a relay. 2 Relay rows honour sharing and the server rechecks consent before storing. 3 Directory rows read `Could not check` unless the guild setting is on (default off), linked to the shared privacy edit. 4 No `all slots enchanted`; enchantable slots only, directory enchants unknown. 5 Reasons use the last hello or snapshot! inside the raid window, not the 1 h liveness. 6 DM sender split out as L-prep-2a, shared with the loot plan. 7 Reminder default off, offered once on the raid page. Count line is now split by source.
Relay: L-prep-3 is gated on the ReloadUI spike, bag-freshness evidence and one real raid where the list matched /ready check. The click-only button and no automatic reload are recorded. That meets my bar.
Watch: row 29 says Blizzard gear goes stale at 20 h. Confirm the DM footer shows the age on the gear field itself.

## guild-later (resolved)
Strength of Earth stated as code against owner, K1 verifies in game before K3 ships. Per-raid labels on bench rows (`counted per raid, not per group`). K2 and K4 stay out of v1. Recruiting says `Guild unknown` on unscanned rows and never `No guild`.

## Value test
Loot plan still flips us off the sheet. Live pulls on deaths and mechanics replaces my log scrub for causes. Prep reminder is a nice reminder until the relay proves out, and the spec gates it correctly.
