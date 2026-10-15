# Rulings I made

Compiled from `.superpowers/sdd/2026-10-11-addon-v2-lanes/progress.md` (634 lines). Bracketed numbers are ledger line numbers. Where the ledger gives a reason or a cost-if-wrong, it is included; where a line is ambiguous it is quoted. Reversals are written "X, later reversed to Y".

## Rulings the user made

- USER RULING (2026-10-14): push only when everything is finished, as one complete deploy, with no interim pushes of the lane branches. End state: all lanes merged into addon/v2-foundation, that merged onto main (with the 70374 build move), and main pushed once so the deploy runs complete. [614]
- This settled the earlier "Ruling pending the user: push the lane branches as a backup; the merge onto main (70338) is its own step after the running tasks land" (the ledger notes none of the v2 branches had ever been pushed and local main was 27 behind origin/main). [579]
- Quoted user question 2026-10-13: "shouldn't we have default builds that ship with the addon...?" The controller's answer was that they do (`Data.defaultBuilds`, `Follow.ensureDefault`), but adoption fails on the live client; the follow-up was a new task L3.8. [401]
- User report 2026-10-13: a Lua error on every open of the character pane ("TextStatusBar.lua:110: attempt to compare a secret number value (execution tainted by ForceTaint_Strong)"). The controller ruled that the addon never writes into a Blizzard frame (see L4.1 fix round 2). [402]
- User report: the Lethal Attacks badge reads "8-36". The controller ruled that a badge names the next run ("8-11", then "36") and the tooltip lists every run (see L3.2-fix-runs). [422]
- User statement (quoted): "a new build version is out". The controller found the client at 1.60.1 70374 and ruled to move straight to 70374 (see cross-lane). [580, 602]
- User confirmed they were not playing, so the client was free for checks again; the user also disabled Baganator for the bag captures and said it may be re-enabled afterwards. [433, 444, 445]
- User note: on this user's setup the character pane binding is "=" rather than "c"; the user rebound the pane to "c" for the live taint verification. [409, 412]

## L0 contracts and fixtures

- The implementer's specified details stand: `|`, `\n` and `%` percent-encoded in FS1 fields, a 255-byte message limit, `saved_at` as local "YYYY-MM-DD HH:MM", `export_seen_at` as RFC 3339 UTC. The brief left them open and these are the client's own limits; cost if wrong is a one-line contract edit. [29]
- The companion keeps emitting the legacy `builds` list beside `build` messages for one release; cost if wrong is duplicate build offers in the dialog for one release. [29]
- `unit` is "dps" only, meaning healing per second when `dpsPerPoint.healing` is true; cost if wrong is a label change. [29]
- `snapshot?`/`snapshot!` go on GUILD when sent from the guild tab and on RAID/PARTY when sent from the group panel; `share` and `push` carry FSB1 build codes; the doc notes the FS1-prefix vs FS1:-export name collision. [30]
- `sim_request.requested_at` is local "YYYY-MM-DD HH:MM" like the other outbox times; cost if wrong is a format note. [34]
- The extra outbox tests the L0.2 implementer added stand because they pin contract 4 from both sides; cost if wrong is two tests to delete. [34]
- L0.2 re-review: findings 1 and 3 addressed, finding 2 open, so fix round 2 asked for every fixture key and recursive embeds (plus apt-get update). [39]

## L1 data lane (worktree fs-data-v2)

### L1.1 default builds
- The point order is greedy and never removes a point: one legal point per level, the first cell (own tree first, tier, column) where that level's ladder build has more ranks than the running order, else the first level-60 target cell whose prerequisites are met. It must end on the guide's 51 points or error; `LadderTalentString` is untouched (it defines the site's bands); divergent levels are recorded per spec. Cost if wrong: a handful of levels on 8 specs differ from the site's own bands until the ladder is fixed. [49]
- The legality check (tier gate, prerequisite, max rank) applies to the first choice as well as the fallback, so the paladin Holy level-33 band is illegal under planner rules; cost if wrong: none. [54]
- Subagent commits carry their own attribution trailer; the controller rewrites trailers to the session's Fable 5.1 line on each lane branch before merge (git filter-branch --msg-filter on the unpushed lane); cost if wrong: none. [52]
- Site data defect handed to the site lane: the ladder's level-33 paladin Holy build ranks Divine Precision before its prerequisite Holy Shock. [52]

### L1.2 BiS picks
- Legacy band files' existing `alternatives` are the fallback when a pick has no runners_up; no manual full leveling-bis regeneration (the nightly job fills chance/runners_up on its next run); the warrior-arms proof run stays committed. Cost if wrong: one spec's picks move a day early. [61]
- Data.lua must load under the 2 MB idle budget, so each BiS pick is packed as one `|`-joined string on disk and `Gear.pick` decodes it to contract 1's logical table with a memo cache; contract 1 gains a physical-encoding note. Cost if wrong: a decode step on hover. [63]

### L1.3 weights and facts
- Carry to L3.3/L4: multiply deltas by dpsPerPoint only with the static `weights` table (same reference stat), never the band's bis_weights. [81]
- Carry to L5.2: match ready-check auras by spell NAME (one rank listed per buff); later widened to match by family name including Arcane Brilliance / Prayer of Fortitude / Gift of the Wild / Greater Blessings. [81, 82]

### L1.4 names and weapon DPS
- The packed pick string stays (extending it would break every alternative in `Gear.decodePick`); two side tables `bisNames[id]` and `bisDps[id]` (weapons only, one decimal, from items/<class>.json `dps`) cover every pick and alternative id. The Fable 5.1 trailer stands for every commit in this plan. [242]

### L1.5 per-band dps-per-point
- L1.5 carries the sim ladder's per-band figure into `dpsPerPoint[spec][band]` (brief amended), because the figure was one per spec measured at level 60 and applied at every level. [259]
- The geared per-band measurement stands over the ladder's weapon-only figure (the ladder is about half at bands 30-40 and has no weapon term). Every band differs by more than 25%, so `dpsPerPoint[spec][band] = { reference, dps, weaponDps }` ships with the spec-level level-60 keys kept for the current reader. [324]

### L1.6 pick coverage
- L1.6 closed at 5e1129cc; lane/data-v2 merges into addon/v2-foundation once L3.5 fix 1 commits, then L3.6 is dispatched. The three low findings go to the L1.x polish list. [360]

### L1.7 rare spawns
- The pipeline may fetch creature rank from its already-pinned classic-db source (`creature_template.Rank` at pin ec4f5961) onto loot.json boss rows, with no hand list. The instance-only label for multi-npc trash drops stands ("Gnomeregan"), not "World drop"; the rare-spawn rule is as briefed. [450]
- Rare-spawn rule = rank 2/4 AND a pooled or open-world spawn, because rank alone would drop VanCleef and the other rank-2 bosses. [480]
- Fix round 1: count only pickable kinds as "other" sources; positive-score coverage with two for paired slots; per-slot exclusion as the contract words it. [487]
- Trinkets with an effect count as rivals even at score 0, accepted. For D3 the implementer's rule is accepted (the instance label only for a shared trash drop whose published name is one of its creatures) over an id list. [515]
- Fix round 2: the rare rule sees the pool the limits leave; tests for the dagger case and the PvP cap; `fillsSlot` narrowed to trinkets. [516]
- The ranker is not run-to-run deterministic (exact-score ties flip); the committed files stand, with a deterministic tie-break (item id ascending after score) carried to L9.2 as a must before release. [531]
- Fix round 3 made now: sorted stat summation, one rounding of scores and dps_delta, determinism tests, and a double regeneration proving byte-identity. [534]

## L2 addon foundation

### L2.1 rules
- The two Skyborne races' class pairs are allowed exactly as data/builds/<build>/combos.json lists them (the "six new pairs" fact describes base races only); cost if wrong: a Skyborne alt sees a false "not for you". [48]
- Unknown item classes return "ok" (never refuse what cannot be judged) and the optional `slot` argument stands; cost if wrong: none visible. [48]

### L2.2 progress
- Includes `Progress.unspentPoints()` (moved from Toast.lua verbatim) so L3.1 can delete Toast; cost if wrong: none, pure move. [22]
- `Gear.upgrades` (bag vs planned) stays until L2.6 deletes GearView; `Progress.bagUpgrades` (bag vs worn) is the chip's and marks' answer. Two questions, two functions, named apart; cost if wrong: one dead function for a week. [53]
- A differing item that cannot be scored reads "behind" with delta nil (the marks draw a dash, no number); cost if wrong: a dash instead of nothing on FS1-code plans. [53]

### L2.3 score and ratings
- No format-1 fallback: the addon and the data addon release together and a format-1 file reads "Update both addons". Cost if wrong: in-game ratings are quiet on this machine until L6.4 ships format 2. [58]
- Score.lua sits before Ratings.lua after Rotation.lua in the TOC (it requires Theme); cost if wrong: none. [58]

### L2.4 comms
- Theme.lua is the only file that may name the addon-message API (no_network_spec allows it there); cost if wrong: none. [66]
- Push verification reads the SENDER's rank from the guild roster or group-leader status (not the receiver's own rank); correct as built. [66]

### L2.6 follow / outbox
- `Follow.ensureDefault(data)`: when a character has no saved build in any slot, load `Data.defaultBuilds[<spec or class default>].leveling` (raid at 60) and mark it `default = true` so the dialog can say "Using the default Arms leveling build". Spec section 5.1 requires zero-setup and the plan omitted the loader; cost if wrong: a wasted default load the player replaces by pasting. [21]
- The outbox (including a sim request) is written only when auto-save is on, same as the export: one switch for everything the companion carries. Cost if wrong: a sim request lost when auto-save is off. [83]
- In-client check blocked: the controller closed two clients it launched by mistake through open_application, one of them the user's session. Lesson ruled: never call open_application for the game; raise the existing window by handle instead; the controller never logs in. [85]

## L3 addon surfaces

### L3.1 chip
- Chip state 0 "pick a path" at level 1-9 with no spec detectable: three class-coloured buttons that call `Follow.ensureDefault` for the chosen spec (spec section 5.1); cost if wrong: one extra chip state to remove. [23]
- `Progress.picksHere` (pure, matches pick source `<instance>: <boss>`) is L3.1's and L4.2 consumes it instead of adding `Gear.picksHere`; `Gear.bandFor` lifted from Tooltip. [89]
- Theme.GLOW_PULSE takes the spec's numbers (0.4-1.0 over 1.2 s, static after 10 s via a third showGlow argument). [89]
- `raidcheck` is fed by `Chip.setRaidCheck` (L5.2 calls it) and never shows until then. [89]
- `lastnight` names the boss with the lowest percentile, once per login. [89]
- combatHide true hides, false fades to 0.25 and ignores the mouse. [89]
- The pick-a-path card is the chip's `path` state with 28 Locale playstyle lines; `Follow.ensureDefault` accepts a forced spec. [89]
- Below level 10 the point state reads "Take X at level 10" without the ring so the path choice is answered at once. [99]
- `hiddenFor`/`pathFor` per-character pref fields stand (prefs are account-wide); the two muted tooltip hint lines stand; the druid's four path buttons stand (feral bear is its own spec). [99]
- Fix round 1: unspent points are pure (level-9 minus spent) and the C_Traits reader is deleted; the point state is absent with 0 points and the tooltip carries "Next at 22: X". [107]
- Off-hand candidates are skipped under a two-hander, a two-hander is scored net of both hands, floor 1.0 per unit. [107]
- The chip grows to 280 then drops the verb (spec 4.1 edited); ring on the icon, frame never gold; 14x14 badge; tooltip hierarchy; `/fs chip` and a chat line on hide; level-60 hide lasts until `/fs chip`; diag silent when the data addon is absent. [107]
- Fix round 2: nextLevel = 10 + Progress.pointsSpent(ranks) when no point is unspent; `levelForPoint` stays for plan labels and the talent-marks strip uses the same function. [109]
- Idle keeps the chip's dark ground and border at full alpha, only the icon at 40 %. [110]
- nextPointLevel = max(level+1, 10 + spent) replaces 9 + order index; the informational next line never counts toward the badge; ellipsis cuts the name not the gain; diag names combat/cinematic hides. [111]
- Raid check and build complete count toward the badge (states the player should notice); only the informational next line is excluded. [113]
- Fix round 3: nextPointLevel returns nil at max level; the keep branch falls back to the ellipsised full line; in combat with combatHide=false the icon alpha is 1 inside the 25 % frame. [115]
- The "Easiest to level" tag is dropped until the data carries a flag (carry to the data lane: a per-class leveling pick from the sim). [120]
- A data addon that is not installed is not a fault: one info line outside "Not working:" ("The Forever Sixty Data addon is not installed; scores show once it is."). [120]
- Fix round 4: drop the tag; register GLOBAL_MOUSE_DOWN once; a refusal is a fault; two nits. [126]
- Fix round 5: diag calls Score.status and spec; the in-combat double "In combat"; a spec for the click-outside refusal. [130]
- Persona round 2 checks: Improved Eviscerate being a Combat talent and Improved Gouge Assassination is correct for this client's Forever trees, so the post-choice line is right and not a data bug. [116, 117]

### L3.2 talent marks
- `Progress.cells(build, ranks, data)` with a next field stands; the TOC order (after Chip) stands; the cap-level strip line without "next at" stands. [137]
- Ruling 7 revised: refresh on show and once deferred; "every node missing" = not built yet (no latch, no note); latch only on a partial miss persisting past the deferred pass. The strip is a left-aligned pill clear of Apply Changes, raised above the window's children, with a one-time diag info line of its geometry. [141]
- The strip hangs on TalentsFrame one strata up (absent on the spellbook tab: wanted). [147]
- The TalentGlow geometry info lines are dev noise even on `/fs diag`, so they go behind `/fs diag verbose`. [150]
- `Follow.buildName` (tab-split fallback) shared by the strip and the minimap stands. [159]
- Fix round 3: "off plan" (spec wording) with the name still "... build"; the apply hint is dropped; segments are never ellipsised (sacrifice order: hint, spend, name). [164]
- Range badges for multi-rank cells (spec 4.2 amended: "a small range of the plan's points on every node the plan touches"), the next-node segment in the strip, and dim digits at 0.75. [172]
- L3.2-fix-runs: a badge names the next run ("8-11", then "36") and the tooltip lists every run; pure `runsFor`/`nextRun` with the hunter's exact order pinned. Closed; merge after the L4.1 redirect commit. [422, 427]

### L3.3 item tooltip
- Delete the tests that pin the old verdict wording, sections shape, band-weights preference and the old BiS row; rename-only for surviving behaviour; hovered item passed as plan with the gain read as minus the delta; U+2212 with an ASCII fallback; other sources as "<source> · <kind>". [187]
- A two-hander compared against the main hand only in the tooltip stands for now (carry to L4.1: a paired-slot verdict shared by chip and tooltip); "cannot use dual wielding" to be reworded "cannot wield two weapons". [192]
- The tooltip applies the two-hander rule ("Not with a two-hander" muted line, no verdict). [195]
- Saved-variables defect: Export.outbox uses the same name and realm as Export.characterKey (full name), and the characters key carries the region ("US") from GetCurrentRegionName/portal (guarded); hello and inbox matching are unaffected. [199]
- Fix round 1: tag "Estimate"/"Your sim" with the spec named (spec 4.3 amended); Upgrade only at 1.0 or more (the chip's threshold); dps_per_point under a sim message only from the message; the outbox character is the full who= name; characterKey carries the region and sameCharacter treats an empty region as a wildcard. [203]
- ForeverSixtyInbox was declared in the main addon's SavedVariables, so no inbox message ever reached the client. A second code-free addon `addon/ForeverSixtyInbox` owns that file; the main addon drops it from SavedVariables and lists `## OptionalDeps: ForeverSixtyInbox, ForeverSixtyData`; packaging ships both addons. [211]
- "Next upgrade: <name> +X DPS" only at or above MIN_GAIN, else "Best for now"; gains compare the items' own stats and enchants are not counted; the "Estimate" tag stays per 4.3; a worn item with no band pick shows the header plus "Equipped" only. [214]
- The addon adopts the sim's leveling-bis rule for weapons: when the reference stat is attack power the weapon's DPS converts at 14 AP per DPS (`attackPowerPerDPS`) and is weighted by that stat's weight; casters and healers ignore it; chip and tooltip share it through `Progress.weaponVerdict`; spec 4.3 gains one sentence. [217]
- D5 (addon-release.yml ships only ForeverSixty/) is fixed in this task, not L9.2: both folders in the release zip and .pkgmeta. [220]
- A worn item with an enchant gets a muted "(enchant not counted)" until an enchant table exists; the gain prints green for a plus and red for a minus; "Best until 30" names the next band's start ("Best for now" only in the cap band). The fixture's dps_per_point is replaced by the sim's real value at the L9.1 end-to-end weights capture; the client's saved file cannot be edited while logged in. [221]
- A pick the client cannot show gives way to the first showable alternative (own source line; the rest "Also good"), else "Best until N". The main zip ships ForeverSixty/ and ForeverSixtyInbox/ only, because ForeverSixtyData has its own nightly CurseForge project (1706707) and bundling the empty shape would clobber players' data (the brief's "as today" was wrong). [226]
- Fix round 4: the Next line is "Next upgrade: <name> +X DPS" and every line of ours is fitted (ellipsis or wrap) to the client's own widest line, never widening it; the comparison name keeps at least its first two words before the ellipsis. [228]
- The worn item's enchant is netted by scanning the client's "Enchanted:" line and parsing "+N <ITEM_MOD_*_SHORT>" pairs; the note survives only when the line cannot be parsed; "Equipped" sits on the header row as a double line; worn items in a non-cap band gain a muted "At 30: <next band's pick> · <source>" line; alternatives show their gain; "Best known for now" when the band's own pick is client-unconfirmed. [229]
- The designer's wraps are adopted: Next becomes "Next upgrade: <name>" / "+1.8 DPS (note)"; verdict becomes "Downgrade -46.0 DPS over" / "<full name>"; the enchant note no longer gets its own line (supersedes the earlier ruling). [230]
- The "At 30:" line keeps at least "<name> · <dungeon or kind>", dropping boss and chance first when short. [241]
- The controller rewrites the fixture's dps_per_point to the shipped sim value so "Your sim" and "Estimate" agree; no "Also good" on a Best-until item is accepted; keep/vendor hints for bag items deferred to L4.1. [243]
- "At 30:" source with a fit rule (short source first, else a second muted line, never ellipsise the name); "Best until level 30"/"At level 30:" wording; spec 4.3 amended (alternatives only under a Next line; the At line carries none); the fixture value is now the shipped 0.2342. [245]
- Hook-level specs with a spec tooltip; `wornEnchant` never caches an unreadable read; every section line goes through the fitter. [246]
- The header never loses its tag: drop the spec word first ("Forever Sixty · Your sim"), then move "Equipped" to the right side of the next row. [248]
- Fix round 7: every verdict with a figure gets a muted breakdown of its top three contributors; the At-line carries the next band pick's gain; "At level 30" wording stays. Carry to L4.1: a one-hander under a worn two-hander should show the pair verdict; carry to L9.1: the first Next-line capture on a character that has one. [254]
- The weapon conversion stays the sim's rule (site and addon agree), the breakdown line makes the figure reconcilable, and a sim-lane ticket validates `dps_per_point` per band against a real level-21 run; spec 4.3 gains the narrow-header sentence. [255]

### L3.4 build dialog
- Option A: the dialog shows Codec's own answer, older codes load, and the capture and spec use a newer-build code refused by `codecNewerBuild`, which is reworded to print no build ids; `build` joins `Codec.MESSAGE_TYPES` with the legacy `builds` list as fallback; `sim_request_state` reads the four named values only. [268]
- Persona round 1 rulings: the button keeps its site-fixed name "Copy my character" but its click reads "Selected - press Ctrl+C, then paste it on the site" (never "Copied"); the companion line becomes "The Forever Sixty app on your PC sends this for you when you log out."; "Sim queued · results on your profile"; the preview labels slots, tags gear "at 60", and says "You can take 12 of these 51 points now". [278]
- The button keeps its site-fixed label but turns secondary; the other designer rulings are folded into fix 1. [279]
- Wrong-class codes are refused by Codec with a codec-level string; `codecWrongPrefix` reworded with no echo; the self-check softened; dismissals by id and code. [282]
- Round 2 rulings: "You copied 22:21. Now paste it on the site."; the placeholder names foreversixty.gg; the "Copy my character" button goes in fix round 2 and the site's copy changes to "In game, type /fs copy, press Ctrl+C, then paste it here."; a muted tab hint at the panel's bottom-left; "(1 of 2)" stands as the chip's convention. [289]
- Designer round 2 rulings folded into the fix 2 brief (items 7-14). [291]
- Round 3 rulings: the line says "same as now" or names the first differing point; the ack line reads "Waiting for the site to see this copy."; a one-line lead "Selected. Press Ctrl+C, then paste it on foreversixty.gg."; every refusal carries verbatim advice; control-state ruling (hover/pressed tokens from the Theme; no focus-visible for buttons; no disabled now). [301, 303]
- Ruling: "Sim done · your sim weights are in use" plus the chip's best next upgrade when a weights message is in use. [314]
- Ruling changes recorded: the copy box keeps focus and selection after Ctrl+C (replaces round 2's drop-focus spec); a button disabled or hidden while held returns to its resting look. [321]
- Final pass rulings: the tab strip never moves (the count is a muted badge in a reserved slot after the paste tab's label); the sim line reads "Sim done from the 10 Oct 20:00 copy · your sim weights are in use" plus "Your next step: <the chip's line>"; "Replaces Arms raid · same points and gear · nothing changes". [325, 326]
- The chip reads `Gear.currentWeights` like the tooltip (one source of weights everywhere). [330]
- Under a weights message in use the line is plain "Sim done · your sim weights are in use" until the API carries `generated_at` and a `done` state (L6.11); request-state lines show only without a weights message. [340]

### L3.5 settings page
- `/fs help` prints the six player commands then one muted line "Testers: /fs diag verbose, inbox, hello, seen, record"; a scroll frame if the capture shows a cut; "Show the chip" on clearing a right-click hide accepted; the tooltip's click lines accepted. [349]
- Round 1 rulings: hints describe the ticked state in one voice; opaque page ground; the last label "Save my character for the companion app"; a seventh switch "Marks on the talent window"; `/fs options` opens the page and is listed; the data line reads "Addon data: current for this patch"; "Share my build" stays on by default and its hint says so. [351, 352]
- A fourth ratings state "newer than this addon" is kept ("too old" would be false); the page needs about 631 px and scrolls (captures decide, Reset must stay visible). [361]
- With marks off nothing replaces the chip, so it stays; spec text to say so in round 2. [363]
- Round 2: ASCII "Options > AddOns"; the Share hint in player words and one line; the stale state surfaces as a warning line under the title only while not current; "Player ratings:" label naming the Forever Sixty Data download; the disabled Reset label in the muted token with a border. [364]
- Describe the chip as it is: it stands aside while the marks are drawn, and stays when "Marks on the talent window" is off; the implementer diagnoses `Chip.inputs`' talents-open source for the N keybind path. "Not installed" counts as not current for the title line (later reversed, see next). [369]
- "Not installed is foot-only" (the designer's reading) wins over the round-2 ruling that it counts for the title line; sentences say "Player ratings (on player tooltips)" and name foreversixty.gg; the Save hint takes the designer's one-line copy. [378]
- L3.5 closed at 6f989280 plus a wording polish commit: "Addon data: for an older patch. Update Forever Sixty at foreversixty.gg." and "For testers: /fs inbox, /fs hello, /fs seen, /fs record". [385]

### L3.6 dps figures
- The missing-band fallback (next band down) is accepted; the spec 4.3 sentence change stands (the 14-AP rule is gone); the chip's "here for you" line naming picks through `bisNames` is carried to L4.x. L3.6 runs in its own worktree `fs-addon-l36` on `lane/addon-l36` so the main checkout stays stable during L3.5 captures. [362, 366]
- L3.6 closed; merges into addon/v2-foundation after the L3.5 fix-2 commit. Minors carried to an addon polish round (L3.7). [368]

### L3.7 polish rounds
- Round A (25 items) dispatched; items 15, 16, 17, 18, 19 dropped to the backlog (locale enchant parsing, DPS scale/percent, arrow head, Baganator widget verify, keep/vendor hints as a product matter). Round B follows with item 5 (talent-mark overlays) as a release must, 28-35 and 40-53. [610]
- Round B rulings: 32 pastes never overwrite a restored raid-slot build (own slot + confirmation line); 45 no protocol bump, the receiver infers the unit from the sender's spec role; 49 a rare-drops switch + skipping unwearable items; 50 a 10 s countdown in the Loot title; 51 inspect age; 52 "+12 score" for healers. Dropped: per-channel badges, own-tooltip weakest component, zone/rank/AFK/ready/role, button label/own row, wizard oil, Ready-check button, DPS reference, component naming (cross-lane). [616]
- Round A readings: the guild gap/hello keep scoring a cold weapon (withholding blanked the cell); the pair line names the best one-hander in the bags; rotation ranks come from the client's spell subtext. [615]
- Round A designer/player: the settings warning is dated by `data.generated` ("too old (made 9 Oct)") instead of the patch, plus the dim description lift and a "≥" on the cold-weapon gap cell go into round B as pre-items; Multi-Shot's missing rank derives from the spellbook's highest learned rank; the tooltip header's "Equipped" on a second row is kept (L4.1 fix-6 carry rule). Wording adopted into round B: "As a pair with <one-hander>:", "Not in your plan. Respec to fix this.", and the Gear-tips hint. [620, 621]
- Round B readings: the companion line under the paste box is accepted pending the personas; the "≥" is own-row only (a hello carries no flag). [624]
- The strip reads the level's points like the flyout ("12 of 15 points · 3 off plan"). [626]
- Healers print HPS everywhere (contract section 3's meaning of a healer's "dps"), "score"/"pts" only without a per-point figure; the Online-only store uses the `same` matcher; a receive-side rate limit on `push`. [629]
- Round B fix 1 landed; strip below level 10 reading the build's length accepted. Polish round B closed at b1f27429; addon lane complete. [632, 634]
- Pre-items for the next round: the dated warning, dim description, "≥" gap and the lows from round A, folded into round B. [618, 620]

### L3.8 default-build adoption
- New task L3.8 in worktree `fs-addon-l38` on `lane/addon-l38` from d3bc356c: reproduce the login path per class in specs, fix the real cause, re-run adoption when ranks arrive late, add a "Build: ..." diag line and a chip tooltip naming the build. [401]
- L3.8 closed at bef633e3; merge `lane/addon-l38` after the L4.1 fix-3 commit, then reload the hunter and check. [416]

## L4 addon surfaces

### L4.1 paper-doll and bag marks
- One verdict source (bagUpgrades/Gear.pick through currentWeights, pinned by a three-way agreement spec); a bagged two-hander marks the main-hand slot only; bag hooks once, no OnUpdate; the `Rules.canEquip` gate lands here; the empty-slot tooltip reuses the Next/source builders; the marks follow the "Gear tips on item tooltips" switch (no eighth switch). Bag keep/vendor hints and chip `bisNames` naming go to L4.2. [386]
- The implementer's reading that the target corner shows only when the tooltip prints "Next upgrade" (gain at least 1.0) is accepted provisionally. [393]
- The controller does not touch the client while the user is playing; L4.1 captures wait for an idle client or the user's word. [395]
- An empty chip line is a defect (the chip must say something or hide); an empty slot with no pick prints the muted "No pick for your band yet" line, matching the web panel's wording. [396, 397]
- Fix round 1: the crash fix plus suite-wide realistic `GetItemInfoInstant` stubs; targets not cached while pending and refreshed on GET_ITEM_INFO_RECEIVED; an unequippable carried pick counts as not carried; the disable stays until /reload and says so; the stricter target-corner rule stands. [398]
- Fix round 2 (taint): the addon never writes into a Blizzard frame; marks live on addon overlay frames parented to the client's buttons, no HookScript on the pane, with a mock-level spec forbidding client-frame writes. [402]
- The taint defect is closed after seven live opens; the talent-pane overlay carry (L4.5) stays. [412]
- Ruling: the corner means "a known next target for your band that you do not wear or carry" (later corrected, see next); an empty-slot filler is always an upgrade (no floor); fix round 3 follows. [413, 414]
- CORRECTION: the "corner on Best-until slots" defect was the controller's misattribution; the corner rule is right. Fix round 4 is redirected to the corner waiting during item-data wait, `bandPick` honouring `missing`, and the doc comment; the leak hunt is dropped. [424]
- The weapon header keeps the spec name and moves "Equipped" to the next row when the two do not fit (spec 4.3 amended). [425]
- Fix round 5: the client's atlas or a cropped arrow at a full 12 px; a legend line; the arrow-slot hover "In your bags / <item> +0.8 DPS / Right-click it in your bags to equip."; bracket inset + shadow; source fit order never cuts the boss; "World drop · check the AH" for BoE; ammo-type flags on ranged alternatives; a "Working:" diag heading. [438]
- Fix round 6: "Equipped" on its own row under the header; the atlas path verified or an addon-drawn chevron; two legend lines of at most 48 chars; the legend follows the marks including self-disable. [449, 451]
- The Waist's gain on the next row stands (never-widen rule). [456, 459]
- The designer overrules the hunter-19 "as ruled" cut: "Also good:" never cuts a name that fits (fix round 7). [459]
- Round 8 persona outcomes: both personas SHIP after round 3. [467]

### L4.2 quest and roll marks
- One verdict source with the tooltip; reward labels and roll lines on addon overlays; "Vendor" only when no choice is useful; the dungeon line via `Progress.picksHere`; bag keep/vendor hints out of scope; the labels follow the Gear-tips switch. [435]
- `ItemTooltip.judge` is the one verdict source; the chip's here line keeps spec 4.1's "Here for you: X (Boss), N more". "Best for <Spec>" needs a gain of at least 1.0 or an empty-slot fill, else no label. [448]

### L4.3 character panel flyout
- The plan's `Theme.addCharacterTab` via `PanelTemplates_SetNumTabs` is overruled by the 1.60 taint facts: an addon-owned "Forever Sixty" button beside the client's tabs plus an addon-owned flyout to the pane's right, scripts only on addon frames. "Sim me" copy names the sim entitlement. [469]
- Keep the default strata (the flyout is beside the compare tooltip, not over it); no redraw on pane open; chip "here for you"/"last night" still print to chat. [475]
- Fix round 1: icons 16 px; flyout 320 px wide with measured single-line fits; no scrolling (rotation at most 8 rows + "N more", last raid at most 6); "at level N" = the next band's unlearned lines (at most 2) appended last; the button keeps the fallback spot; a bag-arrow slot lists the bag item as its target. [477]
- Fix round 2: the flyout anchors right of `CharacterFrameModeTabs` when shown; the button only to a shown tab with a rect; probes under 180 chars; the line box follows the measured font height floored at 12; mark stale and redraw from the button's own OnShow. [485, 486]
- Fix round 3: "Plan:" split, progress against the level's points, off-plan hover, rotation hover with the full condition, button open/hover states, 20x20 close x, and similar. The over-ranked planned talents count is accepted; "· over <worn>" is not on bag rows. [490, 495]
- "· over <worn>" drops first (the action "check the AH" wins). [498]

### L4.4 player tooltip and chat badges
- The unit tooltip moves to `surfaces/PlayerTooltip.lua`; "clicking the score copies the site link" becomes `/fs link [Name-Realm]` opening the build dialog's copy box; chat badges via `ChatFrame_AddMessageEventFilter` behind an eighth switch "Score badges in chat"; `Ratings.lua` deleted; a ForeverSixtyData client fixture. [482]
- The badge goes before the message, not the name (the sender argument carries the client's whisper link, right-click menu and reply target); with no data addon installed the tooltip shows no score line. [492]
- Site links `/character/<region>/<ruleset>/<name-slug>` with a capitals/digits split for run-together realms accepted; Trade/LFG gated by zone channel id (2/26) first; `/fs link` splits at a hyphen only when the suffix is the player's realm. [493, 494]
- Fix round 2: badge "FS 1240 ·" (no brackets); the hint sentence; double-line components + "= total"; "gear gap" unit word; the link box text colour. Cross-lane carry: component maxima/scale and names ("Clears"/"Consumes") must agree between the site score card and the nightly record. [505]
- L4.4 closed at c90f8700; the cross-lane component naming/scale ruling stays carried. [513]

## L5 addon group surfaces

### L5.1 guild tab
- Button + flyout parented to `CommunitiesFrame` (the plan's `Theme.addGuildTab` overruled by the taint facts); default sort online-first then Score desc, "Online only" default on above 40 members; one `snapshot?` per open (at most 1/60 s); officer region by the push rank rule with the addon's own confirm frame; loot council region absent until L5.3. [514]
- Readings accepted: Gear = record gear score, Gap = live gap "14 (live)". [517]
- Fix round 1: one muted data-status line + blank cells instead of 97 "no logs"; the self row live; right-aligned numeric columns; the client's sort arrow texture; the count-line wording; a `GuildTab.seedLive` capture helper. Widths trimmed to the 460 inside accepted; record date left out (no room). [521, 523]
- Fix round 2: `gearGap` scores exactly as the chip, nil when unscorable; count line "N of M shown · K live of N online"; a count-line tooltip; Gap sorts DPS before pts; offline muted cells at least 3:1. Biggest-slot semantics stay with the tooltip saying so. [526, 527]
- The hello now names the build's spec (accepted). Contract carries (L6.3/L7): the hello must send an integer; no unit travels with the gap (a `gapUnit`/encoded unit needs a protocol bump); two-handers scored per hand. [528]
- Fix round 3: one shared per-slot target helper feeds both the Next line and `Comms.gearGap`; a gain under 0.5 shows and sends as 1; uncached picks unscored. [530, 532]
- Ruling reversed: a build code's gear list is a snapshot of equipment, never a target; the per-slot target is always the band's pick (the earlier fix-3 reading that a build's gear list IS the plan is withdrawn). Fix round 4 dispatched. [532, 535]

### L5.2 group panel and ready check
- Button + flyout parented to `RaidFrame` (plus `/fs group`), "Not in a group." solo; ready check against `Data.encounterNeeds`; applicants via L4.4's tooltip; `/fs floor N` flags only; no switch. [540]
- The instance-to-boss map and class-buff provider table are hand tables in the addon for now (carry to the data lane: `encounterNeeds` should carry instance keys and provider classes); `/fs group` does not open the Raid tab; `/fs help` and the settings command list gain the three commands. [541]
- Client fact: `FriendsFrame` has no rect either; the drawn Social window is `SocialUIFrame`; anchor the group button/flyout to it while `RaidFrame:IsShown()`. [543]
- Fix round 1: anchor/parent to the first placed Social frame with `RaidFrame:IsShown()` as the tab test; group-buff variants by the build's real ids (three of the brief's ids were other spells and are dropped; the data lane must carry the variants). A small addon-owned sentinel frame on `RaidFrame` re-evaluates the tab. [544, 545]
- Fix round 2: sentinel frame with its own OnShow/OnHide; solo wording; a clean solo dry run "you are missing nothing" accepted. [547, 550]
- The tab test was wrong on 1.60: the Raid tab is open when `SocialUIFrame.RaidFrame` is shown; the global `RaidFrame` is a legacy frame. The solo chip line is "Ready check: missing flask, food, imbue". [551]
- The solo clean line is "nothing missing"; flyout closes on tab-leave (guarded `closeFlyout()`). [552]
- The `/fs ready` chat clean line solo "Nothing missing." (group keeps "Nobody is missing anything."); `command()`/`reread` close the flyout opened by `/fs group` on a non-Raid tab when the anchor is placed. [554, 557]
- Round 1 persona (solo states) rulings: level gates (flask 50, food 35) in one table; imbue only for warrior/rogue/ret/enhancement/feral; the Natural flasks added; generic = class buffs + consumables only; coverage reworded; "missing" in every Missing line. [560]
- Fix round 4 accepted the Natural flasks' use spells 1293740-43 (1293559-62 were Alchemy recipe spells); the solo second line wraps to two lines; the displayed spec for the imbue gate; the disclosure counts any listed melee member. [562]
- Fix round 5: the side-tab scan counts only shown children with left at or beyond edge-1 and width at most 100; the clean chip uses the tick glyph; `druid-feral-bear` excluded; the recipe ids added back as a hedge; no deferral. [564]
- Drop `druid-feral` from `IMBUE_SPECS`; the flask gate applies only at level 60 or inside a known raid instance; "weapon buff" adopted in the polish round; the solo first line brighter; the four small items fold into L5.3's brief as pre-items. L5.2 closed for the solo states at 22a5fc23. [570, 571]

### L5.3 loot council
- Rulings on the concerns: `/fs help` prints the six player commands plus one muted testers line; "Show the chip" on clearing a right-click hide accepted; the tooltip's click lines accepted. [349]
- The command waits up to about 5 s for the item load and names the item as "Item <id>" meanwhile; the "worn item vs itself" case reads "already worn". [584]
- Fix round 1: a command waiter + a 5 s "unknown item" timeout that clears the request; the site fallback on uncached/nostats/nofigure; the dedupe resets only when the window was closed; epics are always asked whatever the threshold (the threshold only adds rares). [588, 589]
- Fix round 2 rulings: all of the player's items except the carries (menu option/unwearable skip, countdown, inspect age, healer unit wording, DPS reference). [594]
- Drop `keep`; add the colon to the full chip line; a gated member sends no answer and reads "can't score". [600]
- Fix round 4 (tiny): `Chip.inputs` uses `Comms.buildSpec` and `Chip.buildSpec` is deleted. L5.3 closed for the solo states at 5e4f7077; group states stay owed. [606, 608]

## L6 API lane (worktree fs-api-v2)

### L6.1 ScoreV2
- `ScoreV2` as a Store method stands; dungeons cleared stays 0 until a dungeon classification exists in the data; rated fights count from public reports only (a guild-visibility rule is L6.4's call). [87]
- The gear denominator is derived lazily per (build, spec) in `sims` from bis/<spec>.json with that build's weights, cached; a discipline category is satisfied by any consumable of that category from any role; the record's `gear_score` is the gear component in points (0-500) and `attendance` the fraction 0-1; attendance nights by `fought_at` date; a ScoreV2 failure renders the character page with null score fields, never a 500. [90]
- Migration 0035 edited in place is acceptable because lane/api-v2 has never been deployed or merged; cost if wrong: a manual column fix on one database. [94]

### L6.2 sim requests
- Sim requests from the addon require the same `server_sims` entitlement as POST /v1/sims/run; a refused request is reported through the ack in L6.3 (field `sim_request_state`: "queued" | "needs_premium" | "duplicate"). [62]
- `sim_request` content problems answer 200 with `sim_request: "invalid"` (only a structurally malformed outbox gets a 400). [72]
- L6.2 is dispatched before L6.1 because L6.1 consumes sims.GearScore. [24]

### L6.3 inbox v2
- Per-fight last-raid extras are computed in the rater and stored on rating_scores (migration 0036), never at inbox read time; rotation adherence is deferred to a new task L6.5 (`adherence` optional in contract 3); `sim_result.slots` empty for a plain run; `loot_values` prices against the newest export with the newest weights run; ack only for users with a non-revoked device. [93]
- Ruling 1 revised: inbox `character` is the `region/ruleset/name-slug` key; Ruling 8 revised: `dps_per_point` read from data/builds/<build>/dps-per-point.json; Ruling 9 revised: `loot_values` reads data/builds/<build>/loot.json by the last_raid zone. [97]
- The addon's contracts_spec.lua edit stands; a survival loss without a death reads "standing in fire"; an unnamed death cause reads "a death"; fights the engine marked insufficient are left out of last_raid; the warrior-arms FullGearScore pin moves 4838 to 4767 (not a defect). [112]
- Under a two-hander off-hand loot items are skipped; NULL-fact fights are left out of last_raid rather than zeroed; no model-version bump. [118]
- The new logs-engine helper `ActivityIdleMS` (additive, score-neutral) stands. [127]

### L6.4 nightly record
- Nightly character keys are Name-Realm with the canonical ruleset realm; anonymised accounts omitted; no-logs guild members get a nil-score record; guild standing = rank by bosses killed then attendance; gear below 60 measured against the character's own band's best set; tiers stay provisional. [134]
- The contract 5 fixture is the real writer's output so dungeons stay 0 (the addendum's 14 is withdrawn); the main.go gear-data load fix stands. [153]
- Keys are Name-<client realm, spaces/hyphens removed>; skip-and-log per character (abort on more than 50); guild attendance is the mean over scored members. [157]
- `RulesetFromRealm` mapping "Classic Beta PvP" to pvp is correct; rows stored under the old normal key are an ops re-key item. [174]

### L6.5 adherence
- adherence = casts of APL-listed abilities over all ability casts for the spec's level band (excluding auto-attacks, consumables, movement, utility), absent under 10 casts or without an APL; migration 0037; never enters ScoreV2. [206]
- The level band comes from the character's stored level at rating time, band 60 when none; any rank matches via spellranks families. [208]
- Fix round 1 for the cast count (spell-id exclusions, boundary tests, the rotations-empty log). [219]

### L6.6 to L6.14
- L6.6: the chrome snapshot is retaken from lane/web-v2's final HEAD once L8.2 closes. [191]
- L6.7: option (b), the Gear component is unchanged and `gear_gap_total` is the plain sum of gaps with its own openapi wording. [269]
- L6.8: one shared key-based attendance reader for roster and score; keep the contract shape; the roster's window becomes the last 8 non-private reports with a boss fight; verification by logs goes live. [274, 297]
- L6.10: a forged report must never raise a score; the re-rate corrects only the newest snapshot. The nightly writer switches to `presence.EvidenceReport`; rating endpoints and inbox `last_raid` stay display-only readers for now. [334, 358, 389]
- L6.11: the in-game "Sim me" becomes a weights run (migration 0038); weights message per character from the in-game run first; the sweep reports baseline DPS; a failed run does not count toward the hourly duplicate window. L6.11 closed at a441264e. [359, 365, 374, 387]
- L6.12: a second predicate `presence.NotFlagged` for member-facing views; the per-fight ratings endpoint 404s on a flagged report. A public viewer sees `EvidenceReport` counts only, a verified member `NotFlagged`; a non-member never sees a flagged report in Raids. L6.12 closed at ccbe0d22. [453, 464, 470, 478, 483]
- L6.14: `roster[].addon` always present; consent rule: without gear consent `addon.level/spec/gear_gap/build` are null while `seen_at`/`snapshot_at` stay. [566, 569, 572]

## L7 companion (lane/companion-v2)

- The companion follows the API's actual export answer shape; extra fields on known types accepted at runtime (strict only in the golden). Fix round: fetch the inbox right after a successful outbox post and persist the "already sent" outbox key. [140]
- Fix round 1: fetch the inbox right after a stored outbox; persist the sent key; `.row[hidden]`; a 256 KiB outbox bound that never holds the exports back. [144]
- The exports still go up once at launch after a restart (accepted). [151]
- The 256 KiB outbox cap stands as a companion-only guard; fix round 4: 400 only and the mislabelled log. [165]

## L8 web lane (worktree fs-web-v2)

### L8.1 and L8.2
- Every "Copy for the game" button emits an FSB1 build code with item stats; the plan's "sim deltas as stat lines" is overridden; `buildCodeFor` lives in lib/addon/build-code.ts. [91]
- The planner chip button takes the same copy strings through shared constants; enchants and suffixes do not travel in FSB1. [96]
- Fix round 1: chip strings, a SavedSim gameCode wiring test, the data-build stamp caveat stated. [100]
- L8.2: no sparkline at first; creates docs/contracts/fixtures/score-labels.json from the addon's Locale; the "In game" section shows last-seen time only if an existing endpoint exposes it. [106]
- The "In game" section living in the Character.svelte island stands; the astro.config server.fs.allow entry stands. [119]
- Common is white on both sides (the fixture said #9d9d9d; the addon maps common to quality 1). [124]
- Fix round 2: the fifth part's label becomes "Preparation" on both sides (API field name unchanged); a header lookup on every page. Two filled gold buttons on the signed-out home stand; 56 px search-row controls stand. [142, 146, 167]

### L8.4 to L8.6
- L8.4: the "N ranked fights" header count stays removed; the guild line takes the client's angle brackets; the sparkline's vertical scale spans at least 200 points; the e2e fixture's fights gain execution_score. [207]
- Round 2 rulings: the sparkline fills its container with a caption "Up 330 since 10 Nov"; "Holy Priest · 31/20/0" from builds_seen; the owner line is the Character header's sync line; the mock prints "(Player)"; the visitor band moves below the score card. [212]
- Mock prints the level line without the class, the class on its own line, the faction line when known. [215]
- Fix round 3: noscript message; realm nowrap; caption 12px; two-column card from 1280; faction derived from race. [234]
- The score card lists each component's own gap in muted text after "Most to gain"; the Performance Rating panel's consistency and sparkline become L8.5; the gear panel is L8.6. Dots attach to the end of the preceding segment. [261, 285, 286]
- L8.5: the "weakest part" link goes to the newest rated fight; mobile /rankings Execution tests skip below 600px. Important re-ruled: "Latest fight rated 65 of 100 · 6 fights rated"; "counts toward Parses 540 of 600" dropped. [323, 327]
- L8.5 fix round 2 wording rulings (latest fight named, parse vs rating line, row highlight, Experience note). [347, 348]
- L8.6: quality colours from the build's items file; "Best for your band" also for a gap-0 worn item that is not the pick by id; full-column width stands. [367]
- Fix round 1: hold the quality colour until loaded; explanation line under the title with live numbers; "Not scored"; "17% drop". [370, 379]
- Fix round 2 and 3: hover opens only for `pointerType === 'mouse'` and only from a real pointermove; skeleton sized from the real name. "Binds when picked up" is a data carry. [407, 417]
- Fix round 4: below `lg` the note always takes its own line with no dot; the alt-text form silences the dots. [442, 461]
- Fix round 6: the blank row, failed-message shift, non-breaking "main-hand", DESIGN-SYSTEM.md, and the fight page's 404 `flagged`. L8.6 closed at 7dfc1c09. [474, 481]

### L8.3 addon page and roster
- L8.3 needs an API addendum (Task L6.14) and is coded against that shape with the fixture. [566]
- Fix round 1 rulings: commands from Locale.lua by test; one tester line; a three-sentence handshake; the run-together fixed; constant row height with a fixed In-game slot; "gap 12 DPS"; "per addon" only on a difference. [582]
- The In-game slot is fixed only from 1280 (accepted); then moved to lg (1024). The "-" stays (matches the in-game placeholders). [595, 597, 598]
- L8.3 closed at bd622e51. [613]

## L9 release

- L9.1 step 1 closed: weekly loop README and capture checklist (96 captured in week 1, 82 owed); the first "week-2" run is the group-session capture day with the user. [609]
- Integration dry run: take ours for the 70291 data, then the 70374 move regenerates with the merged ranker. [630]
- Integration brief (L9.2): merge order, origin/main merge, 70374 stamp move, release pass, the hunter-thrown data defect, the ranker rounding lows, one push of main; the addon's "Discipline" label becomes "Preparation". [631]
- Integration dispatched in two parts (A: merges + 70374 + toc/CHANGELOG/README; B: the data defect + ranker rounding + re-rank), then the memory capture and the single push with the user. [634]

## Cross-lane and contract

- Pre-flight: L2.6 adds `Follow.ensureDefault`; L2.2 includes `Progress.unspentPoints()`; L3.1 adds chip state 0 (see L2/L3). [21-23]
- The Fable 5.1 trailer stands for every commit in this plan. [52, 242]
- Defect (controller's own brief): Common is white on both sides (fixture and Lua spec change in the L8.2 fix round); carry to L6.4 a rule for the gear component at low-level caps, measured against the character's own band's best set. [124, 134]
- Build move: move straight to 70374 (skipping 70334) as a stamp move; it is a main-branch operation belonging with the v2 merge onto main, scheduled as the first step of the integration. [602]
- Contract carries: the inbox `character` key format (Name-Realm vs region/realm/name) was resolved by L6.3's revised ruling 1. [97, 475]

## Process and tooling

- Addon lanes (L2-L5) run in this checkout one implementer at a time; non-addon lanes run in their own worktrees, one implementer per worktree. [25]
- Lane worktrees: `../fs-data-v2` (lane/data-v2), `../fs-api-v2` (lane/api-v2), later `fs-web-v2`. [41, 91]
- Toolchains: Go in WSL, wowsims-forever fork, `make simdb`, uv, Postgres in WSL (API packages run with `go test -p 1`). [35, 46]
- Subagent commits carry their own trailers; the controller rewrites them to the Fable 5.1 line before merge. [52, 82]
- The controller never enters credentials; never call open_application for the game. [85, 319]
- The controller does not touch the client while the user is playing; keepalive jumps via the raise script. [307, 395, 623]
- The superpowers plugin cache is gone, so review packages are plain `git diff BASE HEAD` files. [576]
- On 2026-10-14 implementers run on Sonnet until Opus is available again (Oct 16). [574]
- Persona agents (ux-designer, wow-player) review captured screenshots each round; fix rounds fold both personas' findings. [103, 104]

## Carries still open

- Data lane `encounterNeeds`: the table is identical for all 19 bosses (only MC's fire resist varies; ZG/AQ/Naxx fall to generic); it should carry instance keys, provider classes, group-buff variants, elixirs-as-flask, consumable sets with levels, and extra buffs/debuffs (Might/Wisdom/Salvation/Divine Spirit). Food level threshold (35 vs about 25 foods) and flask 50 vs 51 also doubtful. [541, 545, 560, 571]
- Hunter-thrown data defect: a thrown weapon (Cracked Blacksmith Hammer) ranked as a hunter's ranged alternative; the ranker's thrown allowance (L1.7) must exclude hunter ranged slots (warriors/rogues keep it). A must before release. [621, 631]
- Ranker determinism and rounding lows are carried into L9.2 part B. [531, 631]
- Group-session captures owed to the user's test group: L5.2 group states, L5.3 group states (master looter's drop, Loot region rows, "no addon"/"can't score", "N of M answered", 60 s linger; README rows 74-76), officer push on the GM character, per-character hide with two alts, the "(staged)" tag, Trade/LFG and incoming whisper badges, epic colour, the worn-slot arrow capture, dungeon line / quest turn-in / group roll captures. [463, 513, 571, 608, 625, 634]
- Web polish backlog: L8.6's nice-to-haves (failed message, "main-hand" split, phone row rhythm, real-Tab and Escape-then-hover proofs to L9.1, "+0 pts"), L8.3 rounds 2-3 nice-to-haves, and the `source-build` pin on the 70374 move. [473, 596, 612, 613]
- Cross-lane component naming/scale: score components' maxima and names ("Clears"/"Consumes" vs the site's) must agree between the site score card and the nightly record. [505, 513]
- Gap-unit protocol note: no unit travels with the gap (a `gapUnit`/encoded unit needs a protocol bump); round B ruled "no protocol bump" and has the receiver infer the unit from the sender's spec role. [528, 616]
- Companion carry: per-member site values for addon-less members; menu option for the quality rule + skipping unwearable items; 10 s answer countdown; inspect age; healer unit wording; DPS reference once per session. [608]
- API carries: score history for the sparkline, a guild field on the character endpoint, the anonymised-character history leak, a per-source "last seen" time in `/v1/me`, L6.13 polish, the progression tab's forged-kill product decision (L6.14 carry), `dps_per_point` per-band validation. [122, 167, 191, 387, 443]
- Data lane audit list: Windows-only data test failures, sim goldens drift against the local engine fork, ItemSparse.Bonding (BoP) and crafted-pick schema carries, the dps-per-point re-measure on the 66 changed alliance sets (L9.1), and rotation audit items. [360, 387, 480, 490]
- Capture-environment carries: "not installed" data addon state shot, other UI scale, longer locale wrap, and the L3.3 first Next-line capture on a character that has one. [254, 385]
- Addon polish remaining: items dropped to the backlog (15 locale enchant parsing, 16 DPS scale/percent, 17 arrow head, 18 Baganator widget, 19 keep/vendor hints), the L4.5 talent-pane overlay (done in round B item 5), README row 37 and similar stale rows. [411, 456, 610]
- Release pass: toc 2.0.0, CHANGELOG, README spike rows, `addon-release.yml`/.pkgmeta shipping both folders, the Fable trailer rewrite, and the single push of main with the user. [631, 634]
