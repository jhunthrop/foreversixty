---
title: Retribution Paladin in Forever
classSlug: paladin
spec: retribution
role: dps
build: 'FS1:1.60.1.70707:paladin:human:050250002/502/0523330100130135:'
raidBuild: 'FS1:1.60.1.70707:paladin:human:050050002/503/05233301001301351:'
recommendedRaces: [human, undead]
statPriority: [Strength, Attack power, Critical strike, Hit, Agility, Melee haste]
description: 'Talents, rotation, stats, gear, races, and professions for Retribution Paladin melee damage in Forever.'
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
  - label: 'wowsims-forever default Retribution priority list'
    url: https://github.com/wowsims/classic/blob/master/ui/retribution_paladin/apls/basic_ret.apl.json
    kind: community
  - label: 'Warcraft Tavern, Classic Retribution Paladin rotation'
    url: https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/
    kind: community
  - label: 'Forever Sixty raid phase data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/data/curated/loot/forever-raid-phases.json
    kind: site
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/races.json'
    url: https://foreversixty.gg/data/1.60.1.69893/races.json
    kind: datamined
---

## Overview

Retribution is the Paladin melee damage tree: keep a Seal up, Judge on cooldown, and add Hammer of Wrath on targets low enough to execute.

### Strengths

- Strongest of the three Paladin specs for solo play and leveling
- Judgement is off the global cooldown and drives the damage loop
- Hammer of Wrath adds an execute below roughly 20% health
- Sanctified Judgement sustains more Judgements over a fight

### Weaknesses

- Exorcism works only against Undead and Demons
- Consecration costs too much mana on a single target without raid mana support
- Every Seal recast costs a global and over 200 mana, so swapping Seals only pays with Unsealed Wrath and raid mana

### In a raid

- **Role:** Melee damage dealer
- **Brings:** Its own damage, plus Blessing of Might and Blessing of Wisdom

## Talents and builds

The leveling build puts 30 points in Retribution and the raid build 31. **Benediction**, **Improved Judgement**, **Conviction** (now 2% critical strike a rank with every spell and attack), **Sanctified Judgement**, **Sacred Arbiter**, **Two-Handed Weapon Specialization**, **Vengeance** and **Champion of the Light** carry the damage. **Seal of Command** is the point Vengeance requires; the leveling rotation never casts it, and the raid build swaps into it for Unsealed Wrath (below). Champion of the Light grew to five ranks and also shortens Hammer of Wrath's cast; Two-Handed Weapon Specialization now raises physical damage only, so your Seals, Judgements and Holy Strike no longer gain from it. **Vindication** no longer raises your own attack power; it stays as a gate point, along with one point of **Repentance**, because the tree needs 25 points before Champion of the Light and Retribution has no other damage talent to fill them.

The rest goes to Holy and Protection. **Divine Strength** and **Improved Seals** (five ranks now) add damage, **Holy Conduit** (now a Holy talent) makes Consecration, Exorcism and Hammer of Wrath cheaper, and **Reverence** keeps some mana regeneration running while you cast. **Divine Intellect** now sits on Protection's first row, which makes it the gate to **Precision**: Intellect feeds Champion of the Light's spell damage and your mana, and Precision is the hit the raid search has favoured since October. The leveling build spends 2 points in Precision and 2 in Reverence; the raid build takes Precision 3/3 and spends its last point on Unsealed Wrath instead of Reverence. The leveling build measures about +3% over a straight carry of its old talents in our level-60 search run, and the raid build was the best the raid search found.

**Unsealed Wrath**, the new capstone, stores your active Seal on each swing, up to three copies, and releases them when a swing lands under a different Seal: each stored copy gets that Seal's normal chance to fire, and Seal of Righteousness fires on every swing. The raid build takes it for its one point of Reverence. Let three swings store Seal of Righteousness, switch to Seal of Command for one swing, which releases three extra Righteousness hits, then switch straight back. That measured about +1.3% in the raid-ready run. An even three-and-three swap gains half as much, and without raid mana every swap is a loss, so the leveling build skips the talent. Open the planner at [/planner?class=paladin](/planner?class=paladin) to build this out.

## Rotation and priority

Apply Seal of Righteousness about a second and a half before the pull so it is active the moment combat starts. Once in combat, Judgement comes first whenever it is off cooldown: it is off the global cooldown, costs 6% of your base mana, unloads the active Seal's Judgement effect and immediately starts its own cooldown again. Exorcism follows on its own 15-second cooldown against Undead and Demons only; it is free Holy damage on those encounters, simply unavailable against anything else. Below roughly 20% target health, Hammer of Wrath takes priority as the execute option, cast at its highest rank available at level 60. Next, reapply Seal of Righteousness just before it expires so uptime never lapses. Putting the seal refresh ahead of Holy Strike was worth about +3.5% in our level-60 search run. Holy Strike, a Forever baseline Paladin attack, comes last, on its own 10-second cooldown, cast whenever it is up: it is what spends the globals that Judgement and seal upkeep alone leave idle.

Without Unsealed Wrath this rotation does not seal-twist. Older community guides for 1.12 Retribution describe alternating Seal of Righteousness and Seal of Command swing-by-swing for a DPS gain. In Forever that twist loses damage on its own: casting a new Seal fully replaces the one that was up, so no swing benefits from both, and each recast costs a full global and over 200 mana at level 60. In our simulator, dropping the swing-by-swing twist and spending those globals on Holy Strike instead was worth about +23%. Keep Seal of Righteousness up and spend every other global on Judgement, Exorcism, Hammer of Wrath and Holy Strike in that priority order. With Unsealed Wrath, the swap described under Talents sits just ahead of the Seal refresh, and it waits whenever you are below 30% mana.

Consecration, part of Forever's baseline kit, is absent from the leveling and bare rotation above and joins it last in the raid-ready one: against a single target it is a loss on a bare character, where its mana cost (scaling steeply with rank) crowds out Judgement, Exorcism and Holy Strike, but with the raid's mana support the rotation search measured it as a clear gain, about +9% in the raid-ready run, and it holds the bare character level when it is only cast while mana is above 65%. While leveling, or fighting several enemies at once, it is still worth reaching for (see Leveling, below).

## Stat priority

## Best in slot

## Enchants and consumables

## Races

## Professions

Blacksmithing suits Retribution well since it can add sockets to weapons and armor, letting a melee Paladin fill in strength or crit the itemization doesn't provide on its own. Enchanting offers weapon and ring enchants a Retribution Paladin can apply without paying another crafter.

## Leveling

Retribution is the strong leveling choice for Paladin — it kills things fastest solo of the three specs, which is why it is worth considering even for a character planning to heal or tank later.

Consecration (baseline from level 20, no talent needed) is worth reaching for while leveling even though this page's single-target rotation leaves it out: questing and farming routinely put two or three mobs on you at once, and Consecration's ground effect hits all of them for the one cast and mana cost, which is a better trade than it is against a single raid target.
