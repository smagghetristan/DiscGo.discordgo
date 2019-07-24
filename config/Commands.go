package config

import (
	"github.com/bwmarrin/discordgo"
)

func addMenu(Menu Menu) {
	Menus = append(Menus, Menu)
}

func addCategory(Category Category) {
	Categories = append(Categories, Category)
}

func AddCommand(Command Command) {
	//Check if the menu exist

	i := 0
	exist := false
	for i = 0; i < len(Menus); i++ {
		if Menus[i].Name == Command.Category.Menu.Name {
			exist = true
			break
		}
	}
	if !exist {
		addMenu(Command.Category.Menu)
	}

	//Check if the Category exist

	i = 0
	exist = false
	for i = 0; i < len(Categories); i++ {
		if Categories[i].Name == Command.Category.Name && Command.Category.Menu.Name == Categories[i].Menu.Name {
			exist = true
			break
		}
	}
	if !exist {
		addCategory(Command.Category)
	}

	Commands = append(Commands, Command)
}

func CreateEmbeds() {
	AllEmbeds := map[string]*discordgo.MessageEmbed{}
	Fields := map[string][]*discordgo.MessageEmbedField{}

	Returned := []*discordgo.MessageEmbed{}

	for i := range Commands {

		if _, ok := AllEmbeds[Commands[i].Category.Menu.Name]; !ok {
			AllEmbeds[Commands[i].Category.Menu.Name] = &discordgo.MessageEmbed{
				Title:       Commands[i].Category.Menu.Name,
				Description: Commands[i].Category.Menu.Description,
				Color:       0xFFDD00,
			}
		}

		k := 0
		exist := false
		for k = 0; k < len(Fields[Commands[i].Category.Menu.Name]); k++ {
			if Fields[Commands[i].Category.Menu.Name][k].Name == Commands[i].Category.Name {
				exist = true
				break
			}
		}
		if !exist {
			Fields[Commands[i].Category.Menu.Name] = append(Fields[Commands[i].Category.Menu.Name], &discordgo.MessageEmbedField{
				Name:  Commands[i].Category.Name,
				Value: Commands[i].ShortDescription,
			})
			AllEmbeds[Commands[i].Category.Menu.Name].Fields = Fields[Commands[i].Category.Menu.Name]
		} else {
			Fields[Commands[i].Category.Menu.Name][k].Value += "\n" + Commands[i].ShortDescription
			AllEmbeds[Commands[i].Category.Menu.Name].Fields = Fields[Commands[i].Category.Menu.Name]
		}
	}

	for key := range AllEmbeds {
		Returned = append(Returned, AllEmbeds[key])
	}

	HelpEmbeds = Returned
}
