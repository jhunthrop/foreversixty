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
     Drawn to logs spec 2026-11-07 section 6.2 (`web-report-queries`): Section `Queries` holding
     the one Plate: the Starting points (Segmented), the SQL as literal code, the primary Run,
     `{n} rows`, then the result as a StatTable under the query's own column names. The parquet's
     column list is the rail's (QueriesRail). -->
<script lang="ts">
  import { sharedQueryLayer } from '../../lib/report/exact';
  import { untrack } from 'svelte';
  import { eventsUrl as eventsUrlFor } from '../../lib/report/load';
  import { formatDuration } from '../../lib/report/format';
  import { QUERY_ABANDONED, QUERY_TEMPLATES, withWindow, type QueryResult } from '../../lib/report/query';
  import type { TimeWindow } from '../../lib/report/window';
  import { BUSY_CLASS } from '../../lib/ui/busy';
  import { FIELD_CONTROL, FIELD_FOCUS, FIELD_LABEL, fieldEdge } from '../../lib/ui/field';
  import type { StatColumn, StatFigure } from '../../lib/ui/stat-table';
  import Button from '../ui/Button.svelte';
  import Section from '../ui/Section.svelte';
  import Segmented from '../ui/Segmented.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';

  let {
    dataBaseUrl,
    fightIndex,
    window: current,
  }: { dataBaseUrl: string; fightIndex: number; window: TimeWindow } = $props();

  const LOADING_NOTE =
    'Loading the query engine and this fight’s event file. A few megabytes, once per fight.';
  const RUNNING_NOTE = 'Running the query.';
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
  const templateOptions = QUERY_TEMPLATES.map((one) => ({ value: one.id, label: one.label }));
  let result = $state<QueryResult | null>(null);
  let running = $state(false);
  let error = $state('');
  /** Whether the answer on screen was paid for with the engine build and the download. */
  let wasCold = $state(false);
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
    wasCold = false;
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

  /** The first column is the row's name; a one-column answer (a count) is a figure alone. */
  const named = $derived((result?.columns.length ?? 0) > 1);
  const resultColumns = $derived<StatColumn[]>(
    (result?.columns ?? [])
      .slice(named ? 1 : 0)
      .map((column, index) => ({ id: `${index}-${column}`, label: column })),
  );
  const cellText = (cell: unknown): string => (cell === null || cell === undefined ? '' : String(cell));
  const figuresOf = (row: readonly unknown[]): StatFigure[] =>
    row.slice(named ? 1 : 0).map((cell) => ({ value: cellText(cell) }));

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
      wasCold = startedCold;
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

<Section title="Queries" testid="queries-view" panel frame="framed" flush>
  <div>
    <div class="flex flex-col gap-2 px-4 py-[14px]">
      <span class={FIELD_LABEL} id="query-starts">Starting points</span>
      <Segmented
        options={templateOptions}
        value={template}
        label="Starting points"
        size="sm"
        stack
        onChange={pick}
        testid="template"
      />
      <p class="text-muted text-[13px] leading-5" data-testid="query-window">
        A starting point is written for the window on screen,
        <span class="tabular-nums">{formatDuration(current.startMs)}</span>
        to
        <span class="tabular-nums">{formatDuration(current.endMs)}</span>. Pick one again after brushing the
        chart to move its bounds.
      </p>

      <label class={['mt-2', FIELD_LABEL]} for="query-sql">Query</label>
      <!-- Literal code, the page's one mono face. -->
      <textarea
        id="query-sql"
        class={[
          FIELD_CONTROL,
          FIELD_FOCUS,
          fieldEdge(false),
          'min-h-[140px] resize-y py-2 font-mono text-[13px] leading-5',
        ]}
        bind:value={sql}
        spellcheck="false"
        data-testid="query-sql"></textarea>

      <div class="flex flex-wrap items-center gap-3">
        <span class={running ? BUSY_CLASS : undefined} aria-busy={running}>
          <Button tone="primary" disabled={running} onclick={() => void run()} testid="query-run">Run</Button>
        </span>
        <!-- Always in the DOM, so the first slow run is announced rather than swallowed: a live
             region added at the same moment as its text is not reliably read out. Polite, not
             assertive -- a download finishing is worth hearing, not worth interrupting for. -->
        <p
          class="text-muted text-[13px] leading-5"
          role="status"
          aria-live="polite"
          data-testid="query-status"
        >
          {#if note !== ''}
            {note}
          {:else if result !== null}
            <span class="tabular-nums">{rowsFound}</span>
            {rowsFound === 1 ? 'row' : 'rows'} in
            <span class="tabular-nums">{Math.round(result.elapsedMs)}</span> ms{wasCold
              ? '. That run included the one-time load.'
              : '.'}
          {/if}
        </p>
      </div>

      {#if error !== ''}
        <p class="text-wipe text-[14px]" role="alert" data-testid="query-error">{error}</p>
      {/if}
    </div>

    {#if result !== null}
      {#if rowsShown === 0}
        <p class="text-muted px-4 pb-4 text-[14px]" data-testid="query-empty">No rows matched.</p>
      {:else}
        <div class="overflow-x-auto">
          <StatTable
            columns={resultColumns}
            nameHead={named ? (result.columns[0] ?? '') : ''}
            bar={false}
            label="Query result"
            testid="query-result"
          >
            {#each result.rows as row, index (index)}
              <StatRow figures={figuresOf(row)} bar={false}>
                {#snippet name()}{named ? cellText(row[0]) : ''}{/snippet}
              </StatRow>
            {/each}
          </StatTable>
        </div>
      {/if}
      {#if rowsFound > rowsShown}
        <p class="text-muted px-4 pb-4 text-[13px]" data-testid="query-truncated">
          The first <span class="tabular-nums">{rowsShown}</span> of
          <span class="tabular-nums">{rowsFound}</span> rows are shown. Add a LIMIT or a GROUP BY to narrow it.
        </p>
      {/if}
    {/if}
  </div>
  <p class="text-muted mt-3 text-[13px] leading-5">
    The file downloads the first time you run a query and nothing leaves the page. These figures are measured,
    not scaled: a query answers the window it asks about exactly, where the tables mark a rescaled split with
    a tilde.
  </p>
</Section>
