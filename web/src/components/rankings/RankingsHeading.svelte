<!-- web/src/components/rankings/RankingsHeading.svelte -->
<!-- The board's heading row (design/specs/2026-10-11-rankings.md section 3, decision 1): the scope
     ("Fury Warriors", "Every class and spec") with the quiet "Another class or spec" link, the
     one-line colour legend under it, and the Characters/Guilds tabs at the right. It repeats no
     count: the header already says it. -->
<script lang="ts">
  import { rankingsCopy } from '../../lib/rankings/copy';
  import { RANK_LEGEND } from '../../lib/rankings/row-view';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import type { RankingsState } from '../../lib/rankings/url';

  const BOARDS = [
    { id: 'character', label: 'Characters' },
    { id: 'guild', label: 'Guilds' },
  ] as const;

  let {
    scope,
    board,
    canChooseAnother,
    onanother,
    onboard,
  }: {
    scope: string;
    board: RankingsState['board'];
    /** Signed in and looking at the character's own board: offer the other boards. */
    canChooseAnother: boolean;
    onanother: () => void;
    onboard: (board: RankingsState['board']) => void;
  } = $props();
</script>

<div class="flex flex-wrap items-end justify-between gap-x-4 gap-y-2" data-testid="rankings-heading">
  <div class="flex flex-col gap-0.5">
    <div class="flex flex-wrap items-baseline gap-x-4">
      <h2 class="section-title" data-testid="rankings-scope">{scope}</h2>
      {#if canChooseAnother}
        <button
          type="button"
          class="text-nav inline-flex min-h-11 items-center text-[13px] underline underline-offset-4"
          data-testid="rankings-another"
          onclick={onanother}>{rankingsCopy.anotherBoard}</button
        >
      {/if}
    </div>
    <p
      class="text-muted flex flex-wrap items-center gap-x-2 gap-y-1.5 text-[12px]"
      data-testid="rankings-legend"
    >
      {rankingsCopy.legendLead}
      {#each RANK_LEGEND as step (step.label)}
        <!-- Tag's shape (22 px, 11 px 700 uppercase, 1 px border) in the parse colour; Tag
             itself has four state tones only. -->
        <span
          class="inline-flex h-[22px] items-center rounded-full border border-current/40 bg-current/10 px-2 text-[11px] font-bold tracking-[.08em] whitespace-nowrap uppercase"
          style:color={step.color}>{step.label}</span
        >
      {/each}
    </p>
  </div>

  <div role="tablist" aria-label={rankingsCopy.boardTabs} class="flex items-center gap-1.5">
    {#each BOARDS as option (option.id)}
      <button
        type="button"
        role="tab"
        class="{SECONDARY_BUTTON_FIXED} px-3.5"
        class:border-gold={board === option.id}
        class:text-gold={board === option.id}
        class:border-line={board !== option.id}
        class:text-nav={board !== option.id}
        aria-selected={board === option.id}
        data-testid={`board-${option.id}`}
        onclick={() => onboard(option.id)}
      >
        {option.label}
      </button>
    {/each}
  </div>
</div>
