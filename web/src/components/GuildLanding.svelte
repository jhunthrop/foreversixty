<!-- web/src/components/GuildLanding.svelte -->
<!-- The bare `/guild` page: where the header's Guild door lands a visitor with no guild of
     their own (spec 2026-10-15-grouped-header.md section 2). A lookup by name, the top
     guilds, and a line for officers. A signed-in member with a verified guild is sent on to
     it once `/v1/me` answers; the lookup is already on screen meanwhile. -->
<script lang="ts">
  import { GUILD_BANDLESS_TOP } from '../lib/guild/raid-layout';
  import { onMount } from 'svelte';
  import { DEFAULT_REGION, DEFAULT_RULESET } from '../lib/character-search';
  import { fetchMeOnce } from '../lib/account/api';
  import { sessionGuildOf } from '../lib/character-selector/session-guild';
  import { guildHref, REGIONS, RULESETS } from '../lib/characters';
  import { guildLandingCopy as copy } from '../lib/guild/copy';
  import { PRIMARY_BUTTON_FIXED } from '../lib/planner/styles';
  import HomeTopGuilds from './HomeTopGuilds.svelte';

  const FIELD =
    'border-line-warm-strong bg-raised rounded-control text-text h-11 w-full border px-3 text-[15px]';

  let name = $state('');
  let region = $state<string>(DEFAULT_REGION);
  let ruleset = $state<string>(DEFAULT_RULESET);
  let status = $state('');

  function onSubmit(event: SubmitEvent): void {
    event.preventDefault();
    const trimmed = name.trim();
    if (trimmed === '') {
      status = copy.emptyName;
      return;
    }
    status = '';
    window.location.assign(guildHref(region, ruleset, trimmed));
  }

  onMount(() => {
    void fetchMeOnce().then((me) => {
      const guild = sessionGuildOf(me);
      if (guild !== null) window.location.replace(guild.href);
    });
  });
</script>

<div class="reveal {GUILD_BANDLESS_TOP} flex flex-col gap-[22px] md:gap-8" data-testid="guild-landing">
  <header class="flex flex-col gap-2">
    <h1 class="heading-page">{copy.heading}</h1>
    <p class="text-muted text-[15px]">{copy.lead}</p>
  </header>

  <form
    role="search"
    class="grid grid-cols-2 items-end gap-3 md:grid-cols-[minmax(0,1fr)_auto_auto_auto]"
    data-testid="guild-lookup"
    onsubmit={onSubmit}
  >
    <div class="col-span-2 flex flex-col gap-2 md:col-span-1">
      <label class="quiet-label" for="guild-lookup-name">{copy.nameLabel}</label>
      <input
        id="guild-lookup-name"
        name="name"
        type="text"
        autocomplete="off"
        spellcheck="false"
        bind:value={name}
        aria-describedby="guild-lookup-status"
        class={FIELD}
      />
    </div>
    <div class="flex flex-col gap-2">
      <label class="quiet-label" for="guild-lookup-region">{copy.regionLabel}</label>
      <select id="guild-lookup-region" name="region" bind:value={region} class={FIELD}>
        {#each REGIONS as option (option)}
          <option value={option}>{option.toUpperCase()}</option>
        {/each}
      </select>
    </div>
    <div class="flex flex-col gap-2">
      <label class="quiet-label" for="guild-lookup-ruleset">{copy.rulesetLabel}</label>
      <select id="guild-lookup-ruleset" name="ruleset" bind:value={ruleset} class={FIELD}>
        {#each RULESETS as option (option.id)}
          <option value={option.id}>{option.label}</option>
        {/each}
      </select>
    </div>
    <button
      type="submit"
      class={`${PRIMARY_BUTTON_FIXED} col-span-2 w-full justify-center px-6 md:col-span-1 md:w-auto`}
    >
      {copy.submit}
    </button>
  </form>
  <p id="guild-lookup-status" class="text-strong min-h-[1lh] text-[13px]" role="status">{status}</p>

  <section class="flex flex-col gap-3" aria-labelledby="guild-landing-top">
    <h2 id="guild-landing-top" class="label text-gold">{copy.topHeading}</h2>
    <HomeTopGuilds />
  </section>

  <p id="claim" class="text-muted scroll-mt-24 text-[13px]" data-testid="guild-landing-claim">
    {copy.officerLine}
  </p>
</div>
