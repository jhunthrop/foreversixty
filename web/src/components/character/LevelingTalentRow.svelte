<!-- web/src/components/character/LevelingTalentRow.svelte -->
<!-- The Talents card's leveling order below 20 (design/specs/2026-11-04-low-level-home.md §4b):
     the next few points of the guide's order as talent icons with their rank, the point to spend
     now lit, each opening the game's own talent tooltip (data-game-tip, the guide chips' keys,
     resolved by the leveling spec's guide bundle through GuideTooltips). -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import type { LevelingPoint } from '../../lib/leveling/leveling-plan';
  import { dataUrl } from '../../lib/planner/load';
  import { talentTipKey } from '../../lib/tooltips/game-tip';

  let { points, lit }: { points: readonly LevelingPoint[]; lit: boolean } = $props();
</script>

<ul class="flex items-center gap-2" data-testid="home-talents-order">
  {#each points as point, index (point.level)}
    <li
      class={[
        'gtip relative flex h-7 w-7 shrink-0 cursor-help items-center justify-center rounded-[4px] border',
        index === 0 && lit
          ? 'border-gold shadow-[0_0_0_1px_var(--color-gold)]'
          : 'border-white/15 opacity-70',
      ]}
      tabindex="0"
      data-game-tip={talentTipKey(point.talentId, point.rank)}
      data-testid={`home-talents-tile-${point.talentId}`}
      aria-label={`Level ${point.level}: ${point.name} ${point.rank}/${point.maxRank}`}
    >
      <img
        src={dataUrl(activeBuild.build, `icons/${point.icon}.webp`)}
        alt=""
        width="26"
        height="26"
        class="rounded-[3px]"
        loading="lazy"
        decoding="async"
      />
      <span
        class="numeric text-strong absolute -right-1.5 -bottom-1.5 rounded-[3px] bg-black/85 px-[3px] text-[10px] leading-[13px]"
        >{point.rank}</span
      >
    </li>
  {/each}
</ul>
