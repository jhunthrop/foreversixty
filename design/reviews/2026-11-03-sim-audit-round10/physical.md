# Sim audit, round 10: physical group (2026-11-02)

This lane covers Arms, Fury, Protection, Assassination, Combat, Subtlety, Beast Mastery, Marksmanship and Survival (melee), on client 1.60.1.70707.

- **Baseline:** fork pin 4d3129cae with site origin/main bb619a188 (the shared `r10/sb-before`).
- **Branches:** fork fixes on `lane/r10-physical` (E), report on `wt/r10-physical` (S). No site code changed.
- **Measurements:** `sim/cmd/spec-breakdown`, 10,000 iterations, seeds 1 and 50001, both factions, on the published sets.
- **Client rows:** the layered 70707 raw tables with the hotfix push (SpellEffect, SpellClassOptions, SpellAuraOptions, SpellMisc, SpellCategories, ItemEffect, ItemXItemEffect), as fetched for round 9 (2026-11-01 21:42Z), and the 70707 talent files. Class masks below are SpellClassOptions rows read against each talent's EffectSpellClassMask.

## Assigned items: the crit split's side effects

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | **Combat's poisons took Hack and Slash's dagger crit.** Hack and Slash gives daggers and fists +1% crit a point "to your melee weapon attacks". The engine put it on the every-table crit, so a Combat rogue's Instant and Deadly Poison and the Lightning procs (all spell-table rolls) took +5% crit. Poisons crit 26% on the raid set; they now crit 21%. Combat moves -0.13 to -0.15% raid and -0.27 to -0.39% bare at 40 to 60. | 13960 is aura 4 (a server dummy, so the text is the only scope); talents/rogue.json rank text "Gives your melee weapon attacks a benefit ... Dagger/Fist: Increases your critical strike chance". Poisons are DefenseType 1 (round 9 crit.md). Test `TestHackAndSlashDaggerGrantsCrit` now asserts melee crit only, and failed first. | CONFIRMED-BUG (single source for the scope: the client text) | E 5849efeca |
| A2 | **Weaponmaster's axe and polearm crit reached spell-table procs.** Same shape as A1: "Gives your melee weapon attacks a benefit ... Axe/Polearm: Increases your critical strike chance by 5%". It was every-table crit; it is now melee crit. No published number moves (the axes in published Arms sets carry no spell proc). Round 9's crit table listed Weaponmaster as aura 290; that row is the old 12700. Forever's node 1290261 is a dummy. | 1290261 aura 4; talents/warrior.json; test `TestWeaponmasterAxeCritIsMeleeOnly` failed first | CONFIRMED-BUG (0 published DPS) | E b71cc6975 |
| A3 | Rogue crit talents, each against its mask:<br>- **Malice** 14138 is aura 290 and its text says "all attacks and Poisons". The poisons should take it, and they do.<br>- **Lethality** 14128 (aura 108, mask 100663310 / 6) covers Sinister Strike, Gouge, Backstab, Mutilate, Ghostly Strike and Hemorrhage. It does not cover finishers. The engine applies it to those five (Gouge is never cast) at 4% a point, so 5/5 crits for 220%.<br>- **Puncturing Wounds** 1224716: +10% a point to Backstab (mask 4) and +5% to Mutilate (mask 0 / 6), plus 15% for an extra combo point. Aligned.<br>- **Improved Ambush** 14079 (mask 512) gives +15% a point. Aligned.<br>- **Cold Blood** reaches only its flagged abilities.<br>- **Seal Fate** (500 ms internal cooldown) fires on crits by builders only. Aligned.<br>- **Relentless Strikes** has no crit in it. It fires on every finisher at 20% a point. Aligned.<br>- Improved Backstab is not in Forever's tree; its line stays commented out. | SpellEffect / SpellClassOptions rows above; `sim/rogue/talents.go`, `backstab.go`, `mutilate.go` | aligned | - |
| A4 | Poisons and the crit table: they roll spell crit (base 0 for a rogue, no Agility). They take Malice and the every-table sources, and since A1 nothing else. They take the merged hit stat. Forever's hit talents are aura 54 plus 55 (Precision 13705; "Increased Hit Chance" 15464), so merged hit is right. | SpellEffect 13705, 15464; `core/spell_result.go:218` | aligned | - |
| A5 | Warrior crit talents:<br>- **Impale** 16493 (aura 108, mask 0xee604cee / 33088 / 1 / 268436480) covers every ability the engine gives it, plus Rend's ticks. Deep Wounds' bleed is not in the mask, and the engine leaves it out. Aligned.<br>- **Improved Overpower** 12290: +25% a point on Overpower (mask 4). Aligned.<br>- **Deep Wounds** fires on physical crits with a non-empty proc mask, from any table (client type mask 0x11154).<br>- **Flurry**'s trigger (12319) is a dummy whose text says "melee critical strike". The engine requires a melee crit. Its charges (12966, type mask 4) are spent by swings, Heroic Strike included.<br>- **Enrage** fires on being hit, which is not crit-dependent.<br>- **Cruelty** is melee crit (round 9).<br>- **Berserker Stance** 7381 and **Recklessness** 1719 are aura 290 (all tables). | rows above; `sim/warrior/talents.go`, `deep_wounds.go`, `stances.go` | aligned | - |
| A6 | **Lacerate and Lacerating Strikes missed Savage Strikes and Predator's Edge.** Lacerate (1299332) and the Lacerating Strikes bleed (1310536) share one class flag, mask 1 bit 4096. Savage Strikes (19159, aura 107 crit chance, 66 / 528384 / 2097152) and Predator's Edge (1310627, aura 108 crit damage, 66 / 528384 / 1075904512) both carry that bit. Lacerate had Savage Strikes' 4% but not Predator's Edge's +30% crit damage. Lacerating Strikes had neither. Both bleeds crit under the owner's rule, so both now take both talents. Survival moves +0.5% raid and +0.1 to +0.5% bare from 30. | SpellClassOptions 1299332, 1310536; SpellEffect 19159, 1310627; test `TestTheBleedsTakeSavageStrikesAndPredatorsEdge` failed first | CONFIRMED-BUG | E 1f6097bcc |
| A7 | **Summon Hawk took Mortal Shots, and neither hawk talent.** The hawk (1293241) carries class mask 2 bit 8192. Mortal Shots' mask (424449 / 2147483648 / 268435968 / 8388608) does not contain that bit. Ferocity's mask (1073741824 / 0 / 8192) does, and both Ferocity and Unleashed Fury say "your pets and hawks". The hawk now takes Ferocity's crit and Unleashed Fury's damage, without Mortal Shots. Ranged Weapon Specialization no longer reaches it either, because the dive is not a weapon attack. No published build takes Summon Hawk. | SpellClassOptions 1293241; SpellEffect 19485, 19598; talents/hunter.json; test `TestSummonHawkTakesTheHawkTalentsNotMortalShots` failed first | CONFIRMED-BUG (0 published DPS) | E 37e83a1a7 |
| A8 | Other hunter crit talents:<br>- **Mortal Shots** covers Auto Shot, Aimed Shot, Multi-Shot, Arcane Shot, Serpent Sting, Volley, Scatter Shot, Concussive Shot and Sniper Shot. The engine's list matches.<br>- **Lethal Attacks** 19426 is aura 290, as the engine has it.<br>- **Savage Strikes** and **Predator's Edge** reach the other melee specials.<br>- **Frenzy** fires on any pet crit, 20% a point.<br>- **Ferocity** gives pet crit, 2% a point.<br>- Killer Instinct is not in Forever's tree; its line stays commented out. | SpellEffect / SpellClassOptions rows; `sim/hunter/talents.go` | aligned | - |
| A9 | The pet's crit after the split: the pet copies the hunter's every-table and melee crit (415429's aura 52 channel). Its white crit is 30% against the hunter's 19.5% ranged, because Ferocity adds 10% and the Agility crit is melee and ranged. That is the plumbing the crit lane owns. Nothing here contradicts it. | BM band-60 raid breakdown | for the crit lane (no finding) | - |
| A10 | **Ammo and pets by level (round 9) at every band.** At band 20 bare, Beast Mastery shoots level-appropriate arrows. Its pet hits for level-20 numbers (the cat's white hit averages 15). The ranks are right for the band: Serpent Sting 13550, Arcane Shot 14282, Claw 16829 and Bite 17256. Band 60 is unchanged from round 9. | `bd/` breakdowns, bare 20 and raid 60 | aligned | - |
| A11 | The bare BM hunter at band 20 is out of mana for 100 of 180 s. It casts Multi-Shot until it is dry, then auto-shoots. That ordering does not change the total. Whether a downrank pays at 20 is the ranks lane's per-band question. | bare 20 breakdown | for the ranks lane (no finding) | - |

## Round 9 holds and opens

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | **Unbridled Wrath on Heroic Strike is settled: it procs, as the engine has it.**<br>- Unbridled Wrath 12322 has proc type mask 4, the melee swing.<br>- Windfury Totem 8516, 10608 and 10610 carry the same mask 4, and Windfury procs from Heroic Strike in Classic.<br>- Heroic Strike is an on-next-swing ability (SpellMisc Attributes_0 0x50014), so its swing raises the swing flag.<br>- Warcraft Wiki's Unbridled Wrath page says Heroic Strike and Cleave "do have a chance to gain extra Rage". One older Blizzard forum bug list claims the opposite, for an unstated version.<br>The engine gives Heroic Strike both the special and the auto proc mask, so Unbridled Wrath, Windfury and Flurry's charges all treat it the same way. | SpellAuraOptions 12322, 8516, 10610; SpellMisc 11567; https://warcraft.wiki.gg/wiki/Unbridled_Wrath_(warrior_talent) (read 2026-11-02); https://eu.forums.blizzard.com/en/wow/t/379954/148 | aligned (two sources; one dissent) | - |
| H2 | Normalized rage: the off-hand factor (0.5) and the crit multiplier (2.0). The client has no rage formula; rage is server code. | `sim/core/rage.go` | OPEN: settle with a level-60 log or a dummy test of rage per off-hand hit and per crit | - |
| H3 | Lacerating Strikes: a new Mongoose Bite replaces the bleed with no rollover. The client row 1310533 is a dummy (aura 4 at 40), and the bleed 1310536 is a plain aura 3 every 3 s. Neither says how a refresh treats the remainder. A6 fixes what reaches its ticks. | SpellEffect 1310533, 1310536 | OPEN: settle with a log of the bleed's tick sizes across two Mongoose Bites inside 21 s | - |
| H4 | Marksmanship's +0.8 to +0.9% build stays held. No fix in this lane moves Marksmanship (raid and bare are level to the decimal). | measurements below | HOLD | - |
| H5 | The Survival guide's raid-card line was not touched (platform lane). | - | - | - |

## Fresh re-audit

What was checked per spec, beyond the rows above. Each published band-60 raid breakdown was read spell by spell: casts, crit, miss and dodge rates, ranks and shares. Every published band-60 raid item with a non-stat equip or use effect was followed to its client spell, and the engine was checked for a model of it.

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | **Arms:**<br>- Mortal Strike, Overpower (77.7% crit = 30 + 50 from Improved Overpower), Execute, Slam, Rend ticks with crits, and Hamstring weaving for Strength of the Champion and Bloodthrill.<br>- Recklessness, Blood Fury and Bloodrage.<br>- Molten Heart's Molten Fury (1249113: aura 99/124, modelled) and Arcanite Champion's Strength of the Champion (16916, modelled).<br>- Deep Wounds 12867 ticks. | `bd/warrior-arms-raid60.txt` | aligned | - |
| F2 | **Fury:**<br>- Execute's 51.7% crit is Recklessness and Death Wish landing in the execute phase.<br>- Bloodthirst, Whirlwind, Heroic Strike queueing (34.6 queued, 32.3 swung), off-hand miss 18% with Furious Precision, and Berserking.<br>- The published raid set carries no item with a non-stat effect. | `bd/warrior-fury-raid60.txt` | aligned | - |
| F3 | **Protection:**<br>- Shield Slam, Revenge, Thunder Clap and Shield Block uptime, as round 9.<br>- Its trinket **Arena Grand Master** (19024) has an equip effect, 1318318: aura 87 at -6, all schools, "damage taken -6%". simdb lists aura 87 as ignored, and the engine has no model for it. The tank fight therefore over-counts damage taken. Through rage from damage taken, it also slightly over-counts the tank's rage. Its use (23506, a 1080 absorb) is also unmodelled.<br>- Guardian Talisman's 1300364 is aura 152 (detection range). It has no fight effect. | ItemXItemEffect → ItemEffect 19024; SpellEffect 1318318, 23506, 1300364; `data/pipeline/simdb/equip.py` IGNORED_AURAS | CONFIRMED-BUG (unmodelled equip effect; items lane: aura 87 is item plumbing) | - |
| F4 | **Assassination:**<br>- Mutilate's two hands (43% crit: base plus Puncturing Wounds' 15%) and Eviscerate at 4 and 5 points.<br>- Cold Blood on Eviscerate, Slice and Dice upkeep, poisons at 21.5% and 19.1% crit, and Instant Poison miss 11.2% (no Precision).<br>- Lightning from the Stormshroud 2-piece, and Revitalize from the 3-piece (23863, 4%). | `bd/rogue-assassination-raid60.txt`; client set rows | aligned | - |
| F5 | **Combat:**<br>- Backstab with daggers (61.9% crit with Puncturing Wounds 3/3), Eviscerate, Blade Flurry and Adrenaline Rush.<br>- Frozen Heart's Chthonic Power (1249110, aura 13, modelled).<br>- Poisons after A1. | `bd/rogue-combat-raid60.txt` | aligned after A1 | - |
| F6 | **Subtlety:**<br>- Hemorrhage, Ghostly Strike pooling (8.7 a fight, from round 9), and Rupture at 4 and 5 points with crits.<br>- Premeditation and Ambush opener, and the Cold Blood Eviscerate (72% crit).<br>- Poisons at 18% (Malice 5). | `bd/rogue-subtlety-raid60.txt` | aligned | - |
| F7 | **Beast Mastery:**<br>- Auto Shot, Multi-Shot, Arcane Shot, Serpent Sting with no clipped ticks (57.6 ticks from 12.7 casts), and Aimed Shot as an opener only.<br>- The cat's Bite and Claw, Bestial Wrath twice, and Rapid Fire.<br>- Aspect of the Hawk's Quick Shots. | `bd/hunter-beast-mastery-raid60.txt` | aligned | - |
| F8 | **Marksmanship:** the same shots, with Lethal Attacks in the 18.8% ranged crit. The cat is at half Beast Mastery's share, as expected without the Beast Mastery talents. | `bd/hunter-marksmanship-raid60.txt` | aligned | - |
| F9 | **Survival:**<br>- No pet; Lone Wolf (415370, +20% with no pet) is in the build. That explains Teebu's Firebolt (1300753, 140 Fire, DefenseType 1) averaging about 207 on the raid set.<br>- Raptor Strike, Mongoose Bite (19.5 a fight from Expose Prey on the marked target), Strider Kick, Wing Clip, Lacerate, Immolation Trap (DefenseType 1, rolls spell crit) and Counterattack.<br>- The bare kit carries Hunter's Mark, so Expose Prey works bare.<br>- The APL's Mongoose Bite note says "5 second window" for both triggers; Expose Prey's window has been 10 s since 70291. The engine has 10 s, so this is a prose slip only. | `bd/hunter-survival-raid60.txt`; SpellCategories 1300753, 13797; `sim/hunter/mongoose_bite.go:93`; `data/curated/apl/hunter-survival.json:89` | aligned (note wording only) | - |

## Before/after

**Before:** fork 4d3129cae with site bb619a188 (`r10/sb-before`). **After:** `lane/r10-physical` 37e83a1a7 with the same site.

Band 60, raid preset (Protection in threat per second). Each cell is seed 1 / seed 50001:

| Spec | alliance % | horde % | seed 1 alliance / horde |
|---|---|---|---|
| Arms | +0.00 / +0.00 | +0.00 / +0.00 | 708.9 → 708.9 / 720.8 → 720.8 |
| Fury | +0.00 / +0.00 | +0.00 / +0.00 | 753.9 → 753.9 / 757.8 → 757.8 |
| Protection (threat) | +0.00 / +0.00 | +0.00 / +0.00 | 1062.9 → 1062.9 / 1051.4 → 1051.4 |
| Assassination | +0.00 / +0.00 | +0.00 / +0.00 | 616.3 → 616.3 / 613.6 → 613.6 |
| Combat | -0.13 / -0.15 | -0.15 / -0.15 | 611.5 → 610.7 / 611.8 → 610.9 |
| Subtlety | +0.00 / +0.00 | +0.00 / +0.00 | 538.5 → 538.5 / 538.3 → 538.3 |
| Beast Mastery | +0.00 / +0.00 | +0.00 / +0.00 | 629.4 → 629.4 / 633.4 → 633.4 |
| Marksmanship | +0.00 / +0.00 | +0.00 / +0.00 | 618.5 → 618.5 / 623.0 → 623.0 |
| Survival | +0.52 / +0.51 | +0.52 / +0.52 | 731.4 → 735.2 / 734.9 → 738.7 |

Bare bands for the specs a fix touches (Combat, Arms, Survival; every other band level):

| Spec, band | alliance % | horde % | seed 1 alliance / horde | cause |
|---|---|---|---|---|
| Combat 40 | -0.27 / -0.27 | -0.27 / -0.27 | 110.3 → 110.0 / 109.2 → 108.9 | A1 |
| Combat 50 | -0.32 / -0.32 | -0.32 / -0.39 | 154.4 → 153.9 / 154.5 → 154.0 | A1 |
| Combat 60 | -0.33 / -0.28 | -0.33 / -0.33 | 214.2 → 213.5 / 213.9 → 213.2 | A1 |
| Survival 30 | +0.13 / +0.13 | +0.12 / +0.12 | 79.3 → 79.4 / 80.3 → 80.4 | A6 |
| Survival 40 | +0.19 / +0.19 | +0.20 / +0.20 | 154.6 → 154.9 / 152.8 → 153.1 | A6 |
| Survival 50 | +0.28 / +0.32 | +0.30 / +0.30 | 217.0 → 217.6 / 201.9 → 202.5 | A6 |
| Survival 60 | +0.50 / +0.50 | +0.51 / +0.51 | 298.7 → 300.2 / 295.2 → 296.7 | A6 |

Combat 20 and 30, Arms 20 to 60 and Survival 20 are level.

## For other lanes

- **Items lane:** Arena Grand Master (19024) is Protection's published trinket.
  - Its equip effect 1318318 (aura 87, -6% damage taken, all schools) is dropped: aura 87 is in simdb's IGNORED_AURAS. Its use, 23506 (a 1080 absorb), is also unmodelled.
  - Five other items carry an aura-87 equip spell: 211500, 211856, 211857 (430432, -5%), 18706 Arena Master (1317285, -5%) and 9509 (1293003, +10% to school 4).
  - It needs either a pipeline stat for "damage taken %" or an engine item effect. That is item plumbing, so it is left to the items lane (F3).
- **Crit lane:** no plumbing finding. A1 and A2 are class talent files and changed only the stat each one writes. Round 9's crit table lists Weaponmaster as aura 290; Forever's node 1290261 is a dummy, so the row should read "melee (text)".
- **Integrator:**
  - TestRotationLadder rows move for rogue-combat (E 5849efeca) and hunter-survival (E 1f6097bcc). They were not regenerated here.
  - No fork golden moved: `go test -tags=with_db ./sim/rogue/... ./sim/warrior/... ./sim/hunter/...` passes at 37e83a1a7.
  - No commit here moves a spec outside this group. Hack and Slash, Weaponmaster and the hunter talents are class-only.

## Still open

- Normalized rage's off-hand factor and crit multiplier (H2).
- Lacerating Strikes' refresh rule (H3).
- Marksmanship's held build (H4).
- Arena Grand Master's damage-taken equip effect, for the items lane (F3).
- Carried from round 9: armor-debuff stacking under aura 674; Eviscerate's and Rupture's attack power; Venom's model; Combat's Slice and Dice gate; the Hawk proc mask; Serpent Sting's and Arcane Shot's spell power; threat values; weapon-proc rates; warlock demon sheets.

Sources (read 2026-11-02): the layered 70707 raw tables above; talents/rogue.json, warrior.json, hunter.json for 1.60.1.70707; https://warcraft.wiki.gg/wiki/Unbridled_Wrath_(warrior_talent); https://eu.forums.blizzard.com/en/wow/t/379954/148.
