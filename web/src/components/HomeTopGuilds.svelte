<!-- web/src/components/HomeTopGuilds.svelte -->
<!-- The home page's Rankings panel live element (spec 2026-09-23 §2 item 3): the top five
     guilds by progression, from `GET /v1/rankings/guilds?kind=progress` -- the one guild
     leaderboard kind that needs no `encounter` (api/internal/rankings/guilds.go: `KindProgress`
     counts distinct encounters killed, ordered desc). Same loading/error/empty/reveal shape
     as `RecentReports.svelte` and `Rankings.svelte`'s own guild board, at panel scale: five
     rows, not eight. -->
<script lang="ts">
  import { guildHref } from '../lib/characters';
  import { rowLink } from '../lib/report/format';
  import { fetchGuildRankings, type GuildRankingRow } from '../lib/rankings/api';
  import { homeTopGuildsCopy } from '../lib/rankings/copy';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  const ROWS = 5;

  let rows = $state<GuildRankingRow[]>([]);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let attempt = $state(0);

  async function load(): Promise<void> {
    status = 'loading';
    try {
      const result = await fetchGuildRankings({ encounter: '', kind: 'progress' });
      rows = result.rows.slice(0, ROWS);
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  $effect(() => {
    void attempt;
    void load();
  });
</script>

{#if status === 'loading'}
  <Skeleton
    lines={ROWS}
    rowHeight="h-9"
    minHeight="min-h-[220px]"
    label={homeTopGuildsCopy.placeholderLabel}
    testid="home-top-guilds-skeleton"
  />
{:else if status === 'failed'}
  <LoadError
    message={homeTopGuildsCopy.failed}
    onRetry={() => (attempt += 1)}
    testid="home-top-guilds-error"
  />
{:else if rows.length === 0}
  <div class="flex flex-col items-start gap-3" data-testid="home-top-guilds-empty">
    <p class="text-muted text-[14px]">{homeTopGuildsCopy.empty}</p>
    <a class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-4" href={homeTopGuildsCopy.emptyHref}>
      {homeTopGuildsCopy.emptyAction}
    </a>
  </div>
{:else}
  <ul class="reveal flex flex-col" data-testid="home-top-guilds">
    {#each rows as row (`${row.rank}-${row.guild.region}-${row.guild.ruleset}-${row.guild.name}`)}
      <li
        class="border-line-soft grid min-h-9 grid-cols-[24px_minmax(0,1fr)_auto] items-center gap-2 border-b py-1 text-[13px] last:border-b-0"
      >
        <span class="text-muted numeric text-[12px]">{row.rank}</span>
        <a
          class="{rowLink} truncate font-semibold"
          href={guildHref(row.guild.region, row.guild.ruleset, row.guild.name)}
        >
          {row.guild.name}
        </a>
        <span class="text-muted numeric text-[12px]">
          {homeTopGuildsCopy.bosses(Math.round(row.value))}
        </span>
      </li>
    {/each}
  </ul>
{/if}
