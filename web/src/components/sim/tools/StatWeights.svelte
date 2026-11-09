<!-- web/src/components/sim/tools/StatWeights.svelte -->
<!-- /sim/weights (mock design/mocks/ws_simtools.py): the run line, the caution as the page's
     standing line, the scale factors on the page's framed panel once a run finishes, and the stats to
     weigh in the rail. Design 7: the caution opens the page, in our words, and links to Top Gear -- it is not
     a footnote, because the whole reason this page exists is that addons ask for a number
     the number itself cannot justify.

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
     own band (lib/sim/published-weights.ts), read through the same table, linking to that set. -->
<script lang="ts">
  import activeBuild from '../../../data/active-build.json';
  import type { Me } from '../../../lib/account/api';
  import { fetchBisFile } from '../../../lib/bis/hover';
  import { publishedWeights, type PublishedWeights } from '../../../lib/sim/published-weights';
  import { requestPresetFor } from '../../../lib/sim/settings';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { bulkCopy, WEIGHTS_STATS_FROM_ENGINE } from '../../../lib/sim/copy';
  import { specLabel } from '../../../lib/sim/spec-label';
  import { hasWeightStats, pickableStatsFor, statLabel, toolScaleView } from '../../../lib/sim/weights';
  import type { Snippet } from 'svelte';
  import Notice from '../../ui/Notice.svelte';
  import Page from '../../ui/Page.svelte';
  import Section from '../../ui/Section.svelte';
  import WeightsScaleTable from '../WeightsScaleTable.svelte';
  import BulkRunBar from './BulkRunBar.svelte';
  import SaveSimForm from './SaveSimForm.svelte';

  let { store, lead }: { store: BulkStore; me: Me | null; lead?: Snippet } = $props();

  const view = $derived(toolScaleView(store.weights, store.character?.spec ?? '', store.referenceStat));
  /** The picker's own list: `store.weightStats` when the spec sent one, the full pinned
   *  vocabulary otherwise (sub-item 4; `weights.ts`'s own fallback rule). */
  const pickable = $derived(pickableStatsFor(store.weightStats));

  /**
   * `store.save` (bulk-store.svelte.ts) has no fallback title of its own, the same as every
   * other tool page's save form -- see ComboResults.svelte's identical comment. A weights
   * run has no winner to headline, so this names the tool and the spec instead, the same
   * two-part title SavedSim.svelte gives a saved weights page's own <h1>.
   */
  const saveTitleFor = $derived(`${bulkCopy.weightsTitle} · ${specLabel(store.character?.spec ?? '')}`);

  let published = $state<PublishedWeights | null>(null);
  $effect(() => {
    const character = store.character;
    const preset = requestPresetFor(store.settings);
    if (character === null) {
      published = null;
      return;
    }
    let live = true;
    fetchBisFile(activeBuild.build, character.spec)
      .then((file) => {
        if (!live) return;
        published =
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

<Page layout="split" panels testid="sim-stat-weights">
  {#snippet bar()}<BulkRunBar {store} />{/snippet}
  {#if lead}{@render lead()}{/if}
  <Notice
    text={bulkCopy.weightsWarning}
    tone="neutral"
    action={{ label: bulkCopy.weightsWarningLink, href: '/sim/gear' }}
    testid="sim-weights-warning"
  />
  {#if store.weights.length > 0}
    <Section
      panel
      frame="framed"
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
    <Section
      panel
      frame="framed"
      title={bulkCopy.weightsPublishedTitle}
      action={{ label: published.label, href: published.href }}
      testid="sim-weights-published"
    >
      <WeightsScaleTable view={published.view} />
    </Section>
  {/if}
  {#snippet rail()}
    <Section title={bulkCopy.weightsPick} panel testid="sim-weights-pick">
      {#if store.referenceStat}
        <p class="text-small text-muted m-0" data-testid="sim-weights-reference">
          {bulkCopy.weightsReference}: {statLabel(store.referenceStat)}
        </p>
      {/if}
      {#if hasWeightStats(store.weightStats)}
        <p class="text-small text-muted m-0" data-testid="sim-weights-stats-note">
          {WEIGHTS_STATS_FROM_ENGINE}
        </p>
      {/if}
      <ul class="m-0 grid list-none grid-cols-2 gap-x-4 p-0">
        {#each pickable as stat (stat.id)}
          <li>
            <label class="text-text flex min-h-11 items-center gap-2 text-[14px]">
              <input
                type="checkbox"
                class="accent-gold h-[18px] w-[18px]"
                data-testid={`sim-weight-pick-${stat.id}`}
                checked={store.stats.includes(stat.id)}
                disabled={stat.id === store.referenceStat}
                onchange={() => togglePick(stat.id)}
              />
              {stat.label}
            </label>
          </li>
        {/each}
      </ul>
    </Section>
  {/snippet}
</Page>
