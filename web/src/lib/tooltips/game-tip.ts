// web/src/lib/tooltips/game-tip.ts
// The game's own spell and talent tooltip, as data: name and rank on the first line, the cost
// beside the range, the cast time beside the cooldown, then the description in the client's
// gold. Browser safe: the guide pages build these at build time (lib/guides/guide-tips.ts) and
// the shared tooltip host (components/BisTooltipHost.svelte) renders them in GameTooltip.svelte.
import type { ItemTooltipModel } from '../items/tooltip';

/** One spell's client rows, as `data/builds/<build>/spelltips/*.json` carries them. */
export interface SpellTipRow {
  name: string;
  /** 0 for a spell the client gives no rank. */
  rank: number;
  /** In the power's client units: rage is stored in tenths. */
  cost: number;
  /** 0 mana, 1 rage, 2 focus, 3 energy. */
  cost_type: number;
  /** A percent of base mana; 0 when the cost is flat. */
  cost_pct: number;
  cast_time_ms: number;
  channeled: boolean;
  passive: boolean;
  cooldown_ms: number;
  /** Maximum range in yards; 0 for self. */
  range_yd: number;
  description: string;
  /** The spell's own client icon; absent on a file written before it carried one. */
  icon?: string;
}

/** A tooltip line pair: the left text and, when the client prints one, the right. */
export interface TipLine {
  left: string;
  right?: string;
}

export interface GameTipModel {
  name: string;
  /** The icon stem, absent when the build ships none. */
  icon?: string;
  /** "Rank 4", "Racial", "Rank 2/5": the muted text right of the name. */
  subtitle?: string;
  lines: TipLine[];
  /** The client's gold description; empty when the client stores none. */
  description: string;
  /** A talent short of its last rank: the next rank's text, under "Next rank". */
  nextRank?: string;
}

/** A talent's tooltip source: every rank's text, so the rank can be chosen where it is shown. */
export interface TalentTipEntry {
  name: string;
  icon?: string;
  maxRank: number;
  ranks: string[];
}

/** The guide's lazily fetched tooltip data (`/data/<build>/guide-tips/<class>/<spec>.json`). */
export interface GuideTipBundle {
  /** `spell:<id>`, `racial:<name>`, `enchant:<id>` -> a ready model. */
  tips: Record<string, GameTipModel>;
  /** Talent id -> its ranks; the key `talent:<id>:<rank>` picks the rank. */
  talents: Record<string, TalentTipEntry>;
  /** Item id -> the shared item tooltip's model, for the guide's gear and consumables. */
  items: Record<string, ItemTooltipModel>;
}

export type GameTipKind = 'spell' | 'talent' | 'racial' | 'enchant';

export function spellTipKey(spellId: number): string {
  return `spell:${spellId}`;
}

export function talentTipKey(talentId: number, rank: number): string {
  return `talent:${talentId}:${rank}`;
}

export function racialTipKey(name: string): string {
  return `racial:${name}`;
}

export function enchantTipKey(enchantId: number): string {
  return `enchant:${enchantId}`;
}

const POWER_NAMES: Readonly<Record<number, string>> = { 0: 'Mana', 1: 'Rage', 2: 'Focus', 3: 'Energy' };
/** The client stores rage in tenths of a point. */
const RAGE_UNITS = 10;
/** The client prints "Melee Range" for anything this close. */
const MELEE_RANGE_YD = 5;

function seconds(ms: number): string {
  const value = ms / 1000;
  return Number.isInteger(value) ? String(value) : value.toFixed(1).replace(/\.0$/, '');
}

/** `SpellPower.PowerType` for a cost paid in health (Bloodrage, Life Tap). */
const HEALTH_POWER = -2;
/** The pools a percent cost can name: "of base mana", "of base health". */
const PERCENT_POOLS: Readonly<Record<number, string>> = { 0: 'mana', [HEALTH_POWER]: 'health' };

/** "30 Rage", "250 Mana", "12% of base mana", "20% of base health"; undefined for a free spell
 *  or a cost in a pool the client names no other way. */
export function costText(row: Pick<SpellTipRow, 'cost' | 'cost_type' | 'cost_pct'>): string | undefined {
  if (row.cost_pct > 0) {
    const pool = PERCENT_POOLS[row.cost_type];
    return pool === undefined ? undefined : `${seconds(row.cost_pct * 1000)}% of base ${pool}`;
  }
  if (row.cost <= 0) return undefined;
  const power = POWER_NAMES[row.cost_type];
  if (power === undefined) return undefined;
  const amount = row.cost_type === 1 ? row.cost / RAGE_UNITS : row.cost;
  return `${amount} ${power}`;
}

/** "Melee Range", "30 yd range"; undefined for a spell on the caster. */
export function rangeText(rangeYd: number): string | undefined {
  if (rangeYd <= 0) return undefined;
  return rangeYd <= MELEE_RANGE_YD ? 'Melee Range' : `${seconds(rangeYd * 1000)} yd range`;
}

/** "Instant", "Channeled", "2.5 sec cast". */
export function castText(row: Pick<SpellTipRow, 'cast_time_ms' | 'channeled'>): string {
  if (row.channeled) return 'Channeled';
  return row.cast_time_ms > 0 ? `${seconds(row.cast_time_ms)} sec cast` : 'Instant';
}

/** "6 sec cooldown", "2 min cooldown", "1.5 min cooldown"; undefined without one. The global
 *  cooldown is not a cooldown the client prints. */
export function cooldownText(cooldownMs: number): string | undefined {
  if (cooldownMs <= 1500) return undefined;
  if (cooldownMs >= 60_000) return `${seconds(cooldownMs / 60)} min cooldown`;
  return `${seconds(cooldownMs)} sec cooldown`;
}

/**
 * The client's description as its tooltip prints it for a character with no talents or
 * auras: `$?s123[with][without]` takes the second branch, `$gmale:female;` both words, and
 * `$Lone:many;` agrees with the number before it. Double spaces after a stop are one space.
 * A sentence the client finishes from the character's own stats (`${$b1*$<mult>}`) is left
 * out (`withoutUnresolved`): a page never prints a raw token.
 */
export function displayDescription(text: string): string {
  let out = text;
  // A conditional's branches can hold one more level of brackets; resolve inner ones first. A
  // chain (`$?s2060[Greater Heal]?s5185[Healing Touch][either]`) reads as its last, general branch.
  const CONDITIONAL =
    /\$\?[!a-z]?\d*(?:[&|]!?[a-z]?\d+)*\[[^[\]]*\](?:\?[!a-z]?\d*(?:[&|]!?[a-z]?\d+)*\[[^[\]]*\])*\[([^[\]]*)\]/gi;
  for (let pass = 0; pass < 3 && CONDITIONAL.test(out); pass += 1) {
    out = out.replace(CONDITIONAL, (_match, no: string) => no);
  }
  out = out.replace(
    /\$[gG]([^:;]*):([^;]*);/g,
    (_match, male: string, female: string) => `${male} or ${female}`,
  );
  out = out.replace(
    /(\d+(?:\.\d+)?)([^$\d]*?)\$[lL]([^:;]*):([^;]*);/g,
    (_match, num: string, between: string, one: string, many: string) =>
      `${num}${between}${Number(num) === 1 ? one : many}`,
  );
  return withoutUnresolved(
    out
      .replace(COLOUR_CODE, '')
      .replace(/\r\n/g, '\n')
      .replace(/([.!?]) {2,}/g, '$1 ')
      .trim(),
  );
}

/** A client token left in the text: `$s1`, `${...}`, `$<mult>`, `$ec1`, `$@spelltooltip1`. */
const UNRESOLVED = /\$[{<@a-zA-Z0-9*/]/;
/** The client's colour escapes, `|cFFFF2020Requires Bear Form|r`: the words stay, the codes go. */
const COLOUR_CODE = /\|c[0-9a-f]{8}|\|r/gi;
/** Where a description's sentences and lines end. */
const SENTENCE_END = /(?<=[.!?])\s+|\n+/;

/** The description with every sentence or line that still holds a client token dropped. */
export function withoutUnresolved(text: string): string {
  if (!UNRESOLVED.test(text)) return text;
  return text
    .split(SENTENCE_END)
    .filter((sentence) => sentence.trim() !== '' && !UNRESOLVED.test(sentence))
    .join(' ')
    .trim();
}

/** A spell's tooltip: name and rank, cost and range, cast and cooldown, the description. */
export function spellTipModel(row: SpellTipRow, icon: string | undefined, subtitle?: string): GameTipModel {
  const lines: TipLine[] = [];
  if (row.passive) {
    lines.push({ left: 'Passive' });
  } else {
    const cost = costText(row);
    const range = rangeText(row.range_yd);
    if (cost !== undefined || range !== undefined) {
      lines.push({ left: cost ?? '', ...(range === undefined ? {} : { right: range }) });
    }
    const cooldown = cooldownText(row.cooldown_ms);
    lines.push({ left: castText(row), ...(cooldown === undefined ? {} : { right: cooldown }) });
  }
  const rankLabel = subtitle ?? (row.rank > 0 ? `Rank ${row.rank}` : undefined);
  return {
    name: row.name,
    ...(icon === undefined ? {} : { icon }),
    ...(rankLabel === undefined ? {} : { subtitle: rankLabel }),
    lines,
    description: displayDescription(row.description),
  };
}

/** A talent at `rank` (0 for untaken): the rank's own text, and the next rank's when there is
 *  one, the way the planner's talent tooltip reads. */
export function talentTipModel(entry: TalentTipEntry, rank: number): GameTipModel {
  const shown = Math.max(0, Math.min(rank, entry.maxRank));
  const current = shown > 0 ? entry.ranks[shown - 1] : entry.ranks[0];
  const next = shown > 0 && shown < entry.maxRank ? entry.ranks[shown] : undefined;
  return {
    name: entry.name,
    ...(entry.icon === undefined ? {} : { icon: entry.icon }),
    subtitle: `Rank ${shown}/${entry.maxRank}`,
    lines: [],
    description: displayDescription(current ?? ''),
    ...(next === undefined ? {} : { nextRank: displayDescription(next) }),
  };
}

/** The model a host's key names in a bundle, or null when the bundle does not carry it. */
export function tipFromBundle(bundle: GuideTipBundle, key: string): GameTipModel | null {
  const talent = /^talent:(\d+):(\d+)$/.exec(key);
  if (talent !== null) {
    const entry = bundle.talents[talent[1]!];
    return entry === undefined ? null : talentTipModel(entry, Number(talent[2]));
  }
  return bundle.tips[key] ?? null;
}
