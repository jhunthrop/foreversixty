<!-- web/src/components/sim/SpecCard.svelte -->
<!-- "Can I trust this" in one row of /sim/specs' Plate (mock design/mocks/ws_simtools.py
     `web-sim-specs`): the spec in its class colour, the measured gap and the state's note under it,
     the state as a Tag at the right, then the rotation, the worst actions and the engine stamp on
     the row's foot. The figure is whatever the validation job measured,
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
  import { formatPercent } from '../../lib/report/format';
  import { resolveActionName, type ActionNames } from '../../lib/sim/action-names';
  import { simCopy } from '../../lib/sim/copy';
  import { classOfSpec, specLabel } from '../../lib/sim/spec-label';
  import { specStateLabel, specStateNote } from '../../lib/sim/spec-state';
  import { classTone } from '../../lib/ui/primitives';
  import Disclosure from '../ui/Disclosure.svelte';
  import Row from '../ui/Row.svelte';
  import Tag from '../ui/Tag.svelte';
  import { loadRotationNotes, rotationNotesFor, type RotationNotes } from '../../lib/sim/rotations';
  import type { SpecFidelity } from '../../lib/sim/types';
  import { engineLabel } from '../../lib/sim/version';

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
  // The notes chunk (rotations.ts) is shared by every card; the disclosure appears once it lands.
  let notes = $state<RotationNotes>({});
  loadRotationNotes().then(
    (loaded) => (notes = loaded),
    (error: unknown) => console.error('sim: the rotation notes did not load', error),
  );
  const rotationSteps = $derived(rotationNotesFor(row.spec, notes));
</script>

<Row
  id={row.spec}
  title={specLabel(row.spec)}
  titleTone={classTone(classOfSpec(row.spec))}
  testid={`spec-${row.spec}`}
  open={false}
>
  {#snippet subContent()}
    <span class="flex flex-col">
      <span class="text-strong tabular-nums" data-testid="spec-figure">{figure}</span>
      <!-- "No parses yet" already says what the note would; one status a row (the Tag). -->
      {#if figure !== simCopy.specNoParses}<span>{specStateNote(row.state)}</span>{/if}
    </span>
  {/snippet}
  {#snippet trail()}
    <Tag tone={row.state === 'validated' ? 'success' : 'muted'} testid="spec-state"
      >{specStateLabel(row.state)}</Tag
    >
  {/snippet}
  {#snippet foot()}
    {#if !compact && rotationSteps.length > 0}
      <div class="w-full">
        <Disclosure
          summary={simCopy.rotationLink}
          testid={`spec-rotation-panel-${row.spec}`}
          summaryTestid={`spec-rotation-trigger-${row.spec}`}
        >
          <ol
            class="m-0 flex list-decimal flex-col gap-1 pl-5 text-[14px]"
            data-testid={`spec-rotation-steps-${row.spec}`}
          >
            {#each rotationSteps as step, index (index)}
              <li>{step}</li>
            {/each}
          </ol>
        </Disclosure>
      </div>
    {/if}
    {#if !compact && row.worst_actions.length > 0}
      <span class="text-small text-muted tabular-nums">
        {simCopy.specWorstActions}:
        {#each row.worst_actions.slice(0, 3) as action, index (action.name)}{index > 0
            ? ' · '
            : ''}{resolveActionName(action.name, actionNames)}
          {action.actual_casts}
          {simCopy.specCast}, {action.sim_casts}
          {simCopy.specSimmed}{/each}
      </span>
    {/if}
    {#if !compact && footer !== ''}
      <span class="text-small text-muted tabular-nums" data-testid="spec-footer">{footer}</span>
    {/if}
  {/snippet}
</Row>
