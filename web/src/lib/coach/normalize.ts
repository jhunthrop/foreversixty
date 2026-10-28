// web/src/lib/coach/normalize.ts
// Turns the coach API's raw JSON into the shapes in types.ts. Like web/src/lib/account's
// normalisers: unknown fields are dropped, a missing or mistyped field takes its empty value,
// and nothing here throws on an older or newer API shape. A shape that cannot be read at all
// (not an object, or no usable `access`) normalises to null.
import type {
  CoachAnchor,
  CoachAnchorKind,
  CoachFull,
  CoachHistoryItem,
  CoachHistoryPage,
  CoachLink,
  CoachRecent,
  CoachReview,
  CoachSource,
  CoachStatus,
  CoachTeaser,
  FixItem,
  LastRaid,
  NightLine,
  ThenItem,
  WellItem,
} from './types';

type Raw = Readonly<Record<string, unknown>>;

const STATUSES: readonly CoachStatus[] = ['ready', 'writing', 'preparing', 'not_reviewable'];
const ANCHOR_KINDS: readonly CoachAnchorKind[] = [
  'cast',
  'aura',
  'death',
  'resource',
  'time',
  'buffs',
  'healing',
];

function asRaw(value: unknown): Raw | null {
  return typeof value === 'object' && value !== null && !Array.isArray(value) ? (value as Raw) : null;
}

const str = (value: unknown): string => (typeof value === 'string' ? value : '');
const num = (value: unknown): number => (typeof value === 'number' && Number.isFinite(value) ? value : 0);
const optNum = (value: unknown): number | undefined =>
  typeof value === 'number' && Number.isFinite(value) ? value : undefined;
const list = (value: unknown): unknown[] => (Array.isArray(value) ? value : []);
const strings = (value: unknown): string[] =>
  list(value).filter((entry): entry is string => typeof entry === 'string');

function mapRaw<T>(value: unknown, read: (raw: Raw) => T | null): T[] {
  return list(value).flatMap((entry) => {
    const raw = asRaw(entry);
    const item = raw === null ? null : read(raw);
    return item === null ? [] : [item];
  });
}

function status(value: unknown): CoachStatus {
  return STATUSES.find((s) => s === value) ?? 'ready';
}

function source(value: unknown): CoachSource {
  return value === 'model' ? 'model' : 'rules';
}

function link(entry: unknown): CoachLink | null {
  if (typeof entry === 'string') return { id: entry, sentence: null };
  const raw = asRaw(entry);
  if (raw === null) return null;
  return { id: str(raw.id), sentence: typeof raw.sentence === 'string' ? raw.sentence : null };
}

function links(value: unknown): CoachLink[] {
  return list(value).flatMap((entry) => {
    const item = link(entry);
    return item === null ? [] : [item];
  });
}

function fixItem(raw: Raw): FixItem {
  return {
    finding: str(raw.finding),
    title: str(raw.title),
    why: str(raw.why),
    drill: str(raw.drill),
    drillText: str(raw.drillText),
    anchors: strings(raw.anchors),
    links: links(raw.links),
  };
}

function thenItem(raw: Raw): ThenItem {
  return {
    finding: str(raw.finding),
    title: str(raw.title),
    why: str(raw.why),
    anchors: strings(raw.anchors),
    links: links(raw.links),
  };
}

function wellItem(raw: Raw): WellItem {
  return { strength: str(raw.strength), text: str(raw.text), anchors: strings(raw.anchors) };
}

function review(value: unknown): CoachReview {
  const raw = asRaw(value) ?? {};
  const fix = asRaw(raw.fixFirst);
  return {
    opening: str(raw.opening),
    fixFirst: fix === null ? null : fixItem(fix),
    thenFix: mapRaw(raw.thenFix, thenItem),
    doingWell: mapRaw(raw.doingWell, wellItem),
    couldntCheck: strings(raw.couldntCheck),
  };
}

function anchor(raw: Raw): CoachAnchor | null {
  const kind = ANCHOR_KINDS.find((k) => k === raw.kind);
  if (kind === undefined) return null;
  const spell = optNum(raw.spell_id);
  const at = optNum(raw.at_ms);
  const to = optNum(raw.to_ms);
  return {
    kind,
    ...(spell === undefined ? {} : { spell_id: spell }),
    ...(typeof raw.spell_name === 'string' && raw.spell_name !== '' ? { spell_name: raw.spell_name } : {}),
    ...(at === undefined ? {} : { at_ms: at }),
    ...(to === undefined ? {} : { to_ms: to }),
    ...(typeof raw.death === 'string' && raw.death !== '' ? { death: raw.death } : {}),
    ...(raw.debuff === true ? { debuff: true } : {}),
    ...(typeof raw.target_guid === 'string' && raw.target_guid !== ''
      ? { target_guid: raw.target_guid }
      : {}),
  };
}

function anchors(value: unknown): Record<string, CoachAnchor> {
  const raw = asRaw(value) ?? {};
  const entries = Object.entries(raw).flatMap(([id, entry]) => {
    const inner = asRaw(entry);
    const read = inner === null ? null : anchor(inner);
    return read === null ? [] : [[id, read] as const];
  });
  return Object.fromEntries(entries);
}

function teaser(raw: Raw): CoachTeaser {
  const first = asRaw(raw.fix_first);
  return {
    status: status(raw.status),
    access: 'teaser',
    finding_count: num(raw.finding_count),
    ...(str(raw.spec) === '' ? {} : { spec: str(raw.spec) }),
    fix_first:
      first === null
        ? null
        : { title: str(first.title), code: str(first.code), spell_id: optNum(first.spell_id) ?? null },
  };
}

function full(raw: Raw, access: 'full' | 'officer'): CoachFull {
  const persona = asRaw(raw.persona) ?? {};
  return {
    report_id: str(raw.report_id),
    fight_index: num(raw.fight_index),
    player_key: str(raw.player_key),
    spec: str(raw.spec),
    role: str(raw.role),
    size: num(raw.size),
    kill: raw.kill === true,
    status: status(raw.status),
    source: source(raw.source),
    busy: raw.busy === true,
    access,
    persona: { label: str(persona.label), status: persona.status === 'starter' ? 'starter' : 'full' },
    reviewed_at: str(raw.reviewed_at),
    review: review(raw.review),
    anchors: anchors(raw.anchors),
    basis_labels: strings(raw.basis_labels),
  };
}

/** A fight answer: a teaser, a full or officer review, or null when it is neither. */
export function normalizeFightCoach(value: unknown): CoachFull | CoachTeaser | null {
  const raw = asRaw(value);
  if (raw === null) return null;
  if (raw.access === 'teaser') return teaser(raw);
  if (raw.access === 'full' || raw.access === 'officer') return full(raw, raw.access);
  return null;
}

function historyItem(raw: Raw): CoachHistoryItem {
  const first = asRaw(raw.fix_first);
  return {
    report_id: str(raw.report_id),
    fight_index: num(raw.fight_index),
    player_key: str(raw.player_key),
    spec: str(raw.spec),
    boss: str(raw.boss),
    kill: raw.kill !== false,
    source: source(raw.source),
    generated_at: str(raw.generated_at),
    fix_first: first === null ? null : { title: str(first.title) },
  };
}

function lastRaid(value: unknown): LastRaid | null {
  const raw = asRaw(value);
  if (raw === null) return null;
  const first = asRaw(raw.first);
  return {
    zone: str(raw.zone),
    fought_at: str(raw.fought_at),
    report_id: str(raw.report_id),
    finding_count: num(raw.finding_count),
    first: first === null ? null : { stem: str(first.stem), on: num(first.on), of: num(first.of) },
    url: str(raw.url),
  };
}

export function normalizeRecent(value: unknown): CoachRecent {
  const raw = asRaw(value) ?? {};
  return {
    last_raid: lastRaid(raw.last_raid),
    pending: raw.pending === true,
    reviews: mapRaw(raw.reviews, historyItem),
  };
}

export function normalizeHistory(value: unknown): CoachHistoryPage {
  const raw = asRaw(value) ?? {};
  const cursor = str(raw.next_cursor);
  return { reviews: mapRaw(raw.reviews, historyItem), next_cursor: cursor === '' ? null : cursor };
}

export function normalizeNightLine(value: unknown): NightLine | null {
  const raw = asRaw(value);
  if (raw === null || str(raw.stem) === '') return null;
  return {
    stem: str(raw.stem),
    on: num(raw.on),
    of: num(raw.of),
    finding_count: num(raw.finding_count),
    fight_index: num(raw.fight_index),
  };
}
