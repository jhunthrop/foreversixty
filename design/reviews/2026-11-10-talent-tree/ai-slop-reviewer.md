Verdict: PASS

Scope: the shared talent tree (TalentTrees, TalentTree, TalentButton, talent-tree.css, GameTooltip) on planner, guide At a glance and Leveling, sim Talents and /sim/talents. Screenshots in after/.

First impression: reads as made by a player. Client art behind the trees, real talent icons, one ring colour per state, the in-game tooltip layout. None of tells 1, 2, 5, 8, 12, 13, 15 appear on this component.

Visual check
- No glow, gradient, blur or halo. The only box-shadows are the inset state ring and a 2px bg-coloured gap behind the rank badge.
- Rank badge is a pill on every learned or available talent (tell 17), but it carries state (3/5, green when spendable, gold when maxed). Justified, keep.
- Trees sit in one panel with column rules, not cards in cards.

Findings, most important first (fixes, none blocking)
1. Tell 22: planner rail repeats "Levels 20 to 29 build" three times (head, "Load the levels 20 to 29 build", tooltip line "Levels 20 to 29 build: 0 of 5"). Say it once in the head; button reads "Load", tooltip line reads "20 to 29: 0 of 5".
2. Tell 21/22: "Differs by 5 points: this build takes Booming Voice 5/5, ... takes Unbridled Wrath 5/5" plus "+9 beyond the build" is a paragraph restating what the Gains and Drops lists show. Use the lists or one short line.
3. Tell 21: planner lead ("Plan every point up to level 60 in the order you take them, check it against the guide builds, then copy it into the game or share the link.") and "Opened from a link. Point order: the main tree first." explain the obvious to a veteran. Cut.
4. Tell 22: sim rows "Your current build 0/20/0" and "Guide build, level 60 . 17/32/2" repeat the points already printed on each tree head. Drop the slash triple.
5. Tell 13 (borderline, passes): "Your saved builds / No saved builds" says the same thing twice; it has a real "Add a build" action, so keep the heading and action, drop the grey line.
6. Data bug visible in the hover capture: Booming Voice reads "...Rage cost by of your Shouts and Howls" (a substituted value is missing). Fix in the description renderer or source before launch; a veteran will see it.
7. Phone (guide-glance-390): three trees at about 18px with badges larger than the icons; nothing legible. Not a tell, but reads as shrunk desktop. Hide badges under a size threshold or stack the trees.

Authentic, keep
- Client tree art behind each tree; prerequisite arrows go gold when met.
- One cell language everywhere: ring scheme, badge, tooltip with icon, "Rank 5/5", gold description. Reads like an addon.
- Real talent names and counts (20 of 51, 17/32/2); gold dot for planned past level; Reset beside the points.
- "Client build 1.60.1.70790" footer is a detail only this product has.
- Hover, focus and tap behave the same on every surface.
