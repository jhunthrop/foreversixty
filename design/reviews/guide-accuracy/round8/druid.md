# Guide accuracy round 8: Druid

Auditor: druid persona. Tree: gr8/audit at 09b98d9e4 (land/guides-r7, all seven rounds applied; READ-ONLY), client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).

Specs audited: class page (`index.md`), Balance (`balance.md`), Feral cat (`feral.md`), Feral bear (`feral-bear.md`), Restoration (`restoration.md`). Before starting I read the round 1 to 7 druid reports with their fix logs, round4/druid-levels.md, the round 2 to 5 race-sims files and ledger.md. Between round 7's audit head (394562f20) and 09b98d9e4, the druid changes are the round 7 fixes only: feral.md:59 (Rend and Tear scope), feral.md:81 (Shifting Power cost with Natural Shapeshifter), feral-bear.md:116 (neutral leveling gear line) and the index.md `updated:` stamp. No commit touched addon-builds.json, talents/druid.json, trainables/druid.json, spelltips/druid.json or the druid APLs.

## Decoded builds (unchanged since round 1, all legal, 51 points)

Re-decoded against talents/druid.json in tab order, prerequisites checked (Nature's Splendor needs Nature's Majesty 2, Vengeance needs Improved Moonfire 2, Lacerate needs Savage Fury 2, Blood Frenzy needs Sharpened Claws 2, Rend and Tear needs Predatory Strikes 3, Primal Bite needs Leader of the Pack, Shifting Power needs Predatory Discipline, Nature's Swiftness needs Naturalist 5, Improved Regrowth needs Improved Rejuvenation 3, Wild Growth needs Living Spirit 3).

- Balance `5232220115101151/0/555301`, raid `5232220115501351/0/505003`.
- Feral `01/50212032123012123001/535033`, raid `05002/50230032023011125001/05503`.
- Bear `0/5523033202310212551/051`.
- Restoration `05312001/0/55530051531132`.

Shipped leveling orders (addon-builds.json, index 0 = level 10), re-decoded: identical to round 5 and 7.
- Balance: Improved Wrath 10-14, Genesis 15-16, Moonglow 17-19, ... Moonkin Form 40.
- Feral: Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17, Brutal Impact 18-19, Shredding Attacks 20-22, ... Shifting Power 40, Genesis 41, Nature's Focus 42-46, Heart of the Wild 47-49, Naturalist 50-54, Natural Shapeshifter 55-57, Subtlety 58-60.
- Bear: Ferocity 10-14, Furor 15-19, Lacerate 25, Primal Bite 40.
- Restoration: Heart of the Wild 10-14, Reflection 15-17, Nature's Focus 18-19.

Every Levels 1 to 19 paragraph and every 20 to 60 talent claim matches these orders, and the 60 respec sentence on feral.md:55 matches the decoded diff (Brutal Impact 2, Feral Charge 1, Blood Frenzy 1 into Feral Instinct +2 and Rend and Tear +2; Nature's Focus 5 and Subtlety 3 into Heart of the Wild +2, Genesis +4, Nature's Majesty +2).

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| DRUID-R8-1 | CONFIRMED | stats | balance.md:105 | "Nature damage feeds only Wrath and Insect Swarm." | Hurricane is a Nature spell: spellconst/druid.json 17402 and 1278759 (Hurricane rank 3) `school_mask: 8`, the same school as Wrath (5176..9912) and Insect Swarm (5570..24977); spelltips/druid.json 17402: "causing 132 Nature damage to enemies every 1 sec". The page's own AoE paragraph (balance.md:99) makes Hurricane your AoE, so "only" contradicts the page. The sentence was written in round 6's stat rewrite; round 7 checked the schools of the four single-target spells but not the "only". | "Nature damage feeds Wrath, Insect Swarm and Hurricane, never Starfire or Moonfire." |

Checked and right (no finding):

**Round 7 edits re-checked**
- feral.md:59 Rend and Tear: rank text is "Increases damage done by your melee abilities on Bleeding targets by 3%" to "15%"; Shred and Ferocious Bite are abilities. Right.
- feral.md:81 Shifting Power: spelltips 1322605 `cost_pct: 55.0`, "Shifting Power's cost is reduced by effects that reduce the cost of Shapeshifting"; Natural Shapeshifter 3/3 "Reduces the mana cost of all shapeshifting by 30%", 3 points in both cat builds. Right, and it now agrees with :67.
- feral-bear.md:116 gear line: no ranking by a nightly value. Right.

**Spells, ranks and levels (trainables/druid.json, spelltips/druid.json, spellconst/druid.json)**
- Wrath at 1 (ranks at 6 and 14), 2.0 s base so 1.5 s with Improved Wrath. Moonfire 4, Entangling Roots 8, Nature's Grasp 10, Faerie Fire 18, Starfire 3.5 s base (3.0 with Improved Starfire 5). Healing Touch 1, Rejuvenation 4, Regrowth 12, Revive 12, Cure Poison 14. Bear Form, Maul, Growl, Demoralizing Roar 10, Enrage 12, Bash 14, Swipe 16, Maul rank 2 at 18. Cat Form, Shred, Rip, Claw 20, Ferocious Bite 24, Cower 28, Ravage and Poised 32, Pounce and Frenzied Regeneration 36, Dire Bear Form, Innervate, Moonkin Form, Hurricane 40, Gift of the Wild 50 and 60.
- Rake 40 energy (35 with Ferocity), Claw 45 (40), Shred 60 (42 with Shredding Attacks 3), Primal Bite 20 rage (15) on 6 s. Bash 60 s cooldown, Brutal Impact 2/2 is +1 s stun and -30 s. Feral Charge bear 15 s; Feral Charge (Cat) 414924 30 s category cooldown. Shifting Power 8 s.
- Dire Bear Form: 120 attack power, 360% armor from items, 600 health (the "flat block of health"). Enrage 10 up front plus 20 over 10 s. Frenzied Regeneration 10 rage a second for 10 s at 1% health a rage. Barkskin 20% physical for 15 s on 60 s. Challenging Roar 6 s. Swipe hits 3. Demoralizing Roar lasts 30 s.
- Omen of Clarity: "not consumed by Wrath". Moonkin Form doubles Omen and "prevents the use of healing spells"; the deep dive says all non-healing spells cast in form, so Rebirth, Innervate and Remove Curse hold. Innervate 400% for 20 s. Tranquility reaches party within 20 yards on 5 min (2 min with Improved Tranquility's 60%). Wild Growth 3.0 s cast, 6 s cooldown, heals the target's party over 7 s.

**Talents (every named one, 70755 rank text)**
- Heart of the Wild 10% Intellect, 20% Bear Stamina, 10% Cat Strength at 5 (6% at 3). Naturalist 0.5 s off Healing Touch and Wild Growth plus 10% all damage. Nature's Focus 70% for Arcane and Nature spells (28% at 2), so Starfire, Healing Touch and Regrowth all qualify. Genesis 10% periodic damage and healing from spells and abilities, so Rip counts. Nature's Majesty 4% spell and melee crit. Nature's Splendor +3 s Moonfire and Rejuvenation, +6 s Regrowth, +2 s Insect Swarm. Eclipse 2 Starfires 0.5 s faster, 4 charges. Nature's Grace 10% cast speed and global cooldown for 3 s. Vengeance +100% Arcane and Nature crit bonus. Moonfury 10%. Nature's Reach 4% hit.
- Furor 5/5: 100% chance of 10 rage into Bear, Cat restores the energy you had plus 10 a second spent outside Bear Form. Feral Swiftness 4% dodge and 30% Cat speed. Feral Instinct 30% Swipe at 3. Thick Hide scales with Defense over five times your level and the form multiplies it. Savage Fury 10% on Claw, Rake, Shred, Maul and Swipe. Blood Frenzy 2/2: 5 rage on bear crits and a combo point on builder crits (50% at 1 point). Predatory Strikes 150% of level. Natural Instinct 20% crit bonus. Natural Reaction 5% dodge and 5 rage a dodge.
- Gift of Nature 10%, Improved Rejuvenation 15%, Bloom 100% Spirit on a natural expiry, Swiftmend 20% base mana on 15 s, Tranquil Spirit 10% Healing Touch and Tranquility, Living Spirit 15% Spirit, Reflection 50%, Improved Regrowth 10% a point. The Wild Growth swap on restoration.md:64 is legal (38 Restoration points before row 7, Living Spirit 3).

**Races** (races.json, combos.json): Elune's Light 10% crit, 15 s, 3 min; Quickness 1% dodge; Endurance 5% health; War Stomp 5 enemies, 2 s. Skyborne: Wind Blessed 2% haste, Read Leyline on the Alliance side only, 50% for 15 s on 2 min with the 60 min Attuned doubling; Skyborne is named only as the paid option on every page. None of the six non-Skyborne new pairs is a druid (human hunter, orc mage, dwarf shaman, undead paladin, gnome priest, troll warlock, per combos.json `new_in_forever`; class 11 has only Night Elf, Tauren and the two Skyborne). The bear race paragraph matches round4/race-sims.md.

**Items and consumables**: Thistle Tea, Mighty Rage Potion, Greater Stoneshield Potion, Major Mana Potion, Demonic Rune, Elemental Sharpening Stone, Dense Weightstone, Brilliant Mana Oil, EZ-Thro Thorium Grenade, Aquarius Belt, Idol of Shifting Tides and Wolfshead Helm (ilvl 45, level 40, blue) all exist in 70755 items.json. The deep dive (fetched 2026-11-04) still confirms potions and Thistle Tea in form, weightstones/Crusader/Windfury on the paw, the bomb exception with SAF-T and EZ-Thro, racials in form, Faerie Fire free in Cat and Bear with a 6 s cooldown, Hurricane with no cooldown and Barkskin in form.

**Card notes**: every rendered druid APL note re-read. Claw "40 to Shred's 54 at 20 and 48 at 21" and "8 to 14 more energy" match the cost math. None states anything false.

**Sources**: every frontmatter link returns 200 on 2026-11-04 (Blizzard deep dive, Blizzard panel recap, Wowhead Skyborne First Look, Talents Forever data.json, /sim/weights, /sim/drops, /bis/druid/feral-bear, /bis/druid/restoration), and each says what its label says.

No em or en dashes. No site, sim, APL, preset or band words in prose. No quoted nightly numbers. No ranking of stats or gear by a nightly value.

## Shoulds (wording, not findings)

- feral.md:73 "so do not pool energy for it": the reason does not lead to the advice. With Predatory Discipline the spare energy adds damage and stays in the bar, which is a reason a high-energy Bite costs nothing, not a reason against holding energy. The advice itself (do not delay the finisher) is fine. Suggested: "With Predatory Discipline the Bite adds damage from your spare energy without spending it, so Bite the moment you reach 5 points; never wait on energy for it."
- balance.md:91 "it always lands on a Starfire or a DoT": a Faerie Fire you cast in Moonkin Form (:97 tells you when) also costs mana and can take the proc. "Almost always" would be exact.

## Decisions

1. **"Only Wrath and Insect Swarm" is a finding, not a wording choice (R8-1):** Hurricane is Nature in the client tables, and the same page makes it your AoE.
2. **Deep dive lines the client has since changed stay as the client has them:** the deep dive calls Lacerate a baseline level 42 ability and says Vengeance and Moonfury cover all damaging spells; 70755 makes Lacerate a talent learned from 25 and limits Vengeance and Moonfury to Arcane and Nature. The pages follow the client, which is right. Not filed.
3. **Swiftmend keeps the HoT and Gift of Nature misses Tranquility:** the round 1, 2 and 5 rulings stand; no page states either falsely.

## Sims requested

No new ones. Carried and still open:
1. From round 4 (S10): druid-feral leveling, the switch from Claw to Shred. Run ladder levels 20 to 25 with the Shred line gated on Shredding Attacks 3/3 (level 22), against the current Ferocious Bite rank 5 gate.
2. From round 3: druid-restoration with Improved Moonfire 1 moved into Improved Regrowth 1 (`05302001/0/555300515311321`). Heal profile, band 60 raid, seeds 1 and 50001, with a ladder-first row that keeps Improved Regrowth.

## Sim findings

- S1 (carried, open at d1d9f2775): `rendAndTearDamagePerRank = 0.02` (sim/druid/feral_talents.go) against the client's 3% a rank.
- S7, S8, S9, S11, S12 and SIM-D1 (carried): Nature's Reach melee hit for the feral raid re-run, Feral Charge (Cat) unmodeled, the Pounce behind-target check, expertise in the bear weights, bear Defense noise, and the feral leveling Nature's Focus and Subtlety points.
- No new engine finding.

## Sources

- Client 1.60.1.70755: talents/druid.json (all ranks and prerequisites), trainables/druid.json, spelltips/druid.json (Hurricane 17402, Cat/Bear/Dire Bear Form, Enrage, Frenzied Regeneration, Barkskin, Feral Charge 1238122, Shifting Power 1322605, Poised, Omen of Clarity, Innervate, Tranquility, Rebirth, Ravage, Pounce, Shred, Claw, Rake, Rip, Ferocious Bite, Primal Bite, Lacerate), spellconst/druid.json (school masks for Hurricane, Wrath, Insect Swarm, Starfire, Moonfire; Feral Charge (Cat) 414924 cooldown), races.json, combos.json, items.json, addon-builds.json, bis/druid-*.json; APLs data/curated/apl/druid-*.json.
- Blizzard, "World of Warcraft: Forever Class Deep Dives: Hunter and Druid", https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid (fetched 2026-11-04).
- Blizzard, "World of Warcraft: Forever Deep Dive Panel Recap", https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200).
- Wowhead, "Skyborne First Look", https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829 (200).
- Talents Forever, https://talentsforever.com/data.json (200).
- Forever Sixty: https://foreversixty.gg/sim/weights, https://foreversixty.gg/sim/drops, https://foreversixty.gg/bis/druid/feral-bear, https://foreversixty.gg/bis/druid/restoration (all 200).

## Fix log (round 8)

- DRUID-R8-1 fixed, balance.md:105: "Nature damage feeds Wrath, Insect Swarm and Hurricane, never Starfire or Moonfire." (spelltips 17402 "Nature damage").
- Decisions 1 to 3 applied (no other change). SHOULDs (feral Bite energy line, Faerie Fire "always") left: wording only.
- Prose check: druid clean. No new SIM findings.
