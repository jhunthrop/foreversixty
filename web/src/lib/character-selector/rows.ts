// web/src/lib/character-selector/rows.ts
// The selector's view model: what `/v1/me` and the stored pointer say, turned into the rows
// the list draws, the closed slot's character and the session state (spec sections 3 and 9).
// Pure functions of their inputs, so the ordering, the stale rule and the pointer cases are
// all tested without a browser.
import type { Me, MeCharacter } from '../account/api';
import { mainCharacter } from '../account/main-character';
import { rulesetLabel } from '../characters';
import type { CurrentCharacter } from '../current-character';
import { relativeTime } from '../dates';
import type { Faction } from '../faction-mark';
import { orderCharacters } from './order';
import { selectorCopy } from './copy';
import { isFailed, isStale } from './stale';

export type RowStatus = 'ok' | 'stale' | 'failed';
export type RowKind = 'account' | 'pointer';

export interface SelectorRow {
  /** The account key for an account row; the pointer's `ref` for a row built from the pointer. */
  key: string;
  kind: RowKind;
  name: string;
  classSlug: string;
  className: string;
  spec?: string;
  level?: number;
  realm?: string;
  /** Omitted for the normal ruleset. */
  ruleset?: string;
  faction?: Faction;
  /** "Addon", "Battle.net" or "Pasted export"; empty when nothing says where it came from. */
  sourceWord: string;
  /** "12 minutes ago"; empty when no sync time is known. */
  age: string;
  status: RowStatus;
  current: boolean;
  /** Set on account rows; the character the pointer is written from. */
  character?: MeCharacter;
}

export type SessionState = 'loading' | 'signed-in' | 'signed-out' | 'expired';

export interface SelectorModel {
  session: SessionState;
  rows: SelectorRow[];
  current: SelectorRow | null;
  /** True when some account character was last synced by the addon ("Get the addon" hides). */
  hasAddonCharacter: boolean;
}

export interface ModelInput {
  /** `undefined` while `/v1/me` is loading; `null` once it answered "signed out". */
  me: Me | null | undefined;
  pointer: CurrentCharacter | null;
  chosenKeys: readonly string[];
  /** The clock, in milliseconds since the epoch (a number, so reactive state can hold it). */
  nowMs: number;
}

function capitalise(slug: string): string {
  return slug.length === 0 ? slug : slug.charAt(0).toUpperCase() + slug.slice(1);
}

function slugOf(className: string | undefined): string {
  return (className ?? '').toLowerCase().replace(/\s+/g, '-');
}

function sourceWordOf(character: MeCharacter): string {
  if (character.build !== undefined) {
    return character.build.source === 'addon' ? selectorCopy.sourceAddon : selectorCopy.sourceBnet;
  }
  if (character.source === 'bnet') return selectorCopy.sourceBnet;
  if (character.source === 'export') return selectorCopy.sourceAddon;
  return '';
}

function statusOf(character: MeCharacter, now: Date): RowStatus {
  if (isFailed(character.build)) return 'failed';
  return isStale(character.build, now) ? 'stale' : 'ok';
}

function rowFromCharacter(character: MeCharacter, current: boolean, now: Date): SelectorRow {
  const captured = character.build?.captured_at;
  const classSlug = slugOf(character.class);
  return {
    key: character.key,
    kind: 'account',
    name: character.name,
    classSlug,
    className: capitalise(classSlug),
    spec: character.spec,
    level: character.level,
    realm: character.realm,
    ruleset: character.ruleset === 'normal' ? undefined : rulesetLabel(character.ruleset),
    faction: character.faction,
    sourceWord: sourceWordOf(character),
    age: captured === undefined ? '' : relativeTime(new Date(captured), now),
    status: statusOf(character, now),
    current,
    character,
  };
}

/** "Name · Fury Warrior" -> its two parts. A bare "Warrior" (a paste with no name) is the class alone. */
function splitPointerLabel(pointer: CurrentCharacter): { name: string; specClass: string } {
  const [head, ...rest] = pointer.label.split(' · ');
  return rest.length === 0 ? { name: head, specClass: '' } : { name: head, specClass: rest.join(' · ') };
}

function specOfLabel(specClass: string, className: string): string | undefined {
  const suffix = ` ${className}`;
  if (specClass.endsWith(suffix)) return specClass.slice(0, -suffix.length);
  return undefined;
}

/** The row for a pointer that names no account character: a pasted export, or an expired session. */
export function rowFromPointer(pointer: CurrentCharacter, sourceWord: string, now: Date): SelectorRow {
  const { name, specClass } = splitPointerLabel(pointer);
  const className = capitalise(pointer.classSlug);
  return {
    key: pointer.ref,
    kind: 'pointer',
    name,
    classSlug: pointer.classSlug,
    className,
    spec: specOfLabel(specClass, className),
    sourceWord,
    age: sourceWord === '' ? '' : relativeTime(new Date(pointer.savedAt), now),
    status: 'ok',
    current: true,
  };
}

type PastedPointer = CurrentCharacter & { source: 'code' | 'addon' };

function isPastedPointer(pointer: CurrentCharacter | null): pointer is PastedPointer {
  return pointer !== null && (pointer.source === 'code' || pointer.source === 'addon');
}

function signedInModel(me: Me, input: ModelInput, now: Date): SelectorModel {
  const { pointer } = input;
  const hasAddonCharacter = me.characters.some((character) => character.build?.source === 'addon');
  if (isPastedPointer(pointer)) {
    const pasted = rowFromPointer(pointer, selectorCopy.sourcePasted, now);
    const ordered = orderCharacters(me.characters, null, input.chosenKeys);
    const rest = ordered.map((character) => rowFromCharacter(character, false, now));
    return { session: 'signed-in', rows: [pasted, ...rest], current: pasted, hasAddonCharacter };
  }
  const pointed = pointer?.source === 'armory' ? pointer.ref : null;
  const known = me.characters.some((character) => character.key === pointed);
  const currentKey = known ? pointed : (mainCharacter(me.characters, me.main_character_key)?.key ?? null);
  const rows = orderCharacters(me.characters, currentKey, input.chosenKeys).map((character) =>
    rowFromCharacter(character, character.key === currentKey, now),
  );
  return { session: 'signed-in', rows, current: rows.find((row) => row.current) ?? null, hasAddonCharacter };
}

function signedOutModel(pointer: CurrentCharacter | null, now: Date): SelectorModel {
  if (isPastedPointer(pointer)) {
    const pasted = rowFromPointer(pointer, selectorCopy.sourcePasted, now);
    return { session: 'signed-out', rows: [pasted], current: pasted, hasAddonCharacter: false };
  }
  if (pointer?.source === 'armory') {
    const kept = rowFromPointer(pointer, '', now);
    return { session: 'expired', rows: [kept], current: kept, hasAddonCharacter: false };
  }
  return { session: 'signed-out', rows: [], current: null, hasAddonCharacter: false };
}

export function buildSelectorModel(input: ModelInput): SelectorModel {
  const now = new Date(input.nowMs);
  if (input.me === undefined) {
    const kept = input.pointer === null ? null : rowFromPointer(input.pointer, '', now);
    return { session: 'loading', rows: kept === null ? [] : [kept], current: kept, hasAddonCharacter: false };
  }
  if (input.me === null) return signedOutModel(input.pointer, now);
  return signedInModel(input.me, input, now);
}

/** Rows whose name contains `query`, ignoring case; an empty query keeps every row. */
export function filterRows(rows: readonly SelectorRow[], query: string): SelectorRow[] {
  const needle = query.trim().toLowerCase();
  return needle === '' ? [...rows] : rows.filter((row) => row.name.toLowerCase().includes(needle));
}

/** `{Spec} {Class} · {level}`, or `{Spec} · {level}` when `withClass` is false; absent parts are dropped. */
export function closedSecondLine(row: SelectorRow, withClass: boolean): string {
  const identity = [row.spec, withClass || row.spec === undefined ? row.className : undefined]
    .filter((part): part is string => part !== undefined && part !== '')
    .join(' ');
  return [identity, row.level === undefined ? '' : String(row.level)]
    .filter((part) => part !== '')
    .join(' · ');
}

/** List line 2: `{Spec} {Class} · {level} · {Realm}`, with `· {Ruleset}` unless Normal. */
export function rowLineTwo(row: SelectorRow): string {
  const specClass = [row.spec, row.className].filter((part) => part !== undefined && part !== '').join(' ');
  const where = [row.realm, row.ruleset].filter((part) => part !== undefined).join(' · ');
  return [specClass, row.level === undefined ? '' : String(row.level), where]
    .filter((part) => part !== '')
    .join(' · ');
}

export interface RowLineThree {
  text: string;
  /** The ember clause after the plain age on a stale row. */
  tail: string;
  failed: boolean;
}

/** List line 3: source and age, plus the stale tail, or the failed sentence. */
export function rowLineThree(row: SelectorRow): RowLineThree {
  if (row.status === 'failed') {
    return { text: selectorCopy.failedLine(row.age), tail: '', failed: true };
  }
  const text = [row.sourceWord, row.age].filter((part) => part !== '').join(' · ');
  if (row.status !== 'stale') return { text, tail: '', failed: false };
  const bnet = row.sourceWord === selectorCopy.sourceBnet;
  return {
    text,
    tail: bnet ? selectorCopy.staleBnetTail : selectorCopy.staleAddonTail,
    failed: false,
  };
}

export type RingMark = 'stale' | 'failed' | 'session';

/** The mark on the closed crest: an expired session, else failed, else stale, else none. */
export function ringMarkOf(model: SelectorModel): RingMark | null {
  if (model.session === 'expired') return 'session';
  if (model.current === null || model.current.status === 'ok') return null;
  return model.current.status;
}

export function markTooltip(mark: RingMark, row: SelectorRow | null): string {
  if (mark === 'stale') return selectorCopy.markTooltip.stale(row?.age ?? '');
  return selectorCopy.markTooltip[mark];
}

/** `{Name}, {Spec} {Class}, level {n}. {Mark sentence} Change character.` */
export function closedAriaLabel(row: SelectorRow, mark: RingMark | null): string {
  const specClass = [row.spec, row.className].filter((part) => part !== undefined && part !== '').join(' ');
  const identity = [row.name, specClass, row.level === undefined ? '' : `level ${row.level}`]
    .filter((part) => part !== '')
    .join(', ');
  const markSentence = mark === null ? '' : ` ${selectorCopy.markAria[mark]}`;
  return `${identity}.${markSentence} Change character.`;
}
