<!-- web/src/components/setup/SetupSurfaces.svelte -->
<!-- /setup's Section "What it does in the game" (setup spec 4, rulings U5): the eleven in-game
     surfaces as open Rows in a RowGrid (three columns from 1024, two from 768), each with its
     sentence and its client captures in the Row's foot (240 px high, lazy, each a link to its
     file). On a phone the grid folds behind one summary row (`surfacesShow`); open from 768. -->
<script lang="ts">
  import { addonCopy } from '../../lib/addon/copy';
  import Capture from '../ui/Capture.svelte';
  import Row from '../ui/Row.svelte';
  import RowGrid from '../ui/RowGrid.svelte';
  import Section from '../ui/Section.svelte';

  const SURFACE_CAPTURE_MAX_PX = 240;
</script>

<Section title={addonCopy.surfacesTitle} testid="setup-surfaces">
  <details class="setup-surf-fold">
    <summary class="label text-gold">{addonCopy.surfacesShow}</summary>
  </details>
  <div class="setup-surf">
    <RowGrid cols={3}>
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
  .setup-surf-fold summary {
    display: flex;
    min-height: 44px;
    align-items: center;
    cursor: pointer;
  }
  /* From 768 the fold is gone and the grid sits right under the Section's hairline. */
  @media (min-width: 768px) {
    .setup-surf {
      margin-top: -12px;
    }
  }
  @media (max-width: 767px) {
    .setup-surf-fold {
      display: block;
    }
    .setup-surf-fold:not([open]) + .setup-surf {
      display: none;
    }
  }
</style>
