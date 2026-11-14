# brand-consistency-reviewer, round 1 (Me with no site activity)

Verdict: PASS

The no-site-activity state matches the web's words and the hub's type, colour and row anatomy. I found no must-fix items.

Checked against `me-empty-l25.png`, `me-empty-l5.png` and `me-empty-l60.png`, the spec, `Locale.lua` and `MeFrame.lua`. I did not open the `-point` or `-setup` captures or the two hub comparison boards, and I did not check the site line's casing against the web.

**Must-fix**
None.

**Matches the web and the hub**
- **Colour:** all colours in `MeFrame.lua` go through `Theme.HEX` and `Theme.rgb`. I found no hardcoded hex in the new sections.
- **Colour roles:**
  - Item names carry quality colours.
  - Gains are success green.
  - Missing enchants are warning red.
  - Labels and notes are muted.
  - Gold appears only on the `foreversixty.gg/app` link in the Me column, and on Right now's act-now rows ("Pick your spec" / "Builds").
- **Type:** headings are cream title face, with a muted right-hand note. Labels and notes are sentence case. There are no forced caps, boxes or rules, which keeps it flat like the web.
- **Copy reuse:** `L.gearBestOf` ("%d of %d best in slot") and `L.gearEnchantMissing` / `L.gearEnchantsMissing` are existing keys, shared with the Gear tab.
- **Web words:** "best in slot" matches `apps/shared/copy` (`shared.paperdoll.best`, `shared.home.onBis`), and "on plan" matches the web's `addon/copy.ts`.
- **Gains:** the gain format matches the web (`+3.9 DPS`, one decimal).
- **Casing and punctuation:** there are no em dashes.
- **Crest:** the header crest is circular and ringed.
- **Item cells:** item icons have quality edges.
- **Row anatomy:** label, value and right-aligned note follow the hub's part-row pattern. The columns line up across levels 5, 25 and 60.
- **Fallbacks:** the Level 5 and Level 60 captures are coherent. Missing rows drop out without leaving gaps. The level-60 site line is the "first logged raid" variant, as the spec says.

**Shoulds**
1. Add the new `Locale.lua` strings (`meStanding`, `meRested`, `meMaxLevel` and similar) to the shared-copy audit, or note that they are addon-local. The "Nothing to chase" and "Level 10, your first talent point" lines have no web or shared equivalent, so they can drift. Keep them terse and game-word only.
2. The Level 25 capture shows `15 of 15 on plan` while the spec example says `11 of 16`. That is probably only a different fixture. Confirm that the figure comes from `BuildsPane.figures` and that the two sources agree.
3. Level 60 shows the Gear note `9 of 16 best in slot` beside "Max level" and `51 of 51 on plan` without making clear that "best in slot" is band-60 scoped. The web paperdoll uses the same words, so this is fine as is. It is polish only.
4. The `Experience` and `Level` row labels are muted at x 0 with a wide gutter (x 86). Where only two rows exist (Level 60), the left column looks sparse beside a full right column. The spec accepts this. If density reviewers complain, the lever is the omitted rows 3 and 4, not new chrome.
5. The Level 25 `2 enchants missing` row sits flush left with no icon, as the spec says. Warning red on the count plus muted slot names reads fine. Optionally check that its text is not clipped at narrow slot lists.
