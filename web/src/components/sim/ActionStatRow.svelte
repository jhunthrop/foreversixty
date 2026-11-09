<!-- web/src/components/sim/ActionStatRow.svelte -->
<!-- One engine action as a row of an open StatTable (sim spec 2026-11-06 §4.5: the Rage sources,
     the aura uptimes): its icon, its name (the tooltip's keyboard stop), a muted sub line, the
     Meter and the figures. The icon is the action's own, else the buff it is; a key with no icon
     keeps the socket, never a broken image. -->
<script lang="ts">
  import { actionTip, resolveIconUrl } from '../../lib/sim/action-mark';
  import type { CardTables } from '../../lib/sim/card-tables.svelte';
  import type { StatFigure } from '../../lib/ui/stat-table';
  import Meter from '../ui/Meter.svelte';
  import StatRow from '../ui/StatRow.svelte';

  let {
    action,
    label,
    tables,
    figures,
    fraction,
    tone = 'accent',
    note,
    noteTone = 'muted',
    fallbackIcon = null,
    testid,
  }: {
    action: string | null;
    label: string;
    tables: CardTables;
    figures: readonly StatFigure[];
    /** The Meter's fill, 0 to 1. */
    fraction: number;
    tone?: 'accent' | 'muted';
    /** The sub line under the name. */
    note?: string;
    /** "−17.2 lost" reads in the wipe colour. */
    noteTone?: 'muted' | 'wipe';
    fallbackIcon?: string | null;
    testid?: string;
  } = $props();

  const ICON_SIZE = 24;
  const iconUrl = $derived(resolveIconUrl(action, tables.visuals, fallbackIcon));
  const opens = $derived(actionTip(action, (id) => tables.spells.fact(id) !== undefined));
</script>

<StatRow {figures} {testid}>
  {#snippet icon()}
    {#if iconUrl !== null}<img
        src={iconUrl}
        alt=""
        width={ICON_SIZE}
        height={ICON_SIZE}
        loading="lazy"
        decoding="async"
        class="mt-[2px] block rounded-[4px]"
        data-game-tip={opens.tip}
        data-bis-item={opens.item}
      />{/if}
  {/snippet}
  {#snippet name()}
    <!-- svelte-ignore a11y_no_noninteractive_tabindex -->
    <span
      data-game-tip={opens.tip}
      data-bis-item={opens.item}
      tabindex={opens.tip !== undefined || opens.item !== undefined ? 0 : undefined}>{label}</span
    >
  {/snippet}
  {#snippet sub()}{#if note !== undefined}<span class={noteTone === 'wipe' ? 'text-wipe' : ''}>{note}</span
      >{/if}{/snippet}
  {#snippet meter()}<Meter value={fraction} {tone} />{/snippet}
</StatRow>
