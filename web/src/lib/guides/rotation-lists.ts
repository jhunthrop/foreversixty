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
import { aplGates, gateWords, type AplGate } from '../bis/apl-gates';
import type { RacialCooldown } from './race-table';

export interface RotationItem {
  /** Absent for a step that names no spell ("Then follow the priority list"). */
  spellId?: number;
  name: string;
  rank?: number;
  /** The icon file stem, absent when the build ships none. */
  icon?: string;
  /** One clause, sentence case, no final stop. */
  clause: string;
  /** The key of the spell's game tooltip (lib/tooltips/game-tip.ts), set once the guide's
   *  tooltip data is known to carry it. */
  tipKey?: string;
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
  /** Its base cast time in milliseconds (`spelltips/<class>.json`); absent or 0 for instant. */
  castMs?: number;
}

export type SpellLookup = (spellId: number) => SpellFacts | undefined;

const COOLDOWN_FLOOR_MS = 30_000;
const MAX_CLAUSE = 90;
/** A clause up to this long is one sentence a player reads whole: it wraps, never cut
 *  ("...and Multi-Shot is about to come…" on Beast Mastery's Rapid Fire, 2026-11-04). Only a
 *  run-on past it falls back to the cut at MAX_CLAUSE. */
const MAX_WHOLE_CLAUSE = 160;

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
  if (text.length > MAX_WHOLE_CLAUSE) {
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

export interface PrepullCast {
  spellId: number;
  secondsBefore: number;
  /** What the APL gates the cast on (`aplGates`): a talent known or not, an aura up or not. */
  gates: AplGate[];
}

/** Pre-pull casts in the APL, in time order, each with the gates its condition holds. */
export function prepullCasts(apl: unknown): PrepullCast[] {
  const list = isNode(apl) && isNode(apl.rotation) ? apl.rotation.prepullActions : undefined;
  if (!Array.isArray(list)) return [];
  const casts: PrepullCast[] = [];
  for (const entry of list) {
    if (!isNode(entry) || !isNode(entry.action) || !isNode(entry.action.castSpell)) continue;
    const spell = entry.action.castSpell.spellId;
    const at =
      isNode(entry.doAtValue) && isNode(entry.doAtValue.const) ? entry.doAtValue.const.val : undefined;
    if (!isNode(spell) || typeof spell.spellId !== 'number' || typeof at !== 'string') continue;
    const seconds = Number.parseFloat(at);
    if (!Number.isFinite(seconds)) continue;
    casts.push({
      spellId: spell.spellId,
      secondsBefore: Math.abs(seconds),
      gates: aplGates(entry.action.condition),
    });
  }
  return casts.sort((a, b) => b.secondsBefore - a.secondsBefore);
}

/** A pre-pull cast gated on an aura being on the target goes out before the pull only when an
 *  earlier pre-pull cast puts that aura there (Hunter's Mark); otherwise the target carries no
 *  aura yet and the cast is one the fight makes later (Retribution's Seal of Command waits for
 *  Judgement of the Crusader), so the opener does not list it. */
function canGoOutBeforePull(cast: PrepullCast, earlier: readonly PrepullCast[]): boolean {
  return cast.gates.every(
    (gate) =>
      gate.kind !== 'active' ||
      !gate.onTarget ||
      !gate.positive ||
      earlier.some((before) => before.spellId === gate.spellId),
  );
}

/** When a pre-pull cast starts. One the APL times to land on the pull (no earlier than its
 *  base cast takes) starts its cast time before the pull, at the cast time the build's own
 *  talents give it (Arcane's Frostbolt: 3 s with no Improved Frostbolt, where the APL's 2.5 s is
 *  Frost's). Any other cast keeps the APL's own time (a summon a minute out, a Lightwell placed
 *  early). */
function prepullSeconds(
  cast: PrepullCast,
  facts: SpellFacts,
  castTimeCuts: ReadonlyMap<string, number>,
): number {
  const castMs = facts.castMs ?? 0;
  if (castMs <= 0 || cast.secondsBefore * 1000 > castMs) return cast.secondsBefore;
  return Math.max(0, castMs - (castTimeCuts.get(facts.name) ?? 0)) / 1000;
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

/** "On the pull" for a cast at the pull itself, else "1.5 s before the pull". */
function whenText(secondsBefore: number): string {
  return secondsBefore === 0 ? 'On the pull' : `${secondsText(secondsBefore)} before the pull`;
}

/** ", with Unsealed Wrath, once Judgement of the Crusader is on the target" for a gated step;
 *  empty when ungated. A gate whose spell the data cannot name is left out. */
function gateText(gates: readonly AplGate[], spellFacts: SpellLookup): string {
  return gates
    .flatMap((gate) => {
      const name = spellFacts(gate.spellId)?.name;
      return name === undefined ? [] : [`, ${gateWords(gate, name)}`];
    })
    .join('');
}

/** A priority row's clause with the talents it is gated on that its note does not already
 *  name ("With Heating Up" on an Arcane Fireball line with no note). Only talent gates: an
 *  aura or resource condition is what the note explains. */
function withTalentGates(clause: string, gates: readonly AplGate[], spellFacts: SpellLookup): string {
  const words = gates.flatMap((gate) => {
    const name = gate.kind === 'known' ? spellFacts(gate.spellId)?.name : undefined;
    return name === undefined || clause.toLowerCase().includes(name.toLowerCase())
      ? []
      : [gateWords(gate, name)];
  });
  if (words.length === 0) return clause;
  const joined = words.join(', ');
  return clause === '' ? upperFirst(joined) : `${clause}, ${joined}`;
}

export function rotationListsFrom(
  lines: readonly CollapsedLine[],
  apl: unknown,
  spellFacts: SpellLookup,
  /** Further cooldown spells of the spec (`specCooldowns`), after the ones its priority names. */
  extraCooldowns: readonly RotationItem[] = [],
  /** What the build's talents take off a spell's cooldown, by spell name (`talentCooldownCuts`). */
  cooldownCuts: ReadonlyMap<string, CooldownCut> = new Map(),
  /** The racial cooldowns of the class's races (`racialCooldownsFor`), named on the trinket line. */
  racialCooldowns: readonly RacialCooldown[] = [],
  /** The guide's role; a healer's trinket line is worded by healing. */
  role: GuideRole = 'dps',
  /** What the build's talents take off a spell's cast time, by spell name (`talentCastTimeCuts`). */
  castTimeCuts: ReadonlyMap<string, number> = new Map(),
): RotationLists {
  const priority = lines.map((line) => {
    const item = lineItem(line);
    return { ...item, clause: withTalentGates(item.clause, line.gates ?? [], spellFacts) };
  });

  // The opener is the APL's pre-pull casts and nothing else: with none, there is no opener to
  // print (a row that is not a spell is not an opener step).
  // A pre-pull rank often has no icon of its own; the priority line for the same spell does.
  const lineIcons = new Map(
    priority.filter((item) => item.icon !== undefined).map((item) => [item.name, item.icon]),
  );
  const opener: RotationItem[] = [];
  const casts = prepullCasts(apl);
  for (const [index, cast] of casts.entries()) {
    const facts = spellFacts(cast.spellId);
    if (facts === undefined || !canGoOutBeforePull(cast, casts.slice(0, index))) continue;
    opener.push({
      spellId: cast.spellId,
      name: facts.name,
      rank: facts.rank,
      icon: facts.icon ?? lineIcons.get(facts.name),
      clause: `${whenText(prepullSeconds(cast, facts, castTimeCuts))}${gateText(cast.gates, spellFacts)}`,
    });
  }

  const cooldowns: RotationItem[] = [];
  const seen = new Set<string>();
  const candidates = [...opener, ...priority.filter((item) => item.spellId !== undefined)];
  for (const item of candidates) {
    const facts = spellFacts(item.spellId!);
    if (facts === undefined) continue;
    const cooldownMs = cutCooldown(facts.cooldownMs, cooldownCuts.get(facts.name));
    if (cooldownMs < COOLDOWN_FLOOR_MS || seen.has(item.name)) continue;
    seen.add(item.name);
    const fromLine = priority.find((line) => line.name === item.name)?.clause;
    const when = `${formatCooldown(cooldownMs)} cooldown`;
    const plain = fromLine === undefined || fromLine === '' || /^on cooldown$/i.test(fromLine);
    cooldowns.push({ ...item, clause: plain ? when : `${when}. ${fromLine}` });
  }
  for (const item of extraCooldowns) {
    if (seen.has(item.name)) continue;
    seen.add(item.name);
    cooldowns.push(item);
  }
  if (usesCooldownPass(apl)) {
    cooldowns.push({ name: COOLDOWN_PASS_NAME, clause: cooldownPassClause(racialCooldowns, role) });
  }
  return { opener, priority, cooldowns };
}

export type GuideRole = 'dps' | 'healer' | 'tank';

const COOLDOWN_PASS_NAME = 'Trinkets, racials and potions';
const COOLDOWN_PASS_CLAUSE = 'On cooldown, with your biggest damage';
const HEALER_COOLDOWN_PASS_CLAUSE = 'On cooldown, when your healing matters most';

/** "On cooldown, with your biggest damage: Eureka! (Gnome), Berserking (Troll)". */
function cooldownPassClause(racials: readonly RacialCooldown[], role: GuideRole): string {
  const lead = role === 'healer' ? HEALER_COOLDOWN_PASS_CLAUSE : COOLDOWN_PASS_CLAUSE;
  if (racials.length === 0) return lead;
  const named = racials.map((entry) => `${entry.racial} (${entry.race})`).join(', ');
  return `${lead}: ${named}`;
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

/** The class's talent spells by spell id: an active talent (Demonic Sacrifice, Soul Link) has
 *  no spell-constant row, but a rotation still casts it. */
function talentSpells(build: string, classSlug: string): Map<number, { name: string; icon: string }> {
  const file = path.join(repoRoot(), 'data/builds', build, 'talents', `${classSlug}.json`);
  const byId = new Map<number, { name: string; icon: string }>();
  if (!existsSync(file)) return byId;
  for (const tree of (JSON.parse(readFileSync(file, 'utf8')) as TalentNameFile).trees) {
    for (const talent of tree.talents) byId.set(talent.spell_id, { name: talent.name, icon: talent.icon });
  }
  return byId;
}

/** A spell lookup over the build's spell constants and icon table, falling back to the class's
 *  talent spells for a talent with no constants row. */
export function spellLookupFor(build: string, classSlug: string): SpellLookup {
  const constants = spellConst(build, classSlug);
  const talents = talentSpells(build, classSlug);
  const icons = spellIcons(build);
  // A pre-pull rank id often has no icon of its own; another rank of the same spell does.
  const iconByName = new Map<string, string>();
  for (const [id, row] of Object.entries(constants?.spells ?? {})) {
    const icon = icons[id];
    if (icon !== undefined && !iconByName.has(row.name)) iconByName.set(row.name, icon);
  }
  const castTimes = classSpellCastTimes(build, classSlug);
  return (spellId) => {
    const row = constants?.spells[String(spellId)];
    if (row === undefined) {
      const talent = talents.get(spellId);
      if (talent === undefined) return undefined;
      return {
        name: talent.name,
        icon: iconFor(build, icons[String(spellId)] ?? talent.icon),
        cooldownMs: 0,
      };
    }
    const icon = icons[String(spellId)] ?? iconByName.get(row.name);
    const iconExists =
      icon !== undefined && existsSync(path.join(repoRoot(), 'data/builds', build, 'icons', `${icon}.webp`));
    return {
      name: row.name,
      rank: row.rank > 0 ? row.rank : undefined,
      icon: iconExists ? icon : undefined,
      cooldownMs: Math.max(row.cooldown_ms, row.category_cooldown_ms),
      castMs: castTimes.get(spellId) ?? 0,
    };
  };
}

/** The class's spells' base cast times, by spell id, from the build's spell tooltips. */
function classSpellCastTimes(build: string, classSlug: string): Map<number, number> {
  const file = path.join(repoRoot(), 'data/builds', build, 'spelltips', `${classSlug}.json`);
  if (!existsSync(file)) return new Map();
  const spells = (
    JSON.parse(readFileSync(file, 'utf8')) as { spells: Record<string, { cast_time_ms?: number }> }
  ).spells;
  return new Map(Object.entries(spells).map(([id, tip]) => [Number(id), tip.cast_time_ms ?? 0]));
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
/** A spell with a cast this long (Astral Recall's 10 s, a summoning ritual) is never a fight cooldown. */
const MAX_COOLDOWN_CAST_MS = 3_000;
const MAX_LEVEL = 60;

interface TrainableFile {
  trainables: {
    name: string;
    active: boolean;
    cooldown_ms: number;
    cast_time_ms?: number;
    ranks: { id: number; rank: number; level: number }[];
  }[];
}

interface TalentNameFile {
  trees: {
    talents: { name: string; icon: string; spell_id: number; ranks?: { description?: string }[] }[];
  }[];
}

/** The talents a guide's builds take, by name, with the highest rank either build takes. */
export type LitTalents = ReadonlyMap<string, number>;

/** A talent's cut to one spell's cooldown: a flat amount ("by 2 min") and a share ("by 90%"). */
export interface CooldownCut {
  ms: number;
  percent: number;
}

/** The client's wordings of a talent's cooldown cut, each naming the spells (group 1) and the
 *  amount (2, 3): "Reduces the cooldown of your Blessing of Protection by 2 min", "...on your
 *  Rapid Fire ability by 1 min" (Rapid Killing), "...of your Trap and Deterrence abilities by
 *  20%" (Survivalist's Discipline), "...of your Divine Shield, Divine Protection, and Templar's
 *  Bulwark spells by 30 sec" (Sacred Duty), and "Reduces threat caused by Tranquility by 50% and
 *  its cooldown by 30%" (Improved Tranquility). */
const COOLDOWN_CUTS: readonly RegExp[] = [
  /Reduces the cooldown (?:of|on) your ([A-Z][A-Za-z' ,]+?)(?: spells?| abilit(?:y|ies))? by (\d+(?:\.\d+)?) ?(sec|min|%)/g,
  /Reduces [a-z ]+ (?:caused by|of) ([A-Z][A-Za-z' ]+?) by \d+(?:\.\d+)?% and its cooldown by (\d+(?:\.\d+)?) ?(sec|min|%)/g,
];
/** "Divine Shield, Divine Protection, and Templar's Bulwark" -> its three names. */
const NAME_LIST = /,\s*(?:and\s+)?|\s+and\s+/;

/** The talents any of `trees` spends points in, at the highest rank any of them takes. */
export function litTalentsOf(
  trees: readonly { talents: readonly { name: string; rank: number }[] }[],
): LitTalents {
  const lit = new Map<string, number>();
  for (const talent of trees.flatMap((tree) => tree.talents)) {
    if (talent.rank > 0) lit.set(talent.name, Math.max(lit.get(talent.name) ?? 0, talent.rank));
  }
  return lit;
}

/** A cooldown less a talent's cut, never below zero. */
export function cutCooldown(ms: number, cut: CooldownCut | undefined): number {
  if (cut === undefined) return ms;
  return Math.max(0, (ms - cut.ms) * (1 - cut.percent / 100));
}

/**
 * What the lit talents take off each spell's cooldown, read from the client's own talent rank
 * text at the rank the build takes ("Reduces the cooldown of your Blessing of Protection by 2
 * min": Guardian's Favor 2/2; "by 90%": Soul Fire), by spell name. The client's spell constants
 * carry the base cooldown only.
 */
export function talentCooldownCuts(
  build: string,
  classSlug: string,
  lit: LitTalents,
): Map<string, CooldownCut> {
  const cuts = new Map<string, CooldownCut>();
  for (const text of talentRankTexts(build, classSlug, lit)) {
    for (const match of COOLDOWN_CUTS.flatMap((pattern) => [...text.matchAll(pattern)])) {
      const amount = Number(match[2]);
      const add: CooldownCut =
        match[3] === '%'
          ? { ms: 0, percent: amount }
          : { ms: amount * (match[3] === 'min' ? 60_000 : 1000), percent: 0 };
      for (const name of match[1]!.split(NAME_LIST)) {
        const held = cuts.get(name.trim()) ?? { ms: 0, percent: 0 };
        cuts.set(name.trim(), { ms: held.ms + add.ms, percent: held.percent + add.percent });
      }
    }
  }
  return cuts;
}

/** The client's wording of a talent's lasting cast-time cut: "Reduces the casting time of your
 *  Frostbolt spell by 0.5 sec" (Improved Frostbolt), "...of your Shadow Bolt, Immolate, and
 *  Incinerate spells by 0.5 sec" (Bane), "Reduces the cast time of Starfire by 0.5 sec". A
 *  cut that a proc or an active grants ("reduce the cast time of your next ...", "Gain
 *  mastery ..., reducing the cast time") is worded otherwise and not read. */
const CAST_TIME_CUT =
  /[Rr]educes the cast(?:ing)? time of (?:your )?([A-Z][A-Za-z' ,]+?)(?: spells?| abilit(?:y|ies))? by (\d+(?:\.\d+)?) sec/g;

/** What the lit talents take off each spell's cast time, in milliseconds, by spell name. */
export function talentCastTimeCuts(build: string, classSlug: string, lit: LitTalents): Map<string, number> {
  const cuts = new Map<string, number>();
  for (const text of talentRankTexts(build, classSlug, lit)) {
    for (const match of text.matchAll(CAST_TIME_CUT)) {
      for (const name of match[1]!.split(NAME_LIST)) {
        cuts.set(name.trim(), (cuts.get(name.trim()) ?? 0) + Number(match[2]) * 1000);
      }
    }
  }
  return cuts;
}

/** Each lit talent's own description at the rank the build takes. */
function talentRankTexts(build: string, classSlug: string, lit: LitTalents): string[] {
  const file = path.join(repoRoot(), 'data/builds', build, 'talents', `${classSlug}.json`);
  if (!existsSync(file)) return [];
  return (JSON.parse(readFileSync(file, 'utf8')) as TalentNameFile).trees.flatMap((tree) =>
    tree.talents.flatMap((talent) => {
      const rank = lit.get(talent.name) ?? 0;
      const text = rank > 0 ? talent.ranks?.[rank - 1]?.description : undefined;
      return text === undefined ? [] : [text];
    }),
  );
}

interface GuideCooldownsFile {
  specs: Record<string, string[]>;
}

/** The baseline trainer cooldowns a spec's guide lists (`data/curated/guide-cooldowns.json`). */
export function guideBaselineCooldowns(spec: string): ReadonlySet<string> {
  const file = path.join(repoRoot(), 'data/curated/guide-cooldowns.json');
  if (!existsSync(file)) return new Set();
  return new Set((JSON.parse(readFileSync(file, 'utf8')) as GuideCooldownsFile).specs[spec] ?? []);
}

interface ClassSpellLimitsFile {
  raceLocked: Record<string, Record<string, string[]>>;
  notLearnable: Record<string, Record<string, string>>;
}

/** The class trainer spells a guide's cooldown list leaves out: the race-locked ones (the
 *  Races section covers those) and the ones no player of the class can learn
 *  (`data/curated/class-spell-limits.json`). */
export function limitedClassSpells(classSlug: string): ReadonlySet<string> {
  const file = path.join(repoRoot(), 'data/curated/class-spell-limits.json');
  if (!existsSync(file)) return new Set();
  const limits = JSON.parse(readFileSync(file, 'utf8')) as ClassSpellLimitsFile;
  return new Set([
    ...Object.keys(limits.raceLocked[classSlug] ?? {}),
    ...Object.keys(limits.notLearnable[classSlug] ?? {}),
  ]);
}

/**
 * The cooldown spells a spec's guide lists beyond its rotation, from the client's own data:
 * the talent spells the spec's builds actually spend points in, and the baseline trainer
 * spells with a long cooldown (three minutes up) that the spec's curated list names
 * (`guideBaselineCooldowns`), so a stun, a taunt or a spell another spec presses stays off.
 * Each cooldown is shown less what the build's talents take off it (`talentCooldownCuts`).
 * A spell the APL already names is left to `rotationListsFrom`. Rank is the highest learned
 * by level 60.
 */
export function specCooldowns(
  build: string,
  classSlug: string,
  spec: string,
  litTalents: LitTalents,
  spellFacts: SpellLookup,
): RotationItem[] {
  const baseline = guideBaselineCooldowns(spec);
  const cuts = talentCooldownCuts(build, classSlug, litTalents);
  const root = repoRoot();
  const trainFile = path.join(root, 'data/builds', build, 'trainables', `${classSlug}.json`);
  const talentFile = path.join(root, 'data/builds', build, 'talents', `${classSlug}.json`);
  if (!existsSync(trainFile)) return [];
  const trainables = (JSON.parse(readFileSync(trainFile, 'utf8')) as TrainableFile).trainables;
  const talentIcons = new Map<string, string>();
  const talentNames = new Set<string>();
  const talentSpells: { name: string; icon: string; spellId: number }[] = [];
  if (existsSync(talentFile)) {
    for (const tree of (JSON.parse(readFileSync(talentFile, 'utf8')) as TalentNameFile).trees) {
      for (const talent of tree.talents) {
        talentNames.add(talent.name);
        talentIcons.set(talent.name, talent.icon);
        talentSpells.push({ name: talent.name, icon: talent.icon, spellId: talent.spell_id });
      }
    }
  }
  const items: (RotationItem & { cooldownMs: number; lit: boolean })[] = [];
  const seen = new Set<string>(limitedClassSpells(classSlug));
  for (const spell of trainables) {
    if (!spell.active || seen.has(spell.name) || (spell.cast_time_ms ?? 0) > MAX_COOLDOWN_CAST_MS) continue;
    const isTalent = talentNames.has(spell.name);
    if (isTalent ? !litTalents.has(spell.name) : !baseline.has(spell.name)) continue;
    const learned = spell.ranks.filter((rank) => rank.level <= MAX_LEVEL);
    const top = learned.reduce<(typeof learned)[number] | undefined>(
      (best, rank) => (best === undefined || rank.rank >= best.rank ? rank : best),
      undefined,
    );
    if (top === undefined) continue;
    const facts = spellFacts(top.id);
    const cooldownMs = cutCooldown(Math.max(spell.cooldown_ms, facts?.cooldownMs ?? 0), cuts.get(spell.name));
    const floor = isTalent ? TALENT_COOLDOWN_FLOOR_MS : BASELINE_COOLDOWN_FLOOR_MS;
    if (cooldownMs < floor || cooldownMs > COOLDOWN_CEILING_MS) continue;
    seen.add(spell.name);
    items.push({
      spellId: top.id,
      name: spell.name,
      rank: facts?.rank,
      // A rank the sim's spell constants do not carry (Dash rank 3) still has the client's icon.
      icon:
        facts?.icon ??
        iconFor(build, spellIcons(build)[String(top.id)]) ??
        iconFor(build, talentIcons.get(spell.name)),
      clause: `${formatCooldown(cooldownMs)} cooldown`,
      cooldownMs,
      lit: isTalent,
    });
  }
  // A talent spell the trainer data does not list (Adrenaline Rush, Cold Blood, Preparation)
  // is read from the talent file's own spell.
  for (const talent of talentSpells) {
    if (seen.has(talent.name) || !litTalents.has(talent.name)) continue;
    const facts = spellFacts(talent.spellId);
    const cooldownMs = cutCooldown(facts?.cooldownMs ?? 0, cuts.get(talent.name));
    if (cooldownMs < TALENT_COOLDOWN_FLOOR_MS || cooldownMs > COOLDOWN_CEILING_MS) continue;
    seen.add(talent.name);
    items.push({
      spellId: talent.spellId,
      name: talent.name,
      rank: facts?.rank,
      icon: facts?.icon ?? iconFor(build, talent.icon),
      clause: `${formatCooldown(cooldownMs)} cooldown`,
      cooldownMs,
      lit: true,
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
