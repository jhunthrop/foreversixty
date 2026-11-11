import DesignSystem
import SwiftUI

/// What the applicant hands back.
struct ApplicantActions {
    var accept: () -> Void = {}
    var decline: () -> Void = {}
    var copy: (String) -> Void = { _ in }
}

/// The applicant (flows spec F3, mocks ios-raid-applicant, -accepted): the numbers, raid ready,
/// guild history and their note, Accept into guild / Decline anchored at the bottom; after Accept,
/// the `/ginvite` line with Copy.
struct ApplicantView: View {
    let content: ApplicantContent
    var busy = false
    var error: String?
    var actions = ApplicantActions()

    var body: some View {
        RaidFlowScroll {
            RaidFlowHead(title: content.name, titleColor: ClassColor.of(content.className), facts: content.facts) {
                ClassCrest(className: content.className, size: FSSize.crest56, flat: true)
            }
            if let accepted = content.accepted {
                acceptedBody(accepted)
            } else {
                sheet
            }
        }
        .safeAreaInset(edge: .bottom) {
            if content.accepted == nil { decision.parityPinned() }
        }
    }

    @ViewBuilder
    private var sheet: some View {
        ApplicantFigures(figures: content.figures)
        if !content.ready.isEmpty {
            FSSection(Copy.raidApplicantReady) { lines(content.ready, slim: true) }.padding(.top, RaidFlowSpacing.section)
        }
        if !content.history.isEmpty {
            FSSection(Copy.raidApplicantHistory) { lines(content.history, slim: false) }.padding(.top, RaidFlowSpacing.section)
        }
        if let note = content.note {
            FSSection(Copy.raidApplicantNote) {
                Text(note).type(FSType.prose).foregroundStyle(FSColor.text)
                    .fixedSize(horizontal: false, vertical: true)
                    .padding(.top, RaidFlowSpacing.line)
            }
            .padding(.top, RaidFlowSpacing.section)
        }
    }

    private func lines(_ lines: [HomeLine], slim: Bool) -> some View {
        VStack(spacing: 0) {
            ForEach(lines, id: \.title) { line in
                FSRow(title: line.title, sub: line.sub.isEmpty ? nil : line.sub, density: slim ? .slim : .regular, inset: 0) { EmptyView() }
            }
        }
    }

    /// Accept into guild / Decline over the role line, above the tab bar.
    private var decision: some View {
        VStack(alignment: .leading, spacing: FSSpace.space2) {
            HStack(spacing: FSSpace.space3) {
                FSButton(Copy.raidApplicantAccept, kind: .primary, action: actions.accept)
                FSButton(Copy.raidApplicantDecline, kind: .quiet, action: actions.decline)
            }
            .disabled(busy)
            SmallText(error ?? Copy.raidApplicantRole, color: error == nil ? nil : FSColor.wipe)
        }
        .padding(.top, FSSpace.space3)
        .padding(.bottom, FSSpace.space1)
        .padding(.horizontal, FSSpace.gutterPhone)
        .screenGround()
        .overlay(alignment: .top) { Hairline() }
    }

    @ViewBuilder
    private func acceptedBody(_ accepted: String) -> some View {
        Text(accepted).type(FSType.lead).foregroundStyle(FSColor.strong)
            .fixedSize(horizontal: false, vertical: true)
            .padding(.top, FSSpace.space3)
        HStack(spacing: FSSpace.space2) {
            Text(content.invite)
                .font(.system(size: 14, weight: .medium, design: .monospaced))
                .foregroundStyle(FSColor.strong)
                .frame(maxWidth: .infinity, alignment: .leading)
                .textSelection(.enabled)
            Button(Copy.raidApplicantCopy) { actions.copy(content.invite) }
                .buttonStyle(FSButtonStyle(fullWidth: false, small: true))
        }
        .padding(.vertical, 10)
        .padding(.horizontal, 12)
        .overlay(RoundedRectangle(cornerRadius: FSRadius.control).strokeBorder(FSColor.line, lineWidth: FSStroke.hairline))
        .padding(.top, FSSpace.space2)
        SmallText(content.after).padding(.top, FSSpace.space3)
    }
}

/// `.figs` on a phone: two columns, the label over the figure and its unit.
struct ApplicantFigures: View {
    let figures: [ApplicantFigure]

    static let columnGap: CGFloat = 24
    static let rowGap: CGFloat = 20
    /// The figure's 32 pt line (Open Sans's natural line at 28 is 38, M0's metrics).
    @ScaledMetric(relativeTo: .title2) private var figureLine: CGFloat = 32

    var body: some View {
        Grid(alignment: .leading, horizontalSpacing: Self.columnGap, verticalSpacing: Self.rowGap) {
            ForEach(Array(stride(from: 0, to: figures.count, by: 2)), id: \.self) { start in
                GridRow {
                    ForEach(figures[start ..< min(start + 2, figures.count)]) { figure in cell(figure) }
                }
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
    }

    private func cell(_ figure: ApplicantFigure) -> some View {
        VStack(alignment: .leading, spacing: 2) {
            // The figures sit on the painted ground: the label and unit in mutedGround (Android FlowFigures).
            Text(figure.label)
                .type(FSType.label)
                .foregroundStyle(FSColor.mutedGround)
            HStack(alignment: .firstTextBaseline, spacing: 4) {
                Text(figure.value).type(FSType.figure).foregroundStyle(FSColor.strong).frame(height: figureLine)
                if let unit = figure.unit {
                    Text(unit).font(.custom(FSFont.bodySemibold, size: 13, relativeTo: .footnote)).foregroundStyle(FSColor.mutedGround)
                }
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .accessibilityElement(children: .combine)
    }
}
