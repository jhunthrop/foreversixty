<!-- web/src/components/report/QueriesView.svelte -->
<!-- The exact answer to anything. Nothing here runs until someone asks: opening the tab
     loads no engine and no events file; pressing Run loads both, once, and every query
     after that is answered from memory. The first run therefore says how long it took to
     load as well as how long the query took, so the two are never confused.

     Every step is awaited and slow -- a cold WebAssembly start, two to ten megabytes of
     Parquet, then the query -- which makes this the most race-prone panel on the page. A
     fight switch or a view change can land between pressing Run and the answer arriving,
     so an answer is painted only if the fight it was asked about is still the fight on
     screen and no newer run has started. src/lib/report/query.ts owns the other half of
     that discipline: one engine, one query at a time, and nothing left running after
     close().
     Drawn to round 5 (design/mocks/directions/report.py, b1-report-queries): the panel `Queries`
     with the starting points as chips, the SQL as literal code, Run and `{n} rows · {ms} ms`; then
     the panel `Result`, a StatTable under the query's own column names. The parquet's column
     list is the rail's (QueriesRail). -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import { splitUnitName } from '../../lib/characters';
  import { sharedQueryLayer } from '../../lib/report/exact';
  import { classColorVar } from '../../lib/report/format';
  import { untrack } from 'svelte';
  import { eventsUrl as eventsUrlFor } from '../../lib/report/load';
  import { QUERY_ABANDONED, QUERY_TEMPLATES, withWindow, type QueryResult } from '../../lib/report/query';
  import type { TimeWindow } from '../../lib/report/window';
  import { FIELD_CONTROL, FIELD_FOCUS, fieldEdge } from '../../lib/ui/field';
  import type { StatColumn, StatFigure } from '../../lib/ui/stat-table';
  import Button from '../ui/Button.svelte';
  import Chips from '../ui/Chips.svelte';
  import CopyCsv from './CopyCsv.svelte';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import { grouped, upTo } from '../../lib/format/number';

  let {
    dataBaseUrl,
    fightIndex,
    window: current,
    classByName = new Map<string, string>(),
  }: {
    dataBaseUrl: string;
    fightIndex: number;
    window: TimeWindow;
    /** A unit's full log name to its class: a player's name in a result reads as on every table. */
    classByName?: ReadonlyMap<string, string>;
  } = $props();

  const LOADING_NOTE = 'Loading this fight’s events.';
  const RUNNING_NOTE = 'Running.';
  const FAILED_NOTE = 'That query did not run.';

  // One engine for the page, shared with the tables' exact measure (exact.ts): a second
  // DuckDB worker would be a second thirty-megabyte heap for the same fight.
  const layer = sharedQueryLayer();

  /** The file every run reads, and the identity a settled run is checked against. */
  const eventsUrl = $derived(eventsUrlFor(dataBaseUrl, fightIndex));

  // untrack for the reason ReportView.svelte untracks inlineMeta: this is a one-shot
  // starting value, not a binding, and reading a prop straight into $state is what
  // state_referenced_locally warns about. Picking a template re-reads the window.
  let sql = $state(untrack(() => withWindow(QUERY_TEMPLATES[0].sql, current)));
  /** The starting point on screen; cleared when the window moves, so picking it again rewrites
   *  its bounds (a Segmented reports only a change). */
  let template = $state<string | null>(QUERY_TEMPLATES[0].id);
  let windowKey = untrack(() => `${current.startMs}-${current.endMs}`);
  $effect(() => {
    const key = `${current.startMs}-${current.endMs}`;
    if (key === windowKey) return;
    windowKey = key;
    template = null;
  });
  const templateChips = $derived(
    QUERY_TEMPLATES.map((one) => ({ id: one.id, label: one.label, on: one.id === template })),
  );
  let result = $state<QueryResult | null>(null);
  let running = $state(false);
  let error = $state('');
  /** The one live region: what a screen reader hears while a slow run is happening. */
  let note = $state('');
  /** Bumped by every run, so an older answer cannot overwrite a newer one. */
  let token = 0;

  const rowsShown = $derived(result?.rows.length ?? 0);
  const rowsFound = $derived(result?.total ?? rowsShown);

  // A result belongs to the fight it was asked about. Switching fight while nothing is in
  // flight leaves no request to guard, so the stale answer is cleared here instead --
  // otherwise fight 3's rows would sit under fight 1's name with nothing to say so.
  $effect(() => {
    void eventsUrl;
    result = null;
    error = '';
    note = '';
  });

  // Svelte tears every live effect down on destroy, so leaving the Queries view, switching
  // mode, or the island itself unmounting all reach this: the DuckDB worker is terminated
  // rather than left holding its WebAssembly heap for the rest of the session.
  // The engine outlives this view: the tables' exact measure may want it next, and a
  // fight switch re-registers the file. The island's unmount is the page's end anyway.

  function useTemplate(id: string): void {
    const found = QUERY_TEMPLATES.find((candidate) => candidate.id === id);
    if (found === undefined) return;
    sql = withWindow(found.sql, current);
  }

  function pick(id: string): void {
    template = id;
    useTemplate(id);
  }

  const cellText = (cell: unknown): string => (cell === null || cell === undefined ? '' : String(cell));
  /** An id or a timestamp reads as written; every other number with its thousands commas. */
  const RAW_COLUMN = /(_id|_guid|time_unix_nano|_flags|_raid)$/;
  /** DuckDB hands a HUGEINT sum over as a string of digits: a number all the same. */
  const NUMERIC = /^-?\d+(\.\d+)?$/;
  const isNumber = (cell: unknown): boolean =>
    typeof cell === 'number' || typeof cell === 'bigint' || (typeof cell === 'string' && NUMERIC.test(cell));
  /** A melee swing has no spell: the log leaves spell_name empty, the tables call it Melee. */
  function cellShown(cell: unknown, column: string): string {
    if (column === 'spell_name' && cellText(cell) === '') return 'Melee';
    if (isNumber(cell) && !RAW_COLUMN.test(column)) {
      const value = typeof cell === 'string' ? Number(cell) : (cell as number | bigint);
      return upTo(Number(value), 1);
    }
    return cellText(cell);
  }

  /**
   * The result on the report's table: the first column names the row; the other text columns
   * (a spell, a target, a miss type) join it on its line, so a word is never right-aligned like a
   * figure; the number columns are the figures. A one-column answer is a figure alone.
   */
  const shape = $derived.by(() => {
    const columns = result?.columns ?? [];
    const rows = result?.rows ?? [];
    const textual = columns.map(
      (column, index) =>
        RAW_COLUMN.test(column) || rows.some((row) => row[index] !== null && !isNumber(row[index])),
    );
    const lead = columns.length > 1 ? 0 : -1;
    const words = columns.map((_, index) => index).filter((index) => index !== lead && textual[index]);
    const figures = columns.map((_, index) => index).filter((index) => index !== lead && !textual[index]);
    return { columns, lead, words, figures };
  });
  const nameHead = $derived(
    shape.lead === -1 ? '' : [shape.lead, ...shape.words].map((index) => shape.columns[index]).join(' · '),
  );
  const resultColumns = $derived<StatColumn[]>(
    shape.figures.map((index) => ({
      id: `${index}-${shape.columns[index]}`,
      label: shape.columns[index] ?? '',
    })),
  );
  const figuresOf = (row: readonly unknown[]): StatFigure[] =>
    shape.figures.map((index) => ({ value: cellShown(row[index], shape.columns[index] ?? '') }));
  /** Below 1024 a row shows its first figure; the rest go on the name's line, each with its column. */
  const phoneLine = (row: readonly unknown[]): string =>
    shape.figures
      .slice(1)
      .map((index) => `${shape.columns[index]} ${cellShown(row[index], shape.columns[index] ?? '')}`)
      .join(' · ');
  const csvLines = (): string[][] =>
    result === null
      ? []
      : [[...result.columns], ...result.rows.map((row) => row.map((cell) => cellText(cell)))];

  async function run(): Promise<void> {
    // Both captured before the first await and re-checked after the last: the visitor can
    // change fight or press Run again while a query is in flight, and an answer for a
    // fight nobody is looking at must be dropped, not painted.
    const mine = (token += 1);
    const wanted = eventsUrl;
    const startedCold = !layer.ready;

    running = true;
    error = '';
    note = startedCold ? LOADING_NOTE : RUNNING_NOTE;
    try {
      const answer = await layer.run(wanted, sql);
      if (mine !== token || wanted !== eventsUrl) return;
      result = answer;
      note = '';
    } catch (thrown) {
      // The rejection path needs the same guard as the success path: a stale failure that
      // painted an alert would describe a fight that is no longer on screen.
      if (mine !== token || wanted !== eventsUrl) return;
      const message = thrown instanceof Error ? thrown.message : FAILED_NOTE;
      // Not a failure the visitor caused or can act on: the layer drops a run whose view
      // has already moved on, and there is nothing to report about it.
      if (message === QUERY_ABANDONED) return;
      result = null;
      note = '';
      error = message;
    } finally {
      if (mine === token) running = false;
    }
  }
</script>

<Section title="Queries" testid="queries-view" panel>
  <div class="flex flex-col gap-3">
    <Chips chips={templateChips} label="Starting points" onToggle={pick} testid="template" />
    <!-- Literal code, the page's one mono face. -->
    <textarea
      id="query-sql"
      aria-label="Query"
      class={[
        FIELD_CONTROL,
        FIELD_FOCUS,
        fieldEdge(false),
        'min-h-[172px] resize-y py-2 font-mono text-[13px] leading-5',
      ]}
      bind:value={sql}
      spellcheck="false"
      data-testid="query-sql"></textarea>

    <div class="flex flex-wrap items-center gap-x-[14px] gap-y-2">
      <Button tone="primary" busy={running} busyLabel="Running…" onclick={() => void run()} testid="query-run"
        >Run</Button
      >
      <!-- Always in the DOM, so the first slow run is announced rather than swallowed: a live
           region added at the same moment as its text is not reliably read out. Polite, not
           assertive -- a download finishing is worth hearing, not worth interrupting for. -->
      <p class="text-muted text-[13px] leading-5" role="status" aria-live="polite" data-testid="query-status">
        {#if note !== ''}
          {note}
        {:else if result !== null}
          <span class="tabular-nums">{grouped(rowsFound)}</span>
          {rowsFound === 1 ? 'row' : 'rows'}{#if rowsFound > rowsShown}
            · <span class="tabular-nums" data-testid="query-truncated">{rowsShown}</span> shown{/if}
          · <span class="tabular-nums">{Math.round(result.elapsedMs)}</span> ms
        {/if}
      </p>
    </div>

    {#if error !== ''}
      <p class="text-wipe text-[14px]" role="alert" data-testid="query-error">{error}</p>
    {/if}
  </div>
</Section>

{#if result !== null}
  <Section title="Result" testid="query-result-panel" panel flush>
    {#snippet head()}{#if rowsShown > 0}<CopyCsv lines={csvLines} inHead />{/if}{/snippet}
    {#if rowsShown === 0}
      <EmptyState inset message="No rows matched." testid="query-empty" />
    {:else}
      <div class="overflow-x-auto">
        <StatTable
          columns={resultColumns}
          {nameHead}
          bar={false}
          icons={false}
          dense
          label="Query result"
          testid="query-result"
        >
          {#each result.rows as row, index (index)}
            <StatRow figures={figuresOf(row)} bar={false}>
              {#snippet name()}{#if shape.lead !== -1}{@const first = cellText(
                    row[shape.lead],
                  )}{#if classByName.has(first)}<span style:color={classColorVar(classByName.get(first))}
                      >{splitUnitName(first).name}</span
                    >{:else}{first}{/if}{#each shape.words as index (index)}<span
                      class="text-text font-normal"
                      >{` · ${cellShown(row[index], shape.columns[index] ?? '')}`}</span
                    >{/each}{/if}{/snippet}
              {#snippet sub()}{#if shape.figures.length > 1}<span class="ui-st-phone">{phoneLine(row)}</span
                  >{/if}{/snippet}
            </StatRow>
          {/each}
        </StatTable>
      </div>
    {/if}
  </Section>
{/if}
