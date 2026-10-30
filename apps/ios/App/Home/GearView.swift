import DesignTokens
import Models
import SwiftUI

/// F2.7 Your gear: all 17 slots in character-frame order, worn then the pick.
struct GearView: View {
    let rows: [GearRow]
    let intro: String
    let build: String
    /// Opens an item's tooltip; the second value is the item worn in that slot.
    var open: (HomeItemLook, Int?) -> Void = { _, _ in }

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 16) {
                Text(intro).font(HomeType.body).foregroundStyle(FSColor.muted)
                VStack(spacing: 0) {
                    ForEach(Array(rows.enumerated()), id: \.element.id) { index, row in
                        if index > 0 { Divider().overlay(FSColor.lineSoft) }
                        GearRowView(row: row, build: build, open: open)
                    }
                }
                .quietCard()
            }
            .padding(16)
        }
        .background(FSColor.bg)
        .navigationTitle(Copy.gearTitle)
        .navigationBarTitleDisplayMode(.large)
    }
}

struct GearRowView: View {
    let row: GearRow
    let build: String
    let open: (HomeItemLook, Int?) -> Void

    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let layout = typeSize.isAccessibilitySize
            ? AnyLayout(VStackLayout(alignment: .leading, spacing: 8))
            : AnyLayout(HStackLayout(alignment: .center, spacing: 12))
        layout {
            MicroLabel(row.slotName).frame(minWidth: 62, alignment: .leading)
            HStack(alignment: .center, spacing: 12) {
                worn
                Spacer(minLength: 0)
                trailing
            }
        }
        .padding(.horizontal, 16)
        .padding(.vertical, 12)
        .frame(minHeight: 72)
    }

    @ViewBuilder
    private var worn: some View {
        Button { if let look = row.worn { open(look, look.id) } } label: {
            HStack(alignment: .center, spacing: 12) {
                ItemIconView(icon: row.worn?.icon, quality: row.worn?.quality, build: build, size: 36, dashed: row.worn == nil)
                VStack(alignment: .leading, spacing: 2) {
                    Text(row.wornName).font(HomeType.meta)
                        .foregroundStyle(row.worn == nil ? FSColor.muted : ItemQuality(value: row.worn?.quality).text)
                    if let level = row.wornLevel { Text(level).font(HomeType.meta).foregroundStyle(FSColor.muted) }
                    if let pick = row.pick {
                        Text(Copy.itemsPickArrow(item: pick.name ?? Copy.itemsNotInData)).font(HomeType.meta)
                            .foregroundStyle(ItemQuality(value: pick.quality).text)
                    }
                }
                .multilineTextAlignment(.leading)
            }
            .contentShape(Rectangle())
        }
        .buttonStyle(.plain)
        .disabled(row.worn == nil)
        .accessibilityElement(children: .combine)
        .accessibilityLabel([row.slotName, row.wornName, row.wornLevel, row.pick.map { Copy.itemsPickArrow(item: $0.name ?? "") }]
            .compactMap(\.self).joined(separator: ", "))
        .accessibilityHint(row.worn == nil ? "" : Copy.itemsA11yHint)
    }

    @ViewBuilder
    private var trailing: some View {
        if row.isBestInSlot {
            HStack(spacing: 4) {
                Image(systemName: "checkmark").font(.system(size: 10, weight: .bold)).accessibilityHidden(true)
                Text(Copy.gearBis.uppercased()).font(HomeType.micro)
            }
            .foregroundStyle(FSColor.kill)
            .padding(.horizontal, 8)
            .frame(minHeight: 24)
            .background(FSColor.kill.opacity(0.12), in: .rect(cornerRadius: FSRadius.pill))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.pill).stroke(FSColor.kill.opacity(0.35)))
        } else if let pick = row.pick {
            Button { open(pick, row.worn?.id) } label: {
                HStack(spacing: 8) {
                    Image(systemName: "chevron.right").font(.system(size: 12)).foregroundStyle(FSColor.muted)
                    ItemIconView(icon: pick.icon, quality: pick.quality, build: build, size: 36)
                }
                .frame(minWidth: 44, minHeight: 44)
                .contentShape(Rectangle())
            }
            .buttonStyle(.plain)
            .accessibilityLabel(Copy.tooltipA11yTile(item: pick.name ?? "", quality: ItemQuality(value: pick.quality).name, slot: row.slotName))
            .accessibilityHint(Copy.itemsA11yHint)
        }
    }
}
