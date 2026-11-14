VERDICT: PASS

Round 2 (commit b7201f852, re-read on brand-rotation.png, brand-me.png, brand-guild.png).

1. Resolved: Rotation "How you played it" now reads "Dead early 42 s" in full. Me "Last raid" (Died and Consumables columns) and the Guild leaderboard and Tonight plate show every shipped string without truncation at the 6 px margins and 10 px plate padding.
2. Resolved: the spec states heights per tab (Me 527, Builds 519, Gear 645, Rotation 613, Guild 514) and confirms 645 fits a 768 px screen at scale 1.0. Title rule gap 4 px.
3. NICE, carried over: the header band's right side holds only the close glyph; one linked figure (Forever Score, item level or next raid) could sit there.
4. NICE, carried over: Gear Sources counts could take share bars; Rotation plates end above the footer (shipped layout).

No BLOCKER or SHOULD remains.

---
Round 1 summary: FIX. SHOULD 1: "Dead early 4..." truncated on Rotation because margins and gaps cost about 16 px of width. SHOULD 2: window height grew by 25 to 30 px with no stated contract. Both fixed in round 2.
