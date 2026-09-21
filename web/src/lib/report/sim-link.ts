// web/src/lib/report/sim-link.ts
// The per-combatant "Sim" link the report shows beside plannerLinkFor's own link
// (planner-link.ts): the same fight-sourced character, opened in the simulator instead of
// the planner. The ref extends the fight-level "Sim this fight" link's own ref shape
// (`<report_id>:<fight_index>`) with the combatant's GUID exactly as the report's roster
// carries it (docs/superpowers/specs/2026-09-21-one-product-design.md section 2) --
// Lane B's `fromLoggedFight` reads the optional third part.
import { simFightHref } from '../handoff-links';

export interface SimLink {
  href: string;
  label: string;
}

export function simLinkFor(reportId: string, fightIndex: number, guid: string): SimLink {
  return { href: simFightHref(reportId, fightIndex, guid), label: 'Sim' };
}
