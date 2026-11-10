<!-- web/src/components/sim/SpecGrid.svelte -->
<!-- /sim/specs (mock design/mocks/ws_simtools.py `web-sim-specs`): the damage specs the simulator
     covers on the page's framed panel, one row each, then the healers and tanks it does not, open
     under their own section, so the page reads "these work, these do not" and never interleaves.
     The simulated list always holds one row per damage spec (specSkeletonSlots()), loading,
     failed or answered, so the reserved height never changes: loading draws placeholder rows,
     a failed read takes the first row's place with the reason and Try again, and an answer
     reads "Not yet" for a spec the API said nothing about (mergeSpecRows). sim/specs.astro
     renders this component with no rows as its static shell, so first paint reserves the same
     height. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { nonDpsSpecs } from '../../lib/sim/spec-label';
  import { mergeSpecRows } from '../../lib/sim/spec-state';
  import { specSkeletonSlots } from '../../lib/sim/spec-skeleton';
  import { topRaidSets, type BisBenchmarks } from '../../lib/sim/bis-benchmark';
  import Button from '../ui/Button.svelte';
  import Page from '../ui/Page.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import type { SpecFidelity } from '../../lib/sim/types';
  import SpecCard from './SpecCard.svelte';
  import UnsimulatedSpecCard from './UnsimulatedSpecCard.svelte';

  let {
    rows,
    error,
    onretry,
    benchmarks = {},
  }: {
    rows: SpecFidelity[] | null;
    error: string | null;
    onretry: () => void;
    /** The published sets (lib/sim/bis-benchmark.ts): each damage row's raid BiS DPS at 60. */
    benchmarks?: BisBenchmarks;
  } = $props();

  /** The level whose raid sets the list reads. */
  const TOP_SET_BAND = 60;
  const topSets = $derived(topRaidSets(benchmarks, TOP_SET_BAND));

  const skeletonSlots = specSkeletonSlots();
  /** How long a `#<spec>` arrival keeps its card in view while the list above it settles. */
  const HASH_FOLLOW_MS = 1500;
  // The 8 healer/tank cards carry no fidelity data (they are never in `rows`, task-2-brief.md's
  // second group) so they need no loading state of their own -- derived once from the
  // canonical spec list and rendered immediately, the same way sim/specs.astro's static
  // shell renders this same list before this island ever
  // mounts.
  const unsimulatedSpecs = nonDpsSpecs();

  // "/sim/specs#warrior-fury" from RotationDisclosure's own fidelity-detail link (opened in
  // a new tab, both from the settings bar and from RotationCard -- final whole-branch review
  // C1 stopped the settings bar's own copy of this control from navigating the current tab
  // at all): the browser's own fragment scroll fires once, on the shell's first paint, long
  // before this island fetches and renders a single card -- so the scroll has to happen
  // here, once the row it names actually exists in the DOM, rather than being left to a
  // navigation the browser already finished acting on.
  $effect(() => {
    if (rows === null) return;
    const id = window.location.hash.slice(1);
    if (id === '') return;
    const target = document.getElementById(id);
    if (target === null) return;
    // The cards above settle a frame or two after the rows land (their disclosures mount), so
    // the card is followed while the list still grows, for at most HASH_FOLLOW_MS, and never
    // past the reader's own first scroll.
    const grid = target.closest('[data-testid="specs-grid"]') ?? target;
    const follow = new ResizeObserver(() => target.scrollIntoView());
    const stop = (): void => follow.disconnect();
    target.scrollIntoView();
    follow.observe(grid);
    window.addEventListener('wheel', stop, { once: true, passive: true });
    window.addEventListener('touchstart', stop, { once: true, passive: true });
    const timer = window.setTimeout(stop, HASH_FOLLOW_MS);
    return () => {
      window.clearTimeout(timer);
      stop();
    };
  });
</script>

<Page layout="ledger" panels>
  <Section
    title={simCopy.specsSimulatedHeading}
    meta={topSets.size > 0 ? simCopy.specsTopSetMeta : undefined}
    panel
    frame="framed"
    flush
    testid="specs-simulated"
  >
    <div>
      <div data-testid="specs-grid">
        {#if rows === null}
          {#each skeletonSlots as slot, index (slot)}
            {#if index === 0 && error !== null}
              <div role="alert" data-testid="specs-error">
                <Row title={error}>
                  {#snippet trail()}
                    <Button size="sm" onclick={() => onretry()} testid="specs-retry"
                      >{simCopy.tryAgain}</Button
                    >
                  {/snippet}
                </Row>
              </div>
            {:else}
              <div aria-hidden="true">
                <Row title="">
                  {#snippet titleContent()}<span class="skeleton-block block h-3 w-40"></span>{/snippet}
                  {#snippet subContent()}<span class="skeleton-block mt-2 block h-3 w-64"></span>{/snippet}
                </Row>
              </div>
            {/if}
          {/each}
        {:else}
          {#each mergeSpecRows(rows) as row (row.spec)}
            <SpecCard {row} top={topSets.get(row.spec) ?? null} />
          {/each}
        {/if}
      </div>
    </div>
  </Section>
  <Section title={simCopy.specsUnsimulatedHeading} panel testid="specs-unsimulated">
    <div data-testid="specs-grid-unsimulated">
      {#each unsimulatedSpecs as spec (spec.spec)}
        <UnsimulatedSpecCard {spec} />
      {/each}
    </div>
  </Section>
</Page>
