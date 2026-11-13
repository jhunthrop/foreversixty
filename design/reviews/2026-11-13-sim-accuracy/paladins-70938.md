# Sim accuracy 2026-11-13: Retribution and Protection paladin on 1.60.1.70938

Lane `sim-paladins70938`, brief `design/reviews/sim-accuracy/brief.md`. Site from origin/main 96c7b4fd8 (active build 1.60.1.70938), fork from master 45c009ce3 (the fractional-values pin). Scope: Retribution and Protection at 50 and 60, bare and raid 60: the engine against the 70938 tables and the server's hotfix push for every seal, judgement, aura, blessing, strike, talent and set bonus the raid rotations and builds use; the twist and seal timing against a level 63 boss; the Protection threat model against the warrior and the bear; logs at the highest levels there are.

**Totals: 2 confirmed bugs fixed, 0 paladin numbers moved. (F1, new data) Gloves of Holy Might: the server's push 114008 of 13 November widens its trigger to melee specials and ranged attacks, which the 70938 client and the engine left out (shared item: every entry that wears it +0.0 to +0.9%; Retribution bare 40 +0.3 to +0.4%, bare 50 +0.5%; MR106). (F2, old) Seal of Command's proc carried the Era client's per-rank ids 20944 to 20947, where every Forever rank triggers one spell, 20424, the id the logs name on every proc (no number moves). Every Retribution and Protection cell (raid 60 and bare 20 to 60, both factions) is identical to the hundredth before and after. The 70938 client and its push carry no paladin change: 601 paladin spells identical to 70790 but for three fractional amounts the engine already reads, and all 273 pushed paladin rows already in the shipped tables. Logs: the three private reports of 10 to 12 November (new to the paladin lanes) confirm Seal of Command at 7 a minute from white swings only, Seal of Righteousness on every swing, and Unsealed Wrath's two open readings (Holy Strike stores nothing; one 15 second timer). No paladin at 50 or above has a fight in any log. Engine pin 45c009ce3 to 5603877e4.**

For the coordinator: the data pipeline layered **no** push row at 70938 for any class (`hotfix-push.json`: Spell 0 changed, 1,253 withheld; SpellEffect 0 changed, 999 withheld). wago.tools lists every 70938 enUS push entry with `data: None`; the deDE, frFR, koKR and zhTW copies carry the values. `PushCollector` keeps enUS only. For paladins nothing is lost (the shipped tables carry every pushed row), but push 114008 (F1) reached the engine only by hand here, and any later push will be missed the same way. It is issue #27's root cause on the spell tables; commented there.

## Method

1. **Client.** `data/builds/1.60.1.70938/spellconst/paladin.json` (and the fork's vendored copy) against 70790's, field by field: 601 spells, no spell added or removed, three amounts differ only by their fraction (Seal of Fury rank 3 486.9, Judgement of Fury ranks 1 and 2 194.4, 36.9 and 316.8), which the engine reads since 45c009ce3. `talents/paladin.json`, `trainables/paladin.json` and the gametables are identical but for the build stamp. The fork's conformance report at 70938: 0 paladin mismatches.
2. **The push.** wago.tools' hotfix page, every 70938 entry whose record belongs to a paladin spell (record id to SpellID through the 70938 DB2 exports): 273 records in 14 tables, each compared with the shipped 70938 row from a locale that carries values. All 273 equal. The pushes after 70790's newest (113732) touch paladins only in text (Unsealed Wrath, Swift Judgement, Shieldsworn, Improved Hammer of Justice, Turn Evil: the same numbers as the talent texts) and one SpellEquippedItems row giving Unsealed Wrath a weapon requirement (class 2, any subclass), already in the shipped table. Then every class: the 295 records pushed after 113732. Two rows differ from the shipped tables: Gloves of Holy Might's trigger (F1) and an NPC spell (437589, an escort script); 82 CurvePoint rows are new identity curves (0 to 1) that no shipped trait reads; 47 entries carry no values in any locale and cannot be checked.
3. **Engine.** Every file the two cards and raid builds touch, at 45c009ce3, against the rows above and the client's aura options: Seal of Command (ProcChance 100 with the server's 7 a minute, ProcTypeMask 4 = white swings, 1 s recovery), Seal of Righteousness and Seal of Fury (mask 4, every swing), Unsealed Wrath (mask 4; Echo of Fury, Command, Righteousness, Justice are 1311701 to 1311705, 3 stacks), Seal of the Crusader (+306 and 2.4 a level attack power, 40% attack speed; Judgement of the Crusader 161 Holy, 40 s, refreshed by the paladin's melee hits), Judgement and its ranged table, Holy Strike (93 + 3.2 a level, 0.429, 50% weapon, 10 s), Consecration, Exorcism, Hammer of Wrath (498, 0.429, 6 s), Hammer of the Righteous (one-handers only since raid-searches B1), Vengeance (1% a rank a stack, 5 stacks, 30 s), Righteous Fury (+33% all threat and +90% Holy, multiplied), Improved Righteous Fury (3% a rank), Iron Creed, Holy Shield (4 charges, 75/110/153/221 at 0.08, x1.2 threat), Redoubt, Reckoning (1.5 s from blocks), Shield Specialization (3 s), Shieldsworn, Templar's Bulwark, Windfury Totem (ProcTypeMask 0x14, 100 ms; Seal of Righteousness and Seal of Fury procs carry SuppressEquipProcs, Seal of Command's proc does not, as the client's melee-class proc row). Aligned except F1 and F2.
4. **Logs.** Our seven reports, read from the bucket (`r2`, read-only): z2opt2x2gs5u, gxrgvz4yzsi2, 3sdrcwdgo2yi (Desolace, 10 Nov), edcgw6rodqs2 (Scarlet Monastery, 12 Nov), eygg26pg4gr2 (Hillsbrad, 12 Nov), ctxicu4cuyek, 4q5el2xsgoy7; foreverlogs.gg's list still ends at 2854 (2855 onward 404); no Warcraft Logs Forever partition. The highest paladin with a fight is 45 (one fight, three casts); the samples are 28 to 35. The advanced log drops events for some segments (a swing with no `SWING_DAMAGE_LANDED` often has no seal damage either), so per-swing rates use only swings carrying `SWING_DAMAGE_LANDED`.
5. **Sims** on the Windows box through `tools/remote-test.sh`, 10,000 iterations, seed 1, one sim at a time, nice 15, GOMAXPROCS 5 (`paladins-70938/`: `full.sh` and `base.txt` for whole breakdowns, `jobs.sh` with `paladins.txt` and `holy-might.txt` for before and after). Fork tests in a clone on the box (`-tags with_db`).

## Confirmed and fixed

| # | finding | evidence | origin | fix |
|---|---|---|---|---|
| F1 | **Gloves of Holy Might ignored melee specials and ranged attacks.** The engine took its trigger from client 70938's SpellAuraOptions 249242 (1330855): ProcTypeMask 69636 (0x11004: white swing or harmful spell). The server's push 114008 (13 November, the deDE, frFR, koKR and zhTW entries; enUS has no values) makes it 69972 (0x11154), adding melee specials (0x10), ranged auto attacks (0x40) and ranged specials (0x100); chance 15 and 20 s unchanged. The push wins over the client. Now the trigger reads every melee and ranged attack and harmful spell. Red first: `TestGlovesOfHolyMightProcsOnWhiteSwings...` with a Wing Clip now expected to proc (0 Smites from 362 hits at the red commit). | wago.tools hotfix entries for record 249242, push 114008 | new data (13 Nov push) | E 2f2ae88ae (red), d3b4ed013 |
| F2 | **Seal of Command's proc carried the Era client's ids.** Ranks 2 to 5 cast procs 20944 to 20947, which the Forever client does not carry; every rank's effect 0 triggers 20424 (70% weapon damage, 0.29), and the logs name 20424 on every proc under ranks 1, 2 and 3 (458 procs). One proc spell now serves every rank (its damage never depended on the rank). Nothing moves: every Retribution cell is identical to the hundredth. The site's curated names and icons for 20944 to 20947 stay for saved results. Red first: `TestSealOfCommandProcIsTheClientsOneSpellAtEveryRank`. | spellconst 20375 to 20920 effect 0; our logs | old (the Era port) | E a84383ada (red), 5d0d0aadd |

Before and after (old pin 45c009ce3 against 5603877e4; `paladins-70938/holy-might.txt`, every published 70938 entry whose hands pick is item 867):

| spec | cells | change |
|---|---|---|
| Retribution | bare 40 Alliance / Horde, bare 50 Alliance | +0.29 / +0.39%, +0.48% |
| Arms | bare 40, 50, 60, raid 60 | +0.70 to +0.89%, +0.58 to +0.66%, +0.43 to +0.45%, +0.02 to +0.22% |
| Fury | bare 40, 50, 60 | +0.23 to +0.36%, +0.22 to +0.33%, +0.03 to +0.20% |
| Beast Mastery | raid 60 | +0.65 to +0.66% |
| Survival | bare 40 to 60, raid 60 | +0.22 to +0.32%, +0.25% |
| Assassination, Combat, Subtlety | bare 40 to 60, Assassination raid 60 | -0.04 to +0.32% |
| Enhancement | bare 40 | +0.33 to +0.49% |
| Feral | bare 40 | +0.21 to +0.25% |
| Retribution and Protection, every other cell | raid 60, bare 20 to 60, both factions (`paladins.txt`) | 0.00% |

The gloves now also fire from Holy Strike, Hammer of Wrath, Judgement of Command and Seal of Command's proc. A re-rank may move the hands slot wherever 867 was close (MR106).

## Logs against the engine

| # | check | evidence | verdict |
|---|---|---|---|
| L1 | **Seal of Command procs at 7 a minute from white swings.** Over 18 paladins at 16 to 33 with at least 14 landed swings under the seal, 413 procs (casts of 20424) against 403.2 expected from 7 x weapon speed / 60 (speed from each one's median swing gap). On swings with full advanced logging, 287 procs in 804 swings (0.36 a swing on 2.7 to 3.8 s weapons). Every proc sits on a white swing; none follows a Holy Strike. | 3sdrcwdgo2yi, edcgw6rodqs2, eygg26pg4gr2, gxrgvz4yzsi2, ctxicu4cuyek | aligned |
| L2 | **Seal of Righteousness hits on every landed swing.** 93 of 97 fully logged swings under it carry its damage within 0.25 s. | 3sdrcwdgo2yi (three paladins at 30 to 33) | aligned |
| L3 | **Unsealed Wrath stores only white swings.** 62 of 62 Echo stacks were added at a white swing; none at any of the same paladins' 33 Holy Strikes. This settles the engine comment's open reading. | Echo of Command, Righteousness, Fury (1311703, 1311704, 1311701): eygg26pg4gr2, 3sdrcwdgo2yi, gxrgvz4yzsi2 | aligned |
| L4 | **The copies share one 15 s timer that each stack restarts.** Three Echoes expired 15.04, 15.06 and 15.08 s after their last stack; a swing at three stacks refreshes it. | as L3 | aligned |
| L5 | **A swing under a different Seal releases the copies and starts a new stack.** Every release (11) removes the Echo and applies the new Seal's at the same swing. Where the log carries the seal damage (one fight), two stored Righteousness copies gave two Righteousness hits under Seal of Command, and three stored Commands gave one Command proc under Righteousness (expected 1.2 at 7 a minute on a 3.6 s weapon). | eygg26pg4gr2 fight 46 | aligned (n small) |
| L6 | **Judgement cadence.** Retribution gaps bottom out at 7.9 to 8.0 s (Improved Judgement 2), Protection at 9.9 to 10.0 s; Swift Judgement on its own minute. | all reports | aligned |
| L7 | **Holy Strike cadence.** Gaps bottom out at 9.9 to 10.0 s on every paladin but one, a level 35 Protection paladin at 6.05 s: Shieldsworn's 4 s off, as the talent says. | all reports | aligned |
| L8 | **Holy Shield's block damage is logged under the rank's own id** (20925 at 75 to 84 on 116 blocks, 20927 at 110); Shieldsworn's strike under 456544 (110 to 111). The engine names the block damage 20955 to 20957 and 27179 (ids the client lacks); see H1. | edcgw6rodqs2, 3sdrcwdgo2yi | naming only, held |
| L9 | **Shares at 28 to 35, for reference only** (no 50+ log): logged Retribution with Seal of Command, white 41 to 85% (median about 64%), Seal of Command 4 to 12%, Judgement of Command 2 to 21%, Holy Strike 3 to 14%, Consecration 1 to 14%; Protection white 30 to 76%, Holy Shield 3 to 15%, Seal of Fury uptime about half of active time. The sim's leveling card at 31 (Seal of Righteousness, no Seal of Command talent) reads white 35.8%, Righteousness 20.4%, Judgement 20.1%, Holy Strike 17.3%. Players at these levels mostly use Seal of Command and trash pulls; nothing here is a like-for-like test. | as above | not comparable |

## The twist and seal timing against a level 63 boss (raid 60, Alliance, sim)

833.35 DPS. Seal of Command up 74.7% and Seal of Righteousness 25.3% (12.0 and 11.3 casts in 180 s, a flip every 16 s), Unsealed Wrath active 94.8% (22.9 activations), 0.0 s out of mana. White swings 28.2% of damage plus 11.0% from extra attacks (Windfury Totem 16.4 procs); Seal of Command's proc 21.4%, Holy Strike 12.1%, Consecration 8.5%, Judgement of Command 6.5% (16.1) and Righteousness 2.7% (6.2) of 22.3 Judgements (one every 8 s), Hammer of Wrath 5.5%, Righteousness' swing damage 4.2%. Against the +3 boss: 0.0% miss (raid hit), 6.5% dodge on every melee attack, 39.8% glancing on white swings, none on specials; Judgement on the ranged table. A glancing swing is a landed swing, so it stores and releases copies as the logs show (L3, L5); a dodged swing does neither, and the card's flip-back waits for a landed one. The store-three-then-flip card and its 80% mana pace stand.

## Protection threat against the warrior and the bear

Raid 60 Alliance, Thunder Clap on every tank's boss (MR99): paladin 1,180.6 threat a second, taking 298.1 (warrior 1,667 and bear 1,939 from the tank follow-ups). The paladin's threat: Holy Shield's damage 261, Holy Strike 207, Hammer of the Righteous 177, white swings 158, Judgement of Fury 118, Consecration 101, Seal of Fury 84, extra attacks 52, mana gains 22. Every multiplier the client or Blizzard states is applied and aligned (Righteous Fury x1.33 with x1.9 on Holy, Holy Shield x1.2, Iron Creed +30% Holy Strike); no factor is invented. The order rests, as for the other two tanks, on the absence of any Forever in-game threat reading (O4 of tanks-60); nothing in 70938 or its push changes a paladin threat number.

## Held and open

- **H1, Holy Shield's damage id.** The logs name the block damage under the rank's own spell (20925, 20927); the engine uses 20955 to 20957 and 27179. Giving the damage spell the rank's id needs a tag to keep it apart from the cast (the APL resolves the first spell with an id), and the result page would then read it as a variant ("Holy Shield (2)") where today the curated engine names read "Holy Shield". Nothing moves either way; held for a page-side variant rule.
- **O1, the push layer.** No 70938 push row was layered for any class (see the top). For the coordinator; commented on issue #27.
- **O2, no paladin above 45 in any log**, so Hammer of Wrath, the twist at 60, Light's Vigil, Shieldsworn's block and Sanctified Judgement's 60% (L14 of 11 Nov) still have no high-level log.
- **O4 of tanks-60** (threat factors without a Forever number) stands for the paladin as well.

## Gates

- Fork, on the box: the red commits fail as stated (`TestSealOfCommandProcIsTheClientsOneSpellAtEveryRank` lists 20944 to 20947 registered; the gloves test 0 Smites from a Wing Clip); at d3b4ed013 the paladin, common, core, hunter, warrior, rogue, shaman and druid suites pass; at 5603877e4 `sim/conformance`, every paladin package and `item_effects` pass.
- Site, on the box with the pin: `go vet ./...` and `go test ./request/ ./rankchoices/ ./leveling/ ./adapter/ ./cmd/spec-breakdown/`.
- No guide or card note names the gloves' trigger or the Seal of Command proc id, so no prose moves and no wow-player round is needed.

## Files

- Fork (master fast-forwarded to 5603877e4): a84383ada, 5d0d0aadd (F2), 2f2ae88ae, d3b4ed013 (F1), 5603877e4 (Unsealed Wrath's comment records L3 and L4).
- Site: the pin (`sim/enginever/version.go`, `data/proto/ENGINE_SHA`, `web/src/lib/sim/version.ts`, the three `web/src/fixtures/sim/*.json`); `paladins-70938/` (runners, job lists); this report; the ledger row and MR106.
- No regens: no BiS rank, data run or simdb rebuild; `Data.lua`, `addon-data.json` and `bis/*.json` untouched.
