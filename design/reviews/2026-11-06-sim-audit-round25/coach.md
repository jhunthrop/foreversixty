# Sim-accuracy round 25: coach lane

Scope: the four coach seeds (guide-accuracy round 26, "Routed (coach code)"), then a fresh re-audit of the F1 and F7 code they touch for every class's talent-changed cooldowns and defensives. Site api only; no engine change, no sim runs needed (no DPS moves). Branch wt/r25-coach on int/r25 (ab2029586).

## Seeds

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **F1 judged talent-shortened cooldowns at base.** A Fury warrior from 37 has an 8 s Whirlwind (Improved Whirlwind 2), a Marksmanship leveler a 3 minute Rapid Fire from 26 (Rapid Killing 2), a Beast Mastery leveler 4 minutes from 57. F1 said "Measured against Whirlwind's 10-second cooldown" and counted possible casts from 10 s, so the stated cooldown was false and the lost count too low. | `f1_core_casts.go` `assessAbility` used `sc.CooldownMS`; client talents/warrior.json "Reduces the cooldown of your Whirlwind ability by 1 sec / 2 sec", talents/hunter.json Rapid Killing "by 1 min / 2 min"; HUNTER-R26-6, WARRIOR-R26-9 | CONFIRMED-BUG | old (coach F1, 2026-10-28), guides-handoff | S 790e8a11a |
| C2 | **How the fix reads talents.** The log never carries talents (`CannotTalents`). The character's stored build does: `addon_exports` (the addon export or the Battle.net import, whichever is newer). The coach now decodes its FS1 tree strings against the talent trees of the export's own data build (tier then column order, the format's digit order), only when the export is the row's class and its build is a version string. A table of every talent that shortens a judged cooldown feeds F1, its ceiling and F7's ready check; each value is checked rank by rank against the client text (`TestCooldownTalentsMatchTheClientText`) and every family against the rule files. With no build on file, F1 keeps the base cooldown and the basis says so: "...base cooldown... Your talents aren't on file, and a talent that shortens it would only add to the lost casts." The talents are today's, like the level the coach already uses. | `talent_effects.go`, `talents_export.go`, `service.go` `talentsOf`, `cmd/api/coach.go` | (fix) | - | S 790e8a11a |
| C3 | **F7 offered Evasion at spell deaths.** Evasion (5277) is one effect, aura 49 (dodge). F7 named it first at nearly every rogue death, fire and shadow included. Now a spell whose every client effect is a dodge or parry aura is melee-only, and F7 skips it unless the killing hit was a melee swing or an ability the player's damage-taken row records as physical school. An unknown school is not proven physical. Vanish is still offered. | spellconst/rogue.json 5277 effects `[(6, 49)]`; ROGUE-R26-11; pages "does nothing against spells" | CONFIRMED-BUG | old, guides-handoff | S 3dd071bf1 |
| C4 | **Shadowburn shard floor.** Shadowburn costs a Soul Shard and the log shows no shards, so a warlock who ran dry was told the shard-less Shadowburns were lost "while you had the mana". F1 now judges Shadowburn only where every cast refunds its shard: Shadow and Flame 5/5 (426316, 20% a rank) on the build on file, or, with talents unknown, from level 44. Both Destruction orders fill it by then (raid 42, leveling 44); a test reads addon-builds.json and fails if the orders move. Unknown level and unknown talents: not judged. | `TestCoreCasts_ShadowburnNeedsTheShardRefund`, `TestShardGateLevelIsWhereTheShippedOrdersFillTheTalent`; WARLOCK-R26-11 | CONFIRMED-BUG (the guide lane filed it LIKELY; the false sentence is certain whenever the bag runs out below 44) | old, guides-handoff | S 790e8a11a |
| C5 | **`TestEveryReportVisibilityConditionIsTheEvidenceRule` red on int/r25.** Two Go comparisons in reports/handler.go: `updated.Visibility != rep.Visibility` (:539, the patch starting the visibility hook) and `s.Visibility == nil` (:569, the hook's nil check, not a report's visibility). Both came in eaecf8cc3 (2026-11-04, raid attendance recompute) without an allowlist entry. Neither reads evidence. The handler is right; the test's list was stale. Added both with their reasons; green. | blame: handler eaecf8cc3, test last 601052dcc | CONFIRMED-BUG (test) | old (eaecf8cc3) | S 46d8a3db2 |

## Re-audit: F1 and F7 for every class

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Fire mage Fire Blast judged at 8 s; Improved Fire Blast 2 (25 and 26 on both Fire orders, 55 and 56 on Arcane raid) makes it 6 s. Same shape as C1. | talents/mage.json "by 1 sec / 2 sec"; spellconst 10199 category 8000 | CONFIRMED-BUG | old | S 790e8a11a |
| A2 | Retribution Judgement judged at 10 s; Improved Judgement 2 (15 and 16, both orders) makes it 8 s. | talents/paladin.json "by 1 sec / 2 sec"; spellconst 20271 10000 | CONFIRMED-BUG | old | S 790e8a11a |
| A3 | Protection paladin Holy Strike judged at 10 s; Shieldsworn (40, both orders) makes it 6 s. | talents/paladin.json "by 4 sec"; spellconst 10333 category 10000 | CONFIRMED-BUG | old | S 790e8a11a |
| A4 | Deterrence (19263, auras 47 parry and 49 dodge) is melee-only like Evasion; F7 offered it at spell deaths to every hunter spec. Fixed by C3's data-derived flag. | spellconst/hunter.json 19263 | CONFIRMED-BUG | old | S 3dd071bf1 |
| A5 | F7's ready proof used base cooldowns for Sacred Duty (Divine Shield, Divine Protection, Templar's Bulwark; Holy 51-52, Protection 19-21), Survivalist's Discipline (Deterrence, 37-38), Endurance (Evasion, Assassination raid 52-53), Elusiveness (Vanish, Subtlety 25 and 33) and Improved Shield Wall. That proves "ready" later than the truth, never earlier, so no false claim, but missed true ones. Now read from the build on file. | rogue round 26 D7 | Aligned (conservative); improved | - | S 790e8a11a |
| A6 | Cooldown resets the coach cannot see: Preparation, Cold Snap (Ice Block, Ice Barrier), Improved Stormstrike's reset on your own dodge or parry, Lightning Mastery. Each makes F1/F7 lenient, never false. | talents/*.json | Aligned (lenient) | - | - |
| A7 | Every other rule-file family checked for a shortening talent (all talents whose client text names a cooldown): none for Bestial Wrath, Arcane Power, Presence of Mind, Combustion, Cold Blood, Adrenaline Rush, Blade Flurry, Ghostly Strike, Stormstrike, Conflagrate, Bane of Doom, Mortal Strike, Bloodthirst, Death Wish, Bloodrage, Shield Slam, Inner Focus, Mana Tide Totem, Primal Bite, Raptor Strike, Strider Kick, Holy Shield, Barkskin, Frenzied Regeneration, Ice Block, Ice Barrier, Lay on Hands, Death Coil, Last Stand, Feign Death. Hotfix-push carries no row for any table talent. | talents/*.json, hotfix-push.json | Aligned | - | - |
| A8 | Other defensives against spell deaths: Shield Wall (aura 87), Barkskin, Ice Block, Ice Barrier, Divine Shield and Protection (school immunity), Templar's Bulwark (absorb), Death Coil, Last Stand, Frenzied Regeneration, Vanish and Feign Death all help against a spell death; only Evasion and Deterrence are dodge or parry only. | spellconst effects | Aligned | - | - |

## Before/after

No DPS moves: site api only, no engine, APL, build or card change. F1/F7 behaviour is covered by the new tests (each failed first: the three F1 talent/shard tests on the old `f1_core_casts.go`, the two Evasion tests before the F7 change).

## Files to regenerate

None. RulesVersion is now coach-rules-v3, so every cached coach pack rebuilds once on its next view and its review regenerates once (model cost, one per viewed pack). Coach profiles do not filter on the version.

## Measure at refresh

None.

## For other lanes

- Integrator/platform: `cmd/api/coach.go` now wires `coach.ExportTalents{DB: pool, BuildsDir: TreeDataDir}`. A talents read reads one `addon_exports` row per built pack. The coach decodes talents only for data builds present in the image's TreeDataDir; an export from a build the image lacks reads as "talents unknown".

## For the guide lane

- The coach F1 basis can now read "...base cooldown... Your talents aren't on file..." (copy.go `BasisCooldownBase`), shown only when a talent could shorten that cooldown and no build is on file.
- Suggest the coach page tell players that importing their character (Battle.net or the addon) lets the coach use their talents' cooldowns.

## Still open

- Talents are read at pack-build time, so a respec between the fight and the first view uses today's talents (same rule as the level). Exact talents per fight need talents in the log, which Forever does not write.
- Fury execute hold (WARRIOR-R26-10) and Bloodrage ceilings (WARRIOR-R26-11) are LIKELY and outside this round's seeds; not touched.
- Logs: GET https://api.foreversixty.gg/v1/reports/recent at 2026-11-06 16:43 UTC lists only the 20 Oct test report; nothing a log could settle here.
