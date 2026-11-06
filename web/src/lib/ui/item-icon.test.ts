import { describe, expect, it } from 'vitest';
import { itemIconFace } from './item-icon';

const BASE = { icon: 'inv_helmet_09', deferred: false, mounted: false, failed: false };

describe('itemIconFace', () => {
  it('draws the image when it has an icon', () => {
    expect(itemIconFace(BASE)).toBe('image');
  });

  it('draws the dark socket when there is no icon', () => {
    expect(itemIconFace({ ...BASE, icon: undefined })).toBe('socket');
    expect(itemIconFace({ ...BASE, icon: '' })).toBe('socket');
  });

  it('falls back to the socket when the image fails, never a broken image', () => {
    expect(itemIconFace({ ...BASE, failed: true })).toBe('socket');
    expect(itemIconFace({ ...BASE, deferred: true, mounted: true, failed: true })).toBe('socket');
  });

  it('keeps a deferred icon to its frame until the island mounts', () => {
    expect(itemIconFace({ ...BASE, deferred: true })).toBe('frame');
    expect(itemIconFace({ ...BASE, deferred: true, mounted: true })).toBe('image');
  });
});
