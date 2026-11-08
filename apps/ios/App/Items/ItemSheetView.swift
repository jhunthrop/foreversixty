import DesignSystem
import Models
import Networking
import SwiftUI

/// The item sheet (flows spec F5, mock ios-item-sheet), loading its files and the guild's answer.
/// Opens at the medium detent, grows to large.
struct ItemSheetScreen: View {
    let request: ItemSheetRequest
    let catalog: ItemCatalog
    /// Nil outside a guild or signed out.
    let needs: (any ItemNeedsSource)?
    /// Lane F12: `More about this item`, which pushes the item page; nil while search is off.
    var moreAbout: (() -> Void)?

    @State private var phase: Phase = .loading
    @State private var guild: ItemSheetGuild = .none

    enum Phase: Equatable {
        case loading
        case loaded(ItemSheetContent)
        case failed
    }

    var body: some View {
        Group {
            switch phase {
            case let .loaded(content):
                ItemSheetView(content: content.with(guild: guild), build: request.tooltip.build, moreAbout: moreAbout)
            case .loading, .failed:
                // The tooltip's own loading and failed states, with Try again.
                ItemTooltipContent(
                    request: request.tooltip, phase: phase == .failed ? .failed : .loading, retry: { Task { await load() } }
                )
            }
        }
        .presentationDetents([.medium, .large])
        .presentationDragIndicator(.visible)
        .presentationBackground(FSColor.raised)
        .task { await load() }
        .task { await loadGuild() }
    }

    private func load() async {
        phase = .loading
        do {
            let sources = try await catalog.sources(for: request.tooltip)
            let loot = await catalog.loot(build: request.tooltip.build)
            phase = .loaded(ItemSheetBuilder.content(request, sources: sources, loot: loot, guild: .none))
        } catch {
            phase = .failed
        }
    }

    private func loadGuild() async {
        guard let needs, let id = request.guildID else { return }
        guild = .loading
        do {
            guild = try await ItemSheetBuilder.guildRows(needs.needs(guild: id, item: request.tooltip.itemID))
        } catch {
            guild = .failed
        }
    }
}

/// The sheet's body, pure over its content so the parity tests draw it directly. The top 17 pt
/// hold the system's drag indicator (the mock's grab: 8 above, 5, 4 below).
struct ItemSheetView: View {
    let content: ItemSheetContent
    let build: String
    var moreAbout: (() -> Void)?

    static let grabRoom: CGFloat = 17

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                ItemLineView(
                    name: content.look.name ?? Copy.itemsNotInData, quality: content.look.quality, icon: content.look.icon,
                    source: content.slotLabel, gain: content.gain.map { ($0.figure, $0.unit) },
                    caption: content.gain == nil ? nil : Copy.itemsSheetOverWorn, build: build
                )
                .accessibilityAddTraits(.isHeader)
                lines
                if let source = content.source { section(Copy.itemsSheetWhere, [source]) }
                if let runnerUp = content.runnerUp { section(Copy.itemsSheetRunnerUp, [runnerUp]) }
                guild
                if let moreAbout {
                    SectionLink(title: Copy.itemPageMoreAbout, action: moreAbout).padding(.top, FSSpace.space4)
                }
            }
            .padding(.top, Self.grabRoom)
            .padding(.horizontal, FSSpace.gutterPhone)
            .padding(.bottom, FSSpace.space6)
        }
        .background(FSColor.raised)
    }

    private var lines: some View {
        VStack(alignment: .leading, spacing: 0) {
            ForEach(Array(content.lines.enumerated()), id: \.offset) { _, line in
                Text(line.text).type(FSType.small).foregroundStyle(color(line.tone)).fixedSize(horizontal: false, vertical: true)
            }
        }
        .padding(.vertical, FSSpace.space2)
        .accessibilityElement(children: .combine)
    }

    private func color(_ tone: ItemSheetLine.Tone) -> Color {
        switch tone {
        case .plain: FSColor.muted
        case .green: FSColor.parseGreen
        case .gold: FSColor.parseGold
        }
    }

    @ViewBuilder
    private var guild: some View {
        switch content.guild {
        case .none:
            EmptyView()
        case .loading:
            FSSection(Copy.itemsSheetGuild) { FSSkeleton(height: FSSize.row).padding(.top, FSSpace.space2) }
        case let .loaded(rows, note):
            FSSection(Copy.itemsSheetGuild) {
                VStack(alignment: .leading, spacing: 0) {
                    ForEach(rows) { ItemSheetRowView(row: $0, build: build) }
                    if let note { SmallText(note).padding(.top, FSSpace.space2) }
                }
            }
        case .failed:
            FSSection(Copy.itemsSheetGuild) { SmallText(Copy.itemsSheetGuildFailed).padding(.top, FSSpace.space2) }
        }
    }

    private func section(_ title: String, _ rows: [ItemSheetRow]) -> some View {
        FSSection(title) {
            VStack(alignment: .leading, spacing: 0) {
                ForEach(rows) { ItemSheetRowView(row: $0, build: build) }
            }
        }
    }
}

/// A row of the sheet: an item or a crest, the title in its quality or class colour, its line.
private struct ItemSheetRowView: View {
    let row: ItemSheetRow
    let build: String

    var body: some View {
        FSRow(title: row.title, titleColor: titleColor, sub: row.sub, inset: 0) {
            switch row.lead {
            case let .item(icon, quality):
                ItemIconView(icon: icon, quality: quality, build: build, size: FSSize.icon40).parityItem(.image)
            case let .crest(className):
                ClassCrest(className: className, size: FSSize.crest40, flat: true)
            }
        }
    }

    private var titleColor: Color {
        if let quality = row.titleQuality { return ItemQuality(value: quality).text }
        if let className = row.titleClass { return ClassColor.of(className) }
        return FSColor.strong
    }
}
