// web/src/lib/guides/rail-stats.ts
// The spec guide's Stat priority rail card (rebuild spec §4.C.3, round-1 mock review
// finding 4): one line per stat the guide's own `statPriority` names, each carrying its
// real per-point scale number -- `StatWeightsPanel`'s own convention (scale factor
// normalized so the top SIGNIFICANT stat at this band reads exactly 1.00), computed fresh
// over just these stats rather than reused from `panel-view.ts`'s own table (that table
// excludes haste as a row entirely; this rail's own mock-board numbers confirm haste is
// meant to compete for the anchor here like any other stat -- see this module's own test
// for the worked Fury example, melee haste normalizing to 1.00).
//
// Ordering (2026-10-05 audit finding): the card's row order comes from the sim, not from
// the guide's own `statPriority` array -- the per-point significant stats sorted by scale
// factor descending, then haste last ("N per 1%"), then insignificant stats last ("Not
// significant"). A guide's written order can disagree with the sim (real example: the Fury
// warrior guide lists Strength and Agility ahead of Critical strike and Hit, but both are
// `insignificant` at band 60) -- the card always follows the sim in that case.
import { statLabelForSpec } from '../sim/weights-display';
import type { BisStatWeight } from '../bis/types';

export interface RailStatRow {
  label: string;
  /** `undefined` means "Not significant" -- the row is never dropped (tenet 4). */
  value: number | undefined;
  /** Haste has no rating in this client and is per 1%, never point-for-point against a
   *  primary stat (owner correction 2026-09-30, `panel-view.ts` `isHasteStat`): it sits
   *  outside the normalization and the rail prints it "per 1%". */
  perPercent?: boolean;
}

function isHasteStat(stat: string): boolean {
  return stat === 'melee_haste' || stat === 'spell_haste';
}

/**
 * `statPriority` (the guide's own frontmatter order, e.g. `["Attack power", "Strength", …]`)
 * matched against `weights` by its class-aware display label (`statLabelForSpec`) -- the
 * same vocabulary `StatPriorityPills` already renders, so a label that does not match any
 * weight row (should not happen on real data) reads as "Not significant" rather than throw.
 */
/**
 * Every guide's own `statPriority` frontmatter (content convention, confirmed across all 27
 * spec guides) says "Critical strike"; `statLabelForSpec`'s own engine vocabulary says
 * "Crit" (`sim/copy.ts`'s pinned `statLabel.crit`) -- the one place these two, otherwise
 * identical, vocabularies diverge. A one-entry alias rather than a second label table:
 * every other guide label (Attack power, Strength, Hit, Melee haste, Spell power, Spell
 * haste, Spell penetration, Intellect, Spirit, MP5, Healing power, …) already matches
 * `statLabelForSpec`'s own word for word.
 */
const GUIDE_LABEL_ALIASES: Readonly<Record<string, string>> = { 'Critical strike': 'Crit' };

/** Sort tier for a matched row: 0 = per-point significant stat (sorted by scale factor
 *  descending within this tier), 1 = haste (reported per 1%, always after the per-point
 *  stats), 2 = insignificant or unmatched (reported "Not significant", always last). */
function sortTier(weight: BisStatWeight | undefined): 0 | 1 | 2 {
  if (weight === undefined || (weight.insignificant ?? false)) return 2;
  return isHasteStat(weight.stat) ? 1 : 0;
}

function scaleOf(weight: BisStatWeight): number {
  return weight.scale_factor ?? weight.weight;
}

export function railStatRows(
  statPriority: readonly string[],
  weights: readonly BisStatWeight[],
  spec: string,
  hasteScaleFactor: number | null = null,
): RailStatRow[] {
  const byLabel = new Map(weights.map((w) => [statLabelForSpec(w.stat, spec), w]));
  const unordered = statPriority.map((label) => ({
    label,
    weight: byLabel.get(GUIDE_LABEL_ALIASES[label] ?? label),
  }));
  const significant = unordered
    .filter((row): row is { label: string; weight: BisStatWeight } => sortTier(row.weight) === 0)
    .map((row) => scaleOf(row.weight));
  const top = significant.length > 0 ? Math.max(...significant) : undefined;

  // The card's own order (2026-10-05 audit finding): the sim's scale factors order the
  // per-point stats, never the guide's own written `statPriority` order. `Array#sort` is
  // stable, so ties within a tier (and the whole insignificant/unmatched tier, which has no
  // secondary key) keep the guide's original relative order.
  const matched = [...unordered].sort((a, b) => {
    const tierA = sortTier(a.weight);
    const tierB = sortTier(b.weight);
    if (tierA !== tierB) return tierA - tierB;
    if (tierA !== 0) return 0;
    return scaleOf(b.weight!) - scaleOf(a.weight!);
  });

  return matched.map(({ label, weight }) => {
    if (weight !== undefined && isHasteStat(weight.stat)) {
      // Live defect 2026-10-05: haste was competing for the 1.00 anchor, so the guide's own
      // first stat read 0.25 and its last read 1.00. Haste is reported per 1%, the same
      // number the BiS page's caption prints (`band.haste_scale_factor`, else its own scale).
      const raw = weight.scale_factor ?? weight.weight;
      const perPercent = hasteScaleFactor ?? (top === undefined ? undefined : raw / top);
      return { label, value: (weight.insignificant ?? false) ? undefined : perPercent, perPercent: true };
    }
    if (weight === undefined || (weight.insignificant ?? false) || top === undefined) {
      return { label, value: undefined };
    }
    const raw = weight.scale_factor ?? weight.weight;
    return { label, value: raw / top };
  });
}
