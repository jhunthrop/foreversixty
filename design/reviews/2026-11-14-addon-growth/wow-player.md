# wow-player verdict, addon growth lane, 2026-11-14

Verdict: FIX (four small must-fixes, all in text; the TOC change itself is fine to ship).

## Seen
- addons-list-tooltip.png (before vs after), ForeverSixty.toc, 2026-11-14-growth-listing.md, curseforge-description.md.
- The capture is a mock of the AddOns list, not a game screenshot. It reads like the client's tooltip: gold title, white notes, three lines of wrapping. Fine.

## The player's question
"What does this addon do for me?" The new Notes answer it in the first words ("Gear upgrades, BiS, talents and rotation"). The old "Your build, in game. Import your character into the planner" answered a different question (how do I set it up). The new one is better.

## Must-fix
1. Summary says "A green arrow on every bag item, quest reward and roll that beats what you wear". Read plainly, every bag item gets an arrow. Only the ones that beat what you wear do. Say "A green arrow on any bag item, quest reward or roll that beats what you wear".
2. "with the gain" has no unit. The live line said "gain in DPS", and the description says "+2.6 DPS" (HPS for healers). Write "with the gain in DPS" or "+2.6 DPS" in the summary. A bare "gain" is a number without a unit.
3. "BiS for each slot and what to press, for every Forever class and spec." Check that Rotation really covers every spec, tanks and healers included. The approved description only promises "What to press at your level" and gives "Healers read HPS" for tooltips. If any spec has no rotation, cut "every Forever class and spec" to what is true.
4. The listing file's "Live" cells (summary, categories "Tooltip, Guild, Map & Minimap") do not match curseforge-description.md (summary "A green arrow on the bag item... with the gain in DPS", categories "Boss Encounters, Class, Guild, Tooltip, Minimap, Quests & Leveling, Combat"). One of them is stale. Before pasting into the dashboards, say which categories are actually set today, and keep the "gain in DPS" wording from the approved file.

## Should
- Notes are two fragments and a command: "A green arrow on gear that beats yours, your next talent marked. /fs opens it." Smoother: "A green arrow on gear that beats yours and your next talent marked. /fs opens the window." Also "rotation" in Notes is the weakest-evidenced word; keep it only if item 3 holds.
- `## Category: Class` is not something I can confirm the Forever (1.60) client reads. Retail 11.x reads it. If Forever ignores it, it is harmless, but do not claim the grouped list works until someone sees it in game. The capture cannot prove this.
- Name: "Forever Sixty: Gear Upgrades, BiS & Talents" is the one I would click, because it says what it does and BiS is the word I type. Keep the proposed one. Alternative 2 ("Class Helper") is generic and reads like a placeholder. The ampersand and the colon are normal on CurseForge.
- Dropping Map & Minimap is right. Dropping "Combat" and "Boss Encounters" from the proposal is also right, since the addon is not a boss mod or a combat tool. Keep Guild: raid sign-ups are a real feature.
- Summary length: it sits near CurseForge's 200-character cap once edited. Count it after fixes 1 and 2.
- Summary leaves out raid night (sign-ups, jobs) and the Forever Score. That is fine for a one-liner, as long as the description keeps them, which it does.

## Claims check against curseforge-description.md
- Green arrow on bag items, quest rewards, rolls: true (Bags, quest rewards and rolls sections).
- Next talent marked: true (talent window).
- BiS per slot: true (Gear tab, "Best in slot" per slot).
- What to press: true for the Rotation tab, scope per spec unverified (item 3).
- "Install with the CurseForge or Wago app and Forever Sixty Data comes with it": true, it is a required dependency.
- Bag arrow works in the default bags and five named bag addons; "with any other bag addon, /fs diag says". The summary does not overclaim here.

## Round 2: SHIP

- Must-fixes 1 to 4 are closed: "any bag item", "gain in DPS", the "every class and spec" claim is gone, Live and approved cells are labelled with their source.
- Notes now read cleanly and wrap to four lines in the capture without clipping; the listing quotes the same text as the TOC.
- Still true, not blocking: `## Category: Class` is unseen in the Forever client (the listing says so). Rotation for tank and healer specs should be spot-checked in game before the store line goes live. "BiS and Talents" is the name I would click.
