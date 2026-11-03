<!-- web/src/components/GameTooltip.svelte -->
<!-- The game's spell and talent tooltip: name with its rank on the right, the cost beside the
     range, the cast time beside the cooldown, then the description in the client's gold and,
     for a talent short of its last rank, the next rank's text. The same panel, placement and
     open rules as ItemTooltip.svelte: the shared host (BisTooltipHost.svelte) lazy-loads this
     component and mounts it only while open, so a page carries none of its weight until a
     reader asks for a tooltip. Pure presentation over a GameTipModel (lib/tooltips/game-tip.ts). -->
<script lang="ts">
  import type { GameTipModel } from '../lib/tooltips/game-tip';
  import { dataUrl } from '../lib/planner/load';

  let { model, build, id }: { model: GameTipModel; build: string; id: string } = $props();

  let iconBroken = $state(model.icon === undefined);
</script>

<div
  {id}
  role="tooltip"
  data-testid="game-tooltip"
  class="border-line bg-raised rounded-panel absolute z-30 flex w-[clamp(260px,calc(100vw_-_32px),340px)] flex-col gap-1.5 border p-3 text-left text-[13px] font-normal whitespace-normal normal-case shadow-[0_12px_30px_rgba(0,0,0,.45)]"
  style:left="var(--item-tooltip-shift, 0px)"
  style:top="var(--item-tooltip-top, 100%)"
  style:bottom="var(--item-tooltip-bottom, auto)"
  style:margin-top="var(--item-tooltip-mt, 8px)"
  style:margin-bottom="var(--item-tooltip-mb, 0px)"
  style:max-height="var(--item-tooltip-max-h, none)"
  style:overflow-y="var(--item-tooltip-overflow-y, visible)"
>
  <div class="flex items-start gap-2">
    {#if !iconBroken && model.icon !== undefined}
      <img
        src={dataUrl(build, `icons/${model.icon}.webp`)}
        alt=""
        width="36"
        height="36"
        loading="lazy"
        decoding="async"
        class="rounded-control border-line h-9 w-9 shrink-0 border object-cover"
        onerror={() => (iconBroken = true)}
        onload={(event) => {
          if (event.currentTarget.naturalWidth <= 1) iconBroken = true;
        }}
      />
    {/if}
    <div class="flex min-w-0 flex-1 items-baseline justify-between gap-2">
      <span class="text-strong text-[15px] font-semibold" data-testid="game-tooltip-name">{model.name}</span>
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
      <span class="label text-muted block">Next rank</span>
      {model.nextRank}
    </p>
  {/if}
</div>
