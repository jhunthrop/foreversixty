<!-- web/src/components/guild/billing/GuildRecords.svelte -->
<!-- Records (spec 3.2.7): every officer, every state, as two link rows of the Records card. The two CSV downloads are plain links to the
     API origin (the file name comes from its Content-Disposition), so they work with no script. -->
<script lang="ts">
  import { exportHref, type ExportKind } from '../../../lib/billing/guild-billing-api';
  import { guildBillingCopy as copy } from '../../../lib/billing/guild-billing-copy';
  import SettingsRow from '../../ui/SettingsRow.svelte';

  let { guildId }: { guildId: number } = $props();

  const DOWNLOADS: readonly { kind: ExportKind; label: string }[] = [
    { kind: 'attendance', label: copy.records.attendance },
    { kind: 'loot', label: copy.records.loot },
  ];
</script>

<div data-testid="guild-records">
  {#each DOWNLOADS as download (download.kind)}
    <SettingsRow
      glyph="download"
      label={download.label}
      href={exportHref(guildId, download.kind)}
      testid={`guild-export-${download.kind}`}
    />
  {/each}
</div>
