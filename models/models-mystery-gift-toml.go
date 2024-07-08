package models

type MysteryGift struct {
	Name          string      `toml:"name"`
	GiftType      string      `toml:"giftType"`
	Pokemon       PokemonGift `toml:"pokemonGift"`
	BeginningDate string      `toml:"beginningDate"`
	EndDate       string      `toml:"endDate"`
}
type PokemonGift struct {
	OriginalTrainer  string      `toml:"originalTrainer"`
	ID               string      `toml:"id"`
	Species          string      `toml:"species"`
	Types            []string    `toml:"types"`
	Abilities        []Abilities `toml:"abilities"`
	Moves            []Moves     `toml:"moves"`
	Stats            Stats       `toml:"stats"`
	AssetData        AssetData   `toml:"assetData"`
	Nickname         string      `toml:"nickname"`
	LocationOfOrigin string      `toml:"locationOfOrigin"`
	Shiny            bool        `toml:"shiny"`
}
