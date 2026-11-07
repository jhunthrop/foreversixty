<!-- web/src/components/SignInPrompt.svelte -->
<!-- What a signed-out visitor sees wherever a page needs an account (the guild home, raid, claim,
     recruit and member pages, the go-links admin): one sentence saying what signing in is for, and
     the Battle.net button under it. No frame of its own: it sits inside whatever Section or Plate
     the page already has, so it never adds a second Plate (web layout spec 2026-11-05 §3A.2). -->
<script lang="ts">
  import { battlenetStartUrl } from '../lib/account/api';
  import { myReportsCopy } from '../lib/reports/my-reports-copy';
  import Button from './ui/Button.svelte';

  let {
    line,
    next = '/account?signed_in=1',
    testid,
  }: {
    line: string;
    /** Where Battle.net sends the visitor back to (the page they were on). */
    next?: string;
    testid: string;
  } = $props();
</script>

<div class="flex flex-col items-start gap-3" data-testid={testid}>
  <p class="max-w-[72ch] text-[15px] leading-6">{line}</p>
  <Button href={battlenetStartUrl(next)}>{myReportsCopy.signInAction}</Button>
</div>
