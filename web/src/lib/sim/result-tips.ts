// web/src/lib/sim/result-tips.ts
// The saved result's tooltip keys and how each resolves: the page's spells from the class bundle
// and its one explanatory note (build ruling R7's Crit header) through the same shared panel. The
// Talents section's trees carry their own text and their own host (ui/TalentTrees).
import type { GameTipModel } from '../tooltips/game-tip';
import { damageCopy } from './result-copy';
import type { SpellInfo } from './spell-info';

/** The Crit column header's note. */
export const CRIT_NOTE_TIP = 'note:crit';

const CRIT_NOTE: GameTipModel = { name: damageCopy.headers.crit, lines: [], description: damageCopy.critTip };

const SPELL_TIP = /^spell:(\d+)$/;

/** The model for a `data-game-tip` key on this page; null when the build carries nothing for it. */
export function resultTip(key: string, spells: SpellInfo): GameTipModel | null {
  if (key === CRIT_NOTE_TIP) return CRIT_NOTE;
  const match = SPELL_TIP.exec(key);
  return match === null ? null : (spells.fact(Number(match[1]))?.tip ?? null);
}
