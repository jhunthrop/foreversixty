<!-- web/src/components/report/AbilityBar.svelte -->
<!-- The amount bar, segmented by ability. Colour comes from one token -- the actor's class
     colour -- at five decreasing opacities, largest ability first, with a hairline of page
     background between segments. Using a palette here would either invent hex values the
     design system forbids or repurpose the class colours, which it forbids harder. -->
<script lang="ts">
  import { formatAmount } from '../../lib/report/format';
  import type { Ability } from '../../lib/report/types';

  let {
    abilities,
    total,
    peak,
    color,
  }: { abilities: Ability[]; total: number; peak: number; color: string } = $props();

  const OPACITIES = [1, 0.85, 0.7, 0.58, 0.48];
  const ordered = $derived([...abilities].sort((a, b) => b.total - a.total));
  const widthPct = $derived(peak <= 0 ? 0 : Math.max((total / peak) * 100, 0.5));
</script>

<div class="bg-line-soft h-[8px] w-full" style={`max-width: ${widthPct}%`} aria-hidden="true">
  <div class="flex h-full w-full">
    {#each ordered as ability, index (ability.spell_id)}
      <span
        class="h-full"
        style={`flex: ${Math.max(ability.total, 0)} 0 0; background: ${color}; opacity: ${OPACITIES[Math.min(index, OPACITIES.length - 1)]}; border-right: 1px solid var(--color-bg)`}
        title={`${ability.name} ${formatAmount(ability.total)}`}
      ></span>
    {/each}
  </div>
</div>
