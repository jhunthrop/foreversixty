// web/src/components/premium/interval.ts
// /premium is static, so its interval bar is rendered with Yearly checked; a ?interval=monthly
// arrival (a reload, a shared link) checks Monthly in place before anyone uses the bar. The
// phone summary line follows. The prices themselves switch on the root's pre-paint `data-iv`.
import { checkRadio } from '../../lib/ui/control-bar';
import type { Interval } from '../../lib/billing/plans';

export function syncIntervalRadios(interval: Interval, root: ParentNode = document): void {
  const radio = root.querySelector<HTMLElement>(
    `[role="radiogroup"][data-param="interval"] [role="radio"][data-value="${interval}"]`,
  );
  if (radio === null || radio.getAttribute('aria-checked') === 'true') return;
  checkRadio(radio);
  const summary = radio
    .closest('[data-control-bar]')
    ?.querySelector<HTMLElement>('[data-control-bar-summary]');
  const label = radio.textContent?.trim();
  if (summary && label) summary.textContent = label;
}
