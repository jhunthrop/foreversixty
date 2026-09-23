// web/src/lib/account/build-pill.ts
// The build-source pill shown on the account page's Characters rows and the simulator
// landing's "Your characters" rows (spec 2026-09-22 §3.1): "Battle.net · 2 days ago",
// "Addon · today", or the muted "No build yet" for a character the site holds no export
// for. One function so the two lists can never say this two different ways.
import { relativeTime } from '../sim/sources';
import { characterListCopy } from './character-list-copy';
import type { MeCharacter } from './api';

export interface BuildPill {
  label: string;
  /** Which pill colour to use, or null for the muted no-pill text. */
  pillClass: 'pill-blizzard' | 'pill-site' | null;
}

export function buildSourcePill(build: MeCharacter['build'], now: Date = new Date()): BuildPill {
  if (build === undefined) return { label: characterListCopy.noBuildYet, pillClass: null };
  const name =
    build.source === 'blizzard' ? characterListCopy.battlenetSource : characterListCopy.addonSource;
  const pillClass = build.source === 'blizzard' ? 'pill-blizzard' : 'pill-site';
  return { label: `${name} · ${relativeTime(build.captured_at, now)}`, pillClass };
}
