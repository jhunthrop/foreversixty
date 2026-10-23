// web/src/lib/guides/rotation-lists.ts
// The guide's Rotation block: an "Opener" list, a "Priority" list and a "Cooldowns" list,
// each line a spell icon, its name and one clause. The priority comes from the same published
// lines the BiS page's "Play it" panel prints (the curated APL, `data/curated/apl/<spec>.json`,
// resolved to the level 60 ranks); the opener comes from the APL's pre-pull actions and its
// cooldown pass; the cooldowns are the lines whose spell carries a cooldown of half a minute or
// more in the client's spell constants. Nothing here is typed per guide.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { CollapsedLine } from '../bis/rotation-collapse';

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
  let text = first.trim().replace(/[.\s]+$/, '');
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
): RotationLists {
  const priority = lines.map(lineItem);

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
  if (usesCooldownPass(apl)) {
    opener.push({ name: 'On-use cooldowns', clause: 'Trinkets, racials and potions, as the pull starts' });
  }
  opener.push({ name: 'Priority list', clause: 'Then work down it from the top' });

  const cooldowns: RotationItem[] = [];
  const seen = new Set<string>();
  const candidates = [
    ...opener.filter((item) => item.spellId !== undefined),
    ...priority.filter((item) => item.spellId !== undefined),
  ];
  for (const item of candidates) {
    const facts = spellFacts(item.spellId!);
    if (facts === undefined || facts.cooldownMs < COOLDOWN_FLOOR_MS || seen.has(item.name)) continue;
    seen.add(item.name);
    const fromLine = priority.find((line) => line.name === item.name)?.clause;
    const when = `${formatCooldown(facts.cooldownMs)} cooldown`;
    const plain = fromLine === undefined || fromLine === '' || /^on cooldown$/i.test(fromLine);
    cooldowns.push({ ...item, clause: plain ? when : `${when}. ${fromLine}` });
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

function repoRoot(): string {
  return path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
}

interface SpellConstFile {
  spells: Record<string, { name: string; rank: number; cooldown_ms: number; category_cooldown_ms: number }>;
}

const constCache = new Map<string, SpellConstFile | undefined>();
let iconCache: { build: string; icons: Record<string, string> } | undefined;

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

function spellIcons(build: string): Record<string, string> {
  if (iconCache?.build === build) return iconCache.icons;
  const file = path.join(repoRoot(), 'data/builds', build, 'spellicons.json');
  const icons = existsSync(file) ? (JSON.parse(readFileSync(file, 'utf8')) as Record<string, string>) : {};
  iconCache = { build, icons };
  return icons;
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
