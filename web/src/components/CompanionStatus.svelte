<!-- web/src/components/CompanionStatus.svelte -->
<!-- Logs landing spec (2026-10-04) §4.D.1: one honest status line per paired device, above
     the companion panel's existing `logsCompanionCopy.pointer` line. Reads `fetchMeOnce`/
     `listDevices` -- both already shared through the client cache (web/src/lib/data/
     query.ts), the same way `Account.svelte`'s own `pairing`/`account` modes read them, so
     this island costs no second network round trip, only a second cache read. Nothing at
     all for a signed-out visitor (no devices to have), and nothing for a signed-in visitor
     with none paired yet -- the pairing block below is that case's own call to action. -->
<script lang="ts">
  import { fetchMeOnce, listDevices, type Device } from '../lib/account/api';
  import { deviceStatusLine } from '../lib/reports/device-status';

  let signedIn = $state(false);
  let devices = $state<Device[]>([]);

  $effect(() => {
    void (async () => {
      const me = await fetchMeOnce();
      signedIn = me !== null;
      if (me !== null) devices = await listDevices();
    })();
  });
</script>

{#if signedIn && devices.length > 0}
  <ul class="reveal flex flex-col gap-1" data-testid="companion-status">
    {#each devices as device (device.id)}
      <li class="text-[13px]" style="color: #c9c2b2" data-testid="companion-status-line">
        {deviceStatusLine(device)}
      </li>
    {/each}
  </ul>
{/if}
