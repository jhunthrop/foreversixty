# Forever Sixty guild page — experience spec

Author: ux-designer. References: `docs/tenets.md` (14 tenets, tenet 14 "player helper, not a
reference site" and tenet 11 "hierarchy over density" above all), `design/DESIGN-SYSTEM.md`
(principle 3 — no source pills/claim politics in a header; principle 7 — one crest language;
the Character row recipe; Pick row's "link, never restate" idiom), the two rebuilds this pass
already shipped as format precedent (`design/specs/2026-10-04-logs-landing.md` — the
"no-one-class, build a bespoke band" decision this page copies verbatim for the same reason;
`design/specs/2026-10-04-guides.md` — the full-bleed band / inner-1344px-centred-at-every-width
convention, confirmed fixed at 2000px by `git show 3ae56259`, "every band inner and body is now
centred at 1344px"), the two persona reviews that set this page's requirements
(`design/reviews/2026-10-04-addon-guild-leader.md`, `design/reviews/2026-10-04-addon-guild-member.md`
— every numbered item below cites one), the guild model
(`docs/superpowers/specs/2026-09-21-guild-membership-design.md`, all five security amendments —
`§2`, `§3`, `§4.1`'s rulings 8/9 are load-bearing for this spec's copy), and the page as it ships
today (`web/src/pages/guild/[...path].astro`, `web/src/components/{Guild,GuildShell,GuildStatus,
GuildRosterHandoff,GuildClaim,GuildSettings}.svelte`, `web/src/lib/guild/{api,copy,roster}.ts`,
`api/internal/guilds/home.go`, `api/internal/rankings/guilds.go`). A database seed
(`api/cmd/seedguild`, a concurrent, separate lane) is building the real 24-raider Molten
Core/Onyxia roster this spec's mock boards describe by hand; this spec does not touch `api/` and
makes no claim about that seed's own field values beyond what `home.go`/`guilds.go` already
return today.

Every pixel size, colour and copy string below is quoted verbatim from the design system,
computed from a real response shape with its field named, or cited as an existing component the
build lane must reuse, never invent (tenet 10).

**Data notes, binding on the seed too.** Forever launches 4 November 2026; raids open 9 December
2026 (`docs/tenets.md`'s own companion memory, confirmed by the coordinator) — no raid night, no
first-kill date and no "Updated" stamp anywhere on this page, the mock boards, or
`api/cmd/seedguild`'s own seed data may predate 9 December 2026 (tenet 8: nothing unverified, and
a Molten Core kill before the raid exists is not just unverified, it is impossible). This spec's
mock boards use four raid nights, all after that date: **2026-12-10** (the roster's earliest,
rolled-off night — Lucifron and Magmadar's first kills, no longer in the trailing "this week"
list), **2026-12-15** ("Molten Core, week 2," Gehennas/Garr/Baron Geddon's first kills),
**2026-12-17** ("Onyxia"), **2026-12-22** ("Molten Core, week 3," Shazzrah/Sulfuron Harbinger's
first kills) — the header's "Updated" stamp reads `Updated Dec 22`, the latest of the four. The
seed lane should treat these same four dates (or an equally post-9-December set) as its own
fixture's raid-night calendar, so a reviewer comparing the real seeded page against these boards
never finds a date mismatch that is not also a real content difference.

---

## 1. Purpose and the one sentence

Tenet 14 applies to this page exactly as it does to Logs, not literally: a guild is not one
class or one character, so the front door this page offers is "where do I, a member of this
guild, stand this week" (the member review's own verdict line), with "what is this guild doing"
underneath it for everyone else. The question changes with who is looking:

| Viewer | The one question | What answers it |
|---|---|---|
| Verified member | "Where do I stand against my guild's same-spec raiders this week, and what gear gap do I chase before the next pull?" (member review §5, verbatim) | One line in the header band naming the viewer's own rank among same-spec, same-class raiders by item level, with a link out to the planner — never a restated gear list (§4.A.1). |
| Unverified member | "Why can't I see more, and what do I do about it?" | One line, in the same header slot, stating the fact and the fix: an officer approves from the roster below (§4.A.1). |
| Officer/leader | Everything a verified member sees, plus: "Is this guild claimed, who is waiting on me, and how do people join?" (leader review §2, the only two things a leader review called genuinely useful) | A compact Officer tools strip, claim badge + pending-approvals count + invite link, directly under the header, never buried on a second page (§4.B). |
| Signed out / not a member | "What is this guild, and how do I get in?" | Identity, progression, roster bests and public reports, unchanged mechanism, plus one `SignInPrompt` in the same header slot a member's standing line would occupy (§4.A.1). |

**Must be visible with no scrolling**, 360px phone / 1280px desktop:

| | 360px phone | 1280px desktop |
|---|---|---|
| Must show | Eyebrow, h1 (guild name), the facts line, the one role-based line (standing / not-verified / sign-in) | All of the above, plus the Officer tools strip (officer viewer only) and the top of "This week's raid nights" |

The phone figure is the narrowest the site designs for (`web/lighthouserc.json`'s own mobile
emulation width). What does not fit is the next region down, never a second copy of what is
already on screen (tenet 11).

---

## 2. Reference

- **Warcraft Logs' own guild page** (progression per boss, a recent-reports list, a roster):
  the shape this page already has and keeps (`rankings/guilds.go`'s `Progression`/`RosterBest`/
  `GuildReport`) — what we do differently and better (tenet 1): every row is a real panel with
  a crest, a class colour and a quality-grade tooltip hand-off, never WCL's plain text table.
- **Raider.IO's guild profile** (a roster table with rank, item level, last-active, sortable
  columns): the exact shape §4.C's roster table borrows for the sortable-header idiom — what we
  do differently: Raider.IO shows every member's gear to any visitor; this page's roster is
  member-only and every gear figure is consent-gated per row (membership spec §3.2), honoured
  server-side in the query itself, never hidden client-side after being fetched.
- **The in-game guild roster frame** (`GuildView.lua`, read in full per this spec's own task):
  one rating number, name only, no class icon, no spec, no item level, no link anywhere — the
  leader review names this exact row a tenet-1/11 defect ("a list where the client... would show
  a panel"). This page's own roster row is the fix this spec commits to shipping on the web
  surface first (the addon's own `GuildView.lua` is out of this spec's scope — it is a web page
  spec — but §9 names the identical row-shape fix for whoever specs that surface next).
- **Why no `ArtPanel`+`ClassHeader`** (the same reasoning `design/specs/2026-10-04-logs-landing.md`
  §4.A gives for Logs, restated here because it is the same shape of gap): `ClassHeader`'s whole
  content — `SpecTabs`, `BandTabs`, `FactionToggle`, a race select — answers "which spec, which
  band, which faction, which race," and a guild has 24 of each at once, not one. Reusing the
  component with every control hidden would wear a class page's clothes over a page that is not
  about a class. This page gets its own band, `GuildHero.svelte` (new, its own design review per
  tenet 10 — the one genuinely new component this spec introduces), built from the same
  primitives (`.label` eyebrow, the 28×1px gold rule, Cinzel h1, the Secondary button token)
  `LogsHero.svelte` already proved out for exactly this "no one class" case.

---

## 3. What this rebuild keeps and what it replaces

**Kept, unchanged mechanism (restyled chrome, regrouped placement):**
- `fetchGuild`/`GuildPage` (public progression, roster-bests, all-reports) — kept verbatim.
- `fetchGuildHome`/`GuildHome` (reports, roster, claim) and every mutation in `lib/guild/api.ts`
  (`approveCharacter`, `removeCharacter`, `contestClaim`, `claimGuild`, `confirmClaim`,
  `releaseClaim`, `rotateInvite`, `updateGuildSettings`) — kept verbatim, no new endpoint.
- `orderRoster`/`unverifiedRosterCount` (`lib/guild/roster.ts`) — kept verbatim; the new sortable
  columns (§4.C) sort only *within* the verified group this function already produces, never
  ahead of an unverified row (tenet: an officer must see who is waiting before any other order).
- `GuildRosterHandoff.svelte`, `GuildStatus.svelte`, `GuildClaim.svelte`, `GuildSettings.svelte`,
  `GuildShell.svelte`'s routing, `SignInPrompt.svelte` — kept verbatim, same mounts.
- `guildHomeCopy`/`guildClaimCopy`/`guildSettingsCopy` — kept verbatim except where §6 marks a
  line new, moved, or removed.

**Replaced:**
- The bare `<h1 class="section-title">{guild.name}</h1>` plus its one-line facts paragraph
  (`Guild.svelte`'s current header) → the full-bleed `GuildHero` band (§4.A) carrying the same
  two facts plus the one role-based line every viewer gets (§4.A.1) — the one piece of this page
  that is still a plain title where the rest of the 2026-10-04 rebuild pass already ships a
  named header band.
- **The duplicated "not verified" line** (the owner's own screenshot: `notVerifiedNote` in the
  header's action slot *and* `reportsUnverifiedNote` again inside the reports panel, both saying
  the same fact in different words) → **one** line, in the header's role-based slot only
  (§4.A.1); the reports panel's own copy of it is deleted, not restyled.
- The flat `<ul>` roster list (`Guild.svelte`'s current `data-testid="guild-home-roster"` markup:
  a wrapped flex row per member, no table, no sort) → a real sortable table, panel-boxed, with a
  crest per row (§4.C) — the single clearest "plain list where a panel was expected" defect on
  this page (tenet 1), and the exact row the leader review calls out against `GuildView.lua`'s
  own name-and-rating-only row.
- The bare "Claim this guild" text link in the reports panel's corner → a standalone Officer
  tools strip (§4.B), claim badge + pending-approvals count + invite row, so an officer's two
  genuinely useful actions (leader review §2) are visible without a second page load.
- Nothing today groups the unverified roster rows with one bulk action — officers approve one
  row at a time. §4.C adds **Approve all**, the member review's and leader review's shared ask
  (member review §4: officer plumbing should not cost the raid-night-to-raid-night member a
  second look; leader review §3: "nobody renders it per-raider" applies just as much to
  approving twenty people one click at a time).

---

## 4. Region-by-region spec

Design tokens cited are from `design/DESIGN-SYSTEM.md` unless noted. Page gutter 48px desktop /
18px phone, content max width 1344px, section gap 32px desktop / 22px phone, **centred at every
viewport width including 2000px** (the guides-page centring fix, `git show 3ae56259`, applies
identically here — the band's own inner content and every panel below it share one `margin:0
auto` wrapper, never a left-flush container on a wide screen). No sky band, no `ArtPanel` tree
art (principle 5: atmosphere lives in the header only, and a guild has no one class to paint —
same ruling Logs and the guide pages' own `ArtPanel` fallback already state). The band ships flat
`--bg-raised`, identical to Logs' and the guide pages' own fallback, until an art lane exists for
a guild-level banner — not re-litigated here (§9).

### 4.A Header band — new, full-bleed, `GuildHero.svelte`

- **Shows:** eyebrow, h1 (guild name), an Updated stamp, the facts line, one role-based line.
- **Copy, verbatim:**
  - Eyebrow: `Guild` — `.label` token, gold `#e5b955`, 28×1px gold rule before it, identical
    token to every other page's own-name eyebrow (`Logs`, `Guides`).
  - h1: `{guild.name}` — Cinzel 700, 22px, `--text-strong` `#f2eee4` (no class colour: a guild is
    not one class, same reasoning Logs' own report-title h1 states for "a report has many
    characters in it").
  - Updated stamp, 12px `--text-muted` `#9a9484`: `Updated {day(mostRecentReport.created_at)}` —
    read off whichever is newest between `home.reports[0]` (member) and `data.reports[0]`
    (public); when neither exists, the stamp is omitted entirely, never a fabricated date
    (tenet 8) — this is the one page in the 2026-10-04 pass whose Updated stamp can legitimately
    be absent, since an unraided guild has nothing to date yet.
  - Facts line, 14px mono for numbers, `--text-muted` otherwise: `{rulesetLabel} {region} ·
    {killed} bosses down · {pulls} pulls` — the exact fragment `Guild.svelte` already renders
    (`rulesetLabel(resolved.ruleset)`, `resolved.region.toUpperCase()`, `killed`, `pulls`),
    reused unchanged so the band and the Progression panel below never disagree on the count.
  - Role-based line — see §4.A.1.
- **Component:** new `GuildHero.svelte`, mounted `client:idle` on `guild/[...path].astro` the
  same way `LogsHero.svelte` is mounted on `logs.astro` — above the fold, names the page's LCP
  candidate (the h1), so its first paint must not wait on the role-based line's own fetch
  (§4.A.1, §8). `Guild.svelte`'s own `data`/`home`/`myGuildMembership` state already contains
  everything this component reads; it takes them as props rather than fetching a second time.
- **Data:** `GuildPage` (existing, `fetchGuild`), `GuildHome` (existing, `fetchGuildHome`,
  member/officer only), `Me.guilds[]` match (existing, `fetchMeOnce`) — no new endpoint.

#### 4.A.1 The role-based line — answers both reviews' §5 verdicts directly

One line, four mutually exclusive branches, in the exact slot the header's facts line sits
above. Never two of these render at once (the literal fix for the duplicated "not verified"
line, §3):

1. **Verified member, `home.roster` contains a row matching one of `myCharacterKeys`:** rank the
   viewer's own row against every other `home.roster` row sharing the **same `class` and
   `spec`** (gear/gear_bags consent only — a `roster`-consent peer has no `item_level` to rank
   by and is excluded from the count, never shown as a false tie), ordered by `item_level`
   descending, ties broken by name for a stable order. Copy, computed:
   `{row.name} · {rank} of {total} {spec} {class}s this week by item level ({item_level} ilvl)` ·
   `Gear in the planner` (the existing `rosterPlannerHref`/`GuildRosterHandoff` link, reused
   unchanged — **never** a restated gear list or a computed gap number on this page: BiS and the
   planner already own that job, tenet 11's "link, never restate," the same rule the one-job-per-
   page ruling already states site-wide). When the viewer's own `item_level` is undefined (their
   own consent is `roster`), the line drops the rank clause and reads only
   `{row.name} · {spec} {class} · Set your gear consent to see where you stand` → `/account`
   (the existing "My guilds" consent control, §6).
2. **Verified member, no same-spec peer** (`total === 1`, the viewer alone in that class/spec):
   `{row.name} · {spec} {class} · the only {spec} {class} in this guild this week` · `Gear in
   the planner` — states the true fact rather than a meaningless "1st of 1."
3. **Unverified member** (a `guild_characters` row exists, `verified_at` is null): `You are not
   verified yet. An officer can approve you from the roster below.` — **this is the only place
   on the page this fact appears**; the reports panel's own former copy of it is deleted (§3).
4. **Not a member / signed out:** `SignInPrompt.svelte`, reused verbatim, its existing `line`
   prop set to `See where you stand in this guild.` — this is the "Battle.net button style" this
   spec's own task names: `SignInPrompt`'s existing `Sign in with Battle.net` button
   (`SECONDARY_BUTTON_FIXED`, `border-line-warm-strong`, `text-strong`, 44px), the one Battle.net
   affordance already on this page family (`GuildClaim`'s signed-out branch uses the identical
   component) — **not** a second, new filled-gold button: the design system reserves that token
   for the signed-out home page alone, and `SignInPrompt` already carries the Battle.net identity
   and copy this task's wording describes.

**State:** `Skeleton` (one line) while `/v1/me`/`fetchGuildHome` resolve, reserving the measured
height of branch 1's own two-line wrap at 360px so arrival never shifts the facts line below it
(§8). A failure here never blocks the header's own h1/facts line — same "deferred, never
blocking" rule Logs' own per-player hook line states — it simply omits the role-based line.

### 4.B Officer tools strip — new, `GuildOfficerTools.svelte`, officer/leader viewer only

Directly under the header band, full width, one compact panel row — **never shown to a plain
member** (member review §4, verbatim: "I'm not an officer; this is loot-council/admin plumbing
bolted onto what should be my personal gear helper"). Three fields in one row, 44px tall,
`panel` chrome (`bg-raised` `#0d111a`, `border` `#262e40`, `radius-panel` 6px):

- **Claim badge**, left: a `pill` (`pill-site` token, gold) reading `Claimed` / `Pending` /
  `Contested`, or, when `unclaimed`, a text link `Claim this guild` → the existing `/claim`
  sub-page (kept mechanism, `GuildClaim.svelte` already owns the full flow — this strip never
  duplicates its buttons, only surfaces the state and a door to it, tenet 11).
- **Pending approvals**, centre, shown only when `waitingForApprovalCount > 0`:
  `{count} waiting for approval` as a text link that scrolls the page to the roster's own
  unverified group (`#guild-roster-unverified`, an anchor the roster table already needs for
  its own heading) — the exact fact the leader review calls "genuinely useful... without tabbing
  out to the browser mid-raid," now one click from the top of the page instead of a scroll past
  the reports panel to find it.
- **Invite link**, right: `Invite link` label, the existing `guildSettingsCopy.inviteWarning`
  sentence shortened to fit one row (`Anyone with this link can join as a member.`), and the
  existing `Rotate invite link` button (`rotateInvite`, kept mechanism) — moved here from being
  settings-page-only, since the leader review names this exact action as one of the only two
  things worth opening this surface for. `GuildSettings.svelte`'s own page is unchanged and still
  owns the one-time-reveal token display and the default-visibility/officer-threshold controls
  this strip does not duplicate (tenet 11: the deeper editing surface stays one page away, this
  strip is the glance-and-act version).
- **Frozen notice**, replaces the whole row when `claim.frozen`: `guildHomeCopy.frozenNotice`
  (kept verbatim, unchanged), `role="alert"`.

**State:** `Skeleton` (one row, 44px) while `fetchGuildHome` resolves; omitted entirely for a
non-officer viewer (not a loading state they ever see, since `canManage` is already known from
the same fetch the header's own role-based line reads).

### 4.C Roster — new sortable table, `GuildRosterTable.svelte`, member/officer viewer only

**Never shown to a public visitor** (membership spec §4.2: no raw roster is public today, only
`roster_best`'s per-encounter aggregate — this spec does not change that exposure boundary).

- **Panel header row:** `Roster` (`section-title`, 18px) · the existing
  `guildHomeCopy.waitingForApproval(count)` line, officer viewer only, now paired with one new
  button, **`Approve all`** (`SECONDARY_BUTTON_FIXED`), calling `approveCharacter` once per
  unverified row in sequence (no new endpoint — the existing per-row mutation, looped), disabled
  while any approval is in flight, each row's own disabled/`aria-busy` state following the same
  `rosterBusy` pattern the single `Approve` button already uses. On a partial failure (row 4 of
  7 fails), the rows that succeeded stay approved (optimistic per-row update, unchanged
  mechanism) and `rosterActionError` names which row failed by name, never a generic "some
  failed" with no identity (tenet: an officer must know exactly who to retry).
- **Unverified group**, pinned first (`orderRoster`, kept, unchanged): each row carries the
  existing `Unverified` pill (`pill-site`, 11px) and, for an officer, the existing per-row
  `Approve` button beside the new bulk one — a single stray unverified row still approves on its
  own with no need to trigger the bulk action.
- **Verified group**, below a 1px `border-line-soft` divider, **sortable** by three columns —
  new. Column headers (`Rank`, `Item level`, `Last seen`), each a 36px-tall button with a small
  mono ▲/▼ glyph that appears on the active column only; clicking toggles
  ascending/descending; default (no explicit sort yet) is the API's own order (rank-grouped,
  then name — unchanged). **"Last seen" is a two-bucket sort, stated honestly, not a true
  timestamp sort**: the API returns only `logged_recently: boolean` (membership spec RULING 9,
  the 24-hour proxy), so this column sorts logged-recently-true rows ahead of the rest and
  nothing finer — §10.1 proposes the field this column would need to do more.
- **Row, each** (the design system's own Character row recipe, `design/DESIGN-SYSTEM.md`
  "Character row" — crest, name in class colour, muted descriptor, a stat on the right — applied
  here rather than invented):
  - Crest, 36px circle, class ring (`classCrestSrc(row.class)`, the Svelte-island inline-`<img>`
    pattern `class-crest.ts`'s own header comment specifies for exactly this case — a `.astro`
    component cannot be imported into this Svelte file). A row with no `class` (the API's own
    `omitempty` case, a `roster`-consent row) renders the crest's own neutral fallback ring, not
    a broken image.
  - `{row.name}` in `classColorVar(row.class)` (existing function, reused) — Cinzel 14px 700, or
    plain `--text` when `class` is absent.
  - `{row.spec ?? ''}` 13px `--text-muted`, muted descriptor.
  - `{row.rank}` — a `pill` only for `officer`/`leader` (member is the unmarked default, no pill
    — visual noise on twenty members every raid night is a tenet-11 defect of its own).
  - `{row.item_level}` mono, 13px, labelled `ilvl` (existing `guildHomeCopy.itemLevelLabel`) —
    omitted, not zero, when the row's own consent withholds it (the API never selects the column
    in that case, so this is simply absent data, never a dash standing in, tenet 8).
  - `Logged in the last day` pill (`pill-site`, existing copy) when `logged_recently`.
  - `GuildRosterHandoff` (existing component, unchanged) at `gear`/`gear_bags` consent.
  - Officer-only: `Approve` (unverified rows) / `Remove` (`may_remove: true` rows) — existing
    buttons, unchanged mechanism, now inside a table row instead of a flex-wrap list item.
- **Empty state** (`soloRoster`, kept logic): officer sees
  `guildHomeCopy.emptyRosterOfficer` + action **`Invite your guildmates`** → the Officer tools
  strip's own invite row (an in-page anchor, `#guild-invite`, not a second navigation — the copy
  changes from the existing bare `Guild settings` label to name the actual next step, per this
  spec's own task wording); a plain member sees `guildHomeCopy.emptyRosterMember`, no action (a
  member cannot invite).

### 4.D This week's raid nights — kept mechanism, elevated chrome, one heading rule

- **Member/officer viewer:** `home.reports` (existing, trailing-7-day window, membership spec
  RULING 8), heading `This week's raid nights` (renamed from the existing bare "This week's
  reports" — "raid nights" is the word both persona reviews use throughout; a dungeon-only report
  still shows under the same heading, since the spec's own dungeon-framing line elsewhere on the
  site already states dungeons log the same way).
- **Public visitor:** the existing public `data.reports` list (unscoped — every report this guild
  has, not "this week," since the public endpoint carries no week filter) under the honest
  heading `Recent raid nights` — never claiming "this week" for data that is not scoped to one
  (tenet 8). This is the same rows `Guild.svelte` already renders as "All reports," renamed and
  promoted to the primary reports panel for a visitor who has no member-only "this week" list at
  all, rather than kept as a second, lower-priority section repeating the same rows a member
  would also see twice (the existing dual-heading fix from UX review defect 2 generalised one
  step further: one reports panel per viewer role, never two).
- **Row, each:** title (`reportLabel`, existing fallback chain title → zone → "Untitled
  report") → `/reports/{id}`, date (mono, muted), `{kill_count} kills · {wipeCount} wipes`
  (existing `reportSummary`/`reportWipeCount`) — panel-boxed list, 1px `border-line-soft`
  dividers, unchanged row shape, promoted out of a bare `<ul>` into the same `panel` chrome every
  other region on this page now uses (tenet 1).
- **Empty state, member/officer:** `No reports this week yet.` (kept) + new action
  **`Upload a raid log`** → `/logs` (the task's own named next step; `EmptyState`'s existing
  `action` prop, unchanged component). **Empty state, public:** unchanged, no action (a stranger
  cannot upload on this guild's behalf).

### 4.E Progression — kept data and row shape, panel chrome, public and member alike

Unchanged `data.progression` (`rankings/guilds.go`'s `Progression`), unchanged row fields
(encounter link, pull count, first-kill date or `not killed`) — the one change is wrapping the
existing `<ul>` in the same `panel` box every other region now uses, and widening the row's own
`grid-cols` slightly so the date column never truncates at 1280px (confirmed against the real
longest value, `2026-12-09`, 10 characters, inside the existing 104px column — already fits, no
change needed there; recorded so a future reviewer does not re-flag it).

### 4.F Roster bests — kept data and row shape, panel chrome, renamed per the leader review

Unchanged `data.roster_best` (`RosterBest`) — the leader review's own phrase, "the single thing
nobody else does for this game," names this list as the guild's real best-parse record; heading
changes from the bare `Roster bests` to `This guild's best parses` (clearer antecedent for
"roster" when the member-only roster table above it is now a visible, different list on the same
page — the two lists must never read as the same thing twice).

---

## 5. States, every region

| Region | Loading | Empty | Error | Public visitor | Member | Officer/leader |
|---|---|---|---|---|---|---|
| Header hero (§4.A) | `Skeleton`, h1+facts line reserved height | N/A (facts always render once `data` resolves) | `GuildStatus` failed, existing retry | identity + facts, no role line until `/v1/me` resolves | identity + facts + standing/unverified line | identity + facts + standing line (officers are members of their own guild too) |
| Role-based line (§4.A.1) | `Skeleton` one line, deferred, never blocks h1 | N/A (`SignInPrompt` is itself the "nothing to show" case for a stranger) | Omitted on failure, never blocks the band | `SignInPrompt` | standing or unverified-note branch | standing branch (an officer's own rank/ilvl line, same rule) |
| Officer tools (§4.B) | `Skeleton` one row | N/A (always renders once `canManage`) | `rosterActionError`-style inline message, existing pattern | Not rendered | Not rendered | Claim badge, approvals count, invite row |
| Roster table (§4.C) | `Skeleton` (N rows) | `emptyRosterOfficer`/`emptyRosterMember` + action | `rosterActionError`, existing | Not rendered | Full table, no Approve/Remove/Approve-all | Full table with every officer action |
| Raid nights (§4.D) | `Skeleton` (5 rows) | `No reports this week yet.` + `Upload a raid log` (member); unchanged, no action (public) | `GuildStatus` failed, retry | `Recent raid nights`, `data.reports` | `This week's raid nights`, `home.reports` | Same as member |
| Progression (§4.E) | `Skeleton` | `No pulls recorded yet.` (kept) | `GuildStatus` failed, retry | Renders | Renders | Renders |
| Roster bests (§4.F) | `Skeleton` | Section omitted when empty (kept, `data.roster_best.length > 0` guard) | `GuildStatus` failed, retry | Renders | Renders | Renders |

Every interactive element (sort header, Approve/Remove/Approve-all, Rotate invite, claim-badge
link) has default, hover (`120ms` border/colour transition, the site's own hover budget),
focus-visible (2px gold outline, the existing token every button on the site already uses),
active, disabled (`BUSY_CLASS`, existing), and — for Approve/Remove/Approve-all specifically —
a loading state (`aria-busy`, existing `rosterBusy` pattern extended to the bulk action). Touch
and keyboard parity: every row action keeps the existing 44px phone / 36px desktop hit height
(`design/DESIGN-SYSTEM.md`'s own hit-target rule); a sort-header button is reachable by Tab in
document order, left to right, and toggles on both Enter/Space and click (tenet 9).

---

## 6. Verbatim copy table

| Region | String | Source |
|---|---|---|
| Header eyebrow | `Guild` | New, matches every other page's own-name eyebrow |
| Header Updated stamp | `Updated {date}` | New, `UpdatedStamp`-style convention, omitted when no report exists |
| Header facts line | `{rulesetLabel} {region} · {killed} bosses down · {pulls} pulls` | `Guild.svelte`, kept verbatim |
| Standing line (ranked) | `{name} · {rank} of {total} {spec} {class}s this week by item level ({ilvl} ilvl)` + `Gear in the planner` | New; link label reused verbatim from `guildHomeCopy.openPlanner` |
| Standing line (alone in spec) | `{name} · {spec} {class} · the only {spec} {class} in this guild this week` | New |
| Standing line (no gear consent) | `{name} · {spec} {class} · Set your gear consent to see where you stand` | New, links to `/account` |
| Standing line (unverified) | `You are not verified yet. An officer can approve you from the roster below.` | `guildHomeCopy.notVerifiedNote`, reused, now shown once |
| Sign-in line | `See where you stand in this guild.` | New, passed to `SignInPrompt`'s existing `line` prop |
| Officer tools claim badge | `Claimed` / `Pending` / `Contested` / `Claim this guild` | New pill text + existing `guildHomeCopy.claimLink` |
| Officer tools approvals | `{count} waiting for approval` | `guildHomeCopy.waitingForApproval`, kept verbatim |
| Officer tools invite warning | `Anyone with this link can join as a member.` | Shortened from `guildSettingsCopy.inviteWarning` |
| Officer tools invite button | `Rotate invite link` | `guildSettingsCopy.rotateButton`, kept verbatim |
| Roster bulk action | `Approve all` | New |
| Roster empty, officer | `You're the only member the site knows about. Share the invite link to bring the rest of the guild in.` + `Invite your guildmates` | `guildHomeCopy.emptyRosterOfficer` kept; action label changed from `Guild settings` |
| Roster empty, member | `You're the only member the site knows about.` | `guildHomeCopy.emptyRosterMember`, kept verbatim |
| Raid nights heading, member | `This week's raid nights` | Renamed from `guildHomeCopy.reportsHeading` |
| Raid nights heading, public | `Recent raid nights` | Renamed from `guildHomeCopy.allReportsHeading` |
| Raid nights empty, member | `No reports this week yet.` + `Upload a raid log` | `guildHomeCopy.noReports` kept; new action |
| Progression heading | `Progression` | `guildHomeCopy.progressionHeading`, kept verbatim |
| Progression empty | `No pulls recorded yet.` | `guildHomeCopy.noProgression`, kept verbatim |
| Roster bests heading | `This guild's best parses` | Renamed from the bare `Roster bests` |
| Frozen notice | `This guild's claim is contested. Officer actions are frozen until a moderator resolves it.` | `guildHomeCopy.frozenNotice`, kept verbatim |
| Board caption (mock only) | `Mock roster — 24 characters, no real guild data yet.` | New, this spec's own mock boards only, never shipped copy |

---

## 7. Phone layout, 390px (and 360px minimum)

Region order, top to bottom: Header band (eyebrow, h1, Updated stamp, facts line, role-based
line — never gated behind a scroll) → Officer tools strip (officer viewer only; its three fields
stack to one column, each its own 44px row, claim badge first) → Roster panel (sort-header
buttons horizontal-scroll like every other pill row on phone, `.level-scale` pattern already
used elsewhere; each row's crest/name/spec stack above its stat/action row rather than wrapping
mid-line) → This week's raid nights → Progression → This guild's best parses.

**What collapses:** the roster table's three-column header row becomes a horizontally-scrollable
strip of sort buttons (same mechanism the guide pages' `SpecTabs` already uses at 390px); every
other panel already stacks to one column with no mechanism change.

**What never collapses:** the header's h1, facts line and role-based line (§1's own no-scroll
requirement); every 44px hit target; the roster row's crest, name and class colour (never
dropped for space — tenet 4, "nothing clipped... when it is the point," and a roster row's class
identity is the point).

---

## 8. Performance and polish limits

- **No layout shift on hydration.** The header's `Skeleton` reserves the ready role-based line's
  own measured two-line height at 360px (branch 1's wrap is the tallest), so its late arrival
  changes opacity only (the site's existing `.reveal` convention), never height.
- **No missing-icon flash.** Every roster-row crest resolves through the existing
  `classCrestSrc`/`ClassCrest` background-ring CSS (`class-crest.ts`'s own documented fix: the
  ring and `--color-raised` background paint before the image loads, so a slow crest never
  flashes transparent) — this page adds no new icon asset.
- **LCP candidate is the header h1**, present in the server-rendered skeleton's reserved text
  size before any island hydrates, same convention `logs-header-h1`/`planner-header-h1` already
  document. The role-based line's fetch (a third read after `/v1/me`/`fetchGuildHome`, both
  already in flight for the rest of the page) is explicitly deferred and must not appear in any
  Lighthouse trace's critical path.
- **Lighthouse budget, proposed (`web/lighthouserc.json` carries no `/guild.html` row today —
  this is a new row this spec proposes, §10.3):** performance ≥0.90, accessibility/SEO ≥0.95,
  LCP ≤2200ms, TBT ≤100ms, CLS ≤0.05 — the same tier `/logs.html`/`/planner.html`/`/setup.html`
  already carry, chosen because this page, like those three, has authenticated, per-row
  interactive controls (approve/remove/sort) rather than the flatter 1700ms-budget tier the pure
  content pages (`/guides`, `/bis`) carry.

---

## 9. Out of scope

Named with a one-line design sketch each, per this spec's own task — all three are explicitly
out of scope in the guild membership design's own §7 ("officer tools... exist only as design
docs"), and both persona reviews ask for them by name:

- **Loot council helper** (leader review §3, the review's own "single thing that would move my
  guild to require the addon"): a `/guild/<...>/loot` page reading this page's own consent-gated
  roster plus a night's drop list (once drops are logged — no such data contract exists yet) to
  rank who an item upgrades, item icon and quality colour leading each row, attendance count
  beside each name. Needs a new drops-per-report data contract this spec does not have.
- **Raid readiness board** (leader review §3: "we already have the data (`bags=` in FS1);
  nobody renders it per-raider"): a quiet pill per roster row reading that character's own
  enchant/consumable gaps straight from its last export's existing `bags=`/`sets=` FS1 sections
  — no new page, no new endpoint, just a new column this spec's own roster table (§4.C) could
  grow once that decode exists client-side.
- **Attendance trend** (leader review §3): a sparkline column on the same roster table, pulls
  attended over the last N reports, built from `fights.players` joined to the same `reports` the
  Progression panel's pull counts already read — no new table, a per-character aggregate query.
- **The addon's own `GuildView.lua` roster row** (leader review §3's class-icon/spec/item-level
  defect, member review §4's "I never want to see the claim line as a member") — a separate,
  in-game surface; this spec is the web page only. The fix this spec ships here (a real roster
  row with a crest, class colour, spec and item level) is the reference the addon's own next
  spec should match, named here so it is not re-derived from nothing.
- **Recruitment tooling** (leader review §3: "nothing anywhere... correctly out of scope per the
  proposal"). Not sketched further — no review asks for a specific shape.
- **Per-guild banner art** for the header band — ships `ArtPanel`'s own flat fallback, same
  ruling already standing for `/bis` and `/guides` (§4, not re-litigated here).

---

## 10. Rulings proposed for the owner

### 10.1 `logged_recently` should become a real timestamp (or a relative-time field) for sorting

**Proposal:** `RosterRow` gains `logged_at?: string` (the same `addon_exports.updated_at` value
`logged_recently` is already computed from server-side, simply exposed instead of collapsed to a
boolean), so the roster table's "Last seen" sort (§4.C) can be a true recency order instead of a
two-bucket split.

**Reasoning:** the member review's own "did my logs upload" ask and the leader review's "who
hasn't logged/synced recently" ask both want a real recency signal; `logged_recently`'s 24-hour
boolean (membership spec RULING 9) already reads this exact column server-side, so exposing it
costs one struct field, not a new query — the same "small, additive response shape change" the
membership spec's own fourth amendment already used for `zone` and `may_remove`.

### 10.2 The role-based standing line's same-spec rank should read item level, not DPS, until a weekly roster-best exists

**Proposal:** §4.A.1's rank is explicitly item-level-based, not a DPS/parse rank, even though the
member review's own wording ("4th of 6 Fury Warriors") reads like a performance rank.

**Reasoning:** `RosterBest` (the only performance data this guild page has) is a per-encounter,
all-time best parse, not a "this week" figure — ranking same-spec raiders by it would silently
mix a raid night from a month ago with one from last night, misleading the exact player this
line exists to help. Item level, read off `HomeRoster`'s own already-correct, consent-gated,
current column, is the one same-spec comparison this page can make honestly today (tenet 8). A
true "this week's parse rank" is a real, named follow-up once a weekly-scoped `fight_metrics`
aggregate exists — not invented here as a number this page cannot actually back.

### 10.3 Add a `/guild.html` row to `web/lighthouserc.json`

**Proposal:** stated in full in §8 — no `/guild/*` URL exists in the budget config today
(confirmed by reading the file), so this page ships with no Lighthouse gate at all unless one is
added. Proposed tier matches `/logs.html`/`/planner.html`/`/setup.html` (0.90 performance, 2200ms
LCP), for the reason stated there.

---

## 11. Acceptance screenshots

Every screenshot is of the **rendered page**, never a diff (tenet 6). Two reviewers
(ux-designer, wow-player) SHIP on these before merge. A same-scale side-by-side with the mock
boards (`design/mocks/renders/guild-{member,officer,public,phone}.png`) is required for every
state below.

**Viewports:** 360px phone, 390px phone, 1024px, 1280px desktop, 1440px desktop, 1920px desktop,
and one capture at **2000px** confirming the band and every panel stay centred at 1344px rather
than flush left (the guides-page centring regression, named so it is never repeated here).

**States to capture, each at every viewport above:**
1. Verified member, ranked among same-spec peers: full page, header standing line resolved with
   a real rank/ilvl and a working `Gear in the planner` link.
2. Verified member, alone in their spec: header standing line's "the only ... in this guild"
   branch.
3. Unverified member: header's single not-verified line, confirming it does **not** also appear
   inside the raid-nights panel.
4. Officer/leader: Officer tools strip (claim badge, pending-approvals count, invite row), the
   roster table with Approve/Remove/Approve-all all visible and at least one unverified row
   grouped above the sortable verified group.
5. Officer, claim contested: the frozen notice replacing the Officer tools row, and every
   Approve/Remove/Rotate control disabled.
6. Public visitor, not a member: identity, facts, `SignInPrompt` in the role-based slot,
   Progression, "This guild's best parses," "Recent raid nights" — confirming no roster table,
   no Officer tools strip, and no "this week" claim over unscoped data.
7. Empty states: a freshly-claimed guild with one member (`emptyRosterOfficer` with its new
   `Invite your guildmates` action; `emptyRosterMember` with none); zero raid nights
   (`No reports this week yet.` with its new `Upload a raid log` action); zero pulls
   (`No pulls recorded yet.`).
8. Roster table sorted by each of its three columns in turn, confirming the unverified group
   never moves out of first position regardless of sort.
9. Header hero's own loading skeleton at 1280px and 360px, confirming reserved height matches
   the ready standing line's measured height (tenet 13).
10. Any one hover/focus-visible state on a sort-header button and on `Approve all`, to verify
    tenet 9's keyboard/touch parity.
11. 360px phone: the full page top to bottom, confirming the §7 region order and that every 44px
    control held its hit target.

**Lighthouse budget:** the proposed `/guild.html` row (§10.3) — performance ≥0.90,
accessibility/SEO ≥0.95, LCP ≤2200ms, TBT ≤100ms, CLS ≤0.05.
