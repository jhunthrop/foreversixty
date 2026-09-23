// web/src/lib/report/og-meta.ts
// What the Worker writes into a shell's head. Pure, so the copy is reviewable and tested
// without a Workers runtime; src/worker.ts only wires the values into HTMLRewriter.
//
// Voice: reference, not pitch. A description states the numbers and the date and stops.
//
// The sim shell's kind words and its empty-result headline come from sim/copy.ts, not from
// a fourth copy here: this module already imports sim/bulk-types, sim/combos and
// sim/weights, and sim/combos itself imports bulkCopy, so there is no dependency left to
// avoid (final whole-branch review, Minor 2).
import { rulesetLabel, type CharacterPath } from '../characters';
import { requestKind, type BulkResult, type WeightsResult } from '../sim/bulk-types';
import { headlineFor } from '../sim/combos';
import { bulkCopy, KIND_TITLES } from '../sim/copy';
import { encounterLabel } from '../sim/encounter';
import { confidenceBand, formatMargin } from '../sim/estimate';
import { specLabel } from '../sim/spec-label';
import type { SimResult } from '../sim/types';
import { statLabel } from '../sim/weights';
import type { ReportMeta } from './types';

export const SITE_BASE_URL = 'https://foreversixty.gg';

export type ShellKind = 'report' | 'rankings' | 'character' | 'guild' | 'sim';

export interface ShellMeta {
  /** The whole <title>, including the site suffix. */
  title: string;
  /** The page's own heading, without the site suffix: what the shell paints as its h1
   *  before the island mounts (lib/report/skeleton.ts). Omitted where the shell has no
   *  heading of its own. */
  heading?: string;
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
export const SITE_CARD = `${SITE_BASE_URL}/og/index.png`;

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
  const heading = meta.title === '' ? meta.zone : meta.title;
  return {
    heading,
    title: `${heading} · Forever Sixty`,
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
  // "N bosses down" counts distinct bosses killed at least once, not total kill events: a
  // boss farmed repeatedly in one progression row is still one boss down.
  const kills = rows.filter((row) => row.kills > 0).length;
  const pulls = rows.reduce((total, row) => total + (row.pull_count ?? 0), 0);
  return {
    title: `${head.guild.name} · ${rulesetLabel(path.ruleset)} ${path.region.toUpperCase()} · Forever Sixty`,
    description: `${plural(kills, 'boss down', 'bosses down')} over ${plural(pulls, 'pull', 'pulls')}.`,
    image: SITE_CARD,
    canonical: `${SITE_BASE_URL}/guild/${path.region}/${path.ruleset}/${path.slug}`,
  };
}

export function guildClaimShellMeta(path: CharacterPath, guildName: string): ShellMeta {
  return {
    title: `Claim ${guildName} · Forever Sixty`,
    description: `Claim officer control of ${guildName}.`,
    image: SITE_CARD,
    canonical: `${SITE_BASE_URL}/guild/${path.region}/${path.ruleset}/${path.slug}/claim`,
  };
}

export function guildSettingsShellMeta(path: CharacterPath, guildName: string): ShellMeta {
  return {
    title: `${guildName} settings · Forever Sixty`,
    description: `Officer settings for ${guildName}.`,
    image: SITE_CARD,
    canonical: `${SITE_BASE_URL}/guild/${path.region}/${path.ruleset}/${path.slug}/settings`,
  };
}

/**
 * Never guild-scoped (plan ruling 6): there is no unauthenticated endpoint to look up a
 * guild by invite token — accepting the invite is the only route that reads one, and it is
 * destructive and session-gated, so the Worker's shell (which must never carry a session)
 * cannot call it.
 */
export function guildInviteShellMeta(token: string): ShellMeta {
  return {
    title: 'Join a guild · Forever Sixty',
    description: 'Join a Forever Sixty guild by invite link.',
    image: SITE_CARD,
    canonical: `${SITE_BASE_URL}/guild/invite/${token}`,
  };
}

/**
 * `simShellMeta`'s own composition, before a member's own name (`result.title`) overrides
 * it -- split out so that override is one place, below, rather than three (a branch per
 * kind). Every existing kind's title/description is exactly what it always was; nothing in
 * this function changed for the title defect fix.
 *
 * The description is the whole run in one sentence: the engine it came from, the figure and
 * its 95% band, how many iterations bought that band, and the settings. Someone deciding
 * whether to open the link has every number that would change their mind. A bulk kind
 * (gear/talents/drops) says the combination count and the leader's own headline instead --
 * there is no single DPS figure worth leading with when the page is a ranked table -- and
 * weights says the top few stats rather than a figure that was never the point of the run.
 */
function fallbackSimShellMeta(result: SimResult): ShellMeta {
  const kind = requestKind(result.request);
  const canonical = `${SITE_BASE_URL}/sim/${result.sim_id ?? ''}`;
  const spec = specLabel(result.request.spec);

  if (kind === 'weights') {
    const weights = (result as WeightsResult).weights ?? [];
    const top = weights
      .slice(0, 3)
      .map((row) => `${statLabel(row.stat)} ${row.weight.toFixed(2)}`)
      .join(' · ');
    return {
      title: `${KIND_TITLES.weights} · ${spec} · Forever Sixty`,
      description: `Simulated on engine ${result.engine_version}: ${top}.`,
      image: SITE_CARD,
      canonical,
    };
  }

  if (kind !== 'run') {
    const bulk = result as BulkResult;
    const combos = bulk.combos ?? [];
    // Contract 10.1 A6 fills Substitution.Name for items too, so the unfurl can name the
    // winning change without the per-class item file this function must never fetch. An
    // empty result reads "no combinations", matching the API's own headline rule (contract
    // 10.6) rather than inventing a second wording for the same state.
    const headline = combos.length === 0 ? bulkCopy.noCombinations : headlineFor(bulk);
    return {
      title: `${KIND_TITLES[kind]} · ${spec} · Forever Sixty`,
      description:
        `Simulated on engine ${result.engine_version}: ${combos.length} combinations, ` + `${headline}.`,
      image: SITE_CARD,
      canonical,
    };
  }

  const dps = Math.round(result.dps.mean).toLocaleString('en-US');
  const band = formatMargin(confidenceBand(result.dps));
  const iterations = result.request.iterations.toLocaleString('en-US');
  // Whether the player carried a raid buff, read from the stored fact rather than guessed:
  // a SimResult carries no buff preset, so this is true exactly when the summary records at
  // least one BUFF-type aura track.
  const buffed = result.summary.auras.some((track) => track.type === 'BUFF');
  const settings = encounterLabel(result.request.encounter, buffed);
  return {
    title: `${spec}, ${dps} DPS · Forever Sixty`,
    description: `Simulated on engine ${result.engine_version}: ${dps} DPS ± ${band} over ${iterations} iterations, ${settings}.`,
    image: SITE_CARD,
    canonical,
  };
}

/**
 * A saved sim's unfurl. There is no rendered card for a sim at launch -- the API draws one
 * per report and per build, and a third renderer is its work, not this lane's -- so the
 * site's own card stands in, the way the rankings and character shells already do.
 *
 * Takes no `apiBase`, unlike reportShellMeta: a sim's image is always the site's own card,
 * never a per-sim render, so there is nothing here for an API origin to build.
 *
 * Defect fix: `result.title` -- "Name this sim" (SaveSimForm.svelte), carried through by
 * `GET /v1/sims/{id}`'s own `GetOutput` (api/internal/sims/handler.go) -- used to reach
 * neither this function's input (SimResult had no such field) nor its output, so a member's
 * own name never appeared in `<title>`, `og:title` or `og:description`: every shared link
 * read the same composed spec-and-DPS sentence regardless of what was typed. It leads both
 * now, when given -- `title` verbatim (src/worker.ts's `rewriteHead` still routes it through
 * HTMLRewriter's own `setInnerContent`/`setAttribute`, which escape by default; nothing here
 * switches either into raw-HTML mode), `description` with it as the opening sentence, ahead
 * of the same numbers `fallbackSimShellMeta` already composed.
 */
export function simShellMeta(result: SimResult): ShellMeta {
  const meta = fallbackSimShellMeta(result);
  if (result.title === undefined || result.title === '') return meta;
  return {
    ...meta,
    title: `${result.title} · Forever Sixty`,
    description: `${result.title}. ${meta.description}`,
  };
}
