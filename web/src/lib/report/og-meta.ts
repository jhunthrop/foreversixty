// web/src/lib/report/og-meta.ts
// What the Worker writes into a shell's head. Pure, so the copy is reviewable and tested
// without a Workers runtime; src/worker.ts only wires the values into HTMLRewriter.
//
// Voice: reference, not pitch. A description states the numbers and the date and stops.
import { rulesetLabel, type CharacterPath } from '../characters';
import type { ReportMeta } from './types';

export const SITE_BASE_URL = 'https://foreversixty.gg';

export type ShellKind = 'report' | 'rankings' | 'character' | 'guild';

export interface ShellMeta {
  /** The whole <title>, including the site suffix. */
  title: string;
  description: string;
  /** Absolute og:image url. */
  image: string;
  /** Absolute canonical and og:url. */
  canonical: string;
}

/**
 * Only the report has a rendered card: the API draws one per report the way it draws build
 * cards. The other three shells take the site's own static card, which is already built by
 * src/pages/og/[...path].png.ts and needs no per-row rendering.
 */
const SITE_CARD = `${SITE_BASE_URL}/og/index.png`;

function isoDate(value: string): string {
  const at = new Date(value);
  return Number.isNaN(at.getTime()) ? 'an unknown date' : at.toISOString().slice(0, 10);
}

function plural(count: number, one: string, many: string): string {
  return `${count} ${count === 1 ? one : many}`;
}

export function titleize(slug: string): string {
  return slug
    .split('-')
    .filter((part) => part !== '')
    .map((part) => part.charAt(0).toUpperCase() + part.slice(1))
    .join(' ');
}

export function reportShellMeta(meta: ReportMeta, apiBase: string): ShellMeta {
  const kills = meta.fights.filter((fight) => fight.kill).length;
  const killPhrase = kills === 0 ? 'no boss kills' : plural(kills, 'boss kill', 'boss kills');
  return {
    title: `${meta.title === '' ? meta.zone : meta.title} · Forever Sixty`,
    description: `${plural(meta.fights.length, 'fight', 'fights')} in ${meta.zone}, ${killPhrase}, logged ${isoDate(meta.created_at)}.`,
    image: `${apiBase}/reports/${meta.id}/card.png`,
    canonical: `${SITE_BASE_URL}/reports/${meta.id}`,
  };
}

/** The head of GET /v1/rankings — only the two fields the unfurl needs. */
export interface RankingsHead {
  total: number;
  updated_at: string;
}

export function rankingsShellMeta(slug: string, head: RankingsHead): ShellMeta {
  const name = titleize(slug);
  return {
    title: `${name} rankings · Forever Sixty`,
    description: `${plural(head.total, 'ranked kill', 'ranked kills')} of ${name}, updated ${isoDate(head.updated_at)}.`,
    image: SITE_CARD,
    canonical: `${SITE_BASE_URL}/rankings/${slug}`,
  };
}

/** The head of GET /v1/characters/{region}/{ruleset}/{name}. */
export interface CharacterHead {
  character: { name: string; region: string; ruleset: string; class?: string };
  best?: unknown[];
  history?: unknown[];
}

export function characterShellMeta(path: CharacterPath, head: CharacterHead): ShellMeta {
  const fights = head.history?.length ?? 0;
  const encounters = head.best?.length ?? 0;
  return {
    title: `${head.character.name} · ${rulesetLabel(path.ruleset)} ${path.region.toUpperCase()} · Forever Sixty`,
    description: `${plural(fights, 'ranked fight', 'ranked fights')} across ${plural(encounters, 'encounter', 'encounters')}.`,
    image: SITE_CARD,
    canonical: `${SITE_BASE_URL}/character/${path.region}/${path.ruleset}/${path.slug}`,
  };
}

/** The head of GET /v1/guilds/{region}/{ruleset}/{name}. */
export interface GuildHead {
  guild: { name: string; region: string; ruleset: string };
  progression?: { kills: number; pull_count?: number }[];
}

export function guildShellMeta(path: CharacterPath, head: GuildHead): ShellMeta {
  const rows = head.progression ?? [];
  const kills = rows.reduce((total, row) => total + row.kills, 0);
  const pulls = rows.reduce((total, row) => total + (row.pull_count ?? 0), 0);
  return {
    title: `${head.guild.name} · ${rulesetLabel(path.ruleset)} ${path.region.toUpperCase()} · Forever Sixty`,
    description: `${plural(kills, 'boss down', 'bosses down')} over ${plural(pulls, 'pull', 'pulls')}.`,
    image: SITE_CARD,
    canonical: `${SITE_BASE_URL}/guild/${path.region}/${path.ruleset}/${path.slug}`,
  };
}
