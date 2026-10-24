<!-- web/src/components/account/CharacterList.svelte -->
<!-- /account's Characters panel (brief 2026-09-22 §B3): the design system's Character row
     inside a StatePanel, in place of the plain text rows the previous lane shipped. A pure
     render of the two props Account.svelte already has in hand -- `me.characters` and
     `me.bnet_imported_at` -- it makes no fetch of its own, and neither does each row's
     `CharacterRowLink` (spec 2026-09-22 §3.4): unlike `CharacterHandoffLinks` (still used,
     one row at a time, by the hero band and the character page), a list of N characters
     must not cost N `sim-input` fetches just to draw its rows. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import { battlenetStartUrl } from '../../lib/account/api';
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import { characterHref } from '../../lib/characters';
  import { relativeTime } from '../../lib/dates';
  import CharacterRow from '../character/CharacterRow.svelte';
  import CharacterRowLink from './CharacterRowLink.svelte';
  import EmptyState from '../ui/EmptyState.svelte';
  import StatePanel from '../ui/StatePanel.svelte';

  let {
    characters,
    bnetImportedAt,
    mainKey = undefined,
    currentKey = undefined,
    onSetMain = undefined,
    supporter = false,
  }: {
    characters: MeCharacter[];
    bnetImportedAt?: string;
    /** The chosen main's key; the row carries a Main pill and every other row a Set as main control. */
    mainKey?: string;
    /** The character the page speaks for (the header's current one): its row is marked. */
    currentKey?: string;
    /** The account shows the Supporter mark: every row carries the chip. */
    supporter?: boolean;
    /** Called with the key when Set as main is pressed; the caller records it and updates `mainKey`. */
    onSetMain?: (key: string) => Promise<void>;
  } = $props();
  let settingMain = $state('');
  let setMainError = $state('');
  async function setMain(key: string): Promise<void> {
    if (onSetMain === undefined) return;
    settingMain = key;
    setMainError = '';
    try {
      await onSetMain(key);
    } catch {
      setMainError = characterListCopy.setAsMainFailed;
    } finally {
      settingMain = '';
    }
  }

  /** `?next=` on the refresh link (spec §7.2/§7.3): a second Battle.net login re-runs the
   *  import and lands back here with the toast query param. */
  const REFRESH_HREF = battlenetStartUrl('/account?refreshed=1');
</script>

{#snippet importedAside()}
  {#if bnetImportedAt !== undefined}
    <span data-testid="bnet-imported">
      {characterListCopy.importedFrom(relativeTime(new Date(bnetImportedAt)))}
      <a href={REFRESH_HREF}>{characterListCopy.refreshFromBattlenet}</a>
    </span>
  {/if}
{/snippet}

<StatePanel
  label={characterListCopy.heading}
  aside={bnetImportedAt === undefined ? undefined : importedAside}
  testid="account-characters"
>
  {#if characters.length === 0}
    <EmptyState
      message={characterListCopy.empty}
      action={{ label: characterListCopy.getSetUp, href: characterListCopy.getSetUpHref }}
      secondary={{ label: characterListCopy.pasteAnExport, href: characterListCopy.pasteHref }}
      testid="account-characters-empty"
    />
  {:else}
    <ul class="flex flex-col">
      {#each characters as character (character.key)}
        <CharacterRow
          {character}
          descriptor="full"
          guildLine
          {supporter}
          href={characterHref(character.region, character.ruleset, character.name)}
          descriptorTestid="character-descriptor"
        >
          {#snippet action()}
            {#if currentKey === character.key}
              <span class="pill pill-site" data-testid="character-current-pill"
                >{characterListCopy.current}</span
              >
            {/if}
            {#if mainKey === character.key}
              <span class="pill pill-site" data-testid="character-main-pill">{characterListCopy.main}</span>
            {:else if onSetMain !== undefined}
              <button
                type="button"
                class="text-nav border-line-warm rounded-control inline-flex min-h-11 items-center border px-3 text-[13px] font-semibold md:min-h-10"
                onclick={() => void setMain(character.key)}
                disabled={settingMain !== ''}
                aria-busy={settingMain === character.key}
                data-testid="character-set-main"
              >
                {characterListCopy.setAsMain}
              </button>
            {/if}
            <CharacterRowLink {character} />
          {/snippet}
        </CharacterRow>
      {/each}
    </ul>
  {/if}
  {#if setMainError !== ''}
    <p class="text-[13px]" role="alert" data-testid="character-set-main-error">{setMainError}</p>
  {/if}
  {#if characters.length > 0}
    <p class="text-muted text-[13px]">
      {characterListCopy.introBattlenetLine}
      <a class="text-text underline" href="/setup">{characterListCopy.introInstallAddonLink}</a>
      {characterListCopy.introAddonTail}
      <a class="text-text underline" href="/setup#paste">{characterListCopy.introPasteLink}</a>
      {characterListCopy.introPasteTail}
    </p>
  {/if}
</StatePanel>
