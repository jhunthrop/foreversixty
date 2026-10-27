package discordbot

// pressCopy is what a press on the raid post, the character select and /signup answer
// (experience spec §3.3.2, §6.2 "Press: …"); the tests read each string back out of the spec.
// {Name}-style markers are filled by fill.
var pressCopy = struct {
	NotLinked string
	// Refusals before the sign-up.
	Locked, Cancelled, Ended string
	NotMember, SetupButton   string
	WaitingOne, WaitingMany  string
	RateLimited, Failure     string
	// The character select.
	SelectPrompt, SelectPlaceholder, OptionDescription, OptionMain, Showing string
	// Outcomes. InSaved plus Count is the spec's In sentence; CountOpen is the count of a raid
	// whose size is unknown. {Date} is the raid's day, "Thu 10 Dec", in the guild's raid zone,
	// so a guild with two raids of one title a week can tell which the press landed on.
	InSaved, Count, CountOpen, LateSaved, NoteButton, BenchSaved, OutSaved, Moved string
	// Already is for In and Late; AlreadyBench and AlreadyOut have no role.
	Already, AlreadyBench, AlreadyOut string
	// NoRaid answers /signup with nothing planned.
	NoRaid string
}{
	NotLinked:         "Link your Discord account to Forever Sixty to sign up from here, then press the button again.",
	Locked:            "Sign-ups are locked. An officer can still change yours.",
	Cancelled:         "This raid was cancelled.",
	Ended:             "This raid has ended.",
	NotMember:         "Sign-ups are for members of {Guild}. Save your character on Forever Sixty and an officer can approve it.",
	SetupButton:       "Get set up",
	WaitingOne:        "{Name} is waiting for approval. You can sign it up once an officer approves it.",
	WaitingMany:       "Your characters are waiting for approval. You can sign them up once an officer approves them.",
	RateLimited:       "Too many sign-up changes. Wait a minute, then try again.",
	Failure:           "That did not save. Try again.",
	SelectPrompt:      "Sign up as {State} for {Title}, <t:{unix}:f>, with which character?",
	SelectPlaceholder: "Choose a character",
	OptionDescription: "{Spec} {Class} · {Role}",
	OptionMain:        " · your main",
	Showing:           "Showing 25 of {n}. Sign up the others on the raid page.",
	InSaved:           "{Name} is In as {Role} for {Title}, {Date}.",
	Count:             " {signed} of {size} signed up.",
	CountOpen:         " {n} signed up.",
	LateSaved:         "{Name} is Late as {Role} for {Title}, {Date}. Add a note for the officers on the raid page.",
	NoteButton:        "Add a note",
	BenchSaved:        "{Name} is on the Bench for {Title}, {Date}.",
	OutSaved:          "{Name} is Out for {Title}, {Date}.",
	Moved:             " Your other character {Other} moved to Bench.",
	Already:           "{Name} is already {State} as {Role} for {Title}, {Date}.",
	AlreadyBench:      "{Name} is already on the Bench for {Title}, {Date}.",
	AlreadyOut:        "{Name} is already Out for {Title}, {Date}.",
	NoRaid:            "No raid planned yet.",
}

// roleWords are the sign-up roles as the sentences say them.
var roleWords = map[string]string{"tank": "Tank", "healer": "Healer", "dps": "DPS"}
