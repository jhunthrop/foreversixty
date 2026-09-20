<!-- web/src/components/sim/SpecCard.svelte -->
<!-- "Can I trust this" in one card. The figure is whatever the validation job measured,
     published either way -- the design's whole point is that the answer lives on the site
     rather than in a footnote, so a bad gap is shown as plainly as a good one.
     The cast-gap rows read actual first: the player's own number is the one they recognise
     from their own log, and the simmed one is the claim being made about it.

     Task 6 (newcomer BLOCKER): this is the page RotationCard's "what it does" drawer links
     out to for the fidelity detail, so it answers "what does this rotation do" itself too,
     from the same rotations.ts the drawer reads -- one source, so the two can never
     disagree about what a spec's rotation is. Hidden in `compact` the same way
     worst_actions and the footer already are, and absent entirely for a spec with no
     curated step notes rather than an empty disclosure with nothing to open. -->
<script lang="ts">
  // formatPercent is the site's one percentage formatter: one decimal and a % sign. Every
  // other caller in report/format.ts pre-multiplies by 100 exactly as this does, so a
  // hand-rolled toFixed(1) here would be a second formatter that drifts the day the first
  // one changes.
  import { classColorVar, formatPercent } from '../../lib/report/format';
  import { resolveActionName, type ActionNames } from '../../lib/sim/action-names';
  import { simCopy } from '../../lib/sim/copy';
  import { classOfSpec, specLabel } from '../../lib/sim/spec-label';
  import { specPillClass, specStateLabel, specStateNote } from '../../lib/sim/spec-state';
  import { rotationNotesFor } from '../../lib/sim/rotations';
  import type { SpecFidelity } from '../../lib/sim/types';
  import { engineLabel } from '../../lib/sim/version';
  import Disclosure from './Disclosure.svelte';

  let {
    row,
    actionNames = null,
    compact = false,
  }: {
    row: SpecFidelity;
    /**
     * The build's name table (Task 23), for `worst_actions[].name`. `api sim-validate`
     * compares a native engine result against a fight's summary, so a name there may be an
     * engine action key. Null on `/sim/specs`, where cards for nine classes are shown at
     * once and loading nine tables for three rows of text is not worth the bytes; the
     * in-page unsupported card passes the store's, which is already loaded.
     */
    actionNames?: ActionNames | null;
    compact?: boolean;
  } = $props();

  const colour = $derived(classColorVar(classOfSpec(row.spec)));
  const figure = $derived(
    row.median_gap === null
      ? simCopy.specNoParses
      : `${simCopy.specMedianGap} ${formatPercent(row.median_gap * 100)} ${simCopy.specOver} ${row.parses} ${row.parses === 1 ? simCopy.specParse : simCopy.specParses}`,
  );
  const footer = $derived(
    [
      row.engine_version === '' ? '' : engineLabel(row.engine_version),
      row.updated_at === null ? '' : row.updated_at.slice(0, 10),
    ]
      .filter((part) => part !== '')
      .join(' · '),
  );
  const rotationSteps = $derived(rotationNotesFor(row.spec));
</script>

<article
  id={row.spec}
  class="border-line bg-card-top rounded-panel flex flex-col gap-3 border p-4"
  data-testid={`spec-${row.spec}`}
>
  <div class="flex flex-wrap items-center gap-2">
    <h3 class="text-[15px] font-semibold" style={`color: ${colour}`}>{specLabel(row.spec)}</h3>
    <span class={specPillClass(row.state)} data-testid="spec-state">{specStateLabel(row.state)}</span>
  </div>

  <p class="tabular text-strong font-mono text-[13px]" data-testid="spec-figure">{figure}</p>
  <p class="text-muted text-[13px]">{specStateNote(row.state)}</p>

  {#if !compact && rotationSteps.length > 0}
    <!-- Final whole-branch review, I2: these three testids were bare -- "spec-rotation-trigger"
         and friends, not "spec-rotation-trigger-<spec>" -- so all 20 damage-spec cards on this
         page carried the identical testid and the identical accessible name at once. Today's
         tests scope through `spec-<slug>` (the card's own testid) so nothing failed, but any
         unscoped locator is an instant Playwright strict-mode violation, and a screen reader
         user tabbing the page hears 20 indistinguishable "what it does" buttons with no
         difference between them. Suffixed with `row.spec`, matching the `id` the <article>
         above already sets. -->
    <Disclosure
      label={simCopy.rotationLink}
      id={`spec-rotation-${row.spec}`}
      triggerClass="label text-nav text-[12px] underline decoration-dotted underline-offset-2"
      panelClass="border-line-soft rounded-panel flex flex-col gap-2 border p-3"
      triggerTestId={`spec-rotation-trigger-${row.spec}`}
      panelTestId={`spec-rotation-panel-${row.spec}`}
    >
      <ol
        class="flex list-decimal flex-col gap-1 pl-5 text-[13px]"
        data-testid={`spec-rotation-steps-${row.spec}`}
      >
        {#each rotationSteps as step, index (index)}
          <li>{step}</li>
        {/each}
      </ol>
    </Disclosure>
  {/if}

  {#if !compact && row.worst_actions.length > 0}
    <div class="flex flex-col gap-1">
      <span class="label text-muted">{simCopy.specWorstActions}</span>
      {#each row.worst_actions.slice(0, 3) as action (action.name)}
        <span class="tabular text-muted font-mono text-[12px]">
          {resolveActionName(action.name, actionNames)}
          {action.actual_casts}
          {simCopy.specCast}, {action.sim_casts}
          {simCopy.specSimmed}
        </span>
      {/each}
    </div>
  {/if}

  {#if !compact && footer !== ''}
    <p class="tabular text-muted mt-auto font-mono text-[12px]" data-testid="spec-footer">{footer}</p>
  {/if}
</article>
