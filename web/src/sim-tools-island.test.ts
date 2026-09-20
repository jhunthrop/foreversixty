// @vitest-environment jsdom
// web/src/sim-tools-island.test.ts
import { describe, expect, it } from 'vitest';
import { toolFrom } from './sim-tools-island';

describe('toolFrom', () => {
  it('prefers the mount’s own stamp', () => {
    const element = document.createElement('div');
    element.dataset.simTool = 'drops';
    expect(toolFrom(element, '/sim/gear')).toBe('drops');
  });

  it('falls back to the path', () => {
    const element = document.createElement('div');
    expect(toolFrom(element, '/sim/weights')).toBe('weights');
    expect(toolFrom(element, '/sim/weights.html')).toBe('weights');
  });

  it('refuses anything not a tool, and lands on gear rather than throwing', () => {
    const element = document.createElement('div');
    element.dataset.simTool = 'javascript:alert(1)';
    expect(toolFrom(element, '/sim/nonesuch')).toBe('gear');
  });
});
