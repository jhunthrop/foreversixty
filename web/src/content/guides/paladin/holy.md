---
title: Holy Paladin in Forever
classSlug: paladin
spec: holy
role: healer
build: 'FS1:1.60.1.70707:paladin:human:50320215012101251/50322/0503:'
recommendedRaces: [human, undead]
statPriority: [Healing power, Intellect, Spirit, MP5, Critical strike]
description: 'Talents, rotation, stats, gear, races, and professions for Holy Paladin healing in Forever.'
updated: 2026-10-31
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/talents/paladin.json'
    url: https://foreversixty.gg/data/1.60.1.69893/talents/paladin.json
    kind: datamined
  - label: 'Forever Sixty stat weights, data/curated/specs.json'
    url: https://foreversixty.gg/data/curated/specs.json
    kind: site
  - label: 'Forever Sixty raid phase data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/data/curated/loot/forever-raid-phases.json
    kind: site
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/races.json'
    url: https://foreversixty.gg/data/1.60.1.69893/races.json
    kind: datamined
---

## Overview

Holy is the Paladin healing tree: reactive single-target healing built on Flash of Light, Holy Shock and Divine Favor, with mana sustained by Reverence and Illumination.

### Strengths

- Holy Shock and Divine Favor give instant and guaranteed-critical burst heals
- Flash of Light is the most efficient heal per mana for tank and party
- Illumination and Reverence sustain mana while casting
- Spiritual Focus protects heals from damage pushback

### Weaknesses

- Mana runs short in a long fight, so Flash of Light carries most of the healing
- Healing is single-target, with no group heal to match Prayer of Healing
- Heals are reactive, without the periodic healing Priest and Druid lean on

### In a raid

- **Role:** Tank and raid healer, reactive single-target healing
- **Brings:** Greater Blessings and an aura for the group

## Talents and builds

The build is 31 points in Holy, 12 in Protection and 8 in Retribution. Chosen by simming candidate builds on the same gear against the same fight. Every figure behind this ranking comes from one stated incoming-damage profile (the scaled Phase 1 tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

- **Spiritual Focus** opens the tree and protects your casts from damage pushback; **Reverence** keeps part of your regeneration running while you cast.
- **Healing Light** strengthens Holy Light, Flash of Light and Holy Shock.
- **Illumination** returns mana on a critical heal, **Infusion of Light** makes the next Holy Light faster after a Flash of Light or Holy Shock crit, and **Divine Favor** guarantees a critical heal on demand.
- **Holy Shock** is the instant heal, and **Boundless Light** makes it crit far more often on anyone you have not sworn **Oath of Light** to.
- **Holy Power** now makes your heals' critical strikes heal for more, rather than crit more often; **Light's Vigil** is the capstone.
- **Divine Intellect** moved to Protection's first row and still lifts your mana pool.
- In Retribution, **Benediction** makes Holy Shock and your other instants cheaper, and **Conviction** now adds critical strike to every spell, heals included, which feeds Illumination and Holy Power. Taking those eight points instead of twenty in Protection measured about +5% healing in our simulator.

Oath of Light is taken, but the client does not yet say how much it adds, so the simulator counts it as nothing. Open the planner at [/planner?class=paladin](/planner?class=paladin) to adjust it.

## Rotation and priority

The curated rotation opens with **Greater Blessing of Light** before the pull and then leans on **Flash of Light** as the staple: it is the most efficient heal per mana for both the tank and the party. **Divine Favor** into **Holy Light** answers a tank who is nearly dead. **Holy Shock** and a second Holy Light line only fire while your mana is keeping up with the share of the fight that is left; a healer who is short on mana, which is the usual case, heals as a Flash of Light healer.

**Light's Vigil** is in the engine but left out of the rotation: against the profile's pulses it cost more mana than it returned. The rotation does not downrank.

## Stat priority

In simulator-derived order, per point of stat: **MP5** first, then **Intellect**, with **healing power** and **Spirit** close together behind and **critical strike** after them. Intellect feeds both your pool and Illumination's mana return. Items trade several points of one for few of another, so read the list as "which stat is cheap to give up".

Spell power adds nothing to a heal. Every figure behind this ranking comes from one stated incoming-damage profile (the scaled Phase 1 tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

## Best in slot

## Enchants and consumables

## Races

## Professions

Alchemy pairs naturally with Holy's reliance on mana consumables, since it lets a healer supply their own mana potions and elixirs rather than buying them out. Enchanting or Tailoring both give access to caster-stat enchants and crafted cloth or leather pieces that carry healing power directly.

## Leveling

Retribution is the stronger leveling spec for Paladin — it kills things faster solo, where Holy's strengths barely apply outside a group. If leveling Holy anyway, expect a slower solo pace than Retribution offers.
