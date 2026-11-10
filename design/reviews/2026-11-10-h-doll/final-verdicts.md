# Lane h-doll: final verdicts

Each reviewer's last verdict on the final captures in after/:
- ux-designer, round 4 (2bea2fd82): VERDICT: SHIP. It asked for the report list widening that landed in 7a60803b9.
- wow-player, round 4 (2bea2fd82): VERDICT: SHIP. Blockers: none.
- info-density-reviewer, round 4 (2bea2fd82): VERDICT: DENSE.
- ai-slop-reviewer, round 6 (5ba7cae6c): VERDICT: PASS.

Rounds 5 and 6 changed only the report frame's width and how the centre tooltip breaks an enchant. The ai-slop-reviewer judged both on the final report capture.

## Routed to the coordinator (outside this lane)
- **Shared ItemTooltip:**
  - The stat lines read "+28 Crit" / "+20 Hit" where the client uses rating sentences.
  - Weapon damage and speed lines are missing.
- **Data accuracy:**
  - The Lionheart Helm's required level, item level and bind.
  - The identical "Presence of Might" text on the helm and legs (a real Zandalar kit, used on both slots in the capture set).
- **CharacterStrip identity row:**
  - "Item level 60.0" carries a decimal.
  - The "Addon export" source pill is tell 17.
- **Report Gear at the pull:**
  - The action "Open in the planner ›" should point to Best in slot (one job per page).
  - The phone strip drops the class (shared PaperdollFrame list strip).
- **Needs data:**
  - missing-enchant flags (per-slot enchantable data);
  - gems and bonus ids;
  - deltas against the last pull or run.
- **Spec:** web layout spec §4.2's report line "Item level 66 · Enchanted" now reads "Item level 66" with the named enchant on its own green line ("Enchanted" only when the id has no name).
