import DesignSystem
import SwiftUI

/// `ui/ItemRow` (`.irow`): the item's icon in its quality border, the name in its quality colour
/// (the tooltip trigger), where it comes from, and the gain at the right. 58 min, padded 8 by 16,
/// no hairline. Home's Next upgrades and Gear's lists draw items with it.
struct ItemLineView: View {
    let name: String
    let quality: Int?
    let icon: String?
    let source: String
    /// The gain's figure and unit (`+14.5`, `DPS`); nil for none.
    var gain: (figure: String, unit: String)?
    let build: String

    static let minHeight: CGFloat = 58
    /// Measured on the mocks' rows: a one-line-source row is 62 high.
    static let iconDescent: CGFloat = 6

    var body: some View {
        HStack(spacing: FSSpace.space3) {
            ItemIconView(icon: icon, quality: quality, build: build, size: FSSize.icon40)
                .parityItem(.image)
                // The web's `.irow` icon is inline, so its row carries the line's descent under it.
                .padding(.bottom, Self.iconDescent)
            VStack(alignment: .leading, spacing: 0) {
                Text(name).type(FSType.item).foregroundStyle(ItemQuality(value: quality).text)
                Text(source).type(RowType.sub).foregroundStyle(FSColor.muted)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            if let gain {
                HStack(alignment: .firstTextBaseline, spacing: FSSpace.space1) {
                    Text(gain.figure).type(FSType.figureSm).foregroundStyle(FSColor.strong)
                    Text(gain.unit)
                        .font(.custom(FSFont.bodySemibold, size: 13, relativeTo: .footnote))
                        .foregroundStyle(FSColor.muted)
                }
                .fixedSize()
            }
        }
        .padding(.horizontal, FSSpace.space4)
        .padding(.vertical, FSSpace.space2)
        .frame(minHeight: Self.minHeight)
        .contentShape(.rect)
        .accessibilityElement(children: .combine)
        .accessibilityHint(Copy.itemsA11yHint)
        .parityItem(.row)
    }
}
