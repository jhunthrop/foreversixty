<!-- web/src/components/report/ThisPull.svelte -->
<!-- The rail's This pull panel (directions b1): only what the band does not say (when it started,
     a boss pull's phases), then the pull's two actions, Sim this fight and Copy link (a link to
     exactly this view); below 768 those two live in the band's Report actions menu. Over the whole
     night it is `The night` (round 5 b1-report-night): the pulls, kills, wipes, time in combat and
     deaths, and Copy link. -->
<script lang="ts">
  import { reportCopy } from '../../lib/report/copy';
  import { simCopy } from '../../lib/sim/copy';
  import { button as buttonClass } from '../../lib/ui/button';
  import { CopyState } from '../../lib/ui/copy-state.svelte';
  import type { Fact } from '../../lib/ui/facts';
  import Facts from '../ui/Facts.svelte';
  import Section from '../ui/Section.svelte';

  let {
    facts,
    simHref,
    title = 'This pull',
  }: {
    facts: readonly Fact[];
    simHref: string | null;
    /** The panel's title: `The night` over a combined view. */
    title?: string;
  } = $props();

  const copier = new CopyState();
  $effect(() => () => copier.dispose());
</script>

<Section {title} panel testid="report-this-pull">
  <Facts items={facts} />
  <!-- Below 768 the band's Report actions menu holds these two (ReportActions). -->
  <div class="flex flex-wrap gap-2 pt-4 max-md:hidden">
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
      onclick={() => void copier.copy(window.location.href)}>{copier.label(reportCopy.copyLink)}</button
    >
  </div>
</Section>
