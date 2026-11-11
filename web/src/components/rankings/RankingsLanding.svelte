<!-- web/src/components/rankings/RankingsLanding.svelte -->
<!-- The /rankings landing (K4, coordinator round 2026-11-09: a raider's ranking tool, dense but
     clean): the lead panel lists each ranked encounter, its name a gold link to the board, its zone,
     size and kill count, then its top three parses inline (rank in its percentile colour, crest,
     name in class colour, DPS, percentile). Beside it, Leaders by spec: every spec's best on the raid
     boss with the most kills, healers by HPS. Every board read stands alone: one that fails or comes
     back empty says so in one line, and the rest still show. The newest board's updated_at goes up
     to the title band's stamp. -->
<script lang="ts">
  import { fetchRankings, type RankedEncounter, type RankingsPage } from '../../lib/rankings/api';
  import { encounterPickerCopy, rankingsCopy } from '../../lib/rankings/copy';
  import {
    LANDING_ENCOUNTERS,
    leaderEncounter,
    specLeaders,
    topParses,
    type LandingRow,
    type SpecLeaders,
  } from '../../lib/rankings/landing';
  import type { RankHead } from '../../lib/ui/rank-list';
  import RankList from '../ui/RankList.svelte';
  import RankingsEncounters from './RankingsEncounters.svelte';
  import RankRow from '../ui/RankRow.svelte';
  import Row from '../ui/Row.svelte';
  import Ruled from '../ui/Ruled.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  let {
    encounters,
    phase,
    hrefFor,
    specHref,
    onupdated,
  }: {
    /** The ranked encounters, most kills first. */
    encounters: readonly RankedEncounter[];
    /** The board's phase filter ('' for every phase). */
    phase: string;
    hrefFor: (slug: string) => string;
    /** A spec's board on an encounter (a leader row's link). */
    specHref: (slug: string, classSlug: string, specName: string) => string;
    /** The newest board's updated_at, once one answers. */
    onupdated?: (iso: string) => void;
  } = $props();

  type Board = { status: 'loading' } | { status: 'failed' } | { status: 'ready'; page: RankingsPage };

  const shown = $derived(encounters.slice(0, LANDING_ENCOUNTERS));
  /** The encounters past the busiest few: a plain list under them, so every board is a link away. */
  const more = $derived(encounters.slice(LANDING_ENCOUNTERS));
  const leaderAt = $derived(leaderEncounter(encounters));
  let boards = $state<Record<string, Board>>({});
  let hps = $state<Board>({ status: 'loading' });

  function read(slug: string, metric: 'dps' | 'hps'): Promise<RankingsPage> {
    return fetchRankings({ encounter: slug, metric, phase, distinct: 'character' }).then((page) => {
      onupdated?.(page.updated_at);
      return page;
    });
  }

  $effect(() => {
    const slugs = shown.map((encounter) => encounter.slug);
    const leader = leaderAt?.slug ?? null;
    let stale = false;
    boards = Object.fromEntries(slugs.map((slug) => [slug, { status: 'loading' } as Board]));
    hps = { status: 'loading' };
    for (const slug of slugs) {
      read(slug, 'dps').then(
        (page) => {
          if (!stale) boards = { ...boards, [slug]: { status: 'ready', page } };
        },
        () => {
          if (!stale) boards = { ...boards, [slug]: { status: 'failed' } };
        },
      );
    }
    if (leader !== null) {
      read(leader, 'hps').then(
        (page) => {
          if (!stale) hps = { status: 'ready', page };
        },
        () => {
          if (!stale) hps = { status: 'failed' };
        },
      );
    }
    return () => {
      stale = true;
    };
  });

  const headsOf = (first: string, amount: string): RankHead[] => [
    { label: '#' },
    { label: '' },
    { label: first },
    { label: '' },
    { label: amount, align: 'end' },
    { label: rankingsCopy.percentileLabel, align: 'end' },
    { label: 'Date', align: 'end' },
    { label: '' },
  ];
  const PARSE_HEADS = headsOf('Character', 'DPS');

  const leaderBoard = $derived(leaderAt === null ? undefined : boards[leaderAt.slug]);
  const leaders = $derived.by((): SpecLeaders | null => {
    if (leaderAt === null || leaderBoard?.status !== 'ready' || hps.status === 'loading') return null;
    const slug = leaderAt.slug;
    return specLeaders(leaderBoard.page, hps.status === 'ready' ? hps.page : null, (cls, spec) =>
      specHref(slug, cls, spec),
    );
  });
</script>

{#snippet rows(list: readonly LandingRow[], testid: string, heads: readonly RankHead[] | undefined)}
  <!-- clear of the panel's framed corners on a phone -->
  <div class="pb-2 max-md:pr-[6px]">
    <RankList {heads} figures={2} figureWidth={72} links={0} compact phoneFigures={1} {testid}>
      {#each list as view (view.key)}
        <RankRow
          rank={view.rank}
          rankColor={view.rankColor}
          name={view.name}
          href={view.href}
          classSlug={view.classSlug}
          sub={view.sub}
          fraction={view.fraction}
          figures={view.figures}
          gap={view.gap}
        >
        </RankRow>
      {/each}
    </RankList>
  </div>
{/snippet}

{#snippet line(text: string, testid: string)}
  <p class="text-small text-muted px-5 pb-3 max-md:px-[14px]" data-testid={testid}>{text}</p>
{/snippet}

<Section title={encounterPickerCopy.heading} panel frame="lead" flush testid="encounter-picker">
  <Ruled testid="rankings-encounters">
    {#each shown as encounter, index (encounter.id)}
      {@const board = boards[encounter.slug]}
      <div>
        <div class="px-5 max-md:px-[14px]">
          <Row
            title={encounter.name}
            sub={encounterPickerCopy.where(encounter.zone, encounter.size) || undefined}
            href={hrefFor(encounter.slug)}
            testid="rankings-encounter"
          >
            {#snippet titleContent()}<span class="text-gold">{encounter.name}</span>{/snippet}
            {#snippet trail()}<span class="text-strong text-[14px] tabular-nums"
                >{encounterPickerCopy.ranked(encounter.kills)} <span class="text-gold">›</span></span
              >{/snippet}
          </Row>
        </div>
        {#if board === undefined || board.status === 'loading'}
          <div class="px-5 pb-3">
            <Skeleton lines={3} rowHeight="h-9" testid="rankings-landing-skeleton" />
          </div>
        {:else if board.status === 'failed'}
          {@render line(encounterPickerCopy.topFailed, `rankings-top-${encounter.slug}`)}
        {:else if board.page.rows.length === 0}
          {@render line(encounterPickerCopy.topNone, `rankings-top-${encounter.slug}`)}
        {:else}
          {@render rows(
            topParses(board.page),
            `rankings-top-${encounter.slug}`,
            index === 0 ? PARSE_HEADS : undefined,
          )}
        {/if}
      </div>
    {/each}
    {#if more.length > 0}
      <div>
        <p class="text-label text-muted px-5 pt-3 font-bold max-md:px-[14px]">{encounterPickerCopy.more}</p>
        <RankingsEncounters encounters={more} {hrefFor} testid="rankings-more" />
      </div>
    {/if}
  </Ruled>
</Section>

{#if leaderAt !== null}
  <Section
    title={encounterPickerCopy.leadersHeading}
    meta={leaderAt.name}
    panel
    frame="plain"
    flush
    testid="rankings-leaders"
  >
    {#if leaders === null && (leaderBoard?.status ?? 'loading') !== 'failed'}
      <div class="px-5 pb-3"><Skeleton lines={4} rowHeight="h-9" testid="rankings-landing-skeleton" /></div>
    {:else if leaders === null || leaders.damage.length + leaders.healing.length === 0}
      {@render line(
        leaders === null ? encounterPickerCopy.topFailed : encounterPickerCopy.topNone,
        'rankings-leaders-line',
      )}
    {:else}
      {#if leaders.damage.length > 0}
        {@render rows(
          leaders.damage,
          'rankings-leaders-damage',
          headsOf(encounterPickerCopy.damageHead, 'DPS'),
        )}
      {/if}
      {#if leaders.healing.length > 0}
        {@render rows(
          leaders.healing,
          'rankings-leaders-healing',
          headsOf(encounterPickerCopy.healingHead, 'HPS'),
        )}
      {/if}
    {/if}
  </Section>
{/if}
