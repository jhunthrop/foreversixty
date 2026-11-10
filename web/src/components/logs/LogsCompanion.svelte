<!-- web/src/components/logs/LogsCompanion.svelte -->
<!-- /logs' rail panel "The companion" (logs spec 2026-11-07 section 4; directions b1-logs): one line and
     the button (the primary once signed in; signed out, Sign in is) to the companion's own step on /setup; signed in, the pairing block under it (a label and
     `Get a pairing code`; the companion-first line is /setup's, not this card's). A paired device's status line sits over
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

<Section title="The companion" anchor="companion" panel testid="logs-companion">
  <div class="flex flex-col items-start gap-3">
    <CompanionStatus {devices} />
    <p class="text-[15px] leading-6" data-testid="logs-companion-pointer">{logsCompanionCopy.pointer}</p>
    <Button
      tone={signedIn ? 'primary' : 'secondary'}
      href={logsCompanionCopy.pointerHref}
      testid="logs-companion-link">{logsCompanionCopy.pointerLink}</Button
    >
    {#if signedIn}
      <PairingBlock {devices} note={false} />
    {/if}
  </div>
</Section>
