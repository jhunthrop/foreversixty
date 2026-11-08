<!-- web/src/components/setup/CompanionDownload.svelte -->
<!-- Row 3's download (lib/setup/downloads.ts, lib/setup/release.ts). Windows or Mac: one button for
     the visitor's own build, linking straight to the release's file, its version and size under it,
     and a quiet "Other platforms" disclosure with the rest. A system the page cannot name (the
     server's render, Linux, ChromeOS): every build side by side, none gold, so nothing is a wrong
     guess before hydration. The meta line keeps its height while the manifest is in flight. -->
<script lang="ts">
  import { addonCopy } from '../../lib/addon/copy';
  import { downloadFor, otherDownloads, type CompanionDownload } from '../../lib/setup/downloads';
  import { downloadHref, releaseMeta, type CompanionRelease } from '../../lib/setup/release';
  import type { VisitorPlatform } from '../../lib/setup/visitor-platform';
  import { LINK } from '../../lib/ui/primitives';
  import Button from '../ui/Button.svelte';

  let {
    platform,
    release,
    primary,
  }: {
    platform: VisitorPlatform;
    /** The feed's newest release; null until (or unless) the manifest answers. */
    release: CompanionRelease | null;
    /** Row 3 is the page's next step: the visitor's own build is the one gold button. */
    primary: boolean;
  } = $props();

  const own = $derived(downloadFor(platform));
  const others = $derived(otherDownloads(platform));
  const label = (download: CompanionDownload): string => addonCopy.setupCompanionDownload(download.platform);
</script>

{#if own !== null}
  {@const meta = releaseMeta(own.asset, release)}
  <div class="flex flex-col gap-1" data-testid="setup-download">
    <div class="flex flex-wrap items-center gap-x-4 gap-y-2">
      <Button
        tone={primary ? 'primary' : 'secondary'}
        href={downloadHref(own.asset, release)}
        rel="noopener"
        testid="setup-action-download">{label(own)}</Button
      >
      {#if others.length > 0}
        <details class="setup-other" data-testid="setup-other-platforms">
          <summary class="{LINK} inline-flex min-h-11 cursor-pointer items-center md:min-h-0"
            >{addonCopy.setupOtherPlatforms}</summary
          >
          <ul class="mt-2 flex flex-col gap-1">
            {#each others as other (other.asset)}
              {@const otherMeta = releaseMeta(other.asset, release)}
              <li class="text-small">
                <a
                  class={LINK}
                  href={downloadHref(other.asset, release)}
                  rel="noopener"
                  data-testid="setup-download-{other.os}">{label(other)} ›</a
                >{#if otherMeta !== null}<span class="text-muted"> {otherMeta}</span>{/if}
              </li>
            {/each}
          </ul>
        </details>
      {/if}
    </div>
    <span class="text-small text-muted min-h-5" data-testid="setup-download-meta">{meta ?? ''}</span>
  </div>
{:else}
  <div
    class="flex flex-col gap-1"
    role="group"
    aria-label={addonCopy.setupDownloadCompanion}
    data-testid="setup-download"
  >
    <div class="flex flex-wrap items-center gap-3">
      {#each others as download (download.asset)}
        <Button
          href={downloadHref(download.asset, release)}
          rel="noopener"
          testid="setup-download-{download.os}">{label(download)}</Button
        >
      {/each}
    </div>
    <span class="text-small text-muted min-h-5" data-testid="setup-download-meta"
      >{release === null ? '' : addonCopy.setupCompanionVersion(release.version)}</span
    >
  </div>
{/if}
