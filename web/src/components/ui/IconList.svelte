<!-- web/src/components/ui/IconList.svelte -->
<!-- A wrapping list of icon + name with no chip border (sim spec 2026-11-06 §4.5, ruling S8; it
     replaces the 30 px bordered buff chips, which drew a box each). An optional 12 px caps label
     heads it. A `dim` entry is one with no effect on this spec: muted text, its icon at 62%. An
     entry whose icon is missing keeps the dark socket at the icon's size (ui/SpellIcon).
     An entry with a `tip` opens the shared game tooltip from its icon and name. -->
<script lang="ts">
  import type { IconListItem } from '../../lib/ui/facts';
  import type { SpellIconSize } from '../../lib/ui/spell-icon';
  import SpellIcon from './SpellIcon.svelte';

  let {
    items,
    label,
    iconSize = 22,
    testid,
  }: { items: readonly IconListItem[]; label?: string; iconSize?: SpellIconSize; testid?: string } = $props();
</script>

{#if label}<h3 class="text-label text-muted mt-4 font-bold">{label}</h3>{/if}
<div class="flex flex-wrap gap-x-[18px] gap-y-2 pt-[10px]" data-testid={testid}>
  {#each items as item (item.id)}
    <!-- svelte-ignore a11y_no_noninteractive_tabindex -->
    <span
      class={[
        'inline-flex min-h-7 items-center gap-2 text-[14px] leading-5',
        item.dim ? 'text-muted' : 'text-text',
      ]}
      data-game-tip={item.tip}
      tabindex={item.tip === undefined ? undefined : 0}
      data-dim={item.dim ? 'true' : undefined}
    >
      <span class={['inline-flex shrink-0', item.dim && 'opacity-[.62]']}
        ><SpellIcon url={item.icon} size={iconSize} /></span
      >
      {item.name}
    </span>
  {/each}
</div>
