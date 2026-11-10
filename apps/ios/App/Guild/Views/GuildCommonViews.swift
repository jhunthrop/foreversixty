import DesignSystem
import SwiftUI

/// A pushed guild screen's frame: the band head with the title and an optional trailing control,
/// the offline and stale strips under the band, then the padded content (N3, as every pushed screen).
struct GuildScreenScroll<Content: View, Trailing: View>: View {
    let title: String
    let offlineStrip: String?
    @ViewBuilder let trailing: () -> Trailing
    @ViewBuilder let content: () -> Content
    /// Lane F15: the stale strip, when Forever Sixty did not answer.
    var stale: StaleStrip?

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                BandHead(title: title, lead: { EmptyView() }, facts: { EmptyView() }, trailing: trailing)
                if let offlineStrip { OfflineStrip(text: offlineStrip) }
                if let stale { stale }
                VStack(alignment: .leading, spacing: 16) {
                    content()
                }
                .padding(16)
            }
        }
        .screenGround()
    }
}

extension GuildScreenScroll where Trailing == EmptyView {
    init(title: String, offlineStrip: String?, stale: StaleStrip? = nil, @ViewBuilder content: @escaping () -> Content) {
        self.init(title: title, offlineStrip: offlineStrip, trailing: { EmptyView() }, content: content, stale: stale)
    }
}

extension View {
    /// A pushed guild screen's bar (N3): the back control on the clear bar; the band head names the
    /// screen, and the title still names the next screen's back button.
    func guildPushedBar(_ title: String) -> some View {
        pushedBandBar(title)
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
                    RoundedRectangle(cornerRadius: FSRadius.legacyPill).frame(width: 32, height: 32)
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
        Text(text).font(ScaledType.meta).foregroundStyle(FSColor.mutedGround)
            .fixedSize(horizontal: false, vertical: true)
            .padding(.top, 6)
    }
}
