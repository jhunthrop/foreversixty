<!-- web/src/components/report/ThisPull.svelte -->
<!-- The rail's This pull panel (directions b1): only what the band does not say (when it started,
     the scope's count, its deaths; a boss pull's phases), then the pull's two actions, Sim this
     fight and Copy link (a link to exactly this view). -->
<script lang="ts">
  import { reportCopy } from '../../lib/report/copy';
  import { simCopy } from '../../lib/sim/copy';
  import { button as buttonClass } from '../../lib/ui/button';
  import type { Fact } from '../../lib/ui/facts';
  import Facts from '../ui/Facts.svelte';
  import Section from '../ui/Section.svelte';

  let {
    facts,
    simHref,
    note = '',
  }: {
    facts: readonly Fact[];
    simHref: string | null;
    /** One muted line under the facts (no rating can exist for this pull). */
    note?: string;
  } = $props();

  const COPIED_FOR_MS = 2000;
  let copied = $state('');

  async function copyLink(): Promise<void> {
    try {
      await navigator.clipboard.writeText(window.location.href);
      copied = reportCopy.copied;
    } catch {
      copied = reportCopy.copyFailed;
    }
    setTimeout(() => (copied = ''), COPIED_FOR_MS);
  }
</script>

<Section title="This pull" panel testid="report-this-pull">
  <Facts items={facts} />
  {#if note !== ''}<p class="text-small text-muted pt-3" data-testid="report-this-pull-note">{note}</p>{/if}
  <div class="flex flex-wrap gap-2 pt-4">
    {#if simHref !== null}
      <a class={buttonClass('secondary', 'sm')} href={simHref} data-testid="report-sim-fight"
        >{simCopy.simThisFight}</a
      >
    {/if}
    <button
      type="button"
      class={buttonClass('ghost', 'sm')}
      title="Copy a link to exactly this view"
      data-testid="copy-link"
      onclick={() => void copyLink()}>{copied === '' ? reportCopy.copyLink : copied}</button
    >
  </div>
</Section>
