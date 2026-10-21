// The hero's example card (home rebuild spec §3.A.2), whose first line used to be the hand-written
// "Beta is live · level cap 30" and went false the day the beta ended. It follows the clock the way
// the timeline does (lib/home/timeline.ts): the beta entry says when the beta is live, the launch
// entry says what the card reads after it, and a small script re-reads both against the visitor's
// clock, since the page is built once per deploy.
import { homeHeroStatePanelCopy as copy } from '../home-landing-copy';
import { timelineRows, type DateEntry } from './timeline';

export type HeroCardMode = 'beta' | 'launch' | 'launched';

export interface HeroCard {
  mode: HeroCardMode;
  /** The gold line: "Beta is live", "Launch Nov 4 · 15:00 PST". */
  eyebrow: string;
  /** The muted tag beside it: "level cap 30", "in 15 days". */
  tag: string;
  /** The class line's second half: the bands the example's picks come from. */
  bandsLine: string;
}

const DAY_MS = 24 * 60 * 60 * 1000;

/** Whole calendar days from `now` to a date-only ISO string, both read in UTC. */
function daysUntil(iso: string, now: Date): number {
  const target = Date.parse(iso);
  const today = Date.UTC(now.getUTCFullYear(), now.getUTCMonth(), now.getUTCDate());
  return Math.round((target - today) / DAY_MS);
}

/**
 * The card for `now`: the beta's while it is live (or has not opened), the launch's from the day
 * after the beta ends until launch day is over, then "Launched". An entry list with no launch
 * keeps the beta wording rather than guess.
 */
export function heroCardFor(dates: readonly DateEntry[], now: Date = new Date()): HeroCard {
  const rows = timelineRows(dates, now);
  const beta = rows.find((row) => row.role === 'beta');
  const launch = rows.find((row) => row.role === 'launch');
  const betaCard: HeroCard = {
    mode: 'beta',
    eyebrow: copy.betaEyebrow,
    tag: copy.levelCapLabel,
    bandsLine: copy.betaBandsLine,
  };
  if (beta === undefined || launch === undefined || beta.status === 'live' || beta.status === 'next') {
    return betaCard;
  }
  if (daysUntil(launch.iso, now) < 0) {
    return {
      mode: 'launched',
      eyebrow: copy.launchedEyebrow(launch.key),
      tag: copy.launchedTag,
      bandsLine: copy.bandsLine,
    };
  }
  return {
    mode: 'launch',
    eyebrow: copy.launchEyebrow(launch.key, launch.note),
    tag: copy.launchCountdown(daysUntil(launch.iso, now)),
    bandsLine: copy.bandsLine,
  };
}
