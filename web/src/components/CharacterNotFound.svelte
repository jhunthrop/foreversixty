<!-- web/src/components/CharacterNotFound.svelte -->
<!-- The character page for a name the site has no record of, shown to a visitor (W.8): one
     plain sentence instead of an empty character card with a score of nothing, the lookup form
     to try another name, and the way to get a character onto the site. The browser title says
     the same, so a tab or a bookmark never claims the character exists. -->
<script lang="ts">
  import { battlenetStartUrl } from '../lib/account/api';
  import { wireSearchForm } from '../lib/character-search-dom';
  import {
    characterHref,
    characterKey,
    REGIONS,
    rulesetLabel,
    RULESETS,
    type CharacterPath,
  } from '../lib/characters';
  import { hasSessionHint } from '../lib/page-load';
  import type { CharacterSuggestion } from '../lib/character-search-api';
  import { characterHeaderCopy } from '../lib/character-profile';
  import { homeCharacterSearchCopy as lookup } from '../lib/home-landing-copy';
  import { scoreCopy } from '../lib/score/copy';
  import { PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import Select from './ui/Select.svelte';

  /** `known`: the site holds this character but has nothing on it, and the viewer does not hold it
   *  (P.1): its own sentence, and only a quiet "Get the addon and companion" link under the lookup box. */
  let { path, name, known = false }: { path: CharacterPath; name: string; known?: boolean } = $props();

  const FIELD = 'border-line-warm-strong bg-raised rounded-control text-text w-full border px-3 text-[15px]';
  const LINK = 'text-nav underline';

  $effect(() => {
    document.title = known ? `${name} · Forever Sixty` : scoreCopy.notFoundTitle;
  });

  /** The same name the site holds on another ruleset or region (a lookup that kept the wrong
   *  default): read once, on the client, never for a character the site already holds here. */
  let elsewhere = $state<CharacterSuggestion[]>([]);
  /** The search has answered (a match list, or none). */
  let searched = $state(false);
  const signedIn = typeof document !== 'undefined' && hasSessionHint();
  const here = characterKey(path.region, path.ruleset, path.slug);
  $effect(() => {
    if (known || name.trim() === '') return;
    const controller = new AbortController();
    void import('../lib/character-search-api')
      .then(({ searchCharacters }) => searchCharacters(name, { limit: 5, signal: controller.signal }))
      .then((rows) => {
        const want = name.trim().toLowerCase();
        elsewhere = rows.filter((row) => row.key !== here && row.name.toLowerCase().startsWith(want));
        searched = true;
      })
      .catch(() => {});
    return () => controller.abort();
  });

  /** Routes the form's submit to the profile it names, as the home page's lookup does. */
  function wire(form: HTMLFormElement): void {
    wireSearchForm(form);
  }
</script>

<div class="reveal flex flex-col gap-5" data-testid="character-not-found" id="character">
  <header class="flex flex-col gap-2">
    <h1 class="heading-page" data-testid="character-name">
      {known ? name : scoreCopy.notFoundTitle}
    </h1>
    <p class="text-[15px]" data-testid="character-not-found-line">
      {known
        ? scoreCopy.noDataLine(name)
        : scoreCopy.notFoundLine(name, characterHeaderCopy.chip(path.ruleset, path.region))}
    </p>
  </header>
  {#if !known && searched && elsewhere.length === 0}
    <p class="text-muted -mt-3 text-[14px]" data-testid="character-not-found-nowhere">
      {scoreCopy.notFoundNowhere}
    </p>
  {/if}
  {#if elsewhere.length > 0}
    <div class="flex flex-col gap-2" data-testid="character-not-found-elsewhere">
      <p class="quiet-label">{scoreCopy.notFoundElsewhere}</p>
      <ul class="flex flex-col gap-2">
        {#each elsewhere as row (row.key)}
          <li>
            <a
              class="{PRIMARY_BUTTON_FIXED} w-full justify-between px-4 md:w-auto"
              href={characterHref(row.region, row.ruleset, row.name)}
            >
              <span>{row.name}</span>
              <span class="text-[13px] font-semibold tracking-normal normal-case opacity-75"
                >{rulesetLabel(row.ruleset)} {row.region.toUpperCase()}</span
              >
            </a>
          </li>
        {/each}
      </ul>
    </div>
  {/if}
  {#if !known || !signedIn}
    <div
      class="border-line-warm bg-raised rounded-card flex flex-col items-start gap-3 border p-4"
      data-testid="character-not-found-signin"
    >
      {#if signedIn}
        <p class="text-[14px]">{scoreCopy.notFoundSignedIn}</p>
      {:else}
        <p class="text-[14px]">{known ? scoreCopy.knownYours : scoreCopy.notFoundNewCharacter}</p>
        <a
          class="{elsewhere.length > 0
            ? SECONDARY_BUTTON_FIXED
            : PRIMARY_BUTTON_FIXED} w-full justify-center px-6 md:w-auto"
          href={battlenetStartUrl('/account')}
          data-testid="character-not-found-signin-button"
        >
          {scoreCopy.notFoundSignIn}
        </a>
      {/if}
    </div>
  {/if}
  <form
    use:wire
    class="grid grid-cols-2 gap-3 md:grid-cols-[minmax(0,1fr)_112px_160px_auto] md:items-end"
    data-character-search
    data-status-id="character-not-found-error"
    data-empty={lookup.empty}
    data-invalid={lookup.invalid}
    novalidate
    data-testid="character-not-found-form"
  >
    <div class="col-span-2 flex flex-col gap-2 md:col-span-1">
      <label class="quiet-label" for="character-not-found-name">{scoreCopy.notFoundLookupLabel}</label>
      <input
        id="character-not-found-name"
        name="name"
        type="search"
        autocomplete="off"
        spellcheck="false"
        placeholder={lookup.namePlaceholder}
        class="{FIELD} h-11 md:h-14"
      />
    </div>
    <div class="flex flex-col gap-2">
      <label class="quiet-label" for="character-not-found-region">{lookup.regionLabel}</label>
      <Select
        id="character-not-found-region"
        name="region"
        class="max-md:[&_.sel]:[--sel-h:44px]"
        value={path.region}
        options={REGIONS.map((region) => ({ value: region, label: region.toUpperCase() }))}
      />
    </div>
    <div class="flex flex-col gap-2">
      <label class="quiet-label" for="character-not-found-ruleset">{lookup.rulesetLabel}</label>
      <Select
        id="character-not-found-ruleset"
        name="ruleset"
        class="max-md:[&_.sel]:[--sel-h:44px]"
        value={path.ruleset}
        options={RULESETS.map((ruleset) => ({ value: ruleset.id, label: ruleset.label }))}
      />
    </div>
    <button
      type="submit"
      class="{PRIMARY_BUTTON_FIXED} col-span-2 w-full justify-center px-6 md:col-span-1 md:h-14 md:w-auto"
    >
      {lookup.submit}
    </button>
  </form>
  <p id="character-not-found-error" class="text-strong min-h-[1lh] text-[13px]" role="status"></p>
  {#if known}
    <p class="text-muted text-[14px]" data-testid="character-not-found-next">
      <a class={LINK} href={scoreCopy.noLogsNext.addonHref}>{scoreCopy.noDataAddon}</a>
    </p>
  {:else}
    <p class="text-muted text-[14px]" data-testid="character-not-found-next">
      {scoreCopy.noLogsNext.lead}
      <a class={LINK} href={scoreCopy.noLogsNext.addonHref}>{scoreCopy.noLogsNext.addon}</a
      >{scoreCopy.separator}<a class={LINK} href={scoreCopy.noLogsNext.uploadHref}
        >{scoreCopy.noLogsNext.upload}</a
      >
    </p>
  {/if}
</div>
