<!-- web/src/components/guild/planner/PlannerRow.svelte -->
<!-- One sign-up on the raid planner (web layout spec 2026-11-05 §7.2): the class crest, the name in
     its class colour, `Holy Paladin · Healer`, the account's other answered characters and the note,
     then the state. A member reads the state as a Tag; an officer gets the in-row In / Late / Bench /
     Out control, which writes the same sign-up the Discord buttons write (its own line on phone).
     While a change saves the control shows it pressed; a refusal reverts it. -->
<script lang="ts">
  import type { GuildSignupState } from '../../../lib/guild/api';
  import { rowSub, stateTag } from '../../../lib/guild/raid-planner-view';
  import type { RosterRow } from '../../../lib/guild/raid-roster';
  import { classTone, ROW_ITEM } from '../../../lib/ui/primitives';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
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
    onSetState,
  }: {
    row: RosterRow;
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
  <Row title={row.name} titleTone={classTone(row.class)} selected={row.own} wrapTrail={officer}>
    {#snippet lead()}<ClassCrestRing characterClass={row.class} size={28} />{/snippet}
    {#snippet subContent()}
      <span class="block truncate" title={sub ?? rowSub(row, roles)}>{sub ?? rowSub(row, roles)}</span>
      {#if row.also !== null}<span class="block break-words" data-testid="raid-roster-also">{row.also}</span
        >{/if}
      {#if row.note !== ''}<span class="text-text block break-words" data-testid="raid-roster-note"
          >{row.note}</span
        >{/if}
      {#if refusal !== null}<span class="text-wipe block" role="alert" data-testid="raid-roster-refusal"
          >{refusal}</span
        >{/if}
    {/snippet}
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
      {:else if row.state !== null}
        {@const tag = stateTag(row.state)}
        <Tag tone={tag.tone} testid="raid-roster-state">{tag.text}</Tag>
      {/if}
    {/snippet}
  </Row>
</li>
