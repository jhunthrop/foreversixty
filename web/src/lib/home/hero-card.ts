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
  /** The muted tag beside it: "level cap 30", "First raids Dec 9"; empty while the tiles say it all. */
  tag: string;
  /** The time left to the next milestone, for the three tiles; null when there is nothing to count to. */
  countdown: Countdown | null;
  /** The class line's second half: the bands the example's picks come from. */
  bandsLine: string;
}

export interface Countdown {
  days: number;
  hours: number;
  minutes: number;
}

const MINUTE_MS = 60 * 1000;
const HOUR_MS = 60 * MINUTE_MS;
const DAY_MS = 24 * HOUR_MS;

/** The whole days, hours and minutes from `now` to `target`; null once the moment has passed. */
export function countdownTo(target: Date, now: Date): Countdown | null {
  const left = target.getTime() - now.getTime();
  if (left < 0) return null;
  return {
    days: Math.floor(left / DAY_MS),
    hours: Math.floor((left % DAY_MS) / HOUR_MS),
    minutes: Math.floor((left % HOUR_MS) / MINUTE_MS),
  };
}

/** The moment a dated entry starts: its exact instant when it has one, else its day's start (UTC). */
function momentOf(entry: DateEntry): Date {
  return new Date(entry.at ?? entry.iso);
}

/**
 * The card for `now`: the beta's while it is live (or has not opened), the launch's from the day
 * after the beta ends until launch, then the first raids' until they open. Each counts down in
 * the tiles. An entry list with no launch keeps the beta wording rather than guess.
 */
export function heroCardFor(dates: readonly DateEntry[], now: Date = new Date()): HeroCard {
  const rows = timelineRows(dates, now);
  const beta = rows.find((row) => row.role === 'beta');
  const launch = rows.find((row) => row.role === 'launch');
  const raids = rows.find((row) => row.role === 'raids');
  const betaCard: HeroCard = {
    mode: 'beta',
    eyebrow: copy.betaEyebrow,
    tag: copy.levelCapLabel,
    countdown: null,
    bandsLine: copy.betaBandsLine,
  };
  if (beta === undefined || launch === undefined || beta.status === 'live' || beta.status === 'next') {
    return betaCard;
  }
  const untilLaunch = countdownTo(momentOf(launch), now);
  if (untilLaunch === null) {
    const untilRaids = raids === undefined ? null : countdownTo(momentOf(raids), now);
    return {
      mode: 'launched',
      eyebrow: copy.launchedEyebrow(launch.key),
      tag: raids !== undefined && untilRaids !== null ? copy.firstRaidsTag(raids.key) : copy.launchedTag,
      countdown: untilRaids,
      bandsLine: copy.bandsLine,
    };
  }
  return {
    mode: 'launch',
    eyebrow: copy.launchEyebrow(launch.key, launch.note),
    tag: '',
    countdown: untilLaunch,
    bandsLine: copy.bandsLine,
  };
}
