<!-- web/src/components/sim/SimPreRun.svelte -->
<!-- /sim with a character before its run lands (spec 2026-11-06-sim §4.3, §4.4): not run yet, the
     doll is the page's one Plate (Gear), then the talents; the rail the rotation (no cast counts
     yet) and Your sims. While a run is in flight the Plate holds the ability table's skeleton at
     its final geometry (the heads and 11 rows at 56 px), and the doll steps out of it into an open
     Gear section (ruling S3), so the result draws into the space it already has. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { SimRequest } from '../../lib/sim/types';
  import Page from '../ui/Page.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';
  import SimGearSection from './SimGearSection.svelte';
  import SimRotationSection from './SimRotationSection.svelte';
  import SimTalentsSection from './SimTalentsSection.svelte';

  let {
    request,
    running,
    railTail,
  }: {
    request: SimRequest;
    running: boolean;
    /** The rail's last section (Your sims, signed in). */
    railTail?: Snippet;
  } = $props();

  /** The ability table's rows at the fixture's size: the result lands in the space held. */
  const SKELETON_ROWS = 11;
</script>

<Page layout="split" testid="sim-pre-run">
  {#if running}
    <Section title="Damage by ability" meta="Draws when the run finishes" testid="sim-damage-pending">
      <Plate label="Damage by ability">
        <div aria-busy="true" data-testid="sim-table-skeleton">
          <div class="flex h-10 items-center gap-3 px-4">
            <span class="skeleton-block h-2.5 w-16"></span><span class="ml-auto"></span><span
              class="skeleton-block h-2.5 w-40"
            ></span>
          </div>
          {#each Array.from({ length: SKELETON_ROWS }, (_, index) => index) as index (index)}
            <div class="flex h-14 items-center gap-3 px-4">
              <span class="skeleton-block h-7 w-7 rounded-[6px]"></span><span class="skeleton-block h-3 w-40"
              ></span><span class="ml-auto"></span><span class="skeleton-block hidden h-2 w-[140px] md:block"
              ></span><span class="skeleton-block h-3 w-12"></span><span
                class="skeleton-block hidden h-3 w-40 md:block"
              ></span>
            </div>
          {/each}
        </div>
      </Plate>
    </Section>
  {/if}
  <SimGearSection {request} plated={!running} />
  <SimTalentsSection {request} />
  {#snippet rail()}
    <SimRotationSection {request} />
    {#if railTail}{@render railTail()}{/if}
  {/snippet}
</Page>
