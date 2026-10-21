<!-- web/src/components/SetupCompanionDone.svelte -->
<!-- The companion step's done line on /setup: once the signed-in account holds a connected
     device, "Done. Connected as <device>, last seen <when>." Signed out, or with no device, it
     draws nothing; the row's height is reserved only for a signed-in visitor (the
     `setup-when-in` hint), so a signed-out page has no gap and a signed-in one does not jump. -->
<script lang="ts">
  import { listDevices } from '../lib/account/api';
  import { companionDoneLine } from '../lib/setup/progress';

  let line = $state<string | null>(null);

  $effect(() => {
    let live = true;
    listDevices()
      .then((devices) => {
        if (live) line = companionDoneLine(devices);
      })
      .catch(() => {
        // Signed out, or the read failed: the step simply has no done line.
      });
    return () => {
      live = false;
    };
  });
</script>

<p class="setup-when-in min-h-[22px] text-[14px]" data-testid="setup-companion-done">
  {#if line !== null}{line}{/if}
</p>
