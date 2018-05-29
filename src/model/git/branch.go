package git

import "errors"

type Branch struct {
	Id			int
	Name		string
	HeadId		int
	CommitIds	[]int
	StartTime	uint32
	EndTime		uint32
}

func (branch *Branch)MergeWith(other *Branch) (Commit, error) {
	// 获取 Head Commit
	dstHead, err := LoadCommit(branch.HeadId)
	if err != nil {
		return Commit{}, errors.New("Wrong Commit Id ")
	}
	srcHead, err := LoadCommit(other.HeadId)
	if err != nil {
		return Commit{}, errors.New("Wrong Commit Id ")
	}

	// 尝试 Merge 两颗树


}