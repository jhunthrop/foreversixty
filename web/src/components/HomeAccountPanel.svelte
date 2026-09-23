<!-- web/src/components/HomeAccountPanel.svelte -->
<!-- The home page's signed-in swap (spec 2026-09-22 §3.2): while this is loading or the
     visitor is signed out, it renders nothing and the Astro shell's own server-rendered
     sentence + button (index.astro) stays exactly where it is -- the same "server shell
     first, island swaps in place" trick SessionNav.svelte's header link uses, so the panel
     never shows two competing versions. Both this island's root and index.astro's
     signed-out block share `[grid-area:1/1]` in a shared grid wrapper, so once this mounts
     signed-in it visually occludes the signed-out row (a solid background over the same
     cell) instead of the two stacking and reflowing the page underneath. -->
<script lang="ts">
  import { fetchMeOnce, type Me } from '../lib/account/api';
  import { readCurrent } from '../lib/current-character';
  import { classColorVar } from '../lib/report/format';
  import { HOME_SIGNED_OUT_ID, homePanelCopy } from '../lib/home-panel-copy';

  let me = $state<Me | null>(null);
  let ready = $state(false);

  $effect(() => {
    void fetchMeOnce()
      .then((result) => {
        me = result;
        ready = true;
      })
      .catch(() => {
        ready = true;
      });
  });

  /**
   * The grid-overlay CLS trick (this component's root and index.astro's signed-out block
   * share the same `[grid-area:1/1]` cell) only ever covers the signed-out row visually --
   * it stays mounted underneath, so without this a signed-in keyboard/screen-reader user
   * could still tab to, or hear, a duplicate "Sign in with Battle.net" link sitting behind
   * the visible strip. Reaches outside this component's own root via `document`, the same
   * cross-island DOM-reach pattern `syncTabHrefs` in `lib/sim/tabs.ts` uses to coordinate
   * with a sibling shell element it doesn't own. `ready && me !== null` never reverts to
   * signed-out within one mount today (`fetchMeOnce` resolves once), but the else branch
   * clears both attributes anyway so this stays correct if that ever changes.
   */
  $effect(() => {
    const signedOut = document.getElementById(HOME_SIGNED_OUT_ID);
    if (signedOut === null) return;
    if (ready && me !== null) {
      signedOut.setAttribute('inert', '');
      signedOut.setAttribute('aria-hidden', 'true');
    } else {
      signedOut.removeAttribute('inert');
      signedOut.removeAttribute('aria-hidden');
    }
  });

  const pointer = $derived(readCurrentIfReady());
  function readCurrentIfReady() {
    if (!ready || me === null) return null;
    return readCurrent();
  }
  const current = $derived(
    pointer === null ? null : me?.characters.find((c) => c.key === pointer.ref || pointer.ref === ''),
  );
  const displayName = $derived(pointer?.label ?? me?.characters[0]?.name ?? '');
  const colour = $derived(classColorVar(current?.class));
</script>

{#if ready && me !== null}
  <div
    class="flex flex-wrap items-center gap-3 bg-[var(--color-bg)] [grid-area:1/1]"
    data-testid="home-account-panel"
  >
    <span class="rounded-control h-7 w-7 shrink-0" style={`background: ${colour}`} aria-hidden="true"></span>
    <span class="text-[15px] font-semibold" style={`color: ${colour}`}>{displayName}</span>
    <a class="text-nav text-[13px] font-semibold" href="/planner">{homePanelCopy.openInPlanner}</a>
    <a class="text-nav text-[13px] font-semibold" href="/sim">{homePanelCopy.openInSimulator}</a>
    <a class="text-nav text-[13px] font-semibold" href="/logs">{homePanelCopy.logs}</a>
    <a class="text-nav text-[13px] font-semibold" href="/account">{homePanelCopy.yourCharacters}</a>
  </div>
{/if}
