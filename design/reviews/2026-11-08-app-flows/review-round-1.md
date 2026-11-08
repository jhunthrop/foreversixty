# App flows review, round 1: raider, officer, leveler

Reviewer: wow-player persona, 2026-11-08, on the round-1 captures in this folder (iOS throughout; Android spot-checked on `android-home-raid-day.png`) against `design/specs/2026-11-08-app-flows.md`. Flows with no mock in the round (F6 links, sim result, tiers) were counted from the spec. How each finding was handled is in `review-round-1-dispositions.md`.

## 1. Flow walk and tap counts

| Flow | Persona | Path (taps) | Taps | Budget | Result |
|---|---|---|---|---|---|
| F1 answer raid, In/Bench/Out | Raider M1 | `ios-notifications.png` push, press In (action button) | 0 | 0 | PASS |
| F1 answer raid, Late | Raider M1 | Push, Late, then 15/30/60 chip | 1 chip | 0 | PASS with caveat (S2, S5) |
| F1 answer from launch | Raider M1 | Launch, Home raid card buttons (`ios-home-raid-day.png`) | 1 | 1 from launch | PASS |
| F1 answer on a normal day | Raider M1 | Launch, "Answer" pill (`ios-home-60.png`), buttons on the raid page | 2 | 1 | OVER: the pill is a link, not buttons |
| See who's coming | Raider M1 | Push shows "18 in, 3 late" only; Home shows role bars only | 0 or 1 | 0 | PASS for counts, FAIL for names (S6) |
| F2 raid night, glance | Raider M2 | Live Activity (`ios-live-activity.png`) | 0 | 0 | PASS, content thinner than the spec (S4) |
| F2 raid night, open | Raider M2 | Launch, "Raid night" row (`ios-home-raid-live.png`) | 1 | 1 | PASS |
| F2 after raid | Raider M2 | Activity ends, tap, summary; the ended state is not mocked | n/a | 1 | Unverifiable (S4) |
| F4 loot, report, best | Raider M3 | Push or Home "Last night" row (`ios-home-after-raid.png`), the item sheet or report | 1 | 1 | PASS |
| F3 raid short | Officer M6 | Push opens "Not answered 5", then "Mention in Discord" (`ios-guild-raid-unanswered.png`) | 1 after the push | none | PASS |
| F3 from launch | Officer M6 | Launch (Home has no officer content), Guild tab, "Ping on Discord" (`ios-guild-officer.png`) | 2 | none | Acceptable, see S3 |
| F3 applicant, lock screen | Officer M6 | Accept or Decline on the push | 0 | none | PASS, see M3 |
| F3 applicant, in app | Officer M6 | Push, `ios-guild-applicant.png`, Accept | 1 | none | PASS |
| F4 band or sim push | Leveler M3 | Push lands on the BiS band or the sim result | 1 | 1 | PASS on paper |
| F5 next upgrade | Leveler M4 | Launch, Home top card (`ios-home-34.png`) | 0 | 1 | PASS |
| F5 item sheet | Leveler M4 | Home upgrade row, sheet | 1 | 1 | PASS |
| F5 what to run tonight | Leveler M4 | Home, "Best in slot ›", Where to go | 2 | 1 | OVER: no instance summary on Home (M4) |
| F8 first run | Leveler | Sign in, Pick main, "Turn on notifications", system Allow | 4 plus the Battle.net login | 3 | OVER by one: the system Allow is not counted |
| F9 switch character | All | Crest, character row (`ios-menu.png`) | 2 | 2 | PASS |

## 2. Findings

### MUST
- **M1 One story across every mock.** Home's raid is Wed 9 Dec 8:00 pm; the guild member and officer pages say Thu 10 Dec and show a raid on Mon 7 Dec, before raids open. The after-raid summary and the item sheet show Ragnaros and Molten Core while the raid was Onyxia's Lair and the first raids are Barrow Deeps, Hyjal Summit and Onyxia's Lair. Role counts and item level disagree between Home and the guild page. Fix: one scenario, loot and report from the raid actually raided.
- **M2 The selected character is not the one signing up.** Home and the menu say Thandrel; the guild pages sign up Hollowmend and Dunmaro; the push and confirmation never name the character, and the player has two 60s in the guild. Fix: "as Thandrel, Damage" on the push, the confirmation and the raid card; Pick your main says sign-ups are for the main and how to change it.
- **M3 Accepting an applicant from the lock screen** with no note, history or gear seen, and no next step (the invite is in game, /ginvite). Fix: Decline stays in the notification, Review opens the page; Accept shows the /ginvite line with Copy; attunement, enchants and BiS gap on the applicant page; a sticky bar with spaced buttons.
- **M4 "What to run tonight" is not on Home** (`ios-home-34.png`): no instance grouping or drop chance. Fix: "Run Razorfen Kraul: 2 upgrades" under Your upgrades, to Where to go; drop chance beside each source.

### SHOULD
- **S1 Item sheet has no tooltip**: add the client's tooltip block; drop the "Not tradeable" noise; drop chance on runners-up; "has it / needs it" on the guild rows.
- **S2 In-app Late has no minutes**: Late opens the 15/30/60 chips in the app too.
- **S3 Officer has nothing on Home on raid day**: "3 not answered · Mention" on the raid card for an officer; one verb, Mention in Discord.
- **S4 Live Activity thinner than F2**: roles in the line, the ended state, a ticked Before you pull checklist, and "2 not ready" on the raid card.
- **S5 Lock-screen chips need a Notification Content Extension** (iOS shows actions only after a long press): state the choice.
- **S6 Who's coming is counts only**: names of the healers and tanks on Raid night.
- **S7 Raid-live Home**: "Pulled 1 h 12 m ago", not a clock; the head one line; no duplicate Raid night row; the feed on Your guild.
- **S8 Officer Not answered page**: short role first, "Mark In" not "Confirm", a preview of the Discord mention and a no-bot state; badge counts that reconcile.
- **S9 Times need a zone**: "8:00 pm CST" everywhere.
- **S10 Leveler Home**: the feed defaults to Your spec at the band; Find a guild on the 60 no-guild Home; hide Raids and Officer notification groups without a guild.
- **S11 First run**: 4 taps with the Allow; the permission page's copy runs into stacked buttons; ask in context for a leveler.
- **S12 Notification settings**: the toggle on the title row, not a repeated label; the raid-day nudge hour.

### NIT
- N1 The menu leaves the tab bar live under it. N2 The selected character has no Main tag. N3 The level-40 push must be class-aware (mail or plate) and name the mount. N4 "33 d" drops under its label. N5 Mock the Late confirmation.

## 3. For the accuracy reviewer
Felbough Arm, Cartilage Shapers and Greenhammer may be Forever-new; Belt of Might / Magmadar, Girdle of the Fallen Crusader / Onyxia and Don Julio's Band / Stormpike Guard need checking. The raid-ready text (Drakefire Amulet, flask or elixir, 40-player Onyxia, fire resistance on tanks) reads correctly.

## 4. Verdicts
- Raider: **HOLD** (M1, M2, S4).
- Officer: **HOLD** (M3, with S3 and S8).
- Leveler: **HOLD** (M4, S10).
