package api

import domainMusic "github.com/zhorzh-p/LyricLibrary/internal/domain/music"

func domainToSongData(song domainMusic.Song) SongData {
	verses := make([]VerseData, len(song.Verses))

	for i, vers := range song.Verses {
		verses[i] = domainToVerseData(vers)
	}

	return SongData{
		ID:          song.ID,
		Name:        song.Name,
		ReleaseDate: song.ReleaseDate,
		Group:       domainToGroupData(song.Group),
		Verses:      verses,
	}
}

func domainToGroupData(group domainMusic.Group) GroupData {
	return GroupData{
		Name: group.Name,
	}
}

func domainToVerseData(verse domainMusic.Verse) VerseData {
	return VerseData{
		Text:  verse.Text,
		Order: verse.Order,
	}
}
