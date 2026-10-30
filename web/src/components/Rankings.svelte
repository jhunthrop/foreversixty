<!-- web/src/components/Rankings.svelte -->
<!-- The Rankings page (design/specs/2026-10-11-rankings.md): "Where does my character stand among
     characters like it?" The band, the shared character header (signed in: your rank on this
     encounter, one action; signed out: its one muted line), the board's heading, one filter
     panel and the board. Character boards and guild boards share the filters and the page's
     shape; only the rows differ. Every filter is in the URL (lib/rankings/url.ts), so a filtered
     board is a link.

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
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { formatAmount, rowLink } from '../lib/report/format';
  import { titleize } from '../lib/report/og-meta';
  import {
    fetchCharacter,
    fetchEncounters,
    fetchGuildRankings,
    fetchRankings,
    type EncounterOption,
    type GuildRankingRow,
    type RankingMetric,
    type RankingsPage,
  } from '../lib/rankings/api';
  import { encounterPickerCopy, rankingsEmptyCopy, rankingsCopy } from '../lib/rankings/copy';
  import { applyDefaults, reaim, rankingsDefaults, resolveDefaultEncounter } from '../lib/rankings/defaults';
  import { RANKINGS_LOADING_MIN_H } from '../lib/rankings/layout';
  import { PENDING_ANSWER, ownParsesAction, ownRankSpec } from '../lib/rankings/own-rank';
  import { PHASES } from '../lib/rankings/phases';
  import { cohortName, filterSummary, scopeHeading } from '../lib/rankings/scope';
  import {
    RANKING_METRICS,
    parseRankingsState,
    rankingsSearch,
    requiresEncounter,
    type RankingsState,
  } from '../lib/rankings/url';
  import CharacterHeader from './character/CharacterHeader.svelte';
  import RankingsBoard from './rankings/RankingsBoard.svelte';
  import RankingsFilters from './rankings/RankingsFilters.svelte';
  import RankingsBoards from './rankings/RankingsBoards.svelte';
  import RankingsLegend from './rankings/RankingsLegend.svelte';
  import RankingsHero from './rankings/RankingsHero.svelte';
  import ActionLink from './ui/ActionLink.svelte';
  import Card from './ui/Card.svelte';
  import { ACTION_LINK, ACTION_LINK_CHEVRON } from '../lib/ui/link';
  import Glyph from './ui/Glyph.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  // The prerendered fixture page passes the slug; the Worker-served shell has none, so the
  // island reads it out of the path. One component, both routes.
  let { slug = '' }: { slug?: string } = $props();

  const SLUG_PATH = /^\/rankings\/([a-z0-9-]{1,64})\/?$/;
  const PHASE_NAMES: ReadonlyMap<string, string> = new Map(PHASES.map((phase) => [phase.id, phase.label]));
  const reads = { encounters: fetchEncounters, character: fetchCharacter };
  const now = new Date();

  const session = createCharacterCardState();
  const viewer = $derived(session.character);
  const settled = $derived(session.status === 'ready' || session.status === 'failed');
  const signedIn = $derived(settled ? viewer !== null : null);
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
  let otherBoard = $state(false);
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
  const showClassAndSpec = $derived(viewer === null || ownDefaults.class === '' || otherBoard || !ownBoard);
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
    otherBoard = false;
    aimed = true;
    show(
      { ...aimedState, ...chosen.filters },
      chosen.encounter ?? (keepEncounter ? encounterSlug : resolved.slug),
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

  const summary = $derived(
    filterSummary({ encounterName, state, defaultMetric: ownDefaults.metric }, PHASE_NAMES),
  );
  const updatedStamp = $derived(updatedAt === null ? null : formatDate(new Date(updatedAt)));

  /** The character board's rows run the card's full width; every other state sits in the padded body. */
  const boardShown = $derived(
    !needsPicker && status === 'ready' && state.board !== 'guild' && page !== null && page.rows.length > 0,
  );

  function backToOwnBoard(): void {
    otherBoard = false;
    patch({ class: ownDefaults.class, spec: ownDefaults.spec });
  }
</script>

<div data-testid="rankings" id="rankings" class="flex flex-col">
  <RankingsHero {signedIn} updated={updatedStamp} />

  <div class="flex flex-col gap-4 px-[var(--page-inset)] pt-1 pb-10 min-[600px]:gap-[22px]">
    <CharacterHeader action={({ character }) => ownParsesAction(character)} {answer} source={session} />

    {#if viewer !== null && ownDefaults.class !== '' && !ownBoard}
      <p class="text-muted text-[13px]" data-testid="rankings-showing-note">
        {rankingsCopy.showingNote(
          state.class === ''
            ? rankingsCopy.everyClassAndSpec.toLowerCase()
            : scopeHeading(state.class, state.spec),
          ownCohort,
        )}
        <button
          type="button"
          class="text-gold hover:text-gold-hover inline-flex min-h-11 items-center underline underline-offset-4"
          data-testid="rankings-back"
          onclick={backToOwnBoard}>{rankingsCopy.backTo(`${ownCohort}s`)}</button
        >
      </p>
    {/if}

    <Card
      glyph="trophy"
      title={scopeHeading(state.class, state.spec)}
      titleTestid="rankings-scope"
      body="none"
      testid="rankings-heading"
    >
      {#snippet aside()}
        <RankingsBoards board={state.board} onboard={(board) => patch({ board })} />
      {/snippet}
      <div class="border-line-soft flex flex-col gap-2 border-t px-5 py-4">
        {#if viewer !== null && ownDefaults.class !== '' && !showClassAndSpec}
          <button
            type="button"
            class={ACTION_LINK}
            data-testid="rankings-another"
            onclick={() => (otherBoard = true)}
          >
            {rankingsCopy.anotherBoard}<Glyph name="chevron" size={16} class={ACTION_LINK_CHEVRON} />
          </button>
        {/if}
        <RankingsFilters
          filters={state}
          {encounters}
          {encounterSlug}
          {summary}
          {showClassAndSpec}
          noEncounters={aimed && encounters.length === 0 && encounterSlug === ''}
          onpatch={patch}
          onencounter={chooseEncounter}
        />
      </div>
      <div class="border-line-soft border-t" class:px-5={!boardShown} class:py-4={!boardShown}>
        {#if needsPicker}
          <div data-testid="encounter-picker" class="flex flex-col gap-3">
            {#if encounters.length === 0}
              <p class="text-muted text-[14px]" data-testid="rankings-no-encounters">
                {encounterPickerCopy.noneYet}
              </p>
              <p class="text-muted text-[14px]">{encounterPickerCopy.tryGuildsProgress}</p>
              <ActionLink href={guildsBoardHref} data-testid="rankings-try-guilds" onclick={openGuildsBoard}>
                {state.board === 'guild'
                  ? encounterPickerCopy.openProgressBoard
                  : encounterPickerCopy.openGuildsBoard}
              </ActionLink>
            {:else}
              <h2 class="label text-muted">{encounterPickerCopy.heading}</h2>
              <ul class="flex flex-col" data-testid="encounter-picker-rows">
                {#each encounters as option (option.id)}
                  <li class="border-line-soft border-b py-2 text-[14px]">
                    <a class={rowLink} href={`/rankings/${option.slug}${rankingsSearch(state, ownDefaults)}`}
                      >{option.name}</a
                    >
                  </li>
                {/each}
              </ul>
            {/if}
          </div>
        {:else if status === 'loading'}
          <Skeleton
            lines={8}
            rowHeight="h-11"
            minHeight={RANKINGS_LOADING_MIN_H}
            testid="rankings-skeleton"
          />
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
                  <span class="numeric text-[12px]">{row.rank}</span>
                  <a class={rowLink} href={guildHref(row.guild.region, row.guild.ruleset, row.guild.name)}
                    >{row.guild.name}</a
                  >
                  <span class="numeric text-right">{formatAmount(Math.round(row.value))}</span>
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
          <RankingsBoard
            {page}
            {metricLabel}
            ownKey={ownBoard && viewer !== null ? viewer.key : null}
            {now}
          />

          <div class="flex items-center gap-3 px-5 py-4">
            <button
              type="button"
              class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3 disabled:opacity-50"
              disabled={state.page <= 1}
              data-testid="rankings-prev"
              onclick={() => patch({ page: state.page - 1 })}
            >
              Previous
            </button>
            <span class="text-muted numeric text-[13px]">Page {state.page} of {lastPage}</span>
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
      {#snippet foot()}
        <RankingsLegend />
      {/snippet}
    </Card>
  </div>
</div>
