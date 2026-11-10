<!-- web/src/components/setup/SetupAddon.svelte -->
<!-- /setup's Section "The addon" (setup spec 4, rulings U4): the data addon's line, CurseForge and
     Wago as quiet links once the addon step is ticked (its buttons have left the plate; a second PC
     still installs), then under a hairline the hand install: what the two zips hold, where they go,
     both downloads through /dl/ and the current data build. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { addonCopy } from '../../lib/addon/copy';
  import { ADDON_DOWNLOAD_PATH, DATA_DOWNLOAD_PATH } from '../../lib/downloads/paths';
  import { CURSEFORGE_URL, WAGO_URL } from '../../lib/setup/links';
  import { LINK } from '../../lib/ui/primitives';
  import Divided from '../ui/Divided.svelte';
  import Section from '../ui/Section.svelte';

  let { addonDone }: { addonDone: boolean } = $props();

  const installLinks = [
    { label: addonCopy.installCurseForge, href: CURSEFORGE_URL, testid: 'setup-curseforge-again' },
    { label: addonCopy.installWago, href: WAGO_URL, testid: 'setup-wago' },
  ];
  const handLinks = [
    { label: addonCopy.installZip, href: ADDON_DOWNLOAD_PATH, testid: 'setup-zip' },
    { label: addonCopy.installDataRelease, href: DATA_DOWNLOAD_PATH, testid: 'setup-data-release' },
  ];
</script>

<Section title={addonCopy.setupAddonTitle} anchor="addon" testid="setup-addon" panel>
  <div class="setup-body">
    <p class="text-small text-muted" data-testid="setup-data-addon">{addonCopy.installDataHint}</p>
    {#if addonDone}
      <div class="setup-acts gap-5">
        {#each installLinks as link (link.href)}
          <a class={LINK} href={link.href} rel="noopener" data-testid={link.testid}>{link.label} ›</a>
        {/each}
      </div>
    {/if}
    <Divided testid="setup-hand-install">
      <span class="text-label text-muted font-bold">{addonCopy.setupHandInstall}</span>
      <p class="text-small text-muted">{addonCopy.setupHandInstallBody}</p>
      <div class="setup-acts gap-5">
        {#each handLinks as link (link.href)}
          <a class={LINK} href={link.href} rel="noopener" data-testid={link.testid}>{link.label} ›</a>
        {/each}
      </div>
      <p class="text-small text-muted">{addonCopy.currentDataBuild(activeBuild.build)}</p>
    </Divided>
  </div>
</Section>
