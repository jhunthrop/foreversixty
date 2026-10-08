---
title: Holy Paladin in Forever
classSlug: paladin
spec: holy
role: healer
build: 'FS1:1.60.1.70009:paladin:human:05320003225111051/55325/0:'
recommendedRaces: [human, undead]
statPriority: [Healing power, Intellect, Spirit, MP5, Critical strike]
description: 'Talents, rotation, stats, gear, races, and professions for Holy Paladin healing in Forever.'
updated: 2026-10-07
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

Holy is Forever's Paladin healing tree: single-target direct healing backed by Holy Shock and Light's Vigil for burst, with mana sustained through Reverence and Illumination rather than heavy regeneration gear alone. In a group it fills the same seat Holy fills in 1.12 — reactive raid and tank healing rather than the periodic-heal style Priest or Druid lean on. The beta caps at level 30, so nothing here about level 60 raid healing, mana curves over a full encounter, or best-in-slot gear has actually been played; it is a projection built from the demo talent trees and 1.12 Holy Paladin knowledge, confirmed only where Blizzard's own recap says so directly.

## Talents and builds

The build is nearly all Holy, with the remaining points in Protection where they cost nothing a healer uses. Chosen by simming candidate builds on the same gear against the same fight. Every figure behind this ranking comes from one stated incoming-damage profile (the Onyxia-sized tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

- **Divine Intellect** lifts your mana pool, and **Reverence** keeps part of your regeneration running while you cast.
- **Healing Light** strengthens Holy Light, Flash of Light and Holy Shock; **Spiritual Focus** protects those casts from damage pushback.
- **Illumination** returns mana on a critical heal, **Infusion of Light** makes the next Holy Light faster after a Flash of Light or Holy Shock crit, and **Divine Favor** guarantees a critical heal on demand.
- **Holy Shock** and **Holy Power** give the instant heal and the extra critical chance; **Light's Vigil** is the capstone.

The sim found any build that keeps that core scores the same, because the other Holy talents change nothing a healer casts. Open the planner at [/planner?class=paladin](/planner?class=paladin) to adjust it.

## Rotation and priority

The curated rotation opens with **Greater Blessing of Light** before the pull and then leans on **Flash of Light** as the staple: it is the most efficient heal per mana for both the tank and the party. **Divine Favor** into **Holy Light** answers a tank who is nearly dead. **Holy Shock** and a second Holy Light line only fire while your mana is keeping up with the share of the fight that is left; a healer who is short on mana, which is the usual case, heals as a Flash of Light healer.

**Light's Vigil** is in the engine but left out of the rotation: against the profile's pulses it cost more mana than it returned. The rotation does not downrank.

## Stat priority

In simulator-derived order, per point of stat: **MP5** first, then **Intellect**, with **healing power** and **Spirit** close together behind and **critical strike** after them. Intellect feeds both your pool and Illumination's mana return. Items trade several points of one for few of another, so read the list as "which stat is cheap to give up".

Spell power adds nothing to a heal. Every figure behind this ranking comes from one stated incoming-damage profile (the Onyxia-sized tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

## Gear

Look for pieces that lead with the stats at the top of the stat priority above; spell power on a healing item adds nothing to a heal. Beyond that principle, this section cannot get specific yet: the beta caps at level 30, nothing raids until the first tier opens on 9 December, and this build's own item data has most raid loot tables only partially re-itemized or not itemized at all for Forever, so a real Holy pre-raid or first-raid gear list would be guessing. This section fills in once raid loot data lands.

## Enchants and consumables

Target healing power and spirit on weapon and bracer enchants, matching the stat priority above. This build's enchant data confirms an Enchant Weapon - Healing Power option and healing-power enchants for bracers and gloves specifically, so a Holy Paladin has verified enchant slots to chase for its top stat. For consumables, look for elixirs or flasks that boost spirit or intellect over general-purpose ones; this build's consumable data lists spirit- and intellect-specific elixirs, though which one is strongest for Forever's mana curve isn't something this site can state with confidence yet.

## Races

Holy Paladin can only be Human or Dwarf on Alliance, or Undead on Horde — this build's race and class combination table confirms no other race can train Paladin at all. Between the two Alliance options, Human is the better healing pick: its reworked The Human Spirit racial grants a flat 5% Spirit, which feeds mana regeneration directly and compounds with Reverence's in-combat regen, while Dwarf's reworked Stoneform is a physical damage-reduction cooldown that a backline healer rarely gets to use. On Horde, Undead is currently the only playable Paladin race; its reworked Cannibalize now restores mana as well as health, giving a Holy Undead a way to top off mana between pulls that neither Alliance race has, and Will of the Forsaken's fear, charm, and sleep cleanse is useful raid utility even though it no longer grants brief immunity the way it did in 1.12.

## Professions

Alchemy pairs naturally with Holy's reliance on mana consumables, since it lets a healer supply their own mana potions and elixirs rather than buying them out. Enchanting or Tailoring both give access to caster-stat enchants and crafted cloth or leather pieces that carry healing power directly. This is general Classic-era community convention rather than anything Forever-specific, since profession bonuses have not been shown to change for Forever.

## Leveling

Retribution is the stronger leveling spec for Paladin — it kills things faster solo, where Holy's strengths barely apply outside a group. If leveling Holy anyway, expect a slower solo pace than Retribution offers. All of this is restated 1.12 knowledge rather than tested Forever guidance: the beta caps at level 30, well short of where a full leveling route would matter.
