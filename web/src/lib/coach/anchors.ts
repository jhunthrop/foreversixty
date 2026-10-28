// web/src/lib/coach/anchors.ts
// Where a coach evidence anchor opens in the report. Every link goes through url.ts's own
// encoder (`reportSearch` over `withState(defaultState(firstFight), ...)`), so a coach link and a
// hand-made one are the same string. A death keeps DeathsTab's literal anchor form, which is the
// engine's DeathAnchor format. Unknown kinds and anchors missing what their kind needs are null.
import { defaultState, reportSearch, withState, type ReportState, type Tab } from '../report/url';
import type { CoachAnchor } from './types';

/** The window a single moment opens on the timelines, when the anchor names one time only. */
export const POINT_WINDOW_MS = 5000;

type StatePatch = Partial<ReportState>;

function abilityPatch(tab: Tab, guid: string, anchor: CoachAnchor): StatePatch | null {
  if (anchor.spell_id === undefined) return null;
  return { tab, source: guid, ability: anchor.spell_id };
}

function windowPatch(guid: string, anchor: CoachAnchor): StatePatch | null {
  if (anchor.at_ms === undefined) return null;
  const end =
    anchor.to_ms !== undefined && anchor.to_ms > anchor.at_ms ? anchor.to_ms : anchor.at_ms + POINT_WINDOW_MS;
  return { view: 'timelines', source: guid, start: anchor.at_ms, end };
}

function patchFor(guid: string, anchor: CoachAnchor): StatePatch | null {
  switch (anchor.kind) {
    case 'cast':
      return abilityPatch('casts', guid, anchor);
    case 'healing':
      return abilityPatch('healing', guid, anchor);
    case 'aura':
      return anchor.spell_id === undefined
        ? null
        : { tab: anchor.debuff === true ? 'debuffs' : 'buffs', source: guid };
    case 'buffs':
      return { tab: 'buffs', source: guid };
    case 'resource':
    case 'time':
      return windowPatch(guid, anchor);
    default:
      return null;
  }
}

/**
 * DeathsTab's own anchor is `death-<guid>-<at_ms>`. The pack carries only `at_ms` (no identifier
 * leaves the pack), so it is rebuilt from the viewed player's guid; an older API that still sends
 * the finished `death` string is honoured.
 */
function deathHref(fightIndex: number, guid: string, anchor: CoachAnchor): string | null {
  const id = anchor.death ?? (anchor.at_ms === undefined ? '' : `death-${guid}-${anchor.at_ms}`);
  return id === '' ? null : `?fight=${fightIndex}&tab=deaths#${id}`;
}

/** The report query string a coach anchor opens, or null when it cannot be built. */
export function coachAnchorHref(
  fightIndex: number,
  firstFight: number,
  playerGuid: string,
  anchor: CoachAnchor,
): string | null {
  if (playerGuid === '') return null;
  if (anchor.kind === 'death') return deathHref(fightIndex, playerGuid, anchor);
  const patch = patchFor(playerGuid, anchor);
  if (patch === null) return null;
  const search = reportSearch(
    withState(defaultState(firstFight), { fight: fightIndex, ...patch }),
    firstFight,
  );
  return anchor.kind === 'aura' && anchor.debuff !== true
    ? `${search}#aura-${playerGuid}-${anchor.spell_id}`
    : search;
}
