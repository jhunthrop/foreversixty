<!-- web/src/components/report/ReportView.svelte -->
<!-- The report island's root. It owns three things and delegates everything else:
       the URL state, which is the page's whole state (src/lib/report/url.ts)
       the data, which is one meta fetch, one report.json and one summary per fight
       the layout, which is a selector column beside a content column on desktop and one
       stacked column on phone.
     Tasks 10 to 17 add panels inside the content column. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import {
    REPORT_LOAD_FAILED,
    fetchAccessUrl,
    fetchReportFile,
    fetchReportMeta,
    fetchSummary,
  } from '../../lib/report/load';
  import { resolveFightIndex } from '../../lib/report/fights';
  import { formatDuration } from '../../lib/report/format';
  import {
    defaultState,
    parseReportState,
    reportSearch,
    withState,
    type ReportState,
  } from '../../lib/report/url';
  import type { FightEntry, ReportFile, ReportMeta, Summary } from '../../lib/report/types';
  import FightSelector from './FightSelector.svelte';
  import ModeBar from './ModeBar.svelte';

  let { reportId, inlineMeta = null }: { reportId: string; inlineMeta?: ReportMeta | null } = $props();

  // untrack because inlineMeta is a one-shot bootstrap, not a binding: the shell renders
  // it once into data-report and never changes it, and reading a prop straight into
  // $state is the pattern Svelte warns about (state_referenced_locally).
  let meta = $state<ReportMeta | null>(untrack(() => inlineMeta));
  let file = $state<ReportFile | null>(null);
  let summary = $state<Summary | null>(null);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let error = $state('');
  /** Where report.json and the fight files come from: /logs-data/… or a signed url. */
  let dataBase = $state('');
  // Read only by fight loads and never rendered, so SvelteMap's per-key tracking would be
  // machinery for nothing: what the view re-reads is `summary`, which is $state.
  // eslint-disable-next-line svelte/prefer-svelte-reactivity
  const summaries = new Map<number, Summary>();

  const fights = $derived<FightEntry[]>(meta?.fights ?? file?.fights ?? []);
  const firstFight = $derived(fights.length > 0 ? fights[0].index : 1);
  let state = $state<ReportState>(defaultState(1));
  const fight = $derived<FightEntry | null>(fights.find((f) => f.index === state.fight) ?? null);
  const roster = $derived(
    (summary?.roster ?? []).map((row) => ({ guid: row.guid, name: row.name, class: row.class })),
  );

  function readUrl(): void {
    const parsed = parseReportState(window.location.search, firstFight);
    // parseReportState cannot check ?fight= against the report -- report.json has not
    // loaded when the url is parsed -- so `?fight=0` (fight_index is 1-based), a number
    // past the end and a gap all reach here. Showing the first fight beats an empty page
    // and a 404 on fights/0/summary.json; the window goes with it, because milliseconds
    // from a fight that does not exist mean nothing.
    const resolved = resolveFightIndex(fights, parsed.fight, firstFight);
    state =
      resolved === parsed.fight ? parsed : withState(parsed, { fight: resolved, start: null, end: null });
  }

  function writeUrl(): void {
    const search = reportSearch(state, firstFight);
    // replaceState, not pushState: brushing the chart and flicking between tabs would
    // otherwise bury the visitor's real back destination under a hundred entries. A link
    // shared from the address bar still carries the whole state.
    window.history.replaceState(null, '', `${window.location.pathname}${search}`);
  }

  function patch(next: Partial<ReportState>): void {
    // Changing fight drops the window: milliseconds from one fight's start mean nothing in
    // another's.
    const clearsWindow = next.fight !== undefined && next.fight !== state.fight;
    state = withState(state, clearsWindow ? { ...next, start: null, end: null } : next);
    writeUrl();
  }

  async function loadReport(): Promise<void> {
    status = 'loading';
    error = '';
    try {
      // inlineMeta, not `meta`: reading the state this function also writes would make it
      // a dependency of the effect below and run the whole load a second time.
      const resolved = inlineMeta ?? (await fetchReportMeta(reportId));
      meta = resolved;
      let base = resolved.data_base_url;
      // Private and guild reports are refused at /logs-data/, so they carry a signed base
      // url the API hands out to people who may read them.
      if (resolved.visibility === 'private' || resolved.visibility === 'guild') {
        base = await fetchAccessUrl(reportId);
      }
      dataBase = base;
      file = await fetchReportFile(base);
      readUrl();
      // A report with no fights at all has nothing to summarise, and asking for
      // fights/0/summary.json would fail the whole page over a file that cannot exist.
      if (fight !== null) await loadFight(state.fight);
      status = 'ready';
    } catch (thrown) {
      status = 'failed';
      error = thrown instanceof Error ? thrown.message : REPORT_LOAD_FAILED;
    }
  }

  /**
   * Picking a fight does not cancel the one before it, so two of these can be in flight
   * and settle in either order. `index === state.fight` is the whole guard: an answer for
   * a fight nobody is looking at is cached and otherwise dropped, rather than painted
   * under the selected fight's name. Cancelling the request instead would buy nothing --
   * the answer is worth keeping, it is only the assignment that is wrong.
   */
  async function loadFight(index: number): Promise<void> {
    const cached = summaries.get(index);
    if (cached !== undefined) {
      summary = cached;
      return;
    }
    const loaded = await fetchSummary(dataBase, index);
    summaries.set(index, loaded);
    if (index === state.fight) summary = loaded;
  }

  $effect(() => {
    void loadReport();
  });

  // Re-read the URL when the visitor uses the browser's own back and forward.
  $effect(() => {
    const onPop = (): void => {
      readUrl();
      void loadFight(state.fight).catch(() => {});
    };
    window.addEventListener('popstate', onPop);
    return () => window.removeEventListener('popstate', onPop);
  });

  // One fetch per fight, and only when the selection actually moves.
  $effect(() => {
    const wanted = state.fight;
    if (status !== 'ready' || dataBase === '') return;
    if (fight === null) return;
    if (summary?.fight_index === wanted) return;
    // Cleared on success as well as set on failure: an alert left over from the fight
    // before this one would describe the wrong fight, which is the same lie in reverse.
    // Both handlers re-check the selection for the reason loadFight does -- a stale
    // rejection must not fail the fight on screen, and a stale success must not retract
    // the current fight's alert.
    void loadFight(wanted).then(
      () => {
        if (wanted === state.fight) error = '';
      },
      (thrown: unknown) => {
        if (wanted !== state.fight) return;
        error = thrown instanceof Error ? thrown.message : REPORT_LOAD_FAILED;
      },
    );
  });
</script>

{#if status === 'failed'}
  <p class="px-[18px] text-[14px] md:px-0" role="alert" data-testid="report-error">{error}</p>
{:else if status === 'loading' || meta === null}
  <p class="text-muted px-[18px] text-[14px] md:px-0">Loading the report.</p>
{:else}
  <header class="flex flex-col gap-1 px-[18px] md:px-0">
    <h1 class="section-title text-[18px]" data-testid="report-title">
      {meta.title === '' ? meta.zone : meta.title}
    </h1>
    <p class="text-muted text-[13px]" data-testid="report-subtitle">
      {meta.zone} · <span class="tabular font-mono">{fights.length} fights</span> · {meta.status}
      {#if fight}· {fight.name}
        <span class="tabular font-mono">{formatDuration(fight.duration_ms)}</span>{/if}
    </p>
  </header>

  {#if error !== ''}
    <!-- The fight selector and the url have already moved by the time a fight's summary
         fails, so without this the previous fight's numbers sit under the new fight's
         label. A live report whose next fight is not written yet is the ordinary way to
         reach it. One line until a later task owns a real error panel. -->
    <p class="text-muted px-[18px] text-[13px] md:px-0" role="alert" data-testid="report-fight-error">
      {error}
    </p>
  {/if}

  <div class="grid grid-cols-1 gap-[22px] px-[18px] md:grid-cols-[280px_minmax(0,1fr)] md:gap-8 md:px-0">
    <FightSelector {fights} selected={state.fight} onSelect={(index) => patch({ fight: index })} />

    <div class="flex min-w-0 flex-col gap-[22px] md:gap-6">
      <ModeBar {state} {roster} onPatch={patch} />
      <!-- Task 10 inserts TimeChart here; Tasks 11 to 17 insert the panels below it. -->
      <p class="text-muted text-[14px]" data-testid="report-placeholder">
        {summary === null ? 'Loading the fight.' : `${summary.roster.length} players in this fight.`}
      </p>
    </div>
  </div>
{/if}
