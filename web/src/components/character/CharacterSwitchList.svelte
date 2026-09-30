<!-- web/src/components/character/CharacterSwitchList.svelte -->
<!-- The character list the spine bar's Switch control opens (spec 2026-09-25 section 4.1:
     "opens the same character list the account menu shows") -- the same CharacterRow every
     other character list in the app draws, wired to the exact write-and-broadcast the
     account menu's own switch already performs (`writeCurrent`/`CURRENT_CHARACTER_CHANGED`),
     left to the caller so this stays a pure render of the list. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import { currentCharacterCopy } from '../../lib/current-character-copy';
  import CharacterRow from './CharacterRow.svelte';

  let {
    characters,
    currentKey,
    onswitch,
    /** Home rebuild spec §3.B.4: the home page's own "Switch character" panel wants the
     *  full race/class descriptor line, not the spine bar's realm/ruleset one -- an
     *  additive, backward-compatible prop (default unchanged) so the spine bar
     *  (CurrentCharacterBar.svelte) keeps its own behaviour with no caller change. */
    descriptor = 'realm',
    testid = 'current-character-bar-switch-list',
    rowTestidPrefix = 'current-character-bar-switch-row',
  }: {
    characters: MeCharacter[];
    currentKey: string | null;
    onswitch: (character: MeCharacter) => void;
    descriptor?: 'full' | 'realm' | 'none';
    testid?: string;
    rowTestidPrefix?: string;
  } = $props();
</script>

<ul class="flex flex-col" data-testid={testid}>
  {#each characters as character (character.key)}
    <CharacterRow {character} {descriptor} testid={`${rowTestidPrefix}-${character.key}`}>
      {#snippet action()}
        {#if character.key === currentKey}
          <span class="text-muted text-[12px]" data-testid="current-character-bar-switch-current">
            {currentCharacterCopy.switchCurrentMarker}
          </span>
        {:else}
          <button
            type="button"
            class="text-nav hover:text-strong flex min-h-11 items-center px-2 text-[13px]"
            data-testid={`current-character-bar-switch-${character.key}`}
            onclick={() => onswitch(character)}
          >
            {currentCharacterCopy.switchAction}
          </button>
        {/if}
      {/snippet}
    </CharacterRow>
  {/each}
</ul>
