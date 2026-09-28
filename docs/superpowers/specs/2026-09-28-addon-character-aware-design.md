# The addon, character aware and worth keeping on: a design pass

**Date:** 2026-09-28. **Owner ask:** "the addon isn't character aware. Let's do a design pass on the addon to make sure that it provides a premium level of ingame value to both novice and advanced players. What features is it missing?"

## 0. What the addon can know

No network. Data reaches it three ways: the nightly Forever Sixty Data addon (a public Lua file: talent trees, stat weights, per-character ratings), the companion's inbox in SavedVariables (private, per account, written from the site when the companion syncs; today only builds), and a pasted build code. Data leaves it as the export string (talents, gear, bags, bank, professions, guild, level, name) and through the companion reading SavedVariables at logout. Every feature below names which channel feeds it; a feature that needs a channel that does not exist says so.

## 1. Character awareness (the defect, and the model)

Today `ForeverSixtyDB.follow` is one loaded build for the whole account: an alt shows the main's build in the tracker, the talent glow, the Follow tab and the Gear scoring. The companion's inbox builds carry a `character` key the addon ignores, so a build queued on the site for one character is offered to all.

The model:

- **Per character** (key `region/realm/name`, the key the export already writes; a spec segment can be appended later without a migration): the loaded build (code, name, source), named build slots, tracker shown/complete state, dismissed inbox entries, the last rating card seen.
- **Per account**: window position and tab, minimap angle, tooltip/toast/chat toggles, auto-save, the advanced-detail toggle.
- **Login on a character the addon has never seen**: nothing loaded, tracker and toast hidden, Overview reads "No build loaded for <Name> yet. Paste a code, or queue one on foreversixty.gg." Nothing errors.
- **Login with a build queued for this character**: a one-time banner on Overview, "A build arrived for <Name>." with **Load it**; builds queued for other characters never show; unaddressed builds show for every character.
- **Named build slots per character** (advanced): "Raid", "Leveling", "PvP"; switching is one click; the tracker follows the active slot.

Lane `addon-character` (running) delivers the per-character follow state, the migration and the inbox filter; slots and the banner are the next addon change.

## 2. Novice: what keeps the addon on while leveling 1 to 60

Ranked by how often it answers "what do I do now" without alt-tabbing.

1. **Level-up talent toast** (exists): "Level 12. Take Improved Rend, rank 2 of 3." Fed by the loaded build.
2. **Next-point tracker** (exists): "Next: Improved Rend, 12 of 51 points."
3. **Rotation toast at new-ability levels** (missing): "Level 20: Sunder Armor opens your rotation now." and, on the Overview, a "Your rotation at level N" card with the three to five abilities in order. Fed by per-level rotation lines the accuracy program's ladder is producing for every spec (`docs/superpowers/specs/2026-09-28-rotation-accuracy-program-design.md`, Phase 3): the pipeline emits them into Data.lua per class per level band. This is the biggest novice gap and it costs the site nothing new once Phase 3 lands.
4. **"Planned for your Chest" tooltip** (exists): stops a new player vendoring what the build wants.
5. **Talent glow** (exists): no misclicks in an unfamiliar trait window.

Not for the addon: quest routes and zone guides. The guardrails in section 5 apply.

## 3. Advanced: what a raider keeps it on for

1. **Gear deltas in the tooltip and the Gear tab** (exists; gap: one weight set per spec). Add the site's other weight sets (single target vs multi-target) as a per-character choice.
2. **Your top upgrades, in game** (missing): the five best upgrades from this character's last simulator run and where each drops. This is private data and belongs in the companion's inbox, which grows from "builds" to typed messages: `upgrades`, `rating`, `weights`. The site writes them for saved characters at sync; the addon shows them on the Overview and the Gear tab, dated ("from your sim on Sept 27").
3. **Personal rating card** (missing; the data is already in the data addon for guild rosters): "Last rated fight: output 82 · survival 91 · mechanics 74", on the Overview, with the fight and date.
4. **Guild state in the Guild tab** (missing): claim state, the officer's pending approvals count, and the roster's "logged today" from the data addon (all public on the site already). Officers check the site today; put the same answer in game.
5. **Caps and breakpoints in the tooltip** (missing): "Hit capped: this ring's hit is wasted." when the site's weights carry caps. Fed by the data addon's weights, which need a `caps` field per spec.

## 4. Cross-cutting

- **Progressive disclosure**: the same five surfaces show to a level-12 alt and a raider. One account-wide toggle, "Show advanced detail" (off until the account has a level-60 export), gates stat-weight breakdowns and cap call-outs.
- **The toast is underused**: level-up talent only. Rotation milestones (novice) and "planned set complete" (advanced) use the same mechanism.
- **The Overview sync card** now shows the code head; with typed inbox messages it becomes the place a character's news lands (build arrived, upgrades updated, rating card).

## 5. Guardrails

The addon is the last-mile executor of decisions made on the site, per character. It is not a second planner, a live damage meter, a log parser, a loot or DKP system, or a raid-coordination tool: it has no network and the companion syncs at logout, so anything needing live cross-player state would be a second, drifting source of truth.

## 6. Order of work

- **Wave A (running):** per-character follow state, migration, inbox addressed by character; then named build slots and the "build arrived" banner.
- **Wave B:** rotation toasts and the rotation card (needs the ladder's per-level rotation lines in Data.lua), personal rating card, advanced-detail toggle.
- **Wave C:** typed inbox messages (upgrades, weights with caps) from the site + companion, tooltip caps, guild state in the Guild tab.

Every wave ships to the beta box for an in-game check before release; the addon's busted suite and luacheck gate each commit.
