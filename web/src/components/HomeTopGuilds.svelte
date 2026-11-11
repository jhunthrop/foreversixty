<!-- web/src/components/HomeTopGuilds.svelte -->
<!-- The guild landing's Top guilds (web layout spec 2026-11-05 §6.1): the top five guilds by
     progression, from `GET /v1/rankings/guilds?kind=progress` (the one guild leaderboard kind that
     needs no `encounter`), as ranked Rows (where the guild lives, its newest kill, the bosses down), in the landing's
     Top guilds panel.
     Loading holds the five rows' height; a failure and an empty board are open lines. -->
<script lang="ts">
  import EmptyState from './ui/EmptyState.svelte';
  import { guildHref, rulesetLabel } from '../lib/characters';
  import { fetchGuildRankings, type GuildRankingRow } from '../lib/rankings/api';

  import { homeTopGuildsCopy } from '../lib/rankings/copy';
  import { viewerZone, weekdayDayMonth } from '../lib/dates';
  import LoadError from './ui/LoadError.svelte';
  import Row from './ui/Row.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  const ROWS = 5;

  let {
    label = homeTopGuildsCopy.placeholderLabel,
    viewerTimezone = viewerZone(),
  }: { label?: string; viewerTimezone?: string } = $props();

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
  <EmptyState message={homeTopGuildsCopy.empty} testid="home-top-guilds-empty" />
{:else}
  <div role="group" aria-label={label} data-testid="home-top-guilds">
    {#each rows as row (`${row.rank}-${row.guild.region}-${row.guild.ruleset}-${row.guild.name}`)}
      {@const where = `${row.guild.region.toUpperCase()} ${rulesetLabel(row.guild.ruleset)}`}
      {@const killed = homeTopGuildsCopy.lastKill(weekdayDayMonth(new Date(row.fought_at), viewerTimezone))}
      <Row
        title={row.guild.name}
        rank={String(row.rank)}
        sub={`${where} · ${killed}`}
        href={guildHref(row.guild.region, row.guild.ruleset, row.guild.name)}
      >
        {#snippet trail()}
          <span class="text-text text-right text-[14px] font-semibold tabular-nums"
            >{homeTopGuildsCopy.bossesDown(Math.round(row.value))}</span
          >
        {/snippet}
      </Row>
    {/each}
  </div>
{/if}
