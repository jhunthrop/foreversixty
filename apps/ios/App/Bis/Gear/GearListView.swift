import DesignSystem
import SwiftUI
import UIKit

/// What the Gear tab's rows hand back.
struct GearActions {
    var openPicker: () -> Void = {}
    var openItem: (GearItemRef) -> Void = { _ in }
}

/// Gear for a signed-in player (flows spec 4.3) in the round-5 look (spec 2026-11-09 Paperdoll "Mobile
/// Gear tab", mocks ios-gear and m-ios-gear-frame): under the band's large title, the paperdoll strip
/// (the character's portrait or the seal, the spec, the set's figure), the line naming the list with
/// Change, every slot in two columns, then Where to go tonight. Home holds the top three upgrades.
struct GearListView: View {
    let content: GearContent
    let build: String
    var portrait: UIImage?
    var actions = GearActions()

    /// The item icon inside the 44 pt socket.
    static let socketIcon: CGFloat = 38

    var body: some View {
        ScrollView {
            BandPage {
                EmptyView()
            } content: {
                PaperdollHeader(
                    title: content.title, figure: content.figure, unit: content.unit, portrait: portrait,
                    slots: content.slots.map(paperdollSlot),
                    onSelect: { id in content.slots.first { $0.slot == id }?.pick.map { actions.openItem($0.item) } }
                ) {
                    GearListLine(facts: content.facts, summary: content.summary, change: actions.openPicker)
                } icon: { item in
                    let look = looks[item.id]
                    ItemIconView(icon: look?.icon, quality: look?.quality, build: build, size: Self.socketIcon).parityItem(.image)
                }
                if !content.tonight.isEmpty {
                    PanelSection(Copy.gearWhereTonight) {
                        ForEach(Array(content.tonight.enumerated()), id: \.element.id) { index, place in
                            Button { actions.openItem(place.item) } label: {
                                PanelRow(place.title, sub: Self.placeSub(place), last: index == content.tonight.count - 1) {
                                    ItemIconView(icon: place.icon, quality: place.quality, build: build, size: FSSize.icon40).parityItem(.image)
                                }
                            }
                            .buttonStyle(RowPressStyle())
                        }
                    }
                }
            }
        }
        .tabRootScroll()
        .screenGround()
    }

    /// What the place gives, then `2 upgrades · +20.4 DPS` on its own line.
    static func placeSub(_ place: GearPlaceRow) -> Text {
        guard let figure = place.figure else { return Text(place.sub) }
        return Text("\(place.sub)\n\(figure)")
    }

    private var looks: [Int: HomeItemLook] {
        Dictionary(content.slots.compactMap { $0.pick.map { ($0.look.id, $0.look) } }, uniquingKeysWith: { first, _ in first })
    }

    /// A slot as the paperdoll's list draws it: the pick in its quality colour, its gain or Equipped.
    private func paperdollSlot(_ row: GearSlotRow) -> PaperdollListSlot {
        guard let pick = row.pick else { return PaperdollListSlot(id: row.slot, label: row.label, item: nil) }
        let detail: String? = switch pick.trail {
        case .worn: Copy.gearWorn
        case let .gain(gain): "\(gain.figure) \(gain.unit)"
        case .noGain: Copy.gearCellNoGain
        case .notModelled: Copy.gearNotModelled
        case nil: nil
        }
        let isGain = if case .gain = pick.trail { true } else { false }
        return PaperdollListSlot(id: row.slot, label: row.label, item: PaperdollListSlot.Item(
            id: pick.look.id, name: pick.look.name ?? Copy.itemsNotInData, quality: ItemQuality(value: pick.look.quality).text,
            detail: detail, isGain: isGain
        ))
    }
}

/// The line that names the list, `60 raid · Alliance · 4 of 17 slots on the list · Change` (14/20).
struct GearListLine: View {
    let facts: String
    /// `+157.6 DPS from 13 upgrades · synced 2 h ago`, muted under the line; a stale age in the warning tone.
    var summary: GearSummary?
    let change: () -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: FSSpace.space1) {
            line
            if let summary {
                summaryText(summary)
                    .type(RowType.sub)
                    .foregroundStyle(FSColor.muted)
                    .frame(maxWidth: .infinity, alignment: .leading)
            }
        }
    }

    private func summaryText(_ summary: GearSummary) -> Text {
        let synced = summary.synced.map { Text($0).foregroundColor(summary.stale ? FSColor.warn : FSColor.muted) }
        guard let figures = summary.figures else { return synced ?? Text("") }
        let lead = Text(figures).foregroundColor(FSColor.muted)
        guard let synced else { return lead }
        return Text("\(lead)\(Text(GearPresenter.separator).foregroundColor(FSColor.muted))\(synced)")
    }

    private var line: some View {
        Button(action: change) {
            Text.line(facts, link: Copy.gearChange, size: FSType.factLine.size)
                .type(FSType.factLine)
                .foregroundStyle(FSColor.text)
                .multilineTextAlignment(.leading)
                .frame(maxWidth: .infinity, alignment: .leading)
                .parityRegion(.controls)
                .contentShape(.rect)
        }
        .buttonStyle(.plain)
        .accessibilityHint(Copy.bisPickerTitle)
    }
}
