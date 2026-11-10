<!-- web/src/components/setup/SetupRail.svelte -->
<!-- /setup's rail (setup spec 4, rulings U9, U11): Try an export open (the Korea way in, and a Mac's while there is no Mac build),
     the 14 /fs commands as a key-binding list (what it does at the left, the command at the right
     as Code), the companion handshake, and the two flows (out of the game, into it). -->
<script lang="ts">
  import { addonCopy } from '../../lib/addon/copy';
  import type { Fact } from '../../lib/ui/facts';
  import AddonPasteBox from '../AddonPasteBox.svelte';
  import Facts from '../ui/Facts.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let { signedIn }: { signedIn: boolean | null } = $props();

  const commands: Fact[] = addonCopy.commands.map(({ command, description }) => ({
    id: command,
    label: description,
    code: command,
  }));
  const flows = [
    { title: addonCopy.flowOutTitle, body: addonCopy.flowOutBody },
    { title: addonCopy.flowInTitle, body: addonCopy.flowInBody },
  ];
</script>

<Section title={addonCopy.setupTryExport} anchor="paste" testid="setup-try-export" panel>
  <div class="setup-body">
    <p class="text-small text-muted" data-testid="setup-export-what">{addonCopy.setupTryExportBody}</p>
    <AddonPasteBox {signedIn} />
  </div>
</Section>

<Section title={addonCopy.commandsTitle} testid="setup-commands" panel>
  <div class="setup-commands"><Facts items={commands} /></div>
</Section>

<Section title={addonCopy.handshakeTitle} testid="setup-handshake" panel>
  <div class="setup-body">
    {#each addonCopy.handshake as sentence (sentence)}<p class="text-small text-muted">{sentence}</p>{/each}
  </div>
</Section>

<Section title={addonCopy.setupHowItWorks} testid="setup-flows" panel>
  <div>
    {#each flows as flow (flow.title)}<Row title={flow.title} sub={flow.body} open />{/each}
  </div>
</Section>

<style>
  .setup-commands :global(dt) {
    color: var(--color-text);
    padding-right: 12px;
  }
  /* A command never breaks and keeps its width: the description takes the rest and wraps. */
  .setup-commands :global(dl) {
    grid-template-columns: minmax(0, 1fr) auto;
  }
  .setup-commands :global(dd) {
    white-space: nowrap;
  }
</style>
