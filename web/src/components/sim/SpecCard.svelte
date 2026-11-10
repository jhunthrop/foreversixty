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
  import { oneDecimal } from '../../lib/sim/result-copy';
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
  import { classCrestSrc } from '../../lib/class-crest';
  import type { SpecTopSet } from '../../lib/sim/bis-benchmark';

  let {
    row,
    actionNames = null,
    compact = false,
    top = null,
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
    /** The spec's raid BiS set at 60 and its place among the specs (/sim/specs); null hides it. */
    top?: SpecTopSet | null;
  } = $props();

  const noParses = $derived(row.median_gap === null);
  const figure = $derived(
    noParses
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
  const showRotation = $derived(!compact && rotationSteps.length > 0);
</script>

{#snippet title()}
  <!-- What it does sits on the name's line (ux-designer, round 5); opened, its steps take the
       line under the name. -->
  <span class="flex flex-wrap items-center gap-x-4 [&_details[open]]:basis-full">
    <span>{specLabel(row.spec)}</span>
    {#if showRotation}
      <span class="contents font-normal">
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
      </span>
    {/if}
  </span>
{/snippet}

{#snippet status()}
  <!-- A spec with no parses says so once, in its Tag; on the in-page card (compact) the line
       stays, as the only status there. -->
  <span class="flex flex-col">
    <span class="text-strong tabular-nums" data-testid="spec-figure">{figure}</span>
    {#if !noParses}<span>{specStateNote(row.state)}</span>{/if}
  </span>
{/snippet}

{#snippet details()}
  {#if row.worst_actions.length > 0}
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
  {#if footer !== ''}
    <span class="text-small text-muted tabular-nums" data-testid="spec-footer">{footer}</span>
  {/if}
{/snippet}

{#snippet crest()}
  <img src={classCrestSrc(classOfSpec(row.spec))} alt="" width="28" height="28" class="rounded-full" />
{/snippet}

<Row
  id={row.spec}
  lead={compact ? undefined : crest}
  density={compact ? 'default' : 'slim'}
  title={specLabel(row.spec)}
  titleTone={classTone(classOfSpec(row.spec))}
  titleContent={title}
  subContent={!noParses || compact ? status : undefined}
  foot={!compact && (row.worst_actions.length > 0 || footer !== '') ? details : undefined}
  testid={`spec-${row.spec}`}
  open={false}
>
  {#snippet trail()}
    <!-- A spec no parse has measured carries no Tag on the list: it runs, and "Not yet" beside
         every unmeasured row read as "not simulated". The in-page card (compact) keeps it. -->
    <span class="flex items-center gap-3">
      {#if compact || !noParses}
        <Tag tone={row.state === 'validated' ? 'success' : 'muted'} testid="spec-state"
          >{specStateLabel(row.state)}</Tag
        >
      {/if}
      {#if top !== null}
        <span class="text-muted w-8 text-right text-[13px] tabular-nums" data-testid="spec-top-rank"
          >{simCopy.specsTopSetRank(top.rank)}</span
        >
        <span
          class="text-strong w-[92px] text-right text-[15px] font-bold tabular-nums"
          data-testid="spec-top-dps"
          >{oneDecimal(top.dps)} <span class="text-muted text-[13px] font-normal">DPS</span></span
        >
      {/if}
    </span>
  {/snippet}
</Row>
