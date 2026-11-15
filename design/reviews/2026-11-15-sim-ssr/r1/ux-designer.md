Verdict: FIX (round 1, commit 44ed408f4)

Hydration is clean: CLS 0, heights match (1611 at 1440, 3311 at 390).

Must fix
1. Guide and BiS touch targets on phone: 44 px tall but about 30 and 18 px wide with a 12 px gap. Give both min-w-11 justify-center below md.
2. Three equal gold links per spec line: spec name, Guide and BiS all gold. Make Guide and BiS secondary (muted at rest, gold on hover/focus); the spec name stays the one gold link.

Should fix
3. Stat line orphans at 390 (a lone "Hit rating 0.83" on a second line).
4. Phone rhythm: gap inside a spec block looks like the gap between blocks; use gap-3.
5. Desktop: ml-auto pushes Guide/BiS about 300 px from the DPS figure.
6. Phone header: check "All stat weights ›" is 44 px tall.

Good: Row align="start" fixes uneven centring; class order with no bars matches the tier-list rule; no-JS only about 21 px taller.
