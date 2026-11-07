<!-- web/src/components/account/AccountYou.svelte -->
<!-- /account's You section in the rail (premium spec 2026-11-08 section 7.7): the pseudonym Switch
     and what it does, then how to delete the account (the privacy page's own route, by email, until
     a self-serve delete exists: gap G6). -->
<script lang="ts">
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Switch from '../ui/Switch.svelte';

  let {
    anonymize,
    busy,
    error = '',
    onAnonymize,
  }: {
    anonymize: boolean;
    busy: boolean;
    /** The switch's save failed. */
    error?: string;
    onAnonymize: (wanted: boolean) => void;
  } = $props();

  const labelId = 'account-pseudonym-label';
</script>

<Section title={accountPageCopy.youLabel} anchor="you" testid="account-you">
  <div class="flex flex-col gap-3">
    <div class="flex min-h-11 items-center justify-between gap-3">
      <span id={labelId} class="text-[15px] leading-5">{accountPageCopy.pseudonymLabel}</span>
      <Switch
        checked={anonymize}
        labelledby={labelId}
        disabled={busy}
        onchange={onAnonymize}
        testid="anonymize"
      />
    </div>
    <p class="text-small text-muted">{accountPageCopy.pseudonymNote}</p>
    {#if error !== ''}<p class="text-wipe text-[13px]" role="alert" data-testid="account-you-error">
        {error}
      </p>{/if}
    <div>
      <Row
        title={accountPageCopy.deleteTitle}
        sub={accountPageCopy.deleteBody}
        open
        testid="account-delete"
      />
    </div>
  </div>
</Section>
