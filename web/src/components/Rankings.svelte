<!-- web/src/components/Rankings.svelte -->
<!-- The Rankings page (design/specs/2026-11-07-logs-reports.md section 7; the rules are
     2026-10-11-rankings.md's; the look is the directions mock b1-rankings*): "Where does my
     character stand on this boss?" A split Page of panels: the compact title band (Rankings and
     the Updated stamp), the ControlBar (Board, then Encounter, Class, Spec, Ruleset, Metric, and
     Region, Phase, Faction and Today only behind Filters, K2), then the board's lead panel: the
     cohort's heading, the percentile legend and a RankList of each character's best kill, your row
     in place or pinned on top (K5), and the pager. The rail holds Your rank (signed in) and the
     tier list pointer; on a phone it comes first. With no encounter named and none to aim at, the
     page is the list of ranked encounters (K4). Every filter is in the URL (lib/rankings/url.ts).

     Who is looking sets where the board starts (lib/rankings/defaults.ts): the viewer's own
     class, spec, ruleset and role's metric, the current phase, and the encounter they most
     recently killed. The header selector re-aims the board in place. The encounter is a filter
     like the rest, in the path (/rankings/<slug>) so a board is still a link. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import type { MeCharacter } from '../lib/account/api';
  import { formatDate } from '../lib/dates';
  import { createCharacterCardState } from '../lib/bis/character-card-state.svelte';
  import { classCrestSrc } from '../lib/class-crest';
  import { titleize } from '../lib/report/og-meta';
  import {
    fetchCharacter,
    fetchGuildRankings,
    fetchRankedEncounters,
    fetchRankings,
    type GuildRankingRow,
    type RankedEncounter,
    type RankingMetric,
    type RankingRow as RankingRowData,
    type RankingsPage,
  } from '../lib/rankings/api';
  import { encounterPickerCopy, rankingsEmptyCopy, rankingsCopy } from '../lib/rankings/copy';
  import { applyDefaults, reaim, rankingsDefaults, resolveDefaultEncounter } from '../lib/rankings/defaults';
  import { guildFigureLabel, guildRowView } from '../lib/rankings/guild-row-view';
  import { RANKINGS_LOADING_MIN_H } from '../lib/rankings/layout';
  import {
    classOptions,
    encounterOptions,
    metricOptions,
    rulesetOptions,
    specOptions,
  } from '../lib/rankings/options';
  import { rankRowView } from '../lib/rankings/rank-row-view';
  import { RANK_LEGEND, ownRowPlacement, showRowCrests, withoutIdleRows } from '../lib/rankings/row-view';
  import { LINK, type ControlGroup } from '../lib/ui/primitives';
  import type { RankHead } from '../lib/ui/rank-list';
  import { PENDING_ANSWER, ownRankSpec } from '../lib/rankings/own-rank';
  import { cohortName, scopeHeading } from '../lib/rankings/scope';
  import { scoreCopy } from '../lib/score/copy';
  import { seedBoard, type RankingsSeed } from '../lib/rankings/snapshot';
  import {
    GUILD_KINDS,
    RANKING_METRICS,
    parseRankingsState,
    rankingsSearch,
    requiresEncounter,
    type GuildKind,
    type RankingsState,
  } from '../lib/rankings/url';
  import RankingsAnswer from './rankings/RankingsAnswer.svelte';
  import RankingsEmpty from './rankings/RankingsEmpty.svelte';
  import RankingsEncounters from './rankings/RankingsEncounters.svelte';
  import RankingsLanding from './rankings/RankingsLanding.svelte';
  import RankingsMoreFilters from './rankings/RankingsMoreFilters.svelte';
  import ControlBar from './ui/ControlBar.svelte';
  import Button from './ui/Button.svelte';
  import Divided from './ui/Divided.svelte';
  import Page from './ui/Page.svelte';
  import PageHeader from './ui/PageHeader.svelte';
  import RankList from './ui/RankList.svelte';
  import RankRow from './ui/RankRow.svelte';
  import Ruled from './ui/Ruled.svelte';
  import Section from './ui/Section.svelte';
  import Skeleton from './ui/Skeleton.svelte';
  import Tag from './ui/Tag.svelte';
  import { grouped } from '../lib/format/number';

  // A prerendered encounter page passes the slug; the Worker-served shell has none, so the
  // island reads it out of the path. One component, both routes. `seed` is the build's rankings
  // snapshot for this page (lib/rankings/snapshot.ts): its rows are drawn, server-side and until the
  // live board answers, for the board the snapshot read.
  let { slug = '', seed = null }: { slug?: string; seed?: RankingsSeed | null } = $props();

  const SLUG_PATH = /^\/rankings\/([a-z0-9-]{1,64})\/?$/;
  const reads = { encounters: fetchRankedEncounters, character: fetchCharacter };
  const now = new Date();

  const session = createCharacterCardState();
  const viewer = $derived(session.character);
  const settled = $derived(session.status === 'ready' || session.status === 'failed');
  const ownDefaults = $derived(rankingsDefaults(viewer, now));

  const initialSlug = untrack(() =>
    slug !== ''
      ? slug
      : typeof window === 'undefined'
        ? ''
        : (SLUG_PATH.exec(window.location.pathname)?.[1] ?? ''),
  );
  const urlNamedAnEncounter = initialSlug !== '';
  let encounterSlug = $state(initialSlug);
  let encounters = $state<RankedEncounter[]>(untrack(() => seed?.encounters ?? []));
  let state = $state<RankingsState>(
    typeof window === 'undefined' ? parseRankingsState('') : parseRankingsState(window.location.search),
  );
  let aimed = $state(false);
  let page = $state<RankingsPage | null>(null);
  let guildRows = $state<GuildRankingRow[]>([]);
  let updatedAt = $state<string | null>(null);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let error = $state('');
  let attempt = $state(0);
  /** Stamps the board's request; a later one makes an earlier answer stale. */
  let loadTicket = 0;
  /** Stamps the aiming; a later switch makes an earlier one stale. */
  let aimTicket = 0;
  /** What the visitor chose before the board knew who was looking: it lands on top of the defaults. */
  let early: { filters: Partial<RankingsState>; encounter?: string } = { filters: {} };

  const encounterName = $derived(
    encounterSlug === ''
      ? rankingsCopy.everyEncounter
      : (encounters.find((encounter) => encounter.slug === encounterSlug)?.name ?? titleize(encounterSlug)),
  );
  const needsPicker = $derived(aimed && encounterSlug === '' && requiresEncounter(state));
  const metricLabel = $derived(
    RANKING_METRICS.find((metric) => metric.id === state.metric)?.label ?? 'Value',
  );
  /** The viewer's own board: their class and spec, so their row is on it and can be marked. */
  const ownBoard = $derived(
    viewer !== null &&
      ownDefaults.class !== '' &&
      state.class === ownDefaults.class &&
      state.spec === ownDefaults.spec,
  );
  const ownCohort = $derived(cohortName(ownDefaults.class, ownDefaults.spec));
  const lastPage = $derived(page === null ? 1 : Math.max(1, Math.ceil(page.total / page.per_page)));
  /** The snapshot's rows for the board on screen while the live one has not answered. */
  const seedPage = $derived(page === null ? seedBoard(seed, encounterSlug, state) : null);
  /** The board the rows are drawn from: the live one, or the snapshot's until it answers. */
  const shownPage = $derived(page ?? seedPage);

  function urlFor(next: RankingsState, nextSlug: string): string {
    const path = nextSlug === '' ? '/rankings' : `/rankings/${nextSlug}`;
    return `${path}${rankingsSearch(next, ownDefaults)}`;
  }

  function show(next: RankingsState, nextSlug: string = encounterSlug): void {
    state = next;
    encounterSlug = nextSlug;
    window.history.replaceState(null, '', urlFor(next, nextSlug));
  }

  /** Any filter change returns to page one: page three of a different filter is nothing. */
  function patch(next: Partial<RankingsState>): void {
    const changed = { ...state, ...next, page: next.page ?? 1 };
    if (aimed) return show(changed);
    // The server-drawn filters are live before /v1/me answers: keep the choice, show it, and
    // put it on top of the defaults once the board knows who is looking.
    early = { ...early, filters: { ...early.filters, ...next } };
    state = changed;
  }

  function chooseEncounter(next: string): void {
    if (aimed) return show({ ...state, page: 1 }, next);
    early = { ...early, encounter: next };
    encounterSlug = next;
  }

  /**
   * Aims the board at whoever the session says is looking: on first load, only the filters the
   * link did not name; on a header switch, the new character's own, back on page one. The
   * encounter follows too unless the link named one.
   */
  async function aim(character: MeCharacter | null, first: boolean): Promise<void> {
    const ticket = ++aimTicket;
    status = 'loading';
    const defaults = rankingsDefaults(character, now);
    const resolved = await resolveDefaultEncounter(character, reads).catch(() => ({
      slug: '',
      encounters: [] as RankedEncounter[],
    }));
    if (ticket !== aimTicket) return;
    encounters = resolved.encounters;
    const keepEncounter = first && urlNamedAnEncounter;
    const aimedState = first
      ? applyDefaults(state, window.location.search, defaults)
      : reaim(state, defaults);
    const chosen = early;
    early = { filters: {} };
    aimed = true;
    show(
      { ...aimedState, ...chosen.filters },
      // Signed out with no encounter named, the page is the encounter list (K4); signed in it
      // opens the encounter the character last killed.
      chosen.encounter ?? (keepEncounter ? encounterSlug : character === null ? '' : resolved.slug),
    );
  }

  let aimedFor: string | null | undefined;
  $effect(() => {
    if (!settled) return;
    const character = viewer;
    const key = character?.key ?? null;
    if (untrack(() => aimedFor) === key) return;
    const first = untrack(() => aimedFor) === undefined;
    aimedFor = key;
    void untrack(() => aim(character, first));
  });

  /**
   * Every filter lives in the URL, so changing one always produces a brand new `state` object
   * (the immutable update in `show` above). A request is stamped with a ticket when it starts,
   * and before any write to shared state -- on the success path and on the rejection path alike
   * -- the stamp is compared with the latest: a fetch started for a filter combination nobody is
   * looking at any more is dropped on both paths instead of painting its answer under the
   * current filters' name. The same hazard report-tabs.spec.ts's `heldRoute()` covers for the
   * report island.
   */
  $effect(() => {
    void attempt;
    if (!aimed || needsPicker) {
      if (needsPicker) {
        page = null;
        guildRows = [];
        status = 'ready';
      }
      return;
    }
    const requested = state;
    const requestedSlug = encounterSlug;
    const own = untrack(() => (ownBoard && viewer !== null ? viewer.key : null));
    const ticket = ++loadTicket;
    status = 'loading';
    const load =
      requested.board === 'guild'
        ? fetchGuildRankings({ encounter: requestedSlug, kind: requested.kind, phase: requested.phase }).then(
            (result) => {
              if (ticket !== loadTicket) return;
              guildRows = result.rows;
              page = null;
              status = 'ready';
            },
          )
        : fetchRankings({
            encounter: requestedSlug,
            metric: requested.metric as RankingMetric,
            spec: requested.spec,
            class: requested.class,
            phase: requested.phase,
            region: requested.region,
            ruleset: requested.ruleset,
            faction: requested.faction,
            since: requested.since,
            page: requested.page,
            distinct: 'character',
            ...(own === null ? {} : { character: own }),
          }).then((result) => {
            if (ticket !== loadTicket) return;
            page = withoutIdleRows(result, requested.metric);
            updatedAt = result.updated_at;
            guildRows = [];
            status = 'ready';
          });

    void load.catch((thrown: unknown) => {
      if (ticket !== loadTicket) return;
      status = 'failed';
      error = thrown instanceof Error ? thrown.message : 'Rankings did not load.';
    });
  });

  // Your rank follows the board's own filters but not its page, so paging never re-asks it: it is
  // rebuilt only when the key of what it reads changes.
  const answerKey = $derived(
    JSON.stringify([
      encounterSlug,
      state.metric,
      state.phase,
      state.region,
      state.ruleset,
      state.faction,
      state.since,
      ownDefaults.class,
      ownDefaults.spec,
    ]),
  );
  const answer = $derived.by(() => {
    void answerKey;
    if (!aimed) return PENDING_ANSWER;
    return untrack(() =>
      ownRankSpec({
        encounterSlug,
        encounterName,
        state,
        own: { class: ownDefaults.class, spec: ownDefaults.spec },
      }),
    );
  });

  const updatedStamp = $derived(updatedAt === null ? null : formatDate(new Date(updatedAt)));

  // ---- the page's frame (spec 2026-11-07 section 7, directions b1-rankings)
  const TWO_METRICS = new Set(['dps', 'hps']);
  const BOARDS = [
    { value: 'character', label: 'Characters' },
    { value: 'guild', label: 'Guilds' },
  ];
  const isIndex = $derived(encounterSlug === '' && (needsPicker || !aimed));
  const isGuild = $derived(state.board === 'guild');
  const guildKind = $derived(state.kind as GuildKind);
  const cohort = $derived(
    state.class === '' ? rankingsCopy.everyClassAndSpec : scopeHeading(state.class, state.spec),
  );
  /** "Fury Warrior": the cohort in the singular, as an empty board names it. */
  const cohortOne = $derived(state.class === '' ? '' : `${cohortName(state.class, state.spec)} `);
  const boardGroup = $derived<ControlGroup>({
    label: rankingsCopy.boardLabel,
    param: 'board',
    kind: 'segmented',
    value: state.board,
    options: BOARDS,
  });
  const groups = $derived.by((): ControlGroup[] => {
    const select = (
      label: string,
      param: string,
      value: string,
      options: readonly { value: string; label: string }[],
    ): ControlGroup => ({
      label,
      param,
      kind: 'select',
      value,
      options: options.map((option) => ({ value: option.value, label: option.label })),
    });
    const metric: ControlGroup = TWO_METRICS.has(state.metric)
      ? {
          label: rankingsCopy.metricLabel,
          param: 'metric',
          kind: 'segmented',
          value: state.metric,
          options: metricOptions().filter((option) => TWO_METRICS.has(option.value)),
        }
      : select(rankingsCopy.metricLabel, 'metric', state.metric, metricOptions());
    // The Progress board spans every boss: its encounter choice starts with all of them.
    const spansAll = isGuild && !requiresEncounter(state);
    const encounterChoices = [
      ...(spansAll ? [{ value: '', label: rankingsCopy.everyEncounter }] : []),
      ...encounterOptions(encounters, encounterSlug),
    ];
    return [
      boardGroup,
      select(rankingsCopy.encounterLabel, 'encounter', encounterSlug, encounterChoices),
      ...(isGuild
        ? [
            select(
              'Kind',
              'kind',
              state.kind,
              GUILD_KINDS.map((kind) => ({ value: kind.id, label: kind.label })),
            ),
          ]
        : [
            select(rankingsCopy.classLabel, 'class', state.class, classOptions()),
            select(rankingsCopy.specLabel, 'spec', state.spec, specOptions(state.class)),
            select(rankingsCopy.rulesetLabel, 'ruleset', state.ruleset, rulesetOptions()),
            metric,
          ]),
    ];
  });

  function onBar(param: string, value: string): void {
    if (param === 'encounter') return chooseEncounter(value);
    if (param === 'class') return patch({ class: value, spec: '' });
    patch({ [param]: value } as Partial<RankingsState>);
  }

  const valueHead = $derived(metricLabel === 'Damage' ? 'DPS' : metricLabel);
  const heads = $derived<RankHead[]>([
    { label: '#' },
    { label: '' },
    { label: 'Character' },
    { label: '' },
    { label: valueHead, align: 'end' },
    { label: rankingsCopy.percentileLabel, align: 'end' },
    { label: 'Executed', align: 'end' },
    { label: 'Length', align: 'end' },
    { label: 'Date', align: 'end' },
    { label: '' },
  ]);
  const guildHeads = $derived<RankHead[]>([
    { label: '#' },
    { label: '' },
    { label: 'Guild' },
    { label: '' },
    { label: guildFigureLabel(guildKind), align: 'end' },
    { label: 'Date', align: 'end' },
    { label: '' },
  ]);
  const ownKey = $derived(ownBoard && viewer !== null ? viewer.key : null);
  const placement = $derived(
    page === null || ownKey === null ? 'none' : ownRowPlacement(page.rows, page.character_rank),
  );
  const boardView = $derived(
    shownPage === null
      ? null
      : {
          total: shownPage.total,
          top: shownPage.rows[0]?.value ?? 0,
          unit: metricLabel === 'Damage' ? 'DPS' : state.metric === 'hps' ? 'HPS' : '',
          multiClass: showRowCrests(shownPage.rows),
        },
  );
  const guildBest = $derived(guildRows[0]?.value ?? 0);
  const pageHref = (n: number): string => urlFor({ ...state, page: n }, encounterSlug);
  const encounterHref = (next: string): string => `/rankings/${next}${rankingsSearch(state, ownDefaults)}`;
  const widen = $derived(
    state.class === '' || isGuild
      ? undefined
      : {
          label: rankingsCopy.widen,
          href: urlFor({ ...state, class: '', spec: '', page: 1 }, encounterSlug),
        },
  );
  const otherEncounters = $derived(encounters.filter((encounter) => encounter.slug !== encounterSlug));
  const boardTitle = $derived(isGuild ? rankingsCopy.guildTitle[guildKind] : cohort);
  /** Nothing ranked where the visitor is looking: the rail offers the companion, what fills it. */
  const nothingRanked = $derived(
    aimed &&
      status === 'ready' &&
      (isIndex
        ? encounters.length === 0
        : isGuild
          ? guildRows.length === 0
          : page === null || page.rows.length === 0),
  );
  const showYourRank = $derived(viewer !== null && !isIndex && !isGuild && ownDefaults.class !== '');
  /** "212 ranked", unless Your rank already says "14th of 212 Fury Warriors" beside it. */
  const boardMeta = $derived(
    isGuild || shownPage === null || shownPage.total === 0 || (showYourRank && placement !== 'none')
      ? undefined
      : `${grouped(shownPage.total)} ranked`,
  );

  /** The landing's boards each carry updated_at: the band's stamp shows the newest. */
  function stampUpdated(iso: string): void {
    if (updatedAt === null || iso > updatedAt) updatedAt = iso;
  }

  function backToOwnBoard(): void {
    patch({ class: ownDefaults.class, spec: ownDefaults.spec });
  }
</script>

<div data-testid="rankings" id="rankings" class="flex flex-col">
  <Page layout="split" splitFrom={1280} panels railFirst testid="rankings-page">
    {#snippet header()}
      <PageHeader title={rankingsCopy.h1} lead={rankingsCopy.lead}>
        {#snippet meta()}
          <!-- Specs, not characters: the tier list, a link in the band (never a card stranded in the rail).
               It leads the line, so the stamp arriving after it moves nothing; no date is printed
               until the board names a real one. -->
          <a
            class="{LINK} [@media(pointer:coarse)]:-my-3 [@media(pointer:coarse)]:inline-block [@media(pointer:coarse)]:py-3"
            href={rankingsCopy.tierHref}
            data-testid="rankings-tier-link">{rankingsCopy.tierLink} ›</a
          >
          {#if updatedStamp !== null}<span data-testid="rankings-updated"
              >{rankingsCopy.updated(updatedStamp)}</span
            >{/if}
        {/snippet}
      </PageHeader>
    {/snippet}
    {#snippet bar()}
      <ControlBar
        groups={isIndex ? [boardGroup] : groups}
        inline={isIndex}
        label="Ranking choices"
        labelsFrom={1920}
        summary={isIndex
          ? undefined
          : `${encounterName} · ${isGuild ? rankingsCopy.guildTitle[guildKind] : state.spec !== '' ? state.spec : state.class === '' ? rankingsCopy.everyClass : cohort}`}
        onchange={onBar}
        testid="rankings-filters"
      >
        {#snippet tail()}{#if !isIndex && !isGuild}<RankingsMoreFilters
              filters={state}
              defaultPhase={ownDefaults.phase}
              onpatch={patch}
            />{/if}{/snippet}
      </ControlBar>
    {/snippet}
    {#snippet rail()}
      {#if showYourRank && viewer !== null}
        <!-- A hairline panel, the board the one framed panel; the answer and its next step only: the
             best kill's figures are the marked row's (h-audit round 8). -->
        <Section title={rankingsCopy.answerLabel} panel testid="rankings-your-rank">
          <RankingsAnswer spec={answer} character={viewer} />
          {#if !ownBoard}
            <div class="mt-3">
              <Divided>
                <p class="text-small text-muted" data-testid="rankings-showing-note">
                  {rankingsCopy.showingNote(cohort.toLowerCase(), ownCohort)}
                  <button type="button" class={LINK} data-testid="rankings-back" onclick={backToOwnBoard}
                    >{rankingsCopy.backTo(`${ownCohort}s`)}</button
                  >
                </p>
              </Divided>
            </div>
          {/if}
        </Section>
      {/if}
      <!-- From 1280, beside the board. Below it the rail stacks first, and this arriving after the
           board answered pushed the board down (CLS 0.09 at 360): there it follows the board. -->
      {#if nothingRanked}<div class="max-xl:hidden">{@render companion('rankings-companion')}</div>{/if}
    {/snippet}

    {#if isIndex && (aimed || seed !== null) && encounters.length > 0}
      <!-- Before the board knows who is looking, the snapshot's landing is for a signed-out visitor
           only: a hinted session sees the picker's skeleton until its own encounter opens. -->
      <div class="contents" data-signed-out-only={aimed ? undefined : true}>
        <RankingsLanding
          {encounters}
          {seed}
          live={aimed}
          phase={state.phase}
          hrefFor={encounterHref}
          specHref={(slug, cls, spec) =>
            urlFor({ ...state, board: 'character', class: cls, spec, page: 1 }, slug)}
          onupdated={stampUpdated}
        />
      </div>
      {#if !aimed}<div class="contents" data-signed-in-only>
          {@render pickerSkeleton('encounter-picker-pending')}
        </div>{/if}
    {:else if isIndex}
      {#if !aimed}{@render pickerSkeleton('encounter-picker')}{:else}
        <Section title={encounterPickerCopy.heading} panel frame="lead" flush testid="encounter-picker">
          <RankingsEmpty
            line={encounterPickerCopy.noneYet}
            fill={encounterPickerCopy.fill}
            testid="rankings-no-encounters"
          />
        </Section>
      {/if}
    {:else}
      <Section
        title={boardTitle}
        meta={boardMeta}
        action={widen !== undefined && page !== null && page.rows.length > 0 ? widen : undefined}
        panel
        frame="lead"
        flush
        testid="rankings-heading"
      >
        {#if status === 'loading' && seedPage !== null}
          <!-- The snapshot's rows until the live board answers; before the board knows who is looking
               they are for a signed-out visitor only (a hinted session opens its own board). -->
          <div class="contents" data-signed-out-only={aimed ? undefined : true}>
            {@render legend()}
            {@render board(seedPage.rows, heads)}
          </div>
          {#if !aimed}<div class="contents" data-signed-in-only>{@render boardSkeleton()}</div>{/if}
        {:else if status === 'loading'}
          {@render boardSkeleton()}
        {:else if status === 'failed'}
          <!-- The failed state holds the loading reserve, so nothing under the panel moves. -->
          <div
            class="text-muted {RANKINGS_LOADING_MIN_H} px-5 pb-4 text-[14px] leading-5"
            role="alert"
            data-testid="rankings-error"
          >
            <p>{error || 'Rankings did not load.'}</p>
            <div class="mt-3">
              <button type="button" class={LINK} onclick={() => (attempt += 1)}>Try again</button>
            </div>
          </div>
        {:else if isGuild}
          {#if guildRows.length === 0}
            <RankingsEmpty
              line={encounterSlug === ''
                ? rankingsEmptyCopy.guildsAnywhere
                : rankingsEmptyCopy.guilds(encounterName)}
              fill={rankingsEmptyCopy.guildsFill}
            />
          {:else}
            <RankList heads={guildHeads} compact figures={1} links={0} label={boardTitle} testid="guild-rows">
              {#each guildRows as row (`${row.rank}-${row.guild.region}-${row.guild.ruleset}-${row.guild.name}`)}
                {@const view = guildRowView(row, { kind: guildKind, best: guildBest })}
                <RankRow
                  rank={view.rank}
                  name={view.name}
                  href={view.href}
                  sub={view.sub}
                  fraction={view.fraction}
                  barTone="accent"
                  figures={view.figures}
                  gap={view.gap}
                  links={view.links}
                  testid={`guild-ranking-${row.rank}`}
                />
              {/each}
            </RankList>
          {/if}
        {:else if page === null || page.rows.length === 0 || boardView === null}
          <RankingsEmpty
            line={rankingsEmptyCopy.board(cohortOne.trim() || 'ranked', encounterName)}
            fill={rankingsEmptyCopy.boardFill}
            {widen}
          >
            {#if otherEncounters.length > 0}
              <p class="text-label text-muted px-5 pt-3 font-bold max-md:px-[14px]">
                {rankingsEmptyCopy.elsewhere}
              </p>
              <RankingsEncounters
                encounters={otherEncounters}
                hrefFor={encounterHref}
                testid="rankings-elsewhere"
              />
            {/if}
          </RankingsEmpty>
        {:else}
          {@render legend()}
          {#if placement === 'pinned' && page.character_rank}
            <!-- The viewer's best kill off this page: captioned, above the board, at its true rank;
                 the one column head serves both lists. -->
            <p class="text-label text-muted px-5 pb-1 font-bold max-md:px-[14px]">{rankingsCopy.ownBest}</p>
            <RankList
              {heads}
              figures={4}
              figureWidth={72}
              links={1}
              compact
              phoneFigures={2}
              label={rankingsCopy.ownBest}
              testid="ranking-pinned"
            >
              {@render rankRow(
                { ...page.character_rank.row, rank: page.character_rank.rank },
                true,
                'ranking-pinned-row',
                rankingsCopy.ownBest,
              )}
            </RankList>
          {/if}
          {@render board(page.rows, placement === 'pinned' ? undefined : heads)}
          {#if lastPage > 1}
            <Divided>
              <div
                class="flex w-full items-center justify-end gap-4 px-5 pt-3 pb-4 text-[14px]"
                data-testid="rankings-pager"
              >
                <span class="text-small text-muted tabular-nums">Page {state.page} of {lastPage}</span>
                {#if state.page > 1}
                  <a
                    class={LINK}
                    href={pageHref(state.page - 1)}
                    data-testid="rankings-prev"
                    onclick={(event) => {
                      event.preventDefault();
                      patch({ page: state.page - 1 });
                    }}>Previous</a
                  >
                {:else}<span class="text-muted" data-testid="rankings-prev" aria-disabled="true"
                    >Previous</span
                  >{/if}
                {#if state.page < lastPage}
                  <a
                    class={LINK}
                    href={pageHref(state.page + 1)}
                    data-testid="rankings-next"
                    onclick={(event) => {
                      event.preventDefault();
                      patch({ page: state.page + 1 });
                    }}>Next</a
                  >
                {:else}<span class="text-muted" data-testid="rankings-next" aria-disabled="true">Next</span
                  >{/if}
              </div>
            </Divided>
          {/if}
        {/if}
      </Section>
    {/if}
    {#if nothingRanked}<div class="xl:hidden">{@render companion('rankings-companion-below')}</div>{/if}
  </Page>
</div>

{#snippet companion(testid: string)}
  <Section title={scoreCopy.explainerLogsHeading} panel {testid}>
    <p class="text-small text-muted">{scoreCopy.explainerLogsBody}</p>
    <p class="mt-3">
      <Button href={scoreCopy.explainerLogsHref}>{scoreCopy.explainerLogsButton}</Button>
    </p>
  </Section>
{/snippet}

{#snippet rankRow(row: RankingRowData, you: boolean, testid: string, label?: string)}
  {#if boardView !== null}
    {@const view = rankRowView(row, boardView)}
    <RankRow
      rank={view.rank}
      rankColor={view.rankColor}
      name={view.name}
      href={view.href}
      classSlug={view.classSlug}
      sub={view.sub}
      note={view.note}
      fraction={view.fraction}
      figures={view.figures}
      gap={view.gap}
      links={view.links}
      selected={you}
      {label}
      {testid}
    >
      {#snippet lead()}<img
          src={classCrestSrc(view.classSlug)}
          alt=""
          width="28"
          height="28"
          class="rounded-full"
        />{/snippet}
      {#snippet tag()}
        {#if you}<Tag tone="yours">You</Tag>{/if}
      {/snippet}
    </RankRow>
  {/if}
{/snippet}

{#snippet board(rows: readonly RankingRowData[], columnHeads: RankHead[] | undefined)}
  <Ruled>
    <RankList
      heads={columnHeads}
      figures={4}
      figureWidth={72}
      links={1}
      compact
      phoneFigures={2}
      label={cohort}
      testid="ranking-board"
    >
      {#each rows as row (`${row.report_id}-${row.fight_index}-${row.player.key}`)}
        {@render rankRow(row, placement === 'in-place' && row.player.key === ownKey, `ranking-${row.rank}`)}
      {/each}
    </RankList>
  </Ruled>
{/snippet}

{#snippet boardSkeleton()}
  <div class="px-5 pb-4">
    <Skeleton lines={8} rowHeight="h-11" minHeight={RANKINGS_LOADING_MIN_H} testid="rankings-skeleton" />
  </div>
{/snippet}

{#snippet pickerSkeleton(testid: string)}
  <Section title={encounterPickerCopy.heading} panel frame="lead" flush {testid}>
    <div class="px-5 pb-4"><Skeleton lines={3} rowHeight="h-11" testid="rankings-skeleton" /></div>
  </Section>
{/snippet}

{#snippet legend()}
  <p
    class="text-small text-muted flex flex-wrap items-baseline gap-x-[14px] gap-y-1 px-5 pb-3 max-md:gap-x-[9px] max-md:px-[14px]"
    data-testid="rankings-legend"
  >
    <span>{rankingsCopy.legendLead}</span>
    {#each RANK_LEGEND as step (step.label)}<b class="font-bold" style:color={step.color}>{step.label}</b
      >{/each}
  </p>
{/snippet}
