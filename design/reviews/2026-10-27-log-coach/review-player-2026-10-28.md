# Log coach: player gate (2026-10-28)

Reviewer: wow-player (Classic since 2019, raided through Naxx, reads Details! and the logs after every raid night).
What I looked at: the 162 captures in this folder (index: `.superpowers/sdd/2026-10-27-log-coach/task-41-run-report.md`), the mock boards, spec `design/specs/2026-10-28-log-coach.md` §4 and §8, and `data/curated/coach/warrior-fury.{md,json}`.

## Overall verdict: NOT SHIPPABLE

The structure is right. A Fix first block, a short ranked Then list, Doing well, and a Couldn't check card that never hides its lines: that is the Details!-plus-Wowhead-mistakes layout I would want. The officer, signed-out, whole-night, account-line and keyboard states are close to done.

But on the fixture, nearly every sentence fails the "can I check it on the log?" test. A guildmate would screenshot and mock these in guild chat:

- **Times that cannot exist.** The rail says `Kill 1:00`, yet the coach says `Death at 3:20`, `The longest gaps were 1:23 and 2:41`, and gives chips `Timelines · 1:23–1:31` and `Resources · 2:00`.
- **A link sentence that contradicts itself.** `The Bloodthirst gap at 1:23 includes your death at 3:20.`
- **Wrong boss count.** The account line says `Onyxia's Lair … on 5 of 8 bosses`. Onyxia's Lair has one boss.
- **A Fury warrior told about a Protection talent.** `Boss damage killed you at 3:20 with Last Stand ready.` The same page also says `Defensive cooldowns: the coach doesn't know your spec's list yet.`
- **Warrior lines on a Holy Priest.** The priest's review praises `Battle Shout was up for the whole pull.` and says it is `Based on: … top Fury Warriors`.
- **Chips with no spell.** They read `· Casts`, `· Buffs`, `· Debuffs`, `· Healing`.
- **No spell icons anywhere.** Not in the Fix first title, the Then titles, the row strip, the teaser or inline names. The mock has all of them.
- **The class where the spec belongs.** `Your Warrior coach` and `Top Warrior logs for this boss`.
- **A horizontal scrollbar at 1440.** The rail spills out of the page (the capture is 1548 px wide).

Some of this comes from the fixture rather than the components (the fake `Sanguine Depths` / `Warden Kelthas (5)` night and the Onyxia loot in a Molten Core post). But these captures are the only rendered proof, and a player judges what is on screen. The fixture has to be a believable Classic night: the spec's own Tyrim / Onyxia (40) / Kill 3:52 worked example is ready to use. Recapture on that before a second gate.

The coach's voice also does not show. The model review (`Your Fury coach`) is word for word the rule-only review (`Your Fury review`). The mock's why, "Bloodthirst is the button the Fury rotation is built around … so that one gap cost you twice", reads like the best warrior in the guild talking. The rendered why, "Bloodthirst went unused 9 times while you had the rage", reads like a parser.

## Rows

### A1: report row strip, Premium owner
- **FIX** `log-coach-a1-row-premium-1440.png`, `-390.png` (and 1024/1920/2000): the strip reads `Fix first · Bloodthirst: 9 casts lost` with no Bloodthirst icon. §3.2 asks for a 20 px spell icon before the title. Players scan logs by icon, so add it.
- Data: `9 casts lost` in a 1:00 kill. Bloodthirst has a 6 s cooldown, so about 10 casts fit in a minute, and losing 9 means the warrior pressed it once. Yet the panel says `91% of your casts were on the written rotation`. Those two cannot both be true. Make the fixture consistent.

### A2: strip states
- **FIX** `log-coach-a2-strip-preparing-1440.png`, `-390.png`: `Top Warrior logs for this boss are still loading.` Key `rowPreparing` is `Top {Spec} logs…`, so this should read `Top Fury logs for this boss are still loading.` A Fury player does not compare against Protection warriors.
- **FIX** `log-coach-a2-strip-free-1440.png`, `-390.png`: `2 things to fix · First: Bloodthirst: 9 casts lost` has two colons back to back and no icon. Suggested `rowFree`: `{n} things to fix · {title} first · Read the full review with Premium`, with the spell icon before the title. Apply the same change to `rowFreeOne` and `teaserFirst`.
- SHIP: `free-no-first`, `premium-no-first`, `clean` (`Nothing to fix on this pull that the log can show.` reads honestly), `error`, `loading`, `hover`, `focus` at 390 and 1440.

### A3: officer, stranger, signed out
- SHIP: `log-coach-a3-officer-opened-row-1440/390` has a `Coach review` button in Morrowlyn's opened row, and only there. `stranger` and `signedOut` show no strip and no Coach tab.

### A4: Coach panel, model review
- **FIX** `log-coach-a4-panel-model-1440.png`: the page is 1548 px wide and the rail is cut off at the right edge. At `-2000.png` the rail sits outside the centred 1344 px column (spec §6: 760 + 40 + 320 inside the column). Keep the rail inside the column.
- **FIX** all widths: the Fix first title `Bloodthirst: 9 casts lost` has no 32 px icon (§3.1 item 4). The Then titles have no 24 px icon. Spell names in the why are plain text, not icon tokens (§3.1 "Names inline"). The mock has all three.
- **FIX** all widths: the chips read `· Casts`, `· Buffs`, `· Debuffs`, `· Healing`, with no spell. Keys `anchorCast`, `anchorBuffs` and `anchorDebuffs` are `{Spell} · Casts` and so on, with the 16 px icon. Expect `Bloodthirst · Casts`, and for the flask `Flask of the Titans · Buffs`, or drop that chip because `Buffs · at the pull` already covers it.
- **FIX** all widths: Whirlwind's chip is `· Debuffs`. Whirlwind applies no debuff. Use `Whirlwind · Casts` alone.
- **FIX** all widths: `The longest gaps were 1:23 and 2:41.` reads as an 83-second gap. Say when and how long, the way the mock does (`longest at 1:12–1:25`). Key `DetailCoreCasts` becomes `{lost} of a possible {possible} went unused while you had the {resource}. Longest gaps {gap1} and {gap2}.`, with each gap filled as a `m:ss–m:ss` range. Also render the key's own words; the panel shows `Bloodthirst went unused 9 times…`, which is not the §8 sentence.
- **FIX** all widths: `Death at 3:20` and `Boss damage killed you at 3:20 with Last Stand ready.` on a 1:00 kill.
  - Last Stand is a Protection talent that a 17/34/0 or 20/31/0 Fury warrior does not have. The log cannot show talents, so the coach cannot claim it was "ready".
  - `Boss damage` is not a killer name. Key `DetailDiedWithReadySince` wants `{Killer} killed you at {t}. {Spell} had been ready since {since}.`, with the killer as the spell or `Onyxia's melee`.
  - For Fury, the honest version is `DetailDiedNotUsed` with a real Fury button, or no defensive clause at all.
- **FIX** all widths: Couldn't check says `Defensive cooldowns: the coach doesn't know your spec's list yet.` while a Then item names a defensive. Show one or the other, never both.
- **FIX** all widths: `Spell ranks: this log doesn't tell us your level.` on a raid boss pull reads lazy. Every raider on a raid boss is 60. Either assume 60 on raid zones, or reword `CannotLevelRanks` to `Spell ranks: this log doesn't record your level, so ranks weren't checked.`
- **FIX** all widths: the rail reads `Fury Warrior · Warden Kelthas (5) · Kill 1:00` and the footer `top Fury Warriors on Warden Kelthas (5)`. A 5-player raid boss and a Shadowlands-sounding zone (`Sanguine Depths`) make me close the tab. Recapture on the Onyxia (40) worked example. The footer also needs the pool size, as in the mock: `Top 25 Fury warriors on Onyxia (40)`. Note too that Couldn't check shows no `Talents and race: Forever logs don't record them.`, which the mock rightly carries.
- **FIX** all widths: Doing well reads `91% of your casts were on the written rotation.` with no reference point. Key `StrengthAdherence` ends `Top {n}: {ref}%.` Without it, 91% means nothing.

### A5: writing to ready
- **FIX** `log-coach-a5-swap-writing-1440.png`, `-390.png`, and the busy and ready-rules states: the link line `The Bloodthirst gap at 1:23 includes your death at 3:20.` (`LinkDeathInGap`). A gap at 1:23–1:31 cannot contain a death at 3:20. The link must only fire when the death time lies inside the gap range, and the fixture must respect that.
- **FIX** `log-coach-a5-swap-ready-1440.png`, `-390.png`: the "coach" text is identical to the rule-only text, so the swap delivers nothing. Feed the ready state the worked model review from the spec (the mock's why and drill). The no-shift measurement passes (y values equal), so this is about content, not layout.

### A6: panel states
- **FIX** `log-coach-a6-state-loading`, `preparing`, `not-reviewable`, `error`, `teaser-two`, `teaser-one`, `teaser-empty` (-1440, -390): the header is `Your Warrior coach`. `titleModel` is `Your {Spec} coach`, so it should be `Your Fury coach`.
- **FIX** `log-coach-a6-state-not-available-1440`, `-390`: the header is `Your Warrior coach`. §4.2 uses `titleGeneric`, so it should be `Your coach`.
- **FIX** `log-coach-a6-state-preparing-1440`, `-390`: `Top Warrior logs for this boss…` should be `Top Fury logs for this boss are still loading. Check again in a few minutes.` (`preparing`).
- **FIX** `log-coach-a6-state-not-reviewable-1440`, `-390`: the panel shows the generic `notReviewable` line but none of the answer's `couldntCheck` reasons. §4.2 asks for them. For a trash pull, "too short to review" is false; use `NotReviewableTrash` (`The coach reviews boss pulls only.`) or `NotReviewableShort` as the reason.
- **FIX** `log-coach-a6-state-no-findings-1440`, `-390`: the opening says `A kill with one habit to fix: Bloodthirst sat unused…` under a review with nothing to fix. It must be `rulesNothing`: `Nothing to fix on this pull that the log can show.`
- **FIX** `log-coach-a6-state-raid-buffs-only-1440`, `-390`: the opening names Bloodthirst, and Then lists `Death at 3:20`, `No flask at the pull` and `Whirlwind: 4 casts lost`. None of those is a raid buff, so this is not the raid-buffs-only state. Render the state from §4.2: no Fix first, an opening about raid buffs, and Then holding `No Mark of the Wild at the pull` or `2 raid buffs missing` (`TitleRaidBuff`/`TitleRaidBuffs`, `DetailRaidBuff`).
- **FIX** `log-coach-a6-state-teaser-two`, `teaser-one` (-1440, -390): the teaser has no Bloodthirst icon (spec §10.2 sends `spell_id` for exactly this), and the title reads `First: Bloodthirst: 9 casts lost`. Use the icon and `Bloodthirst: 9 casts lost` under the gold `2 things to fix` label. The `teaserFirst` key becomes `{title}`. Otherwise it is a fair taste: a free owner learns what to fix first and is shown one button.
- **FIX** `log-coach-a6-state-healer-1440`, `-390`:
  - Doing well is the Fury warrior's lines (`91% of your casts were on the written rotation.`, `Battle Shout was up for the whole pull.`). The footer reads `top Fury Warriors`. The healer fixture must carry its own strengths and basis, e.g. `StrengthOverheal` `Flash Heal (Rank 7): 12% overheal. Top 25: 15%.`
  - The why `More than a third of your Heal never landed on a hurt target.` misdescribes overheal. Overheal is healing beyond missing health; the target was hurt. Say `38% of your Heal's healing went past full health. Top 25 Holy priests: {ref}%.`
  - The title needs the rank, the way players say it: `Heal (Rank 2): 38% overheal`.
  - The drill `Hold Heal until the tank is under 70% and use Flash Heal for the spikes` is not how a Classic Holy priest plays. They downrank, start the cast early and cancel it when the target is topped. The opening `Flash Heal carried the throughput` praises the least mana-efficient heal. Take the drill from the priest-holy persona.
  - "rotation" is the wrong word for a healer's Doing well.
- **FIX** `log-coach-a6-state-busy`, `ready-rules`, `starter` (-1440, -390): these carry the same A4 and A5 content defects (impossible times, Last Stand, empty chips, the false link line). The state chrome itself is right: the busy line with its warm rule, `Your Fury coach · early version` with `This coach reads your log against our guide's Fury rotation.`, and `Your Fury review` for rule-only.
- SHIP (chrome and copy): `officer` (`Baelgrim's Fury review`, `Officer view`, `Written to Baelgrim. You see it as an officer of the guild.`, no History), `signed-out`, `whole-night`, `error`'s body line, `teaser-empty`'s body line. Their content defects are covered above.

### A8: account last-raid line
- **FIX** `log-coach-a8-account-premium-1440.png`, `-390.png`: `Last raid · Onyxia's Lair, Sat 24 Oct · Fix first: Bloodthirst casts lost, on 5 of 8 bosses · Open review`. Onyxia's Lair is one boss. With 8 bosses this is `Molten Core`. For Onyxia use `lastRaidPremiumOne`: `… Fix first: Bloodthirst casts lost · Open review`.
- **FIX** `log-coach-a8-account-free-1440.png`, `-390.png`: the same zone problem. It should read Molten Core when the night counts several findings across bosses.
- SHIP: `pending` (`Reading your pulls…`), `clean`, `none` (`Last raid · no log in the last 14 days`, muted, honest), `error`.

### A9: keyboard
- SHIP: the focus rings on the strip, the chip, the heading and the Premium link are all visible gold rings. (The chip in focus still reads `· Casts`; see A4.)

### A11: Discord
- **FIX** `log-coach-a11-discord-report-coach-one-*`, `-two-*`, `-pending-*` (desktop 640, phone 390): the `Your coach` field comes after `Loot`, which dangles `10 items. The first 8:` into the coach field. §3.4 and the §9 checklist say `Your coach` goes after `Kill order` and before `Loot`.
- **FIX** same files: the link shows raw, as `[Open your coach's review](https://foreversixty.gg/…#coach)`. Real Discord renders masked links in embed field values, so either discord-render.ts lacks markdown or the field is wrong. Prove it on the private test server, as §9 A11 asks. A raw URL in the raid channel looks broken. Also check that `/guild/us/normal/olympus-xxvii/raids/7#coach` lands on the coach (D1 wants a URL that reopens the panel).
- **FIX** `log-coach-a11-discord-report-coach-two-*`: `Duskstrider · Fix first: Aimed Shot clipped — on all 8 bosses.` is not a §8.1 stem. Hunters say "Aimed Shot clipping your Auto Shot". Use a real stem (`NightCoreCasts` `Aimed Shot casts lost`, or add a keyed stem with that wording). Zulmara's line reads well.
- **FIX** fixture across all four: the loot embeds read `Item`, `item level 74`, `to Pyrewisp · Onyxia`, with broken images, in a Molten Core report. Not coach copy, but a player opens this post first.
- **FIX** `log-coach-a11-discord-report-ready-coach-*`: the post says `Kill times, loot and the biggest changes on Forever Sixty. Coach reviews for tonight's pulls are ready on Forever Sixty and in /report.` Two "Forever Sixty"s read like an ad. Change `raidNight` to `Coach reviews for tonight's pulls are ready; open the report or use /report.` It still names and counts nobody. On "would I open it": yes, if the `/report` line names my own Fix first, which it does.

## Data I do not believe (for the accuracy reviewer)
1. `data/curated/coach/warrior-fury.json` lists `Last Stand` (12975, `talent: true`) as Fury's only defensive. Last Stand is deep in the Protection tree, and no raiding Fury build takes it. The log cannot show talents, so F7 will tell Fury warriors that a button they don't have "was ready". Make Fury's defensive list `[]` (unknown, which yields `cannot_defensives`) unless it is seen cast this pull.
2. Fixture times beyond the fight length: the death at 3:20, the gaps at 1:23 and 2:41, and the resources chip at 2:00, all on a 1:00 kill.
3. `9 casts lost` beside `91% on the written rotation` in a 60-second fight.
4. The holy-priest fixture reuses the Fury strengths and basis.
5. `Onyxia's Lair … 5 of 8 bosses`, and `Warden Kelthas (5)` as a raid boss.
6. Spec §10.7 already flags `Rend` in the `DetailAdherence` example. Confirm it never reaches a Fury player.

## Re-read 2026-10-28

Read against the gate-fix captures (`.superpowers/sdd/2026-10-27-log-coach/coach-gate-fix-report.md`, `task-44-report.md`). The fixture is now a believable night: Baelgrim, Fury, Onyxia (40), Kill 3:52, gaps 1:12–1:25 and 2:44–2:55, rage cap 2:45–2:54; Sunwick, Holy, Flame Breath at 3:20. Every time on the panel fits inside the fight. The model review now sounds like a guild's best warrior talking: "Bloodthirst is the button the Fury rotation is built around … so that one gap cost you twice." The rule-only review gives the §8 sentences: "9 of a possible 33 went unused while you had the rage. Longest gaps 1:12–1:25 and 2:44–2:55."

### Overall verdict: SHIP WITH FIXES

### Per row
- A1 strip icon: FIXED (20 px Bloodthirst icon, all widths). The 9-of-33 numbers are now consistent.
- A2 preparing: FIXED (`Top Fury logs for this boss are still loading.`).
- A2 free: FIXED (`3 things to fix · [icon] Bloodthirst: 9 casts lost first · Read the full review with Premium`).
- A4 rail overflow: FIXED. 1440 is 1440 wide and the rail sits inside the column at 2000.
- A4 icons: FIXED for the Fix first title, the why token and the chip. STILL OPEN, minor: `Heroic Strike` in the Rage-full why, and `Whirlwind` / `Heroic Strike` in the Practice line, are plain text (`log-coach-a4-panel-model-1440.png`). Confirm whether this is data (no spellicons entry) or the token pass.
- A4 empty chips: FIXED (`Bloodthirst · Casts`, `Timelines · 1:12–1:25`, `Resources · 2:45–2:54`). The Whirlwind `· Debuffs` chip is gone.
- A4 gaps sentence: FIXED (DetailCoreCasts as keyed, with ranges).
- A4 Last Stand / killer / 3:20 death: FIXED. Fury is alive the whole fight and there is no defensive clause.
- A4 defensives shown both ways: FIXED.
- A4 spell ranks on a raid boss: FIXED (the line is gone on this pull).
- A4 zone, boss, pool size, talents line: FIXED (`Onyxia (40)`, `Top 25 Fury warriors on Onyxia (40)`, `Talents and race: Forever logs don't record them.`).
- A4 adherence without a reference: FIXED (`Battle Shout was up 100% of the fight, against 97% for the top 25.`).
- A5 link line: FIXED (`The Bloodthirst gap at 2:44–2:55 overlaps the rage cap at 2:45–2:54.`, and it is true).
- A5 swap delivers nothing: FIXED. Ready shows the model's words.
- A6 titles (`Your Fury coach`, `Your coach` for not-available): FIXED.
- A6 not-reviewable reasons: FIXED (`The coach reviews boss pulls only.`, `This pull is under 30 seconds, too short to show a rotation.`).
- A6 no-findings opening: FIXED (`Nothing to fix on this pull that the log can show.`). NEW, STILL OPEN: its footer reads `Reviewed 24 Oct · Based on: .` (`log-coach-a6-state-no-findings-1440.png`). That is an empty basis with a stray full stop. Either print the Doing well bases (`Top 25 Fury warriors on Onyxia (40)`) or drop `Based on:` when there is none.
- A6 raid-buffs-only: FIXED (`No Mark of the Wild at the pull … A druid in the raid provides it.`).
- A6 teasers: FIXED (icon, gold `3 things to fix`, title alone).
- A6 healer: FIXED. The pack is a real Holy priest's: `Greater Heal (Rank 4): 44% overheal`, downrank advice, a flask and Inner Fire under Doing well, a Holy basis, and no "rotation". The Greater Heal icon is missing by data (ruled).
- A6 busy / ready-rules / starter content: FIXED.
- A8 zone: FIXED (`Molten Core … on 5 of 8 bosses`, `Onyxia's Lair, Sat 24 Oct · Fix first: Bloodthirst casts lost`).
- A11 field order: FIXED (Your coach sits before Loot).
- A11 masked link: FIXED in the render. STILL OPEN as proof only: there is no capture from the private test server, and no check that `#coach` reopens the panel.
- A11 Aimed Shot stem: FIXED (`Aimed Shot casts lost — on all 8 bosses`).
- A11 raidNight: FIXED (`Coach reviews for tonight's pulls are ready; open the report or use /report.`).
- A11 loot placeholders: out of scope (rating lane), still visible.

### A12 History, first read
It answers "is my Fix first the same every week?", which is the reason to keep a coach. Rows are one link each, newest first, with a Wipe tag. Fine as a start. What a player would trip on:
- **STILL OPEN** `log-coach-a12-history-older-error-*-390/1440`: under a list, `The coach couldn't read this pull just now.` is wrong ("this pull"?). It needs a list key, e.g. `historyOlderError`: `Older reviews didn't load. Try again.`
- **STILL OPEN** `Log only` on nearly every row (`log-coach-a12-history-rows-list-1440.png`) means nothing to a raider. Nowhere on the page is it explained. Either drop the tag or word it `Rules only` with the same meaning the panel's `Your Fury review` title carries.
- Should, not blocker: twenty flat rows read like a parser dump. Group by night (`Molten Core · 17 Oct`, then the bosses), put the 16 px spell icon before each title as the strip does, and the trend jumps out ("Bloodthirst on 5 of 8 again").
- Capture data (parked): Fury rows under Elyra Duskvale, a Holy priest. The same page still carries `Warden Kelthas` and `Deep Warden` in Every ranked fight, plus `0/9/0` talents at level 60.

### Privacy, Log coach paragraph
A player understands it. What leaves: one fight's numbers, with no name, account, guild or realm. Who gets it: the Vercel AI Gateway, which keeps nothing, then DeepSeek, which writes the review. Officers opening your review also send it. Two small things:
- The parenthesis is circular: `the numbers of one fight (… and the fight's own numbers)`. Say `one fight's numbers: spec, class, …, spell names, casts, uptimes and timings`.
- A raider will ask "can I turn it off for me?" The paragraph only covers the site switching the model off. Add one line on whether a player can opt out, or say plainly that they cannot.

### Data I do not believe (for accuracy)
1. The Molten Core kill order (Discord, A12 history) runs Lucifron … Golemagg, Ragnaros with no Sulfuron Harbinger and no Majordomo Executus, and calls that `8 kills` / `all 8 bosses`. Ragnaros cannot be summoned without Majordomo, and MC has 10 bosses. Use a real night, either 10/10 or a believable partial without Ragnaros.
2. Zulmara gets `Fix first: Bloodthirst casts lost` in the Discord post, while the account mock lists Zulmara as a Marksmanship hunter. Check the class.
3. Logs table fixture: `4,400` damage and `19.0 per sec`, `1.9% active` on a 3:52 Onyxia kill, plus the page title `Onyxia's Lair, fixture night`. That is capture data, but it sits above the coach strip in every A1 and A2 capture.
