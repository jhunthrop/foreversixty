# Stage 2 UX verdict (site alignment)

Reviewed from the stage2-* captures at 1440, 390 and 2000 against stage1-account-* and the before-* set. Fixture build: missing icons and API-less error states not counted.

## Verdict: HOLD

The alignment itself lands. Every route now opens on the account's page head (gold rule eyebrow, 44 px uppercase h1, no navy band) and sits on the shared warm card with compact or full CardHead, ActionLinks and Tags. /bis, guides, planner, guild, premium, setup, logs, report, tiers and home read as the same product as /account at 1440. Two 390 defects (clipped button, overlapping head) are visible and must be fixed first.

## Blockers

1. /rankings, stage2-rankings-390.png: the card head "FURY WARRIORS" wraps to two lines and the Characters | Guilds control overlaps the title text ("WARRIORS" is drawn under the CHARACTERS button). The aside must wrap under the title at 390, indented to the title edge (CardHead rule, below 480 px), or the control must move to a TabStrip below the head.
2. /bis/warrior/fury, stage2-bis-warrior-fury-390.png: in the "This set" card the second button "TALENTS IN PLANNER" runs past the card's right edge and is cut off. Buttons must wrap or stack full width inside the card at 390.

## Shoulds

1. /sim 390 and 1440 (stage2-sim-*): the character strip "PLAN  SIM  LOGS  RANKINGS FOR WARRIOR  FORGET" renders as bare tracked-caps text with no tab or link affordance, a second row under the TabStrip. Give it ActionLink or ghost-button treatment, or fold it into the page head meta.
2. /sim 390: the tool TabStrip shows TALENTS at the edge with the rest off-screen; check the 24 px edge fade and scroll-into-view are present (the capture shows no fade).
3. /bis 1440 and 390: the answer ("The list") starts about 970 px down at 1440, and after three cards at 390. This is the same placement as stage 1; flag only because the new heads did not buy it back. Consider stat weights and Play it below the list at 390.
4. /bis 390: the h1 wraps "FURY WARRIOR · 20 to 29 ·" then "Alliance" on a separate line. Put the range and faction on the meta line, not in the h1 row.
5. /planner 1440: a dead hole of about 200 px under the left column's stat tiles, beside the tall right rail, before the trees. Independent columns per §3.6 allow this, but the tools should start higher; move the rail's gear card below the trees or let the trees start at the head.
6. /home 1440 and 390: an empty gap of about 90 px between "Your upgrades" and "Look up a character".
7. /guide 1440: the page head and build cards start at x=96, but the prose column starts at x=208 with the TOC at the right. The left edge breaks between the head and the body. Align the prose to the column edge or centre the head's job line with it.
8. /tiers 1440: the tier table is bare on the page ground with its own tier gutter bands, not in a Card (spec §3.7 "table in a card"). It is the one data surface that does not read as the account material.
9. /rankings: the Characters | Guilds control is an outline segmented control in a card head, not the shared TabStrip. Fine as a control, but it is the fifth tab-like style on the site.
10. /report 390: the h1 truncates with an ellipsis ("SANGUINE DEPTHS, FIXTU..."). Let it wrap to two lines.
11. /guild roster 390: "WAITING FOR APPROVAL" wraps to three lines beside Approve all, and the roster card's meta text is right-aligned and cramped under the title. Stack the aside under the title below 480 px.
12. /guild roster 390: the leader row shows a tall empty block under the stat line (fixture data may cause it; verify with real rows).

## Checked and fine

No navy band on any route. Account, guild overview, settings, raids, premium, setup, about, guides index, home (hero sky stays, by §3.7) and the 2000 captures hold the column and centre. Card heads, Tags (SUPPORTER, CONNECTED, CURRENT, MAIN), ActionLinks ("See all", "Open the planner") and icon tiles match /account.
