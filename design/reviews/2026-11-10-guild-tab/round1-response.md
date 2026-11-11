# Guild tab: response to review round 1 (2026-11-11)

Round 1 verdicts: guild-officer-reviewer NOT YET, info-density-reviewer THIN, ux-designer NOT YET, wow-player NOT YET,
ai-slop-reviewer FAIL, brand-consistency-reviewer DRIFTS. What changed, on both apps unless marked, and what is ruled.

## Fixed
| Finding | Change |
|---|---|
| Officer 1, density 2, 3: roster says too little | Recently on carries `item level N` (roster `item_level` joined on character key); officers get **Missing raids** (roster `attendance_30`: present at under half of at least four nights, `2 of 8 nights`) |
| Officer 2, density 1, wow 9: loot is one line | Loot lists the newest night's drops (the web loot page's `GET …/loot/nights?limit=1`, app-tagged, fixture `guild.loot.nights`, contract test): item in its quality colour with its icon, the winner in class colour or `Not awarded yet`, the night's day as meta |
| Officer 5: officer rows below Your job | Needs you, Officer and Assignments sit right under the raid head; Your job follows |
| Density 4, wow 4, ux 15: Thornwick twice | Close to 60 leaves out names Recently on shows |
| Density 5, wow 10: progression | The tier's raids by name (`Onyxia's Lair`), `1 of 1 down · 6 pulls` |
| Density 6, brand 4: head says nothing at 60 | The spread shows at every level (at 60, 55 to 59, then your level under 60 or `raiders` at 60); the head line always carries the realm (`US PvP`) |
| UX 1: quiet line contrast on iOS | `Raids open at 60 · you're N` in the text ink on both apps |
| UX 2: Android `Holy Paladin Paladin` | The applicant line takes the spec alone when it already names the class, as iOS does |
| UX 5: unselected Guild icon not captured | `tabbar-01-home`, `tabbar-02-gear`, `tabbar-03-guild` on both apps |
| UX 7, slop 6, wow 14: `Roster ›` twice | Who is coming drops its link on the Guild tab; the Roster section's link is the one way |
| UX 8: ax5 guild head | The spread stacks at the accessibility sizes; the crest top-aligns (iOS) |
| UX 9, brand: crest | The crest sits in the one ringed circle (faction ring, raised fill) on both apps |
| UX 12: no chevron on rows that open pages | Recent logs and Progression rows carry the chevron (iOS; Android rows open on tap as Pages rows do) |
| UX 14, slop 3: explainer lines | `By name or realm` and `Claim it: …` are gone; the row reads `Search guilds` |
| UX 16: Applied cream on Android | Muted, as iOS and the mock |
| Slop 2: qualifiers | `Seen` is gone from the age; `6 pulls` not `6 pulls this tier`; the encounter group reads `Onyxia`, not `Onyxia assignments` under Assignments |
| Slop 4, wow 7: `Preview mention` | `Ping on Discord`, the web planner's words (Home and the Guild tab) |
| Wow 8: empty slots hidden | Every template slot shows; an empty one is muted with `0 of 2` |
| Brand 1: assignment names coloured by class | Others muted, yours in gold, the web planner's Tag tones |
| Brand 5: hardcoded type sizes | The head's name is PanelTitle's role (21/26), the figures FSType.itemLg / FsType.itemLg |
| Wow 1, 11: fixture facts | Kaelra a Frost Mage and Thalrin a Marksmanship Hunter everywhere; Scholomance 10/10 |

## Ruled or out of scope
- **Answer echo (UX 3)**: the fixture's sign-up write answers one canned raid, so a tap on Late snaps back to In on the device;
  the real flow is in the flows tests (unchanged by this lane).
- **Site hand-offs signed out (UX 4)**: Plan a raid, Setup, Search guilds and Open in planner open the live site in the
  system browser, which carries no app session; the app cannot sign the browser in. The links land on the right pages.
- **Card per block (slop 1, tell 8)**: the round-5 panel system is the apps' design language for every tab (the owner's
  directions spec); the Guild tab keeps it. Loot and Progression are no longer one-line cards in the officer's view
  (Loot lists drops).
- **Tab icon weight (UX 6, brand 6)**: the glyph is the site's own Guild door at the nav set's 24 pt with the same
  thicken; it stays the site's glyph, as the brief asks. Flagged for the owner if the solid banner reads heavy.
- **Before you pull, raid notes, Flask of the Titans (wow 2, 3)**: the Raid tab's existing sections and fixtures, not this
  lane's.
- **Onyxia template facts (wow 12)**: api/internal/guilds/raid_assign_templates.go, for the accuracy reviewer.
- **Android Calendar to the site (UX 11)**: unchanged from before this lane.
- **Needs data**: per-raider deaths and avoidable damage; boss kills with dates; an account-wide list of your applications.
