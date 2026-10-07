# Racials settled from the client tables (build 1.60.1.70009)

Fork branch `racials-client` (head d1af3529f), site branch `racials-client`.
Sources: data/builds/1.60.1.70009/raw SpellName, Spell (Description_lang),
SpellEffect, SpellMisc + SpellDuration, SpellCooldowns, SpellAuraOptions;
text cross-checked against nether.wowhead.com/forever/tooltip/spell/<id>.
Engine code: sim/core/racials.go, racial_blood_fury.go,
racial_touch_of_the_grave.go. A test per changed racial pins the numbers
(sim/core/racials_test.go).

## Racial by racial

| Racial | Client rows | Engine before | Engine now |
|---|---|---|---|
| Touch of the Grave (Undead) | 1260201: ProcChance 10, ICD 1000 ms, ProcTypeMask 69972 (0x11154); drain 1260198: effect 9 health leech, 5 of caster max Health, Shadow | nothing registered | proc on landed melee, ranged and spell hits (no periodic ticks), 10%, 1 s ICD, flat Shadow damage = 5% of max Health |
| Eureka! (Gnome) | 1259821 (Mana): charges 3, 15 s (duration 8), 120 s cooldown, -10% cost, +10% damage; variants Rage 1259813, Energy 1259812, healer 1259823 | id 1_000_101, 50% cost, permanent until 3 stacks spent | client id, 10% cost, 3 charges, 15 s |
| Blood Fury (Orc) | 20572: auras 166, 167, 317 at 10; 15 s; 120 s | +25% of base AP plus strength and agility AP, attack power only | +10% of total attack power, ranged attack power and spell power (dynamic multiplier, so it follows later buffs); school powers get 10% snapshotted at gain |
| Berserking (Troll) | 20554: auras 319, 140, 65 at 10; 10 s; 180 s; no power cost | health-scaled 10 to 30%, 1/(1-x) for mana users (11.1%), 5 rage, 10 energy or 7% mana | flat 10% casting and attack speed, no cost, 10 s, 3 min |
| Expansive Mind (Gnome) | 20591 mana, 1259802 rage, 1259803 energy: aura 178 at 5 | 5% each | unchanged, confirmed |
| Quickness (Night Elf) | 20582: dodge 1, speed 2 | 1% dodge | unchanged, confirmed |
| Elune's Light (Night Elf) | 1259799: aura 290 at 10, 15 s, 180 s | right numbers, id 58984 | id 1259799 |
| Wind Blessed (Skyborne) | 1259710: auras 342 and 65 at 1 | melee and cast haste | adds ranged haste |
| Elemental Insight (Skyborne) | 1259707: aura 168 at 5, Elementals | 5% | unchanged, confirmed |
| Beast Slaying (Troll) | 20557: aura 168 at 5, Beasts | 5% damage and a further 5% crit multiplier | 5% damage only |
| Big Game Hunter (Dwarf) | 1259721: aura 168 at 5, Beasts | no effect, unconfirmed | 5% versus Beasts |

Touch of the Grave decisions: the ProcTypeMask has no periodic bit, so DoT
ticks never trigger it; the drain is a health leech with no coefficient, so
it is flat (no crit, no spell power, no attacker or target modifiers, no
resistance) and equals exactly 5% of the undead's maximum Health. The heal
to the caster is not modelled.

## What moved

Fork goldens (relative to the previous fork head): Orc DPS warrior about
-2.4% on average (-1.7% to -3.1%; Blood Fury fell from 25% of a base to 10%
of the total), Gnome mage -0.4% on average (to -1.6%), Troll mage +0.2%,
Troll elemental and enhancement shaman about +1.0% to +1.3% (cost and the
11.1% reading gone), Orc enhancement -0.0%.

Site ladder goldens, level 60. The ladder runs each spec as the lowest race
id the class may be (Human for the mages) and Undead for priest, so only
priest-shadow, druid and shaman rows moved: priest-shadow 189.1 to 195.0
(+3.1%); druid-balance and druid-feral only relabel Elune's Light 58984 to
1259799; shaman-elemental, shaman-enhancement about +0.4%.

One-off runs with the mage race forced (not committed), relative level 60:

| Spec | Undead | Gnome |
|---|---|---|
| mage-arcane | +0.9% | -1.5% |
| mage-fire | +5.1% | -1.4% |
| mage-frost | +2.6% | -2.7% |
| priest-shadow (Undead) | +3.1% | n/a |

Gnome falls because Eureka! now costs 10% instead of 50% and lasts 15 s;
Undead rises on Touch of the Grave, most for fire (many cheap direct hits).

## Still open

- Unconfirmed (3): the Undead fourth racial (no source names it), Tauren
  Cultivation and Plainsrunning (which is the second active is unread).
- Eureka! is modelled for magic schools only. The client has Rage and
  Energy rows (1259813, 1259812) and a healer row (1259823), and the damage
  effects (misc 0 direct, 22 periodic) are not school gated; gnome rogues,
  warriors and physical abilities get nothing. The charge is still spent on
  any costed cast.
- Berserking keeps six custom-percentage cooldowns (tags 3 to 6) because
  shipped APLs reference spell 26297 with those tags; they are not client
  rows and could be deleted with the APLs.
- Blood Fury school powers are snapshotted at gain; the generic spell power
  and spell damage follow live changes.
- The Berserking and Blood Fury changes were not requested in detail but
  follow the client rows; the previous values were vanilla leftovers.
