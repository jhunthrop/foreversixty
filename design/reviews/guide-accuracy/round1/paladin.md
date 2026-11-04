# Guide accuracy round 1: Paladin

Auditor: paladin (Holy, Protection, Retribution, class index). Audit tree origin/main 64168181f, client 1.60.1.70755, engine pin 2eaae03c6. Read-only; no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`.

70753 to 70755 for paladin: talents/paladin.json changes one row only, Two-Handed Weapon Specialization goes from "physical damage" to "Physical and Holy damage" (2/4/6%). trainables/paladin.json is unchanged apart from the build stamp. Main's guides already carry the 70755 wording.

## Builds decoded (tab order = talents[] order in talents/paladin.json; all four legal on 70755: 51 points, tier gates and prerequisites met)

**Retribution `build`** `050250002/502/0523330100130135` (14/7/30)
- Holy: Divine Strength 5, Reverence 2, Improved Seals 5, Holy Conduit 2
- Protection: Divine Intellect 5, Precision 2
- Retribution: Benediction 5, Improved Judgement 2, Conviction 3, Sanctified Judgement 3, Vindication 3, Seal of Command 1, Sacred Arbiter 1, Two-Handed Weapon Specialization 3, Repentance 1, Vengeance 3, Champion of the Light 5

**Retribution `raidBuild`** `050050002/503/05233301001301351` (12/8/31): as above minus Reverence 2, plus Precision 3rd point and Unsealed Wrath 1.

**Protection `build`** `05005/15302102513501031/5003` (10/33/8)
- Holy: Divine Strength 5, Improved Seals 5
- Protection: Divine Intellect 1, Redoubt 5, Precision 3, Sacred Duty 2, Templar's Bulwark 1, Shield Specialization 2, Toughness 5, Swift Judgement 1, Improved Righteous Fury 3, Anticipation 5, Holy Shield 1, Iron Creed 3, Shieldsworn 1
- Retribution: Deflection 5, Conviction 3

**Holy `build`** `50320215012101251/50322/0503` (31/12/8)
- Holy: Spiritual Focus 5, Healing Light 3, Reverence 2, Purifying Power 2, Divine Favor 1, Illumination 5, Divine Precision 1, Infusion of Light 2, Oath of Light 1, Holy Shock 1, Boundless Light 2, Holy Power 5, Light's Vigil 1
- Protection: Divine Intellect 5, Precision 3, Guardian's Favor 2, Sacred Duty 2
- Retribution: Benediction 5, Conviction 3

Every talent the prose names is in the matching build with the stated points. Spell levels, ranks and costs checked against trainables/paladin.json and spellconst/paladin.json: Holy Strike 6 (10 s, 20 mana at rank 8), Judgement 4 (10 s, off the GCD, 6% base mana), Seal of Fury and Righteous Fury 10 (RF 30 min, 30% base mana), Consecration 20 (rank 1 135 mana), Exorcism 20, Blessing of Kings 20, Hammer of the Righteous 40, Hammer of Wrath 44 (1.0 s cast, 6 s cd), Holy Wrath 50, Holy Shield 30/40/50/60 (rank 4 1239570, 240 mana), Holy Shock rank 1 160 vs rank 4 325 mana, Flash of Light rank 1 35 mana, Holy Light rank 9 is a book rank (book-ranks.json), Light's Vigil 1340 at 60, Judgement of the Crusader/Wisdom 40 s, Repentance 6 s / 1 min, Blessings and Greater Blessings 1 h. Blessing of Light, Salvation and Sanctuary and Sanctity Aura have no trainer row. No removed or SoD-rune spells named. Rotations match the APLs (data/curated/apl/paladin-*.json) step for step, including the Unsealed Wrath twist, Prot's -3.5 s Seal / -2 s Holy Shield opener and Holy's -1.5 s Oath of Light. Consecration ranks match rank-choices.json.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix | also on r14 |
|---|---|---|---|---|---|---|---|
| PALADIN-1 | CONFIRMED | stats | retribution.md:9 | `statPriority: [Strength, Attack power, Critical strike, Hit, Agility, Melee haste]` leaves out Intellect and Spell power, so the stat table hides the second-best stat at raid 60 and the best at bare 60 | bis/paladin-retribution.json band 60 scale factors: raid Hit 2.16, Intellect 1.07, Strength 1.00, Crit 0.99, Agility 0.77, Spell power 0.57, AP 0.41; bare Intellect 2.88, Hit 1.84, Crit 1.29, Spell power 1.14, Strength 1.00 | `statPriority: [Hit, Intellect, Strength, Critical strike, Agility, Spell power, Attack power, Melee haste]` | fixed |
| PALADIN-2 | CONFIRMED | stats | retribution.md:103 | "Strength is your main stat after that ... Intellect also counts" | same weights: Intellect is ahead of Strength per point in both 60 presets | "Hit comes first until the cap: every miss is a lost swing and a lost Seal proc, and Precision and Artisan's Precision both help. Intellect is next. Champion of the Light turns 60% of it into spell damage for your Seals, Judgements and Consecration, and it is your mana bar, which is what decides how much you press in a long fight. Then Strength, which Divine Strength raises by 10%, and crit, which feeds Vengeance as well as itself." (drop "The weights above are per point at level 60", see PALADIN-10) | fixed |
| PALADIN-3 | CONFIRMED | talent | retribution.md:62 | Conviction "does not reach Consecration or Exorcism" | engine talents.go:69-83: since hotfix push 113139 Conviction is aura 290, crit on every table, applied as `stats.Crit` (stats.go:34-38, all tables); memory forever-dots-crit: Consecration is a non-channelled DoT and always can crit; Consecration ticks through `dot.OutcomeTick` (consecration.go). Exorcism is the one exception: exorcism.go:86-92 never crits (server, beta logs 0 in 641) | "**Conviction** (3 points): 6% crit on everything you do while you hold a melee weapon: swings, Holy Strike, Seal procs, Judgements and Consecration's ticks. Exorcism is the exception, it never crits in Forever." | partly: r14 line 62 says "Consecration and Exorcism included", which contradicts exorcism.go:86-92. Keep main's Exorcism half |
| PALADIN-4 | CONFIRMED | spell | protection.md:109, holy.md:92 | "Lay on Hands is a full heal on a 20 minute cooldown" with no word that it empties your mana | engine lay_on_hands.go:58 `SpendMana(sim, paladin.CurrentMana(), ...)`, the Classic spell's drain; spellconst 633/2800/10310 heal max health and give the target 0/250/550 mana | Prot: "Lay on Hands heals you to full on a 20 minute cooldown, but it takes every point of your mana with it, and with no mana you have no Holy Shield and no threat. Save it for the swing that would kill you, and have a mana potion ready after it." Holy: "Lay on Hands is a full heal on a 20 minute cooldown, and it empties your mana bar. Use it on someone who is about to die when you could not have saved them with Divine Favor and Holy Light, ideally late in the fight." | not fixed |
| PALADIN-5 | CONFIRMED | leveling | protection.md:129 | Leveling gives no Protection talent order; at 20 the owner ruling is Templar's Bulwark | memory forever-refresh-hold, round 14 owner rulings: "Bulwark at Prot pal b20"; main's bis/paladin-protection.json band 20 talents `15302` (Divine Intellect 1, Redoubt 5, Precision 3, Sacred Duty 2) predate the ruling | add: "If you tank as Protection from 20, take Redoubt and Precision first, then Templar's Bulwark as your 11th point ahead of a second point in Sacred Duty: a health bar of absorb saves a low-level tank far more often than a shorter cooldown. Holy Shield is your 21st point at 30 and Shieldsworn your 31st at 40." | fixed (Bulwark sentence; talent walk not added) |
| PALADIN-6 | CONFIRMED | stats | protection.md:8-22 | statPriority has no Block value, so the table drops a stat that outranks Stamina | bis/paladin-protection.json band 60 raid scale factors: Block value 9.05, Stamina 6.96 (bare 8.07 vs 9.25) | add `Block value` after `Block` | fixed |
| PALADIN-7 | LIKELY | stats | retribution.md:72 | "hit is worth more to you than any other stat per point until you reach the cap" | true only at raid 60 (Hit 2.16 top). Bare 60: Intellect 2.88 > Hit 1.84; bands 20 to 40 Hit sits below Strength and Spell power | "...and leads to **Precision**. With raid buffs at 60, hit is your best stat per point until you reach the cap." | not fixed |
| PALADIN-8 | CONFIRMED | stale-70755 | retribution.md:17,29; protection.md:30,39; holy.md:16,25 | source labels read `data/builds/1.60.1.70753/...` while the URLs point at 70755 | active build 1.60.1.70755 | relabel `data/builds/1.60.1.70755/talents/paladin.json` and `.../races.json` | not fixed (r14 moves the URLs and build strings back to 70753; the integrator must re-point all of them to 70755 on rebase) |
| PALADIN-9 | CONFIRMED | stale-70755 | r14 only: retribution.md:65-66 on int/r14 | r14 rewrites Two-Handed Weapon Specialization to "6% more physical damage ... It no longer touches your Seals, Judgements or Holy Strike" and Vengeance to "up to 9%" | talents/paladin.json 70755 node 105697: "Increases the Physical and Holy damage you deal by 6%"; engine talents.go:85-98 school mask Physical+Holy. Vengeance node 105693 rank 3: 3% a stack, 5 stacks; engine talents.go:120-128 `0.01*rank*stacks`, max 5 = 15% | keep main's lines 65 and 66 as they are when r14 lands | r14 regresses (main is right) |
| PALADIN-10 | LIKELY | voice | retribution.md:74, protection.md:84, holy.md:70 | "Open the planner at [/planner?class=paladin](/planner?class=paladin) to change it." | brief rule 9: no talk about the site | delete the sentence | not fixed |
| PALADIN-11 | LIKELY | voice | retribution.md:103; protection.md:113 | "The weights above are per point at level 60." / "The table weighs every stat by how much it lowers the damage you take and how spiky it is, then by threat." | method talk, not player advice | Ret: drop the sentence (PALADIN-2 text). Prot: start with "Expertise and Defense come first." as r14 does | Ret not fixed; Prot fixed |
| PALADIN-12 | NEEDS-SIM | consumable/race | retribution.md:111 | "**Human** is the Alliance pick: Ruthless is 1% crit on everything" | raid 60 per-percent weights: hit 52.2, crit 33.7 per 1%. Artisan's Precision (1% hit, races.json dwarf) is worth about half again Ruthless (1% crit) until hit cap, and raid hit_to_cap is still 3 points short with Precision | NEEDS-SIM: Dwarf vs Human Retribution at 60, raid and bare, guide gear and raidBuild. If Dwarf wins beyond error: "**Dwarf** is the Alliance pick until you reach the hit cap: Artisan's Precision is 1% hit, and Stoneform takes the edge off a hard pull. **Human** is close behind with Ruthless, 1% crit on everything." | n/a |
| PALADIN-13 | LIKELY | consumable/race | holy.md:104 | "**Human** brings The Human Spirit, 10% Spirit, and Perception." leaves out Ruthless | races.json human: Ruthless +1% crit with all spells and attacks | "**Human** brings Ruthless, 1% crit on every heal and so more Illumination refunds, and The Human Spirit, 10% Spirit." | not fixed |
| PALADIN-14 | LIKELY | voice | holy.md:43 | Light's Vigil "costs as much as ten Flash of Lights" | Vigil rank 3 1340 mana; Flash of Light rank 6 140, rank 1 (the guide's own staple) 35, so it is 38 of the casts you actually make | "...Light's Vigil, which costs as much as ten top-rank Flash of Lights" | not fixed |
| PALADIN-15 | LIKELY | talent | index.md:29 | "Retribution's Unsealed Wrath brings a new kind of swap back at 60." | Unsealed Wrath is the 31st Retribution point (tier 6), so it is available from 40; it pays only with raid attack power (APL gates at 1200 AP) | "Retribution's capstone, Unsealed Wrath, brings a new kind of swap back for raiders." | not fixed |
| PALADIN-16 | LIKELY | consumable/race | retribution.md:115, protection.md:125 | Professions never name Elixir of Holy Power, Forever's new paladin elixir | proto/common.proto:495-500: item 21546, 40 Holy spell damage for 30 minutes (the old Greater Firepower recipe 21547 makes it); carried for Retribution and Protection in data/curated/presets.json raid consumes | Ret: "Herbalism and Alchemy keep you in mana potions and Elixirs of Holy Power, Forever's 40 Holy damage elixir, which reaches every Seal proc, Judgement and Consecration tick." Prot: add "Elixir of Holy Power adds to every Holy threat button you have." | not fixed |

Claims checked and correct (no finding): every talent description line in all three guides; Shieldsworn bypassing Shield Specialization's 3 s limit; Holy Shield 4 charges and 10 s; Judgement of Fury taunt 3 s (20232); Swift Judgement free and off the GCD (1310994); Templar's Bulwark full-health absorb 8 s, Forbearance 1 min, Sacred Duty -60 s / -10 s; Holy Power 175%; Divine Favor 2 min; Blessing of Protection 3 min with Guardian's Favor; Holy Shock rank 1 half the cost; Consecration rank 1 below 60 and rank 4 at bare 60 for Ret; index facts (Seal duration 30 s, Judgement keeps the Seal, Blessings 1 h, the 11/16/21/31 milestone rows, Holy Shield on the 21-point row, Human/Dwarf/Undead only per combos.json); racials per races.json (Stoneform 10% physical, Artisan's Precision 1% hit, Will of the Forsaken no immunity, Cannibalize mana, Touch of the Grave 15 s); Major Mana Potion 13444, Demonic Rune 12662, Flask of Distilled Wisdom 13511 exist. No em or en dashes. No sub-40 plate claim.

## Sim findings

- **S1 (NEEDS-SIM), Holy band 30 skips Oath of Light.** bis/paladin-holy.json band 30 talents `50320215012000000` put the 19th point into Divine Precision (Holy spell hit, nothing for a healer) and the 20th and 21st into Infusion of Light. They leave out Oath of Light, which the guide calls "the biggest point in the tree" and which the band-40 kit and the APL both use. Question: at 30, bare, heal profile, does Oath of Light (on the tank, pre-pull) in place of Divine Precision 1 heal more beyond error? If not, find out why Oath is weak at 30 in the engine (the client's "up to 0" scaling text).
- **S2 (NEEDS-SIM), Retribution race.** See PALADIN-12: Dwarf vs Human, raid and bare 60.
- No engine-vs-game conflicts found for the guide claims. Conviction (aura 290, all tables), Two-Handed Weapon Specialization (Physical+Holy on 70755), Vengeance (15%), Consecration crits and Lay on Hands' drain all agree with the guides once PALADIN-3 is fixed.

## Owner questions

- **O1.** Exorcism never critting comes from the beta logs (exorcism.go:86-92), not from a client flag. Is it still true on live? PALADIN-3's fix depends on it.
- **O2.** Does Lay on Hands still drain all of the paladin's mana in Forever? The engine and Classic say yes; spellconst carries no drain attribute to check against.
- **O3.** Kings vs Heart of the Lion stacking is still open (presets.json note). All three guides recommend Greater Blessing of Kings, which is fine either way, but the raid preset leaves Kings out.

## Fix log

Branch gr1/paladin, commit 8813cd492 (off the r14 merge ff0f466b0). No tests run.

- PALADIN-1: fixed by r14 (retribution.md:9 already matches the band 60 raid order; Melee haste stays last since its factor is per 1%, outside the normalization).
- PALADIN-2: fixed by r14 (retribution.md:103).
- PALADIN-3: fixed (retribution.md:62). r14's "Consecration and Exorcism included" removed. Line now names swings, Holy Strike, Seal procs, Judgements and Consecration's ticks, and says nothing about Exorcism crits either way (owner question O1).
- PALADIN-4: fixed (protection.md:109, holy.md:92). Mana drain checked in the pinned engine, lay_on_hands.go:58.
- PALADIN-5: fixed by r14 (protection.md:129, Bulwark sentence).
- PALADIN-6: fixed by r14, then re-ordered. protection.md:8-21 is now the band 60 raid order: Defense, Block, Dodge, Parry, Agility, Block value, Stamina, Hit, Strength, Crit, Armor, AP. Expertise dropped because no item in items.json carries it. Prose at protection.md:113 no longer leads with Expertise and now covers Agility and block value.
- PALADIN-7: fixed (retribution.md:72). Checked against the bis weights: bare 60 Int 2.88 > Hit 1.84; raid 60 Hit is top.
- PALADIN-8: fixed (all six labels in retribution/protection/holy now read data/builds/1.60.1.70755). No 70753 left in the paladin guides.
- PALADIN-9: no regression on the merged tree. retribution.md:65-66 still say Physical and Holy 6% and 15%.
- PALADIN-10: fixed (planner sentences removed at retribution.md:74, protection.md:84, holy.md:70).
- PALADIN-11: fixed by r14 (both).
- PALADIN-12: left (NEEDS-SIM, race pick).
- PALADIN-13: fixed (holy.md:104). races.json human says Ruthless is +1% crit on all spells and attacks.
- PALADIN-14: fixed (holy.md:43). trainables: Flash of Light 140, Light's Vigil 1340.
- PALADIN-15: fixed (index.md:29). Unsealed Wrath is tier 6 in talents/paladin.json, so it can be taken from 40.
- PALADIN-16: fixed (retribution.md:115, protection.md:125). presets.json carries item 21546, 40 Holy spell damage.
- Extra: holy.md:8 statPriority reordered to the band 60 raid order (MP5, Intellect, Healing power, Critical strike, Spirit). Spell haste is insignificant and left out. The prose at holy.md:96 already matched.

Confirmed: 11 (PALADIN-1,2,3,4,5,6,8,9 from the audit, plus Holy statPriority, Prot Expertise, and the r14 Exorcism-crit regression inside PALADIN-3). Fixed in this commit: 7 confirmed plus 5 verified LIKELY. Fixed by r14: 1,2,5,11. PALADIN-9 needed no change.

SIM findings / build proposals (no build strings changed):
- S1: Holy band 30 talents skip Oath of Light for Divine Precision. Sim Oath of Light vs Divine Precision at 30.
- S2 / PALADIN-12: Dwarf vs Human Retribution at 60, raid and bare.
- Owner O1 (Exorcism crit) and O2 (Lay on Hands drain on live) are still open.

## Fix pass 2

MUSTs:
- MUST 1 (Holy leveling order): fixed, holy.md:112. Prose now follows the Leveling strip: "this build picks up Oath of Light in your 30s and Holy Shock in your 40s". The strip walks the guide build (lib/guides/leveling-builds.ts); band 30 at 21 points has no Oath (see S1). Oath earlier is a SIM finding (S1, restated below).
- MUST 2 (mana pacing in sim terms): fixed. Holy: holy.md:80 "keep your mana percent no more than 5 points under the percent of the fight left (halfway through, above 45% mana). Once you fall more than 5 points under, heal with Flash of Light alone", matching the condition currentManaPercent > remainingTimePercent - 0.05; card note paladin-holy.json:195 reworded the same way (notes only). Retribution: retribution.md:95 "your mana percent is at least four fifths of the percent of the fight left (halfway through, 40% mana or more)", matching currentManaPercent >= 0.8 x remainingTimePercent; card note paladin-retribution.json:415 reworded the same way (notes only).
- MUST 3 (Protection Salvation line): fixed, protection.md:59 "Forever has no Blessing of Salvation, so your damage dealers keep all their threat and you have to out-threat them without that help". Claim supported: trainables/paladin.json lists Freedom, Kings, Might, Protection, Sacrifice, Wisdom and the three Greaters, no Salvation (and none in spellranks.json).

SHOULDs:
- Protection races: fixed, protection.md:121. The reviewer's "Human's racials do nothing for a tank" is NOT true per races.json: Will to Survive removes all stuns (3 min). Line now names Will to Survive (stun break) and Ruthless (1% crit, threat), and says The Human Spirit does little for a tank.
- Retribution Consecration rank: fixed, retribution.md:95 "At 60, use rank 4 without raid buffs and rank 5, the top rank, with them" (rank-choices.json: rank 4 at 60 bare; card authors 20924 rank 5; spellranks Consecration tops at rank 5).
- Site-talk lines ("Raid loot is not itemized...", "Crit, hit and haste enchants are not listed yet.", "Numbers from client build ..."): not in any paladin guide markdown; they render from components. Left (R2).

New SIM findings:
- S1 (restated, NEEDS-SIM): the Holy leveling walk reaches Oath of Light only in the 30 to 39 range and Holy Shock in 40 to 49. Test moving Oath of Light ahead of Divine Precision (and Infusion of Light) so it lands at 25 to 29; keep if it heals more beyond error at bands 30 and 40.

## Owner questions resolved by research

Open items listed: O1 (Exorcism crit), O2 (Lay on Hands mana), O3 (Kings vs Heart of the Lion), PALADIN-12 / S2 (Dwarf vs Human Retribution), S1 / Fix pass 2 S1 (Oath of Light earlier in the Holy leveling order), plus the integrator's added item L1 (Forever's low-rank healing and damage bonus rule against every rank the guides recommend). Resolved 5 of 6; one left with the owner.

Sources:
- Blizzard, World of Warcraft: Forever Class Deep Dives, Paladin: https://news.blizzard.com/en-us/article/24302099 (also https://worldofwarcraft.blizzard.com/news/24302099)
- Warcraft Tavern, All Changes to Paladins in WoW Forever: https://www.warcrafttavern.com/forever/news/all-changes-to-paladins-in-wow-forever/
- Warcraft Tavern, Forever Holy Paladin Guide: https://www.warcrafttavern.com/forever/guides/holy-paladin/
- wow.gg, WoW Forever Retribution Paladin (beta): https://wow.gg/guides/wow-forever-retribution-paladin-beta
- wowtbc.gg, Retribution Paladin PvE Guide, WoW Forever: https://wowtbc.gg/warcraftforever/class-guides/retribution-paladin/
- Kami-labs, Holy Paladin Build, WoW Forever: https://kami-labs.fr/en/wow-forever/build-paladin-sacre-wow-forever/
- Wowhead Forever spell pages: Lay on Hands https://www.wowhead.com/forever/spell=633/lay-on-hands and https://www.wowhead.com/forever/spell=10310/lay-on-hands ; Exorcism https://www.wowhead.com/forever/spell=879/exorcism ; Classic Era comparison https://www.wowhead.com/classic/spell=633/lay-on-hands and https://www.wowhead.com/classic/spell=10310/lay-on-hands
- wago.tools SpellMisc exports: https://wago.tools/db2/SpellMisc?build=1.60.1.70755 (Forever) and build 1.15.8.63829 (Era), read 2026-11-03
- Blizzard forums (SoD), Blessing of Kings doesn't stack with Aspect of Lion BUT: https://us.forums.blizzard.com/en/wow/t/blessing-of-kings-doesnt-stack-with-aspect-of-lion-but/1738967
- Game Rant, WoW Season of Discovery: Best SoD Runes, Ranked (Heart of the Lion "doesn't stack" with Kings): https://gamerant.com/wow-season-of-discovery-best-sod-runes-ranked-each-class/
- Warcraft Wiki, Heart of the Lion: https://warcraft.wiki.gg/wiki/Heart_of_the_Lion
- Low-rank rule sources as gathered in round1/priest.md: Blizzard Forever Beta Development Notes (Oct 1, Nov 3) https://us.forums.blizzard.com/en/wow/t/forever-beta-development-notes/2360696 ; ForeverChanges downranking https://foreverchanges.pro/downranking ; ForeverWisp https://www.foreverwisp.com/guides/wow-forever-downranking-spell-ranks ; Master of Warcraft https://www.masterofwarcraft.net/2026/10/wow-forever-spell-downranking-nerf-healers.html
- Forever beta logs as summarized in engine commit f1c7fa888 (wowsims-forever, 2026-11-02): foreverlogs.gg, 39 paladins at 20 to 30.
- 70755 data: trainables/paladin.json (rank learn levels), talents/paladin.json (Holy tiers), spellconst/paladin.json (Oath of Light 1328548 per-level scaling, Lay on Hands 633/2800/10310).

### O1: can Exorcism crit in Forever?
**Answer: RESOLVED, no.** No web source addresses it (Deep Dive, Warcraft Tavern, wow.gg, wowtbc.gg are silent; the "always crits on demons" posts are a later expansion or a player wish). The client gives no answer either way: SpellMisc Attributes_2 for 879 to 10314 is 0 in both 70755 and Era, so no cannot-crit flag, and Wowhead Forever lists no such flag. The deciding evidence is the beta logs: 0 crits in 641 Exorcism hits from 39 paladins, against 18 of 367 Judgement of Righteousness crits by the same players and about 35 expected Exorcism crits. That is far outside chance, so it is the server's rule.
- Prose changed: retribution.md:62 adds "Exorcism is the exception: it never crits in Forever." (the engine pin already rolls hit only, exorcism.go).
- Owner: re-check on a live 60 log after launch; nothing to decide.

### O2: does Lay on Hands drain all mana in Forever?
**Answer: RESOLVED, no, it costs nothing.** Era's SpellMisc Attributes_1 for 633, 2800 and 10310 is 131074 (bit 0x2, "Uses all power", Wowhead Classic shows the flag); Forever 70755 has 131072 for all three ranks, the flag removed, and Wowhead Forever no longer lists "Uses all power". The target still receives 550 mana (rank 3). Blizzard's Deep Dive only says the cooldown went to 20 minutes. Warcraft Tavern's Forever Holy guide still says "it costs all of your mana", but that page predates the racial rework in other places (Dwarf crit damage) and contradicts the Forever client rows, which win.
- Prose changed: protection.md:109 ("in Forever it costs no mana, so you keep your Holy Shield and your threat after it"), holy.md:92 ("in Forever it no longer empties your mana bar, so there is no reason to hold it for the end of a fight").
- **RESOLVED, needs sim round:** engine sim/paladin/lay_on_hands.go:58, drop `paladin.SpendMana(sim, paladin.CurrentMana(), layOnHandsManaMetrics)`; keep the heal and the target mana gain. Prot card note at paladin-protection.json:61 is still right.

### O3: does Blessing of Kings stack with Heart of the Lion?
**Answer: RESOLVED, they do not stack.** Forever uses the Season of Discovery spell (409583; its SpellMisc row is identical in Era and 70755), and in SoD the two did not stack: Game Rant's rune ranking ("which doesn't stack") and the SoD forum thread (Kings replaces the Lion aura on the target; only the hunter's own extra 10% stays). No Forever source reports a change. The raid preset already leaves Kings out, which is correct; no preset change.
- Prose changed: protection.md:64 and :88 (Might with a hunter in the raid; Kings only where no hunter's Heart of the Lion covers you), retribution.md:54, holy.md:50 (Kings only where no Heart of the Lion covers the raid).
- Remove the "if they do stack" hedge from the presets.json note when the integrator next edits it (not my file).

### PALADIN-12 / S2: Dwarf vs Human for Retribution
**Stays with the owner (taste inside the error).** No Forever guide names a winner; wowtbc.gg says "each race for a spec are within 1% DPS of each other", which matches Blizzard's "tuned for similar offensive power". The math leans Dwarf in raids: Artisan's Precision (1% hit) beats Ruthless (1% crit) per point until the hit cap (raid 60 weights 52.2 vs 33.7 per 1%), and the raid build is still short of the cap; Human's Spirit helps only while leveling.
- Options: (a) keep Human (guide build strings and recommendedRaces today); (b) Dwarf for the raid build, Human for leveling; (c) Dwarf for both. Recommendation (b), if the sim round shows Dwarf ahead beyond error at raid 60. Prose not changed: retribution.md:111 describes both truthfully, and flipping it needs the build-string race change.
- Needs sim round: Dwarf vs Human Retribution at 60, raid and bare, raidBuild.

### S1: Oath of Light earlier in the Holy leveling order
**Answer: RESOLVED, Oath of Light at 25.** No Forever guide publishes a Holy leveling order (Warcraft Tavern and Kami-labs give level 30 or 60 builds only). 70755 data settles it: Oath of Light, Divine Precision and Infusion of Light share tier 3 (15 points, so the 16th point at 25 is the first legal one), Divine Precision is only Holy spell hit and silence reduction (nothing for a heal), Infusion of Light needs Flash of Light or Holy Shock crits, and Oath of Light's bonus scales with level (points per level 6.67 / 1.92 / 3.83 on Holy Light / Flash of Light / Holy Shock), so at 30 it is still half the 60 bonus, a large share of a level 30 Flash of Light.
- Prose changed: holy.md:112 now says take Oath of Light at 25, ahead of Divine Precision and Infusion of Light.
- **RESOLVED, needs sim round:** bis/paladin-holy.json band 30 talents `50320215012000000` to `50320215002100000` (Divine Precision 1 to 0, Oath of Light 0 to 1; tier 3 reads Divine Precision 0, Infusion of Light 2, Oath of Light 1), and the Holy leveling walk (lib/guides/leveling-builds.ts order / addon-builds) puts Oath of Light as the 16th point. Until then the Leveling strip shows Oath in the 30s and the prose is the game-correct side.

### L1: Forever's low-rank rule against every recommended rank
**Answer: RESOLVED.** Blizzard's Oct 1 note: ranks "vastly below your current level receive less benefit from your Spell Damage and Spell Healing"; measured: full bonus for 17 levels after a rank's learn level, then 5% less per level, none from 37 levels on (sources above, as in round1/priest.md). Oath of Light's flat amount is its own aura, not healing power, so it should still land in full on any rank (unmeasured). At 60 by 70755 learn levels: Flash of Light rank 1 (20) keeps none, rank 4 (42) about 95%, ranks 5 and 6 all; Holy Shock rank 1 (30) about 35%, rank 3 (48) and 4 all; Holy Light rank 8 (54) all; Consecration rank 4 (50) and 5 all, rank 1 (20) none from 57.
- Prose changed: holy.md:9 (description drops "rank 1"), :36, :81 (tank staple Flash of Light rank 4, rank 1 named as the trap), :82, :84 (Holy Shock rank 3 or 4), :96 (healing power reaches only ranks near your level), :116 (common mistake now rank 1 at 60); retribution.md:121 (Consecration rank 1 through your 30s, then the newest rank you can afford). Protection recommends no low rank.
- **RESOLVED, needs sim round:** (a) engine: the low-rank healing and spell damage bonus curve (as priest/druid); (b) Holy card: Flash of Light on the tank 19750 r1 (paladin-holy.json:575) to 19941 r4, Holy Shock 1311606 r1 (lines 246, 306, 366, 425, 484, 543) to 20929 r3, rerun the rank ranker with the curve on; card notes at paladin-holy.json:255 and :584 still describe rank 1 and are rewritten with the action change (actions untouched this round, so notes left matching the card); (c) Retribution rank-choices: rerun the leveling Consecration rank with the curve (rank 1 below 60 today).

Left for the owner: PALADIN-12 only (race, recommendation Dwarf for the raid build if the sim agrees).

### O2 withdrawn (2026-11-04, after round 1 landed)
The "Lay on Hands costs no mana in Forever" answer was wrong. The 70755 client's own Spell.Description_lang for 633 and 10310 reads "Drains all of the Paladin's remaining Mana when used, but does not interrupt Mana regeneration." The removed Era "uses all power" flag only explains why regeneration keeps running. The engine's SpendMana(CurrentMana()) (lay_on_hands.go:58) is right, so the "needs sim round" item is dropped. protection.md:109 and holy.md:92 revert to the drain wording from sim round 15 (design/reviews/2026-11-04-sim-audit-round15/paladin.md, wave 2). Lesson for every round: a client description beats an inferred flag; read the spell text before claiming a cost or effect change.
