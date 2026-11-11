# brand-consistency-reviewer, round 2: the Guild tab vs the web guild pages at 390 px

Date 2026-11-11. Checked: GuildTabViews.swift, GuildTabParts.kt, assignments-view.ts, parity/ios-full-guild-tab-officer.png, round1-response.md.

## Round 1 findings re-judged
1. (SHOULD) Assignment name colours: FIXED. iOS `AssignedNames.text` and Android's equivalent use FSColor.gold for own, FSColor.muted for the rest, matching `nameTone` (gold / muted). The officer render shows only `Thandrel` in gold. Off-role: `off_role` is not in the app models, so the warn tone is absent. Accepted as a recorded omission (round1-response): the screens are read-only and the apps do not carry the field. NICE follow-up: add it when the contract carries it.
2. (NICE) Your job as a section: stays a spec 2.3a ruling. ALLOWED.
3. (NICE) Spread: shown at every level, label-under-figure, figures in gold for own. Accepted; spec's `close` wording vs `55 to 59` stays a spec/build wording note only.
4. (NICE) Head line always carries the realm (`Alliance · US PvP`). FIXED. Counting `members` vs the web's `raiders` is the recorded fixture/contract choice.
5. (NICE) Type sizes: head name is PanelTitle, figures FSType.itemLg / FsType.itemLg. FIXED. A leftover: iOS `GuildLineRow` has an inline `TypeRole(size: 15, lineHeight: 20)` for the strong trail and `.system(size: 13)` for the chevron. Not a hex, tiny, shared row pattern; NICE, move to a token role when next touched.
6. (NICE) Tab glyph weight: the site's own glyph, flagged for the owner. Accepted.

## New checks
- Colour: no hex in either file; class colours, quality colours, gold, muted all FSColor / FsColor.
- Crest: ringed circle, as web. Flat variant at 32 for rows (under 44), correct.
- Copy: `Ping on Discord`, `Open in planner`, `0 of 2`, `1 of 1` match the web planner words and shared.assign keys. No em dashes seen.
- Surfaces: hairline panels, no gradients or shadows beyond the apps' round-5 panel system (owner ruling).

VERDICT: CONSISTENT
