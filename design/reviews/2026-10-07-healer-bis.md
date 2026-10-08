# Healer BiS (2026-10-07)

What was built, the numbers it produces, and every assumption behind them. The rule the whole lane keeps:
**a healing sim ranks gear under a stated incoming-damage profile. It never says one healer class beats
another, and every published healer number carries the profile's name.**

## 1. The profile and why

`data/curated/heal-profile.json` (id `onyxia-sized`, "Onyxia-sized tank hits and raid pulses") is the one
statement every healer is ranked under. Five fake raid members: a tank and four party members.

| Part | Figure | Why |
| --- | --- | --- |
| Fight | 300 s, no variation | A long boss fight; mana only matters if the fight is long. No variation so "lasts the fight" means the same fight every run. |
| Tank | 9,500 health; a 1,150 hit landing about every 3.3 s; each hit rolls 25 % either side | A raid tank in Phase 1 gear. Hits are after mitigation; misses, dodges and parries are averaged into the interval. |
| Party members | 5,000 health each | Between a cloth and a plate wearer. |
| Raid pulses | 300 damage to 3 of the 4 members every 5 s | A breath, a wing buffet, a periodic volley. |

Total about 530 damage a second: **one healer's share of a boss fight, not the raid's**. It is sized so that a
raid-ready healer is throughput- and mana-bound (cannot cover it) but not hopeless. The first sizing (about 900
a second) put every set out of mana within a minute and could only rank bursts, so it was halved.

**Assumption, stated plainly:** no client table states a creature's damage. Every figure is a judgement sized to
a Phase 1 raid boss (the first tier opens 9 December: Onyxia, Barrow Deeps, Hyjal Summit). The figures live in the
file with their reasons, travel into every healer's BiS file (`heal_profile`) and render on the page under a
disclosure.

How it reaches the engine: `RaidDamageModel` on the raid (fork `proto/common.proto`), laid out by
`core.AddHealingFakeRaid` (a second, empty party so the tank, the fifth fake member, is not in the healer's
party; the healer is raid player 0, members 1-4, tank 5). The fake members get health bars, the tank takes
`tank_hit_damage` every swing, the others take pulses on `pulse_members` random members, and active absorb shields
(now a real pool in `core/shield.go`) soak damage first.

## 2. Engine (fork branch `healers`)

Every healing spell below is registered level-aware from the client's tables (trainables with a learn row plus
the live talent trees) and reads `declared, matches` in the conformance golden: base heal range at the
character's level, spell-power coefficient, cost, cast time, GCD, required level, duration.

| Spec | Spells registered | Talents modelled | Not modelled (reason) |
| --- | --- | --- | --- |
| Priest (Holy, Discipline) | Lesser Heal, Heal, Flash Heal, Greater Heal, Renew, Prayer of Healing, Power Word: Shield, Desperate Prayer; talent-gated Binding Heal, Penance (heal), Holy Nova, Prayer of Mending, Power Infusion; Inner Fire | Improved Renew, Improved Healing, Divine Fury, Spiritual Healing, Spiritual Guidance, Inspiration, Litany of Light, Holy Specialization, Twin Disciplines, Mental Agility, Mental Strength, Improved PW:S, Soul Warding, Renewed Hope, Divine Aegis, Improved Inner Fire, Inner Focus, Power Infusion | Circle of Healing and Lightwell are not in the live tree or need a summoned object; Penance's damage half; Blessed Recovery, Spirit of Redemption, Holy Reach, Twilight Focus, Martyrdom, Power in Light, Holy Precision |
| Holy Paladin | Holy Light 1-9, Flash of Light 1-6, Holy Shock heal 1-4, Light's Vigil 1-3, Blessing of Light | Healing Light, Spiritual Focus, Infusion of Light, Illumination, Divine Favor, Holy Power, Light's Vigil, Champion of the Light | Judgement/Seal of Light and Wisdom (the healer never melees), Sacred Shield (taught at level 80), Cleanse/Purify (nothing to dispel) |
| Restoration Shaman | Healing Wave 1-10, Lesser Healing Wave 1-6, Chain Heal 1-3, Riptide 1-3, Mana Tide Totem 1-3, Healing Stream/Mana Spring corrected to the client | Improved Healing Wave, Tidal Focus, Tidal Mastery, Purification, Healing Way, Mindfulness, Restorative Totems, Totemic Focus, Nature's Swiftness, Water Shield, Riptide, Mana Tide | Improved Reincarnation, Ancestral Healing, Healing Focus; Earth Shield (no learn row); Tidal Waves and the rest are rune-only |
| Restoration Druid | Healing Touch 1-10, Regrowth 1-9, Rejuvenation 1-10, Tranquility 1-4, Swiftmend, Wild Growth 1-3, Innervate | Gift of Nature, Gift of the Earthmother, Tranquil Spirit, Improved Rejuvenation, Improved Tranquility, Improved Regrowth, Naturalist, Genesis, Nature's Splendor, Nature's Swiftness, Living Spirit, Swiftmend, Wild Growth, Reflection (fixed to the client's 17/33/50 %) | Nature's Focus, Subtlety, Natural Shapeshifter, Furor (no effect on a healer); Lifebloom, Nourish, Tree of Life are SoD runes; Omen of Clarity has no published proc rate |

Core changes made for healing (all in the fork, documented in `PORTING.md`):

- **Healing stat.** `Spell.HealingPower` is the healing stat alone. The site's item data states damage and
  healing as separate stats (an old "damage and healing" item is spell power N *and* healing N; a Forever
  healing item is healing N and spell power N/3), so adding spell power counted the same bonus twice. Forever's
  one-third rule stays the global `HealingPower -> SpellDamage` dependency; Spiritual Guidance, Champion of the
  Light and Mental Quickness grant the healing stat.
- **Effective healing.** `GainHealth` returns what landed; results carry `effective_hps` and per-action
  `effective_healing` (shields count for what they absorbed). Stat-weight sweeps with a damage model weigh
  effective healing, and spirit and mp5 get the wider step intellect has (a point of either is below the noise).
- A harmful major cooldown is only used with an enemy as the current target (Smolderweb's Eye panicked at a
  friendly tank dummy).
- Conformance reads heal, periodic-heal, party-aura-heal effects and (when the ability file declares an amount)
  absorb effects.

## 3. The scoring rule

Healers are ranked by an engine runner that speaks healing (`score_heal.go`), so the ranker's tournaments,
verification runs and sweeps are unchanged:

- A run's score is **effective healing per second**, scaled by the **mana guard**: the share of the fight the
  mana lasted, squared, capped at one. A set whose mana lasts the fight scores its effective healing per second;
  one that empties at four fifths of it scores 0.64 of it; at half, a quarter. A set that empties therefore ranks
  below a similar set that lasts. (It is a smooth penalty, not a strict lexicographic order: a set with far more
  healing could still outscore a lasting one that barely heals. Within the same spec and talents the sets differ
  by single-digit percent, where the guard decides.)
- Mana lasts is the engine's own time-to-out-of-mana, projected from the mana left for an iteration that never
  ran dry.
- Weights are per point of effective healing per second, for healing power, intellect, spirit, mp5, crit and
  spell haste (specs.json; the reference stat is healing power), anchored at intellect as every caster is.
- An item's `healing` is weighed by the `healing_power` row; spell power weighs nothing.
- `set_dps` is the effective healing per second, so nothing that reads it breaks.

Band entries add `role`, `profile` and `metrics{hps, raw_hps, overheal_pct, mana_lasts_sec, hpm}`; the file adds
`heal_profile`.

## 4. Per-spec results, band 60 (scratch run, 100 weight iterations, 1,000 iterations per set)

Under `onyxia-sized`, 300 s. "Raid" is the Phase 1 raid-ready preset with the healer consumables (Flask of
Distilled Wisdom, Mageblood, Nightfin Soup, Brilliant Mana Oil, Major Mana Potion, Demonic Rune); "bare" is the
character and its class kit only. These are not a comparison of the specs: each runs its own curated rotation,
talents and gear.

| Spec | Preset | Side | HPS (effective) | Raw HPS | Overheal | Mana lasts | Healing per mana |
| --- | --- | --- | --- | --- | --- | --- | --- |
| priest-holy | bare | alliance gnome | 208 | 208 | 0.0 % | 86 s | 3.19 |
| priest-holy | raid | alliance gnome | 337 | 343 | 1.6 % | 161 s | 4.11 |
| priest-holy | raid | horde troll | 328 | 333 | 1.6 % | 152 s | 4.08 |
| priest-discipline | bare | alliance human | 321 | 330 | 2.7 % | 118 s | 6.24 |
| priest-discipline | raid | alliance human | 426 | 448 | 4.9 % | 216 s | 5.84 |
| priest-discipline | raid | horde undead | 423 | 445 | 4.9 % | 214 s | 5.80 |
| paladin-holy | bare | alliance human | 177 | 177 | 0.0 % | 89 s | 4.02 |
| paladin-holy | raid | alliance human | 309 | 309 | 0.0 % | 180 s | 4.30 |
| paladin-holy | raid | horde undead | 318 | 318 | 0.0 % | 179 s | 4.45 |
| shaman-restoration | bare | alliance dwarf | 231 | 232 | 0.2 % | 150 s | 4.09 |
| shaman-restoration | raid | alliance dwarf | 373 | 378 | 1.4 % | 405 s | 5.31 |
| shaman-restoration | raid | horde tauren | 373 | 378 | 1.4 % | 405 s | 5.31 |
| druid-restoration | bare | alliance night-elf | 317 | 332 | 4.5 % | 230 s | 4.50 |
| druid-restoration | raid | alliance night-elf | 435 | 453 | 4.1 % | 274 s | 5.22 |
| druid-restoration | raid | horde tauren | 438 | 457 | 4.1 % | 274 s | 5.32 |

Reading them: every set is mana-bound under this profile, which is why the mana stats lead the weights. At
raid-ready gear one spec's mana lasts the whole fight (the shaman's); the others run dry before it ends, by 25 % (the
druid) to 45 % (the Holy priest) of the fight, although both priest rotations now pace their mana. Overheal is low everywhere because the profile asks for
more than one healer can give. Raid-ready gear adds roughly 35-75 % effective healing over the bare character,
mostly through the mana consumables.

Weights at the raid band (per point, scale factor with intellect = 1): mp5 leads for every spec (2.2-4.8);
spirit is second for the priests and the druid, intellect for the paladin and shaman; healing power is about
0.5-1.3 and crit about 0.2-0.3. Spell haste is not significant for any (no healing gear carries it), so the
pages show it as a caption only.

## 5. Guides

The Talents, Rotation and Stat priority sections of the five healer guides were rewritten as relative claims only
(no number that renders from data), each pointing to the profile and saying the ranking compares no healer
class with another. **The guide builds were replaced:** the old ones were recollections (a Holy priest with Spell
Warding, Wand Specialization and Power in Light). `sim/cmd/heal-builds` sims legal candidate builds on the same
gear and fight; the best of each spec's candidates is the guide build. Notable findings: for the paladin and
shaman every build that keeps the core scores the same (the other points change nothing a healer casts); for the
druid, dropping Wild Growth and Improved Regrowth for other points scored higher under this rotation (a statement
about the rotation's mana pacing, not about the spells); the Holy priest scored higher with Holy deepest, the
Discipline priest with Discipline deepest.

## 6. Assumptions (all of them)

1. **The profile's figures** (section 1) are judgements, not measurements.
2. **Spirit regeneration, mp5 and the five-second rule** use the engine's formulas (the priest has its own); the
   client publishes no spirit regeneration table.
3. **Healing power is the healing stat alone** (section 2); a source granting "damage and healing" grants both.
4. **No healing-power elixir** exists in the client's item table, so the healer preset has none; Greater
   Mageblood Elixir is not in the engine's consumable list, so Mageblood stands in.
5. **Chain Heal** jumps twice at half the previous heal each, to the most injured members (the client constants
   carry neither); **Wild Growth** ticks fall 5 % each; **periodic heals do not crit** (nothing states they do).
6. **Prayer of Mending** triggers on the holder's next non-periodic heal only; the fake raid's damage has no hook
   for the "when damaged" half. **Blessing of Light** counts through each spell's coefficient (the client states
   only an amount); the alternative is a flat 400.
7. **No downrank penalty** is modelled (the client states none), and the rotations use top ranks: the site
   rewrites every ranked spell id to the highest learned rank below level 60.
8. **Healing crit chance ignores school bonus crit** (core limit), so set bonuses like Ten Storms' nature crit
   do nothing for heals; **item set bonuses** for the healer classes are untouched.
9. **The healer's target is a friend**: damage-spell conformance rows that read an enemy aura (Faerie Fire,
   Earth/Flame/Frost Shock, Judgement of the Crusader, Seal of Righteousness at some levels) show `mismatch` in
   the healer presets' golden rows; no healing row does.
10. **The Restoration Shaman's Tidal Focus hit** (+1 % per rank) now applies to Elemental builds that spend in it,
    so Elemental's published weights shift slightly. Reflection was corrected for Balance/Feral (results unchanged).
11. **Mana consumable use** is the engine's major-cooldown pass (potion, rune, Inner Focus); no healer rotation has
    its own potion lines.
12. **The priests do not last the fight** under this profile even paced, and I did not under-heal to hit a
    number; the guard penalises them within their own gear choices only.
13. **Known test debt:** the Holy Paladin ladder reports `unresolved_id` for Holy Light rank 9 and one more rank
    in conditions at low levels (informational, non-strict); the hunter conformance golden and four DPS test
    packages (mage, elemental, enhancement, dps_warrior) were stale on the base branch and are untouched.

## 7. What shipped, where

- Fork `healers`: raid damage model and effective healing, healing-stat fix, `healsim` harness, four healer
  specs with tests, conformance presets/goldens, healer APLs, sweep and cooldown fixes.
- Site `healers`: `heal-profile.json`, `request.HealProfile`, healer consumable group, `specOptions` for the five
  specs, `inproc.HealingRun`, `adapter.HealingWeights`, `score_heal.go`, `heal-builds`, curated rotations,
  ladder goldens, web healer panel with unit threading, fixture and Playwright spec, guides.
- Not committed by design: `data/builds/*/bis/*.json` (the nightly publishes them), `sim/go.mod`'s local replace.
