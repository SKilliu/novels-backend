package utils

func GetVotingResults(novelOne, novelTwo float32) (votingResNovelOne, votingResNovelTwo int32) {
	if novelOne+novelTwo <= 5 {
		votingResNovelOne = 51
		votingResNovelTwo = 51
	} else {
		votingResNovelOne = int32((novelOne / (novelOne + novelTwo)) * 100.0)
		votingResNovelTwo = int32((novelTwo / (novelOne + novelTwo)) * 100.0)
	}

	return
}
