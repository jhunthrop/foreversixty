# Sim accuracy 2026-11-12: the healers' pushback and regen (Discipline, Holy, Restoration druid, Restoration shaman)

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `sim-healers2`. Site from origin/main 62e835b00, fork from master 4c595ee6b (the pin). Client on the site: 1.60.1.70790 with its hotfix push; live client 1.60.1.70888. Scope: the two classes of bug the Holy paladin lane found this morning (`holy-paladin.md`), checked across the other four healers at the bare bands 20 to 60 and the raid 60 presets: (1) pushback, what the heal profile's pulses do to each healer's casts, whether each healer's pushback-resist talent is requested and applied by the engine on the client's own mask, whether a party aura the engine drops is asked for, and whether the live server's pushback rule matches; (2) regen and procs, whether any mana return the four cards rely on comes from a client row the live server never casts, the way Illumination's 20272 did.

**Totals: 0 confirmed engine bugs. Neither class of bug repeats in the four healers.** (1) Every pushback talent the guide builds take (Twilight Focus 3/3 on both priests, Nature's Focus 5/5, Healing Focus 3/3) is wired into the engine on the mask the client gives the talent's own modifier (14913, 17063 and 16181 decoded against the three classes' family masks below), and the engine measurably applies it: emptying the talent costs the raid 60 healers Discipline -1.7%, the Restoration shaman -2.6%, the Restoration druid -0.4 to -0.5% and the Holy priest nothing (its bar is dry at 124 s of the 300, so a lost second of casting costs it no healing) and the bare 60 healers Discipline -2.2%, the druid -1.1 to -1.3%, the shaman -0.7 to -0.8% and the Holy priest nothing. No healer's request names a buff or aura the engine drops: the priest's Inner Fire and the druid's self Innervate are applied, the shaman asks for nothing, and Concentration Aura is in no healer's request and no raid preset (the proto has no party field for it; the ceiling a never-pushed-back healer reaches is +0.8% for Discipline, +1.2% for the shaman, +0.3% for the druid and nothing for the Holy priest at raid 60, open O1). The live server pushes a cast back about one second on the first hit, as the engine does: of 253 chained casts in the public logs with a creature hit inside the cast window, 20% landed 0.7 to 1.4 s late against 3% of the 3,381 without a hit, and the half-second bucket does not move (5% against 2%). (2) Every mana return the four cards rely on is a row the server casts or a talent's own aura: Litany of Light (the 1317006 talent itself, 10% of the base cost at 2/2), Inner Focus (14751, cost -100 and crit +25 on its own masks), Meditation, Reflection and Mindfulness (the talents' 17/33/50 while casting), Innervate (29166, aura 134 at 100 and aura 110 at 400, the engine's full regen at five times Spirit), Omen's Clearcasting (16870, logged), Water Shield (408511, logged), Mana Tide (16191 to 1325257, four ticks; the breakdown's five income events are the cost plus the four ticks), Dark Sacrifice (the 11 Nov lane), Blessing of Wisdom and Mana Spring (logged at 12/18/30 and 4/8 a tick). No priest, druid or shaman row opens a casting-regen aura from a spell the server never casts. 0 new data: Blizzard's posts through 12 Nov (hotfix thread posts 1 to 9, the tuning thread) carry no healer pushback or regen line beyond the three already landed (Twilight Focus in Shadowform, Omen after a spec swap, Nature's Swiftness on Wild Growth), and the 70888 rows of the three talents are the 70790 rows. Engine pin unchanged at 4c595ee6b; no build, card, rank row or guide line moves, so nothing goes to the "Measure at refresh" table and no wow-player pass was needed.**

## Method

1. **Blizzard since 10 Nov.** "WoW: Forever Hotfixes" (t/2403040) as the Discourse JSON, posts 1 to 9 (post 9 is 12 Nov 19:10 UTC, the Maelstrom Weapon tooltip line going live early), and "Class Tuning Incoming" (t/2404092, post 1 and the 19 replies the JSON carries). Every line naming a priest, druid, shaman, pushback, regeneration, Spirit or a healer spell was read; the matches are in the totals.
2. **Client and push.** wago.tools SpellEffect, SpellAuraOptions and SkillLineAbility rows for 1.60.1.70790 and 70888 with the push layered (`useHotfixes=1`) for the talents and spells below, their EffectSpellClassMask decoded against the site's `spellconst/{priest,druid,shaman}.json` family masks into spell names; the site's `talents/*.json` text for every talent the four builds take; `hotfix-push.json` (no row for any of them); the `gametables` folder (basemp, hppersta, combatratings, playerexpectedstat: no regen table, so the Classic formulas stand). The engine read line by line: `core/{cast,mana,raid_damage,buffs,periodic_action}.go`, `priest/{talents,healing_talents,mana_abilities,priest}.go`, `druid/{talents,talents_restoration,innervate,healing}.go`, `shaman/{talents,talents_restoration,water_shield,mana_tide_totem,shaman}.go`, and `sim/request/request.go` for what each healer's request asks.
3. **Real fights.** Our API's reports table, read through the Cloud SQL proxy read-only once the owner's gcloud login was renewed: seven reports in all, the five uploaded since 10 Nov (three private, two public) hold no fights at all, and the two complete ones (z2opt2x2gs5u, 20 Oct; gxrgvz4yzsi2, 9 Nov) carry the level 20 to 27 passers-by the earlier lanes read, so nothing there settles a healer question. foreverlogs.gg's 26 public reports (2815 to 2854, 5 to 12 Nov): every priest, druid and shaman with 1,500 or more events (27 characters, 86,000 events as source and 46,000 as target, `meta_events` 500 a page in both directions, cached in the lane's scratch). Nine heal: Luminare (2817, level 14 by ranks), Shadow (2819, 20) and Net (2821, 22) among the priests; Surgeon (2847, 20), Yingarna (2836, 18), Super (2823, 24) and Premo (2838, 26) among the druids; Swerds (2820, 20) and Tallak (2839, 40) among the shamans. No logged healer is a Holy priest with Litany of Light, a druid with Innervate or a shaman with Mana Tide, so those three rest on the client rows. The pushback measurement uses every caster's cast-time spells (heals and nukes, 19 players with hits inside a cast).
4. **Sims** on the Windows box through `tools/remote-test.sh` (`--dir sim`, `--pin-engine` at 4c595ee6b and `--engine-sha` for the scratch branch; `nice -n 15`, GOMAXPROCS=5, one job at a time; runner and job files in `healers-pushback-regen/`): `spec-breakdown -heal` for the four healers at raid 60 and bare 60 (both factions, seeds 1 and 50001) and the bare bands 20 to 50 (both factions) on the pin, the guide builds with the pushback talent emptied, and the published entries on a scratch engine that never pushes a cast back (fork branch `scratch/healers2-no-pushback`, 90c598cbc, never to be pinned). 10,000 iterations.

## Pushback: engine against the client, the request and the live server

| # | check | engine | client / request / log | verdict |
|---|---|---|---|---|
| P1 | The pulses reach the healer | `raid_damage.go`: three of the five (the healer among them) every 5 s, each landing hit calls `PushBackCast` | `heal-profile.json` pulse.reaches_healer true at every band | as designed (round 6) |
| P2 | The pushback itself | 1.0 s on the first hit of a cast, 0.2 s less each hit after, never under 0.2; a channel loses the same; one roll per hit against the spell's chance plus the unit's | mangos-classic `Spell::Delayed`; **the logs: see the table below** | aligned |
| P3 | Twilight Focus (14913, both priests 3/3) | 23/47/70% on Lesser Heal, Heal, Flash Heal, Greater Heal, Prayer of Healing, Penance, Smite, Holy Fire, Mind Blast, Mind Flay | effect 0 aura 108 misc 9 at 70, mask 0x00B63D10 / 0x00818484 / 0x80 / 0x400F0400 = Lesser Heal, Heal, Flash Heal, Greater Heal, Prayer of Healing, Penance, Binding Heal, Smite, Holy Fire, Mind Blast, Mind Flay, Mana Burn, Starshards, Resurrection, Mind Control; Lightwell, Renew, Prayer of Mending and Power Word: Shield are not in it; identical in 70888 | aligned (Binding Heal and Starshards are not registered for a healer; Lightwell's 1.5 s prepull cast is rightly unprotected) |
| P4 | Nature's Focus (17063, the druid 5/5) | 14% a rank on Healing Touch, Regrowth, Tranquility, Wrath, Starfire | effect 0 aura 108 misc 9 at 14, mask 0x10400065 / 0x02000000 / 0xA200 / 0x10000 = Healing Touch, Regrowth, Tranquility, Wrath, Starfire, Hurricane, Entangling Roots, Hibernate, Rebirth, Revive, Nourish, Starfall; Wild Growth is not in it; identical in 70888 | aligned |
| P5 | Healing Focus (16181, the shaman 3/3) | 23/47/70% on Healing Wave, Lesser Healing Wave, Chain Heal | effect 0 aura 108 misc 9 at 70, mask 0x1C0 / 0 / 0x100000 / 0 = Healing Wave, Lesser Healing Wave, Chain Heal (and Healing Rain); identical in 70888 | aligned |
| P6 | Martyrdom (14531) | not modelled: the pulses never crit the healer | aura 42 triggering 27828 on a melee or ranged crit taken; in neither priest build | aligned |
| P7 | Auras and buffs the requests ask for | priest: Inner Fire on, Power Infusion to self; druid: Innervate to self; shaman: nothing; the raid preset's 14 buffs | every one is applied (`request.go` 686 to 862, the preset's buffs through `Layer`); no healer asks for a paladin aura, and `proto.RaidBuffs` has no Concentration Aura field | aligned; the aura is a preset question (O1) |
| P8 | The talent's effect in the sim | emptying the talent (the points left unspent) at raid 60: Holy +0.06 / +0.13% (err 0.06), Discipline -1.72 / -1.72% (0.06), Restoration druid -0.37 / -0.47% (0.06), Restoration shaman -2.56 / -2.57% (0.07); bare 60: -0.04 / -0.08% (err 0.08), -2.24 / -2.22% (0.07), -1.14 / -1.27% (0.08), -0.80 / -0.74% (0.06) | the engine applies what it reads | aligned |
| P9 | No pushback at all (the ceiling a party Concentration Aura could reach, 35 on top of 70) | raid 60: Holy -0.08 / -0.02% (err 0.06), Discipline +0.77 / +0.82% (0.06), druid +0.28 / +0.27% (0.04), shaman +1.22 / +1.17% (0.07); bare 60: +0.04 / +0.08% (err 0.08), +1.03 / +1.03% (0.07), +0.66 / +0.64% (0.07), +0.17 / +0.23% (0.06) | scratch engine 90c598cbc | measured for O1; nothing published moves |

**The live server's pushback, from the logs.** For every cast-time spell a logged priest, druid or shaman chained (the next cast landing within the client cast time plus 2.5 s of the previous one), the cast landed late by the gap minus the client cast time, and a creature's hit on the caster inside the last cast time plus 1.2 s before it landed marks the cast as hit. Latency and the global cooldown put most unhit casts inside 0.35 s. A one-second pushback shows as a bump in the 0.7 to 1.4 s bucket on the hit casts only; a half-second rule would show in the 0.35 to 0.7 s bucket.

| class, spell | chains | on time (under 0.35 s) | 0.35 to 0.7 s late | 0.7 to 1.4 s late | later |
|---|---|---|---|---|---|
| Druid Regrowth, a hit inside the cast | 60 | 40 (67%) | 1 (2%) | **14 (23%)** | 5 (8%) |
| Druid Regrowth, no hit | 770 | 676 (88%) | 25 (3%) | 36 (5%) | 33 (4%) |
| Druid Healing Touch, a hit inside the cast | 55 | 42 (76%) | 3 (5%) | **8 (15%)** | 2 (4%) |
| Druid Healing Touch, no hit | 339 | 297 (88%) | 10 (3%) | 15 (4%) | 17 (5%) |
| Shaman Lesser Healing Wave, a hit inside the cast | 42 | 26 (62%) | 3 (7%) | **5 (12%)** | 8 (19%) |
| Shaman Lesser Healing Wave, no hit | 468 | 427 (91%) | 12 (3%) | 13 (3%) | 16 (3%) |
| Priest Lesser Heal, a hit inside the cast | 22 | 10 (45%) | 2 (9%) | **9 (41%)** | 1 (5%) |
| Priest Lesser Heal, no hit | 400 | 374 (94%) | 8 (2%) | 14 (4%) | 4 (1%) |
| Shaman Healing Wave, a hit inside the cast | 12 | 7 (58%) | 1 (8%) | **4 (33%)** | 0 |
| Shaman Healing Wave, no hit | 118 | 105 (89%) | 5 (4%) | 3 (3%) | 5 (4%) |
| Priest Smite, Mind Blast, Druid Wrath, Starfire (hit) | 43 | 30 (70%) | 2 (5%) | **8 (19%)** | 3 (7%) |
| the same, no hit | 652 | 615 (94%) | 11 (2%) | 9 (1%) | 17 (3%) |
| **every cast-time spell, a hit inside the cast** | **253** | 168 (66%) | 13 (5%) | **51 (20%)** | 21 (8%) |
| **every cast-time spell, no hit** | **3,381** | 3,098 (92%) | 79 (2%) | 105 (3%) | 99 (3%) |

The excess sits in the one-second bucket (20% against 3%) and nowhere near the half-second one (5% against 2%), so the live server's first-hit pushback is the one second the engine uses. The share pushed, about 17 points of the hit casts, is what a 23 to 70% resist leaves on players whose points in the talent at 14 to 40 are unknown (the window also marks hits that landed just before the cast started, which dilutes it); the magnitude, not the share, is what this test can read. 19 players, 26 reports, 5 to 12 Nov.

## Regen and procs: engine against the client and the logs

| # | mechanic | engine | client / log | verdict |
|---|---|---|---|---|
| R1 | Spirit regeneration | priest 12.5 + Spirit/4, druid 15 + Spirit/5, shaman 17 + Spirit/5 a 2 s tick (vmangos `OCTRegenMPPerSpirit`); the five-second rule from the cast's completion (`SpendCost` at `OnComplete`, `FiveSecondRuleRefreshTime` now + 5 s); mp5 always | no regen gametable in 70790 (basemp, hppersta, combatratings, playerexpectedstat only); no public log carries a periodic regen tick | engine's own, open since round 7 (O4) |
| R2 | Meditation (14521), Reflection (17106), Mindfulness (1223033) | 17/33/50% of Spirit regen while casting (`SpiritRegenRateCasting`) | the talents' own text at every rank | aligned |
| R3 | Innervate (29166, the druid on itself) | Spirit regen times five and the whole of it while casting for 20 s, 6 min, 5% of base mana, replaces Read Leyline | aura 134 at 100 and aura 110 at 400, duration 20000, cooldown 360000 | aligned (no logged druid is 40; L4) |
| R4 | Inner Focus (14751, both priests) | the next priest spell free and +25% crit if it can crit, 3 min; autocast at the pull and every 3 min | aura 108 misc 14 at -100 on a mask that holds every priest spell with a cost (Lightwell, Prayer of Mending, Renew and Power Word: Shield included), aura 107 misc 7 at 25 on the mask without Renew; the server casts the talent spell itself | aligned |
| R5 | Litany of Light (1317006, Holy 2/2) | 10% of the base cost back on a heal whose previous heal was a different spell (5% a rank), from the talent's own id; 86 returns a fight at raid 60 | rank text 5% and 10%; no logged Holy priest has the talent (the oldest is 22) | aligned on the row; open until a Holy priest at 30 or more logs (L5) |
| R6 | Omen of Clarity (16864) and Clearcasting (16870) | 4% a cast on every druid heal, 10 s between procs, one free cast | SpellAuraOptions 129438: ProcChance 100, ProcCategoryRecovery 10000, ProcTypeMask 0x14014; the 11-12 healers lane's 4.1% on 5,082 casts; Clearcasting applications on five logged druids here | aligned |
| R7 | Water Shield (408510 / 408511, the shaman) | 3 globes (4 with Totem of Sustaining), 2% of maximum mana a globe, 3.5 s between globes, off hits taken and heal crits | effect 0 aura 42 at 100 triggering 408511 (effect 137 at 2); the 11-12 lane's 21 logged gains | aligned |
| R8 | Mana Tide Totem (16190 to 1325256) | 13 s totem, four ticks every 3 s of 88/133/197/290 to the party, 5 min; the breakdown's "5 events" on the row are the 60-mana cost plus the four ticks (one ActionID) | 1325256 duration 13000, category 300000; 16191 effect 35 aura 24 at 88 every 3000 ms | aligned (no logged shaman has the talent; L6) |
| R9 | Blessing of Wisdom, Mana Spring | 12/18/30 and 4/8 every 5 s and 2 s, the raid preset's and the shaman's own, two Mana Springs never stacking | logged on nine casters: Wisdom 19742 at 12, 19850 at 18, 19853 at 30; Mana Spring 5677 at 4, 10493 at 8 | aligned |
| R10 | Dark Sacrifice (Undead priests) | five ticks of the rank's grown amount plus Spirit, 10 min | the 11 Nov priests lane (D1) | aligned |
| R11 | Shadowfiend (401977) | registered at level 1 from the trainables, never cast: no card lists it and it is no major cooldown | SkillLineAbility 49886: skill line 78, class mask 16, AcquireMethod 0, rank 1, both clients; the priest guide index says the rune is not in Forever; no logged priest (19 on 11 Nov, 7 here, to level 30) ever casts it | no number rests on it; a data note (O2) |
| R12 | Any casting-regen aura from a row the server never casts (the Illumination pattern) | every `SpiritRegenRateCasting` and `ForceFullSpiritRegen` writer in the three classes: Meditation, Spirit Tap (Shadow only), Reflection, Mindfulness, Improved Stormstrike (Enhancement only), Innervate | each is a talent's own aura or 29166 | none found |

## Before and after

No engine change, so no before and after. For the record and for O1, the published entries on the pin (10,000 iterations, effective HPS, the combined error), then what the pushback weighs:

| cell | Holy priest | Discipline priest | Restoration druid | Restoration shaman |
|---|---|---|---|---|
| raid 60 Alliance s1 / s50001 | 870.08 / 869.77 (+/- 0.28) | 603.89 / 603.71 (0.19) | 1028.13 / 1028.35 (0.24) | 889.85 / 889.80 (0.31) |
| raid 60 Horde s1 / s50001 | 881.68 / 882.00 (0.30) | 596.88 / 596.33 (0.22) | 1024.22 / 1023.93 (0.24) | 895.55 / 894.94 (0.31) |
| bare 60 Alliance s1 / s50001 | 417.51 / 417.54 (0.16) | 287.42 / 287.26 (0.10) | 495.83 / 496.01 (0.19) | 373.88 / 373.62 (0.11) |
| bare 60 Horde s1 / s50001 | 424.34 / 424.70 (0.17) | 287.30 / 287.53 (0.11) | 494.24 / 494.31 (0.19) | 375.93 / 375.91 (0.11) |
| bare 20 Alliance / Horde | 41.97 / 40.61 | 46.39 / 45.43 | 75.87 / 75.69 | 73.09 / 72.82 |
| bare 30 Alliance / Horde | 76.68 / 79.89 | 94.71 / 96.51 | 132.15 / 130.81 | 134.73 / 135.18 |
| bare 40 Alliance / Horde | 175.92 / 184.46 | 144.56 / 149.01 | 243.82 / 243.35 | 210.41 / 210.75 |
| bare 50 Alliance / Horde | 258.80 / 272.68 | 206.99 / 210.58 | 377.48 / 375.07 | 281.29 / 282.46 |

The mana lasts 124 s (Holy), 91 s (Discipline), 274 s (druid) and 230 s (shaman) of the 300 s raid fight on the Alliance sets.

**What the pushback weighs** (Alliance, seeds 1 and 50001, relative to the rows above):

| variant | Holy priest | Discipline priest | Restoration druid | Restoration shaman |
|---|---|---|---|---|
| the pushback talent emptied, raid 60 | +0.06 / +0.13% (err 0.06) | -1.72 / -1.72% (0.06) | -0.37 / -0.47% (0.06) | -2.56 / -2.57% (0.07) |
| the pushback talent emptied, bare 60 | -0.04 / -0.08% (err 0.08) | -2.24 / -2.22% (0.07) | -1.14 / -1.27% (0.08) | -0.80 / -0.74% (0.06) |
| no pushback at all (scratch 90c598cbc), raid 60 | -0.08 / -0.02% (err 0.06) | +0.77 / +0.82% (0.06) | +0.28 / +0.27% (0.04) | +1.22 / +1.17% (0.07) |
| no pushback at all, bare 60 | +0.04 / +0.08% (err 0.08) | +1.03 / +1.03% (0.07) | +0.66 / +0.64% (0.07) | +0.17 / +0.23% (0.06) |

Where it comes from. The pulses land on the healer three times in five every 5 s, so about one hit in two casts; with the talent, 30% of those hits cost a second (70% with no talent), and a never-pushed-back healer gets the whole second back. The shaman's raid fight is Chain Heal and Healing Wave casts end to end with mana to spare (230 s), so it gains the most from the ceiling and loses the most without the talent; Discipline's Penance channel and Flash Heals are cheaper per second than the mana they save, so its talent is worth 1.7 to 2.2%; the druid's bar is mostly instant Rejuvenations and Swiftmends, which no hit touches; and the Holy priest is bound by its bar (dry at 124 s raid, 47 s bare), so casting time has no value to it at all: without the talent it even gains a hair at raid 60 (+0.06 to +0.13%, inside twice the error) because a pushed-back cast spends its mana later. None of these is a bug: they are what the engine does with the client's rows, and the ceiling row is the measure for O1.

## Logs against the engine

| # | check | evidence | verdict |
|---|---|---|---|
| L1 | **First-hit pushback is about one second** | the table above: 51 of 253 hit casts 0.7 to 1.4 s late against 105 of 3,381 | aligned |
| L2 | **Blessing of Wisdom and Mana Spring tick at the engine's amounts** | 2,700 Wisdom ticks of 12, 18 and 30 (ranks 2, 3 and 4) and 1,100 Mana Spring ticks of 4 and 8 across nine casters | aligned |
| L3 | **Clearcasting lands on druid heals** | Surgeon 53, Super 24, Premo 11 applications on themselves while healing | aligned (the rate: the 11-12 lane) |
| L4 | Innervate, Litany of Light, Mana Tide, Inner Focus's free cast, the Spirit tick | no logged healer is high enough (Innervate 40, Mana Tide 40, Litany about 30) and the stream carries no periodic regen tick or mana cost | no data |
| L5 | Water Shield | the 11-12 lane (Eltato, 21 gains of 2%); Tallak at 40 (293 heals) runs without the talent, so adds nothing | aligned |
| L6 | No priest casts Shadowfiend | 7 priests here (to 30), 19 on 11 Nov | consistent with the guide, against the SkillLineAbility row (O2) |

## Builds, rotations and rank rows

Nothing moved: no engine change, no row re-judged. The four cards' casts under the pulses are the pin's.

## Gates

- No engine change; the pin stays 4c595ee6b. The scratch branch `scratch/healers2-no-pushback` (90c598cbc) is a measuring branch and is never pinned.
- Site: review files, the ledger row and the lane's job files only; no Go, TypeScript, data or guide file changed, so no unit, type, lint or Playwright gate applies beyond the markdown. Rebased on origin/main and pushed to main.

## Files

- `design/reviews/2026-11-12-sim-accuracy/healers-pushback-regen/` (the box runner `jobs.sh`, `jobs-one.sh`, the job files `jobs-60.txt`, `jobs-bands.txt`, `jobs-notalent.txt`, `jobs-nopush.txt`); this report; the ledger row.
- Fork: `scratch/healers2-no-pushback` at 90c598cbc, pushed, never pinned.

## Still open

- **O1, Concentration Aura in the raid preset.** An Alliance (and, in Forever, a Horde) raid has paladins, and a Holy paladin's own aura keeps its own casts; whether the healer party in the raid preset stands under one is a preset question, not an engine bug: the request has no field for it and the preset lists none. Its ceiling is the "no pushback at all" row above (35 on top of 70 is never a pushback for every heal the four cards cast). If the owner wants it, `proto.RaidBuffs` gains a field, `core.ConcentrationAuraAura` runs for every player under it, the preset lists it with its reason, and the four healers' raid 60 picks re-rank (a new MR row then).
- **O2, Shadowfiend.** SkillLineAbility 49886 says a trainer teaches 401977 to priests at level 1 in both clients, the priest guide says the rune is not in Forever, and no logged priest casts it. The engine registers it and no card uses it, so no number depends on which is right; the trainables file lists it as learnable at 1, which the guides' rotation lists read. A priest trainer screenshot or a level 60 priest log settles it.
- **O3, the regen rows that need a higher healer log:** Litany of Light's amount and its "different spell" rule, Innervate's five times Spirit, Mana Tide's four ticks, Inner Focus spending on the next spell cast. The engine reads the client rows; the first Holy priest at 30, Restoration druid at 40 and Restoration shaman at 40 with the talents close them.
- **O4, the Spirit tick itself** (12.5 + Spirit/4, 15 + Spirit/5, 17 + Spirit/5): no gametable in the client, no periodic regen event in the public stream; our own report pipeline with advanced logging is the way to read it.
