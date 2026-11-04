# Guide accuracy round 12: Warrior

Audited on the round 12 audit tree (main a8b147a6f: guide rounds 1 to 11, sim round 15, leveling-order, guide-header-build), client 1.60.1.70755, engine pin 9f7f1f6e4. Files: `web/src/content/guides/warrior/{index,arms,fury,protection}.md` and every card note in `data/curated/apl/warrior-{arms,fury,protection}.json`. I read the round 1 to 11 warrior reports and fix logs, the ledger, and the round 4 and round 5 levels notes.

What changed since round 11's tree (d54de48aa): no warrior guide, card, `addon-builds.json`, BiS or client data file changed. Round 11 had no warrior fixes. The only change that reaches these pages is guide-header-build (3bd24d3a2 and the merges after it):
- `[spec].astro` and `GuideActionRail.astro`: the header build card now shows `raidBuild` where a guide has one, titled "Raid build".
- `guide-rotation.ts`: the rotation cards, cooldown cuts and pre-pull timings read that same build, and a "For the raid build." caption sits above the cards.
- `KeyTalentList.astro`: key-talent chips show both builds' ranks where they differ.

For warrior, only Arms has a `raidBuild`. Fury and Protection still show their `build`.

## Builds decoded (all legal on 70755, unchanged since round 11)

- Arms `build` `13323213032215201/0505/205`: IHS 1, Deflection 3, Improved Rend 3, Improved Charge 2, ITM 3, Improved Overpower 2, AM 1, Deep Wounds 3, 2H Spec 3, Impale 2, Bloodthrill 2, Sweeping Strikes 1, Weaponmaster 5, Improved Slam 2, MS 1 / Cruelty 5, Unbridled Wrath 5 / Improved Bloodrage 2, Iron Will 5.
- Arms `raidBuild` `30303213132215201/05050005012/1` (now the header card): IHS 3, Improved Rend 3, ITM 3, Improved Overpower 2, AM 1, Deep Wounds 3, Spearing Strike 1, 2H Spec 3, Impale 2, Bloodthrill 2, Sweeping Strikes 1, Weaponmaster 5, Improved Slam 2, MS 1 / Cruelty 5, Unbridled Wrath 5, Enrage 5, Reckless Swings 1, Improved Execute 2 / Improved Bloodrage 1.
- Fury `build` `30323013002/050530025122100231/2`: IHS 3, Improved Rend 3, Improved Charge 2, ITM 3, AM 1, Deep Wounds 3, Impale 2 / Cruelty 5, Unbridled Wrath 5, Furious Precision 3, Enrage 2, DW Spec 5, Reckless Swings 1, Improved Execute 2, Improved Berserker Rage 2, Death Wish 1, Improved Whirlwind 2, Flurry 3, Bloodthirst 1 / Improved Bloodrage 2.
- Protection `build` `35300003/0/250533120330001351`: IHS 3, Deflection 5, Improved Rend 3, Deep Wounds 3 / Improved Bloodrage 2, Shield Spec 5, Anticipation 5, Improved Revenge 3, Improved Thunder Clap 3, Last Stand 1, Master of Defense 2, Defiance 3, Improved Sunder Armor 3, Concussion Blow 1, Focused Rage 3, Bastion 5, Shield Slam 1.

Shipped orders are as listed in round 11; `addon-builds.json` is unchanged.

## Findings

None. No statement on the four pages or in the warrior card notes is false on 70755. None contradicts the page it sits on, and none contradicts the new header card.

## Header-build check (round 12 focus)

- **Arms.** The header card, the rotation cards ("For the raid build.") and the key-talent chips now show the raid build. The prose agrees with all of them:
  - Every key-talent bullet that differs between the builds names both: Improved Charge "(2 points, level 60 build only)" and Improved Bloodrage "two points in the level 60 build ... one point in the raid build".
  - The two build paragraphs are labelled "The level 60 build" and "The raid build", so neither reads as the card above it.
  - The Rotation text places Reckless Swings ("in the raid build") and Spearing Strike in the same order as cards [7] and [10].
  - Both of those cards say "raid build only", which is consistent with the new caption.
  - No sentence on the page refers to "the build above" or to the card.
- **Fury and Protection** have no `raidBuild`. Their header card still shows `build`, and their prose describes that build.
- The rail's rotation lines are now the first four rows of the priority card, so they carry the same notes. That is a component change; no warrior note is affected.

## Shoulds (wording, not findings)

- Arms Bloodrage card [2] sits under the new "For the raid build." caption but still adds "half again with the level 60 build's two". This is true, not a contradiction. Optional trim: "Whenever rage is under 80: 10 rage at once and 10 more over 10 seconds, a quarter more with Improved Bloodrage, back every minute."
- Carried from round 11, still open:
  - arms.md:125 leaves out the level 19 point, the first point of Improved Tactical Mastery.
  - arms.md:91 "Mortal Strike on cooldown and Execute on every other global" leaves out Reckless Swings, which sits above Execute in the raid-build cards and works in Berserker Stance.
  - The Reckless Swings health-cost wording differs between Fury and Arms.
- fury.md:54 "No healing cut or debuff of its own for the raid": it means a spec-specific debuff, but a Fury warrior still brings Sunder Armor and Demoralizing Shout. Optional: "No healing cut of its own for the raid".

## Checked and right (re-derived this round)

- **Spell text and constants** (`spelltips/warrior.json`, `spellconst/warrior.json`, `trainables/warrior.json`):
  - Mortal Strike 4: 50% healing cut, 6000 ms category cooldown.
  - Bloodthirst 4: 45% AP, 6000 ms.
  - Sweeping Strikes: "next 5 melee attacks".
  - Whirlwind: 4 targets, both weapons, 10000 ms.
  - Death Wish: 20% physical, 5% damage taken, 30 s, 180000 ms, 10 rage.
  - Recklessness: 100% crit chance, 15 s, 1800000 ms ("once a fight").
  - Berserker Rage: 30000 ms, level 30.
  - Victory Rush: 10% heal, within 20 s, 30000 ms, level 20.
  - Shield Block: 75%, 7 s, 2 attacks, 5000 ms.
  - Shield Wall: 60%, 12 s, 900000 ms.
  - Last Stand: 30%, 20 s.
  - Concussion Blow: 5 s.
  - Challenging Shout: 6 s, 600000 ms.
  - Thunder Clap: 4 targets.
  - Revenge: 5 rage, which Focused Rage makes 2.
  - Taunt: "no effect if the target is already attacking you".
  - Charge 3: 15 rage, "Cannot be used in combat".
  - Execute: converts extra rage.
  - Tactical Mastery: 10 rage at 14.
- **Trainer levels**: Charge and Rend 4, Thunder Clap 6, Hamstring 8, Bloodrage, Defensive Stance, Sunder Armor and Taunt 10, Overpower and Shield Bash 12, Revenge and Demoralizing Shout 14, Shield Block 16, Whirlwind 36.
- **Talent text at the build's ranks**:
  - IHS rank 1 "by 1 Rage". Improved Charge rank 2 "by 6". Improved Bloodrage 25% / 50%. ITM "+15".
  - Shield Spec rank 5 "100% chance to generate 5 Rage", lower ranks a chance, so "from the fifth point" holds.
  - Improved Revenge 60%. Improved Thunder Clap 6. Improved Sunder Armor 3. Anticipation 20. Defiance 15%. Bastion 10%. Master of Defense 5 rage on dodge or parry.
  - Improved Whirlwind 2 s and 30%. Flurry 30% for 3 swings. DW Spec 25% / 50%. Furious Precision 10%.
  - Improved Shield Wall "by 11.0 min", so 4 minutes.
- **Improved Shield Wall swap** (protection.md:83): moving 2 points from Improved Thunder Clap (tier 2) to Improved Shield Wall (tier 5) leaves 25 points above tier 5, 31 above Bastion and 36 above Shield Slam, so the build stays legal. Thunder Clap then costs 4 more.
- **Racials** (`spelltips/shared.json`): Quickness 1% dodge and 2% speed, Endurance 5%, Stoneform (bleed, poison and disease immunity, 10% physical, 8 s), Elune's Light 10% crit 15 s, Blood Fury 10% AP 15 s, Ruthless 1% crit, Artisan's Precision 1% hit, Wind Blessed 2% haste.
- **Source links**, fetched this round: the Blizzard recap, Talents Forever data.json, both Warcraft Tavern Classic and Forever warrior pages, the Warcraft Tavern Fury guide, Wowhead Skyborne First Look, Power Up Gaming, and foreversixty.gg /sim/weights, /sim/drops, /planner and /bis/warrior/protection all return 200. Icy Veins returns 403 (unknown, not dead).
- No em or en dashes, no quoted nightly numbers, no stat ranking in prose.

## Decisions

- **Zero confirmed findings for warrior this round.** No warrior text changed since round 11. The one change that reaches these pages, the header card following `raidBuild`, matches the Arms prose: the prose already labels its level 60 and raid builds separately and ties the raid-only buttons to the raid build.
- **Race lines stay as they are.** They are still unmeasured on the new builds, and nothing contradicts a measured result or the client.

## Sims requested

- Still open, carried: race re-runs on the Arms `raidBuild` and the Fury `build`. Run raid 60 for every race on both factions, 10,000 iterations, seeds 1 and 50001. Lines affected: arms.md:109-111 and fury.md:112-114.
- PROVISIONAL: any prose-vs-BiS-row mismatch after the final re-rank. None found on this tree.
- Still open from earlier rounds: Dwarf against Human breakdown, Blood Fury held timing, Stoneform rotation line, and the Protection Improved Shield Wall build (Shield Wall is never cast in the tank sim).

## Sim findings

None new.

## Sources

- Client data 1.60.1.70755:
  - `talents/warrior.json`
  - `trainables/warrior.json`
  - `spelltips/warrior.json`
  - `spellconst/warrior.json` (21553, 23894, 12328, 1719, 18499, 402927, 2565, 871, 12975, 12809, 1161, 1680, 12292, 11601, 2687, 1329555, 1310222)
  - `spelltips/shared.json` racials
  - `addon-builds.json`
- Commits 3bd24d3a2, d273fa004 and a8b147a6f (guide-header-build): `web/src/pages/guides/[class]/[spec].astro`, `web/src/lib/guides/guide-rotation.ts`, `GuideActionRail.astro`, `KeyTalentList.astro`, `copy.ts` ("Raid build", "For the raid build.").
