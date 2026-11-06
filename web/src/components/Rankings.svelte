<!-- web/src/components/Rankings.svelte -->
<!-- The Rankings page (design/specs/2026-11-07-logs-reports.md section 7; the rules are
     2026-10-11-rankings.md's): "Where does my character stand on this boss?" A ledger Page: the
     PageHeader (the encounter as the title; signed in, your rank as its answer), the ControlBar
     (Encounter, Class, Spec, Ruleset, Metric, Board, and Region, Phase, Faction and Today only
     behind Filters, K2), the Tier list line, then the cohort's Section holding the one Plate: a
     RankList of each character's best kill, your row in place or pinned on top (K5), and the
     pager; the rank colour legend closes the page. With no encounter named and none to aim at,
     the page is the encounter list (K4). Every filter is in the URL (lib/rankings/url.ts).

     Who is looking sets where the board starts (lib/rankings/defaults.ts): the viewer's own
     class, spec, ruleset and role's metric, the current phase, and the encounter they most
     recently killed. The header selector re-aims the board in place. The encounter is a filter
     like the rest, in the path (/rankings/<slug>) so a board is still a link. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import type { MeCharacter } from '../lib/account/api';
  import { formatDate } from '../lib/dates';
  import { createCharacterCardState } from '../lib/bis/character-card-state.svelte';
  import { guildHref } from '../lib/characters';
  import { classCrestSrc } from '../lib/class-crest';
  import { formatAmount } from '../lib/report/format';
  import { titleize } from '../lib/report/og-meta';
  import {
    fetchCharacter,
    fetchEncounters,
    fetchGuildRankings,
    fetchRankings,
    type EncounterOption,
    type GuildRankingRow,
    type RankingMetric,
    type RankingRow as RankingRowData,
    type RankingsPage,
  } from '../lib/rankings/api';
  import { encounterPickerCopy, rankingsEmptyCopy, rankingsCopy } from '../lib/rankings/copy';
  import { applyDefaults, reaim, rankingsDefaults, resolveDefaultEncounter } from '../lib/rankings/defaults';
  import { RANKINGS_LOADING_MIN_H } from '../lib/rankings/layout';
  import {
    classOptions,
    encounterOptions,
    metricOptions,
    rulesetOptions,
    specOptions,
  } from '../lib/rankings/options';
  import { rankRowView } from '../lib/rankings/rank-row-view';
  import { RANK_LEGEND, ownRowPlacement, showRowCrests } from '../lib/rankings/row-view';
  import { REGION } from '../lib/ui/regions';
  import { LINK, type ControlGroup } from '../lib/ui/primitives';
  import type { RankHead } from '../lib/ui/rank-list';
  import { PENDING_ANSWER, ownRankSpec } from '../lib/rankings/own-rank';
  import { cohortName, scopeHeading } from '../lib/rankings/scope';
  import {
    GUILD_KINDS,
    RANKING_METRICS,
    parseRankingsState,
    rankingsSearch,
    requiresEncounter,
    type RankingsState,
  } from '../lib/rankings/url';
  import RankingsAnswer from './rankings/RankingsAnswer.svelte';
  import RankingsMoreFilters from './rankings/RankingsMoreFilters.svelte';
  import ControlBar from './ui/ControlBar.svelte';
  import Notice from './ui/Notice.svelte';
  import Page from './ui/Page.svelte';
  import PageHeader from './ui/PageHeader.svelte';
  import Plate from './ui/Plate.svelte';
  import RankList from './ui/RankList.svelte';
  import RankRow from './ui/RankRow.svelte';
  import Row from './ui/Row.svelte';
  import Section from './ui/Section.svelte';
  import Skeleton from './ui/Skeleton.svelte';
  import Tag from './ui/Tag.svelte';

  // The prerendered fixture page passes the slug; the Worker-served shell has none, so the
  // island reads it out of the path. One component, both routes.
  let { slug = '' }: { slug?: string } = $props();

  const SLUG_PATH = /^\/rankings\/([a-z0-9-]{1,64})\/?$/;
  const reads = { encounters: fetchEncounters, character: fetchCharacter };
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
  let encounters = $state<EncounterOption[]>([]);
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
    encounters.find((encounter) => encounter.slug === encounterSlug)?.name ?? titleize(encounterSlug),
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

  const guildsBoardHref = $derived(
    `/rankings${rankingsSearch({ ...state, board: 'guild', kind: 'progress', page: 1 }, ownDefaults)}`,
  );

  function openGuildsBoard(event: MouseEvent): void {
    event.preventDefault();
    patch({ board: 'guild', kind: 'progress' });
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
      encounters: [] as EncounterOption[],
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
            page = result;
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

  // The header's answer follows the board's own filters but not its page, so paging never
  // re-asks it: it is rebuilt only when the key of what it reads changes.
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

  // ---- the page's frame (spec 2026-11-07 section 7)
  const TWO_METRICS = new Set(['dps', 'hps']);
  const BOARDS = [
    { value: 'character', label: 'Characters' },
    { value: 'guild', label: 'Guilds' },
  ];
  const isIndex = $derived(encounterSlug === '' && (needsPicker || !aimed));
  const cohort = $derived(
    state.class === '' ? rankingsCopy.everyClassAndSpec : scopeHeading(state.class, state.spec),
  );
  const rulesetName = $derived(
    rulesetOptions().find((option) => option.value === state.ruleset)?.label ?? rankingsCopy.allRulesets,
  );
  const rulesetMeta = $derived(state.ruleset === '' ? rulesetName : `${rulesetName} ruleset`);
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
    return [
      select(
        rankingsCopy.encounterLabel,
        'encounter',
        encounterSlug,
        encounterOptions(encounters, encounterSlug),
      ),
      ...(state.board === 'guild'
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
      {
        label: rankingsCopy.boardLabel,
        param: 'board',
        kind: 'segmented',
        value: state.board,
        options: BOARDS,
      },
    ];
  });

  function onBar(param: string, value: string): void {
    if (param === 'encounter') return chooseEncounter(value);
    if (param === 'class') return patch({ class: value, spec: '' });
    patch({ [param]: value } as Partial<RankingsState>);
  }

  const heads = $derived<RankHead[]>([
    { label: '#' },
    { label: '' },
    { label: 'Character' },
    { label: 'Share of the top' },
    { label: metricLabel === 'Damage' ? 'DPS' : metricLabel, align: 'end' },
    { label: 'Item level', align: 'end' },
    { label: 'Executed', align: 'end' },
    { label: 'Length', align: 'end' },
    { label: 'Date', align: 'end' },
    { label: '' },
  ]);
  const multiClass = $derived(page !== null && showRowCrests(page.rows));
  const ownKey = $derived(ownBoard && viewer !== null ? viewer.key : null);
  const placement = $derived(
    page === null || ownKey === null ? 'none' : ownRowPlacement(page.rows, page.character_rank),
  );
  const boardView = $derived(
    page === null
      ? null
      : {
          total: page.total,
          top: page.rows[0]?.value ?? 0,
          unit: metricLabel === 'Damage' ? 'DPS' : state.metric === 'hps' ? 'HPS' : '',
          now,
        },
  );
  const pageHref = (n: number): string => urlFor({ ...state, page: n }, encounterSlug);

  function backToOwnBoard(): void {
    patch({ class: ownDefaults.class, spec: ownDefaults.spec });
  }
</script>

<div data-testid="rankings" id="rankings" class="flex flex-col">
  <Page layout="ledger" testid="rankings-page">
    {#snippet header()}
      <PageHeader
        title={isIndex ? rankingsCopy.h1 : encounterName}
        answer={!isIndex && viewer !== null && state.board !== 'guild' ? rankAnswer : undefined}
      >
        {#snippet meta()}
          {#if isIndex}<span>{rankingsCopy.jobSignedOut}</span>
          {:else}<span>{rankingsCopy.h1} · {state.board === 'guild' ? 'Guilds' : cohort} · {rulesetMeta}</span
            >{/if}
          {#if updatedStamp !== null}<span>Updated {updatedStamp}</span>{/if}
        {/snippet}
      </PageHeader>
    {/snippet}
    {#snippet bar()}
      {#if !isIndex}
        <ControlBar
          {groups}
          label="Ranking choices"
          labelsFrom={1920}
          summary={`${encounterName} · ${state.spec !== '' ? state.spec : state.class === '' ? rankingsCopy.everyClass : cohort}`}
          onchange={onBar}
          testid="rankings-filters"
        >
          {#snippet tail()}<RankingsMoreFilters
              filters={state}
              defaultPhase={ownDefaults.phase}
              onpatch={patch}
            />{/snippet}
        </ControlBar>
      {/if}
      <div class="full-bleed" data-region={REGION.summary}>
        <div class="page-column px-[var(--page-inset)]">
          <Notice text={rankingsCopy.tierQuestion} tone="neutral">
            {#snippet tail()}<a
                class="{LINK} inline-flex items-center max-md:-my-[10px] max-md:min-h-11"
                href="/tiers"
                data-testid="rankings-tier-link">{rankingsCopy.tierLink} ›</a
              >{/snippet}
          </Notice>
          {#if viewer !== null && ownDefaults.class !== '' && !ownBoard && !isIndex}
            <p class="text-small text-muted" data-testid="rankings-showing-note">
              {rankingsCopy.showingNote(cohort.toLowerCase(), ownCohort)}
              <button type="button" class={LINK} data-testid="rankings-back" onclick={backToOwnBoard}
                >{rankingsCopy.backTo(`${ownCohort}s`)}</button
              >
            </p>
          {/if}
        </div>
      </div>
    {/snippet}

    {#if isIndex}
      <Section
        title={encounterPickerCopy.heading}
        meta={encounters.length > 0 ? `${encounters.length} encounters ranked` : undefined}
        testid="encounter-picker"
      >
        <Plate label={encounterPickerCopy.heading}>
          {#if !aimed}
            <div class="p-4"><Skeleton lines={2} rowHeight="h-11" testid="rankings-skeleton" /></div>
          {:else if encounters.length === 0}
            <div class="text-muted p-4 text-[14px] leading-5" data-testid="rankings-no-encounters">
              <p>{encounterPickerCopy.noneYet}</p>
              <p>{encounterPickerCopy.tryGuildsProgress}</p>
              <div class="mt-3">
                <a
                  class={LINK}
                  href={guildsBoardHref}
                  data-testid="rankings-try-guilds"
                  onclick={openGuildsBoard}
                  >{state.board === 'guild'
                    ? encounterPickerCopy.openProgressBoard
                    : encounterPickerCopy.openGuildsBoard} ›</a
                >
              </div>
            </div>
          {:else}
            <div data-testid="rankings-encounters">
              {#each encounters as option (option.id)}
                <Row
                  title={option.name}
                  sub="Characters and guilds, every class and spec"
                  href={`/rankings/${option.slug}${rankingsSearch(state, ownDefaults)}`}
                >
                  {#snippet trail()}<span class="text-gold text-[14px] font-semibold">Open the board ›</span
                    >{/snippet}
                </Row>
              {/each}
            </div>
          {/if}
        </Plate>
      </Section>
      {@render legend(false)}
    {:else}
      <Section
        title={state.board === 'guild' ? 'Guilds' : cohort}
        meta={page !== null && state.board !== 'guild'
          ? `${page.total} ranked${page.deduped === true ? ', each character’s best kill' : ''}`
          : undefined}
        testid="rankings-heading"
      >
        <Plate label={rankingsCopy.h1}>
          {#if status === 'loading'}
            <div class="p-4">
              <Skeleton
                lines={8}
                rowHeight="h-11"
                minHeight={RANKINGS_LOADING_MIN_H}
                testid="rankings-skeleton"
              />
            </div>
          {:else if status === 'failed'}
            <div class="text-muted p-4 text-[14px] leading-5" role="alert" data-testid="rankings-error">
              <p>{error || 'Rankings did not load.'}</p>
              <div class="mt-3">
                <button type="button" class={LINK} onclick={() => (attempt += 1)}>Try again</button>
              </div>
            </div>
          {:else if state.board === 'guild'}
            {#if guildRows.length === 0}
              {@render empty()}
            {:else}
              <div data-testid="guild-rows">
                {#each guildRows as row (`${row.rank}-${row.guild.region}-${row.guild.ruleset}-${row.guild.name}`)}
                  <Row
                    title={row.guild.name}
                    rank={String(row.rank)}
                    href={guildHref(row.guild.region, row.guild.ruleset, row.guild.name)}
                  >
                    {#snippet trail()}<span class="text-strong font-bold tabular-nums"
                        >{formatAmount(Math.round(row.value))}</span
                      >{/snippet}
                  </Row>
                {/each}
              </div>
            {/if}
          {:else if page === null || page.rows.length === 0 || boardView === null}
            {@render empty()}
          {:else}
            <RankList
              heads={multiClass ? heads : heads.filter((_, index) => index !== 1)}
              figures={4}
              links={3}
              leads={multiClass}
              label={cohort}
              testid="ranking-board"
            >
              {#if placement === 'pinned' && page.character_rank}
                {@render rankRow(
                  { ...page.character_rank.row, rank: page.character_rank.rank },
                  true,
                  'ranking-pinned-row',
                )}
              {/if}
              {#each page.rows as row (`${row.report_id}-${row.fight_index}-${row.player.key}`)}
                {@render rankRow(
                  row,
                  placement === 'in-place' && row.player.key === ownKey,
                  `ranking-${row.rank}`,
                )}
              {/each}
            </RankList>
            {#if lastPage > 1}
              <div
                class="border-line-soft flex items-center justify-end gap-4 border-t px-4 py-3 text-[14px]"
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
            {/if}
          {/if}
        </Plate>
      </Section>
      {@render legend(true)}
    {/if}
  </Page>
</div>

{#snippet rankAnswer()}
  {#if viewer !== null}<RankingsAnswer spec={answer} character={viewer} />{/if}
{/snippet}

{#snippet empty()}
  <div class="text-muted p-4 text-[14px] leading-5" data-testid="rankings-empty">
    <p>{rankingsEmptyCopy.message}</p>
    <div class="mt-3"><a class={LINK} href={rankingsEmptyCopy.href}>{rankingsEmptyCopy.action} ›</a></div>
  </div>
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
      {#snippet lead()}
        {#if multiClass}<img
            src={classCrestSrc(view.classSlug)}
            alt=""
            width="32"
            height="32"
            class="rounded-full"
          />{/if}
      {/snippet}
      {#snippet tag()}
        {#if you}<Tag tone="gold">You</Tag>{/if}
      {/snippet}
    </RankRow>
  {/if}
{/snippet}

{#snippet legend(region: boolean)}
  <div class="pt-1" data-region={region ? REGION.footerNote : undefined}>
    <p
      class="text-small text-muted flex flex-wrap items-baseline gap-x-[14px] gap-y-1"
      data-testid="rankings-legend"
    >
      {rankingsCopy.legendLead}
      {#each RANK_LEGEND as step (step.label)}<b class="font-bold" style:color={step.color}>{step.label}</b
        >{/each}
    </p>
  </div>
{/snippet}
