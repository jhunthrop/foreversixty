# Lane P15 (content), review round 1 (2026-11-09)

Captures: design/reviews/2026-11-09-f-content/ (side-by-side-*, build-*, guides/, paperdoll/), commit e118797de.

## ux-designer
VERDICT: SHIP WITH FIXES
- MUST 1: guides/after-warrior-fury-bis-390.png, the prose renders after the lane and the doll is gone on a phone. Prose first.
- MUST 2: side-by-side-404-1440.png, the build adds the character lookup and spec lines the mock lacks; sign them into the mock or cut. Cut "It may have moved as the game changes."
- SHOULD: the band link "60" wraps alone at 390; the About band's rule and 33 px offset (PageHeader, batch E); the changelog's pitch line under the title; tiny row chevrons and same-grey source pills on the changelog; the unavailable panel 60 px narrower than the mock.

## wow-player
SHIP WITH FIXES
- MUST 1: at 390 put the Best in slot prose above the slot list and decide on the doll.
- MUST 2: at 1440 the lane list must carry all weapon slots (Off hand and Ranged had wrapped under the doll).
- MUST 3: cut the changelog intro "Everything that changed in World of Warcraft: Forever and on this site, dated and sourced."
- SHOULD: cut "It may have moved as the game changes."; the short right lanes on Leveling and Rotation; the About sources link; confirm Lionheart Helm "Requires Level 56"; the changelog chevrons.

## ai-slop-reviewer
Verdict: PASS WITH FIXES
- The guide layout passes; the doll fills the left column and the lane rows wrap under the prose.
- Fixes: cut "Picks for a Night Elf (Alliance)." and "Your best gear before the first raid opens." on the level 60 lane (the doll's title says both); cut the 404's second sentence and the lookup hints; cut the changelog subtitle and the month counters; put the doll on the 390 view.
- Should: colour or shrink the repeated changelog source pills; sticky the short Leveling and Rotation lanes or leave them.

## info-density-reviewer
**Verdict: DENSE WITH ADDITIONS**
- MUST: none. No guide section leaves the left column over 200 px beside a tall lane, in 6 guides at 1440 and 2000.
- SHOULD: the right column stands empty beside long prose (Leveling, Rotation): make the lane sticky or fill it; the doll and lane repeat the weapons; drop "Your best gear before the first raid opens."; the band links could share the lane head; Overview and Professions lanes are short or empty.

## Round 2 changes
- Best in slot reads prose first below 1280 (ReadingSection `proseFirst`), so a phone gets the prose, then the list. The doll steps aside under 423 px on purpose: the frame needs 423 px, and its list form would repeat the lane's 17 rows.
- The slot list never wraps: the weapons stay in the lane with the rest; only the links under it wrap.
- The level 60 lane drops "Picks for a ... (...)." and the pre-raid note (the doll's title bar carries both); other bands keep their caption.
- A short lane beside long prose stays pinned in view while the prose scrolls (reading-wrap `data-lane-sticky`, when it fits the viewport).
- The 404 line is "Nothing exists at this address."; the directions mock b4-404 now draws the spec lines and the lookup the page keeps. The lookup's hints are the shared CharacterSearch's (home and /characters), not this lane's.
- The changelog's title band drops the pitch line. The month counts stay: they follow the Show filter (Game, Site) and the info-density reviewer counts them as data.
- Not changed here: the PageHeader rule and band height (batch E owns PageHeader), the changelog's Dated chevrons and Tag pills (shared components; noted for the integrator).
