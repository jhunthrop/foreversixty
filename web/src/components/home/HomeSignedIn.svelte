<!-- web/src/components/home/HomeSignedIn.svelte -->
<!-- The signed-in home (home spec 2026-11-06 §4): one island over the session read the header
     uses. With the session hint on (html[data-session='1']) the server-rendered signed-out page is
     hidden before paint and this island's skeleton holds the place; once /v1/me answers it draws
     the selected character, or the no-character page, or, signed out after all (the hint was
     stale, or the read failed), sets HOME_OUT_ATTR so the signed-out page shows again. -->
<script lang="ts">
  import { createCharacterCardState } from '../../lib/bis/character-card-state.svelte';
  import { HOME_OUT_ATTR } from '../../lib/home/home-layout';
  import { homeCopy, homeHrefs } from '../../lib/home/home-view';
  import { REGION } from '../../lib/ui/regions';
  import Button from '../ui/Button.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import HomeCharacter from './HomeCharacter.svelte';
  import HomeClassPicker from './HomeClassPicker.svelte';
  import HomeSummary from './HomeSummary.svelte';
  import HomeUpgrades from './HomeUpgrades.svelte';

  const session = createCharacterCardState();
  const settled = $derived(session.status === 'ready' || session.status === 'failed');
  const me = $derived(session.me);
  const signedOut = $derived(settled && me === null);

  $effect(() => {
    document.documentElement.toggleAttribute(HOME_OUT_ATTR, signedOut);
  });
</script>

<div class="home-in" data-testid="home-signed-in-root">
  {#if !settled}
    <Page layout="split" testid="home-loading">
      {#snippet header()}
        <div
          data-region={REGION.head}
          class="home-head-skeleton flex min-h-[118px] flex-col justify-end gap-3 pt-5 pb-4"
          role="status"
          aria-busy="true"
        >
          <span class="skeleton-block block size-14 rounded-full" aria-hidden="true"></span>
          <span class="skeleton-block block h-7 w-[280px]" aria-hidden="true"></span>
        </div>
      {/snippet}
      {#snippet bar()}<HomeSummary loading />{/snippet}
      <HomeUpgrades status="loading" upgrades={null} bis={null} build="" onRetry={() => {}} />
    </Page>
  {:else if me !== null && session.character !== null}
    {#key session.character.key}
      <HomeCharacter {me} character={session.character} />
    {/key}
  {:else if me !== null}
    <Page layout="ledger" testid="home-no-character">
      {#snippet header()}
        <PageHeader title={homeCopy.signedOutTitle}>
          {#snippet meta()}<span>{homeCopy.signedInAs(me.user.battletag ?? '')}</span>{/snippet}
        </PageHeader>
      {/snippet}
      {#snippet bar()}<HomeSummary text={homeCopy.noCharacter} />{/snippet}
      <Section title={homeCopy.getIn} testid="home-get-in">
        <Row title={homeCopy.checkAgain} sub={homeCopy.checkAgainSub} open>
          {#snippet trail()}
            <Button tone="primary" size="sm" onclick={() => session.refresh()} testid="home-check-again"
              >{homeCopy.checkAgain}</Button
            >
          {/snippet}
        </Row>
        <Row title={homeCopy.paste} sub={homeCopy.pasteSub} href={homeHrefs.paste} open />
        <Row title={homeCopy.getAddon} sub={homeCopy.getAddonSub} href={homeHrefs.addon} open />
      </Section>
      <HomeClassPicker />
    </Page>
  {/if}
</div>

<style>
  .home-head-skeleton {
    padding-inline: var(--page-inset);
  }
</style>
