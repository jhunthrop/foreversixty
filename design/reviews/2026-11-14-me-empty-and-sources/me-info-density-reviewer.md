# info-density-reviewer, final (round 4): Me with no site activity

Captures: me-empty-l25.png, me-empty-l25-point.png, me-empty-l5.png, me-empty-l60.png (design/mocks/addon_hub_me_empty.py). Earlier rounds: me-info-density-reviewer-round-1.md to -round-3.md.

Verdict: PASS

I checked the round 4 captures `me-empty-l25.png` and `me-empty-l60.png`. The layout holds at both levels, the source line no longer cuts mid-word, and nothing is invented.

**Must-fix**
None.

**Shoulds**
1. **The drop chance is dropped to make the line fit, which is the opposite of what I asked.** The Impaling Harpoon row now reads `The Deadmines · Captain Greenskin`, and the `30%` is gone. My round 3 note asked to drop the area name and keep the chance. The chance is the number a raider acts on, and the area is implied by the boss. Order the drop to fall back as: area, then boss, then chance. For this row it would be `Captain Greenskin · 30% drop`. It costs one comparison in the existing cut, and it is the same data.
2. **The level 60 right column still has about 330 px of empty ground** under Your sim, while the left runs to the footer. Gear sits left only at 60, as intended, and the spec accepts it. If the owner wants it level, the real fillers are `+N more` upgrades (`GearTab.upgrades(doll).count`) and a `+N DPS` total (`.total`). Optional.
3. **Professions cap.** `Skinning 112 of 150` if the fifth return of `GetSkillLineInfo` gives the max rank. Optional.

**Needs data**
- Gold, played time and reputation are not read by the addon.
- Score, logs and sim stay behind the companion call to action.

Files: /Users/jh/code/forever-me-tab/design/reviews/2026-11-14-me-empty-and-sources/me-empty-l25.png, /Users/jh/code/forever-me-tab/design/reviews/2026-11-14-me-empty-and-sources/me-empty-l60.png.
