<!-- web/src/components/HomeTopGuilds.svelte -->
<!-- The guild landing's Top guilds (web layout spec 2026-11-05 §6.1): the top five guilds by
     progression, from `GET /v1/rankings/guilds?kind=progress` (the one guild leaderboard kind that
     needs no `encounter`), as ranked Rows with a Meter against the leader in the page's one Plate.
     Loading holds the five rows' height; a failure and an empty board are open lines. -->
<script lang="ts">
  import { guildHref, rulesetLabel } from '../lib/characters';
  import { fetchGuildRankings, type GuildRankingRow } from '../lib/rankings/api';
  import { homeTopGuildsCopy } from '../lib/rankings/copy';
  import LoadError from './ui/LoadError.svelte';
  import Meter from './ui/Meter.svelte';
  import Plate from './ui/Plate.svelte';
  import Row from './ui/Row.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  const ROWS = 5;

  let { label = homeTopGuildsCopy.placeholderLabel }: { label?: string } = $props();

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

  const top = $derived(Math.max(1, ...rows.map((row) => row.value)));
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
  <p class="text-muted text-[14px]" data-testid="home-top-guilds-empty">{homeTopGuildsCopy.empty}</p>
{:else}
  <Plate {label} testid="home-top-guilds">
    {#each rows as row (`${row.rank}-${row.guild.region}-${row.guild.ruleset}-${row.guild.name}`)}
      {@const bosses = homeTopGuildsCopy.bosses(Math.round(row.value))}
      <Row
        title={row.guild.name}
        rank={String(row.rank)}
        sub={`${row.guild.region.toUpperCase()} ${rulesetLabel(row.guild.ruleset)}`}
        href={guildHref(row.guild.region, row.guild.ruleset, row.guild.name)}
      >
        {#snippet trail()}
          <span class="hidden w-24 md:block"
            ><Meter value={row.value / top} label={`${row.guild.name}: ${bosses}`} /></span
          >
          <span class="text-text w-[72px] text-right text-[14px] font-semibold tabular-nums">{bosses}</span>
        {/snippet}
      </Row>
    {/each}
  </Plate>
{/if}
