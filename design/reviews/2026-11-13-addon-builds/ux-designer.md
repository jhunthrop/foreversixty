# ux-designer verdict, addon Builds pane and panel Talents row (2026-11-13)

Verdict: FIX. Structure, anatomy, labels, Follow buttons and the panel row match the spec. The why line, the Following mark, the row ground and the tooltip do not yet reach the planner rail.

Must fix
1. The why does not say why (all four builds-* boards). One 12 px line cuts it at about 55 characters, so BM reads "all five points of Efficie...", "The raid build moves two points out of Efficiency and the four i...". The first-sentence rule also picks non-whys: Holy "The rest is utility.", Restoration "It is built for..." (no antecedent), Protection "Toughness (10% armor from items) and a point of Divine Intelle..." (fragment), Retribution "never casts it" (no antecedent). Several repeat the row label ("keep the level 60 build:"). Change: write a stand-alone `why` of 70 characters or fewer per build (guide frontmatter or a sentence that passes a "starts with a subject, no pronoun lead, under 70" check), and never cut a why mid-clause. Spec 3 is amended accordingly.
2. Following is only a gold word (builds-hunter-42, hunter-60-raid, druid-24). Add a persistent marker on the followed row: a 2 px gold left edge and a gold ground at about 8% alpha, label stays body colour. The owner asked for a clear Following state; one word among six rows is not it. The cursor edge (white) stays distinct.
3. Row ground and edge have no inner padding (builds-row-states: Hover, Cursor, Held). The ground starts at the label's x, and the cursor bar touches the "L" of "Level 60 build". Extend the ground 6 px left of the text (and keep the right edge under the Follow button), edge 2 px, 4 px gap to text.
4. Row-states board omits states the spec lists: the Follow button hover (edge brightens) looks identical to default, the pressed 1 px label drop is not visible, and Held ground is almost the default ground. Re-render with a hovered button, a pressed button, and a Held ground at least 6% lighter or darker than default. Add a keyboard-cursor board that shows the hint `Enter to follow`; no board shows it.
5. Hover tooltip (row-states) is a flat white wall: label, split and why in one colour and size. Make the label the title (white, 14 px), the split muted grey, the why in normal body text with 4 px above, and add the last line `Click to follow` (or `Following` on the followed row).
6. builds-druid-24-site: `and 3 more` sits 2 px above the paste box (needs 6 px under it, spec 1.2) and is not aligned with the site row's text column. The footer hint (`Tab to switch`) is missing on this board while every other board has it; the dialog bottom must be the same on every board.

Polish
7. Spec name and build label are the same colour, size and weight, with no gap between them (hunter-42: "Beast Mastery ... / Level 60 build"). Give the spec line 2 px more air below it, or set the build label at 13 px. Four muted lines a row read as a wall; the label must be the first thing the eye lands on in each row.
8. The site row's 32 px crest (druid board) is heavier than the own-build icons (16 px) and the heading crest (20 px). Secondary block, so use 20 px.
9. panel-talents-42 (fixed): line 1 to the buttons' top is about 2 px. Take Panel.NAME_DROP to 4 px or lift line 1 by 1 px. Verify that "Change" and "Talents" survive a redraw (the blank button is fixed in the render; the test is what holds it).

Checklist: one question answered without scroll at 420 wide: yes (hunter 42 is about 560 px tall, taller than a 480 target; acceptable per spec 1.2). Both kinds per spec: yes. One-click Follow: yes. Paste secondary and last: yes. Ringed crests: yes. Gold used once: yes (but see 2). Whys readable: no. States complete: no. Tooltip hierarchy: no.

## Round 2

Verdict: SHIP. All six round 1 must-fixes verified on the re-rendered boards.
1. Whys: complete, stand-alone, none cut (hunter-42, druid-24). Fixed.
2. Following: 2 px gold edge plus gold ground, 6 px left of text, label stays body colour. Fixed.
3. Row ground and edge sit 6 px left of the text on Hover, Cursor, Held, Following. Fixed.
4. Held ground is clearly darker than default; Follow hover is lighter; `Enter to follow` shows on hunter-42 (cursor on the raid row). Fixed.
5. Tooltip: white title, grey split, body why, grey `Click to follow`. Fixed.
6. Druid: `and 3 more` aligns with the site text column and has air above the paste box; site crest is 20 px. `Tab to switch` hidden with own builds is the amended spec. Fixed.
Panel row: 4 px gap, Change and Talents render at 340 (left board still shows the blank button; it is labelled as the replayed capture).

Polish (not blocking)
- No pressed-state render (1 px label drop); add it to builds-row-states.
- "Following" and the why line run flush to the ground's right edge; give them 6 px right padding like the left.
- Spec line and build label still sit tight; 2 px more air under the spec line.
