---
title: Retribution Paladin in Forever
classSlug: paladin
spec: retribution
role: dps
build: 'FS1:1.60.1.69893:paladin:human:0/55325/55223331211000021:'
recommendedRaces: [human, undead]
statPriority: [Attack power, Strength, Agility, Critical strike, Hit, Melee haste]
description: 'Talents, rotation, stats, gear, races, and professions for Retribution Paladin melee damage in Forever.'
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

Retribution is Forever's Paladin melee damage tree, built around keeping a Seal active, Judging it off cooldown, and weaving in Hammer of Wrath once a target is low enough to execute. In a group it is the spec that turns the Paladin's kit into raw damage rather than healing or mitigation, and it is also the strongest of the three specs for solo play, since neither Holy nor Protection is built to kill things quickly alone. The beta caps at level 30, so nothing here about level 60 raid damage, real boss encounters, or best-in-slot gear has actually been played; it is a projection from the demo talent trees, this site's own rotation modeling, and 1.12 Retribution Paladin knowledge, confirmed only where Blizzard's own recap says so directly.

## Talents and builds

Verified against this build's own Paladin talent data:

1. **Seal of Command** — the tree's signature damage Seal, adding a chance for extra Holy damage on weapon swings and judging for a burst of Holy damage on demand.
2. **Sanctified Judgement** — gives Judgement a chance to refund part of the judged Seal's mana cost, up to a guaranteed 60% refund at rank 3, which keeps a melee-focused build from running dry on mana.
3. **Vindication** — melee hits have a chance to reduce the target's attack power while raising your own, a self-buff on top of a minor debuff.
4. **Sacred Arbiter** — increases Holy Strike's damage and makes it refresh all active Judgement effects on the target, tying the baseline attack directly into Retribution's damage.
5. **Instrument of Law** — shortens Hammer of Wrath's cast time, which matters specifically for landing it cleanly inside a shrinking execute window.
6. **Twist of Light** — the tree's capstone talent; its tooltip describes swapping off a Seal granting an echo that applies the old Seal's effect on the next melee hit, which on paper is what would make seal twisting viable without losing a swing to bad timing. This site's own rotation does not seal twist (see Rotation and priority, below) — this build's own engine has no code path that reads this talent at all, so it currently does nothing in a sim regardless of how it plays live. Until that lands, this last point is better spent in Holy on the Holy Shock talent instead.

This build puts 31 points in Retribution to reach Twist of Light, or would stop one short of it for a player taking Holy Shock instead, with the remaining 20 in Protection for baseline survivability — Toughness, Redoubt, and Anticipation all maxed, then Precision and Guardian's Favor for the rest of the budget — since Retribution has few defensive tools of its own. That split is a projection — the beta cap of 30 has not let anyone test it. Open the planner at [/planner?class=paladin](/planner?class=paladin) to build this out.

## Rotation and priority

This site's own rotation model for Retribution Paladin has an implemented priority list, unlike Holy and Protection's rotation data, which are still unwritten stubs. Apply Seal of Righteousness about a second and a half before the pull so it is active the moment combat starts. Once in combat, Judgement comes first whenever it is off cooldown — it is a free global that unloads the active Seal's Judgement effect and immediately starts its own cooldown again. Exorcism follows on its own 15-second cooldown against Undead and Demons only — free Holy damage on those encounters, simply unavailable against anything else. Below roughly 20% target health, Hammer of Wrath takes priority as the execute option, cast at its highest rank available at level 60. Holy Strike — Forever's new baseline Paladin attack, added to the client this week — comes next, on its own 10-second cooldown, cast whenever it is up: it is what spends the GCDs that Judgement and seal upkeep alone leave idle. Reapply Seal of Righteousness just before it expires so uptime never lapses; that line also catches the seal being consumed by Judgement's own cast, which fires this same reapply on the very next decision.

This rotation does not seal-twist. Older community guides for 1.12 Retribution (and an earlier version of this page) describe alternating Seal of Righteousness and Seal of Command swing-by-swing for a DPS gain. Measured directly against Forever's own engine, that twist is a DPS **loss** here: casting a new Seal always fully replaces whatever Seal was active a moment before (there is no overlap window for a swing to benefit from both, unlike the timing trick 1.12's version of the technique exploited), and Judgement always ends by reapplying Seal of Righteousness regardless of which Seal was up when it fired, so twisting for Judgement's benefit does not work either. Each Seal recast still costs a full global and 200+ mana at level 60, so twisting on every swing mostly just burns mana for nothing. A forever-sim comparison at level 60 in band-60 gear found the old twist-every-swing rotation at 153.8 DPS, the same rotation with the twist removed at 165.6 DPS, and this page's current rotation (twist removed, Holy Strike added) at 190.0 DPS — a 23.5% gain over the twisting version. Just keep Seal of Righteousness up and spend every other global on Judgement, Exorcism, Hammer of Wrath and Holy Strike in that priority order.

Consecration, also new in Forever's baseline kit this week, is deliberately absent from this priority list. It is a strong tool while leveling or fighting several enemies at once (see Leveling, below) but a DPS loss against a single target: its mana cost scales steeply with rank (135 at rank 1 up to 565 at rank 5), and spending a global and that much mana on a ground effect that does nothing extra to one target crowds out Judgement, Exorcism and Holy Strike. A forever-sim comparison confirmed this directly — adding Consecration on cooldown to the level-60 rotation above measured 184.3 DPS, lower than the 190.0 without it.

## Stat priority

Ordered by this site's own simulator-derived weights:

1. **Attack power** — the primary driver of Retribution's melee and Seal proc damage.
2. **Strength** — converts directly into attack power, Retribution's top stat.
3. **Agility** — adds crit chance and a small amount of attack power and armor.
4. **Crit** — increases the value of every Seal proc and melee swing, and feeds Reckoning if any points are spent in Protection.
5. **Hit** — keeps Judgement and Holy Strike landing reliably, since a missed Judgement is a missed damage window.
6. **Melee haste** — more swings per minute, ranked below the stats above it once accuracy and raw power are covered.

## Gear

Look for pieces that lead with attack power and strength, matching the stat priority above. Beyond that principle, this section cannot get specific yet: the beta caps at level 30, nothing raids until the first tier opens on 9 December, and this build's own item data has most raid loot tables only partially re-itemized or not itemized at all for Forever, so a real Retribution pre-raid or first-raid gear list would be guessing. This section fills in once raid loot data lands.

## Enchants and consumables

Target strength on weapon, bracer, and glove enchants, and consider the Crusader weapon enchant for its proc, matching the stat priority above; this build's enchant data confirms both Enchant Weapon - Strength and Enchant Weapon - Crusader exist. For consumables, look for elixirs or flasks that boost strength or agility over general-purpose ones; this build's consumable data lists an Elixir of Greater Strength specifically, though which combination is strongest for Forever's itemization isn't something this site can state with confidence yet.

## Races

Retribution Paladin can only be Human or Dwarf on Alliance, or Undead on Horde — this build's race and class combination table confirms no other race can train Paladin at all. Between the two Alliance options, Human is the better damage pick if wielding a two-handed sword: its reworked Sword Specialization grants a flat crit chance bonus with swords, which applies to every Judgement and melee swing. Dwarf's Mace Specialization offers the same kind of bonus but only for maces, so the choice mostly follows weapon type rather than one race being flatly stronger. On Horde, Undead is currently the only playable Paladin race; Will of the Forsaken's fear, charm, and sleep cleanse is a useful tool for a melee spec that otherwise has few ways to shrug off crowd control mid-fight, even though it no longer grants brief immunity the way it did in 1.12.

## Professions

Blacksmithing suits Retribution well since it can add sockets to weapons and armor, letting a melee Paladin fill in strength or crit the itemization doesn't provide on its own. Enchanting offers weapon and ring enchants a Retribution Paladin can apply without paying another crafter. This is general Classic-era community convention rather than anything Forever-specific, since profession bonuses have not been shown to change for Forever.

## Leveling

Retribution is the strong leveling choice for Paladin — it kills things fastest solo of the three specs, which is why it is worth considering even for a character planning to heal or tank later. This is restated 1.12 knowledge rather than tested Forever guidance: the beta caps at level 30, well short of where a full leveling route would matter.

Consecration (baseline from level 20, no talent needed) is worth reaching for while leveling even though this page's single-target rotation leaves it out: questing and farming routinely put two or three mobs on you at once, and Consecration's ground effect hits all of them for the one cast and mana cost, which is a better trade than it is against a single raid target.
