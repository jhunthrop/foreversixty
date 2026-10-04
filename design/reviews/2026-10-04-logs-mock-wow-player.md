# Logs landing mock boards — wow-player review (2026-10-04)

Boards reviewed: `design/mocks/renders/logs-signed-in.png`, `logs-signed-out.png`,
`logs-phone-signed-in.png`. Data: `design/mocks/data/logs-sample-report.json`.
Spec: `design/specs/2026-10-04-logs-landing.md` §10.

## Verdict: FIX

1. **Signed-in hero is not actually the visitor's newest report.** `Your reports`
   lists Sanguine Depths (2026-09-15) above `Deadmines, first clear` (2026-09-28)
   under the caption "Newest first," but 09-28 is newer and sits second. The
   header hero mirrors the stale 09-15 report (same title, same 63 fights/12
   kills, same last-kill line) instead of the real newest row — it's a literal
   reuse of the canonical sample report's data, not Zulmara's own rows[0]. This
   breaks the page's entire job ("how did last night go") and is the first thing
   a raider scanning top-to-bottom would catch. Expect: hero = rows[0] of a
   correctly newest-first-sorted list, per spec §4.A's own hero rule. Confidence:
   high — directly visible, and contradicts the panel's own "Newest first" line.

2. **The sample report reads as verbatim retail content, not Classic.**
   "Sanguine Depths" with final boss "Nalthor the Rimebinder" and a Mistweaver
   Monk in the roster is a Shadowlands Revendreth Mythic+ dungeon, not anything
   that exists in Classic or Forever. This is bigger than the already-filed "no
   Monk class" blocker (spec §4.C.3) — the whole zone/boss pairing looks pulled
   straight from a live retail log. This same report is also the signed-in hero
   (finding 1), so the retail-ism is not confined to the signed-out sample.
   Expect: regenerate the sample against a real Classic/Forever zone and boss
   with a legal Forever roster (no Monk), or point `sample-report.json` at a
   different already-public report, per the spec's own named blocker (§4.C.3,
   §10 finding 5). Confidence: high on the Monk (confirmed illegal class);
   medium-high on Sanguine Depths/Nalthor being retail-sourced.

3. **Report rows overlap their own text on phone.** At 390px, `Sanguine Depths`
   and `Deadmines, first clear` wrap across two–three lines while the
   date/fight-count block stays pinned beside line one, producing a visible
   collision (the title runs straight into the date, e.g. "Sanguine" against
   "2026-09-15" with no space). This is a screenshot-and-mock-in-guild-chat bug.
   Expect: the row either stacks title above stats on narrow widths or wraps
   without colliding, same as any other long string elsewhere on the site.
   Confidence: medium-high — seen directly in the phone render.

4. **Dungeon-framing sentence is orphaned at the bottom on phone.** "Logs are
   for group content at any level: a dungeon run logs the same way a raid
   does." sits after the Upload panel, immediately before the footer, instead
   of right under the character spine the way it does on desktop and the way
   §7's region order specifies. Reads like leftover text next to the footer.
   Expect: same position/order as desktop. Confidence: high — directly visible
   in the phone render's text order.

5. **Companion status dot risks re-introducing the "live" claim the copy
   deliberately avoids.** The green dot next to "MacBook Pro · macOS · last
   seen 4 min ago" isn't in the spec's copy (§4.D.1 explicitly bans any
   present-tense "Connected" claim since `last_seen_at` can be minutes stale
   mid-raid). A green dot reads exactly like Discord's "online now" indicator
   to any player, quietly contradicting the honest relative-time text beside
   it. Expect: drop the dot, or tie its state explicitly to the same recency
   threshold the text commits to. Confidence: medium.

6. **Minor copy redundancy on the sample hero.** The h1 itself bakes "sample
   log" into the title text ("Sanguine Depths, sample log") on top of the
   adjacent `Sample` pill — the same fact stated twice. Low severity. Expect:
   either the title drops its own "sample log" suffix or the pill is dropped.
   Confidence: low-medium.

## What's working (for the record)

- The per-player hook line (name, class colour, DPS, `Gear in the planner`,
  `Sim`) is exactly the thing a raider would click after a good pull — present
  on both signed-in and signed-out boards, correctly deferred/static on the
  sample. This is the site's real differentiator over WCL and it finally shows
  up on the landing page.
- The pairing code block states what happens next ("This panel confirms the
  moment it pairs") — closes the old "did that work?" gap.
- The upload panel is honestly usable five minutes before raid: real limits
  (10s first fight, 4 GB, `/combatlog` restart note), drag-drop, visibility
  toggle.
- Signed-out "Recent public reports" now shows its own honest empty state
  instead of one fixture account masquerading as site traffic.
- Phone nav collapses to a hamburger instead of dumping the full site nav
  above the fold — the old stack complaint is gone.
- No marketing hero, slogan, or feature-pitch language anywhere on either
  board; the page leads with a real report, not the tool's own name.
