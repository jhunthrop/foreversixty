import DesignSystem
import DesignTokens
import Models
import SwiftUI

/// What a tap on Home asks the screen to do.
enum HomeIntent: Equatable {
    case openItem(HomeItemLook, sourceLine: String?, wornID: Int?)
    case openGear
    case action(HomeAction.Kind)
    case openNext(HomeNextRow.Kind)
    case retry
}

/// Home's content under the selector, pure over its inputs so the snapshot tests render every
/// state (spec F2.1 to F2.5).
struct HomeAnswerView: View {
    let phase: HomePhase
    let presenter: HomePresenter
    let build: String
    var offlineStamp: String?
    var send: (HomeIntent) -> Void = { _ in }

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            if let offlineStamp { OfflineStrip(text: Copy.offlineStrip(time: offlineStamp)) }
            switch phase {
            case .loading:
                HomeSkeleton()
            case let .loaded(home, _):
                content(presenter.model(home))
            case let .failed(cached):
                failed(cached.map(presenter.model))
            }
        }
    }

    @ViewBuilder
    private func content(_ model: HomeAnswerModel) -> some View {
        HomeHero(model: model, send: send)
        VStack(alignment: .leading, spacing: 0) {
            switch model.body {
            case let .upgrades(rows, footer):
                UpgradesCard(rows: rows, footer: footer, bandHead: model.bandHead, build: build, send: send)
                    .padding(.top, 16)
            case .allMatch, .message:
                EmptyView()
            }
            if let already = model.already {
                Text(already).font(ScaledType.meta).foregroundStyle(FSColor.muted).padding(.top, 12)
            }
            NextRows(title: model.nextTitle, rows: model.next, send: send).padding(.top, 22)
        }
        .padding(.horizontal, 16)
        .padding(.bottom, 24)
    }

    @ViewBuilder
    private func failed(_ cached: HomeAnswerModel?) -> some View {
        VStack(alignment: .leading, spacing: 14) {
            Text(presenter.heroMeta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            MicroLabel(Copy.homeAnswerLabel, color: FSColor.nav)
            Text(Copy.homeAnswerFailed).font(ScaledType.lead).foregroundStyle(FSColor.wipe)
                .accessibilityAddTraits(.isStaticText)
            SecondaryButton(title: Copy.commonTryAgain, systemImage: "arrow.clockwise") { send(.retry) }
        }
        .padding(EdgeInsets(top: 14, leading: 16, bottom: 18, trailing: 16))
        .frame(maxWidth: .infinity, alignment: .leading)
        .heroBand()
        NextRows(title: cached?.nextTitle ?? Copy.homeNextTitle(name: presenter.character.name),
                 rows: cached?.next ?? presenter.nextRowsWithoutAnswer(), send: send)
            .padding(.horizontal, 16)
            .padding(.top, 22)
            .padding(.bottom, 24)
    }
}

/// The hero band: who, the label, the two figures, the note and the sync line.
struct HomeHero: View {
    let model: HomeAnswerModel
    let send: (HomeIntent) -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            Text(model.heroMeta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            MicroLabel(Copy.homeAnswerLabel, color: FSColor.nav).padding(.top, 14)
            answer.padding(.top, 8)
            if let note = model.note {
                Text(note).font(ScaledType.meta).foregroundStyle(FSColor.muted).padding(.top, 6)
            }
            if let synced = model.synced {
                Label { Text(synced) } icon: { Image(systemName: "arrow.triangle.2.circlepath").font(.system(size: 12)) }
                    .font(ScaledType.meta).foregroundStyle(FSColor.nav).padding(.top, 10)
            }
        }
        .padding(EdgeInsets(top: 14, leading: 16, bottom: 18, trailing: 16))
        .frame(maxWidth: .infinity, alignment: .leading)
        .heroBand()
    }

    @ViewBuilder
    private var answer: some View {
        switch model.body {
        case .upgrades:
            figures.accessibilityElement(children: .combine)
        case let .allMatch(text):
            Text(text).font(ScaledType.lead).foregroundStyle(FSColor.strong)
        case let .message(text, buttons, note):
            VStack(alignment: .leading, spacing: 14) {
                Text(text).font(ScaledType.lead).foregroundStyle(FSColor.strong)
                if !buttons.isEmpty {
                    ViewThatFits(in: .horizontal) {
                        HStack(spacing: 10) { actionButtons(buttons) }
                        VStack(alignment: .leading, spacing: 10) { actionButtons(buttons) }
                    }
                }
                if let note { Text(note).font(ScaledType.meta).foregroundStyle(FSColor.muted) }
            }
        }
    }

    private func actionButtons(_ buttons: [HomeAction]) -> some View {
        ForEach(buttons, id: \.title) { button in
            if button.primary {
                PrimaryButton(title: button.title) { send(.action(button.kind)) }
            } else {
                SecondaryButton(title: button.title) { send(.action(button.kind)) }
            }
        }
    }

    private var figures: some View {
        ViewThatFits(in: .horizontal) {
            HStack(alignment: .firstTextBaseline, spacing: 22) { figureTexts }
            VStack(alignment: .leading, spacing: 4) { figureTexts }
        }
    }

    @ViewBuilder
    private var figureTexts: some View {
        if let count = model.count { figure(count) }
        if let gain = model.gain {
            HStack(alignment: .firstTextBaseline, spacing: 6) {
                Text(gain.figure).font(ScaledType.stat).foregroundStyle(FSColor.strong)
                Text(gain.unit).font(ScaledType.body).foregroundStyle(FSColor.muted)
            }
        }
    }

    /// `8 upgrades`: the number in `stat`, the word in body muted.
    private func figure(_ text: String) -> some View {
        let parts = text.split(separator: " ", maxSplits: 1).map(String.init)
        return HStack(alignment: .firstTextBaseline, spacing: 6) {
            Text(parts.first ?? text).font(ScaledType.stat).foregroundStyle(FSColor.strong)
            if parts.count > 1 { Text(parts[1]).font(ScaledType.body).foregroundStyle(FSColor.muted) }
        }
    }
}

/// The Your upgrades card: head row, five rows in the server's order, the footer to Your gear.
struct UpgradesCard: View {
    let rows: [HomeUpgradeRow]
    let footer: String
    let bandHead: String?
    let build: String
    let send: (HomeIntent) -> Void

    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            HStack {
                MicroLabel(Copy.homeUpgradesTitle)
                Spacer()
                if let bandHead { Text(bandHead).font(ScaledType.meta).foregroundStyle(FSColor.muted) }
            }
            .padding(.horizontal, 16)
            .padding(.vertical, 12)
            ForEach(rows) { row in
                Divider().overlay(FSColor.lineSoft)
                Button { send(.openItem(row.pick, sourceLine: row.sourceLine, wornID: row.wornID)) } label: {
                    UpgradeRowView(row: row, build: build, gainBelow: typeSize.isAccessibilitySize)
                }
                .buttonStyle(RowPress())
                .accessibilityHint(Copy.itemsA11yHint)
            }
            Divider().overlay(FSColor.lineSoft)
            Button { send(.openGear) } label: {
                HStack {
                    Text(footer).font(ScaledType.bodyStrong).foregroundStyle(FSColor.gold)
                    Spacer()
                    Image(systemName: "chevron.right").foregroundStyle(FSColor.gold).accessibilityHidden(true)
                }
                .padding(.horizontal, 16)
                .frame(minHeight: 52)
                .contentShape(Rectangle())
            }
            .buttonStyle(RowPress())
        }
        .primaryCard()
    }
}

/// One upgrade row: the pick's icon, name, slot and source, what it replaces, and the gain.
struct UpgradeRowView: View {
    let row: HomeUpgradeRow
    let build: String
    /// At the largest sizes the gain moves under the name (spec §10, XL.1).
    let gainBelow: Bool

    var body: some View {
        HStack(alignment: .top, spacing: 12) {
            ItemIconView(icon: row.pick.icon, quality: row.pick.quality, build: build, size: 44)
            VStack(alignment: .leading, spacing: 2) {
                Text(row.pick.name ?? Copy.itemsNotInData).font(ScaledType.body)
                    .foregroundStyle(ItemQuality(value: row.pick.quality).text)
                Text(row.meta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                Text(row.replaces).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                switch row.gain {
                case let .words(words):
                    Text(words).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                case let .measured(figure, unit) where gainBelow:
                    Text("\(figure) \(unit)").font(ScaledType.mono).foregroundStyle(FSColor.strong)
                case .measured:
                    EmptyView()
                }
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            if case let .measured(figure, unit) = row.gain, !gainBelow {
                VStack(alignment: .trailing, spacing: 0) {
                    Text(figure).font(ScaledType.mono).foregroundStyle(FSColor.strong)
                    Text(unit.uppercased()).font(ScaledType.micro).foregroundStyle(FSColor.muted)
                }
            }
        }
        .multilineTextAlignment(.leading)
        .padding(.horizontal, 16)
        .padding(.vertical, 12)
        .frame(minHeight: 72, alignment: .top)
        .contentShape(Rectangle())
        .accessibilityElement(children: .combine)
    }
}

/// The §8 list-row pressed wash.
struct RowPress: ButtonStyle {
    func makeBody(configuration: Configuration) -> some View {
        configuration.label.background(configuration.isPressed ? FSColor.gold.opacity(0.10) : .clear)
    }
}

/// "Next for {name}": Talents and Sims (spec F2.1 item 4).
struct NextRows: View {
    let title: String
    let rows: [HomeNextRow]
    let send: (HomeIntent) -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            MicroLabel(title)
            VStack(spacing: 0) {
                ForEach(Array(rows.enumerated()), id: \.element.id) { index, row in
                    if index > 0 { Divider().overlay(FSColor.lineSoft) }
                    NextRowView(row: row)
                }
            }
            .quietCard()
        }
    }
}

struct NextRowView: View {
    let row: HomeNextRow

    var body: some View {
        HStack(alignment: .center, spacing: 12) {
            Image(systemName: row.kind == .talents ? "sparkles" : "flask")
                .font(.system(size: 16))
                .foregroundStyle(FSColor.goldHover)
                .frame(width: 36, height: 36)
                .background(FSColor.gold.opacity(0.12), in: .rect(cornerRadius: FSRadius.pill))
                .accessibilityHidden(true)
            VStack(alignment: .leading, spacing: 2) {
                Text(row.title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                Text(row.line).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            switch row.tag {
            case .matches: TagView(text: Copy.homeNextMatches, color: FSColor.kill)
            case .differs: TagView(text: Copy.homeNextDiffers, color: FSColor.gold)
            case nil: EmptyView()
            }
        }
        .padding(.horizontal, 16)
        .padding(.vertical, 10)
        .frame(minHeight: 56)
        .accessibilityElement(children: .combine)
    }
}

/// F2.2: the band and the card as skeletons at their real heights, five rows, and the line.
struct HomeSkeleton: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            VStack(alignment: .leading, spacing: 10) {
                bar(width: 220, height: 13)
                bar(width: 160, height: 11).padding(.top, 4)
                bar(width: 200, height: 30)
                bar(width: 260, height: 13)
                bar(width: 180, height: 13)
            }
            .padding(EdgeInsets(top: 14, leading: 16, bottom: 18, trailing: 16))
            .frame(maxWidth: .infinity, alignment: .leading)
            .heroBand()
            VStack(spacing: 0) {
                ForEach(0..<HomePresenter.visibleRows, id: \.self) { _ in
                    HStack(alignment: .top, spacing: 12) {
                        RoundedRectangle(cornerRadius: FSRadius.pill).fill(FSColor.cardTop).frame(width: 44, height: 44)
                        VStack(alignment: .leading, spacing: 6) {
                            bar(width: 180, height: 14)
                            bar(width: 140, height: 11)
                            bar(width: 200, height: 11)
                        }
                        Spacer()
                    }
                    .padding(.horizontal, 16)
                    .frame(height: 72)
                }
            }
            .padding(.vertical, 8)
            .quietCard()
            .padding(.horizontal, 16)
            .padding(.top, 16)
            Text(Copy.homeAnswerLoading).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .padding(16)
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.homeAnswerLoading)
    }

    private func bar(width: CGFloat, height: CGFloat) -> some View {
        RoundedRectangle(cornerRadius: 4).fill(FSColor.cardTop).frame(width: width, height: height)
    }
}
