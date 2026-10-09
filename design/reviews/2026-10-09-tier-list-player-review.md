# Tier list mock: WoW player review (2026-10-09, round 1)

Spec: design/specs/2026-10-09-tier-list.md. Renders: design/mocks/renders/tier-list*.png.
Canvas for the owner: https://claude.ai/artifact/RFTPQM5vXCzGycNfFURezY

**Verdict: NOT YET.** Layout clean, next step clear (gold BIS on every row, the
signed-in callout answers "where is my spec" first). Least honest about tanks
and healers; the DPS letters carry more certainty than the data supports.

Findings, most misleading first:

1. Healer page: only "effective healing per second, overhealing not counted";
   no fight length, mana or raid utility; "Paladin −24%" reads as heals 24%
   worse; which boss profile is never said.
2. The honesty box sits under 20 rows (3,700 px down on phone), after the
   player has read "D"; it names cleave/adds/movement but none of what a
   raider asks (curses, Windfury, Shadow Weaving, Faerie Fire, Moonkin aura,
   Judgement of Wisdom, mana over 180 s).
3. Letters anchored to the top spec, which carries "Less settled"; Fury −4.8%
   is S and Demonology −5.5% is A on a 0.7-point difference; no error shown.
4. "Raid-ready, Phase 1" never says this is a full set including raid drops
   nobody has at launch; nothing about today's gear or a fresh 60.
5. Tanks ranked on effective health while the Taken column disagrees (Paladin
   3rd takes less than the Bear 2nd); "Share of the top" bar reads as "81% as
   good"; boss unnamed; "TAKEN /S ↓" has no unit word and the arrow is never
   explained; "Eff. health" three times.
6. Phone tank: caption says "the best of each is lit"; nothing is lit.
7. "Less settled" is vague, tooltip-only (none on phone), wraps on Destruction
   at phone width, and sits on the top spec.
8. Faction toggle changes nothing visible; race never named; no Horde render.
9. Tank and healer pages lack the "Fury Warrior is on the DPS list" pointer.
10. Phone is ~3,900 px: every row carries the 44 px link strip.
11. Spec name is not a link; GUIDE and PLANNER are muted 12 px grey; the
    planner link does not say which build.
12. Shared crests per class: Fury/Arms, three warlocks, three priests tell
    apart only by grey class text.
13. "Damage per second", "DPS", "DPS" three times; "BEHIND" head has no unit.
14. "Updated Oct 9" with no build in view.
15. States sheet's loading example says "28 reserved rows", shows 3.

Missing that a raider expects: fight length and boss for tank and healer;
error bars or "ties within X%"; the race simmed per faction; a note that the
order differs at current gear; one strongest raid-utility line per spec; an
on-screen explanation of "Less settled".

Three questions for the designer: (1) why anchor the ladder to a spec the
page flags less settled, and what does a player see about how firm a letter
is; (2) for tanks and healers, which boss, how long, what is in effective
health and HPS, why Paladin 3rd on a stat where it takes less damage, and
why no healer caveat; (3) what does the faction toggle change, which race per
side, and where does a fresh 60 see how the order applies to them.

# Round 2 (2026-10-09, after the designer's revision)

**Verdict: SHIP**, on two conditions. 11 of 15 round-one findings fixed, 3
partly, 1 open by design.

Fixed: healer boss/length/caveats (300 s stand-in Phase 1 fight, overhealing
not counted, mana and utility not ranked, "less healing in this fight, not a
verdict on the class"); honesty lines above the list; letters replaced by
10/20/30% rulers and "≈ tie" inside 1% (the site's gear margin, stated);
fresh-60 line linking the leveling bands; tanks sorted on damage taken per
second with units and the stand-in boss named; phone tank highlights; faction
toggle renders Horde with race under every spec; cross-role pointer; spec
name is the BiS link; units once.

Partly: "Stat weights less certain" is visible text but its explanation is a
tooltip (none on phone) and it sits on the row every gap is measured against;
the states sheet's skeleton caption (20, 3 or 5) vs 6 rows drawn.

Open, accepted: phone length (link strip on every row), shared crests (one-
crest rule), no build beside Updated (provenance rule).

New from the revision: (1) tank order disagrees with the BiS index, which
headlines effective health (Bear 32,796 EH ranks third here) — the condition
that decides the verdict; (2) "≈ tie" marks only the lower row of a pair and
chains drift past the margin (Horde Survival to Fire 1.2%); (3) the long box
repeats the notes above the list; (4) Affliction/Destruction rows cramped by
the weights note; (5) healer head says "Healing per second", HPS reads as raw.

Conditions to ship: make the tank sort key agree with the BiS index (sort the
index on damage taken too, or show both there); mark both rows of a tie and
fix the states caption. Polish: Horde tank and healer renders, drop the
duplicated box lines, the cramped rows.

# Built page (2026-10-09, pre-merge screenshots)

Player: **SHIP**; both conditions met (tanks on damage taken per second with
the figure on the BiS index; tie marks on both rows). Designer: NOT YET on
six points, resolved before merge: the tie marks the designer read as wrong
were right (each marked row has a partner within 1% below it; a vitest pins
the Alliance rows 6–14); the tie note uses the spec's wording; the bottom
box has its two bullets back and the notes above are shortened instead; the
phone header, which never had a menu on this site (eight doors stacked ~230
px), is now a 56 px bar with a Menu button, desktop unchanged, API chrome
refreshed, full Playwright green; tank figures are whole numbers on the tier
list and the BiS index; phone units match the desktop heads. Signed-in,
cross-role pointer, hover, focus and skeleton states captured
(tier-list-built-*.png, git-ignored). No runtime error panel exists: the page
is prerendered and a signed-in fetch failure omits the callout. Merged to
main as 8d6d1716 with Tier List in the nav between Rankings and Guides.
