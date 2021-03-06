package libs

// Options global options
type Options struct {
	RootFolder       string
	SignFolder       string
	PassiveFolder    string
	ResourcesFolder  string
	ThirdPartyFolder string
	ScanID           string
	ConfigFile       string
	FoundCmd         string
	PassiveOutput    string
	PassiveSummary   string
	Output           string
	SummaryOutput    string
	LogFile          string
	Proxy            string
	Selectors        string
	Params           []string
	Headers          []string
	Signs            []string
	Excludes         []string
	SelectedSigns    []string
	ParallelSigns    []string
	SelectedPassive  string
	GlobalVar        map[string]string

	Level           int
	Concurrency     int
	Threads         int
	Delay           int
	SaveRaw         bool
	Timeout         int
	Refresh         int
	Retry           int
	Verbose         bool
	Debug           bool
	NoBackGround    bool
	NoOutput        bool
	EnablePassive   bool
	DisableParallel bool
	Server          Server
}

// Server options for api server
type Server struct {
	DBPath       string
	Bind         string
	JWTSecret    string
	Cors         string
	DefaultSign  string
	SecretCollab string
	Username     string
	Password     string
	Key          string
}

// Job define job for running routine
type Job struct {
	URL  string
	Sign Signature
}

// PJob define job for running routine
type PJob struct {
	Req  Request
	ORec Record
	Sign Signature
}
