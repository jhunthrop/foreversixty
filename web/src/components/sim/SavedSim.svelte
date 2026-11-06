<!-- web/src/components/sim/SavedSim.svelte -->
<!-- /sim/<sim_id> (spec 2026-11-06-sim §4.5, rulings S1 and S5): a `run` result is a document:
     the head (the crest, the title by the shipped rule, the run's facts, "Saved {date}") with its
     answer (the figure, the band labelled once as 95%, Open in the sim and Copy link), then the
     report body. No tool tabs and no run line: Open in the sim is the way to run it. A bulk or
     weights result keeps its own view (SavedKind.svelte). Opening a result marks that server run
     seen, once (experience spec 2026-10-24 3.6). -->
<script lang="ts">
  import { onDestroy, onMount } from 'svelte';
  import { createLazyComponent } from '../../lib/report/lazy-component.svelte';
  import { resultAnswer } from '../../lib/sim/answer';
  import { requestKind } from '../../lib/sim/bulk-types';
  import { raceWords } from '../../lib/sim/head-meta';
  import { markSimSeenOnOpen } from '../../lib/sim/mark-seen-on-open';
  import { simPageCopy as copy } from '../../lib/sim/page-copy';
  import { bandCopy, grouped } from '../../lib/sim/result-copy';
  import {
    bandHeading,
    buffsFact,
    savedDate,
    talentSplit,
    variationPercent,
  } from '../../lib/sim/result-band';
  import { durationLabel } from '../../lib/sim/settings';
  import { specLabel } from '../../lib/sim/spec-label';
  import type { SimResult, SpecFidelity } from '../../lib/sim/types';
  import { createCopyFeedback, type CopyFeedbackState } from '../../lib/ui/copy-feedback';
  import Button from '../ui/Button.svelte';
  import RunReport from './RunReport.svelte';
  import SimAnswer from './SimAnswer.svelte';
  import SimCharacterHead from './SimCharacterHead.svelte';
  import SimFidelityCaption from './SimFidelityCaption.svelte';

  let {
    result,
    onrerun,
    simId = '',
    fidelity = null,
  }: {
    result: SimResult;
    onrerun: () => void;
    simId?: string;
    /** The spec list's row for the run's spec: the answer's fidelity caption (§4.6). */
    fidelity?: SpecFidelity | null;
  } = $props();

  onMount(() => {
    void markSimSeenOnOpen(simId !== '' ? simId : (result.sim_id ?? ''));
  });

  const kind = $derived(requestKind(result.request));
  // A bulk or weights result is rare and carries its own view: a chunk of its own.
  const savedKindLazy = createLazyComponent(() => import('./SavedKind.svelte'));
  $effect(() => {
    if (kind !== 'run') savedKindLazy.load();
  });
  const model = $derived({
    ...resultAnswer(result, specLabel(result.request.spec)),
    label: copy.answer.savedLabel(grouped(result.iterations_run)),
  });

  const lines = $derived.by(() => {
    const { character, encounter } = result.request;
    const fight = bandCopy.fight(
      durationLabel(encounter.duration_sec),
      variationPercent(encounter.variation),
    );
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
    const saved = savedDate(result.created_at);
    return saved === null ? [line1] : [line1, copy.head.saved(saved)];
  });

  /** How long "Copy failed" shows before the label returns. */
  const COPY_FAILED_MS = 3000;
  let copyState = $state<CopyFeedbackState>('idle');
  let failedTimer: ReturnType<typeof setTimeout> | undefined;
  const copier = createCopyFeedback(
    (text) => navigator.clipboard.writeText(text),
    (next) => {
      copyState = next;
      clearTimeout(failedTimer);
      if (next === 'failed') failedTimer = setTimeout(() => (copyState = 'idle'), COPY_FAILED_MS);
    },
  );
  onDestroy(() => {
    copier.dispose();
    clearTimeout(failedTimer);
  });
  const copyLabel = $derived(
    copyState === 'copied'
      ? copy.answer.copied
      : copyState === 'failed'
        ? copy.answer.copyFailed
        : copy.answer.copyLink,
  );
</script>

{#if kind === 'run'}
  <div data-testid="sim-saved">
    <SimCharacterHead name={bandHeading(result)} classSlug={result.request.character.class} {lines}>
      {#snippet answer()}
        <SimAnswer {model}>
          {#snippet extra()}<SimFidelityCaption {fidelity} />{/snippet}
          {#snippet actions()}
            <Button tone="primary" size="sm" onclick={onrerun} testid="sim-run-yourself"
              >{copy.answer.openInSim}</Button
            >
            <Button size="sm" onclick={() => void copier.copy(window.location.href)} testid="sim-saved-copy"
              ><span aria-live="polite">{copyLabel}</span></Button
            >
          {/snippet}
        </SimAnswer>
      {/snippet}
    </SimCharacterHead>
    <RunReport {result} {onrerun} />
  </div>
{:else}
  <div class="flex flex-col gap-[22px] pt-[22px] md:gap-8 md:pt-8">
    {#if savedKindLazy.current}<savedKindLazy.current {result} {onrerun} />{/if}
  </div>
{/if}
