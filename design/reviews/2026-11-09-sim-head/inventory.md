# Page head inventory (2026-11-09, unified head lane)

The owner's report: on /sim (Go Hogs, Fury Warrior 35, a browser run of a spec no parse has
measured) and on /bis (Felsnap, Beast Mastery Hunter 34), the answer block at the right stacked five
to seven rows, and the shared row's `align-items: end` pushed the title to the bottom of the band.
The root cause was shared, so the fix is one contract in `web/src/components/ui/PageHeader.svelte`.

## The contract (PageHeader)

- One top padding and one title row everywhere. The crest, title and meta are **top-aligned** with
  the answer. The old `align-items: end` and home's 120 px hero floor are gone, and so is the unused
  `minRowHeight`.
- **Answer** = a Figure (or a who-line), at most one short line under it, held to a 26rem measure.
- **Actions** sit beside the answer on its last line from 900 px, so they never add a row. On a
  phone they go under it.
- **notice** slot: one line under the row for a state the page reports (the finished server sim).
- **crumb** slot: a way back (a raid's guild) on the band above the title, never above the band on
  the page ground.
- **pending**: the title isn't known yet, so a skeleton holds the row and the h1 is for screen
  readers only. This replaces the hand-built loading heads.

Paths are relative to `web/src/`. Slots: C = crest, T = title, M = meta, L = lead, A = answer,
X = actions, Tb = tabs, N = notice, K = crumb.

| Page / state | Component | Slots | What was wrong | Now |
|---|---|---|---|---|
| `/sim` with a character (live result) | sim/SimCharacterHead via sim/SimView | C T M A X Tb | SimAnswer held 152/169 px and stacked a label, figure, error, a 1000 px fidelity sentence, then Save/Top Gear. The title was pushed to the bottom and set in cream while /character used the class colour | Answer: label, figure, error with the "Not yet" Tag on the same line (its sentence is the Tag's tooltip, it's read to screen readers, and the rotation card states it in full). Save and Top Gear sit beside it. Title in the class colour |
| `/sim/<id>` run | sim/SavedSim | C T M A X | Same stack, with Open in sim and Copy link under the figure | Same contract; the buttons sit beside the figure |
| `/sim/<id>` weights/bulk | sim/SavedKind | C T M A X | Buttons stacked; a `pt-[22px] md:pt-8` wrapper pushed the band 22–32 px off the bar | Buttons beside; band flush |
| `/sim/<id>` loading | lib/sim/skeleton.ts | bespoke HTML | Not on the band; old 169 px answer geometry | Draws the band and the new geometry (a static twin, since it ships as HTML before the island) |
| `/sim` landing, tools (gear, drops, talents, weights, specs) | sim/SimLandingPage, sim/SimToolHead | T M Tb | none | unchanged |
| `/bis` landing (signed in) | pages/bis/index.astro + bis/BisLandingAnswer (ui/CharacterAnswer) | T M A N | CharacterAnswer with a figure reserved 224 px: label, who, figure, caption, meter, note sentence, 44 px button. "Best in slot" sat at the bottom-left under a big empty area | Label for screen readers only. Who-line over figure and Meter on one line, the note in the figure's tooltip, the button beside. The finished-sim line moved into the head's notice slot |
| `/bis/<class>/<spec>` | bis/BisSpecPage | none (title in PaperdollFrame) | No PageHeader: the paperdoll-led page (owner 2026-11-05) titles the doll frame | Unchanged on purpose: the doll *is* the head there. Flagged for the owner |
| `/` signed out | pages/index.astro | T X (hero) | 120 px hero floor with an end-aligned title | Floor gone; title top-aligned |
| `/` signed in | home/HomeCharacter | C T M A (hero) | About 100 px of dead art between the nav and the name row. The answer read "122.1 DPS" with no context. `.home-answer` forced its own alignment at every width | Floor gone. The figure has one linked line under it ("Simmed on import · 14 h ago ›", to that result). Alignment follows the head |
| `/` loading | home/HomeSignedIn | bespoke | Hand-built head off the band (`min-h-[118px]`) | PageHeader `pending` on the hero |
| `/character` ready | Character | C T M A | none beyond the shared row | unchanged |
| `/character` loading | Character (`headSkeleton`) | bespoke | Hand-built head off the band | PageHeader with the name from the address (`pending` only when there is none) |
| `/guides`, `/guides/<class>`, `/guides/<class>/<spec>`, `/404` | guides/GuideAnswer, 404 (ui/CharacterAnswer) | C T M A | 3–4 rows: a "Your character" label, then who and links. The title was end-aligned | Label for screen readers only; who and links; top-aligned |
| `/tiers` | tiers/TierPage | T M | none | unchanged |
| `/planner` | planner/Planner | C T M A | none beyond the shared row (points figure, one line, Sim link) | top-aligned |
| `/ratings` | score/RatingsPage | T M A Tb | A hand-built 36 px score and a 360 px spacer span in place of Figure | Figure (label, value, unit, tier colour); spacer deleted |
| `/premium` | pages/premium.astro | T L A | none beyond the shared row | top-aligned |
| `/premium/checkout` | checkout/CheckoutHead | T A | up to 4 rows (figure, Tag, size, trial, Change plan) | top-aligned; the plan card's rows are the plan's facts |
| `/setup`, `/account` | setup/SetupPage, account/AccountHead | T M A | SetupAnswer plus CompanionDownload, about 4 rows | top-aligned |
| `/logs`, `/reports`, `/connect`, `/login`, `/characters`, `/changelog`, legal, `/app`, admin | various | T M L (X/Tb) | none | unchanged |
| report detail | report/ReportHead, ReportSkeleton | T M A Tb | none | unchanged |
| `/rankings` | Rankings | T M | none | unchanged |
| guild page | Guild | C T M Tb | Role lines, sign-in prompt and claim row stack under the band in a nested `<header>` | Unchanged in this lane (they are body rows on the page ground); flagged |
| guild raid planner / raid night report | guild/planner/RaidPlanner, guild/report/RaidNightReport | K C T M (A/X) Tb | The back link sat above the band on the earth ground, a strip between the bar and the band | Back link in the crumb slot on the band |
| guild raid loading | GuildRaid | bespoke (GuildBandFrame) | Hand-built head off the band | PageHeader: crumb, the raid's kind as the title, skeleton meta, tabs |
| guild raid failed / not visible | GuildRaid | C T | `md:pt-7` pushed the band off the bar | Band flush when there is a head |
| guild claim (and its chunk loading) | GuildClaim, GuildShell | K T | Back link above the band; `md:pt-7` strip | Crumb on the band; flush |
| guild settings (and chunk loading) | GuildSettings, GuildShell | T M | `md:pt-7` strip between the bar and the band | flush |
| guild member profile | guild/recruit/GuildMemberProfile | C T M A | Grid of N figures | top-aligned (figures are the answer) |

## CSS outside PageHeader that targets the head

| Where | What | Status |
|---|---|---|
| ui/CharacterAnswer `.ui-page-header-answer:has(> .ui-cans…)` | hides the answer column signed out | kept (it still resolves: the answer column sits inside `.ui-page-header-side`) |
| ui/CharacterAnswer `.ui-cans[data-figure]{min-height:224px}` | the /bis reserve | 156 px below 900 and 96 px from 900, at the compact shape |
| sim/SimAnswer `min-height:152/169px; min-width:360px` | answer floor | 112 px (the running state: label, figure, caption, Meter), min width 240 |
| home/HomeCharacter `.home-answer` | alignment at every width | inherits the head's |
| account/AccountHead, setup/SetupAnswer, pages/premium.astro | reserve and hide rules on `head-answer` | kept; they still resolve |
| guild/GuildRoute.astro, pages/guild/[...path].astro | phone inset on `.ui-page-header` | kept |
| checkout/CheckoutHead `.is-pending h1` | hides the title until the step is read | kept |
| report/ReportView `.ui-tsm-strip` mask | phone tab fade | kept |

## Later rounds (owner rulings and reviewer findings)

- **Lead sentences removed from heads:** /changelog, Content pages (privacy, terms, refunds, about, sources), /premium and /404. Each head is now its title plus its Updated stamp. The /404 `job` and the setup subline are deleted from the copy source.
- **/sim:** a result has no lane or saved label. The band reads "± 0.3 DPS" (no "(95%)"). The calibration Tag reads "No parses yet" or "Calibrating" and links the spec's row on /sim/specs. An undated addon export has no "Gear from" line. The saved head reads "Saved 4 Nov" and shows the year only when it isn't this year.
- **Home:** the import sim's line is "Sim 8 h ago ›", linking that result.
- **/setup:** the download is the head's actions, beside the answer. The import stamp is said once, in the Sign in row.
- **/bis:** the short line is the stale-sync caption after an old sync, otherwise the no-pick note.
- **Below 900:** the head's answer and actions are the row's own items, so a hidden answer leaves no gap.

## The "<X> is the <Y>:" sweep (owner 2026-11-09)

These were removed from content and card surfaces:
- The nine class intro paragraphs in `web/src/content/guides/*/index.md`. The class landing now fails the build on any intro paragraph (`lib/guides/class-landing.ts`). They were:
  - druid: "The druid is the class that fills any role..."
  - hunter: "Hunter is a pure damage class in Forever..."
  - mage: "Mage is pure ranged damage with the best toolkit..."
  - paladin: "Paladin is the plate hybrid..."
  - priest: "Two healing specs and one shadow caster..."
  - rogue: "Rogue is still energy, combo points and stealth..."
  - shaman: "Shaman is the totem class..."
  - warlock: "Warlock is a ranged caster that pays for mana with health..."
  - warrior: "Warrior is the pure melee class..."
- The spec card lines on every class page, such as "Fury is the dual-wield warrior: ...". These were read from each guide's Overview by `cardLineFor` and `specCardLine`, and both are deleted. Cards are now the icon, the name, the role and "Read the guide ›".

These were kept on purpose. Each spec guide's own Overview opening sentence is in-guide instructional prose, which the owner ruling says to keep inside the guide body: balance, beast-mastery, survival, retribution, holy (priest), assassination, combat, subtlety, elemental, enhancement, restoration (shaman), arms, fury and protection.

Other "X is the Y:" hits are code comments and in-rotation instructions (rotation-lines.json, race-table), not summary blurbs.

The pattern is tell 23 in `.claude/agents/ai-slop-reviewer.md`.
