// web/src/lib/report/cast-key.ts
// The key a cast row and its measured counts share: the caster and the spell. Split out of
// exact.ts on its own -- exact.ts pulls in the DuckDB query layer (query.ts), which
// dynamically imports the multi-hundred-KB duckdb-wasm package, and CastTable.svelte is part
// of the report island's eager landing view. A static import of this one pure function must
// not drag the rest of exact.ts, and therefore query.ts, into the initial bundle with it.
export function castKey(row: { guid: string; spell_id: number }): string {
  return `${row.guid}|${row.spell_id}`;
}
