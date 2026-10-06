# Persona review: the retail raider's first week with Forever Sixty

Reviewer persona: **Tessara**, 31, Fury Warrior (Human), joining a guild clearing the first
Forever tier (raids open 9 Dec 2026). Retail raider since Wrath, two tiers as a raid lead,
purple parses on Warcraft Logs, Raidbots Top Gear before every pull night, Wowhead tab always
open, WeakAuras/Details!/DBM/Pawn/RCLootCouncil/SimC addon installed on live. Read against the
repo at `/Users/jh/code/forever` as built: home (`web/src/pages/index.astro`), BiS
(`web/src/pages/bis/[class]/[spec].astro`), planner (`Planner.svelte`), simulator
(`web/src/pages/sim.astro`), logs (`web/src/pages/logs.astro`), rankings, guild centre
(`Guild.svelte`), the Fury guide (`content/guides/warrior/fury.md`), the addon spec
(`design/specs/2026-10-04-addon.md`) and README, the talent-search report
(`design/reviews/talent-search/paladin-retribution.md`), `docs/tenets.md`, `docs/STRATEGY.md`,
and renders under `design/mocks/renders/` (guides-spec.png, guild-overview-officer.png,
addon-tooltip.png, addon-overview.png). Also pulled live `foreversixty.gg/`, `/bis/warrior/fury`,
`/logs` to confirm the build matches git.

## 1. Who I am

Fury Warrior, Human, 31. Ran a Heroic guild two tiers on retail — I know what "ready for pull"
looks like on paper and in practice. Raidbots before every lockout, Logs open in a second
monitor during the raid, Pawn on every roll, RCLootCouncil on loot night, SimC double-checks
anything Raidbots looks odd on. I don't trust a number I can't trace to a sim or a log. I'm
allergic to "beta" labels that hide instead of explain, and I will screenshot a bad table into
guild chat.

## 2. First 10 minutes, as a diary

Land on `/`. The hero band says "Play your class better" with a nine-class crest picker right
under it — correct front door, not a feature pitch. Click the Warrior crest
(`home-class-picker-warrior`); it routes to `/guides/warrior`, not straight to Fury — one extra
click I didn't expect since I already know my spec, but minor.

Land on the guide (`guides-spec.png` is the real render). Top row: Build (talent icon grid),
Rotation·60 (Bloodrage → Death Wish → Battle Shout → Bloodthirst, each with a one-line
condition), Stat Priority (Attack power 0.15, Strength —, Agility —, Critical strike 0.14, Hit
0.03, Melee haste 1.00). I stop on "Not significant" beside Strength and Agility — odd for a
Fury weight table, flagged below (§4 of the sheet I'd hand accuracy). Below: Overview, Talents,
Rotation, Stat priority, Gear, Enchants, Races, Professions, Leveling, each stating plainly what's
beta-verified versus projected ("the beta only reaches level 30... not confirmed play"). One
honest sentence per claim, not a badge on every line — I like that.

"LOAD THIS BUILD" round-trips to `/planner?class=warrior` pre-filled. "SIM THIS BUILD" lands on
`/bis/warrior/fury`'s own band — fine, that's where I actually want to be: The List (slot / pick
/ runners-up), a weight rail in the same unit convention as the guide, a Play It rotation panel,
and a footer grouping "Where to get it" by source plus "New at this band." Best surface on the
site against my own stack — does what Sixty Upgrades does, with the real item tooltip and icon
right there, not a spreadsheet name.

Then I look for what I actually open every raid night: the simulator. `/sim` signed out shows an
example card and a sign-in prompt, with one honest scope line: "Damage specs are simulated;
healing and tanking specs are not simulated yet." Right way to say "not done." But I haven't run
MY character in these ten minutes — needs Battle.net sign-in plus a tracked character or addon
export. That's where a new visitor bounces to "come back once I'm in-game."

`/logs`: "Your reports" / "Recent public reports," live logging via companion, upload for an
already-recorded log. Correct shape, nothing to look at yet (no raids until December), but the
framing sentence told me what it's for before I had to guess.

Ten minutes in, the one thing I'd screenshot: both the home class grid and the guide route me
through an extra click before my own spec's numbers show. Everything else moved faster than
Wowhead does for the same question.

## 3. Versus my stack

| My tool | Forever Sixty surface | Verdict | Why |
|---|---|---|---|
| Warcraft Logs | `/logs`, `/rankings` | Same (unproven) | Right shape, free, no ads — but zero real raid data exists before 9 Dec, can't judge parsing yet |
| Raidbots Top Gear | `/sim` | Worse, today | Damage specs only; needs sign-in + a tracked character or export to run MY gear, where Raidbots takes a pasted string cold |
| Wowhead | `/guides/warrior/fury`, `/bis/...`, planner | Better | Both start Forever coverage at zero, but this guide already has a real tooltip, real icons, and a stated beta-vs-projection line Wowhead never bothers with |
| Pawn | Addon tooltip verdict (`design/specs/2026-10-04-addon.md` §4.1) | Better, once shipped | Pawn prints a dry "Upgrade"; this spec's verdict carries a delta, a sourced BiS row, and a live compare tooltip — ahead of Pawn's own idiom, not yet confirmed in-game |
| RCLootCouncil | Guild `Loot` tab (`guild-overview-officer.png`) | Same, partial | "8 class drops · next pick: Helm of Wrath" is a useful glance; can't tell from this if it replaces RCLC's vote flow or just reports on it |
| SimC addon | No in-game sim; browser sim is web-only | Worse | SimC sims instantly against my live gear in-client; this sim is a separate tab, and the addon has no network to trigger one itself |

## 4. The addon

Would I install it beside my 12? Yes, provisionally, for two things: the tooltip verdict (beats
Pawn on the spec) and the next-talent tracker/glow (replaces a WeakAura I'd otherwise hand-build
leveling). Those earn a slot day one.

What gets it disabled in a week: if the tooltip's BiS row doesn't draw on every hover the spec
promises — the spec itself calls this "the single biggest gap against tenet 1" and says it must
land before an in-game build counts as matching it (§4.1). Silence on a bag item where Pawn would
speak loses the slot by Thursday; the honest "UNRATED, no guess" state is fine, a missing verdict
where one should exist is not.

Two more trust checks before week one ends: the tier-indexing bug the spec itself caught ("tier
0" vs the real window's "Tier 1," §4.5.3) reads as a bug next to my own talent frame if it ships
unfixed; and the Guild officer block only points at the website ("Review on the site"), honest
about having no network, but it means tabbing out for approvals same as RCLC already does.

## 5. Would I use it

**Sometimes**, leaning yes once the simulator takes my own gear without a round-trip. The one
thing that decides it: pasting live gear/talents and getting a trustworthy DPS number as fast as
Raidbots does today. The guide/BiS/planner trio already beats Wowhead+Sixty-Upgrades. The sim is
the gate.

## 6. Stickiness

Daily: tracker + tooltip, passive, no site visit needed, if accurate. Weekly pre-launch:
guide/BiS re-reads as talents firm up. Per raid night (Dec 9+): logs then rankings the morning
after — the biggest lever, not live yet. Per tier: BiS refresh, loot tab, readiness checks
("4 checks failing" on `guild-overview-officer.png` is exactly what a raid lead opens before
reset). Stops me: one wrong sim number I catch by hand, or a silent tooltip where the spec says
one shouldn't be silent. Honest guess: **2–3 visits/week** in month 1 (pre-launch planning),
**5–7/week** by month 4 if logs and rankings become the raid-night habit Warcraft Logs already is.

## 7. Value over time

Compounds: my own logs and parses, my guild's roster/readiness/loot history, one data model
shared by planner/addon/sim (`docs/STRATEGY.md`'s own thesis) — the thing Logs' history does for
me today. Decays: guide prose and BiS picks once the community's memorized them, same as Wowhead.
The Ret Paladin talent-search finding (guide sims 17.1% behind a re-spent variant, "beyond
error," `design/reviews/talent-search/paladin-retribution.md`) is proof this already happens —
fine, as long as the site re-runs the search and republishes; a guide left stale after that is
the fastest way to lose me. By the first tier the site needs: logs as fast as Logs, a sim that
takes my real gear with no extra steps, and guides re-verified by search the way Ret already was.

## 8. Three fixes, ranked

1. **Let the sim take my real gear/talents in one step** (addon export or Battle.net import), no
   sign-in detour — the one surface where Raidbots still wins, and the gate on whether I use the
   site at all.
2. **Ship the addon tooltip's BiS row on every tracked-slot hover**, not just exact-worn matches
   — the spec's own "biggest gap against tenet 1," and what gets the addon uninstalled week one.
3. **Re-run talent search against every spec's guide before launch and republish the winner**, the
   way Ret Paladin already did — a guide 17% off a provably better build is the first thing a
   theorycrafter in my guild posts to Discord.
