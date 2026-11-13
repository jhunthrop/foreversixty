# Sim accuracy 2026-11-13: Discipline priest (the owner's "it doesn't look like we model disc correctly")

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `sim-disc`. Site from origin/main d730bbcac, fork from master 5515ae0ef (the pin). Client on the site: 1.60.1.70790 with its hotfix push; live client 1.60.1.70888. Scope: Discipline end to end against the tier list's raid 60 healer row (Restoration druid 1,024.0, Restoration shaman 895.8, Holy priest 882.0, Holy paladin 780.8, Discipline 596.6, 12 Nov), with Holy beside it because the two share the bug.

**Totals: 1 confirmed bug, old, fixed (a card bug shared by both priest cards). 0 engine bugs, 0 new data. The rest of the gap is real under the heal profile, and the report says where it comes from and what Discipline brings that HPS does not show.** Engine pin 5515ae0ef to 186440a08 (the fork carries the card copy, `ui/healing_priest/apls`, nothing else).

## What the HPS metric counts for a Discipline priest

- **Shields count for what they soak, never for what they would have soaked.** `core/shield.go` books the absorbed part of every shield (Power Word: Shield, Divine Aegis) as the shielding spell's effective healing at the moment damage hits it; the published HPS is effective healing. The amount applied goes into raw HPS only, so a shield that expires or is overwritten unused counts nothing. Overheal from heals never counts.
- **Divine Aegis is in.** 15% of every non-periodic crit (Penance's three pulses included: they are direct heals, not ticks) as a shield capped at 20 a level, adding to a live one and renewing it. At raid 60 it is 3.2% of Discipline's healing; 89% of what it applies is soaked.
- **Weakened Soul gates only the tank's shield.** The card shields the tank (under 90%, Weakened Soul off, on the 90% mana pace); Renewed Hope's Flash Heal, Greater Heal and Penance cut 4 s off it (once per Penance cast, Blizzard's bug ruling of 10 Oct). The tank is hit every 3.3 s for about 2,600, so a 1,220 shield is gone on the next swing: 6.6 shields a fight, 7,988 applied and 7,888 soaked (99%). The profile does not leave shields to expire; if anything it soaks them more completely than a real raid (L2).
- **The fight is a mana fight.** The profile asks one healer to cover a stated share of a raid, 300 s, and every healer's raid set runs dry inside it: Discipline at 91 s, Holy at 124 s, the shaman at 230 s, the druid at 274 s. Raw and effective HPS differ by 3% for Discipline: almost nothing overheals, so a shield's real advantage (it never overheals and lands before the hit) earns it nothing over a heal of the same size.

## Engine against the client (1.60.1.70790 with its push; 70888 identical)

Every Discipline row the raid build and card touch, read against `spellconst`, `talents` and `trainables` for 70790, the push (`hotfix-push.json`: no priest row) and the earlier lanes' 70888 diff (priests 2026-11-11, shadow 2026-11-12, healers pushback 2026-11-12: no priest row moved; Blizzard's posts through 12 Nov have no Discipline line). The engine read: `sim/priest/{power_word_shield,penance,heals,healing_talents,talents,priest}.go`, `core/{shield,raid_damage,metrics_aggregator}.go`.

| # | row | engine | client | verdict |
|---|---|---|---|---|
| C1 | Power Word: Shield ranks 1 to 10 (17 to 10901) | client amount plus points per level to MaxLevel, plus 0.1 of healing; 30 s; 4 s category cooldown; Weakened Soul 15 s | rank 10: 928 base, 4.3 a level, 0.1, 500 mana, category 4000, duration 30000; 6788 duration 15000 | aligned |
| C2 | Improved Power Word: Shield (14748) | 7/14/20% on the absorb | "by 20%" at rank 3 | aligned |
| C3 | Soul Warding (402000) | -4 s and -15% cost | text | aligned |
| C4 | Twin Disciplines (1225132) | +5% on Renew, Penance, Holy Nova, Prayer of Mending, Desperate Prayer and the racial heals; not on the shield | effect masks 0x19480020/0x18012/0/0x200000 and 0xa08060/0/0/0x10 exclude the shield's family bit 0 | aligned on the mask (L3 asks nothing of it) |
| C5 | Penance ranks 1 to 4 (402174 to 1316995) | instant pulse and two more a second apart, 0.19 each, 12 s, 385 mana at rank 4 | cast 1316995 cost 385, category 12000; helper 1316992 2000 ms, period 1000, heal 1316991 425 at 0.19 | aligned |
| C6 | Renewed Hope (425280) | +2% crit a rank on Flash Heal, Lesser Heal, Heal, Greater Heal and Penance on a Weakened Soul target, -1 s a rank off it | 10 and -5000 at rank 5 | aligned |
| C7 | Divine Aegis (431622, 431624) | 5% a rank of a non-periodic crit, capped at 20 a level, 12 s | 15 and 20 at rank 3, 431624 12000 ms | aligned |
| C8 | Meditation (14521), Mental Strength (18551), Mental Agility (14520), Improved Healing (14912) | 17/33/50% casting regen; 2% Intellect a rank; 3/7/10% on the client mask (not Penance); 5% a rank on Lesser Heal, Heal, Greater Heal, Penance, Prayer of Mending | texts and masks | aligned |
| C9 | Inner Focus (14751) | next spell free, +25% crit on a non-periodic spell, 3 min | -100 and 25, 180000 | aligned (healers pushback lane R4) |
| C10 | Power Infusion (10060) | +20% spell damage and healing, 15 s, 3 min, 20% base mana | aura 79 and 136 at 20 | aligned; not in the published build (the guide's own call) |
| C11 | Divine Spirit, Spirit to spell power | Divine Spirit is a raid buff in the preset; Discipline has no Spirit-to-healing talent (Spiritual Guidance is Holy's tier 5) | talents/priest.json | aligned |
| C12 | Litany of Light, Twilight Focus | Holy talents; the Discipline build takes Twilight Focus 3 (healers pushback lane P3), cannot reach Litany (Holy tier 5) | talents/priest.json | aligned |

No Forever-new Discipline mechanic is missing: the tree's new throughput is Penance, Divine Aegis, Renewed Hope and Twin Disciplines, all modelled on the client rows. Forever adds no atonement-style mechanic (Power in Light is damage only).

## Logs

- **No level 60 log of any priest exists.** Raids open 9 Dec; foreverlogs.gg's 26 public reports (2815 to 2854) top out at 43, and the "Discipline" tags there are mostly leveling priests shooting wands (Emanuele, 2848: 426 Shoots, 6 heals). Our own reports (seven, five with fights after the backfill: 3sdrcwdgo2yi, ctxicu4cuyek, eygg26pg4gr2, edcgw6rodqs2, 4q5el2xsgoy7) were read from their raw chunks in R2: 471k lines, all open world (Desolace world PvP, Stormwind, a few dungeon bosses), priests 21 to 36. A Discipline-versus-Holy comparison at 60 in the same fight is not possible yet, so the 32% cannot be confirmed or refuted by a log.
- **L1, shield amounts.** 15 priests applied 85,776 of Power Word: Shield (the absorb is in SPELL_AURA_APPLIED). Each priest's shields are one size per rank and the sizes fit the client row plus 0.1 of the logged spell power times Improved Power Word: Shield (for example 6066 at 461 to 386, 6065 at 377, 3747 at 293 to 298); without the talents in the log (only edcgw6rodqs2 carries COMBATANT_INFO, a Shadow priest with neither shield talent whose 386 is the row plus his power), the talent multiplier cannot be read exactly. Nothing contradicts C1 and C2.
- **L2, shield waste.** Of those 85,776 applied, 48,341 were soaked (56%) and 21,029 were still on the target when the aura fell (expired or the fight moved). The six Penance casters soaked 25 to 83% of theirs (Lil-Bald 38%, Hp-Pottion 40%, Micro-Penances 62%). The sim soaks 99% of the tank's shields: the profile is generous to the shield, not hard on it.
- **L3, the shield's share.** In the leveling logs a Penance priest's shields are 34 to 55% of what she healed and soaked (Lil-Bald 4,696 of 10,788; Micro-Penances 4,011 of 7,254); in the sim's raid fight Power Word: Shield is 4.4% and Divine Aegis 3.2%. The difference is the fight, not the spell: open world and five-man fights are short, burst-heavy and mana-light, where an instant 1.5 s shield on whoever is about to be hit is the right cast; a 300 s fight that runs every healer dry prices the shield by its healing per mana (about 3.2, against Greater Heal rank 1's 5.0, Penance's 6.3 and Renew's 6.5).

## The bug

| # | bug | origin | fix | evidence |
|---|---|---|---|---|
| F1 | **The tank's emergency Flash Heal was never paced.** Both priest cards fire Flash Heal on a tank under 35% with no mana condition, while the party's emergency Flash Heals have waited on 70% of the pace since round 6. In the raid fight the bar runs dry at 91 s and the tank is under 35% often after that, so Discipline cast 27 Flash Heals a fight at 3.48 healing per mana (the worst on its card), 10,000 mana, more than any other spell it casts; each displaced a Greater Heal at 4.8 to 5.0. | old (the card since round 6) | the tank's Flash Heal line waits on the card's own filler pace (mana at least 90% of the share of the fight left) on both cards; Greater Heal on a tank under 50% stays unpaced, so the emergency still has an answer when the bar is behind | `spec-breakdown -heal` (below); the variants that measured it |

**Rotation variants measured first** (raid 60, Alliance, seed 1, 5,000 iterations, effective HPS; the files are in `discipline/`):

| variant | Discipline | Holy |
|---|---|---|
| the card as published | 603.8 | 870.5 |
| shields on the four party members too (under 90%, paced) | 589.3 (-2.4%) | |
| shields on the party whenever Weakened Soul is off (paced) | 583.3 (-3.4%) | |
| shields on the party, unpaced | 534.8 (-11.4%) | |
| the tank shield on Weakened Soul and pace alone | 600.2 (-0.6%) | |
| no Flash Heal at all | 661.2 (+9.5%) | 879.9 (+1.1%) |
| no party Flash Heal | 604.7 (+0.1%) | |
| tank Flash Heal under 20% only | 624.4 (+3.4%) | 884.2 (+1.6%) |
| tank Flash Heal on 70% of the pace | 652.1 (+8.0%) | 890.7 (+2.3%) |
| **tank Flash Heal on 90% of the pace (adopted)** | **654.4 (+8.4%)** | **892.9 (+2.6%)** |

The adopted line keeps Flash Heal on the card (5 casts a fight at raid 60, the first minute and a half, while the bar keeps pace) and gains beyond the error at every cell; dropping it entirely would gain a hair more but leaves no fast heal at all, which the adoption rule does not allow. Shielding the party loses at every gate: a party member takes about 83 damage a second from the pulses, so a 1,220 shield at 375 mana is worse healing per mana than the Greater Heal it displaces, and Weakened Soul is never what holds the shield back. The rotation search (`sim/cmd/rotation-search`) does not run healers (`errSkipped`), so these were measured by hand on `spec-breakdown -rotation`.

## Before and after

`spec-breakdown -heal`, the published band entries, 10,000 iterations, effective HPS; before is the pin 5515ae0ef on the published card, after the same engine on the new card (the fork change is the card copy only). Combined error under 0.07% in every cell.

| cell | Discipline before | after | gain | Holy before | after | gain |
|---|---|---|---|---|---|---|
| raid 60 Alliance s1 | 603.89 | 654.40 | +8.36% | 870.08 | 892.68 | +2.60% |
| raid 60 Alliance s50001 | 603.71 | 654.22 | +8.37% | 869.77 | 892.31 | +2.59% |
| raid 60 Horde s1 | 596.88 | 648.70 | +8.68% | 881.68 | 900.77 | +2.17% |
| raid 60 Horde s50001 | 596.33 | 648.62 | +8.77% | 882.00 | 900.92 | +2.15% |
| bare 60 Alliance | 287.42 | 303.60 | +5.63% | 417.51 | 418.63 | +0.27% |
| bare 60 Horde | 287.30 | 303.01 | +5.47% | 424.34 | 428.29 | +0.93% |
| bare 50 Alliance / Horde | 206.99 / 210.58 | 209.46 / 213.40 | +1.19 / +1.34% | 258.80 / 272.68 | 258.62 / 272.45 | -0.07 / -0.08% (within error) |
| bare 40 Alliance / Horde | 144.56 / 149.01 | 148.76 / 151.86 | +2.91 / +1.91% | 175.92 / 184.46 | 175.92 / 184.44 | 0.00 / -0.01% |
| bare 30 Alliance / Horde | 94.71 / 96.51 | 95.45 / 97.13 | +0.78 / +0.64% | 76.68 / 79.89 | 76.75 / 82.08 | +0.09 / +2.74% |
| bare 20 Alliance / Horde | 46.39 / 45.43 | 47.28 / 46.32 | +1.92 / +1.96% | 41.97 / 40.61 | 42.76 / 41.44 | +1.88 / +2.04% |

Raid 60 tank Flash Heals: Discipline 27 to 5 a fight, Holy 24 to 16; the mana lasts 4 s longer for Discipline and 2 s for Holy. On the tier list's Alliance raid row, Discipline moves from 31% behind Holy to 27% behind, and from 41.7% to about 36% behind the top. The published numbers move at the next ranking (MR94).

## Is the rest of the gap real?

Yes, under this profile, and the engine is right about why. Three things make it:

1. **Mana.** The fight is longer than any priest's bar. Holy gets 12 mana a second back from Litany of Light (a Holy talent Discipline cannot reach) and its best spell, Prayer of Mending, returns 11 healing per mana; Discipline's best spells (Penance 6.3, Renew 6.5) are on a cooldown or a slot, and its fillers return 3.5 to 5. Holy's spend is 41k against Discipline's 37k, at 6.4 healing per mana against 5.0.
2. **Throughput while the mana lasts is close.** Over the first 90 s (`-duration 90`, raid 60 Alliance): Restoration druid 1,092, Holy 1,149, Restoration shaman 995, Discipline 955. Discipline is 17% behind Holy and 4% behind the shaman there; the 300 s fight's mana turns that into 27%.
3. **The profile rewards no overheal-free healing.** Almost nothing overheals in the raid fight (Discipline 3%), because the damage is sized so healers cover 60 to 80% of it. A shield's real-raid edge (it lands before the hit and cannot overheal) is worth nothing when every heal lands on a hurt target anyway.

What Discipline brings that HPS does not show, for the tier page:
- **Shields before the hit.** An instant absorb on whoever is about to take a spike keeps a tank or a cloth wearer alive through damage no heal can answer in time; the sim only credits what the shield soaks, never a death it prevented.
- **Power Infusion.** 20% more spell damage and healing on a mage, warlock or healer for 15 s every 3 minutes: raid damage the healer's own number never shows, and the published build leaves it out because it is worth more on someone else.
- **Divine Aegis and Renewed Hope.** Its crits leave shields and its heals shorten Weakened Soul, so a Discipline priest's tank shield comes back sooner than any other priest's.
- **Penance on short fights.** Instant, cheap, on a 12 s cooldown: in a five-man pull or a short burn it is the best healing per mana a priest has, which a 300 s raid fight undersells.

Suggested tier page line for the Discipline row (the other lane owns the page; nothing here edits it): "Discipline trades long-fight mana for shields that land before the hit and Power Infusion for a caster; this list counts neither."

## Held and open

- **H1, Flash Heal off the cards entirely.** +9.5% Discipline and +1.1% Holy at raid 60 on Alliance seed 1, but it removes the fast heal; the adoption rule keeps it. Not adopted.
- **O1, the profile's overheal.** The profile is sized so healers cover 60 to 80% of what comes in, which makes the raid fight a mana fight with almost no overheal. That is its stated rule (`heal-profile.json` tuning), not a bug, but it is the main reason a shield healer ranks last; whether the healer list should also show a shorter or burstier fight is the owner's call.
- **O2, a level 60 Discipline log.** The first raid logs (9 Dec) settle the shield share, the shield waste and Penance's share in a real raid.
- **O3, shield sizes against talents.** Only one priest log carries COMBATANT_INFO, a Shadow priest; a Discipline priest's log with talents would pin Improved Power Word: Shield's multiplier exactly.

## Guide and card

- Both cards: the tank Flash Heal line's condition and note; the top-level notes record the change. `make apl-sync` wrote the two copies (`sim/request/apl`, the fork's `ui/healing_priest/apls`); fork commit 186440a08.
- Guides: Discipline's Flash Heal paragraph (keep it for a tank about to die while the mana keeps up; behind, a Greater Heal started before the hit), Holy's Flash Heal paragraph likewise; both `updated` stamps read 2026-11-13.
- wow-player: SHIP WITH FIXES for both classes, both fixes taken (the Discipline note's wording, the Holy note's pace covering the party as well); `discipline/wow-player.md`.

## Measure at refresh

- **MR94.** priest-discipline and priest-holy, every bare band 20 to 60 and raid 60: picks, weights (mana stats against healing power move with a card that spends less on Flash Heal), races, the published HPS and the tier list's healer row.

## Gates

See the ledger row and `discipline/gates.sh` (apl-check against the fork, the request suite with the two priest ladder goldens regenerated, the spec-breakdown tests, rank choices, leveling, the guide unit tests, lint, types, the curated APL tests) and the fork's priest suites. The guard test `TestPriestTankFlashHealWaitsOnTheManaPace` (`sim/request/priest_tank_flash_test.go`) fails on the published cards and passes on the new ones.

## Files

- `discipline/` (job runner, job files, the rotation variants and their generators, gates).
- `sim/cmd/spec-breakdown/heal.go`: a "Healing source" table in `-heal` (healing and shields applied against what landed, per action), which is how the shields' soak and Divine Aegis were read.
