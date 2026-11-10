<!-- web/src/components/sim/tools/WeightsShell.svelte -->
<!-- /sim/weights' server render (sim/weights.astro), the island's no-character state drawn
     without a store: the source switcher, the run line's band with its Run held off, the caution,
     the default spec's published weights and the stats to weigh. The island mounts the same parts
     over it (ToolsView: SourceSwitcher, then StatWeights with `defaultWeights`), so the table a
     search engine reads is the one a player sees, figure for figure. -->
<script lang="ts">
  import { bulkCopy } from '../../../lib/sim/copy';
  import { simPageCopy } from '../../../lib/sim/page-copy';
  import type { PublishedWeights } from '../../../lib/sim/published-weights';
  import { pickableStatsFor, weightsEngineIterations, weightsIterationsFor } from '../../../lib/sim/weights';
  import Button from '../../ui/Button.svelte';
  import ControlBar from '../../ui/ControlBar.svelte';
  import Notice from '../../ui/Notice.svelte';
  import Page from '../../ui/Page.svelte';
  import SourceSwitcher from '../SourceSwitcher.svelte';
  import WeightsPicks from './WeightsPicks.svelte';
  import WeightsPublishedSection from './WeightsPublishedSection.svelte';

  let { published }: { published: PublishedWeights } = $props();

  const nothing = (): void => {};
  /** BulkRunBar's cost line at the store's opening state: Fast, in the browser, nothing ticked. */
  const openingCost = weightsEngineIterations(weightsIterationsFor('fast', 'browser'), 0);
</script>

<div data-testid="sim-tools-view">
  <!-- Held busy (every field and Load disabled) until the island mounts: a paste typed into the
       server's form would be dropped when the island replaces it. -->
  <SourceSwitcher
    busy={true}
    message={null}
    signedIn={false}
    hasCharacters={false}
    onaddon={nothing}
    onbuild={nothing}
    onfight={nothing}
    onsignin={nothing}
  />
  <Page layout="split" panels testid="sim-stat-weights">
    {#snippet bar()}
      <ControlBar groups={[]} label={simPageCopy.bar.label} labelsFrom={1920}>
        {#snippet action()}<Button tone="primary" disabled>{bulkCopy.runBulk}</Button>{/snippet}
      </ControlBar>
      <div class="flex flex-col gap-1 px-[var(--page-inset)] pt-3">
        <p class="m-0 text-[14px] leading-5">
          <span class="text-muted">{bulkCopy.weightsCostNote(openingCost)}</span>
        </p>
      </div>
    {/snippet}
    <Notice
      text={bulkCopy.weightsWarning}
      tone="neutral"
      action={{ label: bulkCopy.weightsWarningLink, href: '/sim/gear' }}
      testid="sim-weights-warning"
    />
    <WeightsPublishedSection {published} isDefault />
    {#snippet rail()}
      <WeightsPicks pickable={pickableStatsFor(undefined)} picked={[]} reference="" />
    {/snippet}
  </Page>
</div>
