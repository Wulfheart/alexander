package orders


type FullOrders struct{
	Moves []Move 
	SupportMoves []SupportMove 
	SupportHolds []SupportHold 
	Holds Hold 
	Convoy []Convoy 
	Builds []Build 
	Disbands Disband 
}
