package main


func sharesTransaction(transaction Transaction)int{
	if len(transaction.Shares) != len(transaction.IDs){
		return -1
	}
	valid := checkValidity(transaction)
	if valid == false{
		return 404
	}
	var totalShares int
	for _, share := range transaction.Shares{
		totalShares += share
	}
	for index,id := range transaction.IDs{
		if id == transaction.PaidBy{
			continue
		}
		var splitValue float32 = float32(transaction.Shares[index])/float32(totalShares)
		splitValue *= float32(transaction.Total)
		if bank[transaction.PaidBy] == nil{
			bank[transaction.PaidBy] = make(map[int]float32)
		}
		bank[transaction.PaidBy][id] += splitValue
		if bank[id]==nil{
			bank[id]=make(map[int]float32)
		}
		bank[id][transaction.PaidBy] -= splitValue
	}
	return 0
}