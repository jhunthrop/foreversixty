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
    POLL_INTERVAL_MS,
    REPORT_LOAD_FAILED,
    createPoller,
    fetchAccessUrl,
    fetchLive,
    fetchReportFile,
    fetchReportMeta,
    eventsUrl,
    fetchSummary,
    withFreshBase,
  } from '../../lib/report/load';
  import { defaultFightIndex, resolveFightIndex } from '../../lib/report/fights';
  import { aggregateNight, nightSummary, type Night } from '../../lib/report/night';
  import { formatDate, formatDuration, outcomeLabel } from '../../lib/report/format';
  import {
    ALL_FIGHTS,
    FLAG_LETTERS,
    type FlagKey,
    defaultState,
    parseReportState,
    reportSearch,
    withState,
    type ReportState,
  } from '../../lib/report/url';
  import type { Actor, FightEntry, ReportFile, ReportMeta, RosterRow, Summary } from '../../lib/report/types';
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
  import CompareMode from './CompareMode.svelte';
  import DeathsTab from './DeathsTab.svelte';
  import EventsView from './EventsView.svelte';
  import ExchangeTable from './ExchangeTable.svelte';
  import FightSelector from './FightSelector.svelte';
  import FilterBar from './FilterBar.svelte';
  import ModeBar from './ModeBar.svelte';
  import NightView from './NightView.svelte';
  import Glossary from './Glossary.svelte';
  import RaidCooldowns from './RaidCooldowns.svelte';
  import QueriesView from './QueriesView.svelte';
  import RankingsMode from './RankingsMode.svelte';
  import ResourceGraphs from './ResourceGraphs.svelte';
  import SummaryTab from './SummaryTab.svelte';
  import ThreatTable from './ThreatTable.svelte';
  import TimeChart from './TimeChart.svelte';
  import TimelinesView from './TimelinesView.svelte';
  import {
    applyActorFilters,
    bossGuids,
    bossGuidsOf,
    friendlyGuids,
    playerGuids,
    type ReportFilters,
  } from '../../lib/report/filters';
  import { createPercentileLoader, percentileKey, type Placement } from '../../lib/report/percentile';
  import activeBuild from '../../data/active-build.json';
  import { loadTalents } from '../../lib/planner/load';
  import { encounterSlug as slugFor } from '../../lib/rankings/api';
  import { phaseAt } from '../../lib/rankings/phases';
  import { inSource, scopeSource } from '../../lib/report/source';
  import {
    measureExact,
    measureTable,
    sharedQueryLayer,
    type ExactSplit,
    type ExactTotals,
    type TargetScope,
  } from '../../lib/report/exact';
  import { resolveTreeSizes } from '../../lib/report/tree-sizes';
  import { classSlugOf } from '../../lib/report/planner-link';

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
  // A fight whose summary is on its way. A second request for the same fight while the
  // first is in flight joins it instead of asking R2 again: the fight-selection effect
  // and the popstate handler can both ask for the fight on screen within one tick, and
  // each extra request is a round trip the visitor waits on for nothing.
  // eslint-disable-next-line svelte/prefer-svelte-reactivity
  const inflight = new Map<number, Promise<void>>();

  const fights = $derived<FightEntry[]>(meta?.fights ?? file?.fights ?? []);
  /** The fight a bare url opens on: the first boss pull, not the first trash. */
  const firstFight = $derived(defaultFightIndex(fights));
  /** The whole night rather than one fight: every boss pull folded together. */
  const nightMode = $derived(state.fight === ALL_FIGHTS);
  let night = $state<Night | null>(null);
  /** The night as one Summary-shaped object, for the fight tabs to show. */
  let nightFold = $state<Summary | null>(null);
  let nightLoading = $state(false);
  /** What the tables read: the whole night, or the selected fight. */
  const base = $derived<Summary | null>(nightMode ? nightFold : summary);
  let state = $state<ReportState>(defaultState(1));
  const fight = $derived<FightEntry | null>(fights.find((f) => f.index === state.fight) ?? null);
  /**
   * The encounter slug the rankings pages use, from the selected fight's own name. This
   * is `rankings/api.ts`'s one shared derivation -- Task 19's `/rankings/<encounter-slug>`
   * route uses the same function, so the two cannot drift apart on a name with punctuation.
   */
  const currentEncounterSlug = $derived(slugFor(fight?.name ?? ''));
  const roster = $derived(
    ((nightMode ? nightFold : summary)?.roster ?? []).map((row) => ({
      guid: row.guid,
      name: row.name,
      class: row.class,
    })),
  );

  // Named timeWindow, not window: a `const window` in a Svelte <script> shadows the
  // global one, and this component uses window.location, window.history and
  // window.addEventListener.
  const timeWindow = $derived(windowOf(state, base?.duration_ms ?? 0));
  /** The bosses' unit names over the night, for the debuff table's per-spell line. */
  const bossUnitNames = $derived.by(() => {
    const units = file?.units ?? [];
    const guids = bossGuidsOf(
      units,
      fights.filter((entry) => entry.kind === 'encounter').map((entry) => entry.name),
    );
    return new Set(units.filter((unit) => guids.has(unit.guid)).map((unit) => unit.name));
  });
  /** GUID to unit name from report.json, for the auras' casters. */
  const unitNames = $derived(new Map((file?.units ?? []).map((unit) => [unit.guid, unit.name])));
  /** The player GUIDs from report.json, for the source scope and the filters. */
  const playerSet = $derived(playerGuids(file?.units ?? []));
  /** The players and their pets and totems: what the enemies scope leaves out. */
  const friendlySet = $derived(friendlyGuids(file?.units ?? []));
  /**
   * Every table below reads this, never `summary`: one rescope per window change, then
   * the Source control's scope over it, so picking one player narrows the whole page.
   */
  /**
   * The exact split of one row inside the window, from the fight's events through the
   * shared DuckDB engine. One fight at a time: the night has no single event file.
   */
  const tableKind = $derived(
    state.tab === 'damage-taken' ? 'damage-taken' : state.tab === 'healing' ? 'healing' : 'damage-done',
  );
  /** A player's pets, to their owner, from the report's units: what a measure folds a pet's lines under. */
  const petOwners = $derived(
    new Map(
      (file?.units ?? [])
        .filter((unit) => unit.owner_guid !== undefined && playerSet.has(unit.owner_guid))
        .map((unit) => [unit.guid, unit.owner_guid as string]),
    ),
  );
  const measureOptions = $derived({ pets: petOwners, countOverkill: filters.countOverkill });
  function measureRow(actor: Actor): Promise<ExactSplit> {
    return measureExact(
      sharedQueryLayer(),
      eventsUrl(dataBase, state.fight),
      tableKind,
      actor.guid,
      cutWindow,
      filters.target === '' ? null : filters.target,
      measureOptions,
    );
  }

  /**
   * The whole table measured from the fight's events for this tab, window and filter:
   * null until asked, and dropped the moment any of those changes, since it answered a
   * different question. A window met by a target or boss filter is the case that needs
   * it: the summary can only prorate the one by the other.
   */
  let tableExact = $state<Map<string, ExactTotals> | null>(null);
  let tableMeasuring = $state(false);
  let tableMeasureError = $state('');
  /** The measure that is wanted now; an answer for an older question is dropped. */
  let measureToken = 0;
  $effect(() => {
    // Read everything the measure depends on, so a change in any of it re-arms it. The
    // measure is the default the moment the table is prorated: every reviewer read a
    // prorated window as fact, and one read the wrong killer off it. The delay lets a
    // drag settle before the events are asked.
    const wanted =
      actorTableApproximate &&
      !nightMode &&
      state.view === 'tables' &&
      (state.tab === 'damage-done' || state.tab === 'damage-taken' || state.tab === 'healing');
    void [tableKind, cutWindow, tableScope, filters.countOverkill, state.fight];
    tableExact = null;
    tableMeasureError = '';
    const token = ++measureToken;
    if (!wanted) return;
    const timer = setTimeout(() => void runTableMeasure(token), 350);
    return () => clearTimeout(timer);
  });
  /**
   * What the filter narrowed the other side to, for the measure: a target by GUID and
   * name, or the bosses. Read off the summary's own pairs, never the measured table: the
   * measure must not depend on its own answer.
   */
  const tableScope = $derived.by((): TargetScope | null => {
    if (filters.target !== '') {
      const source =
        state.tab === 'damage-taken'
          ? scoped?.damage_taken
          : state.tab === 'healing'
            ? scoped?.healing
            : scoped?.damage_done;
      // eslint-disable-next-line svelte/prefer-svelte-reactivity
      const names = new Set<string>();
      for (const actor of source ?? []) {
        for (const pair of actor.targets) if (pair.guid === filters.target) names.add(pair.name);
      }
      return { guids: [filters.target], names: [...names] };
    }
    if (filters.bossOnly && state.tab !== 'healing') return { guids: [...filterContext.bosses], names: [] };
    return null;
  });
  async function runTableMeasure(token = measureToken): Promise<void> {
    tableMeasuring = true;
    tableMeasureError = '';
    try {
      const measured = await measureTable(
        sharedQueryLayer(),
        eventsUrl(dataBase, state.fight),
        tableKind,
        cutWindow,
        tableScope,
        measureOptions,
      );
      if (token === measureToken) tableExact = measured;
    } catch (thrown) {
      if (token === measureToken)
        tableMeasureError = thrown instanceof Error ? thrown.message : 'The measurement did not run.';
    } finally {
      if (token === measureToken) tableMeasuring = false;
    }
  }

  /** The window, cut at the last death when the filter asks for it. */
  const cutWindow = $derived.by(() => {
    if (base === null || !state.flags.includes('ignoreAfterDeath') || base.deaths.length === 0)
      return timeWindow;
    const last = Math.max(...base.deaths.map((death) => death.at_ms));
    return clampWindow(
      { startMs: timeWindow.startMs, endMs: Math.min(timeWindow.endMs, last) },
      base.duration_ms,
    );
  });
  /** The fight in the window, before the source scope: what the events view reads whole. */
  const windowed = $derived(base === null ? null : scopeSummary(base, cutWindow));
  const scoped = $derived(
    windowed === null ? null : scopeSource(windowed, state.source, playerSet, friendlySet),
  );
  const presets = $derived(summary === null ? [] : windowPresets(summary));
  /** What the chart draws: the table the tab shows, under the source scope. */
  const chartActors = $derived.by(() => {
    if (scoped === null) return [];
    const table =
      state.tab === 'damage-taken'
        ? scoped.damage_taken
        : state.tab === 'healing'
          ? scoped.healing
          : scoped.damage_done;
    return table.filter((actor) => inSource(actor.guid, state.source, playerSet, friendlySet));
  });
  const chartSeries = $derived(combinedSeries(chartActors));
  /** On the summary, damage taken and healing ride behind the damage line. */
  const chartExtra = $derived(
    scoped === null || state.tab !== 'summary'
      ? []
      : [
          {
            label: 'Damage taken',
            series: combinedSeries(
              scoped.damage_taken.filter((actor) =>
                inSource(actor.guid, state.source, playerSet, friendlySet),
              ),
            ),
            token: 'var(--color-death)',
          },
          {
            label: 'Healing',
            series: combinedSeries(
              scoped.healing.filter((actor) => inSource(actor.guid, state.source, playerSet, friendlySet)),
            ),
            token: 'var(--color-kill)',
          },
        ],
  );
  const chartLabel = $derived(
    state.tab === 'damage-taken' ? 'Damage taken' : state.tab === 'healing' ? 'Healing' : 'Damage',
  );
  const windowIsWhole = $derived(base !== null && isFullWindow(timeWindow, base.duration_ms));

  /**
   * True while the report is still being written -- the report's own status says so, or a
   * fight report.json already knows about is still open -- which is what the poll below
   * exists for. Not the same question as "is the fight on screen still open": that is
   * `fightIsLive`, just below, and it is what the header badge answers. A report can stay
   * live (still gaining fights) after the fight the visitor is looking at has already
   * closed, and the poll has to keep running for that reason even once the badge has gone.
   */
  const isLive = $derived(meta?.status === 'live' || fights.some((entry) => entry.in_progress));
  /** Drives the "Live" badge: the fight actually on screen, not the report as a whole. */
  const fightIsLive = $derived(fight?.in_progress === true);

  /** The filter bar, read off the url state so a copied link carries it. */
  const filters = $derived<ReportFilters>({
    target: state.target,
    ability: state.ability,
    bossOnly: state.flags.includes('bossOnly'),
    playersOnly: state.flags.includes('playersOnly'),
    countOverkill: state.flags.includes('countOverkill'),
    ignoreAfterDeath: state.flags.includes('ignoreAfterDeath'),
  });

  function setFilters(next: ReportFilters): void {
    patch({
      target: next.target,
      ability: next.ability,
      flags: (Object.keys(FLAG_LETTERS) as FlagKey[]).filter((key) => next[key]),
    });
  }
  /** 'Copied' for a moment after the link is copied; '' otherwise. */
  let copied = $state('');

  async function copyLink(): Promise<void> {
    try {
      await navigator.clipboard.writeText(window.location.href);
      copied = 'Copied';
    } catch {
      copied = 'Copy failed';
    }
    setTimeout(() => (copied = ''), 2000);
  }
  /** Set when the url named a fight the report does not have, cleared on the next pick. */
  let missingFight = $state<number | null>(null);

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
  let percentiles = $state(new Map<string, Placement>());
  const loader = createPercentileLoader();

  const filterContext = $derived({
    // The whole night's bosses are every encounter's; one pull's is its own.
    bosses: nightMode
      ? bossGuidsOf(
          file?.units ?? [],
          fights.filter((entry) => entry.kind === 'encounter').map((entry) => entry.name),
        )
      : bossGuids(file?.units ?? [], fight?.name ?? ''),
    players: playerSet,
    deaths: scoped?.deaths ?? [],
  });

  /**
   * What an empty Parse cell means, for the tables to say so: '' on trash, where there is
   * nothing to rank; 'wipe' on a wipe; a dash on a kill nothing of that spec has been
   * ranked on yet.
   */
  const parseFallback = $derived(
    fight?.encounter_id === undefined || fight.in_progress ? '' : fight.kill ? '–' : 'wipe',
  );
  /** The actor tables' fallback: Damage Taken has no parse at all, and its cells say so. */
  const tableParseFallback = $derived(state.tab === 'damage-taken' ? 'none' : parseFallback);

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
    // Only classes this build has talent data for: a log from another client can carry
    // classes Forever does not, and asking for their file is a 404 on every load.
    const classes = [
      ...new Set(
        (summary?.roster ?? [])
          .map((row) => row.class)
          .filter((name): name is string => name !== undefined && classSlugOf(name) !== null),
      ),
    ];
    // resolveTreeSizes (src/lib/report/tree-sizes.ts) only ever returns a class this build
    // truly has no talent data for (a 404) or one it fetched successfully -- never a class
    // that merely failed to load this time, so a network blip or a 5xx cannot pin that class
    // to gear-only links forever: it stays out of `treeSizes` and this effect tries it again
    // the next time it runs.
    void resolveTreeSizes(classes, new Set(treeSizes.keys()), loadTalents, activeBuild.build).then(
      (entries) => {
        if (entries.length === 0) return;
        treeSizes = new Map([...treeSizes, ...entries]);
      },
    );
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
          ? source.filter((actor) => !friendlySet.has(actor.guid))
          : source.filter((actor) => actor.guid === state.source);
    // "Boss damage only" has no meaning for healing, whose targets are players: applied
    // there it blanked the table.
    const applied = state.tab === 'healing' ? { ...filters, bossOnly: false } : filters;
    const filtered = applyActorFilters(bySource, applied, filterContext);
    if (tableExact === null) return filtered;
    // Measured rows: the events' own totals and pairs replace the prorated ones, and a
    // row the events do not mention did nothing in this window to these targets.
    const measured = tableExact;
    return filtered
      .map((actor): Actor => {
        const found = measured.get(actor.guid);
        return {
          ...actor,
          total: found?.effective ?? 0,
          effective: found?.effective ?? 0,
          overheal: undefined,
          targets: found?.targets ?? [],
          measured: true,
        };
      })
      .sort((a, b) => b.effective - a.effective);
  });

  const metricLabel = $derived(state.tab === 'healing' ? 'Healing' : 'Damage');

  /**
   * The metric a table tab's Parse column places every row on, the way Warcraft Logs
   * does: damage per second on Damage Done and healing per second on Healing, each row
   * among ranked kills by its own spec, so a healer has a damage parse and a tank has
   * one too. Damage Taken has no parse: taking more is not doing better.
   */
  function tabMetric(tab: string, row: RosterRow): { metric: string; value: number } | null {
    if (tab === 'healing') return { metric: 'hps', value: row.hps };
    if (tab === 'damage-taken') return null;
    return { metric: 'dps', value: row.dps };
  }

  /** The Summary tab's headline parse: healers on healing, everyone else, tanks too, on damage. */
  function roleMetric(row: RosterRow): { metric: string; value: number } {
    if (row.role === 'healer') return { metric: 'hps', value: row.hps };
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
    // A fight still being written is not a parse, for the same reason a brushed window is
    // not: the numbers are a partial fight's, and the cache key carries the rounded dps,
    // which moves on every five-second tick -- so a live 25-player pull asked for 25 fresh
    // t-digest lookups every five seconds, per viewer, and never hit the cache once.
    // Only kills are ranked, so only kills are asked about: a wipe's dps placed among the
    // kills read as "0" beside a tooltip saying wipes have no parse.
    if (
      summary === null ||
      encounterId === undefined ||
      !windowIsWhole ||
      fight?.in_progress === true ||
      fight?.kill !== true
    ) {
      percentiles = new Map();
      return;
    }
    const tab = state.tab;
    // The API files a ranking row under the phase the fight happened in, so the lookup
    // has to name that phase: a report's own `phase` when the API sends one, otherwise
    // the fight's start read against the same boundaries. Asking for `launch` on a beta
    // kill found nothing, and every Parse column sat empty.
    const phase = meta?.phase ?? phaseAt(fight?.start ?? '');
    const difficulty = fight?.difficulty ?? 0;

    // One query per roster row that has a spec, kept beside its GUID so the answers can be
    // put back on the right rows. Summary picks each row's role metric; a table tab asks
    // for its own metric for every row.
    const wanted = summary.roster
      .filter((row) => row.spec !== undefined && row.spec !== '')
      .flatMap((row) => {
        const picked = tab === 'summary' ? roleMetric(row) : tabMetric(tab, row);
        if (picked === null) return [];
        return [
          {
            guid: row.guid,
            query: {
              encounterId,
              difficulty,
              spec: row.spec ?? '',
              phase,
              metric: picked.metric,
              value: Math.round(picked.value * 100) / 100,
            },
          },
        ];
      });

    void loader.load(wanted.map((entry) => entry.query)).then((answers) => {
      if (state.fight !== wantedFight || state.tab !== wantedTab) return;
      // A plain Map, not SvelteMap: this is a throwaway local built up once and then
      // assigned whole to `percentiles` (already $state) below, the same reasoning the
      // `summaries` cache above gives for its own eslint-disable.
      // eslint-disable-next-line svelte/prefer-svelte-reactivity
      const next = new Map<string, Placement>();
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
    // Said out loud rather than swapped silently: a link to fight 20 that opens on fight
    // 1 with no word about it reads as the wrong fight, not as a missing one.
    missingFight = fights.length > 0 && resolved !== parsed.fight ? parsed.fight : null;
    state =
      resolved === parsed.fight ? parsed : withState(parsed, { fight: resolved, start: null, end: null });
  }

  function writeUrl(push: boolean): void {
    const search = reportSearch(state, firstFight);
    const url = `${window.location.pathname}${search}`;
    // A fight change is a page in its own right, so it is pushed and the back button
    // returns to the pull before it. Brushing the chart and flicking between tabs replace:
    // pushing those would bury the visitor's real back destination under a hundred
    // entries. A link shared from the address bar carries the whole state either way.
    if (push) window.history.pushState(null, '', url);
    else window.history.replaceState(null, '', url);
  }

  function patch(next: Partial<ReportState>): void {
    // Changing fight drops the window: milliseconds from one fight's start mean nothing in
    // another's.
    const changesFight = next.fight !== undefined && next.fight !== state.fight;
    if (next.fight !== undefined) missingFight = null;
    // Picking a fight, mode, view, tab or source is a page in its own right, so it is
    // pushed and the back button undoes it. Brushing the chart replaces: a drag writes
    // the url on every pointer move, and pushing those would bury the real back
    // destination under a hundred entries.
    const pushes =
      (['fight', 'mode', 'view', 'tab', 'source'] as const).some(
        (key) => next[key] !== undefined && next[key] !== state[key],
      ) ||
      // The first window set on a whole fight is pushed too, so Back returns to the
      // whole fight rather than to the fight before it; moving an existing window
      // replaces, since a drag writes on every pointer move.
      (next.start !== undefined && next.start !== null && state.start === null);
    state = withState(state, changesFight ? { ...next, start: null, end: null } : next);
    writeUrl(pushes);
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
      // The whole night is loaded by its own effect once the page is ready.
      if (fight !== null) await loadFight(state.fight);
      status = 'ready';
    } catch (thrown) {
      status = 'failed';
      error = thrown instanceof Error ? thrown.message : REPORT_LOAD_FAILED;
    }
  }

  /**
   * Every read of a fight file, through the one re-signing retry.
   *
   * `dataBase` is a signed url for a private or guild report and it expires in ten
   * minutes, which is a fraction of how long a report page stays open. load.ts's
   * withFreshBase turns the 403 that follows into one fresh `/access` call and one retry,
   * and the base that worked is kept for the next read.
   */
  async function fromDataBase<T>(work: (base: string) => Promise<T>): Promise<T> {
    const { value, base } = await withFreshBase(dataBase, () => fetchAccessUrl(reportId), work);
    if (base !== dataBase) dataBase = base;
    return value;
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
    const pending = inflight.get(index);
    if (pending !== undefined) return pending;
    const request = fetchFight(index).finally(() => inflight.delete(index));
    inflight.set(index, request);
    return request;
  }

  async function fetchFight(index: number): Promise<void> {
    // An open fight has no summary.json at all -- the engine writes it when the fight
    // closes -- so asking for one 404s and renders "No report with that id" over a report
    // that is perfectly fine, for as long as the pull lasts. The snapshot is live.json,
    // and it is never cached: it is a few seconds old by definition.
    const entry = fights.find((candidate) => candidate.index === index);
    if (entry?.in_progress === true) {
      const live = await fromDataBase((base) => fetchLive(base, index));
      if (live !== null) {
        if (index === state.fight) summary = live;
        return;
      }
      // Null means the fight closed between report.json being read and this request, so
      // the immutable summary exists after all.
    }
    const loaded = await fromDataBase((base) => fetchSummary(base, index));
    summaries.set(index, loaded);
    if (index === state.fight) summary = loaded;
  }

  $effect(() => {
    void loadReport();
  });

  // The shell reserves height on #report (min-h-[200svh] in src/pages/reports/[id].astro)
  // so the footer does not jump once this component's real content replaces the loading
  // placeholder -- but Svelte's mount() only ever manages #report's children, never the
  // element itself, so nothing else ever removes that class. Left alone it is a permanent
  // minimum height rather than a one-time reservation: any report shorter than ~200svh
  // would show a standing gap above the footer for as long as the page stayed open. This
  // clears it the first time `status` leaves 'loading' -- by then loadReport has already
  // set `summary` (or failed), so the real content is already in the DOM and removing the
  // reservation causes no further shift. The class name is duplicated from the shell on
  // purpose rather than shared: the two files sides of this are a plain Astro page and a
  // component with no build-time link between them.
  $effect(() => {
    if (status === 'loading') return;
    document.getElementById('report')?.classList.remove('min-h-[200svh]');
  });

  // Re-read the URL when the visitor uses the browser's own back and forward.
  $effect(() => {
    const onPop = (): void => {
      readUrl();
      if (state.fight !== ALL_FIGHTS) void loadFight(state.fight).catch(() => {});
    };
    window.addEventListener('popstate', onPop);
    return () => window.removeEventListener('popstate', onPop);
  });

  // The whole night: every boss pull's summary, a few at a time, folded into `night` as
  // they land so the table fills in rather than waiting for the slowest fetch. Pulls the
  // cache already holds cost nothing; the ones this fetches are cached for the fight
  // views. Re-run when the selection returns here after a pull was opened, which is free
  // once everything is cached.
  $effect(() => {
    if (status !== 'ready' || !nightMode) return;
    const encounters = fights.filter((entry) => entry.kind === 'encounter' && !entry.in_progress);
    let cancelled = false;
    night = aggregateNight(encounters, summaries);
    nightFold = nightSummary(encounters, summaries);
    const queue = encounters.filter((entry) => !summaries.has(entry.index)).map((entry) => entry.index);
    if (queue.length === 0) return;
    nightLoading = true;
    const workers = Array.from({ length: Math.min(4, queue.length) }, async () => {
      for (let index = queue.shift(); index !== undefined; index = queue.shift()) {
        if (cancelled) return;
        try {
          const loaded = await fromDataBase((base) => fetchSummary(base, index as number));
          summaries.set(index, loaded);
        } catch {
          /* One pull that will not load leaves a gap the summary line reports. */
        }
        if (!cancelled) {
          night = aggregateNight(encounters, summaries);
          nightFold = nightSummary(encounters, summaries);
        }
      }
    });
    void Promise.all(workers).then(() => {
      if (!cancelled) nightLoading = false;
    });
    return () => {
      cancelled = true;
      nightLoading = false;
    };
  });

  // One fetch per fight, and only when the selection actually moves.
  $effect(() => {
    const wanted = state.fight;
    // `dataBase` is read untracked: it is a guard, not a trigger. A re-sign assigns it
    // from inside the very loadFight this effect started, before that load has set
    // `summary`, so tracking it re-ran the effect and fetched the same fight's summary a
    // second time on every expiry. `status` turns 'ready' only after `dataBase` is set,
    // so the first run still sees a base.
    if (status !== 'ready' || untrack(() => dataBase) === '') return;
    if (fight === null) return;
    if (summary?.fight_index === wanted) {
      // The selection came back to the fight whose numbers are on screen (a failed pick
      // in between), so the alert from that pick no longer describes what is shown.
      error = '';
      return;
    }
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

  /**
   * The live poll. Stops in three ways: this effect's own cleanup runs when the component
   * unmounts (Svelte tears down every live effect on destroy, so the poller is stopped with
   * it); `isLive` turning false (the fight on screen closed and report.json has nothing
   * else open) reruns this effect, which runs the same cleanup and then returns before a
   * new poller starts; and a failed report load (`status !== 'ready'`) or a report with no
   * data base yet never starts one. A closed report never enters this effect at all, so it
   * costs nothing.
   */
  $effect(() => {
    if (!isLive || dataBase === '' || status !== 'ready') return;

    const poller = createPoller(async () => {
      // report.json first: it is what turns a live fight into a closed one and adds the
      // next fight to the selector.
      const next = await fromDataBase(fetchReportFile);
      file = next;
      if (meta !== null) meta = { ...meta, fights: next.fights };

      const selected = next.fights.find((entry) => entry.index === state.fight);
      if (selected === undefined) return;
      if (selected.in_progress) {
        // Captured before the await, re-checked after: the visitor can switch fights while
        // this request is in flight, and a live snapshot of fight A must never land on
        // fight B, the same discipline loadFight and Effect 3 above already use.
        const wantedFight = selected.index;
        const live = await fromDataBase((base) => fetchLive(base, selected.index));
        if (live !== null && wantedFight === state.fight) {
          summary = live;
          // Cleared here as well as in Effect 3 and in the closing branch below: Effect 3
          // returns early while `summary.fight_index` already matches the selection, so a
          // failure from a fight the visitor has since come back from would otherwise sit
          // under the header for as long as this fight stays open.
          error = '';
        }
        return;
      }
      // The fight closed while we were watching: the immutable summary replaces the
      // snapshot, and the cache entry with it. loadFight carries its own `index ===
      // state.fight` guard, so a fight switch mid-request is handled there too.
      summaries.delete(selected.index);
      await loadFight(selected.index);
      // Effect 3's own success handler clears `error` the same way: a "did not load" alert
      // left over from an earlier failed fetch of this same fight must not outlive the
      // poll quietly loading it correctly, the same lie in reverse an alert that outlived
      // its fight already is.
      if (selected.index === state.fight) error = '';
    }, POLL_INTERVAL_MS);

    poller.start();
    return () => poller.stop();
  });
</script>

{#if status === 'failed'}
  <p class="px-[18px] text-[14px] md:px-0" role="alert" data-testid="report-error">{error}</p>
{:else if status === 'loading' || meta === null}
  <p class="text-muted px-[18px] text-[14px] md:px-0">Loading the report.</p>
{:else}
  <header class="flex flex-col gap-1 px-[18px] md:px-0">
    <div class="flex flex-wrap items-center gap-2">
      <h1 class="section-title text-[18px]" data-testid="report-title">
        {meta.title === '' ? meta.zone : meta.title}
      </h1>
      <!-- Announced politely, not stolen focus: a fight going live or settling is worth a
           screen reader hearing about on its own time, not interrupting whatever the
           visitor is doing. role="status" plus aria-live="polite" is the same pairing
           SummaryBar.svelte uses for its own background result. -->
      <span role="status" aria-live="polite">
        {#if fightIsLive}
          <span class="pill pill-site" data-testid="report-live">Live</span>
        {/if}
      </span>
      <button
        type="button"
        class="border-line-warm rounded-control text-text ml-auto inline-flex h-11 items-center border px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:h-9"
        title="Copy a link to exactly this view"
        data-testid="copy-link"
        onclick={() => void copyLink()}
      >
        {copied === '' ? 'Copy link' : copied}
      </button>
    </div>
    <p class="text-muted text-[13px]" data-testid="report-subtitle">
      {#if meta.zone !== ''}{meta.zone} ·{/if}
      {#if fights.length > 0}<span class="tabular font-mono">{formatDate(fights[0].start)}</span> ·{/if}
      <span class="tabular font-mono">{fights.length} fights</span> · {meta.status}
      {#if nightMode}· All boss pulls{/if}
      {#if fight}· {fight.name}
        {#if fight.kind === 'encounter' && !fight.in_progress}
          <span
            class="font-semibold {fight.kill ? 'text-kill' : 'text-wipe'}"
            title={fight.kill ? 'The boss died' : 'The percentage is the boss’s health when the pull ended'}
            data-testid="report-fight-outcome">{outcomeLabel(fight)}</span
          >
        {/if}
        <span class="tabular font-mono">{formatDuration(fight.duration_ms)}</span>
        {#if fight.deaths > 0}
          · <span class="text-death tabular font-mono">{fight.deaths}</span>
          {fight.deaths === 1 ? 'death' : 'deaths'}
        {/if}{/if}
    </p>
  </header>

  {#if missingFight !== null}
    <p class="text-muted px-[18px] text-[13px] md:px-0" role="status" data-testid="report-missing-fight">
      This report has no fight <span class="tabular font-mono">{missingFight}</span>, so the first fight is
      showing.
    </p>
  {/if}

  {#if error !== ''}
    <!-- The fight selector and the url have already moved by the time a fight's summary
         fails, so without this the previous fight's numbers sit under the new fight's
         label. A live report whose next fight is not written yet is the ordinary way to
         reach it. One line until a later task owns a real error panel. -->
    <p class="text-muted px-[18px] text-[13px] md:px-0" role="alert" data-testid="report-fight-error">
      {error}
    </p>
  {/if}

  <div class="grid grid-cols-1 gap-[22px] px-[18px] md:grid-cols-[300px_minmax(0,1fr)] md:gap-8 md:px-0">
    <FightSelector {fights} selected={state.fight} onSelect={(index) => patch({ fight: index })} />

    <div class="flex min-w-0 flex-col gap-[22px] md:gap-6">
      <!-- Sticky on phone only: the desktop layout keeps the selector column beside the
           content and the whole bar is a short scroll from anything. On a phone a forty-row
           table puts the tabs a long way off the top of the screen, so the bar rides under
           the page header instead. The negative margin takes it out to the viewport edges
           so its background covers the rows sliding under it, and the padding puts the
           18px gutter back on its own children. -->
      <div class="bg-bg sticky top-0 z-10 -mx-[18px] px-[18px] py-2 md:static md:mx-0 md:px-0 md:py-0">
        <ModeBar {state} {roster} onPatch={patch} {nightMode} />
      </div>
      <Glossary />
      <!-- The chart and its presets are a fight's: nothing draws a chart over a night. -->
      {#if summary !== null && !nightMode}
        <TimeChart
          series={chartSeries}
          extra={chartExtra}
          durationMs={summary.duration_ms}
          window={timeWindow}
          deaths={summary.deaths.map((death) => ({ at_ms: death.at_ms, name: death.name }))}
          label={chartLabel}
          onWindow={setWindow}
        />
        <div class="flex gap-2 overflow-x-auto md:flex-wrap" data-testid="window-presets">
          <!-- Keyed by position, not by label: a battle-rez puts the same name in
               `deaths` twice, and two buttons labelled "Before Thalgrit died" would be a
               duplicate key, which Svelte throws on rather than renders. The list is
               rebuilt wholesale whenever the fight changes, so position is stable. -->
          {#each presets as preset, position (position)}
            {@const active =
              preset.window === null
                ? windowIsWhole
                : timeWindow.startMs === preset.window.startMs && timeWindow.endMs === preset.window.endMs}
            <button
              type="button"
              class="rounded-control inline-flex h-11 shrink-0 items-center border px-3 text-[12px] font-bold tracking-[0.06em] whitespace-nowrap uppercase md:h-9 {active
                ? 'border-gold bg-card-top text-strong'
                : 'border-line-soft text-nav'}"
              aria-pressed={active}
              onclick={() => setWindow(preset.window)}
            >
              {preset.label}
            </button>
          {/each}
        </div>
      {/if}
      {#if scoped !== null && state.mode === 'analyze' && state.view === 'tables'}
        {#if state.tab === 'summary' && nightMode}
          {#if night !== null}
            <NightView
              {night}
              loading={nightLoading}
              onSelect={(index) => patch({ fight: index })}
              onSelectPlayer={(guid) => patch({ source: guid })}
            />
          {/if}
        {:else if state.tab === 'summary'}
          <SummaryTab
            summary={scoped}
            durationMs={scoped.duration_ms}
            {percentiles}
            {parseFallback}
            approximate={!windowIsWhole}
            dataBuild={activeBuild.build}
            {classOf}
            {treeSizesFor}
            onSelectPlayer={(guid) => patch({ source: guid })}
            players={playerSet}
            onTab={(tab) => patch({ tab })}
          />
        {:else if state.tab === 'damage-done' || state.tab === 'damage-taken' || state.tab === 'healing'}
          <FilterBar
            {filters}
            actors={tabActors}
            onChange={setFilters}
            showBossOnly={state.tab !== 'healing'}
          />
          <ActorTable
            actors={tabActors}
            durationMs={scoped.duration_ms}
            {metricLabel}
            {percentiles}
            parseFallback={tableParseFallback}
            pairsLabel={state.tab === 'damage-taken' ? 'Sources' : 'Targets'}
            measure={nightMode ? undefined : measureRow}
            approximate={actorTableApproximate}
            amountApproximate={filtersScale && !windowIsWhole}
          />
          {#if actorTableApproximate}
            <p class="text-muted text-[12px]" data-testid="approximate-note">
              {#if tableExact !== null}
                <span class="text-kill" data-testid="table-measured"
                  >Totals, shares, per-second figures and targets are measured from the fight’s events for
                  this window and filter.</span
                > A row’s ability split is measured the same way when it is opened.
              {:else if nightMode}
                Over the whole night a tilde marks a figure split across abilities and targets in proportion
                to the window and any target or boss filter{filtersScale && !windowIsWhole
                  ? ', totals included'
                  : ''}. Open a pull to read its figures from the fight’s events.
              {:else if tableMeasuring}
                <span data-testid="table-measuring">Measuring this window from the fight’s events…</span> A tilde
                marks a figure still prorated from the summary.
              {:else if tableMeasureError !== ''}
                <span class="text-wipe" role="alert">{tableMeasureError}</span> A tilde marks a figure
                prorated from the summary{filtersScale && !windowIsWhole ? ', totals included' : ''}.
                <button
                  type="button"
                  class="text-gold inline-flex min-h-11 items-center underline-offset-2 hover:underline md:min-h-0"
                  data-testid="measure-table"
                  onclick={() => void runTableMeasure(++measureToken)}>Try the measure again</button
                >
              {:else}
                A tilde marks a figure split across abilities and targets in proportion to the window and any
                active target or boss filter. Totals and per-second figures are exact.
                <button
                  type="button"
                  class="text-gold inline-flex min-h-11 items-center underline-offset-2 hover:underline md:min-h-0"
                  data-testid="measure-exactly"
                  onclick={() => patch({ view: 'queries' })}>Measure this window exactly in Queries</button
                >.
              {/if}
            </p>
          {/if}
        {:else if state.tab === 'buffs'}
          <RaidCooldowns tracks={scoped.auras} window={cutWindow} deaths={scoped.deaths} names={unitNames} />
          <AuraTable tracks={scoped.auras} durationMs={scoped.duration_ms} kind="BUFF" />
        {:else if state.tab === 'debuffs'}
          <AuraTable
            tracks={scoped.auras}
            durationMs={scoped.duration_ms}
            kind="DEBUFF"
            bossNames={bossUnitNames}
          />
        {:else if state.tab === 'casts'}
          <CastTable
            rows={scoped.casts}
            durationMs={scoped.duration_ms}
            startMs={timeWindow.startMs}
            {classOf}
            approximate={!windowIsWhole}
          />
        {:else if state.tab === 'interrupts'}
          <ExchangeTable
            rows={scoped.interrupts}
            emptyText="Nothing was interrupted in the whole fight."
            casts={base?.casts ?? []}
            players={playerSet}
          />
        {:else if state.tab === 'dispels'}
          <ExchangeTable rows={scoped.dispels} emptyText="Nothing was dispelled in the whole fight." />
        {:else if state.tab === 'resources'}
          <ResourceGraphs
            tracks={scoped.resources}
            durationMs={scoped.duration_ms}
            deaths={scoped.deaths.map((death) => ({
              guid: death.guid,
              at_ms: death.at_ms - timeWindow.startMs,
            }))}
          />
        {:else if state.tab === 'threat'}
          <ThreatTable
            rows={scoped.threat}
            {classOf}
            approximate={!windowIsWhole}
            totalThreat={windowed?.threat.reduce((sum, row) => sum + row.threat, 0)}
          />
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
      {#if nightMode && (state.view !== 'tables' || state.mode !== 'analyze')}
        <p class="text-muted text-[14px]" data-testid="night-tables-only">
          Timelines, events, queries, compare and rankings are one pull's. Pick a pull from the list to see
          them, or
          <button
            type="button"
            class="text-gold inline-flex min-h-11 items-center underline-offset-2 hover:underline md:min-h-0"
            onclick={() => patch({ mode: 'analyze', view: 'tables' })}>go back to the night's tables</button
          >.
        </p>
      {/if}
      {#if scoped !== null && !nightMode && state.mode === 'analyze' && state.view === 'timelines'}
        <TimelinesView
          summary={scoped}
          window={timeWindow}
          {classOf}
          bossName={fight?.kind === 'encounter' ? fight.name : ''}
          players={playerSet}
          allCasts={summary?.casts ?? []}
        />
      {/if}
      {#if scoped !== null && !nightMode && state.mode === 'analyze' && state.view === 'events'}
        <EventsView
          summary={windowed ?? scoped}
          {classOf}
          inScope={(guid) => inSource(guid, state.source, playerSet, friendlySet)}
        />
      {/if}
      <!-- `scoped` only to say a summary has loaded, the same guard its three siblings
           use; the Queries view reads the fight's events.parquet, not the summary, and
           takes the window so a starting point is written for what is on screen. -->
      {#if scoped !== null && !nightMode && state.mode === 'analyze' && state.view === 'queries'}
        <QueriesView dataBaseUrl={dataBase} fightIndex={state.fight} window={timeWindow} />
      {/if}
      {#if state.mode === 'compare' && summary !== null && !nightMode}
        <CompareMode
          {fights}
          current={state.fight}
          dataBaseUrl={dataBase}
          left={summary}
          window={windowIsWhole ? null : cutWindow}
          rightIndex={state.compareWith}
          metric={state.compareMetric}
          onPatch={patch}
        />
      {/if}
      {#if state.mode === 'rankings' && fight !== null}
        <RankingsMode {fight} {reportId} encounterSlug={currentEncounterSlug} />
      {/if}
    </div>
  </div>
{/if}
