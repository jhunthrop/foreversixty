# Sim accuracy 2026-11-13: the raid 60 searches on enchanted gear, every spec but the rogues

Lane `raid-searches`, brief `design/reviews/sim-accuracy/brief.md`. Site from origin/main 25f3696db, rebased on b3a5ddb19. Every search ran on client 1.60.1.70790 (the site's active build; the 70938 client move had not reached the site's data) and engine 79358b3c2, the pin when the lane started; the follow-ups ran on 3db538a02 (the fix below on 79358b3c2) and the gates on the new pin 70dd8eb24.

**Totals: 34 raid searches of each kind (17 DPS specs, both factions), no build or rotation adopted: every gain beyond error is under the 1% bar or fails the rule on its own terms. 2 bugs fixed, both old, neither on a published number: (B1) the engine let any weapon cast Hammer of the Righteous, where the client limits it to a one-handed axe, mace or sword, so the Retribution search "found" +16.2% by pressing it with a two-hander; (T1) `heal-builds`, the healers' build comparison, wore no enchants, the same gap the rogue lane closed in the two searches. Engine pin 70dd8eb24. No adoption, so no guide, card or wow-player round, and no MR row.**

## Method

- `rotation-search` and `talent-search` at level 60, `-preset raid`, on the tools that wear the band's published enchants (18d046e19) and the guide's raid build (b30035273), Alliance and Horde (every spec's raid race differs by faction; no raid build does), seed 7, 300 iterations to screen and 800 to confirm (rotation), the ranker's weights level for the talent finalists. Talent keeps from `design/reviews/talent-search/raid-keeps.json`. One search at a time on the Windows box, `nice -n 15`, `GOMAXPROCS=5`, through `tools/remote-test.sh`, in visibility order: Marksmanship, Feral, Fury, Survival, Frost, Balance, Elemental, Affliction, then the rest. Script `raid-searches/run.sh`; every report under `raid-searches/<faction>/<rotation|talents>/`; the log `raid-searches/search.log.txt`. All 68 searches took 25 minutes of box time.
- The adoption rule ([[forever-talent-adoption-rule]]): a winner replaces the curated build or rotation only with a gain of at least 1% beyond error, every unmodeled guide talent kept, nothing the rotation casts dropped. A gain at the bar is confirmed at 5,000 iterations and a second seed before it lands (`raid-searches/confirm.sh`).
- Fire again on the new pin 70dd8eb24, which carries the crit-flags lane's Ignite fix (its ticks never crit) that landed mid-lane: the curated rotation is still the best found and the guide build still holds, both factions (baselines 813.1 Alliance and 800.7 Horde, -1.2% and -1.0% on Ignite; output `raid-searches/fire-70dd8eb24.out.txt`). No other searched spec casts a spell that fix touches in its raid rotation.
- Baselines now sit on the published raid `set_dps` for the specs nothing has moved since the last rank (Balance 738.6 against 737.9, Feral 1089.0 against 1090.0); the warriors and hunters read lower by the pending MR96 and MR97 re-ranks.

## Results

Rotation: the curated rotation against the search's best. Talents: the guide's raid build against the best variant ("drops an unmodeled guide talent" means the winner fails the rule's keep clause). Baseline DPS at 800 iterations, seed 7.

| spec | faction | baseline | rotation search | talent search | result |
|---|---|---|---|---|---|
| Marksmanship | Alliance | 1088.0 | +0.7% (+8.0 ±2.2) | guide best | held: under the 1% bar |
| Marksmanship | Horde | 1067.5 | +0.5% (+4.9 ±2.2) | guide best | held: under the 1% bar |
| Feral | Alliance | 1089.0 | +0.4% (+3.9 ±1.8) | guide best | held: under the 1% bar |
| Feral | Horde | 1076.9 | curated best (+0.0 ±1.9) | guide best | within error |
| Fury | Alliance | 1004.6 | +0.8% (+8.5 ±2.5) | guide: 2 Enrage -> Blood Craze: +0.4% (+4.2 ±2.6) | held: under the 1% bar |
| Fury | Horde | 1009.6 | curated best (+0.0 ±2.7) | guide best | within error |
| Survival | Alliance | 975.1 | +0.6% (+6.0 ±2.7) | guide best | held: under the 1% bar |
| Survival | Horde | 968.6 | +0.3% (+2.9 ±2.7) | guide best | held: under the 1% bar |
| Frost | Alliance | 696.2 | curated best (+0.0 ±1.7) | guide holds (best +0.5 ±1.7) | within error |
| Frost | Horde | 687.9 | +0.4% (+3.0 ±1.7) | guide holds (best +0.2 ±1.7) | held: under the 1% bar |
| Balance | Alliance | 738.6 | curated best (+1.7 ±1.8) | guide best | within error |
| Balance | Horde | 723.3 | curated best (+0.0 ±1.8) | guide best | within error |
| Elemental | Alliance | 749.5 | curated best (+0.0 ±1.6) | deep Enhancement: +0.6% (+4.2 ±1.5), drops an unmodeled guide talent | held: under 1%, and the winner drops an unmodeled guide talent |
| Elemental | Horde | 751.6 | curated best (+0.0 ±1.5) | guide, all non-damage points re-spent: +0.4% (+3.0 ±1.5), drops an unmodeled guide talent | held: under 1%, and the winner drops an unmodeled guide talent |
| Affliction | Alliance | 768.2 | curated best (+0.0 ±1.4) | guide, modeled and survival-only non-damage points re-spent + 2 Grim Reach -> Improved Life Tap: +0.2% (+1.8 ±1.4), drops an unmodeled guide talent | held: under 1%, and the winner drops an unmodeled guide talent |
| Affliction | Horde | 757.9 | curated best (+0.0 ±1.4) | guide, modeled and survival-only non-damage points re-spent + 2 Grim Reach -> Improved Life Tap: +0.3% (+2.3 ±1.4), drops an unmodeled guide talent | held: under 1%, and the winner drops an unmodeled guide talent |
| Beast Mastery | Alliance | 971.3 | curated best (+0.0 ±1.4) | guide, modeled and survival-only non-damage points re-spent + 1 Efficiency -> Improved Stings: +0.6% (+5.7 ±1.4) | held: under the 1% bar |
| Beast Mastery | Horde | 962.8 | curated best (+0.0 ±1.4) | guide, modeled and survival-only non-damage points re-spent + 1 Efficiency -> Improved Stings: +0.6% (+5.4 ±1.4) | held: under the 1% bar |
| Arms | Alliance | 936.9 | +1.0% (+9.4 ±2.8) | guide best | held: +0.8 to +0.9% at 5,000 iterations, seeds 7 and 50001, under the 1% bar |
| Arms | Horde | 924.1 | +1.0% (+9.6 ±2.9) | guide holds (best +0.3 ±2.9) | held: +0.8 to +0.9% at 5,000 iterations, seeds 7 and 50001, under the 1% bar |
| Retribution | Alliance | 835.4 | +16.2% (+135.1 ±2.7) | guide best | held, engine bug: the winner cast Hammer of the Righteous with a two-hander; fixed (B1), the curated rotation is best on the fix |
| Retribution | Horde | 840.0 | +16.2% (+136.2 ±2.7) | guide best | held, engine bug: the winner cast Hammer of the Righteous with a two-hander; fixed (B1), the curated rotation is best on the fix |
| Shadow | Alliance | 797.6 | +4.0% (+31.9 ±1.4) | guide best | held: the winner drops Vampiric Embrace, the party heal a solo sim only sees as a cost |
| Shadow | Horde | 785.6 | +3.7% (+29.2 ±1.4) | guide best | held: the winner drops Vampiric Embrace, the party heal a solo sim only sees as a cost |
| Enhancement | Alliance | 874.8 | +0.9% (+8.0 ±3.3) | guide, modeled and survival-only non-damage points re-spent + 2 Maelstrom Weapon -> Ancestral Knowledge: +0.8% (+7.0 ±3.3) | held: under the 1% bar |
| Enhancement | Horde | 856.8 | curated best (+2.7 ±3.3) | guide holds (best +3.3 ±3.4) | within error |
| Arcane | Alliance | 803.9 | curated best (+1.9 ±2.2) | guide best | within error |
| Arcane | Horde | 789.0 | +0.8% (+6.6 ±2.2) | deep Arcane: +0.4% (+3.1 ±2.3) | held: under the 1% bar |
| Fire | Alliance | 822.6 | curated best (+0.0 ±2.4) | guide holds (best +0.9 ±2.4) | within error |
| Fire | Horde | 809.1 | curated best (+1.0 ±2.2) | guide holds (best +1.0 ±2.2) | within error |
| Demonology | Alliance | 817.0 | curated best (+0.0 ±2.3) | guide holds (best +0.6 ±2.3) | within error |
| Demonology | Horde | 800.8 | +0.3% (+2.1 ±2.1) | guide, modeled and survival-only non-damage points re-spent + 1 Demonic Embrace -> Improved Sayaad: +0.3% (+2.3 ±2.1) | held: under the 1% bar |
| Destruction | Alliance | 885.2 | curated best (+0.0 ±1.9) | guide holds (best +0.2 ±2.0) | within error |
| Destruction | Horde | 879.4 | curated best (+0.0 ±1.9) | guide best | within error |

Tanks (Protection warrior, Protection paladin, Bear) and healers (Holy and Discipline priest, Holy paladin, Restoration shaman and druid) have no search: both tools exit 3 ("DPS is not its objective"), by design. The healers' build comparison is `heal-builds`, and it had the same bug (T1, fixed); its comparisons are hand-picked, not searched, so nothing re-runs by itself.

### The gains under the bar (held)

| spec | change | gain | note |
|---|---|---|---|
| Arms | Slam above Execute in the execute phase | +0.8 to +0.9% (5,000 iterations, seeds 7 and 50001, both factions; ±1.1) | the 800-iteration +1.0% did not hold |
| Enhancement (Alliance) | Flame Shock on cooldown; Ancestral Knowledge for 2 Maelstrom Weapon | +0.9%, +0.8% | Horde within error on both |
| Fury (Alliance) | Pummel on cooldown; Blood Craze for 2 Enrage | +0.8%, +0.4% | Horde within error |
| Arcane (Horde) | a Scorch line; deep Arcane | +0.8%, +0.4% | Alliance within error |
| Marksmanship | Arcane Shot above Serpent Sting (Alliance), Serpent Sting above Sniper Shot (Horde) | +0.7%, +0.5% | the two factions disagree on the swap |
| Beast Mastery | Improved Stings for 1 Efficiency | +0.6% both factions | clean, under the bar |
| Elemental | deep Enhancement, or every non-damage point re-spent | +0.6%, +0.4% | drops an unmodeled guide talent |
| Survival | Immolation Trap above Strider Kick | +0.6%, +0.3% | |
| Feral (Alliance) | no Ravage opener | +0.4% | Horde within error |
| Frost (Horde) | Ice Barrier refreshed under 3 s | +0.4% | Alliance within error |
| Demonology (Horde) | Corruption refreshed under 3 s; Improved Sayaad | +0.3%, +0.3% | Alliance within error |
| Affliction | Improved Life Tap for Grim Reach | +0.2 to +0.3% | drops Grim Reach (unmodeled range) |

### Shadow: Vampiric Embrace (held)

The search drops Vampiric Embrace (15286) for +4.0% Alliance and +3.7% Horde: one GCD and 40 mana every 30 s (its 70425 cooldown) that a solo sim scores only as a cost, where the real payoff is 20% of the priest's Shadow damage healed to the party. The talent is on the raid keeps for that reason and the rotation keeps the cast. This is the tool's own caveat on group buffs, not a rotation change.

## B1: Hammer of the Righteous with any weapon (CONFIRMED-BUG, engine, fixed)

**Finding.** The Retribution raid rotation search inserted Hammer of the Righteous (407632, learned at 40 on the Protection line) at the top for +135 DPS (+16.2%, both factions). The Retribution raid set wears a two-hander. The client's SpellEquippedItems row 251475 for 407632 (1.60.1.70888, hotfixed, wago.tools) is item class 2 with subclass mask 145: one-handed axes (bit 0), one-handed maces (bit 4) and one-handed swords (bit 7). The fork's `paladin/hammer_of_the_righteous.go` had no weapon condition.

**Fix.** `hammerOfTheRighteousWeaponAllowed` holds the cast to a one-handed axe, mace or sword in the main hand (`ExtraCastCondition`). Red first: `TestHammerOfTheRighteousNeedsAOneHandedAxeMaceOrSword` fails to build at 629e43041 and passes at 3db538a02; on fork master as 07f5a9f99 (red) and 70dd8eb24 (fix), on top of the crit-flags lane's 8f61b9e6a.

**Reach.** No published number moves. Retribution's curated rotation never casts it (on the fix the search's insertion probe reads +0.0 and the curated rotation is the best found, both factions). Every Protection band wears a one-handed mace, sword or axe (20 to 60, bare and raid), and the fork's Protection golden (a one-handed mace) is unchanged; the fork's paladin, Holy, Protection and Retribution suites pass with the item database. The site's Protection rotation ladder moves: it wields the highest item level weapon at each level, and at 40 and 50 that is Ruthless Shiv (7758), a polearm, so Hammer of the Righteous leaves its cast list there (54.7 to 48.1 and 69.8 to 63.4 on that synthetic ladder; golden regenerated). Holy carries no Hammer line.

## T1: heal-builds wore no enchants (CONFIRMED-BUG, tool, fixed)

**Finding.** `sim/cmd/heal-builds` built its gear from the band's `item_id`s alone; the healer raid bands wear healing enchants (Priest Holy raid 60: 1509 on head and legs, 1891 chest, 1901 wrist, 8212 main hand). Every healer build comparison since the tool existed ran without them.

**Fix.** `bisBand.gear()` wears `enchant_id`, as `spec-breakdown` and the two searches do. Red first: `TestBuildRequestWearsTheBandsPublishedEnchants` fails at e80f024f6 (`gear enchants map[head:0 main_hand:0]`) and passes at c0f2d1618.

**Reach.** No published number moves: `heal-builds` only compares hand-picked builds. The healer guides' builds were chosen on unenchanted gear; their next healer lane should re-run its comparison (Open).

## Held

- Every gain in the tables above: under 1%, or a dropped unmodeled talent, or (Shadow) a group buff.
- Arms' Slam over Execute: +0.8 to +0.9% at 5,000 iterations; a candidate for the next warrior lane to read against a log.

## Open

- Tanks and healers: no raid talent search exists for either role (`talent-search` and `rotation-search` skip them by design). The healers' `heal-builds` comparisons should be re-run on the fixed tool by the next healer lane.
- Hammer of the Righteous shares its cooldown with Crusader Strike in the client text ("Cooldown shared with Crusader Strike."); no Forever paladin build learns Crusader Strike today, so unchecked.

## Sim changes in this landing

- **Fork:** 07f5a9f99 (red), 70dd8eb24 (Hammer of the Righteous weapon rule), on master.
- **Site:** e80f024f6 (red), c0f2d1618 and the slot type (heal-builds wears the published enchants); `engine: pin 70dd8eb24`; the Protection ladder golden; this report, `raid-searches/` (scripts, reports, logs) and the ledger row. No build, rotation, guide, card or BiS change, so no wow-player round and no MR row.

## Sources

- wago.tools 1.60.1.70888 with `useHotfixes=1`: SpellEquippedItems (row 251475), Spell (407632 description), SpellName.
- The search reports under `raid-searches/`.
