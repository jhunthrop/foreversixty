---
title: Holy Paladin in Forever
classSlug: paladin
spec: holy
role: healer
build: 'FS1:1.60.1.69893:paladin:human:255321003025101001/5/0:'
recommendedRaces: [human, undead]
statPriority: [Healing power, Spell power, Spirit, MP5, Intellect, Critical strike]
description: 'Talents, rotation, stats, gear, races, and professions for Holy Paladin healing in Forever.'
updated: 2026-09-24
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

Verified against this build's own Paladin talent data:

1. **Reverence** — lets a percentage of mana regeneration continue while casting, up to 30% at rank 3, which is the foundation of Holy's sustain since Paladins have no periodic heal to lean on between casts.
2. **Illumination** — a critical heal has a chance, up to 100% at rank 5, to refund mana equal to half the spell's base cost, turning crit rate into a second mana stat.
3. **Holy Shock** — an instant, short-cooldown heal or damage bolt, the tree's main tool for topping someone off without eating a global on a slow cast.
4. **Infusion of Light** — a Holy Shock or Flash of Light crit shortens the cast time of the next Holy Light, letting a big heal come out faster right after a proc.
5. **Divine Favor** — an activated cooldown that guarantees a critical effect on the next Flash of Light, Holy Light, or Holy Shock, useful for a spike heal or to bank a guaranteed Illumination proc.
6. **Light's Vigil** — the tree's capstone talent; marking an ally with it drops the cooldown off your very next Holy Shock and turns that cast into a party-wide heal, a strong cooldown-neutral burst tool.

A rough point split at level 60 would put close to 31 points in Holy to reach Light's Vigil, with the remaining points going toward Protection's early survivability talents like Toughness rather than Retribution, which has little to offer a pure healer build. That split is a projection — the beta cap of 30 has not let anyone test it. Open the planner at [/planner?class=paladin](/planner?class=paladin) to build this out.

## Rotation and priority

This site's own rotation data for Holy Paladin is an unwritten stub as of this writing, so the following is written from general 1.12 Holy Paladin practice plus the baseline changes Forever confirmed, not from a simulated priority list. Holy healing in Classic-style content is reactive rather than a fixed rotation: watch raid and tank health, cast Flash of Light for a fast partial heal or Holy Light for a slower full heal depending on how much damage needs covering, and use Holy Shock on cooldown when nobody needs it as free damage or, more often, to catch a spike before a slower cast would land. Light's Vigil is worth tagging a tank with early in a pull so its free Holy Shock is ready when damage picks up. In quiet moments, Judgement is free mana-wise and does not cost you a Seal anymore, so there is no reason not to fire it off cooldown for a small mana return and threat reduction. Divine Favor is best saved for a heal you know needs to land, since its guaranteed crit also feeds Illumination.

## Stat priority

Ordered by this site's own simulator-derived weights:

1. **Healing power** — the direct multiplier on every heal this spec casts.
2. **Spell power** — contributes to the same healing coefficient as healing power on Forever's itemization.
3. **Spirit** — feeds mana regeneration directly, and Reverence extends some of that regen into the casting window itself.
4. **MP5** — flat mana-per-five regeneration, useful once fights get long enough that burst healing power stops being the bottleneck.
5. **Intellect** — more mana pool and a small amount of spell crit, which also feeds Illumination.
6. **Crit** — increases Illumination procs and Infusion of Light windows, but ranks below the stats that raise raw throughput.

## Gear

Look for pieces that lead with healing power, then spell power, spirit, and mp5, in that order, matching the stat priority above. Beyond that principle, this section cannot get specific yet: the beta caps at level 30, nothing raids until the first tier opens on 9 December, and this build's own item data has most raid loot tables only partially re-itemized or not itemized at all for Forever, so a real Holy pre-raid or first-raid gear list would be guessing. This section fills in once raid loot data lands.

## Enchants and consumables

Target healing power and spirit on weapon and bracer enchants, matching the stat priority above. This build's enchant data confirms an Enchant Weapon - Healing Power option and healing-power enchants for bracers and gloves specifically, so a Holy Paladin has verified enchant slots to chase for its top stat. For consumables, look for elixirs or flasks that boost spirit or intellect over general-purpose ones; this build's consumable data lists spirit- and intellect-specific elixirs, though which one is strongest for Forever's mana curve isn't something this site can state with confidence yet.

## Races

Holy Paladin can only be Human or Dwarf on Alliance, or Undead on Horde — this build's race and class combination table confirms no other race can train Paladin at all. Between the two Alliance options, Human is the better healing pick: its reworked The Human Spirit racial grants a flat 5% Spirit, which feeds mana regeneration directly and compounds with Reverence's in-combat regen, while Dwarf's reworked Stoneform is a physical damage-reduction cooldown that a backline healer rarely gets to use. On Horde, Undead is currently the only playable Paladin race; its reworked Cannibalize now restores mana as well as health, giving a Holy Undead a way to top off mana between pulls that neither Alliance race has, and Will of the Forsaken's fear, charm, and sleep cleanse is useful raid utility even though it no longer grants brief immunity the way it did in 1.12.

## Professions

Alchemy pairs naturally with Holy's reliance on mana consumables, since it lets a healer supply their own mana potions and elixirs rather than buying them out. Enchanting or Tailoring both give access to caster-stat enchants and crafted cloth or leather pieces that carry healing power directly. This is general Classic-era community convention rather than anything Forever-specific, since profession bonuses have not been shown to change for Forever.

## Leveling

Retribution is the stronger leveling spec for Paladin — it kills things faster solo, where Holy's strengths barely apply outside a group. If leveling Holy anyway, expect a slower solo pace than Retribution offers. All of this is restated 1.12 knowledge rather than tested Forever guidance: the beta caps at level 30, well short of where a full leveling route would matter.
