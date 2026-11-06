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
  import { saveSnapshot, snapshotOf, type RefreshResult } from '../../lib/account/refresh-snapshot';
  import Spinner from '../ui/Spinner.svelte';
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
    viaBattlenet = true,
    refreshResult = null,
  }: {
    /** What the last Battle.net refresh did, on the return to /account?refreshed=1 (Account.svelte);
     *  null otherwise. */
    refreshResult?: RefreshResult | null;
    /** Signed in with Battle.net: an empty list means Battle.net has no Forever character yet. */
    viaBattlenet?: boolean;
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
   *  import and lands back here with `?refreshed=1`. */
  const REFRESH_HREF = battlenetStartUrl('/account?refreshed=1');
  /** A success line gives way to the plain import time after this long; an error stays. */
  const RESULT_HOLD_MS = 8000;

  // The refresh leaves the page for Battle.net: until it does, the link is busy and the rows
  // dim under a progress bar, and what the account holds now is kept for the return to compare.
  let refreshing = $state(false);
  let resultShown = $state(true);
  $effect(() => {
    if (refreshResult === null || refreshResult.kind === 'error') return;
    resultShown = true;
    const timer = setTimeout(() => (resultShown = false), RESULT_HOLD_MS);
    return () => clearTimeout(timer);
  });
  const result = $derived(resultShown ? refreshResult : null);

  // A page restored from the back-forward cache (the player came back from Battle.net with the
  // browser's Back) is not refreshing any more.
  $effect(() => {
    const reset = (event: PageTransitionEvent): void => {
      if (event.persisted) refreshing = false;
    };
    window.addEventListener('pageshow', reset);
    return () => window.removeEventListener('pageshow', reset);
  });

  function startRefresh(event: MouseEvent): void {
    if (refreshing) {
      event.preventDefault();
      return;
    }
    refreshing = true;
    saveSnapshot(snapshotOf({ characters, bnet_imported_at: bnetImportedAt }), sessionStorage);
  }

  const announcement = $derived(
    refreshing
      ? characterListCopy.refreshingAnnounce
      : result === null
        ? ''
        : result.kind === 'error'
          ? characterListCopy.refreshError
          : result.kind === 'unchanged'
            ? characterListCopy.noChanges(result.count)
            : characterListCopy.updatedJustNow(result.count, result.added),
  );
</script>

{#snippet refreshLink(label: string)}
  {#if refreshing}
    <span
      class="text-gold inline-flex items-center gap-1.5"
      aria-disabled="true"
      data-testid="bnet-refreshing"
    >
      <Spinner />{characterListCopy.refreshing}
    </span>
  {:else}
    <a href={REFRESH_HREF} onclick={startRefresh} data-testid="bnet-refresh">{label}</a>
  {/if}
{/snippet}

{#snippet importedAside()}
  <span class="text-muted inline-flex flex-wrap items-center justify-end gap-x-1.5">
    {#if result !== null && !refreshing}
      {#if result.kind === 'error'}
        <span class="text-wipe" data-testid="bnet-refresh-result">{characterListCopy.refreshError}</span>
        {@render refreshLink(characterListCopy.tryAgain)}
      {:else}
        <span class="text-text" data-testid="bnet-refresh-result">{announcement}</span>
      {/if}
    {:else if bnetImportedAt !== undefined}
      {#if !refreshing}<span data-testid="bnet-imported"
          >{characterListCopy.importedFrom(relativeTime(new Date(bnetImportedAt)))}</span
        >{/if}
      {@render refreshLink(characterListCopy.refreshFromBattlenet)}
    {:else if viaBattlenet && characters.length === 0}
      <!-- The empty state says "refresh from Battle.net": the link it names, before any import. -->
      {@render refreshLink(characterListCopy.refreshFromBattlenet)}
    {/if}
  </span>
{/snippet}

<section class="card-surface relative overflow-hidden" data-testid="account-characters">
  {#if refreshing}
    <div class="refresh-bar" aria-hidden="true" data-testid="bnet-refresh-bar"></div>
  {/if}
  <p class="sr-only" role="status" aria-live="polite" data-testid="bnet-refresh-status">{announcement}</p>
  <CardHead glyph="user" title={characterListCopy.heading} aside={importedAside} />
  {#if characters.length === 0}
    <div class="border-line-soft border-t px-5 py-4">
      <EmptyState
        message={viaBattlenet ? characterListCopy.empty : characterListCopy.emptyByEmail}
        action={{ label: characterListCopy.getSetUp, href: characterListCopy.getSetUpHref }}
        secondary={{ label: characterListCopy.pasteAnExport, href: characterListCopy.pasteHref }}
        testid="account-characters-empty"
      />
    </div>
  {:else}
    <ul
      class={`m-0 flex list-none flex-col p-0 transition-opacity duration-[120ms] ${refreshing ? 'opacity-60' : ''}`}
      aria-busy={refreshing ? 'true' : undefined}
    >
      {#each characters as character (character.key)}
        <CharacterRow
          {character}
          settingsRow
          descriptor="full"
          guildLine
          href={characterHref(character.region, character.ruleset, character.name)}
          descriptorTestid="character-descriptor"
        >
          {#snippet tags()}
            {#if currentKey === character.key}
              <Tag testid="character-current-pill">{characterListCopy.current}</Tag>
            {/if}
            {#if mainKey === character.key}
              <Tag tone="muted" testid="character-main-pill">{characterListCopy.main}</Tag>
            {/if}
          {/snippet}
          {#snippet action()}
            {#if mainKey !== character.key && onSetMain !== undefined}
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
  <!-- The last raid is a coach hint, not about any one character: under the rows, above the
       help line. With no characters it would read as a stray line. -->
  {#if characters.length > 0}<CoachLastRaidLine premium={logCoach} />{/if}
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
