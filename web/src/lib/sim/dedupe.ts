// web/src/lib/sim/dedupe.ts
// A name is this lane's whole identity for a talent loadout or a named set alike --
// `TalentLoadout`/`GearSet` (bulk-types.ts) carry nothing else the engine could tell two
// apart by -- so two entries sharing one can never both render. TalentCandidates.svelte and
// NamedSets.svelte each key an `{#each}` block and a `sim-loadout-<name>`/`sim-set-<name>`
// test id purely by name; two elements sharing one test id is not just a confusing screen,
// it fails Playwright's strict-mode "one match" rule the moment a test looks for that id.
// One shared helper so both components resolve a collision the same way, rather than each
// inventing its own precedence rule.

/**
 * Keeps the first entry seen for each name and drops every later one whose name is already
 * taken -- across as many lists as the caller wants to combine, by threading one `taken` set
 * through them in call order. Pass a pre-seeded set to give a source the caller does not
 * render here (or a prior call's own kept names) first claim on a name.
 */
export function dedupeByName<T extends { name: string }>(taken: Set<string>, entries: readonly T[]): T[] {
  const kept: T[] = [];
  for (const entry of entries) {
    if (taken.has(entry.name)) continue;
    taken.add(entry.name);
    kept.push(entry);
  }
  return kept;
}
