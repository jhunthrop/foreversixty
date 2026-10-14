# Capture checklist (L9.1)

One line per capture: every surface of the spec's section 4 table, in every state it has, at level 1, level 20, level 60
and in combat. Run it each week from `docs/superpowers/reviews/README.md`; the personas review what it produces.

How to read a line:

- **Produces** is the file the capture is saved under in the week's folder: `<surface>-<level>-<state>.jpg`, the level
  token being `l1`, `l20`, `l60` or `combat`.
- **Driven by** is the README spike row (`addon/README.md`, the table numbered 1 to 76) or the spec section
  (`docs/superpowers/specs/2026-10-11-addon-v2-design.md`) that says what the state looks like and how to reach it.
- **Week 1** is the status against `docs/superpowers/reviews/week1/chip-captures.md` only. `captured:` names the files that
  note names, as they were saved then (`hunter-N-...`, `chip-...`, `marks...`, `tip...`, `dialog...`, `settings...`); a line
  is never marked captured on any other evidence. `owed:` says what blocks it. Week 1 names stay as saved: they are not
  renamed to the Produces column.
- `l20` means the 20 to 29 band: week 1 captured it on the level 21 warrior Obnoxious Yell and the level 24 hunter Bow
  Jackzon, and a line says which. Nothing in week 1 was captured at level 20 itself, at level 60, or in combat.
- Combat lines say what the spec asks of the surface in combat; where the spec is silent the line records what the surface
  does, so the personas can judge it.
- Fixtures: `[data]` = `addon/fixtures/ForeverSixtyData-l44` is junctioned in; `[inbox]` = `addon/fixtures/ForeverSixtyInbox-l53`
  is (both restored afterwards, see the README).

## Chip (4.1)

| Produces | Driven by | Week 1 |
|---|---|---|
| chip-l1-pick-a-path.jpg | 4.1, level-1 alt | captured: chip2-alt-screen-pick-a-path.jpg, chip2-alt-path-chip.jpg |
| chip-l1-path-card.jpg | 4.1, click the chip | captured: chip2-alt-path-card.jpg, chip2-alt-screen-path-card.jpg |
| chip-l1-path-tooltip.jpg | 4.1 hover | captured: chip2-alt-path-tooltip.jpg |
| chip-l1-next-talent-line.jpg | 4.1 state 7, below level 10 the muted next-talent line | captured: chip3-alt-next-line.jpg, chip3-alt-tooltip.jpg |
| chip-l20-point-to-spend.jpg | 4.1 state 1, ring and "+N" | captured: chip5-point-state.jpg, chip6-point-rank.jpg, chip-point.jpg (level 21 warrior) |
| chip-l20-point-tooltip.jpg | 4.1 hover, every waiting state | captured: chip5-point-tooltip.jpg, chip-tooltip.jpg (level 21 warrior) |
| chip-l20-point-ring-pulse.jpg | 4.1 pulse 0.4 to 1.0 over 1.2 s, static after 10 s | owed: stills cannot show the pulse (chip5-point-after-6s.jpg looked the same); needs a burst of timed shots or a short clip |
| chip-l20-upgrade-in-bags.jpg | 4.1 state 2; README rows 37, 39 | captured: hunter-16-gloves-in-bag-chip-equip-line.jpg, hunter-18-blizzard-bag-screen.jpg (level 24 hunter, a worn piece bagged) |
| chip-l20-loot-here-for-you.jpg | 4.1 state 3; row 44 | owed: needs a run into The Deadmines (or another dungeon) at 20 to 29 |
| chip-l20-raid-check.jpg | 4.1 state 4; rows 66, 35 | captured: hunter-37-3-chip-ready-check.jpg, hunter-39-2-chip-clean-tick.jpg (`/fs ready` dry run, solo; the real leader's ready check is chip-l60-raid-check-missing) |
| chip-l20-last-night.jpg | 4.1 state 5; row 73 `[inbox]` | captured: hunter-40-last-raid-chip-tooltip.jpg, hunter-43-loot-lines-and-chip-short-form.jpg, hunter-44-last-raid-hover-with-spec.jpg |
| chip-l20-build-complete.jpg | 4.1 state 6, check icon for 5 s | owed: needs the last point of a build spent; no week 1 file names it |
| chip-l20-idle.jpg | 4.1 state 7, class icon at 40 % | captured: chip4-main-idle.jpg, chip3-main-idle.jpg (level 21 warrior) |
| chip-l20-idle-tooltip.jpg | 4.1 hover, headed by the build in force | captured: chip4-main-tooltip.jpg (level 21 warrior); hunter-11-chip-tooltip-build.jpg (level 24 hunter, build named) |
| chip-l20-legend-tooltip.jpg | 4.4 marks' legend, two lines | captured: hunter-19-chip-legend.jpg, hunter-20-chip-legend.jpg, hunter-22-chip-legend-colours.jpg; off: hunter-21-chip-legend-off.jpg |
| chip-l20-plus-n-badge-two-states.jpg | 4.1 "+N" badge when more than one state waits | captured: chip-point.jpg (level 21 warrior, "+1" with a bag upgrade also waiting) |
| chip-l20-right-click-hide.jpg | 4.1 "Hide until next level" | captured: chip-hidden-after-right-click.jpg, chip2-hidden-chat-line.jpg, chip2-fs-chip-restore.jpg |
| chip-l20-shift-drag-move.jpg | 4.1 shift-drag | owed: the capture tool could not shift-drag (settings round 2 note); needs a real shift-drag |
| chip-l20-stands-aside-talent-window.jpg | 4.1 / 4.2 chip stands aside while the window is open | captured: marks4-screen.jpg, marks4-chip-after-close.jpg, settings3-6-chip-hidden-under-talents.jpg, settings3-6-chip-back-after-close.jpg |
| chip-l20-hidden-by-switch.jpg | 4.11 "Show the chip" off | captured: settings-4-chip-gone.jpg, settings-4-chip-back.jpg |
| chip-l60-idle.jpg | 4.1 state 7, no text from level 10 | owed: needs a level-60 character |
| chip-l60-upgrade-in-bags.jpg | 4.1 state 2 with sim weights | owed: needs a level-60 character and spare gear that beats a worn piece |
| chip-l60-raid-check-missing.jpg | 4.1 state 4; rows 66, 67 | owed: needs a level-60 character, a group and a real ready check with a missing imbue |
| chip-l60-last-night.jpg | 4.1 state 5 from a logged raid | owed: needs a level-60 character (fixture `[inbox]` last_raid is keyed to Bow Jackzon) |
| chip-combat-faded.jpg | 4.1 fades to 25 % in combat, ignores the mouse | owed: needs a hostile mob (round 1: no mob in the city) |
| chip-combat-switch-off.jpg | 4.11 "Hide the chip in combat" off, so only faded | owed: needs a hostile mob |
| chip-combat-equip-disabled.jpg | 4.1 state 2 disabled with "In combat" under lockdown | owed: needs a hostile mob and a bagged upgrade |

## Talent marks (4.2)

| Produces | Driven by | Week 1 |
|---|---|---|
| marks-l1-window-no-points.jpg | 4.2 at level 1, nothing to spend | owed: not shot; the level-1 alt (Shorty Small) exists and needs a session |
| marks-l20-ring-badges-strip.jpg | 4.2 ring, order numbers, strip; rows 22, 37 | captured: marks5-screen.jpg, marks5-trees.jpg, marks5-strip.jpg (level 21 warrior, 12 unspent) |
| marks-l20-range-badges.jpg | 4.2 point range "7 to 11" | captured: marks5-arms-badges.jpg, marks5-fury-badge.jpg |
| marks-l20-run-badges.jpg | 4.2 non-consecutive runs, "8 to 11" then "36" | captured: hunter-15-lethal-attacks-badge-8-11.jpg, hunter-15-lethal-attacks-tooltip-runs.jpg (level 24 hunter) |
| marks-l20-node-tooltip.jpg | 4.2 node tooltip lists the runs | captured: marks4-node-tooltip.jpg, hunter-15-lethal-attacks-tooltip-runs.jpg |
| marks-l20-off-plan.jpg | 4.2 point off plan, count warning-coloured | captured: marks4-offplan-strip.jpg, marks4-offplan-node-tooltip.jpg, marks3-staged-off-build-strip.jpg |
| marks-l20-zero-unspent-next.jpg | 4.2 strip with 0 unspent | captured: hunter-11-talents-saved-build.jpg (level 24 hunter, "Next: Lethal Attacks") |
| marks-l20-no-build.jpg | 4.2 with no build followed | captured: hunter-4-talents-no-build.jpg |
| marks-l20-switch-off.jpg | 4.11 "Marks on the talent window" off | captured: settings2-6-talents-no-marks.jpg, settings2-6-talents-marks-back.jpg |
| marks-l20-first-open-after-reload.jpg | 4.2 first open finds the buttons; row 22 | captured: marks2-screen.jpg (fixed in round 2), marks1-second-open-trees.jpg (the pre-fix miss) |
| marks-l60-window-complete.jpg | 4.2 at 51 of 51 points, "0 off plan" | owed: needs a level-60 character with a spent build |
| marks-combat-window.jpg | 4.2 talent window opened in combat | owed: needs a hostile mob; spec silent, record what the client allows |

## Item tooltip (4.3)

| Produces | Driven by | Week 1 |
|---|---|---|
| tip-l1-worn-item.jpg | 4.3 on a level-1 alt's worn item | owed: not shot; needs a session on the level-1 alt |
| tip-l20-bag-downgrade.jpg | 4.3 verdict "Downgrade" with the breakdown line | captured: tip8-mace.jpg, tip7-mace.jpg (level 21 warrior) |
| tip-l20-bag-upgrade.jpg | 4.3 verdict "Upgrade +N over ..." | owed: needs spare gear that beats a worn piece (round 1 to 8 notes "an Upgrade in a bag" never captured) |
| tip-l20-bag-sidegrade.jpg | 4.3 "Sidegrade (+0.3 DPS)" | owed: needs spare gear within 1.0 of a worn piece |
| tip-l20-fills-empty-slot.jpg | 4.3 "Fills an empty slot" | captured: hunter-10-bag-gloves-tooltip.jpg, hunter-16-gloves-in-bag-chip-equip-line.jpg |
| tip-l20-not-for-you.jpg | 4.3 "Not with your two-hander" / "Not for you until 40: plate armor" | captured: tip7-shield.jpg ("Not with your two-hander"); owed: the plate line ("until 40"), needs a plate item for a non-plate class |
| tip-l20-worn-next-upgrade.jpg | 4.3 Next upgrade, the gain on its own row, Also good rows | captured: hunter-20-waist-tooltip.jpg, hunter-20-ranged-tooltip.jpg, hunter-22-ranged-also-good-rows.jpg, hunter-22-waist-also-good-rows.jpg |
| tip-l20-worn-best-until.jpg | 4.3 "Best until level 30" with the At line and source | captured: tip8-chest.jpg, tip7-chest.jpg, hunter-5-back-tooltip.jpg, hunter-19-mainhand-header-fits.jpg |
| tip-l20-next-compare-tooltip.jpg | 4.3 the next upgrade's compare tooltip beside it | captured: hunter-14-ranged-next-upgrade.jpg, hunter-16-empty-hands-hover-next.jpg |
| tip-l20-no-pick.jpg | 4.3 "No pick for your band yet" | captured: hunter-12-mainhand-no-pick-line.jpg, hunter-16-empty-trinket-no-pick.jpg |
| tip-l20-header-equipped-collision.jpg | 4.3 "Equipped" on its own row when the header is wide | captured: hunter-19-waist-hover-equipped-collision.jpg, hunter-19-ranged-hover-equipped-collision.jpg (pre-fix), hunter-20-waist-tooltip.jpg (fixed) |
| tip-l20-estimate-header.jpg | 4.3 "Estimate" tag | captured: hunter-5-back-tooltip.jpg, hunter-9-hands-corner-tooltip.jpg |
| tip-l20-your-sim-header.jpg | 4.3 "Your sim" tag; row 47 `[inbox]` | captured: tip7-chest.jpg, tip7-worn-mainhand.jpg (level 21 warrior, controller's sim-weights fixture) |
| tip-l20-enchant-not-counted.jpg | 4.3 "(enchant not counted)" | captured: tip4-chest.jpg (pre-fix text) |
| tip-l20-non-gear-blank.jpg | 4.3 non-gear shows no section | captured: tip1-hearthstone.jpg |
| tip-l20-in-your-bags-hover.jpg | 4.4 empty slot hover "In your bags" | captured: hunter-19-in-your-bags-hover.jpg |
| tip-l20-vendor-window.jpg | 4.3 tooltip in a vendor window | owed: week 1 notes name it as never captured; needs a vendor visit |
| tip-l60-worn-sim.jpg | 4.3 worn item, "Your sim", Next upgrade with sim weights | owed: needs a level-60 character (row 47) |
| tip-l60-best-for-now.jpg | 4.3 "Best for now" in the top band | owed: needs a level-60 character; tip3-worn-ranged.jpg showed it once on the level 21 warrior with the pre-fix wording |
| tip-combat-hover.jpg | 4.3 tooltip while in combat | owed: needs a hostile mob |

## Paper doll, bag marks, quest rewards, rolls (4.4)

| Produces | Driven by | Week 1 |
|---|---|---|
| slots-l1-paper-doll.jpg | 4.4 at level 1 | owed: not shot; needs a session on the level-1 alt |
| slots-l20-empty-slot-arrow.jpg | 4.4 green arrow on an empty slot; rows 37, 39 | captured: hunter-16-empty-hands-arrow.jpg, hunter-20-hands-arrow-4x.jpg, hunter-18-empty-hands-arrow-fullres.jpg |
| slots-l20-bag-arrow.jpg | 4.4 arrow on the bag item; rows 35, 36 | captured: hunter-20-bag-arrow.jpg, hunter-18-blizzard-bag-mark.jpg |
| slots-l20-beats-worn-arrow.jpg | 4.4 arrow on a filled slot a bag item beats | owed: needs spare gear that beats a worn piece (the captures so far use an empty slot) |
| slots-l20-gold-corners.jpg | 4.4 target corner on a slot with a next upgrade; row 37 | captured: hunter-14-right-column-corners-fullres.jpg, hunter-15-corners-1s-after-reload.jpg, hunter-15-corners-6s-after-reload.jpg |
| slots-l20-reequip-clears.jpg | 4.4 marks clear on equipping | captured: hunter-16-reequipped-arrow-gone.jpg, hunter-18-reequipped.jpg |
| slots-l20-bag-addon.jpg | 4.4 Baganator: no mark, diag says so; row 38 | captured: hunter-16-diag.jpg, hunter-10-paperdoll-empty-hands-no-arrow.jpg |
| slots-l20-switch-off.jpg | 4.11 "Gear tips on item tooltips" off hides the marks | captured: hunter-18-switch-off-pane.jpg, hunter-18-switch-off-bag.jpg, hunter-18-switch-on-pane.jpg, hunter-18-switch-on-bag.jpg |
| slots-l60-paper-doll.jpg | 4.4 at level 60 | owed: needs a level-60 character |
| slots-combat-bag-click.jpg | 4.4 marks in combat; equipping is locked down | owed: needs a hostile mob |
| rewards-l1-quest-reward.jpg | 4.4 "Best for <spec>" on a starter quest's choices; rows 40, 41 | owed: not shot; needs a quest giver with a reward choice at level 1 |
| rewards-l20-quest-best-for.jpg | 4.4 "Best for Arms" on the preferred choice; row 41 | owed: needs a quest turn-in with three reward choices (week 1 notes: never captured) |
| rewards-l20-quest-vendor.jpg | 4.4 "Vendor" muted when none scores; row 41 | owed: needs a quest turn-in where no choice scores |
| rewards-l20-quest-log-no-label.jpg | 4.4 no label in the quest log; row 41 | owed: needs the same quest giver session |
| rewards-l20-roll-frame.jpg | 4.4 verdict line beside the roll frame; rows 42, 43 | owed: needs the user's test group (a group roll on gear) |
| rewards-l60-quest-and-roll.jpg | 4.4 at level 60 | owed: needs a level-60 character and the user's test group |
| rewards-combat-roll.jpg | 4.4 a roll frame that opens in combat | owed: needs the user's test group and a hostile mob |

## Player tooltip and score (4.6)

| Produces | Driven by | Week 1 |
|---|---|---|
| ptip-l1-score-line.jpg | 4.6 at level 1 (unlogged line) | owed: not shot; needs a session on the level-1 alt |
| ptip-l20-own-score.jpg | 4.6 "Forever Score 1240 - Rare - as of 13 Oct"; row 51 `[data]` | captured: hunter-28-own-tooltip-game-shot.jpg, hunter-27-own-tooltip-score.jpg |
| ptip-l20-shift-components.jpg | 4.6 Shift held, five lines and the total; row 53 | captured: hunter-28-shift-real-game-shot.jpg, hunter-27-shift-components.jpg |
| ptip-l20-unlogged.jpg | 4.6 "no logs yet", never a zero; row 52 | captured: hunter-28-unlogged-game-shot.jpg, hunter-27-unlogged-tooltip.jpg |
| ptip-l20-no-data-addon.jpg | 4.6 no score line at all without the data addon; row 52 | owed: not shot; needs the empty `addon/ForeverSixtyData` in place |
| ptip-l20-epic-tier.jpg | 4.6 epic colour (Smoke Gouda, 1610); row 54 | owed: Smoke Gouda was out of target range; needs them in range |
| ptip-l20-live-line.jpg | 4.6 / 4.9 "Level n Spec, gear gap n (live)"; row 54 | owed: needs a guildmate running Forever Sixty online |
| ptip-l20-fs-link.jpg | 4.6 `/fs link` opens the copy tab; row 57 | captured: hunter-27-fs-link-screen.jpg, hunter-27-fs-link-crop.jpg |
| ptip-l20-hook-diag.jpg | rows 48 to 50 | captured: hunter-27-diag-fixture.jpg (diag only; the probes of rows 49, 50 are not named in the note) |
| ptip-l60-score.jpg | 4.6 at level 60 | owed: needs a level-60 character |
| ptip-combat-hover.jpg | 4.6 tooltip on a player in combat | owed: needs a hostile mob |

## Chat badges (4.6, L4.4)

| Produces | Driven by | Week 1 |
|---|---|---|
| badges-l1-guild-chat.jpg | 4.6 badge at level 1 | owed: needs the level-1 alt in a guild or a channel with a scored sender |
| badges-l20-guild-line.jpg | 4.6 "FS 1240 - " before the message; row 55 | captured: hunter-28-guild-badge-fs.jpg, hunter-27-guild-badge.jpg (pre-fix brackets) |
| badges-l20-trade.jpg | 4.6 Trade with two badged names; row 55 | owed: needs a city's Trade channel and Smoke Gouda speaking in range (week 1: hunter not in a city) |
| badges-l20-lfg.jpg | 4.6 LFG channel | owed: needs an LFG channel with a scored sender |
| badges-l20-whisper.jpg | 4.6 incoming whisper keeps the reply target; rows 50, 55 | owed: needs a scored sender whispering (Smoke Gouda) |
| badges-l20-off.jpg | 4.11 "Score badges in chat" off; row 56 | captured: hunter-27-guild-badge-off.jpg |
| badges-l60-guild-chat.jpg | 4.6 at level 60 | owed: needs a level-60 character |
| badges-combat-chat.jpg | 4.6 badge while in combat | owed: needs a hostile mob and a speaking sender |

## "Forever Sixty" character-panel tab (4.5)

| Produces | Driven by | Week 1 |
|---|---|---|
| tab-l1-flyout.jpg | 4.5 at level 1: only the regions that have content | owed: not shot; needs a session on the level-1 alt |
| tab-l20-button.jpg | 4.5 button 112 by 24 under the pane; row 45 | captured: hunter-24-panel-button-screen.jpg, hunter-24-panel-button-crop.jpg |
| tab-l20-button-hover.jpg | row 45 hover hint | captured: hunter-26-button-hover-hint-screen.jpg, hunter-26-button-hover-hint-crop.jpg |
| tab-l20-button-open-state.jpg | row 45 label in body white while open | captured: hunter-26-button-open-state.jpg |
| tab-l20-flyout.jpg | 4.5 targets, rotation, build; row 46 | captured: hunter-24-flyout-screen.jpg, hunter-26-flyout-crop.jpg |
| tab-l20-target-row-hover.jpg | 4.5 next targets, item tooltip and compare | captured: hunter-24-flyout-row-hover.jpg, hunter-26-target-row-hover.jpg |
| tab-l20-rotation-hover.jpg | 4.5 rotation condition tooltip, "at level N" rows | captured: hunter-26-rotation-hover.jpg, hunter-26-rapid-fire-hover.jpg |
| tab-l20-off-plan-hover.jpg | 4.5 build region "3 off plan" | captured: hunter-26-offplan-hover.jpg |
| tab-l20-build-dialogs.jpg | 4.5 "Use a different build", "Copy my character" | captured: hunter-25-build-paste-dialog-screen.jpg, hunter-25-copy-dialog-screen.jpg |
| tab-l20-reopen.jpg | row 46 flyout survives close and reopen | captured: hunter-24-flyout-after-reopen.jpg, hunter-25-button-closed-flyout.jpg |
| tab-l60-sim-last-raid.jpg | 4.5 Sim and Last raid regions; row 47 `[inbox]` | owed: needs a level-60 character |
| tab-l60-sim-me-again.jpg | 4.5 "Sim me again" queues a request | owed: needs a level-60 character and the companion |
| tab-combat-button.jpg | 4.5 / row 69 the button is never made in combat | owed: needs a hostile mob |

## Guild tab (4.7)

| Produces | Driven by | Week 1 |
|---|---|---|
| guild-l1-no-guild.jpg | 4.7 unguilded character | owed: needs a guildless character; not named in the note |
| guild-l20-button.jpg | row 59 button under the guild frame | captured: hunter-29-guild-frame-button.jpg |
| guild-l20-roster-online.jpg | 4.7 roster, online only; row 60 | captured: hunter-29-guild-flyout-online-crop.jpg, hunter-30-guild-flyout-default-crop.jpg, hunter-31-guild-default-crop.jpg |
| guild-l20-roster-scored.jpg | 4.7 score tiers; row 60 `[data]` | captured: hunter-29-guild-flyout-scored-crop.jpg, hunter-30-guild-fixture-tiers-crop.jpg, hunter-30-guild-fixture-mixed-tiers.jpg |
| guild-l20-roster-all-scrolled.jpg | 4.7 97 rows, scroll and count | captured: hunter-30-guild-roster-all.jpg, hunter-30-guild-roster-scrolled.jpg, hunter-31-guild-scrolled-count.jpg |
| guild-l20-sorted.jpg | 4.7 sortable headers | captured: hunter-29-guild-roster-sorted-level.jpg, hunter-30-guild-sorted-spec.jpg, hunter-31-guild-score-desc-game-shot.jpg |
| guild-l20-header-tooltips.jpg | 4.7 Gear, Gap and count line tooltips | captured: hunter-31-guild-gear-tooltip.jpg, hunter-31-guild-gap-tooltip.jpg, hunter-31-guild-live-tooltip.jpg |
| guild-l20-live-row.jpg | 4.7 / 4.9 a member with the addon, live | owed: hunter-30-guild-live-row.jpg was seeded with `ForeverSixtyGuildTab.seedLive`, not a real guildmate; needs a real guildmate with the addon |
| guild-l20-self-row-gap.jpg | 4.7 gap equals the tooltip's biggest upgrade | captured: hunter-33-guild-self-row-4dps.jpg |
| guild-l20-officer-region.jpg | 4.7 officer region; row 61 | owed: needs the user's guild-master character and an inbox `guild` message |
| guild-l20-push-confirm.jpg | 4.7 / 4.9 push confirm, Push and Cancel; row 62 | owed: needs the user's guild-master character |
| guild-l60-roster.jpg | 4.7 at level 60 | owed: needs a level-60 character |
| guild-combat-frame.jpg | 4.7 guild frame in combat | owed: needs a hostile mob |

## Group panel and loot council (4.8, 4.9)

| Produces | Driven by | Week 1 |
|---|---|---|
| group-l1-solo-flyout.jpg | 4.8 solo at level 1 | owed: not shot; needs a session on the level-1 alt |
| group-l20-solo-button.jpg | row 64 button on the Raid tab | captured: hunter-35-group-button-raid-tab.jpg, hunter-37-1-raid-tab-flyout-hint.jpg |
| group-l20-solo-flyout.jpg | row 64 "Not in a group." and the hint | captured: hunter-35-group-flyout-solo.jpg, hunter-38-1-solo-flyout-clear-of-tabs.jpg, hunter-39-1-solo-flyout-59-units.jpg |
| group-l20-friends-tab-rests.jpg | row 64 button rests on other tabs | captured: hunter-37-2-friends-tab-no-button.jpg, hunter-37-4-open-on-friends.jpg, hunter-37-5-raid-tab-wake.jpg |
| group-l20-party-of-two.jpg | 4.8 rows, headers, sort; row 65 | owed: needs the user's test group |
| group-l20-inspect-no-addon.jpg | 4.8 inspect fills a member; row 68 | owed: needs the user's test group (a member without the addon in range) |
| group-l20-ready-check-dry-run.jpg | row 66 `/fs ready` solo | captured: hunter-35-chip-ready-check.jpg, hunter-37-3-chip-ready-check.jpg |
| group-l20-diag.jpg | rows 63, 64 `/fs diag` group line | captured: hunter-35-diag-group.jpg |
| group-l60-ready-check-imbue.jpg | 4.8 ready check lists missing imbue; row 66 | owed: needs a level-60 character and a group |
| group-l60-coverage.jpg | 4.8 Coverage list; row 67 | owed: needs a level-60 character in Molten Core, Onyxia's Lair or Blackwing Lair |
| group-combat-inspect-paused.jpg | 4.8 "inspect paused" in combat; row 69 | owed: needs the user's test group and a hostile mob |
| loot-l20-fs-loot-solo.jpg | 4.9 `/fs loot` on and off the list; rows 70 to 72 | captured: hunter-40-fs-loot-solo.jpg, hunter-41-fs-loot-uncached-and-planned.jpg |
| loot-l20-drop-block.jpg | 4.9 master looter's "Drop:" region; row 74 | owed: needs the user's test group |
| loot-l20-member-no-addon.jpg | 4.9 "waiting", "(inspect)", "no addon"; row 75 | owed: needs the user's test group |
| loot-l20-second-drop-stack.jpg | 4.9 stacking, "and N more", 60 s clear; row 76 | owed: needs the user's test group |
| loot-l60-drop-block.jpg | 4.9 at level 60 | owed: needs a level-60 character and the user's test group |
| loot-combat-drop.jpg | 4.9 a drop that arrives in combat | owed: needs the user's test group and a hostile mob |

## Build dialog (4.10)

| Produces | Driven by | Week 1 |
|---|---|---|
| dialog-l1-paste-empty.jpg | 4.10 empty, at level 1 | owed: not shot; needs a session on the level-1 alt |
| dialog-l20-paste-empty.jpg | 4.10 "Following", "Next talent", paste box | captured: dialog3-1-empty.jpg, dialog2-1-empty.jpg, hunter-25-build-paste-dialog-screen.jpg |
| dialog-l20-paste-preview.jpg | 4.10 good code, preview, "Use this build" | captured: dialog3-2-preview.jpg, dialog2-2-preview.jpg |
| dialog-l20-site-row.jpg | 4.10 site build row, gold edge | captured: dialog7-1-row.jpg, dialog7-2-row-preview.jpg, dialog6-1-row.jpg, dialog6-2-row-preview.jpg |
| dialog-l20-use-pressed.jpg | 4.10 Use hover and pressed | owed: dialog5-3-use-hover.jpg and dialog5-3-use-pressed.jpg show no visible change under the capture tool; needs a human hover and press |
| dialog-l20-use-result.jpg | 4.10 "Now following ..." and the chip | captured: dialog-3-chat.jpg, dialog-3-chip.jpg, dialog-7-use-chat.jpg |
| dialog-l20-refused.jpg | 4.10 bad code, advice line | captured: dialog4-4-refused-advice.jpg, dialog3-4-wrong-class.jpg, dialog-4-refused.jpg |
| dialog-l20-copy.jpg | 4.10 Copy my character, selected export | captured: dialog7-5-copy.jpg, dialog5-5-copy.jpg, hunter-42-copy-dialog-hunter.jpg |
| dialog-l20-copied.jpg | 4.10 green "You copied ..." | captured: dialog6-5-copied.jpg, dialog5-5-copied.jpg, dialog3-5-copied.jpg |
| dialog-l20-ack-states.jpg | 4.10 "Seen by the site", "Yours is newer", "Waiting for the site" | captured: dialog-6-seen.jpg, dialog6-5-copy.jpg, dialog7-5-copy.jpg |
| dialog-l20-sim-line.jpg | 4.10 "Sim queued for ..." / "Sim done ..." | captured: dialog8-5-sim-line.jpg, dialog5-5-copy.jpg |
| dialog-l20-no-companion.jpg | 4.10 "No companion found" | captured: hunter-42-copy-dialog-hunter.jpg, hunter-25-copy-dialog-screen.jpg |
| dialog-l20-tabs-escape.jpg | 4.10 Tab switches, Escape goes back then closes | captured: dialog7-tabs-paste.jpg, dialog7-escape-empty.jpg, dialog-10-tab-paste.jpg |
| dialog-l20-waiting-badge.jpg | 4.10 "1 waiting" in its own slot | captured: dialog7-1-row.jpg, dialog6-5-copy.jpg |
| dialog-l20-keyboard-free.jpg | row 24 movement keys after Escape | owed: the notes say keys were free (round 1 facts) but name no capture |
| dialog-l60-copy-sim-done.jpg | 4.10 "Sim done from the ... copy" at 60; row 47 | owed: needs a level-60 character |
| dialog-combat-open.jpg | 4.10 opening `/fs` in combat | owed: needs a hostile mob |

## Settings (4.11)

| Produces | Driven by | Week 1 |
|---|---|---|
| settings-l1-page.jpg | 4.11 at level 1 | owed: not shot; needs a session on the level-1 alt |
| settings-l20-page-top.jpg | 4.11 eight switches, title warning; rows 17, 18 | captured: settings4-1-page-top.jpg, hunter-27-settings-eight-switches.jpg, hunter-27-settings-page-crop.jpg |
| settings-l20-scroll-foot.jpg | 4.11 scroll to the foot, reopen at top | captured: settings4-2-foot-reached.jpg, settings4-3-reopened-at-top.jpg |
| settings-l20-reset-positions.jpg | 4.11 Reset positions, "Positions reset" | captured: settings3-4-reset-pressed.jpg, settings3-4-reset-after.jpg, settings2-7-button-moved.jpg, settings2-7-reset-after.jpg |
| settings-l20-warnings.jpg | 4.11 older patch and data too old | captured: settings4-6-addon-data-warning-title.jpg, settings4-6-both-warnings-foot.jpg, settings3-5-title-warning.jpg |
| settings-l20-data-not-installed.jpg | 4.11 "not installed" foot | owed: spec-only so far (settings round 4 facts); needs the empty `addon/ForeverSixtyData` in place |
| settings-l20-data-current.jpg | 4.11 "current" foot with the fixture `[data]` | owed: hunter-27-diag-fixture.jpg shows only the diag line; needs the page's foot with the fixture in place |
| settings-l20-share-off.jpg | 4.11 / 4.9 share off, `/fs hello` refused | captured: settings4-4-share-off.jpg, settings4-4-hello-refused.jpg |
| settings-l20-gear-tips-hint.jpg | 4.11 "Gear tips" hint with the player-tooltip clause | captured: hunter-28-settings-gear-tips-hint.jpg |
| settings-l20-help-commands.jpg | 4.11 `/fs help` and the command list | captured: settings4-5-fs-help.jpg, settings3-8-fs-help.jpg |
| settings-l60-page.jpg | 4.11 at level 60 | owed: needs a level-60 character |
| settings-combat-open.jpg | 4.11 opening the page in combat | owed: needs a hostile mob; the client may block the Settings window, record that |

## Launcher (4.1, 4.11)

| Produces | Driven by | Week 1 |
|---|---|---|
| launcher-l1-minimap-button.jpg | 4.11 minimap button at level 1 | owed: not shot; needs a session on the level-1 alt |
| launcher-l20-minimap-tooltip.jpg | 4.11 minimap button hover; row 26 | captured: settings-6-minimap-tooltip.jpg, settings2-2-minimap-hint-glyph.jpg |
| launcher-l20-minimap-moved.jpg | 4.11 drag and reset | captured: settings2-7-button-moved.jpg, settings2-7-button-back.jpg |
| launcher-l20-left-click-dialog.jpg | 4.11 left click opens the build dialog | owed: no file named; needs a left click on the button |
| launcher-l20-right-click-settings.jpg | 4.11 right click opens the Settings page | captured: settings-1-screen.jpg, settings2-1-screen.jpg |
| launcher-l20-fs-options.jpg | 4.11 `/fs options` | captured: settings2-8-fs-options.jpg |
| launcher-l20-compartment-entry.jpg | 4.11 addon compartment entry; row 26 | owed: no file named; needs the compartment menu open |
| launcher-l20-keybind.jpg | row 27 "Toggle Forever Sixty" in Key Bindings | owed: no file named |
| launcher-l20-addons-list-icon.jpg | row 28 LX seal in the AddOns list | owed: no file named |
| launcher-l60-minimap-button.jpg | 4.11 at level 60 | owed: needs a level-60 character |
| launcher-combat-minimap-button.jpg | 4.11 minimap button in combat | owed: needs a hostile mob |
