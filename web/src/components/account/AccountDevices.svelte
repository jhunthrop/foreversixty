<!-- web/src/components/account/AccountDevices.svelte -->
<!-- /account's Devices section in the rail (premium spec 2026-11-08 section 7.5): a Row per paired
     companion (its name, then the platform in words and when it was last seen, deviceStatusLine;
     never the raw platform), Revoke asking first in the row. With none: why, the download, and the
     pairing code as the fallback; with some: Pair another computer. A shown code is the Code big
     under its one line. Each row says when it was seen, so the section has no Updated stamp. -->
<script lang="ts">
  import { smallButton } from '../../lib/account/small-button';
  import { companionsOf, type Device, type PairingCode } from '../../lib/account/api';
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import { devicesCopy, pairingEntryLabel } from '../../lib/account/devices-copy';
  import { deviceStatusLine } from '../../lib/reports/device-status';
  import { pairingCopy } from '../../lib/reports/copy';
  import { button } from '../../lib/ui/button';
  import { LINK } from '../../lib/ui/primitives';
  import Code from '../ui/Code.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import ConfirmButtons from '../ui/ConfirmButtons.svelte';

  let {
    devices,
    pairing,
    pairedDevice,
    busy,
    error = '',
    now = undefined,
    onPair,
    onRevoke,
  }: {
    devices: readonly Device[];
    /** The code on screen, or null. */
    pairing: PairingCode | null;
    /** The companion that just used the code. */
    pairedDevice: Device | null;
    busy: boolean;
    /** A pair or revoke that failed. */
    error?: string;
    now?: Date;
    onPair: () => void;
    onRevoke: (id: string) => void;
  } = $props();

  /** The device whose Revoke is asking; null when no row is. */
  let revoking = $state<string | null>(null);
  /** The line under the name: the status line without the name it starts with. */
  const statusOf = (device: Device): string =>
    deviceStatusLine(device, now).split(' · ').slice(1).join(' · ');
</script>

<Section title={accountPageCopy.devicesLabel} anchor="devices" testid="account-devices" panel>
  <div class="flex flex-col gap-3">
    {#if devices.length > 0}
      <div>
        {#each devices as device (device.id)}
          {#if revoking === device.id}
            <Row title={devicesCopy.revokeQuestion(device.name)} open wrapTrail>
              {#snippet titleContent()}<span data-testid="revoke-question"
                  >{devicesCopy.revokeQuestion(device.name)}</span
                >{/snippet}
              {#snippet trail()}
                <ConfirmButtons
                  confirm={devicesCopy.revoke}
                  cancel={devicesCopy.cancel}
                  {busy}
                  onconfirm={() => {
                    revoking = null;
                    onRevoke(device.id);
                  }}
                  oncancel={() => (revoking = null)}
                  testid="revoke"
                />
              {/snippet}
            </Row>
          {:else}
            <Row title={device.name} sub={statusOf(device)} open wrapTrail testid="account-device">
              {#snippet trail()}
                <button
                  type="button"
                  class={smallButton()}
                  disabled={busy}
                  onclick={() => (revoking = device.id)}
                  data-testid="revoke">{devicesCopy.revoke}</button
                >
              {/snippet}
            </Row>
          {/if}
        {/each}
      </div>
    {/if}
    {#if pairedDevice !== null}
      <p class="text-[15px] leading-6" data-testid="pairing-success">
        {pairingCopy.success(pairedDevice.name)}
      </p>
    {:else if pairing !== null}
      <p class="text-small text-muted" data-testid="pairing-fallback-line">{devicesCopy.codeFallbackLine}</p>
      <Code text={pairing.code} big testid="pairing-code" />
      <p class="text-small text-muted">
        Type this into the companion within {Math.round(pairing.expires_in / 60)} minutes.
      </p>
    {:else}
      {#if devices.length === 0}<p class="text-small text-muted max-w-[80ch]">
          {accountPageCopy.noDevices}
        </p>{/if}
      <div class="flex flex-wrap items-center gap-x-4 gap-y-[10px]">
        {#if companionsOf(devices).length === 0}
          <a
            class={button('secondary')}
            href={accountPageCopy.downloadCompanionHref}
            data-testid="account-download-companion">{accountPageCopy.downloadCompanion}</a
          >
        {/if}
        <button
          type="button"
          class="{LINK} text-[15px]"
          onclick={onPair}
          disabled={busy}
          data-testid="account-get-pairing-code"
          >{pairingEntryLabel(devices.length, accountPageCopy.pairADevice)} ›</button
        >
      </div>
    {/if}
    {#if error !== ''}<p class="text-wipe text-[13px]" role="alert" data-testid="account-devices-error">
        {error}
      </p>{/if}
  </div>
</Section>
