<!-- web/src/components/GameTooltip.svelte -->
<!-- The game's spell and talent tooltip: name with its rank on the right, the cost beside the
     range, the cast time beside the cooldown, then the description in the client's gold and,
     for a talent short of its last rank, the next rank's text. The same panel, placement and
     open rules as ItemTooltip.svelte: the shared host (BisTooltipHost.svelte) lazy-loads this
     component and mounts it only while open, so a page carries none of its weight until a
     reader asks for a tooltip. The frame is the Overlay primitive's tooltip kind and the spell icon is
     ui/ItemIcon in its plain frame. Pure presentation over a GameTipModel (lib/tooltips/game-tip.ts). -->
<script lang="ts">
  import type { GameTipModel } from '../lib/tooltips/game-tip';
  import { TOOLTIP_PANEL_PLACEMENT } from '../lib/items/tooltip-position';
  import ItemIcon from './ui/ItemIcon.svelte';
  import Overlay from './ui/Overlay.svelte';

  let { model, build, id }: { model: GameTipModel; build: string; id: string } = $props();
</script>

<Overlay
  kind="tooltip"
  {id}
  testid="game-tooltip"
  class="absolute z-30 flex w-[clamp(260px,calc(100vw_-_32px),340px)] flex-col gap-1.5 px-[14px] py-3 text-left text-[13px] font-normal whitespace-normal normal-case"
  style={TOOLTIP_PANEL_PLACEMENT}
>
  <div class="flex items-start gap-2">
    {#if model.icon !== undefined}
      <ItemIcon icon={model.icon} {build} size={36} />
    {/if}
    <div class="flex min-w-0 flex-1 items-baseline justify-between gap-2">
      <span class="text-strong text-[15px] font-bold" data-testid="game-tooltip-name">{model.name}</span>
      {#if model.subtitle !== undefined}
        <span class="text-muted numeric shrink-0 text-[12px]">{model.subtitle}</span>
      {/if}
    </div>
  </div>
  {#each model.lines as line, index (index)}
    <div class="text-strong flex flex-wrap justify-between gap-x-3 text-[12px]">
      <span>{line.left}</span>
      {#if line.right !== undefined}<span>{line.right}</span>{/if}
    </div>
  {/each}
  {#if model.description !== ''}
    <p class="text-gold text-[12px] leading-snug whitespace-pre-line" data-testid="game-tooltip-description">
      {model.description}
    </p>
  {/if}
  {#if model.nextRank !== undefined}
    <p class="text-muted text-[12px] leading-snug whitespace-pre-line">
      <span class="text-muted block font-semibold">Next rank</span>
      {model.nextRank}
    </p>
  {/if}
</Overlay>
