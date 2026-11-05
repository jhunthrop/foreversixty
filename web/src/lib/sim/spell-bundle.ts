// web/src/lib/sim/spell-bundle.ts
// Build-time half of lib/sim/spell-info.ts: one class's spell facts from the client's own
// `spelltips/<class>.json`. Node only (it reads the build directory); the published file is what
// the page reads.
import { readBuildJson } from '../guides/build-files';
import { spellTipModel, type SpellTipRow } from '../tooltips/game-tip';
import type { SimSpellBundle, SimSpellFact } from './spell-info';

interface ClassSpellTips {
  spells: Record<string, SpellTipRow>;
}

function fact(row: SpellTipRow): SimSpellFact {
  return { name: row.name, rank: row.rank, cooldown_ms: row.cooldown_ms, tip: spellTipModel(row, row.icon) };
}

export function simSpellBundle(build: string, classSlug: string): SimSpellBundle {
  const rows = readBuildJson<ClassSpellTips>(build, `spelltips/${classSlug}.json`)?.spells ?? {};
  return {
    spells: Object.fromEntries(Object.entries(rows).map(([id, row]) => [id, fact(row)])),
  };
}
