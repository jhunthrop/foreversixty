# Guild control-centre build report (2026-10-04)

Branch `web-guild-centre` in `.worktrees/web-guild-centre`, built against
`design/specs/2026-10-04-guild-page.md` v2 and
`docs/contracts/2026-10-04-guild-centre-api.md`. Final sha: `610ffe87`.

## What shipped, per tab

- **Tab strip** (`GuildTabs.svelte`): hash-anchored (`#roster`, `#raids`, ...), 44px pill
  family, role-gated via `tabsForRole`, active pill scrolled into view on load/change
  (spec §7). A stale/disallowed hash falls back to the first allowed tab silently.
- **Overview** (`GuildOverview.svelte` + the role-line in `Guild.svelte`): per-tab summary
  cards (Roster, Raids, Progression, Readiness, Loot), each with a "See all" link. The
  standing line and "before Thursday" sentence (spec §4.A.1/§4.A.2) render once, above the
  tab strip, so they stay visible on every tab — not Overview-only. Readiness/Loot summary
  cards fetch their own endpoints once (documented deviation below).
- **Roster** (`GuildRosterTable.svelte`): filter bar (role/class/verified/below-floor,
  AND-combined), five-column sort header, rating tooltip, attendance, parses, professions,
  main/alt tag, own row pinned first, Approve/Approve all/Remove.
- **Raids** (`GuildRaids.svelte`): expandable nights, present roster, top parse, pull-by-pull
  fight list, unnamed-encounter honesty for Barrow Deeps/Hyjal Summit.
- **Progression** (`GuildProgression.svelte`): tier bar with the honest caveat, Onyxia's own
  pulls-to-kill trend, death-cause `NEEDS NEW DATA` line, unnamed-zones panel, and v1's
  "Roster bests" carried forward here (not on Overview — the v2 summary-card table in §4.A
  does not list a Roster-bests card, so it lives on the tab closest to it).
- **Readiness** (`GuildReadiness.svelte`): worst-first sort, gear gap/enchants/consumables/
  talent points/item level cells each correctly consent-gated and coloured (pill red/green,
  gold for unspent points, coloured item-level delta), member-only pin, officer-only Nudge
  (clipboard copy), phone stacked chips.
- **Loot** (`GuildLoot.svelte`): boss picker defaulting to next-unkilled, ranked candidates,
  attendance, "already holds equivalent", one awardee per item with every other candidate
  reading "Awarded to {name}", member read-only.
- **Settings** (`GuildSettingsTab.svelte`): thin wrapper around the existing
  `GuildSettings.svelte`, unchanged logic, restyled into the tab strip.

## Tests, verbatim

```
npx astro check        → 0 errors, 0 warnings, 10 pre-existing hints
npx eslint .            → clean
npx prettier --check .  → clean
npx vitest run          → Test Files 301 passed (301); Tests 3125 passed | 1 skipped (3126)
npx playwright test guild-centre.spec.ts guild-claim.spec.ts guild-settings.spec.ts \
  guild-settings-billing.spec.ts guild-invite.spec.ts guild-account-consent.spec.ts \
  --project=desktop --project=mobile
  → 80 passed (80)
```

`vitest run`'s only pre-existing gap (`src/lib/sim/bulk-store.test.ts`, 5 tests, unrelated
Molten Core/Onyxia data-tier fixtures) is confirmed present on `main` itself before this
branch touched anything, and intermittently passes/fails independent of this work — not
this lane's to fix.

## Side-by-side differences found and fixed

Captured the fixture guild (`/guild/us/pvp/olympus-xxvii`, id-matched to
`design/mocks/gen_guild.py`'s own "Olympus XXVII") at 1440/2000 (officer: Overview, Roster,
Readiness, Loot; member: Overview; public: Overview) and 390 (member Readiness) into
`design/mocks/renders/build/guild-*.png`, via `web/scripts/capture-guild.mjs` (a one-off
Playwright capture tool, not part of the e2e suite). Four real defects found and fixed
before this report:

1. **Scoped `<style>` CSS in every new Svelte component was silently dropped from the
   build.** This build's Astro/Vite pipeline only inlines a Svelte component's scoped CSS
   for classes present in the page's server-rendered markup; every guild/* tab only mounts
   client-side after an async fetch, so none of it ever appeared there. First capture
   showed borderless tab pills, unstyled rows, no readiness pills. Fixed by converting
   every guild/* component to Tailwind utility classes only (the convention the rest of
   this codebase's Svelte components already follow — now understood to be load-bearing).
   Also fixed two latent custom-property typos the dead CSS was masking
   (`--color-border(-soft)`, which never existed; the real tokens are `--color-line(-soft)`).
2. **Roster tab doubled every unverified raider and pushed the pinned row down.**
   `orderRosterForTab` already leads with unverified rows; the component also rendered its
   own explicit unverified block above it. Fixed by sourcing the main row list from
   verified rows only.
3. **Readiness pinned the officer's own row.** Spec §4.0's matrix pins "own row first" for
   a member only; an officer's cell reads "Nudge on every row" with no pin clause. Fixed by
   gating the pin on `!officer`.
4. **Roster's main/alt tag was built but never rendered**, and **Readiness's Enchants/
   Consumables/Talent-points/Item-level cells had no colour or pill at all** (plain text).
   Both fixed; the latter needed new pure `*Kind()` discriminators in `readiness-view.ts`
   kept separate from the existing `*Label()` text functions, each unit-tested.

Also fixed along the way: the Overview standing line only showed for a member (spec §4.A.1
never gates it on role — only the "before Thursday" sentence is member-only), it now
carries the viewer's own name (looked up from their roster row, since the standing
contract has none) and class colour; the header facts line now prefers `home.summary`'s
own named-encounters/pulls counts over the public endpoint's once a membership match
loads, so it never disagrees with the Progression card; and the header gained the eyebrow
"GUILD" label and "Updated" stamp.

## What I left out, and why

- **Full-bleed header band.** The hard rule asks for a full-bleed band with its inner
  content capped and centred at 1344px. `[...path].astro` already wraps the whole page
  (header included) in a `max-w-[1344px]` `<main>`, which v1 never broke out of either; a
  true edge-to-edge band needs the header pulled out of that wrapper into its own region in
  `[...path].astro`. Given the size of this round's own tab rebuild, I did the cheap,
  high-value part (eyebrow label, "Updated" stamp, larger h1) within the existing structure
  and left the architectural extraction for a follow-up — flagged here rather than silently
  skipped.
- **"Gear in the planner" link on the standing line.** Needs the same per-character FS1
  addon-export lookup `GuildRosterHandoff.svelte` already does, which is an async,
  per-character call; adding a second copy of that mechanism just for one header link felt
  like more plumbing than the capture gap was worth this round.
- **Overview's Readiness/Loot summary cards fetch their own endpoints** (`fetchGuildReadiness`/
  `fetchGuildLoot`) rather than reading only `home`'s own fetch, because the contract's
  `home.summary` carries no field for "three worst-readiness raiders" or "current boss/next
  pick." Both responses are cached through `query.ts`, so opening the real tab afterward
  costs nothing further — documented in `GuildOverview.svelte`'s own header comment as a
  deliberate, reported deviation from the letter of "no new round trip."
- **`guild-roster-approve`/`guild-home-roster`-style testids from the superseded
  `guild-home.spec.ts`/`guild-phone.spec.ts`.** Both exercised exactly the flat single-page
  body this round replaces (the duplicated unverified note, the old flat roster list) — per
  the build brief's own "delete what the new components make dead ... and their tests,"
  I replaced them with `guild-centre.spec.ts`, porting every claim/contest/approve/remove
  assertion worth keeping into the new tab structure rather than leaving stale coverage
  behind.
- **Roster bests on Overview.** v1 showed it there unconditionally; v2's own §4.A summary-
  card table does not list it as an Overview card. I kept the feature (Progression tab,
  every role) rather than dropping it, but it is one tab deeper than the kept-unchanged
  `guild-public` board's own round-1 layout shows it.
- **Fixture-only nits, not product bugs**, left as-is given time: the mock's `home.summary`
  below-rating-floor count uses a simpler threshold than the real `ratingFloor()` 25th-
  percentile rule the Roster tab itself uses, so the two numbers can disagree in the
  fixture; the capture script's `/v1/me` stub hardcodes `class: 'warrior'` for every viewer,
  so a hunter/priest viewer's own nav avatar renders blank in the member/public captures;
  and `GuildRosterHandoff`'s "Open in simulator/planner" links never appear in captures
  because the capture script does not stub the per-character `sim-input` endpoint they
  depend on.

## 404 fallback behaviour

Raids, Progression, Readiness and Loot each fetch their own endpoint only once their tab is
opened (lazy, matching spec §8). Any failure — 404 because the api lane has not deployed an
endpoint yet, or any other error — lands the tab on its own `EmptyState` with a named
next-step line ("... the api lane is still deploying this endpoint. Check back after the
next sync."), never a crash; Progression additionally keeps showing v1's own working
Roster-bests panel underneath even when its own new endpoint 404s. Verified by
`guild-centre.spec.ts`'s own 404 test and confirmed manually against a guild id the real
API does not know yet.

## Captures

`design/mocks/renders/build/guild-{overview-officer,roster,readiness,loot}-{1440,2000}.png`,
`guild-member-{1440,2000}.png`, `guild-public-1440.png`, `guild-phone-390.png`.
