<!-- web/src/components/report/ReportView.svelte -->
<!-- The report island's root (logs spec 2026-11-07 section 6). It owns three things and delegates
     everything else:
       the URL state, which is the page's whole state (src/lib/report/url.ts)
       the data, which is one meta fetch, one report.json and one summary per fight
       the frame: the PageHeader with its answer and the one tab strip (ReportHead), the control
       bar (ReportBar), then a split Page whose main column holds the tab's one Plate first and its
       open Sections after, and whose rail holds This pull and the tab's own rail Sections. -->
<script lang="ts">
  import { iconUrl } from '../../lib/ui/assets';
  import { untrack } from 'svelte';
  import { simCompareFightHref } from '../../lib/handoff-links';
  import Notice from '../ui/Notice.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';
  import TabStripMore from '../ui/TabStripMore.svelte';
  import Tag from '../ui/Tag.svelte';
  import {
    POLL_INTERVAL_MS,
    REPORT_LOAD_FAILED,
    REPORT_NOT_FOUND,
    createPoller,
    fetchAccessUrl,
    fetchLive,
    fetchReportFile,
    fetchReportFileFor,
    fetchReportMeta,
    fetchReportMetaAfterUpload,
    eventsUrl,
    fetchSummary,
    withFreshBase,
  } from '../../lib/report/load';
  import {
    combinedFights,
    defaultFightIndex,
    resolveFightIndex,
    sampleFightIndex,
  } from '../../lib/report/fights';
  import { aggregateNight, nightFacts, nightSummary, type Night } from '../../lib/report/night';
  import { schoolToken, TRASH_PARSE } from '../../lib/report/format';
  import {
    FLAG_LETTERS,
    MISSING_FIGHT,
    SOURCE_ENEMIES,
    SOURCE_EVERYONE,
    SOURCE_FRIENDLIES,
    legacyTab,
    trashScopeKey,
    type FlagKey,
    defaultState,
    parseReportState,
    reportSearch,
    withState,
    type ReportState,
  } from '../../lib/report/url';
  import {
    abilityKey,
    type Ability,
    type Actor,
    type FightEntry,
    type ReportFile,
    type ReportMeta,
    type RosterRow,
    type Summary,
  } from '../../lib/report/types';
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
  import ActorTable from '../coach/GatedActorTable.svelte';
  import DeathsTab from './DeathsTab.svelte';
  import NightView from './NightView.svelte';
  import PlayerView from './PlayerView.svelte';
  import RatingPanel from './RatingPanel.svelte';
  import ReportBar from './ReportBar.svelte';
  import ReportHead from './ReportHead.svelte';
  import SummaryPanels from './SummaryPanels.svelte';
  import ThisPull from './ThisPull.svelte';
  import SummaryDeaths from './SummaryDeaths.svelte';
  import KeyMechanics from './KeyMechanics.svelte';
  import GroupBuffs from './GroupBuffs.svelte';
  import TargetsRail from './TargetsRail.svelte';
  import CastsRail from './CastsRail.svelte';
  import DispelsRail from './DispelsRail.svelte';
  import Glossary from './Glossary.svelte';
  import EnemyCastsRail from './EnemyCastsRail.svelte';
  import { enemyCasts, spellSchools } from '../../lib/report/exchange-rails';
  import GroupDebuffsRail from './GroupDebuffsRail.svelte';
  import MechanicsRail from './MechanicsRail.svelte';
  import QueriesRail from './QueriesRail.svelte';
  import RatingRail from './RatingRail.svelte';
  import ThreatRail from './ThreatRail.svelte';
  import type { ReportRatings } from '../../lib/rating/types';
  import type { ReportRatingsStatus } from '../../lib/rating/report-ratings.svelte';
  import { reportFacts, reportHeading } from '../../lib/report/heading';
  import { hasRankedKill } from '../../lib/report/columns';
  import { quietMinutes } from '../../lib/report/quiet';
  import { isSampleReport } from '../../lib/report/sample';
  import { pageTitle } from '../../lib/seo';
  import { reportCredit } from '../../lib/report/credit';
  import { homeRealmOf, setHomeRealm, unitLabel } from '../../lib/report/names.svelte';
  import { unitClassMap } from '../../lib/report/roster-class';
  import SummaryTab from './SummaryTab.svelte';
  import TimeChart from './TimeChart.svelte';
  import { loadChartChunk } from '../../lib/chart/uplot-loader';
  import { deathKey, rateMarkers, threatMarkers } from '../../lib/chart/report-markers';
  import type { ChartEvent, ChartPhase } from '../../lib/chart/types';
  import type { ThreatPlayer } from '../../lib/chart/threat';
  import { onCurrentCharacterChange, readCurrent, type CurrentCharacter } from '../../lib/current-character';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import Row from '../ui/Row.svelte';
  import Spinner from '../ui/Spinner.svelte';
  import Crest from '../ui/Crest.svelte';
  import { REGION } from '../../lib/ui/regions';
  import { button as buttonClass } from '../../lib/ui/button';
  import { REPORT_LAZY_MIN_H } from '../../lib/report/layout';
  import { provideCoachRow } from '../../lib/coach/report-context';
  import CoachTabHost from '../coach/CoachTabHost.svelte';
  import CoachRowLine from '../coach/CoachRowLine.svelte';
  import { createLazyComponent, type LazyLoadState } from '../../lib/report/lazy-component.svelte';
  // Compare, Mechanics, Rankings, Timelines, Events and Queries are not the landing mode or
  // view (Analyze / tables / summary is), so each ships as its own chunk fetched the first
  // time its mode or view is actually picked, instead of sitting in every visitor's initial
  // bundle. See lazy-component.svelte.ts.
  const compareModeLazy = createLazyComponent(() => import('./CompareMode.svelte'));
  const mechanicsModeLazy = createLazyComponent(() => import('./MechanicsMode.svelte'));
  const timelinesViewLazy = createLazyComponent(() => import('./TimelinesView.svelte'));
  const eventsViewLazy = createLazyComponent(() => import('./EventsView.svelte'));
  const queriesViewLazy = createLazyComponent(() => import('./QueriesView.svelte'));
  const ratingTabLazy = createLazyComponent(() => import('./RatingTab.svelte'));
  // The tab tables past Summary and the three actor tabs ship as their own chunks too, fetched the
  // first time their tab opens (the island's gzip budget, spec 13).
  const threatTableLazy = createLazyComponent(() => import('./ThreatTable.svelte'));
  const castTableLazy = createLazyComponent(() => import('./CastTable.svelte'));
  const auraTabLazy = createLazyComponent(() => import('./AuraTab.svelte'));
  const exchangeTableLazy = createLazyComponent(() => import('./ExchangeTable.svelte'));
  const resourceGraphsLazy = createLazyComponent(() => import('./ResourceGraphs.svelte'));
  import {
    abilityOptions,
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
  import { fightGroupScope, groupScope, raidFlaggedPlayers, scopeNoun } from '../../lib/report/group';
  import { splitUnitName } from '../../lib/characters';
  import { reportCopy } from '../../lib/report/copy';
  import { justUploadedAt } from '../../lib/reports/just-uploaded';
  import {
    ACTOR_TABS,
    NIGHT_TABS,
    TABS_MD,
    TABS_XL,
    listedTabs,
    reportFigure,
    reportMeta,
    thisPullFacts,
    reportMetaTail,
  } from '../../lib/report/frame';
  import { pullNumbers } from '../../lib/report/fights';
  import { createCoachViewerState } from '../../lib/coach/viewer-state.svelte';
  import { coachTabListed } from '../../lib/coach/viewer';
  import { loadSpellIcons, NO_SPELL_ICONS, type SpellIconOf } from '../../lib/guild/report-spell-icons';
  // Value imports from exact.ts are deliberately absent: it pulls in the DuckDB query layer
  // (src/lib/report/query.ts), which dynamically imports the multi-hundred-KB duckdb-wasm
  // package. None of that belongs in the initial bundle everyone pays for on load -- every
  // call site below imports exact.ts itself, dynamically, the first time a measured feature
  // is actually used. Only the types are imported statically; a type-only import is erased
  // at build and costs nothing.
  import type { CastCounts, ExactSplit, DeadSpan, ExactTotals, TargetScope } from '../../lib/report/exact';
  import { resolveTreeSizes } from '../../lib/report/tree-sizes';
  import ReportSkeleton from './ReportSkeleton.svelte';
  import { classSlugOf, plannerLinkFor } from '../../lib/report/planner-link';
  import { simLinkFor } from '../../lib/report/sim-link';

  let { reportId, inlineMeta = null }: { reportId: string; inlineMeta?: ReportMeta | null } = $props();

  // untrack because inlineMeta is a one-shot bootstrap, not a binding: the shell renders
  // it once into data-report and never changes it, and reading a prop straight into
  // $state is the pattern Svelte warns about (state_referenced_locally).
  let meta = $state<ReportMeta | null>(untrack(() => inlineMeta));
  let file = $state<ReportFile | null>(null);
  /** The engine that wrote the report's files, riding on their urls so a re-parse is a cache miss. */
  const engineVersion = $derived(file?.engine_version ?? meta?.engine_version ?? '');
  let summary = $state<Summary | null>(null);
  /**
   * 'processing', 'waiting' and 'unparsed' are the screens of a report whose files do not exist
   * (yet, or ever): 'waiting' is a live report whose first fight has not been uploaded.
   */
  let status = $state<'loading' | 'ready' | 'failed' | 'processing' | 'waiting' | 'unparsed'>('loading');
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
  /** The fight a bare url opens on: the first boss pull, not the first trash; the sample's own rule for the sample. */
  const firstFight = $derived(
    isSampleReport(reportId) ? sampleFightIndex(fights) : defaultFightIndex(fights),
  );
  /** A Parse column exists only on a report with a boss kill to rank; the sample has none. */
  // The bundled sample is no ranked night: it never asks the rankings, so no Parse column and no
  // "could not be reached" line over a log that was never ranked.
  const showParse = $derived(!isSampleReport(reportId) && hasRankedKill(fights));
  /**
   * The bundled sample paints its first table from the fight's own summary, before report.json
   * (half a megabyte of units) has arrived: on a phone on a slow line the serial pair kept
   * "Loading the report." up for seconds. The meta already names the players, which is all the
   * first table needs of the units; the rest of them fill in behind it.
   */
  const summaryFirst = $derived(isSampleReport(reportId));
  /** The subtitle's words, the same ones the /logs sample row reads (lib/report/heading.ts). */
  const facts = $derived(reportFacts({ zone: meta?.zone ?? '', created_at: meta?.created_at ?? '', fights }));
  /** The pulls a combined view folds together (every boss pull, or one zone's trash); null for one fight. */
  const combined = $derived(combinedFights(fights, state.fight));
  /** A combined view rather than one fight: the whole night's boss pulls, or a zone's trash. */
  const nightMode = $derived(combined !== null);
  /** The combined view is a zone's trash, not the night's bosses. */
  const trashMode = $derived(trashScopeKey(state.fight) !== null);
  let night = $state<Night | null>(null);
  /** The night as one Summary-shaped object, for the fight tabs to show. */
  let nightFold = $state<Summary | null>(null);
  let nightLoading = $state(false);
  /** What the tables read: the whole night, or the selected fight. */
  const base = $derived<Summary | null>(nightMode ? nightFold : summary);
  let state = $state<ReportState>(defaultState(1));
  const fight = $derived<FightEntry | null>(fights.find((f) => f.index === state.fight) ?? null);
  provideCoachRow(() => ({ reportId, fight, firstFight }));
  /** The rankings board of the boss on screen (decision R3b: the Rankings mode retired to a link). */
  const rankingsHref = $derived(fight?.kind === 'encounter' ? `/rankings/${slugFor(fight.name)}` : null);
  const roster = $derived(
    ((nightMode ? nightFold : summary)?.roster ?? []).map((row) => ({
      guid: row.guid,
      name: row.name,
      class: row.class,
      spec: row.spec,
      role: row.role,
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
  /** The Interrupts and Dispels tabs' empty line, naming the scope it was empty under. */
  function exchangeEmpty(verb: 'interrupted' | 'dispelled'): string {
    const scope = nightMode ? 'tonight' : 'this pull';
    // With nothing kicked, say whether anything could have been: an instant cast has no cast bar.
    const castBars = enemyCasts((windowed ?? base)?.casts ?? [], friendlySet).some((row) => row.castMs > 0);
    const tail = verb === 'interrupted' && !castBars ? ' No enemy cast had a cast bar.' : '';
    const noun = verb === 'interrupted' ? 'interrupts' : 'dispels';
    if (state.source === SOURCE_FRIENDLIES || state.source === SOURCE_EVERYONE)
      return `No ${noun} ${scope}.${tail}`;
    if (state.source === SOURCE_ENEMIES) return `The enemies ${verb} nothing ${scope}.`;
    const name = splitUnitName(unitNames.get(state.source) ?? state.source).name;
    return `${name} ${verb} nothing ${scope}.${tail}`;
  }
  /** The player GUIDs from report.json, for the source scope and the filters. */
  const playerSet = $derived(
    file === null && summaryFirst ? new Set(meta?.players ?? []) : playerGuids(file?.units ?? []),
  );
  /**
   * The pulls the scope is read from: the one on show, or the night's boss pulls. report.json's
   * own lines once it has arrived (the units the fallback reads come from the same file), the
   * meta's before.
   */
  const scopeFights = $derived.by(() => {
    const lines = file?.fights ?? fights;
    return nightMode
      ? lines.filter((entry) => entry.kind === 'encounter')
      : lines.filter((entry) => entry.index === state.fight);
  });
  /**
   * Who the default scope is (lib/report/group.ts). From engine 0.6.0 each pull says who was in
   * the group while it ran, so it is read from there. An older report falls back to the
   * report-wide flags: the logger's group when outsiders outnumber it (an open-world log), else
   * every player; before report.json the sample reads its meta's group.
   */
  const groupScopeNow = $derived(
    fightGroupScope(scopeFights, raidFlaggedPlayers(file?.units ?? [])) ??
      (file === null && summaryFirst
        ? groupScope(
            (meta?.players ?? []).map((guid) => ({ guid, kind: 'player', flags: 0 })),
            meta?.group ?? [],
            fight?.players,
          )
        : groupScope(file?.units ?? [], undefined, nightMode ? undefined : fight?.players)),
  );
  /** Everyone in range is the friendlies over every player seen; the scope rules read it so. */
  const scopeKey = $derived(state.source === SOURCE_EVERYONE ? SOURCE_FRIENDLIES : state.source);
  /** The players the friendlies scope keeps: the group when the default narrows to it. */
  const scopePlayers = $derived(
    state.source === SOURCE_FRIENDLIES && groupScopeNow.scoped ? groupScopeNow.group : playerSet,
  );
  /** What the scope's players are called on the chart and the tables. */
  const noun = $derived(scopeNoun(groupScopeNow, state.source));
  /** The players and their pets and totems: what the enemies scope leaves out. */
  const friendlySet = $derived(
    file === null && summaryFirst ? new Set(meta?.players ?? []) : friendlyGuids(file?.units ?? []),
  );
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
  const measureOptions = $derived({
    pets: petOwners,
    countOverkill: filters.countOverkill,
    exclude: ignoringDead ? deadSpans : [],
    ability: filters.ability ?? undefined,
  });
  async function measureRow(actor: Actor): Promise<ExactSplit> {
    const { measureExact, sharedQueryLayer } = await import('../../lib/report/exact');
    return measureExact(
      sharedQueryLayer(),
      eventsUrl(dataBase, state.fight, engineVersion),
      tableKind,
      actor.guid,
      cutWindow,
      tableScope,
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
  let tableMeasureError = $state('');
  /** The measure that is wanted now; an answer for an older question is dropped. */
  let measureToken = 0;
  $effect(() => {
    // Read everything the measure depends on, so a change in any of it re-arms it. The
    // measure is the default the moment the table is prorated: every reviewer read a
    // prorated window as fact, and one read the wrong killer off it. The delay lets a
    // drag settle before the events are asked.
    const wanted = actorTableApproximate && !nightMode && ACTOR_TABS.has(state.tab);
    void [tableKind, cutWindow, tableScope, filters.countOverkill, ignoringDead, state.fight];
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
    tableMeasureError = '';
    try {
      const { measureTable, sharedQueryLayer } = await import('../../lib/report/exact');
      const measured = await measureTable(
        sharedQueryLayer(),
        eventsUrl(dataBase, state.fight, engineVersion),
        tableKind,
        cutWindow,
        tableScope,
        measureOptions,
      );
      if (token === measureToken) tableExact = measured;
    } catch (thrown) {
      if (token === measureToken)
        tableMeasureError = `The measure did not run, so these are the summary’s prorated figures${
          thrown instanceof Error ? ` (${thrown.message})` : ''
        }.`;
    }
  }

  /**
   * The cast table's window counts, read from the fight's cast lines: the summary keeps
   * only whole-fight starts and failures, so a brushed window's cancelled casts need the
   * events. Null until read, and dropped when the window or fight changes.
   */
  let castExact = $state<Map<string, CastCounts> | null>(null);
  let castMeasureError = $state('');
  let castToken = 0;
  $effect(() => {
    const wanted = !windowIsWhole && !nightMode && state.tab === 'casts';
    void [cutWindow, state.fight];
    castExact = null;
    castMeasureError = '';
    const token = ++castToken;
    if (!wanted) return;
    const timer = setTimeout(() => void runCastMeasure(token), 350);
    return () => clearTimeout(timer);
  });
  async function runCastMeasure(token: number): Promise<void> {
    try {
      const { measureCasts, sharedQueryLayer } = await import('../../lib/report/exact');
      const measured = await measureCasts(
        sharedQueryLayer(),
        eventsUrl(dataBase, state.fight, engineVersion),
        cutWindow,
      );
      if (token === castToken) castExact = measured;
    } catch (thrown) {
      if (token === castToken)
        castMeasureError = `The window’s casts did not load, so these are the summary’s scaled figures${
          thrown instanceof Error ? ` (${thrown.message})` : ''
        }.`;
    }
  }

  /**
   * The one ability drawn behind the main chart's series: a look, not a view, so it is
   * not in the url. Measured over the whole fight once and sliced by the brush like every
   * other line, and dropped whenever the fight, the tab or the source scope changes,
   * because it answered a question about the table that was on screen then. A window is
   * not one of those: reading the line against a stretch is the whole point of it, and
   * the brush only slices a series that is already measured over the whole fight.
   */
  let chartedAbility = $state<{ key: string; label: string; token: string; series: number[] } | null>(null);
  let chartedError = $state('');
  /** The measure that is wanted now; an answer for an older pick is dropped. */
  let chartedToken = 0;
  /**
   * What the pick belongs to, as one string. `patch` replaces the whole `state` object on
   * every url change, brushing included, so an effect that reads `state.fight` directly
   * woke on every pointer move of a drag and cleared the line the reader had just asked
   * for. Comparing the three values against the last ones they were dropped for keeps the
   * clearing to a real change of fight, tab or source scope.
   */
  const chartedScope = $derived(`${state.fight}|${state.tab}|${state.source}|${String(nightMode)}`);
  let chartedFor = chartedScope;
  $effect(() => {
    if (chartedScope === chartedFor) return;
    chartedFor = chartedScope;
    chartedAbility = null;
    chartedError = '';
    chartedToken += 1;
  });
  async function toggleAbilityOnChart(actor: Actor, ability: Ability): Promise<void> {
    const key = `${actor.guid}|${abilityKey(ability)}`;
    if (chartedAbility?.key === key) {
      chartedAbility = null;
      return;
    }
    const token = ++chartedToken;
    chartedError = '';
    try {
      const { measureAbilitySeries, sharedQueryLayer } = await import('../../lib/report/exact');
      const series = await measureAbilitySeries(
        sharedQueryLayer(),
        eventsUrl(dataBase, state.fight, engineVersion),
        tableKind,
        actor.guid,
        ability.spell_id,
        ability.via ?? '',
        base?.duration_ms ?? 0,
        measureOptions,
      );
      if (token !== chartedToken) return;
      chartedAbility = {
        key,
        label: `${splitUnitName(actor.name).name} · ${ability.name}`,
        token: schoolToken(ability.school),
        series,
      };
    } catch (thrown) {
      if (token !== chartedToken) return;
      chartedError = `That ability's line did not load${thrown instanceof Error ? ` (${thrown.message})` : ''}.`;
    }
  }
  /** True where there is a chart to put an ability on: one pull, one player's own view. */
  const abilityChartAvailable = $derived(!nightMode && playerView);
  /** The zero-height mark above the tab's table, for the phone to scroll to on a tab change. */
  let tabAnchor = $state<HTMLElement | undefined>(undefined);
  let shownTab = state.tab;
  /** Scrolls a phone to the mark; a desktop has both in view and never moves. */
  function bringIntoView(anchor: HTMLElement | undefined): void {
    if (typeof window === 'undefined' || !window.matchMedia('(max-width: 767px)').matches) return;
    const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
    anchor?.scrollIntoView({ block: 'start', behavior: reduced ? 'auto' : 'smooth' });
  }
  $effect(() => {
    const tab = state.tab;
    if (tab === shownTab) return;
    shownTab = tab;
    // The Threat tab's chart is the tab's answer (spec §B1.6): on a phone the Source row goes to
    // the top, 8 px down, so the whole chart fits the first screen; every other tab scrolls to its table.
    const sourceRow =
      tab === 'threat' ? (document.getElementById('report-source')?.closest('label') ?? null) : null;
    if (sourceRow instanceof HTMLElement) {
      sourceRow.style.scrollMarginTop = '8px';
      bringIntoView(sourceRow);
    } else bringIntoView(tabAnchor);
  });
  /** The window, cut at the last death when the filter asks for it. */
  // The window itself is never cut: "ignore events after a death" leaves each player's own
  // dead spans out of the measure instead (see `deadSpans`), since a raised player fights on.
  const cutWindow = $derived(timeWindow);
  /** Each player's dead spans: from a death to their first cast after it, or the fight's end. */
  const deadSpans = $derived.by((): DeadSpan[] => {
    if (base === null || nightMode) return [];
    return base.deaths.map((death) => {
      const raised = base.casts
        .filter((row) => row.guid === death.guid)
        .flatMap((row) => row.sequence)
        .filter((at) => at > death.at_ms)
        .sort((a, b) => a - b)[0];
      return { guid: death.guid, startMs: death.at_ms, endMs: raised ?? base.duration_ms };
    });
  });
  const ignoringDead = $derived(
    state.flags.includes('ignoreAfterDeath') && !nightMode && deadSpans.length > 0,
  );
  /** The window's length less the time this player spent dead inside it. */
  function aliveMs(guid: string): number {
    const dead = deadSpans
      .filter((span) => span.guid === guid)
      .reduce(
        (sum, span) =>
          sum +
          Math.max(0, Math.min(span.endMs, cutWindow.endMs) - Math.max(span.startMs, cutWindow.startMs)),
        0,
      );
    return Math.max(0, cutWindow.endMs - cutWindow.startMs - dead);
  }
  /** The fight in the window, before the source scope: what the events view reads whole. */
  const windowed = $derived(base === null ? null : scopeSummary(base, cutWindow));
  const scoped = $derived(
    windowed === null ? null : scopeSource(windowed, scopeKey, scopePlayers, friendlySet),
  );
  // The death presets follow the source pick: with one player chosen, "20 s before X
  // died" offers that player's deaths, not everyone's.
  const presets = $derived(
    scoped === null ? [] : windowPresets(scoped, base?.duration_ms ?? scoped.duration_ms),
  );
  /** The main chart's metric: the table's own on the three rate tabs, damage done elsewhere (spec §B5.2). */
  /** The by-ability panel each amount tab shows under its chart. */
  const ABILITY_METRIC = { 'damage-done': 'done', healing: 'healing', 'damage-taken': 'taken' } as const;
  const chartMetric = $derived<'damage-done' | 'healing' | 'damage-taken'>(
    state.tab === 'damage-taken' ? 'damage-taken' : state.tab === 'healing' ? 'healing' : 'damage-done',
  );
  /** The source scope over the whole fight: the chart's lines run across the fight and the window zooms them. */
  const wholeScoped = $derived(
    base === null || nightMode ? null : scopeSource(base, scopeKey, scopePlayers, friendlySet),
  );
  /** A chart's actors over the whole fight, under the source scope. */
  function wholeChartActorsFor(metric: 'damage-done' | 'healing' | 'damage-taken'): Actor[] {
    if (wholeScoped === null) return [];
    const table =
      metric === 'damage-taken'
        ? wholeScoped.damage_taken
        : metric === 'healing'
          ? wholeScoped.healing
          : wholeScoped.damage_done;
    return table.filter((actor) => inSource(actor.guid, scopeKey, scopePlayers, friendlySet));
  }
  /** The brush's shape on every chart: the raid's damage per second over the whole fight. */
  const brushShape = $derived(
    base === null
      ? []
      : combinedSeries(
          base.damage_done.filter((actor) => scopePlayers.size === 0 || scopePlayers.has(actor.guid)),
        ),
  );
  const chartPhases = $derived<ChartPhase[]>(
    (base?.phases ?? []).map((phase) => ({ name: phase.name, startMs: phase.start_ms, endMs: phase.end_ms })),
  );
  /** One player as the Source, by GUID. */
  const sourcePlayer = $derived(playerSet.has(state.source) ? state.source : undefined);
  /** The viewer's own character in this fight, matched by name and class from the current character. */
  let youGuid = $state<string | undefined>(undefined);
  /** The current character as last read; a choice in the header selector re-reads it. */
  let current = $state<CurrentCharacter | null>(null);
  $effect(() => {
    current = readCurrent();
    return onCurrentCharacterChange(() => {
      current = readCurrent();
    });
  });
  $effect(() => {
    const rows = base?.roster ?? [];
    const you = current;
    if (you === null) {
      youGuid = undefined;
      return;
    }
    const name = you.label.split(' · ')[0]?.trim().toLowerCase() ?? '';
    youGuid = rows.find(
      (row) =>
        splitUnitName(row.name).name.toLowerCase() === name &&
        (row.class ?? '').toLowerCase().replace(/\s+/g, '-') === you.classSlug,
    )?.guid;
  });
  /** A death's link on the Deaths tab, keeping the fight and the window (spec §B6.6). */
  function deathHref(death: { guid: string; at_ms: number }): string {
    return `${window.location.pathname}${reportSearch(
      withState(state, { tab: 'deaths', openDeaths: [deathKey(death)] }),
      firstFight,
    )}`;
  }
  function openDeaths(deaths: readonly ChartEvent[]): void {
    const keys = deaths.map((event) => event.key.replace(/^death-/, ''));
    const span =
      deaths.length > 1 ? { start: deaths[0].atMs, end: deaths[deaths.length - 1].atMs + 1000 } : {};
    patch({ tab: 'deaths', openDeaths: keys, ...span });
  }
  const iconOf = (_spellId: number, stem?: string): string | null => iconUrl(activeBuild.build, stem);
  /**
   * The chart marks the scope's own deaths and cooldowns: under the group scope an outsider's death
   * is not drawn over the group's lines. Every other source marks as before (rateMarkers narrows a
   * single player itself).
   */
  function inMarkedScope(guid: string): boolean {
    return state.source !== SOURCE_FRIENDLIES || !groupScopeNow.scoped || groupScopeNow.group.has(guid);
  }
  const mainMarkers = $derived(
    base === null || nightMode
      ? []
      : rateMarkers({
          deaths: base.deaths.filter((death) => inMarkedScope(death.guid)),
          casts: base.casts.filter((row) => inMarkedScope(row.owner_guid ?? row.guid)),
          sourcePlayer,
          deathHref,
          iconOf,
        }),
  );
  const threatPlayers = $derived(
    new Map<string, ThreatPlayer>(
      (base?.roster ?? []).map((row) => [
        row.guid,
        {
          name: row.name,
          className: classOf.get(row.guid),
          role: row.role === 'tank' || row.role === 'healer' || row.role === 'dps' ? row.role : 'other',
        },
      ]),
    ),
  );
  const chartSyncKey = $derived(`${reportId}:${state.fight}`);
  // The chart chunk loads in parallel with the summary, so a chart never waits on it (spec §B11).
  $effect(() => {
    loadChartChunk().catch(() => {
      /* the chart shows its own error state and Try again */
    });
  });
  // Whole means the cut window too: "ignore events after a death" narrows what the tables
  // sum just as a brush does, and a narrowed table is measured, not prorated.
  const windowIsWhole = $derived(base !== null && isFullWindow(cutWindow, base.duration_ms));

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
  /** Set when the url named a fight the report does not have, cleared on the next pick. */
  let missingFight = $state<number | null>(null);
  /** Set when the url named a source the report has no unit for; cleared on the next pick. */
  let unknownSource = $state<string | null>(null);
  /** Set when the url named an ability this tab has no row for; cleared on the next pick. */
  let unknownAbility = $state<string | null>(null);
  // An ability named in the url rather than numbered -- the dropdown's own words, pasted
  // into a link -- resolves to the id of that name in this tab; anything else is said out
  // loud, instead of the whole table showing under a filter that silently did nothing.
  $effect(() => {
    const name = state.abilityName;
    if (name === '' || scoped === null) return;
    const wanted = name.toLowerCase();
    const found = abilityOptions(tabSource).find((option) => option.name.toLowerCase() === wanted);
    if (found !== undefined) {
      patch({ ability: Number.parseInt(found.id, 10), abilityName: '' });
      return;
    }
    unknownAbility = name;
    patch({ abilityName: '' });
  });
  // A source that is not a scope word and not a GUID the report knows: a name typed by
  // hand resolves to the player of that name; anything else is said out loud and the
  // scope falls back to everyone, instead of every tab answering "nothing" as a fact.
  $effect(() => {
    const source = state.source;
    if (
      file === null ||
      source === SOURCE_FRIENDLIES ||
      source === SOURCE_EVERYONE ||
      source === SOURCE_ENEMIES
    )
      return;
    if (unitNames.has(source)) return;
    const wanted = source.toLowerCase();
    // Players first: a boss's mirror images borrow players' names, and a typed name means
    // the player, not the clone that happened to be logged first.
    const matches = (unit: { name: string }): boolean =>
      splitUnitName(unit.name).name.toLowerCase() === wanted || unit.name.toLowerCase() === wanted;
    const units = file.units ?? [];
    const named = units.find((unit) => playerSet.has(unit.guid) && matches(unit)) ?? units.find(matches);
    if (named !== undefined) {
      patch({ source: named.guid });
      return;
    }
    unknownSource = source;
    patch({ source: SOURCE_FRIENDLIES });
  });

  // Two independent sources of scaling on an Actor-shaped row: the window (its share of
  // the actor's series, window.ts) and a target or boss filter (its share of the actor's
  // targets, filters.ts's `targetShare`). Either one alone is enough to make the
  // per-ability and per-target splits approximate, so a whole-fight window with "Boss
  // damage only" engaged still needs the mark. Ability, players-only, count-overkill and
  // ignore-after-death do not scale anything -- they subset or add exact figures -- so
  // they are not part of this. Over-claiming (marking a row that individually happens to
  // be exact) is the safe direction here; under-claiming is not.
  // An ability filter counts too: the summary cannot split sources or mitigation by one
  // ability, and the measured path can.
  const filtersScale = $derived(filters.target !== '' || filters.bossOnly || filters.ability !== null);
  /**
   * The night's totals are prorated only by a target or boss filter: an ability's total
   * is exact in every pull's summary and exact in their sum, so an ability filter alone
   * leaves the night's amounts bare.
   */
  const nightProrates = $derived(nightMode && (filters.target !== '' || filters.bossOnly));
  const actorTableApproximate = $derived(
    !windowIsWhole || filtersScale || ignoringDead || filters.countOverkill,
  );
  let percentiles = $state(new Map<string, Placement>());
  /** The last percentile load hit a failure (not a 404): the empty cells mean "could not ask". */
  let percentilesUnavailable = $state(false);
  /** A percentile load is in flight: the empty cells mean "asking", not "nothing ranked". */
  let percentilesPending = $state(false);
  /** Bumped by the retry line; the percentile effect reads it, so a bump asks again. */
  let parseAttempt = $state(0);
  const loader = createPercentileLoader();

  const filterContext = $derived({
    // The whole night's bosses are every encounter's; one pull's is its own; trash has none.
    bosses: nightMode
      ? bossGuidsOf(
          file?.units ?? [],
          (combined ?? []).filter((entry) => entry.kind === 'encounter').map((entry) => entry.name),
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
    trashMode || (!nightMode && fight !== null && fight.kind !== 'encounter' && !fight.in_progress)
      ? TRASH_PARSE
      : fight?.encounter_id === undefined || fight.in_progress
        ? ''
        : !windowIsWhole
          ? 'window'
          : filtersScale
            ? 'filter'
            : fight.kill
              ? percentilesPending
                ? '…'
                : percentilesUnavailable
                  ? '?'
                  : '–'
              : 'wipe',
  );
  /** The actor tables' fallback: Damage Taken has no parse at all, and its cells say so. */
  const tableParseFallback = $derived(
    trashMode ? TRASH_PARSE : nightMode ? 'night' : state.tab === 'damage-taken' ? 'none' : parseFallback,
  );

  /** GUID to class, for the tables whose rows are not Actors. Off the whole fight rather
      than the scoped one: a colour is a property of the player, not of the current source
      scope, and Mechanics draws the whole raid however the scope is set. */
  const unitClasses = $derived(unitClassMap(file?.units ?? []));
  const credit = $derived(reportCredit(meta));
  /** The heading: leads with the logging character when the report is credited. */
  const heading = $derived(
    meta === null ? '' : reportHeading(meta, credit.credited ? credit.name : undefined),
  );
  /** The page clock, re-read on every poll tick so "quiet" is judged against now, not against when the page loaded. */
  let now = $state(Date.now());
  /** Minutes since the companion was last heard, once past the threshold; null while it is fresh or unknown. */
  const quiet = $derived(quietMinutes(meta, now));
  $effect(() => {
    // A credited report names its character in the tab as well; the others keep the title the shell gave them.
    if (credit.credited && heading !== '') document.title = pageTitle(`/reports/${reportId}`, heading);
  });
  const classOf = $derived(
    new Map(
      (base?.roster ?? [])
        .map((row) => [row.guid, row.class || unitClasses.get(row.guid)?.class] as const)
        .filter((entry): entry is [string, string] => Boolean(entry[1])),
    ),
  );
  // The log writes each player's realm after the name; the report's own realm is the one most
  // of them carry, and only a player from elsewhere shows theirs.
  $effect(() => {
    setHomeRealm(
      homeRealmOf((file?.units ?? []).filter((unit) => playerSet.has(unit.guid)).map((unit) => unit.name)),
    );
  });

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
  /** The tab's rows under the source scope, before the filters: what the filter pickers list. */
  const tabSource = $derived.by((): Actor[] => {
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
      scopeKey === SOURCE_FRIENDLIES
        ? source.filter((actor) => filterContext.players.has(actor.guid))
        : state.source === 'enemies'
          ? source.filter((actor) => !friendlySet.has(actor.guid))
          : source.filter((actor) => actor.guid === state.source);
    return bySource;
  });

  /** Which Actor[] the current tab shows: the source scope, then the filters, then any measure. */
  const tabActors = $derived.by(() => {
    // "Boss damage only" has no meaning for healing, whose targets are players: applied
    // there it blanked the table.
    const applied = state.tab === 'healing' ? { ...filters, bossOnly: false } : filters;
    const filtered = applyActorFilters(tabSource, applied, filterContext, tableExact !== null);
    if (tableExact === null) return filtered;
    // Measured rows: the events' own totals and pairs replace the prorated ones, and a
    // row the events do not mention did nothing in this window to these targets.
    const measured = tableExact;
    return filtered
      .map((actor): Actor => {
        const found = measured.get(actor.guid);
        return {
          ...actor,
          total: found?.total ?? 0,
          effective: found?.effective ?? 0,
          overheal: actor.overheal === undefined ? undefined : (found?.overheal ?? 0),
          targets: found?.targets ?? [],
          measured: true,
          mitigated: found?.mitigated ?? { absorbed: 0, blocked: 0, misses: {} },
          // The summary's active time counts a corpse's HoT ticks; with the dead spans left
          // out, active time can be at most the time alive in the window.
          active_ms: ignoringDead
            ? Math.min(Math.max(0, actor.active_ms - (found?.deadActiveMs ?? 0)), aliveMs(actor.guid))
            : actor.active_ms,
        };
      })
      .filter((actor) => actor.total > 0 || actor.effective > 0)
      .sort((a, b) => b.effective - a.effective);
  });

  /**
   * Players with no row in this table: under a death window that is usually someone who
   * was already dead, and a table that silently lacks them reads as if they never existed.
   */
  const filtersNarrow = $derived(filters.target !== '' || filters.ability !== null || filters.bossOnly);
  const absentPlayers = $derived.by((): { name: string; deadSince: number | null }[] => {
    if (scoped === null || (windowIsWhole && !filtersNarrow)) return [];
    const shown = new Set(tabActors.map((actor) => actor.guid));
    return scoped.roster
      .filter((row) => playerSet.has(row.guid) && !shown.has(row.guid))
      .filter((row) => inSource(row.guid, scopeKey, scopePlayers, friendlySet))
      .map((row) => ({
        name: splitUnitName(row.name).name,
        deadSince: deadSince(row.guid, cutWindow.endMs),
      }));
  });
  /** Rows of players who were dead at the window's end: a HoT ticking on a corpse is not a live row. */
  const deadRows = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const out = new Map<string, number>();
    if (windowIsWhole || nightMode) return out;
    for (const actor of tabActors) {
      if (!playerSet.has(actor.guid)) continue;
      const since = deadSince(actor.guid, cutWindow.endMs);
      if (since !== null) out.set(actor.guid, since);
    }
    return out;
  });
  /** When this player died before `at` and cast nothing since, the instant they died; else null. */
  function deadSince(guid: string, at: number): number | null {
    if (base === null) return null;
    const death = [...base.deaths].filter((entry) => entry.guid === guid && entry.at_ms <= at).pop();
    if (death === undefined) return null;
    const castSince = base.casts
      .filter((row) => row.guid === guid)
      .some((row) => row.sequence.some((cast) => cast > death.at_ms && cast <= at));
    return castSince ? null : death.at_ms;
  }

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
    void parseAttempt;
    // A fight still being written is not a parse, for the same reason a brushed window is
    // not: the numbers are a partial fight's, and the cache key carries the rounded dps,
    // which moves on every five-second tick -- so a live 25-player pull asked for 25 fresh
    // t-digest lookups every five seconds, per viewer, and never hit the cache once.
    // Only kills are ranked, so only kills are asked about: a wipe's dps placed among the
    // kills read as "0" beside a tooltip saying wipes have no parse.
    if (
      summary === null ||
      encounterId === undefined ||
      fight?.kind !== 'encounter' ||
      !windowIsWhole ||
      fight?.in_progress === true ||
      fight?.kill !== true
    ) {
      percentiles = new Map();
      percentilesUnavailable = false;
      percentilesPending = false;
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

    percentilesPending = true;
    void loader.load(wanted.map((entry) => entry.query)).then(({ placements, unavailable }) => {
      if (state.fight !== wantedFight || state.tab !== wantedTab) return;
      percentilesPending = false;
      percentilesUnavailable = unavailable;
      // A plain Map, not SvelteMap: this is a throwaway local built up once and then
      // assigned whole to `percentiles` (already $state) below, the same reasoning the
      // `summaries` cache above gives for its own eslint-disable.
      // eslint-disable-next-line svelte/prefer-svelte-reactivity
      const next = new Map<string, Placement>();
      for (const entry of wanted) {
        const found = placements.get(percentileKey(entry.query));
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

  function patch(raw: Partial<ReportState>): void {
    // A sibling still speaking the old words ({ mode, view, tab }) lands on the tab it means.
    const next: Partial<ReportState> = legacyTab(raw);
    // Changing fight drops the window: milliseconds from one fight's start mean nothing in
    // another's.
    const changesFight = next.fight !== undefined && next.fight !== state.fight;
    if (next.fight !== undefined) missingFight = null;
    if (next.source !== undefined && next.source !== SOURCE_FRIENDLIES) unknownSource = null;
    if (next.ability !== undefined) unknownAbility = null;
    // Picking a fight, mode, view, tab or source is a page in its own right, so it is
    // pushed and the back button undoes it. Brushing the chart replaces: a drag writes
    // the url on every pointer move, and pushing those would bury the real back
    // destination under a hundred entries.
    const pushes =
      (['fight', 'tab', 'source'] as const).some(
        (key) => next[key] !== undefined && next[key] !== state[key],
      ) ||
      // The first window set on a whole fight is pushed too, so Back returns to the
      // whole fight rather than to the fight before it; moving an existing window
      // replaces, since a drag writes on every pointer move.
      (next.start !== undefined && next.start !== null && state.start === null);
    state = withState(state, changesFight ? { ...next, start: null, end: null } : next);
    writeUrl(pushes);
  }

  async function loadReport(continuing = false, known: ReportMeta | null = null): Promise<void> {
    if (!continuing) status = 'loading';
    error = '';
    try {
      // inlineMeta, not `meta`: reading the state this function also writes would make it
      // a dependency of the effect below and run the whole load a second time.
      let resolved: ReportMeta;
      const uploadedAt = justUploadedAt(reportId);
      if (known !== null) {
        // A poll has just read a newer meta than the one the page shipped with.
        resolved = known;
      } else if (inlineMeta !== null) {
        resolved = inlineMeta;
      } else if (uploadedAt === null) {
        resolved = await fetchReportMeta(reportId);
      } else {
        // Just uploaded from this tab: a 404 is "not visible yet" for a short while, and
        // the page shows the processing screen meanwhile rather than an error.
        resolved = await fetchReportMetaAfterUpload(reportId, {
          uploadedAt,
          onWaiting: () => {
            if (status === 'loading') status = 'processing';
          },
        });
      }
      meta = resolved;
      // A report still being parsed has no files to read; a failed one never will.
      if (resolved.status === 'processing') {
        status = 'processing';
        return;
      }
      if (resolved.status === 'failed') {
        status = 'unparsed';
        return;
      }
      let base = resolved.data_base_url;
      // Private and guild reports are refused at /logs-data/, so they carry a signed base
      // url the API hands out to people who may read them.
      if (resolved.visibility === 'private' || resolved.visibility === 'guild') {
        base = await fetchAccessUrl(reportId);
      }
      dataBase = base;
      if (summaryFirst && continuing === false && resolved.fights.length > 0) {
        // untracked: this runs before the first await, inside the load effect, and these
        // read and write the state the effect must not depend on.
        untrack(() => readUrl());
        // The whole night (fight=all) and a zone's trash have no summary.json of their own;
        // the combined-view effect folds their pulls once the page is ready.
        if (untrack(() => fight) !== null) await untrack(() => loadFight(state.fight));
        status = 'ready';
        const rest = await fetchReportFileFor(resolved.status, base);
        if (rest !== null) file = rest;
        return;
      }
      const loaded = await fetchReportFileFor(resolved.status, base);
      if (loaded === null) {
        // A live report with no fight uploaded yet: the poll below keeps asking for report.json.
        file = null;
        status = 'waiting';
        return;
      }
      file = loaded;
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
   * The parse poll: while the report is processing, ask for it again every five seconds
   * (the answer is never cached while processing, see load.ts) and carry on into the normal
   * load the moment it is anything else, so the page fills in without a reload. Stops when
   * `status` leaves 'processing' (this effect's cleanup) or the component unmounts.
   */
  $effect(() => {
    if (status !== 'processing' || meta === null) return;
    const poller = createPoller(async () => {
      const next = await fetchReportMeta(reportId);
      meta = next;
      if (next.status !== 'processing') await loadReport(true, next);
    }, POLL_INTERVAL_MS);
    poller.start();
    return () => poller.stop();
  });

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
    const loaded = await fromDataBase((base) => fetchSummary(base, index, engineVersion));
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
    // The shell's reservation is sized for a full report; a report with no fights has none.
    if (status === 'processing' || status === 'waiting' || status === 'unparsed') {
      document
        .getElementById('report')
        ?.classList.remove('min-h-[1420px]', 'md:min-h-[1780px]', 'lg:min-h-[1480px]');
    }
  });

  // Re-read the URL when the visitor uses the browser's own back and forward.
  $effect(() => {
    const onPop = (): void => {
      readUrl();
      if (combinedFights(fights, state.fight) === null) void loadFight(state.fight).catch(() => {});
    };
    window.addEventListener('popstate', onPop);
    return () => window.removeEventListener('popstate', onPop);
  });

  // A combined view (the whole night's boss pulls, or one zone's trash): every pull's summary,
  // a few at a time, folded into `night` as they land so the table fills in rather than waiting
  // for the slowest fetch. Pulls the cache already holds cost nothing; the ones this fetches are
  // cached for the fight views. Re-run when the selection returns here after a pull was opened,
  // which is free once everything is cached.
  $effect(() => {
    if (status !== 'ready' || combined === null) return;
    const encounters = combined.filter((entry) => !entry.in_progress);
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
          const loaded = await fromDataBase((base) => fetchSummary(base, index as number, engineVersion));
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
   * new poller starts; and a failed report load or a report with no data base yet never
   * starts one. A closed report never enters this effect at all, so it costs nothing.
   *
   * The waiting screen (a live report whose first fight has not been uploaded) rides the
   * same poller rather than a second timer: its tick asks for report.json and, once that
   * exists, hands over to a normal load.
   */
  $effect(() => {
    const waiting = status === 'waiting';
    if (dataBase === '' || !(waiting || (isLive && status === 'ready'))) return;

    const poller = createPoller(async () => {
      now = Date.now();
      if (waiting) {
        const first = await fromDataBase((base) => fetchReportFileFor(meta?.status ?? 'live', base, 'fresh'));
        if (first !== null) {
          await loadReport(true, meta);
          return;
        }
        // Nothing yet. The meta may have moved on (the log ended with no fight in it, or
        // failed), and then the waiting screen would be a lie.
        const latest = await fetchReportMeta(reportId);
        meta = latest;
        if (latest.status === 'failed') status = 'unparsed';
        else if (latest.status === 'complete') await loadReport(true, latest);
        return;
      }
      // report.json first: it is what turns a live fight into a closed one and adds the
      // next fight to the selector.
      const next = await fromDataBase((base) => fetchReportFile(base, 'fresh'));
      file = next;
      // The report's own status moves too (live to complete); it is not cached while live.
      const latest = await fetchReportMeta(reportId).catch(() => null);
      if (meta !== null) meta = { ...(latest ?? meta), fights: next.fights };

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

  // Each of these mirrors the `{#if}` its component renders under, below: the import starts
  // the moment the mode or view it belongs to is actually selected, not before, and `load()`
  // is a no-op once the chunk has resolved so re-entering a mode already visited this session
  // costs nothing further.
  $effect(() => {
    if (state.tab === 'compare' && summary !== null && !nightMode) compareModeLazy.load();
  });
  $effect(() => {
    if (state.tab === 'mechanics' && base !== null) mechanicsModeLazy.load();
  });
  $effect(() => {
    if (scoped !== null && !nightMode && state.tab === 'timeline') timelinesViewLazy.load();
  });
  $effect(() => {
    if (scoped !== null && !nightMode && state.tab === 'events') eventsViewLazy.load();
  });
  $effect(() => {
    if (scoped !== null && !nightMode && state.tab === 'queries') queriesViewLazy.load();
  });
  $effect(() => {
    if (scoped !== null && !nightMode && state.tab === 'rating') ratingTabLazy.load();
  });
  $effect(() => {
    if (scoped === null) return;
    if (state.tab === 'threat') threatTableLazy.load();
    if (state.tab === 'casts') castTableLazy.load();
    if (state.tab === 'buffs' || state.tab === 'debuffs') auraTabLazy.load();
    if (state.tab === 'interrupts' || state.tab === 'dispels') exchangeTableLazy.load();
    if (state.tab === 'resources') resourceGraphsLazy.load();
  });
  /** The Rating tab's answer, for its rail (RatingRail). */
  let ratingState = $state<{ data: ReportRatings | null; status: ReportRatingsStatus }>({
    data: null,
    status: 'idle',
  });

  // ---- the frame (logs spec 2026-11-07 6.1) ----
  /** The one-line notices under the bar (section 6.1): each says what the page did instead. */
  const notices = $derived.by((): { id: string; text: string }[] => {
    const out: { id: string; text: string }[] = [];
    if (unknownSource !== null)
      out.push({
        id: 'report-unknown-source',
        text: `This report has no player or unit called ${unknownSource}, so every friendly is showing.`,
      });
    if (unknownAbility !== null)
      out.push({
        id: 'report-unknown-ability',
        text: `This tab has no ability called ${unknownAbility}, so every ability is showing.`,
      });
    if (missingFight !== null)
      out.push({
        id: 'report-missing-fight',
        text:
          missingFight === MISSING_FIGHT
            ? 'This link names a fight the page cannot read, so its opening pull is showing.'
            : trashScopeKey(missingFight) !== null
              ? `This report has no trash group starting at fight ${trashScopeKey(missingFight)}, so its opening pull is showing.`
              : `This report has no fight ${missingFight}, so its opening pull is showing.`,
      });
    // A pull the logger's group was not in reads everyone in range (lib/report/group.ts), and
    // says so, so the scope word over the table is never a surprise.
    if (groupMissing) out.push({ id: 'report-group-missing', text: reportCopy.groupMissing });
    if (error !== '') out.push({ id: 'report-fight-error', text: error });
    if (parsesUnknownHere) out.push({ id: 'parse-unavailable-line', text: '' });
    return out.filter((notice) => notice.text !== '' || notice.id === 'parse-unavailable-line');
  });
  /** The tabs whose tables use the glossary's words (Parse, Active, Share, threat, uptime). */
  const GLOSSARY_TABS: ReadonlySet<string> = new Set([
    'summary',
    'damage-done',
    'healing',
    'damage-taken',
    'deaths',
    'threat',
    'casts',
    'buffs',
    'debuffs',
    'interrupts',
    'dispels',
    'resources',
  ]);
  /** The tables' Parse column: a boss pull's, and gone while the rankings did not answer, so a
   *  column of question marks never stands in for figures; the line over the table says so and
   *  offers Try again. */
  const parseColumn = $derived(showParse && fight?.kind === 'encounter' && !percentilesUnavailable);
  /** The tabs with a Parse column: the line that parses did not load shows on them alone. */
  const PARSE_TABS: ReadonlySet<string> = new Set(['summary', 'damage-done', 'healing']);
  const parsesUnknownHere = $derived(
    percentilesUnavailable && showParse && fight?.kind === 'encounter' && PARSE_TABS.has(state.tab),
  );
  /** The report has a logger's group, and this pull ran without any of it. */
  const groupMissing = $derived(
    !nightMode &&
      fight !== null &&
      state.source === SOURCE_FRIENDLIES &&
      groupScopeNow.noun === 'everyone' &&
      (fights.some((entry) => (entry.group?.length ?? 0) > 0) || (meta?.group?.length ?? 0) > 0),
  );
  /** "2 minutes ago": how long a report has been processing. */
  const uploadedAge = $derived.by(() => {
    const at = Date.parse(meta?.created_at ?? '');
    if (Number.isNaN(at)) return 'just now';
    const minutes = Math.max(0, Math.round((now - at) / 60_000));
    return minutes === 0 ? 'just now' : `${minutes} ${minutes === 1 ? 'minute' : 'minutes'} ago`;
  });
  /** The Coach tab is listed only for a signed-in owner of a character in the pull, or an officer. */
  const coachViewer = createCoachViewerState();
  const tabs = $derived(listedTabs(coachTabListed(coachViewer.viewer, roster)));
  const tabXl = $derived(TABS_XL - (tabs.some((tab) => tab.id === 'coach') ? 0 : 1));
  /** A url on this page with `next` applied: what every in-page link carries. */
  function hrefWith(next: Partial<ReportState>): string {
    return `${typeof window === 'undefined' ? '' : window.location.pathname}${reportSearch(withState(state, next), firstFight) || '?'}`;
  }
  const tabItems = $derived(
    tabs.map((tab) => ({
      id: tab.id,
      label: tab.label,
      href: hrefWith({ tab: tab.id }),
      testid: `tab-${tab.id}`,
    })),
  );
  const playerHref = (guid: string): string => hrefWith({ source: guid });
  /** At pull's Gear link: the player's own view on Damage done, whose "Gear at the pull" is the
   *  paperdoll (PlayerView). A source on the Summary tab alone scopes the summary and shows no gear. */
  const gearHref = (guid: string): string => hrefWith({ source: guid, tab: 'damage-done' });
  /** The Source list's players: the scope's own (the group on a crowded log, everyone in range under that word). */
  const sourceRoster = $derived(
    state.source === SOURCE_ENEMIES
      ? roster
      : roster.filter(
          (row) => inSource(row.guid, scopeKey, scopePlayers, friendlySet) || row.guid === state.source,
        ),
  );
  /** The players and deaths in the Source scope, for the band's answer and its players tail. */
  const scopeRoster = $derived(
    (wholeScoped?.roster ?? []).filter((row) => inSource(row.guid, scopeKey, scopePlayers, friendlySet)),
  );
  const scopeDeaths = $derived(
    (wholeScoped?.deaths ?? []).filter((death) => inSource(death.guid, scopeKey, scopePlayers, friendlySet)),
  );
  /** The rows the band's figure sums: the source player alone on their view, else the scope. */
  const figureRows = $derived(
    sourcePlayer !== undefined ? scopeRoster.filter((row) => row.guid === sourcePlayer) : scopeRoster,
  );
  const figure = $derived(
    fight === null || nightMode || summary === null
      ? null
      : reportFigure({
          tab: state.tab,
          player: playerView,
          bossPull: !isTrashPull,
          damageDone: figureRows.reduce((sum, row) => sum + row.damage_done, 0),
          healing: figureRows.reduce((sum, row) => sum + row.healing_done, 0),
          damageTaken: figureRows.reduce((sum, row) => sum + row.damage_taken, 0),
          deaths: scopeDeaths.length,
          durationMs: summary.duration_ms,
        }),
  );
  const pullFacts = $derived(
    fight === null || nightMode ? [] : thisPullFacts({ fight, phases: summary?.phases ?? [] }),
  );
  const isTrashPull = $derived(fight !== null && fight.kind !== 'encounter');
  /** A pull a rating can exist for: a boss pull of a report the API keeps (never the bundled sample). */
  const rateable = $derived(!isTrashPull && !isSampleReport(reportId));
  const live = $derived<'live' | 'quiet' | null>(fightIsLive ? (quiet === null ? 'live' : 'quiet') : null);
  // The player on screen when one is the Source, else the fight's own first-DPS rule.
  const simHref = $derived(
    fight === null || nightMode ? null : simCompareFightHref(reportId, state.fight, sourcePlayer),
  );
  /** The band's meta line: the pull on screen, or the night (zone, date, fight count). */
  const headLine = $derived(
    fight !== null && !nightMode
      ? reportMeta(facts.zone, fight, pullNumbers(fights).get(fight.index)?.pull)
      : [facts.zone, fights.length > 0 ? facts.date : '', facts.fights]
          .filter((part) => part !== '')
          .join(' · '),
  );
  /** The meta's last fact (players, or mobs died): a desktop fact, so a phone's meta is one line. */
  const headTail = $derived(fight !== null && !nightMode ? reportMetaTail(fight, scopeRoster.length) : '');
  /** One player as the Source on an actor tab: the player's own view (decision R5). */
  const playerView = $derived(sourcePlayer !== undefined && ACTOR_TABS.has(state.tab) && !nightMode);
  const playerRow = $derived(base?.roster.find((row) => row.guid === sourcePlayer));
  const rosterByGuid = $derived(new Map((base?.roster ?? []).map((row) => [row.guid, row])));
  const tabTitle = $derived<'Targets' | 'Healed' | 'Sources'>(
    state.tab === 'healing' ? 'Healed' : state.tab === 'damage-taken' ? 'Sources' : 'Targets',
  );
  /** A night shows only the tabs that fold across pulls (section 6.3); the rest say pick a pull
   *  (Coach says so itself, in its own words). */
  const onePullTab = $derived(nightMode && !NIGHT_TABS.has(state.tab) && state.tab !== 'coach');
  /** One player's worn set at the pull as a planner build, for Gear at the pull. */
  function plannerHrefFor(guid: string | undefined): string | null {
    const combatant = base?.combatants.find((row) => row.guid === guid);
    if (combatant === undefined) return null;
    const className = classOf.get(combatant.guid);
    return (
      plannerLinkFor({
        dataBuild: activeBuild.build,
        className,
        treeSizes: treeSizesFor(className),
        combatant,
      })?.href ?? null
    );
  }
  /** The tabs the sibling lane draws: their main content in the main column, This pull in the rail. */
  let spellIcon = $state<SpellIconOf>(NO_SPELL_ICONS);
  $effect(() => {
    void loadSpellIcons(activeBuild.build, { chain: true }).then((icons) => {
      spellIcon = icons;
    });
  });
  /** A same-page link (`?tab=…`, `?source=…`) moves in place: no reload, Back still undoes it. */
  function followInPage(event: MouseEvent): void {
    if (event.defaultPrevented || event.button !== 0 || event.metaKey || event.ctrlKey || event.shiftKey)
      return;
    const anchor = (event.target as Element | null)?.closest('a');
    const href = anchor?.getAttribute('href') ?? '';
    if (anchor === null || anchor === undefined || anchor.target !== '') return;
    const url = new URL(href, window.location.href);
    if (url.pathname !== window.location.pathname || url.origin !== window.location.origin) return;
    event.preventDefault();
    const parsed = parseReportState(url.search, firstFight);
    patch({ ...parsed, start: parsed.start, end: parsed.end });
  }
</script>

{#snippet lazyFallback(lazy: LazyLoadState, minHeight: string)}
  <!-- While a tab's chunk is in flight a Skeleton holds its place; a failed chunk says so inside
       the Plate with Try again, a fresh attempt (lazy-component.svelte.ts). -->
  {#if lazy.error !== ''}
    <Plate label={reportCopy.viewFailed}>
      <div class="px-4 py-3">
        <LoadError message={reportCopy.viewFailed} onRetry={() => lazy.load()} testid="lazy-view-error" />
      </div>
    </Plate>
  {:else}
    <Skeleton {minHeight} testid="lazy-view-skeleton" />
  {/if}
{/snippet}

{#snippet thisPull()}
  {#if nightMode && night !== null}
    <ThisPull
      title={trashMode ? 'This trash' : 'The night'}
      facts={nightFacts(night, trashMode, fights.filter((entry) => entry.kind !== 'encounter').length)}
      simHref={null}
    />
  {:else if fight !== null && !nightMode}
    <ThisPull facts={pullFacts} {simHref} />
  {/if}
{/snippet}

{#snippet chart(metric: 'damage-done' | 'healing' | 'damage-taken')}
  {#if summary !== null && !nightMode}
    <TimeChart
      {metric}
      actors={wholeChartActorsFor(metric)}
      roster={summary.roster}
      durationMs={summary.duration_ms}
      window={timeWindow}
      phases={chartPhases}
      events={mainMarkers}
      {deadSpans}
      {brushShape}
      onWindow={setWindow}
      {sourcePlayer}
      enemies={state.source === SOURCE_ENEMIES}
      {youGuid}
      chosen={state.players}
      preset={state.playersPreset === null ? 'split' : state.playersPreset}
      onChoose={(players) => patch({ players })}
      onPreset={(preset) => patch({ players: null, playersPreset: preset === 'split' ? null : preset })}
      logScale={state.logScale}
      onLogScale={(logScale) => patch({ logScale })}
      targetName={state.target === ''
        ? undefined
        : splitUnitName(unitNames.get(state.target) ?? state.target).name}
      syncKey={chartSyncKey}
      onOpenDeath={openDeaths}
      extra={chartedAbility === null ? null : { label: chartedAbility.label, series: chartedAbility.series }}
      {noun}
      openOnPhone={state.tab === 'deaths'}
    />
  {/if}
{/snippet}

{#snippet stateHead(title: string, line: string)}
  <PageHeader {title} titleTestid="report-title">
    {#snippet meta()}
      {#if status === 'waiting'}<Tag tone="success" testid="report-live"
          >{quiet === null ? reportCopy.liveTag : reportCopy.quietBadge}</Tag
        ><span>{reportCopy.loggingNow}</span>{:else if line !== ''}<span data-testid="report-subtitle"
          >{line}</span
        >{/if}
    {/snippet}
    {#snippet tabs()}
      <TabStripMore
        tabs={[
          {
            id: 'summary',
            label: 'Summary',
            href: '#',
            disabled: true,
            attrs: {
              title:
                status === 'processing'
                  ? reportCopy.parsing
                  : status === 'waiting'
                    ? reportCopy.waiting
                    : reportCopy.parseFailed,
            },
          },
        ]}
        active=""
        md={1}
        xl={1}
        label={reportCopy.tabsLabel}
        testid="report-tabs-disabled"
      />
    {/snippet}
  </PageHeader>
{/snippet}

{#if status === 'failed'}
  <Page layout="ledger" testid="report-failed">
    {#snippet header()}<PageHeader title={reportCopy.eyebrow} />{/snippet}
    <Plate label={reportCopy.eyebrow}>
      <div
        class="flex flex-wrap items-center gap-3 px-4 py-4 text-[14px]"
        role="alert"
        data-testid="report-error"
      >
        {#if error === REPORT_NOT_FOUND}
          <span class="text-muted">{reportCopy.reportMissing}</span>
          <a class="text-gold font-semibold" href="/reports">{reportCopy.openYourReports} ›</a>
        {:else}
          <span class="text-muted"
            >{error === '' || error === REPORT_LOAD_FAILED ? reportCopy.reportFailed : error}</span
          >
          <button
            type="button"
            class={buttonClass('secondary', 'sm')}
            data-testid="report-error-retry"
            onclick={() => void loadReport()}>{reportCopy.tryAgain}</button
          >
        {/if}
      </div>
    </Plate>
  </Page>
{:else if status === 'processing' || status === 'waiting' || status === 'unparsed'}
  <div class="reveal" data-testid="report-processing" data-state={status}>
    <Page layout="ledger">
      {#snippet header()}
        {@render stateHead(
          status === 'waiting' && meta !== null
            ? heading
            : meta?.title || meta?.zone || reportCopy.yourUpload,
          status === 'processing'
            ? `Uploaded ${uploadedAge}`
            : status === 'unparsed'
              ? (meta?.zone ?? '')
              : '',
        )}
      {/snippet}
      <Plate label={reportCopy.fightsLabel}>
        <div class="px-4 py-4 text-[14px]" data-testid="report-fights-strip">
          {#if status === 'processing' || status === 'waiting'}
            <p
              class="flex items-center gap-3"
              role="status"
              data-testid={status === 'waiting' ? 'report-waiting' : 'report-parsing'}
            >
              {#if !(status === 'waiting' && quiet !== null)}<span data-testid="report-spinner"
                  ><Spinner /></span
                >{/if}
              {status === 'waiting'
                ? quiet === null
                  ? reportCopy.waiting
                  : reportCopy.quiet(quiet)
                : reportCopy.parsing}
            </p>
          {:else}
            <p role="alert" data-testid="report-parse-failed">
              {reportCopy.parseFailed}
              <a class="text-gold font-semibold" href="/setup#logging" data-testid="report-parse-failed-link"
                >{reportCopy.parseFailedLink}</a
              >
            </p>
          {/if}
        </div>
      </Plate>
    </Page>
  </div>
{:else if status === 'loading' || meta === null}
  <ReportSkeleton title={meta === null ? undefined : heading} />
{:else}
  <!-- The ready state's own wrapper, the only place `.reveal` may sit. A same-page link inside
       it moves in place (followInPage). -->
  <!-- svelte-ignore a11y_click_events_have_key_events, a11y_no_static_element_interactions -->
  <div class="reveal" onclick={followInPage}>
    <Page layout="split" panels testid="report-page">
      {#snippet header()}
        <ReportHead
          title={heading}
          line={headLine}
          lineTail={headTail}
          {simHref}
          pullActions={fight !== null && !nightMode}
          sample={isSampleReport(reportId)}
          {credit}
          {live}
          {figure}
        >
          {#snippet tabs()}
            <!-- The wrapper only scopes the phone strip's edge fade (below), never a box of its own. -->
            <div class="report-tabs contents">
              <TabStripMore
                tabs={tabItems}
                active={state.tab}
                md={TABS_MD}
                xl={tabXl}
                label={reportCopy.tabsLabel}
                moreLabel={reportCopy.more}
                testid="report-tabs"
                onSelect={(tab) => patch({ tab: tab as ReportState['tab'] })}
              />
            </div>
          {/snippet}
        </ReportHead>
      {/snippet}
      {#snippet bar()}
        <ReportBar
          {state}
          {fights}
          roster={sourceRoster}
          scope={groupScopeNow}
          window={timeWindow}
          durationMs={base?.duration_ms ?? 0}
          {presets}
          actorTab={ACTOR_TABS.has(state.tab) && !playerView}
          {filters}
          filterActors={tabSource}
          afterDeathAvailable={!nightMode}
          onPatch={patch}
          onWindow={setWindow}
          onFilters={setFilters}
        />
        {#if notices.length > 0}
          <div data-region={REGION.summary}>
            {#each notices.filter((notice) => notice.text !== '') as notice (notice.id)}
              <div class="px-[var(--page-inset)]">
                <Notice text={notice.text} tone="neutral" testid={notice.id} />
              </div>
            {/each}
            {#if parsesUnknownHere}
              <p class="px-[var(--page-inset)] pt-2 text-[15px]" data-testid="parse-unavailable">
                {reportCopy.parsesUnknown}
                <button type="button" class="text-gold ml-2 font-semibold" onclick={() => (parseAttempt += 1)}
                  >{reportCopy.tryAgain}</button
                >
              </p>
            {/if}
          </div>
        {/if}
      {/snippet}
      {#snippet rail()}
        {@render thisPull()}
        {#if !onePullTab && scoped !== null}
          {#if state.tab === 'summary' && !nightMode}
            {#if scoped.deaths.length > 0}
              <SummaryDeaths summary={scoped} tabHref={hrefWith({ tab: 'deaths' })} />
            {/if}
            {#if rateable}
              <RatingPanel
                {reportId}
                fightIndex={state.fight}
                roster={scoped.roster.map((row) => ({ guid: row.guid, name: row.name, class: row.class }))}
                tabHref={hrefWith({ tab: 'rating' })}
              />
            {/if}
            {#if !isTrashPull && summary !== null}
              <KeyMechanics {summary} {spellIcon} tabHref={hrefWith({ tab: 'mechanics' })} />
            {/if}
            <GroupBuffs
              summary={scoped}
              durationMs={scoped.duration_ms}
              players={playerSet}
              {spellIcon}
              tabHref={hrefWith({ tab: 'buffs' })}
            />
            {#if fight?.kind === 'encounter' && scoped.combatants.length > 0}
              <Section title="At pull" panel testid="at-pull-card">
                {#each scoped.combatants as combatant (combatant.guid)}
                  {@const label = unitLabel(combatant.name).text}
                  <Row
                    title={label}
                    sub={`${combatant.spec ?? classOf.get(combatant.guid) ?? ''} · ilvl ${combatant.item_level ?? 0}`}
                    open
                  >
                    {#snippet lead()}<Crest cls={classOf.get(combatant.guid)} size={28} />{/snippet}
                    {#snippet trail()}<span class="flex gap-3" data-testid="at-pull-links"
                        ><a
                          class="text-gold text-[14px] font-semibold [@media(pointer:coarse)]:-my-3 [@media(pointer:coarse)]:inline-block [@media(pointer:coarse)]:py-3"
                          href={gearHref(combatant.guid)}>Gear ›</a
                        ><a
                          class="text-gold text-[14px] font-semibold [@media(pointer:coarse)]:-my-3 [@media(pointer:coarse)]:inline-block [@media(pointer:coarse)]:py-3"
                          href={simLinkFor(reportId, state.fight, combatant.guid).href}
                          data-testid="combatant-sim-link">Sim ›</a
                        ></span
                      >{/snippet}
                  </Row>
                {/each}
              </Section>
            {/if}
            {#if rankingsHref !== null}
              <a
                class="text-gold text-[14px] font-semibold [@media(pointer:coarse)]:-my-3 [@media(pointer:coarse)]:inline-block [@media(pointer:coarse)]:py-3"
                href={rankingsHref}
                data-testid="report-rankings-link">Rankings for {fight?.name} ›</a
              >
            {/if}
          {:else if ACTOR_TABS.has(state.tab)}
            <TargetsRail title={tabTitle} actors={tabActors} />
          {:else if state.tab === 'threat'}
            <ThreatRail
              taunts={scoped.taunts}
              target={state.target}
              pairs={scoped.threat_by_target ?? []}
              names={unitNames}
              {classOf}
              sourceName={playerSet.has(state.source) ? unitNames.get(state.source) : undefined}
              approximate={!windowIsWhole}
              scopeNoun={nightMode ? 'night' : 'pull'}
              durationMs={summary?.duration_ms ?? scoped.duration_ms}
              onWindow={setWindow}
            />
          {:else if state.tab === 'casts' || state.tab === 'timeline'}
            <CastsRail
              casts={scoped.casts}
              players={scopePlayers}
              names={unitNames}
              {classOf}
              durationMs={scoped.duration_ms}
            />
          {:else if state.tab === 'debuffs'}
            <DispelsRail dispels={scoped.dispels} {spellIcon} tabHref={hrefWith({ tab: 'dispels' })} />
          {:else if state.tab === 'interrupts'}
            <EnemyCastsRail
              casts={(windowed ?? base)?.casts ?? []}
              players={friendlySet}
              schools={spellSchools(base)}
              {spellIcon}
            />
          {:else if state.tab === 'dispels'}
            <GroupDebuffsRail
              auras={(windowed ?? base)?.auras ?? []}
              players={scopePlayers}
              {spellIcon}
              tabHref={hrefWith({ tab: 'debuffs' })}
            />
          {:else if state.tab === 'queries'}
            <QueriesRail />
          {:else if state.tab === 'rating'}
            <RatingRail ratings={ratingState.data} {roster} />
          {:else if state.tab === 'deaths' && !isTrashPull && summary !== null}
            <!-- Beside the deaths: the boss abilities that did the hurting, with the Mechanics › link. -->
            <KeyMechanics {summary} {spellIcon} tabHref={hrefWith({ tab: 'mechanics' })} />
          {/if}
          {#if state.tab === 'mechanics' && base !== null}
            <MechanicsRail summary={base} {classOf} />
          {/if}
        {/if}
      {/snippet}

      <div bind:this={tabAnchor} class="-mb-[var(--section-gap)] scroll-mt-2" aria-hidden="true"></div>
      {#if onePullTab}
        <Plate label={reportCopy.onePullLabel}>
          <p
            class="text-muted px-4 py-4 text-[14px]"
            data-testid={state.tab === 'rating' ? 'rating-tab-night' : 'night-tables-only'}
          >
            {state.tab === 'rating' ? reportCopy.ratingNight : reportCopy.onePullOnly}
          </p>
        </Plate>
      {:else if scoped !== null}
        {#if state.tab === 'summary' && nightMode}
          {#if night !== null}
            <NightView
              {night}
              inScope={(guid) => inSource(guid, scopeKey, scopePlayers, friendlySet)}
              trash={trashMode}
              loading={nightLoading}
              {unitClasses}
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
            showParse={parseColumn}
            approximate={!windowIsWhole}
            {unitClasses}
            {playerHref}
            onSelectPlayer={(guid) => patch({ source: guid })}
          />
          {@render chart('damage-done')}
          <!-- Two columns from 1440: the long list on the left, the two short ones stacked on the right. -->
          <div class="report-abilities">
            <SummaryPanels
              metric="done"
              tabHref={hrefWith({ tab: 'damage-done' })}
              summary={scoped}
              durationMs={scoped.duration_ms}
              players={playerSet}
              approximate={!windowIsWhole}
              {spellIcon}
            />
            <div class="report-abilities-stack">
              <SummaryPanels
                metric="taken"
                tabHref={hrefWith({ tab: 'damage-taken' })}
                summary={scoped}
                durationMs={scoped.duration_ms}
                players={playerSet}
                approximate={!windowIsWhole}
                {spellIcon}
              />
              {#if scoped.healing.some((actor) => playerSet.has(actor.guid) && actor.effective > 0)}
                <SummaryPanels
                  metric="healing"
                  tabHref={hrefWith({ tab: 'healing' })}
                  summary={scoped}
                  durationMs={scoped.duration_ms}
                  players={playerSet}
                  approximate={!windowIsWhole}
                  {spellIcon}
                />
              {/if}
            </div>
          </div>
        {:else if ACTOR_TABS.has(state.tab) && playerView}
          <PlayerView
            tab={state.tab as 'damage-done' | 'healing' | 'damage-taken'}
            actor={tabActors[0]}
            row={playerRow}
            combatant={undefined}
            durationMs={scoped.duration_ms}
            approximate={actorTableApproximate}
            measure={nightMode ? undefined : measureRow}
            {spellIcon}
            dataBuild={activeBuild.build}
            onChart={abilityChartAvailable ? toggleAbilityOnChart : undefined}
            charted={chartedAbility?.key ?? ''}
          />
          {#if chartedError !== ''}
            <p class="text-small text-wipe" role="alert" data-testid="ability-chart-error">{chartedError}</p>
          {/if}
          {#if tabActors[0] !== undefined}
            <!-- The coach's line for this player (components/coach, placed as shipped): the row is open. -->
            <CoachRowLine actor={tabActors[0]} rowOpen />
          {/if}
          {@render chart(chartMetric)}
          <PlayerView
            part="gear"
            tab={state.tab as 'damage-done' | 'healing' | 'damage-taken'}
            actor={tabActors[0]}
            row={playerRow}
            combatant={scoped.combatants.find((row) => row.guid === sourcePlayer)}
            durationMs={scoped.duration_ms}
            {spellIcon}
            dataBuild={activeBuild.build}
            plannerHref={plannerHrefFor(sourcePlayer)}
          />
        {:else if ACTOR_TABS.has(state.tab)}
          <ActorTable
            tab={state.tab as 'damage-done' | 'healing' | 'damage-taken'}
            actors={tabActors}
            durationMs={scoped.duration_ms}
            roster={rosterByGuid}
            players={playerSet}
            {percentiles}
            parseFallback={tableParseFallback}
            showParse={parseColumn}
            amountApproximate={(filtersScale && !windowIsWhole && tableExact === null) || nightProrates}
            approximate={actorTableApproximate && tableExact === null}
            absent={absentPlayers}
            {windowIsWhole}
            deadAt={deadRows}
            selected={state.source}
            {playerHref}
            onSelectPlayer={(guid) => patch({ source: guid })}
          />
          {#if actorTableApproximate && !nightMode}
            <!-- A window or a target filter: the summary can only prorate, so the table is measured from
                 the fight's events, and this line says which figures are which. -->
            <p class="text-small text-muted" data-testid="approximate-note">
              {#if tableExact !== null}<span class="text-kill" data-testid="table-measured"
                  >Measured from the fight’s events for this window and filter.</span
                >{:else if tableMeasureError !== ''}<span class="text-wipe" role="alert"
                  >{tableMeasureError}</span
                >{:else}Measuring this window from the fight’s events… a tilde marks a figure still prorated.{/if}
            </p>
          {/if}
          {@render chart(chartMetric)}
          {#if !filtersScale && !nightMode}
            <!-- What the amount was made of: the scope's abilities, as the Summary lists them. -->
            <SummaryPanels
              metric={ABILITY_METRIC[chartMetric]}
              summary={scoped}
              durationMs={scoped.duration_ms}
              players={playerSet}
              approximate={!windowIsWhole}
              {spellIcon}
            />
          {/if}
        {:else if state.tab === 'buffs' || state.tab === 'debuffs'}
          {#if auraTabLazy.current}
            <auraTabLazy.current
              table={{
                tracks: scoped.auras,
                durationMs: scoped.duration_ms,
                startMs: timeWindow.startMs,
                kind: state.tab === 'buffs' ? 'BUFF' : 'DEBUFF',
                names: unitNames,
                bossNames: bossUnitNames,
                friendly: friendlySet,
                tauntSpells: new Set((scoped.taunts ?? []).map((taunt) => taunt.spell_id)),
                windowed: !windowIsWhole,
              }}
              cooldowns={state.tab === 'buffs'
                ? {
                    tracks: scoped.auras,
                    casts: scoped.casts,
                    pulls: scoped.pulls ?? [],
                    window: cutWindow,
                    deaths: scoped.deaths,
                    names: unitNames,
                    syncKey: nightMode ? '' : chartSyncKey,
                  }
                : undefined}
            />
          {:else}
            {@render lazyFallback(auraTabLazy, REPORT_LAZY_MIN_H.table)}
          {/if}
        {:else if state.tab === 'casts'}
          {#if castTableLazy.current}
            <castTableLazy.current
              rows={scoped.casts}
              everyone={base?.casts ?? scoped.casts}
              durationMs={scoped.duration_ms}
              startMs={timeWindow.startMs}
              {classOf}
              approximate={!windowIsWhole}
              measured={castExact ?? undefined}
              measureError={castMeasureError}
              whole={base === null ? undefined : scopeSource(base, scopeKey, scopePlayers, friendlySet).casts}
              names={unitNames}
            />
          {:else}
            {@render lazyFallback(castTableLazy, REPORT_LAZY_MIN_H.table)}
          {/if}
        {:else if state.tab === 'interrupts' || state.tab === 'dispels'}
          {#if exchangeTableLazy.current}
            {#if state.tab === 'interrupts'}
              <exchangeTableLazy.current
                kind="interrupt"
                rows={scoped.interrupts}
                everyone={windowed?.interrupts ?? scoped.interrupts}
                emptyText={exchangeEmpty('interrupted')}
                windowed={!windowIsWhole}
                casts={base?.casts ?? []}
                players={playerSet}
                {classOf}
              />
            {:else}
              <exchangeTableLazy.current
                kind="dispel"
                rows={scoped.dispels}
                everyone={windowed?.dispels ?? scoped.dispels}
                emptyText={exchangeEmpty('dispelled')}
                windowed={!windowIsWhole}
                auras={base?.auras ?? []}
                players={playerSet}
                {classOf}
              />
            {/if}
          {:else}
            {@render lazyFallback(exchangeTableLazy, REPORT_LAZY_MIN_H.table)}
          {/if}
        {:else if state.tab === 'resources'}
          {#if resourceGraphsLazy.current}
            <resourceGraphsLazy.current
              tracks={scoped.resources}
              durationMs={scoped.duration_ms}
              windowed={!windowIsWhole}
              deaths={scoped.deaths.map((death) => ({
                guid: death.guid,
                at_ms: death.at_ms - timeWindow.startMs,
              }))}
              syncKey={nightMode ? '' : chartSyncKey}
              windowStartMs={timeWindow.startMs}
              {classOf}
            />
          {:else}
            {@render lazyFallback(resourceGraphsLazy, REPORT_LAZY_MIN_H.table)}
          {/if}
        {:else if state.tab === 'threat'}
          {#if threatTableLazy.current}
            <threatTableLazy.current
              rail={false}
              rows={scoped.threat}
              players={playerSet}
              pairs={scoped.threat_by_target ?? []}
              everyonePairs={windowed?.threat_by_target ?? scoped.threat_by_target ?? []}
              taunts={scoped.taunts}
              names={unitNames}
              {classOf}
              approximate={!windowIsWhole}
              target={state.target}
              sourceName={playerSet.has(state.source)
                ? unitNames.get(state.source)
                : state.source === SOURCE_ENEMIES
                  ? 'the enemies'
                  : undefined}
              scopeNoun={nightMode ? 'night' : 'pull'}
              durationMs={summary?.duration_ms ?? scoped.duration_ms}
              window={cutWindow}
              {nightMode}
              onPatch={patch}
              onWindow={setWindow}
              totalThreat={windowed?.threat
                .filter((row) => playerSet.has(row.guid))
                .reduce((sum, row) => sum + row.threat, 0)}
              chart={base === null || nightMode
                ? undefined
                : {
                    pairs: base.threat_by_target ?? [],
                    players: threatPlayers,
                    taunts: base.taunts ?? [],
                    targetChanges: base.target_changes,
                    damageTaken: base.damage_taken,
                    events: (enemyGuids) =>
                      threatMarkers({
                        deaths: base?.deaths ?? [],
                        taunts: base?.taunts ?? [],
                        targetChanges: base?.target_changes,
                        enemyGuids,
                        deathHref,
                      }),
                    phases: chartPhases,
                    sourcePlayer,
                    youGuid,
                    chosen: state.players,
                    onChoose: (players) => patch({ players }),
                    brushShape,
                    syncKey: chartSyncKey,
                    onOpenDeath: openDeaths,
                  }}
            />
          {:else}
            {@render lazyFallback(threatTableLazy, REPORT_LAZY_MIN_H.table)}
          {/if}
        {:else if state.tab === 'deaths'}
          <DeathsTab
            deaths={scoped.deaths}
            casts={base?.casts ?? []}
            open={state.openDeaths}
            onPatch={patch}
            durationMs={summary?.duration_ms ?? scoped.duration_ms}
            pulls={scoped.pulls ?? []}
            combatants={scoped.combatants}
            {classOf}
            dataBuild={activeBuild.build}
            {treeSizesFor}
            onWindow={setWindow}
            {spellIcon}
          />
          {@render chart('damage-taken')}
        {:else if state.tab === 'coach'}
          <CoachTabHost
            {reportId}
            {fight}
            {firstFight}
            {roster}
            source={scopeKey}
            {nightMode}
            onSelectPlayer={(guid) => patch({ source: guid })}
          />
        {:else if state.tab === 'rating'}
          {#if ratingTabLazy.current}
            <ratingTabLazy.current
              {reportId}
              fightIndex={state.fight}
              {roster}
              source={scopeKey}
              onSelectPlayer={(guid) => patch({ source: guid })}
              onRatings={(next) => (ratingState = next)}
            />
          {:else}
            {@render lazyFallback(ratingTabLazy, REPORT_LAZY_MIN_H.rating)}
          {/if}
        {:else if state.tab === 'timeline'}
          {#if timelinesViewLazy.current}
            <timelinesViewLazy.current
              summary={scoped}
              window={timeWindow}
              {classOf}
              bossName={fight?.kind === 'encounter' ? fight.name : ''}
              players={playerSet}
              allCasts={summary?.casts ?? []}
              taunts={scoped.taunts ?? []}
              syncKey={chartSyncKey}
              auraOrder={[
                ...new Set(
                  (summary?.auras ?? [])
                    .filter((track) => inSource(track.target_guid, scopeKey, scopePlayers, friendlySet))
                    .map((track) => track.name),
                ),
              ].sort((a, b) => a.localeCompare(b))}
            />
          {:else}
            {@render lazyFallback(timelinesViewLazy, REPORT_LAZY_MIN_H.timelines)}
          {/if}
        {:else if state.tab === 'events'}
          {#if eventsViewLazy.current}
            <eventsViewLazy.current
              summary={windowed ?? scoped}
              {classOf}
              names={unitNames}
              inScope={(guid) => inSource(guid, scopeKey, scopePlayers, friendlySet)}
              off={state.eventsOff}
              search={state.find}
              onPatch={patch}
              loadStream={async () => {
                const { loadEventStream, sharedQueryLayer } = await import('../../lib/report/exact');
                return loadEventStream(
                  sharedQueryLayer(),
                  eventsUrl(dataBase, state.fight, engineVersion),
                  cutWindow,
                );
              }}
            />
          {:else}
            {@render lazyFallback(eventsViewLazy, REPORT_LAZY_MIN_H.events)}
          {/if}
        {:else if state.tab === 'queries'}
          {#if queriesViewLazy.current}
            <!-- The Starting points wrap inside the column (the mock's `.qstarts`), never under the rail. -->
            <div class="report-queries flex min-w-0 flex-col gap-[var(--section-gap)]">
              <queriesViewLazy.current
                dataBaseUrl={dataBase}
                fightIndex={state.fight}
                window={timeWindow}
                classByName={new Map((summary?.roster ?? []).map((row) => [row.name, row.class ?? '']))}
              />
            </div>
          {:else}
            {@render lazyFallback(queriesViewLazy, REPORT_LAZY_MIN_H.queries)}
          {/if}
        {:else if state.tab === 'compare' && summary !== null}
          {#if compareModeLazy.current}
            <compareModeLazy.current
              {fights}
              current={state.fight}
              dataBaseUrl={dataBase}
              {engineVersion}
              left={summary}
              window={windowIsWhole ? null : cutWindow}
              rightIndex={state.compareWith}
              metric={state.compareMetric}
              vs={state.compareVs}
              source={scopeKey}
              onPatch={patch}
            />
          {:else}
            {@render lazyFallback(compareModeLazy, REPORT_LAZY_MIN_H.compare)}
          {/if}
        {/if}
      {/if}
      {#if state.tab === 'mechanics' && base !== null}
        <!-- The whole fight, never the window and never the source scope: `base`, not `scoped`. -->
        {#if mechanicsModeLazy.current}
          <mechanicsModeLazy.current
            summary={base}
            {classOf}
            {nightMode}
            trash={!nightMode && fight?.kind !== 'encounter'}
            onPatch={patch}
            hrefFor={(next) => reportSearch(withState(state, legacyTab(next)), firstFight) || '?'}
          />
        {:else}
          {@render lazyFallback(mechanicsModeLazy, REPORT_LAZY_MIN_H.mechanics)}
        {/if}
      {:else if state.tab === 'mechanics' && nightMode && nightLoading}
        <p class="text-muted text-[14px]" data-testid="mechanics-night-loading">Folding the night’s pulls…</p>
      {/if}
      <!-- Last in the column: a glossary is a footnote, read when a word stops someone; only under
           the tabs whose tables use its words (logs lane review 2026-11-09). -->
      {#if GLOSSARY_TABS.has(state.tab)}<Glossary />{/if}
    </Page>
  </div>
{/if}

<style>
  /* Below 768 the tab strip is one line that scrolls sideways: its right edge fades out, so the
     cut tab reads as more to scroll, not as a broken label (logs lane review 2026-11-09). */
  @media (max-width: 767px) {
    .report-tabs :global(.ui-tsm-strip) {
      mask-image: linear-gradient(to right, var(--color-bg) calc(100% - 120px), transparent);
    }
  }
  /* Damage done by ability beside taken and healing (stacked) where the main column has the room (1440 up). */
  .report-abilities,
  .report-abilities-stack {
    display: grid;
    gap: var(--section-gap);
    grid-template-columns: minmax(0, 1fr);
    align-content: start;
  }
  /* The halves keep the ability's name readable at every width: a floor on the name, a cap on the bar. */
  .report-abilities :global(.ui-st) {
    --st-name: minmax(120px, 2fr);
    --st-bar: minmax(0, 120px);
  }
  @media (min-width: 1440px) {
    .report-abilities {
      grid-template-columns: repeat(2, minmax(0, 1fr));
      align-items: start;
    }
  }
  /* Side by side under 1760 the halves are narrow: the share figure stays, the share bar goes,
     so the ability's name keeps its room. */
  @media (min-width: 1440px) and (max-width: 1759px) {
    .report-abilities :global(.ui-st) {
      --st-bar: 0px;
    }
    .report-abilities :global(.ui-st-bar) {
      visibility: hidden;
      overflow: hidden;
      min-width: 0;
    }
  }
  .report-queries :global(.ui-seg) {
    flex-wrap: wrap;
  }
  .report-queries :global(.ui-seg-item) {
    min-width: 0;
  }
</style>
