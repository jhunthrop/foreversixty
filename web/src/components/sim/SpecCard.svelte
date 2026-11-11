<!-- web/src/components/sim/SpecCard.svelte -->
<!-- "Can I trust this" in one row of /sim/specs' Plate (mock design/mocks/ws_simtools.py
     `web-sim-specs`): the spec in its class colour, the measured gap and the state's note under it,
     the state as a Tag at the right, then the worst actions and the engine stamp on
     the row's foot. The figure is whatever the validation job measured,
     published either way -- the design's whole point is that the answer lives on the site
     rather than in a footnote, so a bad gap is shown as plainly as a good one.
     The cast-gap rows read actual first: the player's own number is the one they recognise
     from their own log, and the simmed one is the claim being made about it.

     The rotation reads on /sim, in its rail: this list carries no toggle for it (h-audit
     2026-11-09: an unlabelled caret on every row). -->
<script lang="ts">
  // percent is the site's one percentage formatter: one decimal and a % sign. Every
  // other caller in report/format.ts pre-multiplies by 100 exactly as this does, so a
  // hand-rolled toFixed(1) here would be a second formatter that drifts the day the first
  // one changes.
  import { percent } from '../../lib/format/number';
  import { oneDecimal } from '../../lib/format/number';
  import { resolveActionName, type ActionNames } from '../../lib/sim/action-names';
  import { simCopy } from '../../lib/sim/copy';
  import { classOfSpec, specLabel } from '../../lib/sim/spec-label';
  import { specStateLabel, specStateNote } from '../../lib/sim/spec-state';
  import { classTone } from '../../lib/ui/primitives';
  import Row from '../ui/Row.svelte';
  import Tag from '../ui/Tag.svelte';
  import type { SpecFidelity } from '../../lib/sim/types';
  import { engineLabel } from '../../lib/sim/version';
  import { specLeadIcon } from '../../lib/sim/spec-icon';
  import Crest from '../ui/Crest.svelte';
  import Meter from '../ui/Meter.svelte';
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
      : `${simCopy.specMedianGap} ${percent(row.median_gap * 100)} ${simCopy.specOver} ${row.parses} ${row.parses === 1 ? simCopy.specParse : simCopy.specParses}`,
  );
  const footer = $derived(
    [
      row.engine_version === '' ? '' : engineLabel(row.engine_version),
      row.updated_at === null ? '' : row.updated_at.slice(0, 10),
    ]
      .filter((part) => part !== '')
      .join(' · '),
  );
  /** The spec's own talent-tab icon; the class crest when the build has not published it. */
  let iconFailed = $state(false);
  const lead = $derived(specLeadIcon(row.spec, iconFailed));
</script>

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
  {#if lead === null}<Crest cls={classOfSpec(row.spec)} size={28} />{:else}<img
      src={lead}
      alt=""
      width="28"
      height="28"
      class="rounded-[5px]"
      onerror={() => (iconFailed = true)}
    />{/if}
{/snippet}

<Row
  id={row.spec}
  lead={compact ? undefined : crest}
  density={compact ? 'default' : 'slim'}
  title={specLabel(row.spec)}
  titleTone={classTone(classOfSpec(row.spec))}
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
        <span class="text-muted w-12 text-right text-[13px] tabular-nums" data-testid="spec-top-rank"
          >{simCopy.specsTopSetRank(top.rank)}</span
        >
        <!-- The section's meta names the figure and its unit once ("Raid BiS DPS at 60"). -->
        <span class="hidden w-[240px] md:block" aria-hidden="true"><Meter value={top.share} /></span>
        <span
          class="text-strong w-[64px] text-right text-[15px] font-bold tabular-nums"
          data-testid="spec-top-dps">{oneDecimal(top.dps)}</span
        >
      {/if}
    </span>
  {/snippet}
</Row>
