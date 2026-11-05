# Guild gating, claim and invites (experience spec)

Author: ux-designer, 2026-11-05. Round-1 guild audit fixes UX-3, UX-4, UX-5, UX-10, UX-15, UX-17, WP-8, WP-10, WP-21 (UX-12 and UX-13 as listed by the coordinator; the round-1 audit text is not on disk, so this spec covers what the captures show).
Amends `design/specs/2026-10-25-officer-tools-gate.md` §3.2 and §7 (the sentence words), `design/specs/2026-10-27-launch-pricing.md` §3.2.8 (Guild-tier gate), `design/specs/2026-10-26-guild-raids.md` §6.1 (`gate`), and the join page of `design/specs/2026-10-27-subscribe-and-join.md` (E9).
Captures audited: `scratchpad/guild-r1/shots/{pre,claimed,free,paid}-*` and `only-{anon,stranger}-invite-*`.
Code read: `web/src/components/guild/{GuildPlanGate,GuildInviteLink,GuildTabs,GuildRosterTable}.svelte`, `web/src/components/{GuildJoin,GuildSettings}.svelte`, `web/src/components/ui/{TabStrip,CopyLinkField}.svelte`, `web/src/lib/guild/{copy,raid-copy}.ts`, `api/internal/guilds/{invite,settings,claim}.go`, `api/internal/addon/outbox.go`, `api/internal/entitlements/entitlements.go`.
No new control and no new token (tenet 10).

## 1. The one question per viewer

| Viewer | Question | Answer, no scroll at 360 and 1280 |
|---|---|---|
| Guild master, guild unclaimed | "How do I take charge of this guild here?" | The band's claim row (§3.1), on every tab, under the facts line. |
| Officer, a gated control | "Why is this greyed, and who can turn it on?" | One sentence above the control, naming the plan and the person or step that unlocks it (§3.2). |
| Officer, recruiting | "Where is the link I paste in guild chat?" | Settings `#invite`: `Make a new link`, then the link in a read-only field with `Copy` (§3.3). One click from Overview's `Copy the invite link`. |
| Invited player | "Which guild, and which of my characters joins?" | The join card (§3.4): the guild, the joining character's crest and name, one `Join` button. The inviter joins the card when §4.1 lands. |

## 2. Reference, and what we borrow

- **The in-game greyed spell button** (officer-tools gate §2): the control stays in its slot, dimmed. Kept.
- **Discord's server invite** (`discord.gg/...` preview): "{User} invited you to join {Server}", the inviter's avatar, then one `Accept Invite`. Borrowed: the guild named first and one primary; the inviter line follows once §4.1 records who made the link.
- **Discord's server settings, Invites**: a link you can always copy, with `Revoke`. Borrowed: the link is copyable on every visit, and making a new one is a separate, confirmed act.
- **The in-game Guild Info frame**: the guild master is a character name in class colour, never an account tag. Borrowed for the claim line.

## 3. Layout, region by region

### 3.1 The claim row (band, before a claim exists)

Replaces the band's bare gold `Claim this guild` link (capture `pre-gm-overview-1440`). Drawn under the raid line, above the tab strip, on every tab, while `guild.claimed === false`.

- Layout: one row, `flex items-center justify-between gap-4`, min 44 px, wraps below 768 (sentence, then the control full width).
- Guild master rank (rank 0 on the viewer's verified or unverified row): 14 px `--text` sentence, then `button('secondary','md')` with the gold-outline recipe `Claim {Guild}`, to `/guild/{r}/{rs}/{slug}/claim`. Not filled: the page's one fill belongs to the Before raids open card (spec A §3.1).
- Officer rank (not GM): sentence, then `ActionLink` `Copy the claim link for your guild master`. It copies the claim page URL (the `copy-feedback.ts` recipe: label reads `Copied` for 4 s; failure falls back to `CopyLinkField` inline holding the URL, selected).
- Member: no row. The Before raids open card carries their next step.
- A claim pending: the row reads the pending sentence with `Open the claim` (`ActionLink`) for officers and GM; members see nothing.
- States: the row renders with the band, from the same `/home`; no loading state of its own. Height is held by the band reserve (`band-hold.ts`) like the raid line.

### 3.2 Gate sentences: one per feature, naming the plan

**Plans and features** (from `entitlements.go` and launch pricing):

| Feature | Unlocked by | Gated controls |
|---|---|---|
| Plan raids and events, sign-up officer controls | Guild or Raid guild (`officer_views`) | `Plan an event`, the raid page officer row (Copy stays live) |
| Discord channel | Guild or Raid guild (`officer_views`) | Connect, Send test message |
| Loot council (`Award`, `Undo`) | Raid guild (`raid_guild_tools`) | `Award`, `Undo` |
| Readiness `Nudge` | Raid guild (`raid_guild_tools`) | `Nudge` |

**The sentence** is the existing `GuildPlanGate` element (13 px `--text`, 20 px line, id `officer-tools-gate-{tab}`), officers and the GM only, members never. It is two parts: what the feature needs, then what this viewer can do about it. The second part depends on the viewer, never on the tab.

| Part 1 (by feature) | |
|---|---|
| Raids tab, raid page | `Planning raids and events is part of the Guild and Raid guild plans.` |
| Settings · Discord | `Posting to Discord is part of the Guild and Raid guild plans.` |
| Loot tab | `Awarding loot is part of Raid guild.` |
| Readiness tab | `Nudge is part of Raid guild.` |

| Part 2 (by viewer and billing) | Text | Link target |
|---|---|---|
| Claim holder, no plan, trial unused | link `Start the 7-day trial` | settings `#billing` |
| Claim holder, no plan, trial used or lapsed | link `Subscribe your guild` | settings `#billing` |
| Claim holder on Guild, a Raid guild feature | link `Move to Raid guild` | settings `?upgrade=1#billing` (the Move confirm opens) |
| Claim holder on a Guild plan Forever Sixty granted (no billing contact), a Raid guild feature | `Your plan was granted by Forever Sixty. ` + link `Ask us on Discord` + ` to move it to Raid guild.` (amended 2026-11-05, guild loop round 3: settings has no Move control on a grant, so the Move link led nowhere) | the community Discord invite, new tab |
| Officer, someone else holds the claim | `Ask ` + crest 16 px + `{GM character}` (class colour) + ` to turn it on.` The crest sits 4 px from `Ask` and 4 px from the name (amended 2026-11-05, guild round 2: round 1 drew it flush against `Ask`). | none |
| Guild unclaimed, viewer is GM | link `Claim the guild first` | claim page |
| Guild unclaimed, viewer is an officer | `Your guild master claims the guild first.` | none |

**Preselect.** From a Raid guild feature (Loot, Readiness), the trial and subscribe links carry `?plan=raid_guild` (`parseBillingPreselect` already reads it), so `#billing` opens with Raid guild selected. From a Guild feature (Raids, Discord) they carry no plan, and the chooser keeps its fit-by-size default.

Example, a free guild's officer on Loot: `Awarding loot is part of Raid guild. Ask [crest] Thornhide to turn it on.`

Site-wide, `Subscribe the guild` becomes `Subscribe your guild` (another lane owns the rename; this spec uses the new words).

Deleted strings: `Award and Nudge are part of the guild plan. ` + `Subscribe the guild` (it named the wrong plan for Award and baited a Guild purchase that unlocks neither), `Loot council and the readiness board are part of Raid guild. ` + `Move to Raid guild` (the board itself is readable on Guild; only Nudge is gated), `Planning and running raids are part of the guild plan. Subscribe the guild.`

**Gated controls** keep their slot, label and size (officer-tools gate §3.4) and take the design system's disabled state: native `disabled`, 45% opacity, `cursor-not-allowed`, no hover, out of the tab order, `aria-describedby` the sentence. One change: a gated control that is a primary (`Plan an event`, `Plan a raid` today) renders as `button('secondary','md')` while gated. A 45% gold fill still reads as a live gold button (capture `free-officer-raids-1440`); the gold belongs to the sentence's link, the one thing the viewer can act on.

### 3.3 Settings: invite link, rank field, claim line

**Order of the Members and access card** (`Card`, `CardHead` `lock`): Invite link first (`id="invite"`, `scroll-mt-20`), then Officer ranks, then Default report visibility. Recruiting is the weekly job; the other two are set once.

**Invite link row** (officers and GM; a `SettingsRow` with the field under it). Built with no migration; §4.1's stored link is held for the owner.
- Label `Invite link`, sub `Anyone with this link joins {Guild} as a verified member: no officer has to approve them, and they can sign up for raids and see the roster, loot and readiness at once. A character with an officer rank in game still waits for an officer.` (the officer clause amended 2026-11-05, guild loop round 3: the sub promised every joiner skips approval, §3.4 "The promise holds after the addon syncs")
- **Just made, this session.** `RotateInvite` already returns the URL once. After `Make a new link` (or `Make the invite link`) answers, the row shows `CopyLinkField` (read-only, the full URL in mono, `Copy` reading `Copied` for 4 s), focus on `Copy`, and under it 13 px muted `Copy it now. For safety we keep only a fingerprint, so it cannot be shown again after you leave this page.` The URL lives in component state only, never in storage.
- **A link exists, later visit.** No field. 13 px muted: `Made {12 Nov}. It still works, but it cannot be shown again. Make a new one to copy it.` Then `Make a new link` (`button('secondary','sm')`).
- **No link yet.** 13 px muted `Make a link to share in guild chat or Discord.`, then `Make the invite link` (`button('secondary','sm')`, busy `Making…`): the Guild plan card's trial or subscribe button is the page's one fill, and two filled gold buttons on one page read as two first steps (amended 2026-11-05, guild round 2). No confirm: nothing breaks.
- `Make a new link` (a link exists) opens the confirm in place (the `GuildInviteLink` recipe): `Make a new link? The old one stops working, so anyone you sent it to needs the new one.` with `Make a new link` (`button('primary','sm')`, busy `Making…`) and `Cancel` (ghost). Focus moves to `Cancel`; Escape cancels.
- Gated? No. Invites are free on every plan.
- Error: 13 px `--wipe` line under the row, `That did not work. Try again.`
- Loading: the row's text is a 13 px `Skeleton` bar; no field reserve (the field appears only on a press).
- When §4.1 ships: the later-visit state becomes the field with `Copy` and `Made {date} by {Name}.`; nothing else changes.

**Officer ranks row** (UX-15). Rank names are not available: the addon roster snapshot carries name, level, spec, gear gap and build, no rank names (`addon/outbox.go` `RosterMember`), and the roster says `Ranks show once the addon sends them`. So the number field stays. Its label is `Officer ranks`. Its sub reads `0 is the guild master, 1 the rank under it, and so on. This rank and every rank above it can approve members and award loot.` A rank-name select replaces it in the lane that adds rank names to the snapshot; not this round.

**Claim line** (WP-10). `Claimed by hunthrop#1894` is replaced on the settings page and the band by: `Claimed by ` + `ClassCrestRing` 20 px + the claim holder's character in this guild, in class colour, 600. The character is the holder's lowest-rank-number character on this guild's roster, ties to their main. With none on the roster: `Claimed by the guild master`. No BattleTag on any guild surface, for any viewer. `Release on the claim page` stays, for the holder only.

**Plural** (WP-21): `1 verified members today.` reads `1 verified member today.` through `counted()`.

### 3.4 The join page (`GuildJoin.svelte`)

Capture `only-stranger-invite-1440`: a 480 px card hard left, `Lifecycle Test · PvP US invited you.` (a guild cannot invite), no character.

**Layout.** One `card-surface` card on the form measure (`measure-form`, 480 px), centred in the 1344 px column, 32 px from the header (18 px gutter on phone, full width). Top to bottom, 16 px gaps:

1. **Guild row**: `FactionCrest` 44 px, the guild name in `heading-plan` (24 px, the heading face, its own case), then 13 px muted `{Realm} · {n} members on Forever Sixty`.
2. **Invite line**, 15 px `--text`: `You are invited to join {Guild}.` Only while the invite can be taken: never to a member already in the guild, and gone from the card once the join goes through (amended 2026-11-05, guild loop round 3). The inviter is not named until §4.1 lands (no column records who made the link). Then it becomes `ClassCrestRing` 24 px + `{Inviter}` in class colour 600 + ` invited you to join.`
3. **Joining as**, a `SettingsRow` (div, not a link): icon slot `ClassCrestRing` 32 px; label `You join as {Name}` in class colour; sub `Level {52} {Night Elf} {Rogue}` (each part only when known).
   - First choice: the account's character that is on this guild's roster (the guild's addon snapshot or a roster row names it). Several: the highest level.
   - None on the roster, but the account has characters: label `None of your characters is in {Guild} in game yet.`; sub `Officers see {Name} until one of yours shows on the roster.` ({Name} is the `invitedRosterRows` pick: the main with an export, else the newest export.)
   - No characters at all: icon slot the neutral ring; label `You join with your account`; sub `Your characters show once you sync the addon or import from Battle.net.`
4. **Consent line**, 13 px muted, unchanged words.
5. **Actions**: `Join {Guild}` (`button('primary','md')`, busy `Joining…`), `Not now` (`button('ghost','md')`). The card ends at the actions: room for the error line under them is held only while a join is in flight or after one failed (amended 2026-11-05, guild round 2; round 1 held it at rest and left about 50 px of dead space).

**What verified means**, one 13 px muted line above the consent line: `Joining by this link verifies you at once. No officer has to approve you.`

**The promise holds after the addon syncs** (amended 2026-11-05, guild round 3). The invite verifies the account, not one row: when a character of an invite-verified account syncs into the guild (addon export), or the account already has characters on the guild's roster when it joins, every such unapproved member-rank character takes the invite's verification (`verified_by = 'invite'`) and the account's placeholder row (`account:{id}`) is dropped, in the same transaction (`guilds.AdoptInviteVerification`). A member-rank character of an invited account never comes back `Waiting for approval`. A character whose export carries an officer or leader rank is not verified by the invite: the rank is the export's own word and the link can be passed on, so it waits for an officer like any officer-rank character (Approve all names it, `including 1 officer`), and the account's placeholder stays, keeping it a verified member meanwhile. An invite never hands out officer tools.

**States.** Signed out: rows 1 and 2, then `Sign in to join this guild.` and the existing sign-in button as the primary. Already in: rows 1 and 3 with `You are already in {Guild}.` and `Open the guild` (primary). Invalid link: row 1 absent, `EmptyState` `This invite link no longer works. Ask a guild officer for a new one.` Loading: the card at its ready height (320 px desktop, 380 phone) with `Skeleton` rows. Joined: today's joined block with the plan line.

### 3.5 Roster rows with no account (UX-10)

Capture `claimed-gm-roster-1440`: an empty circle, `Arms in game`.

- Crest: `ClassCrestRing` 36 px from the row's class (§4.2). Unknown class: the component's neutral ringed disc, never an empty outline.
- Name in class colour (muted `--text` when class unknown).
- Descriptor: `{Spec} {Class}` (`Arms Warrior`, `Protection Paladin`); class unknown: the spec alone (`Protection`). The words `in game` are dropped from every roster descriptor; the `No account yet` state says it.
- `No account yet` is a `Tag` (tone `muted`), not loose text.
- Right side unchanged: `Level 58 · seen just now`.

### 3.5a Roster counts are characters (amended 2026-11-05, guild round 3)

The e2e walk read `Waiting for approval 4` on the panel, `2 waiting for approval` on the Overview and `8 on the roster` over a list of 4. Every roster count is now characters, because the list and every approval are per character; the round-1 rule that waiting counts accounts is superseded. `counts.waiting` and the Overview's `summary.waiting_for_approval` are the characters an officer has not approved (the panel's rows). `counts.on_roster` is the Roster list itself: approved characters, invited members and names in game without an account, never the waiting characters the panel above it lists. So the count line reads `4 on the roster · 2 in game without an account · 4 waiting for approval`, the first figure matching the rows under it. `raiders` stays accounts. The `/v1/me` guild card's `roster_count` (everyone on the guild) is `on_roster + waiting`. Web: approving a waiting character moves it onto the list (`on_roster` up one, `waiting` down one); declining one takes it off `waiting` only.

### 3.6 Phone tab strip (UX-17)

`TabStrip.svelte` already scrolls with 24 px edge fades. At 390 the Readiness tab loads with a 6 px sliver of `Overview` at the left (`free-member-readiness-390`): a stub, not a hint.
- On load and on select, the current tab scrolls so at least 48 px of the previous tab shows (`scroll-padding-inline: 48px` on the strip; `scrollTabIntoView` honours it). On the first tab, scroll 0.
- The end fade shows whenever the strip can scroll further, and the start fade likewise; both 24 px, as today.
- No arrows, no new control.

## 4. API

### 4.1 Invite links that can be shown again: HELD FOR OWNER

Nothing ships a migration tonight. This section is the proposal for the owner; the interim in §3.3 and §3.4 is what builds now.
- Migration `0070_guild_invite_shown`: `alter table guilds add column invite_token_enc bytea, add column invite_rotated_by bigint references users(id) on delete set null;`
- `RotateInvite` also stores the token encrypted with the webhook envelope (`api/internal/notify/crypto.go`, the same key) and the rotating user. The hash stays the lookup key.
- Settings (officers only) gains `invite: {url|null, made_at, made_by: {name, class}|null}`. Members never receive it.
- With the migration: the join preview gains `inviter: {name, class}|null`.

### 4.1b Join preview, no migration (builds now)
- The join preview gains `joins_as: {name, class, race, level, on_roster: bool}|null` (§3.4 row 3's pick) and `members` (verified accounts, the billing count).
- Security note for the reviewer: the token is a membership capability like a Discord invite; it is shown only to verified officers, encrypted at rest, and rotation still revokes it at once.

### 4.2 Class for snapshot members (no migration)
- The addon's `snapshot!` member line gains `class` (the classFilename from `GetGuildRosterInfo`); `RosterMember.class` is stored in the existing jsonb.
- Until a snapshot carries it, the API derives class from an unambiguous spec slug (`arms`, `fury` to Warrior; `beast-mastery`, `marksmanship`, `survival` to Hunter; and so on). `protection`, `holy`, `restoration` stay unknown.

### 4.3 Claim holder
- `settings.claimed_by` becomes `{name, class, character_key}|null`; the BattleTag field is removed from every guild response.

## 5. Copy, verbatim

| Where | String |
|---|---|
| Claim row, GM | `You lead {Guild} in game. Claim it here to approve members and share the invite link.` / `Claim {Guild}` |
| Claim row, officer | `Nobody has claimed {Guild} here yet. Your guild master claims it.` / `Copy the claim link for your guild master` / `Copied` |
| Claim row, pending | `A claim on {Guild} is waiting for a second officer.` / `Open the claim` |
| Gate part 1 | `Planning raids and events is part of the Guild and Raid guild plans.`; `Posting to Discord is part of the Guild and Raid guild plans.`; `Awarding loot is part of Raid guild.`; `Nudge is part of Raid guild.` |
| Gate part 2 | `Start the 7-day trial`; `Subscribe your guild`; `Move to Raid guild`; `Ask {GM} to turn it on.`; `Claim the guild first`; `Your guild master claims the guild first.` |
| Invite row | `Invite link`; `Anyone with this link joins {Guild} as a verified member: no officer has to approve them, and they can sign up for raids and see the roster, loot and readiness at once. A character with an officer rank in game still waits for an officer.`; `Copy it now. For safety we keep only a fingerprint, so it cannot be shown again after you leave this page.`; `Made {date}. It still works, but it cannot be shown again. Make a new one to copy it.`; (after §4.1) `Made {date} by {Name}.`; `Make a new link`; `Make a new link? The old one stops working, so anyone you sent it to needs the new one.`; `Making…`; `Cancel`; `Make a link to share in guild chat or Discord.`; `Make the invite link`; `That did not work. Try again.` |
| Officer ranks | `Officer ranks`; `0 is the guild master, 1 the rank under it, and so on. This rank and every rank above it can approve members and award loot.` |
| Claim line | `Claimed by {Name}`; `Claimed by the guild master` |
| Members today | `{n} verified member today.` / `{n} verified members today.` |
| Join | `{Realm} · {n} members on Forever Sixty`; `You are invited to join {Guild}.`; (after §4.1) `{Inviter} invited you to join.`; `You join as {Name}`; `None of your characters is in {Guild} in game yet.`; `Officers see {Name} until one of yours shows on the roster.`; `Joining by this link verifies you at once. No officer has to approve you.`; `Level {L} {Race} {Class}`; `You join with your account`; `Your characters show once you sync the addon or import from Battle.net.`; `Join {Guild}`; `Joining…`; `Not now` |
| Roster | `{Spec} {Class}`; `No account yet` |

## 6. Hierarchy

- Band: the raid line leads; the claim row is secondary (outline button). The Before raids open card holds the page's fill.
- Gate: the sentence is body text; its link is the one gold element in that region; the gated control is visibly dead.
- Join card: the joining character row is the subject; `Join` is the page's one fill.

## 7. Phone and desktop

- Never collapses: the gate sentence (wraps, its link never splits, 44 px tap area per gate spec §3.2), the claim button, the invite URL (the field scrolls inside itself; `Copy` stays visible), the inviter and joining names.
- Collapses: the claim row and invite row stack below 768; the join card goes full width.

## 8. Performance

No new request on the guild page; the claim row and claimed-by line ride `/home` and the settings read. Crests are the existing set; the join card reserves its height. Lighthouse budgets unchanged.

## 9. Acceptance screenshots (390, 1440; 360 for rows 2 and 6)

| # | State |
|---|---|
| 1 | Band claim row: GM, officer, pending; member without it |
| 2 | Raids tab gated, free guild: claim holder, officer without the claim, unclaimed |
| 3 | Loot and Readiness gated on Guild tier (claim holder) and on free (officer) |
| 4 | Settings: just made (field with Copy), a later visit (no field, `Make a new link`), the confirm open, no link yet, the error line |
| 5 | Settings claim line with crest; a member's view of the band (no BattleTag anywhere) |
| 6 | Join page: a character on the roster, characters but none on the roster, no characters, signed out, invalid |
| 7 | Roster with two no-account rows (class known, class unknown) |
| 8 | Tab strip at 360 and 390 on Readiness and Settings, showing the previous tab's 48 px |
| 9 | Focus-visible on every new link and button; disabled controls skipped by Tab |

E2E: no page under `/guild/` renders a `#` followed by four digits; every gated control is `disabled` and described by the sentence; `Copy` after `Make a new link` writes the full invite URL; a reload shows the later-visit state with no URL in the DOM; a Loot-gate trial link lands on `#billing` with Raid guild selected.
