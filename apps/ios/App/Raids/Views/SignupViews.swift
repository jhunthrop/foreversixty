import DesignSystem
import Models
import SwiftUI

/// The four segments (spec F4.3): In, Late, Bench, Out; 2 x 2 from AX1.
struct SegmentsView: View {
    let characterName: String
    let block: SignupBlock
    let enabled: Bool
    let press: (SignupState) -> Void
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        Group {
            if typeSize >= .accessibility1 {
                Grid(horizontalSpacing: 0, verticalSpacing: 0) {
                    GridRow { segment(.signedIn); segment(.late) }
                    GridRow { segment(.bench); segment(.out) }
                }
            } else {
                HStack(spacing: 0) { ForEach(SignupState.allCases) { segment($0) } }
            }
        }
        .background(FSColor.raised)
        .clipShape(RoundedRectangle(cornerRadius: FSRadius.control))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.line))
        .opacity(enabled ? 1 : 0.5)
        .accessibilityElement(children: .contain)
        .accessibilityLabel(Copy.raidsA11ySegments(name: characterName))
    }

    private func segment(_ state: SignupState) -> some View {
        let checked = block.shown == state
        let pending = block.pending == state
        return Button { press(state) } label: {
            Text(state.label)
                .font(ScaledType.segment)
                .lineLimit(1)
                .minimumScaleFactor(0.6)
                .foregroundStyle(checked ? FSColor.strong : FSColor.muted)
                .frame(maxWidth: .infinity, minHeight: 48)
                .background(checked ? (state.isComing ? FSColor.kill.opacity(0.15) : FSColor.lineWarm) : .clear)
                .overlay(alignment: .bottom) {
                    if checked { Rectangle().fill(state.isComing ? FSColor.kill : FSColor.muted).frame(height: 2) }
                }
                .overlay {
                    if pending {
                        RoundedRectangle(cornerRadius: FSRadius.control - 2)
                            .strokeBorder(FSColor.gold, style: StrokeStyle(lineWidth: 1.5, dash: [4, 3]))
                            .padding(4)
                    }
                }
                .overlay(alignment: .trailing) {
                    if state != .out { Rectangle().fill(FSColor.line).frame(width: 1) }
                }
        }
        .buttonStyle(SegmentPressStyle(state: state))
        .disabled(!enabled)
        .accessibilityAddTraits(checked ? .isSelected : [])
        .accessibilityValue(checked && block.busy ? Copy.raidsA11ySaving : pending ? Copy.raidsA11yPending : "")
    }
}

private struct SegmentPressStyle: ButtonStyle {
    let state: SignupState

    func makeBody(configuration: Configuration) -> some View {
        configuration.label.background(
            configuration.isPressed ? (state.isComing ? FSColor.kill.opacity(0.09) : FSColor.lineWarm.opacity(0.6)) : .clear
        )
    }
}

/// The 20 pt status line, announced politely when it changes. On the raid page it is always
/// reserved so nothing shifts; on the next raid card it draws only with text (UX should 1).
struct StatusLineView: View {
    let status: StatusLine?
    var reserves = true

    var body: some View {
        if reserves || status != nil { line }
    }

    private var line: some View {
        Text(status?.text ?? " ")
            .font(ScaledType.meta)
            .foregroundStyle(color)
            .frame(maxWidth: .infinity, minHeight: 20, alignment: .leading)
            .fixedSize(horizontal: false, vertical: true)
            .onChange(of: status) { _, new in
                guard let text = new?.text else { return }
                AccessibilityNotification.Announcement(text).post()
            }
    }

    private var color: Color {
        switch status {
        case .saved: FSColor.kill
        case .error: FSColor.wipe
        case .info, nil: FSColor.muted
        }
    }
}

/// The who line: crest 36, the name in class colour, the spec icon 18 with `{Spec} · {Role}`.
struct WhoLine: View {
    let name: String
    let className: String
    let specName: String?
    let specIcon: String?
    let role: String
    var trailing: AnyView?
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        if typeSize.isAccessibilitySize {
            // Large text: the crest above the name, Change spec below at full width (spec 10).
            VStack(alignment: .leading, spacing: 6) {
                ClassCrest(className: className, size: 36)
                labels
                trailing
            }
        } else {
            HStack(spacing: 12) {
                ClassCrest(className: className, size: 36)
                labels
                Spacer(minLength: 8)
                trailing
            }
        }
    }

    private var labels: some View {
        VStack(alignment: .leading, spacing: 2) {
            NameText(text: name, font: ScaledType.name, color: ClassColor.of(className))
            HStack(spacing: 6) {
                GameIcon(stem: specIcon, size: 18)
                Text(specName.map { Copy.raidsWho(spec: $0, role: RaidRole.label(role)) } ?? RaidRole.label(role))
                    .font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
        }
        .accessibilityElement(children: .combine)
    }
}

/// One character's block on the sign-up card (spec F4.2 item 5).
struct SignupBlockView: View {
    let controller: SignupController
    let character: RaidViewerCharacter
    var showsNote = true
    /// The raid page reserves the status line; the next raid card draws it only with text.
    var reservesStatus = true
    @FocusState private var noteFocused: Bool

    var body: some View {
        let key = character.characterKey
        let block = controller.block(key)
        VStack(alignment: .leading, spacing: 8) {
            WhoLine(
                name: character.name, className: character.className,
                specName: character.specName, specIcon: character.specIcon, role: character.role,
                trailing: character.raidSpec != nil && controller.segmentsEnabled(key)
                    ? AnyView(Button(Copy.specChange) { controller.changeSpec(key) }.buttonStyle(GoldLinkStyle()).font(ScaledType.meta))
                    : nil
            )
            if character.verified {
                SegmentsView(characterName: character.name, block: block, enabled: controller.segmentsEnabled(key)) { state in
                    Task { await controller.press(key, state) }
                }
                StatusLineView(status: controller.restingStatus(key), reserves: reservesStatus)
                if showsNote, block.shown != nil {
                    noteField(key, block: block)
                }
            } else {
                Text(Copy.raidWaitingCharacter).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
        }
    }

    private func noteField(_ key: String, block: SignupBlock) -> some View {
        let text = controller.notes[key] ?? ""
        return VStack(alignment: .leading, spacing: 4) {
            TextField(
                Copy.raidNotePlaceholder,
                text: Binding(get: { controller.notes[key] ?? "" }, set: { controller.editNote(key, $0) })
            )
            .font(ScaledType.body)
            .foregroundStyle(FSColor.text)
            .focused($noteFocused)
            .submitLabel(.done)
            .onSubmit { Task { await controller.saveNote(key) } }
            .onChange(of: noteFocused) { _, focused in
                if !focused { Task { await controller.saveNote(key) } }
            }
            .padding(.horizontal, 14)
            .frame(minHeight: 48)
            .background(Color.black.opacity(0.35), in: .rect(cornerRadius: FSRadius.control))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(
                block.noteFailed ? FSColor.wipe : noteFocused ? FSColor.gold : FSColor.lineCard
            ))
            .disabled(!controller.segmentsEnabled(key))
            if block.noteBusy {
                Text(Copy.raidsNoteSaving).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            } else if block.noteFailed {
                Text(Copy.raidNoteFailed).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
            } else if text.count >= SignupController.noteCounterFrom {
                Text(Copy.raidNoteLeft(n: SignupController.noteLimit - text.count))
                    .font(ScaledType.monoMeta).foregroundStyle(FSColor.muted)
            }
        }
    }
}

/// The reminder ask (spec F4.4), inset in the sign-up card after the first In or Late.
struct ReminderAskView: View {
    let answer: (Bool) -> Void

    var body: some View {
        HStack(alignment: .top, spacing: 12) {
            Image(systemName: "bell")
                .foregroundStyle(FSColor.gold)
                .frame(width: 32, height: 32)
                .background(FSColor.gold.opacity(0.12), in: .rect(cornerRadius: FSRadius.control))
                .accessibilityHidden(true)
            VStack(alignment: .leading, spacing: 8) {
                Text(Copy.askRaidTitle).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                Text(Copy.askRaidBody).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                HStack(spacing: 12) {
                    Button(Copy.askTurnOnRaid) { answer(true) }.buttonStyle(BrandButtonStyle(.accent, height: 48, fullWidth: false))
                    Button(Copy.askNotNow) { answer(false) }
                        .font(ScaledType.button).foregroundStyle(FSColor.nav)
                        .frame(minHeight: 44)
                }
            }
        }
        .raidCardBox(14).quietCard()
        .transition(.opacity.combined(with: .move(edge: .top)))
    }
}

/// "Which spec tonight?" (spec F4.3, the guided-cards spec P.10).
struct SpecSheet: View {
    let ask: SpecAsk
    let character: RaidViewerCharacter
    let raidTitle: String
    let date: String
    let pick: (RaidSpecOption) -> Void

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 14) {
                Text(Copy.specTitle.uppercased()).font(ScaledType.sec).foregroundStyle(FSColor.strong)
                    .accessibilityAddTraits(.isHeader)
                Text(lead).font(ScaledType.body).foregroundStyle(FSColor.text)
                VStack(spacing: 0) {
                    ForEach(character.specOptions ?? [], id: \.spec) { option in
                        Button { pick(option) } label: { row(option) }
                            .buttonStyle(RowPressStyle())
                        if option.spec != character.specOptions?.last?.spec { Divider().overlay(FSColor.lineSoft) }
                    }
                }
                .raidCardBox(0).quietCard()
                Text(Copy.specCloseNote).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
            .padding(16)
        }
        .background(GameColors.sheet)
        .presentationDetents([.medium, .large])
        .presentationDragIndicator(.visible)
    }

    private var lead: AttributedString {
        let sentence: String = switch ask {
        case let .ask(_, state): Copy.specLead(name: character.name, state: state.label, title: raidTitle, date: date)
        case .change: Copy.specLeadChange(name: character.name, title: raidTitle, date: date)
        }
        var text = AttributedString(sentence)
        if let range = text.range(of: character.name) {
            text[range].font = ScaledType.bodyStrong
            text[range].foregroundColor = FSColor.strong
        }
        return text
    }

    private func row(_ option: RaidSpecOption) -> some View {
        let checked = ask == .change(characterKey: character.characterKey) && character.raidSpec == option.spec
        return HStack(spacing: 12) {
            GameIcon(stem: option.icon, size: 40)
            VStack(alignment: .leading, spacing: 2) {
                Text(option.name).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                Text(RaidRole.label(option.role)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
            Spacer()
            Image(systemName: checked ? "checkmark" : "chevron.right")
                .foregroundStyle(checked ? FSColor.kill : FSColor.muted)
                .accessibilityHidden(true)
        }
        .padding(.horizontal, 14)
        .frame(minHeight: 64)
        .contentShape(Rectangle())
        .accessibilityElement(children: .combine)
        .accessibilityAddTraits(checked ? .isSelected : [])
    }
}
