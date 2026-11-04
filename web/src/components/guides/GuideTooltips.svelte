<!-- web/src/components/guides/GuideTooltips.svelte -->
<!-- The spec guide's one tooltip island: the site's shared host (BisTooltipHost.svelte, the
     same hover, focus, tap and Escape rules as /bis and the character pages) over this guide's
     tooltip bundle (lib/guides/guide-tips.ts). The bundle is fetched once, the first time a
     reader points at, focuses or taps a talent, spell, racial, enchant or item, so the page
     carries no tooltip data of its own. -->
<script lang="ts">
  import BisTooltipHost from '../BisTooltipHost.svelte';
  import type { ItemTooltipModel } from '../../lib/items/tooltip';
  import { tipFromBundle, type GameTipModel, type GuideTipBundle } from '../../lib/tooltips/game-tip';

  let { build, bundleUrl }: { build: string; bundleUrl: string } = $props();

  let bundle: Promise<GuideTipBundle | null> | null = null;

  function loadBundle(): Promise<GuideTipBundle | null> {
    bundle ??= fetch(bundleUrl)
      .then((response) => (response.ok ? (response.json() as Promise<GuideTipBundle>) : null))
      .catch(() => {
        bundle = null; // a failed fetch is retried on the next open
        return null;
      });
    return bundle;
  }

  async function resolveItem(itemId: number): Promise<ItemTooltipModel | null> {
    return (await loadBundle())?.items[String(itemId)] ?? null;
  }

  async function resolveTip(key: string): Promise<GameTipModel | null> {
    const loaded = await loadBundle();
    return loaded === null ? null : tipFromBundle(loaded, key);
  }
</script>

<BisTooltipHost {build} modelsElementId="guide-tooltip-models" resolve={resolveItem} {resolveTip} />
