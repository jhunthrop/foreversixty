<!-- web/src/components/Rankings.svelte -->
<!-- One encounter's rankings. Character boards and guild boards share the filter bar and
     the page's shape; only the row differs, which is why they are one component.
     Every filter is in the URL (src/lib/rankings/url.ts), so a filtered board is a link. -->
<script lang="ts">
  import {
    REGIONS,
    RULESETS,
    characterHref,
    guildHref,
    parseCharacterKey,
    rulesetLabel,
    splitUnitName,
  } from '../lib/characters';
  import { readCurrent } from '../lib/current-character';
  import { SECONDARY_BUTTON, SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import {
    classColorVar,
    formatAmount,
    formatDuration,
    percentileToken,
    rowLink,
  } from '../lib/report/format';
  import { titleize } from '../lib/report/og-meta';
  import {
    fetchEncounters,
    fetchGuildRankings,
    fetchRankings,
    type EncounterOption,
    type GuildRankingRow,
    type RankingRow,
    type RankingsPage,
  } from '../lib/rankings/api';
  import { encounterPickerCopy, rankingsEmptyCopy } from '../lib/rankings/copy';
  import {
    applyCurrentCharacterPrefilter,
    pinCurrentCharacterRow,
  } from '../lib/rankings/current-character-prefilter';
  import { RANKINGS_LOADING_MIN_H, RANKING_ROW_GRID } from '../lib/rankings/layout';
  import { PHASES } from '../lib/rankings/phases';
  import {
    FACTIONS,
    GUILD_KINDS,
    RANKING_METRICS,
    parseRankingsState,
    rankingsSearch,
    requiresEncounter,
    type RankingsState,
  } from '../lib/rankings/url';
  import { simCopy } from '../lib/sim/copy';
  import { executionHref, executionLabel, executionTitle } from '../lib/sim/execution';
  import CurrentCharacterBar from './CurrentCharacterBar.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  // The prerendered fixture page passes the slug; the Worker-served shell has none, so
  // the island reads it out of the path. One component, both routes.
  let { slug = '' }: { slug?: string } = $props();

  const resolvedSlug = $derived(
    slug !== ''
      ? slug
      : typeof window === 'undefined'
        ? ''
        : (/^\/rankings\/([a-z0-9-]{1,64})\/?$/.exec(window.location.pathname)?.[1] ?? ''),
  );
  // Bare /rankings names no encounter yet, so the page says what it is rather than
  // rendering an empty heading while the picker below offers what does exist.
  const encounter = $derived(resolvedSlug === '' ? 'Rankings' : titleize(resolvedSlug));
  let state = $state<RankingsState>(
    typeof window === 'undefined'
      ? parseRankingsState('')
      : applyCurrentCharacterPrefilter(
          parseRankingsState(window.location.search),
          window.location.search,
          readCurrent(),
        ),
  );
  let page = $state<RankingsPage | null>(null);
  const metricLabel = $derived(
    RANKING_METRICS.find((metric) => metric.id === state.metric)?.label ?? 'Value',
  );
  /** The pointer's own armory key at the time the board loaded: the row it names is pinned
   *  first by pinCurrentCharacterRow and marked "You" in the list. */
  let currentKey = $state<string | null>(null);
  let guildRows = $state<GuildRankingRow[]>([]);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let error = $state('');
  let attempt = $state(0);

  /**
   * A bare /rankings with no encounter in the URL cannot ask GET /v1/rankings or
   * /v1/rankings/guilds for a board that needs one (requiresEncounter) -- there is
   * nothing to name. Rather than firing that request and rendering its 400, the board
   * this filter combination would need is swapped for the encounter picker below.
   */
  const needsPicker = $derived(resolvedSlug === '' && requiresEncounter(state));
  let encounters = $state<EncounterOption[]>([]);
  let encountersStatus = $state<'idle' | 'loading' | 'ready' | 'failed'>('idle');

  async function loadEncounters(): Promise<void> {
    encountersStatus = 'loading';
    try {
      const result = await fetchEncounters();
      encounters = result.rows;
      encountersStatus = 'ready';
    } catch {
      encountersStatus = 'failed';
    }
  }

  const select = 'border-line-warm bg-raised rounded-control text-text h-11 px-2 text-[13px] md:h-9';

  function patch(next: Partial<RankingsState>): void {
    // Any filter change returns to page one: page three of a different filter is nothing.
    state = { ...state, ...next, page: next.page ?? 1 };
    window.history.replaceState(null, '', `${window.location.pathname}${rankingsSearch(state)}`);
  }

  /**
   * Every filter lives in the URL, so changing one always produces a brand new `state`
   * object (the immutable update in `patch` above) rather than mutating the one in flight.
   * That gives the effect below a cheap, exact staleness check: capture the object this
   * run is fetching for, and before any write to shared state -- on the success path and
   * on the rejection path alike -- compare it against the live `state` by reference. A
   * fetch started for a filter combination nobody is looking at any more is dropped on
   * both paths instead of painting its answer under the current filters' name. This is the
   * same hazard report-tabs.spec.ts's `heldRoute()` and RankingsMode.svelte's own
   * `wantedFight`/`wantedMetric` guard cover for the report island; here the whole state
   * object stands in for "the fight and the metric" because every filter, not just two
   * fields, can change at once.
   */
  $effect(() => {
    void attempt;
    if (needsPicker) {
      // Nothing to fetch until an encounter is chosen: clear any board this filter
      // combination is not going to answer for, and read the picker's own list once.
      page = null;
      guildRows = [];
      status = 'ready';
      if (encountersStatus === 'idle') void loadEncounters();
      return;
    }
    const requested = state;
    status = 'loading';
    const load =
      requested.board === 'guild'
        ? fetchGuildRankings({ encounter: resolvedSlug, kind: requested.kind, phase: requested.phase }).then(
            (result) => {
              if (state !== requested) return;
              guildRows = result.rows;
              page = null;
              status = 'ready';
            },
          )
        : fetchRankings({
            encounter: resolvedSlug,
            metric: requested.metric,
            spec: requested.spec,
            class: requested.class,
            phase: requested.phase,
            region: requested.region,
            ruleset: requested.ruleset,
            faction: requested.faction,
            since: requested.since,
            page: requested.page,
          }).then((result) => {
            if (state !== requested) return;
            const current = readCurrent();
            currentKey = current?.source === 'armory' ? current.ref : null;
            page = { ...result, rows: pinCurrentCharacterRow(result.rows, current) };
            guildRows = [];
            status = 'ready';
          });

    void load.catch((thrown: unknown) => {
      if (state !== requested) return;
      status = 'failed';
      error = thrown instanceof Error ? thrown.message : 'Rankings did not load.';
    });
  });

  const lastPage = $derived(page === null ? 1 : Math.max(1, Math.ceil(page.total / page.per_page)));

  function plannerHref(row: RankingRow): string | null {
    return row.build_id === undefined ? null : `/b/${row.build_id}`;
  }

  /**
   * A row's own region and ruleset, read out of `player.key` rather than the row's guild --
   * an unguilded row still fought in a real region and ruleset, and defaulting them (the
   * way `us`/`normal` would) links a real player at someone else's server. Null when
   * `player.key` does not parse, rather than a guess: fix round 1 caught this pointing an
   * unguilded row at the wrong region.
   */
  function characterRowHref(row: RankingRow): string | null {
    const key = parseCharacterKey(row.player.key);
    return key === null ? null : characterHref(key.region, key.ruleset, row.player.name);
  }
</script>

<div class="flex flex-col gap-[22px] md:gap-6" data-testid="rankings" id="rankings">
  <CurrentCharacterBar spine />
  <header class="flex flex-col gap-1">
    <h1 class="section-title text-[18px]">{encounter}</h1>
    <p class="text-muted text-[13px]" data-testid="rankings-count">
      {#if page !== null}
        <span class="tabular font-mono">{page.total}</span> ranked kills · updated
        <span class="tabular font-mono">{page.updated_at.slice(0, 10)}</span>
      {:else if state.board === 'guild'}
        <span class="tabular font-mono">{guildRows.length}</span> guilds
      {/if}
    </p>
  </header>

  <div class="flex flex-wrap items-center gap-x-4 gap-y-2" data-testid="rankings-filters">
    <div role="tablist" aria-label="Board" class="flex items-center gap-1">
      {#each [{ id: 'character', label: 'Characters' }, { id: 'guild', label: 'Guilds' }] as board (board.id)}
        <button
          type="button"
          role="tab"
          class="{SECONDARY_BUTTON} px-3"
          class:border-line-warm-strong={state.board === board.id}
          class:border-line-soft={state.board !== board.id}
          class:text-strong={state.board === board.id}
          class:text-nav={state.board !== board.id}
          aria-selected={state.board === board.id}
          data-testid={`board-${board.id}`}
          onclick={() => patch({ board: board.id as RankingsState['board'] })}
        >
          {board.label}
        </button>
      {/each}
    </div>

    {#if state.board === 'character'}
      <label class="label text-muted flex items-center gap-2" for="rankings-metric">
        Metric
        <select
          id="rankings-metric"
          class={select}
          value={state.metric}
          data-testid="filter-metric"
          onchange={(event) => patch({ metric: (event.currentTarget as HTMLSelectElement).value })}
        >
          {#each RANKING_METRICS as metric (metric.id)}<option value={metric.id}>{metric.label}</option
            >{/each}
        </select>
      </label>
    {:else}
      <label class="label text-muted flex items-center gap-2" for="rankings-kind">
        Board
        <select
          id="rankings-kind"
          class={select}
          value={state.kind}
          onchange={(event) => patch({ kind: (event.currentTarget as HTMLSelectElement).value })}
        >
          {#each GUILD_KINDS as kind (kind.id)}<option value={kind.id}>{kind.label}</option>{/each}
        </select>
      </label>
    {/if}

    <label class="label text-muted flex items-center gap-2" for="rankings-ruleset">
      Ruleset
      <select
        id="rankings-ruleset"
        class={select}
        value={state.ruleset}
        data-testid="filter-ruleset"
        onchange={(event) => patch({ ruleset: (event.currentTarget as HTMLSelectElement).value })}
      >
        <option value="">Every ruleset</option>
        {#each RULESETS as ruleset (ruleset.id)}<option value={ruleset.id}>{ruleset.label}</option>{/each}
      </select>
    </label>

    <label class="label text-muted flex items-center gap-2" for="rankings-region">
      Region
      <select
        id="rankings-region"
        class={select}
        value={state.region}
        onchange={(event) => patch({ region: (event.currentTarget as HTMLSelectElement).value })}
      >
        <option value="">Every region</option>
        {#each REGIONS as region (region)}<option value={region}>{region.toUpperCase()}</option>{/each}
      </select>
    </label>

    <label class="label text-muted flex items-center gap-2" for="rankings-phase">
      Phase
      <select
        id="rankings-phase"
        class={select}
        value={state.phase}
        onchange={(event) => patch({ phase: (event.currentTarget as HTMLSelectElement).value })}
      >
        <option value="">Every phase</option>
        {#each PHASES as phase (phase.id)}<option value={phase.id}>{phase.label}</option>{/each}
      </select>
    </label>

    <label class="label text-muted flex items-center gap-2" for="rankings-faction">
      Faction
      <select
        id="rankings-faction"
        class={select}
        value={state.faction}
        onchange={(event) => patch({ faction: (event.currentTarget as HTMLSelectElement).value })}
      >
        <option value="">Both</option>
        {#each FACTIONS as faction (faction)}<option value={faction}
            >{faction === 'horde' ? 'Horde' : 'Alliance'}</option
          >{/each}
      </select>
    </label>

    <label class="flex min-h-11 items-center gap-2 text-[13px] md:min-h-0">
      <input
        type="checkbox"
        checked={state.since === 'today'}
        onchange={(event) =>
          patch({ since: (event.currentTarget as HTMLInputElement).checked ? 'today' : '' })}
      />
      Today only
    </label>
  </div>

  {#if needsPicker}
    <div data-testid="encounter-picker" class="flex flex-col gap-3">
      {#if encountersStatus === 'idle' || encountersStatus === 'loading'}
        <Skeleton lines={3} rowHeight="h-4" testid="rankings-picker-skeleton" />
      {:else if encountersStatus === 'failed'}
        <LoadError
          message={encounterPickerCopy.failed}
          onRetry={() => void loadEncounters()}
          testid="rankings-picker-error"
        />
      {:else if encounters.length === 0}
        <p class="text-muted text-[14px]" data-testid="rankings-no-encounters">
          {encounterPickerCopy.noneYet}
        </p>
        <p class="text-muted text-[14px]">
          {encounterPickerCopy.tryGuildsProgress}
          <button
            type="button"
            class={rowLink}
            data-testid="rankings-try-guilds"
            onclick={() => patch({ board: 'guild', kind: 'progress' })}
          >
            {state.board === 'guild' ? encounterPickerCopy.progressButton : encounterPickerCopy.guildsButton}
          </button>
        </p>
      {:else}
        <h2 class="label text-muted">{encounterPickerCopy.heading}</h2>
        <ul class="flex flex-col" data-testid="encounter-picker-rows">
          {#each encounters as option (option.id)}
            <li class="border-line-soft border-b py-2 text-[14px]">
              <a class={rowLink} href={`/rankings/${option.slug}${rankingsSearch(state)}`}>{option.name}</a>
            </li>
          {/each}
        </ul>
      {/if}
    </div>
  {:else if status === 'loading'}
    <Skeleton lines={8} rowHeight="h-11" minHeight={RANKINGS_LOADING_MIN_H} testid="rankings-skeleton" />
  {:else if status === 'failed'}
    <LoadError message={error} onRetry={() => (attempt += 1)} testid="rankings-error" />
  {:else if state.board === 'guild'}
    {#if guildRows.length === 0}
      <EmptyState
        message={rankingsEmptyCopy.message}
        action={{ label: rankingsEmptyCopy.action, href: rankingsEmptyCopy.href }}
        testid="rankings-empty"
      />
    {:else}
      <ul class="reveal flex flex-col" data-testid="guild-rows">
        {#each guildRows as row (`${row.rank}-${row.guild.region}-${row.guild.ruleset}-${row.guild.name}`)}
          <li
            class="border-line-soft grid min-h-11 grid-cols-[40px_minmax(0,1fr)_auto] items-center gap-3 border-b px-2 py-2 text-[14px]"
          >
            <span class="tabular font-mono text-[12px]">{row.rank}</span>
            <a class={rowLink} href={guildHref(row.guild.region, row.guild.ruleset, row.guild.name)}
              >{row.guild.name}</a
            >
            <span class="tabular text-right font-mono">{formatAmount(Math.round(row.value))}</span>
          </li>
        {/each}
      </ul>
    {/if}
  {:else if page === null || page.rows.length === 0}
    <EmptyState
      message={rankingsEmptyCopy.message}
      action={{ label: rankingsEmptyCopy.action, href: rankingsEmptyCopy.href }}
      testid="rankings-empty"
    />
  {:else}
    <!-- The column line: the board's one header, drawn only at md and up, where every
         cell has its own column. On a phone the row folds its figures into labelled lines
         of its own, so a header would name columns that are not there. -->
    <div
      class="{RANKING_ROW_GRID} label text-muted border-line-soft hidden gap-x-3 border-b px-2 py-1 md:grid"
      data-testid="ranking-columns"
    >
      <span>#</span>
      <span>Character</span>
      <span>Guild</span>
      <span class="text-right">Size</span>
      <span class="text-right">{metricLabel}</span>
      <span class="text-right">Executed</span>
      <span class="text-right">Date</span>
      <span class="text-right">Length</span>
      <span class="text-right">Build</span>
    </div>
    <ul class="reveal flex flex-col" data-testid="ranking-rows">
      {#each page.rows as row (`${row.report_id}-${row.fight_index}-${row.player.key}`)}
        {@const buildHref = plannerHref(row)}
        {@const characterLinkHref = characterRowHref(row)}
        <li
          class="{RANKING_ROW_GRID} border-line-soft grid min-h-11 items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px]"
          data-testid={`ranking-${row.rank}`}
        >
          <span
            class="tabular font-mono text-[12px]"
            style={`color: ${percentileToken(Math.max(0, 100 - ((row.rank - 1) / Math.max(page.total, 1)) * 100))}`}
          >
            {row.rank}
          </span>
          <span class="flex min-w-0 items-center gap-2">
            {#if characterLinkHref !== null}
              <a
                class="{rowLink} truncate font-semibold"
                style={`color: ${classColorVar(row.player.class)}`}
                href={characterLinkHref}
                data-testid="ranking-character"
              >
                {splitUnitName(row.player.name).name}
              </a>
            {:else}
              <!-- player.key did not parse into a region and ruleset: an unlinked name is
                   the legible failure, not a guessed link that would point at the wrong
                   character. -->
              <span class="truncate font-semibold" style={`color: ${classColorVar(row.player.class)}`}>
                {splitUnitName(row.player.name).name}
              </span>
            {/if}
            {#if currentKey !== null && row.player.key === currentKey}
              <span class="pill pill-sample shrink-0" data-testid="ranking-you">You</span>
            {/if}
          </span>
          <span class="text-muted truncate text-[13px]">
            {#if row.guild}
              <a class={rowLink} href={guildHref(row.guild.region, row.guild.ruleset, row.guild.name)}
                >{row.guild.name}</a
              >
              <span class="hidden md:inline">
                · {rulesetLabel(row.guild.ruleset)} {row.guild.region.toUpperCase()}</span
              >
            {/if}
          </span>
          <span class="text-muted tabular hidden text-right font-mono text-[13px] md:inline">{row.size}</span>
          <span class="tabular text-right font-mono">{formatAmount(Math.round(row.value))}</span>
          <!-- Desktop gets its own cell; on phone the same figure joins the row's own phone-only
               line below, because a ninth column at phone width would push the name to one word. -->
          {#if row.execution_score === null}
            <span
              class="text-muted tabular hidden text-right font-mono text-[13px] md:inline"
              title={executionTitle(null)}
              aria-label={executionTitle(null)}
              data-testid="ranking-execution">{executionLabel(null)}</span
            >
          {:else}
            <a
              class="tabular hidden min-h-11 items-center justify-end text-right font-mono text-[13px] md:inline-flex"
              href={executionHref(row.report_id, row.fight_index)}
              title={executionTitle(row.execution_score)}
              data-testid="ranking-execution">{executionLabel(row.execution_score)}</a
            >
          {/if}
          <span class="text-muted tabular hidden text-right font-mono text-[13px] md:inline">
            {row.fought_at.slice(0, 10)}
          </span>
          <span class="text-muted tabular hidden text-right font-mono text-[13px] md:inline">
            {formatDuration(row.duration_ms)}
          </span>
          <span class="flex items-center justify-end gap-2 text-[13px]">
            {#if buildHref !== null}
              <a class={rowLink} href={buildHref} data-testid="ranking-build">{row.talent_split}</a>
            {:else}
              <span class="text-muted tabular font-mono">{row.talent_split}</span>
            {/if}
            <a
              class={rowLink}
              href={`/reports/${row.report_id}?fight=${row.fight_index}`}
              data-testid="ranking-report">Report</a
            >
          </span>
          <span class="text-muted label col-span-full md:hidden" data-testid="ranking-execution-phone">
            {row.execution_score === null
              ? simCopy.executionUnscored
              : `${executionLabel(row.execution_score)} executed`}
          </span>
          {#if row.state !== 'ok'}
            <!-- Plain text, not a colour swatch: the row's own background or a coloured dot
                 would say nothing to a screen reader, and the design system's rule against
                 colour-only signalling applies to moderation exactly the way it applies
                 everywhere else. Read in order, a removed row says "...duration, removed"
                 -- part of the same sentence the rest of the row already reads as, not a
                 separate alert. -->
            <span class="pill pill-sample col-span-full" data-testid="ranking-state"
              >{row.state.replace('_', ' ')}</span
            >
          {/if}
        </li>
      {/each}
    </ul>

    <div class="flex items-center gap-3">
      <button
        type="button"
        class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3 disabled:opacity-50"
        disabled={state.page <= 1}
        data-testid="rankings-prev"
        onclick={() => patch({ page: state.page - 1 })}
      >
        Previous
      </button>
      <span class="text-muted tabular font-mono text-[13px]">Page {state.page} of {lastPage}</span>
      <button
        type="button"
        class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3 disabled:opacity-50"
        disabled={state.page >= lastPage}
        data-testid="rankings-next"
        onclick={() => patch({ page: state.page + 1 })}
      >
        Next
      </button>
    </div>
  {/if}
</div>
