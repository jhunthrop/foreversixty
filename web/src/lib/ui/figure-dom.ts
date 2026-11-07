// web/src/lib/ui/figure-dom.ts
// Painting a server-rendered ui/Figure and ui/Meter from a page script (no island): the tier
// list's answer and the /bis landing's count write their numbers into the static markup.

/** The figure's value is its first text node, its unit the span after it, its caption the line
 *  under it (ui/Figure.svelte). */
export function writeFigure(figure: HTMLElement, value: string, unit: string, caption?: string): void {
  const line = figure.querySelector<HTMLElement>('.ui-figure > span');
  if (line === null) return;
  const text = [...line.childNodes].find((node) => node.nodeType === Node.TEXT_NODE);
  if (text !== undefined) text.nodeValue = value;
  const unitNode = line.querySelector('span');
  if (unitNode !== null) unitNode.textContent = unit;
  const captionNode = line.nextElementSibling;
  if (caption !== undefined && captionNode instanceof HTMLElement) captionNode.textContent = caption;
}

/** The meter's fill, 0 to 1, clamped (ui/Meter.svelte without segments). */
export function writeMeter(meter: HTMLElement, value: number): void {
  const fill = meter.querySelector<HTMLElement>('.ui-meter > i');
  const fraction = Math.min(1, Math.max(0, Number.isFinite(value) ? value : 0));
  if (fill !== null) fill.style.width = `${(fraction * 100).toFixed(1)}%`;
}
