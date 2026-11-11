import DesignSystem
import Models
import Networking
import Routing
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
    @State private var nextRaid: ItemNextRaid?
    @Environment(\.openAppLink) private var openAppLink
    @Environment(\.dismiss) private var dismiss

    enum Phase: Equatable {
        case loading
        case loaded(ItemSheetContent)
        case failed
    }

    var body: some View {
        Group {
            switch phase {
            case let .loaded(content):
                ItemSheetView(
                    content: content.with(guild: guild, nextRaid: request.guildID.flatMap {
                        ItemSheetBuilder.nextRaid(nextRaid, guild: $0, zone: content.raidZone)
                    }),
                    build: request.tooltip.build, moreAbout: moreAbout,
                    openRaid: openAppLink.map { open in { link in dismiss(); open(link) } }
                )
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
            let answer = try await needs.needs(guild: id, item: request.tooltip.itemID)
            guild = ItemSheetBuilder.guildRows(answer)
            nextRaid = answer.nextRaid
        } catch {
            guild = .failed
        }
    }
}

/// The sheet's body (round 5, m-ios-item), pure over its content so the parity tests draw it
/// directly: the client's tooltip with the gain closing it, then one list (where it drops, the
/// guild's next raid there, what it replaces, the runner-up, who in the guild needs or has it). The
/// top 23 pt hold the system's drag indicator (the mock's grab: 9 above, 5, 9 below).
struct ItemSheetView: View {
    let content: ItemSheetContent
    let build: String
    var moreAbout: (() -> Void)?
    /// Opens the next raid's page; nil leaves the line as text.
    var openRaid: ((DeepLink) -> Void)?

    static let grabRoom: CGFloat = 23

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                ItemTooltipPanel(model: content.tooltip, gain: content.gain.map { "\($0.figure) \($0.unit)" })
                    .accessibilityAddTraits(.isHeader)
                    .parityRegion(.inspector)
                list
                    .padding(.top, FSSpace.space3)
                    .parityRegion(.section(Copy.gearSheetWhere))
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

    /// One line of the sheet's list.
    private enum Line: Identifiable {
        case row(ItemSheetRow)
        case nextRaid(ItemSheetNextRaid)

        var id: String {
            switch self {
            case let .row(row): row.id
            case .nextRaid: "next-raid"
            }
        }
    }

    private var lines: [Line] {
        var lines: [Line] = content.source.map { [.row($0)] } ?? []
        if let next = content.nextRaid { lines.append(.nextRaid(next)) }
        lines += [content.replaces, content.runnerUp].compactMap(\.self).map(Line.row)
        return lines + guildRows.map(Line.row)
    }

    private var list: some View {
        let lines = lines
        let note = guildNote
        return VStack(alignment: .leading, spacing: 0) {
            ForEach(Array(lines.enumerated()), id: \.element.id) { index, line in
                let last = note == nil && content.guild != .loading && index == lines.count - 1
                switch line {
                case let .row(row):
                    ItemSheetRowView(row: row, build: build, last: last)
                case let .nextRaid(next):
                    Button { openRaid?(next.link) } label: { PanelRow(next.line, last: last, chevron: openRaid != nil) }
                        .buttonStyle(RowPressStyle())
                        .disabled(openRaid == nil)
                }
            }
            if content.guild == .loading { FSSkeleton(height: FSSize.row).padding(.vertical, FSSpace.space2) }
            if let note { Text(note).type(FSType.small).foregroundStyle(FSColor.muted).padding(.vertical, FSSpace.space2) }
        }
    }

    private var guildRows: [ItemSheetRow] {
        if case let .loaded(rows, _) = content.guild { return rows }
        return []
    }

    private var guildNote: String? {
        switch content.guild {
        case let .loaded(_, note): note
        case .failed: Copy.itemsSheetGuildFailed
        case .none, .loading: nil
        }
    }
}

/// A row of the sheet: an item or a crest, the title in its quality or class colour, its line.
private struct ItemSheetRowView: View {
    let row: ItemSheetRow
    let build: String
    let last: Bool

    var body: some View {
        PanelRow(row.title, titleColor: titleColor, sub: row.sub.map { Text($0) }, last: last) {
            switch row.lead {
            case let .item(icon, quality):
                ItemIconView(icon: icon, quality: quality, build: build, size: FSSize.icon40).parityItem(.image)
            case let .crest(className):
                ClassCrest(className: className, size: FSSize.crest40, flat: true)
            case .none:
                EmptyView()
            }
        }
    }

    private var titleColor: Color {
        if let quality = row.titleQuality { return ItemQuality(value: quality).text }
        if let className = row.titleClass { return ClassColor.of(className) }
        return FSColor.strong
    }
}
