import DesignSystem
import Models
import SwiftUI
import UIKit

/// The officer row (spec F4.18), drawn only when `viewer.can` says so.
struct OfficerRow: View {
    let controller: SignupController
    let clock: RaidClock

    var body: some View {
        let facts = controller.facts
        let locked = facts.chip == .locked
        let ended = controller.raid.viewer.toolsEnded == true
        VStack(alignment: .leading, spacing: 10) {
            SectionLabel(text: Copy.officerLabel)
            ViewThatFits(in: .horizontal) {
                HStack(spacing: 10) { lockButton(locked: locked, ended: ended); copyButton(ended: ended) }
                VStack(alignment: .leading, spacing: 10) { lockButton(locked: locked, ended: ended); copyButton(ended: ended) }
            }
            Text(ended ? Copy.officerGate : locked ? Copy.officerLockedNote : Copy.officerLockNote)
                .font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.muted)
            if let error = controller.officerError {
                Text(error).font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.wipe)
            }
        }
    }

    private func lockButton(locked: Bool, ended: Bool) -> some View {
        Button { Task { await controller.toggleLock() } } label: {
            HStack(spacing: 8) {
                if controller.officerBusy {
                    ProgressView().controlSize(.small)
                } else {
                    Image(systemName: locked ? "lock.open" : "lock").accessibilityHidden(true)
                }
                Text(locked ? Copy.officerUnlock : Copy.officerLock)
            }
        }
        .buttonStyle(SecondaryButtonStyle())
        .disabled(ended || !controller.facts.canLock || controller.offline)
        .opacity(ended ? 0.5 : 1)
    }

    private func copyButton(ended: Bool) -> some View {
        Button {
            UIPasteboard.general.string = controller.copyList(clock: clock)
        } label: {
            HStack(spacing: 8) {
                Image(systemName: controller.copied ? "checkmark" : "doc.on.clipboard").accessibilityHidden(true)
                Text(controller.copied ? Copy.officerCopied : Copy.officerCopy)
            }
        }
        .buttonStyle(SecondaryButtonStyle())
        .disabled(ended)
        .opacity(ended ? 0.5 : 1)
    }
}

/// Sign-ups (spec F4.2 item 8): groups by role, then Late, Bench, Out, then Not signed up. An
/// officer's rows end in their place button (F4.19); Not signed up become rows with Add.
struct RosterView: View {
    let controller: SignupController
    let facts: RaidFacts

    var body: some View {
        let roster = RosterGroups(raid: controller.raid, overrides: controller.rosterOverrides)
        let officer = facts.canMove && facts.isSignable
        VStack(alignment: .leading, spacing: 14) {
            HStack(alignment: .firstTextBaseline, spacing: 8) {
                Text(Copy.raidSignups.uppercased()).font(Theme.Typography.h2).foregroundStyle(Theme.Palette.strong)
                    .accessibilityAddTraits(.isHeader)
                Text(facts.ratio).font(Theme.Typography.mono).foregroundStyle(Theme.Palette.muted)
            }
            ForEach(roster.groups) { group in
                VStack(alignment: .leading, spacing: 8) {
                    SectionLabel(text: group.kind.label, trailing: "\(group.rows.count)")
                    if group.rows.isEmpty, let empty = group.kind.emptyLine {
                        Text(empty).font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.muted)
                            .padding(14).quietCard()
                    } else {
                        rows(group.rows, officer: officer)
                    }
                }
            }
            if !roster.notSigned.isEmpty {
                if officer {
                    VStack(alignment: .leading, spacing: 8) {
                        SectionLabel(text: Copy.raidGroupNotSigned, trailing: "\(roster.notSigned.count)")
                        rows(roster.notSigned, officer: true)
                    }
                } else {
                    notSignedLine(roster.notSigned)
                }
            }
        }
        .animation(.default, value: roster)
    }

    private func rows(_ rows: [RosterRow], officer: Bool) -> some View {
        VStack(spacing: 0) {
            ForEach(rows) { row in
                RosterRowView(row: row, officer: officer, error: controller.rowErrors[row.characterKey]) { state in
                    Task { await controller.move(row, to: state) }
                }
                if row.id != rows.last?.id { Divider().overlay(Theme.Palette.lineSoft) }
            }
        }
        .quietCard()
    }

    private func notSignedLine(_ rows: [RosterRow]) -> some View {
        let names = rows.enumerated().reduce(Text("")) { line, item in
            let name = Text(item.element.name).foregroundStyle(Theme.Palette.classColor(item.element.className))
            return item.offset == 0 ? Text("\(name)") : Text("\(line), \(name)")
        }
        return VStack(alignment: .leading, spacing: 6) {
            SectionLabel(text: Copy.raidGroupNotSigned, trailing: "\(rows.count)")
            names.font(Theme.Typography.metaText)
        }
        .accessibilityElement(children: .combine)
    }
}

/// One roster row (min 48): spec icon 28, the name in class colour, the spec, the note.
struct RosterRowView: View {
    let row: RosterRow
    let officer: Bool
    let error: String?
    let move: (SignupState) -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 4) {
            HStack(spacing: 12) {
                GameIcon(stem: row.specIcon, size: 28)
                VStack(alignment: .leading, spacing: 2) {
                    Text(row.name).font(Theme.Typography.bodyStrong).foregroundStyle(Theme.Palette.classColor(row.className))
                    if !row.specName.isEmpty {
                        Text(row.specName).font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.muted)
                    }
                    if !row.note.isEmpty {
                        Text(row.note).font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.text)
                            .fixedSize(horizontal: false, vertical: true)
                    }
                }
                .accessibilityElement(children: .ignore)
                .accessibilityLabel([Copy.raidsA11yRoster(name: row.name, spec: row.specName, className: GameClass(rawValue: row.className)?.name ?? ""), row.note].filter { !$0.isEmpty }.joined(separator: ". "))
                Spacer(minLength: 8)
                if officer { PlaceMenu(row: row, move: move) }
            }
            if let error {
                Text(error).font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.wipe)
            }
        }
        .padding(.horizontal, 14)
        .padding(.vertical, 10)
        .frame(minHeight: 48)
        .background(row.own ? Theme.Palette.rowWash : .clear)
    }
}

/// The officer's place button and its move menu (spec F4.20).
struct PlaceMenu: View {
    let row: RosterRow
    let move: (SignupState) -> Void

    /// The menu's items for a place (the spec's table).
    static func items(for state: SignupState?) -> [(String, SignupState, Bool)] {
        switch state {
        case .signedIn: [(Copy.officerMoveBench, .bench, false), (Copy.officerMoveRemove, .out, true)]
        case .bench: [(Copy.officerMoveIn, .signedIn, false), (Copy.officerMoveRemove, .out, true)]
        case .late: [(Copy.officerMoveIn, .signedIn, false), (Copy.officerMoveBench, .bench, false)]
        case .out, nil: [(Copy.officerMoveAddIn, .signedIn, false), (Copy.officerMoveAddBench, .bench, false)]
        }
    }

    var body: some View {
        let place = row.state?.label ?? Copy.officerAdd
        Menu {
            Section(Copy.officerMenuHead(name: row.name, state: row.state?.label ?? Copy.raidChipNotSigned, role: RaidRole.label(row.role))) {
                ForEach(Self.items(for: row.state), id: \.0) { item in
                    Button(role: item.2 ? .destructive : nil) { move(item.1) } label: { Text(item.0) }
                }
            }
        } label: {
            HStack(spacing: 4) {
                Text(place).font(Theme.Typography.segment)
                Image(systemName: "chevron.down").font(.system(size: 11, weight: .semibold)).accessibilityHidden(true)
            }
            .foregroundStyle(Theme.Palette.strong)
            .padding(.horizontal, 12)
            .frame(minWidth: 56, minHeight: 44)
            .background(Theme.Palette.background.opacity(0.4), in: .rect(cornerRadius: Theme.Radius.control))
            .overlay(RoundedRectangle(cornerRadius: Theme.Radius.control).stroke(Theme.Palette.line))
        }
        .accessibilityLabel(Copy.raidsA11yPlace(name: row.name, state: place))
    }
}
