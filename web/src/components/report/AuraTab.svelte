<!-- web/src/components/report/AuraTab.svelte -->
<!-- The Buffs and Debuffs tabs in one lazily loaded chunk (logs spec 2026-11-07 6.2): the aura
     table (its own Section and Plate), and on Buffs the raid cooldowns under it. -->
<script lang="ts">
  import type { ComponentProps } from 'svelte';
  import AuraTable from './AuraTable.svelte';
  import RaidCooldowns from './RaidCooldowns.svelte';

  let {
    table,
    cooldowns,
  }: {
    table: ComponentProps<typeof AuraTable>;
    /** Buffs only: the raid cooldowns lanes under the table. */
    cooldowns?: ComponentProps<typeof RaidCooldowns>;
  } = $props();
</script>

<AuraTable {...table} />
{#if cooldowns}<RaidCooldowns {...cooldowns} />{/if}
