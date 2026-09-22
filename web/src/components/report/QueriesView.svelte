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
     close(). -->
<script lang="ts">
  import { sharedQueryLayer } from '../../lib/report/exact';
  import { untrack } from 'svelte';
  import { eventsUrl as eventsUrlFor } from '../../lib/report/load';
  import { formatDuration } from '../../lib/report/format';
  import { QUERY_ABANDONED, QUERY_TEMPLATES, withWindow, type QueryResult } from '../../lib/report/query';
  import type { TimeWindow } from '../../lib/report/window';
  import { BUSY_CLASS } from '../../lib/ui/busy';

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
    const template = QUERY_TEMPLATES.find((candidate) => candidate.id === id);
    if (template !== undefined) sql = withWindow(template.sql, current);
  }

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

<div class="flex flex-col gap-3" data-testid="queries-view">
  <p class="text-muted text-[13px]">
    SQL over this fight’s event file, in your browser. The file downloads the first time you run a query and
    nothing leaves the page. These figures are measured, not scaled: a query answers the window it asks about
    exactly, where the tables mark a rescaled split with a tilde.
  </p>

  <div class="flex flex-wrap gap-2" role="group" aria-label="Starting points">
    {#each QUERY_TEMPLATES as template (template.id)}
      <button
        type="button"
        class="border-line-soft rounded-control text-nav inline-flex min-h-11 items-center border px-3 text-left text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-9"
        onclick={() => useTemplate(template.id)}
        data-testid={`template-${template.id}`}
      >
        {template.label}
      </button>
    {/each}
  </div>

  <p class="text-muted text-[12px]" data-testid="query-window">
    A starting point is written for the window on screen,
    <span class="tabular font-mono">{formatDuration(current.startMs)}</span>
    to
    <span class="tabular font-mono">{formatDuration(current.endMs)}</span>. Pick one again after brushing the
    chart to move its bounds.
  </p>

  <label class="label text-muted" for="query-sql">Query</label>
  <textarea
    id="query-sql"
    class="border-line-warm bg-raised rounded-control text-text min-h-[140px] border p-3 font-mono text-[13px]"
    bind:value={sql}
    spellcheck="false"
    data-testid="query-sql"></textarea>

  <button
    type="button"
    class={`border-line-warm-strong rounded-control text-strong inline-flex h-11 w-fit items-center border px-4 text-[12px] font-bold tracking-[0.06em] uppercase ${running ? BUSY_CLASS : ''}`}
    onclick={() => void run()}
    disabled={running}
    aria-busy={running}
    data-testid="query-run"
  >
    Run
  </button>

  <!-- Always in the DOM, so the first slow run is announced rather than swallowed: a live
       region added at the same moment as its text is not reliably read out. Polite, not
       assertive -- a download finishing is worth hearing, not worth interrupting for. -->
  <p class="text-muted text-[13px]" role="status" aria-live="polite" data-testid="query-status">
    {#if note !== ''}
      {note}
    {:else if result !== null}
      <span class="tabular font-mono">{rowsFound}</span>
      {rowsFound === 1 ? 'row' : 'rows'} in
      <span class="tabular font-mono">{Math.round(result.elapsedMs)}</span> ms{wasCold
        ? '. That run included the one-time load.'
        : '.'}
    {/if}
  </p>

  {#if error !== ''}
    <p class="text-[14px]" role="alert" data-testid="query-error">{error}</p>
  {/if}

  {#if result !== null}
    {#if rowsShown === 0}
      <p class="text-muted text-[14px]" data-testid="query-empty">No rows matched.</p>
    {:else}
      <div class="overflow-x-auto">
        <table class="w-full text-[13px]" data-testid="query-result">
          <thead>
            <tr>
              {#each result.columns as column, index (index)}
                <th class="label text-muted border-line-soft border-b px-2 py-1 text-left">{column}</th>
              {/each}
            </tr>
          </thead>
          <tbody>
            {#each result.rows as row, index (index)}
              <tr class="border-line-soft border-b">
                {#each row as cell, cellIndex (cellIndex)}
                  <td class="tabular px-2 py-1 font-mono">{cell === null ? '' : String(cell)}</td>
                {/each}
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    {/if}
    {#if rowsFound > rowsShown}
      <p class="text-muted text-[12px]" data-testid="query-truncated">
        The first <span class="tabular font-mono">{rowsShown}</span> of
        <span class="tabular font-mono">{rowsFound}</span> rows are shown. Add a LIMIT or a GROUP BY to narrow it.
      </p>
    {/if}
  {/if}
</div>
