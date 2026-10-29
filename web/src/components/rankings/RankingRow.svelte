<!-- web/src/components/rankings/RankingRow.svelte -->
<!-- One ranked kill (design/specs/2026-10-11-rankings.md section 3, decision 6). One element for
     every width: the grid in lib/rankings/layout.ts shows every column from 1100 px, drops Size
     and Length (the guild moves under the name) below it, and below 600 px folds the rest into
     a muted second line under the name with the figure and Report beside it. The rank keeps its
     parse-ladder colour; the name is its class colour, a link to the character, and its link
     covers the row's name area so the whole of it is one touch target. Build is plain text with
     an underline on hover; Report is the one gold link. The viewer's own row carries a gold
     edge and tint in place. A kill a moderator touched gets one plain line under the name. -->
<script lang="ts">
  import { classCrestSrc } from '../../lib/class-crest';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import { characterHref, guildHref, parseCharacterKey, splitUnitName } from '../../lib/characters';
  import { classColorVar, formatAmount, formatDuration } from '../../lib/report/format';
  import { executionHref, executionLabel, executionTitle } from '../../lib/sim/execution';
  import type { RankingRow } from '../../lib/rankings/api';
  import { rankingsCopy } from '../../lib/rankings/copy';
  import { RANKING_ROW_GRID, TABLE_UP, WIDE_ONLY } from '../../lib/rankings/layout';
  import { moderationLine, rankColor, rankingDate } from '../../lib/rankings/row-view';

  let {
    row,
    total,
    own = false,
    crest = false,
    now,
    testid,
  }: {
    row: RankingRow;
    /** The board's size, for the rank's colour. */
    total: number;
    own?: boolean;
    /** A board of more than one class shows a crest before each name. */
    crest?: boolean;
    now: Date;
    testid: string;
  } = $props();

  const name = $derived(splitUnitName(row.player.name).name);
  const key = $derived(parseCharacterKey(row.player.key));
  const characterLink = $derived(
    key === null ? null : characterHref(key.region, key.ruleset, row.player.name),
  );
  const region = $derived(row.guild?.region ?? key?.region ?? '');
  const buildLink = $derived(row.build_id === undefined ? null : `/b/${row.build_id}`);
  const moderation = $derived(moderationLine(row.state));
  const executed = $derived(executionLabel(row.execution_score));
  const guildLine = $derived(
    [row.guild?.name, region.toUpperCase()].filter((part) => part !== undefined && part !== '').join(' · '),
  );
  const nameClass =
    'truncate font-semibold after:absolute after:inset-0 after:content-[""] focus-visible:outline-2 focus-visible:outline-offset-2';
</script>

{#snippet build(extra: string, buildTestid: string)}
  {#if buildLink !== null}
    <a
      class="text-text numeric relative z-10 text-[13px] underline-offset-4 hover:underline {extra}"
      href={buildLink}
      data-testid={buildTestid}>{row.talent_split}</a
    >
  {:else}
    <span class="text-muted numeric text-[13px] {extra}">{row.talent_split}</span>
  {/if}
{/snippet}

<li
  class="{RANKING_ROW_GRID} border-line-soft hover:bg-card-top data-[own=true]:border-l-gold data-[own=true]:bg-gold/6 relative grid min-h-[60px] items-center gap-x-2 border-b border-l-2 border-l-transparent px-2.5 py-1.5 text-[14px] min-[600px]:min-h-11 min-[600px]:gap-x-3 min-[600px]:py-0 min-[600px]:pr-3.5 min-[600px]:pl-3"
  data-own={own ? 'true' : undefined}
  data-testid={testid}
>
  <span class="numeric text-[12px]" style:color={rankColor(row.rank, total)} data-testid="ranking-rank"
    >{row.rank}</span
  >

  <span class="relative flex min-w-0 flex-col justify-center gap-px py-1" data-testid="ranking-name-cell">
    <span class="flex min-w-0 items-center gap-2.5">
      {#if crest}
        <!-- The ringed crest is drawn here rather than through ClassCrestRing: a row only exists
             after hydration, so a component's own scoped CSS never reaches the page. -->
        <img
          class="size-[22px] shrink-0 rounded-full object-cover min-[600px]:size-7"
          src={classCrestSrc(classSlugFromName(row.player.class))}
          alt=""
          width="28"
          height="28"
          loading="lazy"
          decoding="async"
          style:box-shadow={`0 0 0 2px color-mix(in srgb, ${classColorVar(row.player.class)} 55%, transparent)`}
        />
      {/if}
      {#if characterLink !== null}
        <a
          class={nameClass}
          style:color={classColorVar(row.player.class)}
          href={characterLink}
          data-testid="ranking-character">{name}</a
        >
      {:else}
        <!-- player.key did not parse into a region and ruleset: an unlinked name is the legible
             failure, not a guessed link that would point at the wrong character. -->
        <span class="truncate font-semibold" style:color={classColorVar(row.player.class)}>{name}</span>
      {/if}
    </span>
    <span class="text-muted truncate text-[12px] min-[600px]:hidden" data-testid="ranking-meta">
      {rankingsCopy.sizeMan(row.size)} · {formatDuration(row.duration_ms)} · {executed} ·
      {@render build('', 'ranking-build-phone')}
    </span>
    {#if guildLine !== ''}
      <span class="text-muted hidden truncate text-[12px] min-[600px]:block min-[1100px]:hidden">
        {guildLine}
      </span>
    {/if}
    {#if moderation !== null}
      <span class="text-muted text-[12px]" data-testid="ranking-state">{moderation}</span>
    {/if}
  </span>

  <span class="text-muted truncate text-[13px] {WIDE_ONLY}">
    {#if row.guild}
      <a
        class="text-nav relative z-10 underline-offset-4 hover:underline"
        href={guildHref(row.guild.region, row.guild.ruleset, row.guild.name)}>{row.guild.name}</a
      >
      · {row.guild.region.toUpperCase()}
    {/if}
  </span>
  <span class="text-muted numeric text-right text-[13px] {WIDE_ONLY}">{row.size}</span>
  <span class="numeric text-right" data-testid="ranking-value">{formatAmount(Math.round(row.value))}</span>

  {#if row.execution_score === null}
    <span
      class="text-muted numeric text-right text-[13px] {TABLE_UP}"
      title={executionTitle(null)}
      aria-label={executionTitle(null)}
      data-testid="ranking-execution">{executed}</span
    >
  {:else}
    <a
      class="numeric relative z-10 hidden min-h-11 items-center justify-end text-right text-[13px] min-[600px]:inline-flex min-[600px]:min-h-0 min-[600px]:self-stretch"
      href={executionHref(row.report_id, row.fight_index)}
      title={executionTitle(row.execution_score)}
      data-testid="ranking-execution">{executed}</a
    >
  {/if}
  <span class="text-muted numeric text-right text-[13px] {WIDE_ONLY}">{formatDuration(row.duration_ms)}</span>
  <span class="text-muted numeric text-right text-[13px] {TABLE_UP}">{rankingDate(row.fought_at, now)}</span>
  {@render build(`${TABLE_UP} text-right`, 'ranking-build')}
  <a
    class="text-gold hover:text-gold-hover relative z-10 inline-flex min-h-11 items-center justify-end text-[13px] font-semibold min-[600px]:min-h-0 min-[600px]:self-stretch"
    href={`/reports/${row.report_id}?fight=${row.fight_index}`}
    data-testid="ranking-report">Report</a
  >
</li>
