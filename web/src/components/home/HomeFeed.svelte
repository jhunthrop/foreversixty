<!-- web/src/components/home/HomeFeed.svelte -->
<!-- "Latest on Forever" (home B spec 2026-11-07 P1 to P11, F1; round 5 draws it with no filters):
     what is happening on Forever, from GET /v1/feed, newest first, each line's time updating each
     minute, and below five lines the floor of real dated site events (lib/home/floor.ts, built at
     build time). Its own skeleton and its own error. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { invalidate } from '../../lib/data/query';
  import { minuteClock } from '../../lib/ui/minute-clock.svelte';
  import {
    feedCopy,
    feedPath,
    fetchFeed,
    homeRows,
    whenText,
    withFloor,
    type FeedFilter,
    type FeedRow,
  } from '../../lib/home/feed';
  import { HOME_FEED_FETCH, signedInRows } from '../../lib/home/feed-mix';
  import { API_BASE_URL } from '../../lib/planner/config';
  import FeedRowUi from '../ui/FeedRow.svelte';
  import ItemIcon from '../ui/ItemIcon.svelte';
  import EmptyState from '../ui/EmptyState.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Section from '../ui/Section.svelte';
  import FeedSentence from './FeedSentence.svelte';
  import FeedSkeleton from '../ui/FeedSkeleton.svelte';

  let {
    signedIn,
    spec,
    band,
    raidOnPlaque = false,
    guild,
    floor,
    build,
  }: {
    signedIn: boolean;
    /** The viewer's spec key (`warrior-fury`) and band: Your spec, and the YOUR SPEC tags. */
    spec?: string;
    band?: number;
    /** The guild's raid is on the plaque above: its post is left out here. */
    raidOnPlaque?: boolean;
    /** The viewer's own guild: its name is left off its members' sims. */
    guild?: string;
    floor: FeedRow[];
    build: string;
  } = $props();

  /** A line this fresh slides in (when motion is allowed). */
  const FRESH_MS = 5 * 60_000;
  const clock = minuteClock(() => undefined);
  /** Home shows everything; the filters left with round 5 (the mock draws none). */
  const filter: FeedFilter = 'all';
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let rows = $state<FeedRow[]>([]);
  /** Signed in, enough rows that seven survive the sims folding (lib/home/feed-mix.ts). */
  const limit = untrack(() => (signedIn ? HOME_FEED_FETCH : undefined));

  function load(): void {
    status = 'loading';
    fetchFeed({ filter, spec, band, signedIn, limit })
      .then((page) => {
        rows = page.rows;
        status = 'ready';
      })
      .catch(() => {
        status = 'failed';
      });
  }

  function retry(): void {
    invalidate(`${API_BASE_URL}${feedPath({ filter, spec, band, limit })}`);
    load();
  }

  $effect(() => {
    untrack(load);
  });

  const shown = $derived(
    signedIn
      ? signedInRows({ rows, floor, spec, ownGuild: guild, raidOnPlaque })
      : withFloor(homeRows(rows, raidOnPlaque), floor, filter, spec),
  );
  const freshCut = $derived(clock.now.getTime() - FRESH_MS);
</script>

<Section title={feedCopy.title} boxed testid="home-feed">
  {#if status === 'loading'}
    <FeedSkeleton testid="home-feed-loading" />
  {:else if status === 'failed'}
    <LoadError message={feedCopy.failed} onRetry={retry} testid="home-feed-error" />
  {:else if shown.length === 0}
    <EmptyState
      message={feedCopy.logsEmpty}
      action={{ label: feedCopy.uploadLog, href: '/logs/upload' }}
      testid="home-feed-empty"
    />
  {:else}
    <div class="flex flex-col" data-testid="home-feed-rows">
      {#each shown as row, index (row.id)}
        <FeedRowUi
          href={row.href}
          when={whenText(row, clock.now)}
          fresh={index === 0 && new Date(row.at).getTime() > freshCut}
          lead={row}
          {build}
          testid="home-feed-row"
        >
          <FeedSentence segments={row.segments} />
        </FeedRowUi>
      {/each}
    </div>
  {/if}
</Section>
