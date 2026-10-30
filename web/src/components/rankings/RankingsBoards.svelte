<!-- web/src/components/rankings/RankingsBoards.svelte -->
<!-- The Characters/Guilds switch (design/specs/2026-10-11-rankings.md section 3, decision 1), the
     board card's one control, in its head's right slot (site alignment spec 2026-10-29 §5.7). -->
<script lang="ts">
  import { rankingsCopy } from '../../lib/rankings/copy';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import type { RankingsState } from '../../lib/rankings/url';

  const BOARDS = [
    { id: 'character', label: 'Characters' },
    { id: 'guild', label: 'Guilds' },
  ] as const;

  let {
    board,
    onboard,
  }: {
    board: RankingsState['board'];
    onboard: (board: RankingsState['board']) => void;
  } = $props();
</script>

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
