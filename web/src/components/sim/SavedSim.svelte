<!-- web/src/components/sim/SavedSim.svelte -->
<!-- /sim/<sim_id> (spec 2026-11-06-sim §4.5, rulings S1 and S5): a `run` result is a document:
     the head (the crest, the title by the shipped rule, the run's facts, "Saved {date}") with its
     answer (the figure, the band labelled once as 95%, Open in the sim and Copy link), then the
     report body. No tool tabs and no run line: Open in the sim is the way to run it. A bulk or
     weights result keeps its own view (SavedKind.svelte). Opening a result marks that server run
     seen, once (experience spec 2026-10-24 3.6). -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { createLazyComponent } from '../../lib/report/lazy-component.svelte';
  import activeBuild from '../../data/active-build.json';
  import { resultAnswer } from '../../lib/sim/answer';
  import { rerunHref } from '../../lib/sim/rerun-href';
  import { requestKind } from '../../lib/sim/bulk-types';
  import { raceWords } from '../../lib/sim/head-meta';
  import { markSimSeenOnOpen } from '../../lib/sim/mark-seen-on-open';
  import { simPageCopy as copy } from '../../lib/sim/page-copy';
  import { bandCopy } from '../../lib/sim/result-copy';
  import { bandHeading, buffsFact, savedDay, talentSplit, variationPercent } from '../../lib/sim/result-band';
  import { clockSeconds } from '../../lib/format/duration';
  import type { PrerenderedNote } from '../../lib/sim/prerender-notes';
  import type { BisBenchmarks } from '../../lib/sim/bis-benchmark';
  import type { SimResult, SpecFidelity } from '../../lib/sim/types';
  import Button from '../ui/Button.svelte';
  import CopyButton from '../ui/CopyButton.svelte';
  import RunReport from './RunReport.svelte';
  import SimAnswer from './SimAnswer.svelte';
  import SimCharacterHead from './SimCharacterHead.svelte';
  import SimFidelityCaption from './SimFidelityCaption.svelte';

  let {
    result,
    onrerun,
    simId = '',
    fidelity = null,
    initialNotes,
    bisBenchmarks = {},
  }: {
    result: SimResult;
    onrerun: () => void;
    simId?: string;
    /** The spec list's row for the run's spec: the answer's fidelity caption (§4.6). */
    fidelity?: SpecFidelity | null;
    /** Worth knowing read at build time for a prerendered run (RunReport `initialNotes`). */
    initialNotes?: readonly PrerenderedNote[];
    /** The published sets a run is read against (lib/sim/bis-benchmark.ts). */
    bisBenchmarks?: BisBenchmarks;
  } = $props();

  onMount(() => {
    void markSimSeenOnOpen(simId !== '' ? simId : (result.sim_id ?? ''));
  });

  const kind = $derived(requestKind(result.request));
  /** "Open in the sim" as a link (it works before the island mounts); a Worth knowing line's own
   *  "Open in the sim" button goes the same way. */
  const openHref = $derived(rerunHref(result, activeBuild.build));
  function openInSim(): void {
    onrerun();
    window.location.href = openHref;
  }
  // A bulk or weights result is rare and carries its own view: a chunk of its own.
  const savedKindLazy = createLazyComponent(() => import('./SavedKind.svelte'));
  $effect(() => {
    if (kind !== 'run') savedKindLazy.load();
  });
  const model = $derived(resultAnswer(result));

  const lines = $derived.by(() => {
    const { character, encounter } = result.request;
    const fight = bandCopy.fight(clockSeconds(encounter.duration_sec), variationPercent(encounter.variation));
    const buffs = buffsFact(result, null);
    const buffsWords =
      buffs === bandCopy.buffsRaid ? 'raid buffs' : buffs === bandCopy.none ? 'no buffs' : 'custom buffs';
    const line1 = copy.head.savedMeta(
      raceWords(character.race),
      character.level,
      talentSplit(character.talents),
      fight,
      buffsWords,
      character.consumes.length,
    );
    const saved = savedDay(result.created_at);
    return saved === null ? [line1] : [line1, copy.head.saved(saved)];
  });
</script>

{#if kind === 'run'}
  <div data-testid="sim-saved">
    <SimCharacterHead name={bandHeading(result)} classSlug={result.request.character.class} {lines}>
      {#snippet answer()}
        <SimAnswer {model}>
          {#snippet badge()}<SimFidelityCaption {fidelity} />{/snippet}
        </SimAnswer>
      {/snippet}
      {#snippet actions()}
        <Button tone="primary" size="sm" href={openHref} onclick={onrerun} testid="sim-run-yourself"
          >{copy.answer.openInSim}</Button
        >
        <CopyButton text={() => window.location.href} label={copy.answer.copyLink} testid="sim-saved-copy" />
      {/snippet}
    </SimCharacterHead>
    <RunReport {result} onrerun={openInSim} {initialNotes} {bisBenchmarks} />
  </div>
{:else}
  <div class="flex flex-col gap-[22px] md:gap-8">
    {#if savedKindLazy.current}<savedKindLazy.current {result} onrerun={openInSim} />{/if}
  </div>
{/if}
