<!-- web/src/components/guild/planner/PlannerRow.svelte -->
<!-- One sign-up on the raid planner (web layout spec 2026-11-05 §7.2): the class crest, the name in
     its class colour, `Holy Paladin · Healer`, the account's other answered characters and the note,
     then the state. A member reads the state as a Tag; an officer gets the in-row In / Late / Bench /
     Out control, which writes the same sign-up the Discord buttons write (its own line on phone).
     While a change saves the control shows it pressed; a refusal reverts it. From 1024 the row
     also carries its figures (PlannerFacts) in fixed columns under the roster's one head line. -->
<script lang="ts">
  import { raidCopy } from '../../../lib/guild/raid-copy';
  import type { GuildSignupState } from '../../../lib/guild/api';
  import { rowSub, stateTag } from '../../../lib/guild/raid-planner-view';
  import type { RosterRow } from '../../../lib/guild/raid-roster';
  import { phoneFacts, signed, type PlannerFacts } from '../../../lib/guild/planner-facts';
  import { PLANNER_FACT_COLUMNS } from './planner-columns';
  import AttendanceStrip from '../dense/AttendanceStrip.svelte';
  import { classTone, ROW_ITEM } from '../../../lib/ui/primitives';
  import Crest from '../../ui/Crest.svelte';
  import Row from '../../ui/Row.svelte';
  import Tag from '../../ui/Tag.svelte';
  import SignupSegments from '../SignupSegments.svelte';

  let {
    row,
    sub,
    roles,
    officer,
    disabled = false,
    describedBy,
    busy = false,
    refusal = null,
    facts = null,
    stateNamed = false,
    onSetState,
  }: {
    row: RosterRow;
    /** The raider's figures; null draws no figure columns. */
    facts?: PlannerFacts | null;
    /** Replaces the spec line (a non-signer's `No Discord linked`). */
    sub?: string;
    /** The night counts roles (a raid or a run, not a guild event). */
    roles: boolean;
    officer: boolean;
    disabled?: boolean;
    /** The gate sentence a read-only control names. */
    describedBy?: string;
    busy?: boolean;
    /** A refused change's line under this row. */
    refusal?: string | null;
    /** The group head above already names this row's state (Late, Bench, Out): no Tag repeats it. */
    stateNamed?: boolean;
    onSetState: (row: RosterRow, state: GuildSignupState) => Promise<boolean>;
  } = $props();

  /** The state pressed while its save is in flight; dropped when it settles. */
  let chosen = $state<GuildSignupState | null>(null);
  const shown = $derived(chosen ?? row.state);

  async function change(state: GuildSignupState): Promise<void> {
    chosen = state;
    await onSetState(row, state);
    chosen = null;
  }
</script>

<li class={ROW_ITEM} aria-busy={busy} data-testid="raid-roster-row">
  <Row
    title={row.name}
    titleTone={classTone(row.class)}
    selected={row.own}
    wrapTrail={officer}
    density={facts === null ? 'default' : 'slim'}
  >
    {#snippet lead()}<Crest cls={row.class} size={28} />{/snippet}
    {#snippet subContent()}
      <span class="block truncate" title={sub ?? rowSub(row, roles)}>{sub ?? rowSub(row, roles)}</span>
      {#if facts !== null && facts.why.length > 0}<span class="text-wipe block" data-testid="raid-roster-why"
          >{facts.why.join(', ')}</span
        >{/if}
      {#if facts !== null}<span class="block lg:hidden" data-testid="raid-roster-facts-phone"
          >{phoneFacts(facts)}</span
        >{/if}
      {#if row.also !== null}<span class="block break-words" data-testid="raid-roster-also">{row.also}</span
        >{/if}
      {#if row.note !== ''}<span class="text-text block break-words" data-testid="raid-roster-note"
          >{row.note}</span
        >{/if}
      {#if refusal !== null}<span class="text-wipe block" role="alert" data-testid="raid-roster-refusal"
          >{refusal}</span
        >{/if}
    {/snippet}
    <!-- In is the roster's default: a raider's row carries a Tag only for Late, Bench and Out, and
         none under a group head that already names its state. -->
    {#snippet trail()}
      {#if officer}
        <SignupSegments
          value={shown}
          name={row.name}
          size="band"
          {disabled}
          {describedBy}
          onChange={(state) => void change(state)}
          testid={`raid-roster-state-${row.character_key}`}
        />
      {:else if row.state !== null && row.state !== 'in' && !stateNamed}
        {@const tag = stateTag(row.state)}
        <Tag tone={tag.tone} testid="raid-roster-state">{tag.text}</Tag>
      {/if}
      {#if facts !== null}
        <span class="hidden {PLANNER_FACT_COLUMNS} text-[14px] lg:grid" data-testid="raid-roster-facts">
          <span class="flex justify-end"
            >{#if facts.cells !== null}<AttendanceStrip cells={facts.cells} />{/if}</span
          >
          <span class="text-strong">
            {facts.itemLevel ?? '–'}{#if facts.itemLevelDelta !== null}
              <span class="text-muted ml-1 font-normal">{signed(facts.itemLevelDelta)}</span>{/if}
          </span>
          {#if facts.short === null}<span class="text-muted">–</span>
          {:else if facts.short === 0}<span class="text-kill" aria-label="Ready">✓</span>
          {:else}<span class="text-wipe" title={facts.why.join(', ')}>{raidCopy.missing(facts.short)}</span
            >{/if}
          <span>{facts.rating ?? '–'}</span>
          <span>{facts.attendancePct === null ? '–' : `${facts.attendancePct}%`}</span>
        </span>
      {/if}
    {/snippet}
  </Row>
</li>
