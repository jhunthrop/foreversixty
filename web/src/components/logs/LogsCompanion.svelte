<!-- web/src/components/logs/LogsCompanion.svelte -->
<!-- /logs' rail Section "The companion" (logs spec 2026-11-07 section 4): one line and the secondary
     button to the companion's own step on /setup; signed in, the pairing block under it (a label,
     `pairingCopy.companionFirst`, `Get a pairing code`). A paired device's status line sits over
     the pointer. The pairing block is companion/PairingBlock.svelte, shared with /setup. -->
<script lang="ts">
  import type { Device } from '../../lib/account/api';
  import { logsCompanionCopy } from '../../lib/reports/copy';
  import CompanionStatus from '../CompanionStatus.svelte';
  import PairingBlock from '../companion/PairingBlock.svelte';
  import Button from '../ui/Button.svelte';
  import Section from '../ui/Section.svelte';

  let { signedIn, devices }: { signedIn: boolean; devices: readonly Device[] } = $props();
</script>

<Section title="The companion" meta="Live, while you raid" anchor="companion" testid="logs-companion">
  <div class="flex flex-col items-start gap-3 pt-1">
    <CompanionStatus {devices} />
    <p class="text-[15px] leading-6" data-testid="logs-companion-pointer">{logsCompanionCopy.pointer}</p>
    <Button href={logsCompanionCopy.pointerHref} testid="logs-companion-link"
      >{logsCompanionCopy.pointerLink}</Button
    >
    {#if signedIn}
      <PairingBlock {devices} />
    {/if}
  </div>
</Section>
