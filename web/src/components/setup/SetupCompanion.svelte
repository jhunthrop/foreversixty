<!-- web/src/components/setup/SetupCompanion.svelte -->
<!-- /setup's Section "The companion" (setup spec 4, rulings U4, U7, U8): what it is, the visitor's
     own build's platforms line and signing line (a Mac visitor without a Mac build reads "Windows
     only" once, on row 3, and neither line here), a quiet
     download link once row 3 is ticked, the three ranked how-to Rows (`/combatlog` as Code), the
     Status capture with its Example tag and, signed in, the pairing block under a hairline (each
     paired device's status line first). -->
<script lang="ts">
  import type { Device } from '../../lib/account/api';
  import { addonCopy } from '../../lib/addon/copy';
  import { deviceStatusLine } from '../../lib/reports/device-status';
  import { COMPANION_PLATFORMS, downloadFor } from '../../lib/setup/downloads';
  import { companionPlatformsLine } from '../../lib/setup/progress';
  import { stepHref } from '../../lib/setup/links';
  import type { VisitorPlatform } from '../../lib/setup/visitor-platform';
  import { LINK } from '../../lib/ui/primitives';
  import PairingBlock from '../companion/PairingBlock.svelte';
  import Capture from '../ui/Capture.svelte';
  import Code from '../ui/Code.svelte';
  import Divided from '../ui/Divided.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  const COMBATLOG = '/combatlog';

  let {
    platform,
    companionDone,
    signedIn,
    devices,
    now,
    onCode,
    onPaired,
  }: {
    platform: VisitorPlatform;
    companionDone: boolean;
    signedIn: boolean;
    devices: readonly Device[];
    now: Date;
    onCode?: (code: string | null) => void;
    onPaired?: (device: Device) => void;
  } = $props();

  /** The visitor's own build; null on a Mac while the feed carries no Mac build. */
  const download = $derived(downloadFor(platform));
  /** The quiet link for a second computer: the visitor's own build, else the Windows one. */
  const another = $derived(download ?? downloadFor('windows'));
  /** `setupCompanionBody` at its sentences: what it is, then the two how-to steps after the pairing one. */
  const [lead, ...rest] = addonCopy.setupCompanionBody.split(/(?<=\.)\s+(?=[A-Z])/);
  const howTo = [addonCopy.setupCompanionPair, ...rest];
  const platformsLine = companionPlatformsLine(COMPANION_PLATFORMS);
  const capture = addonCopy.setupCompanionCapture;
</script>

{#snippet howToTitle(text: string)}
  {@const at = text.indexOf(COMBATLOG)}
  {#if at < 0}{text}{:else}{text.slice(0, at)}<Code text={COMBATLOG} />{text.slice(
      at + COMBATLOG.length,
    )}{/if}
{/snippet}

<Section
  title={addonCopy.setupStep3Title}
  meta={addonCopy.setupCompanionNote}
  anchor="companion"
  testid="setup-companion"
>
  <div class="setup-body">
    <p data-testid="setup-companion-lead">{lead}</p>
    {#if download !== null}
      {#if platformsLine !== null}<p class="text-small text-muted" data-testid="setup-companion-platforms">
          {platformsLine}
        </p>{/if}
      <p class="text-small text-muted" data-testid="setup-companion-first-run">{download.firstRun}</p>
    {/if}
    {#if companionDone && another !== null}
      <div class="setup-acts">
        <a
          class={LINK}
          href={stepHref('download', another.asset) ?? '#'}
          rel="noopener"
          data-testid="setup-download-another">{addonCopy.setupDownloadAnother} ›</a
        >
      </div>
    {/if}
    <div data-testid="setup-companion-howto">
      {#each howTo as text, index (text)}
        <Row
          title={text}
          rank={String(index + 1)}
          open
          testid={index === 0 ? 'setup-companion-pair' : undefined}
        >
          {#snippet titleContent()}{@render howToTitle(text)}{/snippet}
        </Row>
      {/each}
    </div>
    <Capture
      src={capture.src}
      width={capture.width}
      height={capture.height}
      alt={capture.alt}
      tag={capture.exampleLabel}
      eager
      testid="setup-companion-capture"
    />
    {#if signedIn}
      <Divided gap={8} testid="setup-pairing">
        {#each devices as device (device.id)}
          <p class="text-small text-text" data-testid="setup-device-status">
            {deviceStatusLine(device, now)}
          </p>
        {/each}
        <PairingBlock {devices} {onCode} {onPaired} />
      </Divided>
    {/if}
  </div>
</Section>

<style>
  /* The ranked how-to lines are one-line rows: 56 px, as the mock's open Rows. */
  [data-testid='setup-companion-howto'] :global(.ui-row) {
    min-height: 56px;
  }
  /* The pairing block's lines sit 8 px apart under the hairline, as the device lines above them. */
  :global([data-testid='setup-pairing'] > [data-testid='logs-companion-pair']) {
    gap: 8px;
  }
</style>
