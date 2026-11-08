import DesignSystem
import SwiftUI

/// What the Gear tab's rows hand back.
struct GearActions {
    var openPicker: () -> Void = {}
    var openItem: (GearItemRef) -> Void = { _ in }
}

/// Gear for a signed-in player (flows spec 4.3, mock ios-gear): the list it compares against with
/// Change, Where to go tonight, then every slot with Worn or the gain. Home holds the top three
/// upgrades; Gear does not repeat them.
struct GearListView: View {
    let content: GearContent
    let build: String
    var actions = GearActions()

    /// `.ap section { margin-top: 28px }`.
    static let sectionGap: CGFloat = 28

    var body: some View {
        GearPage {
            GearSummaryRow(line: content.summary, change: actions.openPicker)
            if !content.tonight.isEmpty {
                FSSection(Copy.gearTonight) {
                    VStack(alignment: .leading, spacing: 0) {
                        ForEach(content.tonight) { place in
                            Button { actions.openItem(place.item) } label: { GearPlaceRowView(place: place, build: build) }
                                .buttonStyle(RowPressStyle())
                        }
                    }
                }
                .padding(.top, Self.sectionGap)
            }
            FSSection(Copy.gearEverySlot) {
                VStack(alignment: .leading, spacing: 0) {
                    ForEach(content.slots) { GearSlotRowView(row: $0, build: build, open: actions.openItem) }
                }
                .parityRegion(.slots)
            }
            .padding(.top, Self.sectionGap)
        }
    }
}

/// A Gear page's scroll: 8 under the bar, the phone gutter, the content's parity space.
struct GearPage<Content: View>: View {
    @ViewBuilder let content: () -> Content

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0, content: content)
                .padding(.top, FSSpace.space2)
                .padding(.horizontal, FSSpace.gutterPhone)
                .padding(.bottom, FSSpace.space6)
                .frame(maxWidth: .infinity, alignment: .leading)
                .parityContent()
        }
        .background(FSColor.bg)
    }
}

/// `Fury Warrior · 60 raid · Alliance`, `4 of 17 slots on the list`, and Change under them.
struct GearSummaryRow: View {
    let line: HomeLine
    let change: () -> Void

    /// The link's 44 pt target, padded without taking the room (web ActionLink).
    static let targetPad: CGFloat = 12

    var body: some View {
        FSRow(title: line.title, sub: line.sub, inset: 0, trailPlacement: .below) {
            EmptyView()
        } trail: {
            Button(action: change) {
                Text(Copy.gearChange)
                    .type(FSType.item)
                    .foregroundStyle(FSColor.gold)
                    .padding(.vertical, Self.targetPad)
                    .contentShape(.rect)
                    .padding(.vertical, -Self.targetPad)
            }
            .buttonStyle(.plain)
            .accessibilityHint(Copy.bisPickerTitle)
        }
        .parityRegion(.controls)
    }
}

/// A place: its name, the items it gives and how, its best item's icon.
private struct GearPlaceRowView: View {
    let place: GearPlaceRow
    let build: String

    var body: some View {
        FSRow(title: place.title, sub: place.sub, inset: 0) {
            ItemIconView(icon: place.icon, quality: place.quality, build: build, size: FSSize.icon40).parityItem(.image)
        }
    }
}

/// A slot: the pick with its slot and source, and Worn or the gain; or why the slot has none.
private struct GearSlotRowView: View {
    let row: GearSlotRow
    let build: String
    let open: (GearItemRef) -> Void

    var body: some View {
        switch row.content {
        case let .pick(look, source, trail, item):
            Button { open(item) } label: {
                ItemLineView(
                    name: look.name ?? Copy.itemsNotInData, quality: look.quality, icon: look.icon, source: source,
                    gain: trail.map(Self.figure), build: build
                )
            }
            .buttonStyle(RowPressStyle())
        case let .empty(reason):
            FSRow(title: reason, inset: FSSpace.space4) { EmptyView() }
        }
    }

    private static func figure(_ trail: GearSlotRow.Trail) -> (figure: String, unit: String) {
        switch trail {
        case .worn: (Copy.gearWorn, "")
        case let .gain(gain): (gain.figure, gain.unit)
        }
    }
}
