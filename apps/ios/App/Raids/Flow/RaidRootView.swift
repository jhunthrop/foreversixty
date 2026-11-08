import DesignSystem
import SwiftUI

/// What the Raid tab's root hands back.
struct RaidRootActions {
    var raid = NextRaidActions()
    var openPage: (GuildPage) -> Void = { _ in }
    var openLink: (URL) -> Void = { _ in }
    var apply: (RecruitingRow) -> Void = { _ in }
    var claim: () -> Void = {}
    var findGuild: () -> Void = {}
    var signIn: () -> Void = {}
    var retry: () -> Void = {}
}

/// The Raid tab (flows spec 4.2; mocks ios-raid, -officer, -live, -no-raid, -none, -leveler, -out):
/// the next raid with your answer, an officer's Needs you, then the guild's pages; raid night leads
/// while a raid is live. No guild: the guilds recruiting your spec; under 60, when raids start.
struct RaidRootView: View {
    let content: RaidRootContent
    var actions = RaidRootActions()

    static let skeletons = 4
    static let skeletonHeight: CGFloat = 56

    var body: some View {
        RaidFlowScroll {
            switch content {
            case .loading: loading
            case .failed: failed
            case .signedOut: signedOut
            case let .leveler(toGo): leveler(toGo)
            case let .noGuild(recruiting): noGuild(recruiting)
            case .waiting:
                FSSection(Copy.raidsNext) { SmallText(Copy.raidsWaiting).padding(.top, RaidFlowSpacing.line) }
                    .padding(.top, RaidFlowSpacing.section)
            case let .guild(guild): RaidGuildView(content: guild, actions: actions)
            }
        }
    }

    private var loading: some View {
        VStack(spacing: FSSpace.space3) {
            ForEach(0 ..< Self.skeletons, id: \.self) { _ in FSSkeleton(height: Self.skeletonHeight) }
        }
        .padding(.top, RaidFlowSpacing.section)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.homeAnswerLoading)
    }

    private var failed: some View {
        FSRow(title: Copy.raidsFailed, inset: 0, trailPlacement: .below) {
            EmptyView()
        } trail: {
            Button(Copy.commonTryAgain, action: actions.retry).buttonStyle(FSButtonStyle(fullWidth: false, small: true))
        }
        .padding(.top, RaidFlowSpacing.section)
    }

    @ViewBuilder
    private var signedOut: some View {
        PageHeader(title: Copy.homeFindGuildTitle) { SmallText(Copy.raidOutLine) }
        FSButton(Copy.signinButton, kind: .primary, action: actions.signIn).padding(.top, RaidFlowSpacing.buttons)
    }

    @ViewBuilder
    private func leveler(_ toGo: String) -> some View {
        PageHeader(title: Copy.raidLevelerTitle) { SmallText(toGo) }
        SmallText(Copy.raidLevelerLine).padding(.top, RaidFlowSpacing.line)
        FSButton(Copy.homeFindGuildButton, action: actions.findGuild).padding(.top, RaidFlowSpacing.buttons)
    }

    @ViewBuilder
    private func noGuild(_ recruiting: RecruitingContent) -> some View {
        PageHeader(title: Copy.homeFindGuildTitle) { SmallText(Copy.raidRecruitMeta) }
        FSSection(recruiting.heading) {
            if recruiting.rows.isEmpty {
                SmallText(Copy.raidRecruitNone).padding(.top, RaidFlowSpacing.line)
            }
            VStack(spacing: 0) {
                ForEach(recruiting.rows) { row in
                    FSRow(title: row.name, sub: row.line, inset: 0, trailPlacement: .below) {
                        EmptyView()
                    } trail: {
                        Button(Copy.raidRecruitApply) { actions.apply(row) }
                            .buttonStyle(FSButtonStyle(fullWidth: false, small: true))
                    }
                }
            }
        }
        .padding(.top, RaidFlowSpacing.section)
        FSSection(Copy.raidClaimTitle) {
            SmallText(Copy.raidClaimLine).padding(.top, RaidFlowSpacing.line)
            FSButton(Copy.raidClaimButton, action: actions.claim).padding(.top, RaidFlowSpacing.buttons)
        }
        .padding(.top, RaidFlowSpacing.section)
    }
}

/// The Raid tab in a guild: raid night or the next raid, Needs you, the guild's pages.
struct RaidGuildView: View {
    let content: RaidGuildContent
    let actions: RaidRootActions

    var body: some View {
        if let stamp = content.offlineStamp { OfflineStrip(text: Copy.offlineStrip(time: stamp)) }
        if let stale = content.staleStamp { StaleStrip(time: stale, retry: actions.retry) }
        if let night = content.night {
            FSSection(Copy.raidNightTitle) {
                RaidLiveCard(content: night.live, open: nil).padding(.top, RaidFlowSpacing.block)
                RaidRoleRows(roles: night.roles)
            }
            .padding(.top, RaidFlowSpacing.section)
            if let ready = content.ready {
                FSSection(Copy.raidReadyTitle) { ReadyChecklistView(content: ready, open: actions.openLink) }
                    .padding(.top, RaidFlowSpacing.section)
            }
        } else {
            FSSection(Copy.raidsNext) {
                if let next = content.next {
                    NextRaidCardView(content: next, actions: actions.raid).padding(.top, RaidFlowSpacing.block)
                } else {
                    SmallText(Copy.raidNoRaid).padding(.top, RaidFlowSpacing.line)
                }
            }
            .padding(.top, RaidFlowSpacing.section)
        }
        if let needs = content.needsYou {
            FSSection(Copy.raidNeedsYou, meta: needs.meta) { OfficerItemsView(items: needs.items, actions: actions.raid) }
                .padding(.top, RaidFlowSpacing.section)
        }
        GuildPagesSection(content: content.pages, open: actions.openPage).padding(.top, RaidFlowSpacing.section)
    }
}

/// The guild's pages as rows, titled with the guild's name.
struct GuildPagesSection: View {
    let content: GuildPagesContent
    let open: (GuildPage) -> Void

    var body: some View {
        FSSection(content.guild) {
            VStack(spacing: 0) {
                ForEach(content.rows) { row in
                    Button { open(row.page) } label: {
                        FSRow(title: row.page.title, sub: row.sub, inset: 0) { EmptyView() }
                    }
                    .buttonStyle(RowPressStyle())
                }
            }
        }
    }
}
