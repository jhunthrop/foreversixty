import DesignSystem
import Models
import Networking
import Observation
import SwiftUI

/// Loads `GET /v1/me/raid-nights/{raid}` for Your raid night (spec F4.22).
@MainActor
@Observable
final class RaidNightModel {
    enum State: Equatable {
        case loading
        case loaded(RaidNightSummary)
        case failed
    }

    private(set) var state: State
    private(set) var offlineSince: Date?
    let raidID: Int64
    private let service: any RaidService
    private let cache: any RaidCache
    private let now: @Sendable () -> Date

    init(raid: Int64, service: any RaidService, cache: any RaidCache, now: @escaping @Sendable () -> Date, initial: State = .loading) {
        raidID = raid
        self.service = service
        self.cache = cache
        self.now = now
        state = initial
    }

    private var key: String { "night-\(raidID)" }

    func load() async {
        if case .loading = state, let cached = cache.read(RaidNightSummary.self, key: key) { state = .loaded(cached.value) }
        do {
            let night = try await service.night(raid: raidID)
            cache.write(night, key: key, at: now())
            offlineSince = nil
            state = .loaded(night)
        } catch {
            if SignupRefusal.isOffline(error), let cached = cache.read(RaidNightSummary.self, key: key) {
                offlineSince = cached.savedAt
                state = .loaded(cached.value)
            } else if case .loaded = state {
                return
            } else {
                state = .failed
            }
        }
    }
}

/// F4.22 Your raid night: kills, wipes and bests, loot awarded to you, personal bests, coach notes.
struct RaidNightView: View {
    @State private var model: RaidNightModel
    let clock: RaidClock
    let now: () -> Date
    /// Lane F10: pushes the report in the app; nil while the reports switch is off.
    let openReport: ((String) -> Void)?
    @Environment(\.openURL) private var openURL

    init(model: @autoclosure () -> RaidNightModel, clock: RaidClock, now: @escaping () -> Date, openReport: ((String) -> Void)? = nil) {
        _model = State(wrappedValue: model())
        self.clock = clock
        self.now = now
        self.openReport = openReport
    }

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                if let since = model.offlineSince {
                    OfflineStrip(text: Copy.offlineStrip(time: clock.stripTime(since, now: now())))
                }
                VStack(alignment: .leading, spacing: 22) { content }.padding(16)
            }
        }
        .background(FSColor.bg)
        .navigationTitle(Copy.nightTitle)
        .navigationBarTitleDisplayMode(.large)
        .task { await model.load() }
        .refreshable { await model.load() }
    }

    @ViewBuilder
    private var content: some View {
        switch model.state {
        case .loading:
            VStack(alignment: .leading, spacing: 14) {
                RoundedRectangle(cornerRadius: 4).frame(width: 220, height: 13)
                RoundedRectangle(cornerRadius: FSRadius.panel).frame(height: 110)
                RoundedRectangle(cornerRadius: FSRadius.panel).frame(height: 160)
            }
            .foregroundStyle(FSColor.lineSoft)
            .accessibilityHidden(true)
        case .failed:
            ErrorRegion(message: Copy.nightFailed) { Task { await model.load() } }
        case let .loaded(night):
            loaded(night)
        }
    }

    @ViewBuilder
    private func loaded(_ night: RaidNightSummary) -> some View {
        Text(Copy.nightMeta(zone: night.zoneName, weekday: clock.weekday(night.startsAt), date: clock.dayMonth(night.startsAt), guild: night.guildName))
            .font(ScaledType.meta).foregroundStyle(FSColor.muted)
        ViewThatFits(in: .horizontal) {
            HStack(alignment: .center, spacing: 16) {
                figures(night)
                Spacer(minLength: 8)
                who(night, alignment: .trailing)
            }
            // Large text: the figures stack, the character under them.
            VStack(alignment: .leading, spacing: 12) {
                VStack(alignment: .leading, spacing: 8) { figures(night) }
                who(night, alignment: .leading)
            }
        }
        .raidCardBox(16).primaryCard()
        if !night.loot.isEmpty {
            section(Copy.nightLoot) {
                ForEach(night.loot, id: \.itemId) { loot in lootRow(loot) }
            }
        }
        if !night.bests.isEmpty {
            section(Copy.nightBests) {
                ForEach(night.bests, id: \.boss) { best in bestRow(best) }
            }
        }
        if let notes = night.coachNotes, !notes.isEmpty {
            section(Copy.nightCoach) {
                ForEach(Array(notes.enumerated()), id: \.offset) { index, note in
                    HStack(alignment: .top, spacing: 10) {
                        Text("\(index + 1)").font(ScaledType.mono).foregroundStyle(FSColor.gold)
                        Text(note).font(ScaledType.body).foregroundStyle(FSColor.text)
                            .fixedSize(horizontal: false, vertical: true)
                    }
                    .padding(14)
                    .frame(maxWidth: .infinity, alignment: .leading)
                }
            }
        }
        if let openReport {
            Button { openReport(night.reportId) } label: {
                HStack {
                    Text(Copy.nightOpenReport).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                    Spacer()
                    SimChevron()
                }
                .padding(.horizontal, 14)
                .frame(minHeight: 52)
                .contentShape(Rectangle())
            }
            .buttonStyle(RowPressStyle())
            .raidCardBox(0).quietCard()
        }
        Button(Copy.nightReport) { openURL(RaidLinks.report(night.reportId)) }.buttonStyle(GoldLinkStyle())
    }

    @ViewBuilder
    private func figures(_ night: RaidNightSummary) -> some View {
        figure(night.kills, Copy.nightKill(n: night.kills))
        figure(night.wipes, Copy.nightWipes(n: night.wipes))
        figure(night.newBests, Copy.nightBest(n: night.newBests))
    }

    private func who(_ night: RaidNightSummary, alignment: HorizontalAlignment) -> some View {
        VStack(alignment: alignment, spacing: 2) {
            NameText(text: night.character.name, font: ScaledType.name, color: ClassColor.of(night.character.className))
            Text(Copy.raidsWho(spec: night.character.specName, role: RaidRole.label(night.character.role)))
                .font(ScaledType.meta).foregroundStyle(FSColor.muted)
        }
    }

    /// The one figure: the number in `stat`, its word in meta (the plural carries the count).
    private func figure(_ value: Int, _ words: String) -> some View {
        let word = words.replacingOccurrences(of: "\(value) ", with: "")
        return VStack(alignment: .leading, spacing: 0) {
            Text("\(value)").font(ScaledType.stat).foregroundStyle(FSColor.strong)
            Text(word).font(ScaledType.meta).foregroundStyle(FSColor.muted)
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(words)
    }

    private func section(_ title: String, @ViewBuilder rows: () -> some View) -> some View {
        VStack(alignment: .leading, spacing: 8) {
            SectionLabel(text: title)
            VStack(spacing: 0) { rows() }.raidCardBox(0).quietCard()
        }
    }

    private func lootRow(_ loot: RaidNightLoot) -> some View {
        HStack(alignment: .top, spacing: 12) {
            ZStack {
                GameIcon(stem: loot.icon, size: 44, border: ItemQuality(value: loot.quality).border, borderWidth: 1.5)
            }
            .shadow(color: ItemQuality(value: loot.quality).border.opacity(0.22), radius: 5)
            VStack(alignment: .leading, spacing: 2) {
                Text(loot.name).font(ScaledType.bodyStrong).foregroundStyle(ItemQuality(value: loot.quality).text)
                Text(Copy.nightLootLine(slot: loot.slot, boss: loot.boss, officer: loot.awardedBy))
                    .font(ScaledType.meta).foregroundStyle(FSColor.muted)
                if let replaces = loot.replaces {
                    Text(Copy.homeUpgradesReplaces(item: replaces.name, n: replaces.itemLevel))
                        .font(ScaledType.meta).foregroundStyle(FSColor.muted)
                }
            }
        }
        .padding(14)
        .frame(maxWidth: .infinity, minHeight: 72, alignment: .leading)
        .accessibilityElement(children: .combine)
    }

    private func bestRow(_ best: RaidNightBest) -> some View {
        HStack(alignment: .top) {
            VStack(alignment: .leading, spacing: 2) {
                Text(best.boss).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                Text(Copy.nightBestLine(time: Self.killTime(best.killMs))).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
            Spacer()
            VStack(alignment: .trailing, spacing: 2) {
                Text(Copy.raidsNightDps(dps: best.dps.formatted(.number.precision(.fractionLength(1)))))
                    .font(ScaledType.mono).foregroundStyle(FSColor.strong)
                Text(Copy.nightPercentile(ordinal: Self.ordinal(best.percentile)))
                    .font(ScaledType.meta).foregroundStyle(GameColors.parse(best.percentile))
            }
        }
        .padding(14)
        .accessibilityElement(children: .combine)
    }

    /// `4:28` from milliseconds.
    static func killTime(_ ms: Int64) -> String {
        let seconds = Int(ms / 1000)
        return String(format: "%d:%02d", seconds / 60, seconds % 60)
    }

    /// The English ordinal (`91st`, `92nd`, `93rd`, `94th`, `11th`).
    static func ordinal(_ n: Int) -> String {
        let formatter = NumberFormatter()
        formatter.numberStyle = .ordinal
        formatter.locale = Locale(identifier: "en_US")
        return formatter.string(from: NSNumber(value: n)) ?? "\(n)"
    }
}
