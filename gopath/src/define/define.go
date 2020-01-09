package define

type Request struct {
	Word string
}

type Response struct {
	LetterMap map[string]int
}
