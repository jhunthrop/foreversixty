# ux-designer verdict: addon default-build labels (2026-11-13)

Verdict: SHIP.

- Words: "Level 60 build", "Raid build", "Level 60 and raid build" match web/src/lib/planner/copy.ts in all three renders.
- 420 wide: no wrap, no collision. The longest label ("Level 60 and raid build", Survival and Holy/Protection rows) ends near x=386, well clear of the Follow button at x=754.
- "Level 60 and raid build" reads cleanly as one row for both uses, with a single Follow button, and sits at the same weight as its siblings.
- The Following state (gold bar, "Following" at right) does not crowd the label in either hunter render.
- Polish only, no fix needed: in the level 42 render the Beast Mastery "Level 60 build" row is the followed one, which is correct but may read oddly at level 42. The label is the web's word, so keep it.
