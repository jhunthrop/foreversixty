// web/src/lib/guild/report-night-view.ts
// The pure half of the done raid's report page (web layout spec 2026-11-05 §7.2 "Raid report"):
// the answer figures, the pull rows and their death causes, the DPS and HPS rows by boss kill,
// the consumables-by-pull grid and the Discord summary's preview. Every number is the API's; a
// member's report already carries only their own deaths, output and consumables, so nothing here
// ranks or filters other raiders. Strings come from report-night-copy.ts.
import { formatAmount } from '../report/format';
import type { CellTone, GridCell, GridRowModel } from '../ui/cells';
import type {
  PullResult,
  RaidNight,
  RaidPullConsumable,
  RaidPullConsumables,
  RaidReport,
  RaidReportDeath,
  RaidReportHit,
  RaidReportParse,
  RaidReportPull,
} from './attendance-contract';
import { outputFigure, outputUnit } from './output-figure';
import {
  causeCount,
  diedTo,
  hitAmount,
  intoPull,
  pullTitle,
  ratedOf,
  reportNightCopy as copy,
  listWords,
  unknownPullLine,
} from './report-night-copy';
import { SEP } from './raid-copy';
import type { StateTone } from './raid-planner-view';
import { cameRatio, deaths as deathsWord, durationLabel, killTimeLabel } from './attendance-copy';

/** The engine's id for a melee swing. */
export const MELEE_SPELL_ID = 1;
/** The Discord embed's own caps (api/internal/notify/report_embed.go). */
const DISCORD_TOP_OUTPUT = 3;
const DISCORD_MAX_LOOT = 5;
const DISCORD_MAX_KILLS = 6;

export interface HeadlineFigure {
  key: 'kills' | 'wipes' | 'attempts' | 'deaths';
  label: string;
  value: string;
  tone: 'strong' | 'kill' | 'wipe';
}

/** Kills, wipes, attempts and deaths, in that order: the page's answer. */
export function headlineFigures(report: RaidReport): HeadlineFigure[] {
  const { kills, wipes, attempts, deaths } = report.headline;
  return [
    { key: 'kills', label: copy.figures.kills, value: String(kills), tone: kills > 0 ? 'kill' : 'strong' },
    { key: 'wipes', label: copy.figures.wipes, value: String(wipes), tone: wipes > 0 ? 'wipe' : 'strong' },
    { key: 'attempts', label: copy.figures.attempts, value: String(attempts), tone: 'strong' },
    { key: 'deaths', label: copy.figures.deaths, value: String(deaths), tone: 'strong' },
  ];
}

const RESULT_TONE: Readonly<Record<PullResult, StateTone>> = {
  kill: 'success',
  wipe: 'danger',
  clear: 'muted',
};

export interface PullRowModel {
  key: string;
  pull: RaidReportPull;
  title: string;
  sub: string;
  tag: { tone: StateTone; text: string };
  /** The row opens its deaths: some to show, or deaths the log has not been read for. */
  expandable: boolean;
}

/** `died to Deep Breath ×3, Tail Sweep ×2`: the killing blows' abilities, most first. */
export function causeSummary(causes: readonly RaidReportDeath[]): string | null {
  const counts = new Map<string, number>();
  for (const death of causes) {
    const ability = death.killing_blow === null ? null : abilityName(death.killing_blow);
    if (ability !== null) counts.set(ability, (counts.get(ability) ?? 0) + 1);
  }
  if (counts.size === 0) return null;
  const ordered = [...counts.entries()].sort((a, b) => b[1] - a[1] || a[0].localeCompare(b[0]));
  return diedTo(ordered.map(([ability, n]) => causeCount(ability, n)));
}

function pullSub(pull: RaidReportPull): string {
  const parts = [pull.attempt === null ? copy.pulls.trash : null, killTimeLabel(pull.duration_ms)];
  if (pull.deaths > 0) parts.push(deathsWord(pull.deaths));
  const causes = pull.death_causes === null ? null : causeSummary(pull.death_causes);
  if (causes !== null) parts.push(causes);
  return parts.filter((part): part is string => part !== null).join(SEP);
}

export function pullRows(report: RaidReport): PullRowModel[] {
  return report.pulls.map((pull) => ({
    key: String(pull.fight_index),
    pull,
    title: pullTitle(pull.name, pull.attempt),
    sub: pullSub(pull),
    tag: { tone: RESULT_TONE[pull.result], text: copy.result[pull.result] },
    expandable: pull.death_causes === null ? pull.deaths > 0 : pull.death_causes.length > 0,
  }));
}

/** The pull the wide inspector opens on: the first wipe with deaths, else the first with deaths. */
export function defaultPullKey(rows: readonly PullRowModel[]): string | null {
  const open = rows.filter((row) => row.expandable);
  return (open.find((row) => row.pull.result === 'wipe') ?? open[0])?.key ?? null;
}

/** A melee swing reads `Melee`, whatever the log called it. */
export function abilityName(hit: RaidReportHit): string {
  return hit.spell_id === MELEE_SPELL_ID || hit.ability === '' ? copy.pulls.melee : hit.ability;
}

export interface HitModel {
  ability: string;
  spellId: number | null;
  source: string;
  clock: string;
  amount: string;
}

export function hitModel(hit: RaidReportHit): HitModel {
  return {
    ability: abilityName(hit),
    spellId: hit.spell_id === MELEE_SPELL_ID ? null : hit.spell_id,
    source: hit.source,
    clock: killTimeLabel(hit.at_ms),
    amount: hitAmount(formatAmount(hit.amount), hit.overkill > 0 ? formatAmount(hit.overkill) : null),
  };
}

export interface DeathModel {
  key: string;
  name: string;
  class: string;
  at: string;
  killingBlow: HitModel | null;
  /** The hits before the killing blow, oldest first. */
  before: HitModel[];
}

function sameHit(a: RaidReportHit, b: RaidReportHit): boolean {
  return a.at_ms === b.at_ms && a.spell_id === b.spell_id && a.amount === b.amount;
}

export function deathModels(causes: readonly RaidReportDeath[]): DeathModel[] {
  return causes.map((death, index) => {
    const blow = death.killing_blow;
    return {
      key: `${death.name}-${death.at_ms}-${index}`,
      name: death.name,
      class: death.class,
      at: intoPull(killTimeLabel(death.at_ms)),
      killingBlow: blow === null ? null : hitModel(blow),
      before: death.last_hits.filter((hit) => blow === null || !sameHit(hit, blow)).map(hitModel),
    };
  });
}

export interface OutputGroup {
  key: string;
  boss: string;
  rows: (RaidReportParse & { figure: string; unit: string; sub: string })[];
}

/** The DPS and HPS rows by boss kill, in the API's order (by kill, damage before healing, best first). */
export function outputGroups(parses: readonly RaidReportParse[]): OutputGroup[] {
  const groups = new Map<number, OutputGroup>();
  for (const parse of parses) {
    const group = groups.get(parse.fight_index) ?? {
      key: String(parse.fight_index),
      boss: parse.encounter,
      rows: [],
    };
    group.rows.push({
      ...parse,
      figure: outputFigure(parse.value),
      unit: outputUnit(parse.metric),
      sub:
        parse.rating === null
          ? roleWord(parse.role)
          : `${roleWord(parse.role)}${SEP}${ratedOf(parse.rating)}`,
    });
    groups.set(parse.fight_index, group);
  }
  return [...groups.values()];
}

const ROLE_WORDS: Readonly<Record<string, string>> = { tank: 'Tank', healer: 'Healer', dps: 'DPS' };
const roleWord = (role: string): string => ROLE_WORDS[role] ?? role;

export interface ConsumablesModel {
  columns: string[];
  rows: (GridRowModel & { class: string })[];
  /** One muted line per pull the log cannot speak for. */
  unknown: string[];
}

function columnTitle(pull: RaidPullConsumables, bosses: number): string {
  return bosses > 1 ? pullTitle(pull.name, pull.attempt) : `Pull ${pull.attempt}`;
}

const MARK = (value: boolean | null): string => (value === null ? '·' : value ? '✓' : '×');

function cellTone(row: RaidPullConsumable): CellTone {
  if (row.flask === false || row.food === false || row.weapon_enchant === false) return 'miss';
  if (row.flask === true && row.food === true && row.weapon_enchant === true) return 'ok';
  return 'warn';
}

function cellLabel(row: RaidPullConsumable): string {
  const brought = row.brought.length === 0 ? 'nothing the log can name' : row.brought.join(', ');
  const potions = row.potions > 0 ? `, ${row.potions} potion${row.potions === 1 ? '' : 's'}` : '';
  return `Up: ${brought}${potions}`;
}

function consumableCell(pull: RaidPullConsumables, key: string): GridCell {
  if (pull.rows === null) {
    return { tone: 'none', text: '?', label: copy.consumables.reasons[pull.reason ?? 'not_read'] };
  }
  const row = pull.rows.find((r) => r.character_key === key || r.name === key);
  if (row === undefined) return { tone: 'none', text: '–', label: copy.consumables.notInPull };
  return {
    tone: cellTone(row),
    text: `${MARK(row.flask)}${MARK(row.food)}${MARK(row.weapon_enchant)}`,
    label: cellLabel(row),
  };
}

/** Raider by boss pull; a pull with no rows draws `?` cells (never red) and says why below. */
export function consumablesModel(byPull: readonly RaidPullConsumables[]): ConsumablesModel {
  const bosses = new Set(byPull.map((pull) => pull.name)).size;
  const raiders = new Map<string, { name: string; class: string }>();
  for (const pull of byPull) {
    for (const row of pull.rows ?? []) {
      const key = row.character_key === '' ? row.name : row.character_key;
      if (!raiders.has(key)) raiders.set(key, { name: row.name, class: row.class });
    }
  }
  const rows = [...raiders.entries()]
    .sort((a, b) => a[1].name.localeCompare(b[1].name))
    .map(([key, who]) => ({
      key,
      label: who.name,
      class: who.class,
      cells: byPull.map((pull) => consumableCell(pull, key)),
    }));
  return { columns: byPull.map((pull) => columnTitle(pull, bosses)), rows, unknown: unknownLines(byPull) };
}

/** One line per reason the log cannot say, naming every pull it covers: "Onyxia, pull 1 and
 *  Onyxia, pull 2: the log carried no character sheet, ...", never the same sentence twice. */
function unknownLines(byPull: readonly RaidPullConsumables[]): string[] {
  const titles = new Map<string, string[]>();
  for (const pull of byPull) {
    if (pull.rows !== null) continue;
    const reason = copy.consumables.reasons[pull.reason ?? 'not_read'];
    titles.set(reason, [...(titles.get(reason) ?? []), pullTitle(pull.name, pull.attempt)]);
  }
  return [...titles.entries()].map(([reason, names]) => unknownPullLine(listWords(names), reason));
}

export interface RaiderDeathsModel {
  key: string;
  name: string;
  class: string;
  /** Each death of theirs in the night: "Deep Breath, Onyxia, pull 1 at 1:45". */
  deaths: string;
  /** "2 on wipes · 1 on a kill" past one death; '' for one, which its line already says. */
  count: string;
}

/** The officer's deaths by raider, in the order the API sends them (the night's order, never worst
 *  first): each death with what killed them, which pull and when, and a count past one. */
export function raiderDeaths(report: RaidReport): RaiderDeathsModel[] {
  return report.deaths.map((row) => {
    const lines = report.pulls.flatMap((pull) =>
      (pull.death_causes ?? [])
        .filter((death) => death.character_key === row.character_key)
        .map((death) =>
          copy.deaths.line(
            death.killing_blow === null ? copy.deaths.unknown : abilityName(death.killing_blow),
            pullTitle(pull.name, pull.attempt),
            killTimeLabel(death.at_ms),
          ),
        ),
    );
    const total = row.on_wipes + row.on_kills;
    return {
      key: row.character_key === '' ? row.name : row.character_key,
      name: row.name,
      class: row.class,
      deaths: lines.join('; '),
      count: total > 1 ? copy.deaths.count(row.on_wipes, row.on_kills) : '',
    };
  });
}

/** What the Discord post carries, as plain lines (the embed's fields in order; it names nobody
 *  for dying: api/internal/notify/report_embed.go). */
export function discordPreview(title: string, report: RaidReport, night: RaidNight): string[] {
  const killed = report.kill_order.filter((kill) => kill.killed && kill.kill_ms !== null);
  const kills = killed
    .slice(0, DISCORD_MAX_KILLS)
    .map((kill) => `${kill.name} ${killTimeLabel(kill.kill_ms ?? 0)}`);
  if (killed.length > DISCORD_MAX_KILLS) kills.push(copy.discord.more(killed.length - DISCORD_MAX_KILLS));
  const facts = [
    `${report.headline.wipes} wipe${report.headline.wipes === 1 ? '' : 's'}`,
    night.duration_ms > 0 ? durationLabel(night.duration_ms) : null,
    cameRatio(night),
    deathsWord(report.headline.deaths),
  ].filter((part): part is string => part !== null);
  const top = report.parses
    .filter((parse) => parse.metric === 'dps')
    .sort((a, b) => b.value - a.value)
    .slice(0, DISCORD_TOP_OUTPUT)
    .map((parse) =>
      copy.discord.top(parse.name, outputFigure(parse.value), outputUnit(parse.metric), parse.encounter),
    );
  const awards = report.loot.flatMap((group) => group.awards);
  const loot = awards
    .slice(0, DISCORD_MAX_LOOT)
    .map((award) => copy.discord.lootLine(award.item_name, award.name));
  if (awards.length > DISCORD_MAX_LOOT) loot.push(copy.discord.more(awards.length - DISCORD_MAX_LOOT));
  return [
    copy.discord.title(title),
    kills.length === 0 ? copy.discord.noKills : kills.join(', '),
    copy.discord.line(facts),
    ...top,
    ...loot,
  ];
}
