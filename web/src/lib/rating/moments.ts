// web/src/lib/rating/moments.ts
// Where a rating moment jumps to, per docs/superpowers/plans/2026-09-21-rating-web.md's
// Ruling 4: a death uses DeathsTab's own pre-existing anchor convention verbatim; every
// other kind is built here from a spell id plus this player's own guid, since
// ExchangeRow/AuraTrack carry no per-event timestamp for the spec's literal <atMs>
// convention to read.
import { coachAnchorHref } from '../coach/anchors';
import type { CoachAnchor } from '../coach/types';
import { MISSING_FIGHT } from '../report/url';
import type { RatingMoment } from './types';

/** The moment kinds that open through the coach's anchor builder (the report URL encoder). */
const ANCHOR_KINDS = ['cast', 'time', 'resource'] as const;
type AnchorMomentKind = (typeof ANCHOR_KINDS)[number];

const isAnchorKind = (kind: string): kind is AnchorMomentKind =>
  (ANCHOR_KINDS as readonly string[]).includes(kind);

function anchorOf(kind: AnchorMomentKind, moment: RatingMoment): CoachAnchor {
  return { kind, spell_id: moment.spell_id, at_ms: moment.at_ms, to_ms: moment.to_ms };
}

const TAB_BY_KIND: Record<string, { tab: string; prefix: 'exchange' | 'aura' } | undefined> = {
  interrupt: { tab: 'interrupts', prefix: 'exchange' },
  dispel: { tab: 'dispels', prefix: 'exchange' },
  'carried-debuff': { tab: 'debuffs', prefix: 'aura' },
  'utility-uptime': { tab: 'buffs', prefix: 'aura' },
};

export function momentHref(fightIndex: number, playerGuid: string, moment: RatingMoment): string | null {
  if (playerGuid === '') return null;
  if (moment.kind === 'death') {
    if (moment.anchor === undefined || moment.anchor === '') return null;
    return `?fight=${fightIndex}&tab=deaths#${moment.anchor}`;
  }
  if (isAnchorKind(moment.kind)) {
    // MISSING_FIGHT as the first fight keeps `fight=` in the link, as the literal forms here do.
    return coachAnchorHref(fightIndex, MISSING_FIGHT, playerGuid, anchorOf(moment.kind, moment));
  }
  const target = TAB_BY_KIND[moment.kind];
  if (target === undefined || moment.spell_id === undefined) return null;
  return `?fight=${fightIndex}&tab=${target.tab}#${target.prefix}-${playerGuid}-${moment.spell_id}`;
}
