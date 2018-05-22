package git

type Branch struct {
	Id			int
	Name		string
	HeadId		int
	CommitIds	[]int
}
