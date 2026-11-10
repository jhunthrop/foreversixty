# AI-slop review round 2: character and planner pages (1440 and 390)

Reviewer: ai-slop. Pages: character-visitor-60, character-owner-60-premium, planner, planner-you, planner-guide.
Handoff items (Client build footer, yellow dash marker in Point order, tier-gate gutter, paperdoll item level and Upgrades card) are noted, not failed.

## Verdicts
- character-visitor-60: PASS WITH FIXES
- character-owner-60-premium: PASS WITH FIXES
- planner: PASS WITH FIXES
- planner-you: PASS WITH FIXES
- planner-guide: PASS WITH FIXES

No visual tells 1, 2, 5, 12, 13 (planner blank borderline), 15. Reads as made by a player who raids: paperdoll, real item names with quality colour, HPS gains, parse-coloured best per encounter, real trees, point-order icon strip.

## Round 1 fixes confirmed
"Paste in game with /fs." gone; "Show DPS on the card" shortened; talents-loaded grey lines gone; "21 of 25 points"; Builds seen "Holy · 3 Nov"; Activity rows start with the verb; History/Coach rows name content ("Score and parses 7 fights", "Coach reviews"); "Loaded" is the only pill on guide builds, "Matches" is a real state.

## Remaining findings
1. Tell 8 (borderline, all five pages): every block is still the same dark bordered plate with serif heading; character rail stacks four to five (Score, Standing, Builds seen, Activity), planner stacks seven (Talents, Point order, Share, Import, Level build, In the game, Guide builds). Source: components/planner/Planner.svelte, components/character/CharacterOwnerSections.svelte. Fix: fold Share and Import into one plate with a hairline; Builds seen and Activity without plate chrome.
2. Tell 11 dead space: planner 1440 right rail ends ~935 vs left ~1480; planner-you rail ~1005 vs ~1830; planner-guide ~990 vs ~1795; character visitor rail ~1240 vs ~1415. Fix: Guide builds below Point order as a full-width row, or Share and Import as a right-rail pair.
3. Tell 22/16: Forever Score block still shows Score 1,830, Tier Epic, Rating 65 of 100 plus five bars (character, both widths). Fix: drop Tier and Rating rows; colour the score by tier.
4. Tell 22: Standing block (owner) wraps "#3 of 14 healers in Cx66 Olympus, 60 raid" with orphan "raid" at 1440; three unlabelled bold lines. Fix: label left, value right like Score: "#3 of 14 healers", "Holy A tier", "Ony 33 d 8 h".
5. Tell 22: band meta "Supporter", "Addon · 1 day ago" (owner, both widths), plain grey text with no state. Fix: "Synced 1 d ago"; Supporter as crest or dropped.
6. Tell 22: "Show DPS on the card" still has "on the card" (planner pages). lib/sim/planner-copy.ts. Fix: "Show DPS".
7. Tell 22: "51 of this build's points still to take" (planner 1440/390). lib/planner/copy.ts. Fix: "51 points left".
8. Tell 21: "Felsnap is level 34. Then the Beast Mastery guide's order." (planner-you). lib/planner/copy.ts:95. Fix: "Felsnap is 34. Then Beast Mastery guide order." (dash marker is f-selected.)
9. Tell 13 (borderline): planner blank, Point order plate holds only "No points spent yet." Fix: hide until a point is spent, or show the level 60 order dimmed.
10. Tell 13 (mild): planner-you Point order plate has a dashed border (1440/390) that reads as placeholder; planner-guide's is solid. Fix: solid like every other plate.
11. Tell 22: owner 390 Upgrade headline still wraps two lines. Fix: one line, truncate "Devout Skirt, Baron Rivendare".
12. Planner blank: "Copy for the game" and "Share" are dim teal and read as unstyled. Fix: hide until a build exists, or make enabled state obvious.

## Handoff, noted only
"Client build 1.60.1.70790" footer (TalentEditor, f-selected); yellow dash before Felsnap line (PointOrder, f-selected); tier-gate gutter and "n more to tier k" (f-content); paperdoll item level/enchant state and floating Upgrades card (h-doll).

## Authentic, keep
Paperdoll with real icons; 390 gear as a named list; +128.9 HPS lead number; Best per encounter parse colour plus HPS; class-coloured planner title and crest; real trees with ranks; point-order strip with per-level rank; guide builds named with split; Send this build to Felsnap; no glow, gradient, eyebrow or icon-per-bullet.
