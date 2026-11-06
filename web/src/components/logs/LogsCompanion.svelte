<!-- web/src/components/logs/LogsCompanion.svelte -->
<!-- /logs' rail Section "The companion" (logs spec 2026-11-07 section 4): one line and the secondary
     button to the companion's own step on /setup; signed in, the pairing block under it (a label,
     `pairingCopy.companionFirst`, `Get a pairing code`). A paired device's status line sits over
     the pointer. The code, its expiry and the success line replace the button in place; a poll
     watches for the companion using it (lib/reports/pairing-poll.ts). -->
<script lang="ts">
  import {
    companionsOf,
    pairDevice,
    refreshDevices,
    type Device,
    type PairingCode,
  } from '../../lib/account/api';
  import { logsCompanionCopy, pairingCopy } from '../../lib/reports/copy';
  import { createPairingPoll, type PairingPoll } from '../../lib/reports/pairing-poll';
  import CompanionStatus from '../CompanionStatus.svelte';
  import Button from '../ui/Button.svelte';
  import Section from '../ui/Section.svelte';

  const PAIR_ACTION = 'Get a pairing code';
  const PAIR_FAILED = 'The pairing code did not load. Try again.';
  const SECONDS_PER_MINUTE = 60;

  let { signedIn, devices }: { signedIn: boolean; devices: readonly Device[] } = $props();

  let pairing = $state<PairingCode | null>(null);
  let paired = $state<Device | null>(null);
  let busy = $state(false);
  let error = $state('');
  let poll: PairingPoll | null = null;

  async function pair(): Promise<void> {
    busy = true;
    error = '';
    poll?.stop();
    paired = null;
    try {
      const before = new Set(devices.map((device) => device.id));
      const code = await pairDevice();
      pairing = code;
      poll = createPairingPoll(
        async () => companionsOf(await refreshDevices()),
        before,
        (device) => {
          paired = device;
          pairing = null;
        },
        { expiresInMs: code.expires_in * 1000 },
      );
      poll.start();
    } catch {
      error = PAIR_FAILED;
    } finally {
      busy = false;
    }
  }

  $effect(() => () => poll?.stop());

  const label = $derived(
    devices.length > 0 ? logsCompanionCopy.pairAnotherLabel : logsCompanionCopy.pairLabel,
  );
</script>

<Section title="The companion" meta="Live, while you raid" anchor="companion" testid="logs-companion">
  <div class="flex flex-col items-start gap-3 pt-1">
    <CompanionStatus {devices} />
    <p class="text-[15px] leading-6" data-testid="logs-companion-pointer">{logsCompanionCopy.pointer}</p>
    <Button href={logsCompanionCopy.pointerHref} testid="logs-companion-link"
      >{logsCompanionCopy.pointerLink}</Button
    >
    {#if signedIn}
      <div class="flex flex-col items-start gap-3" data-testid="logs-companion-pair">
        <span class="label text-muted">{label}</span>
        {#if paired !== null}
          <p class="text-[14px]" data-testid="pairing-success">{pairingCopy.success(paired.name)}</p>
        {:else if pairing !== null}
          <p class="tabular text-strong font-mono text-[24px]" data-testid="pairing-code">{pairing.code}</p>
          <p class="text-small text-muted">
            Type this into the companion within {Math.round(pairing.expires_in / SECONDS_PER_MINUTE)} minutes. It
            pairs one device and cannot be reused.
          </p>
        {:else}
          <p class="text-muted text-[15px] leading-6" data-testid="pairing-companion-first">
            {pairingCopy.companionFirst}
          </p>
          <Button size="sm" onclick={() => void pair()} disabled={busy} testid="pairing-start"
            >{PAIR_ACTION}</Button
          >
        {/if}
        {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="account-error">{error}</p>{/if}
      </div>
    {/if}
  </div>
</Section>
