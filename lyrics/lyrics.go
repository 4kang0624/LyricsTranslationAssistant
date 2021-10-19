package lyrics

type Lyric struct {
	Original string
	Pronunciation string
	Translation string
}

type music struct {
	lyrics []*Lyric
}

var m *music

func inputOriginalLyrics(data string) *Lyric {
	line := Lyric{data, "", ""}
	return &line
}

func Lyrics()*music {
	return m
}

func (m *music) AddLyric(data string) {
	m.lyrics = append(m.lyrics, inputOriginalLyrics(data))
}


func (m *music) ShowAllLyrics() []*Lyric {
	return m.lyrics
}