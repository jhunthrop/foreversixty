# Guides rebuild — wow-player review (mock, pre-build)

Reviewed: `design/mocks/renders/guides-{index,class,spec,spec-phone}.png` against
`design/specs/2026-10-04-guides.md`, `docs/tenets.md`, and the current live pages
(`design/mocks/renders/current/guides-*`, `bis-warrior-fury-1440.png`).

## Verdicts

- `/guides` index: **SHIP**
- `/guides/<class>` (Warrior landing): **SHIP WITH FIXES**
- `/guides/<class>/<spec>` (Fury Warrior, desktop + phone): **SHIP WITH FIXES**

Yes, I'd keep this tab open leveling and at 60. The 3-up rail genuinely beats Icy Veins — Build,
Rotation and Stat priority answer "what do I do" in one glance, with a real loadable tree and
"Load/Sim this build" buttons Icy Veins has no equivalent of. The old index was a sitemap; this
is a tool. The old spec page was a 3,800px scroll of prose with no identity; this one tells me
in three seconds this is a Warrior page and what the rail says to do.

## Findings

1. **[must fix] The two DPS figures on the class landing are the wrong faction's number.**
   `195.1 DPS` (Arms) and `241.5 DPS` (Fury) on `guides-class.png` are the Horde/Orc and
   Horde/Troll band-60 `set_dps` values (`data/builds/1.60.1.70009/bis/warrior-{arms,fury}.json`:
   195.09 horde-orc, 241.45 horde-troll). Both guides' own frontmatter (`recommendedRaces:
   [human, troll]` for Fury, `[human, orc]` for Arms) and the spec's own §4.D ruling ("the
   guide's own first `recommendedRaces` entry decides which faction") say these should be the
   Alliance/Human numbers: 261.99 (Fury) and 201.62 (Arms), traced to the mock generator's own
   hardcoded `SPEC_SET_DPS = {'Arms': 195.1, 'Fury': 241.5}` comment ("band-60 set DPS (horde...)")
   — contradicts the spec it's illustrating. A player who knows Fury should out-parse Arms by
   racials at the same tier would also notice these two numbers are closer together than the
   real Alliance figures (262 vs 202) suggest they should be. Fix: compute from the first
   recommended race per guide, matching the Leveling strip's own band-60 row (which correctly
   reads 51 points from the Alliance/Human entry).

2. **[blocker] `SpecTabs`' third tab is cut off mid-word on phone with no scroll affordance.**
   `guides-spec-phone.png`, right under the h1: "PROTECT" is sliced at the 390px edge, no fade
   gradient, no peeking edge of the next tab, nothing to signal "swipe me." The spec says
   `SpecTabs` already horizontal-scrolls under 1024px today — the current live phone page
   (`guides-warrior-fury-390.png`) doesn't even have this row, so there's no "same as today" to
   fall back on here. A player sees a cut-off word and reads it as broken chrome, not an
   affordance — exactly the screenshot-and-mock-in-guild-chat case. Needs a visible scroll cue
   (edge fade, or size tabs to guarantee at least a partial next-tab peek) before this ships.

3. **[must fix] The phone spec guide drops "On this page" entirely.** Spec §7.C says
   `SpecTableOfContents`'s existing inline, boxed, collapsible placement rides under the header
   "same as today." The current live phone page has it (`ON THIS PAGE  SHOW`, right under
   Load/Sim). `guides-spec-phone.png` goes straight from the Stat priority rail card into
   "OVERVIEW" with no jump box anywhere — a returning reader loses the one navigation aid a
   3,000px article needs most on a phone. Either the mock skipped building it or it's a real
   regression; either way it needs to show up before this ships.

4. **[polish] The Stat priority rail card is two-thirds dead space.** On `guides-spec.png` the
   three rail cards are equal height (grid/flex stretch), but Build and Rotation fill theirs
   while Stat priority's six pills sit in the top third and leave ~180px of blank panel below —
   visible in a crop at (908,428)-(1440,620) of the 1440 board. Tenet 11 cuts both ways: a card
   that's mostly empty reads as unfinished the same way a wall of equal-weight rows reads as
   dense. Either let the card size to content (cards don't need matched height — Build and
   Rotation already differ from each other slightly) or add the per-stat DPS-per-point numbers
   BiS's own `StatWeightsPanel` shows, trimmed to one line each, to use the space honestly.

5. **[polish] The Leveling strip tells me *how many* points, never *which* points.** Five rows,
   each "{band} · {points} points · Load in planner" — to see whether band 30's build actually
   picks up Flurry or Enrage first I have to leave the page and open a cold planner, which is
   the exact gap §10.1 says this strip exists to close for everything except the build itself.
   A thumbnail of each band's tree (even at the rail's own 120px non-interactive scale) would
   answer "what do I respec into" without the click. Not a blocker — `Load in planner` is one
   tap — but it's the one place on this page that still makes the reader go elsewhere to see
   the actual answer.

6. **Verified correct, called out because they looked suspicious at first glance:** the 32
   Fury / 19 Arms / 0 Protection split at 60 sums the talent string exactly
   (`35311103002000000`=19, `353211005050010051`=32, zero Protection); the five band totals
   11/21/31/41/51 match both factions' own band data and the spec's own "never a single
   truncated level-60 string" claim; the rail's four rotation lines (Bloodrage, Death Wish,
   Battle Shout, Bloodthirst) match `addon-data.json`'s level-60 `rotations["warrior-fury"]`
   verbatim, correctly truncated to four with `Full priority ↓` pointing at the real section.

7. **Race row is honest and useful without portraits** — Human filled/selected, Troll
   outlined in faction red as the Horde alternative, exactly `RacePillRow`'s documented
   "no invented art" ruling (§10.2) carried forward; no complaint here, this is the right call
   until real race art exists.

8. **Index grid is a clear win over the stacked list.** The old `/guides` (`current/
   guides-index-1440.png`) was nine coloured rectangles with a class name and plain text links —
   no icon anywhere, the textbook tenet-1 defect. The new 3×3 grid gives every class its real
   ringed crest at 56px in the class colour inside a proper panel card — reads as a tool, not a
   sitemap, and the signed-in "Your guide: Fury Warrior →" callout is exactly the front-door
   hand-off tenet 14 asks for.
