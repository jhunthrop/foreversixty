// web/src/lib/guides/key-talents.ts
// The "Key talents" list under a guide's builds: each entry is the talents' icons and names
// with a one-line why. The why is the guide's own bullet (`- **Name** (n points) — why`) in
// its Talents section (`- **Name** (n points): why` is the form a guide writes now; the older
// em dash separator still parses); the icons, the ranks and the legal names come from the
// builds' talent data. The bullets speak to both builds, so a talent the level 60 and raid
// builds take at different ranks carries both, each named (`buildRanks`); otherwise its rank
// is the raid build's, or the level 60 build's for a guide with no raid build of its own. A spec whose Talents section has no such bullets gets its build's
// own strongest talents with each talent's description as the why, so the list is never empty
// for a spec with a build.
import type { DecodedTalentView, DecodedTreeView } from './build-tree';
import { upperFirst } from '../text';

export interface KeyTalentRef {
  /** The talent's id, absent when the build has no talent of that name. */
  id?: number;
  name: string;
  /** The icon file stem; absent when the build has no talent of that name. */
  icon?: string;
  /** Points the build spends in it: the raid build's when it takes the talent, else the level
   *  60 build's; absent when neither build has the name. */
  rank?: number;
  maxRank?: number;
  /** Set only when the guide has a raid build and the two builds take the talent at different
   *  ranks: each build's own, so a chip never shows one build's rank beside the other's text. */
  buildRanks?: { leveling: number; raid: number };
}

export interface KeyTalentEntry {
  talents: KeyTalentRef[];
  /** One line, sentence case, no final stop. */
  why: string;
}

export interface ParsedBullet {
  names: string[];
  why: string;
  /** The raw source lines the bullet spans, so a caller can cut them from the prose. */
  raw: string;
}

const DASH_SPLIT =
  /^(.*?\*\*[^*]+\*\*(?:\s*\([^)]*\))?(?:,?\s*(?:and\s+)?\*\*[^*]+\*\*(?:\s*\([^)]*\))?)*)(?:\s+—\s+|:\s+)(.+)$/;

/** A key talent with no why: `- **Name** (n points)`. Its tooltip says what it does, so a blurb
 *  that only restated it is left out (h-audit 2026-11-09, ai-slop tell 23). */
const NAME_ONLY = /^\*\*([^*]+)\*\*(?:\s*\([^)]*\))?\.?$/;

/** Every `- **Name** …: why` (or `— why`) bullet in a Talents section, and every name-only
 *  `- **Name** (n points)` bullet (an empty why), in order. */
export function parseTalentBullets(markdown: string): ParsedBullet[] {
  const bullets: ParsedBullet[] = [];
  for (const line of markdown.split(/\r?\n/)) {
    if (!line.startsWith('- **')) continue;
    const bare = NAME_ONLY.exec(line.slice(2).trim());
    if (bare !== null) {
      bullets.push({ names: [bare[1]!.trim()], why: '', raw: line });
      continue;
    }
    const match = DASH_SPLIT.exec(line.slice(2));
    if (match === null) continue;
    const names = [...(match[1] ?? '').matchAll(/\*\*([^*]+)\*\*/g)].map((m) => m[1]!.trim());
    const why = upperFirst((match[2] ?? '').trim().replace(/\.$/, ''));
    if (names.length === 0 || why === '') continue;
    bullets.push({ names, why, raw: line });
  }
  return bullets;
}

/** The section with the parsed bullets cut out, and the blank lines they leave collapsed. */
export function withoutTalentBullets(markdown: string): string {
  const parsed = new Set(parseTalentBullets(markdown).map((bullet) => bullet.raw));
  if (parsed.size === 0) return markdown;
  const kept = markdown.split(/\r?\n/).filter((line) => !parsed.has(line));
  return kept.join('\n').replace(/\n{3,}/g, '\n\n');
}

function normalise(name: string): string {
  return name
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, ' ')
    .trim();
}

function allTalents(trees: readonly DecodedTreeView[]): DecodedTalentView[] {
  return trees.flatMap((tree) => tree.talents);
}

function findTalent(name: string, talents: readonly DecodedTalentView[]): DecodedTalentView | undefined {
  const wanted = normalise(name);
  return talents.find((talent) => normalise(talent.name) === wanted);
}

function refFor(
  name: string,
  talents: readonly DecodedTalentView[],
  raidTalents?: readonly DecodedTalentView[],
): KeyTalentRef {
  const leveling = findTalent(name, talents);
  const raid = raidTalents === undefined ? undefined : findTalent(name, raidTalents);
  const hit = raid !== undefined && raid.rank > 0 ? raid : (leveling ?? raid);
  if (hit === undefined) return { name };
  const ref: KeyTalentRef = {
    id: hit.id,
    name: hit.name,
    icon: hit.icon,
    rank: hit.rank,
    maxRank: hit.max_rank,
  };
  const levelingRank = leveling?.rank ?? 0;
  const raidRank = raid?.rank ?? 0;
  return raidTalents !== undefined && levelingRank !== raidRank
    ? { ...ref, buildRanks: { leveling: levelingRank, raid: raidRank } }
    : ref;
}

const FALLBACK_COUNT = 6;

function describeRank(talent: DecodedTalentView): string {
  const text = talent.ranks[Math.max(0, talent.rank - 1)]?.description ?? '';
  return text.replace(/[.\s]+$/, '');
}

/** The primary tree's deepest talents the build takes, with the description they carry at the
 *  build's rank. */
function fallbackEntries(trees: readonly DecodedTreeView[]): KeyTalentEntry[] {
  const primary = trees.reduce<DecodedTreeView | undefined>(
    (best, tree) => (best === undefined || tree.points > best.points ? tree : best),
    undefined,
  );
  if (primary === undefined) return [];
  return primary.talents
    .filter((talent) => talent.rank > 0 && describeRank(talent) !== '')
    .sort((a, b) => b.tier - a.tier || b.rank - a.rank)
    .slice(0, FALLBACK_COUNT)
    .map((talent) => ({
      talents: [
        { id: talent.id, name: talent.name, icon: talent.icon, rank: talent.rank, maxRank: talent.max_rank },
      ],
      why: describeRank(talent),
    }));
}

/** `raidTrees`: the guide's own raid build, when it differs from the level 60 build. */
export function keyTalentsFor(
  talentsMarkdown: string,
  levelingTrees: readonly DecodedTreeView[] | undefined,
  raidTrees?: readonly DecodedTreeView[],
): KeyTalentEntry[] {
  const talents = levelingTrees === undefined ? [] : allTalents(levelingTrees);
  const raidTalents = raidTrees === undefined ? undefined : allTalents(raidTrees);
  const bullets = parseTalentBullets(talentsMarkdown);
  if (bullets.length > 0) {
    return bullets.map((bullet) => ({
      talents: bullet.names.map((name) => refFor(name, talents, raidTalents)),
      why: bullet.why,
    }));
  }
  const described = raidTrees ?? levelingTrees;
  return described === undefined ? [] : fallbackEntries(described);
}
