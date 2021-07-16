package main

type Score int
type HighScore Score

func main() {
	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	// hs = s // cannot use s (variable of type Score) as HighScore value in assignment
	// s = i // cannot use i (variable of type int) as Score value in assignment
	s = Score(i)
	hs = HighScore(s)
}
