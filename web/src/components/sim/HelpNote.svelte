<!-- web/src/components/sim/HelpNote.svelte -->
<!-- Task 7 (newcomer BLOCKER: `[data-tooltip],[role=tooltip],abbr,.tooltip` was 0 on a
     phone; the only explanations on /sim were `title=`, which native touch never fires --
     100% of the page's help was unreachable off a mouse). One reusable help affordance,
     built on Disclosure.svelte (Ruling 3's one show/hide primitive) rather than a second
     toggle mechanism: every control in SettingsBar.svelte, SettingsSheet.svelte,
     RunControl.svelte and ReportOptions.svelte that used to carry a `title=` or carried no
     explanation at all now opens one of these instead.

     The trigger's own visible text IS its accessible name ("What <label> means"), the same
     way `simCopy.whatsInIt` and `simCopy.rotationLink` already work as a Disclosure trigger
     -- so a screen reader user hears which control a given button explains without relying
     on visual proximity alone, and there is nothing to mistake for a second link reading
     the same words next to a different control. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { simCopy } from '../../lib/sim/copy';
  import Disclosure from './Disclosure.svelte';

  let {
    label,
    id,
    disabled = false,
    open = $bindable(false),
    children,
  }: {
    /** The control's own visible label, e.g. "Fight style" -- embedded in the trigger's
     *  accessible name via `simCopy.helpTrigger`. */
    label: string;
    /** Unique within the page; becomes `${id}-help`, Disclosure's own id. */
    id: string;
    disabled?: boolean;
    /** Forwarded to Disclosure's own bindable -- every caller in this lane opens closed
     *  (Disclosure's default), but the prop is threaded through rather than dropped so a
     *  test can render both states without a click (this project's vitest config hands
     *  `mount()` Svelte's server entry -- see lib/report/tree-sizes.ts's own note -- so
     *  HelpNote.test.ts renders open and closed by prop, and Playwright covers the tap). */
    open?: boolean;
    /** The note's body. A single `<p>` for a plain sentence; Fight style renders a `<dl>`
     *  of the nine styles instead, since Disclosure's children accept any markup. */
    children: Snippet;
  } = $props();
</script>

<Disclosure
  label={simCopy.helpTrigger(label)}
  id={`${id}-help`}
  {disabled}
  bind:open
  triggerClass="label text-nav text-[12px] underline decoration-dotted underline-offset-2"
  panelClass="border-line-soft rounded-panel text-muted flex max-w-[320px] flex-col gap-2 border p-3 text-[12px]"
  triggerTestId={`${id}-help-trigger`}
  panelTestId={`${id}-help-panel`}
>
  {@render children()}
</Disclosure>
