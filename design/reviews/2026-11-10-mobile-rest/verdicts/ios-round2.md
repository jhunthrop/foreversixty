# iOS secondary screens, round 2 verdicts (2026-11-10)

Goldens recorded on CI (ios.yml record, run 42924962651).

| Group | ux-designer | wow-player | ai-slop | info-density (r1) | brand-consistency | guild-officer |
|---|---|---|---|---|---|---|
| guides | SHIP | SHIP | PASS | DENSE | DRIFT | |
| character | SHIP | SHIP | PASS | DENSE | CONSISTENT | |
| bis-search | SHIP | SHIP | FAIL | DENSE | CONSISTENT | |
| account | SHIP | SHIP | PASS | DENSE | CONSISTENT | |
| guild-raids | SHIP | SHIP | PASS | DENSE | DRIFT | WORTH PAYING FOR |
| sims-reports-talents | NO SHIP | SHIP | FAIL | DENSE | DRIFT | |

Answered, not changed:
- **"15 DPS" and "24 of 40".** 3 + 6 + 15 = 24, so the 24 already includes the one late raider (coming = In plus Late, the agreed count model). The Android sheet the reviewer compared against predated the count change.
- **The composition panel's role counts.** They come from the API's composition block, by an earlier coordinator ruling ("every count and every present or missing comes from the API"). Brand rated the difference NICE.

## Round 3 work (iOS, plus Android's BiS control bar)
- **Bar titles.** Sims, the sim result and Top Gear move onto the N3 bar, with no title at rest.
- **Band heads.** The guide loading and error states and the change entry get one.
- **Raid flow parity hosts.** They render inside a NavigationStack, with back and share.
- **Buttons:**
  - Not answered's row actions and Mention all are outline;
  - the officer row's kinds match Android;
  - the outline button reads at AA.
- **Ink.** mutedGround for Shared build's meta, Your builds' dates and Talents' unspent line.
- **Signed out.** The band runs under the status bar.
- **Delete account** sits on a panel.
- **BiS control bar.** Both platforms use the short range label (`60 · Phase 1 raid`), so it stays one line. The picker keeps the opening date.
- **The spec sheet** gets its title row.

Not changed:
- **BiS set-line wording** ("at best in slot for your level range"). It is a shared copy key in `en.json`, which another lane holds under its copy lock. It goes to that lane as a follow-up: "{dps} {unit} · Updated {date}".
