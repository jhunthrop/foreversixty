<!-- web/src/components/account/AccountHead.svelte -->
<!-- /account's PageHeader (premium spec 2026-11-08 section 7.1): `Your account`, the meta line (the
     BattleTag, how the account signed in, and Sign out, which asks first), and the answer: the
     account's main character as /setup draws it (setupAnswer: crest, name, level line, Open and Best
     in slot). It reads the main, never the header pointer, and never switches it. With no `me` (the
     loading state) the meta and the answer hold their ready boxes. -->
<script lang="ts">
  import { smallButton } from '../../lib/account/small-button';
  import type { Me } from '../../lib/account/api';
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import { devicesCopy } from '../../lib/account/devices-copy';
  import { setupAnswer } from '../../lib/setup/answer';
  import PageHeader from '../ui/PageHeader.svelte';
  import SetupAnswer from '../setup/SetupAnswer.svelte';
  import ConfirmButtons from './ConfirmButtons.svelte';

  let {
    me = null,
    now = new Date(),
    busy = false,
    onSignOut,
  }: {
    /** The signed-in account; null draws the title alone. */
    me?: Me | null;
    now?: Date;
    busy?: boolean;
    onSignOut?: () => void;
  } = $props();

  let asking = $state(false);
  const viaBattlenet = $derived(me?.user.battletag !== undefined && me?.user.battletag !== null);
  const who = $derived(me?.user.battletag ?? me?.user.email ?? '');
  // The answer is the main (no pointer passed): this page shows the account, not the header's pick.
  const answer = $derived.by(() => {
    const found = me === null ? null : setupAnswer(me, null, now);
    return found === null ? null : { ...found, imported: null };
  });
</script>

{#snippet meta()}
  <span data-testid="account-battletag">{who}</span>
  <span>{viaBattlenet ? accountPageCopy.signedInWithBattlenet : accountPageCopy.signedInByEmail}</span>
  {#if asking}
    <span class="text-text" data-testid="account-sign-out-question">{accountPageCopy.signOutAsk}</span>
    <span class="flex flex-wrap gap-[10px]">
      <ConfirmButtons
        confirm={accountPageCopy.signOut}
        cancel={devicesCopy.cancel}
        {busy}
        onconfirm={() => onSignOut?.()}
        oncancel={() => (asking = false)}
        testid="account-sign-out"
      />
    </span>
  {:else}
    <button
      type="button"
      class={smallButton()}
      disabled={busy}
      onclick={() => (asking = true)}
      data-testid="account-sign-out">{accountPageCopy.signOut}</button
    >
  {/if}
{/snippet}

{#snippet answerContent()}
  {#if answer !== null}<SetupAnswer {answer} />{/if}
{/snippet}

{#snippet loadingMeta()}
  <span class="account-meta-reserve skeleton-block" aria-hidden="true"></span>
{/snippet}

{#snippet answerReserve()}
  <span class="account-answer-reserve" aria-hidden="true"></span>
{/snippet}

{#if me === null}
  <PageHeader title={accountPageCopy.title} meta={loadingMeta} answer={answerReserve} testid="account-head" />
{:else}
  <PageHeader
    title={accountPageCopy.title}
    {meta}
    answer={answer === null ? undefined : answerContent}
    testid="account-head"
  />
{/if}

<style>
  /* The loading head holds the ready meta line's box: the Sign out button's height. */
  .account-meta-reserve {
    display: none;
    width: 220px;
    height: 44px;
  }
  :global(html[data-session='1']) .account-meta-reserve {
    display: block;
  }
  /* A hinted visitor with a stored character (an account with characters): the answer's box, held
     until /v1/me answers (the ready answer's measured size), as /setup holds it. */
  .account-answer-reserve {
    display: none;
  }
  /* Without the reserve the answer's column is not there at all (its 16 px gap would stack on a phone). */
  :global(header div.ui-page-header-answer:has(> span.account-answer-reserve)) {
    display: none;
  }
  :global(
    html[data-session='1'][data-pointer='1']
      header
      div.ui-page-header-answer:has(> span.account-answer-reserve)
  ) {
    display: flex;
  }
  :global(html[data-session='1'][data-pointer='1']) .account-answer-reserve {
    display: block;
    width: 300px;
    height: 72px;
  }
  @media (min-width: 1024px) and (pointer: fine) {
    .account-meta-reserve {
      height: 36px;
    }
  }
</style>
