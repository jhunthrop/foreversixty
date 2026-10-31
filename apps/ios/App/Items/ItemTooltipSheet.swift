import DesignSystem
import DesignTokens
import SwiftUI

/// The item tooltip as a sheet (spec F2.6): opens at the medium detent, grows to large. It reads
/// only the build's static files, cached per build, so it needs no API request.
struct ItemTooltipSheet: View {
    let request: ItemTooltipRequest
    let catalog: ItemCatalog

    @State private var phase: Phase = .loading

    enum Phase: Equatable {
        case loading
        case loaded(ItemTooltipModel)
        case failed
    }

    var body: some View {
        ItemTooltipContent(request: request, phase: phase, retry: { Task { await load() } })
            .presentationDetents([.medium, .large])
            .presentationDragIndicator(.visible)
            .presentationBackground(FSColor.raised)
            .task { await load() }
    }

    private func load() async {
        phase = .loading
        do {
            let sources = try await catalog.sources(for: request)
            phase = .loaded(.build(request, sources: sources))
        } catch {
            phase = .failed
        }
    }
}

/// The sheet's body, pure over its phase so the snapshot tests render it directly.
struct ItemTooltipContent: View {
    let request: ItemTooltipRequest
    let phase: ItemTooltipSheet.Phase
    var retry: () -> Void = {}

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 16) {
                switch phase {
                case .loading:
                    header(look: request.fallback, meta: ItemQuality(value: request.fallback.quality).name)
                    Text(Copy.itemsLoading).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                case let .loaded(model):
                    // The panel opens with the name, as the game's tooltip does; the header does not repeat it.
                    header(look: model.look, meta: model.headerMeta, showsName: false)
                    ItemTooltipPanel(model: model)
                    sourceSection(model)
                    nextBest(model)
                    if let footer = model.pickFooter {
                        Text(footer).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    }
                case .failed:
                    header(look: request.fallback, meta: ItemQuality(value: request.fallback.quality).name)
                    Text(Copy.itemsFailed).font(ScaledType.body).foregroundStyle(FSColor.wipe)
                        .accessibilityAddTraits(.isStaticText)
                    SecondaryButton(title: Copy.commonTryAgain, action: retry)
                }
            }
            .padding(.horizontal, 16)
            .padding(.top, 24)
            .padding(.bottom, 32)
        }
        .background(FSColor.raised)
    }

    private func header(look: HomeItemLook, meta: String, showsName: Bool = true) -> some View {
        HStack(alignment: showsName ? .top : .center, spacing: 12) {
            ItemIconView(icon: look.icon, quality: look.quality, build: request.build, size: 56)
            VStack(alignment: .leading, spacing: 2) {
                if showsName {
                    Text(look.name ?? Copy.itemsNotInData)
                        .font(ScaledType.sec)
                        .foregroundStyle(ItemQuality(value: look.quality).text)
                        .fixedSize(horizontal: false, vertical: true)
                }
                Text(meta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
        }
        .accessibilityElement(children: .combine)
        .accessibilityAddTraits(.isHeader)
    }

    @ViewBuilder
    private func sourceSection(_ model: ItemTooltipModel) -> some View {
        if let source = model.source {
            VStack(alignment: .leading, spacing: 6) {
                MicroLabel(Copy.tooltipSource)
                Text(source).font(ScaledType.body).foregroundStyle(FSColor.strong)
            }
            .accessibilityElement(children: .combine)
        }
    }

    @ViewBuilder
    private func nextBest(_ model: ItemTooltipModel) -> some View {
        if model.pickFooter != nil || !model.alternatives.isEmpty {
            VStack(alignment: .leading, spacing: 0) {
                MicroLabel(Copy.tooltipNextBest).padding(.bottom, 8)
                if model.alternatives.isEmpty {
                    Text(Copy.tooltipNothingElse).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                }
                ForEach(model.alternatives) { alt in
                    HStack(spacing: 12) {
                        ItemIconView(icon: alt.look.icon, quality: alt.look.quality, build: request.build, size: 36)
                        VStack(alignment: .leading, spacing: 2) {
                            Text(alt.look.name ?? Copy.itemsNotInData).font(ScaledType.body)
                                .foregroundStyle(ItemQuality(value: alt.look.quality).text)
                            Text(alt.meta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                            if let source = alt.source {
                                Text(source).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                            }
                        }
                    }
                    .padding(.vertical, 10)
                    .accessibilityElement(children: .combine)
                    Divider().overlay(FSColor.lineSoft)
                }
            }
        }
    }
}

/// The game-tooltip panel: near-black, `#3a4256` hairline, the client's line order.
struct ItemTooltipPanel: View {
    let model: ItemTooltipModel

    var body: some View {
        VStack(alignment: .leading, spacing: 4) {
            Text(model.look.name ?? Copy.itemsNotInData).foregroundStyle(ItemQuality(value: model.look.quality).text)
                .font(FSFont.bodyBold(17))
            if !model.known {
                line(Copy.itemsNotInData, color: FSColor.muted)
            }
            if let binds = model.binds { line(binds) }
            if model.unique { line(Copy.tooltipUnique) }
            if model.slotName != nil || model.typeName != nil {
                ViewThatFits(in: .horizontal) {
                    HStack {
                        Text(model.slotName ?? "")
                        Spacer(minLength: 12)
                        Text(model.typeName ?? "")
                    }
                    VStack(alignment: .leading, spacing: 4) {
                        Text(model.slotName ?? "")
                        if let type = model.typeName { Text(type) }
                    }
                }
                .font(ScaledType.tooltip)
                .foregroundStyle(FSColor.strong)
            }
            if let armor = model.armor { line(armor) }
            if let damage = model.weaponDamage, let speed = model.weaponSpeed {
                ViewThatFits(in: .horizontal) {
                    HStack { Text(damage); Spacer(minLength: 12); Text(speed) }
                    VStack(alignment: .leading, spacing: 4) { Text(damage); Text(speed) }
                }
                .font(ScaledType.tooltip).foregroundStyle(FSColor.strong)
            }
            if let dps = model.weaponDPS { line(dps) }
            ForEach(model.stats, id: \.self) { line($0) }
            if let effect = model.effect { line(effect, color: FSColor.parseGreen) }
            if let set = model.setName {
                line(set, color: FSColor.parseGold).padding(.top, 6)
                ForEach(model.setBonuses, id: \.self) { line($0, color: FSColor.muted) }
            }
            if let requires = model.requiresLevel { line(requires).padding(.top, 6) }
            if let level = model.itemLevel { line(level, color: FSColor.parseGold) }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(.vertical, 12)
        .padding(.horizontal, 14)
        .background(Color(tokenHex: 0x0A0D14), in: .rect(cornerRadius: FSRadius.control))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(Color(tokenHex: 0x3A4256)))
        .accessibilityElement(children: .combine)
    }

    private func line(_ text: String, color: Color = FSColor.strong) -> some View {
        Text(text).font(ScaledType.tooltip).foregroundStyle(color).fixedSize(horizontal: false, vertical: true)
    }
}
