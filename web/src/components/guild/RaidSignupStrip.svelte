<!-- web/src/components/guild/RaidSignupStrip.svelte -->
<!-- Your sign-up (spec 3.3, 4 "Page states" and "Races"): one row per character of the viewer in
     this guild, the header selector's current character first. Locked and ended put their line
     above disabled rows; cancelled replaces the rows with its line. The strip owns every row, so
     it can tell the row the API moved to Bench (R3) that it moved. -->
<script lang="ts">
  import type { GuildRaidView } from '../../lib/guild/api';
  import { raidCopy } from '../../lib/guild/raid-copy';
  import { orderedCharacters, type SignupChoice, type SignupHandler } from '../../lib/guild/signup-strip';
  import SignupRow from './SignupRow.svelte';

  let {
    raid,
    locked,
    ended,
    cancelled,
    onSignup,
    currentKey = null,
  }: {
    raid: GuildRaidView;
    locked: boolean;
    ended: boolean;
    cancelled: boolean;
    onSignup: SignupHandler;
    /** The header selector's current character, drawn first. */
    currentKey?: string | null;
  } = $props();

  /** `{Name} moved to Bench.` lines, by character key. */
  let notices = $state<Readonly<Record<string, string>>>({});

  const rows = $derived(orderedCharacters(raid.viewer.characters, currentKey));
  const savedNotes = $derived(new Map(raid.signups.map((s) => [s.character_key, s.note])));
  const stateLine = $derived(locked ? raidCopy.raidStates.locked : ended ? raidCopy.raidStates.ended : null);
  const stateLineId = $derived(`raid-strip-state-${raid.id}`);
  /** A done raid refuses sign-ups like a locked one; its line lives in the band. */
  const rowsDisabled = $derived(stateLine !== null || raid.status === 'done');

  function setNotice(key: string, line: string | null): void {
    const { [key]: _dropped, ...rest } = notices;
    notices = line === null ? rest : { ...rest, [key]: line };
  }

  function noteMoved(name: string): void {
    const moved = raid.viewer.characters.find((c) => c.name === name);
    if (moved !== undefined) setNotice(moved.character_key, raidCopy.stripStatus.movedToBench(name));
  }

  function save(key: string, choice: SignupChoice): ReturnType<SignupHandler> {
    return onSignup(key, choice.state, choice.role, choice.note);
  }
</script>

<section
  class="bg-raised border-line rounded-panel border p-4 shadow-[inset_0_1px_0_rgba(229,185,85,.35)] md:p-5"
  aria-labelledby="raid-signup-heading"
  data-testid="raid-signup-strip"
>
  <h2 id="raid-signup-heading" class="section-title text-[15px]">{raidCopy.subheads.yourSignup}</h2>
  {#if cancelled}
    <p class="text-muted mt-3 text-[14px]" data-testid="raid-strip-state">{raidCopy.raidStates.cancelled}</p>
  {:else}
    {#if stateLine !== null}
      <p id={stateLineId} class="text-text mt-3 text-[14px]" data-testid="raid-strip-state">{stateLine}</p>
    {/if}
    <ul class="mt-1">
      {#each rows as character (character.character_key)}
        <SignupRow
          {character}
          savedNote={savedNotes.get(character.character_key) ?? ''}
          disabled={rowsDisabled}
          describedBy={stateLine !== null ? stateLineId : undefined}
          notice={notices[character.character_key] ?? null}
          onClearNotice={() => setNotice(character.character_key, null)}
          onSave={(choice) => save(character.character_key, choice)}
          onMoved={noteMoved}
        />
      {/each}
    </ul>
  {/if}
</section>
