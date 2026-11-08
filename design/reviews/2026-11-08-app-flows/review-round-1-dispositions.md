# Round 1 dispositions

How each finding of `review-round-1.md` was handled, in the spec (`design/specs/2026-11-08-app-flows.md`) and the mocks (`design/mocks/ws_app_flows.py`, `ws_app.py`). The captures in this folder are the round-2 state.

| # | Finding | Disposition |
|---|---|---|
| M1 | One story across every mock | Fixed. One scenario everywhere: Thandrel, Fury Warrior 60, the main, officer of Olympus XXVII; Onyxia's Lair, Wed 9 Dec, 8:00 pm CST, the day the first raids open; 18 in, 3 late; 2 of 2 tanks, 5 of 8 healers, 11 of 30 damage. The Guild tab's first screen is now app-composed on that scenario (`guild-member`, `guild-officer`). The loot is Onyxia's from data/builds 1.60.1.70790: Deathbringer (8%, item level 75) with Vis'kag the Bloodletter as the runner-up; the report and the best are Onyxia's. The raid page, roster and loot pages pushed from the Guild tab still come from the guild parity fixture until the Guild build step re-points them. |
| M2 | Who is answering | Fixed. `as Thandrel, Damage` on the push, both confirmations, Home's raid card and the Guild tab; Pick your main says sign-ups are for the main and how to switch (E7, F1, F8). |
| M3 | Accept from the lock screen | Fixed. The push offers Review and Decline; the applicant page adds Raid ready (attunement, enchants, best-in-slot gap) and a sticky Accept / Decline bar; after Accept, `/ginvite Mirelle` with Copy (`guild-applicant-accepted`). |
| M4 | What to run tonight on Home | Fixed. One instance line under Your upgrades on every signed-in Home (`Run Razorfen Kraul: 2 upgrades`), to Where to go (F5). Drop chance beside each upgrade's source is in the spec; the web home mock's upgrade rows do not carry it yet. |
| F1 OVER | Answering on a normal day took 2 taps | Fixed. The normal-day next-raid line carries In, Late, Bench, Out. |
| S1 | Item sheet tooltip | Fixed: the client's lines (binds, slot and type, damage and speed, the chance on hit, required level); drop chances; has it / needs it. |
| S2 | In-app Late minutes | Spec: Late opens the same three chips in the app. |
| S3 | Officer on raid day | Fixed: `home-raid-day-officer`; one verb, Mention in Discord. |
| S4 | Live Activity and checklist | Fixed: roles in the Live Activity line, the ended state (Live Activity and notification), the ticked Before you pull checklist on Raid night, the open count on Home's raid card. |
| S5 | Late chips on the lock screen | Spec: an iOS Notification Content Extension; Android's expanded notification. Mock labels say long-pressed / expanded. |
| S6 | Names, not only counts | Fixed on Raid night; spec says tanks and healers by name. |
| S7 | Raid-live Home | Fixed: `Pulled 1 h 12 min ago`, the head one line, one Open raid night link, the feed on Your guild. |
| S8 | Officer Not answered page | Fixed: healers first, Mark In, the Discord preview; the no-bot state is in the spec; the badge is 6 everywhere (5 unanswered + 1 applicant), the Needs you count. |
| S9 | Time zones | Fixed: CST on every raid time (E7). |
| S10 | Leveler Home | Fixed: the feed opens on Your spec; Find a guild on the no-guild 60 Home; the duplicate Join a guild line removed; the spec hides Raids, Guild ladder and Officer settings without a guild. The level-34 web mock's own Find a guild card still right-aligns its body; that is the web home mock's (a follow-up there). |
| S11 | First run | Spec: 3 taps without a guild, 4 with the system Allow; buttons stacked; no page for a leveler, the ask comes in context. |
| S12 | Notification settings | Fixed: switches on the title rows; the raid-day reminder hour. |
| N1 | Menu over a live tab bar | Fixed: the drawer is modal and covers it. |
| N2 | Main tag | Fixed. |
| N3 | Class-aware level-40 push | Spec: mail or plate by class, and the mount. |
| N4 | "33 d" | Fixed: a muted Tag, `In 33 days`. |
| N5 | Late confirmation | Fixed: `You're Late (30 min) for Onyxia`. |
| 3 | Items to verify | The new scenario uses only Onyxia's loot from the build data. Felbough Arm, Sentinel's Medallion and Cartilage Shapers on the level-34 Home come from the web home fixture (`web/tests/e2e/fixtures/home/home-34.json`), outside this round. |
