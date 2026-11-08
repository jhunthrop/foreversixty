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
    /// The words under the gain (`over your worn`); nil for none.
    var caption: String?
    let build: String

    static let minHeight: CGFloat = 58

    var body: some View {
        HStack(spacing: FSSpace.space3) {
            ItemIconView(icon: icon, quality: quality, build: build, size: FSSize.icon40)
                .parityItem(.image)
            VStack(alignment: .leading, spacing: 0) {
                Text(name).type(FSType.item).foregroundStyle(ItemQuality(value: quality).text)
                Text(source).type(RowType.sub).foregroundStyle(FSColor.muted)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            if let gain {
                VStack(alignment: .trailing, spacing: 2) {
                    HStack(alignment: .firstTextBaseline, spacing: FSSpace.space1) {
                        Text(gain.figure).type(FSType.figureSm).foregroundStyle(FSColor.strong)
                        if !gain.unit.isEmpty {
                            Text(gain.unit)
                                .font(.custom(FSFont.bodySemibold, size: 13, relativeTo: .footnote))
                                .foregroundStyle(FSColor.muted)
                        }
                    }
                    if let caption {
                        Text(caption).type(RowType.detail).foregroundStyle(FSColor.muted)
                    }
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
