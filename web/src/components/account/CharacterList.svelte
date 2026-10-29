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
  import { button } from '../../lib/ui/button';
  import { relativeTime } from '../../lib/dates';
  import CharacterRow from '../character/CharacterRow.svelte';
  import CharacterRowLink from './CharacterRowLink.svelte';
  import CoachLastRaidLine from './CoachLastRaidLine.svelte';
  import EmptyState from '../ui/EmptyState.svelte';
  import CardHead from '../ui/CardHead.svelte';
  import Tag from '../ui/Tag.svelte';

  let {
    characters,
    bnetImportedAt,
    mainKey = undefined,
    currentKey = undefined,
    onSetMain = undefined,
    logCoach = false,
  }: {
    /** `me.entitlements.log_coach`: picks the Premium or the free form of the last-raid line. */
    logCoach?: boolean;
    characters: MeCharacter[];
    bnetImportedAt?: string;
    /** The chosen main's key; the row carries a Main pill and every other row a Set as main control. */
    mainKey?: string;
    /** The character the page speaks for (the header's current one): its row is marked. */
    currentKey?: string;
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
    <span class="text-muted" data-testid="bnet-imported">
      {characterListCopy.importedFrom(relativeTime(new Date(bnetImportedAt)))}
      <a href={REFRESH_HREF}>{characterListCopy.refreshFromBattlenet}</a>
    </span>
  {/if}
{/snippet}

<section class="card-surface overflow-hidden" data-testid="account-characters">
  <CardHead glyph="user" title={characterListCopy.heading} aside={importedAside} />
  <CoachLastRaidLine premium={logCoach} />
  {#if characters.length === 0}
    <div class="border-line-soft border-t px-5 py-4">
      <EmptyState
        message={characterListCopy.empty}
        action={{ label: characterListCopy.getSetUp, href: characterListCopy.getSetUpHref }}
        secondary={{ label: characterListCopy.pasteAnExport, href: characterListCopy.pasteHref }}
        testid="account-characters-empty"
      />
    </div>
  {:else}
    <ul class="m-0 flex list-none flex-col p-0">
      {#each characters as character (character.key)}
        <CharacterRow
          {character}
          settingsRow
          descriptor="full"
          guildLine
          href={characterHref(character.region, character.ruleset, character.name)}
          descriptorTestid="character-descriptor"
        >
          {#snippet action()}
            {#if currentKey === character.key}
              <Tag testid="character-current-pill">{characterListCopy.current}</Tag>
            {/if}
            {#if mainKey === character.key}
              <Tag tone="muted" testid="character-main-pill">{characterListCopy.main}</Tag>
            {:else if onSetMain !== undefined}
              <button
                type="button"
                class={button('secondary', 'sm')}
                onclick={() => void setMain(character.key)}
                disabled={settingMain !== ''}
                aria-busy={settingMain === character.key ? 'true' : undefined}
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
    <p class="text-wipe px-5 py-2 text-[13px]" role="alert" data-testid="character-set-main-error">
      {setMainError}
    </p>
  {/if}
  {#if characters.length > 0}
    <p class="text-muted border-line-soft border-t px-5 py-4 text-[13px]">
      {characterListCopy.introBattlenetLine}
      <a class="text-gold underline" href="/setup">{characterListCopy.introInstallAddonLink}</a>
      {characterListCopy.introAddonTail}
      <a class="text-gold underline" href="/setup#paste">{characterListCopy.introPasteLink}</a>
      {characterListCopy.introPasteTail}
    </p>
  {/if}
</section>
