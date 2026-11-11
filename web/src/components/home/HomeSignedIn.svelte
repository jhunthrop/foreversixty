<!-- web/src/components/home/HomeSignedIn.svelte -->
<!-- The signed-in home (home spec 2026-11-06 §4, round 5 2026-11-09): one island over the session read the header
     uses. With the session hint on (html[data-session='1']) the server-rendered signed-out page is
     hidden before paint and this island's skeleton holds the place; once /v1/me answers it draws
     the selected character, or the no-character page, or, signed out after all (the hint was
     stale, or the read failed), sets HOME_OUT_ATTR so the signed-out page shows again. -->
<script lang="ts">
  import { createCharacterCardState } from '../../lib/bis/character-card-state.svelte';
  import { HOME_OUT_ATTR } from '../../lib/home/home-layout';
  import { homeCopy, homeHrefs } from '../../lib/home/home-view';
  import { feedCopy, type FeedRow } from '../../lib/home/feed';
  import type { RaidFact } from '../../lib/home/raid-facts';
  import Button from '../ui/Button.svelte';
  import ClassPicker from '../ui/ClassPicker.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Row from '../ui/Row.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';
  import HomeCharacter from './HomeCharacter.svelte';
  import HomeGrid from './HomeGrid.svelte';
  import HomePlaqueSkeleton from './HomePlaqueSkeleton.svelte';
  import FeedSkeleton from '../ui/FeedSkeleton.svelte';

  let { floor, raidsOpen, raids }: { floor: FeedRow[]; raidsOpen: string | null; raids: RaidFact[] } =
    $props();

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
    <div data-testid="home-loading" role="status">
      <PageHeader title={homeCopy.loadingTitle} art="hero" pending secondHead testid="home-head-loading">
        {#snippet crest()}<span class="skeleton-block block size-12 rounded-full" aria-hidden="true"
          ></span>{/snippet}
      </PageHeader>
      <HomeGrid>
        {#snippet lead()}<HomePlaqueSkeleton />{/snippet}
        <Section title={feedCopy.title} boxed>
          <FeedSkeleton />
        </Section>
      </HomeGrid>
    </div>
  {:else if me !== null && session.character !== null}
    {#key session.character.key}
      <HomeCharacter {me} character={session.character} {floor} {raidsOpen} {raids} />
    {/key}
  {:else if me !== null}
    <Page layout="ledger" testid="home-no-character">
      {#snippet header()}<PageHeader title={homeCopy.getIn} art="hero" />{/snippet}
      <Plate flush={false} label={homeCopy.getIn} testid="home-get-in">
        <Row title={homeCopy.noCharacterYet} sub={homeCopy.checkAgainSub} open>
          {#snippet trail()}
            <Button tone="primary" size="sm" onclick={() => session.refresh()} testid="home-check-again"
              >{homeCopy.checkAgain}</Button
            >
          {/snippet}
        </Row>
        <Row title={homeCopy.paste} open>
          {#snippet trail()}
            <Button size="sm" href={homeHrefs.paste} testid="home-paste">{homeCopy.pasteAction}</Button>
          {/snippet}
        </Row>
        <Row title={homeCopy.getAddon} open>
          {#snippet trail()}
            <Button size="sm" href={homeHrefs.addon} testid="home-setup">{homeCopy.setupAction}</Button>
          {/snippet}
        </Row>
      </Plate>
      <Section title={homeCopy.pickClass} boxed testid="home-class-picker"
        ><ClassPicker testidPrefix="home-class-" /></Section
      >
    </Page>
  {/if}
</div>

<style>
  .home-head-skeleton {
    padding-inline: var(--page-inset);
  }
</style>
