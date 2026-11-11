# wow-player gate, lane G2 (home "What's here" / "Leveling BiS by spec", /rankings)

Seen rendered: home-1440, home-390, rankings-1440, rankings-saltspine-1440 (signed out, real data).

## Must fix
None found.

## Should fix
1. "The most logged bosses on the rankings: Relic Guardian, Saltspine and Shadetooth" sits straight after "The first raid tier opens on 9 Dec with Barrow Deeps, Hyjal Summit and Onyxia's Lair." A raider reads those three as raid bosses. They are Excavation Site: Wetlands 5-player bosses (the rankings page says so). Say "dungeon bosses" or start a new sentence ("Kills are already ranked on Relic Guardian, ..."). The sentence is also not a raid claim, so it should not share the raid paragraph.
2. Link text and heading say "leveling BiS" / "Leveling BiS by spec" while the nav and the owner's rule (2026-10-10) say "BiS" / "Best in slot". The data also holds a second 60 set ("60 raid", 12 band rows per file), so "leveling" undersells it. Use "BiS by spec" and "BiS for each band from 20 to 60".
3. "Leveling BiS by spec" lists each spec by full name only. Feral Druid (Bear) under Tanks and Feral Druid under Damage are right, but the two Feral entries read as one spec to a search engine and a skimmer. Fine as is, just keep the "(Bear)" tag.

## Claims checked, true
- BiS: all 28 spec files (incl. druid-feral-bear) in data/builds/1.60.1.70790/bis cover bands 20/30/40/50/60; "each ten-level band from 20 to 60" holds. Every one of those specs has a guide in web/src/content/guides.
- Sim tools named exist: /sim/gear, /sim/weights, /sim/talents, /sim/drops. Planner, setup, tiers (damage, healer, tank) pages exist.
- First raid names (Barrow Deeps, Hyjal Summit, Onyxia's Lair) and 9 Dec match the countdown card. All six Forever race/class pairs and the class list on screen are consistent.
- Role split is right: Prot Warrior/Paladin and Bear are tanks; Holy/Disc/Resto healers; Survival under Damage (melee in Forever, fine). No em dashes in about.ts.
- Game words only, no explaining the obvious. Tone is a player's, not a pitch.

## Not verified
- "orders every spec on the tier list at 60": did not find a level statement in tiers/index.astro by grep; confirm the tier page is 60-only.
- "bring your character over with its gear and talents": per the addon/companion, not checked in code.

## Data oddities seen (not copy)
- rankings Saltspine: five rows (Mage x3, Paladin, Balance Druid) show 0 DPS with percentile 50 down to 10 and a 0:39 length. A player will mock a "0 DPS" 50th percentile. Likely test or wipe logs; filter or hide zero-DPS rows.
- Rankings intro ("ranked by DPS and HPS for each spec, class, ruleset and guild") is accurate and short.

SHIP WITH FIXES: (1) stop the dungeon-boss sentence reading as raid bosses, (2) "BiS" not "leveling BiS" in link text and heading, (3) confirm the tier-list "at 60" claim; zero-DPS rows on /rankings/saltspine are a separate data fix.
