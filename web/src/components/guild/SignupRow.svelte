<!-- web/src/components/guild/SignupRow.svelte -->
<!-- Phone (< 768): who at the left of line 1 with the 176 px spec picker at its right, the
     segments on line 2, the note on line 3, then the status line: at most 180 px (spec-ask e2e). -->
<!-- One character's row of the sign-up strip (raids spec 3.3, raid comp spec 3.2): who (crest, name,
     `Holy Paladin`), the segments, one spec picker whose (spec, role) pair sets the role, and the
     note, then a reserved 20 px status line. Picking on a signed-up row saves at once; on an
     unsigned row it only sets what the next press sends. In, Late or Bench with no spec picked
     waits: `Pick a spec first.`, the picker takes the wipe border and focus, and the first pick saves
     the held press. The row keeps its own optimistic values while a save is in flight and reverts
     them when the save is refused; the caller applies a successful save's raid before resolving.
     A guild event asks no spec: the row is who (crest, name, class), the segments and the note,
     and a press saves at once (guild events spec 3.5, guild round 3). -->
<script lang="ts">
  import { tick } from 'svelte';
  import { MediaQuery } from 'svelte/reactivity';
  import type { GuildRaidViewerCharacter, GuildSignupState, GuildZoneKind } from '../../lib/guild/api';
  import { FIELD_CHROME } from '../../lib/guild/field-styles';
  import { focusOnMount } from '../../lib/guild/focus-now';
  import {
    classWord,
    pickOf,
    specClassLine,
    specOptions,
    specPicks,
    whoLine,
  } from '../../lib/guild/raid-comp';
  import { kindCountsRoles, raidCopy } from '../../lib/guild/raid-copy';
  import { RAID_WIDE_QUERY } from '../../lib/guild/raid-layout';
  import { specIconSrc } from '../../lib/guild/raid-spec';
  import { createSerialSave } from '../../lib/guild/serial-save';
  import {
    NOTE_MAX,
    SAVED_LINE_MS,
    SPEC_ASK_IDLE_MS,
    defaultPick,
    failureLine,
    noteCounter,
    pressWaitsForSpec,
    roleOfPick,
    stateTakesSpec,
    type SignupChoice,
    type SignupOutcome,
  } from '../../lib/guild/signup-strip';
  import { classColorVar } from '../../lib/report/format';
  import { ROW_ITEM } from '../../lib/ui/primitives';
  import Crest from '../ui/Crest.svelte';
  import Select from '../ui/Select.svelte';
  import SignupSegments from './SignupSegments.svelte';

  let {
    character,
    savedNote,
    disabled,
    describedBy,
    notice,
    seenAs = null,
    kind = 'raid',
    onClearNotice,
    onSave,
    onMoved,
  }: {
    character: GuildRaidViewerCharacter;
    /** The note on the character's sign-up, `''` while unsigned. */
    savedNote: string;
    disabled: boolean;
    describedBy?: string;
    /** A line the strip sets on this row: `{Name} moved to Bench.` */
    notice: string | null;
    /** When this answered character is not its account's roster row: `Officers see Quillfeather (Late).` */
    seenAs?: string | null;
    /** What the night is, for the ended line of a refused save. */
    kind?: GuildZoneKind;
    onClearNotice: () => void;
    onSave: (choice: SignupChoice) => Promise<SignupOutcome>;
    onMoved: (name: string) => void;
  } = $props();

  type RowStatus =
    { kind: 'saved' } | { kind: 'error'; message: string; race: boolean } | { kind: 'pick-first' };

  /** `undefined` while the row shows the props' state. */
  let pendingState = $state<GuildSignupState | undefined>(undefined);
  /** The pick on this row not yet confirmed by the props: unsaved on an unsigned row, or in flight. */
  let pickChoice = $state<string | null>(null);
  let noteEdit = $state<string | null>(null);
  let busy = $state(false);
  /** In, Late or Bench was pressed with no spec picked: the pick, not the press, signs up. */
  let askedState = $state<GuildSignupState | null>(null);
  let rowEl = $state<HTMLLIElement | undefined>(undefined);
  let idleTimer: ReturnType<typeof setTimeout> | undefined;
  let status = $state<RowStatus | null>(null);
  /** The note last sent, so blur after Enter does not save it twice. */
  let sentNote: string | null = null;
  let savedTimer: ReturnType<typeof setTimeout> | undefined;

  const key = $derived(character.character_key);
  /** A raid or a run asks the spec; a guild event does not. */
  const asksSpec = $derived(kindCountsRoles(kind));
  const specId = $derived(`signup-spec-${key}`);
  const statusId = $derived(`signup-status-${key}`);
  /** From 768 the select's face reads `Restoration · Healer`; below, the spec alone and the Who
   *  line carries the role (raid comp spec 3.2). The server renders the desktop face. */
  const wide = new MediaQuery(RAID_WIDE_QUERY, true);
  const options = $derived(specOptions(specPicks(character.class), !wide.current, specIconSrc));
  const shownState = $derived(pendingState === undefined ? character.state : pendingState);
  const shownPick = $derived(
    pickChoice === null ? defaultPick(character) : pickOf(character.class, pickChoice),
  );
  const shownNote = $derived(noteEdit ?? savedNote);
  const counter = $derived(noteCounter(shownNote));
  const who = $derived(
    !asksSpec
      ? classWord(character.class)
      : shownPick === null
        ? specClassLine(character.class, character.spec_key)
        : whoLine(character.class, shownPick.spec.spec, shownPick.role, !wide.current),
  );
  const asking = $derived(askedState !== null);

  $effect(() => () => {
    clearTimeout(savedTimer);
    clearTimeout(idleTimer);
  });

  function cancelAsk(): void {
    clearTimeout(idleTimer);
    askedState = null;
    if (status?.kind === 'pick-first') status = null;
  }

  function armIdle(): void {
    clearTimeout(idleTimer);
    idleTimer = setTimeout(cancelAsk, SPEC_ASK_IDLE_MS);
  }

  /** While a pick is owed: focus inside the row holds the press; leaving it for 60 s, or Escape,
   *  drops it (nothing was saved, so there is nothing to revert on the server). */
  $effect(() => {
    if (askedState === null || rowEl === undefined) return;
    const row = rowEl;
    const hold = (): void => clearTimeout(idleTimer);
    const leave = (event: FocusEvent): void => {
      if (!row.contains(event.relatedTarget as Node | null)) armIdle();
    };
    const escape = (event: KeyboardEvent): void => {
      if (event.key === 'Escape') cancelAsk();
    };
    row.addEventListener('focusin', hold);
    row.addEventListener('focusout', leave);
    row.addEventListener('keydown', escape);
    return () => {
      row.removeEventListener('focusin', hold);
      row.removeEventListener('focusout', leave);
      row.removeEventListener('keydown', escape);
    };
  });

  function showSaved(): void {
    clearTimeout(savedTimer);
    status = { kind: 'saved' };
    savedTimer = setTimeout(() => {
      status = null;
    }, SAVED_LINE_MS);
  }

  async function settle(choice: SignupChoice): Promise<SignupOutcome> {
    sentNote = choice.note;
    try {
      return await onSave(choice);
    } catch {
      return { ok: false, code: 'network', message: '' };
    }
  }

  /** One answer. While a follow-up is owed the row keeps its edits: the follow-up carries them. */
  function settled(outcome: SignupOutcome, choice: SignupChoice, more: boolean): void {
    if (outcome.ok && outcome.moved !== null) onMoved(outcome.moved);
    if (more) return;
    pendingState = undefined;
    // A sent pick is now the props' (or refused, and reverts to them); an unsent one stays.
    if (choice.spec !== undefined) pickChoice = null;
    if (!outcome.ok) {
      sentNote = null;
      status = { kind: 'error', ...failureLine(outcome.code, kind) };
      return;
    }
    if (noteEdit === choice.note) noteEdit = null;
    showSaved();
  }

  /** Saves are serialised per row: a press during a save is sent, merged, when it settles. In, Late
   *  and Bench carry the pick, so state and spec save in one request. */
  const saver = createSerialSave<SignupChoice, SignupOutcome>({
    read: () => {
      if (shownState === null) return null;
      const pick = asksSpec && stateTakesSpec(shownState) ? shownPick : null;
      return {
        state: shownState,
        role: roleOfPick(asksSpec ? shownPick : null, character),
        note: shownNote,
        ...(pick === null ? {} : { spec: pick.spec.spec }),
      };
    },
    send: settle,
    settled,
    onBusy: (next) => (busy = next),
  });

  function save(): void {
    if (askedState !== null) return;
    status = null;
    onClearNotice();
    void saver.request();
  }

  async function holdForSpec(state: GuildSignupState): Promise<void> {
    askedState = state;
    status = { kind: 'pick-first' };
    armIdle();
    await tick();
    document.getElementById(specId)?.focus();
  }

  function pickState(state: GuildSignupState): void {
    if (asksSpec && pressWaitsForSpec(state, shownPick)) {
      void holdForSpec(state);
      return;
    }
    cancelAsk();
    pendingState = state;
    save();
  }

  function pickSpec(value: string): void {
    if (value === '' || pickOf(character.class, value) === null || value === shownPick?.value) return;
    pickChoice = value;
    const held = askedState;
    cancelAsk();
    if (held !== null) pendingState = held;
    if (held !== null || shownState !== null) save();
  }

  function commitNote(): void {
    if (shownState === null || shownNote === (sentNote ?? savedNote)) return;
    save();
  }

  function onNoteKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Enter') return;
    event.preventDefault();
    commitNote();
  }
</script>

{#if !character.verified}
  <li class="{ROW_ITEM} flex min-h-[68px] items-center gap-3 py-3" data-testid={`signup-row-${key}`}>
    <span class="opacity-60"><Crest cls={character.class} size={32} /></span>
    <div class="min-w-0">
      <p class="heading-name text-muted text-[14px]">{character.name}</p>
      <p class="text-muted text-[13px]">{raidCopy.waitingCharacter}</p>
    </div>
  </li>
{:else}
  <li bind:this={rowEl} class="{ROW_ITEM} pt-1.5 md:pt-3" aria-busy={busy} data-testid={`signup-row-${key}`}>
    <div
      class="grid grid-cols-[minmax(0,1fr)_176px] items-center gap-x-2 gap-y-1 md:min-h-[68px] md:gap-4 {asksSpec
        ? 'md:grid-cols-[minmax(180px,1fr)_auto_220px]'
        : 'md:grid-cols-[minmax(180px,1fr)_auto_minmax(200px,1.4fr)]'}"
      data-testid={`signup-grid-${key}`}
    >
      <div
        class="col-start-1 row-start-1 flex min-w-0 items-center gap-3 md:col-span-1 {asksSpec
          ? ''
          : 'col-span-2'}"
      >
        <Crest cls={character.class} size={32} />
        <div class="min-w-0">
          <p class="heading-name truncate text-[14px]" style={`color:${classColorVar(character.class)}`}>
            {character.name}
          </p>
          <p
            class="text-muted text-[12px] leading-4 md:truncate md:leading-normal"
            data-testid={`signup-who-${key}`}
          >
            {who}
          </p>
        </div>
      </div>
      <div class="col-span-2 row-start-2 md:col-span-1 md:col-start-2 md:row-start-1">
        <SignupSegments
          value={shownState}
          pending={askedState}
          name={character.name}
          size="strip"
          {disabled}
          {describedBy}
          onChange={pickState}
          testid={`signup-segments-${key}`}
        />
      </div>
      {#if asksSpec}
        <div class="col-start-2 row-start-1 flex w-[176px] min-w-0 md:col-start-3 md:w-auto">
          <Select
            id={specId}
            class="w-full"
            size="compact"
            label={raidCopy.spec.label(character.name)}
            placeholder={raidCopy.spec.placeholder}
            value={shownPick?.value ?? ''}
            {options}
            {disabled}
            {busy}
            invalid={asking}
            describedBy={disabled ? describedBy : asking ? statusId : undefined}
            testid={specId}
            onchange={(event) => pickSpec((event.currentTarget as HTMLSelectElement).value)}
          />
        </div>
      {/if}
      <div
        class="relative col-span-2 row-start-3 {asksSpec
          ? 'md:col-span-3 md:col-start-1 md:row-start-2'
          : 'md:col-span-1 md:col-start-3 md:row-start-1'}"
      >
        <input
          type="text"
          class="{FIELD_CHROME} pr-16 disabled:opacity-50"
          maxlength={NOTE_MAX}
          placeholder={raidCopy.placeholders.note}
          aria-label={raidCopy.placeholders.note}
          aria-describedby={disabled ? describedBy : undefined}
          value={shownNote}
          {disabled}
          data-testid={`signup-note-${key}`}
          oninput={(event) => (noteEdit = (event.currentTarget as HTMLInputElement).value)}
          onblur={commitNote}
          onkeydown={onNoteKeydown}
        />
        {#if counter !== null}
          <span
            class="text-muted numeric pointer-events-none absolute top-1/2 right-3 -translate-y-1/2 text-[12px]"
            data-testid={`signup-counter-${key}`}>{counter}</span
          >
        {/if}
      </div>
    </div>
    <div id={statusId} class="min-h-5 text-[13px] leading-5" data-testid={`signup-status-${key}`}>
      {#key status}
        {#if status?.kind === 'error' && status.race}
          <p class="text-wipe" role="alert" tabindex="-1" use:focusOnMount>{status.message}</p>
        {:else if status?.kind === 'error'}
          <p class="text-wipe" role="alert">{status.message}</p>
        {:else if status?.kind === 'pick-first'}
          <p class="text-wipe" role="alert">{raidCopy.spec.pickFirst}</p>
        {:else if status?.kind === 'saved'}
          <p class="text-text" role="status">{raidCopy.stripStatus.saved}</p>
        {:else if notice !== null}
          <p class="text-text" role="status">{notice}</p>
        {:else if seenAs !== null}
          <p class="text-muted" data-testid={`signup-seen-as-${key}`}>{seenAs}</p>
        {/if}
      {/key}
    </div>
  </li>
{/if}
