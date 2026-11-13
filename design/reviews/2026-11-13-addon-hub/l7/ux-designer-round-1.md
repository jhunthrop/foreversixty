# Officer tab (L7) review, round 1: roster, readiness, push and master loot

Checked: the 9 renders in l7/ against hub-officer.png and hub-officer-loot.png, spec 3, 4.6 and 8, and the mock docstring's stated deviations.

## Verdict: FIX

The tab answers its question; the stated deviations are sound. One blocker is in the loot session.

1. MUST: the armed `Give to Ilvar?` button covers Ilvar's Won and Att. cells. Keep the armed button inside the Give column, or shift Won and Att. so it never overlaps.
2. SHOULD: Enchants cuts mid-name; cut on whole names with a muted `+1`, full list in the hover.
3. SHOULD: role counts all warning red; draw muted, colour only an empty role.
4. SHOULD: render the pre-30 s nobody-answering state (gain `-` muted, waiting).
5. SHOULD: render `Sent to guild` in the footer, a disabled Give with the client's reason, keyboard focus.
6. SHOULD: `open` in the right column drawn in client gold; draw it plain or say why.
7. NICE: `DPS` and `Wears` headings sit 4 px apart.
8. NICE: in no-companion, put the reason for the Sign-up dashes in the roster's note.
