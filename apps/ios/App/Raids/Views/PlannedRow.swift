import DesignSystem
import Models
import SwiftUI

/// A planned raid's row on the Calendar page (min 76): date block, title, when, count with a 64 pt
/// bar, one chip.
struct PlannedRow: View {
    @Environment(\.dynamicTypeSize) private var typeSize
    let raid: RaidView
    let clock: RaidClock
    let now: Date
    let currentKey: String?
    let open: () -> Void

    var body: some View {
        let facts = RaidFacts(raid: raid, now: now)
        let large = typeSize.isAccessibilitySize
        Button(action: open) {
            HStack(alignment: .top, spacing: 12) {
                // The date block keeps its width; the text column takes the rest.
                DateBlock(date: raid.startsAt, clock: clock, width: large ? 56 : 48, height: 52)
                    .frame(width: large ? 56 : nil)
                VStack(alignment: .leading, spacing: 4) {
                    Text(raid.title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong).lineLimit(2)
                    Text(clock.when(raid.startsAt, now: now)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    HStack(spacing: 8) {
                        Text(raid.size == nil ? facts.countLine : facts.ratio)
                            .font(ScaledType.tabularMeta).foregroundStyle(FSColor.strong)
                        if let fill = facts.fill, !large { FillBar(fill: fill, width: 64) }
                    }
                    // Large text: the chip moves under the title and wraps inside the row.
                    if large { chip(facts.rowChip(currentKey: currentKey), wraps: true) }
                }
                if !large {
                    Spacer(minLength: 8)
                    chip(facts.rowChip(currentKey: currentKey))
                }
            }
            .padding(.vertical, 14)
            .frame(minHeight: 76)
            .contentShape(Rectangle())
        }
        .buttonStyle(RowPressStyle())
        .accessibilityElement(children: .combine)
    }

    @ViewBuilder
    private func chip(_ chip: RowChip, wraps: Bool = false) -> some View {
        let (text, color): (String, Color) = switch chip {
        case let .status(status): (status.label, TagView.color(for: status))
        case let .signed(text): (text, FSColor.kill)
        case .notSigned: (Copy.raidChipNotSigned, FSColor.gold)
        }
        if wraps {
            WrappingTag(text: text, color: color)
        } else {
            TagView(text: text, color: color)
        }
    }
}
