<!-- web/src/components/account/AccountLoading.svelte -->
<!-- /account's loading state (premium spec 2026-11-08 section 7A): the ready page's head (the title
     alone, as ready) and split of panels with skeletons at the final geometry. The two columns carry the shipped reserve (spec
     2026-10-25-discord-notifications 9.1): global.css's "account loading reserve" block sizes
     `.account-left-skeleton` (main) and `.account-rail-skeleton` (the rail) from the heights
     Base.astro reads before paint, so a returning visitor's loading page has the ready page's
     shape. A visitor with no session hint can only resolve to the sign-in page, so that page is
     what the static render shows them (AccountOut), and the island draws the same. The Battle.net
     line in the Characters panel is static copy, drawn here as on the ready page, so the ready
     page's largest text is painted in the first frame instead of after /v1/me (LCP). -->
<script lang="ts">
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import { notificationsCopy } from '../../lib/account/notifications-copy';
  import Page from '../ui/Page.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import AccountOut from './AccountOut.svelte';
  import CharacterListIntro from './CharacterListIntro.svelte';

  const RAIL = [accountPageCopy.youLabel, accountPageCopy.devicesLabel, notificationsCopy.label];
</script>

<div class="account-when-out"><AccountOut /></div>
<div class="account-when-in">
  <Page layout="split" panels testid="account-loading">
    {#snippet header()}<PageHeader title={accountPageCopy.title} testid="account-head" />{/snippet}
    <div class="account-loading session-reserve account-left-skeleton" data-testid="account-left-skeleton">
      <Section title={characterListCopy.heading} panel frame="framed">
        <Skeleton lines={2} rowHeight="h-[56px]" testid="account-characters-skeleton" />
        <CharacterListIntro />
      </Section>
      <Section title={accountPageCopy.planLabel} panel>
        <Skeleton lines={4} rowHeight="h-5" testid="account-plan-skeleton" />
      </Section>
    </div>
    {#snippet rail()}
      <div class="account-loading session-reserve account-rail-skeleton" data-testid="account-rail-skeleton">
        {#each RAIL as title (title)}
          <Section {title} panel>
            <Skeleton lines={3} rowHeight="h-5" testid="account-{title.toLowerCase()}-skeleton" />
          </Section>
        {/each}
      </div>
    {/snippet}
  </Page>
</div>

<style>
  .account-loading {
    display: flex;
    flex-direction: column;
    gap: var(--section-gap);
  }
  :global(html:not([data-session='1'])) .account-when-in,
  :global(html[data-session='1']) .account-when-out {
    display: none;
  }
</style>
