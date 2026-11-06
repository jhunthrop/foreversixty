<!-- web/src/components/guild/loot/LootExport.svelte -->
<!-- Export history (web layout spec 2026-11-05 §7.2, officers): an optional from and to date (the
     night's day, inclusive), then the CSV and JSON downloads of every award in range with its
     MS/OS, votes, reason and who awarded it. Empty dates take the whole history. -->
<script lang="ts">
  import { lootHistoryHref } from '../../../lib/guild/loot-council-api';
  import { LINK } from '../../../lib/ui/primitives';
  import Field from '../../ui/Field.svelte';

  let { guildId, apiBase }: { guildId: number; apiBase?: string } = $props();

  let from = $state('');
  let to = $state('');
  const range = $derived({ from, to });
</script>

<div class="flex flex-col gap-3" data-testid="guild-loot-export">
  <div class="flex flex-wrap gap-3">
    <Field
      id="guild-loot-export-from"
      label="From"
      type="date"
      bind:value={from}
      testid="guild-loot-export-from"
    />
    <Field id="guild-loot-export-to" label="To" type="date" bind:value={to} testid="guild-loot-export-to" />
  </div>
  <p class="flex flex-wrap gap-x-5 gap-y-2 text-[14px]">
    <a class={LINK} href={lootHistoryHref(guildId, 'csv', range, apiBase)} data-testid="guild-loot-export-csv"
      >Download CSV ›</a
    >
    <a
      class={LINK}
      href={lootHistoryHref(guildId, 'json', range, apiBase)}
      data-testid="guild-loot-export-json">Download JSON ›</a
    >
  </p>
</div>
