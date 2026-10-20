<!-- web/src/components/character/ExportNotRead.svelte -->
<!-- The one panel an export that carried no gear and no talents is drawn as (W.3), on the
     character page, /sim and the planner's character card: "Your gear has not been read yet."
     and what to do about it. It stands where the slot list, the "Empty" rows and "0 talent
     points" used to, which read as a verdict on the character. `capturedAt` is the sync's ISO
     time when the page knows it (the owner's /v1/me row); `owned` false is someone else's
     character, whose steps are not the visitor's. -->
<script lang="ts">
  import { formatSyncStamp } from '../../lib/dates';
  import { notReadCopy } from '../../lib/gear/copy';

  let {
    capturedAt = null,
    owned = true,
    testid = 'export-not-read',
  }: { capturedAt?: string | null; owned?: boolean; testid?: string } = $props();

  const stamp = $derived(capturedAt === null ? '' : formatSyncStamp(new Date(capturedAt)));
</script>

<div
  class="border-line-warm bg-raised rounded-panel flex flex-col gap-1 border px-4 py-3 text-[14px]"
  role="status"
  data-testid={testid}
>
  {#if owned}
    <p class="text-strong font-semibold" data-testid="export-not-read-lead">{notReadCopy.lead}</p>
    <p class="text-muted">{notReadCopy.synced(stamp)}</p>
    <p class="text-muted">{notReadCopy.steps}</p>
  {:else}
    <p class="text-strong font-semibold" data-testid="export-not-read-lead">{notReadCopy.visitorLead}</p>
    <p class="text-muted">{notReadCopy.visitorSynced}</p>
  {/if}
</div>
