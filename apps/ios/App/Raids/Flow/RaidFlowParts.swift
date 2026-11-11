import DesignSystem
import SwiftUI

/// `.ap section { margin-top: 28px }`, `.blk { margin-top: 24px }`, the small lines' 8 and the
/// buttons' 20 (`.vstack`).
enum RaidFlowSpacing {
    static let section: CGFloat = 28
    static let block: CGFloat = 24
    static let line: CGFloat = FSSpace.space2
    static let buttons: CGFloat = 20
}

/// A screen of the Raid flows: the gutter, the top inset and the parity content, on the page ground.
struct RaidFlowScroll<Content: View>: View {
    @ViewBuilder let content: () -> Content

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) { content() }
                .padding(.top, FSSpace.space2)
                .padding(.horizontal, FSSpace.gutterPhone)
                .padding(.bottom, FSSpace.space6)
                .parityContent()
        }
        .screenGround()
    }
}

/// A gold text action (`.link`), at the row's foot or after a sentence.
struct RaidLinkButton: View {
    let title: String
    let action: () -> Void

    var body: some View {
        Button(title, action: action)
            .font(.custom(FSFont.bodySemibold, size: 15, relativeTo: .body))
            .foregroundStyle(FSColor.gold)
            .frame(minHeight: 20)
    }
}

/// The Ready check (F2): each item ticked, warned or missing, where to get a miss, the export's age.
struct ReadyChecklistView: View {
    let content: ReadyListContent
    let open: (URL) -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            ForEach(content.rows) { row in
                FSRow(title: row.title, sub: row.line, inset: 0, trailPlacement: .below) {
                    ReadyMark(mark: row.mark)
                } trail: {
                    if let fix = row.fix {
                        RaidLinkButton(title: fix) { if let link = row.link { open(link) } }
                            .disabled(row.link == nil)
                    }
                }
            }
            if let stamp = content.stamp { SmallText(stamp).padding(.top, RaidFlowSpacing.line) }
        }
    }
}

/// A ready-check mark on its filled disc: green ticked, red missing, gold to do; its word to VoiceOver.
struct ReadyMark: View {
    let mark: ReadyRowContent.Mark

    static let size: CGFloat = 22

    var body: some View {
        let (glyph, color, word): (String, Color, String) = switch mark {
        case .ok: ("checkmark", FSColor.kill, Copy.raidReadyOk)
        case .warn: ("exclamationmark", FSColor.gold, Copy.raidReadyWarn)
        case .miss: ("xmark", FSColor.wipe, Copy.raidReadyMiss)
        }
        Image(systemName: glyph)
            .font(.system(size: 11, weight: .heavy))
            .foregroundStyle(FSColor.bg)
            .frame(width: Self.size, height: Self.size)
            .background(color, in: .circle)
            .accessibilityLabel(word)
    }
}

/// `Tanks 2 of 2` over the names, slim rows (raid night).
struct RaidRoleRows: View {
    let roles: [HomeLine]

    var body: some View {
        VStack(spacing: 0) {
            ForEach(roles, id: \.title) { role in
                FSRow(title: role.title, sub: role.sub.isEmpty ? nil : role.sub, density: .slim, inset: 0) { EmptyView() }
            }
        }
    }
}

/// A head with a lead (the zone icon, a crest) and its facts wrapping under the title, on the band
/// under the clear bar (N3), 16 above the content.
struct RaidFlowHead<Lead: View>: View {
    let title: String
    var titleColor: Color = FSColor.strong
    let facts: [String]
    /// How far in from the scroll's gutter the head sits (raid night's `.ph`).
    var inset: CGFloat = 0
    @ViewBuilder let lead: () -> Lead

    var body: some View {
        // On the band under the clear bar (N3): the art runs up under the bars and out to the edges.
        PageHeader(title: title, titleColor: titleColor) {
            lead()
        } meta: {
            WrapLayout(spacing: FSSpace.space4, lineSpacing: 6) {
                ForEach(facts, id: \.self) { SmallText($0) }
            }
        }
        .padding(.horizontal, inset)
        .bandUnderBars(bleed: FSSpace.gutterPhone + inset)
        .padding(.bottom, BandReach.gap)
    }
}

extension RaidFlowHead where Lead == EmptyView {
    init(title: String, facts: [String] = []) {
        self.init(title: title, facts: facts, lead: { EmptyView() })
    }
}
