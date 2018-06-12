package git

type BranchSet struct {
	TreeId			int		`json:"treeId"`
	MainBranchId	int		`json:"mainBranchId"`
	BranchIds		[]int	`json:"branchIds"`
}