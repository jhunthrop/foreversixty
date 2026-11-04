# Guide accuracy round 1: Warlock

Auditor: warlock lane. Tree: origin/main 64168181f, client 1.60.1.70755, engine pin 2eaae03c6.
Guides audited: `web/src/content/guides/warlock/index.md`, `affliction.md`, `demonology.md`, `destruction.md`, plus the rendered rotation notes in `data/curated/apl/warlock-*.json`.

## What 70755 changed for warlock

Diffing `data/builds/1.60.1.70753` against `1.60.1.70755`: trainables are identical. Talents differ in two places only:
- **Wrack** (1316697): 36 a tick became 86 a tick at 60 (the engine notes in `sim/warlock/wrack.go:23-27` also say the coefficient doubled from 0.143 to 0.286).
- **Shadow and Flame** (426316): 2/4/6/8/10% became 4/8/12/16/20%.

Unstable Affliction (427717, 1242971, 1242972) and Haunt (403501, 1293693, 1293694) are still listed only as `source: "class_spell"` in `trainables/warlock.json`. They have no trainer row and no talent row, so they are not learnable. The index's line 31 is correct.

## Builds decoded (talents/warlock.json tab order)

Every build totals 51 points. Every tier gate (5 points a tier) and every prerequisite holds: Wrack needs Siphon Life, Soul Link needs Demonic Sacrifice, Pact needs Soul Link, Fire and Brimstone needs Conflagrate, Incinerate needs Bane of Havoc.

**Affliction leveling** `25522300130201351/200522/005`
- Affliction 35: Improved Life Tap 2, Suppression 5, Improved Corruption 5, Grim Reach 2, Soul Harvest 2, Improved Drains 3, Amplify Curse 1, Pandemic 3, Nightfall 2, Siphon Life 1, Soul Siphon 3, Shadow Mastery 5, Wrack 1
- Demonology 11: Improved Health Funnel 2, Unholy Power 5, Demonic Aegis 2, Improved Voidwalker 2
- Destruction 5: Bane 5

**Affliction raid** `2552001013500105/00052/0550004`
- Affliction 30: Improved Life Tap 2, Suppression 5, Improved Corruption 5, Grim Reach 2, Improved Bane of Agony 1, Amplify Curse 1, Pandemic 3, Malevolence 5, Siphon Life 1, Shadow Mastery 5
- Demonology 7: Unholy Power 5, Demonic Aegis 2
- Destruction 14: Improved Shadow Bolt 5, Bane 5, Ruin 4

**Demonology leveling** `25522/233211310122000135/005`
- Affliction 16: Improved Life Tap 2, Suppression 5, Improved Corruption 5, Grim Reach 2, Soul Harvest 2
- Demonology 30: Improved Health Funnel 2, Improved Imp 3, Demonic Embrace 3, Unholy Power 2, Demonic Aegis 1, Improved Voidwalker 1, Fel Vitality 3, Demonic Energies 1, Demonic Sacrifice 1, Master Summoner 2, Decimation 2, Soul Link 1, Demonic Knowledge 3, Master Demonologist 5
- Destruction 5: Bane 5

**Demonology raid** `05/2035113011220001351/0550005`
- Affliction 5: Suppression 5
- Demonology 31: Improved Health Funnel 2, Demonic Embrace 3, Unholy Power 5, Demonic Aegis 1, Improved Voidwalker 1, Fel Vitality 3, Improved Sayaad 1, Demonic Sacrifice 1, Master Summoner 2, Decimation 2, Soul Link 1, Demonic Knowledge 3, Master Demonologist 5, Demonic Pact 1
- Destruction 15: Improved Shadow Bolt 5, Bane 5, Ruin 5

**Destruction leveling** `25522/0/2053235103101351`
- Affliction 16: Improved Life Tap 2, Suppression 5, Improved Corruption 5, Grim Reach 2, Soul Harvest 2
- Destruction 35: Destructive Reach 2, Bane 5, Molten Skin 3, Cataclysm 2, Aftermath 3, Ruin 5, Shadowburn 1, Agonizing Flames 3, Conflagrate 1, Bane of Havoc 1, Fire and Brimstone 3, Shadow and Flame 5, Incinerate 1

**Destruction raid** `2550000002/0/2053345103101351`
- Affliction 14: Improved Life Tap 2, Suppression 5, Improved Corruption 5, Pandemic 2
- Destruction 37: as leveling, plus Cataclysm 3 and Aftermath 4

## Claims checked and found correct (no finding)

I checked each of these against hard evidence:
- **Curse of the Elements** rank 4 (1311680): -75 resistance and +10% damage taken, both on misc 126 (Holy, Fire, Nature, Frost, Shadow, Arcane), lasts 300 s.
- **Curse of Recklessness** 11717: -505 armor.
- **Amplify Curse** 18288: no GCD, 180 s cooldown, 30 s duration.
- **Life Tap**: the base amount plus Spirit, times Improved Life Tap (`sim/warlock/lifetap.go`, client text "$SPI*1").
- **Drain Soul**: a shard on damage, not only on a kill. The Wowhead Forever tooltip for 11675 reads "Non-trivial targets damaged by Drain Soul have a chance to grant you a Soul Shard".
- **Portal of Summoning** 437169: level 60, 1 h cooldown, five shards (round 12 A3).
- **Shadow Trance**: 10 s.
- **Soul Fire**: 48 to learn, 6 s cast, 60 s category cooldown. With Bane and Decimation it is a 2.4 s cast on a 6 s cooldown.
- **Incinerate**: 2.5 s base, 2.0 s with Bane, learned at 40, +25% on an Immolated target.
- **Conflagrate** 10 s and **Shadowburn** 15 s. Shadowburn takes a Soul Shard reagent (tooltip).
- **Bane of Doom**: 1742 after 60 s, 60 s cooldown, learned at 60.
- **Spellstone** is a weapon enchant in Forever: shadow damage plus cast speed (`sim/warlock/spellstone_test.go`). "Brilliant Wizard Oil first, Spellstone fallback" is consistent with that.
- **Eureka!** on warlocks: the direct-damage mask has no Incinerate, and the periodic mask covers only Drain Life, Drain Soul and Wrack (`sim/core/eureka_auto_gen.go` 1259821). "Your DoTs get nothing" and "does not touch Incinerate" both hold.
- **Racials**: Blood Fury (10% spell power, 15 s, 2 min), Berserking (10%, 3 min), Expansive Mind (5% mana), Command (+5% pet damage), Ruthless, The Human Spirit and Touch of the Grave all match `races.json`.
- **Warlock races**: Human, Orc, Undead, Gnome, and Troll as the new pair. No Skyborne (`combos.json` class 9).
- **Subjugate Demon** (30/44/58). **Summon Incubus** at 20, Succubus at 20, Felhunter at 30, Voidwalker at 10. Siphon Life rank 1 at 30. Wrack at 40.
- **The sacrifices swapped schools**: 18789 Burning Shadow is aura 79 misc 32 (Shadow, from the Imp). 18791 Touch of Fire is misc 4 (Fire, from the Succubus).
- **Every talent number** the prose quotes matches the talent text at the decoded rank. That covers Pandemic +100%, Malevolence 5%, Soul Siphon 12/36%, Decimation 90%/40%/6%, Master Summoner 4 s/40%, Demonic Knowledge 100% of level, Master Demonologist 10%, Soul Link 30%/3%, Aftermath 10% a point, Agonizing Flames 10%, Fire and Brimstone 25%, and Shadow and Flame 20% on 70755.
- **Elixir of Fire Power** is right as the Fire elixir. The item is in `simconsumes.json`. Only the *recipe* for Greater Firepower exists in `items.json`, not the elixir.
- **"Enchanting makes some of the best pre-raid caster pieces"**: Frozen Heart of the Mountain and Torch of Light are Enchanting crafts at band 60 raid in `bis/warlock-affliction.md:173,176`.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix | also on r14 |
|---|---|---|---|---|---|---|---|
| WARLOCK-1 | CONFIRMED | consumable/race | affliction.md:8, :100 | `recommendedRaces: [gnome, troll]`; "Gnome and Troll are the picks. ... Berserking ... Orc Blood Fury is ..." | Owner ruling in memory forever-refresh-hold.md: "Orc = Horde pick for Affliction/Destruction" (round 14, O5). | `recommendedRaces: [gnome, orc]`. Races paragraph: "Gnome and Orc are the picks. Eureka! gives 8 seconds of 10% more damage and 10% lower cost on your direct damage spells every two minutes. Your DoTs get nothing from it, so press it as you start a run of Shadow Bolts. Expansive Mind's 5% more mana is real for a spec that taps. Orc's Blood Fury is 10% spell power for 15 seconds every two minutes, and Command makes your demon hit 5% harder all fight: press Blood Fury with your trinkets once the DoTs are up. Troll's Berserking is 10% cast speed on a three minute cooldown, a step behind Orc. ..." (keep the Undead and Human sentences) | fixed (7014d2c48) |
| WARLOCK-2 | CONFIRMED | consumable/race | destruction.md:8, :112 | `recommendedRaces: [gnome, troll]`; "Gnome and Troll are the picks." | Same ruling. | `recommendedRaces: [gnome, orc]`. "Gnome and Orc are the picks. ... Orc's Blood Fury is 10% spell power for 15 seconds every two minutes, and Command makes your demon hit 5% harder: press Blood Fury at the pull with your trinkets. Troll's Berserking is 10% cast speed on a three minute cooldown, a step behind Orc." | fixed (7014d2c48) |
| WARLOCK-3 | CONFIRMED | talent | affliction.md:62 | "It also moves the point from Improved Drains to Improved Bane of Agony" | The leveling build decodes to Improved Drains **3**, not 1. The raid build drops all 3, plus Improved Health Funnel 2 and Improved Voidwalker 2, and adds Improved Bane of Agony 1. The sentence compares against an old raid build the reader never sees. | "It also drops the three points of Improved Drains, since a raid rotation casts no drain, along with Improved Health Funnel and Improved Voidwalker, and takes a point of Improved Bane of Agony for 5% more Agony damage." | fixed (9733f7eb9 / r14 text) |
| WARLOCK-4 | CONFIRMED | rotation | demonology.md:87 | "the burn is Shadow Bolt, Soul Fire, Shadow Bolt, Soul Fire" | Soul Fire under Decimation has a 6 s cooldown (category 60 s, -90%) and a 2.4 s cast. Shadow Bolt is a 2.5 s cast with Bane. SF + SB = 4.9 s, which is under 6 s, so one Shadow Bolt never covers the cooldown. Even with Berserking (4.45 s) it takes two Shadow Bolts between Soul Fires, whether the cooldown starts at cast start or at cast end. The APL (`warlock-demonology.json` priority 1) fires Soul Fire whenever it is ready under Decimation, which gives SF, SB, SB. | "Decimation also cuts Soul Fire's cooldown to 6 seconds all the time, so the burn is Soul Fire every time it comes off cooldown and two Shadow Bolts in between: each Soul Fire is a 2.4 second cast that costs no shard and beats Shadow Bolt, and each Shadow Bolt keeps Decimation up." | not fixed |
| WARLOCK-5 | CONFIRMED | talent | destruction.md:65 | "Bane of Havoc, Molten Skin and Destructive Reach are path points for this build." | Bane of Havoc is a real path point: Incinerate requires it. Molten Skin 3 and Destructive Reach 2 are not. Without them the tree still clears every gate: tier 1 Bane 5; tier 2 13; tier 3 19; tier 4 23; tier 5 27; Incinerate after 32. That frees 5 points for Aftermath 5/5, Pandemic 3/3 and Malevolence (Affliction then holds 15, so tier 3 opens). Talent text: Molten Skin is "Reduces all damage taken", Destructive Reach is range only. | "Bane of Havoc is the path to Incinerate. Molten Skin and Destructive Reach are survival and range picks, not damage: on a fight where you never take a hit, the fifth point of Aftermath and more Pandemic do more." (Build change: see Sim findings S4.) | not fixed |
| WARLOCK-6 | CONFIRMED | stale-70755 | data/curated/apl/warlock-destruction.json:225, :244 (rendered rotation notes) | "Each hit gives you 20 seconds of 10% more Shadow damage" / "... 10% more Fire damage" | `talents/warlock.json` on 70755: Shadow and Flame rank 5 reads "increases all Shadow damage you deal by 20% for 20 sec ... all Fire damage you deal by 20%". The guide's own line 61 already says 20%. | Conflagrate note: "Conflagrate on cooldown while Immolate is up. Each hit gives you 20 seconds of 20% more Shadow damage, and with 5 points in Shadow and Flame it never eats your Immolate." Shadowburn note: "Shadowburn on cooldown, at any boss health. Each hit gives you 20 seconds of 20% more Fire damage, and with 5 points in Shadow and Flame the Soul Shard comes straight back." | not fixed. r14 also REWRITES destruction.md:36 and :61 back to 10% (its 70753 basis). That is a regression on 70755 and must not land as is. |
| WARLOCK-7 | CONFIRMED | stats | affliction.md:9 | `statPriority: [Spell power, Shadow power, Intellect, Hit, Critical strike, Spell haste, Spell penetration]` | `bis/warlock-affliction.json`, band 60 raid order: spell power, Shadow power, hit, (haste, inside its error), crit, intellect. Intellect is last of the real stats, and the guide's own prose calls it "a small mana gain". The rendered pills follow the data order (`[spec].astro:211`), so the frontmatter only misleads readers of the source and any other consumer. | `statPriority: [Spell power, Shadow power, Hit, Critical strike, Intellect, Spell haste, Spell penetration]` | not fixed |
| WARLOCK-8 | CONFIRMED | stats | demonology.md:9-10 | `[Spell power, Hit, Intellect, Critical strike, Shadow power, ...]` | `bis/warlock-demonology.json`, band 60 raid: spell power, Shadow power, hit, crit, intellect, Fire power (haste on top is noise, error 2.22 against 13.88). The prose (line 101) puts Shadow power second as well. | `statPriority: [Spell power, Shadow power, Hit, Critical strike, Intellect, Fire power, Spell haste, Spell penetration]` | not fixed (r14 changed only the prose) |
| WARLOCK-9 | CONFIRMED | stats | destruction.md:9-10 | `[Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, Shadow power, Fire power]` | `bis/warlock-destruction.json`, band 60 raid: spell power, hit, Fire power, crit, Shadow power, intellect. The frontmatter puts Fire power last and Intellect second, which contradicts both the data and the prose at line 100 ("Fire power is the better school stat"). | `statPriority: [Spell power, Hit, Fire power, Critical strike, Shadow power, Intellect, Spell haste, Spell penetration]` | not fixed |
| WARLOCK-10 | CONFIRMED | voice | affliction.md:96, demonology.md:109, destruction.md:108 | "The table above is what the raid numbers assume." | Sim meta talk (memory feedback-guide-voice). | "The table above is the standard raid kit." | fixed (fcb37cbc8) |
| WARLOCK-11 | LIKELY | talent/leveling | demonology.md:6, :53 | Leveling build `233211...` keeps Improved Imp 3, and the prose says it "keeps ... Improved Imp" | The same guide levels with the Voidwalker or Succubus (line 121). Improved Imp only buffs the Imp's Firebolt and Fire Shield, so 3 points do nothing while Unholy Power sits at 2/5. | Build `25522/203511310122000135/005` (Unholy Power 5, Improved Imp 0). Prose: "The leveling build stops at 30 Demonology points, before Demonic Pact, keeps Improved Corruption, and puts all 5 points in Unholy Power rather than Improved Imp, since you level with the Voidwalker or the Succubus." | fixed (c40362b6e) |
| WARLOCK-12 | LIKELY | talent | demonology.md:67 | "Lash of Pain crits at about a normal caster's rate, more with Malevolence." | The Demonology raid build takes no Malevolence. The talent reads "your Shadow spells", and the engine does not pass warlock crit to demons (`sim/warlock/pet.go:263-266`, "Crit is not inherited"). | "Lash of Pain crits at about a normal caster's rate." | fixed (9733f7eb9) |
| WARLOCK-13 | NEEDS-SIM | rotation | affliction.md:72 (and the APL note, warlock-affliction.json Bane of Agony line) | "Bane of Agony beats Bane of Doom for Affliction. ... Amplify Curse adds 50% to Agony and nothing to Doom ... which is why they run Doom." | Agony replaced Doom in the 2026-10-31 audit together with Siphon Life (+4% combined), when the build still had **Wrack**, whose 10% bonus covers only Corruption and Agony. Wrack has since left the rotation (round 2) and the raid build. Per-minute base damage from the client: Doom 1742, coefficient 4.0, 1 GCD, 300 mana. Agony 552 per 24 s, 12 × 0.133, so 1380 + 4.0 SP per minute over 2.5 GCDs and 537 mana. Improved Bane of Agony 1 (+5%) and one Amplify every 3 minutes (about +6.7% on average) do not obviously close a 360-damage and 1.5-GCD gap. Amplify reaches one Agony cast in seven, so it cannot be "the reason". | Question for the sim round: Affliction raid build as published, 3 min Patchwerk, band 60 raid. Agony maintained vs Doom (with Agony in the last minute, as Demonology and Destruction do it). Either result needs prose without the Amplify reasoning. If Agony holds: "Bane of Agony beats Bane of Doom for Affliction: its ticks start at once, Improved Bane of Agony and Amplify Curse both raise it, and Pandemic doubles its crits." If Doom wins, copy Demonology's Doom paragraph. | not fixed (r14 only rewrote the APL note to the Amplify reason) |
| WARLOCK-14 | NEEDS-SIM | stale-70755 | affliction.md:60, :74, :51 | "Cheap and strong solo, weaker than Shadow Bolt in a raid." / "Wrack stays off your raid bars. ... Shadow Bolt does more in those 6 seconds." | Round 2 measured Wrack about 11% below Shadow Bolt at 36 a tick and coefficient 0.143. The 70755 push made it 86 a tick and 0.286. Rough 6 s comparison at ~500 SP: Wrack 516 + 1.72 SP, plus its 10% on Corruption and Agony, plus Soul Siphon's 36% if taken. Shadow Bolt is 2.4 casts of 268 + 0.857 SP. With Soul Siphon, Wrack may now win, which would turn the raid build back toward Nightfall, Soul Siphon and Wrack. dps-per-point and bis are carried from 70753 (`dps-per-point.carried-from`). | Re-run the Affliction raid talent search and the Wrack-vs-Shadow-Bolt filler on 70755. Hold the line as is until measured. If Wrack wins, the raid paragraphs at :51, :60, :62 and :74 need rewriting. | n/a (r14 is on 70753) |
| WARLOCK-15 | LIKELY | talent | affliction.md:7 | Raid build spends Grim Reach 2 and Improved Bane of Agony 1/2 | Grim Reach is range only. Swapping one Grim Reach point into Improved Bane of Agony 2/2 keeps Shadow Mastery's gate at exactly 25 points (ILT2 Sup5 IC5 GR1 IBoA2 AC1 Pand3 Mal5 SL1) and is strictly more Agony damage (+5%) for a point of range. | `raidBuild: 'FS1:1.60.1.70755:warlock:gnome:2551002013500105/00052/0550004:'` (Grim Reach 1, Improved Bane of Agony 2), and the line at :62 becomes "takes both points of Improved Bane of Agony for 10% more Agony damage". Subject to WARLOCK-13 and WARLOCK-14. | not fixed |
| WARLOCK-16 | LIKELY | leveling | destruction.md:120 (and :62 "your filler from level 40") | "From 40, Incinerate replaces Shadow Bolt." | Incinerate is a tier 6 talent, so it needs 30 Destruction points (level 39) plus the rank at 40. That only works if every point goes into Destruction first. The leveling build shows 16 Affliction points, and nothing says to buy them last. The leveling bands do it that way: `bis/warlock-destruction.json` band 40 is all 31 points in Destruction. | Add to the Leveling paragraph: "Put every point into Destruction until Incinerate at 40, and only then go back for Suppression and Improved Corruption." | not fixed |
| WARLOCK-17 | LIKELY | voice | affliction.md:92, demonology.md:105, destruction.md:104 | "the full list and its alternatives are one click away" | Site UI talk, the same three times. | "The list below shows each piece and where it drops." | not fixed |
| WARLOCK-18 | LIKELY | voice | all three spec guides: Soul Shards (aff :78, demo :93, destro :90), Professions (aff :104, demo :117, destro :116), Enchants (:96/:109/:108), Best in slot (:92/:105/:104) | Word-for-word identical paragraphs across the three specs | Reads as template copy, not three players. Destruction's Fire elixir line differs, but the rest is pasted. | Keep the facts and give each spec one line of its own. Demonology: "every wipe is another Succubus summon, so carry twice the shards". Destruction: "Shadowburn needs one in the bag even though it comes back". Affliction: "a few Drain Soul ticks on trash refill the bag". Drop the repeated sentences. | not fixed |
| WARLOCK-19 | LIKELY | voice | demonology.md:67 | "its melee gets 8/17 of your hit. Demons almost never crit in melee against a boss." | The facts are sourced (Blizzard, `pet.go:256-261`, 234-253), but a fraction like 8/17 reads like engine notes, not a player. | "Its spells hit as often as yours, and its melee gets a bit under half your hit. Its melee almost never crits on a boss; Lash of Pain crits about as often as you do." | not fixed |
| WARLOCK-20 | LIKELY | talent | demonology.md:65 | "Improved Health Funnel, Demonic Aegis and Improved Voidwalker add no damage." | The raid build also carries Demonic Embrace 3 (stamina, no damage) and Improved Sayaad 1, which adds 10% to Lash of Pain (talent text). Neither is mentioned. | "Improved Health Funnel, Demonic Embrace, Demonic Aegis and Improved Voidwalker add no damage; they are path points that keep you and your demon alive. The point in Improved Sayaad is 10% more Lash of Pain." | not fixed |
| WARLOCK-21 | LIKELY | talent | index.md:26 | "New talents. Malevolence ..., Pandemic, Wrack ..., Decimation ..., Demonic Pact ..., Shadow and Flame, Incinerate and Bane of Havoc." | The 70755 tree has more new talents the guides use: Demonic Knowledge, Demonic Aegis, Demonic Energies, Soul Harvest, Improved Drains, Molten Skin, Fire and Brimstone. Conflagrate is reworked (no longer eats Immolate with Shadow and Flame). Not wrong, just short. | Append: "Smaller new picks fill the trees: Demonic Knowledge (spell damage equal to your level with a demon out), Soul Harvest, Improved Drains, Demonic Aegis, Demonic Energies, Molten Skin and Fire and Brimstone." | not fixed |
| WARLOCK-22 | LIKELY | voice | destruction.md:43 | "Incinerate and Conflagrate are expensive, so you Life Tap often" | Incinerate is 325 mana per 2.0 s (162/s), which is close to Shadow Bolt's 380 per 2.5 s (152/s). The heavy spends are Shadowburn (365 a GCD) and Conflagrate (255 a GCD) on cooldown. | "Conflagrate and Shadowburn on cooldown burn mana fast, so you Life Tap often" | not fixed |

## Sim findings (engine and APL vs the real game; for the next sim round)

- **S1 (APL, rotation): Curse of the Elements order.** All three APLs cast Corruption before Curse of the Elements (affliction priority 2 before 3, demonology 3 before 4, destruction 0 before 3). For the raid's Curse of the Elements holder, real play is CoE first. It raises every Fire, Frost, Arcane, Nature, Shadow and Holy hit in the raid from the first global, and a personal sim cannot see that. The main guides (affliction.md:68 "it goes on first", demonology.md:81 "Curse of the Elements first") are right. r14 rewrote both lines to "right after Corruption" to match the APL, which is a regression in real play. Fix the APLs (CoE before Corruption when the holder casts it) and keep the main prose.
- **S2 (talent search, 70755): Wrack.** The Wrack buff (36 to 86 a tick, coefficient 0.143 to 0.286) has not been measured. Re-run the Affliction raid talent search and a Wrack filler line against Shadow Bolt (WARLOCK-14).
- **S3 (rotation): Agony vs Doom for Affliction** on the current raid build, which has no Wrack (WARLOCK-13).
- **S4 (build): Destruction raid build.** Molten Skin 3 and Destructive Reach 2 are optional (WARLOCK-5). Measure variants such as `25500000033/0/0050355103101351` (Aftermath 5, Pandemic 3, Malevolence 3, no Molten Skin or Destructive Reach; 18 + 33 = 51, every gate checked) against the published build under the adoption rule. Keep Molten Skin if the owner wants the survival pick.
- **S5 (build): Affliction raid** Grim Reach 1 + Improved Bane of Agony 2 against Grim Reach 2 + Improved Bane of Agony 1 (WARLOCK-15). Strictly better on paper, so only the size of the gain is open.
- **S6 (data): carried weights.** `bis/` and `dps-per-point.json` are carried from 70753. Destruction's weights and set DPS were computed with Shadow and Flame at 10%, and Affliction's leveling bands with the old Wrack. The refresh must re-measure both. Prose is number-free, so nothing in the guides moves.

## Owner questions

1. **Demonology Horde race.** The Orc ruling names Affliction and Destruction. Demonology still recommends Troll and calls Orc "the strong Horde alternative" (demonology.md:8, :113). Command's +5% pet damage counts for more with Soul Link and Unholy Power 5. Should Demonology also be Gnome and Orc?
2. **Curse of the Elements order** for the holder: raid-first (main prose) or personal-DPS order (APL, and r14's prose)? See S1.
3. **r14 integration.** r14's warlock guide diff is based on 70753. It sets Shadow and Flame back to 10% (destruction.md:36, :61) and stamps every build `FS1:1.60.1.70753`. On 70755 both are wrong. The integrator must keep the 70755 stamps and the 20% while taking r14's race, Improved Drains, Improved Imp, Malevolence and voice fixes.
4. **Destruction Molten Skin.** Molten Skin is 6% less damage taken. Is it a deliberate survival pick for the raid build, or should the build go all damage (S4)?

## Fix log

Branch gr1/warlock, commit 0a4dd65d8 (off the new origin/main with round 14 on 70755). Re-read on that tree: every build is stamped `FS1:1.60.1.70755`, and destruction.md:36 and :61 already read 20% for Shadow and Flame, so the r14 regressions warned about in WARLOCK-6 and owner question 3 did not land. No new r14 regression found.

| id | result |
|---|---|
| WARLOCK-1 | fixed by r14 (affliction.md:8, :100 already Gnome and Orc) |
| WARLOCK-2 | fixed by r14 (destruction.md:8, :112) |
| WARLOCK-3 | fixed by r14 (affliction.md:62) |
| WARLOCK-4 | fixed (demonology.md:87): Soul Fire on cooldown, two Shadow Bolts between |
| WARLOCK-5 | fixed (destruction.md:65): Havoc is the path to Incinerate; Molten Skin (6% less damage taken) and Destructive Reach (20% range) are survival and range picks, not damage. Build string unchanged, see S4 |
| WARLOCK-6 | fixed (data/curated/apl/warlock-destruction.json:225, :244 notes, 10% to 20%); guide lines 36 and 61 already 20% on main |
| WARLOCK-7 | fixed (affliction.md:9) to the band 60 raid order |
| WARLOCK-8 | fixed (demonology.md:9-10) |
| WARLOCK-9 | fixed (destruction.md:9-10) |
| WARLOCK-10 | fixed by r14 (all three: "the standard raid kit") |
| WARLOCK-11 | fixed by r14 (build and prose demonology.md:6, :53) |
| WARLOCK-12 | fixed by r14 (demonology.md:67) |
| WARLOCK-13 | softened (affliction.md:72): dropped the unproven "Amplify Curse is why the others run Doom" causal claim; kept the published choice of Agony. Agony vs Doom stays a SIM item (S3). The Affliction APL note (warlock-affliction.json:106) still gives the Amplify reason; left, this lane's note edits were limited to the Destruction file |
| WARLOCK-14 | left (NEEDS-SIM, S2): the Wrack lines are not contradicted by measured evidence on 70755 yet |
| WARLOCK-15 | left (build change, S5); prose already says "one point of Improved Bane of Agony", which matches the build |
| WARLOCK-16 | fixed (destruction.md:120): verified Incinerate is tier 6 (30 Destruction points), so every point goes to Destruction before 40 |
| WARLOCK-17 | fixed (affliction.md:92, demonology.md:105, destruction.md:104) |
| WARLOCK-18 | fixed (Soul Shards demonology.md:93, destruction.md:90; Professions affliction.md:104, demonology.md:117, destruction.md:116): each spec now has its own line; the Enchants paragraphs already differ in their elixir line and stay |
| WARLOCK-19 | fixed (demonology.md:67) |
| WARLOCK-20 | fixed (demonology.md:65): verified Improved Sayaad rank 1 is 10% Lash of Pain in talents/warlock.json |
| WARLOCK-21 | fixed (index.md:26) |
| WARLOCK-22 | fixed (destruction.md:43): verified spellconst costs Incinerate r3 325, Shadow Bolt r10 380, Conflagrate r6 255, Shadowburn r6 365 |

Confirmed 10 (no r14 regressions added). Fixed in this lane: 4, 5, 6, 7, 8, 9 confirmed, plus LIKELY 16, 17, 18, 19, 20, 21, 22 and the softened NEEDS-SIM 13. Fixed by r14: 1, 2, 3, 10, 11, 12.

### SIM findings and build proposals (no build strings changed)

- S1: Curse of the Elements before Corruption for the curse holder in all three APLs. r14 prose says "right after Corruption" (affliction.md:68, demonology.md:81, destruction.md:71) to match the APL; left pending owner question 2.
- S2: Wrack on 70755 (86 a tick, coefficient 0.286): re-run the Affliction raid talent search and Wrack vs Shadow Bolt filler (WARLOCK-14).
- S3: Agony vs Doom on the Affliction raid build (WARLOCK-13). If Doom wins, rewrite affliction.md:72 and the APL note.
- S4: Destruction raid build: measure `25500000033/0/0050355103101351` (Aftermath 5, Pandemic 3, Malevolence 3, no Molten Skin or Destructive Reach) against the published `2550000002/0/2053345103101351` (WARLOCK-5, owner question 4).
- S5: Affliction raid build: `2551002013500105/00052/0550004` (Grim Reach 1, Improved Bane of Agony 2) against the published build (WARLOCK-15).
- S6: bis weights and dps-per-point carried from 70753; re-measure Destruction (Shadow and Flame 20%) and Affliction leveling (Wrack). statPriority frontmatter was set from the carried band 60 raid weights and should be rechecked after the refresh.

Owner question 1 (Demonology Horde race) still open; demonology.md:8 and :113 unchanged.

## Fix pass 2

MUSTs:
1. Trees/icons not rendering: left (capture artifact, integrator ruling).
2. "The raid build moves N points from ..." miscount (Affliction 10 of 15, Demonology 11 of 12): left, R2. The line is generated by web/src/lib/guides/build-diff.ts:72 via GuideBuilds.astro, not guide prose. SITE finding: the diff summary drops talents (decoded: Affliction moves 15 = Nightfall 2, Soul Siphon 3, Wrack 1, Soul Harvest 2, Improved Drains 3, Improved Health Funnel 2, Improved Voidwalker 2 into Malevolence 5, Improved Shadow Bolt 5, Ruin 4, Improved Bane of Agony 1; Demonology moves 12 = Improved Corruption 5, Improved Life Tap 2, Grim Reach 2, Soul Harvest 2, Demonic Energies 1 into Demonic Pact 1, Improved Sayaad 1, Improved Shadow Bolt 5, Ruin 5). Prose checked against the decoded builds: affliction.md:62 and destruction.md:53 already correct; demonology.md:63 now names what the raid build gives up.
3. Engineering goggles vs Tailoring/Enchanting prose: fixed. Item facts from items/warlock.json (15999 Spellpower Goggles Xtreme Plus, 27 spell power, BoE; 14111 Felcloth Hood, 30 Shadow power, 10 Int) and bis sources (Engineering / Tailoring crafted; Felcloth Hood is the Demonology runner-up, Dreamweave Circlet the Destruction one). demonology.md:105, :117; destruction.md:104, :116. No engineering skill requirement exists in items data, so the prose says "an engineer's craft", not "requires Engineering".
4. Demonology Troll vs Orc: left, owner question 1 still open.

SHOULDs:
- "Until Demonic Pact at 60": fixed, demonology.md:121 (Pact is tier 7, the 31st point, open from 40; the leveling build skips it).
- Apology lines: fixed per R4, demonology.md:65 (toll to reach Pact; what each buys at the ranks taken: 9% Stamina, 15% Demon Armor, 10% Voidwalker, Health Funnel), destruction.md:65 (safety and reach).
- Duplicate race lines: left. "Orc is the Horde pick: ..." comes from recommendedRacesLine in web/src/lib/guides/copy.ts, so only one of the two is guide prose (ruling: drop only if both are prose). SITE finding: the generated line repeats the Races prose opener.
- R1 site talk also removed: "cap shown above" / "the table above shows" (affliction.md:88, demonology.md:101, destruction.md:100), "The table above is the standard raid kit" (affliction.md:96, demonology.md:109, destruction.md:108), "The list below shows each piece" (affliction.md:92, demonology.md:105, destruction.md:104).

New SIM findings: none this pass (no build or action changes raised beyond S1 to S6).
