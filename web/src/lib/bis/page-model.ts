// web/src/lib/bis/page-model.ts
// The BiS spec page's compact, serializable model (web layout spec 2026-11-05 §5, owner decision
// D10): one entry per faction × band × preset, each holding only what the paperdoll, its detail,
// the answer in the page head and the weights need, plus one item dictionary for the page. The page
// builds it at build time from the published file (panel-view.ts already resolves sources, empty
// reasons and weights), ships it as one JSON block, and the island derives every view from it with
// no network request (R2). Every field the nightly adds later (`by_kind`, `ranking`, `run`,
// `filtered_set_dps`, `set_dps_error`) is read when present and absent otherwise.
import { oneDecimal } from '../format/number';
import type { ItemTooltipModel } from '../items/tooltip';
import { bisCopy } from './copy';
import type { HealerSetView } from './heal-view';
import type { HitCapLine } from './hit-cap';
import type { ExpertiseCapLine } from './expertise-cap';
import type { BandInfo, RowView } from './panel-view';
import { raidCaptionFor, effectLineFor } from './preset-caption';
import { presetOf } from './presets';
import { bandModelKey } from './model-key';
import { bandTopLevel } from './rotation-view';
import type { ScaleRow } from './scale-view';
import { filterKindOfSource, pickKinds } from './source-filter';
import { resolveSourceCell, type SourceCell } from './source-cell';
import { fullSourceLine, standingOf, type InstanceLevels } from './source-line';
import type { TankHeadline } from './tank-view';
import type {
  BisAlternative,
  BisBand,
  BisByKindEntry,
  BisFile,
  BisFilteredSetDps,
  BisFilteredSetKey,
  BisPresetId,
  BisSlot,
  Faction,
  ItemDetail,
  LootQuestsFile,
  SourceFilterKind,
} from './types';
import type { LootFile } from '../sim/loot';

import { grouped } from '../format/number';
import { upperFirst } from '../text';
/** The unit a band's figures are in: damage, healing, or a tank's score. */
export type GainUnit = 'DPS' | 'HPS' | 'score';

/** One item, as every view on the page shows it. */
export interface BisItem {
  name: string;
  quality: number;
  icon?: string;
  ilvl?: number;
  req?: number;
  stats?: Record<string, number>;
  /** `equip` for bind on equip. */
  bind?: string;
  effect?: string;
}

/** Where one candidate comes from: the full line, its loot kind, and its group for Where to go. */
export interface BisSourceInfo {
  line: string;
  /** The resolved source kind (`dungeon`, `crafted`, `world_drop`, ...). */
  kind: string;
  /** The Where to go group's name: the instance, the profession, the faction; absent for the
   *  kinds that group by kind alone (quest, vendor, world drop). */
  place?: string;
  /** A reputation source's standing, lower case. */
  standing?: string;
}

/** A runner-up or a trinket the tournament measured, against the pick. */
export interface BisCandidate {
  id: number;
  delta: number;
}

/** The source-filter fields of one slot: the subset of `BisSlot` `source-filter.ts` reads. */
export type BisFilterSlot = Pick<
  BisSlot,
  | 'slot'
  | 'item_id'
  | 'source_kind'
  | 'kinds'
  | 'two_hand'
  | 'not_with'
  | 'by_kind'
  | 'standing'
  | 'kind_standings'
>;

export interface BisRowModel {
  slot: string;
  label: string;
  /** The pick; absent for an empty slot. */
  id?: number;
  /** An empty slot's reason (the shipped copy, data-driven). */
  empty?: string;
  /** The two-hander that takes this off hand. */
  coveredBy?: number;
  /** The enchant the raid-ready set wears here, by name. */
  enchant?: string;
  setBonus?: string;
  /** Runners-up, best first, DPS against the pick (negative is worse). */
  alts: BisCandidate[];
  /** Trinkets: the tournament that decided the pick (`BisSlot.ranking`). */
  ranking?: BisCandidate[];
  /** Every candidate this row can show, by id: the pick, the runners-up, the per-kind options. */
  sources: Record<number, BisSourceInfo>;
  filter: BisFilterSlot;
}

export interface BisAnswerModel {
  /** The set's figure, one decimal, in `unit`. Absent on a tank band (its four figures). */
  figure?: string;
  error?: number;
  filtered?: Partial<Record<BisFilteredSetKey, BisFilteredSetDps>>;
  tank?: TankHeadline;
  healer?: HealerSetView;
  /** The trust line's facts, in order (§5.10 item 6). */
  trust: string[];
  simmed: string;
  simHref?: string;
  rankingSimHref?: string;
  plannerHref: string;
  /** The raid preset's effect names as one plain line, for the trust line's hover title. */
  trustTitle?: string;
}

export interface BisWeightsModel {
  /** "Stack agility, then hit rating, then ranged attack power." Absent when nothing is significant. */
  sentence?: string;
  rows: ScaleRow[];
  note: string;
  haste?: string;
  hitCap?: HitCapLine;
  expertiseCap?: ExpertiseCapLine;
}

export interface BisBandModel {
  key: string;
  faction: Faction;
  band: number;
  preset: BisPresetId;
  /** "20 to 29", "60". */
  label: string;
  unit: GainUnit;
  race: string;
  answer: BisAnswerModel;
  weights: BisWeightsModel;
  rows: BisRowModel[];
  /** True when the file publishes no `by_kind`: filters re-pick over the listed runners-up. */
  listedOnly: boolean;
  /** A two-handed main hand: "Legionite Glaive is a two-hander: the off hand stays empty". */
  weaponNote?: string;
}

export interface BisPageData {
  spec: string;
  build: string;
  bands: number[];
  models: Record<string, BisBandModel>;
  items: Record<number, BisItem>;
  /** The inlined default slice: the full set is the page's JSON endpoint. */
  partial?: boolean;
}

/** The page's JSON endpoint: every model and item, and every item's tooltip (the page inlines only
 *  its slice's tooltips). */
export interface BisFullData extends BisPageData {
  tooltips: Record<number, ItemTooltipModel>;
}

export { bandModelKey };

/** Enchant names by (id, slot): the same id names one enchant per slot in enchants.json. */
export type EnchantNames = (id: number, slot: string) => string | undefined;

export interface PageModelDeps {
  file: BisFile;
  details: ReadonlyMap<number, ItemDetail>;
  loot: LootFile & Partial<LootQuestsFile>;
  instanceLevels: InstanceLevels;
  enchantName: EnchantNames;
  unit: GainUnit;
  simHrefFor: (band: BisBand) => string | undefined;
  rankingSimHrefFor: (band: BisBand) => string | undefined;
  plannerHrefFor: (band: BisBand) => string;
}

function sourceInfo(
  cand: Pick<BisSlot, 'item_id' | 'source_kind'> & { source?: string },
  faction: Faction,
  deps: PageModelDeps,
): BisSourceInfo {
  const cell = resolveSourceCell(cand, faction, deps.loot, cand.source ?? cand.source_kind);
  return {
    line: fullSourceLine(cell, {
      instanceLevels: deps.instanceLevels,
      bonding: deps.details.get(cand.item_id)?.bonding,
    }),
    kind: cell.kind === 'unknown' ? cand.source_kind : cell.kind,
    place: placeOf(cell),
    standing: standingOf(cell),
  };
}

function placeOf(cell: SourceCell): string | undefined {
  switch (cell.kind) {
    case 'dungeon':
    case 'raid':
      return cell.instance;
    case 'crafted':
      return cell.name.replace(/\b\w/g, (letter) => letter.toUpperCase());
    case 'rep':
      return cell.faction;
    default:
      return undefined;
  }
}

/** A file published before `by_kind`: each kind's best listed runner-up stands in for it, so the
 *  filters re-pick over what the page already shows (labelled as such). */
function listedByKind(slot: BisSlot, sources: Record<number, BisSourceInfo>): BisSlot['by_kind'] {
  const own = new Set(pickKinds(slot));
  const out: Partial<Record<SourceFilterKind, BisByKindEntry>> = {};
  (slot.alternatives ?? []).forEach((alt: BisAlternative, index) => {
    const kind = filterKindOfSource(alt.source_kind);
    if (kind === undefined || own.has(kind) || out[kind] !== undefined) return;
    out[kind] = {
      item_id: alt.item_id,
      rank: index + 1,
      source_kind: alt.source_kind,
      source: alt.source,
      dps_delta: alt.dps_delta,
      ...(sources[alt.item_id]?.standing ? { standing: sources[alt.item_id]!.standing } : {}),
    };
  });
  return out;
}

/** Every by_kind entry (and its `next` chain) with the kind it was listed under: the nightly omits
 *  `source_kind` where the key already says it. */
function byKindEntries(byKind: BisSlot['by_kind']): (BisByKindEntry & { source_kind: string })[] {
  const out: (BisByKindEntry & { source_kind: string })[] = [];
  for (const [kind, entry] of Object.entries(byKind ?? {})) {
    for (let node: BisByKindEntry | undefined = entry; node !== undefined; node = node.next) {
      out.push({ ...node, source_kind: node.source_kind ?? kind });
    }
  }
  return out;
}

function enchantLabel(name: string): string {
  return name.replace(/^Enchant [\w ]+? - /, '');
}

function rowModel(
  row: RowView,
  raw: BisSlot | undefined,
  faction: Faction,
  deps: PageModelDeps,
  mainHand: BisSlot | undefined,
): BisRowModel {
  const sources: Record<number, BisSourceInfo> = {};
  const base = { slot: row.slot, label: row.slotLabel, alts: [] as BisCandidate[], sources };
  const hasPick = raw !== undefined && typeof raw.item_id === 'number' && !row.empty;
  if (!hasPick) {
    const coveredBy =
      row.slot === 'off_hand' && mainHand !== undefined && isTwoHand(mainHand, deps.details)
        ? mainHand.item_id
        : undefined;
    return {
      ...base,
      empty: row.emptyCopy,
      coveredBy,
      filter: { slot: row.slot, item_id: 0, source_kind: '', by_kind: raw?.by_kind },
    };
  }
  sources[raw.item_id] = sourceInfo(raw, faction, deps);
  for (const alt of raw.alternatives ?? []) sources[alt.item_id] = sourceInfo(alt, faction, deps);
  const published = raw.by_kind !== undefined;
  const byKind = published ? raw.by_kind : listedByKind(raw, sources);
  for (const entry of byKindEntries(byKind)) {
    sources[entry.item_id] ??= sourceInfo(
      { item_id: entry.item_id, source_kind: entry.source_kind, source: entry.source },
      faction,
      deps,
    );
  }
  for (const entry of raw.ranking ?? []) {
    sources[entry.item_id] ??= sourceInfo(
      { item_id: entry.item_id, source_kind: entry.source_kind ?? '' },
      faction,
      deps,
    );
  }
  const enchant = raw.enchant_id === undefined ? undefined : deps.enchantName(raw.enchant_id, raw.slot);
  return {
    ...base,
    id: raw.item_id,
    alts: (raw.alternatives ?? []).map((alt) => ({ id: alt.item_id, delta: alt.dps_delta })),
    ...(raw.ranking
      ? { ranking: raw.ranking.map((entry) => ({ id: entry.item_id, delta: entry.dps_delta })) }
      : {}),
    ...(enchant ? { enchant: enchantLabel(enchant) } : {}),
    ...(row.setBonus ? { setBonus: row.setBonus.line } : {}),
    filter: {
      slot: raw.slot,
      item_id: raw.item_id,
      source_kind: raw.source_kind,
      kinds: raw.kinds,
      two_hand: raw.two_hand ?? isTwoHand(raw, deps.details),
      not_with: raw.not_with,
      by_kind: byKind,
      standing: raw.standing ?? sources[raw.item_id]?.standing,
      ...(raw.kind_standings ? { kind_standings: raw.kind_standings } : {}),
    },
  };
}

function isTwoHand(slot: BisSlot, details: ReadonlyMap<number, ItemDetail>): boolean {
  return slot.two_hand ?? details.get(slot.item_id)?.two_hand ?? false;
}

/** "Stack agility, then hit rating, then ranged attack power." from the significant weights. */
export function stackSentence(rows: readonly ScaleRow[]): string | undefined {
  const top = rows
    .filter((row) => row.significant)
    .sort((a, b) => b.scaleFactor - a.scaleFactor)
    .slice(0, 3)
    .map((row) => row.label.toLowerCase());
  if (top.length === 0) return undefined;
  return `Stack ${top.join(', then ')}.`;
}

function trustFacts(file: BisFile, raw: BisBand, info: BandInfo): string[] {
  const facts = [
    `Level ${bandTopLevel(raw.band)} ${upperFirst(raw.race)}`,
    `${raw.talent_points} talent points`,
  ];
  const run = file.run;
  if (run !== undefined) {
    const target = run.target_type.replace(/_/g, ' ');
    const level = raw.target_level === undefined ? '' : `level ${raw.target_level} `;
    facts.push(`${level}${target}, ${run.duration_sec} s`);
  }
  const preset = presetOf(raw);
  const meta = file.presets?.[preset];
  facts.push(preset === 'raid' && meta !== undefined ? raidCaptionFor(meta) : bisCopy.presetBareCaption);
  if (run !== undefined) facts.push(`${grouped(run.iterations.set)} runs`);
  if (info.setDpsPartial) facts.push(bisCopy.setDpsPartialNote(info.setDpsPartialCount, info.unit));
  return facts;
}

function answerModel(file: BisFile, raw: BisBand, info: BandInfo, deps: PageModelDeps): BisAnswerModel {
  const preset = presetOf(raw);
  const meta = preset === 'raid' ? file.presets?.raid : undefined;
  return {
    figure: info.tank === undefined ? (info.healer?.figure ?? oneDecimal(raw.set_dps)) : undefined,
    ...(raw.set_dps_error === undefined ? {} : { error: raw.set_dps_error }),
    ...(raw.filtered_set_dps === undefined ? {} : { filtered: raw.filtered_set_dps }),
    ...(info.tank === undefined ? {} : { tank: info.tank }),
    ...(info.healer === undefined ? {} : { healer: info.healer }),
    trust: trustFacts(file, raw, info),
    simmed: bisCopy.simmedLine(file.generated_at.slice(0, 10)),
    simHref: info.healer === undefined ? deps.simHrefFor(raw) : undefined,
    rankingSimHref: deps.rankingSimHrefFor(raw),
    plannerHref: deps.plannerHrefFor(raw),
    ...(meta === undefined ? {} : { trustTitle: effectLineFor(meta) }),
  };
}

/** One faction × band × preset, from its resolved `BandInfo` and its published entry. */
export function bandModelFor(info: BandInfo, raw: BisBand, deps: PageModelDeps): BisBandModel {
  const faction = raw.faction as Faction;
  const bySlot = new Map(raw.slots.map((slot) => [slot.slot, slot]));
  const mainHand = bySlot.get('main_hand');
  const rows = info.rows.map((row) => rowModel(row, bySlot.get(row.slot), faction, deps, mainHand));
  const twoHander =
    mainHand !== undefined && typeof mainHand.item_id === 'number' && isTwoHand(mainHand, deps.details);
  return {
    key: bandModelKey(faction, raw.band, presetOf(raw)),
    faction,
    band: raw.band,
    preset: presetOf(raw),
    label: bisCopy.bandRangeLabel(raw.band),
    unit: deps.unit,
    race: raw.race,
    answer: answerModel(deps.file, raw, info, deps),
    weights: {
      sentence: stackSentence(info.scaleRows),
      rows: info.scaleRows,
      note: info.scaleNoteLine,
      ...(info.hasteCaptionLine === undefined ? {} : { haste: info.hasteCaptionLine }),
      ...(info.hitCap === undefined ? {} : { hitCap: info.hitCap }),
      ...(info.expertiseCap === undefined ? {} : { expertiseCap: info.expertiseCap }),
    },
    rows,
    listedOnly: !raw.slots.some((slot) => slot.by_kind !== undefined),
    ...(twoHander ? { weaponNote: `${mainHand.item_name} is a two-hander: the off hand stays empty.` } : {}),
  };
}

/** Every id a set of band models names, for the item dictionary and the tooltip models. */
export function modelItemIds(models: Iterable<BisBandModel>): Set<number> {
  const ids = new Set<number>();
  for (const model of models) {
    for (const row of model.rows) {
      if (row.id !== undefined) ids.add(row.id);
      if (row.coveredBy !== undefined) ids.add(row.coveredBy);
      for (const id of Object.keys(row.sources)) ids.add(Number(id));
    }
  }
  return ids;
}

/** The item dictionary for the ids the models name; ids the build's item file lacks fall back to
 *  the name the BiS file published. */
export function itemDictionary(
  ids: Iterable<number>,
  details: ReadonlyMap<number, ItemDetail>,
  fallbackName: (id: number) => string | undefined,
): Record<number, BisItem> {
  const out: Record<number, BisItem> = {};
  for (const id of ids) {
    const detail = details.get(id);
    if (detail === undefined) {
      out[id] = { name: fallbackName(id) ?? `Item ${id}`, quality: 1 };
      continue;
    }
    out[id] = {
      name: detail.name,
      quality: detail.quality,
      icon: detail.icon,
      ilvl: detail.item_level,
      req: detail.required_level,
      ...(Object.keys(detail.stats).length > 0 ? { stats: detail.stats } : {}),
      ...(detail.bonding ? { bind: detail.bonding } : {}),
      ...(detail.effect_text ? { effect: detail.effect_text } : {}),
    };
  }
  return out;
}
