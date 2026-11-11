<!-- web/src/components/sim/tools/StatWeights.svelte -->
<!-- /sim/weights (mock design/mocks/ws_simtools.py): the run line, the stats to weigh as one row of
     ticks (the explainer notice and its Top Gear link were cut, h-audit 2026-11-09), then the
     scale factors on the page's framed panel.

     The reference stat always leads `store.stats`: the engine normalises its weights to it,
     so the request must weigh it. It is seeded from the spec's own
     `reference_stat` (weights.ts's `referenceFor`/`defaultStatsFor`, run by the store) and
     is never itself untickable here -- the checkbox for it is disabled, not merely
     defaulted, so `WeightsSpec.Reference` (required, contract 10.8) can never go empty
     through this control. `store.stats` cannot reach [] this way; only `setStats([])`
     directly (the store's own pre-send refusal, `bulkCopy.weightsNeedStats`) can.

     The results read as the published weights do (owner, 2026-09-30, round 11): the
     SimulationCraft scale-factor table (WeightsScaleTable.svelte over weights.ts's
     `toolScaleView`), top stat = 1.00 with DPS per point and the error, haste per 1% in the
     caption, no Pawn string. A row the engine flagged `insignificant` greys there. The
     picker above them offers only
     `store.weightStats` when the spec sent one (sub-item 4): a stat this 1.60 sim does not
     model for this spec is never on screen to tick in the first place.

     Before a run, the main column holds the weights the ranker published for the character's
     own band (lib/sim/published-weights.ts), read through the same table, linking to that set.
     With no character loaded it holds `defaultWeights`, the default spec's level 60 figures the
     page computed at build time and server-rendered (WeightsShell.svelte). -->
<script lang="ts">
  import activeBuild from '../../../data/active-build.json';
  import type { Me } from '../../../lib/account/api';
  import { fetchBisFile } from '../../../lib/bis/hover';
  import {
    DEFAULT_WEIGHTS_SPEC,
    publishedWeights,
    type PublishedWeights,
  } from '../../../lib/sim/published-weights';
  import { requestPresetFor } from '../../../lib/sim/settings';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { bulkCopy } from '../../../lib/sim/copy';
  import { specLabel } from '../../../lib/sim/spec-label';
  import { pickableStatsFor, pickableStatsForSpec, toolScaleView } from '../../../lib/sim/weights';
  import type { Snippet } from 'svelte';
  import Page from '../../ui/Page.svelte';
  import Section from '../../ui/Section.svelte';
  import WeightsScaleTable from '../WeightsScaleTable.svelte';
  import BulkRunBar from './BulkRunBar.svelte';
  import SaveSimForm from './SaveSimForm.svelte';
  import WeightsPicks from './WeightsPicks.svelte';
  import WeightsPublishedSection from './WeightsPublishedSection.svelte';

  let {
    store,
    lead,
    defaultWeights = null,
  }: {
    store: BulkStore;
    me: Me | null;
    lead?: Snippet;
    /** The default spec's published figures, shown while no character is loaded. */
    defaultWeights?: PublishedWeights | null;
  } = $props();

  const view = $derived(toolScaleView(store.weights, store.character?.spec ?? '', store.referenceStat));
  /** The picker's own list: `store.weightStats` when the spec sent one, the full pinned
   *  vocabulary otherwise (sub-item 4; `weights.ts`'s own fallback rule). With no character it is
   *  the default table's spec's, as the server render draws it (WeightsShell). */
  const pickable = $derived(
    store.character === null
      ? pickableStatsForSpec(defaultWeights?.spec ?? DEFAULT_WEIGHTS_SPEC)
      : pickableStatsFor(store.weightStats),
  );

  /**
   * `store.save` (bulk-store.svelte.ts) has no fallback title of its own, the same as every
   * other tool page's save form -- see ComboResults.svelte's identical comment. A weights
   * run has no winner to headline, so this names the tool and the spec instead, the same
   * two-part title SavedSim.svelte gives a saved weights page's own <h1>.
   */
  const saveTitleFor = $derived(`${bulkCopy.weightsTitle} · ${specLabel(store.character?.spec ?? '')}`);

  let characterPublished = $state<PublishedWeights | null>(null);
  const published = $derived(store.character === null ? defaultWeights : characterPublished);
  $effect(() => {
    const character = store.character;
    const preset = requestPresetFor(store.settings);
    if (character === null) {
      characterPublished = null;
      return;
    }
    let live = true;
    fetchBisFile(activeBuild.build, character.spec)
      .then((file) => {
        if (!live) return;
        characterPublished =
          file === null
            ? null
            : publishedWeights(file, {
                spec: character.spec,
                race: character.race_slug,
                level: character.level,
                preset,
              });
      })
      .catch((error: unknown) => console.error('sim weights: the published weights did not load', error));
    return () => {
      live = false;
    };
  });

  /** A stopped run has nothing finished worth naming and saving. */
  const canSave = $derived(store.result?.aborted !== true);

  function togglePick(id: string): void {
    if (id === store.referenceStat) return;
    const next = store.stats.includes(id) ? store.stats.filter((stat) => stat !== id) : [...store.stats, id];
    // The reference stat always leads (see the file comment above).
    const reference = store.referenceStat;
    store.setStats(reference === '' ? next : [reference, ...next.filter((stat) => stat !== reference)]);
  }
</script>

<Page layout="ledger" panels testid="sim-stat-weights">
  {#snippet bar()}<BulkRunBar {store} />{/snippet}
  {#if lead}{@render lead()}{/if}
  <WeightsPicks {pickable} picked={store.stats} reference={store.referenceStat} ontoggle={togglePick} />
  {#if store.weights.length > 0}
    <Section
      panel
      frame="plain"
      title={bulkCopy.weightsResultsTitle}
      meta={specLabel(store.character?.spec ?? '')}
      testid="sim-weights"
    >
      <!-- Contract 10.9's error-is-a-lower-bound caveat, read right before the numbers it caveats. -->
      <p class="text-small text-muted m-0 pb-3" data-testid="sim-weights-error-caveat">
        {bulkCopy.weightsErrorCaveat(store.weights.map((row) => row.stat))}
      </p>
      <WeightsScaleTable {view} />
      <!-- Contract 10.6: saved sims of every kind are public at /sim/<id>, saved through the one
           SaveSimForm Top Gear and talent compare use. -->
      <div class="pt-4">
        <SaveSimForm onsave={(title) => store.save(title)} titleFor={saveTitleFor} {canSave} />
      </div>
    </Section>
  {:else if published !== null}
    <WeightsPublishedSection {published} isDefault={store.character === null} />
  {/if}
</Page>
