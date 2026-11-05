// web/src/lib/sim/spell-info.ts
// What the saved result page knows about a spell beyond its name: the icon, the rank, the
// cooldown and the game tooltip (spec 2026-11-04-sim-results 3.5, 3.6). Two static build files,
// both fetched once and both optional: a build without them still renders every row, with the
// name in plain text and the icon box empty, never a broken image.
import { dataUrl, loadOptional } from '../planner/load';
import type { GameTipModel } from '../tooltips/game-tip';
import { parseActionKey } from './action-names';

export interface SimSpellFact {
  name: string;
  /** 0 for a spell the client gives no rank. */
  rank: number;
  cooldown_ms: number;
  tip: GameTipModel;
}

/** `/data/<build>/sim-spells/<class>.json`. */
export interface SimSpellBundle {
  spells: Record<string, SimSpellFact>;
}

export interface SpellInfo {
  /** The icon's url for an action key, or null when the build has none for it. */
  iconUrl: (key: string) => string | null;
  /** The client's rank and cooldown for a spell id. */
  fact: (spellId: number) => SimSpellFact | undefined;
}

export const NO_SPELL_INFO: SpellInfo = { iconUrl: () => null, fact: () => undefined };

/** The one load: icons from the build's spell table, facts from the class's bundle. */
export async function loadSpellInfo(build: string, classSlug: string): Promise<SpellInfo> {
  const [icons, bundle] = await Promise.all([
    loadOptional<Record<string, string>>(dataUrl(build, 'spellicons.json'), {}),
    loadOptional<SimSpellBundle>(dataUrl(build, `sim-spells/${classSlug}.json`), { spells: {} }),
  ]);
  return {
    iconUrl: (key) => {
      const parsed = parseActionKey(key);
      if (parsed === null || parsed.kind !== 'spell') return null;
      const stem = icons[parsed.label];
      return stem === undefined ? null : dataUrl(build, `icons/${stem}.webp`);
    },
    fact: (spellId) => bundle.spells[String(spellId)],
  };
}
