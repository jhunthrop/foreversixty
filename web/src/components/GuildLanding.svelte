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
  import { button } from '../lib/ui/button';
  import Card from './ui/Card.svelte';
  import Glyph from './ui/Glyph.svelte';
  import PageHead from './ui/PageHead.svelte';
  import Select from './ui/Select.svelte';
  import HomeTopGuilds from './HomeTopGuilds.svelte';

  const regionOptions = REGIONS.map((option) => ({ value: option, label: option.toUpperCase() }));
  const rulesetOptions = RULESETS.map((option) => ({ value: option.id, label: option.label }));

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
  <PageHead eyebrow={copy.eyebrow} title={copy.heading} job={copy.lead} />

  <Card glyph="search" title={copy.lookupHeading} body="none" testid="guild-lookup-card">
    <form
      role="search"
      class="border-line-soft grid grid-cols-2 items-end gap-3 border-t px-5 py-4 md:grid-cols-[minmax(0,1fr)_auto_auto_auto]"
      data-testid="guild-lookup"
      onsubmit={onSubmit}
    >
      <div class="col-span-2 flex flex-col gap-2 md:col-span-1">
        <label class="quiet-label" for="guild-lookup-name">{copy.nameLabel}</label>
        <div class="relative">
          <Glyph
            name="search"
            size={16}
            class="text-muted pointer-events-none absolute top-1/2 left-3 -translate-y-1/2"
          />
          <input
            id="guild-lookup-name"
            name="name"
            type="text"
            autocomplete="off"
            spellcheck="false"
            placeholder={copy.namePlaceholder}
            bind:value={name}
            aria-describedby="guild-lookup-status"
            class="rounded-control border-line-card text-strong placeholder:text-muted hover:border-gold/30 focus-visible:outline-gold h-11 w-full border bg-white/[.045] pr-3 pl-9 text-[15px] font-semibold focus-visible:outline-2 focus-visible:outline-offset-2"
          />
        </div>
      </div>
      <div class="flex flex-col gap-2">
        <label class="quiet-label" for="guild-lookup-region">{copy.regionLabel}</label>
        <Select
          id="guild-lookup-region"
          name="region"
          size="compact"
          bind:value={region}
          options={regionOptions}
        />
      </div>
      <div class="flex flex-col gap-2">
        <label class="quiet-label" for="guild-lookup-ruleset">{copy.rulesetLabel}</label>
        <Select
          id="guild-lookup-ruleset"
          name="ruleset"
          size="compact"
          bind:value={ruleset}
          options={rulesetOptions}
        />
      </div>
      <button type="submit" class="{button('primary')} col-span-2 w-full md:col-span-1 md:w-auto">
        {copy.submit}
      </button>
    </form>
    <p
      id="guild-lookup-status"
      class="text-strong min-h-[1lh] px-5 pb-4 text-[13px] empty:hidden"
      role="status"
    >
      {status}
    </p>
  </Card>

  <Card glyph="trophy" title={copy.topHeading} body="none" testid="guild-landing-top">
    <div class="border-line-soft border-t px-5 py-2"><HomeTopGuilds /></div>
    {#snippet foot()}
      <p id="claim" class="text-muted scroll-mt-24 text-[13px]" data-testid="guild-landing-claim">
        {copy.officerLine}
      </p>
    {/snippet}
  </Card>
</div>
