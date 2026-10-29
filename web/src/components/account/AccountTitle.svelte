<!-- web/src/components/account/AccountTitle.svelte -->
<!-- The account page's eyebrow and h1, one component so the loading, failed, signed-out and ready
     states draw the same box (the loading reserve pins its height). -->
<script lang="ts">
  import { accountPageCopy } from '../../lib/account/account-page-copy';
</script>

<div class="flex flex-col gap-2">
  <p class="text-gold flex items-center gap-3 text-[11px] leading-none font-bold tracking-[.18em] uppercase">
    <span class="bg-gold/60 h-px w-6" aria-hidden="true"></span>{accountPageCopy.eyebrow}
  </p>
  <h1 class="heading-page">{accountPageCopy.title}</h1>
</div>
