# Sim results build contract (coordinator, 2026-11-04)

Shared by lanes A, B, C, E so the lanes never edit the same file.

- Types for every new field live in `web/src/lib/sim/types.ts` (already on land/simview). Lane E may add doc comments but never renames a field.
- Fixtures: `web/src/fixtures/sim/result-full.json` (new shape, lane E owns and regenerates), `web/src/fixtures/sim/result.json` (older shape, untouched).
- Prerendered fixture route: lane A adds `simfullfixt1` (from `result-full.json`) beside `simfixtureab` in `lib/report/shell-paths.ts` and `pages/sim/[id].astro`.
- Section components, each `{ result: SimResult; names: ActionNames | null }` props plus anything it loads itself, each lazily mounted by lane A from `SavedSim.svelte` (lane A is the only lane that edits SavedSim.svelte):
  - lane B: `web/src/components/sim/SimTimeline.svelte` (section id `timeline`).
  - lane C: `web/src/components/sim/SimGearCard.svelte` (`gear`, includes the final_stats block), `SimTalentsCard.svelte` (`talents`), `SimResourceCard.svelte` (`resources`), `SimAurasCard.svelte` (`auras`), and the existing `RotationCard.svelte` (`rotation`).
- Lane A ships placeholder files with those names (each renders nothing but a `data-testid="<id>-pending"` element) so its branch builds; lanes B and C replace them. On merge the integrator takes B's and C's versions.
- Lanes B and C may add a capture-only mount in their own branch's SavedSim.svelte in a separate commit titled `chore(lane-local): capture mount`; the integrator drops those commits.
- Skeleton heights: each lazy card's final height is in the spec (§4.2, §3.7); lane A reserves it in the lazy fallback, and lanes B and C report their measured heights so the integrator can reconcile.
