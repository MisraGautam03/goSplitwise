package main

func checkValidity(transaction Transaction)bool{
	for _, id := range transaction.IDs{
		_,valid := mapID[id]
		if valid == false{
			return false
		}
	}
	_,valid := mapID[transaction.PaidBy]
	if valid == false{
		return false
	}
	return true
}
