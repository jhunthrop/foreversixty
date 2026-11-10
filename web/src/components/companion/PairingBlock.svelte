<!-- web/src/components/companion/PairingBlock.svelte -->
<!-- The fallback pairing block (logs spec 2026-11-07 section 4, setup spec U8): a label,
     `pairingCopy.companionFirst` and `Get a pairing code`. The code, its expiry and the success line
     replace the button in place; a poll watches for the companion using it
     (lib/reports/pairing-poll.ts). /logs' rail and /setup's companion Section draw this one block. -->
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
  import Button from '../ui/Button.svelte';
  import Code from '../ui/Code.svelte';

  const SECONDS_PER_MINUTE = 60;

  let {
    devices,
    onCode,
    onPaired,
    note = true,
  }: {
    devices: readonly Device[];
    /** The companion-first line above the code button; /logs leaves it out (its card says Get the companion). */
    note?: boolean;
    /** The code while it is out and unused; null once used, expired or never asked for. */
    onCode?: (code: string | null) => void;
    onPaired?: (device: Device) => void;
  } = $props();

  let pairing = $state<PairingCode | null>(null);
  let paired = $state<Device | null>(null);
  let busy = $state(false);
  let error = $state('');
  let poll: PairingPoll | null = null;

  function setCode(code: PairingCode | null): void {
    pairing = code;
    onCode?.(code?.code ?? null);
  }

  async function pair(): Promise<void> {
    busy = true;
    error = '';
    poll?.stop();
    paired = null;
    try {
      const before = new Set(devices.map((device) => device.id));
      const code = await pairDevice();
      setCode(code);
      poll = createPairingPoll(
        async () => companionsOf(await refreshDevices()),
        before,
        (device) => {
          paired = device;
          setCode(null);
          onPaired?.(device);
        },
        { expiresInMs: code.expires_in * 1000, onExpire: () => setCode(null) },
      );
      poll.start();
    } catch {
      error = pairingCopy.failed;
    } finally {
      busy = false;
    }
  }

  $effect(() => () => poll?.stop());

  const label = $derived(
    devices.length > 0 ? logsCompanionCopy.pairAnotherLabel : logsCompanionCopy.pairLabel,
  );
</script>

<div class="flex flex-col items-start gap-3" data-testid="logs-companion-pair">
  <span class="text-strong text-[15px] font-bold">{label}</span>
  {#if paired !== null}
    <p class="text-[14px]" data-testid="pairing-success">{pairingCopy.success(paired.name)}</p>
  {:else if pairing !== null}
    <Code text={pairing.code} big testid="pairing-code" />
    <p class="text-small text-muted">
      {pairingCopy.codeLine(Math.round(pairing.expires_in / SECONDS_PER_MINUTE))}
    </p>
  {:else}
    {#if note}<p class="text-muted text-[14px] leading-5" data-testid="pairing-companion-first">
        {pairingCopy.companionFirst}
      </p>{/if}
    <Button size="sm" onclick={() => void pair()} disabled={busy} testid="pairing-start"
      >{pairingCopy.action}</Button
    >
  {/if}
  {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="account-error">{error}</p>{/if}
</div>
