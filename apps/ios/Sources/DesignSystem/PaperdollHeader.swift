import SwiftUI
import UIKit

/// One slot of the phone paperdoll's list: the slot's word, and the item in it or nothing.
public struct PaperdollSlot: Identifiable, Sendable {
    /// What a filled slot shows.
    public struct Item: Sendable {
        public let id: Int
        public let name: String
        /// The client's quality colour for the name (the items' `ItemQuality`).
        public let quality: Color
        /// The second line after the slot's word: a gain (`+5.9 DPS`) or `Worn`.
        public let detail: String?
        /// A gain draws its detail in the gain green.
        public let isGain: Bool

        public init(id: Int, name: String, quality: Color, detail: String? = nil, isGain: Bool = false) {
            self.id = id
            self.name = name
            self.quality = quality
            self.detail = detail
            self.isGain = isGain
        }
    }

    /// The slot key (`head`, `finger1`, `off_hand`).
    public let id: String
    /// The slot's word (`Head`, `Off Hand`).
    public let label: String
    public let item: Item?
    /// The off hand under a two-hander: the main hand's icon greyed, and the caller's word for it
    /// (`Two-handed`, from the app's catalogue).
    public let coveredBy: Item?
    public let coveredNote: String?

    public init(id: String, label: String, item: Item?, coveredBy: Item? = nil, coveredNote: String? = nil) {
        self.id = id
        self.label = label
        self.item = item
        self.coveredBy = coveredBy
        self.coveredNote = coveredNote
    }
}

/// The paperdoll's phone form (spec round 5, Paperdoll "Phone"; web ui/PaperdollFrame.svelte's
/// list form): under 36 pt slots the frame gives way to its one-line header strip (the medallion
/// with the character's portrait, or the LX seal; the title; the set's DPS) and the slots in two
/// columns on the panel. A tap on a filled slot calls `onSelect`. `icon` draws a slot's item icon
/// (the app's `ItemIcon`), which sits in the 44 pt socket.
public struct PaperdollHeader<Icon: View>: View {
    private let title: String
    private let figure: String?
    private let unit: String
    private let portrait: UIImage?
    private let slots: [PaperdollSlot]
    private let selected: String?
    private let onSelect: ((String) -> Void)?
    private let icon: (PaperdollSlot.Item) -> Icon

    static var medallionSize: CGFloat { 40 }
    static var cellSize: CGFloat { 44 }
    static var iconInset: CGFloat { 3 }
    static var titleSize: CGFloat { 19 }
    static var figureSize: CGFloat { 17 }
    static var coveredOpacity: Double { 0.55 }

    public init(
        title: String,
        figure: String?,
        unit: String = "DPS",
        portrait: UIImage?,
        slots: [PaperdollSlot],
        selected: String? = nil,
        onSelect: ((String) -> Void)? = nil,
        @ViewBuilder icon: @escaping (PaperdollSlot.Item) -> Icon
    ) {
        self.title = title
        self.figure = figure
        self.unit = unit
        self.portrait = portrait
        self.slots = slots
        self.selected = selected
        self.onSelect = onSelect
        self.icon = icon
    }

    public var body: some View {
        VStack(spacing: 0) {
            strip
            list
        }
        .panel()
    }

    private var strip: some View {
        HStack(spacing: 10) {
            PaperdollMedallion(portrait: portrait, size: Self.medallionSize)
            Text(title)
                .font(.custom(FSFont.title, size: Self.titleSize, relativeTo: .title3))
                .foregroundStyle(Theme.Palette.title)
                .fixedSize(horizontal: false, vertical: true)
                .accessibilityAddTraits(.isHeader)
            Spacer(minLength: 0)
            if let figure {
                (Text(figure).foregroundStyle(FSColor.strong)
                    + Text(" " + unit).font(.custom(FSFont.bodySemibold, size: Self.figureSize * 0.66, relativeTo: .caption))
                    .foregroundStyle(FSColor.text))
                    .font(.custom(FSFont.bodySemibold, size: Self.figureSize, relativeTo: .headline))
                    .monospacedDigit()
                    .lineLimit(1)
                    .fixedSize()
                    .shadow(color: .black, radius: 1, y: 1)
            }
        }
        .padding(.horizontal, FSSpace.space3)
        .padding(.vertical, FSSpace.space2)
        .frame(maxWidth: .infinity, alignment: .leading)
        .background(Theme.Palette.bar)
        .overlay(alignment: .bottom) { Color.black.frame(height: FSStroke.hairline) }
        .accessibilityElement(children: .combine)
    }

    private var list: some View {
        LazyVGrid(
            columns: [GridItem(.flexible(), spacing: 10, alignment: .leading), GridItem(.flexible(), alignment: .leading)],
            alignment: .leading, spacing: 6
        ) {
            ForEach(slots) { slot in cell(slot) }
        }
        .padding(FSSpace.space3)
        .parityRegion(.slots)
    }

    @ViewBuilder
    private func cell(_ slot: PaperdollSlot) -> some View {
        let shown = slot.item ?? slot.coveredBy
        let row = HStack(spacing: FSSpace.space2) {
            socket(shown, covered: slot.item == nil, selected: slot.id == selected)
            VStack(alignment: .leading, spacing: 0) {
                if let item = slot.item {
                    Text(item.name)
                        .font(.custom(FSFont.bodySemibold, size: 13, relativeTo: .footnote))
                        .foregroundStyle(item.quality)
                        .fixedSize(horizontal: false, vertical: true)
                    sub(slot.label, detail: item.detail, gain: item.isGain)
                } else {
                    sub(slot.label, detail: nil, gain: false)
                    if let note = slot.coveredNote { sub(note, detail: nil, gain: false) }
                }
            }
            .frame(maxWidth: .infinity, alignment: .leading)
        }
        .frame(minHeight: FSTarget.min)
        .contentShape(.rect)
        .parityItem(.row)
        if let item = slot.item, let onSelect {
            Button { onSelect(slot.id) } label: { row }
                .buttonStyle(.plain)
                .accessibilityLabel("\(item.name), \(slot.label)")
                .accessibilityAddTraits(slot.id == selected ? .isSelected : [])
        } else {
            row.accessibilityElement(children: .combine)
        }
    }

    private func sub(_ label: String, detail: String?, gain: Bool) -> some View {
        var line = Text(label).foregroundStyle(FSColor.muted)
        if let detail {
            line = line + Text(" · ").foregroundStyle(FSColor.muted)
                + Text(detail).foregroundStyle(gain ? FSColor.gain : FSColor.muted).fontWeight(gain ? .bold : .regular)
        }
        return line
            .font(.custom(FSFont.body, size: 12, relativeTo: .caption))
            .lineLimit(1)
    }

    private func socket(_ item: PaperdollSlot.Item?, covered: Bool, selected: Bool) -> some View {
        ZStack {
            RoundedRectangle(cornerRadius: Theme.Radius.icon).fill(Theme.Palette.well)
            RoundedRectangle(cornerRadius: Theme.Radius.icon).strokeBorder(FSColor.line, lineWidth: FSStroke.hairline)
            if let item {
                icon(item)
                    .padding(Self.iconInset)
                    .saturation(covered ? 0 : 1)
                    .opacity(covered ? Self.coveredOpacity : 1)
            }
        }
        .frame(width: Self.cellSize, height: Self.cellSize)
        .overlay {
            if selected {
                RoundedRectangle(cornerRadius: Theme.Radius.icon + 3)
                    .strokeBorder(FSColor.cream, lineWidth: FSStroke.focus)
                    .padding(-3)
            }
        }
    }
}

/// The paperdoll's medallion: the character's Battle.net portrait, else the LX seal (owner,
/// 2026-11-08), never the class crest. A circle at `size`.
public struct PaperdollMedallion: View {
    private let portrait: UIImage?
    private let size: CGFloat

    public init(portrait: UIImage?, size: CGFloat) {
        self.portrait = portrait
        self.size = size
    }

    public var body: some View {
        Group {
            if let image = portrait ?? BrandArt.seal {
                Image(uiImage: image).resizable().scaledToFill()
            } else {
                Theme.Palette.well
            }
        }
        .frame(width: size, height: size)
        .clipShape(.circle)
        .accessibilityHidden(true)
    }
}
