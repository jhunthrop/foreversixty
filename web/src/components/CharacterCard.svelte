<!-- web/src/components/CharacterCard.svelte -->
<!-- The best-in-slot header row's "Your character" panel (bis rebuild spec §4.B, §5). A
     Svelte island, not the `.astro` name the spec's own §7 suggests, because every one of
     its states (loading, error, signed out, signed in) depends on the session -- runtime
     data an Astro page's build-time frontmatter cannot have, the exact same reason
     `HomeAccountPanel.svelte` is a Svelte component and not an `.astro` one. Mounted
     `client:idle` (never `client:load`, spec §8's own Lighthouse note): this region sits
     well below the header's own h1, this page's LCP element, so a slightly deferred
     hydration costs nothing there.

     Cannot import `ClassCrest.astro`/`FactionMark.astro` (a framework component can't reach
     an `.astro` file -- see those two components' own header comments); inlines the
     identical `<img>` from the same `classCrestSrc`/`factionMarkSrc` path functions and
     `classColorVar` instead, the documented pattern every other Svelte island on the site
     already follows for the same reason. -->
<script lang="ts">
  import { createCharacterCardState } from '../lib/bis/character-card-state.svelte';
  import { classCrestSrc } from '../lib/class-crest';
  import { factionMarkSrc } from '../lib/faction-mark';
  import { classColorVar } from '../lib/report/format';
  import { classSlugFromName } from '../lib/report/tree-sizes';
  import { relativeTime } from '../lib/sim/sources';
  import { battlenetStartUrl } from '../lib/account/api';
  import { homePanelCopy } from '../lib/home-panel-copy';
  import { bisCopy } from '../lib/bis/copy';
  import { PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import Skeleton from './ui/Skeleton.svelte';
  import LoadError from './ui/LoadError.svelte';

  let {
    nextPath,
    switchHref = '/account',
    sendListHref = '/addon',
    installHref = '/setup#pair',
  }: {
    nextPath: string;
    switchHref?: string;
    sendListHref?: string;
    installHref?: string;
  } = $props();

  const state = createCharacterCardState();
  const character = $derived(state.character);
  const synced = $derived(character?.build !== undefined);
  const syncedRelative = $derived(
    character?.build === undefined ? '' : relativeTime(character.build.captured_at),
  );

  /**
   * The header's own race select defaults to the visitor's OWN character race once signed
   * in (ux-designer review round 1: "the header race select must default to the visitor's
   * character race ... never the faction default; the two panels must agree") -- reaches
   * outside this island's own root to both `ClassHeader` copies (`[data-testid="bis-race-
   * select"]`, one per faction, only one ever visible at a time) the same cross-island DOM-
   * reach pattern `HomeAccountPanel.svelte`'s own `inert` effect and `syncTabHrefs` already
   * use, since the select lives in a plain Astro component this island cannot otherwise
   * reach. Matched by the OPTION's own text against `character.race` rather than a slug, so
   * this needs no second race-name-to-slug mapping of its own: `ClassHeader` already renders
   * each option's real display name as its text content. A race the current page's class
   * cannot be (no matching option) leaves the select on its own per-faction default rather
   * than force a selection that does not exist here. */
  $effect(() => {
    const race = character?.race;
    if (race === undefined) return;
    document.querySelectorAll<HTMLSelectElement>('[data-testid="bis-race-select"]').forEach((select) => {
      const match = Array.from(select.options).find(
        (option) => option.textContent?.trim().toLowerCase() === race.toLowerCase(),
      );
      if (match !== undefined) select.value = match.value;
    });
  });
</script>

{#if state.status === 'loading'}
  <div class="panel-box character-card" data-testid="bis-character-card-loading">
    <Skeleton lines={3} testid="bis-character-card-skeleton" />
    <div class="character-card-tile-skeleton" aria-hidden="true">
      <Skeleton lines={1} rowHeight="h-10" />
      <Skeleton lines={1} rowHeight="h-10" />
      <Skeleton lines={1} rowHeight="h-10" />
    </div>
  </div>
{:else if state.status === 'failed'}
  <div class="panel-box character-card" data-testid="bis-character-card-error">
    <LoadError message={bisCopy.characterCardLoadError} onRetry={() => state.refresh()} />
  </div>
{:else if character === null}
  <!-- Signed out, or signed in with zero characters -- design system's own rule: "Signed
       out, the card is replaced by the sign-in button and one line of copy." -->
  <div class="panel-box character-card character-card-signed-out" data-testid="bis-character-card-signed-out">
    <a class={`${PRIMARY_BUTTON_FIXED} px-4`} href={battlenetStartUrl(nextPath)}>
      {homePanelCopy.signInButton}
    </a>
    <p class="text-muted character-card-signed-out-line">{homePanelCopy.signedOutLine}</p>
  </div>
{:else}
  {@const color = classColorVar(character.class)}
  {@const faction = character.faction}
  <div class="panel-box character-card" data-testid="bis-character-card">
    <div class="character-card-label-row">
      <span class="label character-card-label-text">
        <!-- A "Sample" pill is the approved mock's own illustrative labelling for its
             stand-in character, never a real runtime state of this card (ux-designer/
             wow-player review round 1): a signed-in visitor's own character -- synced or
             not -- is real, and is never labelled that way. Removed entirely, not
             conditioned on whether the character is synced, since both real states must be
             free of it; its own copy string left this module with it. -->
        {bisCopy.yourCharacterLabel}
      </span>
      <a class="character-card-switch" href={switchHref}>{bisCopy.switchLabel}</a>
    </div>
    <div class="character-card-identity">
      <img
        src={classCrestSrc(classSlugFromName(character.class ?? ''))}
        alt=""
        width="40"
        height="40"
        class="character-card-crest"
        style={`--c:${color}`}
      />
      <span class="character-card-identity-text">
        <span class="font-display character-card-name" style={`color:${color}`}>{character.name}</span>
        <span class="text-muted character-card-descriptor">
          {#if faction !== undefined}
            <img src={factionMarkSrc(faction)} alt="" width="16" height="16" class="character-card-faction" />
          {/if}
          {bisCopy.characterCardIdentityLine(character.level ?? 0, character.race ?? '')}
          ·
          {synced ? bisCopy.syncedRelative(syncedRelative) : bisCopy.notSyncedYetLabel}
        </span>
      </span>
    </div>
    {#if synced}
      <p class="text-muted character-card-fallback-line" data-testid="bis-character-card-no-worn-gear">
        {bisCopy.gearSyncNotHereYet}
      </p>
      <a
        class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text justify-center px-4`}
        href={sendListHref}
      >
        {bisCopy.sendListToAddon}
      </a>
    {:else}
      <p class="text-muted character-card-fallback-line" data-testid="bis-character-card-not-synced">
        {bisCopy.installAddonToCompare}
      </p>
      <a
        class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text justify-center px-4`}
        href={installHref}
      >
        {bisCopy.installTheAddon}
      </a>
    {/if}
  </div>
{/if}

<style>
  .panel-box {
    background: color-mix(in srgb, var(--color-raised) 85%, transparent);
    border: 1px solid var(--color-line);
    border-radius: var(--radius-panel);
  }
  .character-card {
    display: flex;
    flex-direction: column;
    gap: 12px;
    padding: 16px 20px;
    min-height: 196px;
  }
  .character-card-tile-skeleton {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 8px;
  }
  .character-card-signed-out {
    justify-content: center;
    gap: 10px;
  }
  .character-card-signed-out-line {
    font-size: 13px;
  }
  .character-card-label-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  .character-card-label-text {
    display: flex;
    align-items: center;
    gap: 8px;
  }
  .character-card-switch {
    font-size: 12px;
    font-weight: 700;
    letter-spacing: 0.06em;
    text-transform: uppercase;
  }
  .character-card-identity {
    display: flex;
    align-items: center;
    gap: 12px;
  }
  .character-card-crest {
    flex-shrink: 0;
    border-radius: 999px;
    box-shadow: 0 0 0 2px var(--c);
    background: var(--color-raised);
    object-fit: cover;
  }
  .character-card-identity-text {
    display: flex;
    flex-direction: column;
    min-width: 0;
  }
  .character-card-name {
    font-size: 15px;
    font-weight: 700;
  }
  .character-card-descriptor {
    display: flex;
    align-items: center;
    gap: 4px;
    font-size: 12px;
  }
  .character-card-faction {
    object-fit: contain;
  }
  .character-card-fallback-line {
    font-size: 13px;
  }
</style>
