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
  import type { FightEntry, ReportFile, ReportMeta, RosterRow, Summary } from '../../lib/report/types';
  import {
    clampWindow,
    combinedSeries,
    isFullWindow,
    scopeSummary,
    windowMs,
    windowOf,
    windowPresets,
    type TimeWindow,
  } from '../../lib/report/window';
  import ActorTable from './ActorTable.svelte';
  import AuraTable from './AuraTable.svelte';
  import CastTable from './CastTable.svelte';
  import DeathsTab from './DeathsTab.svelte';
  import ExchangeTable from './ExchangeTable.svelte';
  import FightSelector from './FightSelector.svelte';
  import FilterBar from './FilterBar.svelte';
  import ModeBar from './ModeBar.svelte';
  import ResourceGraphs from './ResourceGraphs.svelte';
  import SummaryTab from './SummaryTab.svelte';
  import ThreatTable from './ThreatTable.svelte';
  import TimeChart from './TimeChart.svelte';
  import {
    DEFAULT_FILTERS, applyActorFilters, bossGuids, playerGuids, type ReportFilters,
  } from '../../lib/report/filters';
  import { createPercentileLoader, percentileKey } from '../../lib/report/percentile';
  import activeBuild from '../../data/active-build.json';
  import { loadTalents } from '../../lib/planner/load';
  import type { TalentFile } from '../../lib/planner/types';

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

  // Named timeWindow, not window: a `const window` in a Svelte <script> shadows the
  // global one, and this component uses window.location, window.history and
  // window.addEventListener.
  const timeWindow = $derived(windowOf(state, summary?.duration_ms ?? 0));
  /** Every table below reads this, never `summary`: one rescope per window change. */
  const scoped = $derived(summary === null ? null : scopeSummary(summary, timeWindow));
  const presets = $derived(summary === null ? [] : windowPresets(summary));
  const chartSeries = $derived(combinedSeries(summary?.damage_done ?? []));
  const windowIsWhole = $derived(summary !== null && isFullWindow(timeWindow, summary.duration_ms));

  let filters = $state<ReportFilters>(DEFAULT_FILTERS);

  // Two independent sources of scaling on an Actor-shaped row: the window (its share of
  // the actor's series, window.ts) and a target or boss filter (its share of the actor's
  // targets, filters.ts's `targetShare`). Either one alone is enough to make the
  // per-ability and per-target splits approximate, so a whole-fight window with "Boss
  // damage only" engaged still needs the mark. Ability, players-only, count-overkill and
  // ignore-after-death do not scale anything -- they subset or add exact figures -- so
  // they are not part of this. Over-claiming (marking a row that individually happens to
  // be exact) is the safe direction here; under-claiming is not.
  const filtersScale = $derived(filters.target !== '' || filters.bossOnly);
  const actorTableApproximate = $derived(!windowIsWhole || filtersScale);
  let percentiles = $state(new Map<string, number>());
  const loader = createPercentileLoader();

  const filterContext = $derived({
    bosses: bossGuids(file?.units ?? [], fight?.name ?? ''),
    players: playerGuids(file?.units ?? []),
    deaths: scoped?.deaths ?? [],
  });

  /** GUID to class, for the tables whose rows are not Actors. */
  const classOf = $derived(
    new Map((scoped?.roster ?? []).filter((row) => row.class).map((row) => [row.guid, row.class as string])),
  );

  /**
   * Talents per tree, per class, fetched once per class the fight actually contains. The
   * planner already publishes these files under /data/<build>/talents/<class>.json, so the
   * report reuses them rather than shipping a second copy of the tree shapes.
   */
  // Replaced wholesale below, never keyed: SvelteMap's per-key tracking is machinery this
  // does not need.
  let treeSizes = $state(new Map<string, number[]>());

  $effect(() => {
    const classes = new Set(
      (summary?.roster ?? []).map((row) => row.class).filter((name): name is string => name !== undefined),
    );
    const wanted = [...classes].filter((name) => !treeSizes.has(name));
    if (wanted.length === 0) return;
    void Promise.all(
      wanted.map(async (name) => {
        const slug = name.toLowerCase().replace(/\s+/g, '-');
        try {
          const file: TalentFile = await loadTalents(activeBuild.build, slug);
          return [name, file.trees.map((tree) => tree.talents.length)] as const;
        } catch {
          // No talent data for this class in this build: the link falls back to gear only.
          return [name, []] as const;
        }
      }),
    ).then((entries) => {
      treeSizes = new Map([...treeSizes, ...entries]);
    });
  });

  const treeSizesFor = (className: string | undefined): number[] =>
    className === undefined ? [] : (treeSizes.get(className) ?? []);

  /** Which Actor[] the current tab shows, scoped to `state.source` and then filtered. */
  const tabActors = $derived.by(() => {
    if (scoped === null) return [];
    const source =
      state.tab === 'damage-done'
        ? scoped.damage_done
        : state.tab === 'damage-taken'
          ? scoped.damage_taken
          : state.tab === 'healing'
            ? scoped.healing
            : [];
    const bySource =
      state.source === 'friendlies'
        ? source.filter((actor) => filterContext.players.has(actor.guid))
        : state.source === 'enemies'
          ? source.filter((actor) => !filterContext.players.has(actor.guid))
          : source.filter((actor) => actor.guid === state.source);
    return applyActorFilters(bySource, filters, filterContext);
  });

  const metricLabel = $derived(state.tab === 'healing' ? 'Healing' : 'Damage');

  /**
   * The Summary tab's headline percentile is each row's own role metric -- DPS for a dps
   * row, HPS for a healer, damage taken for a tank (spec section 3: "Role metric (DPS,
   * HPS, damage taken for tanks)") -- not a blanket DPS for everyone on the product's
   * default landing tab. The Damage Done, Damage Taken and Healing tabs are unaffected:
   * they already rank every row by that table's own metric.
   */
  function roleMetric(row: RosterRow): { metric: string; value: number } {
    if (row.role === 'healer') return { metric: 'hps', value: row.hps };
    if (row.role === 'tank') return { metric: 'damage_taken', value: row.dtps };
    return { metric: 'dps', value: row.dps };
  }

  // Percentiles mean a fight's whole-fight role metric against the rankings, so they are
  // asked for only on an encounter kill at the full window. A brushed window's number is
  // not a parse, and saying otherwise would be worse than saying nothing.
  $effect(() => {
    const encounterId = fight?.encounter_id;
    // Captured now, checked when the request resolves: a fight or tab change while it is
    // in flight must not paint stale answers under the new selection, the same guard
    // loadFight uses on `index === state.fight`.
    const wantedFight = state.fight;
    const wantedTab = state.tab;
    if (summary === null || encounterId === undefined || !windowIsWhole) {
      percentiles = new Map();
      return;
    }
    const tab = state.tab;
    const phase = meta?.phase ?? 'launch';
    const difficulty = fight?.difficulty ?? 0;

    // One query per roster row that has a spec, kept beside its GUID so the answers can be
    // put back on the right rows. Summary picks each row's own role metric; the other
    // three tabs rank by that table's metric, exactly as before.
    const wanted = summary.roster
      .filter((row) => row.spec !== undefined && row.spec !== '')
      .map((row) => {
        const { metric, value } =
          tab === 'summary'
            ? roleMetric(row)
            : { metric: tab === 'healing' ? 'hps' : 'dps', value: tab === 'healing' ? row.hps : row.dps };
        return {
          guid: row.guid,
          query: { encounterId, difficulty, spec: row.spec ?? '', phase, metric, value: Math.round(value) },
        };
      });

    void loader.load(wanted.map((entry) => entry.query)).then((answers) => {
      if (state.fight !== wantedFight || state.tab !== wantedTab) return;
      // A plain Map, not SvelteMap: this is a throwaway local built up once and then
      // assigned whole to `percentiles` (already $state) below, the same reasoning the
      // `summaries` cache above gives for its own eslint-disable.
      // eslint-disable-next-line svelte/prefer-svelte-reactivity
      const next = new Map<string, number>();
      for (const entry of wanted) {
        const found = answers.get(percentileKey(entry.query));
        if (found !== undefined) next.set(entry.guid, found);
      }
      percentiles = next;
    });
  });

  function setWindow(next: TimeWindow | null): void {
    const duration = summary?.duration_ms ?? 0;
    // Clamped for the reason resolveFightIndex clamps the fight: a drag, a preset or a
    // pasted URL can land past either end. A window that covers the fight is no window at
    // all, so it leaves the URL rather than sitting there as start=0&end=<duration>.
    const clamped = next === null ? null : clampWindow(next, duration);
    patch(
      clamped === null || windowMs(clamped) === 0 || isFullWindow(clamped, duration)
        ? { start: null, end: null }
        : { start: clamped.startMs, end: clamped.endMs },
    );
  }

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
      <!-- Tasks 11 to 17 insert the panels below the chart. -->
      {#if summary !== null}
        <TimeChart
          series={chartSeries}
          durationMs={summary.duration_ms}
          window={timeWindow}
          deaths={summary.deaths.map((death) => ({ at_ms: death.at_ms, name: death.name }))}
          label="Damage"
          onWindow={setWindow}
        />
        <div class="flex flex-wrap gap-2" data-testid="window-presets">
          <!-- Keyed by position, not by label: a battle-rez puts the same name in
               `deaths` twice, and two buttons labelled "Before Thalgrit died" would be a
               duplicate key, which Svelte throws on rather than renders. The list is
               rebuilt wholesale whenever the fight changes, so position is stable. -->
          {#each presets as preset, position (position)}
            <button
              type="button"
              class="border-line-soft rounded-control text-nav inline-flex h-11 items-center border px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:h-9"
              onclick={() => setWindow(preset.window)}
            >
              {preset.label}
            </button>
          {/each}
        </div>
      {/if}
      {#if scoped !== null && state.mode === 'analyze' && state.view === 'tables'}
        {#if state.tab === 'summary'}
          <SummaryTab
            summary={scoped}
            durationMs={scoped.duration_ms}
            {percentiles}
            approximate={!windowIsWhole}
            dataBuild={activeBuild.build}
            {classOf}
            {treeSizesFor}
          />
        {:else if state.tab === 'damage-done' || state.tab === 'damage-taken' || state.tab === 'healing'}
          <FilterBar {filters} actors={tabActors} onChange={(next) => (filters = next)} />
          <ActorTable
            actors={tabActors}
            durationMs={scoped.duration_ms}
            {metricLabel}
            {percentiles}
            approximate={actorTableApproximate}
          />
          {#if actorTableApproximate}
            <p class="text-muted text-[12px]" data-testid="approximate-note">
              A tilde marks a figure split across abilities and targets in proportion to the window and
              any active target or boss filter. Totals and per-second figures are exact. Queries answers
              the split exactly.
            </p>
          {/if}
        {:else if state.tab === 'buffs'}
          <AuraTable tracks={scoped.auras} durationMs={scoped.duration_ms} kind="BUFF" />
        {:else if state.tab === 'debuffs'}
          <AuraTable tracks={scoped.auras} durationMs={scoped.duration_ms} kind="DEBUFF" />
        {:else if state.tab === 'casts'}
          <CastTable
            rows={scoped.casts}
            durationMs={scoped.duration_ms}
            startMs={timeWindow.startMs}
            {classOf}
            approximate={!windowIsWhole}
          />
        {:else if state.tab === 'interrupts'}
          <ExchangeTable rows={scoped.interrupts} emptyText="Nothing was interrupted in the whole fight." />
        {:else if state.tab === 'dispels'}
          <ExchangeTable rows={scoped.dispels} emptyText="Nothing was dispelled in the whole fight." />
        {:else if state.tab === 'resources'}
          <ResourceGraphs tracks={scoped.resources} durationMs={scoped.duration_ms} />
        {:else if state.tab === 'threat'}
          <ThreatTable rows={scoped.threat} {classOf} approximate={!windowIsWhole} />
        {:else if state.tab === 'deaths'}
          <DeathsTab
            deaths={scoped.deaths}
            durationMs={summary?.duration_ms ?? scoped.duration_ms}
            combatants={scoped.combatants}
            {classOf}
            dataBuild={activeBuild.build}
            {treeSizesFor}
            onWindow={setWindow}
          />
        {/if}
      {/if}
    </div>
  </div>
{/if}
