package lyrics

import "sync"

type Lyric struct {
	Original string
	Pronunciation string
	Translation string
}

type music struct {
	lyrics []*Lyric
}

var m *music
var once sync.Once

func inputOriginalLyrics(data string) *Lyric {
	line := Lyric{data, "", ""}
	return &line
}

func Lyrics()*music {
	if m == nil {
		once.Do(func() {
			m = &music{}
			m.AddLyric("")
		})
	}
	return m
}

func (m *music) AddLyric(data string) {
	m.lyrics = append(m.lyrics, inputOriginalLyrics(data))
}


func (m *music) ShowAllLyrics() []*Lyric {
	return m.lyrics
}

func (m *music) InputPronunciation(data string) {

}