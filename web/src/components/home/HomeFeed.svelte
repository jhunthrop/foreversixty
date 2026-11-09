<!-- web/src/components/home/HomeFeed.svelte -->
<!-- "Latest on Forever" (home B spec 2026-11-07 P1 to P11, F1): what is happening on Forever, from
     GET /v1/feed, with the filters a signed-in player gets (All, Your spec, Your guild, Logs; signed
     out All only), each line's time updating each minute, and below five lines the floor of real
     dated site events (lib/home/floor.ts, built at build time). Its own skeleton and its own error. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { invalidate } from '../../lib/data/query';
  import { minuteClock } from '../../lib/guild/minute-clock.svelte';
  import {
    feedCopy,
    feedFilters,
    feedPath,
    fetchFeed,
    homeRows,
    whenText,
    withFloor,
    type FeedFilter,
    type FeedRow,
  } from '../../lib/home/feed';
  import { API_BASE_URL } from '../../lib/planner/config';
  import { LINK } from '../../lib/ui/primitives';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import FeedRowUi from '../ui/FeedRow.svelte';
  import ItemIcon from '../ui/ItemIcon.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Section from '../ui/Section.svelte';
  import Segmented from '../ui/Segmented.svelte';
  import FeedSentence from './FeedSentence.svelte';
  import FeedSkeleton from '../ui/FeedSkeleton.svelte';

  let {
    signedIn,
    spec,
    band,
    inGuild = false,
    raidOnPlaque = false,
    floor,
    build,
  }: {
    signedIn: boolean;
    /** The viewer's spec key (`warrior-fury`) and band: Your spec, and the YOUR SPEC tags. */
    spec?: string;
    band?: number;
    inGuild?: boolean;
    /** The guild's raid is on the plaque above: its post is left out here. */
    raidOnPlaque?: boolean;
    floor: FeedRow[];
    build: string;
  } = $props();

  /** The row's crest or icon (round 5 mock: 30 px leads on the boxed feed). */
  const LEAD_PX = 32;
  /** A line this fresh slides in (when motion is allowed). */
  const FRESH_MS = 5 * 60_000;
  const clock = minuteClock(() => undefined);
  const filters = $derived(feedFilters(signedIn, inGuild));
  let filter = $state<FeedFilter>('all');
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let rows = $state<FeedRow[]>([]);

  function load(): void {
    status = 'loading';
    const asked = filter;
    fetchFeed({ filter: asked, spec, band, signedIn })
      .then((page) => {
        if (asked !== filter) return;
        rows = page.rows;
        status = 'ready';
      })
      .catch(() => {
        if (asked === filter) status = 'failed';
      });
  }

  function retry(): void {
    invalidate(`${API_BASE_URL}${feedPath({ filter, spec, band })}`);
    load();
  }

  $effect(() => {
    void filter;
    untrack(load);
  });

  const shown = $derived(withFloor(homeRows(rows, raidOnPlaque), floor, filter, spec));
  const freshCut = $derived(clock.now.getTime() - FRESH_MS);
  const options = $derived(filters.map((value) => ({ value, label: feedCopy.filters[value] })));
</script>

<Section title={feedCopy.title} boxed testid="home-feed">
  {#snippet head()}
    {#if options.length > 0}
      <Segmented
        {options}
        value={filter}
        label={feedCopy.filterLabel}
        size="sm"
        onChange={(value) => (filter = value)}
        testid="home-feed-filter"
      />
    {/if}
  {/snippet}
  {#if status === 'loading'}
    <FeedSkeleton testid="home-feed-loading" />
  {:else if status === 'failed'}
    <LoadError message={feedCopy.failed} onRetry={retry} testid="home-feed-error" />
  {:else if shown.length === 0}
    <p class="text-muted py-3 text-[15px] leading-6" data-testid="home-feed-empty">
      {#if filter === 'guild'}{feedCopy.guildEmpty}{:else}{feedCopy.logsEmpty}
        <a class={LINK} href="/logs/upload">{feedCopy.uploadLog} ›</a>{/if}
    </p>
  {:else}
    <div class="flex flex-col" data-testid="home-feed-rows">
      {#each shown as row, index (row.id)}
        <FeedRowUi
          href={row.href}
          when={whenText(row, clock.now)}
          fresh={index === 0 && new Date(row.at).getTime() > freshCut}
          testid="home-feed-row"
        >
          {#snippet lead()}
            {#if row.crest !== null}
              <ClassCrestRing characterClass={row.crest} size={LEAD_PX} />
            {:else}
              <ItemIcon icon={row.icon ?? undefined} quality={row.quality ?? 1} {build} size={LEAD_PX} />
            {/if}
          {/snippet}
          <FeedSentence segments={row.segments} />
        </FeedRowUi>
      {/each}
    </div>
  {/if}
</Section>
