<!-- web/src/components/EmailLinkAlternative.svelte -->
<!-- The email-link way in, said once for every sign-in surface (the sim's account card, the
     guild page's and the logs page's sign-in prompts): "or use an email link", the last words a
     link, placed straight after the Battle.net button it is an alternative to. A sentence of its
     own ("Use an email link instead.") read as dangling, with nothing before it to be instead of. -->
<script lang="ts">
  import { emailLinkCopy } from '../lib/account/signin-copy';

  import ActionLink from './ui/ActionLink.svelte';

  let {
    testid = 'email-link-alternative',
    action = false,
  }: {
    testid?: string;
    /** True where the link is a standalone action (the sim's account card): an `ActionLink`
     *  after the muted "or", rather than an underlined link in the sentence. */
    action?: boolean;
  } = $props();
</script>

{#if action}
  <p class="text-muted flex items-center gap-1 text-[13px]" data-testid={testid}>
    {emailLinkCopy.lead}
    <ActionLink href={emailLinkCopy.href}>{emailLinkCopy.link}</ActionLink>
  </p>
{:else}
  <p class="text-muted text-[13px]" data-testid={testid}>
    {emailLinkCopy.lead}
    <a class="text-nav underline" href={emailLinkCopy.href}>{emailLinkCopy.link}</a>
  </p>
{/if}
