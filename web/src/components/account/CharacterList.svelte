<!-- web/src/components/account/CharacterList.svelte -->
<!-- /account's Characters section (premium spec 2026-11-08 section 7.2): the page's framed panel, a Row
     per character (the ringed crest, the name in its class colour with Current and Main, the
     descriptor and the guild, then the build source, Set as main and Open), the last raid as the
     panel's foot row when there is one, and under it, in the panel, the Battle.net line and the reports link. The
     head's action refreshes from Battle.net: the page leaves for Battle.net and comes back to
     `?refreshed=1`, where the caller compares and hands the result in. A pure render of what the
     caller holds: no fetch of its own but the last-raid line's. -->
<script lang="ts">
  import { smallButton } from '../../lib/account/small-button';
  import { battlenetStartUrl, type MeCharacter } from '../../lib/account/api';
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import { characterDescriptor } from '../../lib/account/character-descriptor';
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import { saveSnapshot, snapshotOf, type RefreshResult } from '../../lib/account/refresh-snapshot';
  import { characterHref } from '../../lib/characters';
  import { lowLevelCopy, splitByLevel } from '../../lib/account/level-filter';
  import { classTone, LINK } from '../../lib/ui/primitives';
  import Crest from '../ui/Crest.svelte';
  import EmptyState from '../ui/EmptyState.svelte';
  import ShowAll from '../ui/ShowAll.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Tag from '../ui/Tag.svelte';
  import CharacterListIntro from './CharacterListIntro.svelte';
  import CoachLastRaidLine from './CoachLastRaidLine.svelte';

  /** Where Refresh lands: a second Battle.net login re-runs the import and comes back here. */
  const REFRESH_HREF = battlenetStartUrl('/account?refreshed=1');
  /** A success line gives way to the plain import time after this long; an error stays. */
  const RESULT_HOLD_MS = 8000;
  const REPORTS_HREF = '/reports';

  let {
    characters,
    bnetImportedAt,
    mainKey = undefined,
    currentKey = undefined,
    onSetMain = undefined,
    logCoach = false,
    viaBattlenet = true,
    refreshResult = null,
    now = undefined,
  }: {
    characters: MeCharacter[];
    bnetImportedAt?: string;
    /** The chosen main's key: its row carries Main, every other row Set as main. */
    mainKey?: string;
    /** The header's character: its row carries Current. */
    currentKey?: string;
    /** Records the main; the caller updates `mainKey`. A rejection prints the failure line. */
    onSetMain?: (key: string) => Promise<void>;
    /** `me.entitlements.log_coach`: the Premium or the free form of the last-raid line. */
    logCoach?: boolean;
    /** Signed in with Battle.net: an empty list means Battle.net has no Forever character yet. */
    viaBattlenet?: boolean;
    /** What the last Battle.net refresh did, on the return to `?refreshed=1`; null otherwise. */
    refreshResult?: RefreshResult | null;
    now?: Date;
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

  let refreshing = $state(false);
  let resultShown = $state(true);
  $effect(() => {
    if (refreshResult === null || refreshResult.kind === 'error') return;
    resultShown = true;
    const timer = setTimeout(() => (resultShown = false), RESULT_HOLD_MS);
    return () => clearTimeout(timer);
  });
  const result = $derived(resultShown ? refreshResult : null);

  // Back from Battle.net with the browser's Back: the page is not refreshing any more.
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
  /** The head's muted meta: the success line for a while, else when Battle.net last brought them in. */
  const meta = $derived.by(() => {
    // The empty Plate says where characters come from; the head keeps only its Refresh (spec 7.2).
    if (refreshing || result?.kind === 'error' || characters.length === 0) return undefined;
    if (result !== null) return announcement;
    return bnetImportedAt === undefined ? undefined : characterListCopy.updatedFrom(bnetImportedAt, now);
  });
  /** Refresh is offered once there is something to refresh, or Battle.net is the account's way in. */
  const canRefresh = $derived(bnetImportedAt !== undefined || (viaBattlenet && characters.length === 0));

  let showLow = $state(false);
  /** Level 10 and up by default; the current character and the main always show. */
  const listed = $derived(splitByLevel(characters, showLow, [currentKey, mainKey]));

  const sourceOf = (character: MeCharacter): string => {
    if (character.build === undefined) return characterListCopy.noBuildYet;
    return character.build.source === 'addon'
      ? characterListCopy.addonSource
      : characterListCopy.battlenetSource;
  };
</script>

{#snippet refreshLink(label: string)}
  <a
    class="{LINK} text-[14px] [@media(pointer:coarse)]:-my-3 [@media(pointer:coarse)]:inline-block [@media(pointer:coarse)]:py-3"
    href={REFRESH_HREF}
    onclick={startRefresh}
    data-testid="bnet-refresh">{label} ›</a
  >
{/snippet}

{#snippet head()}
  {#if refreshing}
    <span class="text-small text-gold" aria-disabled="true" data-testid="bnet-refreshing"
      >{characterListCopy.refreshing}</span
    >
  {:else if result?.kind === 'error'}
    <span class="text-small text-wipe" data-testid="bnet-refresh-result"
      >{characterListCopy.refreshError}</span
    >
    {@render refreshLink(characterListCopy.tryAgain)}
  {:else if canRefresh}
    {@render refreshLink(characterListCopy.refreshFromBattlenet)}
  {/if}
{/snippet}

<Section
  title={characterListCopy.heading}
  {meta}
  {head}
  anchor="characters"
  testid="account-characters"
  panel
  frame="framed"
  flush
>
  {#if announcement !== ''}<p class="sr-only" role="status" data-testid="bnet-refresh-status">
      {announcement}
    </p>{/if}
  <div>
    <div
      class="account-plate-body"
      class:account-refreshing={refreshing}
      aria-busy={refreshing ? 'true' : undefined}
    >
      {#if refreshing}<span class="refresh-bar" aria-hidden="true" data-testid="bnet-refresh-bar"></span>{/if}
      {#if characters.length === 0}
        <div class="px-5 pt-4 pb-5">
          <EmptyState
            message={viaBattlenet ? characterListCopy.emptyLead : characterListCopy.emptyByEmailLead}
            fill={viaBattlenet ? characterListCopy.emptyKorea : characterListCopy.emptyByEmailFill}
            action={{ label: characterListCopy.getSetUp, href: characterListCopy.getSetUpHref }}
            secondary={{ label: characterListCopy.pasteAnExport, href: characterListCopy.pasteHref }}
            testid="account-characters-empty"
          />
        </div>
      {:else}
        {#each listed.shown as character (character.key)}
          {@const isMain = mainKey === character.key}
          <Row
            title={character.name}
            titleTone={classTone(character.class ?? '')}
            wrapTrail
            testid="account-character-row"
          >
            {#snippet lead()}<Crest cls={character.class ?? ''} size={40} />{/snippet}
            {#snippet titleContent()}
              <span class="inline-flex flex-wrap items-center gap-2"
                ><span data-testid="character-name">{character.name}</span>
                {#if currentKey === character.key}<Tag testid="character-current-pill"
                    >{characterListCopy.current}</Tag
                  >{/if}
                {#if isMain}<Tag tone="muted" testid="character-main-pill">{characterListCopy.main}</Tag>{/if}
              </span>
            {/snippet}
            {#snippet subContent()}
              <span data-testid="character-descriptor">{characterDescriptor(character)}</span>
              {#if character.guild !== undefined}<br /><span data-testid="character-guild"
                  >{character.guild.name}</span
                >{/if}
            {/snippet}
            {#snippet trail()}
              {#if character.item_level !== undefined}<span
                  class="text-strong text-[15px] whitespace-nowrap tabular-nums"
                  data-testid="character-item-level">{characterListCopy.itemLevel(character.item_level)}</span
                >{/if}
              <Tag tone="muted" testid="character-source">{sourceOf(character)}</Tag>
              {#if !isMain && onSetMain !== undefined}
                <button
                  type="button"
                  class={smallButton()}
                  disabled={settingMain !== ''}
                  aria-busy={settingMain === character.key ? 'true' : undefined}
                  onclick={() => void setMain(character.key)}
                  data-testid="character-set-main">{characterListCopy.setAsMain}</button
                >
              {/if}
              <a
                class="{LINK} text-[15px] whitespace-nowrap"
                href={characterHref(character.region, character.ruleset, character.name)}
                data-testid="character-open">{characterListCopy.open} ›</a
              >
            {/snippet}
          </Row>
        {/each}
        {#if listed.lowCount > 0}
          <div class="px-5 py-2">
            <ShowAll
              shown={listed.shown.length + listed.hidden - listed.lowCount}
              total={listed.shown.length + listed.hidden}
              expanded={showLow}
              onToggle={() => (showLow = !showLow)}
              label={lowLevelCopy.show(listed.lowCount)}
              fewerLabel={lowLevelCopy.hide}
              testid="account-low-level-toggle"
            />
          </div>
        {/if}
        <CoachLastRaidLine premium={logCoach} />
      {/if}
    </div>
  </div>
  {#if setMainError !== '' || characters.length > 0}
    <div class="account-characters-foot">
      {#if setMainError !== ''}
        <p class="text-wipe pt-2 text-[13px]" role="alert" data-testid="character-set-main-error">
          {setMainError}
        </p>
      {/if}
      {#if characters.length > 0}
        <CharacterListIntro />
        <p class="pt-2">
          <a class={LINK} href={REPORTS_HREF} data-testid="account-reports-link"
            >{accountPageCopy.reportsLink} ›</a
          >
        </p>
      {/if}
    </div>
  {/if}
</Section>

<style>
  /* The panel's text under the rows, at the panel's 20 px inset (14 on a phone). */
  .account-characters-foot {
    padding: 0 20px 10px;
  }
  @media (max-width: 767px) {
    .account-characters-foot {
      padding: 0 14px 8px;
    }
  }
  .account-plate-body {
    position: relative;
    transition: opacity 120ms;
  }
  .account-refreshing > :global(.ui-row) {
    opacity: 0.6;
  }
</style>
