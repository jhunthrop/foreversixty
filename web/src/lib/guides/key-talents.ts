// web/src/lib/guides/key-talents.ts
// The "Key talents" list under a guide's builds: each entry is the talents' icons and names
// with a one-line why. The why is the guide's own bullet (`- **Name** (n points) — why`) in
// its Talents section; the icons, the ranks and the legal names come from the leveling
// build's talent data. A spec whose Talents section has no such bullets gets its build's
// own strongest talents with each talent's description as the why, so the list is never empty
// for a spec with a build.
import type { DecodedTalentView, DecodedTreeView } from './build-tree';

export interface KeyTalentRef {
  name: string;
  /** The icon file stem; absent when the build has no talent of that name. */
  icon?: string;
  /** Points the leveling build spends in it, absent when the name is not in the build. */
  rank?: number;
  maxRank?: number;
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
  /^(.*?\*\*[^*]+\*\*(?:\s*\([^)]*\))?(?:,?\s*(?:and\s+)?\*\*[^*]+\*\*(?:\s*\([^)]*\))?)*)\s+—\s+(.+)$/;

function upperFirst(text: string): string {
  return text === '' ? text : text[0]!.toUpperCase() + text.slice(1);
}

/** Every `- **Name** … — why` bullet in a Talents section, in order. */
export function parseTalentBullets(markdown: string): ParsedBullet[] {
  const bullets: ParsedBullet[] = [];
  for (const line of markdown.split(/\r?\n/)) {
    if (!line.startsWith('- **')) continue;
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

function refFor(name: string, talents: readonly DecodedTalentView[]): KeyTalentRef {
  const wanted = normalise(name);
  const hit = talents.find((talent) => normalise(talent.name) === wanted);
  return hit === undefined
    ? { name }
    : { name: hit.name, icon: hit.icon, rank: hit.rank, maxRank: hit.max_rank };
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
      talents: [{ name: talent.name, icon: talent.icon, rank: talent.rank, maxRank: talent.max_rank }],
      why: describeRank(talent),
    }));
}

export function keyTalentsFor(
  talentsMarkdown: string,
  levelingTrees: readonly DecodedTreeView[] | undefined,
): KeyTalentEntry[] {
  const talents = levelingTrees === undefined ? [] : allTalents(levelingTrees);
  const bullets = parseTalentBullets(talentsMarkdown);
  if (bullets.length > 0) {
    return bullets.map((bullet) => ({
      talents: bullet.names.map((name) => refFor(name, talents)),
      why: bullet.why,
    }));
  }
  return levelingTrees === undefined ? [] : fallbackEntries(levelingTrees);
}
