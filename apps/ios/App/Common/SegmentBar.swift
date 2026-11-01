import DesignSystem
import SwiftUI

/// A screen's segments (spec §4.5): equal cells on the raised ground, the checked one with the
/// gold 12% wash and its 2 pt gold underline; stacked at the accessibility sizes so a label is
/// never cut. A change plays the selection haptic.
struct SegmentBar<Item: Hashable>: View {
    let items: [Item]
    let selection: Item
    let title: (Item) -> String
    let select: (Item) -> Void
    @Environment(\.dynamicTypeSize) private var typeSize

    static var minHeight: CGFloat { 44 }
    /// Side by side, a long label shrinks a little before it would wrap; stacked it never needs to.
    static var minimumScale: CGFloat { 0.85 }

    var body: some View {
        let layout = typeSize.isAccessibilitySize ? AnyLayout(VStackLayout(spacing: 0)) : AnyLayout(HStackLayout(spacing: 0))
        layout {
            ForEach(items, id: \.self) { item in
                let checked = item == selection
                Button { select(item) } label: {
                    Text(title(item))
                        .font(ScaledType.segment)
                        .foregroundStyle(checked ? FSColor.strong : FSColor.muted)
                        .lineLimit(1)
                        .minimumScaleFactor(Self.minimumScale)
                        .padding(.horizontal, 8)
                        .padding(.vertical, typeSize.isAccessibilitySize ? 10 : 0)
                        .frame(maxWidth: .infinity, minHeight: Self.minHeight)
                        .overlay(alignment: .top) {
                            if typeSize.isAccessibilitySize, item != items.first { FSColor.line.frame(height: 1) }
                        }
                        .background(checked ? FSColor.gold.opacity(0.12) : .clear)
                        .overlay(alignment: .bottom) { if checked { FSColor.gold.frame(height: 2) } }
                        .contentShape(Rectangle())
                }
                .buttonStyle(.plain)
                .accessibilityAddTraits(checked ? .isSelected : [])
            }
        }
        .background(FSColor.raised)
        .clipShape(.rect(cornerRadius: FSRadius.control))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.line, lineWidth: 1))
        .sensoryFeedback(.selection, trigger: selection)
    }
}
