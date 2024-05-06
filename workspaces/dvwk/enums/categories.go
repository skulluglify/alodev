package models

type CategoryKind int

const (
	Applications CategoryKind = iota
	Articles
	Extensions
	Documents
	Files
	Images
	Links
	Others
	Plugins
	Posts
	Products
	Reels
	Videos
)

func (categoryKind CategoryKind) ToString() string {
	switch categoryKind {
	case Applications:
		return "Applications"
	case Articles:
		return "Articles"
	case Extensions:
		return "Extensions"
	case Documents:
		return "Documents"
	case Files:
		return "Files"
	case Images:
		return "Images"
	case Links:
		return "Links"
	case Others:
		return "Others"
	case Plugins:
		return "Plugins"
	case Reels:
		return "Reels"
	case Posts:
		return "Posts"
	case Products:
		return "Products"
	case Videos:
		return "Videos"
	default:
		return "Others"
	}
}

func CategoryKindFromString(categoryKind string) CategoryKind {
	switch categoryKind {
	case "Applications":
		return Applications
	case "Articles":
		return Articles
	case "Extensions":
		return Extensions
	case "Documents":
		return Documents
	case "Files":
		return Files
	case "Images":
		return Images
	case "Links":
		return Links
	case "Others":
		return Others
	case "Plugins":
		return Plugins
	case "Posts":
		return Posts
	case "Products":
		return Products
	case "Reels":
		return Reels
	case "Videos":
		return Videos
	default:
		return Others
	}
}
