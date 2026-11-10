<!-- web/src/components/GuildLanding.svelte -->
<!-- The bare `/guild` page (web layout spec 2026-11-05 §6.1 Landing, §7.1 signed-in landing): where
     the header's Guild door lands a visitor with no guild of their own. PageHeader, then a split
     Page: Open a guild (a FormRow) and Top guilds (the one Plate); the rail holds Your guilds when
     the account has any, then Run a guild? (with the sign-in button signed out) or the claim the
     account holds. A signed-in visitor whose current character has a guild, or whose account has
     exactly one, is sent on to it once `/v1/me` answers; the lookup is already on screen meanwhile,
     and starts on the account's first guild's region and ruleset (guild round 3). -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { battlenetStartUrl, type Me } from '../lib/account/api';
  import { DEFAULT_REGION, DEFAULT_RULESET } from '../lib/character-search';
  import { guildHref, REGIONS, RULESETS } from '../lib/characters';
  import { guildLandingCopy as copy } from '../lib/guild/copy';
  import { heldClaimRow, lookupDefaults, myGuildRows, type MyGuildRow } from '../lib/guild/landing-mine';
  import { FIELD_LABEL } from '../lib/ui/field';
  import Button from './ui/Button.svelte';
  import Field from './ui/Field.svelte';
  import FormRow from './ui/FormRow.svelte';
  import Page from './ui/Page.svelte';
  import PageHeader from './ui/PageHeader.svelte';
  import Section from './ui/Section.svelte';
  import Select from './ui/Select.svelte';
  import HomeTopGuilds from './HomeTopGuilds.svelte';
  import GuildLandingMine from './guild/GuildLandingMine.svelte';

  /** The session, when a test hands one in; the page reads its own on mount. */
  let { me: givenMe = null }: { me?: Me | null } = $props();

  const regionOptions = REGIONS.map((option) => ({ value: option, label: option.toUpperCase() }));
  const rulesetOptions = RULESETS.map((option) => ({ value: option.id, label: option.label }));

  let name = $state('');
  let region = $state<string>(DEFAULT_REGION);
  let ruleset = $state<string>(DEFAULT_RULESET);
  let status = $state('');
  let mine = $state<MyGuildRow[]>(myGuildRows(givenMe));
  /** The guild whose claim the account holds: the rail says so instead of asking. */
  let held = $state<MyGuildRow | null>(heldClaimRow(givenMe));
  /** Null until `/v1/me` answers (the sign-in button waits for it), then whether there is a session. */
  let signedIn = $state<boolean | null>(givenMe === null ? null : true);
  /** The player changed a select: the session's defaults never overwrite their choice. */
  let placeTouched = false;

  function applySession(me: Me | null): void {
    signedIn = me !== null;
    mine = myGuildRows(me);
    held = heldClaimRow(me);
    const defaults = lookupDefaults(me);
    if (defaults === null || placeTouched) return;
    region = defaults.region;
    ruleset = defaults.ruleset;
  }

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

  // The current character's guild, when it has one, is this landing's answer (GuildShell follows
  // later choices of character the same way).
  /** The lookup sent before this island hydrated (or with no script): the server-rendered form
   *  submits /guild?name=&region=&ruleset= natively, and the landing finishes the job here. */
  function finishNativeLookup(): boolean {
    const params = new URLSearchParams(window.location.search);
    if (!params.has('name')) return false;
    const typed = (params.get('name') ?? '').trim();
    const pickedRegion = params.get('region') ?? '';
    const pickedRuleset = params.get('ruleset') ?? '';
    const knownRegion = (REGIONS as readonly string[]).includes(pickedRegion) ? pickedRegion : region;
    const knownRuleset = RULESETS.some((option) => option.id === pickedRuleset) ? pickedRuleset : ruleset;
    if (typed === '') {
      window.history.replaceState(null, '', window.location.pathname);
      status = copy.emptyName;
      return false;
    }
    window.location.replace(guildHref(knownRegion, knownRuleset, typed));
    return true;
  }

  onMount(() => {
    if (finishNativeLookup()) return;
    void import('../lib/character-selector/follow-page-run')
      .then(({ openCurrentGuild }) => openCurrentGuild())
      .then(applySession);
  });

  const next = typeof location === 'undefined' ? '/guild' : location.pathname;
</script>

<div data-testid="guild-landing">
  {#snippet header()}
    <PageHeader title={copy.heading} testid="guild-landing-head" />
  {/snippet}
  {#snippet rail()}
    <GuildLandingMine rows={mine} />
    <Section
      title={held === null ? copy.runHeading : copy.heldHeading}
      panel
      anchor="claim"
      testid="guild-landing-claim"
    >
      <div class="flex flex-col items-start gap-3">
        <p class="text-text text-[14px] leading-[22px]" data-testid="guild-landing-claim-line">
          {held === null ? copy.officerLine : copy.holdsClaim(held.name)}
        </p>
        {#if signedIn === false}
          <Button href={battlenetStartUrl(next)} testid="guild-landing-sign-in">{copy.signIn}</Button>
        {/if}
      </div>
    </Section>
  {/snippet}
  <Page layout="split" panels {header} {rail} testid="guild-landing-page">
    <Section title={copy.lookupHeading} panel frame="framed" testid="guild-lookup-card">
      <FormRow label={copy.lookupHeading} onsubmit={onSubmit} testid="guild-lookup">
        <Field
          id="guild-lookup-name"
          name="name"
          label={copy.nameLabel}
          placeholder={copy.namePlaceholder}
          autocomplete="off"
          bind:value={name}
          grow
        />
        <!-- One child of the FormRow, so a phone keeps region and ruleset side by side under the name. -->
        <div class="flex gap-3 [&>*]:flex-1 md:[&>*]:flex-none">
          <div class="flex flex-col gap-[6px]">
            <label class={FIELD_LABEL} for="guild-lookup-region">{copy.regionLabel}</label>
            <Select
              id="guild-lookup-region"
              name="region"
              size="compact"
              bind:value={region}
              onchange={() => (placeTouched = true)}
              options={regionOptions}
            />
          </div>
          <div class="flex flex-col gap-[6px]">
            <label class={FIELD_LABEL} for="guild-lookup-ruleset">{copy.rulesetLabel}</label>
            <Select
              id="guild-lookup-ruleset"
              name="ruleset"
              size="compact"
              bind:value={ruleset}
              onchange={() => (placeTouched = true)}
              options={rulesetOptions}
            />
          </div>
        </div>
        {#snippet submit()}
          <Button tone="primary" type="submit">{copy.submit}</Button>
        {/snippet}
      </FormRow>
      <p class="text-strong min-h-[1lh] pt-2 text-[13px] empty:hidden" role="status">{status}</p>
    </Section>
    <Section title={copy.topHeading} panel testid="guild-landing-top">
      <HomeTopGuilds label={copy.topHeading} />
    </Section>
  </Page>
</div>
