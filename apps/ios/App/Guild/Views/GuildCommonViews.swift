import DesignSystem
import SwiftUI

/// F16.7: the Roster and Raid history rows under the Raids tab's guild head.
struct GuildTabRows {
    /// `counts.on_roster` from the guild home; nil until it answers.
    let onRoster: Int?
    let openRoster: () -> Void
    let openHistory: () -> Void
}

struct GuildTabRowsCard: View {
    let rows: GuildTabRows

    var body: some View {
        VStack(spacing: 0) {
            Button(action: rows.openRoster) {
                AccountRow(icon: "person.2", title: Copy.raidsRosterRow, subtitle: rows.onRoster.map { Copy.raidsRosterRowSub(n: $0) }, chevron: true)
            }
            Divider().overlay(FSColor.lineSoft)
            Button(action: rows.openHistory) {
                AccountRow(icon: "doc.text", title: Copy.raidsHistoryRow, subtitle: Copy.raidsHistoryRowSub, chevron: true)
            }
        }
        .buttonStyle(RowPressStyle())
        .quietCard()
    }
}

/// A pushed guild screen's frame: the offline strip, the condensed-caps title (v1 §4, as Sims
/// and the raid page) with an optional trailing control, then the padded content.
struct GuildScreenScroll<Content: View, Trailing: View>: View {
    let title: String
    let offlineStrip: String?
    @ViewBuilder let trailing: () -> Trailing
    @ViewBuilder let content: () -> Content

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                if let offlineStrip { OfflineStrip(text: offlineStrip) }
                VStack(alignment: .leading, spacing: 16) {
                    HStack(alignment: .center, spacing: 12) {
                        Text(title.uppercased()).font(ScaledType.title).foregroundStyle(FSColor.strong)
                            .fixedSize(horizontal: false, vertical: true)
                            .accessibilityAddTraits(.isHeader)
                        Spacer(minLength: 0)
                        trailing()
                    }
                    content()
                }
                .padding(16)
            }
        }
        .background(FSColor.bg)
    }
}

extension GuildScreenScroll where Trailing == EmptyView {
    init(title: String, offlineStrip: String?, @ViewBuilder content: @escaping () -> Content) {
        self.init(title: title, offlineStrip: offlineStrip, trailing: { EmptyView() }, content: content)
    }
}

extension View {
    /// A pushed guild screen's bar: the back control only; the title draws in the content. The
    /// title still names the next screen's back button.
    func guildPushedBar(_ title: String) -> some View {
        navigationTitle(title)
            .navigationBarTitleDisplayMode(.inline)
            .toolbar(removing: .title)
    }
}

/// An empty or no-match card: the sentence and its note.
struct GuildMessageCard: View {
    let title: String
    let note: String

    var body: some View {
        VStack(alignment: .leading, spacing: 6) {
            Text(title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
            Text(note).font(ScaledType.meta).foregroundStyle(FSColor.muted)
        }
        .raidCardBox(16).quietCard()
        .accessibilityElement(children: .combine)
    }
}

/// A list's skeleton rows at their real height.
struct GuildSkeletonRows: View {
    let count: Int
    let height: CGFloat

    var body: some View {
        VStack(spacing: 0) {
            ForEach(0..<count, id: \.self) { index in
                HStack(spacing: 12) {
                    RoundedRectangle(cornerRadius: FSRadius.pill).frame(width: 32, height: 32)
                    VStack(alignment: .leading, spacing: 6) {
                        RoundedRectangle(cornerRadius: 4).frame(width: 120, height: 15)
                        RoundedRectangle(cornerRadius: 4).frame(width: 180, height: 12)
                    }
                    Spacer()
                    RoundedRectangle(cornerRadius: 4).frame(width: 70, height: 12)
                }
                .padding(.horizontal, 14)
                .frame(height: height)
                if index < count - 1 { Divider().overlay(FSColor.lineSoft) }
            }
        }
        .foregroundStyle(FSColor.lineSoft)
        .quietCard()
        .accessibilityHidden(true)
    }
}

/// The quiet footer line under a guild screen.
struct GuildFooter: View {
    let text: String

    var body: some View {
        Text(text).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            .fixedSize(horizontal: false, vertical: true)
            .padding(.top, 6)
    }
}
