package game

type Builder interface {
	//Builds a building in the given Rental Group.
	Build(Group)
}
