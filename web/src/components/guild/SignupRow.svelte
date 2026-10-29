<!-- web/src/components/guild/SignupRow.svelte -->
<!-- One character's row of the sign-up strip (spec 3.3, 4): who, the segments, the role select
     and the note, then a reserved 20 px status line. The row keeps its own optimistic values while
     a save is in flight and reverts them when the save is refused. The caller applies a
     successful save's raid before resolving, so the props carry the saved values afterwards. -->
<script lang="ts">
  import type { GuildRaidRole, GuildRaidViewerCharacter, GuildSignupState } from '../../lib/guild/api';
  import { focusOnMount } from '../../lib/guild/focus-now';
  import { raidCopy } from '../../lib/guild/raid-copy';
  import {
    pressAsksSpec,
    savedSpec,
    showsSpecSelect,
    specOfClass,
    specsForClass,
  } from '../../lib/guild/raid-spec';
  import { createSerialSave } from '../../lib/guild/serial-save';
  import {
    NOTE_MAX,
    ROLE_FIELD_LABEL,
    SAVED_LINE_MS,
    SPEC_ASK_IDLE_MS,
    SPEC_RESERVE_ASK,
    SPEC_RESERVE_SAVED,
    failureLine,
    initialRole,
    noteCounter,
    type SignupChoice,
    type SignupOutcome,
  } from '../../lib/guild/signup-strip';
  import { classColorVar } from '../../lib/report/format';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import FilterSelect from '../rankings/FilterSelect.svelte';
  import SignupSpec from './SignupSpec.svelte';
  import SignupSegments from './SignupSegments.svelte';

  let {
    character,
    savedNote,
    disabled,
    describedBy,
    notice,
    raidTitle,
    raidDate,
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
    /** The raid's title and date, for the spec ask's lead sentence. */
    raidTitle: string;
    raidDate: string;
    onClearNotice: () => void;
    onSave: (choice: SignupChoice) => Promise<SignupOutcome>;
    onMoved: (name: string) => void;
  } = $props();

  type RowStatus = { kind: 'saved'; kept: string | null } | { kind: 'error'; message: string; race: boolean };

  const ROLE_OPTIONS = (['tank', 'healer', 'dps'] as const).map((role) => ({
    value: role,
    label: raidCopy.roleLabels[role],
  }));

  /** `undefined` while the row shows the props' state. */
  let pendingState = $state<GuildSignupState | undefined>(undefined);
  let roleChoice = $state<GuildRaidRole | null>(null);
  let noteEdit = $state<string | null>(null);
  let busy = $state(false);
  /** In, Late or Bench was pressed with no saved spec: the pick, not the press, signs up. Until
   *  then the segment reads pending and nothing is saved. */
  let askedState = $state<GuildSignupState | null>(null);
  let rowEl = $state<HTMLLIElement | undefined>(undefined);
  let idleTimer: ReturnType<typeof setTimeout> | undefined;
  /** The spec picked and in flight; `null` once the props carry the saved one. */
  let specChoice = $state<string | null>(null);
  let status = $state<RowStatus | null>(null);
  /** The note last sent, so blur after Enter does not save it twice. */
  let sentNote: string | null = null;
  let savedTimer: ReturnType<typeof setTimeout> | undefined;

  const key = $derived(character.character_key);
  const shownState = $derived(pendingState === undefined ? character.state : pendingState);
  const shownRole = $derived(roleChoice ?? initialRole(character));
  const shownNote = $derived(noteEdit ?? savedNote);
  const counter = $derived(noteCounter(shownNote));
  const saved = $derived(savedSpec(character));
  const specState = $derived(askedState ?? shownState);
  const showSpec = $derived(showsSpecSelect(character.class, specState));
  /** An unsigned row keeps the room the spec block takes, so a press shifts nothing below it. */
  const reservesSpec = $derived(character.state === null && specsForClass(character.class).length > 1);
  const chosenSpec = $derived(specChoice ?? saved?.spec ?? '');
  const specLead = $derived(
    saved === null
      ? raidCopy.spec.lead(character.name, raidCopy.segments[specState ?? 'in'], raidTitle, raidDate)
      : raidCopy.spec.leadChange(character.name, raidTitle, raidDate),
  );

  $effect(() => () => {
    clearTimeout(savedTimer);
    clearTimeout(idleTimer);
  });

  function cancelAsk(): void {
    clearTimeout(idleTimer);
    askedState = null;
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

  function showSaved(choice: SignupChoice): void {
    clearTimeout(savedTimer);
    const picked = choice.spec === undefined ? null : specOfClass(character.class, choice.spec);
    status = {
      kind: 'saved',
      kept: picked === null ? null : raidCopy.spec.kept(picked.name, character.name),
    };
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
    roleChoice = null;
    specChoice = null;
    if (!outcome.ok) {
      sentNote = null;
      status = { kind: 'error', ...failureLine(outcome.code) };
      return;
    }
    if (noteEdit === choice.note) noteEdit = null;
    showSaved(choice);
  }

  /** Saves are serialised per row: a press during a save is sent, merged, when it settles. A
   *  refusal reverts to the confirmed (props) state, which the caller updates on each success. */
  const saver = createSerialSave<SignupChoice, SignupOutcome>({
    read: () =>
      shownState === null
        ? null
        : {
            state: shownState,
            role: shownRole,
            note: shownNote,
            ...(specChoice === null ? {} : { spec: specChoice }),
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

  function isRole(value: string): value is GuildRaidRole {
    return ROLE_OPTIONS.some((option) => option.value === value);
  }

  function pickState(state: GuildSignupState): void {
    if (pressAsksSpec(character, state)) {
      askedState = state;
      status = null;
      armIdle();
      return;
    }
    cancelAsk();
    pendingState = state;
    save();
  }

  function pickSpec(slug: string): void {
    const spec = specOfClass(character.class, slug);
    if (spec === null) return;
    specChoice = slug;
    roleChoice = spec.role;
    if (askedState !== null) pendingState = askedState;
    cancelAsk();
    save();
  }

  function pickRole(value: string): void {
    if (!isRole(value)) return;
    roleChoice = value;
    if (shownState !== null) save();
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
  <li
    class="border-line-soft flex min-h-[68px] items-center gap-3 border-b py-3 last:border-b-0"
    data-testid={`signup-row-${key}`}
  >
    <span class="opacity-60"><ClassCrestRing characterClass={character.class} size={32} /></span>
    <div class="min-w-0">
      <p class="heading-name text-muted text-[14px]">{character.name}</p>
      <p class="text-muted text-[13px]">{raidCopy.waitingCharacter}</p>
    </div>
  </li>
{:else}
  <li
    bind:this={rowEl}
    class="border-line-soft border-b pt-3 last:border-b-0"
    aria-busy={busy}
    data-testid={`signup-row-${key}`}
  >
    <div
      class="grid grid-cols-[minmax(0,1fr)_120px] items-center gap-2 md:min-h-[68px] md:grid-cols-[minmax(180px,1fr)_auto_152px_minmax(200px,1.4fr)] md:gap-4"
    >
      <div class="col-start-1 row-start-1 flex min-w-0 items-center gap-3">
        <ClassCrestRing characterClass={character.class} size={32} />
        <div class="min-w-0">
          <p class="heading-name truncate text-[14px]" style={`color:${classColorVar(character.class)}`}>
            {character.name}
          </p>
          <p class="text-muted truncate text-[12px]">{character.spec}</p>
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
      <div class="col-start-2 row-start-1 flex md:col-start-3">
        <FilterSelect
          id={`signup-role-${key}`}
          label={ROLE_FIELD_LABEL}
          labelHidden
          value={shownRole}
          options={ROLE_OPTIONS}
          onchange={pickRole}
          {disabled}
          {describedBy}
          testid={`signup-role-${key}`}
        />
      </div>
      <div class="relative col-span-2 row-start-3 md:col-span-1 md:col-start-4 md:row-start-1">
        <input
          type="text"
          class="border-line-warm bg-raised text-text rounded-control placeholder:text-muted hover:border-line-warm-strong focus-visible:outline-gold h-11 w-full min-w-0 border pr-16 pl-3 text-[13px] transition-colors duration-[120ms] focus-visible:outline-2 disabled:cursor-not-allowed disabled:opacity-50"
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
            class="text-muted pointer-events-none absolute top-1/2 right-3 -translate-y-1/2 font-mono text-[12px] tabular-nums"
            data-testid={`signup-counter-${key}`}>{counter}</span
          >
        {/if}
      </div>
    </div>
    {#if showSpec}
      <SignupSpec
        id={`signup-spec-${key}`}
        classSlug={character.class}
        chosen={chosenSpec}
        saved={saved !== null}
        lead={specLead}
        {busy}
        {disabled}
        onPick={pickSpec}
      />
    {:else if reservesSpec}
      <div
        class={saved === null ? SPEC_RESERVE_ASK : SPEC_RESERVE_SAVED}
        aria-hidden="true"
        data-testid={`signup-spec-reserve-${key}`}
      ></div>
    {/if}
    <div class="min-h-5 pt-0.5 pb-2 text-[13px] leading-5" data-testid={`signup-status-${key}`}>
      {#key status}
        {#if status?.kind === 'error' && status.race}
          <p class="text-wipe" role="alert" tabindex="-1" use:focusOnMount>{status.message}</p>
        {:else if status?.kind === 'error'}
          <p class="text-wipe" role="alert">{status.message}</p>
        {:else if status?.kind === 'saved'}
          <p class="text-text" role="status">
            {raidCopy.stripStatus.saved}{status.kept === null ? '' : ` ${status.kept}`}
          </p>
        {:else if notice !== null}
          <p class="text-text" role="status">{notice}</p>
        {/if}
      {/key}
    </div>
  </li>
{/if}
