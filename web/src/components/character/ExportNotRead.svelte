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

<!-- No frame (web layout spec 2026-11-05 §3A): the sentence leads in the body face 600, the two
     lines under it muted; it sits inside the section that would have held the slots. -->
<div class="flex flex-col gap-1 py-2 text-[14px]" role="status" data-testid={testid}>
  <p class="text-strong text-[15px] font-semibold" data-testid="export-not-read-lead">
    {owned ? notReadCopy.lead : notReadCopy.visitorLead}
  </p>
  {#if owned}
    <p class="text-muted">{notReadCopy.synced(stamp)}</p>
    <p class="text-muted">{notReadCopy.steps}</p>
  {:else}
    <p class="text-muted">{notReadCopy.visitorSynced}</p>
  {/if}
</div>
