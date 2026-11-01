import DesignSystem
import Models
import SwiftUI

/// F11.2 Filters: a sheet at the large detent, live like the BiS picker (every choice applies at
/// once). Signed in there is no class or spec filter: the character sets it.
struct FiltersSheet: View {
    let filters: RankingsFilters
    let setup: RankingsSetup
    let apply: (RankingsFilters) -> Void
    let reset: () -> Void
    @State private var showsEncounters = false

    var body: some View {
        NavigationStack {
            ScrollView {
                FiltersView(filters: filters, setup: setup, apply: apply, reset: reset) { showsEncounters = true }
            }
            .background(GameColors.sheet)
            .toolbar(.hidden, for: .navigationBar)
            .navigationDestination(isPresented: $showsEncounters) {
                ScrollView {
                    EncounterList(encounters: setup.encounters, selected: filters.encounter) { slug in
                        var next = filters
                        next.encounter = slug
                        apply(next)
                        showsEncounters = false
                    }
                }
                .background(GameColors.sheet)
                .navigationTitle(Copy.rankFilterEncounter)
                .navigationBarTitleDisplayMode(.inline)
            }
        }
        .presentationDetents([.large])
        .presentationDragIndicator(.visible)
        .presentationBackground(GameColors.sheet)
        .sensoryFeedback(.selection, trigger: filters)
    }
}

/// The sheet's body, pure over the filters so the snapshot tests render it directly.
struct FiltersView: View {
    let filters: RankingsFilters
    let setup: RankingsSetup
    let apply: (RankingsFilters) -> Void
    let reset: () -> Void
    let openEncounters: () -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 18) {
            Text(Copy.rankFilters.uppercased()).font(ScaledType.sec).foregroundStyle(FSColor.strong)
                .accessibilityAddTraits(.isHeader)
            section(Copy.rankFilterEncounter) { encounterRow }
            section(Copy.rankFilterMeasure) {
                chips(RankMetric.allCases.map { ($0.rawValue, RankingsCopy.metric($0)) }, selected: filters.metric.rawValue) { id in
                    change { $0.metric = RankMetric(rawValue: id) ?? .dps }
                }
            }
            section(Copy.rankFilterRuleset) {
                chips((RankingsCopy.rulesets + [""]).map { ($0, RankingsCopy.ruleset($0)) }, selected: filters.ruleset) { id in
                    change { $0.ruleset = id }
                }
            }
            section(Copy.rankFilterRegion) {
                chips([("", Copy.rankRegionAll)] + RankingsCopy.regions.map { ($0, $0.uppercased()) }, selected: filters.region) { id in
                    change { $0.region = id }
                }
            }
            section(Copy.rankFilterPhase) {
                chips(RankingsCopy.phaseChoices(setup.phases).map { ($0.id, $0.label) }, selected: filters.phase) { id in
                    change { $0.phase = id }
                }
            }
            section(Copy.rankFilterFaction) {
                chips([("", Copy.rankFactionBoth)] + RankingsCopy.factions.map { ($0, RankingsCopy.faction($0)) }, selected: filters.faction) { id in
                    change { $0.faction = id }
                }
            }
            Toggle(isOn: Binding(get: { filters.todayOnly }, set: { on in change { $0.todayOnly = on } })) {
                Text(Copy.rankFilterToday).font(ScaledType.body).foregroundStyle(FSColor.text)
            }
            .tint(FSColor.gold)
            .frame(minHeight: 44)
            SecondaryButton(title: Copy.rankReset, action: reset)
        }
        .padding(16)
    }

    private func change(_ edit: (inout RankingsFilters) -> Void) {
        var next = filters
        edit(&next)
        apply(next)
    }

    private var encounterRow: some View {
        Button(action: openEncounters) {
            HStack {
                Text(setup.encounterName(filters.encounter)).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                    .fixedSize(horizontal: false, vertical: true)
                Spacer()
                Image(systemName: "chevron.right").font(ScaledType.meta.weight(.semibold)).foregroundStyle(FSColor.muted)
                    .accessibilityHidden(true)
            }
            .padding(.horizontal, 14)
            .frame(minHeight: 52)
            .background(GameColors.controlFill, in: .rect(cornerRadius: FSRadius.control))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.line))
            .contentShape(Rectangle())
        }
        .buttonStyle(.plain)
        .accessibilityLabel("\(Copy.rankFilterEncounter), \(setup.encounterName(filters.encounter))")
    }

    private func chips(_ options: [(id: String, label: String)], selected: String, choose: @escaping (String) -> Void) -> some View {
        WrapLayout(spacing: 8) {
            ForEach(options, id: \.id) { option in
                let isOn = option.id == selected
                Button { choose(option.id) } label: {
                    Text(option.label)
                        .font(isOn ? ScaledType.bodyStrong : ScaledType.meta)
                        .foregroundStyle(isOn ? FSColor.gold : FSColor.nav)
                        .padding(.horizontal, 12)
                        .frame(minHeight: 44)
                        .pillSurface(fill: isOn ? FSColor.gold.opacity(0.12) : .clear, line: isOn ? FSColor.gold : FSColor.line)
                }
                .buttonStyle(.plain)
                .accessibilityAddTraits(isOn ? .isSelected : [])
            }
        }
    }

    private func section(_ title: String, @ViewBuilder _ body: () -> some View) -> some View {
        VStack(alignment: .leading, spacing: 8) {
            MicroLabel(title).accessibilityAddTraits(.isHeader)
            body()
        }
    }
}

/// The encounters with ranked kills (W3-C6): name, then `zone · size players · kills`.
struct EncounterList: View {
    let encounters: [RankedEncounter]
    let selected: String
    let choose: (String) -> Void

    var body: some View {
        VStack(spacing: 0) {
            ForEach(Array(encounters.enumerated()), id: \.element.slug) { index, encounter in
                if index > 0 { Divider().overlay(FSColor.lineSoft) }
                Button { choose(encounter.slug) } label: {
                    HStack(spacing: 12) {
                        VStack(alignment: .leading, spacing: 2) {
                            Text(encounter.name).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                                .fixedSize(horizontal: false, vertical: true)
                            Text(RankingsCopy.encounterMeta(encounter)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                                .fixedSize(horizontal: false, vertical: true)
                        }
                        .frame(maxWidth: .infinity, alignment: .leading)
                        if encounter.slug == selected {
                            Image(systemName: "checkmark").foregroundStyle(FSColor.gold).accessibilityHidden(true)
                        }
                    }
                    .padding(.horizontal, 16)
                    .padding(.vertical, 10)
                    .frame(minHeight: 60)
                    .contentShape(Rectangle())
                }
                .buttonStyle(RowPressStyle())
                .accessibilityAddTraits(encounter.slug == selected ? .isSelected : [])
            }
        }
        .quietCard()
        .padding(16)
    }
}
