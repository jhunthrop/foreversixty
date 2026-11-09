<!-- web/src/components/app-beta/StoreButtons.svelte -->
<!-- /app once src/data/apps.json has a store link: one button per live store, and a phone whose
     store is live is sent straight there (the QR's job then: scan, land in the store). A desktop,
     or a phone whose store is not live yet, stays on the page with the buttons. -->
<script lang="ts">
  import { appBetaCopy as copy } from '../../lib/app-beta/copy';
  import { phoneOS, storeFor, type StoreLink } from '../../lib/app-beta/platform';
  import { button } from '../../lib/ui/button';

  let { stores }: { stores: readonly StoreLink[] } = $props();

  $effect(() => {
    const store = storeFor(stores, phoneOS(navigator.userAgent, navigator.maxTouchPoints));
    if (store !== null) window.location.replace(store.url);
  });
</script>

<div class="flex flex-wrap items-center gap-3" data-testid="app-beta-stores">
  {#each stores as store, index (store.platform)}
    <a
      class={button(index === 0 ? 'primary' : 'secondary')}
      href={store.url}
      rel="noopener"
      data-testid={`app-beta-store-${store.platform}`}>{copy.stores[store.platform]}</a
    >
  {/each}
</div>
