package graph

type Ecosystem string

const (
	EcosystemNPM   Ecosystem = "npm"
	EcosystemGo    Ecosystem = "golang"
	EcosystemCargo Ecosystem = "crates.io"
)

type DependencyNode struct {
	ID        string    `json:"id"` // e.g. "npm:axios:0.21.1"
	Ecosystem Ecosystem `json:"ecosystem"`
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	Direct    bool      `json:"direct"` // Direct project dependency or transitive
}

type Graph struct {
	RootNodes []string // Project-level direct dependencies
	Nodes     map[string]*DependencyNode
	Adjacency map[string][]string // NodeID -> Slice of Child NodeIDs
}
