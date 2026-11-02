---
title: Survival Hunter in Forever
classSlug: hunter
spec: survival
role: dps
build: 'FS1:1.60.1.70707:hunter:dwarf:5/005005001/500230231051222151:'
raidBuild: 'FS1:1.60.1.70707:hunter:dwarf:5/005005201/500230231051022151:'
recommendedRaces: [dwarf, troll]
statPriority: [Agility, Attack power, Strength, Critical strike, Hit, Attack speed]
description: 'Survival Hunter overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
updated: 2026-11-01
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'Forever Sixty build data, data/builds/1.60.1.70374/talents/hunter.json'
    url: https://foreversixty.gg/data/1.60.1.70707/talents/hunter.json
    kind: datamined
  - label: 'Forever Sixty build data, data/builds/1.60.1.70374/spellconst/hunter.json'
    url: https://foreversixty.gg/data/1.60.1.70707/spellconst/hunter.json
    kind: datamined
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/combos.json'
    url: https://foreversixty.gg/data/1.60.1.69893/combos.json
    kind: datamined
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/races.json'
    url: https://foreversixty.gg/data/1.60.1.69893/races.json
    kind: datamined
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/enchants.json'
    url: https://foreversixty.gg/data/1.60.1.69893/enchants.json
    kind: datamined
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/simconsumes.json'
    url: https://foreversixty.gg/data/1.60.1.69893/simconsumes.json
    kind: datamined
  - label: 'This site, stat weights'
    url: https://foreversixty.gg/sim/weights
    kind: site
  - label: 'This site, default rotation'
    url: https://foreversixty.gg/sim/specs
    kind: site
  - label: 'This site, character planner'
    url: https://foreversixty.gg/planner
    kind: site
  - label: 'wowhead Forever, Mongoose Bite and Counterattack tooltips'
    url: https://www.wowhead.com/forever/spell=1495
    kind: community
  - label: 'wowsims-forever default hunter priority list'
    url: https://github.com/wowsims/classic/blob/master/ui/hunter/apls/p1.apl.json
    kind: community
  - label: 'Warcraft Tavern, Survival Hunter rotation guide'
    url: https://www.warcrafttavern.com/wow-classic/guides/pve-survival-hunter-rotations-cooldowns/
    kind: community
---

## Overview

Survival is a melee spec in Forever: the Hunter stands in weapon range for the whole fight and uses Raptor Strike, Mongoose Bite and bleeds instead of shots.

### Strengths

- Counterattack is a free strike that cannot be blocked, dodged or parried
- Lacerating Strikes adds a bleed to Mongoose Bite
- Raptor Strike replaces the next auto-attack without costing a swing
- Strider Kick adds a third melee special on a short cooldown

### Weaknesses

- Stays in melee range and fires no ranged shots
- Counterattack and Mongoose Bite fire only when their windows open
- Wing Clip filler pays off only with a Windfury Totem in the group

### In a raid

- **Role:** Melee DPS, played without a pet (Lone Wolf)
- **Brings:** Hunter's Mark on the target; no raid-wide buff

## Talents and builds

Reading the Survival tree in priority order for a melee build: **Savage Strikes** and **Predator's Edge** are direct melee damage — crit chance and crit damage on every melee special, plus a bonus to off-hand damage if this spec carries a second weapon. **Deterrence** unlocks **Counterattack**, a free, unavoidable strike after a parry (cannot be blocked, dodged, or parried) — a genuine "free damage" proc this build takes specifically to use. **Expose Prey** is a damage talent in its own right: each melee or ranged attack against a target carrying your Hunter's Mark has a 5% (10% at 2/2) chance to activate Mongoose Bite for 10 seconds, and that window is the only way a boss that never attacks you lets you cast it at all — the two points measure about +7% in our level-60 search run, so keep Hunter's Mark up. It also unlocks the capstone, **Lacerating Strikes**, which adds a bleed to Mongoose Bite worth 40% of that hit's damage over 21 seconds. **Strider Kick** is a third melee special on its own short cooldown. **Improved Wing Clip** and **Entrapment** are left at 0 here: they raise the reliability of control tools this rotation doesn't use against a stationary target, and the points are worth more spent reaching Counterattack and Lacerating Strikes sooner. The remaining points go into Marksmanship. **Careful Aim** (5/5) turns Intellect into melee and ranged attack power, and it opens **Lone Wolf**: 20% more damage with every attack while no pet is out. Forever's pets take only a tenth of your attack power and get no raid buffs, so for a melee Survival the pet is worth less than Lone Wolf, and this build plays without one. **Mortal Shots** is left out: its crit damage bonus covers ranged abilities only, and Survival fights with melee strikes. Its points go to Lone Wolf, a second point each in **Clever Traps** and **Resourcefulness**, and **Survivalist's Discipline**; dropping the pet for Lone Wolf is the bigger part of that respec's gain in our level-60 search run. **Lethal Attacks** (5/5) is a flat 5% critical strike chance on all attacks. It also opens the tree's second row, so the first row's other two talents are not needed to get there: **Hawk Eye** adds range to ranged weapons and **Improved Concussive Shot** a stun to a shot this rotation never fires, and neither does anything for a melee hunter. This build moves all five of those points to **Deadly Aspects** (5/5) in Beast Mastery. With Aspect of the Beast up, that talent's haste proc comes from your melee attacks and speeds them up; Hawk Eye's range is worth nothing to a hunter who never shoots, and its three points there measure about +4% on their own. **Efficiency** cuts the mana cost of melee abilities too, but **Resourcefulness** already cuts it by 60% and keeps half your mana regeneration running through the fight, so a Survival Hunter does not run dry and Efficiency gains nothing. The raid build goes further. With Judgement of Wisdom on the target, a raid Survival Hunter has more mana than it can spend, so the raid build drops Resourcefulness for both points of **Rapid Killing** (two minutes off Rapid Fire's cooldown). That measures a little under +1% over the leveling build in our raid run. Without Judgement of Wisdom it runs out of mana, so it is the raid build only. Open the planner at [/planner?class=hunter](/planner?class=hunter) to build this out.

## Rotation and priority

This site's simulator starts Survival in melee range and keeps it there for the whole fight. Before the pull, Aspect of the Beast goes up (trained at 30, with new ranks at 40, 50 and 60): in Forever it is the melee aspect, raising melee attack power by 110 at rank 4. That is twice the melee attack power Aspect of the Falcon gives at 60, and a melee hunter has no use for Falcon's ranged half. Hunter's Mark stays on the target for Expose Prey; the raid preset already carries it. Rapid Fire goes out with the other cooldowns, because in Forever it speeds up melee swings as well as shots. Raptor Strike is queued on the swing timer (it replaces the next melee auto-attack rather than costing its own swing) with no mana gate at all; an earlier version held it back under 20% mana to protect Mongoose Bite, and the search measured that gate as a small loss. Counterattack is a bare cast that only lands while its parry-triggered window is open — the same "free line, falls through otherwise" shape this site's own Warrior Overpower line uses — so it is checked early, right after Raptor Strike's queue. Mongoose Bite is a counterattack: Forever's tooltip says it can only be performed after you dodge, and Expose Prey opens the same 5-second window, so the line is gated on that window (one cast spends it) and fires the moment it opens. Lacerate, the trainable bleed, is recast whenever its own bleed has fallen off, above Strider Kick — with Mongoose Bite held to its window, the global it used to take is free, and the search measured this line at about +5%. Strider Kick comes next on its own 8-second cooldown, ahead of Immolation Trap; it costs a slice of base mana per cast, and the search placed it above the trap by a few percent once that cost was in the model. Beyond that there is one more filler, and only in a raid: with a Windfury Totem in the group, spend any global with nothing else ready on Wing Clip (80 mana for 50 damage), because the cheap main-hand hit rolls the totem's extra attack and your weapon procs. It measured about +9% in the raid-ready run and a loss without the totem, so this site's simulator only casts it when a totem reaches the hunter; without one, a global with nothing else ready goes to the swing timer, the same shape live Classic Survival play has always had. Ranged shots (Aimed Shot, Multi-Shot, Arcane Shot) and Serpent Sting are not part of this rotation at all: they require standing outside melee range, which this build never does.

## Stat priority

## Best in slot

Prioritize weapon slots: a main-hand (and off-hand, if this build carries one — Predator's Edge's off-hand damage bonus only matters with one equipped) over a ranged weapon, which this rotation no longer uses for anything but a stat stick.

## Enchants and consumables

## Races

## Professions

Community convention for Hunter favors **Skinning** for the cheap, plentiful leather this class already collects while killing beasts to level, often paired with **Leatherworking** for self-crafted agility gear, or **Blacksmithing** for a melee-focused Hunter's own weapon and armor work.

## Leveling

Beast Mastery is generally regarded as the strongest leveling spec for Hunter; a melee-committed Survival needs deeper Survival investment (Deterrence for Counterattack, Expose Prey for Lacerating Strikes) than a leveling character can usually afford early, so it comes online later than Beast Mastery does, and later than it did as a part-ranged spec.
