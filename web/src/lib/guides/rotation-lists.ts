// web/src/lib/guides/rotation-lists.ts
// The guide's Rotation block: an "Opener" list, a "Priority" list and a "Cooldowns" list,
// each line a spell icon, its name and one clause. The priority comes from the same published
// lines the BiS page's "Play it" panel prints (the curated APL, `data/curated/apl/<spec>.json`,
// resolved to the level 60 ranks); the opener comes from the APL's pre-pull actions and its
// cooldown pass; the cooldowns are the lines whose spell carries a cooldown of half a minute or
// more in the client's spell constants. Nothing here is typed per guide.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import type { CollapsedLine } from '../bis/rotation-collapse';
import { repoRoot, spellIcons } from '../data/build-files';

export interface RotationItem {
  /** Absent for a step that names no spell ("Then follow the priority list"). */
  spellId?: number;
  name: string;
  rank?: number;
  /** The icon file stem, absent when the build ships none. */
  icon?: string;
  /** One clause, sentence case, no final stop. */
  clause: string;
}

export interface RotationLists {
  opener: RotationItem[];
  priority: RotationItem[];
  cooldowns: RotationItem[];
}

/** What the lists need to know about a spell the APL names by id. */
export interface SpellFacts {
  name: string;
  rank?: number;
  icon?: string;
  /** The longest of its own and its category cooldown, in milliseconds; 0 for none. */
  cooldownMs: number;
}

export type SpellLookup = (spellId: number) => SpellFacts | undefined;

const COOLDOWN_FLOOR_MS = 30_000;
const MAX_CLAUSE = 90;

type Node = Record<string, unknown>;

function isNode(value: unknown): value is Node {
  return typeof value === 'object' && value !== null && !Array.isArray(value);
}

function upperFirst(text: string): string {
  return text === '' ? text : text[0]!.toUpperCase() + text.slice(1);
}

function escapeRegExp(text: string): string {
  return text.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
}

/**
 * One clause out of a curated note: the note's first sentence or first colon-led phrase, the
 * spell's own name taken off the front (the row already says it), a leading "is" dropped, cut
 * to a word boundary when it runs long. "Ice Lance on a Fingers of Frost charge: the charge
 * makes ..." reads "On a Fingers of Frost charge".
 */
export function clauseOf(note: string, name: string): string {
  const first = note.split(/(?<=[.!?])\s|:\s|\s--\s|\s—\s|;\s/)[0] ?? '';
  // A split inside a parenthesis ("for its five-minute life (client data; Era is 60 s)") leaves
  // the aside unclosed: the clause ends before it.
  const open = first.lastIndexOf('(');
  const unclosed = open >= 0 && first.indexOf(')', open) < 0;
  let text = (unclosed ? first.slice(0, open) : first).trim().replace(/[.\s]+$/, '');
  const lead = new RegExp(`^${escapeRegExp(name)}(?:\\s+\\(rank \\d+\\))?\\s*`, 'i');
  text = text
    .replace(lead, '')
    .replace(/^[,\s]+/, '')
    .replace(/^\([^)]*\)\s*/, '')
    .replace(/^is\s+/i, '')
    .trim();
  if (text.length > MAX_CLAUSE) {
    const cut = text.slice(0, MAX_CLAUSE);
    const comma = text.slice(0, MAX_CLAUSE + 1).lastIndexOf(', ');
    if (comma > 40) {
      // A clean clause end: no ellipsis needed.
      text = cut.slice(0, comma);
    } else {
      const space = cut.lastIndexOf(' ');
      text = `${cut.slice(0, space > 40 ? space : MAX_CLAUSE).replace(/[,\s]+$/, '')}…`;
    }
  }
  return upperFirst(text);
}

/** "30 s", "2 min", "1 min 30 s". */
export function formatCooldown(ms: number): string {
  const seconds = Math.round(ms / 1000);
  if (seconds < 60) return `${seconds} s`;
  const minutes = Math.floor(seconds / 60);
  const rest = seconds % 60;
  return rest === 0 ? `${minutes} min` : `${minutes} min ${rest} s`;
}

function lineItem(line: CollapsedLine): RotationItem {
  return {
    spellId: line.spellId,
    name: line.name,
    rank: line.rank,
    icon: line.icon,
    clause: clauseOf(line.condition, line.name),
  };
}

/** Pre-pull casts in the APL, in time order: `{ spellId, secondsBefore }`. */
export function prepullCasts(apl: unknown): { spellId: number; secondsBefore: number }[] {
  const list = isNode(apl) && isNode(apl.rotation) ? apl.rotation.prepullActions : undefined;
  if (!Array.isArray(list)) return [];
  const casts: { spellId: number; secondsBefore: number }[] = [];
  for (const entry of list) {
    if (!isNode(entry) || !isNode(entry.action) || !isNode(entry.action.castSpell)) continue;
    const spell = entry.action.castSpell.spellId;
    const at =
      isNode(entry.doAtValue) && isNode(entry.doAtValue.const) ? entry.doAtValue.const.val : undefined;
    if (!isNode(spell) || typeof spell.spellId !== 'number' || typeof at !== 'string') continue;
    const seconds = Number.parseFloat(at);
    if (Number.isFinite(seconds)) casts.push({ spellId: spell.spellId, secondsBefore: Math.abs(seconds) });
  }
  return casts.sort((a, b) => b.secondsBefore - a.secondsBefore);
}

/** Whether the APL hands the on-use cooldowns (trinkets, racials, potions) to the engine's
 *  cooldown pass. */
export function usesCooldownPass(apl: unknown): boolean {
  const list = isNode(apl) && isNode(apl.rotation) ? apl.rotation.priorityList : undefined;
  return (
    Array.isArray(list) &&
    list.some((entry) => isNode(entry) && isNode(entry.action) && 'autocastOtherCooldowns' in entry.action)
  );
}

function secondsText(seconds: number): string {
  return `${Number(seconds.toFixed(1))} s`;
}

export function rotationListsFrom(
  lines: readonly CollapsedLine[],
  apl: unknown,
  spellFacts: SpellLookup,
  /** Further cooldown spells of the spec (`specCooldowns`), after the ones its priority names. */
  extraCooldowns: readonly RotationItem[] = [],
): RotationLists {
  const priority = lines.map(lineItem);

  // The opener is the APL's pre-pull casts and nothing else: with none, there is no opener to
  // print (a row that is not a spell is not an opener step).
  // A pre-pull rank often has no icon of its own; the priority line for the same spell does.
  const lineIcons = new Map(
    priority.filter((item) => item.icon !== undefined).map((item) => [item.name, item.icon]),
  );
  const opener: RotationItem[] = [];
  for (const cast of prepullCasts(apl)) {
    const facts = spellFacts(cast.spellId);
    if (facts === undefined) continue;
    opener.push({
      spellId: cast.spellId,
      name: facts.name,
      rank: facts.rank,
      icon: facts.icon ?? lineIcons.get(facts.name),
      clause: `${secondsText(cast.secondsBefore)} before the pull`,
    });
  }

  const cooldowns: RotationItem[] = [];
  const seen = new Set<string>();
  const candidates = [...opener, ...priority.filter((item) => item.spellId !== undefined)];
  for (const item of candidates) {
    const facts = spellFacts(item.spellId!);
    if (facts === undefined || facts.cooldownMs < COOLDOWN_FLOOR_MS || seen.has(item.name)) continue;
    seen.add(item.name);
    const fromLine = priority.find((line) => line.name === item.name)?.clause;
    const when = `${formatCooldown(facts.cooldownMs)} cooldown`;
    const plain = fromLine === undefined || fromLine === '' || /^on cooldown$/i.test(fromLine);
    cooldowns.push({ ...item, clause: plain ? when : `${when}. ${fromLine}` });
  }
  for (const item of extraCooldowns) {
    if (seen.has(item.name)) continue;
    seen.add(item.name);
    cooldowns.push(item);
  }
  if (usesCooldownPass(apl)) {
    cooldowns.push({
      name: 'Trinkets, racials and potions',
      clause: 'On cooldown, with your biggest damage',
    });
  }
  return { opener, priority, cooldowns };
}

// --- disk readers ---------------------------------------------------------------------------

interface SpellConstFile {
  spells: Record<string, { name: string; rank: number; cooldown_ms: number; category_cooldown_ms: number }>;
}

const constCache = new Map<string, SpellConstFile | undefined>();

function spellConst(build: string, classSlug: string): SpellConstFile | undefined {
  const key = `${build}::${classSlug}`;
  if (!constCache.has(key)) {
    const file = path.join(repoRoot(), 'data/builds', build, 'spellconst', `${classSlug}.json`);
    constCache.set(
      key,
      existsSync(file) ? (JSON.parse(readFileSync(file, 'utf8')) as SpellConstFile) : undefined,
    );
  }
  return constCache.get(key);
}

/** A spell lookup over the build's spell constants and icon table. */
export function spellLookupFor(build: string, classSlug: string): SpellLookup {
  const constants = spellConst(build, classSlug);
  const icons = spellIcons(build);
  // A pre-pull rank id often has no icon of its own; another rank of the same spell does.
  const iconByName = new Map<string, string>();
  for (const [id, row] of Object.entries(constants?.spells ?? {})) {
    const icon = icons[id];
    if (icon !== undefined && !iconByName.has(row.name)) iconByName.set(row.name, icon);
  }
  return (spellId) => {
    const row = constants?.spells[String(spellId)];
    if (row === undefined) return undefined;
    const icon = icons[String(spellId)] ?? iconByName.get(row.name);
    const iconExists =
      icon !== undefined && existsSync(path.join(repoRoot(), 'data/builds', build, 'icons', `${icon}.webp`));
    return {
      name: row.name,
      rank: row.rank > 0 ? row.rank : undefined,
      icon: iconExists ? icon : undefined,
      cooldownMs: Math.max(row.cooldown_ms, row.category_cooldown_ms),
    };
  };
}

/** The spec's curated APL, or `undefined` when none is on disk. */
export function loadApl(spec: string): unknown {
  const file = path.join(repoRoot(), 'data/curated/apl', `${spec}.json`);
  return existsSync(file) ? (JSON.parse(readFileSync(file, 'utf8')) as unknown) : undefined;
}

// --- the spec's own cooldown spells ---------------------------------------------------------

const BASELINE_COOLDOWN_FLOOR_MS = 180_000;
const TALENT_COOLDOWN_FLOOR_MS = 60_000;
const COOLDOWN_CEILING_MS = 1_800_000;
const MAX_LEVEL = 60;

interface TrainableFile {
  trainables: {
    name: string;
    active: boolean;
    cooldown_ms: number;
    ranks: { id: number; rank: number; level: number }[];
  }[];
}

interface TalentNameFile {
  trees: { talents: { name: string; icon: string; spell_id: number }[] }[];
}

/**
 * The cooldown spells a spec can use, from the client's own data: the class trainer's active
 * spells with a long cooldown (a baseline spell from three minutes up), and the talent spells
 * the spec's builds actually spend points in. A spell the APL already names is left to
 * `rotationListsFrom`. Rank is the highest learned by level 60.
 */
export function specCooldowns(
  build: string,
  classSlug: string,
  litTalents: ReadonlySet<string>,
  spellFacts: SpellLookup,
): RotationItem[] {
  const root = repoRoot();
  const trainFile = path.join(root, 'data/builds', build, 'trainables', `${classSlug}.json`);
  const talentFile = path.join(root, 'data/builds', build, 'talents', `${classSlug}.json`);
  if (!existsSync(trainFile)) return [];
  const trainables = (JSON.parse(readFileSync(trainFile, 'utf8')) as TrainableFile).trainables;
  const talentIcons = new Map<string, string>();
  const talentNames = new Set<string>();
  if (existsSync(talentFile)) {
    for (const tree of (JSON.parse(readFileSync(talentFile, 'utf8')) as TalentNameFile).trees) {
      for (const talent of tree.talents) {
        talentNames.add(talent.name);
        talentIcons.set(talent.name, talent.icon);
      }
    }
  }
  const items: (RotationItem & { cooldownMs: number; lit: boolean })[] = [];
  const seen = new Set<string>();
  for (const spell of trainables) {
    if (!spell.active || seen.has(spell.name)) continue;
    const isTalent = talentNames.has(spell.name);
    if (isTalent && !litTalents.has(spell.name)) continue;
    const learned = spell.ranks.filter((rank) => rank.level <= MAX_LEVEL);
    const top = learned.reduce<(typeof learned)[number] | undefined>(
      (best, rank) => (best === undefined || rank.rank >= best.rank ? rank : best),
      undefined,
    );
    if (top === undefined) continue;
    const facts = spellFacts(top.id);
    const cooldownMs = Math.max(spell.cooldown_ms, facts?.cooldownMs ?? 0);
    const floor = isTalent ? TALENT_COOLDOWN_FLOOR_MS : BASELINE_COOLDOWN_FLOOR_MS;
    if (cooldownMs < floor || cooldownMs > COOLDOWN_CEILING_MS) continue;
    seen.add(spell.name);
    items.push({
      spellId: top.id,
      name: spell.name,
      rank: facts?.rank,
      icon: facts?.icon ?? iconFor(build, talentIcons.get(spell.name)),
      clause: `${formatCooldown(cooldownMs)} cooldown`,
      cooldownMs,
      lit: isTalent,
    });
  }
  return items
    .sort((a, b) => Number(b.lit) - Number(a.lit) || a.cooldownMs - b.cooldownMs)
    .map(({ cooldownMs: _cooldownMs, lit: _lit, ...item }) => item);
}

function iconFor(build: string, icon: string | undefined): string | undefined {
  if (icon === undefined) return undefined;
  return existsSync(path.join(repoRoot(), 'data/builds', build, 'icons', `${icon}.webp`)) ? icon : undefined;
}
