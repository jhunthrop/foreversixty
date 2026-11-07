<!-- web/src/components/account/AccountLoading.svelte -->
<!-- /account's loading state (premium spec 2026-11-08 section 7A): the ready page's head and split
     with skeletons at the final geometry. The two columns carry the shipped reserve (spec
     2026-10-25-discord-notifications 9.1): global.css's "account loading reserve" block sizes
     `.account-left-skeleton` (main) and `.account-rail-skeleton` (the rail) from the heights
     Base.astro reads before paint, so a returning visitor's loading page has the ready page's
     shape. A visitor with no session hint can only resolve to the sign-in page, so that page is
     what the static render shows them (AccountOut), and the island draws the same. -->
<script lang="ts">
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import { notificationsCopy } from '../../lib/account/notifications-copy';
  import Page from '../ui/Page.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import AccountHead from './AccountHead.svelte';
  import AccountOut from './AccountOut.svelte';

  const RAIL = [accountPageCopy.devicesLabel, notificationsCopy.label, accountPageCopy.youLabel];
</script>

<div class="account-when-out"><AccountOut /></div>
<div class="account-when-in">
  <Page layout="split" testid="account-loading">
    {#snippet header()}<AccountHead />{/snippet}
    <div class="account-loading session-reserve account-left-skeleton" data-testid="account-left-skeleton">
      <Section title={characterListCopy.heading}>
        <Plate>
          <div class="px-4 py-3">
            <Skeleton lines={2} rowHeight="h-[56px]" testid="account-characters-skeleton" />
          </div>
        </Plate>
      </Section>
      <Section title={accountPageCopy.planLabel}>
        <Skeleton lines={4} rowHeight="h-5" testid="account-plan-skeleton" />
      </Section>
    </div>
    {#snippet rail()}
      <div class="account-loading session-reserve account-rail-skeleton" data-testid="account-rail-skeleton">
        {#each RAIL as title (title)}
          <Section {title}>
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
