import DesignSystem
import Models
import SwiftUI

/// F4.1 the Raids tab's content, pure over its state so snapshot tests draw every state.
struct RaidsTabContentView: View {
    @Environment(\.dynamicTypeSize) private var typeSize
    let state: RaidsTabState
    let clock: RaidClock
    let now: () -> Date
    /// The next raid's sign-up, owned by the root so a press survives a redraw.
    let nextController: SignupController?
    let openRaid: (RaidView, Bool) -> Void
    let openNight: (Int64) -> Void
    let openWeb: (URL) -> Void
    let retry: () -> Void

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                if case let .loaded(content) = state, let since = content.offlineSince {
                    OfflineStrip(text: Copy.offlineStrip(time: clock.stripTime(since, now: now())))
                }
                VStack(alignment: .leading, spacing: 22) { stateContent }
                    .padding(16)
            }
        }
        .background(Theme.Palette.background)
    }

    @ViewBuilder
    private var stateContent: some View {
        switch state {
        case .loading:
            RaidsTabSkeleton()
        case .signedOut:
            message(title: Copy.raidsSignedOutTitle, body: Copy.raidsSignedOutBody, note: nil)
        case let .noGuild(name):
            message(title: Copy.raidsNoGuildTitle, body: Copy.raidsNoGuildBody(name: name), note: Copy.raidsNoGuildNote)
        case let .waiting(head, character):
            GuildHeadView(head: head)
            VStack(alignment: .leading, spacing: 12) {
                HStack {
                    WhoLine(name: character.name, className: character.className ?? "", specName: character.spec, specIcon: nil, role: "", trailing: nil)
                    StatusTag(text: Copy.raidsWaitingTag, color: Theme.Palette.accent)
                }
                Text(Copy.raidsWaiting).font(Theme.Typography.bodyText).foregroundStyle(Theme.Palette.text)
            }
            .primaryCard()
        case .failed:
            ErrorRegion(message: Copy.raidsFailed, retry: retry)
        case let .loaded(content):
            loaded(content)
        }
    }

    private func message(title: String, body: String, note: String?) -> some View {
        VStack(alignment: .leading, spacing: 12) {
            Image(systemName: "calendar")
                .font(.system(size: 20))
                .foregroundStyle(Theme.Palette.accent)
                .frame(width: 44, height: 44)
                .background(Theme.Palette.accent.opacity(0.12), in: .rect(cornerRadius: Theme.Radius.control))
                .accessibilityHidden(true)
            Text(title.uppercased()).font(Theme.Typography.h2).foregroundStyle(Theme.Palette.strong)
                .accessibilityAddTraits(.isHeader)
            Text(body).font(Theme.Typography.bodyText).foregroundStyle(Theme.Palette.text)
            if let note { Text(note).font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.muted) }
        }
        .padding(.top, 24)
    }

    @ViewBuilder
    private func loaded(_ content: RaidsTabContent) -> some View {
        let planned = PlannedRaids(content.planned, now: now())
        GuildHeadView(head: content.head)
        if let next = planned.next, let controller = nextController, controller.raid.id == next.id {
            NextRaidCard(controller: controller, clock: clock, now: now, currentKey: content.character.key) { roster in
                openRaid(controller.raid, roster)
            }
        } else if planned.next == nil {
            VStack(alignment: .leading, spacing: 6) {
                Text(Copy.raidsEmpty).font(Theme.Typography.bodyStrong).foregroundStyle(Theme.Palette.strong)
                Text(Copy.raidsEmptyNote).font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.muted)
            }
            .primaryCard()
        }
        if let waiting = content.waiting {
            Button { if let url = content.head.webURL { openWeb(url) } } label: {
                (typeSize.isAccessibilitySize
                    ? AnyLayout(VStackLayout(alignment: .leading, spacing: 8))
                    : AnyLayout(HStackLayout(spacing: 12))) {
                    Image(systemName: "person.badge.clock").foregroundStyle(Theme.Palette.accent)
                        .frame(width: 36, height: 36)
                        .background(Theme.Palette.accent.opacity(0.12), in: .rect(cornerRadius: Theme.Radius.control))
                        .accessibilityHidden(true)
                    VStack(alignment: .leading, spacing: 2) {
                        Text(Copy.raidsPending(n: waiting)).font(Theme.Typography.bodyStrong).foregroundStyle(Theme.Palette.strong)
                        Text(Copy.raidsPendingLine).font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.muted)
                    }
                    Spacer()
                }
                .padding(14)
                .contentShape(Rectangle())
            }
            .buttonStyle(RowPressStyle())
            .quietCard()
        }
        if !planned.also.isEmpty {
            VStack(alignment: .leading, spacing: 8) {
                SectionLabel(text: Copy.raidsPlanned, trailing: "\(planned.also.count)")
                VStack(spacing: 0) {
                    ForEach(planned.also, id: \.id) { raid in
                        PlannedRow(raid: raid, clock: clock, now: now(), currentKey: content.character.key) { openRaid(raid, false) }
                        if raid.id != planned.also.last?.id { Divider().overlay(Theme.Palette.lineSoft) }
                    }
                }
                .quietCard()
            }
        }
        if !content.nights.isEmpty {
            VStack(alignment: .leading, spacing: 8) {
                SectionLabel(text: Copy.raidsNights)
                VStack(spacing: 0) {
                    ForEach(content.nights, id: \.id) { night in
                        NightRow(night: night, clock: clock) { if let id = night.raidId { openNight(id) } }
                        if night.id != content.nights.last?.id { Divider().overlay(Theme.Palette.lineSoft) }
                    }
                }
                .quietCard()
            }
        }
    }
}

/// The guild head: crest 40, the name, `{Faction} · {Region} {Ruleset} · {n} on the roster`.
struct GuildHeadView: View {
    let head: GuildHead
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let layout = typeSize.isAccessibilitySize
            ? AnyLayout(VStackLayout(alignment: .leading, spacing: 8))
            : AnyLayout(HStackLayout(spacing: 12))
        layout {
            GuildCrest(url: head.crestURL, faction: head.faction)
            VStack(alignment: .leading, spacing: 2) {
                Text(head.name.uppercased()).font(Theme.Typography.cardHead).foregroundStyle(Theme.Palette.strong)
                    .accessibilityAddTraits(.isHeader)
                if let line = head.line {
                    Text(line).font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.muted)
                }
            }
        }
        .accessibilityElement(children: .combine)
    }
}

/// The guild crest (circular 40), else the faction mark.
struct GuildCrest: View {
    let url: URL?
    let faction: String?

    var body: some View {
        ZStack {
            Circle().fill(Theme.Palette.raised)
            if let url {
                AsyncImage(url: url) { $0.resizable().scaledToFill() } placeholder: { factionMark }
                    .clipShape(Circle())
            } else {
                factionMark
            }
        }
        .frame(width: 40, height: 40)
        .overlay(Circle().stroke(Theme.Palette.line))
        .accessibilityHidden(true)
    }

    private var factionMark: some View {
        Image(systemName: "shield.lefthalf.filled")
            .foregroundStyle(faction == "alliance" ? Theme.Palette.alliance : Theme.Palette.horde)
    }
}

/// The next raid card (spec F4.1 item 2), the tab's one primary surface.
struct NextRaidCard: View {
    let controller: SignupController
    let clock: RaidClock
    let now: () -> Date
    let currentKey: String?
    /// Pushes the raid page; true scrolls it to the roster.
    let open: (Bool) -> Void
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let raid = controller.raid
        let facts = controller.facts
        VStack(alignment: .leading, spacing: 14) {
            Button { open(false) } label: {
                VStack(alignment: .leading, spacing: 14) {
                    // Large text: the date block above the title instead of beside it.
                    (typeSize.isAccessibilitySize
                        ? AnyLayout(VStackLayout(alignment: .leading, spacing: 10))
                        : AnyLayout(HStackLayout(alignment: .top, spacing: 12))) {
                        DateBlock(date: raid.startsAt, clock: clock)
                        VStack(alignment: .leading, spacing: 4) {
                            SectionLabel(text: Copy.raidsNext, color: Theme.Palette.accent)
                            Text(raid.title.uppercased()).font(Theme.Typography.cardHead).foregroundStyle(Theme.Palette.strong)
                            TimesLine(raid: raid, clock: clock, showsEnd: false)
                            CountdownLine(raid: raid, facts: facts, clock: clock, now: now)
                        }
                        .accessibilityElement(children: .combine)
                        .accessibilityLabel([clock.longDate(raid.startsAt), raid.title].joined(separator: ", "))
                    }
                    CountBlock(facts: facts)
                }
                .contentShape(Rectangle())
            }
            .buttonStyle(.plain)
            Rectangle().fill(Theme.Palette.lineSoft).frame(height: 1)
            SectionLabel(text: Copy.raidYourSignup)
            if let character = facts.viewerCharacters(currentKey: currentKey).first {
                SignupBlockView(controller: controller, character: character, showsNote: false, reservesStatus: false)
            }
            if controller.showsReminderAsk {
                ReminderAskView { turnOn in Task { await controller.answerReminderAsk(turnOn: turnOn) } }
            }
            Button { open(true) } label: {
                HStack(spacing: 4) {
                    Text(Copy.raidsWhoIsComing)
                    Image(systemName: "chevron.right").font(.system(size: 12, weight: .semibold)).accessibilityHidden(true)
                }
            }
            .buttonStyle(GoldLinkStyle())
            .padding(.top, -6)
        }
        .primaryCard()
        .sheet(item: Binding(get: { controller.specAsk }, set: { if $0 == nil { controller.closeSpecSheet() } })) { ask in
            if let character = controller.character(ask.characterKey) {
                SpecSheet(ask: ask, character: character, raidTitle: raid.title, date: clock.shortDate(raid.startsAt)) { option in
                    Task { await controller.pickSpec(option) }
                }
            }
        }
    }
}

/// An Also planned row (min 76): date block, title, when, count with a 64 pt bar, one chip.
struct PlannedRow: View {
    @Environment(\.dynamicTypeSize) private var typeSize
    let raid: RaidView
    let clock: RaidClock
    let now: Date
    let currentKey: String?
    let open: () -> Void

    var body: some View {
        let facts = RaidFacts(raid: raid, now: now)
        let large = typeSize.isAccessibilitySize
        Button(action: open) {
            HStack(alignment: .top, spacing: 12) {
                // The date block keeps its width; the text column takes the rest.
                DateBlock(date: raid.startsAt, clock: clock, width: large ? 56 : 48, height: 52)
                    .frame(width: large ? 56 : nil)
                VStack(alignment: .leading, spacing: 4) {
                    Text(raid.title).font(Theme.Typography.bodyStrong).foregroundStyle(Theme.Palette.strong).lineLimit(2)
                    Text(clock.when(raid.startsAt, now: now)).font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.muted)
                    HStack(spacing: 8) {
                        Text(raid.size == nil ? facts.countLine : facts.ratio)
                            .font(Theme.Typography.monoMeta).foregroundStyle(Theme.Palette.strong)
                        if let fill = facts.fill, !large { FillBar(fill: fill, width: 64) }
                    }
                    // Large text: the chip moves under the title.
                    if large { chip(facts.rowChip(currentKey: currentKey)) }
                }
                if !large {
                    Spacer(minLength: 8)
                    chip(facts.rowChip(currentKey: currentKey))
                }
            }
            .padding(14)
            .frame(minHeight: 76)
            .contentShape(Rectangle())
        }
        .buttonStyle(RowPressStyle())
        .accessibilityElement(children: .combine)
    }

    @ViewBuilder
    private func chip(_ chip: RowChip) -> some View {
        switch chip {
        case let .status(status): StatusTag(chip: status)
        case let .signed(text): StatusTag(text: text, color: Theme.Palette.kill)
        case .notSigned: StatusTag(text: Copy.raidChipNotSigned, color: Theme.Palette.accent)
        }
    }
}

/// A raid nights row: the night's title, its date, kills and wipes.
struct NightRow: View {
    let night: RaidNightRow
    let clock: RaidClock
    let open: () -> Void

    var body: some View {
        Button(action: open) {
            HStack {
                VStack(alignment: .leading, spacing: 2) {
                    Text(night.title).font(Theme.Typography.bodyStrong).foregroundStyle(Theme.Palette.strong)
                    Text(Copy.raidsNightMeta(date: clock.shortDate(night.createdAt), kills: Copy.nightKill(n: night.killCount), wipes: Copy.nightWipes(n: night.wipeCount)))
                        .font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.muted)
                }
                Spacer()
                Image(systemName: "chevron.right").foregroundStyle(Theme.Palette.muted).accessibilityHidden(true)
            }
            .padding(14)
            .frame(minHeight: 56)
            .contentShape(Rectangle())
        }
        .buttonStyle(RowPressStyle())
        .accessibilityElement(children: .combine)
    }
}

/// First load with nothing cached: the head, the card and a row at their real heights.
struct RaidsTabSkeleton: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 22) {
            HStack(spacing: 12) {
                Circle().frame(width: 40, height: 40)
                VStack(alignment: .leading, spacing: 6) { bar(160, 22); bar(220, 13) }
            }
            VStack(alignment: .leading, spacing: 14) {
                HStack(alignment: .top, spacing: 12) {
                    RoundedRectangle(cornerRadius: Theme.Radius.control).frame(width: 52, height: 56)
                    VStack(alignment: .leading, spacing: 6) { bar(70, 11); bar(180, 22); bar(200, 15); bar(150, 15) }
                }
                bar(240, 15)
                RoundedRectangle(cornerRadius: Theme.Radius.control).frame(height: 48)
                bar(90, 20)
            }
            .primaryCard()
        }
        .foregroundStyle(Theme.Palette.lineSoft)
        .accessibilityHidden(true)
    }

    private func bar(_ width: CGFloat, _ height: CGFloat) -> some View {
        RoundedRectangle(cornerRadius: 4).frame(width: width, height: height)
    }
}
