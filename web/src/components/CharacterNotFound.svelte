<!-- web/src/components/CharacterNotFound.svelte -->
<!-- The character page for a name the site has no record of, shown to a visitor (W.8): one
     plain sentence instead of an empty character card with a score of nothing, the lookup form
     to try another name, and the way to get a character onto the site. The browser title says
     the same, so a tab or a bookmark never claims the character exists. -->
<script lang="ts">
  import { wireSearchForm } from '../lib/character-search-dom';
  import { REGIONS, RULESETS, type CharacterPath } from '../lib/characters';
  import { characterHeaderCopy } from '../lib/character-profile';
  import { homeCharacterSearchCopy as lookup } from '../lib/home-landing-copy';
  import { scoreCopy } from '../lib/score/copy';
  import { PRIMARY_BUTTON_FIXED } from '../lib/planner/styles';

  /** `known`: the site holds this character but has nothing on it, and the viewer does not hold it
   *  (P.1): its own sentence, and only a quiet "Get the addon" link under the lookup box. */
  let { path, name, known = false }: { path: CharacterPath; name: string; known?: boolean } = $props();

  const FIELD = 'border-line-warm-strong bg-raised rounded-control text-text w-full border px-3 text-[15px]';
  const LINK = 'text-nav underline';

  $effect(() => {
    document.title = known ? `${name} · Forever Sixty` : scoreCopy.notFoundTitle;
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
      <select id="character-not-found-region" name="region" class="{FIELD} h-11 md:h-14">
        {#each REGIONS as region (region)}
          <option value={region} selected={region === path.region}>
            {region.toUpperCase()}
          </option>
        {/each}
      </select>
    </div>
    <div class="flex flex-col gap-2">
      <label class="quiet-label" for="character-not-found-ruleset">{lookup.rulesetLabel}</label>
      <select id="character-not-found-ruleset" name="ruleset" class="{FIELD} h-11 md:h-14">
        {#each RULESETS as ruleset (ruleset.id)}
          <option value={ruleset.id} selected={ruleset.id === path.ruleset}>
            {ruleset.label}
          </option>
        {/each}
      </select>
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
