# Officer tab (L7), round 1: good raid-night tool, a few rough edges

Checked: spec 4.6 and 8, both boards, the nine L7 renders, the `L.off*` copy, and `LootCouncil.give` (GiveMasterLoot via the candidate index).

## Verdict: FIX

1. MUST: the armed Give button covers the Won and Att. cells; Won is the number an officer checks before handing a drop over. Reserve the width, or keep the armed text short and name the player in the hint line.
2. MUST: the Enchants row cuts off names ("Kethra 1, Velisandre 2, S…"). Show "+N more" so the cut is plain, every name in a hover.
3. SHOULD: "Won" has no unit; a hover reading items from you this week, your own loot only.
4. SHOULD: the Wears column is cut; the full name in the hover.
5. SHOULD: Gap has two blank states (blank and Best); a muted `-` for no report.
6. SHOULD: readiness: No food, repair (needs data), attunement must show when someone lacks it; make the ready-check row reachable.
7. SHOULD: "Online (24)" reads as guild online: "Online raiders".
8. SHOULD: "waiting" on a boss item reads like a roll countdown: "no answers yet".
9. NICE: an "on list" pick with a lower gain than the top row could mislead; say so in a hover.
10. NICE: the guild-chat reminder is right.
11. NICE: Give is safe (two clicks, the client's candidates, a plain refusal).

No site jargon found. Figures are plausible for Forever.
