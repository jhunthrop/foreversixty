---
title: Retribution Paladin in Forever
classSlug: paladin
spec: retribution
role: dps
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
6. **Twist of Light** — the tree's capstone talent; swapping off a Seal grants an echo that applies the old Seal's effect on the next melee hit, which is what makes seal twisting possible without a separate swing-timer addon. It demands tight execution, so a Retribution build that would rather not track seal swaps can spend that last point in Holy on the Holy Shock talent instead.

A rough point split at level 60 would put close to 31 points in Retribution to reach Twist of Light, or stop one short of it for a player skipping the twist, with the remaining points typically going into Protection for baseline survivability talents like Toughness and Anticipation, since Retribution has few defensive tools of its own. That split is a projection — the beta cap of 30 has not let anyone test it. Open the planner at [/planner?class=paladin](/planner?class=paladin) to build this out.

## Rotation and priority

This site's own rotation model for Retribution Paladin has an implemented priority list, unlike Holy and Protection's rotation data, which are still unwritten stubs. Apply a Seal about a second and a half before the pull so it is active the moment combat starts. Once in combat, Judgement comes first whenever it is off cooldown — it is a free global that unloads the active Seal's Judgement effect and immediately starts its own cooldown again. Below roughly 20% target health, Hammer of Wrath takes priority as the execute option, cast at its highest rank available at level 60. Reapply the active Seal just before it expires so uptime never lapses. For a player using the full seal-twisting build, the idea is to always have the more valuable Seal's on-hit effect ready to land on the very next swing: whichever Seal of the pair is not currently up gets cast in the instant just before the swing timer fires, so the model alternates between Seal of Righteousness and Seal of Command from swing to swing rather than holding one for the whole fight. Early in a pull, before either Seal has had a chance to land once, the model simply keeps Seal of Command active rather than risk going without a Seal at all. A player skipping seal twisting simply keeps one Seal active throughout and follows the Judgement and Hammer of Wrath priority above.

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
