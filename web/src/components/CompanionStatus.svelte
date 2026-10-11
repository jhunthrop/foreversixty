<!-- web/src/components/CompanionStatus.svelte -->
<!-- Logs landing spec (2026-10-04) section 4.D.1: one honest status line per paired device, above
     the companion panel's pointer line. A pure render of the devices the page already read: the
     page needs the same list to decide the panel order, so it owns the one read. Nothing at all
     when no device is paired -- the pairing block below is that case's own call to action. -->
<script lang="ts">
  import type { Device } from '../lib/account/api';
  import { deviceStatusLine } from '../lib/reports/device-status';

  let { devices }: { devices: readonly Device[] } = $props();
</script>

{#if devices.length > 0}
  <ul class="reveal flex flex-col gap-1" data-testid="companion-status">
    {#each devices as device (device.id)}
      <li class="text-[13px]" style="color: #c9c2b2" data-testid="companion-status-line">
        {deviceStatusLine(device)}
      </li>
    {/each}
  </ul>
{/if}
