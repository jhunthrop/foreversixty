<!-- web/src/components/setup/SetupCompanion.svelte -->
<!-- /setup's Section "The companion" (setup spec 4, rulings U7, U8): the download right under the
     heading in every state (owner 2026-11-08; a phone: "runs on your PC or Mac" and Copy link), what
     it is, the signing and first-run line of the visitor's own build (every build's while the system
     is not known; none on a phone), the three ranked how-to Rows (`/combatlog` as Code), the
     Status capture with its Example tag and, signed in, the pairing block under a hairline (each
     paired device's status line first). -->
<script lang="ts">
  import type { Device } from '../../lib/account/api';
  import { addonCopy } from '../../lib/addon/copy';
  import { deviceStatusLine } from '../../lib/reports/device-status';
  import { COMPANION_DOWNLOADS, downloadFor } from '../../lib/setup/downloads';
  import type { CompanionRelease } from '../../lib/setup/release';
  import type { VisitorPlatform } from '../../lib/setup/visitor-platform';
  import PairingBlock from '../companion/PairingBlock.svelte';
  import Capture from '../ui/Capture.svelte';
  import Code from '../ui/Code.svelte';
  import Divided from '../ui/Divided.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import CompanionDownload from './CompanionDownload.svelte';

  const COMBATLOG = '/combatlog';

  let {
    platform,
    release = null,
    copied = false,
    onCopyLink,
    signedIn,
    devices,
    now,
    onCode,
    onPaired,
  }: {
    platform: VisitorPlatform;
    release?: CompanionRelease | null;
    /** A phone's Copy link has just copied. */
    copied?: boolean;
    onCopyLink?: () => void;
    signedIn: boolean;
    devices: readonly Device[];
    now: Date;
    onCode?: (code: string | null) => void;
    onPaired?: (device: Device) => void;
  } = $props();

  /** The visitor's own build; null on a phone or while the system is not known. */
  const download = $derived(downloadFor(platform));
  /** The first-run lines: the visitor's own build's, every build's when the system is not known. */
  const firstRuns = $derived(
    download !== null ? [download] : platform === 'unknown' ? COMPANION_DOWNLOADS : [],
  );
  /** `setupCompanionBody` at its sentences: what it is, then the two how-to steps after the pairing one. */
  const [lead, ...rest] = addonCopy.setupCompanionBody.split(/(?<=\.)\s+(?=[A-Z])/);
  const howTo = [addonCopy.setupCompanionPair, ...rest];
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
  panel
>
  <div class="setup-body">
    <CompanionDownload
      {platform}
      {release}
      primary={false}
      {copied}
      {onCopyLink}
      idPrefix="setup-companion"
    />
    <p data-testid="setup-companion-lead">{lead}</p>
    {#each firstRuns as build (build.asset)}
      <p class="text-small text-muted" data-testid="setup-companion-first-run">
        {firstRuns.length > 1 ? `${build.platform}: ${build.firstRun}` : build.firstRun}
      </p>
    {/each}
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
      narrow={capture.narrow}
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
