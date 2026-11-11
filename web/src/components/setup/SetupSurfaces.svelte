<!-- web/src/components/setup/SetupSurfaces.svelte -->
<!-- /setup's Section "What it does in the game" (setup spec 4, rulings U5): the eleven in-game
     surfaces as open Rows in a RowGrid (two columns from 768), each with its
     sentence and its client captures in the Row's foot (240 px high, lazy, each a link to its
     file). On a phone the grid folds behind one button (`surfacesShow`, aria-expanded); open from 768. -->
<script lang="ts">
  import { addonCopy } from '../../lib/addon/copy';
  import { button } from '../../lib/ui/button';
  import Capture from '../ui/Capture.svelte';
  import Row from '../ui/Row.svelte';
  import RowGrid from '../ui/RowGrid.svelte';
  import Section from '../ui/Section.svelte';

  const SURFACE_CAPTURE_MAX_PX = 240;
  const GRID_ID = 'setup-surfaces-grid';

  /** The phone's fold: closed until the button opens it (from 768 the grid always shows). */
  let shown = $state(false);
</script>

<Section title={addonCopy.surfacesTitle} testid="setup-surfaces" panel>
  {#if !shown}
    <button
      type="button"
      class="{button('secondary')} setup-surf-fold"
      aria-expanded="false"
      aria-controls={GRID_ID}
      data-testid="setup-surfaces-show"
      onclick={() => (shown = true)}>{addonCopy.surfacesShow}</button
    >
  {/if}
  <div class="setup-surf" id={GRID_ID} data-shown={shown || undefined}>
    <RowGrid cols={2}>
      {#each addonCopy.surfaces as surface (surface.section)}
        <Row title={surface.title} sub={surface.body} testid="setup-surface-{surface.section}">
          {#snippet foot()}
            {#each surface.images as image (image.src)}
              <Capture
                src={image.src}
                width={image.width}
                height={image.height}
                alt={image.alt}
                maxHeight={SURFACE_CAPTURE_MAX_PX}
              />
            {/each}
          {/snippet}
        </Row>
      {/each}
    </RowGrid>
  </div>
</Section>

<style>
  .setup-surf :global(.ui-row) {
    align-items: start;
    align-content: start;
  }
  /* The captures sit 12 px under the sentence. */
  .setup-surf :global(.ui-row > .col-span-full) {
    padding-top: 12px;
  }
  .setup-surf-fold {
    display: none;
  }
  @media (max-width: 767px) {
    .setup-surf-fold {
      display: inline-flex;
    }
    .setup-surf:not([data-shown]) {
      display: none;
    }
  }
</style>
