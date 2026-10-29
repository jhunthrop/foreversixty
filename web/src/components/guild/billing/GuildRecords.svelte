<!-- web/src/components/guild/billing/GuildRecords.svelte -->
<!-- Records (spec 3.2.7): every officer, every state. The two CSV downloads are plain links to the
     API origin (the file name comes from its Content-Disposition), so they work with no script. -->
<script lang="ts">
  import { exportHref, type ExportKind } from '../../../lib/billing/guild-billing-api';
  import { guildBillingCopy as copy } from '../../../lib/billing/guild-billing-copy';
  import { BILLING_OUTLINE_BUTTON } from '../../../lib/guild/billing-styles';

  let { guildId }: { guildId: number } = $props();

  const DOWNLOADS: readonly { kind: ExportKind; label: string }[] = [
    { kind: 'attendance', label: copy.records.attendance },
    { kind: 'loot', label: copy.records.loot },
  ];
</script>

<div class="flex flex-col gap-2" data-testid="guild-records">
  <h3 class="section-title">{copy.records.heading}</h3>
  <p class="text-muted text-[13px]">{copy.records.body}</p>
  <div class="flex flex-col gap-2 md:flex-row">
    {#each DOWNLOADS as download (download.kind)}
      <a
        href={exportHref(guildId, download.kind)}
        class="{BILLING_OUTLINE_BUTTON} w-full justify-center md:w-fit"
        data-testid={`guild-export-${download.kind}`}
      >
        {download.label}
      </a>
    {/each}
  </div>
</div>
