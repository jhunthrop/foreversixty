<!-- web/src/components/setup/CompanionDownload.svelte -->
<!-- The companion's download, the one drawing of it on /setup (lib/setup/downloads.ts,
     lib/setup/release.ts): row 3's action, the head's answer and the top of The companion (owner
     2026-11-08: "a download link at the top and in the companion section"). Windows or Mac: one
     button for the visitor's own build, linking straight to the release's file, its version and size
     under it, and a quiet "Other platforms" disclosure with the rest. A system the page cannot name
     (the server's render, Linux, ChromeOS): every build side by side, none gold, so nothing is a
     wrong guess before hydration. A phone: "runs on your PC or Mac" and Copy link (row 3 draws its
     own from lib/setup/steps.ts). The meta line keeps its height while the manifest is in flight. -->
<script lang="ts">
  import { addonCopy } from '../../lib/addon/copy';
  import { downloadFor, otherDownloads, type CompanionDownload } from '../../lib/setup/downloads';
  import { downloadHref, releaseMeta, type CompanionRelease } from '../../lib/setup/release';
  import type { VisitorPlatform } from '../../lib/setup/visitor-platform';
  import { LINK } from '../../lib/ui/primitives';
  import Button from '../ui/Button.svelte';
  import Disclosure from '../ui/Disclosure.svelte';

  let {
    platform,
    release,
    primary,
    copied = false,
    onCopyLink,
    idPrefix = 'setup',
    end = false,
  }: {
    platform: VisitorPlatform;
    /** The feed's newest release; null until (or unless) the manifest answers. */
    release: CompanionRelease | null;
    /** Row 3 is the page's next step: the visitor's own build is the one gold button. */
    primary: boolean;
    /** A phone's Copy link has just copied: its button says so. */
    copied?: boolean;
    /** A phone's Copy link was pressed. */
    onCopyLink?: () => void;
    /** Each placement's testids start with it: row 3 `setup`, the head `setup-head`. */
    idPrefix?: string;
    /** Right-aligned from 900, as the head's answer is. */
    end?: boolean;
  } = $props();

  const own = $derived(downloadFor(platform));
  const others = $derived(otherDownloads(platform));
  const label = (download: CompanionDownload): string => addonCopy.setupCompanionDownload(download.platform);
</script>

{#if platform === 'mobile'}
  <div class="flex flex-col items-start gap-2" class:dl-end={end} data-testid="{idPrefix}-copy">
    <span class="text-small text-muted">{addonCopy.setupCompanionMobile}</span>
    <Button onclick={() => onCopyLink?.()} testid="{idPrefix}-action-copylink"
      >{copied ? addonCopy.setupLinkCopied : addonCopy.setupCopyLink}</Button
    >
  </div>
{:else if own !== null}
  {@const meta = releaseMeta(own.asset, release)}
  <div class="flex flex-col gap-1" class:dl-end={end} data-testid="{idPrefix}-download">
    <div class="dl-row flex flex-wrap items-center gap-x-4 gap-y-2">
      <Button
        tone={primary ? 'primary' : 'secondary'}
        href={downloadHref(own.asset, release)}
        download={own.saveAs}
        rel="noopener"
        testid="{idPrefix}-action-download">{label(own)}</Button
      >
      {#if others.length > 0}
        <Disclosure summary={addonCopy.setupOtherPlatforms} link testid="{idPrefix}-other-platforms">
          <ul class="flex flex-col gap-1">
            {#each others as other (other.asset)}
              {@const otherMeta = releaseMeta(other.asset, release)}
              <li class="text-small">
                <a
                  class={LINK}
                  href={downloadHref(other.asset, release)}
                  download={other.saveAs}
                  rel="noopener"
                  data-testid="{idPrefix}-download-{other.os}">{label(other)} ›</a
                >{#if otherMeta !== null}<span class="text-muted"> {otherMeta}</span>{/if}
              </li>
            {/each}
          </ul>
        </Disclosure>
      {/if}
    </div>
    <span class="text-small text-muted min-h-5" data-testid="{idPrefix}-download-meta">{meta ?? ''}</span>
  </div>
{:else}
  <div
    class="flex flex-col gap-1"
    class:dl-end={end}
    role="group"
    aria-label={addonCopy.setupDownloadCompanion}
    data-testid="{idPrefix}-download"
  >
    <div class="dl-row flex flex-wrap items-center gap-x-4 gap-y-2">
      {#each others as download (download.asset)}
        <Button
          href={downloadHref(download.asset, release)}
          download={download.saveAs}
          rel="noopener"
          testid="{idPrefix}-download-{download.os}">{label(download)}</Button
        >
      {/each}
    </div>
    <span class="text-small text-muted min-h-5" data-testid="{idPrefix}-download-meta"
      >{release === null ? '' : addonCopy.setupCompanionVersion(release.version)}</span
    >
  </div>
{/if}

<style>
  @media (min-width: 900px) {
    .dl-end {
      align-items: flex-end;
    }
    .dl-end .dl-row {
      justify-content: flex-end;
    }
  }
</style>
