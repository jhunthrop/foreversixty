<!-- web/src/components/report/SummaryTab.svelte -->
<!-- The Summary tab's lead (logs spec 2026-11-07 6.2; directions b1): the Players panel, framed, over one StatTable, a row
     per player in the scope: the crest, the name in its class colour (the link that opens the
     player's own view), "{spec} {class} · item level {n}", then each amount beside its rate
     (MUST 1: Damage and DPS, Healing and HPS, Taken and Taken / s; a pair whose rate is 0.0 on
     every row is hidden), Active and Deaths. Parse leads the figures on a boss pull. Sorted by
     Damage; every head sorts.

     `roster[].activity_pct` is not rescoped by a brushed window (lib/report/window.ts), so under
     one Active carries the whole-fight mark and a line says why. Damage, healing, taken and deaths
     are the window's own. -->
<script lang="ts">
  import { unitLabel } from '../../lib/report/names.svelte';
  import { specCell, type UnitClass } from '../../lib/report/roster-class';
  import { classColorVar, formatAmount, wholeFightMark } from '../../lib/report/format';
  import { formatRate } from '../../lib/report/frame';
  import type { Placement } from '../../lib/report/percentile';
  import type { RosterRow, Summary } from '../../lib/report/types';
  import type { StatColumn, StatFigure, StatSort } from '../../lib/ui/stat-table';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import Meter from '../ui/Meter.svelte';
  import { roleIcon, roleOf } from '../../lib/report/role-icon';
  import ClassIcon from './ClassIcon.svelte';

  let {
    summary,
    durationMs,
    percentiles = new Map<string, Placement>(),
    parseFallback = '',
    showParse = false,
    approximate = false,
    unitClasses = new Map<string, UnitClass>(),
    playerHref,
    onSelectPlayer,
  }: {
    summary: Summary;
    durationMs: number;
    percentiles?: ReadonlyMap<string, Placement>;
    /** What an empty Parse cell shows: a dash on a kill not ranked yet, '…' while asking, '?' unreachable. */
    parseFallback?: string;
    /** A boss pull: the Parse column leads the figures. */
    showParse?: boolean;
    approximate?: boolean;
    /** Class and where it came from per GUID, off the report's units: what a row without a spec shows. */
    unitClasses?: ReadonlyMap<string, UnitClass>;
    /** The url of one player's view (Source = that player). */
    playerHref: (guid: string) => string;
    onSelectPlayer: (guid: string) => void;
  } = $props();

  interface Pair {
    amount: string;
    rate: string;
    amountOf: (row: RosterRow) => number;
  }
  const PAIRS: readonly Pair[] = [
    { amount: 'Damage', rate: 'DPS', amountOf: (row) => row.damage_done },
    { amount: 'Healing', rate: 'HPS', amountOf: (row) => row.healing_done },
    { amount: 'Taken', rate: 'Taken / s', amountOf: (row) => row.damage_taken },
  ];
  /** A crowd (the sample's 265-player pull) paints its biggest rows first; the rest is one press away. */
  const CROWD = 100;
  const FIRST_ROWS = 50;

  const seconds = $derived(Math.max(durationMs, 1) / 1000);
  const rateOf = (pair: Pair, row: RosterRow): number => pair.amountOf(row) / seconds;
  /** MUST 1: a pair whose rate reads 0.0 on every row says nothing, amount and rate both. */
  const pairs = $derived(
    PAIRS.filter((pair) => summary.roster.some((row) => formatRate(rateOf(pair, row)) !== '0.0')),
  );
  const columns = $derived<StatColumn[]>([
    ...(showParse ? [{ id: 'parse', label: 'Parse', sortable: true }] : []),
    ...pairs.flatMap((pair) => [
      { id: pair.amount, label: pair.amount, sortable: true },
      { id: pair.rate, label: pair.rate, sortable: true },
    ]),
    { id: 'active', label: 'Active', sortable: true },
    { id: 'deaths', label: 'Deaths', sortable: true },
  ]);

  let sort = $state<StatSort>({ column: 'Damage', direction: 'desc' });
  function onsort(column: string): void {
    sort =
      sort.column === column
        ? { column, direction: sort.direction === 'desc' ? 'asc' : 'desc' }
        : { column, direction: 'desc' };
  }
  function sortValue(row: RosterRow, column: string): number {
    if (column === 'parse') return percentiles.get(row.guid)?.percentile ?? -1;
    if (column === 'active') return row.activity_pct;
    if (column === 'deaths') return row.deaths;
    const pair = PAIRS.find((entry) => entry.amount === column || entry.rate === column);
    return pair === undefined ? 0 : pair.amountOf(row);
  }
  const rows = $derived.by(() => {
    const sign = sort.direction === 'desc' ? -1 : 1;
    return [...summary.roster].sort(
      (a, b) =>
        sign * (sortValue(a, sort.column) - sortValue(b, sort.column)) || a.name.localeCompare(b.name),
    );
  });
  let showEveryone = $state(false);
  const shown = $derived(rows.length > CROWD && !showEveryone ? rows.slice(0, FIRST_ROWS) : rows);
  /** The Damage column's share bar runs against the top damage, as Damage done's does. */
  const topDamage = $derived(Math.max(1, ...rows.map((row) => row.damage_done)));
  const activeMark = $derived(wholeFightMark(approximate));

  function figures(row: RosterRow): StatFigure[] {
    const placed = percentiles.get(row.guid);
    const parse: StatFigure[] = showParse
      ? [
          placed === undefined
            ? { value: parseFallback || '–', muted: true }
            : { value: String(Math.floor(placed.percentile)) },
        ]
      : [];
    return [
      ...parse,
      ...pairs.flatMap((pair) => [
        { value: formatAmount(pair.amountOf(row)) },
        { value: formatRate(rateOf(pair, row)) },
      ]),
      { value: `${activeMark}${Math.round(row.activity_pct)}%` },
      { value: String(row.deaths), muted: row.deaths === 0 },
    ];
  }

  function subOf(row: RosterRow): string {
    const cell = specCell(row, unitClasses);
    const who = row.spec && cell.class ? `${row.spec} ${cell.class}` : cell.text;
    return [who === '—' ? '' : who, row.item_level ? `item level ${row.item_level}` : '']
      .filter((part) => part !== '')
      .join(' · ');
  }

  function follow(event: MouseEvent, guid: string): void {
    if (event.metaKey || event.ctrlKey || event.shiftKey || event.button !== 0) return;
    event.preventDefault();
    onSelectPlayer(guid);
  }
</script>

<Section title="Players" panel frame="framed" flush testid="summary-tab">
  <!-- Up to eight figure columns: the name keeps a floor so it never folds letter by letter in
       the narrow main column (1024 to 1279); past that floor the table scrolls sideways. -->
  <div class="summary-roster" data-testid="summary-roster-card">
    <StatTable {columns} nameHead="Player" {sort} {onsort} label="Players">
      {#each shown as row (row.guid)}
        {@const display = unitLabel(row.name)}
        {@const cell = specCell(row, unitClasses)}
        {@const role = roleOf(row.role)}
        <StatRow
          figures={figures(row)}
          phone={`damage · ${formatRate(row.damage_done / seconds)} DPS`}
          testid={`roster-${row.guid}`}
        >
          {#snippet icon()}<ClassIcon className={cell.class} size={28} />{/snippet}
          {#snippet name()}<a
              href={playerHref(row.guid)}
              class="font-bold hover:underline [@media(pointer:coarse)]:-my-[13px] [@media(pointer:coarse)]:inline-block [@media(pointer:coarse)]:py-[13px]"
              style:color={classColorVar(cell.class)}
              title={display.title}
              data-testid="roster-name"
              onclick={(event) => follow(event, row.guid)}>{display.name}</a
            >{#if display.realm !== ''}<span class="text-muted ml-1 text-[12px]">· {display.realm}</span
              >{/if}{/snippet}
          {#snippet sub()}{#if role !== null}{@const icon = roleIcon(role)}<img
                src={icon.url}
                alt={icon.label}
                title={icon.label}
                width="14"
                height="14"
                class="mr-1 inline-block rounded-[2px] align-[-2px]"
                data-testid="roster-role"
              />{/if}{subOf(row) || ' '}{/snippet}
          {#snippet meter()}<Meter
              value={row.damage_done / topDamage}
              tone="class"
              color={classColorVar(cell.class)}
            />{/snippet}
        </StatRow>
      {/each}
    </StatTable>
    {#if shown.length < rows.length}
      <button
        type="button"
        class="text-nav hover:text-strong min-h-11 px-4 text-left text-[14px]"
        data-testid="roster-show-all"
        onclick={() => (showEveryone = true)}>Show all {rows.length} players</button
      >
    {/if}
  </div>
  {#if approximate}
    <p class="text-small text-muted px-5 pt-2" data-testid="summary-wholefight-note">
      Active is marked {activeMark} because the summary does not track it over time: it is the whole fight's figure
      even inside a shorter window. Damage, healing, taken and deaths are this window's own numbers.
    </p>
  {/if}
</Section>

<style>
  .summary-roster {
    overflow-x: auto;
  }
  .summary-roster :global(.ui-st) {
    --st-name: minmax(120px, 2fr);
    --st-bar: minmax(40px, 56px);
    --st-fig: minmax(max-content, 54px);
    min-width: min-content;
  }
  @media (max-width: 1023px) {
    .summary-roster :global(.ui-st) {
      min-width: 0;
    }
  }
</style>
