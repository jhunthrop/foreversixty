# Guild leader review — the addon's Guild tab and guild surfaces
**Reviewer persona:** 40-person guild's GM/loot council lead. **Date:** 2026-10-04.
**Read:** tenets.md; guild-membership-design.md + guild-and-premium-proposal.md;
addon.md §4.5.5/§10; `addon/ForeverSixty/views/GuildView.lua`; `Locale.lua` guild* strings;
`ForeverSixtyData/README.md` + `Data.lua`; `companion/README.md` + `internal/addon/addon.go`;
`web/src/pages/guild/[...path].astro` + `GuildShell.svelte`/`Guild.svelte`;
`api/internal/guilds/*`. No `addon-guild.png` board exists yet in `design/mocks/renders/`.

## 1. What the Guild tab gives me today

In game: my own guild name and rank, read straight from `GetGuildInfo` (client — always
correct, zero setup). If the Forever Sixty **Data** addon is installed and current, a public
standing line (`guild.progress`, `"%d raid nights logged"`, `"%d on the roster"`) and a
roster table sorted by **rating** (`Ratings.forGuild`/`forCharacter`, from the nightly
`ForeverSixtyData.characters`/`.guilds` tables) — unrated members sort to the bottom reading
"No rating yet." If the **companion** is running and paired, one private line below that:
claim state (`unclaimed`/`pending`/`claimed`/`contested`) and, only for an officer/leader,
a pending-approvals count (`ForeverSixtyInbox`'s `guild` message, `GuildView.lua`'s
`stateLine`). On the **site**, `/guild/<region>/<ruleset>/<name>` carries the same public
progression/roster-bests/reports it always has, plus — signed in — a roster with per-row
consent-gated gear/item level, this week's reports, approve/remove buttons for officers, and
claim/settings/invite pages (`Guild.svelte`, `GuildClaim.svelte`, `GuildSettings.svelte`,
`api/internal/guilds`). That is the whole built surface. The proposal's officer tools —
raid readiness board, loot council helper, attendance, raid-night comparison, assignments —
exist only as design docs and as one promise line in billing copy ("Officer tools once they
ship..."); nothing in `api/internal` or `web/src/components` builds any of them, and the
in-game Guild tab has no code path to show them even once they land on the API side.

## 2. What I'd actually use, raid night or between raids

1. **The claim and invite flow on the site** — getting my guild claimed and inviting the
   dozen people who don't run the addon is a real five-minute job this does for free.
2. **The in-game claim/approval state line** — knowing "claim pending" or "3 pending
   approvals" without tabbing out to the browser mid-raid is genuinely useful.
3. That's it. The roster-by-rating table is interesting but not actionable (see §4): it
   tells me who's rated, not who's ready, who needs gear, or who's behind. I would not open
   this tab on raid night for any reason beyond those two.

## 3. What's missing to require this guild-wide

- **Raid readiness** (who's missing enchants/consumables/talent points before pull): **addon**
  is the right home (bags and talents are already in the export) but it does not exist; the
  spec for it (proposal §3.3.1) is unbuilt. We already have the data (`bags=` in FS1); nobody
  renders it per-raider.
- **Loot council helper** (who this item upgrades, ranked, with attendance beside each name):
  **site**, server compute — proposal's own flagship-adjacent pick, explicitly unbuilt
  (`api/internal/guilds` has no such handler). This is the single feature that would move my
  guild to require the addon; right now loot council still runs on a spreadsheet.
- **Attendance** (who showed up, trend over weeks): **site**, needs `reports`/`fights` joined
  to roster — unbuilt. We have the raw data (every report carries a `guild_id` and a fight
  roster) but nothing aggregates it into a per-raider attendance number anywhere I can see.
- **Who hasn't logged/synced recently**: partially exists — the design spec's "who logged"
  (last 24h `addon_exports.updated_at`) is in the home spec (§4.1) but I don't see it rendered
  in `Guild.svelte` today; worth confirming on the rendered page, not just the design doc.
- **Rating-floor / BiS-gap roster view**: the in-game roster shows a rating number with no
  context (what's a good rating? what's this person's class/spec/item level?) — `rosterRow`
  in `GuildView.lua` renders only `name` and a rating string, no class icon, no spec, no item
  level, no link to that character's page. That's a list where the client (and this addon
  elsewhere) would show a panel — a tenet-1/11 defect on this specific row.
- **Recruitment**: nothing anywhere — no applicant comparison, no "post your roster" tool.
  Correctly out of scope per the proposal (premium officer tools don't mention it), but worth
  naming since guild leaders do spend real time on it.

## 4. What I'd remove or hide from members

- The **private claim/approval state line** is already correctly officer-gated for the
  approval count, but the claim-state text itself (`"Claimed"`/`"Contested"`) shows to every
  member via the companion inbox message, not just officers — a contested claim is exactly
  the kind of guild-politics detail a GM would not want broadcast to all 40 raiders' game
  clients before it's resolved. Scope it to officer/leader rank the same way
  `pending_approvals` already is.
- Nothing else needs hiding: the public roster-by-rating is fine for everyone to see (ratings
  are already public-by-design per the proposal), and the rest of the tab is read-only
  information, not a control surface a member could misuse.

## 5. Verdict

RECOMMEND, not require: the claim/invite plumbing and the private claim-state line are
real, working conveniences, but with no readiness board, loot council helper, or attendance
built, the Guild tab has nothing that beats what my officers already do by hand.
The one feature that moves it up a level: ship the loot council helper (who this drop upgrades,
ranked, with attendance beside each name) — that is the single thing nobody else does for
this game, and it's the first thing I'd put in front of my loot council instead of a spreadsheet.
