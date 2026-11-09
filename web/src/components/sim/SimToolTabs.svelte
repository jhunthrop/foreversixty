<!-- web/src/components/sim/SimToolTabs.svelte -->
<!-- The simulator's tool tabs in the page head (spec 2026-11-06-sim §4.1): Quick sim, Top gear,
     Droptimizer, Talents, Weights, Spec support, as route links (ui/TabStrip without onSelect is a
     nav of links). Each tab carries the loaded character (lib/sim/tabs.ts `tabStateFor`), so a tab
     opens the same character in that tool. Not on /sim/<id> (ruling S5). -->
<script lang="ts">
  import {
    SIM_TAB_ATTR,
    SIM_TABS,
    tabHref,
    tabStateFor,
    type SimTabId,
    type TabSource,
  } from '../../lib/sim/tabs';
  import TabStrip from '../ui/TabStrip.svelte';

  let {
    active = 'quick-sim',
    source = null,
    fallbackCode = null,
    keep = false,
  }: {
    /** The tab this page is. */
    active?: SimTabId;
    /** The loaded character's source, or null with none. */
    source?: TabSource | null;
    /** An FS1 code for a source with no ref (a paste). */
    fallbackCode?: string | null;
    /** The page views a character (`keep=1`): every tab keeps the reader's header character. */
    keep?: boolean;
  } = $props();

  const tabs = $derived(
    SIM_TABS.map((tab) => ({
      id: tab.id,
      label: tab.label,
      href: tabHref(tab.href, tabStateFor(source, tab.supportsCode ? fallbackCode : null, keep)),
      testid: `sim-nav-tab-${tab.id}`,
      attrs: { [SIM_TAB_ATTR]: tab.id },
    })),
  );
</script>

<TabStrip {tabs} {active} label="Simulator tools" testid="sim-nav-tabs" wrap />
