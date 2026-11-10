# Onboarding, persona review round 1 (overall round 3)

The six-page onboarding (commit 47aabd7d1), built to the owner's brief: one loop between the site, the addon, the companion and the apps; a QR for the phone app; what is not live marked `coming`. All three personas: **HOLD**. Fixes landed in the next commit; the owner's brief sets the fixed points (the loop pages, first-login auto-open, the QR, `coming` tags), and the dispositions work inside them.

## Clutter-hater: HOLD

| # | Finding | Disposition |
|---|---|---|
| M1 | Cut to one or two pages; pages 1 to 4 and 6 are a pitch | **Partly.** The owner asked for the loop story, so pages 1 to 3 stay. The tour drops to four pages without a guild, five with one: the guild page shows only in a guild, and the done page is folded into the last. `Just the tooltips` on page 1 ends the tour and quiets the addon in one click. |
| M2 | Remove every `coming` tag and every unbuilt feature | **Partly.** The owner asked for "coming" marks, so the brief wins. Page 4's two tagged rows become one muted line, and the QR keeps its tag (the owner asked for it). Nothing unbuilt is claimed as live. |
| M3 | Do not auto-open a big window; offer Not now and Tooltips only first | **Owner's call stands** (first login, once per account). Page 1 now has `Skip tour` and `Just the tooltips` side by side, so the minimalist's way out is the first thing on the first page. It never opens in combat or an instance. |
| S1 | A real tooltip, not a stub | Fixed: the client's player tooltip in full. |
| S2 | No dots for two pages | Not applicable at four or five pages. |
| S3 | Connected state as a panel line | The tour keeps it (the raider wants it); the panel already shows the sim age. |

## Leveller: HOLD

| # | Finding | Disposition |
|---|---|---|
| M1 | Set-up too late; ~3 pages for a solo leveller | **Partly.** A guildless character gets four pages. The order follows the owner's loop, so set-up comes last. A skipped path still has the chip's `Pick a path`. |
| M2 | Playstyle lines, the level-10 line, no pre-pick | **Fixed.** Tiles carry `specPlaystyle` and nothing is pre-picked. The caption says what a path loads (talent order, gear picks, rotation). The first-talent-at-10 line lives in the panel's Pick your path, and the chip says `First talent at 10`. |
| M3 | No raid promise to a guildless leveller | **Fixed.** Companion: `Uploads your fights for reports and the coach, and keeps your builds in step with the site.` |
| M4 | "Phase 2" confusion, the "…" row, the faction banner, where the Data addon comes from | **Fixed.** The example job is `Whelps, west cave`, from the Onyxia template. The tooltip is real. The banner is the player's own faction's (spec; the mock character is Alliance). Data addon: `beside it on CurseForge and Wago`. |
| S1 | Cut the QR | Kept (owner's ask), tagged `coming`. |
| S2 | Half-empty pages | Pages are fixed height so the buttons never move; accepted. |
| S3 | Drop page 6 | Done. |

## Raider: HOLD

| # | Finding | Disposition |
|---|---|---|
| M1 | In-game answer: the tour says coming, but the panel draws buttons | **Fixed in the spec:** the panel draws no answer buttons until the `raid_signup` and `raid_answer` contracts exist, so the tour's `coming` and the panel agree (§3.1, §4.3). |
| M2 | "Companion is running" is not true from a 2-hour-old ack | **Fixed.** `✓ The companion has run` and `The site saw Bow Jackzon 2 hours ago.` After 3 days: the warning colour, `The site has not seen Bow Jackzon in 4 days. Is the companion running?` |
| M3 | Tiles load a leveling build on a 60 | **Fixed.** Tiles load the spec's build for the level (raid at 60, `onboarding-5-set-up-60.png`); a character with a build shows it with `Change`. |
| S1 | Discord only if linked | Fixed (`if your guild has linked it`). |
| S2 | Logging step; can I hide my score | Logging fixed (one sentence on page 3). Hiding the score: nothing built does it, so the tour claims nothing; raised to the owner as an open question (§3.1). |
| S3 | Cut page 1's slogan | Heading is now `How Forever Sixty works`; the loop stays (owner). |
| S4 | Alliance banner | Faction-aware (spec). |
| S5 | URLs exist; QR at low UI scale; the backwards "page 2" pointer | `/setup` is live; `/app` is being built by the web lane. The QR is drawn at a fixed 112 UI units. The pointer moved into the sign-in row (`Upgrades, talents and rotation work without it.`). |
| S6 | Phase 2 wording | The example is now `Whelps, west cave`; the template wording itself is the guild lane's. |
